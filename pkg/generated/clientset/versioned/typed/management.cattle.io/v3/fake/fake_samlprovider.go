/*
Copyright 2025 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSamlProviders implements SamlProviderInterface
type FakeSamlProviders struct {
	Fake *FakeManagementV3
}

var samlprovidersResource = v3.SchemeGroupVersion.WithResource("samlproviders")

var samlprovidersKind = v3.SchemeGroupVersion.WithKind("SamlProvider")

// Get takes name of the samlProvider, and returns the corresponding samlProvider object, and an error if there is any.
func (c *FakeSamlProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.SamlProvider, err error) {
	emptyResult := &v3.SamlProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(samlprovidersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.SamlProvider), err
}

// List takes label and field selectors, and returns the list of SamlProviders that match those selectors.
func (c *FakeSamlProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.SamlProviderList, err error) {
	emptyResult := &v3.SamlProviderList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(samlprovidersResource, samlprovidersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.SamlProviderList{ListMeta: obj.(*v3.SamlProviderList).ListMeta}
	for _, item := range obj.(*v3.SamlProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested samlProviders.
func (c *FakeSamlProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(samlprovidersResource, opts))
}

// Create takes the representation of a samlProvider and creates it.  Returns the server's representation of the samlProvider, and an error, if there is any.
func (c *FakeSamlProviders) Create(ctx context.Context, samlProvider *v3.SamlProvider, opts v1.CreateOptions) (result *v3.SamlProvider, err error) {
	emptyResult := &v3.SamlProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(samlprovidersResource, samlProvider, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.SamlProvider), err
}

// Update takes the representation of a samlProvider and updates it. Returns the server's representation of the samlProvider, and an error, if there is any.
func (c *FakeSamlProviders) Update(ctx context.Context, samlProvider *v3.SamlProvider, opts v1.UpdateOptions) (result *v3.SamlProvider, err error) {
	emptyResult := &v3.SamlProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(samlprovidersResource, samlProvider, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.SamlProvider), err
}

// Delete takes name of the samlProvider and deletes it. Returns an error if one occurs.
func (c *FakeSamlProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(samlprovidersResource, name, opts), &v3.SamlProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSamlProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(samlprovidersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.SamlProviderList{})
	return err
}

// Patch applies the patch and returns the patched samlProvider.
func (c *FakeSamlProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.SamlProvider, err error) {
	emptyResult := &v3.SamlProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(samlprovidersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.SamlProvider), err
}
