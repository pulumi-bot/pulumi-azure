// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Associates a Route Table with a Subnet within a Virtual Network.
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
// 		exampleRouteTable, err := network.NewRouteTable(ctx, "exampleRouteTable", &network.RouteTableArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Routes: network.RouteTableRouteArray{
// 				&network.RouteTableRouteArgs{
// 					Name:               pulumi.String("example"),
// 					AddressPrefix:      pulumi.String("10.100.0.0/14"),
// 					NextHopType:        pulumi.String("VirtualAppliance"),
// 					NextHopInIpAddress: pulumi.String("10.10.1.1"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnetRouteTableAssociation(ctx, "exampleSubnetRouteTableAssociation", &network.SubnetRouteTableAssociationArgs{
// 			SubnetId:     exampleSubnet.ID(),
// 			RouteTableId: exampleRouteTable.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SubnetRouteTableAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSubnetRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewSubnetRouteTableAssociation(ctx *pulumi.Context,
	name string, args *SubnetRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*SubnetRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource SubnetRouteTableAssociation
	err := ctx.RegisterResource("azure:network/subnetRouteTableAssociation:SubnetRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetRouteTableAssociation gets an existing SubnetRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetRouteTableAssociationState, opts ...pulumi.ResourceOption) (*SubnetRouteTableAssociation, error) {
	var resource SubnetRouteTableAssociation
	err := ctx.ReadResource("azure:network/subnetRouteTableAssociation:SubnetRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetRouteTableAssociation resources.
type subnetRouteTableAssociationState struct {
	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	RouteTableId *string `pulumi:"routeTableId"`
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
}

type SubnetRouteTableAssociationState struct {
	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	RouteTableId pulumi.StringPtrInput
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
}

func (SubnetRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetRouteTableAssociationState)(nil)).Elem()
}

type subnetRouteTableAssociationArgs struct {
	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	RouteTableId string `pulumi:"routeTableId"`
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SubnetRouteTableAssociation resource.
type SubnetRouteTableAssociationArgs struct {
	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	RouteTableId pulumi.StringInput
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
}

func (SubnetRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetRouteTableAssociationArgs)(nil)).Elem()
}

type SubnetRouteTableAssociationInput interface {
	pulumi.Input

	ToSubnetRouteTableAssociationOutput() SubnetRouteTableAssociationOutput
	ToSubnetRouteTableAssociationOutputWithContext(ctx context.Context) SubnetRouteTableAssociationOutput
}

func (SubnetRouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetRouteTableAssociation)(nil)).Elem()
}

func (i SubnetRouteTableAssociation) ToSubnetRouteTableAssociationOutput() SubnetRouteTableAssociationOutput {
	return i.ToSubnetRouteTableAssociationOutputWithContext(context.Background())
}

func (i SubnetRouteTableAssociation) ToSubnetRouteTableAssociationOutputWithContext(ctx context.Context) SubnetRouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetRouteTableAssociationOutput)
}

type SubnetRouteTableAssociationOutput struct {
	*pulumi.OutputState
}

func (SubnetRouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetRouteTableAssociationOutput)(nil)).Elem()
}

func (o SubnetRouteTableAssociationOutput) ToSubnetRouteTableAssociationOutput() SubnetRouteTableAssociationOutput {
	return o
}

func (o SubnetRouteTableAssociationOutput) ToSubnetRouteTableAssociationOutputWithContext(ctx context.Context) SubnetRouteTableAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetRouteTableAssociationOutput{})
}
