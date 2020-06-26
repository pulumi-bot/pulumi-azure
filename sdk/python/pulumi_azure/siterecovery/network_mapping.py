# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class NetworkMapping(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the network mapping.
    """
    recovery_vault_name: pulumi.Output[str]
    """
    The name of the vault that should be updated.
    """
    resource_group_name: pulumi.Output[str]
    """
    Name of the resource group where the vault that should be updated is located.
    """
    source_network_id: pulumi.Output[str]
    """
    The id of the primary network.
    """
    source_recovery_fabric_name: pulumi.Output[str]
    """
    Specifies the ASR fabric where mapping should be created.
    """
    target_network_id: pulumi.Output[str]
    """
    The id of the recovery network.
    """
    target_recovery_fabric_name: pulumi.Output[str]
    """
    The Azure Site Recovery fabric object corresponding to the recovery Azure region.
    """
    def __init__(__self__, resource_name, opts=None, name=None, recovery_vault_name=None, resource_group_name=None, source_network_id=None, source_recovery_fabric_name=None, target_network_id=None, target_recovery_fabric_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a site recovery network mapping on Azure. A network mapping decides how to translate connected netwroks when a VM is migrated from one region to another.

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
        # Avoids issues with crearing fabrics simultainusly
        primary_virtual_network = azure.network.VirtualNetwork("primaryVirtualNetwork",
            resource_group_name=primary_resource_group.name,
            address_spaces=["192.168.1.0/24"],
            location=primary_resource_group.location)
        secondary_virtual_network = azure.network.VirtualNetwork("secondaryVirtualNetwork",
            resource_group_name=secondary_resource_group.name,
            address_spaces=["192.168.2.0/24"],
            location=secondary_resource_group.location)
        recovery_mapping = azure.siterecovery.NetworkMapping("recovery-mapping",
            resource_group_name=secondary_resource_group.name,
            recovery_vault_name=vault.name,
            source_recovery_fabric_name="primary-fabric",
            target_recovery_fabric_name="secondary-fabric",
            source_network_id=primary_virtual_network.id,
            target_network_id=secondary_virtual_network.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_network_id: The id of the primary network.
        :param pulumi.Input[str] source_recovery_fabric_name: Specifies the ASR fabric where mapping should be created.
        :param pulumi.Input[str] target_network_id: The id of the recovery network.
        :param pulumi.Input[str] target_recovery_fabric_name: The Azure Site Recovery fabric object corresponding to the recovery Azure region.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_network_id is None:
                raise TypeError("Missing required property 'source_network_id'")
            __props__['source_network_id'] = source_network_id
            if source_recovery_fabric_name is None:
                raise TypeError("Missing required property 'source_recovery_fabric_name'")
            __props__['source_recovery_fabric_name'] = source_recovery_fabric_name
            if target_network_id is None:
                raise TypeError("Missing required property 'target_network_id'")
            __props__['target_network_id'] = target_network_id
            if target_recovery_fabric_name is None:
                raise TypeError("Missing required property 'target_recovery_fabric_name'")
            __props__['target_recovery_fabric_name'] = target_recovery_fabric_name
        super(NetworkMapping, __self__).__init__(
            'azure:siterecovery/networkMapping:NetworkMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, recovery_vault_name=None, resource_group_name=None, source_network_id=None, source_recovery_fabric_name=None, target_network_id=None, target_recovery_fabric_name=None):
        """
        Get an existing NetworkMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_network_id: The id of the primary network.
        :param pulumi.Input[str] source_recovery_fabric_name: Specifies the ASR fabric where mapping should be created.
        :param pulumi.Input[str] target_network_id: The id of the recovery network.
        :param pulumi.Input[str] target_recovery_fabric_name: The Azure Site Recovery fabric object corresponding to the recovery Azure region.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["source_network_id"] = source_network_id
        __props__["source_recovery_fabric_name"] = source_recovery_fabric_name
        __props__["target_network_id"] = target_network_id
        __props__["target_recovery_fabric_name"] = target_recovery_fabric_name
        return NetworkMapping(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
