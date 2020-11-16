# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['EventGridTopic']

warnings.warn("""azure.eventhub.EventGridTopic has been deprecated in favor of azure.eventgrid.Topic""", DeprecationWarning)


class EventGridTopic(pulumi.CustomResource):
    warnings.warn("""azure.eventhub.EventGridTopic has been deprecated in favor of azure.eventgrid.Topic""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['EventGridTopicInputMappingDefaultValuesArgs']]] = None,
                 input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['EventGridTopicInputMappingFieldsArgs']]] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US 2")
        example_topic = azure.eventgrid.Topic("exampleTopic",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventGridTopicInputMappingDefaultValuesArgs']] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventGridTopicInputMappingFieldsArgs']] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.log.warn("EventGridTopic is deprecated: azure.eventhub.EventGridTopic has been deprecated in favor of azure.eventgrid.Topic")
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

            __props__['input_mapping_default_values'] = input_mapping_default_values
            __props__['input_mapping_fields'] = input_mapping_fields
            __props__['input_schema'] = input_schema
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['endpoint'] = None
            __props__['primary_access_key'] = None
            __props__['secondary_access_key'] = None
        super(EventGridTopic, __self__).__init__(
            'azure:eventhub/eventGridTopic:EventGridTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['EventGridTopicInputMappingDefaultValuesArgs']]] = None,
            input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['EventGridTopicInputMappingFieldsArgs']]] = None,
            input_schema: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            primary_access_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_access_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'EventGridTopic':
        """
        Get an existing EventGridTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The Endpoint associated with the EventGrid Topic.
        :param pulumi.Input[pulumi.InputType['EventGridTopicInputMappingDefaultValuesArgs']] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventGridTopicInputMappingFieldsArgs']] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The Primary Shared Access Key associated with the EventGrid Topic.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The Secondary Shared Access Key associated with the EventGrid Topic.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["endpoint"] = endpoint
        __props__["input_mapping_default_values"] = input_mapping_default_values
        __props__["input_mapping_fields"] = input_mapping_fields
        __props__["input_schema"] = input_schema
        __props__["location"] = location
        __props__["name"] = name
        __props__["primary_access_key"] = primary_access_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["tags"] = tags
        return EventGridTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The Endpoint associated with the EventGrid Topic.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="inputMappingDefaultValues")
    def input_mapping_default_values(self) -> pulumi.Output[Optional['outputs.EventGridTopicInputMappingDefaultValues']]:
        """
        A `input_mapping_default_values` block as defined below.
        """
        return pulumi.get(self, "input_mapping_default_values")

    @property
    @pulumi.getter(name="inputMappingFields")
    def input_mapping_fields(self) -> pulumi.Output[Optional['outputs.EventGridTopicInputMappingFields']]:
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

