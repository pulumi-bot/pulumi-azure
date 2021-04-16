// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

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
	case "azure:appinsights/analyticsItem:AnalyticsItem":
		r = &AnalyticsItem{}
	case "azure:appinsights/apiKey:ApiKey":
		r = &ApiKey{}
	case "azure:appinsights/insights:Insights":
		r = &Insights{}
	case "azure:appinsights/smartDetectionRule:SmartDetectionRule":
		r = &SmartDetectionRule{}
	case "azure:appinsights/webTest:WebTest":
		r = &WebTest{}
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
		"appinsights/analyticsItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appinsights/apiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appinsights/insights",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appinsights/smartDetectionRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appinsights/webTest",
		&module{version},
	)
}
