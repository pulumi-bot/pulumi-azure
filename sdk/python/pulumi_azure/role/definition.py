# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("azure.role.Definition has been deprecated in favor of azure.authorization.RoleDefinition", DeprecationWarning)

class Definition(pulumi.CustomResource):
    assignable_scopes: pulumi.Output[list]
    """
    One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
    """
    description: pulumi.Output[str]
    """
    A description of the Role Definition.
    """
    name: pulumi.Output[str]
    """
    The name of the Role Definition. Changing this forces a new resource to be created.
    """
    permissions: pulumi.Output[list]
    """
    A `permissions` block as defined below.

      * `actions` (`list`) - One or more Allowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
      * `dataActions` (`list`) - One or more Allowed Data Actions, such as `*`, `Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
      * `notActions` (`list`) - One or more Disallowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
      * `notDataActions` (`list`) - One or more Disallowed Data Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
    """
    role_definition_id: pulumi.Output[str]
    """
    A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
    """
    scope: pulumi.Output[str]
    """
    The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignable_scopes`. Changing this forces a new resource to be created.
    """
    warnings.warn("azure.role.Definition has been deprecated in favor of azure.authorization.RoleDefinition", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, assignable_scopes=None, description=None, name=None, permissions=None, role_definition_id=None, scope=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a custom Role Definition, used to assign Roles to Users/Principals. See ['Understand role definitions'](https://docs.microsoft.com/en-us/azure/role-based-access-control/role-definitions) in the Azure documentation for more details.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        primary = azure.core.get_subscription()
        example = azure.authorization.RoleDefinition("example",
            scope=primary.id,
            description="This is a custom role created",
            permissions=[{
                "actions": ["*"],
                "notActions": [],
            }],
            assignable_scopes=[primary.id])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] assignable_scopes: One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
        :param pulumi.Input[str] description: A description of the Role Definition.
        :param pulumi.Input[str] name: The name of the Role Definition. Changing this forces a new resource to be created.
        :param pulumi.Input[list] permissions: A `permissions` block as defined below.
        :param pulumi.Input[str] role_definition_id: A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope: The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignable_scopes`. Changing this forces a new resource to be created.

        The **permissions** object supports the following:

          * `actions` (`pulumi.Input[list]`) - One or more Allowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
          * `dataActions` (`pulumi.Input[list]`) - One or more Allowed Data Actions, such as `*`, `Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
          * `notActions` (`pulumi.Input[list]`) - One or more Disallowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
          * `notDataActions` (`pulumi.Input[list]`) - One or more Disallowed Data Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        pulumi.log.warn("Definition is deprecated: azure.role.Definition has been deprecated in favor of azure.authorization.RoleDefinition")
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if assignable_scopes is None:
                raise TypeError("Missing required property 'assignable_scopes'")
            __props__['assignable_scopes'] = assignable_scopes
            __props__['description'] = description
            __props__['name'] = name
            if permissions is None:
                raise TypeError("Missing required property 'permissions'")
            __props__['permissions'] = permissions
            __props__['role_definition_id'] = role_definition_id
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
        super(Definition, __self__).__init__(
            'azure:role/definition:Definition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, assignable_scopes=None, description=None, name=None, permissions=None, role_definition_id=None, scope=None):
        """
        Get an existing Definition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] assignable_scopes: One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
        :param pulumi.Input[str] description: A description of the Role Definition.
        :param pulumi.Input[str] name: The name of the Role Definition. Changing this forces a new resource to be created.
        :param pulumi.Input[list] permissions: A `permissions` block as defined below.
        :param pulumi.Input[str] role_definition_id: A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope: The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignable_scopes`. Changing this forces a new resource to be created.

        The **permissions** object supports the following:

          * `actions` (`pulumi.Input[list]`) - One or more Allowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
          * `dataActions` (`pulumi.Input[list]`) - One or more Allowed Data Actions, such as `*`, `Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
          * `notActions` (`pulumi.Input[list]`) - One or more Disallowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
          * `notDataActions` (`pulumi.Input[list]`) - One or more Disallowed Data Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["assignable_scopes"] = assignable_scopes
        __props__["description"] = description
        __props__["name"] = name
        __props__["permissions"] = permissions
        __props__["role_definition_id"] = role_definition_id
        __props__["scope"] = scope
        return Definition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
