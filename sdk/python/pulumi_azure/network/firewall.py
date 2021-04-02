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

__all__ = ['Firewall']


class Firewall(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 firewall_policy_id: Optional[pulumi.Input[str]] = None,
                 ip_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallIpConfigurationArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_ip_configuration: Optional[pulumi.Input[pulumi.InputType['FirewallManagementIpConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 sku_tier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 threat_intel_mode: Optional[pulumi.Input[str]] = None,
                 virtual_hub: Optional[pulumi.Input[pulumi.InputType['FirewallVirtualHubArgs']]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Firewall.

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
            address_prefixes=["10.0.1.0/24"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_firewall = azure.network.Firewall("exampleFirewall",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.FirewallIpConfigurationArgs(
                name="configuration",
                subnet_id=example_subnet.id,
                public_ip_address_id=example_public_ip.id,
            )])
        ```

        ## Import

        Azure Firewalls can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/firewall:Firewall example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/azureFirewalls/testfirewall
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
        :param pulumi.Input[str] firewall_policy_id: The ID of the Firewall Policy applied to this Firewall.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallIpConfigurationArgs']]]] ip_configurations: An `ip_configuration` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['FirewallManagementIpConfigurationArgs']] management_ip_configuration: A `management_ip_configuration` block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the `subnet_id` in an existing block forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_ranges: A list of SNAT private CIDR IP ranges, or the special string `IANAPrivateRanges`, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Sku name of the Firewall. Possible values are `AZFW_Hub` and `AZFW_VNet`.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_tier: Sku tier of the Firewall. Possible values are `Premium` and `Standard`.  Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] threat_intel_mode: The operation mode for threat intelligence-based filtering. Possible values are: `Off`, `Alert`,`Deny` and `""`(empty string). Defaults to `Alert`.
        :param pulumi.Input[pulumi.InputType['FirewallVirtualHubArgs']] virtual_hub: A `virtual_hub` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: Specifies the availability zones in which the Azure Firewall should be created. Changing this forces a new resource to be created.
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

            __props__['dns_servers'] = dns_servers
            __props__['firewall_policy_id'] = firewall_policy_id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            __props__['management_ip_configuration'] = management_ip_configuration
            __props__['name'] = name
            __props__['private_ip_ranges'] = private_ip_ranges
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku_name'] = sku_name
            __props__['sku_tier'] = sku_tier
            __props__['tags'] = tags
            __props__['threat_intel_mode'] = threat_intel_mode
            __props__['virtual_hub'] = virtual_hub
            __props__['zones'] = zones
        super(Firewall, __self__).__init__(
            'azure:network/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            firewall_policy_id: Optional[pulumi.Input[str]] = None,
            ip_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallIpConfigurationArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            management_ip_configuration: Optional[pulumi.Input[pulumi.InputType['FirewallManagementIpConfigurationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            sku_tier: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            threat_intel_mode: Optional[pulumi.Input[str]] = None,
            virtual_hub: Optional[pulumi.Input[pulumi.InputType['FirewallVirtualHubArgs']]] = None,
            zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Firewall':
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
        :param pulumi.Input[str] firewall_policy_id: The ID of the Firewall Policy applied to this Firewall.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallIpConfigurationArgs']]]] ip_configurations: An `ip_configuration` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['FirewallManagementIpConfigurationArgs']] management_ip_configuration: A `management_ip_configuration` block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the `subnet_id` in an existing block forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_ranges: A list of SNAT private CIDR IP ranges, or the special string `IANAPrivateRanges`, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Sku name of the Firewall. Possible values are `AZFW_Hub` and `AZFW_VNet`.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_tier: Sku tier of the Firewall. Possible values are `Premium` and `Standard`.  Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] threat_intel_mode: The operation mode for threat intelligence-based filtering. Possible values are: `Off`, `Alert`,`Deny` and `""`(empty string). Defaults to `Alert`.
        :param pulumi.Input[pulumi.InputType['FirewallVirtualHubArgs']] virtual_hub: A `virtual_hub` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: Specifies the availability zones in which the Azure Firewall should be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dns_servers"] = dns_servers
        __props__["firewall_policy_id"] = firewall_policy_id
        __props__["ip_configurations"] = ip_configurations
        __props__["location"] = location
        __props__["management_ip_configuration"] = management_ip_configuration
        __props__["name"] = name
        __props__["private_ip_ranges"] = private_ip_ranges
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["sku_tier"] = sku_tier
        __props__["tags"] = tags
        __props__["threat_intel_mode"] = threat_intel_mode
        __props__["virtual_hub"] = virtual_hub
        __props__["zones"] = zones
        return Firewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
        """
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter(name="firewallPolicyId")
    def firewall_policy_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Firewall Policy applied to this Firewall.
        """
        return pulumi.get(self, "firewall_policy_id")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallIpConfiguration']]]:
        """
        An `ip_configuration` block as documented below.
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementIpConfiguration")
    def management_ip_configuration(self) -> pulumi.Output[Optional['outputs.FirewallManagementIpConfiguration']]:
        """
        A `management_ip_configuration` block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the `subnet_id` in an existing block forces a new resource to be created.
        """
        return pulumi.get(self, "management_ip_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Firewall. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateIpRanges")
    def private_ip_ranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of SNAT private CIDR IP ranges, or the special string `IANAPrivateRanges`, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
        """
        return pulumi.get(self, "private_ip_ranges")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        Sku name of the Firewall. Possible values are `AZFW_Hub` and `AZFW_VNet`.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="skuTier")
    def sku_tier(self) -> pulumi.Output[str]:
        """
        Sku tier of the Firewall. Possible values are `Premium` and `Standard`.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_tier")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="threatIntelMode")
    def threat_intel_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The operation mode for threat intelligence-based filtering. Possible values are: `Off`, `Alert`,`Deny` and `""`(empty string). Defaults to `Alert`.
        """
        return pulumi.get(self, "threat_intel_mode")

    @property
    @pulumi.getter(name="virtualHub")
    def virtual_hub(self) -> pulumi.Output[Optional['outputs.FirewallVirtualHub']]:
        """
        A `virtual_hub` block as documented below.
        """
        return pulumi.get(self, "virtual_hub")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the availability zones in which the Azure Firewall should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

