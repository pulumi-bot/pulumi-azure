# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Domain(pulumi.CustomResource):
    endpoint: pulumi.Output[str]
    """
    The Endpoint associated with the EventGrid Domain.
    """
    input_mapping_default_values: pulumi.Output[dict]
    """
    A `input_mapping_default_values` block as defined below.

      * `dataVersion` (`str`) - Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `eventType` (`str`) - Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `subject` (`str`) - Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    input_mapping_fields: pulumi.Output[dict]
    """
    A `input_mapping_fields` block as defined below.

      * `dataVersion` (`str`) - Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `eventTime` (`str`) - Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `eventType` (`str`) - Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `id` (`str`) - Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `subject` (`str`) - Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
      * `topic` (`str`) - Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    input_schema: pulumi.Output[str]
    """
    Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
    """
    primary_access_key: pulumi.Output[str]
    """
    The Primary Shared Access Key associated with the EventGrid Domain.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
    """
    secondary_access_key: pulumi.Output[str]
    """
    The Secondary Shared Access Key associated with the EventGrid Domain.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, input_mapping_default_values=None, input_mapping_fields=None, input_schema=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an EventGrid Domain

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US 2")
        example_domain = azure.eventgrid.Domain("exampleDomain",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[dict] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **input_mapping_default_values** object supports the following:

          * `dataVersion` (`pulumi.Input[str]`) - Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `eventType` (`pulumi.Input[str]`) - Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `subject` (`pulumi.Input[str]`) - Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

        The **input_mapping_fields** object supports the following:

          * `dataVersion` (`pulumi.Input[str]`) - Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `eventTime` (`pulumi.Input[str]`) - Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `eventType` (`pulumi.Input[str]`) - Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `id` (`pulumi.Input[str]`) - Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `subject` (`pulumi.Input[str]`) - Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `topic` (`pulumi.Input[str]`) - Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
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
            opts.version = utilities.get_version()
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
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/domain:Domain")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Domain, __self__).__init__(
            'azure:eventgrid/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, endpoint=None, input_mapping_default_values=None, input_mapping_fields=None, input_schema=None, location=None, name=None, primary_access_key=None, resource_group_name=None, secondary_access_key=None, tags=None):
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The Endpoint associated with the EventGrid Domain.
        :param pulumi.Input[dict] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[dict] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The Primary Shared Access Key associated with the EventGrid Domain.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The Secondary Shared Access Key associated with the EventGrid Domain.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **input_mapping_default_values** object supports the following:

          * `dataVersion` (`pulumi.Input[str]`) - Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `eventType` (`pulumi.Input[str]`) - Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `subject` (`pulumi.Input[str]`) - Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

        The **input_mapping_fields** object supports the following:

          * `dataVersion` (`pulumi.Input[str]`) - Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `eventTime` (`pulumi.Input[str]`) - Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `eventType` (`pulumi.Input[str]`) - Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `id` (`pulumi.Input[str]`) - Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `subject` (`pulumi.Input[str]`) - Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
          * `topic` (`pulumi.Input[str]`) - Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
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
        return Domain(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
