// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a VPN Site.
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
// 		exampleVirtualWan, err := network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVpnSite(ctx, "exampleVpnSite", &network.VpnSiteArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			Links: network.VpnSiteLinkArray{
// 				&network.VpnSiteLinkArgs{
// 					Name:      pulumi.String("link1"),
// 					IpAddress: pulumi.String("10.0.0.1"),
// 				},
// 			},
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
// VPN Sites can be imported using the `resource id`, e.g.
type VpnSite struct {
	pulumi.CustomResourceState

	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	AddressCidrs pulumi.StringArrayOutput `pulumi:"addressCidrs"`
	// The model of the VPN device.
	DeviceModel pulumi.StringPtrOutput `pulumi:"deviceModel"`
	// The name of the VPN device vendor.
	DeviceVendor pulumi.StringPtrOutput `pulumi:"deviceVendor"`
	// One or more `link` blocks as defined below.
	Links VpnSiteLinkArrayOutput `pulumi:"links"`
	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the VPN Site.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	VirtualWanId pulumi.StringOutput `pulumi:"virtualWanId"`
}

// NewVpnSite registers a new resource with the given unique name, arguments, and options.
func NewVpnSite(ctx *pulumi.Context,
	name string, args *VpnSiteArgs, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualWanId == nil {
		return nil, errors.New("missing required argument 'VirtualWanId'")
	}
	if args == nil {
		args = &VpnSiteArgs{}
	}
	var resource VpnSite
	err := ctx.RegisterResource("azure:network/vpnSite:VpnSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSite gets an existing VpnSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSiteState, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	var resource VpnSite
	err := ctx.ReadResource("azure:network/vpnSite:VpnSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSite resources.
type vpnSiteState struct {
	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	AddressCidrs []string `pulumi:"addressCidrs"`
	// The model of the VPN device.
	DeviceModel *string `pulumi:"deviceModel"`
	// The name of the VPN device vendor.
	DeviceVendor *string `pulumi:"deviceVendor"`
	// One or more `link` blocks as defined below.
	Links []VpnSiteLink `pulumi:"links"`
	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the VPN Site.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	VirtualWanId *string `pulumi:"virtualWanId"`
}

type VpnSiteState struct {
	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	AddressCidrs pulumi.StringArrayInput
	// The model of the VPN device.
	DeviceModel pulumi.StringPtrInput
	// The name of the VPN device vendor.
	DeviceVendor pulumi.StringPtrInput
	// One or more `link` blocks as defined below.
	Links VpnSiteLinkArrayInput
	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the VPN Site.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	VirtualWanId pulumi.StringPtrInput
}

func (VpnSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteState)(nil)).Elem()
}

type vpnSiteArgs struct {
	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	AddressCidrs []string `pulumi:"addressCidrs"`
	// The model of the VPN device.
	DeviceModel *string `pulumi:"deviceModel"`
	// The name of the VPN device vendor.
	DeviceVendor *string `pulumi:"deviceVendor"`
	// One or more `link` blocks as defined below.
	Links []VpnSiteLink `pulumi:"links"`
	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the VPN Site.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	VirtualWanId string `pulumi:"virtualWanId"`
}

// The set of arguments for constructing a VpnSite resource.
type VpnSiteArgs struct {
	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	AddressCidrs pulumi.StringArrayInput
	// The model of the VPN device.
	DeviceModel pulumi.StringPtrInput
	// The name of the VPN device vendor.
	DeviceVendor pulumi.StringPtrInput
	// One or more `link` blocks as defined below.
	Links VpnSiteLinkArrayInput
	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the VPN Site.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	VirtualWanId pulumi.StringInput
}

func (VpnSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteArgs)(nil)).Elem()
}
