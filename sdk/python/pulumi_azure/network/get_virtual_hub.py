# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetVirtualHubResult',
    'AwaitableGetVirtualHubResult',
    'get_virtual_hub',
]


class GetVirtualHubResult:
    """
    A collection of values returned by getVirtualHub.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, address_prefix=None, id=None, location=None, name=None, resource_group_name=None, tags=None, virtual_wan_id=None) -> None:
        if address_prefix and not isinstance(address_prefix, str):
            raise TypeError("Expected argument 'address_prefix' to be a str")
        __self__.address_prefix = address_prefix
        """
        The Address Prefix used for this Virtual Hub.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure Region where the Virtual Hub exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the Virtual Hub.
        """
        if virtual_wan_id and not isinstance(virtual_wan_id, str):
            raise TypeError("Expected argument 'virtual_wan_id' to be a str")
        __self__.virtual_wan_id = virtual_wan_id
        """
        The ID of the Virtual WAN within which the Virtual Hub exists.
        """


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


def get_virtual_hub(name: Optional[str] = None, resource_group_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualHubResult:
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
    __ret__ = pulumi.runtime.invoke('azure:network/getVirtualHub:getVirtualHub', __args__, opts=opts).value

    return AwaitableGetVirtualHubResult(
        address_prefix=__ret__.get('addressPrefix'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        virtual_wan_id=__ret__.get('virtualWanId'))
