# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Volume(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
    """
    export_policy_rules: pulumi.Output[list]
    """
    One or more `export_policy_rule` block defined below.

      * `allowedClients` (`list`) - A list of allowed clients IPv4 addresses.
      * `cifsEnabled` (`bool`) - Is the CIFS protocol allowed?
      * `nfsv3Enabled` (`bool`) - Is the NFSv3 protocol allowed?
      * `nfsv4Enabled` (`bool`) - Is the NFSv4 protocol allowed?
      * `protocolsEnabled` (`str`) - A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifs_enabled`, `nfsv3_enabled` and `nfsv4_enabled`.
      * `ruleIndex` (`float`) - The index number of the rule.
      * `unixReadOnly` (`bool`) - Is the file system on unix read only?
      * `unixReadWrite` (`bool`) - Is the file system on unix read and write?
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    mount_ip_addresses: pulumi.Output[list]
    """
    A list of IPv4 Addresses which should be used to mount the volume.
    """
    name: pulumi.Output[str]
    """
    The name of the NetApp Volume. Changing this forces a new resource to be created.
    """
    pool_name: pulumi.Output[str]
    """
    The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
    """
    protocols: pulumi.Output[list]
    """
    The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
    """
    service_level: pulumi.Output[str]
    """
    The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
    """
    storage_quota_in_gb: pulumi.Output[float]
    """
    The maximum Storage Quota allowed for a file system in Gigabytes.
    """
    subnet_id: pulumi.Output[str]
    """
    The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    volume_path: pulumi.Output[str]
    """
    A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, export_policy_rules=None, location=None, name=None, pool_name=None, protocols=None, resource_group_name=None, service_level=None, storage_quota_in_gb=None, subnet_id=None, tags=None, volume_path=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a NetApp Volume.

        ## NetApp Volume Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/16"])
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.2.0/24",
            delegations=[{
                "name": "netapp",
                "serviceDelegation": {
                    "name": "Microsoft.Netapp/volumes",
                    "actions": [
                        "Microsoft.Network/networkinterfaces/*",
                        "Microsoft.Network/virtualNetworks/subnets/join/action",
                    ],
                },
            }])
        example_account = azure.netapp.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pool = azure.netapp.Pool("examplePool",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_name=example_account.name,
            service_level="Premium",
            size_in_tb=4)
        example_volume = azure.netapp.Volume("exampleVolume",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_name=example_account.name,
            pool_name=example_pool.name,
            volume_path="my-unique-file-path",
            service_level="Premium",
            subnet_id=example_subnet.id,
            protocols=["NFSv4.1"],
            storage_quota_in_gb=100)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[list] export_policy_rules: One or more `export_policy_rule` block defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Volume. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_name: The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[list] protocols: The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[float] storage_quota_in_gb: The maximum Storage Quota allowed for a file system in Gigabytes.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] volume_path: A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.

        The **export_policy_rules** object supports the following:

          * `allowedClients` (`pulumi.Input[list]`) - A list of allowed clients IPv4 addresses.
          * `cifsEnabled` (`pulumi.Input[bool]`) - Is the CIFS protocol allowed?
          * `nfsv3Enabled` (`pulumi.Input[bool]`) - Is the NFSv3 protocol allowed?
          * `nfsv4Enabled` (`pulumi.Input[bool]`) - Is the NFSv4 protocol allowed?
          * `protocolsEnabled` (`pulumi.Input[str]`) - A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifs_enabled`, `nfsv3_enabled` and `nfsv4_enabled`.
          * `ruleIndex` (`pulumi.Input[float]`) - The index number of the rule.
          * `unixReadOnly` (`pulumi.Input[bool]`) - Is the file system on unix read only?
          * `unixReadWrite` (`pulumi.Input[bool]`) - Is the file system on unix read and write?
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['export_policy_rules'] = export_policy_rules
            __props__['location'] = location
            __props__['name'] = name
            if pool_name is None:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            __props__['protocols'] = protocols
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_level is None:
                raise TypeError("Missing required property 'service_level'")
            __props__['service_level'] = service_level
            if storage_quota_in_gb is None:
                raise TypeError("Missing required property 'storage_quota_in_gb'")
            __props__['storage_quota_in_gb'] = storage_quota_in_gb
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            if volume_path is None:
                raise TypeError("Missing required property 'volume_path'")
            __props__['volume_path'] = volume_path
            __props__['mount_ip_addresses'] = None
        super(Volume, __self__).__init__(
            'azure:netapp/volume:Volume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, export_policy_rules=None, location=None, mount_ip_addresses=None, name=None, pool_name=None, protocols=None, resource_group_name=None, service_level=None, storage_quota_in_gb=None, subnet_id=None, tags=None, volume_path=None):
        """
        Get an existing Volume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[list] export_policy_rules: One or more `export_policy_rule` block defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[list] mount_ip_addresses: A list of IPv4 Addresses which should be used to mount the volume.
        :param pulumi.Input[str] name: The name of the NetApp Volume. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_name: The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[list] protocols: The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[float] storage_quota_in_gb: The maximum Storage Quota allowed for a file system in Gigabytes.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] volume_path: A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.

        The **export_policy_rules** object supports the following:

          * `allowedClients` (`pulumi.Input[list]`) - A list of allowed clients IPv4 addresses.
          * `cifsEnabled` (`pulumi.Input[bool]`) - Is the CIFS protocol allowed?
          * `nfsv3Enabled` (`pulumi.Input[bool]`) - Is the NFSv3 protocol allowed?
          * `nfsv4Enabled` (`pulumi.Input[bool]`) - Is the NFSv4 protocol allowed?
          * `protocolsEnabled` (`pulumi.Input[str]`) - A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifs_enabled`, `nfsv3_enabled` and `nfsv4_enabled`.
          * `ruleIndex` (`pulumi.Input[float]`) - The index number of the rule.
          * `unixReadOnly` (`pulumi.Input[bool]`) - Is the file system on unix read only?
          * `unixReadWrite` (`pulumi.Input[bool]`) - Is the file system on unix read and write?
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["export_policy_rules"] = export_policy_rules
        __props__["location"] = location
        __props__["mount_ip_addresses"] = mount_ip_addresses
        __props__["name"] = name
        __props__["pool_name"] = pool_name
        __props__["protocols"] = protocols
        __props__["resource_group_name"] = resource_group_name
        __props__["service_level"] = service_level
        __props__["storage_quota_in_gb"] = storage_quota_in_gb
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["volume_path"] = volume_path
        return Volume(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
