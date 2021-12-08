// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v2beta3

import (
	"context"
	"time"

	v2beta3 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/apis/config/v2beta3"
	scheme "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApisixConsumersGetter has a method to return a ApisixConsumerInterface.
// A group's client should implement this interface.
type ApisixConsumersGetter interface {
	ApisixConsumers(namespace string) ApisixConsumerInterface
}

// ApisixConsumerInterface has methods to work with ApisixConsumer resources.
type ApisixConsumerInterface interface {
	Create(ctx context.Context, apisixConsumer *v2beta3.ApisixConsumer, opts v1.CreateOptions) (*v2beta3.ApisixConsumer, error)
	Update(ctx context.Context, apisixConsumer *v2beta3.ApisixConsumer, opts v1.UpdateOptions) (*v2beta3.ApisixConsumer, error)
	UpdateStatus(ctx context.Context, apisixConsumer *v2beta3.ApisixConsumer, opts v1.UpdateOptions) (*v2beta3.ApisixConsumer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta3.ApisixConsumer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta3.ApisixConsumerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta3.ApisixConsumer, err error)
	ApisixConsumerExpansion
}

// apisixConsumers implements ApisixConsumerInterface
type apisixConsumers struct {
	client rest.Interface
	ns     string
}

// newApisixConsumers returns a ApisixConsumers
func newApisixConsumers(c *ApisixV2beta3Client, namespace string) *apisixConsumers {
	return &apisixConsumers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apisixConsumer, and returns the corresponding apisixConsumer object, and an error if there is any.
func (c *apisixConsumers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta3.ApisixConsumer, err error) {
	result = &v2beta3.ApisixConsumer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apisixconsumers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApisixConsumers that match those selectors.
func (c *apisixConsumers) List(ctx context.Context, opts v1.ListOptions) (result *v2beta3.ApisixConsumerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta3.ApisixConsumerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apisixconsumers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apisixConsumers.
func (c *apisixConsumers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apisixconsumers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apisixConsumer and creates it.  Returns the server's representation of the apisixConsumer, and an error, if there is any.
func (c *apisixConsumers) Create(ctx context.Context, apisixConsumer *v2beta3.ApisixConsumer, opts v1.CreateOptions) (result *v2beta3.ApisixConsumer, err error) {
	result = &v2beta3.ApisixConsumer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apisixconsumers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apisixConsumer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apisixConsumer and updates it. Returns the server's representation of the apisixConsumer, and an error, if there is any.
func (c *apisixConsumers) Update(ctx context.Context, apisixConsumer *v2beta3.ApisixConsumer, opts v1.UpdateOptions) (result *v2beta3.ApisixConsumer, err error) {
	result = &v2beta3.ApisixConsumer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apisixconsumers").
		Name(apisixConsumer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apisixConsumer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apisixConsumers) UpdateStatus(ctx context.Context, apisixConsumer *v2beta3.ApisixConsumer, opts v1.UpdateOptions) (result *v2beta3.ApisixConsumer, err error) {
	result = &v2beta3.ApisixConsumer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apisixconsumers").
		Name(apisixConsumer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apisixConsumer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apisixConsumer and deletes it. Returns an error if one occurs.
func (c *apisixConsumers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apisixconsumers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apisixConsumers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apisixconsumers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apisixConsumer.
func (c *apisixConsumers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta3.ApisixConsumer, err error) {
	result = &v2beta3.ApisixConsumer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apisixconsumers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}