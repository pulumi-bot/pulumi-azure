# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetApiResult:
    """
    A collection of values returned by getApi.
    """
    def __init__(__self__, api_management_name=None, description=None, display_name=None, id=None, is_current=None, is_online=None, name=None, path=None, protocols=None, resource_group_name=None, revision=None, service_url=None, soap_pass_through=None, subscription_key_parameter_names=None, version=None, version_set_id=None):
        if api_management_name and not isinstance(api_management_name, str):
            raise TypeError("Expected argument 'api_management_name' to be a str")
        __self__.api_management_name = api_management_name
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A description of the API Management API, which may include HTML formatting tags.
        """
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        The display name of the API.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if is_current and not isinstance(is_current, bool):
            raise TypeError("Expected argument 'is_current' to be a bool")
        __self__.is_current = is_current
        """
        Is this the current API Revision?
        """
        if is_online and not isinstance(is_online, bool):
            raise TypeError("Expected argument 'is_online' to be a bool")
        __self__.is_online = is_online
        """
        Is this API Revision online/accessible via the Gateway?
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        """
        The Path for this API Management API.
        """
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        __self__.protocols = protocols
        """
        A list of protocols the operations in this API can be invoked.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if revision and not isinstance(revision, str):
            raise TypeError("Expected argument 'revision' to be a str")
        __self__.revision = revision
        if service_url and not isinstance(service_url, str):
            raise TypeError("Expected argument 'service_url' to be a str")
        __self__.service_url = service_url
        """
        Absolute URL of the backend service implementing this API.
        """
        if soap_pass_through and not isinstance(soap_pass_through, bool):
            raise TypeError("Expected argument 'soap_pass_through' to be a bool")
        __self__.soap_pass_through = soap_pass_through
        """
        Should this API expose a SOAP frontend, rather than a HTTP frontend?
        """
        if subscription_key_parameter_names and not isinstance(subscription_key_parameter_names, list):
            raise TypeError("Expected argument 'subscription_key_parameter_names' to be a list")
        __self__.subscription_key_parameter_names = subscription_key_parameter_names
        """
        A `subscription_key_parameter_names` block as documented below.
        """
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        __self__.version = version
        """
        The Version number of this API, if this API is versioned.
        """
        if version_set_id and not isinstance(version_set_id, str):
            raise TypeError("Expected argument 'version_set_id' to be a str")
        __self__.version_set_id = version_set_id
        """
        The ID of the Version Set which this API is associated with.
        """
class AwaitableGetApiResult(GetApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiResult(
            api_management_name=self.api_management_name,
            description=self.description,
            display_name=self.display_name,
            id=self.id,
            is_current=self.is_current,
            is_online=self.is_online,
            name=self.name,
            path=self.path,
            protocols=self.protocols,
            resource_group_name=self.resource_group_name,
            revision=self.revision,
            service_url=self.service_url,
            soap_pass_through=self.soap_pass_through,
            subscription_key_parameter_names=self.subscription_key_parameter_names,
            version=self.version,
            version_set_id=self.version_set_id)

def get_api(api_management_name=None,name=None,resource_group_name=None,revision=None,opts=None):
    """
    Use this data source to access information about an existing API Management API.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.apimanagement.get_api(name="search-api",
        api_management_name="search-api-management",
        resource_group_name="search-service",
        revision="2")
    pulumi.export("apiManagementApiId", example.id)
    ```


    :param str api_management_name: The name of the API Management Service in which the API Management API exists.
    :param str name: The name of the API Management API.
    :param str resource_group_name: The Name of the Resource Group in which the API Management Service exists.
    :param str revision: The Revision of the API Management API.
    """
    __args__ = dict()


    __args__['apiManagementName'] = api_management_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['revision'] = revision
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:apimanagement/getApi:getApi', __args__, opts=opts).value

    return AwaitableGetApiResult(
        api_management_name=__ret__.get('apiManagementName'),
        description=__ret__.get('description'),
        display_name=__ret__.get('displayName'),
        id=__ret__.get('id'),
        is_current=__ret__.get('isCurrent'),
        is_online=__ret__.get('isOnline'),
        name=__ret__.get('name'),
        path=__ret__.get('path'),
        protocols=__ret__.get('protocols'),
        resource_group_name=__ret__.get('resourceGroupName'),
        revision=__ret__.get('revision'),
        service_url=__ret__.get('serviceUrl'),
        soap_pass_through=__ret__.get('soapPassThrough'),
        subscription_key_parameter_names=__ret__.get('subscriptionKeyParameterNames'),
        version=__ret__.get('version'),
        version_set_id=__ret__.get('versionSetId'))
