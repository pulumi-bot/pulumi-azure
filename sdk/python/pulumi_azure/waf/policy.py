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

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 managed_rules: pulumi.Input['PolicyManagedRulesArgs'],
                 resource_group_name: pulumi.Input[str],
                 custom_rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyCustomRuleArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_settings: Optional[pulumi.Input['PolicyPolicySettingsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input['PolicyManagedRulesArgs'] managed_rules: A `managed_rules` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyCustomRuleArgs']]] custom_rules: One or more `custom_rules` blocks as defined below.
        :param pulumi.Input[str] location: Resource location. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input['PolicyPolicySettingsArgs'] policy_settings: A `policy_settings` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Web Application Firewall Policy.
        """
        pulumi.set(__self__, "managed_rules", managed_rules)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if custom_rules is not None:
            pulumi.set(__self__, "custom_rules", custom_rules)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_settings is not None:
            pulumi.set(__self__, "policy_settings", policy_settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="managedRules")
    def managed_rules(self) -> pulumi.Input['PolicyManagedRulesArgs']:
        """
        A `managed_rules` blocks as defined below.
        """
        return pulumi.get(self, "managed_rules")

    @managed_rules.setter
    def managed_rules(self, value: pulumi.Input['PolicyManagedRulesArgs']):
        pulumi.set(self, "managed_rules", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="customRules")
    def custom_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyCustomRuleArgs']]]]:
        """
        One or more `custom_rules` blocks as defined below.
        """
        return pulumi.get(self, "custom_rules")

    @custom_rules.setter
    def custom_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyCustomRuleArgs']]]]):
        pulumi.set(self, "custom_rules", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Resource location. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policySettings")
    def policy_settings(self) -> Optional[pulumi.Input['PolicyPolicySettingsArgs']]:
        """
        A `policy_settings` block as defined below.
        """
        return pulumi.get(self, "policy_settings")

    @policy_settings.setter
    def policy_settings(self, value: Optional[pulumi.Input['PolicyPolicySettingsArgs']]):
        pulumi.set(self, "policy_settings", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the Web Application Firewall Policy.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Azure Web Application Firewall Policy instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_policy = azure.waf.Policy("examplePolicy",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            custom_rules=[
                azure.waf.PolicyCustomRuleArgs(
                    name="Rule1",
                    priority=1,
                    rule_type="MatchRule",
                    match_conditions=[azure.waf.PolicyCustomRuleMatchConditionArgs(
                        match_variables=[azure.waf.PolicyCustomRuleMatchConditionMatchVariableArgs(
                            variable_name="RemoteAddr",
                        )],
                        operator="IPMatch",
                        negation_condition=False,
                        match_values=[
                            "192.168.1.0/24",
                            "10.0.0.0/24",
                        ],
                    )],
                    action="Block",
                ),
                azure.waf.PolicyCustomRuleArgs(
                    name="Rule2",
                    priority=2,
                    rule_type="MatchRule",
                    match_conditions=[
                        azure.waf.PolicyCustomRuleMatchConditionArgs(
                            match_variables=[azure.waf.PolicyCustomRuleMatchConditionMatchVariableArgs(
                                variable_name="RemoteAddr",
                            )],
                            operator="IPMatch",
                            negation_condition=False,
                            match_values=["192.168.1.0/24"],
                        ),
                        azure.waf.PolicyCustomRuleMatchConditionArgs(
                            match_variables=[azure.waf.PolicyCustomRuleMatchConditionMatchVariableArgs(
                                variable_name="RequestHeaders",
                                selector="UserAgent",
                            )],
                            operator="Contains",
                            negation_condition=False,
                            match_values=["Windows"],
                        ),
                    ],
                    action="Block",
                ),
            ],
            policy_settings=azure.waf.PolicyPolicySettingsArgs(
                enabled=True,
                mode="Prevention",
                request_body_check=True,
                file_upload_limit_in_mb=100,
                max_request_body_size_in_kb=128,
            ),
            managed_rules=azure.waf.PolicyManagedRulesArgs(
                exclusions=[
                    azure.waf.PolicyManagedRulesExclusionArgs(
                        match_variable="RequestHeaderNames",
                        selector="x-company-secret-header",
                        selector_match_operator="Equals",
                    ),
                    azure.waf.PolicyManagedRulesExclusionArgs(
                        match_variable="RequestCookieNames",
                        selector="too-tasty",
                        selector_match_operator="EndsWith",
                    ),
                ],
                managed_rule_sets=[azure.waf.PolicyManagedRulesManagedRuleSetArgs(
                    type="OWASP",
                    version="3.1",
                    rule_group_overrides=[azure.waf.PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs(
                        rule_group_name="REQUEST-920-PROTOCOL-ENFORCEMENT",
                        disabled_rules=[
                            "920300",
                            "920440",
                        ],
                    )],
                )],
            ))
        ```

        ## Import

        Web Application Firewall Policy can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:waf/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/example-wafpolicy
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyCustomRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_rules: Optional[pulumi.Input[pulumi.InputType['PolicyManagedRulesArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_settings: Optional[pulumi.Input[pulumi.InputType['PolicyPolicySettingsArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Azure Web Application Firewall Policy instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_policy = azure.waf.Policy("examplePolicy",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            custom_rules=[
                azure.waf.PolicyCustomRuleArgs(
                    name="Rule1",
                    priority=1,
                    rule_type="MatchRule",
                    match_conditions=[azure.waf.PolicyCustomRuleMatchConditionArgs(
                        match_variables=[azure.waf.PolicyCustomRuleMatchConditionMatchVariableArgs(
                            variable_name="RemoteAddr",
                        )],
                        operator="IPMatch",
                        negation_condition=False,
                        match_values=[
                            "192.168.1.0/24",
                            "10.0.0.0/24",
                        ],
                    )],
                    action="Block",
                ),
                azure.waf.PolicyCustomRuleArgs(
                    name="Rule2",
                    priority=2,
                    rule_type="MatchRule",
                    match_conditions=[
                        azure.waf.PolicyCustomRuleMatchConditionArgs(
                            match_variables=[azure.waf.PolicyCustomRuleMatchConditionMatchVariableArgs(
                                variable_name="RemoteAddr",
                            )],
                            operator="IPMatch",
                            negation_condition=False,
                            match_values=["192.168.1.0/24"],
                        ),
                        azure.waf.PolicyCustomRuleMatchConditionArgs(
                            match_variables=[azure.waf.PolicyCustomRuleMatchConditionMatchVariableArgs(
                                variable_name="RequestHeaders",
                                selector="UserAgent",
                            )],
                            operator="Contains",
                            negation_condition=False,
                            match_values=["Windows"],
                        ),
                    ],
                    action="Block",
                ),
            ],
            policy_settings=azure.waf.PolicyPolicySettingsArgs(
                enabled=True,
                mode="Prevention",
                request_body_check=True,
                file_upload_limit_in_mb=100,
                max_request_body_size_in_kb=128,
            ),
            managed_rules=azure.waf.PolicyManagedRulesArgs(
                exclusions=[
                    azure.waf.PolicyManagedRulesExclusionArgs(
                        match_variable="RequestHeaderNames",
                        selector="x-company-secret-header",
                        selector_match_operator="Equals",
                    ),
                    azure.waf.PolicyManagedRulesExclusionArgs(
                        match_variable="RequestCookieNames",
                        selector="too-tasty",
                        selector_match_operator="EndsWith",
                    ),
                ],
                managed_rule_sets=[azure.waf.PolicyManagedRulesManagedRuleSetArgs(
                    type="OWASP",
                    version="3.1",
                    rule_group_overrides=[azure.waf.PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs(
                        rule_group_name="REQUEST-920-PROTOCOL-ENFORCEMENT",
                        disabled_rules=[
                            "920300",
                            "920440",
                        ],
                    )],
                )],
            ))
        ```

        ## Import

        Web Application Firewall Policy can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:waf/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/example-wafpolicy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyCustomRuleArgs']]]] custom_rules: One or more `custom_rules` blocks as defined below.
        :param pulumi.Input[str] location: Resource location. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyManagedRulesArgs']] managed_rules: A `managed_rules` blocks as defined below.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyPolicySettingsArgs']] policy_settings: A `policy_settings` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Web Application Firewall Policy.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyCustomRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_rules: Optional[pulumi.Input[pulumi.InputType['PolicyManagedRulesArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_settings: Optional[pulumi.Input[pulumi.InputType['PolicyPolicySettingsArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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

            __props__['custom_rules'] = custom_rules
            __props__['location'] = location
            if managed_rules is None and not opts.urn:
                raise TypeError("Missing required property 'managed_rules'")
            __props__['managed_rules'] = managed_rules
            __props__['name'] = name
            __props__['policy_settings'] = policy_settings
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Policy, __self__).__init__(
            'azure:waf/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyCustomRuleArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            managed_rules: Optional[pulumi.Input[pulumi.InputType['PolicyManagedRulesArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_settings: Optional[pulumi.Input[pulumi.InputType['PolicyPolicySettingsArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyCustomRuleArgs']]]] custom_rules: One or more `custom_rules` blocks as defined below.
        :param pulumi.Input[str] location: Resource location. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyManagedRulesArgs']] managed_rules: A `managed_rules` blocks as defined below.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyPolicySettingsArgs']] policy_settings: A `policy_settings` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Web Application Firewall Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_rules"] = custom_rules
        __props__["location"] = location
        __props__["managed_rules"] = managed_rules
        __props__["name"] = name
        __props__["policy_settings"] = policy_settings
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customRules")
    def custom_rules(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyCustomRule']]]:
        """
        One or more `custom_rules` blocks as defined below.
        """
        return pulumi.get(self, "custom_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedRules")
    def managed_rules(self) -> pulumi.Output['outputs.PolicyManagedRules']:
        """
        A `managed_rules` blocks as defined below.
        """
        return pulumi.get(self, "managed_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policySettings")
    def policy_settings(self) -> pulumi.Output[Optional['outputs.PolicyPolicySettings']]:
        """
        A `policy_settings` block as defined below.
        """
        return pulumi.get(self, "policy_settings")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the Web Application Firewall Policy.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

