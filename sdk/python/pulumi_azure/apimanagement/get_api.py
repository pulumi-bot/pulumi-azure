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
    'GetApiResult',
    'AwaitableGetApiResult',
    'get_api',
]

@pulumi.output_type
class GetApiResult:
    """
    A collection of values returned by getApi.
    """
    def __init__(__self__, api_management_name=None, description=None, display_name=None, id=None, is_current=None, is_online=None, name=None, path=None, protocols=None, resource_group_name=None, revision=None, service_url=None, soap_pass_through=None, subscription_key_parameter_names=None, subscription_required=None, version=None, version_set_id=None):
        if api_management_name and not isinstance(api_management_name, str):
            raise TypeError("Expected argument 'api_management_name' to be a str")
        pulumi.set(__self__, "api_management_name", api_management_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_current and not isinstance(is_current, bool):
            raise TypeError("Expected argument 'is_current' to be a bool")
        pulumi.set(__self__, "is_current", is_current)
        if is_online and not isinstance(is_online, bool):
            raise TypeError("Expected argument 'is_online' to be a bool")
        pulumi.set(__self__, "is_online", is_online)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if revision and not isinstance(revision, str):
            raise TypeError("Expected argument 'revision' to be a str")
        pulumi.set(__self__, "revision", revision)
        if service_url and not isinstance(service_url, str):
            raise TypeError("Expected argument 'service_url' to be a str")
        pulumi.set(__self__, "service_url", service_url)
        if soap_pass_through and not isinstance(soap_pass_through, bool):
            raise TypeError("Expected argument 'soap_pass_through' to be a bool")
        pulumi.set(__self__, "soap_pass_through", soap_pass_through)
        if subscription_key_parameter_names and not isinstance(subscription_key_parameter_names, list):
            raise TypeError("Expected argument 'subscription_key_parameter_names' to be a list")
        pulumi.set(__self__, "subscription_key_parameter_names", subscription_key_parameter_names)
        if subscription_required and not isinstance(subscription_required, bool):
            raise TypeError("Expected argument 'subscription_required' to be a bool")
        pulumi.set(__self__, "subscription_required", subscription_required)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if version_set_id and not isinstance(version_set_id, str):
            raise TypeError("Expected argument 'version_set_id' to be a str")
        pulumi.set(__self__, "version_set_id", version_set_id)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> str:
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the API Management API, which may include HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of the API.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isCurrent")
    def is_current(self) -> bool:
        """
        Is this the current API Revision?
        """
        return pulumi.get(self, "is_current")

    @property
    @pulumi.getter(name="isOnline")
    def is_online(self) -> bool:
        """
        Is this API Revision online/accessible via the Gateway?
        """
        return pulumi.get(self, "is_online")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The Path for this API Management API.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def protocols(self) -> Sequence[str]:
        """
        A list of protocols the operations in this API can be invoked.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def revision(self) -> str:
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter(name="serviceUrl")
    def service_url(self) -> str:
        """
        Absolute URL of the backend service implementing this API.
        """
        return pulumi.get(self, "service_url")

    @property
    @pulumi.getter(name="soapPassThrough")
    def soap_pass_through(self) -> bool:
        """
        Should this API expose a SOAP frontend, rather than a HTTP frontend?
        """
        return pulumi.get(self, "soap_pass_through")

    @property
    @pulumi.getter(name="subscriptionKeyParameterNames")
    def subscription_key_parameter_names(self) -> Sequence['outputs.GetApiSubscriptionKeyParameterNameResult']:
        """
        A `subscription_key_parameter_names` block as documented below.
        """
        return pulumi.get(self, "subscription_key_parameter_names")

    @property
    @pulumi.getter(name="subscriptionRequired")
    def subscription_required(self) -> bool:
        """
        Should this API require a subscription key?
        """
        return pulumi.get(self, "subscription_required")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The Version number of this API, if this API is versioned.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionSetId")
    def version_set_id(self) -> str:
        """
        The ID of the Version Set which this API is associated with.
        """
        return pulumi.get(self, "version_set_id")


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
            subscription_required=self.subscription_required,
            version=self.version,
            version_set_id=self.version_set_id)


def get_api(api_management_name: Optional[str] = None,
            name: Optional[str] = None,
            resource_group_name: Optional[str] = None,
            revision: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiResult:
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
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:apimanagement/getApi:getApi', __args__, opts=opts, typ=GetApiResult).value

    return AwaitableGetApiResult(
        api_management_name=__ret__.api_management_name,
        description=__ret__.description,
        display_name=__ret__.display_name,
        id=__ret__.id,
        is_current=__ret__.is_current,
        is_online=__ret__.is_online,
        name=__ret__.name,
        path=__ret__.path,
        protocols=__ret__.protocols,
        resource_group_name=__ret__.resource_group_name,
        revision=__ret__.revision,
        service_url=__ret__.service_url,
        soap_pass_through=__ret__.soap_pass_through,
        subscription_key_parameter_names=__ret__.subscription_key_parameter_names,
        subscription_required=__ret__.subscription_required,
        version=__ret__.version,
        version_set_id=__ret__.version_set_id)
