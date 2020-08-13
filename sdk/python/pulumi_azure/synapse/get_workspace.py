# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetWorkspaceResult',
    'AwaitableGetWorkspaceResult',
    'get_workspace',
]


@pulumi.output_type
class _GetWorkspaceResult(dict):
    connectivity_endpoints: Mapping[str, str] = pulumi.property("connectivityEndpoints")
    id: str = pulumi.property("id")
    location: str = pulumi.property("location")
    name: str = pulumi.property("name")
    resource_group_name: str = pulumi.property("resourceGroupName")
    tags: Mapping[str, str] = pulumi.property("tags")


class GetWorkspaceResult:
    """
    A collection of values returned by getWorkspace.
    """
    def __init__(__self__, connectivity_endpoints=None, id=None, location=None, name=None, resource_group_name=None, tags=None):
        if connectivity_endpoints and not isinstance(connectivity_endpoints, dict):
            raise TypeError("Expected argument 'connectivity_endpoints' to be a dict")
        __self__.connectivity_endpoints = connectivity_endpoints
        """
        A list of Connectivity endpoints for this Synapse Workspace.
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
        The Azure location where the Synapse Workspace exists.
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


class AwaitableGetWorkspaceResult(GetWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceResult(
            connectivity_endpoints=self.connectivity_endpoints,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_workspace(name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceResult:
    """
    Use this data source to access information about an existing Synapse Workspace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.synapse.get_workspace(name="existing",
        resource_group_name="example-resource-group")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of this Synapse Workspace.
    :param str resource_group_name: The name of the Resource Group where the Synapse Workspace exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:synapse/getWorkspace:getWorkspace', __args__, opts=opts, typ=_GetWorkspaceResult).value

    return AwaitableGetWorkspaceResult(
        connectivity_endpoints=_utilities.get_dict_value(__ret__, 'connectivityEndpoints'),
        id=_utilities.get_dict_value(__ret__, 'id'),
        location=_utilities.get_dict_value(__ret__, 'location'),
        name=_utilities.get_dict_value(__ret__, 'name'),
        resource_group_name=_utilities.get_dict_value(__ret__, 'resourceGroupName'),
        tags=_utilities.get_dict_value(__ret__, 'tags'))
