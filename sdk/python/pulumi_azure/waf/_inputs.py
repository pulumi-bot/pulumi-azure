# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'PolicyCustomRuleArgs',
    'PolicyCustomRuleMatchConditionArgs',
    'PolicyCustomRuleMatchConditionMatchVariableArgs',
    'PolicyManagedRulesArgs',
    'PolicyManagedRulesExclusionArgs',
    'PolicyManagedRulesManagedRuleSetArgs',
    'PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs',
    'PolicyPolicySettingsArgs',
]

@pulumi.input_type
class PolicyCustomRuleArgs:
    action: pulumi.Input[str] = pulumi.input_property("action")
    """
    Type of action.
    """
    match_conditions: pulumi.Input[List[pulumi.Input['PolicyCustomRuleMatchConditionArgs']]] = pulumi.input_property("matchConditions")
    """
    One or more `match_conditions` blocks as defined below.
    """
    priority: pulumi.Input[float] = pulumi.input_property("priority")
    """
    Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
    """
    rule_type: pulumi.Input[str] = pulumi.input_property("ruleType")
    """
    Describes the type of rule.
    """
    name: Optional[pulumi.Input[str]] = pulumi.input_property("name")
    """
    Gets name of the resource that is unique within a policy. This name can be used to access the resource.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, action: pulumi.Input[str], match_conditions: pulumi.Input[List[pulumi.Input['PolicyCustomRuleMatchConditionArgs']]], priority: pulumi.Input[float], rule_type: pulumi.Input[str], name: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] action: Type of action.
        :param pulumi.Input[List[pulumi.Input['PolicyCustomRuleMatchConditionArgs']]] match_conditions: One or more `match_conditions` blocks as defined below.
        :param pulumi.Input[float] priority: Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
        :param pulumi.Input[str] rule_type: Describes the type of rule.
        :param pulumi.Input[str] name: Gets name of the resource that is unique within a policy. This name can be used to access the resource.
        """
        __self__.action = action
        __self__.match_conditions = match_conditions
        __self__.priority = priority
        __self__.rule_type = rule_type
        __self__.name = name

@pulumi.input_type
class PolicyCustomRuleMatchConditionArgs:
    match_values: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("matchValues")
    """
    A list of match values.
    """
    match_variables: pulumi.Input[List[pulumi.Input['PolicyCustomRuleMatchConditionMatchVariableArgs']]] = pulumi.input_property("matchVariables")
    """
    One or more `match_variables` blocks as defined below.
    """
    operator: pulumi.Input[str] = pulumi.input_property("operator")
    """
    Describes operator to be matched.
    """
    negation_condition: Optional[pulumi.Input[bool]] = pulumi.input_property("negationCondition")
    """
    Describes if this is negate condition or not
    """
    transforms: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("transforms")
    """
    A list of transformations to do before the match is attempted.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, match_values: pulumi.Input[List[pulumi.Input[str]]], match_variables: pulumi.Input[List[pulumi.Input['PolicyCustomRuleMatchConditionMatchVariableArgs']]], operator: pulumi.Input[str], negation_condition: Optional[pulumi.Input[bool]] = None, transforms: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] match_values: A list of match values.
        :param pulumi.Input[List[pulumi.Input['PolicyCustomRuleMatchConditionMatchVariableArgs']]] match_variables: One or more `match_variables` blocks as defined below.
        :param pulumi.Input[str] operator: Describes operator to be matched.
        :param pulumi.Input[bool] negation_condition: Describes if this is negate condition or not
        :param pulumi.Input[List[pulumi.Input[str]]] transforms: A list of transformations to do before the match is attempted.
        """
        __self__.match_values = match_values
        __self__.match_variables = match_variables
        __self__.operator = operator
        __self__.negation_condition = negation_condition
        __self__.transforms = transforms

@pulumi.input_type
class PolicyCustomRuleMatchConditionMatchVariableArgs:
    variable_name: pulumi.Input[str] = pulumi.input_property("variableName")
    """
    The name of the Match Variable
    """
    selector: Optional[pulumi.Input[str]] = pulumi.input_property("selector")
    """
    Describes field of the matchVariable collection
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, variable_name: pulumi.Input[str], selector: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] variable_name: The name of the Match Variable
        :param pulumi.Input[str] selector: Describes field of the matchVariable collection
        """
        __self__.variable_name = variable_name
        __self__.selector = selector

@pulumi.input_type
class PolicyManagedRulesArgs:
    managed_rule_sets: pulumi.Input[List[pulumi.Input['PolicyManagedRulesManagedRuleSetArgs']]] = pulumi.input_property("managedRuleSets")
    """
    One or more `managed_rule_set` block defined below.
    """
    exclusions: Optional[pulumi.Input[List[pulumi.Input['PolicyManagedRulesExclusionArgs']]]] = pulumi.input_property("exclusions")
    """
    One or more `exclusion` block defined below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, managed_rule_sets: pulumi.Input[List[pulumi.Input['PolicyManagedRulesManagedRuleSetArgs']]], exclusions: Optional[pulumi.Input[List[pulumi.Input['PolicyManagedRulesExclusionArgs']]]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input['PolicyManagedRulesManagedRuleSetArgs']]] managed_rule_sets: One or more `managed_rule_set` block defined below.
        :param pulumi.Input[List[pulumi.Input['PolicyManagedRulesExclusionArgs']]] exclusions: One or more `exclusion` block defined below.
        """
        __self__.managed_rule_sets = managed_rule_sets
        __self__.exclusions = exclusions

@pulumi.input_type
class PolicyManagedRulesExclusionArgs:
    match_variable: pulumi.Input[str] = pulumi.input_property("matchVariable")
    selector: pulumi.Input[str] = pulumi.input_property("selector")
    """
    Describes field of the matchVariable collection.
    """
    selector_match_operator: pulumi.Input[str] = pulumi.input_property("selectorMatchOperator")
    """
    Describes operator to be matched. Possible values: `Contains`, `EndsWith`, `Equals`, `EqualsAny`, `StartsWith`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, match_variable: pulumi.Input[str], selector: pulumi.Input[str], selector_match_operator: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] selector: Describes field of the matchVariable collection.
        :param pulumi.Input[str] selector_match_operator: Describes operator to be matched. Possible values: `Contains`, `EndsWith`, `Equals`, `EqualsAny`, `StartsWith`.
        """
        __self__.match_variable = match_variable
        __self__.selector = selector
        __self__.selector_match_operator = selector_match_operator

@pulumi.input_type
class PolicyManagedRulesManagedRuleSetArgs:
    version: pulumi.Input[str] = pulumi.input_property("version")
    """
    The rule set version. Possible values: `0.1`, `1.0`, `2.2.9`, `3.0` and `3.1`.
    """
    rule_group_overrides: Optional[pulumi.Input[List[pulumi.Input['PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs']]]] = pulumi.input_property("ruleGroupOverrides")
    """
    One or more `rule_group_override` block defined below.
    """
    type: Optional[pulumi.Input[str]] = pulumi.input_property("type")
    """
    The rule set type. Possible values: `Microsoft_BotManagerRuleSet` and `OWASP`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, version: pulumi.Input[str], rule_group_overrides: Optional[pulumi.Input[List[pulumi.Input['PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs']]]] = None, type: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] version: The rule set version. Possible values: `0.1`, `1.0`, `2.2.9`, `3.0` and `3.1`.
        :param pulumi.Input[List[pulumi.Input['PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs']]] rule_group_overrides: One or more `rule_group_override` block defined below.
        :param pulumi.Input[str] type: The rule set type. Possible values: `Microsoft_BotManagerRuleSet` and `OWASP`.
        """
        __self__.version = version
        __self__.rule_group_overrides = rule_group_overrides
        __self__.type = type

@pulumi.input_type
class PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs:
    disabled_rules: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("disabledRules")
    """
    One or more Rule ID's
    """
    rule_group_name: pulumi.Input[str] = pulumi.input_property("ruleGroupName")
    """
    The name of the Rule Group
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, disabled_rules: pulumi.Input[List[pulumi.Input[str]]], rule_group_name: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] disabled_rules: One or more Rule ID's
        :param pulumi.Input[str] rule_group_name: The name of the Rule Group
        """
        __self__.disabled_rules = disabled_rules
        __self__.rule_group_name = rule_group_name

@pulumi.input_type
class PolicyPolicySettingsArgs:
    enabled: Optional[pulumi.Input[bool]] = pulumi.input_property("enabled")
    """
    Describes if the policy is in enabled state or disabled state. Defaults to `Enabled`.
    """
    file_upload_limit_in_mb: Optional[pulumi.Input[float]] = pulumi.input_property("fileUploadLimitInMb")
    max_request_body_size_in_kb: Optional[pulumi.Input[float]] = pulumi.input_property("maxRequestBodySizeInKb")
    mode: Optional[pulumi.Input[str]] = pulumi.input_property("mode")
    """
    Describes if it is in detection mode or prevention mode at the policy level. Defaults to `Prevention`.
    """
    request_body_check: Optional[pulumi.Input[bool]] = pulumi.input_property("requestBodyCheck")
    """
    Is Request Body Inspection enabled? Defaults to `true`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, enabled: Optional[pulumi.Input[bool]] = None, file_upload_limit_in_mb: Optional[pulumi.Input[float]] = None, max_request_body_size_in_kb: Optional[pulumi.Input[float]] = None, mode: Optional[pulumi.Input[str]] = None, request_body_check: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[bool] enabled: Describes if the policy is in enabled state or disabled state. Defaults to `Enabled`.
        :param pulumi.Input[str] mode: Describes if it is in detection mode or prevention mode at the policy level. Defaults to `Prevention`.
        :param pulumi.Input[bool] request_body_check: Is Request Body Inspection enabled? Defaults to `true`.
        """
        __self__.enabled = enabled
        __self__.file_upload_limit_in_mb = file_upload_limit_in_mb
        __self__.max_request_body_size_in_kb = max_request_body_size_in_kb
        __self__.mode = mode
        __self__.request_body_check = request_body_check

