/*
Copyright 2024 Rancher Labs, Inc.

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

	v1 "github.com/rancher/rancher/pkg/apis/rke.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRKEClusters implements RKEClusterInterface
type FakeRKEClusters struct {
	Fake *FakeRkeV1
	ns   string
}

var rkeclustersResource = v1.SchemeGroupVersion.WithResource("rkeclusters")

var rkeclustersKind = v1.SchemeGroupVersion.WithKind("RKECluster")

// Get takes name of the rKECluster, and returns the corresponding rKECluster object, and an error if there is any.
func (c *FakeRKEClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RKECluster, err error) {
	emptyResult := &v1.RKECluster{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(rkeclustersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RKECluster), err
}

// List takes label and field selectors, and returns the list of RKEClusters that match those selectors.
func (c *FakeRKEClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RKEClusterList, err error) {
	emptyResult := &v1.RKEClusterList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(rkeclustersResource, rkeclustersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RKEClusterList{ListMeta: obj.(*v1.RKEClusterList).ListMeta}
	for _, item := range obj.(*v1.RKEClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rKEClusters.
func (c *FakeRKEClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(rkeclustersResource, c.ns, opts))

}

// Create takes the representation of a rKECluster and creates it.  Returns the server's representation of the rKECluster, and an error, if there is any.
func (c *FakeRKEClusters) Create(ctx context.Context, rKECluster *v1.RKECluster, opts metav1.CreateOptions) (result *v1.RKECluster, err error) {
	emptyResult := &v1.RKECluster{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(rkeclustersResource, c.ns, rKECluster, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RKECluster), err
}

// Update takes the representation of a rKECluster and updates it. Returns the server's representation of the rKECluster, and an error, if there is any.
func (c *FakeRKEClusters) Update(ctx context.Context, rKECluster *v1.RKECluster, opts metav1.UpdateOptions) (result *v1.RKECluster, err error) {
	emptyResult := &v1.RKECluster{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(rkeclustersResource, c.ns, rKECluster, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RKECluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRKEClusters) UpdateStatus(ctx context.Context, rKECluster *v1.RKECluster, opts metav1.UpdateOptions) (result *v1.RKECluster, err error) {
	emptyResult := &v1.RKECluster{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(rkeclustersResource, "status", c.ns, rKECluster, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RKECluster), err
}

// Delete takes name of the rKECluster and deletes it. Returns an error if one occurs.
func (c *FakeRKEClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rkeclustersResource, c.ns, name, opts), &v1.RKECluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRKEClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(rkeclustersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RKEClusterList{})
	return err
}

// Patch applies the patch and returns the patched rKECluster.
func (c *FakeRKEClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RKECluster, err error) {
	emptyResult := &v1.RKECluster{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(rkeclustersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RKECluster), err
}
