# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetVirtualHubResult',
    'AwaitableGetVirtualHubResult',
    'get_virtual_hub',
]

@pulumi.output_type
class GetVirtualHubResult:
    """
    A collection of values returned by getVirtualHub.
    """
    def __init__(__self__, address_prefix=None, id=None, location=None, name=None, resource_group_name=None, tags=None, virtual_wan_id=None):
        if address_prefix and not isinstance(address_prefix, str):
            raise TypeError("Expected argument 'address_prefix' to be a str")
        pulumi.set(__self__, "address_prefix", address_prefix)
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
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if virtual_wan_id and not isinstance(virtual_wan_id, str):
            raise TypeError("Expected argument 'virtual_wan_id' to be a str")
        pulumi.set(__self__, "virtual_wan_id", virtual_wan_id)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> str:
        """
        The Address Prefix used for this Virtual Hub.
        """
        return pulumi.get(self, "address_prefix")

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
        The Azure Region where the Virtual Hub exists.
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
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the Virtual Hub.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualWanId")
    def virtual_wan_id(self) -> str:
        """
        The ID of the Virtual WAN within which the Virtual Hub exists.
        """
        return pulumi.get(self, "virtual_wan_id")


class AwaitableGetVirtualHubResult(GetVirtualHubResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualHubResult(
            address_prefix=self.address_prefix,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            virtual_wan_id=self.virtual_wan_id)


def get_virtual_hub(name: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualHubResult:
    """
    Uses this data source to access information about an existing Virtual Hub.

    ## Virtual Hub Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_virtual_hub(name="example-hub",
        resource_group_name="example-resources")
    pulumi.export("virtualHubId", example.id)
    ```


    :param str name: The name of the Virtual Hub.
    :param str resource_group_name: The Name of the Resource Group where the Virtual Hub exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getVirtualHub:getVirtualHub', __args__, opts=opts, typ=GetVirtualHubResult).value

    return AwaitableGetVirtualHubResult(
        address_prefix=__ret__.address_prefix,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags,
        virtual_wan_id=__ret__.virtual_wan_id)


@_utilities.lift_output_func(get_virtual_hub)
def get_virtual_hub_apply(name: Optional[pulumi.Input[str]] = None,
                          resource_group_name: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVirtualHubResult]:
    ...
