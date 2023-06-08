/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and clustersecret-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/sap/clustersecret-operator/pkg/apis/core.cs.sap.com/v1alpha1"
	corecssapcomv1alpha1 "github.com/sap/clustersecret-operator/pkg/client/applyconfiguration/core.cs.sap.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterSecrets implements ClusterSecretInterface
type FakeClusterSecrets struct {
	Fake *FakeCoreV1alpha1
}

var clustersecretsResource = v1alpha1.SchemeGroupVersion.WithResource("clustersecrets")

var clustersecretsKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterSecret")

// Get takes name of the clusterSecret, and returns the corresponding clusterSecret object, and an error if there is any.
func (c *FakeClusterSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustersecretsResource, name), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}

// List takes label and field selectors, and returns the list of ClusterSecrets that match those selectors.
func (c *FakeClusterSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustersecretsResource, clustersecretsKind, opts), &v1alpha1.ClusterSecretList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterSecretList{ListMeta: obj.(*v1alpha1.ClusterSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterSecrets.
func (c *FakeClusterSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustersecretsResource, opts))
}

// Create takes the representation of a clusterSecret and creates it.  Returns the server's representation of the clusterSecret, and an error, if there is any.
func (c *FakeClusterSecrets) Create(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.CreateOptions) (result *v1alpha1.ClusterSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustersecretsResource, clusterSecret), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}

// Update takes the representation of a clusterSecret and updates it. Returns the server's representation of the clusterSecret, and an error, if there is any.
func (c *FakeClusterSecrets) Update(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.UpdateOptions) (result *v1alpha1.ClusterSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustersecretsResource, clusterSecret), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterSecrets) UpdateStatus(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.UpdateOptions) (*v1alpha1.ClusterSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustersecretsResource, "status", clusterSecret), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}

// Delete takes name of the clusterSecret and deletes it. Returns an error if one occurs.
func (c *FakeClusterSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustersecretsResource, name, opts), &v1alpha1.ClusterSecret{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustersecretsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterSecretList{})
	return err
}

// Patch applies the patch and returns the patched clusterSecret.
func (c *FakeClusterSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustersecretsResource, name, pt, data, subresources...), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterSecret.
func (c *FakeClusterSecrets) Apply(ctx context.Context, clusterSecret *corecssapcomv1alpha1.ClusterSecretApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSecret, err error) {
	if clusterSecret == nil {
		return nil, fmt.Errorf("clusterSecret provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterSecret)
	if err != nil {
		return nil, err
	}
	name := clusterSecret.Name
	if name == nil {
		return nil, fmt.Errorf("clusterSecret.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustersecretsResource, *name, types.ApplyPatchType, data), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeClusterSecrets) ApplyStatus(ctx context.Context, clusterSecret *corecssapcomv1alpha1.ClusterSecretApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSecret, err error) {
	if clusterSecret == nil {
		return nil, fmt.Errorf("clusterSecret provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterSecret)
	if err != nil {
		return nil, err
	}
	name := clusterSecret.Name
	if name == nil {
		return nil, fmt.Errorf("clusterSecret.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustersecretsResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.ClusterSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSecret), err
}
