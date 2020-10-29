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

__all__ = ['ResourceGroupExport']


class ResourceGroupExport(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 delivery_info: Optional[pulumi.Input[pulumi.InputType['ResourceGroupExportDeliveryInfoArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[pulumi.InputType['ResourceGroupExportQueryArgs']]] = None,
                 recurrence_period_end: Optional[pulumi.Input[str]] = None,
                 recurrence_period_start: Optional[pulumi.Input[str]] = None,
                 recurrence_type: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Cost Management Export for a Resource Group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Is the cost management export active? Default is `true`.
        :param pulumi.Input[pulumi.InputType['ResourceGroupExportDeliveryInfoArgs']] delivery_info: A `delivery_info` block as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ResourceGroupExportQueryArgs']] query: A `query` block as defined below.
        :param pulumi.Input[str] recurrence_period_end: The date the export will stop capturing information.
        :param pulumi.Input[str] recurrence_period_start: The date the export will start capturing information.
        :param pulumi.Input[str] recurrence_type: How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
        :param pulumi.Input[str] resource_group_id: The id of the resource group in which to export information.
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

            __props__['active'] = active
            if delivery_info is None:
                raise TypeError("Missing required property 'delivery_info'")
            __props__['delivery_info'] = delivery_info
            __props__['name'] = name
            if query is None:
                raise TypeError("Missing required property 'query'")
            __props__['query'] = query
            if recurrence_period_end is None:
                raise TypeError("Missing required property 'recurrence_period_end'")
            __props__['recurrence_period_end'] = recurrence_period_end
            if recurrence_period_start is None:
                raise TypeError("Missing required property 'recurrence_period_start'")
            __props__['recurrence_period_start'] = recurrence_period_start
            if recurrence_type is None:
                raise TypeError("Missing required property 'recurrence_type'")
            __props__['recurrence_type'] = recurrence_type
            if resource_group_id is None:
                raise TypeError("Missing required property 'resource_group_id'")
            __props__['resource_group_id'] = resource_group_id
        super(ResourceGroupExport, __self__).__init__(
            'azure:costmanagement/resourceGroupExport:ResourceGroupExport',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            delivery_info: Optional[pulumi.Input[pulumi.InputType['ResourceGroupExportDeliveryInfoArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query: Optional[pulumi.Input[pulumi.InputType['ResourceGroupExportQueryArgs']]] = None,
            recurrence_period_end: Optional[pulumi.Input[str]] = None,
            recurrence_period_start: Optional[pulumi.Input[str]] = None,
            recurrence_type: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None) -> 'ResourceGroupExport':
        """
        Get an existing ResourceGroupExport resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Is the cost management export active? Default is `true`.
        :param pulumi.Input[pulumi.InputType['ResourceGroupExportDeliveryInfoArgs']] delivery_info: A `delivery_info` block as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ResourceGroupExportQueryArgs']] query: A `query` block as defined below.
        :param pulumi.Input[str] recurrence_period_end: The date the export will stop capturing information.
        :param pulumi.Input[str] recurrence_period_start: The date the export will start capturing information.
        :param pulumi.Input[str] recurrence_type: How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
        :param pulumi.Input[str] resource_group_id: The id of the resource group in which to export information.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["active"] = active
        __props__["delivery_info"] = delivery_info
        __props__["name"] = name
        __props__["query"] = query
        __props__["recurrence_period_end"] = recurrence_period_end
        __props__["recurrence_period_start"] = recurrence_period_start
        __props__["recurrence_type"] = recurrence_type
        __props__["resource_group_id"] = resource_group_id
        return ResourceGroupExport(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the cost management export active? Default is `true`.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="deliveryInfo")
    def delivery_info(self) -> pulumi.Output['outputs.ResourceGroupExportDeliveryInfo']:
        """
        A `delivery_info` block as defined below.
        """
        return pulumi.get(self, "delivery_info")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> pulumi.Output['outputs.ResourceGroupExportQuery']:
        """
        A `query` block as defined below.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="recurrencePeriodEnd")
    def recurrence_period_end(self) -> pulumi.Output[str]:
        """
        The date the export will stop capturing information.
        """
        return pulumi.get(self, "recurrence_period_end")

    @property
    @pulumi.getter(name="recurrencePeriodStart")
    def recurrence_period_start(self) -> pulumi.Output[str]:
        """
        The date the export will start capturing information.
        """
        return pulumi.get(self, "recurrence_period_start")

    @property
    @pulumi.getter(name="recurrenceType")
    def recurrence_type(self) -> pulumi.Output[str]:
        """
        How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
        """
        return pulumi.get(self, "recurrence_type")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The id of the resource group in which to export information.
        """
        return pulumi.get(self, "resource_group_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

