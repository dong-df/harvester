package upgrade

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"

	"github.com/harvester/harvester/pkg/config"
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
)

const (
	upgradeControllerName = "harvester-upgrade-controller"
	planControllerName    = "harvester-plan-controller"
	jobControllerName     = "harvester-upgrade-job-controller"
	podControllerName     = "harvester-upgrade-pod-controller"
	settingControllerName = "harvester-version-setting-controller"
	vmImageControllerName = "harvester-upgrade-vm-image-controller"
	secretControllerName  = "harvester-upgrade-secret-controller"
	nodeControllerName    = "harvester-upgrade-node-controller"
)

func Register(ctx context.Context, management *config.Management, options config.Options) error {
	if !options.HCIMode {
		return nil
	}

	upgrades := management.HarvesterFactory.Harvesterhci().V1beta1().Upgrade()
	upgradeLogs := management.HarvesterFactory.Harvesterhci().V1beta1().UpgradeLog()
	versions := management.HarvesterFactory.Harvesterhci().V1beta1().Version()
	settings := management.HarvesterFactory.Harvesterhci().V1beta1().Setting()
	plans := management.UpgradeFactory.Upgrade().V1().Plan()
	managedcharts := management.RancherManagementFactory.Management().V3().ManagedChart()
	nodes := management.CoreFactory.Core().V1().Node()
	jobs := management.BatchFactory.Batch().V1().Job()
	pods := management.CoreFactory.Core().V1().Pod()
	vmImages := management.HarvesterFactory.Harvesterhci().V1beta1().VirtualMachineImage()
	vms := management.VirtFactory.Kubevirt().V1().VirtualMachine()
	services := management.CoreFactory.Core().V1().Service()
	namespaces := management.CoreFactory.Core().V1().Namespace()
	clusters := management.ProvisioningFactory.Provisioning().V1().Cluster()
	machines := management.ClusterFactory.Cluster().V1beta1().Machine()
	secrets := management.CoreFactory.Core().V1().Secret()
	pvcs := management.CoreFactory.Core().V1().PersistentVolumeClaim()
	lhSettings := management.LonghornFactory.Longhorn().V1beta2().Setting()
	configMaps := management.CoreFactory.Core().V1().ConfigMap()

	virtSubsrcConfig := rest.CopyConfig(management.RestConfig)
	virtSubsrcConfig.GroupVersion = &schema.GroupVersion{Group: "subresources.kubevirt.io", Version: "v1"}
	virtSubsrcConfig.APIPath = "/apis"
	virtSubsrcConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	virtSubresourceClient, err := rest.RESTClientFor(virtSubsrcConfig)
	if err != nil {
		return err
	}

	controller := &upgradeHandler{
		ctx:                ctx,
		jobClient:          jobs,
		jobCache:           jobs.Cache(),
		nodeCache:          nodes.Cache(),
		namespace:          options.Namespace,
		upgradeClient:      upgrades,
		upgradeCache:       upgrades.Cache(),
		upgradeController:  upgrades,
		upgradeLogClient:   upgradeLogs,
		upgradeLogCache:    upgradeLogs.Cache(),
		versionCache:       versions.Cache(),
		planClient:         plans,
		planCache:          plans.Cache(),
		managedChartClient: managedcharts,
		managedChartCache:  managedcharts.Cache(),
		vmImageClient:      vmImages,
		vmImageCache:       vmImages.Cache(),
		vmClient:           vms,
		vmCache:            vms.Cache(),
		serviceClient:      services,
		pvcClient:          pvcs,
		clusterClient:      clusters,
		clusterCache:       clusters.Cache(),
		lhSettingClient:    lhSettings,
		lhSettingCache:     lhSettings.Cache(),
		vmRestClient:       virtSubresourceClient,
	}
	upgrades.OnChange(ctx, upgradeControllerName, controller.OnChanged)
	upgrades.OnRemove(ctx, upgradeControllerName, controller.OnRemove)

	planHandler := &planHandler{
		namespace:     options.Namespace,
		upgradeClient: upgrades,
		upgradeCache:  upgrades.Cache(),
		nodeCache:     nodes.Cache(),
		planClient:    plans,
	}
	plans.OnChange(ctx, planControllerName, planHandler.OnChanged)

	jobHandler := &jobHandler{
		namespace:      options.Namespace,
		planCache:      plans.Cache(),
		upgradeClient:  upgrades,
		upgradeCache:   upgrades.Cache(),
		machineCache:   machines.Cache(),
		secretClient:   secrets,
		nodeClient:     nodes,
		nodeCache:      nodes.Cache(),
		jobClient:      jobs,
		jobCache:       jobs.Cache(),
		configMapCache: configMaps.Cache(),
	}
	jobs.OnChange(ctx, jobControllerName, jobHandler.OnChanged)

	podHandler := &podHandler{
		namespace:     options.Namespace,
		planCache:     plans.Cache(),
		upgradeClient: upgrades,
		upgradeCache:  upgrades.Cache(),
	}
	pods.OnChange(ctx, podControllerName, podHandler.OnChanged)

	vmImageHandler := &vmImageHandler{
		namespace:     options.Namespace,
		upgradeClient: upgrades,
		upgradeCache:  upgrades.Cache(),
	}
	vmImages.OnChange(ctx, vmImageControllerName, vmImageHandler.OnChanged)

	secretHandler := &secretHandler{
		namespace:     options.Namespace,
		upgradeClient: upgrades,
		upgradeCache:  upgrades.Cache(),
		jobClient:     jobs,
		jobCache:      jobs.Cache(),
		machineCache:  machines.Cache(),
	}
	secrets.OnChange(ctx, secretControllerName, secretHandler.OnChanged)

	nodeHandler := &nodeHandler{
		namespace:     options.Namespace,
		nodeClient:    nodes,
		nodeCache:     nodes.Cache(),
		upgradeClient: upgrades,
		upgradeCache:  upgrades.Cache(),
		secretClient:  secrets,
	}
	nodes.OnChange(ctx, nodeControllerName, nodeHandler.OnChanged)

	versionSyncer := newVersionSyncer(ctx, options.Namespace, versions, nodes, namespaces)

	settingHandler := settingHandler{
		versionSyncer: versionSyncer,
	}
	settings.OnChange(ctx, settingControllerName, settingHandler.OnChanged)

	go versionSyncer.start()

	return nil
}
