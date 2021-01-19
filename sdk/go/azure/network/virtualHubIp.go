// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual Hub IP.
//
// > **NOTE** Virtual Hub IP only supports Standard Virtual Hub without Virtual Wan.
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
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Dynamic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.5.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefix:      pulumi.String("10.5.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualHubIp(ctx, "exampleVirtualHubIp", &network.VirtualHubIpArgs{
// 			VirtualHubId:              exampleVirtualHub.ID(),
// 			PrivateIpAddress:          pulumi.String("10.5.1.18"),
// 			PrivateIpAllocationMethod: pulumi.String("Static"),
// 			PublicIpAddressId:         examplePublicIp.ID(),
// 			SubnetId:                  exampleSubnet.ID(),
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
// Virtual Hub IPs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualHubIp:VirtualHubIp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/ipConfigurations/ipConfig1
// ```
type VirtualHubIp struct {
	pulumi.CustomResourceState

	// The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The private IP address of the IP configuration.
	PrivateIpAddress pulumi.StringPtrOutput `pulumi:"privateIpAddress"`
	// The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
	PrivateIpAllocationMethod pulumi.StringPtrOutput `pulumi:"privateIpAllocationMethod"`
	// The ID of the Public IP Address.
	PublicIpAddressId pulumi.StringPtrOutput `pulumi:"publicIpAddressId"`
	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
}

// NewVirtualHubIp registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubIp(ctx *pulumi.Context,
	name string, args *VirtualHubIpArgs, opts ...pulumi.ResourceOption) (*VirtualHubIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource VirtualHubIp
	err := ctx.RegisterResource("azure:network/virtualHubIp:VirtualHubIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubIp gets an existing VirtualHubIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubIpState, opts ...pulumi.ResourceOption) (*VirtualHubIp, error) {
	var resource VirtualHubIp
	err := ctx.ReadResource("azure:network/virtualHubIp:VirtualHubIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubIp resources.
type virtualHubIpState struct {
	// The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The private IP address of the IP configuration.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
	PrivateIpAllocationMethod *string `pulumi:"privateIpAllocationMethod"`
	// The ID of the Public IP Address.
	PublicIpAddressId *string `pulumi:"publicIpAddressId"`
	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type VirtualHubIpState struct {
	// The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The private IP address of the IP configuration.
	PrivateIpAddress pulumi.StringPtrInput
	// The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
	PrivateIpAllocationMethod pulumi.StringPtrInput
	// The ID of the Public IP Address.
	PublicIpAddressId pulumi.StringPtrInput
	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (VirtualHubIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpState)(nil)).Elem()
}

type virtualHubIpArgs struct {
	// The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The private IP address of the IP configuration.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
	PrivateIpAllocationMethod *string `pulumi:"privateIpAllocationMethod"`
	// The ID of the Public IP Address.
	PublicIpAddressId *string `pulumi:"publicIpAddressId"`
	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
	// The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a VirtualHubIp resource.
type VirtualHubIpArgs struct {
	// The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The private IP address of the IP configuration.
	PrivateIpAddress pulumi.StringPtrInput
	// The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
	PrivateIpAllocationMethod pulumi.StringPtrInput
	// The ID of the Public IP Address.
	PublicIpAddressId pulumi.StringPtrInput
	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
	// The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
}

func (VirtualHubIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpArgs)(nil)).Elem()
}

type VirtualHubIpInput interface {
	pulumi.Input

	ToVirtualHubIpOutput() VirtualHubIpOutput
	ToVirtualHubIpOutputWithContext(ctx context.Context) VirtualHubIpOutput
}

func (*VirtualHubIp) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubIp)(nil))
}

func (i *VirtualHubIp) ToVirtualHubIpOutput() VirtualHubIpOutput {
	return i.ToVirtualHubIpOutputWithContext(context.Background())
}

func (i *VirtualHubIp) ToVirtualHubIpOutputWithContext(ctx context.Context) VirtualHubIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpOutput)
}

func (i *VirtualHubIp) ToVirtualHubIpPtrOutput() VirtualHubIpPtrOutput {
	return i.ToVirtualHubIpPtrOutputWithContext(context.Background())
}

func (i *VirtualHubIp) ToVirtualHubIpPtrOutputWithContext(ctx context.Context) VirtualHubIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpPtrOutput)
}

type VirtualHubIpPtrInput interface {
	pulumi.Input

	ToVirtualHubIpPtrOutput() VirtualHubIpPtrOutput
	ToVirtualHubIpPtrOutputWithContext(ctx context.Context) VirtualHubIpPtrOutput
}

type virtualHubIpPtrType VirtualHubIpArgs

func (*virtualHubIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubIp)(nil))
}

func (i *virtualHubIpPtrType) ToVirtualHubIpPtrOutput() VirtualHubIpPtrOutput {
	return i.ToVirtualHubIpPtrOutputWithContext(context.Background())
}

func (i *virtualHubIpPtrType) ToVirtualHubIpPtrOutputWithContext(ctx context.Context) VirtualHubIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpPtrOutput)
}

// VirtualHubIpArrayInput is an input type that accepts VirtualHubIpArray and VirtualHubIpArrayOutput values.
// You can construct a concrete instance of `VirtualHubIpArrayInput` via:
//
//          VirtualHubIpArray{ VirtualHubIpArgs{...} }
type VirtualHubIpArrayInput interface {
	pulumi.Input

	ToVirtualHubIpArrayOutput() VirtualHubIpArrayOutput
	ToVirtualHubIpArrayOutputWithContext(context.Context) VirtualHubIpArrayOutput
}

type VirtualHubIpArray []VirtualHubIpInput

func (VirtualHubIpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VirtualHubIp)(nil))
}

func (i VirtualHubIpArray) ToVirtualHubIpArrayOutput() VirtualHubIpArrayOutput {
	return i.ToVirtualHubIpArrayOutputWithContext(context.Background())
}

func (i VirtualHubIpArray) ToVirtualHubIpArrayOutputWithContext(ctx context.Context) VirtualHubIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpArrayOutput)
}

// VirtualHubIpMapInput is an input type that accepts VirtualHubIpMap and VirtualHubIpMapOutput values.
// You can construct a concrete instance of `VirtualHubIpMapInput` via:
//
//          VirtualHubIpMap{ "key": VirtualHubIpArgs{...} }
type VirtualHubIpMapInput interface {
	pulumi.Input

	ToVirtualHubIpMapOutput() VirtualHubIpMapOutput
	ToVirtualHubIpMapOutputWithContext(context.Context) VirtualHubIpMapOutput
}

type VirtualHubIpMap map[string]VirtualHubIpInput

func (VirtualHubIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VirtualHubIp)(nil))
}

func (i VirtualHubIpMap) ToVirtualHubIpMapOutput() VirtualHubIpMapOutput {
	return i.ToVirtualHubIpMapOutputWithContext(context.Background())
}

func (i VirtualHubIpMap) ToVirtualHubIpMapOutputWithContext(ctx context.Context) VirtualHubIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpMapOutput)
}

type VirtualHubIpOutput struct {
	*pulumi.OutputState
}

func (VirtualHubIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubIp)(nil))
}

func (o VirtualHubIpOutput) ToVirtualHubIpOutput() VirtualHubIpOutput {
	return o
}

func (o VirtualHubIpOutput) ToVirtualHubIpOutputWithContext(ctx context.Context) VirtualHubIpOutput {
	return o
}

func (o VirtualHubIpOutput) ToVirtualHubIpPtrOutput() VirtualHubIpPtrOutput {
	return o.ToVirtualHubIpPtrOutputWithContext(context.Background())
}

func (o VirtualHubIpOutput) ToVirtualHubIpPtrOutputWithContext(ctx context.Context) VirtualHubIpPtrOutput {
	return o.ApplyT(func(v VirtualHubIp) *VirtualHubIp {
		return &v
	}).(VirtualHubIpPtrOutput)
}

type VirtualHubIpPtrOutput struct {
	*pulumi.OutputState
}

func (VirtualHubIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubIp)(nil))
}

func (o VirtualHubIpPtrOutput) ToVirtualHubIpPtrOutput() VirtualHubIpPtrOutput {
	return o
}

func (o VirtualHubIpPtrOutput) ToVirtualHubIpPtrOutputWithContext(ctx context.Context) VirtualHubIpPtrOutput {
	return o
}

type VirtualHubIpArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubIp)(nil))
}

func (o VirtualHubIpArrayOutput) ToVirtualHubIpArrayOutput() VirtualHubIpArrayOutput {
	return o
}

func (o VirtualHubIpArrayOutput) ToVirtualHubIpArrayOutputWithContext(ctx context.Context) VirtualHubIpArrayOutput {
	return o
}

func (o VirtualHubIpArrayOutput) Index(i pulumi.IntInput) VirtualHubIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubIp {
		return vs[0].([]VirtualHubIp)[vs[1].(int)]
	}).(VirtualHubIpOutput)
}

type VirtualHubIpMapOutput struct{ *pulumi.OutputState }

func (VirtualHubIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualHubIp)(nil))
}

func (o VirtualHubIpMapOutput) ToVirtualHubIpMapOutput() VirtualHubIpMapOutput {
	return o
}

func (o VirtualHubIpMapOutput) ToVirtualHubIpMapOutputWithContext(ctx context.Context) VirtualHubIpMapOutput {
	return o
}

func (o VirtualHubIpMapOutput) MapIndex(k pulumi.StringInput) VirtualHubIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualHubIp {
		return vs[0].(map[string]VirtualHubIp)[vs[1].(string)]
	}).(VirtualHubIpOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubIpOutput{})
	pulumi.RegisterOutputType(VirtualHubIpPtrOutput{})
	pulumi.RegisterOutputType(VirtualHubIpArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubIpMapOutput{})
}
