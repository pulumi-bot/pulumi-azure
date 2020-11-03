// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Directory within an Azure Storage File Share.
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
// 		exampleShare, err := storage.NewShare(ctx, "exampleShare", &storage.ShareArgs{
// 			StorageAccountName: exampleAccount.Name,
// 			Quota:              pulumi.Int(50),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewShareDirectory(ctx, "exampleShareDirectory", &storage.ShareDirectoryArgs{
// 			ShareName:          exampleShare.Name,
// 			StorageAccountName: exampleAccount.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ShareDirectory struct {
	pulumi.CustomResourceState

	// A mapping of metadata to assign to this Directory.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName pulumi.StringOutput `pulumi:"shareName"`
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewShareDirectory registers a new resource with the given unique name, arguments, and options.
func NewShareDirectory(ctx *pulumi.Context,
	name string, args *ShareDirectoryArgs, opts ...pulumi.ResourceOption) (*ShareDirectory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	var resource ShareDirectory
	err := ctx.RegisterResource("azure:storage/shareDirectory:ShareDirectory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShareDirectory gets an existing ShareDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShareDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareDirectoryState, opts ...pulumi.ResourceOption) (*ShareDirectory, error) {
	var resource ShareDirectory
	err := ctx.ReadResource("azure:storage/shareDirectory:ShareDirectory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShareDirectory resources.
type shareDirectoryState struct {
	// A mapping of metadata to assign to this Directory.
	Metadata map[string]string `pulumi:"metadata"`
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName *string `pulumi:"shareName"`
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type ShareDirectoryState struct {
	// A mapping of metadata to assign to this Directory.
	Metadata pulumi.StringMapInput
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName pulumi.StringPtrInput
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
}

func (ShareDirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareDirectoryState)(nil)).Elem()
}

type shareDirectoryArgs struct {
	// A mapping of metadata to assign to this Directory.
	Metadata map[string]string `pulumi:"metadata"`
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName string `pulumi:"shareName"`
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a ShareDirectory resource.
type ShareDirectoryArgs struct {
	// A mapping of metadata to assign to this Directory.
	Metadata pulumi.StringMapInput
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName pulumi.StringInput
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
}

func (ShareDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareDirectoryArgs)(nil)).Elem()
}
