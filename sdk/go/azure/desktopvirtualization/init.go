// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

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
	case "azure:desktopvirtualization/applicationGroup:ApplicationGroup":
		r = &ApplicationGroup{}
	case "azure:desktopvirtualization/hostPool:HostPool":
		r = &HostPool{}
	case "azure:desktopvirtualization/workspace:Workspace":
		r = &Workspace{}
	case "azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation":
		r = &WorkspaceApplicationGroupAssociation{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"desktopvirtualization/applicationGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"desktopvirtualization/hostPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"desktopvirtualization/workspace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"desktopvirtualization/workspaceApplicationGroupAssociation",
		&module{version},
	)
}
