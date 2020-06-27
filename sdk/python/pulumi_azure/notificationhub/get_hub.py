# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetHubResult:
    """
    A collection of values returned by getHub.
    """
    def __init__(__self__, apns_credentials=None, gcm_credentials=None, id=None, location=None, name=None, namespace_name=None, resource_group_name=None, tags=None):
        if apns_credentials and not isinstance(apns_credentials, list):
            raise TypeError("Expected argument 'apns_credentials' to be a list")
        __self__.apns_credentials = apns_credentials
        """
        A `apns_credential` block as defined below.
        """
        if gcm_credentials and not isinstance(gcm_credentials, list):
            raise TypeError("Expected argument 'gcm_credentials' to be a list")
        __self__.gcm_credentials = gcm_credentials
        """
        A `gcm_credential` block as defined below.
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
        The Azure Region in which this Notification Hub exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        __self__.namespace_name = namespace_name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
class AwaitableGetHubResult(GetHubResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHubResult(
            apns_credentials=self.apns_credentials,
            gcm_credentials=self.gcm_credentials,
            id=self.id,
            location=self.location,
            name=self.name,
            namespace_name=self.namespace_name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)

def get_hub(name=None,namespace_name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing Notification Hub within a Notification Hub Namespace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.notificationhub.get_hub(name="notification-hub",
        namespace_name="namespace-name",
        resource_group_name="resource-group-name")
    pulumi.export("id", example.id)
    ```


    :param str name: Specifies the Name of the Notification Hub.
    :param str namespace_name: Specifies the Name of the Notification Hub Namespace which contains the Notification Hub.
    :param str resource_group_name: Specifies the Name of the Resource Group within which the Notification Hub exists.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:notificationhub/getHub:getHub', __args__, opts=opts).value

    return AwaitableGetHubResult(
        apns_credentials=__ret__.get('apnsCredentials'),
        gcm_credentials=__ret__.get('gcmCredentials'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        namespace_name=__ret__.get('namespaceName'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'))
