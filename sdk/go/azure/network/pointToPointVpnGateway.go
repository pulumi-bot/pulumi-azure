// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Point-to-Site VPN Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewPointToPointVpnGateway(ctx, "example", &network.PointToPointVpnGatewayArgs{
// 			Location:                 pulumi.Any(azurerm_resource_group.Example.Location),
// 			ResourceGroupName:        pulumi.Any(azurerm_resource_group.Example.Resource_group_name),
// 			VirtualHubId:             pulumi.Any(azurerm_virtual_hub.Example.Id),
// 			VpnServerConfigurationId: pulumi.Any(azurerm_vpn_server_configuration.Example.Id),
// 			ScaleUnit:                pulumi.Int(1),
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
// Point-to-Site VPN Gateway's can be imported using the `resource id`, e.g.
type PointToPointVpnGateway struct {
	pulumi.CustomResourceState

	// A `connectionConfiguration` block as defined below.
	ConnectionConfiguration PointToPointVpnGatewayConnectionConfigurationOutput `pulumi:"connectionConfiguration"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit pulumi.IntOutput `pulumi:"scaleUnit"`
	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	VpnServerConfigurationId pulumi.StringOutput `pulumi:"vpnServerConfigurationId"`
}

// NewPointToPointVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewPointToPointVpnGateway(ctx *pulumi.Context,
	name string, args *PointToPointVpnGatewayArgs, opts ...pulumi.ResourceOption) (*PointToPointVpnGateway, error) {
	if args == nil || args.ConnectionConfiguration == nil {
		return nil, errors.New("missing required argument 'ConnectionConfiguration'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ScaleUnit == nil {
		return nil, errors.New("missing required argument 'ScaleUnit'")
	}
	if args == nil || args.VirtualHubId == nil {
		return nil, errors.New("missing required argument 'VirtualHubId'")
	}
	if args == nil || args.VpnServerConfigurationId == nil {
		return nil, errors.New("missing required argument 'VpnServerConfigurationId'")
	}
	if args == nil {
		args = &PointToPointVpnGatewayArgs{}
	}
	var resource PointToPointVpnGateway
	err := ctx.RegisterResource("azure:network/pointToPointVpnGateway:PointToPointVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPointToPointVpnGateway gets an existing PointToPointVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPointToPointVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PointToPointVpnGatewayState, opts ...pulumi.ResourceOption) (*PointToPointVpnGateway, error) {
	var resource PointToPointVpnGateway
	err := ctx.ReadResource("azure:network/pointToPointVpnGateway:PointToPointVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PointToPointVpnGateway resources.
type pointToPointVpnGatewayState struct {
	// A `connectionConfiguration` block as defined below.
	ConnectionConfiguration *PointToPointVpnGatewayConnectionConfiguration `pulumi:"connectionConfiguration"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit *int `pulumi:"scaleUnit"`
	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	VpnServerConfigurationId *string `pulumi:"vpnServerConfigurationId"`
}

type PointToPointVpnGatewayState struct {
	// A `connectionConfiguration` block as defined below.
	ConnectionConfiguration PointToPointVpnGatewayConnectionConfigurationPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit pulumi.IntPtrInput
	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	VpnServerConfigurationId pulumi.StringPtrInput
}

func (PointToPointVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*pointToPointVpnGatewayState)(nil)).Elem()
}

type pointToPointVpnGatewayArgs struct {
	// A `connectionConfiguration` block as defined below.
	ConnectionConfiguration PointToPointVpnGatewayConnectionConfiguration `pulumi:"connectionConfiguration"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit int `pulumi:"scaleUnit"`
	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	VpnServerConfigurationId string `pulumi:"vpnServerConfigurationId"`
}

// The set of arguments for constructing a PointToPointVpnGateway resource.
type PointToPointVpnGatewayArgs struct {
	// A `connectionConfiguration` block as defined below.
	ConnectionConfiguration PointToPointVpnGatewayConnectionConfigurationInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit pulumi.IntInput
	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	VpnServerConfigurationId pulumi.StringInput
}

func (PointToPointVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pointToPointVpnGatewayArgs)(nil)).Elem()
}
