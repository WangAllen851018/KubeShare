/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	kubesharev1 "github.com/lsalab-git/KubeShare/pkg/apis/kubeshare/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharePods implements SharePodInterface
type FakeSharePods struct {
	Fake *FakeKubeshareV1
	ns   string
}

var sharepodsResource = schema.GroupVersionResource{Group: "kubeshare.nthu", Version: "v1", Resource: "sharepods"}

var sharepodsKind = schema.GroupVersionKind{Group: "kubeshare.nthu", Version: "v1", Kind: "SharePod"}

// Get takes name of the sharePod, and returns the corresponding sharePod object, and an error if there is any.
func (c *FakeSharePods) Get(name string, options v1.GetOptions) (result *kubesharev1.SharePod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharepodsResource, c.ns, name), &kubesharev1.SharePod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubesharev1.SharePod), err
}

// List takes label and field selectors, and returns the list of SharePods that match those selectors.
func (c *FakeSharePods) List(opts v1.ListOptions) (result *kubesharev1.SharePodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharepodsResource, sharepodsKind, c.ns, opts), &kubesharev1.SharePodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubesharev1.SharePodList{ListMeta: obj.(*kubesharev1.SharePodList).ListMeta}
	for _, item := range obj.(*kubesharev1.SharePodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharePods.
func (c *FakeSharePods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharepodsResource, c.ns, opts))

}

// Create takes the representation of a sharePod and creates it.  Returns the server's representation of the sharePod, and an error, if there is any.
func (c *FakeSharePods) Create(sharePod *kubesharev1.SharePod) (result *kubesharev1.SharePod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharepodsResource, c.ns, sharePod), &kubesharev1.SharePod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubesharev1.SharePod), err
}

// Update takes the representation of a sharePod and updates it. Returns the server's representation of the sharePod, and an error, if there is any.
func (c *FakeSharePods) Update(sharePod *kubesharev1.SharePod) (result *kubesharev1.SharePod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharepodsResource, c.ns, sharePod), &kubesharev1.SharePod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubesharev1.SharePod), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharePods) UpdateStatus(sharePod *kubesharev1.SharePod) (*kubesharev1.SharePod, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharepodsResource, "status", c.ns, sharePod), &kubesharev1.SharePod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubesharev1.SharePod), err
}

// Delete takes name of the sharePod and deletes it. Returns an error if one occurs.
func (c *FakeSharePods) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharepodsResource, c.ns, name), &kubesharev1.SharePod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharePods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharepodsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &kubesharev1.SharePodList{})
	return err
}

// Patch applies the patch and returns the patched sharePod.
func (c *FakeSharePods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubesharev1.SharePod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharepodsResource, c.ns, name, pt, data, subresources...), &kubesharev1.SharePod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubesharev1.SharePod), err
}