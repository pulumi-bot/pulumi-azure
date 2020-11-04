// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Storage Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		_, err = storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Network Rules
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 			ServiceEndpoints: pulumi.StringArray{
// 				pulumi.String("Microsoft.Sql"),
// 				pulumi.String("Microsoft.Storage"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			NetworkRules: &storage.AccountNetworkRulesArgs{
// 				DefaultAction: pulumi.String("Deny"),
// 				IpRules: pulumi.StringArray{
// 					pulumi.String("100.0.0.1"),
// 				},
// 				VirtualNetworkSubnetIds: pulumi.StringArray{
// 					exampleSubnet.ID(),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Account struct {
	pulumi.CustomResourceState

	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier pulumi.StringOutput `pulumi:"accessTier"`
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `StorageV2`.
	AccountKind pulumi.StringPtrOutput `pulumi:"accountKind"`
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS`, `ZRS`, `GZRS` and `RAGZRS`.
	AccountReplicationType pulumi.StringOutput `pulumi:"accountReplicationType"`
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `BlockBlobStorage` and `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier pulumi.StringOutput `pulumi:"accountTier"`
	// Allow or disallow public access to all blobs or containers in the storage account. Defaults to `false`.
	AllowBlobPublicAccess pulumi.BoolPtrOutput `pulumi:"allowBlobPublicAccess"`
	// A `blobProperties` block as defined below.
	BlobProperties AccountBlobPropertiesOutput `pulumi:"blobProperties"`
	// A `customDomain` block as documented below.
	CustomDomain AccountCustomDomainPtrOutput `pulumi:"customDomain"`
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information. Defaults to `true`.
	EnableHttpsTrafficOnly pulumi.BoolPtrOutput `pulumi:"enableHttpsTrafficOnly"`
	// A `identity` block as defined below.
	Identity AccountIdentityOutput `pulumi:"identity"`
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled pulumi.BoolPtrOutput `pulumi:"isHnsEnabled"`
	// Is Large File Share Enabled?
	LargeFileShareEnabled pulumi.BoolOutput `pulumi:"largeFileShareEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The minimum supported TLS version for the storage account. Possible values are `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLS1_0` for new storage accounts.
	MinTlsVersion pulumi.StringPtrOutput `pulumi:"minTlsVersion"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkRules` block as documented below.
	NetworkRules AccountNetworkRulesTypeOutput `pulumi:"networkRules"`
	// The primary access key for the storage account.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The connection string associated with the primary blob location.
	PrimaryBlobConnectionString pulumi.StringOutput `pulumi:"primaryBlobConnectionString"`
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint pulumi.StringOutput `pulumi:"primaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost pulumi.StringOutput `pulumi:"primaryBlobHost"`
	// The connection string associated with the primary location.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint pulumi.StringOutput `pulumi:"primaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost pulumi.StringOutput `pulumi:"primaryDfsHost"`
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint pulumi.StringOutput `pulumi:"primaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost pulumi.StringOutput `pulumi:"primaryFileHost"`
	// The primary location of the storage account.
	PrimaryLocation pulumi.StringOutput `pulumi:"primaryLocation"`
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint pulumi.StringOutput `pulumi:"primaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost pulumi.StringOutput `pulumi:"primaryQueueHost"`
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint pulumi.StringOutput `pulumi:"primaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost pulumi.StringOutput `pulumi:"primaryTableHost"`
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint pulumi.StringOutput `pulumi:"primaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost pulumi.StringOutput `pulumi:"primaryWebHost"`
	// A `queueProperties` block as defined below.
	QueueProperties AccountQueuePropertiesOutput `pulumi:"queueProperties"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary access key for the storage account.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The connection string associated with the secondary blob location.
	SecondaryBlobConnectionString pulumi.StringOutput `pulumi:"secondaryBlobConnectionString"`
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint pulumi.StringOutput `pulumi:"secondaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost pulumi.StringOutput `pulumi:"secondaryBlobHost"`
	// The connection string associated with the secondary location.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint pulumi.StringOutput `pulumi:"secondaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost pulumi.StringOutput `pulumi:"secondaryDfsHost"`
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint pulumi.StringOutput `pulumi:"secondaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost pulumi.StringOutput `pulumi:"secondaryFileHost"`
	// The secondary location of the storage account.
	SecondaryLocation pulumi.StringOutput `pulumi:"secondaryLocation"`
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint pulumi.StringOutput `pulumi:"secondaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost pulumi.StringOutput `pulumi:"secondaryQueueHost"`
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint pulumi.StringOutput `pulumi:"secondaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost pulumi.StringOutput `pulumi:"secondaryTableHost"`
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint pulumi.StringOutput `pulumi:"secondaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost pulumi.StringOutput `pulumi:"secondaryWebHost"`
	// A `staticWebsite` block as defined below.
	StaticWebsite AccountStaticWebsitePtrOutput `pulumi:"staticWebsite"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountReplicationType == nil {
		return nil, errors.New("invalid value for required argument 'AccountReplicationType'")
	}
	if args.AccountTier == nil {
		return nil, errors.New("invalid value for required argument 'AccountTier'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Account
	err := ctx.RegisterResource("azure:storage/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:storage/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier *string `pulumi:"accessTier"`
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `StorageV2`.
	AccountKind *string `pulumi:"accountKind"`
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS`, `ZRS`, `GZRS` and `RAGZRS`.
	AccountReplicationType *string `pulumi:"accountReplicationType"`
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `BlockBlobStorage` and `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier *string `pulumi:"accountTier"`
	// Allow or disallow public access to all blobs or containers in the storage account. Defaults to `false`.
	AllowBlobPublicAccess *bool `pulumi:"allowBlobPublicAccess"`
	// A `blobProperties` block as defined below.
	BlobProperties *AccountBlobProperties `pulumi:"blobProperties"`
	// A `customDomain` block as documented below.
	CustomDomain *AccountCustomDomain `pulumi:"customDomain"`
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information. Defaults to `true`.
	EnableHttpsTrafficOnly *bool `pulumi:"enableHttpsTrafficOnly"`
	// A `identity` block as defined below.
	Identity *AccountIdentity `pulumi:"identity"`
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled *bool `pulumi:"isHnsEnabled"`
	// Is Large File Share Enabled?
	LargeFileShareEnabled *bool `pulumi:"largeFileShareEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The minimum supported TLS version for the storage account. Possible values are `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLS1_0` for new storage accounts.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name *string `pulumi:"name"`
	// A `networkRules` block as documented below.
	NetworkRules *AccountNetworkRulesType `pulumi:"networkRules"`
	// The primary access key for the storage account.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The connection string associated with the primary blob location.
	PrimaryBlobConnectionString *string `pulumi:"primaryBlobConnectionString"`
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint *string `pulumi:"primaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost *string `pulumi:"primaryBlobHost"`
	// The connection string associated with the primary location.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint *string `pulumi:"primaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost *string `pulumi:"primaryDfsHost"`
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint *string `pulumi:"primaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost *string `pulumi:"primaryFileHost"`
	// The primary location of the storage account.
	PrimaryLocation *string `pulumi:"primaryLocation"`
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint *string `pulumi:"primaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost *string `pulumi:"primaryQueueHost"`
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint *string `pulumi:"primaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost *string `pulumi:"primaryTableHost"`
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint *string `pulumi:"primaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost *string `pulumi:"primaryWebHost"`
	// A `queueProperties` block as defined below.
	QueueProperties *AccountQueueProperties `pulumi:"queueProperties"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary access key for the storage account.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The connection string associated with the secondary blob location.
	SecondaryBlobConnectionString *string `pulumi:"secondaryBlobConnectionString"`
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint *string `pulumi:"secondaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost *string `pulumi:"secondaryBlobHost"`
	// The connection string associated with the secondary location.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint *string `pulumi:"secondaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost *string `pulumi:"secondaryDfsHost"`
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint *string `pulumi:"secondaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost *string `pulumi:"secondaryFileHost"`
	// The secondary location of the storage account.
	SecondaryLocation *string `pulumi:"secondaryLocation"`
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint *string `pulumi:"secondaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost *string `pulumi:"secondaryQueueHost"`
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint *string `pulumi:"secondaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost *string `pulumi:"secondaryTableHost"`
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint *string `pulumi:"secondaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost *string `pulumi:"secondaryWebHost"`
	// A `staticWebsite` block as defined below.
	StaticWebsite *AccountStaticWebsite `pulumi:"staticWebsite"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AccountState struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier pulumi.StringPtrInput
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `StorageV2`.
	AccountKind pulumi.StringPtrInput
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS`, `ZRS`, `GZRS` and `RAGZRS`.
	AccountReplicationType pulumi.StringPtrInput
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `BlockBlobStorage` and `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier pulumi.StringPtrInput
	// Allow or disallow public access to all blobs or containers in the storage account. Defaults to `false`.
	AllowBlobPublicAccess pulumi.BoolPtrInput
	// A `blobProperties` block as defined below.
	BlobProperties AccountBlobPropertiesPtrInput
	// A `customDomain` block as documented below.
	CustomDomain AccountCustomDomainPtrInput
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information. Defaults to `true`.
	EnableHttpsTrafficOnly pulumi.BoolPtrInput
	// A `identity` block as defined below.
	Identity AccountIdentityPtrInput
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled pulumi.BoolPtrInput
	// Is Large File Share Enabled?
	LargeFileShareEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The minimum supported TLS version for the storage account. Possible values are `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLS1_0` for new storage accounts.
	MinTlsVersion pulumi.StringPtrInput
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name pulumi.StringPtrInput
	// A `networkRules` block as documented below.
	NetworkRules AccountNetworkRulesTypePtrInput
	// The primary access key for the storage account.
	PrimaryAccessKey pulumi.StringPtrInput
	// The connection string associated with the primary blob location.
	PrimaryBlobConnectionString pulumi.StringPtrInput
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost pulumi.StringPtrInput
	// The connection string associated with the primary location.
	PrimaryConnectionString pulumi.StringPtrInput
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost pulumi.StringPtrInput
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost pulumi.StringPtrInput
	// The primary location of the storage account.
	PrimaryLocation pulumi.StringPtrInput
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost pulumi.StringPtrInput
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost pulumi.StringPtrInput
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost pulumi.StringPtrInput
	// A `queueProperties` block as defined below.
	QueueProperties AccountQueuePropertiesPtrInput
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary access key for the storage account.
	SecondaryAccessKey pulumi.StringPtrInput
	// The connection string associated with the secondary blob location.
	SecondaryBlobConnectionString pulumi.StringPtrInput
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost pulumi.StringPtrInput
	// The connection string associated with the secondary location.
	SecondaryConnectionString pulumi.StringPtrInput
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost pulumi.StringPtrInput
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost pulumi.StringPtrInput
	// The secondary location of the storage account.
	SecondaryLocation pulumi.StringPtrInput
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost pulumi.StringPtrInput
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost pulumi.StringPtrInput
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost pulumi.StringPtrInput
	// A `staticWebsite` block as defined below.
	StaticWebsite AccountStaticWebsitePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier *string `pulumi:"accessTier"`
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `StorageV2`.
	AccountKind *string `pulumi:"accountKind"`
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS`, `ZRS`, `GZRS` and `RAGZRS`.
	AccountReplicationType string `pulumi:"accountReplicationType"`
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `BlockBlobStorage` and `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier string `pulumi:"accountTier"`
	// Allow or disallow public access to all blobs or containers in the storage account. Defaults to `false`.
	AllowBlobPublicAccess *bool `pulumi:"allowBlobPublicAccess"`
	// A `blobProperties` block as defined below.
	BlobProperties *AccountBlobProperties `pulumi:"blobProperties"`
	// A `customDomain` block as documented below.
	CustomDomain *AccountCustomDomain `pulumi:"customDomain"`
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information. Defaults to `true`.
	EnableHttpsTrafficOnly *bool `pulumi:"enableHttpsTrafficOnly"`
	// A `identity` block as defined below.
	Identity *AccountIdentity `pulumi:"identity"`
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled *bool `pulumi:"isHnsEnabled"`
	// Is Large File Share Enabled?
	LargeFileShareEnabled *bool `pulumi:"largeFileShareEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The minimum supported TLS version for the storage account. Possible values are `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLS1_0` for new storage accounts.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name *string `pulumi:"name"`
	// A `networkRules` block as documented below.
	NetworkRules *AccountNetworkRulesType `pulumi:"networkRules"`
	// A `queueProperties` block as defined below.
	QueueProperties *AccountQueueProperties `pulumi:"queueProperties"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `staticWebsite` block as defined below.
	StaticWebsite *AccountStaticWebsite `pulumi:"staticWebsite"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier pulumi.StringPtrInput
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `StorageV2`.
	AccountKind pulumi.StringPtrInput
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS`, `ZRS`, `GZRS` and `RAGZRS`.
	AccountReplicationType pulumi.StringInput
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `BlockBlobStorage` and `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier pulumi.StringInput
	// Allow or disallow public access to all blobs or containers in the storage account. Defaults to `false`.
	AllowBlobPublicAccess pulumi.BoolPtrInput
	// A `blobProperties` block as defined below.
	BlobProperties AccountBlobPropertiesPtrInput
	// A `customDomain` block as documented below.
	CustomDomain AccountCustomDomainPtrInput
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information. Defaults to `true`.
	EnableHttpsTrafficOnly pulumi.BoolPtrInput
	// A `identity` block as defined below.
	Identity AccountIdentityPtrInput
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled pulumi.BoolPtrInput
	// Is Large File Share Enabled?
	LargeFileShareEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The minimum supported TLS version for the storage account. Possible values are `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLS1_0` for new storage accounts.
	MinTlsVersion pulumi.StringPtrInput
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name pulumi.StringPtrInput
	// A `networkRules` block as documented below.
	NetworkRules AccountNetworkRulesTypePtrInput
	// A `queueProperties` block as defined below.
	QueueProperties AccountQueuePropertiesPtrInput
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `staticWebsite` block as defined below.
	StaticWebsite AccountStaticWebsitePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
