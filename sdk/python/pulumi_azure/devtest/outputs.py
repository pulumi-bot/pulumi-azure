# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.output_type
class GlobalVMShutdownScheduleNotificationSettings(dict):
    enabled: bool = pulumi.output_property("enabled")
    """
    Whether to enable pre-shutdown notifications. Possible values are `true` and `false`. Defaults to `false`
    """
    time_in_minutes: Optional[float] = pulumi.output_property("timeInMinutes")
    """
    Time in minutes between 15 and 120 before a shutdown event at which a notification will be sent. Defaults to `30`.
    """
    webhook_url: Optional[str] = pulumi.output_property("webhookUrl")
    """
    The webhook URL to which the notification will be sent. Required if `enabled` is `true`. Optional otherwise.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LinuxVirtualMachineGalleryImageReference(dict):
    offer: str = pulumi.output_property("offer")
    """
    The Offer of the Gallery Image. Changing this forces a new resource to be created.
    """
    publisher: str = pulumi.output_property("publisher")
    """
    The Publisher of the Gallery Image. Changing this forces a new resource to be created.
    """
    sku: str = pulumi.output_property("sku")
    """
    The SKU of the Gallery Image. Changing this forces a new resource to be created.
    """
    version: str = pulumi.output_property("version")
    """
    The Version of the Gallery Image. Changing this forces a new resource to be created.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LinuxVirtualMachineInboundNatRule(dict):
    backend_port: float = pulumi.output_property("backendPort")
    """
    The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
    """
    frontend_port: Optional[float] = pulumi.output_property("frontendPort")
    """
    The frontend port associated with this Inbound NAT Rule.
    """
    protocol: str = pulumi.output_property("protocol")
    """
    The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleDailyRecurrence(dict):
    time: str = pulumi.output_property("time")
    """
    The time each day when the schedule takes effect.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleHourlyRecurrence(dict):
    minute: float = pulumi.output_property("minute")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleNotificationSettings(dict):
    status: Optional[str] = pulumi.output_property("status")
    """
    The status of the notification. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`
    """
    time_in_minutes: Optional[float] = pulumi.output_property("timeInMinutes")
    """
    Time in minutes before event at which notification will be sent.
    """
    webhook_url: Optional[str] = pulumi.output_property("webhookUrl")
    """
    The webhook URL to which the notification will be sent.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleWeeklyRecurrence(dict):
    time: str = pulumi.output_property("time")
    """
    The time when the schedule takes effect.
    """
    week_days: Optional[List[str]] = pulumi.output_property("weekDays")
    """
    A list of days that this schedule takes effect . Possible values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VirtualNetworkSubnet(dict):
    name: Optional[str] = pulumi.output_property("name")
    """
    Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
    """
    use_in_virtual_machine_creation: Optional[str] = pulumi.output_property("useInVirtualMachineCreation")
    """
    Can this subnet be used for creating Virtual Machines? Possible values are `Allow`, `Default` and `Deny`.
    """
    use_public_ip_address: Optional[str] = pulumi.output_property("usePublicIpAddress")
    """
    Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are `Allow`, `Default` and `Deny`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WindowsVirtualMachineGalleryImageReference(dict):
    offer: str = pulumi.output_property("offer")
    """
    The Offer of the Gallery Image. Changing this forces a new resource to be created.
    """
    publisher: str = pulumi.output_property("publisher")
    """
    The Publisher of the Gallery Image. Changing this forces a new resource to be created.
    """
    sku: str = pulumi.output_property("sku")
    """
    The SKU of the Gallery Image. Changing this forces a new resource to be created.
    """
    version: str = pulumi.output_property("version")
    """
    The Version of the Gallery Image. Changing this forces a new resource to be created.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WindowsVirtualMachineInboundNatRule(dict):
    backend_port: float = pulumi.output_property("backendPort")
    """
    The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
    """
    frontend_port: Optional[float] = pulumi.output_property("frontendPort")
    """
    The frontend port associated with this Inbound NAT Rule.
    """
    protocol: str = pulumi.output_property("protocol")
    """
    The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetVirtualNetworkAllowedSubnet(dict):
    allow_public_ip: str = pulumi.output_property("allowPublicIp")
    """
    Indicates if this subnet allows public IP addresses. Possible values are `Allow`, `Default` and `Deny`.
    """
    lab_subnet_name: str = pulumi.output_property("labSubnetName")
    """
    The name of the subnet.
    """
    resource_id: str = pulumi.output_property("resourceId")
    """
    The resource identifier for the subnet.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetVirtualNetworkSubnetOverride(dict):
    lab_subnet_name: str = pulumi.output_property("labSubnetName")
    """
    The name of the subnet.
    """
    resource_id: str = pulumi.output_property("resourceId")
    """
    The resource identifier for the subnet.
    """
    use_in_vm_creation_permission: str = pulumi.output_property("useInVmCreationPermission")
    """
    Indicates if the subnet can be used for VM creation.  Possible values are `Allow`, `Default` and `Deny`.
    """
    use_public_ip_address_permission: str = pulumi.output_property("usePublicIpAddressPermission")
    virtual_network_pool_name: str = pulumi.output_property("virtualNetworkPoolName")
    """
    The virtual network pool associated with this subnet.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

