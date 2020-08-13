# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Workspace']


class Workspace(pulumi.CustomResource):
    application_insights_id: pulumi.Output[str] = pulumi.property("applicationInsightsId")
    """
    The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
    """

    container_registry_id: pulumi.Output[Optional[str]] = pulumi.property("containerRegistryId")
    """
    The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    The description of this Machine Learning Workspace.
    """

    friendly_name: pulumi.Output[Optional[str]] = pulumi.property("friendlyName")
    """
    Friendly name for this Machine Learning Workspace.
    """

    high_business_impact: pulumi.Output[Optional[bool]] = pulumi.property("highBusinessImpact")
    """
    Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
    """

    identity: pulumi.Output['outputs.WorkspaceIdentity'] = pulumi.property("identity")
    """
    An `identity` block defined below.
    """

    key_vault_id: pulumi.Output[str] = pulumi.property("keyVaultId")
    """
    The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
    """

    sku_name: pulumi.Output[Optional[str]] = pulumi.property("skuName")
    """
    SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
    """

    storage_account_id: pulumi.Output[str] = pulumi.property("storageAccountId")
    """
    The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_insights_id: Optional[pulumi.Input[str]] = None,
                 container_registry_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 high_business_impact: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['WorkspaceIdentityArgs']]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Azure Machine Learning Workspace

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_insights = azure.appinsights.Insights("exampleInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_type="web")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium")
        example_account = azure.storage.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_tier="Standard",
            account_replication_type="GRS")
        example_workspace = azure.machinelearning.Workspace("exampleWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_insights_id=example_insights.id,
            key_vault_id=example_key_vault.id,
            storage_account_id=example_account.id,
            identity={
                "type": "SystemAssigned",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_insights_id: The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] container_registry_id: The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of this Machine Learning Workspace.
        :param pulumi.Input[str] friendly_name: Friendly name for this Machine Learning Workspace.
        :param pulumi.Input[bool] high_business_impact: Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
        :param pulumi.Input[pulumi.InputType['WorkspaceIdentityArgs']] identity: An `identity` block defined below.
        :param pulumi.Input[str] key_vault_id: The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
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

            if application_insights_id is None:
                raise TypeError("Missing required property 'application_insights_id'")
            __props__['application_insights_id'] = application_insights_id
            __props__['container_registry_id'] = container_registry_id
            __props__['description'] = description
            __props__['friendly_name'] = friendly_name
            __props__['high_business_impact'] = high_business_impact
            if identity is None:
                raise TypeError("Missing required property 'identity'")
            __props__['identity'] = identity
            if key_vault_id is None:
                raise TypeError("Missing required property 'key_vault_id'")
            __props__['key_vault_id'] = key_vault_id
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku_name'] = sku_name
            if storage_account_id is None:
                raise TypeError("Missing required property 'storage_account_id'")
            __props__['storage_account_id'] = storage_account_id
            __props__['tags'] = tags
        super(Workspace, __self__).__init__(
            'azure:machinelearning/workspace:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            application_insights_id: Optional[pulumi.Input[str]] = None,
            container_registry_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            high_business_impact: Optional[pulumi.Input[bool]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['WorkspaceIdentityArgs']]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_insights_id: The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] container_registry_id: The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of this Machine Learning Workspace.
        :param pulumi.Input[str] friendly_name: Friendly name for this Machine Learning Workspace.
        :param pulumi.Input[bool] high_business_impact: Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
        :param pulumi.Input[pulumi.InputType['WorkspaceIdentityArgs']] identity: An `identity` block defined below.
        :param pulumi.Input[str] key_vault_id: The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_insights_id"] = application_insights_id
        __props__["container_registry_id"] = container_registry_id
        __props__["description"] = description
        __props__["friendly_name"] = friendly_name
        __props__["high_business_impact"] = high_business_impact
        __props__["identity"] = identity
        __props__["key_vault_id"] = key_vault_id
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["storage_account_id"] = storage_account_id
        __props__["tags"] = tags
        return Workspace(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

