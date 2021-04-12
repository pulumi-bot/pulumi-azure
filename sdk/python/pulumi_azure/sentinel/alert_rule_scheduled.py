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

__all__ = ['AlertRuleScheduledArgs', 'AlertRuleScheduled']

@pulumi.input_type
class AlertRuleScheduledArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 log_analytics_workspace_id: pulumi.Input[str],
                 query: pulumi.Input[str],
                 severity: pulumi.Input[str],
                 alert_rule_template_guid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_grouping: Optional[pulumi.Input['AlertRuleScheduledEventGroupingArgs']] = None,
                 incident_configuration: Optional[pulumi.Input['AlertRuleScheduledIncidentConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_frequency: Optional[pulumi.Input[str]] = None,
                 query_period: Optional[pulumi.Input[str]] = None,
                 suppression_duration: Optional[pulumi.Input[str]] = None,
                 suppression_enabled: Optional[pulumi.Input[bool]] = None,
                 tactics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_operator: Optional[pulumi.Input[str]] = None,
                 trigger_threshold: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a AlertRuleScheduled resource.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] query: The query of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] severity: The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[str] alert_rule_template_guid: The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] description: The description of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[bool] enabled: Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input['AlertRuleScheduledEventGroupingArgs'] event_grouping: A `event_grouping` block as defined below.
        :param pulumi.Input['AlertRuleScheduledIncidentConfigurationArgs'] incident_configuration: A `incident_configuration` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] query_frequency: The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        :param pulumi.Input[str] query_period: The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        :param pulumi.Input[str] suppression_duration: If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        :param pulumi.Input[bool] suppression_enabled: Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tactics: A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        :param pulumi.Input[str] trigger_operator: The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        :param pulumi.Input[int] trigger_threshold: The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        pulumi.set(__self__, "query", query)
        pulumi.set(__self__, "severity", severity)
        if alert_rule_template_guid is not None:
            pulumi.set(__self__, "alert_rule_template_guid", alert_rule_template_guid)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if event_grouping is not None:
            pulumi.set(__self__, "event_grouping", event_grouping)
        if incident_configuration is not None:
            pulumi.set(__self__, "incident_configuration", incident_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query_frequency is not None:
            pulumi.set(__self__, "query_frequency", query_frequency)
        if query_period is not None:
            pulumi.set(__self__, "query_period", query_period)
        if suppression_duration is not None:
            pulumi.set(__self__, "suppression_duration", suppression_duration)
        if suppression_enabled is not None:
            pulumi.set(__self__, "suppression_enabled", suppression_enabled)
        if tactics is not None:
            pulumi.set(__self__, "tactics", tactics)
        if trigger_operator is not None:
            pulumi.set(__self__, "trigger_operator", trigger_operator)
        if trigger_threshold is not None:
            pulumi.set(__self__, "trigger_threshold", trigger_threshold)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The friendly name of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @log_analytics_workspace_id.setter
    def log_analytics_workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_analytics_workspace_id", value)

    @property
    @pulumi.getter
    def query(self) -> pulumi.Input[str]:
        """
        The query of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: pulumi.Input[str]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Input[str]:
        """
        The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: pulumi.Input[str]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="alertRuleTemplateGuid")
    def alert_rule_template_guid(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        """
        return pulumi.get(self, "alert_rule_template_guid")

    @alert_rule_template_guid.setter
    def alert_rule_template_guid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_rule_template_guid", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventGrouping")
    def event_grouping(self) -> Optional[pulumi.Input['AlertRuleScheduledEventGroupingArgs']]:
        """
        A `event_grouping` block as defined below.
        """
        return pulumi.get(self, "event_grouping")

    @event_grouping.setter
    def event_grouping(self, value: Optional[pulumi.Input['AlertRuleScheduledEventGroupingArgs']]):
        pulumi.set(self, "event_grouping", value)

    @property
    @pulumi.getter(name="incidentConfiguration")
    def incident_configuration(self) -> Optional[pulumi.Input['AlertRuleScheduledIncidentConfigurationArgs']]:
        """
        A `incident_configuration` block as defined below.
        """
        return pulumi.get(self, "incident_configuration")

    @incident_configuration.setter
    def incident_configuration(self, value: Optional[pulumi.Input['AlertRuleScheduledIncidentConfigurationArgs']]):
        pulumi.set(self, "incident_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="queryFrequency")
    def query_frequency(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        """
        return pulumi.get(self, "query_frequency")

    @query_frequency.setter
    def query_frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_frequency", value)

    @property
    @pulumi.getter(name="queryPeriod")
    def query_period(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        """
        return pulumi.get(self, "query_period")

    @query_period.setter
    def query_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_period", value)

    @property
    @pulumi.getter(name="suppressionDuration")
    def suppression_duration(self) -> Optional[pulumi.Input[str]]:
        """
        If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        """
        return pulumi.get(self, "suppression_duration")

    @suppression_duration.setter
    def suppression_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "suppression_duration", value)

    @property
    @pulumi.getter(name="suppressionEnabled")
    def suppression_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        """
        return pulumi.get(self, "suppression_enabled")

    @suppression_enabled.setter
    def suppression_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "suppression_enabled", value)

    @property
    @pulumi.getter
    def tactics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        """
        return pulumi.get(self, "tactics")

    @tactics.setter
    def tactics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tactics", value)

    @property
    @pulumi.getter(name="triggerOperator")
    def trigger_operator(self) -> Optional[pulumi.Input[str]]:
        """
        The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        """
        return pulumi.get(self, "trigger_operator")

    @trigger_operator.setter
    def trigger_operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_operator", value)

    @property
    @pulumi.getter(name="triggerThreshold")
    def trigger_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "trigger_threshold")

    @trigger_threshold.setter
    def trigger_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trigger_threshold", value)


class AlertRuleScheduled(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_rule_template_guid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_grouping: Optional[pulumi.Input[pulumi.InputType['AlertRuleScheduledEventGroupingArgs']]] = None,
                 incident_configuration: Optional[pulumi.Input[pulumi.InputType['AlertRuleScheduledIncidentConfigurationArgs']]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 query_frequency: Optional[pulumi.Input[str]] = None,
                 query_period: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 suppression_duration: Optional[pulumi.Input[str]] = None,
                 suppression_enabled: Optional[pulumi.Input[bool]] = None,
                 tactics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_operator: Optional[pulumi.Input[str]] = None,
                 trigger_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

        ## Import

        Sentinel Scheduled Alert Rules can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:sentinel/alertRuleScheduled:AlertRuleScheduled example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_rule_template_guid: The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] description: The description of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[bool] enabled: Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input[pulumi.InputType['AlertRuleScheduledEventGroupingArgs']] event_grouping: A `event_grouping` block as defined below.
        :param pulumi.Input[pulumi.InputType['AlertRuleScheduledIncidentConfigurationArgs']] incident_configuration: A `incident_configuration` block as defined below.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] query: The query of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] query_frequency: The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        :param pulumi.Input[str] query_period: The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        :param pulumi.Input[str] severity: The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[str] suppression_duration: If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        :param pulumi.Input[bool] suppression_enabled: Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tactics: A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        :param pulumi.Input[str] trigger_operator: The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        :param pulumi.Input[int] trigger_threshold: The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlertRuleScheduledArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        Sentinel Scheduled Alert Rules can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:sentinel/alertRuleScheduled:AlertRuleScheduled example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param AlertRuleScheduledArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlertRuleScheduledArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_rule_template_guid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_grouping: Optional[pulumi.Input[pulumi.InputType['AlertRuleScheduledEventGroupingArgs']]] = None,
                 incident_configuration: Optional[pulumi.Input[pulumi.InputType['AlertRuleScheduledIncidentConfigurationArgs']]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 query_frequency: Optional[pulumi.Input[str]] = None,
                 query_period: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 suppression_duration: Optional[pulumi.Input[str]] = None,
                 suppression_enabled: Optional[pulumi.Input[bool]] = None,
                 tactics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_operator: Optional[pulumi.Input[str]] = None,
                 trigger_threshold: Optional[pulumi.Input[int]] = None,
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

            __props__['alert_rule_template_guid'] = alert_rule_template_guid
            __props__['description'] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            __props__['event_grouping'] = event_grouping
            __props__['incident_configuration'] = incident_configuration
            if log_analytics_workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'log_analytics_workspace_id'")
            __props__['log_analytics_workspace_id'] = log_analytics_workspace_id
            __props__['name'] = name
            if query is None and not opts.urn:
                raise TypeError("Missing required property 'query'")
            __props__['query'] = query
            __props__['query_frequency'] = query_frequency
            __props__['query_period'] = query_period
            if severity is None and not opts.urn:
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_rule_template_guid: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            event_grouping: Optional[pulumi.Input[pulumi.InputType['AlertRuleScheduledEventGroupingArgs']]] = None,
            incident_configuration: Optional[pulumi.Input[pulumi.InputType['AlertRuleScheduledIncidentConfigurationArgs']]] = None,
            log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query: Optional[pulumi.Input[str]] = None,
            query_frequency: Optional[pulumi.Input[str]] = None,
            query_period: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            suppression_duration: Optional[pulumi.Input[str]] = None,
            suppression_enabled: Optional[pulumi.Input[bool]] = None,
            tactics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            trigger_operator: Optional[pulumi.Input[str]] = None,
            trigger_threshold: Optional[pulumi.Input[int]] = None) -> 'AlertRuleScheduled':
        """
        Get an existing AlertRuleScheduled resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_rule_template_guid: The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] description: The description of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[bool] enabled: Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input[pulumi.InputType['AlertRuleScheduledEventGroupingArgs']] event_grouping: A `event_grouping` block as defined below.
        :param pulumi.Input[pulumi.InputType['AlertRuleScheduledIncidentConfigurationArgs']] incident_configuration: A `incident_configuration` block as defined below.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        :param pulumi.Input[str] query: The query of this Sentinel Scheduled Alert Rule.
        :param pulumi.Input[str] query_frequency: The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        :param pulumi.Input[str] query_period: The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        :param pulumi.Input[str] severity: The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[str] suppression_duration: If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        :param pulumi.Input[bool] suppression_enabled: Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tactics: A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        :param pulumi.Input[str] trigger_operator: The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        :param pulumi.Input[int] trigger_threshold: The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alert_rule_template_guid"] = alert_rule_template_guid
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["event_grouping"] = event_grouping
        __props__["incident_configuration"] = incident_configuration
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

    @property
    @pulumi.getter(name="alertRuleTemplateGuid")
    def alert_rule_template_guid(self) -> pulumi.Output[Optional[str]]:
        """
        The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        """
        return pulumi.get(self, "alert_rule_template_guid")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The friendly name of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventGrouping")
    def event_grouping(self) -> pulumi.Output[Optional['outputs.AlertRuleScheduledEventGrouping']]:
        """
        A `event_grouping` block as defined below.
        """
        return pulumi.get(self, "event_grouping")

    @property
    @pulumi.getter(name="incidentConfiguration")
    def incident_configuration(self) -> pulumi.Output['outputs.AlertRuleScheduledIncidentConfiguration']:
        """
        A `incident_configuration` block as defined below.
        """
        return pulumi.get(self, "incident_configuration")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> pulumi.Output[str]:
        """
        The query of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="queryFrequency")
    def query_frequency(self) -> pulumi.Output[Optional[str]]:
        """
        The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        """
        return pulumi.get(self, "query_frequency")

    @property
    @pulumi.getter(name="queryPeriod")
    def query_period(self) -> pulumi.Output[Optional[str]]:
        """
        The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        """
        return pulumi.get(self, "query_period")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="suppressionDuration")
    def suppression_duration(self) -> pulumi.Output[Optional[str]]:
        """
        If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        """
        return pulumi.get(self, "suppression_duration")

    @property
    @pulumi.getter(name="suppressionEnabled")
    def suppression_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        """
        return pulumi.get(self, "suppression_enabled")

    @property
    @pulumi.getter
    def tactics(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        """
        return pulumi.get(self, "tactics")

    @property
    @pulumi.getter(name="triggerOperator")
    def trigger_operator(self) -> pulumi.Output[Optional[str]]:
        """
        The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        """
        return pulumi.get(self, "trigger_operator")

    @property
    @pulumi.getter(name="triggerThreshold")
    def trigger_threshold(self) -> pulumi.Output[Optional[int]]:
        """
        The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        """
        return pulumi.get(self, "trigger_threshold")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

