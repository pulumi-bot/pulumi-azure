# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ScheduledQueryRulesAlert']


class ScheduledQueryRulesAlert(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertActionArgs']]] = None,
                 authorized_resource_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 frequency: Optional[pulumi.Input[float]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 query_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 throttling: Optional[pulumi.Input[float]] = None,
                 time_window: Optional[pulumi.Input[float]] = None,
                 trigger: Optional[pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertTriggerArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an AlertingAction Scheduled Query Rules resource within Azure Monitor.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertActionArgs']] action: An `action` block as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] authorized_resource_ids: List of Resource IDs referred into query.
        :param pulumi.Input[str] data_source_id: The resource URI over which log search query is to be run.
        :param pulumi.Input[str] description: The description of the scheduled query rule.
        :param pulumi.Input[bool] enabled: Whether this scheduled query rule is enabled.  Default is `true`.
        :param pulumi.Input[float] frequency: Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
        :param pulumi.Input[str] name: The name of the scheduled query rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query: Log search query.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the scheduled query rule instance.
        :param pulumi.Input[float] severity: Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
        :param pulumi.Input[float] throttling: Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
        :param pulumi.Input[float] time_window: Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
        :param pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertTriggerArgs']] trigger: The condition that results in the alert rule being run.
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

            if action is None:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            __props__['authorized_resource_ids'] = authorized_resource_ids
            if data_source_id is None:
                raise TypeError("Missing required property 'data_source_id'")
            __props__['data_source_id'] = data_source_id
            __props__['description'] = description
            __props__['enabled'] = enabled
            if frequency is None:
                raise TypeError("Missing required property 'frequency'")
            __props__['frequency'] = frequency
            __props__['location'] = location
            __props__['name'] = name
            if query is None:
                raise TypeError("Missing required property 'query'")
            __props__['query'] = query
            __props__['query_type'] = query_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['severity'] = severity
            __props__['tags'] = tags
            __props__['throttling'] = throttling
            if time_window is None:
                raise TypeError("Missing required property 'time_window'")
            __props__['time_window'] = time_window
            if trigger is None:
                raise TypeError("Missing required property 'trigger'")
            __props__['trigger'] = trigger
        super(ScheduledQueryRulesAlert, __self__).__init__(
            'azure:monitoring/scheduledQueryRulesAlert:ScheduledQueryRulesAlert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertActionArgs']]] = None,
            authorized_resource_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            data_source_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            frequency: Optional[pulumi.Input[float]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query: Optional[pulumi.Input[str]] = None,
            query_type: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[float]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            throttling: Optional[pulumi.Input[float]] = None,
            time_window: Optional[pulumi.Input[float]] = None,
            trigger: Optional[pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertTriggerArgs']]] = None) -> 'ScheduledQueryRulesAlert':
        """
        Get an existing ScheduledQueryRulesAlert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertActionArgs']] action: An `action` block as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] authorized_resource_ids: List of Resource IDs referred into query.
        :param pulumi.Input[str] data_source_id: The resource URI over which log search query is to be run.
        :param pulumi.Input[str] description: The description of the scheduled query rule.
        :param pulumi.Input[bool] enabled: Whether this scheduled query rule is enabled.  Default is `true`.
        :param pulumi.Input[float] frequency: Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
        :param pulumi.Input[str] name: The name of the scheduled query rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query: Log search query.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the scheduled query rule instance.
        :param pulumi.Input[float] severity: Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
        :param pulumi.Input[float] throttling: Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
        :param pulumi.Input[float] time_window: Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
        :param pulumi.Input[pulumi.InputType['ScheduledQueryRulesAlertTriggerArgs']] trigger: The condition that results in the alert rule being run.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["authorized_resource_ids"] = authorized_resource_ids
        __props__["data_source_id"] = data_source_id
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["frequency"] = frequency
        __props__["location"] = location
        __props__["name"] = name
        __props__["query"] = query
        __props__["query_type"] = query_type
        __props__["resource_group_name"] = resource_group_name
        __props__["severity"] = severity
        __props__["tags"] = tags
        __props__["throttling"] = throttling
        __props__["time_window"] = time_window
        __props__["trigger"] = trigger
        return ScheduledQueryRulesAlert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> 'outputs.ScheduledQueryRulesAlertAction':
        """
        An `action` block as defined below.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="authorizedResourceIds")
    def authorized_resource_ids(self) -> Optional[List[str]]:
        """
        List of Resource IDs referred into query.
        """
        return pulumi.get(self, "authorized_resource_ids")

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> str:
        """
        The resource URI over which log search query is to be run.
        """
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the scheduled query rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether this scheduled query rule is enabled.  Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def frequency(self) -> float:
        """
        Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the scheduled query rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> str:
        """
        Log search query.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="queryType")
    def query_type(self) -> Optional[str]:
        return pulumi.get(self, "query_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the scheduled query rule instance.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def severity(self) -> Optional[float]:
        """
        Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def throttling(self) -> Optional[float]:
        """
        Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
        """
        return pulumi.get(self, "throttling")

    @property
    @pulumi.getter(name="timeWindow")
    def time_window(self) -> float:
        """
        Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
        """
        return pulumi.get(self, "time_window")

    @property
    @pulumi.getter
    def trigger(self) -> 'outputs.ScheduledQueryRulesAlertTrigger':
        """
        The condition that results in the alert rule being run.
        """
        return pulumi.get(self, "trigger")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

