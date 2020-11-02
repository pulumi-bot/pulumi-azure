// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Custom Provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		_, err = core.NewCustomProvider(ctx, "exampleCustomProvider", &core.CustomProviderArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ResourceTypes: core.CustomProviderResourceTypeArray{
// 				&core.CustomProviderResourceTypeArgs{
// 					Name:     pulumi.String("dEf1"),
// 					Endpoint: pulumi.String("https://testendpoint.com/"),
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
//
// ## Import
//
// Custom Provider can be imported using the `resource id`, e.g.
type CustomProvider struct {
	pulumi.CustomResourceState

	// Any number of `action` block as defined below. One of `resourceType` or `action` must be specified.
	Actions CustomProviderActionArrayOutput `pulumi:"actions"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Custom Provider.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Any number of `resourceType` block as defined below. One of `resourceType` or `action` must be specified.
	ResourceTypes CustomProviderResourceTypeArrayOutput `pulumi:"resourceTypes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Any number of `validation` block as defined below.
	Validations CustomProviderValidationArrayOutput `pulumi:"validations"`
}

// NewCustomProvider registers a new resource with the given unique name, arguments, and options.
func NewCustomProvider(ctx *pulumi.Context,
	name string, args *CustomProviderArgs, opts ...pulumi.ResourceOption) (*CustomProvider, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CustomProviderArgs{}
	}
	var resource CustomProvider
	err := ctx.RegisterResource("azure:core/customProvider:CustomProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomProvider gets an existing CustomProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomProviderState, opts ...pulumi.ResourceOption) (*CustomProvider, error) {
	var resource CustomProvider
	err := ctx.ReadResource("azure:core/customProvider:CustomProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomProvider resources.
type customProviderState struct {
	// Any number of `action` block as defined below. One of `resourceType` or `action` must be specified.
	Actions []CustomProviderAction `pulumi:"actions"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Custom Provider.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Any number of `resourceType` block as defined below. One of `resourceType` or `action` must be specified.
	ResourceTypes []CustomProviderResourceType `pulumi:"resourceTypes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Any number of `validation` block as defined below.
	Validations []CustomProviderValidation `pulumi:"validations"`
}

type CustomProviderState struct {
	// Any number of `action` block as defined below. One of `resourceType` or `action` must be specified.
	Actions CustomProviderActionArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Custom Provider.
	ResourceGroupName pulumi.StringPtrInput
	// Any number of `resourceType` block as defined below. One of `resourceType` or `action` must be specified.
	ResourceTypes CustomProviderResourceTypeArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Any number of `validation` block as defined below.
	Validations CustomProviderValidationArrayInput
}

func (CustomProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*customProviderState)(nil)).Elem()
}

type customProviderArgs struct {
	// Any number of `action` block as defined below. One of `resourceType` or `action` must be specified.
	Actions []CustomProviderAction `pulumi:"actions"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Custom Provider.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Any number of `resourceType` block as defined below. One of `resourceType` or `action` must be specified.
	ResourceTypes []CustomProviderResourceType `pulumi:"resourceTypes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Any number of `validation` block as defined below.
	Validations []CustomProviderValidation `pulumi:"validations"`
}

// The set of arguments for constructing a CustomProvider resource.
type CustomProviderArgs struct {
	// Any number of `action` block as defined below. One of `resourceType` or `action` must be specified.
	Actions CustomProviderActionArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Custom Provider.
	ResourceGroupName pulumi.StringInput
	// Any number of `resourceType` block as defined below. One of `resourceType` or `action` must be specified.
	ResourceTypes CustomProviderResourceTypeArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Any number of `validation` block as defined below.
	Validations CustomProviderValidationArrayInput
}

func (CustomProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customProviderArgs)(nil)).Elem()
}
