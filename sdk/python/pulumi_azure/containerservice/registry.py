# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Registry(pulumi.CustomResource):
    admin_enabled: pulumi.Output[bool]
    """
    Specifies whether the admin user is enabled. Defaults to `false`.
    """
    admin_password: pulumi.Output[str]
    """
    The Password associated with the Container Registry Admin account - if the admin account is enabled.
    """
    admin_username: pulumi.Output[str]
    """
    The Username associated with the Container Registry Admin account - if the admin account is enabled.
    """
    georeplication_locations: pulumi.Output[list]
    """
    A list of Azure locations where the container registry should be geo-replicated.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    login_server: pulumi.Output[str]
    """
    The URL that can be used to log into the container registry.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Container Registry. Changing this forces a new resource to be created.
    """
    network_rule_set: pulumi.Output[dict]
    """
    A `network_rule_set` block as documented below.

      * `default_action` (`str`) - The behaviour for requests matching no rules. Either `Allow` or `Deny`. Defaults to `Allow`
      * `ip_rules` (`list`) - One or more `ip_rule` blocks as defined below.
        * `action` (`str`) - The behaviour for requests matching this rule. At this time the only supported value is `Allow`
        * `ipRange` (`str`) - The CIDR block from which requests will match the rule.

      * `virtualNetworks` (`list`) - One or more `virtual_network` blocks as defined below.
        * `action` (`str`) - The behaviour for requests matching this rule. At this time the only supported value is `Allow`
        * `subnet_id` (`str`) - The subnet id from which requests will match the rule.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[str]
    """
    The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
    """
    storage_account_id: pulumi.Output[str]
    """
    The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, admin_enabled=None, georeplication_locations=None, location=None, name=None, network_rule_set=None, resource_group_name=None, sku=None, storage_account_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Container Registry.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West US")
        acr = azure.containerservice.Registry("acr",
            resource_group_name=rg.name,
            location=rg.location,
            sku="Premium",
            admin_enabled=False,
            georeplication_locations=[
                "East US",
                "West Europe",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_enabled: Specifies whether the admin user is enabled. Defaults to `false`.
        :param pulumi.Input[list] georeplication_locations: A list of Azure locations where the container registry should be geo-replicated.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Container Registry. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] network_rule_set: A `network_rule_set` block as documented below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
        :param pulumi.Input[str] storage_account_id: The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **network_rule_set** object supports the following:

          * `default_action` (`pulumi.Input[str]`) - The behaviour for requests matching no rules. Either `Allow` or `Deny`. Defaults to `Allow`
          * `ip_rules` (`pulumi.Input[list]`) - One or more `ip_rule` blocks as defined below.
            * `action` (`pulumi.Input[str]`) - The behaviour for requests matching this rule. At this time the only supported value is `Allow`
            * `ipRange` (`pulumi.Input[str]`) - The CIDR block from which requests will match the rule.

          * `virtualNetworks` (`pulumi.Input[list]`) - One or more `virtual_network` blocks as defined below.
            * `action` (`pulumi.Input[str]`) - The behaviour for requests matching this rule. At this time the only supported value is `Allow`
            * `subnet_id` (`pulumi.Input[str]`) - The subnet id from which requests will match the rule.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_enabled'] = admin_enabled
            __props__['georeplication_locations'] = georeplication_locations
            __props__['location'] = location
            __props__['name'] = name
            __props__['network_rule_set'] = network_rule_set
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['storage_account_id'] = storage_account_id
            __props__['tags'] = tags
            __props__['admin_password'] = None
            __props__['admin_username'] = None
            __props__['login_server'] = None
        super(Registry, __self__).__init__(
            'azure:containerservice/registry:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admin_enabled=None, admin_password=None, admin_username=None, georeplication_locations=None, location=None, login_server=None, name=None, network_rule_set=None, resource_group_name=None, sku=None, storage_account_id=None, tags=None):
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_enabled: Specifies whether the admin user is enabled. Defaults to `false`.
        :param pulumi.Input[str] admin_password: The Password associated with the Container Registry Admin account - if the admin account is enabled.
        :param pulumi.Input[str] admin_username: The Username associated with the Container Registry Admin account - if the admin account is enabled.
        :param pulumi.Input[list] georeplication_locations: A list of Azure locations where the container registry should be geo-replicated.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] login_server: The URL that can be used to log into the container registry.
        :param pulumi.Input[str] name: Specifies the name of the Container Registry. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] network_rule_set: A `network_rule_set` block as documented below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
        :param pulumi.Input[str] storage_account_id: The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **network_rule_set** object supports the following:

          * `default_action` (`pulumi.Input[str]`) - The behaviour for requests matching no rules. Either `Allow` or `Deny`. Defaults to `Allow`
          * `ip_rules` (`pulumi.Input[list]`) - One or more `ip_rule` blocks as defined below.
            * `action` (`pulumi.Input[str]`) - The behaviour for requests matching this rule. At this time the only supported value is `Allow`
            * `ipRange` (`pulumi.Input[str]`) - The CIDR block from which requests will match the rule.

          * `virtualNetworks` (`pulumi.Input[list]`) - One or more `virtual_network` blocks as defined below.
            * `action` (`pulumi.Input[str]`) - The behaviour for requests matching this rule. At this time the only supported value is `Allow`
            * `subnet_id` (`pulumi.Input[str]`) - The subnet id from which requests will match the rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_enabled"] = admin_enabled
        __props__["admin_password"] = admin_password
        __props__["admin_username"] = admin_username
        __props__["georeplication_locations"] = georeplication_locations
        __props__["location"] = location
        __props__["login_server"] = login_server
        __props__["name"] = name
        __props__["network_rule_set"] = network_rule_set
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["storage_account_id"] = storage_account_id
        __props__["tags"] = tags
        return Registry(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
