# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetNamespaceResult',
    'AwaitableGetNamespaceResult',
    'get_namespace',
]

@pulumi.output_type
class GetNamespaceResult:
    """
    A collection of values returned by getNamespace.
    """
    def __init__(__self__, enabled=None, id=None, location=None, name=None, namespace_type=None, resource_group_name=None, servicebus_endpoint=None, sku=None, tags=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_type and not isinstance(namespace_type, str):
            raise TypeError("Expected argument 'namespace_type' to be a str")
        pulumi.set(__self__, "namespace_type", namespace_type)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if servicebus_endpoint and not isinstance(servicebus_endpoint, str):
            raise TypeError("Expected argument 'servicebus_endpoint' to be a str")
        pulumi.set(__self__, "servicebus_endpoint", servicebus_endpoint)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Is this Notification Hub Namespace enabled?
        """
        return pulumi.get(self, "enabled")

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
        The Azure Region in which this Notification Hub Namespace exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard.`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceType")
    def namespace_type(self) -> str:
        """
        The Type of Namespace, such as `Messaging` or `NotificationHub`.
        """
        return pulumi.get(self, "namespace_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="servicebusEndpoint")
    def servicebus_endpoint(self) -> str:
        return pulumi.get(self, "servicebus_endpoint")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.GetNamespaceSkuResult':
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetNamespaceResult(GetNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNamespaceResult(
            enabled=self.enabled,
            id=self.id,
            location=self.location,
            name=self.name,
            namespace_type=self.namespace_type,
            resource_group_name=self.resource_group_name,
            servicebus_endpoint=self.servicebus_endpoint,
            sku=self.sku,
            tags=self.tags)


def get_namespace(name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNamespaceResult:
    """
    Use this data source to access information about an existing Notification Hub Namespace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.notificationhub.get_namespace(name="my-namespace",
        resource_group_name="my-resource-group")
    pulumi.export("servicebusEndpoint", example.servicebus_endpoint)
    ```


    :param str name: Specifies the Name of the Notification Hub Namespace.
    :param str resource_group_name: Specifies the Name of the Resource Group within which the Notification Hub exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:notificationhub/getNamespace:getNamespace', __args__, opts=opts, typ=GetNamespaceResult).value

    return AwaitableGetNamespaceResult(
        enabled=__ret__.enabled,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        namespace_type=__ret__.namespace_type,
        resource_group_name=__ret__.resource_group_name,
        servicebus_endpoint=__ret__.servicebus_endpoint,
        sku=__ret__.sku,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_namespace)
def get_namespace_apply(name: Optional[pulumi.Input[str]] = None,
                        resource_group_name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNamespaceResult]:
    ...
