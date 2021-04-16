// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

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
	case "azure:storage/account:Account":
		r = &Account{}
	case "azure:storage/accountNetworkRules:AccountNetworkRules":
		r = &AccountNetworkRules{}
	case "azure:storage/blob:Blob":
		r = &Blob{}
	case "azure:storage/container:Container":
		r = &Container{}
	case "azure:storage/customerManagedKey:CustomerManagedKey":
		r = &CustomerManagedKey{}
	case "azure:storage/dataLakeGen2Filesystem:DataLakeGen2Filesystem":
		r = &DataLakeGen2Filesystem{}
	case "azure:storage/dataLakeGen2Path:DataLakeGen2Path":
		r = &DataLakeGen2Path{}
	case "azure:storage/encryptionScope:EncryptionScope":
		r = &EncryptionScope{}
	case "azure:storage/managementPolicy:ManagementPolicy":
		r = &ManagementPolicy{}
	case "azure:storage/queue:Queue":
		r = &Queue{}
	case "azure:storage/share:Share":
		r = &Share{}
	case "azure:storage/shareDirectory:ShareDirectory":
		r = &ShareDirectory{}
	case "azure:storage/shareFile:ShareFile":
		r = &ShareFile{}
	case "azure:storage/sync:Sync":
		r = &Sync{}
	case "azure:storage/syncCloudEndpoint:SyncCloudEndpoint":
		r = &SyncCloudEndpoint{}
	case "azure:storage/syncGroup:SyncGroup":
		r = &SyncGroup{}
	case "azure:storage/table:Table":
		r = &Table{}
	case "azure:storage/tableEntity:TableEntity":
		r = &TableEntity{}
	case "azure:storage/zipBlob:ZipBlob":
		r = &ZipBlob{}
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
		"storage/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/accountNetworkRules",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/blob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/container",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/customerManagedKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/dataLakeGen2Filesystem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/dataLakeGen2Path",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/encryptionScope",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/managementPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/queue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/share",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/shareDirectory",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/shareFile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/sync",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/syncCloudEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/syncGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/table",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/tableEntity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"storage/zipBlob",
		&module{version},
	)
}
