# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, api_management_name=None, description=None, display_name=None, external_id=None, id=None, name=None, resource_group_name=None, type=None):
        if api_management_name and not isinstance(api_management_name, str):
            raise TypeError("Expected argument 'api_management_name' to be a str")
        __self__.api_management_name = api_management_name
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of this API Management Group.
        """
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        The display name of this API Management Group.
        """
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        __self__.external_id = external_id
        """
        The identifier of the external Group.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of this API Management Group, such as `custom` or `external`.
        """
class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            api_management_name=self.api_management_name,
            description=self.description,
            display_name=self.display_name,
            external_id=self.external_id,
            id=self.id,
            name=self.name,
            resource_group_name=self.resource_group_name,
            type=self.type)

def get_group(api_management_name=None,name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing API Management Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.apimanagement.get_group(name="my-group",
        api_management_name="example-apim",
        resource_group_name="search-service")
    pulumi.export("groupType", example.type)
    ```


    :param str api_management_name: The Name of the API Management Service in which this Group exists.
    :param str name: The Name of the API Management Group.
    :param str resource_group_name: The Name of the Resource Group in which the API Management Service exists.
    """
    __args__ = dict()


    __args__['apiManagementName'] = api_management_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:apimanagement/getGroup:getGroup', __args__, opts=opts).value

    return AwaitableGetGroupResult(
        api_management_name=__ret__.get('apiManagementName'),
        description=__ret__.get('description'),
        display_name=__ret__.get('displayName'),
        external_id=__ret__.get('externalId'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        type=__ret__.get('type'))
