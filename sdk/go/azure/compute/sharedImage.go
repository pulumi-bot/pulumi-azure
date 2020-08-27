// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Shared Image within a Shared Image Gallery.
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
// 		exampleSharedImageGallery, err := compute.NewSharedImageGallery(ctx, "exampleSharedImageGallery", &compute.SharedImageGalleryArgs{
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
// 		_, err = compute.NewSharedImage(ctx, "exampleSharedImage", &compute.SharedImageArgs{
// 			GalleryName:       exampleSharedImageGallery.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			OsType:            pulumi.String("Linux"),
// 			Identifier: &compute.SharedImageIdentifierArgs{
// 				Publisher: pulumi.String("PublisherName"),
// 				Offer:     pulumi.String("OfferName"),
// 				Sku:       pulumi.String("ExampleSku"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SharedImage struct {
	pulumi.CustomResourceState

	// A description of this Shared Image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula pulumi.StringPtrOutput `pulumi:"eula"`
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName pulumi.StringOutput `pulumi:"galleryName"`
	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifierOutput `pulumi:"identifier"`
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri pulumi.StringPtrOutput `pulumi:"privacyStatementUri"`
	// A `purchasePlan` block as defined below.
	PurchasePlan SharedImagePurchasePlanPtrOutput `pulumi:"purchasePlan"`
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri pulumi.StringPtrOutput `pulumi:"releaseNoteUri"`
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
	Specialized pulumi.BoolPtrOutput `pulumi:"specialized"`
	// A mapping of tags to assign to the Shared Image.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSharedImage registers a new resource with the given unique name, arguments, and options.
func NewSharedImage(ctx *pulumi.Context,
	name string, args *SharedImageArgs, opts ...pulumi.ResourceOption) (*SharedImage, error) {
	if args == nil || args.GalleryName == nil {
		return nil, errors.New("missing required argument 'GalleryName'")
	}
	if args == nil || args.Identifier == nil {
		return nil, errors.New("missing required argument 'Identifier'")
	}
	if args == nil || args.OsType == nil {
		return nil, errors.New("missing required argument 'OsType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &SharedImageArgs{}
	}
	var resource SharedImage
	err := ctx.RegisterResource("azure:compute/sharedImage:SharedImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedImage gets an existing SharedImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedImageState, opts ...pulumi.ResourceOption) (*SharedImage, error) {
	var resource SharedImage
	err := ctx.ReadResource("azure:compute/sharedImage:SharedImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedImage resources.
type sharedImageState struct {
	// A description of this Shared Image.
	Description *string `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula *string `pulumi:"eula"`
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName *string `pulumi:"galleryName"`
	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// An `identifier` block as defined below.
	Identifier *SharedImageIdentifier `pulumi:"identifier"`
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType *string `pulumi:"osType"`
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// A `purchasePlan` block as defined below.
	PurchasePlan *SharedImagePurchasePlan `pulumi:"purchasePlan"`
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
	Specialized *bool `pulumi:"specialized"`
	// A mapping of tags to assign to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
}

type SharedImageState struct {
	// A description of this Shared Image.
	Description pulumi.StringPtrInput
	// The End User Licence Agreement for the Shared Image.
	Eula pulumi.StringPtrInput
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName pulumi.StringPtrInput
	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
	HyperVGeneration pulumi.StringPtrInput
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifierPtrInput
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringPtrInput
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri pulumi.StringPtrInput
	// A `purchasePlan` block as defined below.
	PurchasePlan SharedImagePurchasePlanPtrInput
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
	Specialized pulumi.BoolPtrInput
	// A mapping of tags to assign to the Shared Image.
	Tags pulumi.StringMapInput
}

func (SharedImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageState)(nil)).Elem()
}

type sharedImageArgs struct {
	// A description of this Shared Image.
	Description *string `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula *string `pulumi:"eula"`
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName string `pulumi:"galleryName"`
	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifier `pulumi:"identifier"`
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType string `pulumi:"osType"`
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// A `purchasePlan` block as defined below.
	PurchasePlan *SharedImagePurchasePlan `pulumi:"purchasePlan"`
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
	Specialized *bool `pulumi:"specialized"`
	// A mapping of tags to assign to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SharedImage resource.
type SharedImageArgs struct {
	// A description of this Shared Image.
	Description pulumi.StringPtrInput
	// The End User Licence Agreement for the Shared Image.
	Eula pulumi.StringPtrInput
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName pulumi.StringInput
	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
	HyperVGeneration pulumi.StringPtrInput
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifierInput
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringInput
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri pulumi.StringPtrInput
	// A `purchasePlan` block as defined below.
	PurchasePlan SharedImagePurchasePlanPtrInput
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
	Specialized pulumi.BoolPtrInput
	// A mapping of tags to assign to the Shared Image.
	Tags pulumi.StringMapInput
}

func (SharedImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageArgs)(nil)).Elem()
}
