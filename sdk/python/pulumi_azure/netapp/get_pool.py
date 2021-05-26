# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPoolResult',
    'AwaitableGetPoolResult',
    'get_pool',
]

@pulumi.output_type
class GetPoolResult:
    """
    A collection of values returned by getPool.
    """
    def __init__(__self__, account_name=None, id=None, location=None, name=None, resource_group_name=None, service_level=None, size_in_tb=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if service_level and not isinstance(service_level, str):
            raise TypeError("Expected argument 'service_level' to be a str")
        pulumi.set(__self__, "service_level", service_level)
        if size_in_tb and not isinstance(size_in_tb, int):
            raise TypeError("Expected argument 'size_in_tb' to be a int")
        pulumi.set(__self__, "size_in_tb", size_in_tb)

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
        The Azure Region where the NetApp Pool exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

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
    @pulumi.getter(name="sizeInTb")
    def size_in_tb(self) -> int:
        """
        Provisioned size of the pool in TB.
        """
        return pulumi.get(self, "size_in_tb")


class AwaitableGetPoolResult(GetPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoolResult(
            account_name=self.account_name,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            service_level=self.service_level,
            size_in_tb=self.size_in_tb)


def get_pool(account_name: Optional[str] = None,
             name: Optional[str] = None,
             resource_group_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoolResult:
    """
    Uses this data source to access information about an existing NetApp Pool.

    ## NetApp Pool Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.netapp.get_pool(resource_group_name="acctestRG",
        account_name="acctestnetappaccount",
        name="acctestnetapppool")
    pulumi.export("netappPoolId", example.id)
    ```


    :param str account_name: The name of the NetApp account where the NetApp pool exists.
    :param str name: The name of the NetApp Pool.
    :param str resource_group_name: The Name of the Resource Group where the NetApp Pool exists.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:netapp/getPool:getPool', __args__, opts=opts, typ=GetPoolResult).value

    return AwaitableGetPoolResult(
        account_name=__ret__.account_name,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        service_level=__ret__.service_level,
        size_in_tb=__ret__.size_in_tb)


@_utilities.lift_output_func(get_pool)
def get_pool_apply(account_name: Optional[pulumi.Input[str]] = None,
                   name: Optional[pulumi.Input[str]] = None,
                   resource_group_name: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPoolResult]:
    ...
