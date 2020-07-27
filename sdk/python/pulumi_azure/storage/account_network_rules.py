# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['AccountNetworkRules']


class AccountNetworkRules(pulumi.CustomResource):
    bypasses: pulumi.Output[List[str]] = pulumi.output_property("bypasses")
    """
    Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
    """
    default_action: pulumi.Output[str] = pulumi.output_property("defaultAction")
    """
    Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
    """
    ip_rules: pulumi.Output[List[str]] = pulumi.output_property("ipRules")
    """
    List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
    """
    storage_account_name: pulumi.Output[str] = pulumi.output_property("storageAccountName")
    """
    Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
    """
    virtual_network_subnet_ids: pulumi.Output[List[str]] = pulumi.output_property("virtualNetworkSubnetIds")
    """
    A list of virtual network subnet ids to to secure the storage account.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, bypasses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, default_action: Optional[pulumi.Input[str]] = None, ip_rules: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, storage_account_name: Optional[pulumi.Input[str]] = None, virtual_network_subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages network rules inside of a Azure Storage Account.

        > **NOTE:** Network Rules can be defined either directly on the `storage.Account` resource, or using the `storage.AccountNetworkRules` resource - but the two cannot be used together. Spurious changes will occur if both are used against the same Storage Account.

        > **NOTE:** Only one `storage.AccountNetworkRules` can be tied to an `storage.Account`. Spurious changes will occur if more than `storage.AccountNetworkRules` is tied to the same `storage.Account`.

        > **NOTE:** Deleting this resource updates the storage account back to the default values it had when the storage account was created.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.2.0/24",
            service_endpoints=["Microsoft.Storage"])
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS",
            tags={
                "environment": "staging",
            })
        test = azure.storage.AccountNetworkRules("test",
            resource_group_name=azurerm_resource_group["test"]["name"],
            storage_account_name=azurerm_storage_account["test"]["name"],
            default_action="Allow",
            ip_rules=["127.0.0.1"],
            virtual_network_subnet_ids=[azurerm_subnet["test"]["id"]],
            bypasses=["Metrics"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] bypasses: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
        :param pulumi.Input[str] default_action: Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_rules: List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_name: Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
        :param pulumi.Input[List[pulumi.Input[str]]] virtual_network_subnet_ids: A list of virtual network subnet ids to to secure the storage account.
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

            __props__['bypasses'] = bypasses
            if default_action is None:
                raise TypeError("Missing required property 'default_action'")
            __props__['default_action'] = default_action
            __props__['ip_rules'] = ip_rules
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            __props__['virtual_network_subnet_ids'] = virtual_network_subnet_ids
        super(AccountNetworkRules, __self__).__init__(
            'azure:storage/accountNetworkRules:AccountNetworkRules',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, bypasses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, default_action: Optional[pulumi.Input[str]] = None, ip_rules: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, storage_account_name: Optional[pulumi.Input[str]] = None, virtual_network_subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'AccountNetworkRules':
        """
        Get an existing AccountNetworkRules resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] bypasses: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
        :param pulumi.Input[str] default_action: Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_rules: List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_name: Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
        :param pulumi.Input[List[pulumi.Input[str]]] virtual_network_subnet_ids: A list of virtual network subnet ids to to secure the storage account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bypasses"] = bypasses
        __props__["default_action"] = default_action
        __props__["ip_rules"] = ip_rules
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_account_name"] = storage_account_name
        __props__["virtual_network_subnet_ids"] = virtual_network_subnet_ids
        return AccountNetworkRules(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

