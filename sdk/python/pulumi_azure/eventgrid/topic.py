# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['TopicArgs', 'Topic']

@pulumi.input_type
class TopicArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['TopicInboundIpRuleArgs']]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input['TopicInputMappingDefaultValuesArgs']] = None,
                 input_mapping_fields: Optional[pulumi.Input['TopicInputMappingFieldsArgs']] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Topic resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['TopicInboundIpRuleArgs']]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input['TopicInputMappingDefaultValuesArgs'] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input['TopicInputMappingFieldsArgs'] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if inbound_ip_rules is not None:
            pulumi.set(__self__, "inbound_ip_rules", inbound_ip_rules)
        if input_mapping_default_values is not None:
            pulumi.set(__self__, "input_mapping_default_values", input_mapping_default_values)
        if input_mapping_fields is not None:
            pulumi.set(__self__, "input_mapping_fields", input_mapping_fields)
        if input_schema is not None:
            pulumi.set(__self__, "input_schema", input_schema)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_network_access_enabled is not None:
            pulumi.set(__self__, "public_network_access_enabled", public_network_access_enabled)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TopicInboundIpRuleArgs']]]]:
        """
        One or more `inbound_ip_rule` blocks as defined below.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @inbound_ip_rules.setter
    def inbound_ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TopicInboundIpRuleArgs']]]]):
        pulumi.set(self, "inbound_ip_rules", value)

    @property
    @pulumi.getter(name="inputMappingDefaultValues")
    def input_mapping_default_values(self) -> Optional[pulumi.Input['TopicInputMappingDefaultValuesArgs']]:
        """
        A `input_mapping_default_values` block as defined below.
        """
        return pulumi.get(self, "input_mapping_default_values")

    @input_mapping_default_values.setter
    def input_mapping_default_values(self, value: Optional[pulumi.Input['TopicInputMappingDefaultValuesArgs']]):
        pulumi.set(self, "input_mapping_default_values", value)

    @property
    @pulumi.getter(name="inputMappingFields")
    def input_mapping_fields(self) -> Optional[pulumi.Input['TopicInputMappingFieldsArgs']]:
        """
        A `input_mapping_fields` block as defined below.
        """
        return pulumi.get(self, "input_mapping_fields")

    @input_mapping_fields.setter
    def input_mapping_fields(self, value: Optional[pulumi.Input['TopicInputMappingFieldsArgs']]):
        pulumi.set(self, "input_mapping_fields", value)

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "input_schema")

    @input_schema.setter
    def input_schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_schema", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not public network access is allowed for this server. Defaults to `true`.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @public_network_access_enabled.setter
    def public_network_access_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public_network_access_enabled", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Topic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicInboundIpRuleArgs']]]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['TopicInputMappingDefaultValuesArgs']]] = None,
                 input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['TopicInputMappingFieldsArgs']]] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an EventGrid Topic

        > **Note:** at this time EventGrid Topic's are only available in a limited number of regions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_topic = azure.eventgrid.Topic("exampleTopic",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        EventGrid Topic's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventgrid/topic:Topic topic1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicInboundIpRuleArgs']]]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['TopicInputMappingDefaultValuesArgs']] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[pulumi.InputType['TopicInputMappingFieldsArgs']] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an EventGrid Topic

        > **Note:** at this time EventGrid Topic's are only available in a limited number of regions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_topic = azure.eventgrid.Topic("exampleTopic",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        EventGrid Topic's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventgrid/topic:Topic topic1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
        ```

        :param str resource_name: The name of the resource.
        :param TopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicInboundIpRuleArgs']]]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['TopicInputMappingDefaultValuesArgs']]] = None,
                 input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['TopicInputMappingFieldsArgs']]] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['inbound_ip_rules'] = inbound_ip_rules
            __props__['input_mapping_default_values'] = input_mapping_default_values
            __props__['input_mapping_fields'] = input_mapping_fields
            __props__['input_schema'] = input_schema
            __props__['location'] = location
            __props__['name'] = name
            __props__['public_network_access_enabled'] = public_network_access_enabled
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['endpoint'] = None
            __props__['primary_access_key'] = None
            __props__['secondary_access_key'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/eventGridTopic:EventGridTopic")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Topic, __self__).__init__(
            'azure:eventgrid/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicInboundIpRuleArgs']]]]] = None,
            input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['TopicInputMappingDefaultValuesArgs']]] = None,
            input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['TopicInputMappingFieldsArgs']]] = None,
            input_schema: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            primary_access_key: Optional[pulumi.Input[str]] = None,
            public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_access_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The Endpoint associated with the EventGrid Topic.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicInboundIpRuleArgs']]]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['TopicInputMappingDefaultValuesArgs']] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[pulumi.InputType['TopicInputMappingFieldsArgs']] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The Primary Shared Access Key associated with the EventGrid Topic.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The Secondary Shared Access Key associated with the EventGrid Topic.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["endpoint"] = endpoint
        __props__["inbound_ip_rules"] = inbound_ip_rules
        __props__["input_mapping_default_values"] = input_mapping_default_values
        __props__["input_mapping_fields"] = input_mapping_fields
        __props__["input_schema"] = input_schema
        __props__["location"] = location
        __props__["name"] = name
        __props__["primary_access_key"] = primary_access_key
        __props__["public_network_access_enabled"] = public_network_access_enabled
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["tags"] = tags
        return Topic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The Endpoint associated with the EventGrid Topic.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> pulumi.Output[Optional[Sequence['outputs.TopicInboundIpRule']]]:
        """
        One or more `inbound_ip_rule` blocks as defined below.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @property
    @pulumi.getter(name="inputMappingDefaultValues")
    def input_mapping_default_values(self) -> pulumi.Output[Optional['outputs.TopicInputMappingDefaultValues']]:
        """
        A `input_mapping_default_values` block as defined below.
        """
        return pulumi.get(self, "input_mapping_default_values")

    @property
    @pulumi.getter(name="inputMappingFields")
    def input_mapping_fields(self) -> pulumi.Output[Optional['outputs.TopicInputMappingFields']]:
        """
        A `input_mapping_fields` block as defined below.
        """
        return pulumi.get(self, "input_mapping_fields")

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "input_schema")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> pulumi.Output[str]:
        """
        The Primary Shared Access Key associated with the EventGrid Topic.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not public network access is allowed for this server. Defaults to `true`.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> pulumi.Output[str]:
        """
        The Secondary Shared Access Key associated with the EventGrid Topic.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

