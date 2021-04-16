// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

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
	case "azure:media/asset:Asset":
		r = &Asset{}
	case "azure:media/contentKeyPolicy:ContentKeyPolicy":
		r = &ContentKeyPolicy{}
	case "azure:media/job:Job":
		r = &Job{}
	case "azure:media/liveEvent:LiveEvent":
		r = &LiveEvent{}
	case "azure:media/serviceAccount:ServiceAccount":
		r = &ServiceAccount{}
	case "azure:media/streamingEndpoint:StreamingEndpoint":
		r = &StreamingEndpoint{}
	case "azure:media/streamingLocator:StreamingLocator":
		r = &StreamingLocator{}
	case "azure:media/streamingPolicy:StreamingPolicy":
		r = &StreamingPolicy{}
	case "azure:media/transform:Transform":
		r = &Transform{}
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
		"media/asset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/contentKeyPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/job",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/liveEvent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/serviceAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/streamingEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/streamingLocator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/streamingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"media/transform",
		&module{version},
	)
}
