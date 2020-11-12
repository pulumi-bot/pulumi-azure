// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a connection in an existing Virtual Network Gateway.
//
// ## Example Usage
// ### Site-to-Site connection
//
// The following example shows a connection between an Azure virtual network
// and an on-premises VPN device and network.
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
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		onpremiseLocalNetworkGateway, err := network.NewLocalNetworkGateway(ctx, "onpremiseLocalNetworkGateway", &network.LocalNetworkGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			GatewayAddress:    pulumi.String("168.62.225.23"),
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.1.1.0/24"),
// 			},
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
// 		exampleVirtualNetworkGateway, err := network.NewVirtualNetworkGateway(ctx, "exampleVirtualNetworkGateway", &network.VirtualNetworkGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Type:              pulumi.String("Vpn"),
// 			VpnType:           pulumi.String("RouteBased"),
// 			ActiveActive:      pulumi.Bool(false),
// 			EnableBgp:         pulumi.Bool(false),
// 			Sku:               pulumi.String("Basic"),
// 			IpConfigurations: network.VirtualNetworkGatewayIpConfigurationArray{
// 				&network.VirtualNetworkGatewayIpConfigurationArgs{
// 					PublicIpAddressId:          examplePublicIp.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 					SubnetId:                   exampleSubnet.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualNetworkGatewayConnection(ctx, "onpremiseVirtualNetworkGatewayConnection", &network.VirtualNetworkGatewayConnectionArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			Type:                    pulumi.String("IPsec"),
// 			VirtualNetworkGatewayId: exampleVirtualNetworkGateway.ID(),
// 			LocalNetworkGatewayId:   onpremiseLocalNetworkGateway.ID(),
// 			SharedKey:               pulumi.String("4-v3ry-53cr37-1p53c-5h4r3d-k3y"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### VNet-to-VNet connection
//
// The following example shows a connection between two Azure virtual network
// in different locations/regions.
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
// 		usResourceGroup, err := core.NewResourceGroup(ctx, "usResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		usVirtualNetwork, err := network.NewVirtualNetwork(ctx, "usVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          usResourceGroup.Location,
// 			ResourceGroupName: usResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		usGateway, err := network.NewSubnet(ctx, "usGateway", &network.SubnetArgs{
// 			ResourceGroupName:  usResourceGroup.Name,
// 			VirtualNetworkName: usVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		usPublicIp, err := network.NewPublicIp(ctx, "usPublicIp", &network.PublicIpArgs{
// 			Location:          usResourceGroup.Location,
// 			ResourceGroupName: usResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Dynamic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		usVirtualNetworkGateway, err := network.NewVirtualNetworkGateway(ctx, "usVirtualNetworkGateway", &network.VirtualNetworkGatewayArgs{
// 			Location:          usResourceGroup.Location,
// 			ResourceGroupName: usResourceGroup.Name,
// 			Type:              pulumi.String("Vpn"),
// 			VpnType:           pulumi.String("RouteBased"),
// 			Sku:               pulumi.String("Basic"),
// 			IpConfigurations: network.VirtualNetworkGatewayIpConfigurationArray{
// 				&network.VirtualNetworkGatewayIpConfigurationArgs{
// 					PublicIpAddressId:          usPublicIp.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 					SubnetId:                   usGateway.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		europeResourceGroup, err := core.NewResourceGroup(ctx, "europeResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		europeVirtualNetwork, err := network.NewVirtualNetwork(ctx, "europeVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          europeResourceGroup.Location,
// 			ResourceGroupName: europeResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.1.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		europeGateway, err := network.NewSubnet(ctx, "europeGateway", &network.SubnetArgs{
// 			ResourceGroupName:  europeResourceGroup.Name,
// 			VirtualNetworkName: europeVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.1.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		europePublicIp, err := network.NewPublicIp(ctx, "europePublicIp", &network.PublicIpArgs{
// 			Location:          europeResourceGroup.Location,
// 			ResourceGroupName: europeResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Dynamic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		europeVirtualNetworkGateway, err := network.NewVirtualNetworkGateway(ctx, "europeVirtualNetworkGateway", &network.VirtualNetworkGatewayArgs{
// 			Location:          europeResourceGroup.Location,
// 			ResourceGroupName: europeResourceGroup.Name,
// 			Type:              pulumi.String("Vpn"),
// 			VpnType:           pulumi.String("RouteBased"),
// 			Sku:               pulumi.String("Basic"),
// 			IpConfigurations: network.VirtualNetworkGatewayIpConfigurationArray{
// 				&network.VirtualNetworkGatewayIpConfigurationArgs{
// 					PublicIpAddressId:          europePublicIp.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 					SubnetId:                   europeGateway.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualNetworkGatewayConnection(ctx, "usToEurope", &network.VirtualNetworkGatewayConnectionArgs{
// 			Location:                    usResourceGroup.Location,
// 			ResourceGroupName:           usResourceGroup.Name,
// 			Type:                        pulumi.String("Vnet2Vnet"),
// 			VirtualNetworkGatewayId:     usVirtualNetworkGateway.ID(),
// 			PeerVirtualNetworkGatewayId: europeVirtualNetworkGateway.ID(),
// 			SharedKey:                   pulumi.String("4-v3ry-53cr37-1p53c-5h4r3d-k3y"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualNetworkGatewayConnection(ctx, "europeToUs", &network.VirtualNetworkGatewayConnectionArgs{
// 			Location:                    europeResourceGroup.Location,
// 			ResourceGroupName:           europeResourceGroup.Name,
// 			Type:                        pulumi.String("Vnet2Vnet"),
// 			VirtualNetworkGatewayId:     europeVirtualNetworkGateway.ID(),
// 			PeerVirtualNetworkGatewayId: usVirtualNetworkGateway.ID(),
// 			SharedKey:                   pulumi.String("4-v3ry-53cr37-1p53c-5h4r3d-k3y"),
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
// Virtual Network Gateway Connections can be imported using their `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection exampleConnection /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1
// ```
type VirtualNetworkGatewayConnection struct {
	pulumi.CustomResourceState

	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	// The IKE protocol version to use. Possible
	// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
	// Changing this value will force a resource to be created.
	// > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
	ConnectionProtocol pulumi.StringOutput `pulumi:"connectionProtocol"`
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp pulumi.BoolOutput `pulumi:"enableBgp"`
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId pulumi.StringPtrOutput `pulumi:"expressRouteCircuitId"`
	// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
	ExpressRouteGatewayBypass pulumi.BoolOutput `pulumi:"expressRouteGatewayBypass"`
	// A `ipsecPolicy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy VirtualNetworkGatewayConnectionIpsecPolicyPtrOutput `pulumi:"ipsecPolicy"`
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId pulumi.StringPtrOutput `pulumi:"localNetworkGatewayId"`
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId pulumi.StringPtrOutput `pulumi:"peerVirtualNetworkGatewayId"`
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The routing weight. Defaults to `10`.
	RoutingWeight pulumi.IntOutput `pulumi:"routingWeight"`
	// The shared IPSec key. A key could be provided if a
	// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
	// Only one block can be defined for a connection.
	// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
	TrafficSelectorPolicy VirtualNetworkGatewayConnectionTrafficSelectorPolicyPtrOutput `pulumi:"trafficSelectorPolicy"`
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsecPolicy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors pulumi.BoolOutput `pulumi:"usePolicyBasedTrafficSelectors"`
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId pulumi.StringOutput `pulumi:"virtualNetworkGatewayId"`
}

// NewVirtualNetworkGatewayConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.VirtualNetworkGatewayId == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkGatewayId'")
	}
	if args == nil {
		args = &VirtualNetworkGatewayConnectionArgs{}
	}
	var resource VirtualNetworkGatewayConnection
	err := ctx.RegisterResource("azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkGatewayConnection gets an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	var resource VirtualNetworkGatewayConnection
	err := ctx.ReadResource("azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkGatewayConnection resources.
type virtualNetworkGatewayConnectionState struct {
	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The IKE protocol version to use. Possible
	// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
	// Changing this value will force a resource to be created.
	// > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
	ConnectionProtocol *string `pulumi:"connectionProtocol"`
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp *bool `pulumi:"enableBgp"`
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId *string `pulumi:"expressRouteCircuitId"`
	// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
	ExpressRouteGatewayBypass *bool `pulumi:"expressRouteGatewayBypass"`
	// A `ipsecPolicy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy *VirtualNetworkGatewayConnectionIpsecPolicy `pulumi:"ipsecPolicy"`
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId *string `pulumi:"localNetworkGatewayId"`
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId *string `pulumi:"peerVirtualNetworkGatewayId"`
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The routing weight. Defaults to `10`.
	RoutingWeight *int `pulumi:"routingWeight"`
	// The shared IPSec key. A key could be provided if a
	// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
	SharedKey *string `pulumi:"sharedKey"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
	// Only one block can be defined for a connection.
	// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
	TrafficSelectorPolicy *VirtualNetworkGatewayConnectionTrafficSelectorPolicy `pulumi:"trafficSelectorPolicy"`
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type *string `pulumi:"type"`
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsecPolicy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors *bool `pulumi:"usePolicyBasedTrafficSelectors"`
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId *string `pulumi:"virtualNetworkGatewayId"`
}

type VirtualNetworkGatewayConnectionState struct {
	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey pulumi.StringPtrInput
	// The IKE protocol version to use. Possible
	// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
	// Changing this value will force a resource to be created.
	// > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
	ConnectionProtocol pulumi.StringPtrInput
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp pulumi.BoolPtrInput
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId pulumi.StringPtrInput
	// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
	ExpressRouteGatewayBypass pulumi.BoolPtrInput
	// A `ipsecPolicy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy VirtualNetworkGatewayConnectionIpsecPolicyPtrInput
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId pulumi.StringPtrInput
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The routing weight. Defaults to `10`.
	RoutingWeight pulumi.IntPtrInput
	// The shared IPSec key. A key could be provided if a
	// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
	SharedKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
	// Only one block can be defined for a connection.
	// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
	TrafficSelectorPolicy VirtualNetworkGatewayConnectionTrafficSelectorPolicyPtrInput
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type pulumi.StringPtrInput
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsecPolicy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId pulumi.StringPtrInput
}

func (VirtualNetworkGatewayConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionState)(nil)).Elem()
}

type virtualNetworkGatewayConnectionArgs struct {
	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The IKE protocol version to use. Possible
	// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
	// Changing this value will force a resource to be created.
	// > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
	ConnectionProtocol *string `pulumi:"connectionProtocol"`
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp *bool `pulumi:"enableBgp"`
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId *string `pulumi:"expressRouteCircuitId"`
	// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
	ExpressRouteGatewayBypass *bool `pulumi:"expressRouteGatewayBypass"`
	// A `ipsecPolicy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy *VirtualNetworkGatewayConnectionIpsecPolicy `pulumi:"ipsecPolicy"`
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId *string `pulumi:"localNetworkGatewayId"`
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId *string `pulumi:"peerVirtualNetworkGatewayId"`
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The routing weight. Defaults to `10`.
	RoutingWeight *int `pulumi:"routingWeight"`
	// The shared IPSec key. A key could be provided if a
	// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
	SharedKey *string `pulumi:"sharedKey"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
	// Only one block can be defined for a connection.
	// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
	TrafficSelectorPolicy *VirtualNetworkGatewayConnectionTrafficSelectorPolicy `pulumi:"trafficSelectorPolicy"`
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type string `pulumi:"type"`
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsecPolicy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors *bool `pulumi:"usePolicyBasedTrafficSelectors"`
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId string `pulumi:"virtualNetworkGatewayId"`
}

// The set of arguments for constructing a VirtualNetworkGatewayConnection resource.
type VirtualNetworkGatewayConnectionArgs struct {
	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey pulumi.StringPtrInput
	// The IKE protocol version to use. Possible
	// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
	// Changing this value will force a resource to be created.
	// > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
	ConnectionProtocol pulumi.StringPtrInput
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp pulumi.BoolPtrInput
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId pulumi.StringPtrInput
	// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
	ExpressRouteGatewayBypass pulumi.BoolPtrInput
	// A `ipsecPolicy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy VirtualNetworkGatewayConnectionIpsecPolicyPtrInput
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId pulumi.StringPtrInput
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The routing weight. Defaults to `10`.
	RoutingWeight pulumi.IntPtrInput
	// The shared IPSec key. A key could be provided if a
	// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
	SharedKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
	// Only one block can be defined for a connection.
	// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
	TrafficSelectorPolicy VirtualNetworkGatewayConnectionTrafficSelectorPolicyPtrInput
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type pulumi.StringInput
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsecPolicy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId pulumi.StringInput
}

func (VirtualNetworkGatewayConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionArgs)(nil)).Elem()
}

type VirtualNetworkGatewayConnectionInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput
	ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput
}

func (VirtualNetworkGatewayConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnection)(nil)).Elem()
}

func (i VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return i.ToVirtualNetworkGatewayConnectionOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayConnectionOutput)
}

type VirtualNetworkGatewayConnectionOutput struct {
	*pulumi.OutputState
}

func (VirtualNetworkGatewayConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionOutput)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionOutput{})
}
