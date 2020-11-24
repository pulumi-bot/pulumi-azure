// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a VPN Gateway within a Virtual Hub, which enables Site-to-Site communication.
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
// 		_, err = network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualWan, err := network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVpnGateway(ctx, "exampleVpnGateway", &network.VpnGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			VirtualHubId:      exampleVirtualHub.ID(),
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
// VPN Gateways can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/vpnGateway:VpnGateway gateway1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/vpnGateways/gateway1
// ```
type VpnGateway struct {
	pulumi.CustomResourceState

	// A `bgpSettings` block as defined below.
	BgpSettings VpnGatewayBgpSettingsOutput `pulumi:"bgpSettings"`
	// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Scale Unit for this VPN Gateway. Defaults to `1`.
	ScaleUnit pulumi.IntPtrOutput `pulumi:"scaleUnit"`
	// A mapping of tags to assign to the VPN Gateway.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
}

// NewVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource VpnGateway
	err := ctx.RegisterResource("azure:network/vpnGateway:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGateway gets an existing VpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("azure:network/vpnGateway:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGateway resources.
type vpnGatewayState struct {
	// A `bgpSettings` block as defined below.
	BgpSettings *VpnGatewayBgpSettings `pulumi:"bgpSettings"`
	// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Scale Unit for this VPN Gateway. Defaults to `1`.
	ScaleUnit *int `pulumi:"scaleUnit"`
	// A mapping of tags to assign to the VPN Gateway.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type VpnGatewayState struct {
	// A `bgpSettings` block as defined below.
	BgpSettings VpnGatewayBgpSettingsPtrInput
	// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Scale Unit for this VPN Gateway. Defaults to `1`.
	ScaleUnit pulumi.IntPtrInput
	// A mapping of tags to assign to the VPN Gateway.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	// A `bgpSettings` block as defined below.
	BgpSettings *VpnGatewayBgpSettings `pulumi:"bgpSettings"`
	// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Scale Unit for this VPN Gateway. Defaults to `1`.
	ScaleUnit *int `pulumi:"scaleUnit"`
	// A mapping of tags to assign to the VPN Gateway.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// A `bgpSettings` block as defined below.
	BgpSettings VpnGatewayBgpSettingsPtrInput
	// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Scale Unit for this VPN Gateway. Defaults to `1`.
	ScaleUnit pulumi.IntPtrInput
	// A mapping of tags to assign to the VPN Gateway.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}

type VpnGatewayInput interface {
	pulumi.Input

	ToVpnGatewayOutput() VpnGatewayOutput
	ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput
}

func (VpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil)).Elem()
}

func (i VpnGateway) ToVpnGatewayOutput() VpnGatewayOutput {
	return i.ToVpnGatewayOutputWithContext(context.Background())
}

func (i VpnGateway) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayOutput)
}

type VpnGatewayOutput struct {
	*pulumi.OutputState
}

func (VpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayOutput)(nil)).Elem()
}

func (o VpnGatewayOutput) ToVpnGatewayOutput() VpnGatewayOutput {
	return o
}

func (o VpnGatewayOutput) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayOutput{})
}
