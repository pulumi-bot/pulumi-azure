// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Subnet within a Virtual Network.
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
// 		example, err := network.LookupSubnet(ctx, &network.LookupSubnetArgs{
// 			Name:               "backend",
// 			VirtualNetworkName: "production",
// 			ResourceGroupName:  "networking",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("subnetId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("azure:network/getSubnet:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetArgs struct {
	// Specifies the name of the Subnet.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Virtual Network is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the Virtual Network this Subnet is located within.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// A collection of values returned by getSubnet.
type LookupSubnetResult struct {
	// (Deprecated) The address prefix used for the subnet.
	AddressPrefix string `pulumi:"addressPrefix"`
	// The address prefixes for the subnet.
	AddressPrefixes []string `pulumi:"addressPrefixes"`
	// Enable or Disable network policies for the private link endpoint on the subnet.
	EnforcePrivateLinkEndpointNetworkPolicies bool `pulumi:"enforcePrivateLinkEndpointNetworkPolicies"`
	// Enable or Disable network policies for the private link service on the subnet.
	EnforcePrivateLinkServiceNetworkPolicies bool `pulumi:"enforcePrivateLinkServiceNetworkPolicies"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The ID of the Network Security Group associated with the subnet.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	// The ID of the Route Table associated with this subnet.
	RouteTableId string `pulumi:"routeTableId"`
	// A list of Service Endpoints within this subnet.
	ServiceEndpoints   []string `pulumi:"serviceEndpoints"`
	VirtualNetworkName string   `pulumi:"virtualNetworkName"`
}

func LookupSubnetApply(ctx *pulumi.Context, args LookupSubnetApplyInput, opts ...pulumi.InvokeOption) LookupSubnetResultOutput {
	return args.ToLookupSubnetApplyOutput().ApplyT(func(v LookupSubnetArgs) (LookupSubnetResult, error) {
		r, err := LookupSubnet(ctx, &v, opts...)
		return *r, err

	}).(LookupSubnetResultOutput)
}

// LookupSubnetApplyInput is an input type that accepts LookupSubnetApplyArgs and LookupSubnetApplyOutput values.
// You can construct a concrete instance of `LookupSubnetApplyInput` via:
//
//          LookupSubnetApplyArgs{...}
type LookupSubnetApplyInput interface {
	pulumi.Input

	ToLookupSubnetApplyOutput() LookupSubnetApplyOutput
	ToLookupSubnetApplyOutputWithContext(context.Context) LookupSubnetApplyOutput
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetApplyArgs struct {
	// Specifies the name of the Subnet.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group the Virtual Network is located in.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Specifies the name of the Virtual Network this Subnet is located within.
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (LookupSubnetApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}

func (i LookupSubnetApplyArgs) ToLookupSubnetApplyOutput() LookupSubnetApplyOutput {
	return i.ToLookupSubnetApplyOutputWithContext(context.Background())
}

func (i LookupSubnetApplyArgs) ToLookupSubnetApplyOutputWithContext(ctx context.Context) LookupSubnetApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupSubnetApplyOutput)
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetApplyOutput struct{ *pulumi.OutputState }

func (LookupSubnetApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}

func (o LookupSubnetApplyOutput) ToLookupSubnetApplyOutput() LookupSubnetApplyOutput {
	return o
}

func (o LookupSubnetApplyOutput) ToLookupSubnetApplyOutputWithContext(ctx context.Context) LookupSubnetApplyOutput {
	return o
}

// Specifies the name of the Subnet.
func (o LookupSubnetApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the name of the resource group the Virtual Network is located in.
func (o LookupSubnetApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// Specifies the name of the Virtual Network this Subnet is located within.
func (o LookupSubnetApplyOutput) VirtualNetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetArgs) string { return v.VirtualNetworkName }).(pulumi.StringOutput)
}

// A collection of values returned by getSubnet.
type LookupSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetResult)(nil)).Elem()
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutput() LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutputWithContext(ctx context.Context) LookupSubnetResultOutput {
	return o
}

// (Deprecated) The address prefix used for the subnet.
func (o LookupSubnetResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

// The address prefixes for the subnet.
func (o LookupSubnetResultOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

// Enable or Disable network policies for the private link endpoint on the subnet.
func (o LookupSubnetResultOutput) EnforcePrivateLinkEndpointNetworkPolicies() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetResult) bool { return v.EnforcePrivateLinkEndpointNetworkPolicies }).(pulumi.BoolOutput)
}

// Enable or Disable network policies for the private link service on the subnet.
func (o LookupSubnetResultOutput) EnforcePrivateLinkServiceNetworkPolicies() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetResult) bool { return v.EnforcePrivateLinkServiceNetworkPolicies }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSubnetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Network Security Group associated with the subnet.
func (o LookupSubnetResultOutput) NetworkSecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.NetworkSecurityGroupId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The ID of the Route Table associated with this subnet.
func (o LookupSubnetResultOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.RouteTableId }).(pulumi.StringOutput)
}

// A list of Service Endpoints within this subnet.
func (o LookupSubnetResultOutput) ServiceEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.ServiceEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupSubnetResultOutput) VirtualNetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.VirtualNetworkName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetApplyOutput{})
	pulumi.RegisterOutputType(LookupSubnetResultOutput{})
}
