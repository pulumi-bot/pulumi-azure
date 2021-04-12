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

__all__ = ['TimeSeriesInsightsGen2EnvironmentArgs', 'TimeSeriesInsightsGen2Environment']

@pulumi.input_type
class TimeSeriesInsightsGen2EnvironmentArgs:
    def __init__(__self__, *,
                 id_properties: pulumi.Input[Sequence[pulumi.Input[str]]],
                 resource_group_name: pulumi.Input[str],
                 sku_name: pulumi.Input[str],
                 storage: pulumi.Input['TimeSeriesInsightsGen2EnvironmentStorageArgs'],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 warm_store_data_retention_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TimeSeriesInsightsGen2Environment resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] id_properties: A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
        :param pulumi.Input['TimeSeriesInsightsGen2EnvironmentStorageArgs'] storage: A `storage` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "id_properties", id_properties)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku_name", sku_name)
        pulumi.set(__self__, "storage", storage)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if warm_store_data_retention_time is not None:
            pulumi.set(__self__, "warm_store_data_retention_time", warm_store_data_retention_time)

    @property
    @pulumi.getter(name="idProperties")
    def id_properties(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
        """
        return pulumi.get(self, "id_properties")

    @id_properties.setter
    def id_properties(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "id_properties", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Input[str]:
        """
        Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku_name", value)

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Input['TimeSeriesInsightsGen2EnvironmentStorageArgs']:
        """
        A `storage` block as defined below.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: pulumi.Input['TimeSeriesInsightsGen2EnvironmentStorageArgs']):
        pulumi.set(self, "storage", value)

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
        Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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

    @property
    @pulumi.getter(name="warmStoreDataRetentionTime")
    def warm_store_data_retention_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "warm_store_data_retention_time")

    @warm_store_data_retention_time.setter
    def warm_store_data_retention_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "warm_store_data_retention_time", value)


class TimeSeriesInsightsGen2Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage: Optional[pulumi.Input[pulumi.InputType['TimeSeriesInsightsGen2EnvironmentStorageArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 warm_store_data_retention_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure IoT Time Series Insights Gen2 Environment.

        ## Import

        Azure IoT Time Series Insights Gen2 Environment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] id_properties: A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
        :param pulumi.Input[pulumi.InputType['TimeSeriesInsightsGen2EnvironmentStorageArgs']] storage: A `storage` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TimeSeriesInsightsGen2EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure IoT Time Series Insights Gen2 Environment.

        ## Import

        Azure IoT Time Series Insights Gen2 Environment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
        ```

        :param str resource_name: The name of the resource.
        :param TimeSeriesInsightsGen2EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TimeSeriesInsightsGen2EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage: Optional[pulumi.Input[pulumi.InputType['TimeSeriesInsightsGen2EnvironmentStorageArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 warm_store_data_retention_time: Optional[pulumi.Input[str]] = None,
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

            if id_properties is None and not opts.urn:
                raise TypeError("Missing required property 'id_properties'")
            __props__['id_properties'] = id_properties
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None and not opts.urn:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            if storage is None and not opts.urn:
                raise TypeError("Missing required property 'storage'")
            __props__['storage'] = storage
            __props__['tags'] = tags
            __props__['warm_store_data_retention_time'] = warm_store_data_retention_time
            __props__['data_access_fqdn'] = None
        super(TimeSeriesInsightsGen2Environment, __self__).__init__(
            'azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_access_fqdn: Optional[pulumi.Input[str]] = None,
            id_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            storage: Optional[pulumi.Input[pulumi.InputType['TimeSeriesInsightsGen2EnvironmentStorageArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            warm_store_data_retention_time: Optional[pulumi.Input[str]] = None) -> 'TimeSeriesInsightsGen2Environment':
        """
        Get an existing TimeSeriesInsightsGen2Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_access_fqdn: The FQDN used to access the environment data.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] id_properties: A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
        :param pulumi.Input[pulumi.InputType['TimeSeriesInsightsGen2EnvironmentStorageArgs']] storage: A `storage` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_access_fqdn"] = data_access_fqdn
        __props__["id_properties"] = id_properties
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["storage"] = storage
        __props__["tags"] = tags
        __props__["warm_store_data_retention_time"] = warm_store_data_retention_time
        return TimeSeriesInsightsGen2Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataAccessFqdn")
    def data_access_fqdn(self) -> pulumi.Output[str]:
        """
        The FQDN used to access the environment data.
        """
        return pulumi.get(self, "data_access_fqdn")

    @property
    @pulumi.getter(name="idProperties")
    def id_properties(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
        """
        return pulumi.get(self, "id_properties")

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
        Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output['outputs.TimeSeriesInsightsGen2EnvironmentStorage']:
        """
        A `storage` block as defined below.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="warmStoreDataRetentionTime")
    def warm_store_data_retention_time(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "warm_store_data_retention_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

