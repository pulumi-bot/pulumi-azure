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

__all__ = ['TimeSeriesInsightsReferenceDataSet']


class TimeSeriesInsightsReferenceDataSet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_string_comparison_behavior: Optional[pulumi.Input[str]] = None,
                 key_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TimeSeriesInsightsReferenceDataSetKeyPropertyArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 time_series_insights_environment_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure IoT Time Series Insights Reference Data Set.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_time_series_insights_standard_environment = azure.iot.TimeSeriesInsightsStandardEnvironment("exampleTimeSeriesInsightsStandardEnvironment",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="S1_1",
            data_retention_time="P30D")
        example_time_series_insights_reference_data_set = azure.iot.TimeSeriesInsightsReferenceDataSet("exampleTimeSeriesInsightsReferenceDataSet",
            time_series_insights_environment_id=example_time_series_insights_standard_environment.id,
            location=example_resource_group.location,
            key_properties=[azure.iot.TimeSeriesInsightsReferenceDataSetKeyPropertyArgs(
                name="keyProperty1",
                type="String",
            )])
        ```

        ## Import

        Azure IoT Time Series Insights Reference Data Set can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example/referenceDataSets/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_string_comparison_behavior: The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TimeSeriesInsightsReferenceDataSetKeyPropertyArgs']]]] key_properties: A `key_property` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] time_series_insights_environment_id: The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
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

            __props__['data_string_comparison_behavior'] = data_string_comparison_behavior
            if key_properties is None and not opts.urn:
                raise TypeError("Missing required property 'key_properties'")
            __props__['key_properties'] = key_properties
            __props__['location'] = location
            __props__['name'] = name
            __props__['tags'] = tags
            if time_series_insights_environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'time_series_insights_environment_id'")
            __props__['time_series_insights_environment_id'] = time_series_insights_environment_id
        super(TimeSeriesInsightsReferenceDataSet, __self__).__init__(
            'azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_string_comparison_behavior: Optional[pulumi.Input[str]] = None,
            key_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TimeSeriesInsightsReferenceDataSetKeyPropertyArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            time_series_insights_environment_id: Optional[pulumi.Input[str]] = None) -> 'TimeSeriesInsightsReferenceDataSet':
        """
        Get an existing TimeSeriesInsightsReferenceDataSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_string_comparison_behavior: The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TimeSeriesInsightsReferenceDataSetKeyPropertyArgs']]]] key_properties: A `key_property` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] time_series_insights_environment_id: The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_string_comparison_behavior"] = data_string_comparison_behavior
        __props__["key_properties"] = key_properties
        __props__["location"] = location
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["time_series_insights_environment_id"] = time_series_insights_environment_id
        return TimeSeriesInsightsReferenceDataSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataStringComparisonBehavior")
    def data_string_comparison_behavior(self) -> pulumi.Output[Optional[str]]:
        """
        The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
        """
        return pulumi.get(self, "data_string_comparison_behavior")

    @property
    @pulumi.getter(name="keyProperties")
    def key_properties(self) -> pulumi.Output[Sequence['outputs.TimeSeriesInsightsReferenceDataSetKeyProperty']]:
        """
        A `key_property` block as defined below.
        """
        return pulumi.get(self, "key_properties")

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
        Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeSeriesInsightsEnvironmentId")
    def time_series_insights_environment_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "time_series_insights_environment_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

