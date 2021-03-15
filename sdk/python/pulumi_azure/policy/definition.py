# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['DefinitionArgs', 'Definition']

@pulumi.input_type
class DefinitionArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 mode: pulumi.Input[str],
                 policy_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 management_group_id: Optional[pulumi.Input[str]] = None,
                 management_group_name: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 policy_rule: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Definition resource.
        :param pulumi.Input[str] display_name: The display name of the policy definition.
        :param pulumi.Input[str] mode: The policy mode that allows you to specify which resource
               types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        :param pulumi.Input[str] policy_type: The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of the policy definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy definition. This
               is a JSON string representing additional metadata that should be stored
               with the policy definition.
        :param pulumi.Input[str] name: The name of the policy definition. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field
               is a JSON string that allows you to parameterize your policy definition.
        :param pulumi.Input[str] policy_rule: The policy rule for the policy definition. This
               is a JSON string representing the rule that contains an if and
               a then block.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "policy_type", policy_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if management_group_id is not None:
            warnings.warn("""Deprecated in favour of `management_group_name`""", DeprecationWarning)
            pulumi.log.warn("""management_group_id is deprecated: Deprecated in favour of `management_group_name`""")
        if management_group_id is not None:
            pulumi.set(__self__, "management_group_id", management_group_id)
        if management_group_name is not None:
            pulumi.set(__self__, "management_group_name", management_group_name)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if policy_rule is not None:
            pulumi.set(__self__, "policy_rule", policy_rule)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the policy definition.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        The policy mode that allows you to specify which resource
        types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Input[str]:
        """
        The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the policy definition.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="managementGroupId")
    def management_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "management_group_id")

    @management_group_id.setter
    def management_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "management_group_id", value)

    @property
    @pulumi.getter(name="managementGroupName")
    def management_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "management_group_name")

    @management_group_name.setter
    def management_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "management_group_name", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[str]]:
        """
        The metadata for the policy definition. This
        is a JSON string representing additional metadata that should be stored
        with the policy definition.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy definition. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[str]]:
        """
        Parameters for the policy definition. This field
        is a JSON string that allows you to parameterize your policy definition.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="policyRule")
    def policy_rule(self) -> Optional[pulumi.Input[str]]:
        """
        The policy rule for the policy definition. This
        is a JSON string representing the rule that contains an if and
        a then block.
        """
        return pulumi.get(self, "policy_rule")

    @policy_rule.setter
    def policy_rule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_rule", value)


class Definition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a policy rule definition on a management group or your provider subscription.

        Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        policy = azure.policy.Definition("policy",
            display_name="acceptance test policy definition",
            metadata=\"\"\"    {
            "category": "General"
            }


        \"\"\",
            mode="Indexed",
            parameters=\"\"\"	{
            "allowedLocations": {
              "type": "Array",
              "metadata": {
                "description": "The list of allowed locations for resources.",
                "displayName": "Allowed locations",
                "strongType": "location"
              }
            }
          }

        \"\"\",
            policy_rule=\"\"\"	{
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
            policy_type="Custom")
        ```

        ## Import

        Policy Definitions can be imported using the `policy name`, e.g.

        ```sh
         $ pulumi import azure:policy/definition:Definition examplePolicy /subscriptions/<SUBSCRIPTION_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
        ```

         or

        ```sh
         $ pulumi import azure:policy/definition:Definition examplePolicy /providers/Microsoft.Management/managementgroups/<MANGAGEMENT_GROUP_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
        ```

        :param str resource_name: The name of the resource.
        :param DefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 management_group_id: Optional[pulumi.Input[str]] = None,
                 management_group_name: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 policy_rule: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a policy rule definition on a management group or your provider subscription.

        Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        policy = azure.policy.Definition("policy",
            display_name="acceptance test policy definition",
            metadata=\"\"\"    {
            "category": "General"
            }


        \"\"\",
            mode="Indexed",
            parameters=\"\"\"	{
            "allowedLocations": {
              "type": "Array",
              "metadata": {
                "description": "The list of allowed locations for resources.",
                "displayName": "Allowed locations",
                "strongType": "location"
              }
            }
          }

        \"\"\",
            policy_rule=\"\"\"	{
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
            policy_type="Custom")
        ```

        ## Import

        Policy Definitions can be imported using the `policy name`, e.g.

        ```sh
         $ pulumi import azure:policy/definition:Definition examplePolicy /subscriptions/<SUBSCRIPTION_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
        ```

         or

        ```sh
         $ pulumi import azure:policy/definition:Definition examplePolicy /providers/Microsoft.Management/managementgroups/<MANGAGEMENT_GROUP_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy definition.
        :param pulumi.Input[str] display_name: The display name of the policy definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy definition. This
               is a JSON string representing additional metadata that should be stored
               with the policy definition.
        :param pulumi.Input[str] mode: The policy mode that allows you to specify which resource
               types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        :param pulumi.Input[str] name: The name of the policy definition. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field
               is a JSON string that allows you to parameterize your policy definition.
        :param pulumi.Input[str] policy_rule: The policy rule for the policy definition. This
               is a JSON string representing the rule that contains an if and
               a then block.
        :param pulumi.Input[str] policy_type: The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 management_group_id: Optional[pulumi.Input[str]] = None,
                 management_group_name: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 policy_rule: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
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

            __props__['description'] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if management_group_id is not None and not opts.urn:
                warnings.warn("""Deprecated in favour of `management_group_name`""", DeprecationWarning)
                pulumi.log.warn("""management_group_id is deprecated: Deprecated in favour of `management_group_name`""")
            __props__['management_group_id'] = management_group_id
            __props__['management_group_name'] = management_group_name
            __props__['metadata'] = metadata
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__['mode'] = mode
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['policy_rule'] = policy_rule
            if policy_type is None and not opts.urn:
                raise TypeError("Missing required property 'policy_type'")
            __props__['policy_type'] = policy_type
        super(Definition, __self__).__init__(
            'azure:policy/definition:Definition',
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
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[str]] = None,
            policy_rule: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None) -> 'Definition':
        """
        Get an existing Definition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy definition.
        :param pulumi.Input[str] display_name: The display name of the policy definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy definition. This
               is a JSON string representing additional metadata that should be stored
               with the policy definition.
        :param pulumi.Input[str] mode: The policy mode that allows you to specify which resource
               types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        :param pulumi.Input[str] name: The name of the policy definition. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field
               is a JSON string that allows you to parameterize your policy definition.
        :param pulumi.Input[str] policy_rule: The policy rule for the policy definition. This
               is a JSON string representing the rule that contains an if and
               a then block.
        :param pulumi.Input[str] policy_type: The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["management_group_id"] = management_group_id
        __props__["management_group_name"] = management_group_name
        __props__["metadata"] = metadata
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["policy_rule"] = policy_rule
        __props__["policy_type"] = policy_type
        return Definition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the policy definition.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the policy definition.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="managementGroupId")
    def management_group_id(self) -> pulumi.Output[str]:
        """
        The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "management_group_id")

    @property
    @pulumi.getter(name="managementGroupName")
    def management_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "management_group_name")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[str]:
        """
        The metadata for the policy definition. This
        is a JSON string representing additional metadata that should be stored
        with the policy definition.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        The policy mode that allows you to specify which resource
        types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy definition. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[str]]:
        """
        Parameters for the policy definition. This field
        is a JSON string that allows you to parameterize your policy definition.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyRule")
    def policy_rule(self) -> pulumi.Output[Optional[str]]:
        """
        The policy rule for the policy definition. This
        is a JSON string representing the rule that contains an if and
        a then block.
        """
        return pulumi.get(self, "policy_rule")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        """
        The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "policy_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

