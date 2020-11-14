// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the association between a Network Interface and a Load Balancer's NAT Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/lb"
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
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FrontendIpConfigurations: lb.LoadBalancerFrontendIpConfigurationArray{
// 				&lb.LoadBalancerFrontendIpConfigurationArgs{
// 					Name:              pulumi.String("primary"),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNatRule, err := lb.NewNatRule(ctx, "exampleNatRule", &lb.NatRuleArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			LoadbalancerId:              exampleLoadBalancer.ID(),
// 			Protocol:                    pulumi.String("Tcp"),
// 			FrontendPort:                pulumi.Int(3389),
// 			BackendPort:                 pulumi.Int(3389),
// 			FrontendIpConfigurationName: pulumi.String("primary"),
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
// 		_, err = network.NewNetworkInterfaceNatRuleAssociation(ctx, "exampleNetworkInterfaceNatRuleAssociation", &network.NetworkInterfaceNatRuleAssociationArgs{
// 			NetworkInterfaceId:  exampleNetworkInterface.ID(),
// 			IpConfigurationName: pulumi.String("testconfiguration1"),
// 			NatRuleId:           exampleNatRule.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkInterfaceNatRuleAssociation struct {
	pulumi.CustomResourceState

	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringOutput `pulumi:"ipConfigurationName"`
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId pulumi.StringOutput `pulumi:"natRuleId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceNatRuleAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceNatRuleAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceNatRuleAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceNatRuleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'IpConfigurationName'")
	}
	if args.NatRuleId == nil {
		return nil, errors.New("invalid value for required argument 'NatRuleId'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	var resource NetworkInterfaceNatRuleAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceNatRuleAssociation gets an existing NetworkInterfaceNatRuleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceNatRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceNatRuleAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceNatRuleAssociation, error) {
	var resource NetworkInterfaceNatRuleAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceNatRuleAssociation resources.
type networkInterfaceNatRuleAssociationState struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName *string `pulumi:"ipConfigurationName"`
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId *string `pulumi:"natRuleId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
}

type NetworkInterfaceNatRuleAssociationState struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringPtrInput
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId pulumi.StringPtrInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
}

func (NetworkInterfaceNatRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceNatRuleAssociationState)(nil)).Elem()
}

type networkInterfaceNatRuleAssociationArgs struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName string `pulumi:"ipConfigurationName"`
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId string `pulumi:"natRuleId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceNatRuleAssociation resource.
type NetworkInterfaceNatRuleAssociationArgs struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringInput
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId pulumi.StringInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceNatRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceNatRuleAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceNatRuleAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceNatRuleAssociationOutput() NetworkInterfaceNatRuleAssociationOutput
	ToNetworkInterfaceNatRuleAssociationOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationOutput
}

func (NetworkInterfaceNatRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceNatRuleAssociation)(nil)).Elem()
}

func (i NetworkInterfaceNatRuleAssociation) ToNetworkInterfaceNatRuleAssociationOutput() NetworkInterfaceNatRuleAssociationOutput {
	return i.ToNetworkInterfaceNatRuleAssociationOutputWithContext(context.Background())
}

func (i NetworkInterfaceNatRuleAssociation) ToNetworkInterfaceNatRuleAssociationOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceNatRuleAssociationOutput)
}

type NetworkInterfaceNatRuleAssociationOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceNatRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceNatRuleAssociationOutput)(nil)).Elem()
}

func (o NetworkInterfaceNatRuleAssociationOutput) ToNetworkInterfaceNatRuleAssociationOutput() NetworkInterfaceNatRuleAssociationOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationOutput) ToNetworkInterfaceNatRuleAssociationOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceNatRuleAssociationOutput{})
}
