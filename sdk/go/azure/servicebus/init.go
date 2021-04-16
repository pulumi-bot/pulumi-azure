// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:servicebus/namespace:Namespace":
		r = &Namespace{}
	case "azure:servicebus/namespaceAuthorizationRule:NamespaceAuthorizationRule":
		r = &NamespaceAuthorizationRule{}
	case "azure:servicebus/namespaceNetworkRuleSet:NamespaceNetworkRuleSet":
		r = &NamespaceNetworkRuleSet{}
	case "azure:servicebus/queue:Queue":
		r = &Queue{}
	case "azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule":
		r = &QueueAuthorizationRule{}
	case "azure:servicebus/subscription:Subscription":
		r = &Subscription{}
	case "azure:servicebus/subscriptionRule:SubscriptionRule":
		r = &SubscriptionRule{}
	case "azure:servicebus/topic:Topic":
		r = &Topic{}
	case "azure:servicebus/topicAuthorizationRule:TopicAuthorizationRule":
		r = &TopicAuthorizationRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/namespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/namespaceAuthorizationRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/namespaceNetworkRuleSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/queue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/queueAuthorizationRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/subscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/subscriptionRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/topic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicebus/topicAuthorizationRule",
		&module{version},
	)
}
