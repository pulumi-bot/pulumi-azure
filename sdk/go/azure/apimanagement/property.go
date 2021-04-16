// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an API Management Property.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = apimanagement.NewProperty(ctx, "exampleProperty", &apimanagement.PropertyArgs{
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
//  $ pulumi import azure:apimanagement/property:Property example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/namedValues/example-apimp
// ```
type Property struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The display name of this API Management Property.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrOutput `pulumi:"secret"`
	// A list of tags to be applied to the API Management Property.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The value of this API Management Property.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewProperty registers a new resource with the given unique name, arguments, and options.
func NewProperty(ctx *pulumi.Context,
	name string, args *PropertyArgs, opts ...pulumi.ResourceOption) (*Property, error) {
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
	var resource Property
	err := ctx.RegisterResource("azure:apimanagement/property:Property", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProperty gets an existing Property resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyState, opts ...pulumi.ResourceOption) (*Property, error) {
	var resource Property
	err := ctx.ReadResource("azure:apimanagement/property:Property", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Property resources.
type propertyState struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The display name of this API Management Property.
	DisplayName *string `pulumi:"displayName"`
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Property.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Property.
	Value *string `pulumi:"value"`
}

type PropertyState struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The display name of this API Management Property.
	DisplayName pulumi.StringPtrInput
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Property.
	Tags pulumi.StringArrayInput
	// The value of this API Management Property.
	Value pulumi.StringPtrInput
}

func (PropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyState)(nil)).Elem()
}

type propertyArgs struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The display name of this API Management Property.
	DisplayName string `pulumi:"displayName"`
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Property.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Property.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Property resource.
type PropertyArgs struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The display name of this API Management Property.
	DisplayName pulumi.StringInput
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Property.
	Tags pulumi.StringArrayInput
	// The value of this API Management Property.
	Value pulumi.StringInput
}

func (PropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyArgs)(nil)).Elem()
}

type PropertyInput interface {
	pulumi.Input

	ToPropertyOutput() PropertyOutput
	ToPropertyOutputWithContext(ctx context.Context) PropertyOutput
}

func (*Property) ElementType() reflect.Type {
	return reflect.TypeOf((*Property)(nil))
}

func (i *Property) ToPropertyOutput() PropertyOutput {
	return i.ToPropertyOutputWithContext(context.Background())
}

func (i *Property) ToPropertyOutputWithContext(ctx context.Context) PropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyOutput)
}

func (i *Property) ToPropertyPtrOutput() PropertyPtrOutput {
	return i.ToPropertyPtrOutputWithContext(context.Background())
}

func (i *Property) ToPropertyPtrOutputWithContext(ctx context.Context) PropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyPtrOutput)
}

type PropertyPtrInput interface {
	pulumi.Input

	ToPropertyPtrOutput() PropertyPtrOutput
	ToPropertyPtrOutputWithContext(ctx context.Context) PropertyPtrOutput
}

type propertyPtrType PropertyArgs

func (*propertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Property)(nil))
}

func (i *propertyPtrType) ToPropertyPtrOutput() PropertyPtrOutput {
	return i.ToPropertyPtrOutputWithContext(context.Background())
}

func (i *propertyPtrType) ToPropertyPtrOutputWithContext(ctx context.Context) PropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyPtrOutput)
}

// PropertyArrayInput is an input type that accepts PropertyArray and PropertyArrayOutput values.
// You can construct a concrete instance of `PropertyArrayInput` via:
//
//          PropertyArray{ PropertyArgs{...} }
type PropertyArrayInput interface {
	pulumi.Input

	ToPropertyArrayOutput() PropertyArrayOutput
	ToPropertyArrayOutputWithContext(context.Context) PropertyArrayOutput
}

type PropertyArray []PropertyInput

func (PropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Property)(nil))
}

func (i PropertyArray) ToPropertyArrayOutput() PropertyArrayOutput {
	return i.ToPropertyArrayOutputWithContext(context.Background())
}

func (i PropertyArray) ToPropertyArrayOutputWithContext(ctx context.Context) PropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyArrayOutput)
}

// PropertyMapInput is an input type that accepts PropertyMap and PropertyMapOutput values.
// You can construct a concrete instance of `PropertyMapInput` via:
//
//          PropertyMap{ "key": PropertyArgs{...} }
type PropertyMapInput interface {
	pulumi.Input

	ToPropertyMapOutput() PropertyMapOutput
	ToPropertyMapOutputWithContext(context.Context) PropertyMapOutput
}

type PropertyMap map[string]PropertyInput

func (PropertyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Property)(nil))
}

func (i PropertyMap) ToPropertyMapOutput() PropertyMapOutput {
	return i.ToPropertyMapOutputWithContext(context.Background())
}

func (i PropertyMap) ToPropertyMapOutputWithContext(ctx context.Context) PropertyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMapOutput)
}

type PropertyOutput struct {
	*pulumi.OutputState
}

func (PropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Property)(nil))
}

func (o PropertyOutput) ToPropertyOutput() PropertyOutput {
	return o
}

func (o PropertyOutput) ToPropertyOutputWithContext(ctx context.Context) PropertyOutput {
	return o
}

func (o PropertyOutput) ToPropertyPtrOutput() PropertyPtrOutput {
	return o.ToPropertyPtrOutputWithContext(context.Background())
}

func (o PropertyOutput) ToPropertyPtrOutputWithContext(ctx context.Context) PropertyPtrOutput {
	return o.ApplyT(func(v Property) *Property {
		return &v
	}).(PropertyPtrOutput)
}

type PropertyPtrOutput struct {
	*pulumi.OutputState
}

func (PropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Property)(nil))
}

func (o PropertyPtrOutput) ToPropertyPtrOutput() PropertyPtrOutput {
	return o
}

func (o PropertyPtrOutput) ToPropertyPtrOutputWithContext(ctx context.Context) PropertyPtrOutput {
	return o
}

type PropertyArrayOutput struct{ *pulumi.OutputState }

func (PropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Property)(nil))
}

func (o PropertyArrayOutput) ToPropertyArrayOutput() PropertyArrayOutput {
	return o
}

func (o PropertyArrayOutput) ToPropertyArrayOutputWithContext(ctx context.Context) PropertyArrayOutput {
	return o
}

func (o PropertyArrayOutput) Index(i pulumi.IntInput) PropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Property {
		return vs[0].([]Property)[vs[1].(int)]
	}).(PropertyOutput)
}

type PropertyMapOutput struct{ *pulumi.OutputState }

func (PropertyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Property)(nil))
}

func (o PropertyMapOutput) ToPropertyMapOutput() PropertyMapOutput {
	return o
}

func (o PropertyMapOutput) ToPropertyMapOutputWithContext(ctx context.Context) PropertyMapOutput {
	return o
}

func (o PropertyMapOutput) MapIndex(k pulumi.StringInput) PropertyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Property {
		return vs[0].(map[string]Property)[vs[1].(string)]
	}).(PropertyOutput)
}

func init() {
	pulumi.RegisterOutputType(PropertyOutput{})
	pulumi.RegisterOutputType(PropertyPtrOutput{})
	pulumi.RegisterOutputType(PropertyArrayOutput{})
	pulumi.RegisterOutputType(PropertyMapOutput{})
}
