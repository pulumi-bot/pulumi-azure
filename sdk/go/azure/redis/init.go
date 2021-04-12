// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure"
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
	case "azure:redis/cache:Cache":
		r = &Cache{}
	case "azure:redis/enterpriseCluster:EnterpriseCluster":
		r = &EnterpriseCluster{}
	case "azure:redis/enterpriseDatabase:EnterpriseDatabase":
		r = &EnterpriseDatabase{}
	case "azure:redis/firewallRule:FirewallRule":
		r = &FirewallRule{}
	case "azure:redis/linkedServer:LinkedServer":
		r = &LinkedServer{}
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
		"redis/cache",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"redis/enterpriseCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"redis/enterpriseDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"redis/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"redis/linkedServer",
		&module{version},
	)
}
