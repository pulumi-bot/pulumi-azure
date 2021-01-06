// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual WAN.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
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
// 		_, err = network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
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
// Virtual WAN can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualWan:VirtualWan example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualWans/testvwan
// ```
type VirtualWan struct {
	pulumi.CustomResourceState

	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic pulumi.BoolPtrOutput `pulumi:"allowBranchToBranchTraffic"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	AllowVnetToVnetTraffic pulumi.BoolPtrOutput `pulumi:"allowVnetToVnetTraffic"`
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption pulumi.BoolPtrOutput `pulumi:"disableVpnEncryption"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory pulumi.StringPtrOutput `pulumi:"office365LocalBreakoutCategory"`
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Virtual WAN.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewVirtualWan registers a new resource with the given unique name, arguments, and options.
func NewVirtualWan(ctx *pulumi.Context,
	name string, args *VirtualWanArgs, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource VirtualWan
	err := ctx.RegisterResource("azure:network/virtualWan:VirtualWan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualWan gets an existing VirtualWan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualWan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWanState, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	var resource VirtualWan
	err := ctx.ReadResource("azure:network/virtualWan:VirtualWan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualWan resources.
type virtualWanState struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory *string `pulumi:"office365LocalBreakoutCategory"`
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Virtual WAN.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type *string `pulumi:"type"`
}

type VirtualWanState struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory pulumi.StringPtrInput
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Virtual WAN.
	Tags pulumi.StringMapInput
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type pulumi.StringPtrInput
}

func (VirtualWanState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanState)(nil)).Elem()
}

type virtualWanArgs struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory *string `pulumi:"office365LocalBreakoutCategory"`
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Virtual WAN.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a VirtualWan resource.
type VirtualWanArgs struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory pulumi.StringPtrInput
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Virtual WAN.
	Tags pulumi.StringMapInput
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type pulumi.StringPtrInput
}

func (VirtualWanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanArgs)(nil)).Elem()
}

type VirtualWanInput interface {
	pulumi.Input

	ToVirtualWanOutput() VirtualWanOutput
	ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput
}

func (*VirtualWan) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualWan)(nil))
}

func (i *VirtualWan) ToVirtualWanOutput() VirtualWanOutput {
	return i.ToVirtualWanOutputWithContext(context.Background())
}

func (i *VirtualWan) ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanOutput)
}

func (i *VirtualWan) ToVirtualWanPtrOutput() VirtualWanPtrOutput {
	return i.ToVirtualWanPtrOutputWithContext(context.Background())
}

func (i *VirtualWan) ToVirtualWanPtrOutputWithContext(ctx context.Context) VirtualWanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanPtrOutput)
}

type VirtualWanPtrInput interface {
	pulumi.Input

	ToVirtualWanPtrOutput() VirtualWanPtrOutput
	ToVirtualWanPtrOutputWithContext(ctx context.Context) VirtualWanPtrOutput
}

type virtualWanPtrType VirtualWanArgs

func (*virtualWanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualWan)(nil))
}

func (i *virtualWanPtrType) ToVirtualWanPtrOutput() VirtualWanPtrOutput {
	return i.ToVirtualWanPtrOutputWithContext(context.Background())
}

func (i *virtualWanPtrType) ToVirtualWanPtrOutputWithContext(ctx context.Context) VirtualWanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanOutput).ToVirtualWanPtrOutput()
}

// VirtualWanArrayInput is an input type that accepts VirtualWanArray and VirtualWanArrayOutput values.
// You can construct a concrete instance of `VirtualWanArrayInput` via:
//
//          VirtualWanArray{ VirtualWanArgs{...} }
type VirtualWanArrayInput interface {
	pulumi.Input

	ToVirtualWanArrayOutput() VirtualWanArrayOutput
	ToVirtualWanArrayOutputWithContext(context.Context) VirtualWanArrayOutput
}

type VirtualWanArray []VirtualWanInput

func (VirtualWanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualWan)(nil))
}

func (i VirtualWanArray) ToVirtualWanArrayOutput() VirtualWanArrayOutput {
	return i.ToVirtualWanArrayOutputWithContext(context.Background())
}

func (i VirtualWanArray) ToVirtualWanArrayOutputWithContext(ctx context.Context) VirtualWanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanArrayOutput)
}

// VirtualWanMapInput is an input type that accepts VirtualWanMap and VirtualWanMapOutput values.
// You can construct a concrete instance of `VirtualWanMapInput` via:
//
//          VirtualWanMap{ "key": VirtualWanArgs{...} }
type VirtualWanMapInput interface {
	pulumi.Input

	ToVirtualWanMapOutput() VirtualWanMapOutput
	ToVirtualWanMapOutputWithContext(context.Context) VirtualWanMapOutput
}

type VirtualWanMap map[string]VirtualWanInput

func (VirtualWanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualWan)(nil))
}

func (i VirtualWanMap) ToVirtualWanMapOutput() VirtualWanMapOutput {
	return i.ToVirtualWanMapOutputWithContext(context.Background())
}

func (i VirtualWanMap) ToVirtualWanMapOutputWithContext(ctx context.Context) VirtualWanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanMapOutput)
}

type VirtualWanOutput struct {
	*pulumi.OutputState
}

func (VirtualWanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualWan)(nil))
}

func (o VirtualWanOutput) ToVirtualWanOutput() VirtualWanOutput {
	return o
}

func (o VirtualWanOutput) ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput {
	return o
}

func (o VirtualWanOutput) ToVirtualWanPtrOutput() VirtualWanPtrOutput {
	return o.ToVirtualWanPtrOutputWithContext(context.Background())
}

func (o VirtualWanOutput) ToVirtualWanPtrOutputWithContext(ctx context.Context) VirtualWanPtrOutput {
	return o.ApplyT(func(v VirtualWan) *VirtualWan {
		return &v
	}).(VirtualWanPtrOutput)
}

type VirtualWanPtrOutput struct {
	*pulumi.OutputState
}

func (VirtualWanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualWan)(nil))
}

func (o VirtualWanPtrOutput) ToVirtualWanPtrOutput() VirtualWanPtrOutput {
	return o
}

func (o VirtualWanPtrOutput) ToVirtualWanPtrOutputWithContext(ctx context.Context) VirtualWanPtrOutput {
	return o
}

type VirtualWanArrayOutput struct{ *pulumi.OutputState }

func (VirtualWanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualWan)(nil))
}

func (o VirtualWanArrayOutput) ToVirtualWanArrayOutput() VirtualWanArrayOutput {
	return o
}

func (o VirtualWanArrayOutput) ToVirtualWanArrayOutputWithContext(ctx context.Context) VirtualWanArrayOutput {
	return o
}

func (o VirtualWanArrayOutput) Index(i pulumi.IntInput) VirtualWanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualWan {
		return vs[0].([]VirtualWan)[vs[1].(int)]
	}).(VirtualWanOutput)
}

type VirtualWanMapOutput struct{ *pulumi.OutputState }

func (VirtualWanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualWan)(nil))
}

func (o VirtualWanMapOutput) ToVirtualWanMapOutput() VirtualWanMapOutput {
	return o
}

func (o VirtualWanMapOutput) ToVirtualWanMapOutputWithContext(ctx context.Context) VirtualWanMapOutput {
	return o
}

func (o VirtualWanMapOutput) MapIndex(k pulumi.StringInput) VirtualWanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualWan {
		return vs[0].(map[string]VirtualWan)[vs[1].(string)]
	}).(VirtualWanOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualWanOutput{})
	pulumi.RegisterOutputType(VirtualWanPtrOutput{})
	pulumi.RegisterOutputType(VirtualWanArrayOutput{})
	pulumi.RegisterOutputType(VirtualWanMapOutput{})
}
