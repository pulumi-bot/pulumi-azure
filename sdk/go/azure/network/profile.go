// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Network Profile.
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.1.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.1.0.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("delegation"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.ContainerInstance/containerGroups"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewProfile(ctx, "exampleProfile", &network.ProfileArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ContainerNetworkInterface: &network.ProfileContainerNetworkInterfaceArgs{
// 				Name: pulumi.String("examplecnic"),
// 				IpConfigurations: network.ProfileContainerNetworkInterfaceIpConfigurationArray{
// 					&network.ProfileContainerNetworkInterfaceIpConfigurationArgs{
// 						Name:     pulumi.String("exampleipconfig"),
// 						SubnetId: exampleSubnet.ID(),
// 					},
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
// Network Profile can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/profile:Profile example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkProfiles/examplenetprofile
// ```
type Profile struct {
	pulumi.CustomResourceState

	// A `containerNetworkInterface` block as documented below.
	ContainerNetworkInterface ProfileContainerNetworkInterfaceOutput `pulumi:"containerNetworkInterface"`
	// A list of Container Network Interface ID's.
	ContainerNetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"containerNetworkInterfaceIds"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Network Profile. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerNetworkInterface == nil {
		return nil, errors.New("invalid value for required argument 'ContainerNetworkInterface'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Profile
	err := ctx.RegisterResource("azure:network/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure:network/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// A `containerNetworkInterface` block as documented below.
	ContainerNetworkInterface *ProfileContainerNetworkInterface `pulumi:"containerNetworkInterface"`
	// A list of Container Network Interface ID's.
	ContainerNetworkInterfaceIds []string `pulumi:"containerNetworkInterfaceIds"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Network Profile. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ProfileState struct {
	// A `containerNetworkInterface` block as documented below.
	ContainerNetworkInterface ProfileContainerNetworkInterfacePtrInput
	// A list of Container Network Interface ID's.
	ContainerNetworkInterfaceIds pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Network Profile. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// A `containerNetworkInterface` block as documented below.
	ContainerNetworkInterface ProfileContainerNetworkInterface `pulumi:"containerNetworkInterface"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Network Profile. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// A `containerNetworkInterface` block as documented below.
	ContainerNetworkInterface ProfileContainerNetworkInterfaceInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Network Profile. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((*Profile)(nil))
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

func (i *Profile) ToProfilePtrOutput() ProfilePtrOutput {
	return i.ToProfilePtrOutputWithContext(context.Background())
}

func (i *Profile) ToProfilePtrOutputWithContext(ctx context.Context) ProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfilePtrOutput)
}

type ProfilePtrInput interface {
	pulumi.Input

	ToProfilePtrOutput() ProfilePtrOutput
	ToProfilePtrOutputWithContext(ctx context.Context) ProfilePtrOutput
}

type ProfileOutput struct {
	*pulumi.OutputState
}

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Profile)(nil))
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

type ProfilePtrOutput struct {
	*pulumi.OutputState
}

func (ProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil))
}

func (o ProfilePtrOutput) ToProfilePtrOutput() ProfilePtrOutput {
	return o
}

func (o ProfilePtrOutput) ToProfilePtrOutputWithContext(ctx context.Context) ProfilePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfilePtrOutput{})
}
