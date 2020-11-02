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

__all__ = ['Snapshot']


class Snapshot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_option: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 encryption_settings: Optional[pulumi.Input[pulumi.InputType['SnapshotEncryptionSettingsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Disk Snapshot.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_managed_disk = azure.compute.ManagedDisk("exampleManagedDisk",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb=10)
        example_snapshot = azure.compute.Snapshot("exampleSnapshot",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            create_option="Copy",
            source_uri=example_managed_disk.id)
        ```

        ## Import

        Snapshots can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] disk_size_gb: The size of the Snapshotted Disk in GB.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_resource_id: Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_uri: Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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
            __props__['disk_size_gb'] = disk_size_gb
            __props__['encryption_settings'] = encryption_settings
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['source_resource_id'] = source_resource_id
            __props__['source_uri'] = source_uri
            __props__['storage_account_id'] = storage_account_id
            __props__['tags'] = tags
        super(Snapshot, __self__).__init__(
            'azure:compute/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_option: Optional[pulumi.Input[str]] = None,
            disk_size_gb: Optional[pulumi.Input[int]] = None,
            encryption_settings: Optional[pulumi.Input[pulumi.InputType['SnapshotEncryptionSettingsArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            source_resource_id: Optional[pulumi.Input[str]] = None,
            source_uri: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] disk_size_gb: The size of the Snapshotted Disk in GB.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_resource_id: Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_uri: Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_option"] = create_option
        __props__["disk_size_gb"] = disk_size_gb
        __props__["encryption_settings"] = encryption_settings
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["source_resource_id"] = source_resource_id
        __props__["source_uri"] = source_uri
        __props__["storage_account_id"] = storage_account_id
        __props__["tags"] = tags
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> pulumi.Output[str]:
        """
        Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "create_option")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> pulumi.Output[int]:
        """
        The size of the Snapshotted Disk in GB.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="encryptionSettings")
    def encryption_settings(self) -> pulumi.Output[Optional['outputs.SnapshotEncryptionSettings']]:
        return pulumi.get(self, "encryption_settings")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_resource_id")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

