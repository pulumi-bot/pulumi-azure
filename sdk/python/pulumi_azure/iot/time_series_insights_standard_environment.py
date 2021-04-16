# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TimeSeriesInsightsStandardEnvironmentArgs', 'TimeSeriesInsightsStandardEnvironment']

@pulumi.input_type
class TimeSeriesInsightsStandardEnvironmentArgs:
    def __init__(__self__, *,
                 data_retention_time: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 sku_name: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_key: Optional[pulumi.Input[str]] = None,
                 storage_limit_exceeded_behavior: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TimeSeriesInsightsStandardEnvironment resource.
        :param pulumi.Input[str] data_retention_time: Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] partition_key: The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_limit_exceeded_behavior: Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "data_retention_time", data_retention_time)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku_name", sku_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition_key is not None:
            pulumi.set(__self__, "partition_key", partition_key)
        if storage_limit_exceeded_behavior is not None:
            pulumi.set(__self__, "storage_limit_exceeded_behavior", storage_limit_exceeded_behavior)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataRetentionTime")
    def data_retention_time(self) -> pulumi.Input[str]:
        """
        Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "data_retention_time")

    @data_retention_time.setter
    def data_retention_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_retention_time", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Input[str]:
        """
        Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku_name", value)

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
        Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="partitionKey")
    def partition_key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "partition_key")

    @partition_key.setter
    def partition_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition_key", value)

    @property
    @pulumi.getter(name="storageLimitExceededBehavior")
    def storage_limit_exceeded_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        """
        return pulumi.get(self, "storage_limit_exceeded_behavior")

    @storage_limit_exceeded_behavior.setter
    def storage_limit_exceeded_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_limit_exceeded_behavior", value)

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


@pulumi.input_type
class _TimeSeriesInsightsStandardEnvironmentState:
    def __init__(__self__, *,
                 data_retention_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_key: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_limit_exceeded_behavior: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering TimeSeriesInsightsStandardEnvironment resources.
        :param pulumi.Input[str] data_retention_time: Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] partition_key: The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_limit_exceeded_behavior: Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if data_retention_time is not None:
            pulumi.set(__self__, "data_retention_time", data_retention_time)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition_key is not None:
            pulumi.set(__self__, "partition_key", partition_key)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if sku_name is not None:
            pulumi.set(__self__, "sku_name", sku_name)
        if storage_limit_exceeded_behavior is not None:
            pulumi.set(__self__, "storage_limit_exceeded_behavior", storage_limit_exceeded_behavior)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataRetentionTime")
    def data_retention_time(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "data_retention_time")

    @data_retention_time.setter
    def data_retention_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_retention_time", value)

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
        Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="partitionKey")
    def partition_key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "partition_key")

    @partition_key.setter
    def partition_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition_key", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sku_name", value)

    @property
    @pulumi.getter(name="storageLimitExceededBehavior")
    def storage_limit_exceeded_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        """
        return pulumi.get(self, "storage_limit_exceeded_behavior")

    @storage_limit_exceeded_behavior.setter
    def storage_limit_exceeded_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_limit_exceeded_behavior", value)

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


class TimeSeriesInsightsStandardEnvironment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_retention_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_key: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_limit_exceeded_behavior: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an Azure IoT Time Series Insights Standard Environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_time_series_insights_standard_environment = azure.iot.TimeSeriesInsightsStandardEnvironment("exampleTimeSeriesInsightsStandardEnvironment",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="S1_1",
            data_retention_time="P30D")
        ```

        ## Import

        Azure IoT Time Series Insights Standard Environment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_retention_time: Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] partition_key: The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_limit_exceeded_behavior: Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TimeSeriesInsightsStandardEnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure IoT Time Series Insights Standard Environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_time_series_insights_standard_environment = azure.iot.TimeSeriesInsightsStandardEnvironment("exampleTimeSeriesInsightsStandardEnvironment",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="S1_1",
            data_retention_time="P30D")
        ```

        ## Import

        Azure IoT Time Series Insights Standard Environment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
        ```

        :param str resource_name: The name of the resource.
        :param TimeSeriesInsightsStandardEnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TimeSeriesInsightsStandardEnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_retention_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_key: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_limit_exceeded_behavior: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TimeSeriesInsightsStandardEnvironmentArgs.__new__(TimeSeriesInsightsStandardEnvironmentArgs)

            if data_retention_time is None and not opts.urn:
                raise TypeError("Missing required property 'data_retention_time'")
            __props__.__dict__["data_retention_time"] = data_retention_time
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["partition_key"] = partition_key
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sku_name is None and not opts.urn:
                raise TypeError("Missing required property 'sku_name'")
            __props__.__dict__["sku_name"] = sku_name
            __props__.__dict__["storage_limit_exceeded_behavior"] = storage_limit_exceeded_behavior
            __props__.__dict__["tags"] = tags
        super(TimeSeriesInsightsStandardEnvironment, __self__).__init__(
            'azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_retention_time: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            storage_limit_exceeded_behavior: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'TimeSeriesInsightsStandardEnvironment':
        """
        Get an existing TimeSeriesInsightsStandardEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_retention_time: Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] partition_key: The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_limit_exceeded_behavior: Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TimeSeriesInsightsStandardEnvironmentState.__new__(_TimeSeriesInsightsStandardEnvironmentState)

        __props__.__dict__["data_retention_time"] = data_retention_time
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["partition_key"] = partition_key
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["sku_name"] = sku_name
        __props__.__dict__["storage_limit_exceeded_behavior"] = storage_limit_exceeded_behavior
        __props__.__dict__["tags"] = tags
        return TimeSeriesInsightsStandardEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataRetentionTime")
    def data_retention_time(self) -> pulumi.Output[str]:
        """
        Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "data_retention_time")

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
        Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionKey")
    def partition_key(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "partition_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="storageLimitExceededBehavior")
    def storage_limit_exceeded_behavior(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
        """
        return pulumi.get(self, "storage_limit_exceeded_behavior")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

