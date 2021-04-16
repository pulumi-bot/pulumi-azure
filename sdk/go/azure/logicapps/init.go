// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

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
	case "azure:logicapps/actionCustom:ActionCustom":
		r = &ActionCustom{}
	case "azure:logicapps/actionHttp:ActionHttp":
		r = &ActionHttp{}
	case "azure:logicapps/integrationAccount:IntegrationAccount":
		r = &IntegrationAccount{}
	case "azure:logicapps/interationServiceEnvironment:InterationServiceEnvironment":
		r = &InterationServiceEnvironment{}
	case "azure:logicapps/triggerCustom:TriggerCustom":
		r = &TriggerCustom{}
	case "azure:logicapps/triggerHttpRequest:TriggerHttpRequest":
		r = &TriggerHttpRequest{}
	case "azure:logicapps/triggerRecurrence:TriggerRecurrence":
		r = &TriggerRecurrence{}
	case "azure:logicapps/workflow:Workflow":
		r = &Workflow{}
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
		"logicapps/actionCustom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/actionHttp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/integrationAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/interationServiceEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/triggerCustom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/triggerHttpRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/triggerRecurrence",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"logicapps/workflow",
		&module{version},
	)
}
