// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Container within an Azure Storage Account.
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
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Container struct {
	pulumi.CustomResourceState

	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType pulumi.StringPtrOutput `pulumi:"containerAccessType"`
	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy pulumi.BoolOutput `pulumi:"hasImmutabilityPolicy"`
	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold pulumi.BoolOutput `pulumi:"hasLegalHold"`
	// A mapping of MetaData for this Container.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the Container which should be created within the Storage Account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Resource Manager ID of this Storage Container.
	ResourceManagerId pulumi.StringOutput `pulumi:"resourceManagerId"`
	// The name of the Storage Account where the Container should be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	var resource Container
	err := ctx.RegisterResource("azure:storage/container:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("azure:storage/container:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type containerState struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType *string `pulumi:"containerAccessType"`
	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy *bool `pulumi:"hasImmutabilityPolicy"`
	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold *bool `pulumi:"hasLegalHold"`
	// A mapping of MetaData for this Container.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the Container which should be created within the Storage Account.
	Name *string `pulumi:"name"`
	// The Resource Manager ID of this Storage Container.
	ResourceManagerId *string `pulumi:"resourceManagerId"`
	// The name of the Storage Account where the Container should be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type ContainerState struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType pulumi.StringPtrInput
	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy pulumi.BoolPtrInput
	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold pulumi.BoolPtrInput
	// A mapping of MetaData for this Container.
	Metadata pulumi.StringMapInput
	// The name of the Container which should be created within the Storage Account.
	Name pulumi.StringPtrInput
	// The Resource Manager ID of this Storage Container.
	ResourceManagerId pulumi.StringPtrInput
	// The name of the Storage Account where the Container should be created.
	StorageAccountName pulumi.StringPtrInput
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType *string `pulumi:"containerAccessType"`
	// A mapping of MetaData for this Container.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the Container which should be created within the Storage Account.
	Name *string `pulumi:"name"`
	// The name of the Storage Account where the Container should be created.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType pulumi.StringPtrInput
	// A mapping of MetaData for this Container.
	Metadata pulumi.StringMapInput
	// The name of the Container which should be created within the Storage Account.
	Name pulumi.StringPtrInput
	// The name of the Storage Account where the Container should be created.
	StorageAccountName pulumi.StringInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}
