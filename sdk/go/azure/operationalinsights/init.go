// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

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
	case "azure:operationalinsights/analyticsSolution:AnalyticsSolution":
		r = &AnalyticsSolution{}
	case "azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace":
		r = &AnalyticsWorkspace{}
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
		"operationalinsights/analyticsSolution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"operationalinsights/analyticsWorkspace",
		&module{version},
	)
}
