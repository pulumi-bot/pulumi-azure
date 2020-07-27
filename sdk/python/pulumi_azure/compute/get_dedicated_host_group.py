# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetDedicatedHostGroupResult',
    'AwaitableGetDedicatedHostGroupResult',
    'get_dedicated_host_group',
]


class GetDedicatedHostGroupResult:
    """
    A collection of values returned by getDedicatedHostGroup.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, id=None, location=None, name=None, platform_fault_domain_count=None, resource_group_name=None, tags=None, zones=None) -> None:
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
        The Azure location where the Dedicated Host Group exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if platform_fault_domain_count and not isinstance(platform_fault_domain_count, float):
            raise TypeError("Expected argument 'platform_fault_domain_count' to be a float")
        __self__.platform_fault_domain_count = platform_fault_domain_count
        """
        The number of fault domains that the Dedicated Host Group spans.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones
        """
        The Availability Zones in which this Dedicated Host Group is located.
        """


class AwaitableGetDedicatedHostGroupResult(GetDedicatedHostGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDedicatedHostGroupResult(
            id=self.id,
            location=self.location,
            name=self.name,
            platform_fault_domain_count=self.platform_fault_domain_count,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            zones=self.zones)


def get_dedicated_host_group(name: Optional[str] = None, resource_group_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDedicatedHostGroupResult:
    """
    Use this data source to access information about an existing Dedicated Host Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.compute.get_dedicated_host_group(name="example-dedicated-host-group",
        resource_group_name="example-rg")
    pulumi.export("id", example.id)
    ```


    :param str name: Specifies the name of the Dedicated Host Group.
    :param str resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getDedicatedHostGroup:getDedicatedHostGroup', __args__, opts=opts).value

    return AwaitableGetDedicatedHostGroupResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        platform_fault_domain_count=__ret__.get('platformFaultDomainCount'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        zones=__ret__.get('zones'))
