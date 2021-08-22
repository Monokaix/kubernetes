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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// NamespacesGetter has a method to return a NamespaceInterface.
// A group's client should implement this interface.
type NamespacesGetter interface {
	Namespaces() NamespaceInterface
}

// NamespaceInterface has methods to work with Namespace resources.
type NamespaceInterface interface {
	Create(ctx context.Context, namespace *v1.Namespace, opts metav1.CreateOptions) (*v1.Namespace, error)
	Update(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error)
	UpdateStatus(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Namespace, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NamespaceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Namespace, err error)
	Apply(ctx context.Context, namespace *corev1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Namespace, err error)
	ApplyStatus(ctx context.Context, namespace *corev1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Namespace, err error)
	NamespaceExpansion
}

// namespaces implements NamespaceInterface
type namespaces struct {
	client rest.Interface
}

// newNamespaces returns a Namespaces
func newNamespaces(c *CoreV1Client) *namespaces {
	return &namespaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
// 封装了对namespace的http请求
func (c *namespaces) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Get().
		Resource("namespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *namespaces) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NamespaceList{}
	err = c.client.Get().
		Resource("namespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *namespaces) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("namespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *namespaces) Create(ctx context.Context, namespace *v1.Namespace, opts metav1.CreateOptions) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Post().
		Resource("namespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespace).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a namespace and updates it. Returns the server's representation of the namespace, and an error, if there is any.
func (c *namespaces) Update(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Put().
		Resource("namespaces").
		Name(namespace.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespace).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *namespaces) UpdateStatus(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Put().
		Resource("namespaces").
		Name(namespace.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespace).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *namespaces) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("namespaces").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched namespace.
func (c *namespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Patch(pt).
		Resource("namespaces").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied namespace.
func (c *namespaces) Apply(ctx context.Context, namespace *corev1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Namespace, err error) {
	if namespace == nil {
		return nil, fmt.Errorf("namespace provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(namespace)
	if err != nil {
		return nil, err
	}
	name := namespace.Name
	if name == nil {
		return nil, fmt.Errorf("namespace.Name must be provided to Apply")
	}
	result = &v1.Namespace{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("namespaces").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *namespaces) ApplyStatus(ctx context.Context, namespace *corev1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Namespace, err error) {
	if namespace == nil {
		return nil, fmt.Errorf("namespace provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(namespace)
	if err != nil {
		return nil, err
	}

	name := namespace.Name
	if name == nil {
		return nil, fmt.Errorf("namespace.Name must be provided to Apply")
	}

	result = &v1.Namespace{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("namespaces").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
