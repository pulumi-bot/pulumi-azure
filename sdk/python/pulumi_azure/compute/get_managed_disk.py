# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetManagedDiskResult',
    'AwaitableGetManagedDiskResult',
    'get_managed_disk',
]

@pulumi.output_type
class GetManagedDiskResult:
    """
    A collection of values returned by getManagedDisk.
    """
    def __init__(__self__, create_option=None, disk_encryption_set_id=None, disk_iops_read_write=None, disk_mbps_read_write=None, disk_size_gb=None, id=None, image_reference_id=None, name=None, os_type=None, resource_group_name=None, source_resource_id=None, source_uri=None, storage_account_id=None, storage_account_type=None, tags=None, zones=None):
        if create_option and not isinstance(create_option, str):
            raise TypeError("Expected argument 'create_option' to be a str")
        pulumi.set(__self__, "create_option", create_option)
        if disk_encryption_set_id and not isinstance(disk_encryption_set_id, str):
            raise TypeError("Expected argument 'disk_encryption_set_id' to be a str")
        pulumi.set(__self__, "disk_encryption_set_id", disk_encryption_set_id)
        if disk_iops_read_write and not isinstance(disk_iops_read_write, int):
            raise TypeError("Expected argument 'disk_iops_read_write' to be a int")
        pulumi.set(__self__, "disk_iops_read_write", disk_iops_read_write)
        if disk_mbps_read_write and not isinstance(disk_mbps_read_write, int):
            raise TypeError("Expected argument 'disk_mbps_read_write' to be a int")
        pulumi.set(__self__, "disk_mbps_read_write", disk_mbps_read_write)
        if disk_size_gb and not isinstance(disk_size_gb, int):
            raise TypeError("Expected argument 'disk_size_gb' to be a int")
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_reference_id and not isinstance(image_reference_id, str):
            raise TypeError("Expected argument 'image_reference_id' to be a str")
        pulumi.set(__self__, "image_reference_id", image_reference_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if source_resource_id and not isinstance(source_resource_id, str):
            raise TypeError("Expected argument 'source_resource_id' to be a str")
        pulumi.set(__self__, "source_resource_id", source_resource_id)
        if source_uri and not isinstance(source_uri, str):
            raise TypeError("Expected argument 'source_uri' to be a str")
        pulumi.set(__self__, "source_uri", source_uri)
        if storage_account_id and not isinstance(storage_account_id, str):
            raise TypeError("Expected argument 'storage_account_id' to be a str")
        pulumi.set(__self__, "storage_account_id", storage_account_id)
        if storage_account_type and not isinstance(storage_account_type, str):
            raise TypeError("Expected argument 'storage_account_type' to be a str")
        pulumi.set(__self__, "storage_account_type", storage_account_type)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> str:
        return pulumi.get(self, "create_option")

    @property
    @pulumi.getter(name="diskEncryptionSetId")
    def disk_encryption_set_id(self) -> str:
        """
        The ID of the Disk Encryption Set used to encrypt this Managed Disk.
        """
        return pulumi.get(self, "disk_encryption_set_id")

    @property
    @pulumi.getter(name="diskIopsReadWrite")
    def disk_iops_read_write(self) -> int:
        """
        The number of IOPS allowed for this disk, where one operation can transfer between 4k and 256k bytes.
        """
        return pulumi.get(self, "disk_iops_read_write")

    @property
    @pulumi.getter(name="diskMbpsReadWrite")
    def disk_mbps_read_write(self) -> int:
        """
        The bandwidth allowed for this disk.
        """
        return pulumi.get(self, "disk_mbps_read_write")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> int:
        """
        The size of the Managed Disk in gigabytes.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageReferenceId")
    def image_reference_id(self) -> str:
        """
        The ID of the source image used for creating this Managed Disk.
        """
        return pulumi.get(self, "image_reference_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> str:
        """
        The operating system used for this Managed Disk.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> str:
        """
        The ID of an existing Managed Disk which this Disk was created from.
        """
        return pulumi.get(self, "source_resource_id")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> str:
        """
        The Source URI for this Managed Disk.
        """
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> str:
        """
        The ID of the Storage Account where the `source_uri` is located.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter(name="storageAccountType")
    def storage_account_type(self) -> str:
        """
        The storage account type for the Managed Disk.
        """
        return pulumi.get(self, "storage_account_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[str]:
        """
        A list of Availability Zones where the Managed Disk exists.
        """
        return pulumi.get(self, "zones")


class AwaitableGetManagedDiskResult(GetManagedDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedDiskResult(
            create_option=self.create_option,
            disk_encryption_set_id=self.disk_encryption_set_id,
            disk_iops_read_write=self.disk_iops_read_write,
            disk_mbps_read_write=self.disk_mbps_read_write,
            disk_size_gb=self.disk_size_gb,
            id=self.id,
            image_reference_id=self.image_reference_id,
            name=self.name,
            os_type=self.os_type,
            resource_group_name=self.resource_group_name,
            source_resource_id=self.source_resource_id,
            source_uri=self.source_uri,
            storage_account_id=self.storage_account_id,
            storage_account_type=self.storage_account_type,
            tags=self.tags,
            zones=self.zones)


def get_managed_disk(name: Optional[str] = None,
                     resource_group_name: Optional[str] = None,
                     tags: Optional[Mapping[str, str]] = None,
                     zones: Optional[Sequence[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedDiskResult:
    """
    Use this data source to access information about an existing Managed Disk.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    existing = azure.compute.get_managed_disk(name="example-datadisk",
        resource_group_name="example-resources")
    pulumi.export("id", existing.id)
    ```


    :param str name: Specifies the name of the Managed Disk.
    :param str resource_group_name: Specifies the name of the Resource Group where this Managed Disk exists.
    :param Mapping[str, str] tags: A mapping of tags assigned to the resource.
    :param Sequence[str] zones: A list of Availability Zones where the Managed Disk exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['tags'] = tags
    __args__['zones'] = zones
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getManagedDisk:getManagedDisk', __args__, opts=opts, typ=GetManagedDiskResult).value

    return AwaitableGetManagedDiskResult(
        create_option=__ret__.create_option,
        disk_encryption_set_id=__ret__.disk_encryption_set_id,
        disk_iops_read_write=__ret__.disk_iops_read_write,
        disk_mbps_read_write=__ret__.disk_mbps_read_write,
        disk_size_gb=__ret__.disk_size_gb,
        id=__ret__.id,
        image_reference_id=__ret__.image_reference_id,
        name=__ret__.name,
        os_type=__ret__.os_type,
        resource_group_name=__ret__.resource_group_name,
        source_resource_id=__ret__.source_resource_id,
        source_uri=__ret__.source_uri,
        storage_account_id=__ret__.storage_account_id,
        storage_account_type=__ret__.storage_account_type,
        tags=__ret__.tags,
        zones=__ret__.zones)
