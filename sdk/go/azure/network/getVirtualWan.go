// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Virtual Wan.
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
// 		example, err := network.LookupVirtualWan(ctx, &network.LookupVirtualWanArgs{
// 			Name:              "existing",
// 			ResourceGroupName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		ctx.Export("allowBranchToBranchTraffic", example.AllowBranchToBranchTraffic)
// 		ctx.Export("disableVpnEncryption", data.Azurerm_virtual_wan.Exemple.Disable_vpn_encryption)
// 		ctx.Export("location", data.Azurerm_virtual_wan.Exemple.Location)
// 		ctx.Export("office365LocalBreakoutCategory", data.Azurerm_virtual_wan.Exemple.Office365_local_breakout_category)
// 		ctx.Export("sku", data.Azurerm_virtual_wan.Exemple.Sku)
// 		ctx.Export("tags", data.Azurerm_virtual_wan.Exemple.Tags)
// 		ctx.Export("virtualHubs", data.Azurerm_virtual_wan.Exemple.Virtual_hubs)
// 		ctx.Export("vpnSites", data.Azurerm_virtual_wan.Exemple.Vpn_sites)
// 		return nil
// 	})
// }
// ```
func LookupVirtualWan(ctx *pulumi.Context, args *LookupVirtualWanArgs, opts ...pulumi.InvokeOption) (*LookupVirtualWanResult, error) {
	var rv LookupVirtualWanResult
	err := ctx.Invoke("azure:network/getVirtualWan:getVirtualWan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualWan.
type LookupVirtualWanArgs struct {
	// The name of this Virtual Wan.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Virtual Wan exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getVirtualWan.
type LookupVirtualWanResult struct {
	// Is branch to branch traffic is allowed?
	AllowBranchToBranchTraffic bool `pulumi:"allowBranchToBranchTraffic"`
	// Is VPN Encryption disabled?
	DisableVpnEncryption bool `pulumi:"disableVpnEncryption"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the Virtual Wan exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The Office365 Local Breakout Category.
	Office365LocalBreakoutCategory string `pulumi:"office365LocalBreakoutCategory"`
	ResourceGroupName              string `pulumi:"resourceGroupName"`
	// Type of Virtual Wan (Basic or Standard).
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the Virtual Wan.
	Tags map[string]string `pulumi:"tags"`
	// A list of Virtual Hubs ID's attached to this Virtual WAN.
	VirtualHubIds []string `pulumi:"virtualHubIds"`
	// A list of VPN Site ID's attached to this Virtual WAN.
	VpnSiteIds []string `pulumi:"vpnSiteIds"`
}

func LookupVirtualWanApply(ctx *pulumi.Context, args LookupVirtualWanApplyInput, opts ...pulumi.InvokeOption) LookupVirtualWanResultOutput {
	return args.ToLookupVirtualWanApplyOutput().ApplyT(func(v LookupVirtualWanArgs) (LookupVirtualWanResult, error) {
		r, err := LookupVirtualWan(ctx, &v, opts...)
		return *r, err

	}).(LookupVirtualWanResultOutput)
}

// LookupVirtualWanApplyInput is an input type that accepts LookupVirtualWanApplyArgs and LookupVirtualWanApplyOutput values.
// You can construct a concrete instance of `LookupVirtualWanApplyInput` via:
//
//          LookupVirtualWanApplyArgs{...}
type LookupVirtualWanApplyInput interface {
	pulumi.Input

	ToLookupVirtualWanApplyOutput() LookupVirtualWanApplyOutput
	ToLookupVirtualWanApplyOutputWithContext(context.Context) LookupVirtualWanApplyOutput
}

// A collection of arguments for invoking getVirtualWan.
type LookupVirtualWanApplyArgs struct {
	// The name of this Virtual Wan.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Virtual Wan exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVirtualWanApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanArgs)(nil)).Elem()
}

func (i LookupVirtualWanApplyArgs) ToLookupVirtualWanApplyOutput() LookupVirtualWanApplyOutput {
	return i.ToLookupVirtualWanApplyOutputWithContext(context.Background())
}

func (i LookupVirtualWanApplyArgs) ToLookupVirtualWanApplyOutputWithContext(ctx context.Context) LookupVirtualWanApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupVirtualWanApplyOutput)
}

// A collection of arguments for invoking getVirtualWan.
type LookupVirtualWanApplyOutput struct{ *pulumi.OutputState }

func (LookupVirtualWanApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanArgs)(nil)).Elem()
}

func (o LookupVirtualWanApplyOutput) ToLookupVirtualWanApplyOutput() LookupVirtualWanApplyOutput {
	return o
}

func (o LookupVirtualWanApplyOutput) ToLookupVirtualWanApplyOutputWithContext(ctx context.Context) LookupVirtualWanApplyOutput {
	return o
}

// The name of this Virtual Wan.
func (o LookupVirtualWanApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the Resource Group where the Virtual Wan exists.
func (o LookupVirtualWanApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getVirtualWan.
type LookupVirtualWanResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualWanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanResult)(nil)).Elem()
}

func (o LookupVirtualWanResultOutput) ToLookupVirtualWanResultOutput() LookupVirtualWanResultOutput {
	return o
}

func (o LookupVirtualWanResultOutput) ToLookupVirtualWanResultOutputWithContext(ctx context.Context) LookupVirtualWanResultOutput {
	return o
}

// Is branch to branch traffic is allowed?
func (o LookupVirtualWanResultOutput) AllowBranchToBranchTraffic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) bool { return v.AllowBranchToBranchTraffic }).(pulumi.BoolOutput)
}

// Is VPN Encryption disabled?
func (o LookupVirtualWanResultOutput) DisableVpnEncryption() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) bool { return v.DisableVpnEncryption }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualWanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Region where the Virtual Wan exists.
func (o LookupVirtualWanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Office365 Local Breakout Category.
func (o LookupVirtualWanResultOutput) Office365LocalBreakoutCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Office365LocalBreakoutCategory }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// Type of Virtual Wan (Basic or Standard).
func (o LookupVirtualWanResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Sku }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the Virtual Wan.
func (o LookupVirtualWanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A list of Virtual Hubs ID's attached to this Virtual WAN.
func (o LookupVirtualWanResultOutput) VirtualHubIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) []string { return v.VirtualHubIds }).(pulumi.StringArrayOutput)
}

// A list of VPN Site ID's attached to this Virtual WAN.
func (o LookupVirtualWanResultOutput) VpnSiteIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) []string { return v.VpnSiteIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualWanApplyOutput{})
	pulumi.RegisterOutputType(LookupVirtualWanResultOutput{})
}
