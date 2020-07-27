# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Vault']


class Vault(pulumi.CustomResource):
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[str] = pulumi.output_property("sku")
    """
    Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
    """
    soft_delete_enabled: pulumi.Output[Optional[bool]] = pulumi.output_property("softDeleteEnabled")
    """
    Is soft delete enable for this Vault? Defaults to `true`.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, sku: Optional[pulumi.Input[str]] = None, soft_delete_enabled: Optional[pulumi.Input[bool]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an Recovery Services Vault.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West US")
        vault = azure.recoveryservices.Vault("vault",
            location=rg.location,
            resource_group_name=rg.name,
            sku="Standard",
            soft_delete_enabled=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
        :param pulumi.Input[bool] soft_delete_enabled: Is soft delete enable for this Vault? Defaults to `true`.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['soft_delete_enabled'] = soft_delete_enabled
            __props__['tags'] = tags
        super(Vault, __self__).__init__(
            'azure:recoveryservices/vault:Vault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, sku: Optional[pulumi.Input[str]] = None, soft_delete_enabled: Optional[pulumi.Input[bool]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None) -> 'Vault':
        """
        Get an existing Vault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
        :param pulumi.Input[bool] soft_delete_enabled: Is soft delete enable for this Vault? Defaults to `true`.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["soft_delete_enabled"] = soft_delete_enabled
        __props__["tags"] = tags
        return Vault(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

