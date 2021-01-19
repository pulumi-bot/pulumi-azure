// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Resource Group.
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
// 		_, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
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
// Resource Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/resourceGroup:ResourceGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example
// ```
type ResourceGroup struct {
	pulumi.CustomResourceState

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A mapping of tags which should be assigned to the Resource Group.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil {
		args = &ResourceGroupArgs{}
	}

	var resource ResourceGroup
	err := ctx.RegisterResource("azure:core/resourceGroup:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("azure:core/resourceGroup:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroup resources.
type resourceGroupState struct {
	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
	Name *string `pulumi:"name"`
	// A mapping of tags which should be assigned to the Resource Group.
	Tags map[string]string `pulumi:"tags"`
}

type ResourceGroupState struct {
	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
	Name pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Resource Group.
	Tags pulumi.StringMapInput
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
	Name *string `pulumi:"name"`
	// A mapping of tags which should be assigned to the Resource Group.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
	Name pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Resource Group.
	Tags pulumi.StringMapInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

type ResourceGroupInput interface {
	pulumi.Input

	ToResourceGroupOutput() ResourceGroupOutput
	ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput
}

func (*ResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroup)(nil))
}

func (i *ResourceGroup) ToResourceGroupOutput() ResourceGroupOutput {
	return i.ToResourceGroupOutputWithContext(context.Background())
}

func (i *ResourceGroup) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupOutput)
}

func (i *ResourceGroup) ToResourceGroupPtrOutput() ResourceGroupPtrOutput {
	return i.ToResourceGroupPtrOutputWithContext(context.Background())
}

func (i *ResourceGroup) ToResourceGroupPtrOutputWithContext(ctx context.Context) ResourceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPtrOutput)
}

type ResourceGroupPtrInput interface {
	pulumi.Input

	ToResourceGroupPtrOutput() ResourceGroupPtrOutput
	ToResourceGroupPtrOutputWithContext(ctx context.Context) ResourceGroupPtrOutput
}

type resourceGroupPtrType ResourceGroupArgs

func (*resourceGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil))
}

func (i *resourceGroupPtrType) ToResourceGroupPtrOutput() ResourceGroupPtrOutput {
	return i.ToResourceGroupPtrOutputWithContext(context.Background())
}

func (i *resourceGroupPtrType) ToResourceGroupPtrOutputWithContext(ctx context.Context) ResourceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPtrOutput)
}

// ResourceGroupArrayInput is an input type that accepts ResourceGroupArray and ResourceGroupArrayOutput values.
// You can construct a concrete instance of `ResourceGroupArrayInput` via:
//
//          ResourceGroupArray{ ResourceGroupArgs{...} }
type ResourceGroupArrayInput interface {
	pulumi.Input

	ToResourceGroupArrayOutput() ResourceGroupArrayOutput
	ToResourceGroupArrayOutputWithContext(context.Context) ResourceGroupArrayOutput
}

type ResourceGroupArray []ResourceGroupInput

func (ResourceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResourceGroup)(nil))
}

func (i ResourceGroupArray) ToResourceGroupArrayOutput() ResourceGroupArrayOutput {
	return i.ToResourceGroupArrayOutputWithContext(context.Background())
}

func (i ResourceGroupArray) ToResourceGroupArrayOutputWithContext(ctx context.Context) ResourceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupArrayOutput)
}

// ResourceGroupMapInput is an input type that accepts ResourceGroupMap and ResourceGroupMapOutput values.
// You can construct a concrete instance of `ResourceGroupMapInput` via:
//
//          ResourceGroupMap{ "key": ResourceGroupArgs{...} }
type ResourceGroupMapInput interface {
	pulumi.Input

	ToResourceGroupMapOutput() ResourceGroupMapOutput
	ToResourceGroupMapOutputWithContext(context.Context) ResourceGroupMapOutput
}

type ResourceGroupMap map[string]ResourceGroupInput

func (ResourceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResourceGroup)(nil))
}

func (i ResourceGroupMap) ToResourceGroupMapOutput() ResourceGroupMapOutput {
	return i.ToResourceGroupMapOutputWithContext(context.Background())
}

func (i ResourceGroupMap) ToResourceGroupMapOutputWithContext(ctx context.Context) ResourceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupMapOutput)
}

type ResourceGroupOutput struct {
	*pulumi.OutputState
}

func (ResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroup)(nil))
}

func (o ResourceGroupOutput) ToResourceGroupOutput() ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupPtrOutput() ResourceGroupPtrOutput {
	return o.ToResourceGroupPtrOutputWithContext(context.Background())
}

func (o ResourceGroupOutput) ToResourceGroupPtrOutputWithContext(ctx context.Context) ResourceGroupPtrOutput {
	return o.ApplyT(func(v ResourceGroup) *ResourceGroup {
		return &v
	}).(ResourceGroupPtrOutput)
}

type ResourceGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ResourceGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil))
}

func (o ResourceGroupPtrOutput) ToResourceGroupPtrOutput() ResourceGroupPtrOutput {
	return o
}

func (o ResourceGroupPtrOutput) ToResourceGroupPtrOutputWithContext(ctx context.Context) ResourceGroupPtrOutput {
	return o
}

type ResourceGroupArrayOutput struct{ *pulumi.OutputState }

func (ResourceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGroup)(nil))
}

func (o ResourceGroupArrayOutput) ToResourceGroupArrayOutput() ResourceGroupArrayOutput {
	return o
}

func (o ResourceGroupArrayOutput) ToResourceGroupArrayOutputWithContext(ctx context.Context) ResourceGroupArrayOutput {
	return o
}

func (o ResourceGroupArrayOutput) Index(i pulumi.IntInput) ResourceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceGroup {
		return vs[0].([]ResourceGroup)[vs[1].(int)]
	}).(ResourceGroupOutput)
}

type ResourceGroupMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroup)(nil))
}

func (o ResourceGroupMapOutput) ToResourceGroupMapOutput() ResourceGroupMapOutput {
	return o
}

func (o ResourceGroupMapOutput) ToResourceGroupMapOutputWithContext(ctx context.Context) ResourceGroupMapOutput {
	return o
}

func (o ResourceGroupMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroup {
		return vs[0].(map[string]ResourceGroup)[vs[1].(string)]
	}).(ResourceGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGroupOutput{})
	pulumi.RegisterOutputType(ResourceGroupPtrOutput{})
	pulumi.RegisterOutputType(ResourceGroupArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupMapOutput{})
}
