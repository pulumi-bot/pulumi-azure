# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class TriggerSchedule(pulumi.CustomResource):
    annotations: pulumi.Output[list]
    """
    List of tags that can be used for describing the Data Factory Schedule Trigger.
    """
    data_factory_name: pulumi.Output[str]
    """
    The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
    """
    end_time: pulumi.Output[str]
    """
    The time the Schedule Trigger should end. The time will be represented in UTC.
    """
    frequency: pulumi.Output[str]
    """
    The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
    """
    interval: pulumi.Output[float]
    """
    The interval for how often the trigger occurs. This defaults to 1.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
    """
    pipeline_name: pulumi.Output[str]
    """
    The Data Factory Pipeline name that the trigger will act on.
    """
    pipeline_parameters: pulumi.Output[dict]
    """
    The pipeline parameters that the trigger will act upon.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
    """
    start_time: pulumi.Output[str]
    """
    The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
    """
    def __init__(__self__, resource_name, opts=None, annotations=None, data_factory_name=None, end_time=None, frequency=None, interval=None, name=None, pipeline_name=None, pipeline_parameters=None, resource_group_name=None, start_time=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Trigger Schedule inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        test_pipeline = azure.datafactory.Pipeline("testPipeline",
            resource_group_name=azurerm_resource_group["test"]["name"],
            data_factory_name=azurerm_data_factory["test"]["name"])
        test_trigger_schedule = azure.datafactory.TriggerSchedule("testTriggerSchedule",
            data_factory_name=azurerm_data_factory["test"]["name"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            pipeline_name=test_pipeline.name,
            interval=5,
            frequency="Day")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] annotations: List of tags that can be used for describing the Data Factory Schedule Trigger.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        :param pulumi.Input[str] end_time: The time the Schedule Trigger should end. The time will be represented in UTC.
        :param pulumi.Input[str] frequency: The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        :param pulumi.Input[float] interval: The interval for how often the trigger occurs. This defaults to 1.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] pipeline_name: The Data Factory Pipeline name that the trigger will act on.
        :param pulumi.Input[dict] pipeline_parameters: The pipeline parameters that the trigger will act upon.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
        :param pulumi.Input[str] start_time: The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
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

            __props__['annotations'] = annotations
            if data_factory_name is None:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['end_time'] = end_time
            __props__['frequency'] = frequency
            __props__['interval'] = interval
            __props__['name'] = name
            if pipeline_name is None:
                raise TypeError("Missing required property 'pipeline_name'")
            __props__['pipeline_name'] = pipeline_name
            __props__['pipeline_parameters'] = pipeline_parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['start_time'] = start_time
        super(TriggerSchedule, __self__).__init__(
            'azure:datafactory/triggerSchedule:TriggerSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, annotations=None, data_factory_name=None, end_time=None, frequency=None, interval=None, name=None, pipeline_name=None, pipeline_parameters=None, resource_group_name=None, start_time=None):
        """
        Get an existing TriggerSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] annotations: List of tags that can be used for describing the Data Factory Schedule Trigger.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        :param pulumi.Input[str] end_time: The time the Schedule Trigger should end. The time will be represented in UTC.
        :param pulumi.Input[str] frequency: The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        :param pulumi.Input[float] interval: The interval for how often the trigger occurs. This defaults to 1.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] pipeline_name: The Data Factory Pipeline name that the trigger will act on.
        :param pulumi.Input[dict] pipeline_parameters: The pipeline parameters that the trigger will act upon.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
        :param pulumi.Input[str] start_time: The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["annotations"] = annotations
        __props__["data_factory_name"] = data_factory_name
        __props__["end_time"] = end_time
        __props__["frequency"] = frequency
        __props__["interval"] = interval
        __props__["name"] = name
        __props__["pipeline_name"] = pipeline_name
        __props__["pipeline_parameters"] = pipeline_parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["start_time"] = start_time
        return TriggerSchedule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
