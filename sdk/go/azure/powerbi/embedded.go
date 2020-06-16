// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package powerbi

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a PowerBI Embedded.
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/powerbi"
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
// 		_, err = powerbi.NewEmbedded(ctx, "exampleEmbedded", &powerbi.EmbeddedArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("A1"),
// 			Administrators: pulumi.StringArray{
// 				pulumi.String("azsdktest@microsoft.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Embedded struct {
	pulumi.CustomResourceState

	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	Administrators pulumi.StringArrayOutput `pulumi:"administrators"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the PowerBI Embedded. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: `A1`, `A2`, `A3`, `A4`, `A5`, `A6`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEmbedded registers a new resource with the given unique name, arguments, and options.
func NewEmbedded(ctx *pulumi.Context,
	name string, args *EmbeddedArgs, opts ...pulumi.ResourceOption) (*Embedded, error) {
	if args == nil || args.Administrators == nil {
		return nil, errors.New("missing required argument 'Administrators'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil {
		args = &EmbeddedArgs{}
	}
	var resource Embedded
	err := ctx.RegisterResource("azure:powerbi/embedded:Embedded", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmbedded gets an existing Embedded resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmbedded(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmbeddedState, opts ...pulumi.ResourceOption) (*Embedded, error) {
	var resource Embedded
	err := ctx.ReadResource("azure:powerbi/embedded:Embedded", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Embedded resources.
type embeddedState struct {
	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	Administrators []string `pulumi:"administrators"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the PowerBI Embedded. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: `A1`, `A2`, `A3`, `A4`, `A5`, `A6`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type EmbeddedState struct {
	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	Administrators pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the PowerBI Embedded. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: `A1`, `A2`, `A3`, `A4`, `A5`, `A6`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EmbeddedState) ElementType() reflect.Type {
	return reflect.TypeOf((*embeddedState)(nil)).Elem()
}

type embeddedArgs struct {
	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	Administrators []string `pulumi:"administrators"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the PowerBI Embedded. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: `A1`, `A2`, `A3`, `A4`, `A5`, `A6`.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Embedded resource.
type EmbeddedArgs struct {
	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	Administrators pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the PowerBI Embedded. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: `A1`, `A2`, `A3`, `A4`, `A5`, `A6`.
	SkuName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EmbeddedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*embeddedArgs)(nil)).Elem()
}
