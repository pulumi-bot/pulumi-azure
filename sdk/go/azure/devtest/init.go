// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

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
	case "azure:devtest/globalVMShutdownSchedule:GlobalVMShutdownSchedule":
		r = &GlobalVMShutdownSchedule{}
	case "azure:devtest/lab:Lab":
		r = &Lab{}
	case "azure:devtest/linuxVirtualMachine:LinuxVirtualMachine":
		r = &LinuxVirtualMachine{}
	case "azure:devtest/policy:Policy":
		r = &Policy{}
	case "azure:devtest/schedule:Schedule":
		r = &Schedule{}
	case "azure:devtest/virtualNetwork:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure:devtest/windowsVirtualMachine:WindowsVirtualMachine":
		r = &WindowsVirtualMachine{}
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
		"devtest/globalVMShutdownSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"devtest/lab",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"devtest/linuxVirtualMachine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"devtest/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"devtest/schedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"devtest/virtualNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"devtest/windowsVirtualMachine",
		&module{version},
	)
}
