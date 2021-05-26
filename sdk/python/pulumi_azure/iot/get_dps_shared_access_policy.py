# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDpsSharedAccessPolicyResult',
    'AwaitableGetDpsSharedAccessPolicyResult',
    'get_dps_shared_access_policy',
]

@pulumi.output_type
class GetDpsSharedAccessPolicyResult:
    """
    A collection of values returned by getDpsSharedAccessPolicy.
    """
    def __init__(__self__, id=None, iothub_dps_name=None, name=None, primary_connection_string=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_key=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iothub_dps_name and not isinstance(iothub_dps_name, str):
            raise TypeError("Expected argument 'iothub_dps_name' to be a str")
        pulumi.set(__self__, "iothub_dps_name", iothub_dps_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        pulumi.set(__self__, "primary_connection_string", primary_connection_string)
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        pulumi.set(__self__, "primary_key", primary_key)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        pulumi.set(__self__, "secondary_connection_string", secondary_connection_string)
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="iothubDpsName")
    def iothub_dps_name(self) -> str:
        return pulumi.get(self, "iothub_dps_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> str:
        """
        The primary connection string of the Shared Access Policy.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> str:
        """
        The primary key used to create the authentication token.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> str:
        """
        The secondary connection string of the Shared Access Policy.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> str:
        """
        The secondary key used to create the authentication token.
        """
        return pulumi.get(self, "secondary_key")


class AwaitableGetDpsSharedAccessPolicyResult(GetDpsSharedAccessPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDpsSharedAccessPolicyResult(
            id=self.id,
            iothub_dps_name=self.iothub_dps_name,
            name=self.name,
            primary_connection_string=self.primary_connection_string,
            primary_key=self.primary_key,
            resource_group_name=self.resource_group_name,
            secondary_connection_string=self.secondary_connection_string,
            secondary_key=self.secondary_key)


def get_dps_shared_access_policy(iothub_dps_name: Optional[str] = None,
                                 name: Optional[str] = None,
                                 resource_group_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDpsSharedAccessPolicyResult:
    """
    Use this data source to access information about an existing IotHub Device Provisioning Service Shared Access Policy

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.iot.get_dps_shared_access_policy(name="example",
        resource_group_name=azurerm_resource_group["example"]["name"],
        iothub_dps_name=azurerm_iothub_dps["example"]["name"])
    ```


    :param str iothub_dps_name: Specifies the name of the IoT Hub Device Provisioning service to which the Shared Access Policy belongs.
    :param str name: Specifies the name of the IotHub Shared Access Policy.
    :param str resource_group_name: Specifies the name of the resource group under which the IotHub Shared Access Policy resource exists.
    """
    __args__ = dict()
    __args__['iothubDpsName'] = iothub_dps_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:iot/getDpsSharedAccessPolicy:getDpsSharedAccessPolicy', __args__, opts=opts, typ=GetDpsSharedAccessPolicyResult).value

    return AwaitableGetDpsSharedAccessPolicyResult(
        id=__ret__.id,
        iothub_dps_name=__ret__.iothub_dps_name,
        name=__ret__.name,
        primary_connection_string=__ret__.primary_connection_string,
        primary_key=__ret__.primary_key,
        resource_group_name=__ret__.resource_group_name,
        secondary_connection_string=__ret__.secondary_connection_string,
        secondary_key=__ret__.secondary_key)


@_utilities.lift_output_func(get_dps_shared_access_policy)
def get_dps_shared_access_policy_apply(iothub_dps_name: Optional[pulumi.Input[str]] = None,
                                       name: Optional[pulumi.Input[str]] = None,
                                       resource_group_name: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDpsSharedAccessPolicyResult]:
    ...
