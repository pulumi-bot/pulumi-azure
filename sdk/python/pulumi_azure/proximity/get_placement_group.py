# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetPlacementGroupResult',
    'AwaitableGetPlacementGroupResult',
    'get_placement_group',
]


class GetPlacementGroupResult:
    """
    A collection of values returned by getPlacementGroup.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, id=None, location=None, name=None, resource_group_name=None, tags=None) -> None:
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags


class AwaitableGetPlacementGroupResult(GetPlacementGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlacementGroupResult(
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_placement_group(name: Optional[str] = None, resource_group_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlacementGroupResult:
    """
    Use this data source to access information about an existing Proximity Placement Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.proximity.get_placement_group(name="tf-appsecuritygroup",
        resource_group_name="my-resource-group")
    pulumi.export("proximityPlacementGroupId", example.id)
    ```


    :param str name: The name of the Proximity Placement Group.
    :param str resource_group_name: The name of the resource group in which the Proximity Placement Group exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:proximity/getPlacementGroup:getPlacementGroup', __args__, opts=opts).value

    return AwaitableGetPlacementGroupResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'))
