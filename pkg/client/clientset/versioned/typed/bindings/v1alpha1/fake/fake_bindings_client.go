/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/bindings/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBindingsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeBindingsV1alpha1) ProvisionedServices(namespace string) v1alpha1.ProvisionedServiceInterface {
	return &FakeProvisionedServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBindingsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
