# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ReplicatedVMArgs', 'ReplicatedVM']

@pulumi.input_type
class ReplicatedVMArgs:
    def __init__(__self__, *,
                 recovery_replication_policy_id: pulumi.Input[str],
                 recovery_vault_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 source_recovery_fabric_name: pulumi.Input[str],
                 source_recovery_protection_container_name: pulumi.Input[str],
                 source_vm_id: pulumi.Input[str],
                 target_recovery_fabric_id: pulumi.Input[str],
                 target_recovery_protection_container_id: pulumi.Input[str],
                 target_resource_group_id: pulumi.Input[str],
                 managed_disks: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatedVMManagedDiskArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatedVMNetworkInterfaceArgs']]]] = None,
                 target_availability_set_id: Optional[pulumi.Input[str]] = None,
                 target_network_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReplicatedVM resource.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_recovery_fabric_name: Name of fabric that should contains this replication.
        :param pulumi.Input[str] source_recovery_protection_container_name: Name of the protection container to use.
        :param pulumi.Input[str] source_vm_id: Id of the VM to replicate
        :param pulumi.Input[str] target_recovery_fabric_id: Id of fabric where the VM replication should be handled when a failover is done.
        :param pulumi.Input[str] target_recovery_protection_container_id: Id of protection container where the VM replication should be created when a failover is done.
        :param pulumi.Input[str] target_resource_group_id: Id of resource group where the VM should be created when a failover is done.
        :param pulumi.Input[Sequence[pulumi.Input['ReplicatedVMManagedDiskArgs']]] managed_disks: One or more `managed_disk` block.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[Sequence[pulumi.Input['ReplicatedVMNetworkInterfaceArgs']]] network_interfaces: One or more `network_interface` block.
        :param pulumi.Input[str] target_availability_set_id: Id of availability set that the new VM should belong to when a failover is done.
        :param pulumi.Input[str] target_network_id: Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        """
        pulumi.set(__self__, "recovery_replication_policy_id", recovery_replication_policy_id)
        pulumi.set(__self__, "recovery_vault_name", recovery_vault_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "source_recovery_fabric_name", source_recovery_fabric_name)
        pulumi.set(__self__, "source_recovery_protection_container_name", source_recovery_protection_container_name)
        pulumi.set(__self__, "source_vm_id", source_vm_id)
        pulumi.set(__self__, "target_recovery_fabric_id", target_recovery_fabric_id)
        pulumi.set(__self__, "target_recovery_protection_container_id", target_recovery_protection_container_id)
        pulumi.set(__self__, "target_resource_group_id", target_resource_group_id)
        if managed_disks is not None:
            pulumi.set(__self__, "managed_disks", managed_disks)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_interfaces is not None:
            pulumi.set(__self__, "network_interfaces", network_interfaces)
        if target_availability_set_id is not None:
            pulumi.set(__self__, "target_availability_set_id", target_availability_set_id)
        if target_network_id is not None:
            pulumi.set(__self__, "target_network_id", target_network_id)

    @property
    @pulumi.getter(name="recoveryReplicationPolicyId")
    def recovery_replication_policy_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "recovery_replication_policy_id")

    @recovery_replication_policy_id.setter
    def recovery_replication_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "recovery_replication_policy_id", value)

    @property
    @pulumi.getter(name="recoveryVaultName")
    def recovery_vault_name(self) -> pulumi.Input[str]:
        """
        The name of the vault that should be updated.
        """
        return pulumi.get(self, "recovery_vault_name")

    @recovery_vault_name.setter
    def recovery_vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "recovery_vault_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Name of the resource group where the vault that should be updated is located.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="sourceRecoveryFabricName")
    def source_recovery_fabric_name(self) -> pulumi.Input[str]:
        """
        Name of fabric that should contains this replication.
        """
        return pulumi.get(self, "source_recovery_fabric_name")

    @source_recovery_fabric_name.setter
    def source_recovery_fabric_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_recovery_fabric_name", value)

    @property
    @pulumi.getter(name="sourceRecoveryProtectionContainerName")
    def source_recovery_protection_container_name(self) -> pulumi.Input[str]:
        """
        Name of the protection container to use.
        """
        return pulumi.get(self, "source_recovery_protection_container_name")

    @source_recovery_protection_container_name.setter
    def source_recovery_protection_container_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_recovery_protection_container_name", value)

    @property
    @pulumi.getter(name="sourceVmId")
    def source_vm_id(self) -> pulumi.Input[str]:
        """
        Id of the VM to replicate
        """
        return pulumi.get(self, "source_vm_id")

    @source_vm_id.setter
    def source_vm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_vm_id", value)

    @property
    @pulumi.getter(name="targetRecoveryFabricId")
    def target_recovery_fabric_id(self) -> pulumi.Input[str]:
        """
        Id of fabric where the VM replication should be handled when a failover is done.
        """
        return pulumi.get(self, "target_recovery_fabric_id")

    @target_recovery_fabric_id.setter
    def target_recovery_fabric_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_recovery_fabric_id", value)

    @property
    @pulumi.getter(name="targetRecoveryProtectionContainerId")
    def target_recovery_protection_container_id(self) -> pulumi.Input[str]:
        """
        Id of protection container where the VM replication should be created when a failover is done.
        """
        return pulumi.get(self, "target_recovery_protection_container_id")

    @target_recovery_protection_container_id.setter
    def target_recovery_protection_container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_recovery_protection_container_id", value)

    @property
    @pulumi.getter(name="targetResourceGroupId")
    def target_resource_group_id(self) -> pulumi.Input[str]:
        """
        Id of resource group where the VM should be created when a failover is done.
        """
        return pulumi.get(self, "target_resource_group_id")

    @target_resource_group_id.setter
    def target_resource_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_group_id", value)

    @property
    @pulumi.getter(name="managedDisks")
    def managed_disks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatedVMManagedDiskArgs']]]]:
        """
        One or more `managed_disk` block.
        """
        return pulumi.get(self, "managed_disks")

    @managed_disks.setter
    def managed_disks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatedVMManagedDiskArgs']]]]):
        pulumi.set(self, "managed_disks", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the network mapping.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatedVMNetworkInterfaceArgs']]]]:
        """
        One or more `network_interface` block.
        """
        return pulumi.get(self, "network_interfaces")

    @network_interfaces.setter
    def network_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatedVMNetworkInterfaceArgs']]]]):
        pulumi.set(self, "network_interfaces", value)

    @property
    @pulumi.getter(name="targetAvailabilitySetId")
    def target_availability_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of availability set that the new VM should belong to when a failover is done.
        """
        return pulumi.get(self, "target_availability_set_id")

    @target_availability_set_id.setter
    def target_availability_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_availability_set_id", value)

    @property
    @pulumi.getter(name="targetNetworkId")
    def target_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        """
        return pulumi.get(self, "target_network_id")

    @target_network_id.setter
    def target_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_network_id", value)


class ReplicatedVM(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 managed_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMManagedDiskArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMNetworkInterfaceArgs']]]]] = None,
                 recovery_replication_policy_id: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_recovery_fabric_name: Optional[pulumi.Input[str]] = None,
                 source_recovery_protection_container_name: Optional[pulumi.Input[str]] = None,
                 source_vm_id: Optional[pulumi.Input[str]] = None,
                 target_availability_set_id: Optional[pulumi.Input[str]] = None,
                 target_network_id: Optional[pulumi.Input[str]] = None,
                 target_recovery_fabric_id: Optional[pulumi.Input[str]] = None,
                 target_recovery_protection_container_id: Optional[pulumi.Input[str]] = None,
                 target_resource_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a VM replicated using Azure Site Recovery (Azure to Azure only). A replicated VM keeps a copiously updated image of the VM in another region in order to be able to start the VM in that region in case of a disaster.

        ## Import

        Site Recovery Replicated VM's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:siterecovery/replicatedVM:ReplicatedVM vmreplication /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name/replicationProtectedItems/vm-replication-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMManagedDiskArgs']]]] managed_disks: One or more `managed_disk` block.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMNetworkInterfaceArgs']]]] network_interfaces: One or more `network_interface` block.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_recovery_fabric_name: Name of fabric that should contains this replication.
        :param pulumi.Input[str] source_recovery_protection_container_name: Name of the protection container to use.
        :param pulumi.Input[str] source_vm_id: Id of the VM to replicate
        :param pulumi.Input[str] target_availability_set_id: Id of availability set that the new VM should belong to when a failover is done.
        :param pulumi.Input[str] target_network_id: Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        :param pulumi.Input[str] target_recovery_fabric_id: Id of fabric where the VM replication should be handled when a failover is done.
        :param pulumi.Input[str] target_recovery_protection_container_id: Id of protection container where the VM replication should be created when a failover is done.
        :param pulumi.Input[str] target_resource_group_id: Id of resource group where the VM should be created when a failover is done.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicatedVMArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a VM replicated using Azure Site Recovery (Azure to Azure only). A replicated VM keeps a copiously updated image of the VM in another region in order to be able to start the VM in that region in case of a disaster.

        ## Import

        Site Recovery Replicated VM's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:siterecovery/replicatedVM:ReplicatedVM vmreplication /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name/replicationProtectedItems/vm-replication-name
        ```

        :param str resource_name: The name of the resource.
        :param ReplicatedVMArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicatedVMArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 managed_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMManagedDiskArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMNetworkInterfaceArgs']]]]] = None,
                 recovery_replication_policy_id: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_recovery_fabric_name: Optional[pulumi.Input[str]] = None,
                 source_recovery_protection_container_name: Optional[pulumi.Input[str]] = None,
                 source_vm_id: Optional[pulumi.Input[str]] = None,
                 target_availability_set_id: Optional[pulumi.Input[str]] = None,
                 target_network_id: Optional[pulumi.Input[str]] = None,
                 target_recovery_fabric_id: Optional[pulumi.Input[str]] = None,
                 target_recovery_protection_container_id: Optional[pulumi.Input[str]] = None,
                 target_resource_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['managed_disks'] = managed_disks
            __props__['name'] = name
            __props__['network_interfaces'] = network_interfaces
            if recovery_replication_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'recovery_replication_policy_id'")
            __props__['recovery_replication_policy_id'] = recovery_replication_policy_id
            if recovery_vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_recovery_fabric_name is None and not opts.urn:
                raise TypeError("Missing required property 'source_recovery_fabric_name'")
            __props__['source_recovery_fabric_name'] = source_recovery_fabric_name
            if source_recovery_protection_container_name is None and not opts.urn:
                raise TypeError("Missing required property 'source_recovery_protection_container_name'")
            __props__['source_recovery_protection_container_name'] = source_recovery_protection_container_name
            if source_vm_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_vm_id'")
            __props__['source_vm_id'] = source_vm_id
            __props__['target_availability_set_id'] = target_availability_set_id
            __props__['target_network_id'] = target_network_id
            if target_recovery_fabric_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_recovery_fabric_id'")
            __props__['target_recovery_fabric_id'] = target_recovery_fabric_id
            if target_recovery_protection_container_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_recovery_protection_container_id'")
            __props__['target_recovery_protection_container_id'] = target_recovery_protection_container_id
            if target_resource_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_group_id'")
            __props__['target_resource_group_id'] = target_resource_group_id
        super(ReplicatedVM, __self__).__init__(
            'azure:siterecovery/replicatedVM:ReplicatedVM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            managed_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMManagedDiskArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMNetworkInterfaceArgs']]]]] = None,
            recovery_replication_policy_id: Optional[pulumi.Input[str]] = None,
            recovery_vault_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            source_recovery_fabric_name: Optional[pulumi.Input[str]] = None,
            source_recovery_protection_container_name: Optional[pulumi.Input[str]] = None,
            source_vm_id: Optional[pulumi.Input[str]] = None,
            target_availability_set_id: Optional[pulumi.Input[str]] = None,
            target_network_id: Optional[pulumi.Input[str]] = None,
            target_recovery_fabric_id: Optional[pulumi.Input[str]] = None,
            target_recovery_protection_container_id: Optional[pulumi.Input[str]] = None,
            target_resource_group_id: Optional[pulumi.Input[str]] = None) -> 'ReplicatedVM':
        """
        Get an existing ReplicatedVM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMManagedDiskArgs']]]] managed_disks: One or more `managed_disk` block.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatedVMNetworkInterfaceArgs']]]] network_interfaces: One or more `network_interface` block.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_recovery_fabric_name: Name of fabric that should contains this replication.
        :param pulumi.Input[str] source_recovery_protection_container_name: Name of the protection container to use.
        :param pulumi.Input[str] source_vm_id: Id of the VM to replicate
        :param pulumi.Input[str] target_availability_set_id: Id of availability set that the new VM should belong to when a failover is done.
        :param pulumi.Input[str] target_network_id: Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        :param pulumi.Input[str] target_recovery_fabric_id: Id of fabric where the VM replication should be handled when a failover is done.
        :param pulumi.Input[str] target_recovery_protection_container_id: Id of protection container where the VM replication should be created when a failover is done.
        :param pulumi.Input[str] target_resource_group_id: Id of resource group where the VM should be created when a failover is done.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["managed_disks"] = managed_disks
        __props__["name"] = name
        __props__["network_interfaces"] = network_interfaces
        __props__["recovery_replication_policy_id"] = recovery_replication_policy_id
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["source_recovery_fabric_name"] = source_recovery_fabric_name
        __props__["source_recovery_protection_container_name"] = source_recovery_protection_container_name
        __props__["source_vm_id"] = source_vm_id
        __props__["target_availability_set_id"] = target_availability_set_id
        __props__["target_network_id"] = target_network_id
        __props__["target_recovery_fabric_id"] = target_recovery_fabric_id
        __props__["target_recovery_protection_container_id"] = target_recovery_protection_container_id
        __props__["target_resource_group_id"] = target_resource_group_id
        return ReplicatedVM(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="managedDisks")
    def managed_disks(self) -> pulumi.Output[Optional[Sequence['outputs.ReplicatedVMManagedDisk']]]:
        """
        One or more `managed_disk` block.
        """
        return pulumi.get(self, "managed_disks")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the network mapping.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[Sequence['outputs.ReplicatedVMNetworkInterface']]:
        """
        One or more `network_interface` block.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="recoveryReplicationPolicyId")
    def recovery_replication_policy_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "recovery_replication_policy_id")

    @property
    @pulumi.getter(name="recoveryVaultName")
    def recovery_vault_name(self) -> pulumi.Output[str]:
        """
        The name of the vault that should be updated.
        """
        return pulumi.get(self, "recovery_vault_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Name of the resource group where the vault that should be updated is located.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sourceRecoveryFabricName")
    def source_recovery_fabric_name(self) -> pulumi.Output[str]:
        """
        Name of fabric that should contains this replication.
        """
        return pulumi.get(self, "source_recovery_fabric_name")

    @property
    @pulumi.getter(name="sourceRecoveryProtectionContainerName")
    def source_recovery_protection_container_name(self) -> pulumi.Output[str]:
        """
        Name of the protection container to use.
        """
        return pulumi.get(self, "source_recovery_protection_container_name")

    @property
    @pulumi.getter(name="sourceVmId")
    def source_vm_id(self) -> pulumi.Output[str]:
        """
        Id of the VM to replicate
        """
        return pulumi.get(self, "source_vm_id")

    @property
    @pulumi.getter(name="targetAvailabilitySetId")
    def target_availability_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        Id of availability set that the new VM should belong to when a failover is done.
        """
        return pulumi.get(self, "target_availability_set_id")

    @property
    @pulumi.getter(name="targetNetworkId")
    def target_network_id(self) -> pulumi.Output[str]:
        """
        Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        """
        return pulumi.get(self, "target_network_id")

    @property
    @pulumi.getter(name="targetRecoveryFabricId")
    def target_recovery_fabric_id(self) -> pulumi.Output[str]:
        """
        Id of fabric where the VM replication should be handled when a failover is done.
        """
        return pulumi.get(self, "target_recovery_fabric_id")

    @property
    @pulumi.getter(name="targetRecoveryProtectionContainerId")
    def target_recovery_protection_container_id(self) -> pulumi.Output[str]:
        """
        Id of protection container where the VM replication should be created when a failover is done.
        """
        return pulumi.get(self, "target_recovery_protection_container_id")

    @property
    @pulumi.getter(name="targetResourceGroupId")
    def target_resource_group_id(self) -> pulumi.Output[str]:
        """
        Id of resource group where the VM should be created when a failover is done.
        """
        return pulumi.get(self, "target_resource_group_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

