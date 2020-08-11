# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetNamespaceResult',
    'AwaitableGetNamespaceResult',
    'get_namespace',
]


class GetNamespaceResult:
    """
    A collection of values returned by getNamespace.
    """
    def __init__(__self__, auto_inflate_enabled=None, capacity=None, dedicated_cluster_id=None, default_primary_connection_string=None, default_primary_connection_string_alias=None, default_primary_key=None, default_secondary_connection_string=None, default_secondary_connection_string_alias=None, default_secondary_key=None, id=None, kafka_enabled=None, location=None, maximum_throughput_units=None, name=None, resource_group_name=None, sku=None, tags=None, zone_redundant=None):
        if auto_inflate_enabled and not isinstance(auto_inflate_enabled, bool):
            raise TypeError("Expected argument 'auto_inflate_enabled' to be a bool")
        __self__.auto_inflate_enabled = auto_inflate_enabled
        """
        Is Auto Inflate enabled for the EventHub Namespace?
        """
        if capacity and not isinstance(capacity, float):
            raise TypeError("Expected argument 'capacity' to be a float")
        __self__.capacity = capacity
        """
        The Capacity / Throughput Units for a `Standard` SKU namespace.
        """
        if dedicated_cluster_id and not isinstance(dedicated_cluster_id, str):
            raise TypeError("Expected argument 'dedicated_cluster_id' to be a str")
        __self__.dedicated_cluster_id = dedicated_cluster_id
        """
        The ID of the EventHub Dedicated Cluster where this Namespace exists.
        """
        if default_primary_connection_string and not isinstance(default_primary_connection_string, str):
            raise TypeError("Expected argument 'default_primary_connection_string' to be a str")
        __self__.default_primary_connection_string = default_primary_connection_string
        """
        The primary connection string for the authorization
        rule `RootManageSharedAccessKey`.
        """
        if default_primary_connection_string_alias and not isinstance(default_primary_connection_string_alias, str):
            raise TypeError("Expected argument 'default_primary_connection_string_alias' to be a str")
        __self__.default_primary_connection_string_alias = default_primary_connection_string_alias
        """
        The alias of the primary connection string for the authorization
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
        if default_secondary_connection_string_alias and not isinstance(default_secondary_connection_string_alias, str):
            raise TypeError("Expected argument 'default_secondary_connection_string_alias' to be a str")
        __self__.default_secondary_connection_string_alias = default_secondary_connection_string_alias
        """
        The alias of the secondary connection string for the
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
        if kafka_enabled and not isinstance(kafka_enabled, bool):
            raise TypeError("Expected argument 'kafka_enabled' to be a bool")
        __self__.kafka_enabled = kafka_enabled
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure location where the EventHub Namespace exists
        """
        if maximum_throughput_units and not isinstance(maximum_throughput_units, float):
            raise TypeError("Expected argument 'maximum_throughput_units' to be a float")
        __self__.maximum_throughput_units = maximum_throughput_units
        """
        Specifies the maximum number of throughput units when Auto Inflate is Enabled.
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
        Defines which tier to use.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the EventHub Namespace.
        """
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        __self__.zone_redundant = zone_redundant
        """
        Is this EventHub Namespace deployed across Availability Zones?
        """


class AwaitableGetNamespaceResult(GetNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNamespaceResult(
            auto_inflate_enabled=self.auto_inflate_enabled,
            capacity=self.capacity,
            dedicated_cluster_id=self.dedicated_cluster_id,
            default_primary_connection_string=self.default_primary_connection_string,
            default_primary_connection_string_alias=self.default_primary_connection_string_alias,
            default_primary_key=self.default_primary_key,
            default_secondary_connection_string=self.default_secondary_connection_string,
            default_secondary_connection_string_alias=self.default_secondary_connection_string_alias,
            default_secondary_key=self.default_secondary_key,
            id=self.id,
            kafka_enabled=self.kafka_enabled,
            location=self.location,
            maximum_throughput_units=self.maximum_throughput_units,
            name=self.name,
            resource_group_name=self.resource_group_name,
            sku=self.sku,
            tags=self.tags,
            zone_redundant=self.zone_redundant)


def get_namespace(name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNamespaceResult:
    """
    Use this data source to access information about an existing EventHub Namespace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.eventhub.get_namespace(name="search-eventhubns",
        resource_group_name="search-service")
    pulumi.export("eventhubNamespaceId", example.id)
    ```


    :param str name: The name of the EventHub Namespace.
    :param str resource_group_name: The Name of the Resource Group where the EventHub Namespace exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:eventhub/getNamespace:getNamespace', __args__, opts=opts).value

    return AwaitableGetNamespaceResult(
        auto_inflate_enabled=__ret__.get('autoInflateEnabled'),
        capacity=__ret__.get('capacity'),
        dedicated_cluster_id=__ret__.get('dedicatedClusterId'),
        default_primary_connection_string=__ret__.get('defaultPrimaryConnectionString'),
        default_primary_connection_string_alias=__ret__.get('defaultPrimaryConnectionStringAlias'),
        default_primary_key=__ret__.get('defaultPrimaryKey'),
        default_secondary_connection_string=__ret__.get('defaultSecondaryConnectionString'),
        default_secondary_connection_string_alias=__ret__.get('defaultSecondaryConnectionStringAlias'),
        default_secondary_key=__ret__.get('defaultSecondaryKey'),
        id=__ret__.get('id'),
        kafka_enabled=__ret__.get('kafkaEnabled'),
        location=__ret__.get('location'),
        maximum_throughput_units=__ret__.get('maximumThroughputUnits'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        zone_redundant=__ret__.get('zoneRedundant'))
