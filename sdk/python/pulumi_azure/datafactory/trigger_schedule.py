# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['TriggerScheduleArgs', 'TriggerSchedule']

@pulumi.input_type
class TriggerScheduleArgs:
    def __init__(__self__, *,
                 data_factory_name: pulumi.Input[str],
                 pipeline_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pipeline_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TriggerSchedule resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        :param pulumi.Input[str] pipeline_name: The Data Factory Pipeline name that the trigger will act on.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Schedule Trigger.
        :param pulumi.Input[str] end_time: The time the Schedule Trigger should end. The time will be represented in UTC.
        :param pulumi.Input[str] frequency: The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        :param pulumi.Input[int] interval: The interval for how often the trigger occurs. This defaults to 1.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pipeline_parameters: The pipeline parameters that the trigger will act upon.
        :param pulumi.Input[str] start_time: The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
        """
        pulumi.set(__self__, "data_factory_name", data_factory_name)
        pulumi.set(__self__, "pipeline_name", pipeline_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if frequency is not None:
            pulumi.set(__self__, "frequency", frequency)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pipeline_parameters is not None:
            pulumi.set(__self__, "pipeline_parameters", pipeline_parameters)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Input[str]:
        """
        The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter(name="pipelineName")
    def pipeline_name(self) -> pulumi.Input[str]:
        """
        The Data Factory Pipeline name that the trigger will act on.
        """
        return pulumi.get(self, "pipeline_name")

    @pipeline_name.setter
    def pipeline_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Schedule Trigger.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time the Schedule Trigger should end. The time will be represented in UTC.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def frequency(self) -> Optional[pulumi.Input[str]]:
        """
        The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The interval for how often the trigger occurs. This defaults to 1.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pipelineParameters")
    def pipeline_parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The pipeline parameters that the trigger will act upon.
        """
        return pulumi.get(self, "pipeline_parameters")

    @pipeline_parameters.setter
    def pipeline_parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "pipeline_parameters", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)


class TriggerSchedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Trigger Schedule inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
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

        ## Import

        Data Factory Schedule Trigger can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/triggerSchedule:TriggerSchedule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
        ```

        :param str resource_name: The name of the resource.
        :param TriggerScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pipeline_name: Optional[pulumi.Input[str]] = None,
                 pipeline_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Trigger Schedule inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
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

        ## Import

        Data Factory Schedule Trigger can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/triggerSchedule:TriggerSchedule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Schedule Trigger.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        :param pulumi.Input[str] end_time: The time the Schedule Trigger should end. The time will be represented in UTC.
        :param pulumi.Input[str] frequency: The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        :param pulumi.Input[int] interval: The interval for how often the trigger occurs. This defaults to 1.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] pipeline_name: The Data Factory Pipeline name that the trigger will act on.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pipeline_parameters: The pipeline parameters that the trigger will act upon.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
        :param pulumi.Input[str] start_time: The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pipeline_name: Optional[pulumi.Input[str]] = None,
                 pipeline_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
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

            __props__['annotations'] = annotations
            if data_factory_name is None and not opts.urn:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['end_time'] = end_time
            __props__['frequency'] = frequency
            __props__['interval'] = interval
            __props__['name'] = name
            if pipeline_name is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_name'")
            __props__['pipeline_name'] = pipeline_name
            __props__['pipeline_parameters'] = pipeline_parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['start_time'] = start_time
        super(TriggerSchedule, __self__).__init__(
            'azure:datafactory/triggerSchedule:TriggerSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            frequency: Optional[pulumi.Input[str]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pipeline_name: Optional[pulumi.Input[str]] = None,
            pipeline_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[str]] = None) -> 'TriggerSchedule':
        """
        Get an existing TriggerSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Schedule Trigger.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        :param pulumi.Input[str] end_time: The time the Schedule Trigger should end. The time will be represented in UTC.
        :param pulumi.Input[str] frequency: The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        :param pulumi.Input[int] interval: The interval for how often the trigger occurs. This defaults to 1.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] pipeline_name: The Data Factory Pipeline name that the trigger will act on.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pipeline_parameters: The pipeline parameters that the trigger will act upon.
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

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Schedule Trigger.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Output[str]:
        """
        The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[str]]:
        """
        The time the Schedule Trigger should end. The time will be represented in UTC.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output[Optional[str]]:
        """
        The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[Optional[int]]:
        """
        The interval for how often the trigger occurs. This defaults to 1.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pipelineName")
    def pipeline_name(self) -> pulumi.Output[str]:
        """
        The Data Factory Pipeline name that the trigger will act on.
        """
        return pulumi.get(self, "pipeline_name")

    @property
    @pulumi.getter(name="pipelineParameters")
    def pipeline_parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The pipeline parameters that the trigger will act upon.
        """
        return pulumi.get(self, "pipeline_parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
        """
        return pulumi.get(self, "start_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

