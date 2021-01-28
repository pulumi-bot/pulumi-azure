// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

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
	case "azure:core/customProvider:CustomProvider":
		r, err = NewCustomProvider(ctx, name, nil, pulumi.URN_(urn))
	case "azure:core/resourceGroup:ResourceGroup":
		r, err = NewResourceGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure:core/resourceGroupTemplateDeployment:ResourceGroupTemplateDeployment":
		r, err = NewResourceGroupTemplateDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:core/resourceProviderRegistration:ResourceProviderRegistration":
		r, err = NewResourceProviderRegistration(ctx, name, nil, pulumi.URN_(urn))
	case "azure:core/subscriptionTemplateDeployment:SubscriptionTemplateDeployment":
		r, err = NewSubscriptionTemplateDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:core/templateDeployment:TemplateDeployment":
		r, err = NewTemplateDeployment(ctx, name, nil, pulumi.URN_(urn))
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
		"core/customProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceGroupTemplateDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceProviderRegistration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/subscriptionTemplateDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/templateDeployment",
		&module{version},
	)
}
