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

// FakeOIDCProviders implements OIDCProviderInterface
type FakeOIDCProviders struct {
	Fake *FakeManagementV3
}

var oidcprovidersResource = v3.SchemeGroupVersion.WithResource("oidcproviders")

var oidcprovidersKind = v3.SchemeGroupVersion.WithKind("OIDCProvider")

// Get takes name of the oIDCProvider, and returns the corresponding oIDCProvider object, and an error if there is any.
func (c *FakeOIDCProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.OIDCProvider, err error) {
	emptyResult := &v3.OIDCProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(oidcprovidersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.OIDCProvider), err
}

// List takes label and field selectors, and returns the list of OIDCProviders that match those selectors.
func (c *FakeOIDCProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.OIDCProviderList, err error) {
	emptyResult := &v3.OIDCProviderList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(oidcprovidersResource, oidcprovidersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.OIDCProviderList{ListMeta: obj.(*v3.OIDCProviderList).ListMeta}
	for _, item := range obj.(*v3.OIDCProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oIDCProviders.
func (c *FakeOIDCProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(oidcprovidersResource, opts))
}

// Create takes the representation of a oIDCProvider and creates it.  Returns the server's representation of the oIDCProvider, and an error, if there is any.
func (c *FakeOIDCProviders) Create(ctx context.Context, oIDCProvider *v3.OIDCProvider, opts v1.CreateOptions) (result *v3.OIDCProvider, err error) {
	emptyResult := &v3.OIDCProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(oidcprovidersResource, oIDCProvider, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.OIDCProvider), err
}

// Update takes the representation of a oIDCProvider and updates it. Returns the server's representation of the oIDCProvider, and an error, if there is any.
func (c *FakeOIDCProviders) Update(ctx context.Context, oIDCProvider *v3.OIDCProvider, opts v1.UpdateOptions) (result *v3.OIDCProvider, err error) {
	emptyResult := &v3.OIDCProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(oidcprovidersResource, oIDCProvider, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.OIDCProvider), err
}

// Delete takes name of the oIDCProvider and deletes it. Returns an error if one occurs.
func (c *FakeOIDCProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(oidcprovidersResource, name, opts), &v3.OIDCProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOIDCProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(oidcprovidersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.OIDCProviderList{})
	return err
}

// Patch applies the patch and returns the patched oIDCProvider.
func (c *FakeOIDCProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.OIDCProvider, err error) {
	emptyResult := &v3.OIDCProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(oidcprovidersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.OIDCProvider), err
}
