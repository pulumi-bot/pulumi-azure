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

__all__ = ['Assignment']


class Assignment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enforcement_mode: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['AssignmentIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 policy_definition_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures the specified Policy Definition at the specified Scope. Also, Policy Set Definitions are supported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description to use for this Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enforcement_mode: Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
        :param pulumi.Input[pulumi.InputType['AssignmentIdentityArgs']] identity: An `identity` block.
        :param pulumi.Input[str] location: The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
        :param pulumi.Input[str] name: The name of the Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] not_scopes: A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_definition_id: The ID of the Policy Definition to be applied at the specified Scope.
        :param pulumi.Input[str] scope: The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
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
            __props__['display_name'] = display_name
            __props__['enforcement_mode'] = enforcement_mode
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['not_scopes'] = not_scopes
            __props__['parameters'] = parameters
            if policy_definition_id is None:
                raise TypeError("Missing required property 'policy_definition_id'")
            __props__['policy_definition_id'] = policy_definition_id
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
        super(Assignment, __self__).__init__(
            'azure:policy/assignment:Assignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enforcement_mode: Optional[pulumi.Input[bool]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['AssignmentIdentityArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            not_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            parameters: Optional[pulumi.Input[str]] = None,
            policy_definition_id: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'Assignment':
        """
        Get an existing Assignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description to use for this Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enforcement_mode: Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
        :param pulumi.Input[pulumi.InputType['AssignmentIdentityArgs']] identity: An `identity` block.
        :param pulumi.Input[str] location: The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
        :param pulumi.Input[str] name: The name of the Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] not_scopes: A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_definition_id: The ID of the Policy Definition to be applied at the specified Scope.
        :param pulumi.Input[str] scope: The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["enforcement_mode"] = enforcement_mode
        __props__["identity"] = identity
        __props__["location"] = location
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["not_scopes"] = not_scopes
        __props__["parameters"] = parameters
        __props__["policy_definition_id"] = policy_definition_id
        __props__["scope"] = scope
        return Assignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description to use for this Policy Assignment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enforcementMode")
    def enforcement_mode(self) -> pulumi.Output[Optional[bool]]:
        """
        Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
        """
        return pulumi.get(self, "enforcement_mode")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output['outputs.AssignmentIdentity']:
        """
        An `identity` block.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[str]:
        """
        The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Policy Assignment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notScopes")
    def not_scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
        """
        return pulumi.get(self, "not_scopes")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[str]]:
        """
        Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> pulumi.Output[str]:
        """
        The ID of the Policy Definition to be applied at the specified Scope.
        """
        return pulumi.get(self, "policy_definition_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "scope")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

