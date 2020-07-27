# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'PolicyCustomRule',
    'PolicyCustomRuleMatchCondition',
    'PolicyCustomRuleMatchConditionMatchVariable',
    'PolicyManagedRules',
    'PolicyManagedRulesExclusion',
    'PolicyManagedRulesManagedRuleSet',
    'PolicyManagedRulesManagedRuleSetRuleGroupOverride',
    'PolicyPolicySettings',
]

@pulumi.output_type
class PolicyCustomRule(dict):
    action: str = pulumi.output_property("action")
    """
    Type of action.
    """
    match_conditions: List['outputs.PolicyCustomRuleMatchCondition'] = pulumi.output_property("matchConditions")
    """
    One or more `match_conditions` blocks as defined below.
    """
    name: Optional[str] = pulumi.output_property("name")
    """
    Gets name of the resource that is unique within a policy. This name can be used to access the resource.
    """
    priority: float = pulumi.output_property("priority")
    """
    Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
    """
    rule_type: str = pulumi.output_property("ruleType")
    """
    Describes the type of rule.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyCustomRuleMatchCondition(dict):
    match_values: List[str] = pulumi.output_property("matchValues")
    """
    A list of match values.
    """
    match_variables: List['outputs.PolicyCustomRuleMatchConditionMatchVariable'] = pulumi.output_property("matchVariables")
    """
    One or more `match_variables` blocks as defined below.
    """
    negation_condition: Optional[bool] = pulumi.output_property("negationCondition")
    """
    Describes if this is negate condition or not
    """
    operator: str = pulumi.output_property("operator")
    """
    Describes operator to be matched.
    """
    transforms: Optional[List[str]] = pulumi.output_property("transforms")
    """
    A list of transformations to do before the match is attempted.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyCustomRuleMatchConditionMatchVariable(dict):
    selector: Optional[str] = pulumi.output_property("selector")
    """
    Describes field of the matchVariable collection
    """
    variable_name: str = pulumi.output_property("variableName")
    """
    The name of the Match Variable
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyManagedRules(dict):
    exclusions: Optional[List['outputs.PolicyManagedRulesExclusion']] = pulumi.output_property("exclusions")
    """
    One or more `exclusion` block defined below.
    """
    managed_rule_sets: List['outputs.PolicyManagedRulesManagedRuleSet'] = pulumi.output_property("managedRuleSets")
    """
    One or more `managed_rule_set` block defined below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyManagedRulesExclusion(dict):
    match_variable: str = pulumi.output_property("matchVariable")
    selector: str = pulumi.output_property("selector")
    """
    Describes field of the matchVariable collection.
    """
    selector_match_operator: str = pulumi.output_property("selectorMatchOperator")
    """
    Describes operator to be matched. Possible values: `Contains`, `EndsWith`, `Equals`, `EqualsAny`, `StartsWith`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyManagedRulesManagedRuleSet(dict):
    rule_group_overrides: Optional[List['outputs.PolicyManagedRulesManagedRuleSetRuleGroupOverride']] = pulumi.output_property("ruleGroupOverrides")
    """
    One or more `rule_group_override` block defined below.
    """
    type: Optional[str] = pulumi.output_property("type")
    """
    The rule set type. Possible values: `Microsoft_BotManagerRuleSet` and `OWASP`.
    """
    version: str = pulumi.output_property("version")
    """
    The rule set version. Possible values: `0.1`, `1.0`, `2.2.9`, `3.0` and `3.1`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyManagedRulesManagedRuleSetRuleGroupOverride(dict):
    disabled_rules: List[str] = pulumi.output_property("disabledRules")
    """
    One or more Rule ID's
    """
    rule_group_name: str = pulumi.output_property("ruleGroupName")
    """
    The name of the Rule Group
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyPolicySettings(dict):
    enabled: Optional[bool] = pulumi.output_property("enabled")
    """
    Describes if the policy is in enabled state or disabled state. Defaults to `Enabled`.
    """
    file_upload_limit_in_mb: Optional[float] = pulumi.output_property("fileUploadLimitInMb")
    max_request_body_size_in_kb: Optional[float] = pulumi.output_property("maxRequestBodySizeInKb")
    mode: Optional[str] = pulumi.output_property("mode")
    """
    Describes if it is in detection mode or prevention mode at the policy level. Defaults to `Prevention`.
    """
    request_body_check: Optional[bool] = pulumi.output_property("requestBodyCheck")
    """
    Is Request Body Inspection enabled? Defaults to `true`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


