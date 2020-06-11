# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ManagedDisk(pulumi.CustomResource):
    create_option: pulumi.Output[str]
    """
    The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
    """
    disk_encryption_set_id: pulumi.Output[str]
    """
    The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
    """
    disk_iops_read_write: pulumi.Output[float]
    """
    The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
    """
    disk_mbps_read_write: pulumi.Output[float]
    """
    The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
    """
    disk_size_gb: pulumi.Output[float]
    """
    Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
    """
    encryption_settings: pulumi.Output[dict]
    """
    A `encryption_settings` block as defined below.

      * `diskEncryptionKey` (`dict`) - A `disk_encryption_key` block as defined above.
        * `secretUrl` (`str`) - The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as `id` on the `keyvault.Secret` resource.
        * `sourceVaultId` (`str`) - The ID of the source Key Vault.

      * `enabled` (`bool`) - Is Encryption enabled on this Managed Disk? Changing this forces a new resource to be created.
      * `keyEncryptionKey` (`dict`) - A `key_encryption_key` block as defined below.
        * `keyUrl` (`str`) - The URL to the Key Vault Key used as the Key Encryption Key. This can be found as `id` on the `keyvault.Key` resource.
        * `sourceVaultId` (`str`) - The ID of the source Key Vault.
    """
    image_reference_id: pulumi.Output[str]
    """
    ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
    """
    location: pulumi.Output[str]
    """
    Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
    """
    os_type: pulumi.Output[str]
    """
    Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Managed Disk should exist.
    """
    source_resource_id: pulumi.Output[str]
    """
    The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
    """
    source_uri: pulumi.Output[str]
    """
    URI to a valid VHD file to be used when `create_option` is `Import`.
    """
    storage_account_id: pulumi.Output[str]
    """
    The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
    """
    storage_account_type: pulumi.Output[str]
    """
    The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    zones: pulumi.Output[str]
    """
    A collection containing the availability zone to allocate the Managed Disk in.
    """
    def __init__(__self__, resource_name, opts=None, create_option=None, disk_encryption_set_id=None, disk_iops_read_write=None, disk_mbps_read_write=None, disk_size_gb=None, encryption_settings=None, image_reference_id=None, location=None, name=None, os_type=None, resource_group_name=None, source_resource_id=None, source_uri=None, storage_account_id=None, storage_account_type=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a managed disk.

        ## Example Usage with Create Empty

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US 2")
        example_managed_disk = azure.compute.ManagedDisk("exampleManagedDisk",
            location="West US 2",
            resource_group_name=example_resource_group.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb="1",
            tags={
                "environment": "staging",
            })
        ```

        ## Example Usage with Create Copy

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West US 2")
        source = azure.compute.ManagedDisk("source",
            location="West US 2",
            resource_group_name=example.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb="1",
            tags={
                "environment": "staging",
            })
        copy = azure.compute.ManagedDisk("copy",
            location="West US 2",
            resource_group_name=example.name,
            storage_account_type="Standard_LRS",
            create_option="Copy",
            source_resource_id=source.id,
            disk_size_gb="1",
            tags={
                "environment": "staging",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
        :param pulumi.Input[str] disk_encryption_set_id: The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        :param pulumi.Input[float] disk_iops_read_write: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        :param pulumi.Input[float] disk_mbps_read_write: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        :param pulumi.Input[float] disk_size_gb: Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        :param pulumi.Input[dict] encryption_settings: A `encryption_settings` block as defined below.
        :param pulumi.Input[str] image_reference_id: ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        :param pulumi.Input[str] location: Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Managed Disk should exist.
        :param pulumi.Input[str] source_resource_id: The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        :param pulumi.Input[str] source_uri: URI to a valid VHD file to be used when `create_option` is `Import`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_type: The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A collection containing the availability zone to allocate the Managed Disk in.

        The **encryption_settings** object supports the following:

          * `diskEncryptionKey` (`pulumi.Input[dict]`) - A `disk_encryption_key` block as defined above.
            * `secretUrl` (`pulumi.Input[str]`) - The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as `id` on the `keyvault.Secret` resource.
            * `sourceVaultId` (`pulumi.Input[str]`) - The ID of the source Key Vault.

          * `enabled` (`pulumi.Input[bool]`) - Is Encryption enabled on this Managed Disk? Changing this forces a new resource to be created.
          * `keyEncryptionKey` (`pulumi.Input[dict]`) - A `key_encryption_key` block as defined below.
            * `keyUrl` (`pulumi.Input[str]`) - The URL to the Key Vault Key used as the Key Encryption Key. This can be found as `id` on the `keyvault.Key` resource.
            * `sourceVaultId` (`pulumi.Input[str]`) - The ID of the source Key Vault.
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

            if create_option is None:
                raise TypeError("Missing required property 'create_option'")
            __props__['create_option'] = create_option
            __props__['disk_encryption_set_id'] = disk_encryption_set_id
            __props__['disk_iops_read_write'] = disk_iops_read_write
            __props__['disk_mbps_read_write'] = disk_mbps_read_write
            __props__['disk_size_gb'] = disk_size_gb
            __props__['encryption_settings'] = encryption_settings
            __props__['image_reference_id'] = image_reference_id
            __props__['location'] = location
            __props__['name'] = name
            __props__['os_type'] = os_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['source_resource_id'] = source_resource_id
            __props__['source_uri'] = source_uri
            __props__['storage_account_id'] = storage_account_id
            if storage_account_type is None:
                raise TypeError("Missing required property 'storage_account_type'")
            __props__['storage_account_type'] = storage_account_type
            __props__['tags'] = tags
            __props__['zones'] = zones
        super(ManagedDisk, __self__).__init__(
            'azure:compute/managedDisk:ManagedDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_option=None, disk_encryption_set_id=None, disk_iops_read_write=None, disk_mbps_read_write=None, disk_size_gb=None, encryption_settings=None, image_reference_id=None, location=None, name=None, os_type=None, resource_group_name=None, source_resource_id=None, source_uri=None, storage_account_id=None, storage_account_type=None, tags=None, zones=None):
        """
        Get an existing ManagedDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
        :param pulumi.Input[str] disk_encryption_set_id: The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        :param pulumi.Input[float] disk_iops_read_write: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        :param pulumi.Input[float] disk_mbps_read_write: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        :param pulumi.Input[float] disk_size_gb: Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        :param pulumi.Input[dict] encryption_settings: A `encryption_settings` block as defined below.
        :param pulumi.Input[str] image_reference_id: ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        :param pulumi.Input[str] location: Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Managed Disk should exist.
        :param pulumi.Input[str] source_resource_id: The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        :param pulumi.Input[str] source_uri: URI to a valid VHD file to be used when `create_option` is `Import`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_type: The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A collection containing the availability zone to allocate the Managed Disk in.

        The **encryption_settings** object supports the following:

          * `diskEncryptionKey` (`pulumi.Input[dict]`) - A `disk_encryption_key` block as defined above.
            * `secretUrl` (`pulumi.Input[str]`) - The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as `id` on the `keyvault.Secret` resource.
            * `sourceVaultId` (`pulumi.Input[str]`) - The ID of the source Key Vault.

          * `enabled` (`pulumi.Input[bool]`) - Is Encryption enabled on this Managed Disk? Changing this forces a new resource to be created.
          * `keyEncryptionKey` (`pulumi.Input[dict]`) - A `key_encryption_key` block as defined below.
            * `keyUrl` (`pulumi.Input[str]`) - The URL to the Key Vault Key used as the Key Encryption Key. This can be found as `id` on the `keyvault.Key` resource.
            * `sourceVaultId` (`pulumi.Input[str]`) - The ID of the source Key Vault.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_option"] = create_option
        __props__["disk_encryption_set_id"] = disk_encryption_set_id
        __props__["disk_iops_read_write"] = disk_iops_read_write
        __props__["disk_mbps_read_write"] = disk_mbps_read_write
        __props__["disk_size_gb"] = disk_size_gb
        __props__["encryption_settings"] = encryption_settings
        __props__["image_reference_id"] = image_reference_id
        __props__["location"] = location
        __props__["name"] = name
        __props__["os_type"] = os_type
        __props__["resource_group_name"] = resource_group_name
        __props__["source_resource_id"] = source_resource_id
        __props__["source_uri"] = source_uri
        __props__["storage_account_id"] = storage_account_id
        __props__["storage_account_type"] = storage_account_type
        __props__["tags"] = tags
        __props__["zones"] = zones
        return ManagedDisk(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
