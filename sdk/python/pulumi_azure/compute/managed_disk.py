# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ManagedDisk']


class ManagedDisk(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_option: Optional[pulumi.Input[str]] = None,
                 disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
                 disk_iops_read_write: Optional[pulumi.Input[int]] = None,
                 disk_mbps_read_write: Optional[pulumi.Input[int]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 encryption_settings: Optional[pulumi.Input[pulumi.InputType['ManagedDiskEncryptionSettingsArgs']]] = None,
                 image_reference_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 storage_account_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a managed disk.

        ## Example Usage
        ### With Create Empty

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US 2")
        example_managed_disk = azure.compute.ManagedDisk("exampleManagedDisk",
            location="West US 2",
            resource_group_name=example_resource_group.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb=1,
            tags={
                "environment": "staging",
            })
        ```
        ### With Create Copy

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West US 2")
        source = azure.compute.ManagedDisk("source",
            location="West US 2",
            resource_group_name=example.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb=1,
            tags={
                "environment": "staging",
            })
        copy = azure.compute.ManagedDisk("copy",
            location="West US 2",
            resource_group_name=example.name,
            storage_account_type="Standard_LRS",
            create_option="Copy",
            source_resource_id=source.id,
            disk_size_gb=1,
            tags={
                "environment": "staging",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
        :param pulumi.Input[str] disk_encryption_set_id: The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        :param pulumi.Input[int] disk_iops_read_write: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        :param pulumi.Input[int] disk_mbps_read_write: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        :param pulumi.Input[int] disk_size_gb: Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        :param pulumi.Input[pulumi.InputType['ManagedDiskEncryptionSettingsArgs']] encryption_settings: A `encryption_settings` block as defined below.
        :param pulumi.Input[str] image_reference_id: ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        :param pulumi.Input[str] location: Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Managed Disk should exist.
        :param pulumi.Input[str] source_resource_id: The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        :param pulumi.Input[str] source_uri: URI to a valid VHD file to be used when `create_option` is `Import`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_type: The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A collection containing the availability zone to allocate the Managed Disk in.
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_option: Optional[pulumi.Input[str]] = None,
            disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
            disk_iops_read_write: Optional[pulumi.Input[int]] = None,
            disk_mbps_read_write: Optional[pulumi.Input[int]] = None,
            disk_size_gb: Optional[pulumi.Input[int]] = None,
            encryption_settings: Optional[pulumi.Input[pulumi.InputType['ManagedDiskEncryptionSettingsArgs']]] = None,
            image_reference_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            os_type: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            source_resource_id: Optional[pulumi.Input[str]] = None,
            source_uri: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            storage_account_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            zones: Optional[pulumi.Input[str]] = None) -> 'ManagedDisk':
        """
        Get an existing ManagedDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
        :param pulumi.Input[str] disk_encryption_set_id: The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        :param pulumi.Input[int] disk_iops_read_write: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        :param pulumi.Input[int] disk_mbps_read_write: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        :param pulumi.Input[int] disk_size_gb: Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        :param pulumi.Input[pulumi.InputType['ManagedDiskEncryptionSettingsArgs']] encryption_settings: A `encryption_settings` block as defined below.
        :param pulumi.Input[str] image_reference_id: ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        :param pulumi.Input[str] location: Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Managed Disk should exist.
        :param pulumi.Input[str] source_resource_id: The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        :param pulumi.Input[str] source_uri: URI to a valid VHD file to be used when `create_option` is `Import`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_type: The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A collection containing the availability zone to allocate the Managed Disk in.
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

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> pulumi.Output[str]:
        """
        The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
        """
        return pulumi.get(self, "create_option")

    @property
    @pulumi.getter(name="diskEncryptionSetId")
    def disk_encryption_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        """
        return pulumi.get(self, "disk_encryption_set_id")

    @property
    @pulumi.getter(name="diskIopsReadWrite")
    def disk_iops_read_write(self) -> pulumi.Output[int]:
        """
        The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        """
        return pulumi.get(self, "disk_iops_read_write")

    @property
    @pulumi.getter(name="diskMbpsReadWrite")
    def disk_mbps_read_write(self) -> pulumi.Output[int]:
        """
        The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        """
        return pulumi.get(self, "disk_mbps_read_write")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> pulumi.Output[int]:
        """
        Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="encryptionSettings")
    def encryption_settings(self) -> pulumi.Output[Optional['outputs.ManagedDiskEncryptionSettings']]:
        """
        A `encryption_settings` block as defined below.
        """
        return pulumi.get(self, "encryption_settings")

    @property
    @pulumi.getter(name="imageReferenceId")
    def image_reference_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        """
        return pulumi.get(self, "image_reference_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Managed Disk should exist.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        """
        return pulumi.get(self, "source_resource_id")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> pulumi.Output[str]:
        """
        URI to a valid VHD file to be used when `create_option` is `Import`.
        """
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter(name="storageAccountType")
    def storage_account_type(self) -> pulumi.Output[str]:
        """
        The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        """
        return pulumi.get(self, "storage_account_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[str]]:
        """
        A collection containing the availability zone to allocate the Managed Disk in.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

