# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ScheduledQueryRulesLog(pulumi.CustomResource):
    authorized_resource_ids: pulumi.Output[list]
    criteria: pulumi.Output[dict]
    """
    A `criteria` block as defined below.

      * `dimensions` (`list`) - A `dimension` block as defined below.
        * `name` (`str`) - Name of the dimension.
        * `operator` (`str`) - Operator for dimension values, - 'Include'.
        * `values` (`list`) - List of dimension values.

      * `metricName` (`str`) - Name of the metric.  Supported metrics are listed in the Azure Monitor [Microsoft.OperationalInsights/workspaces](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/metrics-supported#microsoftoperationalinsightsworkspaces) metrics namespace.
    """
    data_source_id: pulumi.Output[str]
    """
    The resource uri over which log search query is to be run.
    """
    description: pulumi.Output[str]
    """
    The description of the scheduled query rule.
    """
    enabled: pulumi.Output[bool]
    """
    Whether this scheduled query rule is enabled.  Default is `true`.
    """
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    The name of the scheduled query rule. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the scheduled query rule instance.
    """
    tags: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, authorized_resource_ids=None, criteria=None, data_source_id=None, description=None, enabled=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a LogToMetricAction Scheduled Query Rules resource within Azure Monitor.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] data_source_id: The resource uri over which log search query is to be run.
        :param pulumi.Input[str] description: The description of the scheduled query rule.
        :param pulumi.Input[bool] enabled: Whether this scheduled query rule is enabled.  Default is `true`.
        :param pulumi.Input[str] name: The name of the scheduled query rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the scheduled query rule instance.

        The **criteria** object supports the following:

          * `dimensions` (`pulumi.Input[list]`) - A `dimension` block as defined below.
            * `name` (`pulumi.Input[str]`) - Name of the dimension.
            * `operator` (`pulumi.Input[str]`) - Operator for dimension values, - 'Include'.
            * `values` (`pulumi.Input[list]`) - List of dimension values.

          * `metricName` (`pulumi.Input[str]`) - Name of the metric.  Supported metrics are listed in the Azure Monitor [Microsoft.OperationalInsights/workspaces](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/metrics-supported#microsoftoperationalinsightsworkspaces) metrics namespace.
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

            __props__['authorized_resource_ids'] = authorized_resource_ids
            if criteria is None:
                raise TypeError("Missing required property 'criteria'")
            __props__['criteria'] = criteria
            if data_source_id is None:
                raise TypeError("Missing required property 'data_source_id'")
            __props__['data_source_id'] = data_source_id
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(ScheduledQueryRulesLog, __self__).__init__(
            'azure:monitoring/scheduledQueryRulesLog:ScheduledQueryRulesLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, authorized_resource_ids=None, criteria=None, data_source_id=None, description=None, enabled=None, location=None, name=None, resource_group_name=None, tags=None):
        """
        Get an existing ScheduledQueryRulesLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] data_source_id: The resource uri over which log search query is to be run.
        :param pulumi.Input[str] description: The description of the scheduled query rule.
        :param pulumi.Input[bool] enabled: Whether this scheduled query rule is enabled.  Default is `true`.
        :param pulumi.Input[str] name: The name of the scheduled query rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the scheduled query rule instance.

        The **criteria** object supports the following:

          * `dimensions` (`pulumi.Input[list]`) - A `dimension` block as defined below.
            * `name` (`pulumi.Input[str]`) - Name of the dimension.
            * `operator` (`pulumi.Input[str]`) - Operator for dimension values, - 'Include'.
            * `values` (`pulumi.Input[list]`) - List of dimension values.

          * `metricName` (`pulumi.Input[str]`) - Name of the metric.  Supported metrics are listed in the Azure Monitor [Microsoft.OperationalInsights/workspaces](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/metrics-supported#microsoftoperationalinsightsworkspaces) metrics namespace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized_resource_ids"] = authorized_resource_ids
        __props__["criteria"] = criteria
        __props__["data_source_id"] = data_source_id
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return ScheduledQueryRulesLog(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
