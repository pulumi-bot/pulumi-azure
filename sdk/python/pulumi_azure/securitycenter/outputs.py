# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'AutomationAction',
    'AutomationSource',
    'AutomationSourceRuleSet',
    'AutomationSourceRuleSetRule',
]

@pulumi.output_type
class AutomationAction(dict):
    def __init__(__self__, *,
                 resource_id: str,
                 type: str,
                 connection_string: Optional[str] = None,
                 trigger_url: Optional[str] = None):
        """
        :param str resource_id: The resource id of the target Logic App, Event Hub namespace or Log Analytics workspace.
        :param str type: Type of Azure resource to send data to. Must be set to one of: `LogicApp`, `EventHub` or `LogAnalytics`.
        :param str connection_string: A connection string to send data to the target Event Hub namespace, this should include a key with send permissions.
        :param str trigger_url: The callback URL to trigger the Logic App that will receive and process data sent by this automation. This can be found in the Azure Portal under "See trigger history"
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "type", type)
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if trigger_url is not None:
            pulumi.set(__self__, "trigger_url", trigger_url)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The resource id of the target Logic App, Event Hub namespace or Log Analytics workspace.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of Azure resource to send data to. Must be set to one of: `LogicApp`, `EventHub` or `LogAnalytics`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[str]:
        """
        A connection string to send data to the target Event Hub namespace, this should include a key with send permissions.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="triggerUrl")
    def trigger_url(self) -> Optional[str]:
        """
        The callback URL to trigger the Logic App that will receive and process data sent by this automation. This can be found in the Azure Portal under "See trigger history"
        """
        return pulumi.get(self, "trigger_url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AutomationSource(dict):
    def __init__(__self__, *,
                 event_source: str,
                 rule_sets: Optional[Sequence['outputs.AutomationSourceRuleSet']] = None):
        """
        :param str event_source: Type of data that will trigger this automation. Must be one of `Alerts`, `Assessments` or `SubAssessments`. Note. assessments are also referred to as recommendations
        :param Sequence['AutomationSourceRuleSetArgs'] rule_sets: A set of rules which evaluate upon event and data interception. This is defined in one or more `rule_set` blocks as defined below.
        """
        pulumi.set(__self__, "event_source", event_source)
        if rule_sets is not None:
            pulumi.set(__self__, "rule_sets", rule_sets)

    @property
    @pulumi.getter(name="eventSource")
    def event_source(self) -> str:
        """
        Type of data that will trigger this automation. Must be one of `Alerts`, `Assessments` or `SubAssessments`. Note. assessments are also referred to as recommendations
        """
        return pulumi.get(self, "event_source")

    @property
    @pulumi.getter(name="ruleSets")
    def rule_sets(self) -> Optional[Sequence['outputs.AutomationSourceRuleSet']]:
        """
        A set of rules which evaluate upon event and data interception. This is defined in one or more `rule_set` blocks as defined below.
        """
        return pulumi.get(self, "rule_sets")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AutomationSourceRuleSet(dict):
    def __init__(__self__, *,
                 rules: Sequence['outputs.AutomationSourceRuleSetRule']):
        """
        :param Sequence['AutomationSourceRuleSetRuleArgs'] rules: One or more `rule` blocks as defined below.
        """
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.AutomationSourceRuleSetRule']:
        """
        One or more `rule` blocks as defined below.
        """
        return pulumi.get(self, "rules")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AutomationSourceRuleSetRule(dict):
    def __init__(__self__, *,
                 expected_value: str,
                 operator: str,
                 property_path: str,
                 property_type: str):
        """
        :param str expected_value: A value that will be compared with the value in `property_path`.
        :param str operator: The comparison operator to use, must be one of: `Contains`, `EndsWith`, `Equals`, `GreaterThan`, `GreaterThanOrEqualTo`, `LesserThan`, `LesserThanOrEqualTo`, `NotEquals`, `StartsWith`
        :param str property_path: The JPath of the entity model property that should be checked.
        :param str property_type: The data type of the compared operands, must be one of: `Integer`, `String`, `Boolean` or `Number`.
        """
        pulumi.set(__self__, "expected_value", expected_value)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "property_path", property_path)
        pulumi.set(__self__, "property_type", property_type)

    @property
    @pulumi.getter(name="expectedValue")
    def expected_value(self) -> str:
        """
        A value that will be compared with the value in `property_path`.
        """
        return pulumi.get(self, "expected_value")

    @property
    @pulumi.getter
    def operator(self) -> str:
        """
        The comparison operator to use, must be one of: `Contains`, `EndsWith`, `Equals`, `GreaterThan`, `GreaterThanOrEqualTo`, `LesserThan`, `LesserThanOrEqualTo`, `NotEquals`, `StartsWith`
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter(name="propertyPath")
    def property_path(self) -> str:
        """
        The JPath of the entity model property that should be checked.
        """
        return pulumi.get(self, "property_path")

    @property
    @pulumi.getter(name="propertyType")
    def property_type(self) -> str:
        """
        The data type of the compared operands, must be one of: `Integer`, `String`, `Boolean` or `Number`.
        """
        return pulumi.get(self, "property_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


