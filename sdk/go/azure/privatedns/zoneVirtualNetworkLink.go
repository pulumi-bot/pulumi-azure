// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage Private DNS zone Virtual Network Links. These Links enable DNS resolution and registration inside Azure Virtual Networks using Azure Private DNS.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatedns"
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
// 		exampleZone, err := privatedns.NewZone(ctx, "exampleZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleZoneVirtualNetworkLink, err := privatedns.NewZoneVirtualNetworkLink(ctx, "exampleZoneVirtualNetworkLink", &privatedns.ZoneVirtualNetworkLinkArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			PrivateDnsZoneName: exampleZone.Name,
// 			VirtualNetworkId:   pulumi.String(azurerm_virtual_network.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ZoneVirtualNetworkLink struct {
	pulumi.CustomResourceState

	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName pulumi.StringOutput `pulumi:"privateDnsZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled pulumi.BoolPtrOutput `pulumi:"registrationEnabled"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringOutput `pulumi:"virtualNetworkId"`
}

// NewZoneVirtualNetworkLink registers a new resource with the given unique name, arguments, and options.
func NewZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *ZoneVirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*ZoneVirtualNetworkLink, error) {
	if args == nil || args.PrivateDnsZoneName == nil {
		return nil, errors.New("missing required argument 'PrivateDnsZoneName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkId == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkId'")
	}
	if args == nil {
		args = &ZoneVirtualNetworkLinkArgs{}
	}
	var resource ZoneVirtualNetworkLink
	err := ctx.RegisterResource("azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneVirtualNetworkLink gets an existing ZoneVirtualNetworkLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneVirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*ZoneVirtualNetworkLink, error) {
	var resource ZoneVirtualNetworkLink
	err := ctx.ReadResource("azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneVirtualNetworkLink resources.
type zoneVirtualNetworkLinkState struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName *string `pulumi:"privateDnsZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled *bool `pulumi:"registrationEnabled"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
}

type ZoneVirtualNetworkLinkState struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName pulumi.StringPtrInput
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled pulumi.BoolPtrInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringPtrInput
}

func (ZoneVirtualNetworkLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneVirtualNetworkLinkState)(nil)).Elem()
}

type zoneVirtualNetworkLinkArgs struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName string `pulumi:"privateDnsZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled *bool `pulumi:"registrationEnabled"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId string `pulumi:"virtualNetworkId"`
}

// The set of arguments for constructing a ZoneVirtualNetworkLink resource.
type ZoneVirtualNetworkLinkArgs struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName pulumi.StringInput
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled pulumi.BoolPtrInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringInput
}

func (ZoneVirtualNetworkLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneVirtualNetworkLinkArgs)(nil)).Elem()
}
