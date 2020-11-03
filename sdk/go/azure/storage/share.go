// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a File Share within Azure Storage.
//
// > **Note:** The storage share supports two storage tiers: premium and standard. Standard file shares are created in general purpose (GPv1 or GPv2) storage accounts and premium file shares are created in FileStorage storage accounts. For further information, refer to the section "What storage tiers are supported in Azure Files?" of [documentation](https://docs.microsoft.com/en-us/azure/storage/files/storage-files-faq#general).
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewShare(ctx, "exampleShare", &storage.ShareArgs{
// 			StorageAccountName: exampleAccount.Name,
// 			Quota:              pulumi.Int(50),
// 			Acls: storage.ShareAclArray{
// 				&storage.ShareAclArgs{
// 					Id: pulumi.String("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
// 					AccessPolicies: storage.ShareAclAccessPolicyArray{
// 						&storage.ShareAclAccessPolicyArgs{
// 							Permissions: pulumi.String("rwdl"),
// 							Start:       pulumi.String("2019-07-02T09:38:21.0000000Z"),
// 							Expiry:      pulumi.String("2019-07-02T10:38:21.0000000Z"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Share struct {
	pulumi.CustomResourceState

	// One or more `acl` blocks as defined below.
	Acls ShareAclArrayOutput `pulumi:"acls"`
	// A mapping of MetaData for this File Share.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the share. Must be unique within the storage account where the share is located.
	Name pulumi.StringOutput `pulumi:"name"`
	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
	Quota pulumi.IntPtrOutput `pulumi:"quota"`
	// The Resource Manager ID of this File Share.
	ResourceManagerId pulumi.StringOutput `pulumi:"resourceManagerId"`
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// The URL of the File Share
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewShare registers a new resource with the given unique name, arguments, and options.
func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOption) (*Share, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	var resource Share
	err := ctx.RegisterResource("azure:storage/share:Share", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShare gets an existing Share resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareState, opts ...pulumi.ResourceOption) (*Share, error) {
	var resource Share
	err := ctx.ReadResource("azure:storage/share:Share", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Share resources.
type shareState struct {
	// One or more `acl` blocks as defined below.
	Acls []ShareAcl `pulumi:"acls"`
	// A mapping of MetaData for this File Share.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the share. Must be unique within the storage account where the share is located.
	Name *string `pulumi:"name"`
	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
	Quota *int `pulumi:"quota"`
	// The Resource Manager ID of this File Share.
	ResourceManagerId *string `pulumi:"resourceManagerId"`
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The URL of the File Share
	Url *string `pulumi:"url"`
}

type ShareState struct {
	// One or more `acl` blocks as defined below.
	Acls ShareAclArrayInput
	// A mapping of MetaData for this File Share.
	Metadata pulumi.StringMapInput
	// The name of the share. Must be unique within the storage account where the share is located.
	Name pulumi.StringPtrInput
	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
	Quota pulumi.IntPtrInput
	// The Resource Manager ID of this File Share.
	ResourceManagerId pulumi.StringPtrInput
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
	// The URL of the File Share
	Url pulumi.StringPtrInput
}

func (ShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareState)(nil)).Elem()
}

type shareArgs struct {
	// One or more `acl` blocks as defined below.
	Acls []ShareAcl `pulumi:"acls"`
	// A mapping of MetaData for this File Share.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the share. Must be unique within the storage account where the share is located.
	Name *string `pulumi:"name"`
	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
	Quota *int `pulumi:"quota"`
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Share resource.
type ShareArgs struct {
	// One or more `acl` blocks as defined below.
	Acls ShareAclArrayInput
	// A mapping of MetaData for this File Share.
	Metadata pulumi.StringMapInput
	// The name of the share. Must be unique within the storage account where the share is located.
	Name pulumi.StringPtrInput
	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
	Quota pulumi.IntPtrInput
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
}

func (ShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareArgs)(nil)).Elem()
}
