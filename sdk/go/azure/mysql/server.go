// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a MySQL Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mysql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewServer(ctx, "exampleServer", &mysql.ServerArgs{
// 			Location:                        exampleResourceGroup.Location,
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			AdministratorLogin:              pulumi.String("mysqladminun"),
// 			AdministratorLoginPassword:      pulumi.String("H@Sh1CoR3!"),
// 			SkuName:                         pulumi.String("B_Gen5_2"),
// 			StorageMb:                       pulumi.Int(5120),
// 			Version:                         pulumi.String("5.7"),
// 			AutoGrowEnabled:                 pulumi.Bool(true),
// 			BackupRetentionDays:             pulumi.Int(7),
// 			GeoRedundantBackupEnabled:       pulumi.Bool(true),
// 			InfrastructureEncryptionEnabled: pulumi.Bool(true),
// 			PublicNetworkAccessEnabled:      pulumi.Bool(false),
// 			SslEnforcementEnabled:           pulumi.Bool(true),
// 			SslMinimalTlsVersionEnforced:    pulumi.String("TLS1_2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Server struct {
	pulumi.CustomResourceState

	// The Administrator Login for the MySQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword pulumi.StringPtrOutput `pulumi:"administratorLoginPassword"`
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled pulumi.BoolOutput `pulumi:"autoGrowEnabled"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntOutput `pulumi:"backupRetentionDays"`
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// For creation modes other than `Default`, the source server ID to use.
	CreationSourceServerId pulumi.StringPtrOutput `pulumi:"creationSourceServerId"`
	// The FQDN of the MySQL Server.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled pulumi.BoolOutput `pulumi:"geoRedundantBackupEnabled"`
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled pulumi.BoolPtrOutput `pulumi:"infrastructureEncryptionEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
	RestorePointInTime pulumi.StringPtrOutput `pulumi:"restorePointInTime"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// Deprecated: this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement pulumi.StringOutput `pulumi:"sslEnforcement"`
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled pulumi.BoolOutput `pulumi:"sslEnforcementEnabled"`
	// The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced pulumi.StringPtrOutput `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb pulumi.IntOutput `pulumi:"storageMb"`
	// Deprecated: all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile ServerStorageProfileOutput `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	if args == nil {
		args = &ServerArgs{}
	}
	var resource Server
	err := ctx.RegisterResource("azure:mysql/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure:mysql/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The Administrator Login for the MySQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled *bool `pulumi:"autoGrowEnabled"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
	CreateMode *string `pulumi:"createMode"`
	// For creation modes other than `Default`, the source server ID to use.
	CreationSourceServerId *string `pulumi:"creationSourceServerId"`
	// The FQDN of the MySQL Server.
	Fqdn *string `pulumi:"fqdn"`
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled *bool `pulumi:"geoRedundantBackupEnabled"`
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled *bool `pulumi:"infrastructureEncryptionEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName *string `pulumi:"skuName"`
	// Deprecated: this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled *bool `pulumi:"sslEnforcementEnabled"`
	// The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced *string `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb *int `pulumi:"storageMb"`
	// Deprecated: all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile *ServerStorageProfile `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version *string `pulumi:"version"`
}

type ServerState struct {
	// The Administrator Login for the MySQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the MySQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword pulumi.StringPtrInput
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled pulumi.BoolPtrInput
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
	CreateMode pulumi.StringPtrInput
	// For creation modes other than `Default`, the source server ID to use.
	CreationSourceServerId pulumi.StringPtrInput
	// The FQDN of the MySQL Server.
	Fqdn pulumi.StringPtrInput
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled pulumi.BoolPtrInput
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
	RestorePointInTime pulumi.StringPtrInput
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName pulumi.StringPtrInput
	// Deprecated: this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement pulumi.StringPtrInput
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled pulumi.BoolPtrInput
	// The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced pulumi.StringPtrInput
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb pulumi.IntPtrInput
	// Deprecated: all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile ServerStorageProfilePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The Administrator Login for the MySQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled *bool `pulumi:"autoGrowEnabled"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
	CreateMode *string `pulumi:"createMode"`
	// For creation modes other than `Default`, the source server ID to use.
	CreationSourceServerId *string `pulumi:"creationSourceServerId"`
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled *bool `pulumi:"geoRedundantBackupEnabled"`
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled *bool `pulumi:"infrastructureEncryptionEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName string `pulumi:"skuName"`
	// Deprecated: this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled *bool `pulumi:"sslEnforcementEnabled"`
	// The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced *string `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb *int `pulumi:"storageMb"`
	// Deprecated: all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile *ServerStorageProfile `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Administrator Login for the MySQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the MySQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword pulumi.StringPtrInput
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled pulumi.BoolPtrInput
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
	CreateMode pulumi.StringPtrInput
	// For creation modes other than `Default`, the source server ID to use.
	CreationSourceServerId pulumi.StringPtrInput
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled pulumi.BoolPtrInput
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
	RestorePointInTime pulumi.StringPtrInput
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName pulumi.StringInput
	// Deprecated: this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement pulumi.StringPtrInput
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled pulumi.BoolPtrInput
	// The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced pulumi.StringPtrInput
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb pulumi.IntPtrInput
	// Deprecated: all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile ServerStorageProfilePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version pulumi.StringInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}
