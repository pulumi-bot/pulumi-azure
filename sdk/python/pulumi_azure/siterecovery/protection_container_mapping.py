# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ProtectionContainerMapping']


class ProtectionContainerMapping(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recovery_fabric_name: Optional[pulumi.Input[str]] = None,
                 recovery_replication_policy_id: Optional[pulumi.Input[str]] = None,
                 recovery_source_protection_container_name: Optional[pulumi.Input[str]] = None,
                 recovery_target_protection_container_id: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Azure recovery vault protection container mapping. A protection container mapping decides how to translate the protection container when a VM is migrated from one region to another.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        primary_resource_group = azure.core.ResourceGroup("primaryResourceGroup", location="West US")
        secondary_resource_group = azure.core.ResourceGroup("secondaryResourceGroup", location="East US")
        vault = azure.recoveryservices.Vault("vault",
            location=secondary_resource_group.location,
            resource_group_name=secondary_resource_group.name,
            sku="Standard")
        primary_fabric = azure.siterecovery.Fabric("primaryFabric",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            location=primary_resource_group.location)
        secondary_fabric = azure.siterecovery.Fabric("secondaryFabric",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            location=secondary_resource_group.location)
        primary_protection_container = azure.siterecovery.ProtectionContainer("primaryProtectionContainer",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            recovery_fabric_name=primary_fabric.name)
        secondary_protection_container = azure.siterecovery.ProtectionContainer("secondaryProtectionContainer",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            recovery_fabric_name=secondary_fabric.name)
        policy = azure.siterecovery.ReplicationPolicy("policy",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            recovery_point_retention_in_minutes=24 * 60,
            application_consistent_snapshot_frequency_in_minutes=4 * 60)
        container_mapping = azure.siterecovery.ProtectionContainerMapping("container-mapping",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            recovery_fabric_name=primary_fabric.name,
            recovery_source_protection_container_name=primary_protection_container.name,
            recovery_target_protection_container_id=secondary_protection_container.id,
            recovery_replication_policy_id=policy.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[str] recovery_fabric_name: Name of fabric that should contains the protection container to map.
        :param pulumi.Input[str] recovery_replication_policy_id: Id of the policy to use for this mapping.
        :param pulumi.Input[str] recovery_source_protection_container_name: Name of the source protection container to map.
        :param pulumi.Input[str] recovery_target_protection_container_id: Id of target protection container to map to.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        """
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

            __props__['name'] = name
            if recovery_fabric_name is None:
                raise TypeError("Missing required property 'recovery_fabric_name'")
            __props__['recovery_fabric_name'] = recovery_fabric_name
            if recovery_replication_policy_id is None:
                raise TypeError("Missing required property 'recovery_replication_policy_id'")
            __props__['recovery_replication_policy_id'] = recovery_replication_policy_id
            if recovery_source_protection_container_name is None:
                raise TypeError("Missing required property 'recovery_source_protection_container_name'")
            __props__['recovery_source_protection_container_name'] = recovery_source_protection_container_name
            if recovery_target_protection_container_id is None:
                raise TypeError("Missing required property 'recovery_target_protection_container_id'")
            __props__['recovery_target_protection_container_id'] = recovery_target_protection_container_id
            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(ProtectionContainerMapping, __self__).__init__(
            'azure:siterecovery/protectionContainerMapping:ProtectionContainerMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            recovery_fabric_name: Optional[pulumi.Input[str]] = None,
            recovery_replication_policy_id: Optional[pulumi.Input[str]] = None,
            recovery_source_protection_container_name: Optional[pulumi.Input[str]] = None,
            recovery_target_protection_container_id: Optional[pulumi.Input[str]] = None,
            recovery_vault_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'ProtectionContainerMapping':
        """
        Get an existing ProtectionContainerMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[str] recovery_fabric_name: Name of fabric that should contains the protection container to map.
        :param pulumi.Input[str] recovery_replication_policy_id: Id of the policy to use for this mapping.
        :param pulumi.Input[str] recovery_source_protection_container_name: Name of the source protection container to map.
        :param pulumi.Input[str] recovery_target_protection_container_id: Id of target protection container to map to.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["recovery_fabric_name"] = recovery_fabric_name
        __props__["recovery_replication_policy_id"] = recovery_replication_policy_id
        __props__["recovery_source_protection_container_name"] = recovery_source_protection_container_name
        __props__["recovery_target_protection_container_id"] = recovery_target_protection_container_id
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        return ProtectionContainerMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the network mapping.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recoveryFabricName")
    def recovery_fabric_name(self) -> str:
        """
        Name of fabric that should contains the protection container to map.
        """
        return pulumi.get(self, "recovery_fabric_name")

    @property
    @pulumi.getter(name="recoveryReplicationPolicyId")
    def recovery_replication_policy_id(self) -> str:
        """
        Id of the policy to use for this mapping.
        """
        return pulumi.get(self, "recovery_replication_policy_id")

    @property
    @pulumi.getter(name="recoverySourceProtectionContainerName")
    def recovery_source_protection_container_name(self) -> str:
        """
        Name of the source protection container to map.
        """
        return pulumi.get(self, "recovery_source_protection_container_name")

    @property
    @pulumi.getter(name="recoveryTargetProtectionContainerId")
    def recovery_target_protection_container_id(self) -> str:
        """
        Id of target protection container to map to.
        """
        return pulumi.get(self, "recovery_target_protection_container_id")

    @property
    @pulumi.getter(name="recoveryVaultName")
    def recovery_vault_name(self) -> str:
        """
        The name of the vault that should be updated.
        """
        return pulumi.get(self, "recovery_vault_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        Name of the resource group where the vault that should be updated is located.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

