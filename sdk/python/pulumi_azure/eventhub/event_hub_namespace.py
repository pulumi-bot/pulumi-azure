# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['EventHubNamespace']


class EventHubNamespace(pulumi.CustomResource):
    auto_inflate_enabled: pulumi.Output[Optional[bool]] = pulumi.output_property("autoInflateEnabled")
    """
    Is Auto Inflate enabled for the EventHub Namespace?
    """
    capacity: pulumi.Output[Optional[float]] = pulumi.output_property("capacity")
    """
    Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Valid values range from `1` - `20`.
    """
    dedicated_cluster_id: pulumi.Output[Optional[str]] = pulumi.output_property("dedicatedClusterId")
    """
    Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
    """
    default_primary_connection_string: pulumi.Output[str] = pulumi.output_property("defaultPrimaryConnectionString")
    """
    The primary connection string for the authorization
    rule `RootManageSharedAccessKey`.
    """
    default_primary_connection_string_alias: pulumi.Output[str] = pulumi.output_property("defaultPrimaryConnectionStringAlias")
    """
    The alias of the primary connection string for the authorization
    rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
    """
    default_primary_key: pulumi.Output[str] = pulumi.output_property("defaultPrimaryKey")
    """
    The primary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    default_secondary_connection_string: pulumi.Output[str] = pulumi.output_property("defaultSecondaryConnectionString")
    """
    The secondary connection string for the
    authorization rule `RootManageSharedAccessKey`.
    """
    default_secondary_connection_string_alias: pulumi.Output[str] = pulumi.output_property("defaultSecondaryConnectionStringAlias")
    """
    The alias of the secondary connection string for the
    authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
    """
    default_secondary_key: pulumi.Output[str] = pulumi.output_property("defaultSecondaryKey")
    """
    The secondary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    maximum_throughput_units: pulumi.Output[float] = pulumi.output_property("maximumThroughputUnits")
    """
    Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
    """
    network_rulesets: pulumi.Output['outputs.EventHubNamespaceNetworkRulesets'] = pulumi.output_property("networkRulesets")
    """
    A `network_rulesets` block as defined below.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[str] = pulumi.output_property("sku")
    """
    Defines which tier to use. Valid options are `Basic` and `Standard`.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    zone_redundant: pulumi.Output[Optional[bool]] = pulumi.output_property("zoneRedundant")
    """
    Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, auto_inflate_enabled: Optional[pulumi.Input[bool]] = None, capacity: Optional[pulumi.Input[float]] = None, dedicated_cluster_id: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, maximum_throughput_units: Optional[pulumi.Input[float]] = None, name: Optional[pulumi.Input[str]] = None, network_rulesets: Optional[pulumi.Input[pulumi.InputType['EventHubNamespaceNetworkRulesetsArgs']]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, sku: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, zone_redundant: Optional[pulumi.Input[bool]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an EventHub Namespace.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            capacity=2,
            tags={
                "environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_inflate_enabled: Is Auto Inflate enabled for the EventHub Namespace?
        :param pulumi.Input[float] capacity: Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Valid values range from `1` - `20`.
        :param pulumi.Input[str] dedicated_cluster_id: Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_throughput_units: Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['EventHubNamespaceNetworkRulesetsArgs']] network_rulesets: A `network_rulesets` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: Defines which tier to use. Valid options are `Basic` and `Standard`.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] zone_redundant: Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_inflate_enabled'] = auto_inflate_enabled
            __props__['capacity'] = capacity
            __props__['dedicated_cluster_id'] = dedicated_cluster_id
            __props__['location'] = location
            __props__['maximum_throughput_units'] = maximum_throughput_units
            __props__['name'] = name
            __props__['network_rulesets'] = network_rulesets
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zone_redundant'] = zone_redundant
            __props__['default_primary_connection_string'] = None
            __props__['default_primary_connection_string_alias'] = None
            __props__['default_primary_key'] = None
            __props__['default_secondary_connection_string'] = None
            __props__['default_secondary_connection_string_alias'] = None
            __props__['default_secondary_key'] = None
        super(EventHubNamespace, __self__).__init__(
            'azure:eventhub/eventHubNamespace:EventHubNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, auto_inflate_enabled: Optional[pulumi.Input[bool]] = None, capacity: Optional[pulumi.Input[float]] = None, dedicated_cluster_id: Optional[pulumi.Input[str]] = None, default_primary_connection_string: Optional[pulumi.Input[str]] = None, default_primary_connection_string_alias: Optional[pulumi.Input[str]] = None, default_primary_key: Optional[pulumi.Input[str]] = None, default_secondary_connection_string: Optional[pulumi.Input[str]] = None, default_secondary_connection_string_alias: Optional[pulumi.Input[str]] = None, default_secondary_key: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, maximum_throughput_units: Optional[pulumi.Input[float]] = None, name: Optional[pulumi.Input[str]] = None, network_rulesets: Optional[pulumi.Input[pulumi.InputType['EventHubNamespaceNetworkRulesetsArgs']]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, sku: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, zone_redundant: Optional[pulumi.Input[bool]] = None) -> 'EventHubNamespace':
        """
        Get an existing EventHubNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_inflate_enabled: Is Auto Inflate enabled for the EventHub Namespace?
        :param pulumi.Input[float] capacity: Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Valid values range from `1` - `20`.
        :param pulumi.Input[str] dedicated_cluster_id: Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] default_primary_connection_string: The primary connection string for the authorization
               rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_primary_connection_string_alias: The alias of the primary connection string for the authorization
               rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] default_primary_key: The primary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_secondary_connection_string: The secondary connection string for the
               authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_secondary_connection_string_alias: The alias of the secondary connection string for the
               authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] default_secondary_key: The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_throughput_units: Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['EventHubNamespaceNetworkRulesetsArgs']] network_rulesets: A `network_rulesets` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: Defines which tier to use. Valid options are `Basic` and `Standard`.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] zone_redundant: Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_inflate_enabled"] = auto_inflate_enabled
        __props__["capacity"] = capacity
        __props__["dedicated_cluster_id"] = dedicated_cluster_id
        __props__["default_primary_connection_string"] = default_primary_connection_string
        __props__["default_primary_connection_string_alias"] = default_primary_connection_string_alias
        __props__["default_primary_key"] = default_primary_key
        __props__["default_secondary_connection_string"] = default_secondary_connection_string
        __props__["default_secondary_connection_string_alias"] = default_secondary_connection_string_alias
        __props__["default_secondary_key"] = default_secondary_key
        __props__["location"] = location
        __props__["maximum_throughput_units"] = maximum_throughput_units
        __props__["name"] = name
        __props__["network_rulesets"] = network_rulesets
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["zone_redundant"] = zone_redundant
        return EventHubNamespace(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

