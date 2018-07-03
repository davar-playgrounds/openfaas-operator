/*
Copyright 2018 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/openfaas-incubator/openfaas-operator/pkg/client/clientset/versioned/typed/openfaas/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOpenfaasV1alpha2 struct {
	*testing.Fake
}

func (c *FakeOpenfaasV1alpha2) Functions(namespace string) v1alpha2.FunctionInterface {
	return &FakeFunctions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOpenfaasV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
