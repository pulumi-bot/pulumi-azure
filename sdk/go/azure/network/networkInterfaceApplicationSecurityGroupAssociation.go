// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the association between a Network Interface and a Application Security Group.
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
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleApplicationSecurityGroup, err := network.NewApplicationSecurityGroup(ctx, "exampleApplicationSecurityGroup", &network.ApplicationSecurityGroupArgs{
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
// 		_, err = network.NewNetworkInterfaceApplicationSecurityGroupAssociation(ctx, "exampleNetworkInterfaceApplicationSecurityGroupAssociation", &network.NetworkInterfaceApplicationSecurityGroupAssociationArgs{
// 			NetworkInterfaceId:         exampleNetworkInterface.ID(),
// 			ApplicationSecurityGroupId: exampleApplicationSecurityGroup.ID(),
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
// Associations between Network Interfaces and Application Security Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkInterfaceApplicationSecurityGroupAssociation:NetworkInterfaceApplicationSecurityGroupAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1"
// ```
type NetworkInterfaceApplicationSecurityGroupAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	ApplicationSecurityGroupId pulumi.StringOutput `pulumi:"applicationSecurityGroupId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceApplicationSecurityGroupAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceApplicationSecurityGroupAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceApplicationSecurityGroupAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceApplicationSecurityGroupAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationSecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationSecurityGroupId'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	var resource NetworkInterfaceApplicationSecurityGroupAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceApplicationSecurityGroupAssociation:NetworkInterfaceApplicationSecurityGroupAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceApplicationSecurityGroupAssociation gets an existing NetworkInterfaceApplicationSecurityGroupAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceApplicationSecurityGroupAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceApplicationSecurityGroupAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceApplicationSecurityGroupAssociation, error) {
	var resource NetworkInterfaceApplicationSecurityGroupAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceApplicationSecurityGroupAssociation:NetworkInterfaceApplicationSecurityGroupAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceApplicationSecurityGroupAssociation resources.
type networkInterfaceApplicationSecurityGroupAssociationState struct {
	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	ApplicationSecurityGroupId *string `pulumi:"applicationSecurityGroupId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
}

type NetworkInterfaceApplicationSecurityGroupAssociationState struct {
	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	ApplicationSecurityGroupId pulumi.StringPtrInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
}

func (NetworkInterfaceApplicationSecurityGroupAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceApplicationSecurityGroupAssociationState)(nil)).Elem()
}

type networkInterfaceApplicationSecurityGroupAssociationArgs struct {
	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	ApplicationSecurityGroupId string `pulumi:"applicationSecurityGroupId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceApplicationSecurityGroupAssociation resource.
type NetworkInterfaceApplicationSecurityGroupAssociationArgs struct {
	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	ApplicationSecurityGroupId pulumi.StringInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceApplicationSecurityGroupAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceApplicationSecurityGroupAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceApplicationSecurityGroupAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceApplicationSecurityGroupAssociationOutput() NetworkInterfaceApplicationSecurityGroupAssociationOutput
	ToNetworkInterfaceApplicationSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceApplicationSecurityGroupAssociationOutput
}

func (*NetworkInterfaceApplicationSecurityGroupAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceApplicationSecurityGroupAssociation)(nil))
}

func (i *NetworkInterfaceApplicationSecurityGroupAssociation) ToNetworkInterfaceApplicationSecurityGroupAssociationOutput() NetworkInterfaceApplicationSecurityGroupAssociationOutput {
	return i.ToNetworkInterfaceApplicationSecurityGroupAssociationOutputWithContext(context.Background())
}

func (i *NetworkInterfaceApplicationSecurityGroupAssociation) ToNetworkInterfaceApplicationSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceApplicationSecurityGroupAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceApplicationSecurityGroupAssociationOutput)
}

func (i *NetworkInterfaceApplicationSecurityGroupAssociation) ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutput() NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput {
	return i.ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutputWithContext(context.Background())
}

func (i *NetworkInterfaceApplicationSecurityGroupAssociation) ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput)
}

type NetworkInterfaceApplicationSecurityGroupAssociationPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutput() NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput
	ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput
}

type NetworkInterfaceApplicationSecurityGroupAssociationOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceApplicationSecurityGroupAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceApplicationSecurityGroupAssociation)(nil))
}

func (o NetworkInterfaceApplicationSecurityGroupAssociationOutput) ToNetworkInterfaceApplicationSecurityGroupAssociationOutput() NetworkInterfaceApplicationSecurityGroupAssociationOutput {
	return o
}

func (o NetworkInterfaceApplicationSecurityGroupAssociationOutput) ToNetworkInterfaceApplicationSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceApplicationSecurityGroupAssociationOutput {
	return o
}

type NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceApplicationSecurityGroupAssociation)(nil))
}

func (o NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput) ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutput() NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput {
	return o
}

func (o NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput) ToNetworkInterfaceApplicationSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceApplicationSecurityGroupAssociationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceApplicationSecurityGroupAssociationPtrOutput{})
}
