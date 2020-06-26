# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetResourcesResult:
    """
    A collection of values returned by getResources.
    """
    def __init__(__self__, id=None, name=None, required_tags=None, resource_group_name=None, resources=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of this Resource.
        """
        if required_tags and not isinstance(required_tags, dict):
            raise TypeError("Expected argument 'required_tags' to be a dict")
        __self__.required_tags = required_tags
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        __self__.resources = resources
        """
        One or more `resource` blocks as defined below.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of this Resource. (e.g. `Microsoft.Network/virtualNetworks`).
        """
class AwaitableGetResourcesResult(GetResourcesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourcesResult(
            id=self.id,
            name=self.name,
            required_tags=self.required_tags,
            resource_group_name=self.resource_group_name,
            resources=self.resources,
            type=self.type)

def get_resources(name=None,required_tags=None,resource_group_name=None,type=None,opts=None):
    """
    Use this data source to access information about existing resources.


    :param str name: The name of the Resource.
    :param dict required_tags: A mapping of tags which the resource has to have in order to be included in the result.
    :param str resource_group_name: The name of the Resource group where the Resources are located.
    :param str type: The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['requiredTags'] = required_tags
    __args__['resourceGroupName'] = resource_group_name
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:core/getResources:getResources', __args__, opts=opts).value

    return AwaitableGetResourcesResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        required_tags=__ret__.get('requiredTags'),
        resource_group_name=__ret__.get('resourceGroupName'),
        resources=__ret__.get('resources'),
        type=__ret__.get('type'))
