# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetHubResult',
    'AwaitableGetHubResult',
    'get_hub',
]

@pulumi.output_type
class GetHubResult:
    """
    A collection of values returned by getHub.
    """
    def __init__(__self__, apns_credentials=None, gcm_credentials=None, id=None, location=None, name=None, namespace_name=None, resource_group_name=None, tags=None):
        if apns_credentials and not isinstance(apns_credentials, list):
            raise TypeError("Expected argument 'apns_credentials' to be a list")
        pulumi.set(__self__, "apns_credentials", apns_credentials)
        if gcm_credentials and not isinstance(gcm_credentials, list):
            raise TypeError("Expected argument 'gcm_credentials' to be a list")
        pulumi.set(__self__, "gcm_credentials", gcm_credentials)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        pulumi.set(__self__, "namespace_name", namespace_name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="apnsCredentials")
    def apns_credentials(self) -> Sequence['outputs.GetHubApnsCredentialResult']:
        """
        A `apns_credential` block as defined below.
        """
        return pulumi.get(self, "apns_credentials")

    @property
    @pulumi.getter(name="gcmCredentials")
    def gcm_credentials(self) -> Sequence['outputs.GetHubGcmCredentialResult']:
        """
        A `gcm_credential` block as defined below.
        """
        return pulumi.get(self, "gcm_credentials")

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
        The Azure Region in which this Notification Hub exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> str:
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")


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


def get_hub(name: Optional[str] = None,
            namespace_name: Optional[str] = None,
            resource_group_name: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHubResult:
    """
    Use this data source to access information about an existing Notification Hub within a Notification Hub Namespace.


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
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:notificationhub/getHub:getHub', __args__, opts=opts, typ=GetHubResult).value

    return AwaitableGetHubResult(
        apns_credentials=__ret__.apns_credentials,
        gcm_credentials=__ret__.gcm_credentials,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        namespace_name=__ret__.namespace_name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags)
