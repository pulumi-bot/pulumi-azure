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

__all__ = ['Definition']


class Definition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorizations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefinitionAuthorizationArgs']]]]] = None,
                 create_ui_definition: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 lock_level: Optional[pulumi.Input[str]] = None,
                 main_template: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_enabled: Optional[pulumi.Input[bool]] = None,
                 package_file_uri: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Managed Application Definition.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_definition = azure.managedapplication.Definition("exampleDefinition",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            lock_level="ReadOnly",
            package_file_uri="https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip",
            display_name="TestManagedApplicationDefinition",
            description="Test Managed Application Definition",
            authorizations=[azure.managedapplication.DefinitionAuthorizationArgs(
                service_principal_id=current.object_id,
                role_definition_id="a094b430-dad3-424d-ae58-13f72fd72591",
            )])
        ```

        ## Import

        Managed Application Definition can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:managedapplication/definition:Definition example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applicationDefinitions/appDefinition1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefinitionAuthorizationArgs']]]] authorizations: One or more `authorization` block defined below.
        :param pulumi.Input[str] create_ui_definition: Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
        :param pulumi.Input[str] description: Specifies the managed application definition description.
        :param pulumi.Input[str] display_name: Specifies the managed application definition display name.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lock_level: Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] main_template: Specifies the inline main template json which has resources to be provisioned.
        :param pulumi.Input[str] name: Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] package_enabled: Is the package enabled? Defaults to `true`.
        :param pulumi.Input[str] package_file_uri: Specifies the managed application definition package file Uri.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['authorizations'] = authorizations
            __props__['create_ui_definition'] = create_ui_definition
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['location'] = location
            if lock_level is None:
                raise TypeError("Missing required property 'lock_level'")
            __props__['lock_level'] = lock_level
            __props__['main_template'] = main_template
            __props__['name'] = name
            __props__['package_enabled'] = package_enabled
            __props__['package_file_uri'] = package_file_uri
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Definition, __self__).__init__(
            'azure:managedapplication/definition:Definition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorizations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefinitionAuthorizationArgs']]]]] = None,
            create_ui_definition: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            lock_level: Optional[pulumi.Input[str]] = None,
            main_template: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            package_enabled: Optional[pulumi.Input[bool]] = None,
            package_file_uri: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Definition':
        """
        Get an existing Definition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefinitionAuthorizationArgs']]]] authorizations: One or more `authorization` block defined below.
        :param pulumi.Input[str] create_ui_definition: Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
        :param pulumi.Input[str] description: Specifies the managed application definition description.
        :param pulumi.Input[str] display_name: Specifies the managed application definition display name.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lock_level: Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] main_template: Specifies the inline main template json which has resources to be provisioned.
        :param pulumi.Input[str] name: Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] package_enabled: Is the package enabled? Defaults to `true`.
        :param pulumi.Input[str] package_file_uri: Specifies the managed application definition package file Uri.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorizations"] = authorizations
        __props__["create_ui_definition"] = create_ui_definition
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["location"] = location
        __props__["lock_level"] = lock_level
        __props__["main_template"] = main_template
        __props__["name"] = name
        __props__["package_enabled"] = package_enabled
        __props__["package_file_uri"] = package_file_uri
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Definition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorizations(self) -> pulumi.Output[Optional[Sequence['outputs.DefinitionAuthorization']]]:
        """
        One or more `authorization` block defined below.
        """
        return pulumi.get(self, "authorizations")

    @property
    @pulumi.getter(name="createUiDefinition")
    def create_ui_definition(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
        """
        return pulumi.get(self, "create_ui_definition")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the managed application definition description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Specifies the managed application definition display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="lockLevel")
    def lock_level(self) -> pulumi.Output[str]:
        """
        Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lock_level")

    @property
    @pulumi.getter(name="mainTemplate")
    def main_template(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the inline main template json which has resources to be provisioned.
        """
        return pulumi.get(self, "main_template")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageEnabled")
    def package_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the package enabled? Defaults to `true`.
        """
        return pulumi.get(self, "package_enabled")

    @property
    @pulumi.getter(name="packageFileUri")
    def package_file_uri(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the managed application definition package file Uri.
        """
        return pulumi.get(self, "package_file_uri")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

