// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mixedreality

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Spatial Anchors Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mixedreality"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mixedreality.NewSpatialAnchorsAccount(ctx, "exampleSpatialAnchorsAccount", &mixedreality.SpatialAnchorsAccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Spatial Anchors Account can be imported using the `resource id`, e.g.
type SpatialAnchorsAccount struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSpatialAnchorsAccount registers a new resource with the given unique name, arguments, and options.
func NewSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, args *SpatialAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &SpatialAnchorsAccountArgs{}
	}
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpatialAnchorsAccount gets an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpatialAnchorsAccount resources.
type spatialAnchorsAccountState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SpatialAnchorsAccountState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountState)(nil)).Elem()
}

type spatialAnchorsAccountArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SpatialAnchorsAccount resource.
type SpatialAnchorsAccountArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountArgs)(nil)).Elem()
}
