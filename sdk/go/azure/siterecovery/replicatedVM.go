// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VM replicated using Azure Site Recovery (Azure to Azure only). A replicated VM keeps a copiously updated image of the VM in another region in order to be able to start the VM in that region in case of a disaster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/siterecovery"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		primaryVirtualNetwork, err := network.NewVirtualNetwork(ctx, "primaryVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: primaryResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("192.168.1.0/24"),
// 			},
// 			Location: primaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primarySubnet, err := network.NewSubnet(ctx, "primarySubnet", &network.SubnetArgs{
// 			ResourceGroupName:  primaryResourceGroup.Name,
// 			VirtualNetworkName: primaryVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("192.168.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryPublicIp, err := network.NewPublicIp(ctx, "primaryPublicIp", &network.PublicIpArgs{
// 			AllocationMethod:  pulumi.String("Static"),
// 			Location:          primaryResourceGroup.Location,
// 			ResourceGroupName: primaryResourceGroup.Name,
// 			Sku:               pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vmNetworkInterface, err := network.NewNetworkInterface(ctx, "vmNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          primaryResourceGroup.Location,
// 			ResourceGroupName: primaryResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("vm"),
// 					SubnetId:                   primarySubnet.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 					PublicIpAddressId:          primaryPublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vmVirtualMachine, err := compute.NewVirtualMachine(ctx, "vmVirtualMachine", &compute.VirtualMachineArgs{
// 			Location:          primaryResourceGroup.Location,
// 			ResourceGroupName: primaryResourceGroup.Name,
// 			VmSize:            pulumi.String("Standard_B1s"),
// 			NetworkInterfaceIds: pulumi.StringArray{
// 				vmNetworkInterface.ID(),
// 			},
// 			StorageImageReference: &compute.VirtualMachineStorageImageReferenceArgs{
// 				Publisher: pulumi.String("OpenLogic"),
// 				Offer:     pulumi.String("CentOS"),
// 				Sku:       pulumi.String("7.5"),
// 				Version:   pulumi.String("latest"),
// 			},
// 			StorageOsDisk: &compute.VirtualMachineStorageOsDiskArgs{
// 				Name:            pulumi.String("vm-os-disk"),
// 				OsType:          pulumi.String("Linux"),
// 				Caching:         pulumi.String("ReadWrite"),
// 				CreateOption:    pulumi.String("FromImage"),
// 				ManagedDiskType: pulumi.String("Premium_LRS"),
// 			},
// 			OsProfile: &compute.VirtualMachineOsProfileArgs{
// 				AdminUsername: pulumi.String("test-admin-123"),
// 				AdminPassword: pulumi.String("test-pwd-123"),
// 				ComputerName:  pulumi.String("vm"),
// 			},
// 			OsProfileLinuxConfig: &compute.VirtualMachineOsProfileLinuxConfigArgs{
// 				DisablePasswordAuthentication: pulumi.Bool(false),
// 			},
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
// 		primaryAccount, err := storage.NewAccount(ctx, "primaryAccount", &storage.AccountArgs{
// 			Location:               primaryResourceGroup.Location,
// 			ResourceGroupName:      primaryResourceGroup.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryVirtualNetwork, err := network.NewVirtualNetwork(ctx, "secondaryVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("192.168.2.0/24"),
// 			},
// 			Location: secondaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnet(ctx, "secondarySubnet", &network.SubnetArgs{
// 			ResourceGroupName:  secondaryResourceGroup.Name,
// 			VirtualNetworkName: secondaryVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("192.168.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryPublicIp, err := network.NewPublicIp(ctx, "secondaryPublicIp", &network.PublicIpArgs{
// 			AllocationMethod:  pulumi.String("Static"),
// 			Location:          secondaryResourceGroup.Location,
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			Sku:               pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = siterecovery.NewReplicatedVM(ctx, "vm_replication", &siterecovery.ReplicatedVMArgs{
// 			ResourceGroupName:                     secondaryResourceGroup.Name,
// 			RecoveryVaultName:                     vault.Name,
// 			SourceRecoveryFabricName:              primaryFabric.Name,
// 			SourceVmId:                            vmVirtualMachine.ID(),
// 			RecoveryReplicationPolicyId:           policy.ID(),
// 			SourceRecoveryProtectionContainerName: primaryProtectionContainer.Name,
// 			TargetResourceGroupId:                 secondaryResourceGroup.ID(),
// 			TargetRecoveryFabricId:                secondaryFabric.ID(),
// 			TargetRecoveryProtectionContainerId:   secondaryProtectionContainer.ID(),
// 			ManagedDisks: siterecovery.ReplicatedVMManagedDiskArray{
// 				&siterecovery.ReplicatedVMManagedDiskArgs{
// 					DiskId: vmVirtualMachine.StorageOsDisk.ApplyT(func(storageOsDisk compute.VirtualMachineStorageOsDisk) (string, error) {
// 						return storageOsDisk.ManagedDiskId, nil
// 					}).(pulumi.StringOutput),
// 					StagingStorageAccountId: primaryAccount.ID(),
// 					TargetResourceGroupId:   secondaryResourceGroup.ID(),
// 					TargetDiskType:          pulumi.String("Premium_LRS"),
// 					TargetReplicaDiskType:   pulumi.String("Premium_LRS"),
// 				},
// 			},
// 			NetworkInterfaces: siterecovery.ReplicatedVMNetworkInterfaceArray{
// 				&siterecovery.ReplicatedVMNetworkInterfaceArgs{
// 					SourceNetworkInterfaceId:  vmNetworkInterface.ID(),
// 					TargetSubnetName:          pulumi.String("network2-subnet"),
// 					RecoveryPublicIpAddressId: secondaryPublicIp.ID(),
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
// Site Recovery Replicated VM's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:siterecovery/replicatedVM:ReplicatedVM vmreplication /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name/replicationProtectedItems/vm-replication-name
// ```
type ReplicatedVM struct {
	pulumi.CustomResourceState

	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVMManagedDiskArrayOutput `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `networkInterface` block.
	NetworkInterfaces           ReplicatedVMNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	RecoveryReplicationPolicyId pulumi.StringOutput                     `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringOutput `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringOutput `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId pulumi.StringOutput `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringPtrOutput `pulumi:"targetAvailabilitySetId"`
	// Network to use when a failover is done (recommended to set if any networkInterface is configured for failover).
	TargetNetworkId pulumi.StringOutput `pulumi:"targetNetworkId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringOutput `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringOutput `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringOutput `pulumi:"targetResourceGroupId"`
}

// NewReplicatedVM registers a new resource with the given unique name, arguments, and options.
func NewReplicatedVM(ctx *pulumi.Context,
	name string, args *ReplicatedVMArgs, opts ...pulumi.ResourceOption) (*ReplicatedVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecoveryReplicationPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryReplicationPolicyId'")
	}
	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceRecoveryFabricName == nil {
		return nil, errors.New("invalid value for required argument 'SourceRecoveryFabricName'")
	}
	if args.SourceRecoveryProtectionContainerName == nil {
		return nil, errors.New("invalid value for required argument 'SourceRecoveryProtectionContainerName'")
	}
	if args.SourceVmId == nil {
		return nil, errors.New("invalid value for required argument 'SourceVmId'")
	}
	if args.TargetRecoveryFabricId == nil {
		return nil, errors.New("invalid value for required argument 'TargetRecoveryFabricId'")
	}
	if args.TargetRecoveryProtectionContainerId == nil {
		return nil, errors.New("invalid value for required argument 'TargetRecoveryProtectionContainerId'")
	}
	if args.TargetResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceGroupId'")
	}
	var resource ReplicatedVM
	err := ctx.RegisterResource("azure:siterecovery/replicatedVM:ReplicatedVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicatedVM gets an existing ReplicatedVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicatedVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicatedVMState, opts ...pulumi.ResourceOption) (*ReplicatedVM, error) {
	var resource ReplicatedVM
	err := ctx.ReadResource("azure:siterecovery/replicatedVM:ReplicatedVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicatedVM resources.
type replicatedVMState struct {
	// One or more `managedDisk` block.
	ManagedDisks []ReplicatedVMManagedDisk `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// One or more `networkInterface` block.
	NetworkInterfaces           []ReplicatedVMNetworkInterface `pulumi:"networkInterfaces"`
	RecoveryReplicationPolicyId *string                        `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName *string `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName *string `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId *string `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId *string `pulumi:"targetAvailabilitySetId"`
	// Network to use when a failover is done (recommended to set if any networkInterface is configured for failover).
	TargetNetworkId *string `pulumi:"targetNetworkId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId *string `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId *string `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId *string `pulumi:"targetResourceGroupId"`
}

type ReplicatedVMState struct {
	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVMManagedDiskArrayInput
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// One or more `networkInterface` block.
	NetworkInterfaces           ReplicatedVMNetworkInterfaceArrayInput
	RecoveryReplicationPolicyId pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringPtrInput
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringPtrInput
	// Id of the VM to replicate
	SourceVmId pulumi.StringPtrInput
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringPtrInput
	// Network to use when a failover is done (recommended to set if any networkInterface is configured for failover).
	TargetNetworkId pulumi.StringPtrInput
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringPtrInput
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringPtrInput
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringPtrInput
}

func (ReplicatedVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicatedVMState)(nil)).Elem()
}

type replicatedVMArgs struct {
	// One or more `managedDisk` block.
	ManagedDisks []ReplicatedVMManagedDisk `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// One or more `networkInterface` block.
	NetworkInterfaces           []ReplicatedVMNetworkInterface `pulumi:"networkInterfaces"`
	RecoveryReplicationPolicyId string                         `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName string `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName string `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId string `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId *string `pulumi:"targetAvailabilitySetId"`
	// Network to use when a failover is done (recommended to set if any networkInterface is configured for failover).
	TargetNetworkId *string `pulumi:"targetNetworkId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId string `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId string `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId string `pulumi:"targetResourceGroupId"`
}

// The set of arguments for constructing a ReplicatedVM resource.
type ReplicatedVMArgs struct {
	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVMManagedDiskArrayInput
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// One or more `networkInterface` block.
	NetworkInterfaces           ReplicatedVMNetworkInterfaceArrayInput
	RecoveryReplicationPolicyId pulumi.StringInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringInput
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringInput
	// Id of the VM to replicate
	SourceVmId pulumi.StringInput
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringPtrInput
	// Network to use when a failover is done (recommended to set if any networkInterface is configured for failover).
	TargetNetworkId pulumi.StringPtrInput
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringInput
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringInput
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringInput
}

func (ReplicatedVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicatedVMArgs)(nil)).Elem()
}

type ReplicatedVMInput interface {
	pulumi.Input

	ToReplicatedVMOutput() ReplicatedVMOutput
	ToReplicatedVMOutputWithContext(ctx context.Context) ReplicatedVMOutput
}

func (*ReplicatedVM) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicatedVM)(nil))
}

func (i *ReplicatedVM) ToReplicatedVMOutput() ReplicatedVMOutput {
	return i.ToReplicatedVMOutputWithContext(context.Background())
}

func (i *ReplicatedVM) ToReplicatedVMOutputWithContext(ctx context.Context) ReplicatedVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMOutput)
}

func (i *ReplicatedVM) ToReplicatedVMPtrOutput() ReplicatedVMPtrOutput {
	return i.ToReplicatedVMPtrOutputWithContext(context.Background())
}

func (i *ReplicatedVM) ToReplicatedVMPtrOutputWithContext(ctx context.Context) ReplicatedVMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMPtrOutput)
}

type ReplicatedVMPtrInput interface {
	pulumi.Input

	ToReplicatedVMPtrOutput() ReplicatedVMPtrOutput
	ToReplicatedVMPtrOutputWithContext(ctx context.Context) ReplicatedVMPtrOutput
}

type replicatedVMPtrType ReplicatedVMArgs

func (*replicatedVMPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicatedVM)(nil))
}

func (i *replicatedVMPtrType) ToReplicatedVMPtrOutput() ReplicatedVMPtrOutput {
	return i.ToReplicatedVMPtrOutputWithContext(context.Background())
}

func (i *replicatedVMPtrType) ToReplicatedVMPtrOutputWithContext(ctx context.Context) ReplicatedVMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMPtrOutput)
}

// ReplicatedVMArrayInput is an input type that accepts ReplicatedVMArray and ReplicatedVMArrayOutput values.
// You can construct a concrete instance of `ReplicatedVMArrayInput` via:
//
//          ReplicatedVMArray{ ReplicatedVMArgs{...} }
type ReplicatedVMArrayInput interface {
	pulumi.Input

	ToReplicatedVMArrayOutput() ReplicatedVMArrayOutput
	ToReplicatedVMArrayOutputWithContext(context.Context) ReplicatedVMArrayOutput
}

type ReplicatedVMArray []ReplicatedVMInput

func (ReplicatedVMArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ReplicatedVM)(nil))
}

func (i ReplicatedVMArray) ToReplicatedVMArrayOutput() ReplicatedVMArrayOutput {
	return i.ToReplicatedVMArrayOutputWithContext(context.Background())
}

func (i ReplicatedVMArray) ToReplicatedVMArrayOutputWithContext(ctx context.Context) ReplicatedVMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMArrayOutput)
}

// ReplicatedVMMapInput is an input type that accepts ReplicatedVMMap and ReplicatedVMMapOutput values.
// You can construct a concrete instance of `ReplicatedVMMapInput` via:
//
//          ReplicatedVMMap{ "key": ReplicatedVMArgs{...} }
type ReplicatedVMMapInput interface {
	pulumi.Input

	ToReplicatedVMMapOutput() ReplicatedVMMapOutput
	ToReplicatedVMMapOutputWithContext(context.Context) ReplicatedVMMapOutput
}

type ReplicatedVMMap map[string]ReplicatedVMInput

func (ReplicatedVMMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ReplicatedVM)(nil))
}

func (i ReplicatedVMMap) ToReplicatedVMMapOutput() ReplicatedVMMapOutput {
	return i.ToReplicatedVMMapOutputWithContext(context.Background())
}

func (i ReplicatedVMMap) ToReplicatedVMMapOutputWithContext(ctx context.Context) ReplicatedVMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMMapOutput)
}

type ReplicatedVMOutput struct {
	*pulumi.OutputState
}

func (ReplicatedVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicatedVM)(nil))
}

func (o ReplicatedVMOutput) ToReplicatedVMOutput() ReplicatedVMOutput {
	return o
}

func (o ReplicatedVMOutput) ToReplicatedVMOutputWithContext(ctx context.Context) ReplicatedVMOutput {
	return o
}

func (o ReplicatedVMOutput) ToReplicatedVMPtrOutput() ReplicatedVMPtrOutput {
	return o.ToReplicatedVMPtrOutputWithContext(context.Background())
}

func (o ReplicatedVMOutput) ToReplicatedVMPtrOutputWithContext(ctx context.Context) ReplicatedVMPtrOutput {
	return o.ApplyT(func(v ReplicatedVM) *ReplicatedVM {
		return &v
	}).(ReplicatedVMPtrOutput)
}

type ReplicatedVMPtrOutput struct {
	*pulumi.OutputState
}

func (ReplicatedVMPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicatedVM)(nil))
}

func (o ReplicatedVMPtrOutput) ToReplicatedVMPtrOutput() ReplicatedVMPtrOutput {
	return o
}

func (o ReplicatedVMPtrOutput) ToReplicatedVMPtrOutputWithContext(ctx context.Context) ReplicatedVMPtrOutput {
	return o
}

type ReplicatedVMArrayOutput struct{ *pulumi.OutputState }

func (ReplicatedVMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicatedVM)(nil))
}

func (o ReplicatedVMArrayOutput) ToReplicatedVMArrayOutput() ReplicatedVMArrayOutput {
	return o
}

func (o ReplicatedVMArrayOutput) ToReplicatedVMArrayOutputWithContext(ctx context.Context) ReplicatedVMArrayOutput {
	return o
}

func (o ReplicatedVMArrayOutput) Index(i pulumi.IntInput) ReplicatedVMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicatedVM {
		return vs[0].([]ReplicatedVM)[vs[1].(int)]
	}).(ReplicatedVMOutput)
}

type ReplicatedVMMapOutput struct{ *pulumi.OutputState }

func (ReplicatedVMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicatedVM)(nil))
}

func (o ReplicatedVMMapOutput) ToReplicatedVMMapOutput() ReplicatedVMMapOutput {
	return o
}

func (o ReplicatedVMMapOutput) ToReplicatedVMMapOutputWithContext(ctx context.Context) ReplicatedVMMapOutput {
	return o
}

func (o ReplicatedVMMapOutput) MapIndex(k pulumi.StringInput) ReplicatedVMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicatedVM {
		return vs[0].(map[string]ReplicatedVM)[vs[1].(string)]
	}).(ReplicatedVMOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicatedVMOutput{})
	pulumi.RegisterOutputType(ReplicatedVMPtrOutput{})
	pulumi.RegisterOutputType(ReplicatedVMArrayOutput{})
	pulumi.RegisterOutputType(ReplicatedVMMapOutput{})
}
