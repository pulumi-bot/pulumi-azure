# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Namespace']


class Namespace(pulumi.CustomResource):
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
    """
    metric_id: pulumi.Output[str] = pulumi.output_property("metricId")
    """
    The Identifier for Azure Insights metrics.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
    """
    primary_connection_string: pulumi.Output[str] = pulumi.output_property("primaryConnectionString")
    """
    The primary connection string for the authorization rule `RootManageSharedAccessKey`.
    """
    primary_key: pulumi.Output[str] = pulumi.output_property("primaryKey")
    """
    The primary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the Azure Relay Namespace.
    """
    secondary_connection_string: pulumi.Output[str] = pulumi.output_property("secondaryConnectionString")
    """
    The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
    """
    secondary_key: pulumi.Output[str] = pulumi.output_property("secondaryKey")
    """
    The secondary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    sku_name: pulumi.Output[str] = pulumi.output_property("skuName")
    """
    The name of the SKU to use. At this time the only supported value is `Standard`.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, sku_name: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an Azure Relay Namespace.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.relay.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Standard",
            tags={
                "source": "example",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Namespace.
        :param pulumi.Input[str] sku_name: The name of the SKU to use. At this time the only supported value is `Standard`.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['metric_id'] = None
            __props__['primary_connection_string'] = None
            __props__['primary_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['secondary_key'] = None
        super(Namespace, __self__).__init__(
            'azure:relay/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, location: Optional[pulumi.Input[str]] = None, metric_id: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, primary_connection_string: Optional[pulumi.Input[str]] = None, primary_key: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, secondary_connection_string: Optional[pulumi.Input[str]] = None, secondary_key: Optional[pulumi.Input[str]] = None, sku_name: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None) -> 'Namespace':
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metric_id: The Identifier for Azure Insights metrics.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The primary connection string for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] primary_key: The primary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Namespace.
        :param pulumi.Input[str] secondary_connection_string: The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] secondary_key: The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] sku_name: The name of the SKU to use. At this time the only supported value is `Standard`.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["metric_id"] = metric_id
        __props__["name"] = name
        __props__["primary_connection_string"] = primary_connection_string
        __props__["primary_key"] = primary_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["secondary_key"] = secondary_key
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        return Namespace(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

