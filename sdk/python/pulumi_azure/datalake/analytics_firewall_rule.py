# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AnalyticsFirewallRule']


class AnalyticsFirewallRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 end_ip_address: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 start_ip_address: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Azure Data Lake Analytics Firewall Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_store = azure.datalake.Store("exampleStore",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_analytics_account = azure.datalake.AnalyticsAccount("exampleAnalyticsAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            default_store_account_name=example_store.name)
        example_analytics_firewall_rule = azure.datalake.AnalyticsFirewallRule("exampleAnalyticsFirewallRule",
            account_name=azurerm_data_lake_analytics["example"]["name"],
            resource_group_name=example_resource_group.name,
            start_ip_address="1.2.3.4",
            end_ip_address="2.3.4.5")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
        :param pulumi.Input[str] end_ip_address: The End IP Address for the firewall rule.
        :param pulumi.Input[str] name: Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Lake Analytics.
        :param pulumi.Input[str] start_ip_address: The Start IP address for the firewall rule.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if end_ip_address is None:
                raise TypeError("Missing required property 'end_ip_address'")
            __props__['end_ip_address'] = end_ip_address
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if start_ip_address is None:
                raise TypeError("Missing required property 'start_ip_address'")
            __props__['start_ip_address'] = start_ip_address
        super(AnalyticsFirewallRule, __self__).__init__(
            'azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            end_ip_address: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            start_ip_address: Optional[pulumi.Input[str]] = None) -> 'AnalyticsFirewallRule':
        """
        Get an existing AnalyticsFirewallRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
        :param pulumi.Input[str] end_ip_address: The End IP Address for the firewall rule.
        :param pulumi.Input[str] name: Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Lake Analytics.
        :param pulumi.Input[str] start_ip_address: The Start IP address for the firewall rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["end_ip_address"] = end_ip_address
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["start_ip_address"] = start_ip_address
        return AnalyticsFirewallRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="endIpAddress")
    def end_ip_address(self) -> pulumi.Output[str]:
        """
        The End IP Address for the firewall rule.
        """
        return pulumi.get(self, "end_ip_address")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Lake Analytics.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="startIpAddress")
    def start_ip_address(self) -> pulumi.Output[str]:
        """
        The Start IP address for the firewall rule.
        """
        return pulumi.get(self, "start_ip_address")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

