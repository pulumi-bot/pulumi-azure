// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Public IP Address.
//
// ## Example Usage
// ### Reference An Existing)
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
// 		example, err := network.GetPublicIP(ctx, &network.GetPublicIPArgs{
// 			Name:              "name_of_public_ip",
// 			ResourceGroupName: "name_of_resource_group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("domainNameLabel", example.DomainNameLabel)
// 		ctx.Export("publicIpAddress", example.IpAddress)
// 		return nil
// 	})
// }
// ```
// ### Retrieve The Dynamic Public IP Of A New VM)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
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
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:             exampleResourceGroup.Location,
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			AllocationMethod:     pulumi.String("Dynamic"),
// 			IdleTimeoutInMinutes: pulumi.Int(30),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkInterface, err := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("testconfiguration1"),
// 					SubnetId:                   exampleSubnet.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Static"),
// 					PrivateIpAddress:           pulumi.String("10.0.2.5"),
// 					PublicIpAddressId:          examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualMachine, err := compute.NewVirtualMachine(ctx, "exampleVirtualMachine", &compute.VirtualMachineArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NetworkInterfaceIds: pulumi.StringArray{
// 				exampleNetworkInterface.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("publicIpAddress", examplePublicIp.IpAddress)
// 		return nil
// 	})
// }
// ```
func GetPublicIP(ctx *pulumi.Context, args *GetPublicIPArgs, opts ...pulumi.InvokeOption) (*GetPublicIPResult, error) {
	var rv GetPublicIPResult
	err := ctx.Invoke("azure:network/getPublicIP:getPublicIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicIP.
type GetPublicIPArgs struct {
	// Specifies the name of the public IP address.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assigned to the resource.
	Tags  map[string]string `pulumi:"tags"`
	Zones []string          `pulumi:"zones"`
}

// A collection of values returned by getPublicIP.
type GetPublicIPResult struct {
	AllocationMethod string `pulumi:"allocationMethod"`
	// The label for the Domain Name.
	DomainNameLabel string `pulumi:"domainNameLabel"`
	// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
	Fqdn string `pulumi:"fqdn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies the timeout for the TCP idle connection.
	IdleTimeoutInMinutes int `pulumi:"idleTimeoutInMinutes"`
	// The IP address value that was allocated.
	IpAddress string `pulumi:"ipAddress"`
	// A mapping of tags to assigned to the resource.
	IpTags map[string]string `pulumi:"ipTags"`
	// The IP version being used, for example `IPv4` or `IPv6`.
	IpVersion         string `pulumi:"ipVersion"`
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ReverseFqdn       string `pulumi:"reverseFqdn"`
	// The SKU of the Public IP.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assigned to the resource.
	Tags  map[string]string `pulumi:"tags"`
	Zones []string          `pulumi:"zones"`
}

func GetPublicIPOutput(ctx *pulumi.Context, args GetPublicIPOutputArgs, opts ...pulumi.InvokeOption) GetPublicIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPublicIPResult, error) {
			args := v.(GetPublicIPArgs)
			r, err := GetPublicIP(ctx, &args, opts...)
			return *r, err
		}).(GetPublicIPResultOutput)
}

// A collection of arguments for invoking getPublicIP.
type GetPublicIPOutputArgs struct {
	// Specifies the name of the public IP address.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assigned to the resource.
	Tags  pulumi.StringMapInput   `pulumi:"tags"`
	Zones pulumi.StringArrayInput `pulumi:"zones"`
}

func (GetPublicIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIPArgs)(nil)).Elem()
}

// A collection of values returned by getPublicIP.
type GetPublicIPResultOutput struct{ *pulumi.OutputState }

func (GetPublicIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIPResult)(nil)).Elem()
}

func (o GetPublicIPResultOutput) ToGetPublicIPResultOutput() GetPublicIPResultOutput {
	return o
}

func (o GetPublicIPResultOutput) ToGetPublicIPResultOutputWithContext(ctx context.Context) GetPublicIPResultOutput {
	return o
}

func (o GetPublicIPResultOutput) AllocationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.AllocationMethod }).(pulumi.StringOutput)
}

// The label for the Domain Name.
func (o GetPublicIPResultOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.DomainNameLabel }).(pulumi.StringOutput)
}

// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
func (o GetPublicIPResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPublicIPResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the timeout for the TCP idle connection.
func (o GetPublicIPResultOutput) IdleTimeoutInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v GetPublicIPResult) int { return v.IdleTimeoutInMinutes }).(pulumi.IntOutput)
}

// The IP address value that was allocated.
func (o GetPublicIPResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// A mapping of tags to assigned to the resource.
func (o GetPublicIPResultOutput) IpTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetPublicIPResult) map[string]string { return v.IpTags }).(pulumi.StringMapOutput)
}

// The IP version being used, for example `IPv4` or `IPv6`.
func (o GetPublicIPResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

func (o GetPublicIPResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetPublicIPResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetPublicIPResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o GetPublicIPResultOutput) ReverseFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.ReverseFqdn }).(pulumi.StringOutput)
}

// The SKU of the Public IP.
func (o GetPublicIPResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIPResult) string { return v.Sku }).(pulumi.StringOutput)
}

// A mapping of tags to assigned to the resource.
func (o GetPublicIPResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetPublicIPResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetPublicIPResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPublicIPResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicIPResultOutput{})
}
