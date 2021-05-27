// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Virtual Network Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := network.LookupVirtualNetworkGateway(ctx, &network.LookupVirtualNetworkGatewayArgs{
// 			Name:              "production",
// 			ResourceGroupName: "networking",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("virtualNetworkGatewayId", example.Id)
// 		return nil
// 	})
// }
// ```
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
	ActiveActive bool                                  `pulumi:"activeActive"`
	BgpSettings  []GetVirtualNetworkGatewayBgpSetting  `pulumi:"bgpSettings"`
	CustomRoutes []GetVirtualNetworkGatewayCustomRoute `pulumi:"customRoutes"`
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
	Name string `pulumi:"name"`
	// Whether a private IP will be used for this  gateway for connections.
	PrivateIpAddressEnabled bool   `pulumi:"privateIpAddressEnabled"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
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

func LookupVirtualNetworkGatewayOutput(ctx *pulumi.Context, args LookupVirtualNetworkGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkGatewayResult, error) {
			args := v.(LookupVirtualNetworkGatewayArgs)
			r, err := LookupVirtualNetworkGateway(ctx, &args, opts...)
			return *r, err
		}).(LookupVirtualNetworkGatewayResultOutput)
}

// A collection of arguments for invoking getVirtualNetworkGateway.
type LookupVirtualNetworkGatewayOutputArgs struct {
	// Specifies the name of the Virtual Network Gateway.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group the Virtual Network Gateway is located in.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVirtualNetworkGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualNetworkGateway.
type LookupVirtualNetworkGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayResult)(nil)).Elem()
}

func (o LookupVirtualNetworkGatewayResultOutput) ToLookupVirtualNetworkGatewayResultOutput() LookupVirtualNetworkGatewayResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayResultOutput) ToLookupVirtualNetworkGatewayResultOutputWithContext(ctx context.Context) LookupVirtualNetworkGatewayResultOutput {
	return o
}

// Is this an Active-Active Gateway?
func (o LookupVirtualNetworkGatewayResultOutput) ActiveActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) bool { return v.ActiveActive }).(pulumi.BoolOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) BgpSettings() GetVirtualNetworkGatewayBgpSettingArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) []GetVirtualNetworkGatewayBgpSetting { return v.BgpSettings }).(GetVirtualNetworkGatewayBgpSettingArrayOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) CustomRoutes() GetVirtualNetworkGatewayCustomRouteArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) []GetVirtualNetworkGatewayCustomRoute { return v.CustomRoutes }).(GetVirtualNetworkGatewayCustomRouteArrayOutput)
}

// The ID of the local network gateway
// through which outbound Internet traffic from the virtual network in which the
// gateway is created will be routed (*forced tunneling*). Refer to the
// [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
func (o LookupVirtualNetworkGatewayResultOutput) DefaultLocalNetworkGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.DefaultLocalNetworkGatewayId }).(pulumi.StringOutput)
}

// Will BGP (Border Gateway Protocol) will be enabled
// for this Virtual Network Gateway.
func (o LookupVirtualNetworkGatewayResultOutput) EnableBgp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) bool { return v.EnableBgp }).(pulumi.BoolOutput)
}

// The Generation of the Virtual Network Gateway.
func (o LookupVirtualNetworkGatewayResultOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Generation }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualNetworkGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// One or two `ipConfiguration` blocks documented below.
func (o LookupVirtualNetworkGatewayResultOutput) IpConfigurations() GetVirtualNetworkGatewayIpConfigurationArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) []GetVirtualNetworkGatewayIpConfiguration {
		return v.IpConfigurations
	}).(GetVirtualNetworkGatewayIpConfigurationArrayOutput)
}

// The location/region where the Virtual Network Gateway is located.
func (o LookupVirtualNetworkGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// The user-defined name of the revoked certificate.
func (o LookupVirtualNetworkGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether a private IP will be used for this  gateway for connections.
func (o LookupVirtualNetworkGatewayResultOutput) PrivateIpAddressEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) bool { return v.PrivateIpAddressEnabled }).(pulumi.BoolOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// Configuration of the size and capacity of the Virtual Network Gateway.
func (o LookupVirtualNetworkGatewayResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Sku }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the resource.
func (o LookupVirtualNetworkGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the Virtual Network Gateway.
func (o LookupVirtualNetworkGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// A `vpnClientConfiguration` block which is documented below.
func (o LookupVirtualNetworkGatewayResultOutput) VpnClientConfigurations() GetVirtualNetworkGatewayVpnClientConfigurationArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) []GetVirtualNetworkGatewayVpnClientConfiguration {
		return v.VpnClientConfigurations
	}).(GetVirtualNetworkGatewayVpnClientConfigurationArrayOutput)
}

// The routing type of the Virtual Network Gateway.
func (o LookupVirtualNetworkGatewayResultOutput) VpnType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.VpnType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkGatewayResultOutput{})
}
