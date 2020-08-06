# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetNetworkWatcherResult:
    """
    A collection of values returned by getNetworkWatcher.
    """
    def __init__(__self__, id=None, location=None, name=None, resource_group_name=None, tags=None):
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
        The supported Azure location where the resource exists.
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
        A mapping of tags assigned to the resource.
        """


class AwaitableGetNetworkWatcherResult(GetNetworkWatcherResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkWatcherResult(
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_network_watcher(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing Network Watcher.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_network_watcher(name=azurerm_network_watcher["example"]["name"],
        resource_group_name=azurerm_resource_group["example"]["name"])
    pulumi.export("networkWatcherId", example.id)
    ```


    :param str name: Specifies the Name of the Network Watcher.
    :param str resource_group_name: Specifies the Name of the Resource Group within which the Network Watcher exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getNetworkWatcher:getNetworkWatcher', __args__, opts=opts).value

    return AwaitableGetNetworkWatcherResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'))
