# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetFirewallResult',
    'AwaitableGetFirewallResult',
    'get_firewall',
]

@pulumi.output_type
class GetFirewallResult:
    """
    A collection of values returned by getFirewall.
    """
    def __init__(__self__, dns_servers=None, firewall_policy_id=None, id=None, ip_configurations=None, location=None, management_ip_configurations=None, name=None, resource_group_name=None, sku_name=None, sku_tier=None, tags=None, threat_intel_mode=None, virtual_hubs=None, zones=None):
        if dns_servers and not isinstance(dns_servers, list):
            raise TypeError("Expected argument 'dns_servers' to be a list")
        pulumi.set(__self__, "dns_servers", dns_servers)
        if firewall_policy_id and not isinstance(firewall_policy_id, str):
            raise TypeError("Expected argument 'firewall_policy_id' to be a str")
        pulumi.set(__self__, "firewall_policy_id", firewall_policy_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_configurations and not isinstance(ip_configurations, list):
            raise TypeError("Expected argument 'ip_configurations' to be a list")
        pulumi.set(__self__, "ip_configurations", ip_configurations)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if management_ip_configurations and not isinstance(management_ip_configurations, list):
            raise TypeError("Expected argument 'management_ip_configurations' to be a list")
        pulumi.set(__self__, "management_ip_configurations", management_ip_configurations)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        pulumi.set(__self__, "sku_name", sku_name)
        if sku_tier and not isinstance(sku_tier, str):
            raise TypeError("Expected argument 'sku_tier' to be a str")
        pulumi.set(__self__, "sku_tier", sku_tier)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if threat_intel_mode and not isinstance(threat_intel_mode, str):
            raise TypeError("Expected argument 'threat_intel_mode' to be a str")
        pulumi.set(__self__, "threat_intel_mode", threat_intel_mode)
        if virtual_hubs and not isinstance(virtual_hubs, list):
            raise TypeError("Expected argument 'virtual_hubs' to be a list")
        pulumi.set(__self__, "virtual_hubs", virtual_hubs)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> Sequence[str]:
        """
        The list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
        """
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter(name="firewallPolicyId")
    def firewall_policy_id(self) -> str:
        """
        The ID of the Firewall Policy applied to the Azure Firewall.
        """
        return pulumi.get(self, "firewall_policy_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> Sequence['outputs.GetFirewallIpConfigurationResult']:
        """
        A `ip_configuration` block as defined below.
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure location where the Azure Firewall exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementIpConfigurations")
    def management_ip_configurations(self) -> Sequence['outputs.GetFirewallManagementIpConfigurationResult']:
        """
        A `management_ip_configuration` block as defined below, which allows force-tunnelling of traffic to be performed by the firewall.
        """
        return pulumi.get(self, "management_ip_configurations")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        """
        The sku name of the Azure Firewall.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="skuTier")
    def sku_tier(self) -> str:
        """
        The sku tier of the Azure Firewall.
        """
        return pulumi.get(self, "sku_tier")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the Azure Firewall.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="threatIntelMode")
    def threat_intel_mode(self) -> str:
        """
        The operation mode for threat intelligence-based filtering.
        """
        return pulumi.get(self, "threat_intel_mode")

    @property
    @pulumi.getter(name="virtualHubs")
    def virtual_hubs(self) -> Sequence['outputs.GetFirewallVirtualHubResult']:
        """
        A `virtual_hub` block as defined below.
        """
        return pulumi.get(self, "virtual_hubs")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[str]:
        """
        The availability zones in which the Azure Firewall is created.
        """
        return pulumi.get(self, "zones")


class AwaitableGetFirewallResult(GetFirewallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallResult(
            dns_servers=self.dns_servers,
            firewall_policy_id=self.firewall_policy_id,
            id=self.id,
            ip_configurations=self.ip_configurations,
            location=self.location,
            management_ip_configurations=self.management_ip_configurations,
            name=self.name,
            resource_group_name=self.resource_group_name,
            sku_name=self.sku_name,
            sku_tier=self.sku_tier,
            tags=self.tags,
            threat_intel_mode=self.threat_intel_mode,
            virtual_hubs=self.virtual_hubs,
            zones=self.zones)


def get_firewall(name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 zones: Optional[Sequence[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallResult:
    """
    Use this data source to access information about an existing Azure Firewall.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_firewall(name="firewall1",
        resource_group_name="firewall-RG")
    pulumi.export("firewallPrivateIp", example.ip_configurations[0].private_ip_address)
    ```


    :param str name: The name of the Azure Firewall.
    :param str resource_group_name: The name of the Resource Group in which the Azure Firewall exists.
    :param Sequence[str] zones: The availability zones in which the Azure Firewall is created.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['zones'] = zones
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getFirewall:getFirewall', __args__, opts=opts, typ=GetFirewallResult).value

    return AwaitableGetFirewallResult(
        dns_servers=__ret__.dns_servers,
        firewall_policy_id=__ret__.firewall_policy_id,
        id=__ret__.id,
        ip_configurations=__ret__.ip_configurations,
        location=__ret__.location,
        management_ip_configurations=__ret__.management_ip_configurations,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        sku_name=__ret__.sku_name,
        sku_tier=__ret__.sku_tier,
        tags=__ret__.tags,
        threat_intel_mode=__ret__.threat_intel_mode,
        virtual_hubs=__ret__.virtual_hubs,
        zones=__ret__.zones)


@_utilities.lift_output_func(get_firewall)
def get_firewall_apply(name: Optional[pulumi.Input[str]] = None,
                       resource_group_name: Optional[pulumi.Input[str]] = None,
                       zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallResult]:
    ...
