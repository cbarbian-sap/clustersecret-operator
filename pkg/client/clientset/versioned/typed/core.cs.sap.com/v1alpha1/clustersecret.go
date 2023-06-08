/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and clustersecret-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/sap/clustersecret-operator/pkg/apis/core.cs.sap.com/v1alpha1"
	corecssapcomv1alpha1 "github.com/sap/clustersecret-operator/pkg/client/applyconfiguration/core.cs.sap.com/v1alpha1"
	scheme "github.com/sap/clustersecret-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterSecretsGetter has a method to return a ClusterSecretInterface.
// A group's client should implement this interface.
type ClusterSecretsGetter interface {
	ClusterSecrets() ClusterSecretInterface
}

// ClusterSecretInterface has methods to work with ClusterSecret resources.
type ClusterSecretInterface interface {
	Create(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.CreateOptions) (*v1alpha1.ClusterSecret, error)
	Update(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.UpdateOptions) (*v1alpha1.ClusterSecret, error)
	UpdateStatus(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.UpdateOptions) (*v1alpha1.ClusterSecret, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterSecret, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterSecretList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSecret, err error)
	Apply(ctx context.Context, clusterSecret *corecssapcomv1alpha1.ClusterSecretApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSecret, err error)
	ApplyStatus(ctx context.Context, clusterSecret *corecssapcomv1alpha1.ClusterSecretApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSecret, err error)
	ClusterSecretExpansion
}

// clusterSecrets implements ClusterSecretInterface
type clusterSecrets struct {
	client rest.Interface
}

// newClusterSecrets returns a ClusterSecrets
func newClusterSecrets(c *CoreV1alpha1Client) *clusterSecrets {
	return &clusterSecrets{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterSecret, and returns the corresponding clusterSecret object, and an error if there is any.
func (c *clusterSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSecret, err error) {
	result = &v1alpha1.ClusterSecret{}
	err = c.client.Get().
		Resource("clustersecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterSecrets that match those selectors.
func (c *clusterSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSecretList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterSecretList{}
	err = c.client.Get().
		Resource("clustersecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterSecrets.
func (c *clusterSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustersecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterSecret and creates it.  Returns the server's representation of the clusterSecret, and an error, if there is any.
func (c *clusterSecrets) Create(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.CreateOptions) (result *v1alpha1.ClusterSecret, err error) {
	result = &v1alpha1.ClusterSecret{}
	err = c.client.Post().
		Resource("clustersecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSecret).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterSecret and updates it. Returns the server's representation of the clusterSecret, and an error, if there is any.
func (c *clusterSecrets) Update(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.UpdateOptions) (result *v1alpha1.ClusterSecret, err error) {
	result = &v1alpha1.ClusterSecret{}
	err = c.client.Put().
		Resource("clustersecrets").
		Name(clusterSecret.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSecret).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterSecrets) UpdateStatus(ctx context.Context, clusterSecret *v1alpha1.ClusterSecret, opts v1.UpdateOptions) (result *v1alpha1.ClusterSecret, err error) {
	result = &v1alpha1.ClusterSecret{}
	err = c.client.Put().
		Resource("clustersecrets").
		Name(clusterSecret.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSecret).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterSecret and deletes it. Returns an error if one occurs.
func (c *clusterSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustersecrets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustersecrets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterSecret.
func (c *clusterSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSecret, err error) {
	result = &v1alpha1.ClusterSecret{}
	err = c.client.Patch(pt).
		Resource("clustersecrets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterSecret.
func (c *clusterSecrets) Apply(ctx context.Context, clusterSecret *corecssapcomv1alpha1.ClusterSecretApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSecret, err error) {
	if clusterSecret == nil {
		return nil, fmt.Errorf("clusterSecret provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterSecret)
	if err != nil {
		return nil, err
	}
	name := clusterSecret.Name
	if name == nil {
		return nil, fmt.Errorf("clusterSecret.Name must be provided to Apply")
	}
	result = &v1alpha1.ClusterSecret{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("clustersecrets").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *clusterSecrets) ApplyStatus(ctx context.Context, clusterSecret *corecssapcomv1alpha1.ClusterSecretApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSecret, err error) {
	if clusterSecret == nil {
		return nil, fmt.Errorf("clusterSecret provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterSecret)
	if err != nil {
		return nil, err
	}

	name := clusterSecret.Name
	if name == nil {
		return nil, fmt.Errorf("clusterSecret.Name must be provided to Apply")
	}

	result = &v1alpha1.ClusterSecret{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("clustersecrets").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
