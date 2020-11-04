// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure recovery vault protection container mapping. A protection container mapping decides how to translate the protection container when a VM is migrated from one region to another.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/siterecovery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primaryResourceGroup, err := core.NewResourceGroup(ctx, "primaryResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryResourceGroup, err := core.NewResourceGroup(ctx, "secondaryResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vault, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          secondaryResourceGroup.Location,
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryFabric, err := siterecovery.NewFabric(ctx, "primaryFabric", &siterecovery.FabricArgs{
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			RecoveryVaultName: vault.Name,
// 			Location:          primaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryFabric, err := siterecovery.NewFabric(ctx, "secondaryFabric", &siterecovery.FabricArgs{
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			RecoveryVaultName: vault.Name,
// 			Location:          secondaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryProtectionContainer, err := siterecovery.NewProtectionContainer(ctx, "primaryProtectionContainer", &siterecovery.ProtectionContainerArgs{
// 			ResourceGroupName:  secondaryResourceGroup.Name,
// 			RecoveryVaultName:  vault.Name,
// 			RecoveryFabricName: primaryFabric.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryProtectionContainer, err := siterecovery.NewProtectionContainer(ctx, "secondaryProtectionContainer", &siterecovery.ProtectionContainerArgs{
// 			ResourceGroupName:  secondaryResourceGroup.Name,
// 			RecoveryVaultName:  vault.Name,
// 			RecoveryFabricName: secondaryFabric.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		policy, err := siterecovery.NewReplicationPolicy(ctx, "policy", &siterecovery.ReplicationPolicyArgs{
// 			ResourceGroupName:                               secondaryResourceGroup.Name,
// 			RecoveryVaultName:                               vault.Name,
// 			RecoveryPointRetentionInMinutes:                 pulumi.Int(24 * 60),
// 			ApplicationConsistentSnapshotFrequencyInMinutes: pulumi.Int(4 * 60),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = siterecovery.NewProtectionContainerMapping(ctx, "container_mapping", &siterecovery.ProtectionContainerMappingArgs{
// 			ResourceGroupName:                     secondaryResourceGroup.Name,
// 			RecoveryVaultName:                     vault.Name,
// 			RecoveryFabricName:                    primaryFabric.Name,
// 			RecoverySourceProtectionContainerName: primaryProtectionContainer.Name,
// 			RecoveryTargetProtectionContainerId:   secondaryProtectionContainer.ID(),
// 			RecoveryReplicationPolicyId:           policy.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProtectionContainerMapping struct {
	pulumi.CustomResourceState

	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of fabric that should contains the protection container to map.
	RecoveryFabricName pulumi.StringOutput `pulumi:"recoveryFabricName"`
	// Id of the policy to use for this mapping.
	RecoveryReplicationPolicyId pulumi.StringOutput `pulumi:"recoveryReplicationPolicyId"`
	// Name of the source protection container to map.
	RecoverySourceProtectionContainerName pulumi.StringOutput `pulumi:"recoverySourceProtectionContainerName"`
	// Id of target protection container to map to.
	RecoveryTargetProtectionContainerId pulumi.StringOutput `pulumi:"recoveryTargetProtectionContainerId"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProtectionContainerMapping registers a new resource with the given unique name, arguments, and options.
func NewProtectionContainerMapping(ctx *pulumi.Context,
	name string, args *ProtectionContainerMappingArgs, opts ...pulumi.ResourceOption) (*ProtectionContainerMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.RecoveryFabricName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryFabricName'")
	}
	if args.RecoveryReplicationPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryReplicationPolicyId'")
	}
	if args.RecoverySourceProtectionContainerName == nil {
		return nil, errors.New("invalid value for required argument 'RecoverySourceProtectionContainerName'")
	}
	if args.RecoveryTargetProtectionContainerId == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryTargetProtectionContainerId'")
	}
	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ProtectionContainerMapping
	err := ctx.RegisterResource("azure:siterecovery/protectionContainerMapping:ProtectionContainerMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectionContainerMapping gets an existing ProtectionContainerMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectionContainerMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionContainerMappingState, opts ...pulumi.ResourceOption) (*ProtectionContainerMapping, error) {
	var resource ProtectionContainerMapping
	err := ctx.ReadResource("azure:siterecovery/protectionContainerMapping:ProtectionContainerMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectionContainerMapping resources.
type protectionContainerMappingState struct {
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// Name of fabric that should contains the protection container to map.
	RecoveryFabricName *string `pulumi:"recoveryFabricName"`
	// Id of the policy to use for this mapping.
	RecoveryReplicationPolicyId *string `pulumi:"recoveryReplicationPolicyId"`
	// Name of the source protection container to map.
	RecoverySourceProtectionContainerName *string `pulumi:"recoverySourceProtectionContainerName"`
	// Id of target protection container to map to.
	RecoveryTargetProtectionContainerId *string `pulumi:"recoveryTargetProtectionContainerId"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ProtectionContainerMappingState struct {
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// Name of fabric that should contains the protection container to map.
	RecoveryFabricName pulumi.StringPtrInput
	// Id of the policy to use for this mapping.
	RecoveryReplicationPolicyId pulumi.StringPtrInput
	// Name of the source protection container to map.
	RecoverySourceProtectionContainerName pulumi.StringPtrInput
	// Id of target protection container to map to.
	RecoveryTargetProtectionContainerId pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
}

func (ProtectionContainerMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerMappingState)(nil)).Elem()
}

type protectionContainerMappingArgs struct {
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// Name of fabric that should contains the protection container to map.
	RecoveryFabricName string `pulumi:"recoveryFabricName"`
	// Id of the policy to use for this mapping.
	RecoveryReplicationPolicyId string `pulumi:"recoveryReplicationPolicyId"`
	// Name of the source protection container to map.
	RecoverySourceProtectionContainerName string `pulumi:"recoverySourceProtectionContainerName"`
	// Id of target protection container to map to.
	RecoveryTargetProtectionContainerId string `pulumi:"recoveryTargetProtectionContainerId"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ProtectionContainerMapping resource.
type ProtectionContainerMappingArgs struct {
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// Name of fabric that should contains the protection container to map.
	RecoveryFabricName pulumi.StringInput
	// Id of the policy to use for this mapping.
	RecoveryReplicationPolicyId pulumi.StringInput
	// Name of the source protection container to map.
	RecoverySourceProtectionContainerName pulumi.StringInput
	// Id of target protection container to map to.
	RecoveryTargetProtectionContainerId pulumi.StringInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
}

func (ProtectionContainerMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerMappingArgs)(nil)).Elem()
}
