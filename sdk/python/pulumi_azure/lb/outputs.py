# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'LoadBalancerFrontendIpConfiguration',
    'OutboundRuleFrontendIpConfiguration',
    'GetBackendAddressPoolBackendIpConfiguration',
    'GetLBFrontendIpConfiguration',
]

@pulumi.output_type
class LoadBalancerFrontendIpConfiguration(dict):
    id: Optional[str] = pulumi.output_property("id")
    """
    The id of the Frontend IP Configuration.
    """
    inbound_nat_rules: Optional[List[str]] = pulumi.output_property("inboundNatRules")
    """
    The list of IDs of inbound rules that use this frontend IP.
    """
    load_balancer_rules: Optional[List[str]] = pulumi.output_property("loadBalancerRules")
    """
    The list of IDs of load balancing rules that use this frontend IP.
    """
    name: str = pulumi.output_property("name")
    """
    Specifies the name of the frontend ip configuration.
    """
    outbound_rules: Optional[List[str]] = pulumi.output_property("outboundRules")
    """
    The list of IDs outbound rules that use this frontend IP.
    """
    private_ip_address: Optional[str] = pulumi.output_property("privateIpAddress")
    """
    Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
    """
    private_ip_address_allocation: Optional[str] = pulumi.output_property("privateIpAddressAllocation")
    """
    The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
    """
    private_ip_address_version: Optional[str] = pulumi.output_property("privateIpAddressVersion")
    """
    The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
    """
    public_ip_address_id: Optional[str] = pulumi.output_property("publicIpAddressId")
    """
    The ID of a Public IP Address which should be associated with the Load Balancer.
    """
    public_ip_prefix_id: Optional[str] = pulumi.output_property("publicIpPrefixId")
    """
    The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
    """
    subnet_id: Optional[str] = pulumi.output_property("subnetId")
    """
    The ID of the Subnet which should be associated with the IP Configuration.
    """
    zones: Optional[str] = pulumi.output_property("zones")
    """
    A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OutboundRuleFrontendIpConfiguration(dict):
    id: Optional[str] = pulumi.output_property("id")
    """
    The ID of the Load Balancer Outbound Rule.
    """
    name: str = pulumi.output_property("name")
    """
    The name of the Frontend IP Configuration.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBackendAddressPoolBackendIpConfiguration(dict):
    id: str = pulumi.output_property("id")
    """
    The ID of the Backend Address Pool.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetLBFrontendIpConfiguration(dict):
    id: str = pulumi.output_property("id")
    """
    The id of the Frontend IP Configuration.
    """
    name: str = pulumi.output_property("name")
    """
    Specifies the name of the Load Balancer.
    """
    private_ip_address: str = pulumi.output_property("privateIpAddress")
    """
    Private IP Address to assign to the Load Balancer.
    """
    private_ip_address_allocation: str = pulumi.output_property("privateIpAddressAllocation")
    """
    The allocation method for the Private IP Address used by this Load Balancer.
    """
    private_ip_address_version: str = pulumi.output_property("privateIpAddressVersion")
    """
    The Private IP Address Version, either `IPv4` or `IPv6`.
    """
    public_ip_address_id: str = pulumi.output_property("publicIpAddressId")
    """
    The ID of a  Public IP Address which is associated with this Load Balancer.
    """
    subnet_id: str = pulumi.output_property("subnetId")
    """
    The ID of the Subnet which is associated with the IP Configuration.
    """
    zones: List[str] = pulumi.output_property("zones")
    """
    A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


