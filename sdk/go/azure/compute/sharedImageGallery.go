// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Shared Image Gallery.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		_, err = compute.NewSharedImageGallery(ctx, "exampleSharedImageGallery", &compute.SharedImageGalleryArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Description:       pulumi.String("Shared images and things."),
// 			Tags: pulumi.StringMap{
// 				"Hello": pulumi.String("There"),
// 				"World": pulumi.String("Example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SharedImageGallery struct {
	pulumi.CustomResourceState

	// A description for this Shared Image Gallery.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Shared Image Gallery. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Shared Image Gallery.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Unique Name for this Shared Image Gallery.
	UniqueName pulumi.StringOutput `pulumi:"uniqueName"`
}

// NewSharedImageGallery registers a new resource with the given unique name, arguments, and options.
func NewSharedImageGallery(ctx *pulumi.Context,
	name string, args *SharedImageGalleryArgs, opts ...pulumi.ResourceOption) (*SharedImageGallery, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &SharedImageGalleryArgs{}
	}
	var resource SharedImageGallery
	err := ctx.RegisterResource("azure:compute/sharedImageGallery:SharedImageGallery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedImageGallery gets an existing SharedImageGallery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedImageGallery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedImageGalleryState, opts ...pulumi.ResourceOption) (*SharedImageGallery, error) {
	var resource SharedImageGallery
	err := ctx.ReadResource("azure:compute/sharedImageGallery:SharedImageGallery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedImageGallery resources.
type sharedImageGalleryState struct {
	// A description for this Shared Image Gallery.
	Description *string `pulumi:"description"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Shared Image Gallery. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Shared Image Gallery.
	Tags map[string]string `pulumi:"tags"`
	// The Unique Name for this Shared Image Gallery.
	UniqueName *string `pulumi:"uniqueName"`
}

type SharedImageGalleryState struct {
	// A description for this Shared Image Gallery.
	Description pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Shared Image Gallery. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Shared Image Gallery.
	Tags pulumi.StringMapInput
	// The Unique Name for this Shared Image Gallery.
	UniqueName pulumi.StringPtrInput
}

func (SharedImageGalleryState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageGalleryState)(nil)).Elem()
}

type sharedImageGalleryArgs struct {
	// A description for this Shared Image Gallery.
	Description *string `pulumi:"description"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Shared Image Gallery. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Shared Image Gallery.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SharedImageGallery resource.
type SharedImageGalleryArgs struct {
	// A description for this Shared Image Gallery.
	Description pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Shared Image Gallery. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Shared Image Gallery.
	Tags pulumi.StringMapInput
}

func (SharedImageGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageGalleryArgs)(nil)).Elem()
}

type SharedImageGalleryInput interface {
	pulumi.Input

	ToSharedImageGalleryOutput() SharedImageGalleryOutput
	ToSharedImageGalleryOutputWithContext(ctx context.Context) SharedImageGalleryOutput
}

func (SharedImageGallery) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedImageGallery)(nil)).Elem()
}

func (i SharedImageGallery) ToSharedImageGalleryOutput() SharedImageGalleryOutput {
	return i.ToSharedImageGalleryOutputWithContext(context.Background())
}

func (i SharedImageGallery) ToSharedImageGalleryOutputWithContext(ctx context.Context) SharedImageGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedImageGalleryOutput)
}

type SharedImageGalleryOutput struct {
	*pulumi.OutputState
}

func (SharedImageGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedImageGalleryOutput)(nil)).Elem()
}

func (o SharedImageGalleryOutput) ToSharedImageGalleryOutput() SharedImageGalleryOutput {
	return o
}

func (o SharedImageGalleryOutput) ToSharedImageGalleryOutputWithContext(ctx context.Context) SharedImageGalleryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SharedImageGalleryOutput{})
}
