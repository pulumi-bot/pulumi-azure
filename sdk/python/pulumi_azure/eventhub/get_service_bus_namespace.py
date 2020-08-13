# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetServiceBusNamespaceResult',
    'AwaitableGetServiceBusNamespaceResult',
    'get_service_bus_namespace',
]

warnings.warn("azure.eventhub.getServiceBusNamespace has been deprecated in favor of azure.servicebus.getNamespace", DeprecationWarning)

@pulumi.output_type
class _GetServiceBusNamespaceResult(dict):
    capacity: float = pulumi.property("capacity")
    default_primary_connection_string: str = pulumi.property("defaultPrimaryConnectionString")
    default_primary_key: str = pulumi.property("defaultPrimaryKey")
    default_secondary_connection_string: str = pulumi.property("defaultSecondaryConnectionString")
    default_secondary_key: str = pulumi.property("defaultSecondaryKey")
    id: str = pulumi.property("id")
    location: str = pulumi.property("location")
    name: str = pulumi.property("name")
    resource_group_name: str = pulumi.property("resourceGroupName")
    sku: str = pulumi.property("sku")
    tags: Mapping[str, str] = pulumi.property("tags")
    zone_redundant: bool = pulumi.property("zoneRedundant")


class GetServiceBusNamespaceResult:
    """
    A collection of values returned by getServiceBusNamespace.
    """
    def __init__(__self__, capacity=None, default_primary_connection_string=None, default_primary_key=None, default_secondary_connection_string=None, default_secondary_key=None, id=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, zone_redundant=None):
        if capacity and not isinstance(capacity, float):
            raise TypeError("Expected argument 'capacity' to be a float")
        __self__.capacity = capacity
        """
        The capacity of the ServiceBus Namespace.
        """
        if default_primary_connection_string and not isinstance(default_primary_connection_string, str):
            raise TypeError("Expected argument 'default_primary_connection_string' to be a str")
        __self__.default_primary_connection_string = default_primary_connection_string
        """
        The primary connection string for the authorization
        rule `RootManageSharedAccessKey`.
        """
        if default_primary_key and not isinstance(default_primary_key, str):
            raise TypeError("Expected argument 'default_primary_key' to be a str")
        __self__.default_primary_key = default_primary_key
        """
        The primary access key for the authorization rule `RootManageSharedAccessKey`.
        """
        if default_secondary_connection_string and not isinstance(default_secondary_connection_string, str):
            raise TypeError("Expected argument 'default_secondary_connection_string' to be a str")
        __self__.default_secondary_connection_string = default_secondary_connection_string
        """
        The secondary connection string for the
        authorization rule `RootManageSharedAccessKey`.
        """
        if default_secondary_key and not isinstance(default_secondary_key, str):
            raise TypeError("Expected argument 'default_secondary_key' to be a str")
        __self__.default_secondary_key = default_secondary_key
        """
        The secondary access key for the authorization rule `RootManageSharedAccessKey`.
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
        The location of the Resource Group in which the ServiceBus Namespace exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        __self__.sku = sku
        """
        The Tier used for the ServiceBus Namespace.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        __self__.zone_redundant = zone_redundant
        """
        Whether or not this ServiceBus Namespace is zone redundant.
        """


class AwaitableGetServiceBusNamespaceResult(GetServiceBusNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceBusNamespaceResult(
            capacity=self.capacity,
            default_primary_connection_string=self.default_primary_connection_string,
            default_primary_key=self.default_primary_key,
            default_secondary_connection_string=self.default_secondary_connection_string,
            default_secondary_key=self.default_secondary_key,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            sku=self.sku,
            tags=self.tags,
            zone_redundant=self.zone_redundant)


def get_service_bus_namespace(name: Optional[str] = None,
                              resource_group_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceBusNamespaceResult:
    """
    Use this data source to access information about an existing ServiceBus Namespace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.servicebus.get_namespace(name="examplenamespace",
        resource_group_name="example-resources")
    pulumi.export("location", example.location)
    ```


    :param str name: Specifies the name of the ServiceBus Namespace.
    :param str resource_group_name: Specifies the name of the Resource Group where the ServiceBus Namespace exists.
    """
    pulumi.log.warn("get_service_bus_namespace is deprecated: azure.eventhub.getServiceBusNamespace has been deprecated in favor of azure.servicebus.getNamespace")
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:eventhub/getServiceBusNamespace:getServiceBusNamespace', __args__, opts=opts, typ=_GetServiceBusNamespaceResult).value

    return AwaitableGetServiceBusNamespaceResult(
        capacity=_utilities.get_dict_value(__ret__, 'capacity'),
        default_primary_connection_string=_utilities.get_dict_value(__ret__, 'defaultPrimaryConnectionString'),
        default_primary_key=_utilities.get_dict_value(__ret__, 'defaultPrimaryKey'),
        default_secondary_connection_string=_utilities.get_dict_value(__ret__, 'defaultSecondaryConnectionString'),
        default_secondary_key=_utilities.get_dict_value(__ret__, 'defaultSecondaryKey'),
        id=_utilities.get_dict_value(__ret__, 'id'),
        location=_utilities.get_dict_value(__ret__, 'location'),
        name=_utilities.get_dict_value(__ret__, 'name'),
        resource_group_name=_utilities.get_dict_value(__ret__, 'resourceGroupName'),
        sku=_utilities.get_dict_value(__ret__, 'sku'),
        tags=_utilities.get_dict_value(__ret__, 'tags'),
        zone_redundant=_utilities.get_dict_value(__ret__, 'zoneRedundant'))
