// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v2/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:sentinel/alertRuleFusion:AlertRuleFusion":
		r, err = NewAlertRuleFusion(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident":
		r, err = NewAlertRuleMsSecurityIncident(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sentinel/alertRuleScheduled:AlertRuleScheduled":
		r, err = NewAlertRuleScheduled(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"sentinel/alertRuleFusion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sentinel/alertRuleMsSecurityIncident",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sentinel/alertRuleScheduled",
		&module{version},
	)
}