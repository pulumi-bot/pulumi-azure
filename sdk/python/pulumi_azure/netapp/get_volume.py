# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetVolumeResult',
    'AwaitableGetVolumeResult',
    'get_volume',
]

@pulumi.output_type
class GetVolumeResult:
    """
    A collection of values returned by getVolume.
    """
    def __init__(__self__, account_name=None, id=None, location=None, mount_ip_addresses=None, name=None, pool_name=None, protocols=None, resource_group_name=None, service_level=None, storage_quota_in_gb=None, subnet_id=None, volume_path=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if mount_ip_addresses and not isinstance(mount_ip_addresses, list):
            raise TypeError("Expected argument 'mount_ip_addresses' to be a list")
        pulumi.set(__self__, "mount_ip_addresses", mount_ip_addresses)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pool_name and not isinstance(pool_name, str):
            raise TypeError("Expected argument 'pool_name' to be a str")
        pulumi.set(__self__, "pool_name", pool_name)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if service_level and not isinstance(service_level, str):
            raise TypeError("Expected argument 'service_level' to be a str")
        pulumi.set(__self__, "service_level", service_level)
        if storage_quota_in_gb and not isinstance(storage_quota_in_gb, int):
            raise TypeError("Expected argument 'storage_quota_in_gb' to be a int")
        pulumi.set(__self__, "storage_quota_in_gb", storage_quota_in_gb)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if volume_path and not isinstance(volume_path, str):
            raise TypeError("Expected argument 'volume_path' to be a str")
        pulumi.set(__self__, "volume_path", volume_path)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure Region where the NetApp Volume exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mountIpAddresses")
    def mount_ip_addresses(self) -> List[str]:
        """
        A list of IPv4 Addresses which should be used to mount the volume.
        """
        return pulumi.get(self, "mount_ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolName")
    def pool_name(self) -> str:
        return pulumi.get(self, "pool_name")

    @property
    @pulumi.getter
    def protocols(self) -> List[str]:
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> str:
        """
        The service level of the file system.
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter(name="storageQuotaInGb")
    def storage_quota_in_gb(self) -> int:
        """
        The maximum Storage Quota in Gigabytes allowed for a file system.
        """
        return pulumi.get(self, "storage_quota_in_gb")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of a Subnet in which the NetApp Volume resides.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="volumePath")
    def volume_path(self) -> str:
        """
        The unique file path of the volume.
        """
        return pulumi.get(self, "volume_path")


class AwaitableGetVolumeResult(GetVolumeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeResult(
            account_name=self.account_name,
            id=self.id,
            location=self.location,
            mount_ip_addresses=self.mount_ip_addresses,
            name=self.name,
            pool_name=self.pool_name,
            protocols=self.protocols,
            resource_group_name=self.resource_group_name,
            service_level=self.service_level,
            storage_quota_in_gb=self.storage_quota_in_gb,
            subnet_id=self.subnet_id,
            volume_path=self.volume_path)


def get_volume(account_name: Optional[str] = None,
               name: Optional[str] = None,
               pool_name: Optional[str] = None,
               resource_group_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeResult:
    """
    Uses this data source to access information about an existing NetApp Volume.

    ## NetApp Volume Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.netapp.get_volume(resource_group_name="acctestRG",
        account_name="acctestnetappaccount",
        pool_name="acctestnetapppool",
        name="example-volume")
    pulumi.export("netappVolumeId", example.id)
    ```


    :param str account_name: The name of the NetApp account where the NetApp pool exists.
    :param str name: The name of the NetApp Volume.
    :param str pool_name: The name of the NetApp pool where the NetApp volume exists.
    :param str resource_group_name: The Name of the Resource Group where the NetApp Volume exists.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['name'] = name
    __args__['poolName'] = pool_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:netapp/getVolume:getVolume', __args__, opts=opts, typ=GetVolumeResult).value

    return AwaitableGetVolumeResult(
        account_name=__ret__.account_name,
        id=__ret__.id,
        location=__ret__.location,
        mount_ip_addresses=__ret__.mount_ip_addresses,
        name=__ret__.name,
        pool_name=__ret__.pool_name,
        protocols=__ret__.protocols,
        resource_group_name=__ret__.resource_group_name,
        service_level=__ret__.service_level,
        storage_quota_in_gb=__ret__.storage_quota_in_gb,
        subnet_id=__ret__.subnet_id,
        volume_path=__ret__.volume_path)
