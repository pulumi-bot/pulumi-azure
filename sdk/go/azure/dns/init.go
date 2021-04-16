// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

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
	case "azure:dns/aRecord:ARecord":
		r = &ARecord{}
	case "azure:dns/aaaaRecord:AaaaRecord":
		r = &AaaaRecord{}
	case "azure:dns/cNameRecord:CNameRecord":
		r = &CNameRecord{}
	case "azure:dns/caaRecord:CaaRecord":
		r = &CaaRecord{}
	case "azure:dns/mxRecord:MxRecord":
		r = &MxRecord{}
	case "azure:dns/nsRecord:NsRecord":
		r = &NsRecord{}
	case "azure:dns/ptrRecord:PtrRecord":
		r = &PtrRecord{}
	case "azure:dns/srvRecord:SrvRecord":
		r = &SrvRecord{}
	case "azure:dns/txtRecord:TxtRecord":
		r = &TxtRecord{}
	case "azure:dns/zone:Zone":
		r = &Zone{}
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
		"dns/aRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/aaaaRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/cNameRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/caaRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/mxRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/nsRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/ptrRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/srvRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/txtRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"dns/zone",
		&module{version},
	)
}
