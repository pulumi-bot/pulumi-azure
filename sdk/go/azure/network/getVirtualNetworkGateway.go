// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Virtual Network Gateway.
func LookupVirtualNetworkGateway(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayResult, error) {
	var rv LookupVirtualNetworkGatewayResult
	err := ctx.Invoke("azure:network/getVirtualNetworkGateway:getVirtualNetworkGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualNetworkGateway.
type LookupVirtualNetworkGatewayArgs struct {
	// Specifies the name of the Virtual Network Gateway.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Virtual Network Gateway is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getVirtualNetworkGateway.
type LookupVirtualNetworkGatewayResult struct {
	// Is this an Active-Active Gateway?
	ActiveActive bool                                 `pulumi:"activeActive"`
	BgpSettings  []GetVirtualNetworkGatewayBgpSetting `pulumi:"bgpSettings"`
	// The ID of the local network gateway
	// through which outbound Internet traffic from the virtual network in which the
	// gateway is created will be routed (*forced tunneling*). Refer to the
	// [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
	DefaultLocalNetworkGatewayId string `pulumi:"defaultLocalNetworkGatewayId"`
	// Will BGP (Border Gateway Protocol) will be enabled
	// for this Virtual Network Gateway.
	EnableBgp bool `pulumi:"enableBgp"`
	// The Generation of the Virtual Network Gateway.
	Generation string `pulumi:"generation"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// One or two `ipConfiguration` blocks documented below.
	IpConfigurations []GetVirtualNetworkGatewayIpConfiguration `pulumi:"ipConfigurations"`
	// The location/region where the Virtual Network Gateway is located.
	Location string `pulumi:"location"`
	// The user-defined name of the revoked certificate.
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configuration of the size and capacity of the Virtual Network Gateway.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Virtual Network Gateway.
	Type string `pulumi:"type"`
	// A `vpnClientConfiguration` block which is documented below.
	VpnClientConfigurations []GetVirtualNetworkGatewayVpnClientConfiguration `pulumi:"vpnClientConfigurations"`
	// The routing type of the Virtual Network Gateway.
	VpnType string `pulumi:"vpnType"`
}
