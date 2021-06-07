// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

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
	case "azure:containerservice/group:Group":
		r = &Group{}
	case "azure:containerservice/kubernetesCluster:KubernetesCluster":
		r = &KubernetesCluster{}
	case "azure:containerservice/kubernetesClusterNodePool:KubernetesClusterNodePool":
		r = &KubernetesClusterNodePool{}
	case "azure:containerservice/registry:Registry":
		r = &Registry{}
	case "azure:containerservice/registryScopeMap:RegistryScopeMap":
		r = &RegistryScopeMap{}
	case "azure:containerservice/registryToken:RegistryToken":
		r = &RegistryToken{}
	case "azure:containerservice/registryWebhook:RegistryWebhook":
		r = &RegistryWebhook{}
	case "azure:containerservice/registryWebook:RegistryWebook":
		r = &RegistryWebook{}
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
		"containerservice/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/kubernetesCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/kubernetesClusterNodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/registry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/registryScopeMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/registryToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/registryWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"containerservice/registryWebook",
		&module{version},
	)
}
