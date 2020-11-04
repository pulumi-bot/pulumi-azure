// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a NetApp Snapshot.
//
// ## NetApp Snapshot Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/netapp"
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
// 		_, err = network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("netapp"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.Netapp/volumes"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/networkinterfaces/*"),
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/join/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := netapp.NewAccount(ctx, "exampleAccount", &netapp.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePool, err := netapp.NewPool(ctx, "examplePool", &netapp.PoolArgs{
// 			AccountName:       exampleAccount.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServiceLevel:      pulumi.String("Premium"),
// 			SizeInTb:          pulumi.Int(4),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVolume, err := netapp.NewVolume(ctx, "exampleVolume", &netapp.VolumeArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AccountName:       exampleAccount.Name,
// 			PoolName:          examplePool.Name,
// 			VolumePath:        pulumi.String("my-unique-file-path"),
// 			ServiceLevel:      pulumi.String("Premium"),
// 			SubnetId:          pulumi.Any(azurerm_subnet.Test.Id),
// 			StorageQuotaInGb:  pulumi.Int(100),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = netapp.NewSnapshot(ctx, "exampleSnapshot", &netapp.SnapshotArgs{
// 			AccountName:       exampleAccount.Name,
// 			PoolName:          examplePool.Name,
// 			VolumeName:        exampleVolume.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName pulumi.StringOutput `pulumi:"poolName"`
	// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	VolumeName pulumi.StringOutput `pulumi:"volumeName"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	var resource Snapshot
	err := ctx.RegisterResource("azure:netapp/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure:netapp/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName *string `pulumi:"poolName"`
	// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	VolumeName *string `pulumi:"volumeName"`
}

type SnapshotState struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName pulumi.StringPtrInput
	// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	VolumeName pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName string `pulumi:"poolName"`
	// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	VolumeName string `pulumi:"volumeName"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName pulumi.StringInput
	// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
	VolumeName pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}
