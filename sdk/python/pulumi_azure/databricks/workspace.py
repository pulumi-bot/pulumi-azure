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

__all__ = ['Workspace']


class Workspace(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_parameters: Optional[pulumi.Input[pulumi.InputType['WorkspaceCustomParametersArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Databricks Workspace

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_workspace = azure.databricks.Workspace("exampleWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku="standard",
            tags={
                "Environment": "Production",
            })
        ```

        ## Import

        Databrick Workspaces can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:databricks/workspace:Workspace workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WorkspaceCustomParametersArgs']] custom_parameters: A `custom_parameters` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_resource_group_name: The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this forces a new resource to be created.
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

            __props__['custom_parameters'] = custom_parameters
            __props__['location'] = location
            __props__['managed_resource_group_name'] = managed_resource_group_name
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['managed_resource_group_id'] = None
            __props__['workspace_id'] = None
            __props__['workspace_url'] = None
        super(Workspace, __self__).__init__(
            'azure:databricks/workspace:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_parameters: Optional[pulumi.Input[pulumi.InputType['WorkspaceCustomParametersArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            managed_resource_group_id: Optional[pulumi.Input[str]] = None,
            managed_resource_group_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None,
            workspace_url: Optional[pulumi.Input[str]] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WorkspaceCustomParametersArgs']] custom_parameters: A `custom_parameters` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_resource_group_id: The ID of the Managed Resource Group created by the Databricks Workspace.
        :param pulumi.Input[str] managed_resource_group_name: The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workspace_id: The unique identifier of the databricks workspace in Databricks control plane.
        :param pulumi.Input[str] workspace_url: The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_parameters"] = custom_parameters
        __props__["location"] = location
        __props__["managed_resource_group_id"] = managed_resource_group_id
        __props__["managed_resource_group_name"] = managed_resource_group_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["workspace_id"] = workspace_id
        __props__["workspace_url"] = workspace_url
        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customParameters")
    def custom_parameters(self) -> pulumi.Output['outputs.WorkspaceCustomParameters']:
        """
        A `custom_parameters` block as documented below.
        """
        return pulumi.get(self, "custom_parameters")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedResourceGroupId")
    def managed_resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the Managed Resource Group created by the Databricks Workspace.
        """
        return pulumi.get(self, "managed_resource_group_id")

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "managed_resource_group_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[str]:
        """
        The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the databricks workspace in Databricks control plane.
        """
        return pulumi.get(self, "workspace_id")

    @property
    @pulumi.getter(name="workspaceUrl")
    def workspace_url(self) -> pulumi.Output[str]:
        """
        The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
        """
        return pulumi.get(self, "workspace_url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

