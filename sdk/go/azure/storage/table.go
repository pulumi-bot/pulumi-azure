// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Table within an Azure Storage Account.
//
// ## Example Usage
//
//
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
// 		_, err = storage.NewTable(ctx, "exampleTable", &storage.TableArgs{
// 			StorageAccountName: exampleAccount.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Table struct {
	pulumi.CustomResourceState

	// One or more `acl` blocks as defined below.
	Acls TableAclArrayOutput `pulumi:"acls"`
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	if args == nil {
		args = &TableArgs{}
	}
	var resource Table
	err := ctx.RegisterResource("azure:storage/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("azure:storage/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// One or more `acl` blocks as defined below.
	Acls []TableAcl `pulumi:"acls"`
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name *string `pulumi:"name"`
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type TableState struct {
	// One or more `acl` blocks as defined below.
	Acls TableAclArrayInput
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// One or more `acl` blocks as defined below.
	Acls []TableAcl `pulumi:"acls"`
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name *string `pulumi:"name"`
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// One or more `acl` blocks as defined below.
	Acls TableAclArrayInput
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}
