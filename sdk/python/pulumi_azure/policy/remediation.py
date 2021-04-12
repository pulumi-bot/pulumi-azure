# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['RemediationArgs', 'Remediation']

@pulumi.input_type
class RemediationArgs:
    def __init__(__self__, *,
                 policy_assignment_id: pulumi.Input[str],
                 scope: pulumi.Input[str],
                 location_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_definition_reference_id: Optional[pulumi.Input[str]] = None,
                 resource_discovery_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Remediation resource.
        :param pulumi.Input[str] policy_assignment_id: The ID of the Policy Assignment that should be remediated.
        :param pulumi.Input[str] scope: The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] location_filters: A list of the resource locations that will be remediated.
        :param pulumi.Input[str] name: The name of the Policy Remediation. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_definition_reference_id: The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] resource_discovery_mode: The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
        """
        pulumi.set(__self__, "policy_assignment_id", policy_assignment_id)
        pulumi.set(__self__, "scope", scope)
        if location_filters is not None:
            pulumi.set(__self__, "location_filters", location_filters)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_definition_reference_id is not None:
            pulumi.set(__self__, "policy_definition_reference_id", policy_definition_reference_id)
        if resource_discovery_mode is not None:
            pulumi.set(__self__, "resource_discovery_mode", resource_discovery_mode)

    @property
    @pulumi.getter(name="policyAssignmentId")
    def policy_assignment_id(self) -> pulumi.Input[str]:
        """
        The ID of the Policy Assignment that should be remediated.
        """
        return pulumi.get(self, "policy_assignment_id")

    @policy_assignment_id.setter
    def policy_assignment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_assignment_id", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        """
        The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="locationFilters")
    def location_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the resource locations that will be remediated.
        """
        return pulumi.get(self, "location_filters")

    @location_filters.setter
    def location_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "location_filters", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Policy Remediation. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        """
        return pulumi.get(self, "policy_definition_reference_id")

    @policy_definition_reference_id.setter
    def policy_definition_reference_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_definition_reference_id", value)

    @property
    @pulumi.getter(name="resourceDiscoveryMode")
    def resource_discovery_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
        """
        return pulumi.get(self, "resource_discovery_mode")

    @resource_discovery_mode.setter
    def resource_discovery_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_discovery_mode", value)


class Remediation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_assignment_id: Optional[pulumi.Input[str]] = None,
                 policy_definition_reference_id: Optional[pulumi.Input[str]] = None,
                 resource_discovery_mode: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Policy Remediation at the specified Scope.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_definition = azure.policy.Definition("exampleDefinition",
            policy_type="Custom",
            mode="All",
            display_name="my-policy-definition",
            policy_rule=\"\"\"    {
            "if": {
              "not": {
                "field": "location",
                "in": "[parameters('allowedLocations')]"
              }
            },
            "then": {
              "effect": "audit"
            }
          }
        \"\"\",
            parameters=\"\"\"    {
            "allowedLocations": {
              "type": "Array",
              "metadata": {
                "description": "The list of allowed locations for resources.",
                "displayName": "Allowed locations",
                "strongType": "location"
              }
            }
          }
        \"\"\")
        example_assignment = azure.policy.Assignment("exampleAssignment",
            scope=example_resource_group.id,
            policy_definition_id=example_definition.id,
            description="Policy Assignment created via an Acceptance Test",
            display_name="My Example Policy Assignment",
            parameters=\"\"\"{
          "allowedLocations": {
            "value": [ "West Europe" ]
          }
        }
        \"\"\")
        example_remediation = azure.policy.Remediation("exampleRemediation",
            scope=example_assignment.scope,
            policy_assignment_id=example_assignment.id,
            location_filters=["West Europe"])
        ```

        ## Import

        Policy Remediations can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

         or

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

         or

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

         or

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] location_filters: A list of the resource locations that will be remediated.
        :param pulumi.Input[str] name: The name of the Policy Remediation. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_assignment_id: The ID of the Policy Assignment that should be remediated.
        :param pulumi.Input[str] policy_definition_reference_id: The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] resource_discovery_mode: The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
        :param pulumi.Input[str] scope: The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RemediationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure Policy Remediation at the specified Scope.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_definition = azure.policy.Definition("exampleDefinition",
            policy_type="Custom",
            mode="All",
            display_name="my-policy-definition",
            policy_rule=\"\"\"    {
            "if": {
              "not": {
                "field": "location",
                "in": "[parameters('allowedLocations')]"
              }
            },
            "then": {
              "effect": "audit"
            }
          }
        \"\"\",
            parameters=\"\"\"    {
            "allowedLocations": {
              "type": "Array",
              "metadata": {
                "description": "The list of allowed locations for resources.",
                "displayName": "Allowed locations",
                "strongType": "location"
              }
            }
          }
        \"\"\")
        example_assignment = azure.policy.Assignment("exampleAssignment",
            scope=example_resource_group.id,
            policy_definition_id=example_definition.id,
            description="Policy Assignment created via an Acceptance Test",
            display_name="My Example Policy Assignment",
            parameters=\"\"\"{
          "allowedLocations": {
            "value": [ "West Europe" ]
          }
        }
        \"\"\")
        example_remediation = azure.policy.Remediation("exampleRemediation",
            scope=example_assignment.scope,
            policy_assignment_id=example_assignment.id,
            location_filters=["West Europe"])
        ```

        ## Import

        Policy Remediations can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

         or

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

         or

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

         or

        ```sh
         $ pulumi import azure:policy/remediation:Remediation example /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.PolicyInsights/remediations/remediation1
        ```

        :param str resource_name: The name of the resource.
        :param RemediationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemediationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_assignment_id: Optional[pulumi.Input[str]] = None,
                 policy_definition_reference_id: Optional[pulumi.Input[str]] = None,
                 resource_discovery_mode: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
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

            __props__['location_filters'] = location_filters
            __props__['name'] = name
            if policy_assignment_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_assignment_id'")
            __props__['policy_assignment_id'] = policy_assignment_id
            __props__['policy_definition_reference_id'] = policy_definition_reference_id
            __props__['resource_discovery_mode'] = resource_discovery_mode
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
        super(Remediation, __self__).__init__(
            'azure:policy/remediation:Remediation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_assignment_id: Optional[pulumi.Input[str]] = None,
            policy_definition_reference_id: Optional[pulumi.Input[str]] = None,
            resource_discovery_mode: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'Remediation':
        """
        Get an existing Remediation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] location_filters: A list of the resource locations that will be remediated.
        :param pulumi.Input[str] name: The name of the Policy Remediation. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_assignment_id: The ID of the Policy Assignment that should be remediated.
        :param pulumi.Input[str] policy_definition_reference_id: The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] resource_discovery_mode: The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
        :param pulumi.Input[str] scope: The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location_filters"] = location_filters
        __props__["name"] = name
        __props__["policy_assignment_id"] = policy_assignment_id
        __props__["policy_definition_reference_id"] = policy_definition_reference_id
        __props__["resource_discovery_mode"] = resource_discovery_mode
        __props__["scope"] = scope
        return Remediation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="locationFilters")
    def location_filters(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of the resource locations that will be remediated.
        """
        return pulumi.get(self, "location_filters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Policy Remediation. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyAssignmentId")
    def policy_assignment_id(self) -> pulumi.Output[str]:
        """
        The ID of the Policy Assignment that should be remediated.
        """
        return pulumi.get(self, "policy_assignment_id")

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> pulumi.Output[Optional[str]]:
        """
        The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        """
        return pulumi.get(self, "policy_definition_reference_id")

    @property
    @pulumi.getter(name="resourceDiscoveryMode")
    def resource_discovery_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
        """
        return pulumi.get(self, "resource_discovery_mode")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        """
        return pulumi.get(self, "scope")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

