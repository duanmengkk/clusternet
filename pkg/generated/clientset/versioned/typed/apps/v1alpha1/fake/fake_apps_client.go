/*
Copyright The Clusternet Authors.

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
	v1alpha1 "github.com/clusternet/clusternet/pkg/generated/clientset/versioned/typed/apps/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAppsV1alpha1) Bases(namespace string) v1alpha1.BaseInterface {
	return newFakeBases(c, namespace)
}

func (c *FakeAppsV1alpha1) Descriptions(namespace string) v1alpha1.DescriptionInterface {
	return newFakeDescriptions(c, namespace)
}

func (c *FakeAppsV1alpha1) FeedInventories(namespace string) v1alpha1.FeedInventoryInterface {
	return newFakeFeedInventories(c, namespace)
}

func (c *FakeAppsV1alpha1) Globalizations() v1alpha1.GlobalizationInterface {
	return newFakeGlobalizations(c)
}

func (c *FakeAppsV1alpha1) HelmCharts(namespace string) v1alpha1.HelmChartInterface {
	return newFakeHelmCharts(c, namespace)
}

func (c *FakeAppsV1alpha1) HelmReleases(namespace string) v1alpha1.HelmReleaseInterface {
	return newFakeHelmReleases(c, namespace)
}

func (c *FakeAppsV1alpha1) Localizations(namespace string) v1alpha1.LocalizationInterface {
	return newFakeLocalizations(c, namespace)
}

func (c *FakeAppsV1alpha1) Manifests(namespace string) v1alpha1.ManifestInterface {
	return newFakeManifests(c, namespace)
}

func (c *FakeAppsV1alpha1) Subscriptions(namespace string) v1alpha1.SubscriptionInterface {
	return newFakeSubscriptions(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
