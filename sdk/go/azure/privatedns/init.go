// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

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
	case "azure:privatedns/aAAARecord:AAAARecord":
		r = &AAAARecord{}
	case "azure:privatedns/aRecord:ARecord":
		r = &ARecord{}
	case "azure:privatedns/cnameRecord:CnameRecord":
		r = &CnameRecord{}
	case "azure:privatedns/linkService:LinkService":
		r = &LinkService{}
	case "azure:privatedns/mxRecord:MxRecord":
		r = &MxRecord{}
	case "azure:privatedns/pTRRecord:PTRRecord":
		r = &PTRRecord{}
	case "azure:privatedns/sRVRecord:SRVRecord":
		r = &SRVRecord{}
	case "azure:privatedns/txtRecord:TxtRecord":
		r = &TxtRecord{}
	case "azure:privatedns/zone:Zone":
		r = &Zone{}
	case "azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink":
		r = &ZoneVirtualNetworkLink{}
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
		"privatedns/aAAARecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/aRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/cnameRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/linkService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/mxRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/pTRRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/sRVRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/txtRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/zone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"privatedns/zoneVirtualNetworkLink",
		&module{version},
	)
}
