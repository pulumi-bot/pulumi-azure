// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
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
//
// ```sh
//  $ pulumi import azure:network/vpnSite:VpnSite example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnSites/site1
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualWanId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualWanId'")
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

type VpnSiteInput interface {
	pulumi.Input

	ToVpnSiteOutput() VpnSiteOutput
	ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput
}

func (*VpnSite) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSite)(nil))
}

func (i *VpnSite) ToVpnSiteOutput() VpnSiteOutput {
	return i.ToVpnSiteOutputWithContext(context.Background())
}

func (i *VpnSite) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteOutput)
}

func (i *VpnSite) ToVpnSitePtrOutput() VpnSitePtrOutput {
	return i.ToVpnSitePtrOutputWithContext(context.Background())
}

func (i *VpnSite) ToVpnSitePtrOutputWithContext(ctx context.Context) VpnSitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSitePtrOutput)
}

type VpnSitePtrInput interface {
	pulumi.Input

	ToVpnSitePtrOutput() VpnSitePtrOutput
	ToVpnSitePtrOutputWithContext(ctx context.Context) VpnSitePtrOutput
}

type vpnSitePtrType VpnSiteArgs

func (*vpnSitePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSite)(nil))
}

func (i *vpnSitePtrType) ToVpnSitePtrOutput() VpnSitePtrOutput {
	return i.ToVpnSitePtrOutputWithContext(context.Background())
}

func (i *vpnSitePtrType) ToVpnSitePtrOutputWithContext(ctx context.Context) VpnSitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteOutput).ToVpnSitePtrOutput()
}

// VpnSiteArrayInput is an input type that accepts VpnSiteArray and VpnSiteArrayOutput values.
// You can construct a concrete instance of `VpnSiteArrayInput` via:
//
//          VpnSiteArray{ VpnSiteArgs{...} }
type VpnSiteArrayInput interface {
	pulumi.Input

	ToVpnSiteArrayOutput() VpnSiteArrayOutput
	ToVpnSiteArrayOutputWithContext(context.Context) VpnSiteArrayOutput
}

type VpnSiteArray []VpnSiteInput

func (VpnSiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSite)(nil))
}

func (i VpnSiteArray) ToVpnSiteArrayOutput() VpnSiteArrayOutput {
	return i.ToVpnSiteArrayOutputWithContext(context.Background())
}

func (i VpnSiteArray) ToVpnSiteArrayOutputWithContext(ctx context.Context) VpnSiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteArrayOutput)
}

// VpnSiteMapInput is an input type that accepts VpnSiteMap and VpnSiteMapOutput values.
// You can construct a concrete instance of `VpnSiteMapInput` via:
//
//          VpnSiteMap{ "key": VpnSiteArgs{...} }
type VpnSiteMapInput interface {
	pulumi.Input

	ToVpnSiteMapOutput() VpnSiteMapOutput
	ToVpnSiteMapOutputWithContext(context.Context) VpnSiteMapOutput
}

type VpnSiteMap map[string]VpnSiteInput

func (VpnSiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpnSite)(nil))
}

func (i VpnSiteMap) ToVpnSiteMapOutput() VpnSiteMapOutput {
	return i.ToVpnSiteMapOutputWithContext(context.Background())
}

func (i VpnSiteMap) ToVpnSiteMapOutputWithContext(ctx context.Context) VpnSiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteMapOutput)
}

type VpnSiteOutput struct {
	*pulumi.OutputState
}

func (VpnSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSite)(nil))
}

func (o VpnSiteOutput) ToVpnSiteOutput() VpnSiteOutput {
	return o
}

func (o VpnSiteOutput) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return o
}

func (o VpnSiteOutput) ToVpnSitePtrOutput() VpnSitePtrOutput {
	return o.ToVpnSitePtrOutputWithContext(context.Background())
}

func (o VpnSiteOutput) ToVpnSitePtrOutputWithContext(ctx context.Context) VpnSitePtrOutput {
	return o.ApplyT(func(v VpnSite) *VpnSite {
		return &v
	}).(VpnSitePtrOutput)
}

type VpnSitePtrOutput struct {
	*pulumi.OutputState
}

func (VpnSitePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSite)(nil))
}

func (o VpnSitePtrOutput) ToVpnSitePtrOutput() VpnSitePtrOutput {
	return o
}

func (o VpnSitePtrOutput) ToVpnSitePtrOutputWithContext(ctx context.Context) VpnSitePtrOutput {
	return o
}

type VpnSiteArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSite)(nil))
}

func (o VpnSiteArrayOutput) ToVpnSiteArrayOutput() VpnSiteArrayOutput {
	return o
}

func (o VpnSiteArrayOutput) ToVpnSiteArrayOutputWithContext(ctx context.Context) VpnSiteArrayOutput {
	return o
}

func (o VpnSiteArrayOutput) Index(i pulumi.IntInput) VpnSiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSite {
		return vs[0].([]VpnSite)[vs[1].(int)]
	}).(VpnSiteOutput)
}

type VpnSiteMapOutput struct{ *pulumi.OutputState }

func (VpnSiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpnSite)(nil))
}

func (o VpnSiteMapOutput) ToVpnSiteMapOutput() VpnSiteMapOutput {
	return o
}

func (o VpnSiteMapOutput) ToVpnSiteMapOutputWithContext(ctx context.Context) VpnSiteMapOutput {
	return o
}

func (o VpnSiteMapOutput) MapIndex(k pulumi.StringInput) VpnSiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpnSite {
		return vs[0].(map[string]VpnSite)[vs[1].(string)]
	}).(VpnSiteOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnSiteOutput{})
	pulumi.RegisterOutputType(VpnSitePtrOutput{})
	pulumi.RegisterOutputType(VpnSiteArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteMapOutput{})
}
