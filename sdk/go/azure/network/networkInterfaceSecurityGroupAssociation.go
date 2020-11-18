// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the association between a Network Interface and a Network Security Group.
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
// 		exampleNetworkSecurityGroup, err := network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkInterfaceSecurityGroupAssociation(ctx, "exampleNetworkInterfaceSecurityGroupAssociation", &network.NetworkInterfaceSecurityGroupAssociationArgs{
// 			NetworkInterfaceId:     exampleNetworkInterface.ID(),
// 			NetworkSecurityGroupId: exampleNetworkSecurityGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkInterfaceSecurityGroupAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringOutput `pulumi:"networkSecurityGroupId"`
}

// NewNetworkInterfaceSecurityGroupAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceSecurityGroupAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceSecurityGroupAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.NetworkSecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupId'")
	}
	var resource NetworkInterfaceSecurityGroupAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceSecurityGroupAssociation gets an existing NetworkInterfaceSecurityGroupAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceSecurityGroupAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceSecurityGroupAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAssociation, error) {
	var resource NetworkInterfaceSecurityGroupAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAssociation resources.
type networkInterfaceSecurityGroupAssociationState struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId *string `pulumi:"networkSecurityGroupId"`
}

type NetworkInterfaceSecurityGroupAssociationState struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringPtrInput
}

func (NetworkInterfaceSecurityGroupAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAssociationState)(nil)).Elem()
}

type networkInterfaceSecurityGroupAssociationArgs struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
}

// The set of arguments for constructing a NetworkInterfaceSecurityGroupAssociation resource.
type NetworkInterfaceSecurityGroupAssociationArgs struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringInput
}

func (NetworkInterfaceSecurityGroupAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceSecurityGroupAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAssociationOutput() NetworkInterfaceSecurityGroupAssociationOutput
	ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationOutput
}

func (NetworkInterfaceSecurityGroupAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociation)(nil)).Elem()
}

func (i NetworkInterfaceSecurityGroupAssociation) ToNetworkInterfaceSecurityGroupAssociationOutput() NetworkInterfaceSecurityGroupAssociationOutput {
	return i.ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(context.Background())
}

func (i NetworkInterfaceSecurityGroupAssociation) ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAssociationOutput)
}

type NetworkInterfaceSecurityGroupAssociationOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceSecurityGroupAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociationOutput)(nil)).Elem()
}

func (o NetworkInterfaceSecurityGroupAssociationOutput) ToNetworkInterfaceSecurityGroupAssociationOutput() NetworkInterfaceSecurityGroupAssociationOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationOutput) ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAssociationOutput{})
}
