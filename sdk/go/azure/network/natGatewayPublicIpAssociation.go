// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the association between a Nat Gateway and a Public IP.
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
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNatGateway, err := network.NewNatGateway(ctx, "exampleNatGateway", &network.NatGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNatGatewayPublicIpAssociation(ctx, "exampleNatGatewayPublicIpAssociation", &network.NatGatewayPublicIpAssociationArgs{
// 			NatGatewayId:      exampleNatGateway.ID(),
// 			PublicIpAddressId: examplePublicIp.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NatGatewayPublicIpAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// The ID of the Public IP which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpAddressId pulumi.StringOutput `pulumi:"publicIpAddressId"`
}

// NewNatGatewayPublicIpAssociation registers a new resource with the given unique name, arguments, and options.
func NewNatGatewayPublicIpAssociation(ctx *pulumi.Context,
	name string, args *NatGatewayPublicIpAssociationArgs, opts ...pulumi.ResourceOption) (*NatGatewayPublicIpAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.NatGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'NatGatewayId'")
	}
	if args.PublicIpAddressId == nil {
		return nil, errors.New("invalid value for required argument 'PublicIpAddressId'")
	}
	var resource NatGatewayPublicIpAssociation
	err := ctx.RegisterResource("azure:network/natGatewayPublicIpAssociation:NatGatewayPublicIpAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGatewayPublicIpAssociation gets an existing NatGatewayPublicIpAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGatewayPublicIpAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayPublicIpAssociationState, opts ...pulumi.ResourceOption) (*NatGatewayPublicIpAssociation, error) {
	var resource NatGatewayPublicIpAssociation
	err := ctx.ReadResource("azure:network/natGatewayPublicIpAssociation:NatGatewayPublicIpAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGatewayPublicIpAssociation resources.
type natGatewayPublicIpAssociationState struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// The ID of the Public IP which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpAddressId *string `pulumi:"publicIpAddressId"`
}

type NatGatewayPublicIpAssociationState struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId pulumi.StringPtrInput
	// The ID of the Public IP which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpAddressId pulumi.StringPtrInput
}

func (NatGatewayPublicIpAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayPublicIpAssociationState)(nil)).Elem()
}

type natGatewayPublicIpAssociationArgs struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId string `pulumi:"natGatewayId"`
	// The ID of the Public IP which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpAddressId string `pulumi:"publicIpAddressId"`
}

// The set of arguments for constructing a NatGatewayPublicIpAssociation resource.
type NatGatewayPublicIpAssociationArgs struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId pulumi.StringInput
	// The ID of the Public IP which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpAddressId pulumi.StringInput
}

func (NatGatewayPublicIpAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayPublicIpAssociationArgs)(nil)).Elem()
}
