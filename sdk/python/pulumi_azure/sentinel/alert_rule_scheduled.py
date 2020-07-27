# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['AlertRuleScheduled']


class AlertRuleScheduled(pulumi.CustomResource):
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    The description of this Sentinel Scheduled Alert Rule.
    """
    display_name: pulumi.Output[str] = pulumi.output_property("displayName")
    """
    The friendly name of this Sentinel Scheduled Alert Rule.
    """
    enabled: pulumi.Output[Optional[bool]] = pulumi.output_property("enabled")
    """
    Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
    """
    log_analytics_workspace_id: pulumi.Output[str] = pulumi.output_property("logAnalyticsWorkspaceId")
    """
    The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
    """
    query: pulumi.Output[str] = pulumi.output_property("query")
    """
    The query of this Sentinel Scheduled Alert Rule.
    """
    query_frequency: pulumi.Output[Optional[str]] = pulumi.output_property("queryFrequency")
    """
    The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
    """
    query_period: pulumi.Output[Optional[str]] = pulumi.output_property("queryPeriod")
    """
    The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
    """
    severity: pulumi.Output[str] = pulumi.output_property("severity")
    """
    The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
    """
    suppression_duration: pulumi.Output[Optional[str]] = pulumi.output_property("suppressionDuration")
    """
    If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
    """
    suppression_enabled: pulumi.Output[Optional[bool]] = pulumi.output_property("suppressionEnabled")
    """
    Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
    """
    tactics: pulumi.Output[Optional[List[str]]] = pulumi.output_property("tactics")
    """
    A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
    """
    trigger_operator: pulumi.Output[Optional[str]] = pulumi.output_property("triggerOperator")
    """
    The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
    """
    trigger_threshold: pulumi.Output[Optional[float]] = pulumi.output_property("triggerThreshold")
    """
    The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, display_name: Optional[pulumi.Input[str]] = None, enabled: Optional[pulumi.Input[bool]] = None, log_analytics_workspace_id: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, query: Optional[pulumi.Input[str]] = None, query_frequency: Optional[pulumi.Input[str]] = None, query_period: Optional[pulumi.Input[str]] = None, severity: Optional[pulumi.Input[str]] = None, suppression_duration: Optional[pulumi.Input[str]] = None, suppression_enabled: Optional[pulumi.Input[bool]] = None, tactics: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, trigger_operator: Optional[pulumi.Input[str]] = None, trigger_threshold: Optional[pulumi.Input[float]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Sentinel Scheduled Alert Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="pergb2018")
        example_alert_rule_scheduled = azure.sentinel.AlertRuleScheduled("exampleAlertRuleScheduled",
            log_analytics_workspace_id=example_analytics_workspace.id,
            display_name="example",
            severity="High",
            query=\"\"\"AzureActivity |
          where OperationName == "Create or Update Virtual Machine" or OperationName =="Create Deployment" |
          where ActivityStatus == "Succeeded" |
          make-series dcount(ResourceId) default=0 on EventSubmissionTimestamp in range(ago(7d), now(), 1d) by Caller
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[bool] enabled: Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] query: The query of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] query_frequency: The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        :param pulumi.Input[str] query_period: The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        :param pulumi.Input[str] severity: The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[str] suppression_duration: If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        :param pulumi.Input[bool] suppression_enabled: Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] tactics: A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        :param pulumi.Input[str] trigger_operator: The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        :param pulumi.Input[float] trigger_threshold: The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
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

            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            if log_analytics_workspace_id is None:
                raise TypeError("Missing required property 'log_analytics_workspace_id'")
            __props__['log_analytics_workspace_id'] = log_analytics_workspace_id
            __props__['name'] = name
            if query is None:
                raise TypeError("Missing required property 'query'")
            __props__['query'] = query
            __props__['query_frequency'] = query_frequency
            __props__['query_period'] = query_period
            if severity is None:
                raise TypeError("Missing required property 'severity'")
            __props__['severity'] = severity
            __props__['suppression_duration'] = suppression_duration
            __props__['suppression_enabled'] = suppression_enabled
            __props__['tactics'] = tactics
            __props__['trigger_operator'] = trigger_operator
            __props__['trigger_threshold'] = trigger_threshold
        super(AlertRuleScheduled, __self__).__init__(
            'azure:sentinel/alertRuleScheduled:AlertRuleScheduled',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, display_name: Optional[pulumi.Input[str]] = None, enabled: Optional[pulumi.Input[bool]] = None, log_analytics_workspace_id: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, query: Optional[pulumi.Input[str]] = None, query_frequency: Optional[pulumi.Input[str]] = None, query_period: Optional[pulumi.Input[str]] = None, severity: Optional[pulumi.Input[str]] = None, suppression_duration: Optional[pulumi.Input[str]] = None, suppression_enabled: Optional[pulumi.Input[bool]] = None, tactics: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, trigger_operator: Optional[pulumi.Input[str]] = None, trigger_threshold: Optional[pulumi.Input[float]] = None) -> 'AlertRuleScheduled':
        """
        Get an existing AlertRuleScheduled resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[bool] enabled: Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] query: The query of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] query_frequency: The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        :param pulumi.Input[str] query_period: The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        :param pulumi.Input[str] severity: The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[str] suppression_duration: If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        :param pulumi.Input[bool] suppression_enabled: Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] tactics: A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        :param pulumi.Input[str] trigger_operator: The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        :param pulumi.Input[float] trigger_threshold: The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["log_analytics_workspace_id"] = log_analytics_workspace_id
        __props__["name"] = name
        __props__["query"] = query
        __props__["query_frequency"] = query_frequency
        __props__["query_period"] = query_period
        __props__["severity"] = severity
        __props__["suppression_duration"] = suppression_duration
        __props__["suppression_enabled"] = suppression_enabled
        __props__["tactics"] = tactics
        __props__["trigger_operator"] = trigger_operator
        __props__["trigger_threshold"] = trigger_threshold
        return AlertRuleScheduled(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

