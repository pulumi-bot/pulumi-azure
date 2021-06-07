// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure"
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
	case "azure:eventgrid/domain:Domain":
		r = &Domain{}
	case "azure:eventgrid/domainTopic:DomainTopic":
		r = &DomainTopic{}
	case "azure:eventgrid/eventSubscription:EventSubscription":
		r = &EventSubscription{}
	case "azure:eventgrid/getSystemTopic:getSystemTopic":
		r = &GetSystemTopic{}
	case "azure:eventgrid/systemTopic:SystemTopic":
		r = &SystemTopic{}
	case "azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription":
		r = &SystemTopicEventSubscription{}
	case "azure:eventgrid/topic:Topic":
		r = &Topic{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version := azure.PkgVersion()
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/domainTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/eventSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/getSystemTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/systemTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/systemTopicEventSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"eventgrid/topic",
		&module{version},
	)
}
