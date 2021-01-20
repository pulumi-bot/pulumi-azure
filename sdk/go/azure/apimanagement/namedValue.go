// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Named Value.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("pub1"),
// 			PublisherEmail:    pulumi.String("pub1@email.com"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewNamedValue(ctx, "exampleNamedValue", &apimanagement.NamedValueArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			DisplayName:       pulumi.String("ExampleProperty"),
// 			Value:             pulumi.String("Example Value"),
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
// API Management Properties can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/namedValue:NamedValue example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/namedValues/example-apimp
// ```
type NamedValue struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The display name of this API Management Named Value.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrOutput `pulumi:"secret"`
	// A list of tags to be applied to the API Management Named Value.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The value of this API Management Named Value.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewNamedValue registers a new resource with the given unique name, arguments, and options.
func NewNamedValue(ctx *pulumi.Context,
	name string, args *NamedValueArgs, opts ...pulumi.ResourceOption) (*NamedValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource NamedValue
	err := ctx.RegisterResource("azure:apimanagement/namedValue:NamedValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamedValue gets an existing NamedValue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamedValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedValueState, opts ...pulumi.ResourceOption) (*NamedValue, error) {
	var resource NamedValue
	err := ctx.ReadResource("azure:apimanagement/namedValue:NamedValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamedValue resources.
type namedValueState struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The display name of this API Management Named Value.
	DisplayName *string `pulumi:"displayName"`
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Named Value.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Named Value.
	Value *string `pulumi:"value"`
}

type NamedValueState struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The display name of this API Management Named Value.
	DisplayName pulumi.StringPtrInput
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Named Value.
	Tags pulumi.StringArrayInput
	// The value of this API Management Named Value.
	Value pulumi.StringPtrInput
}

func (NamedValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedValueState)(nil)).Elem()
}

type namedValueArgs struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The display name of this API Management Named Value.
	DisplayName string `pulumi:"displayName"`
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Named Value.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Named Value.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a NamedValue resource.
type NamedValueArgs struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The display name of this API Management Named Value.
	DisplayName pulumi.StringInput
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Named Value.
	Tags pulumi.StringArrayInput
	// The value of this API Management Named Value.
	Value pulumi.StringInput
}

func (NamedValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedValueArgs)(nil)).Elem()
}

type NamedValueInput interface {
	pulumi.Input

	ToNamedValueOutput() NamedValueOutput
	ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput
}

func (*NamedValue) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedValue)(nil))
}

func (i *NamedValue) ToNamedValueOutput() NamedValueOutput {
	return i.ToNamedValueOutputWithContext(context.Background())
}

func (i *NamedValue) ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValueOutput)
}

type NamedValueOutput struct {
	*pulumi.OutputState
}

func (NamedValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedValue)(nil))
}

func (o NamedValueOutput) ToNamedValueOutput() NamedValueOutput {
	return o
}

func (o NamedValueOutput) ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamedValueOutput{})
}
