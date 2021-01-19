// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the association between a Network Interface and a Load Balancer's Backend Address Pool.
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
// 		exampleBackendAddressPool, err := lb.NewBackendAddressPool(ctx, "exampleBackendAddressPool", &lb.BackendAddressPoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			LoadbalancerId:    exampleLoadBalancer.ID(),
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
// 		_, err = network.NewNetworkInterfaceBackendAddressPoolAssociation(ctx, "exampleNetworkInterfaceBackendAddressPoolAssociation", &network.NetworkInterfaceBackendAddressPoolAssociationArgs{
// 			NetworkInterfaceId:   exampleNetworkInterface.ID(),
// 			IpConfigurationName:  pulumi.String("testconfiguration1"),
// 			BackendAddressPoolId: exampleBackendAddressPool.ID(),
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
// Associations between Network Interfaces and Load Balancer Backend Address Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1"
// ```
type NetworkInterfaceBackendAddressPoolAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringOutput `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringOutput `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceBackendAddressPoolAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceBackendAddressPoolAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceBackendAddressPoolAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceBackendAddressPoolAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendAddressPoolId == nil {
		return nil, errors.New("invalid value for required argument 'BackendAddressPoolId'")
	}
	if args.IpConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'IpConfigurationName'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	var resource NetworkInterfaceBackendAddressPoolAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceBackendAddressPoolAssociation gets an existing NetworkInterfaceBackendAddressPoolAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceBackendAddressPoolAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceBackendAddressPoolAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceBackendAddressPoolAssociation, error) {
	var resource NetworkInterfaceBackendAddressPoolAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceBackendAddressPoolAssociation resources.
type networkInterfaceBackendAddressPoolAssociationState struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId *string `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName *string `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
}

type NetworkInterfaceBackendAddressPoolAssociationState struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringPtrInput
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringPtrInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
}

func (NetworkInterfaceBackendAddressPoolAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceBackendAddressPoolAssociationState)(nil)).Elem()
}

type networkInterfaceBackendAddressPoolAssociationArgs struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId string `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName string `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceBackendAddressPoolAssociation resource.
type NetworkInterfaceBackendAddressPoolAssociationArgs struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringInput
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceBackendAddressPoolAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceBackendAddressPoolAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceBackendAddressPoolAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceBackendAddressPoolAssociationOutput() NetworkInterfaceBackendAddressPoolAssociationOutput
	ToNetworkInterfaceBackendAddressPoolAssociationOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationOutput
}

func (*NetworkInterfaceBackendAddressPoolAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (i *NetworkInterfaceBackendAddressPoolAssociation) ToNetworkInterfaceBackendAddressPoolAssociationOutput() NetworkInterfaceBackendAddressPoolAssociationOutput {
	return i.ToNetworkInterfaceBackendAddressPoolAssociationOutputWithContext(context.Background())
}

func (i *NetworkInterfaceBackendAddressPoolAssociation) ToNetworkInterfaceBackendAddressPoolAssociationOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceBackendAddressPoolAssociationOutput)
}

func (i *NetworkInterfaceBackendAddressPoolAssociation) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutput() NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return i.ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(context.Background())
}

func (i *NetworkInterfaceBackendAddressPoolAssociation) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceBackendAddressPoolAssociationPtrOutput)
}

type NetworkInterfaceBackendAddressPoolAssociationPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceBackendAddressPoolAssociationPtrOutput() NetworkInterfaceBackendAddressPoolAssociationPtrOutput
	ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationPtrOutput
}

type networkInterfaceBackendAddressPoolAssociationPtrType NetworkInterfaceBackendAddressPoolAssociationArgs

func (*networkInterfaceBackendAddressPoolAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (i *networkInterfaceBackendAddressPoolAssociationPtrType) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutput() NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return i.ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(context.Background())
}

func (i *networkInterfaceBackendAddressPoolAssociationPtrType) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceBackendAddressPoolAssociationPtrOutput)
}

// NetworkInterfaceBackendAddressPoolAssociationArrayInput is an input type that accepts NetworkInterfaceBackendAddressPoolAssociationArray and NetworkInterfaceBackendAddressPoolAssociationArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceBackendAddressPoolAssociationArrayInput` via:
//
//          NetworkInterfaceBackendAddressPoolAssociationArray{ NetworkInterfaceBackendAddressPoolAssociationArgs{...} }
type NetworkInterfaceBackendAddressPoolAssociationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceBackendAddressPoolAssociationArrayOutput() NetworkInterfaceBackendAddressPoolAssociationArrayOutput
	ToNetworkInterfaceBackendAddressPoolAssociationArrayOutputWithContext(context.Context) NetworkInterfaceBackendAddressPoolAssociationArrayOutput
}

type NetworkInterfaceBackendAddressPoolAssociationArray []NetworkInterfaceBackendAddressPoolAssociationInput

func (NetworkInterfaceBackendAddressPoolAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (i NetworkInterfaceBackendAddressPoolAssociationArray) ToNetworkInterfaceBackendAddressPoolAssociationArrayOutput() NetworkInterfaceBackendAddressPoolAssociationArrayOutput {
	return i.ToNetworkInterfaceBackendAddressPoolAssociationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceBackendAddressPoolAssociationArray) ToNetworkInterfaceBackendAddressPoolAssociationArrayOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceBackendAddressPoolAssociationArrayOutput)
}

// NetworkInterfaceBackendAddressPoolAssociationMapInput is an input type that accepts NetworkInterfaceBackendAddressPoolAssociationMap and NetworkInterfaceBackendAddressPoolAssociationMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceBackendAddressPoolAssociationMapInput` via:
//
//          NetworkInterfaceBackendAddressPoolAssociationMap{ "key": NetworkInterfaceBackendAddressPoolAssociationArgs{...} }
type NetworkInterfaceBackendAddressPoolAssociationMapInput interface {
	pulumi.Input

	ToNetworkInterfaceBackendAddressPoolAssociationMapOutput() NetworkInterfaceBackendAddressPoolAssociationMapOutput
	ToNetworkInterfaceBackendAddressPoolAssociationMapOutputWithContext(context.Context) NetworkInterfaceBackendAddressPoolAssociationMapOutput
}

type NetworkInterfaceBackendAddressPoolAssociationMap map[string]NetworkInterfaceBackendAddressPoolAssociationInput

func (NetworkInterfaceBackendAddressPoolAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (i NetworkInterfaceBackendAddressPoolAssociationMap) ToNetworkInterfaceBackendAddressPoolAssociationMapOutput() NetworkInterfaceBackendAddressPoolAssociationMapOutput {
	return i.ToNetworkInterfaceBackendAddressPoolAssociationMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceBackendAddressPoolAssociationMap) ToNetworkInterfaceBackendAddressPoolAssociationMapOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceBackendAddressPoolAssociationMapOutput)
}

type NetworkInterfaceBackendAddressPoolAssociationOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceBackendAddressPoolAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (o NetworkInterfaceBackendAddressPoolAssociationOutput) ToNetworkInterfaceBackendAddressPoolAssociationOutput() NetworkInterfaceBackendAddressPoolAssociationOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationOutput) ToNetworkInterfaceBackendAddressPoolAssociationOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationOutput) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutput() NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return o.ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceBackendAddressPoolAssociationOutput) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceBackendAddressPoolAssociation) *NetworkInterfaceBackendAddressPoolAssociation {
		return &v
	}).(NetworkInterfaceBackendAddressPoolAssociationPtrOutput)
}

type NetworkInterfaceBackendAddressPoolAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceBackendAddressPoolAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (o NetworkInterfaceBackendAddressPoolAssociationPtrOutput) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutput() NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationPtrOutput) ToNetworkInterfaceBackendAddressPoolAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationPtrOutput {
	return o
}

type NetworkInterfaceBackendAddressPoolAssociationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceBackendAddressPoolAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (o NetworkInterfaceBackendAddressPoolAssociationArrayOutput) ToNetworkInterfaceBackendAddressPoolAssociationArrayOutput() NetworkInterfaceBackendAddressPoolAssociationArrayOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationArrayOutput) ToNetworkInterfaceBackendAddressPoolAssociationArrayOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationArrayOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceBackendAddressPoolAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceBackendAddressPoolAssociation {
		return vs[0].([]NetworkInterfaceBackendAddressPoolAssociation)[vs[1].(int)]
	}).(NetworkInterfaceBackendAddressPoolAssociationOutput)
}

type NetworkInterfaceBackendAddressPoolAssociationMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceBackendAddressPoolAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkInterfaceBackendAddressPoolAssociation)(nil))
}

func (o NetworkInterfaceBackendAddressPoolAssociationMapOutput) ToNetworkInterfaceBackendAddressPoolAssociationMapOutput() NetworkInterfaceBackendAddressPoolAssociationMapOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationMapOutput) ToNetworkInterfaceBackendAddressPoolAssociationMapOutputWithContext(ctx context.Context) NetworkInterfaceBackendAddressPoolAssociationMapOutput {
	return o
}

func (o NetworkInterfaceBackendAddressPoolAssociationMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceBackendAddressPoolAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkInterfaceBackendAddressPoolAssociation {
		return vs[0].(map[string]NetworkInterfaceBackendAddressPoolAssociation)[vs[1].(string)]
	}).(NetworkInterfaceBackendAddressPoolAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceBackendAddressPoolAssociationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceBackendAddressPoolAssociationPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceBackendAddressPoolAssociationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceBackendAddressPoolAssociationMapOutput{})
}
