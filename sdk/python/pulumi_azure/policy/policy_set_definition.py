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

__all__ = ['PolicySetDefinition']


class PolicySetDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 management_group_id: Optional[pulumi.Input[str]] = None,
                 management_group_name: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 policy_definition_references: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicySetDefinitionPolicyDefinitionReferenceArgs']]]]] = None,
                 policy_definitions: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a policy set definition.

        > **NOTE:**  Policy set definitions (also known as policy initiatives) do not take effect until they are assigned to a scope using a Policy Set Assignment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy set definition.
        :param pulumi.Input[str] display_name: The display name of the policy set definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
        :param pulumi.Input[str] name: The name of the policy set definition. Changing this forces a new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicySetDefinitionPolicyDefinitionReferenceArgs']]]] policy_definition_references: One or more `policy_definition_reference` blocks as defined below.
        :param pulumi.Input[str] policy_definitions: The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
        :param pulumi.Input[str] policy_type: The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
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
            if management_group_id is not None:
                warnings.warn("Deprecated in favour of `management_group_name`", DeprecationWarning)
                pulumi.log.warn("management_group_id is deprecated: Deprecated in favour of `management_group_name`")
            __props__['management_group_id'] = management_group_id
            __props__['management_group_name'] = management_group_name
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['policy_definition_references'] = policy_definition_references
            if policy_definitions is not None:
                warnings.warn("Deprecated in favor of `policy_definition_reference`", DeprecationWarning)
                pulumi.log.warn("policy_definitions is deprecated: Deprecated in favor of `policy_definition_reference`")
            __props__['policy_definitions'] = policy_definitions
            if policy_type is None:
                raise TypeError("Missing required property 'policy_type'")
            __props__['policy_type'] = policy_type
        super(PolicySetDefinition, __self__).__init__(
            'azure:policy/policySetDefinition:PolicySetDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            management_group_id: Optional[pulumi.Input[str]] = None,
            management_group_name: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[str]] = None,
            policy_definition_references: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicySetDefinitionPolicyDefinitionReferenceArgs']]]]] = None,
            policy_definitions: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None) -> 'PolicySetDefinition':
        """
        Get an existing PolicySetDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy set definition.
        :param pulumi.Input[str] display_name: The display name of the policy set definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
        :param pulumi.Input[str] name: The name of the policy set definition. Changing this forces a new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicySetDefinitionPolicyDefinitionReferenceArgs']]]] policy_definition_references: One or more `policy_definition_reference` blocks as defined below.
        :param pulumi.Input[str] policy_definitions: The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
        :param pulumi.Input[str] policy_type: The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["management_group_id"] = management_group_id
        __props__["management_group_name"] = management_group_name
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["policy_definition_references"] = policy_definition_references
        __props__["policy_definitions"] = policy_definitions
        __props__["policy_type"] = policy_type
        return PolicySetDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the policy set definition.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the policy set definition.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="managementGroupId")
    def management_group_id(self) -> pulumi.Output[str]:
        """
        The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "management_group_id")

    @property
    @pulumi.getter(name="managementGroupName")
    def management_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "management_group_name")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[str]:
        """
        The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy set definition. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[str]]:
        """
        Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyDefinitionReferences")
    def policy_definition_references(self) -> pulumi.Output[Sequence['outputs.PolicySetDefinitionPolicyDefinitionReference']]:
        """
        One or more `policy_definition_reference` blocks as defined below.
        """
        return pulumi.get(self, "policy_definition_references")

    @property
    @pulumi.getter(name="policyDefinitions")
    def policy_definitions(self) -> pulumi.Output[str]:
        """
        The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
        """
        return pulumi.get(self, "policy_definitions")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        """
        The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "policy_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

