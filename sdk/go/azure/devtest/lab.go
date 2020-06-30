// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Dev Test Lab.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/devtest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = devtest.NewLab(ctx, "exampleLab", &devtest.LabArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Tags: pulumi.StringMap{
// 				"Sydney": pulumi.String("Australia"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Lab struct {
	pulumi.CustomResourceState

	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId pulumi.StringOutput `pulumi:"artifactsStorageAccountId"`
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId pulumi.StringOutput `pulumi:"defaultPremiumStorageAccountId"`
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId pulumi.StringOutput `pulumi:"defaultStorageAccountId"`
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId pulumi.StringOutput `pulumi:"premiumDataDiskStorageAccountId"`
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
}

// NewLab registers a new resource with the given unique name, arguments, and options.
func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LabArgs{}
	}
	var resource Lab
	err := ctx.RegisterResource("azure:devtest/lab:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLab gets an existing Lab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure:devtest/lab:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lab resources.
type labState struct {
	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId *string `pulumi:"artifactsStorageAccountId"`
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId *string `pulumi:"defaultPremiumStorageAccountId"`
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId *string `pulumi:"defaultStorageAccountId"`
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId *string `pulumi:"premiumDataDiskStorageAccountId"`
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

type LabState struct {
	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId pulumi.StringPtrInput
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId pulumi.StringPtrInput
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId pulumi.StringPtrInput
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId pulumi.StringPtrInput
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier pulumi.StringPtrInput
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Lab resource.
type LabArgs struct {
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}
