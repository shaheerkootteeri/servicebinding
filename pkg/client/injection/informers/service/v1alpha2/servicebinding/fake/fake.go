/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/vmware-labs/service-bindings/pkg/client/injection/informers/factory/fake"
	servicebinding "github.com/vmware-labs/service-bindings/pkg/client/injection/informers/service/v1alpha2/servicebinding"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = servicebinding.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Service().V1alpha2().ServiceBindings()
	return context.WithValue(ctx, servicebinding.Key{}, inf), inf.Informer()
}
