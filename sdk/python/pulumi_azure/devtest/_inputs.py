# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GlobalVMShutdownScheduleNotificationSettingsArgs',
    'LinuxVirtualMachineGalleryImageReferenceArgs',
    'LinuxVirtualMachineInboundNatRuleArgs',
    'ScheduleDailyRecurrenceArgs',
    'ScheduleHourlyRecurrenceArgs',
    'ScheduleNotificationSettingsArgs',
    'ScheduleWeeklyRecurrenceArgs',
    'VirtualNetworkSubnetArgs',
    'WindowsVirtualMachineGalleryImageReferenceArgs',
    'WindowsVirtualMachineInboundNatRuleArgs',
]

@pulumi.input_type
class GlobalVMShutdownScheduleNotificationSettingsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 time_in_minutes: Optional[pulumi.Input[float]] = None,
                 webhook_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether to enable pre-shutdown notifications. Possible values are `true` and `false`. Defaults to `false`
        :param pulumi.Input[float] time_in_minutes: Time in minutes between 15 and 120 before a shutdown event at which a notification will be sent. Defaults to `30`.
        :param pulumi.Input[str] webhook_url: The webhook URL to which the notification will be sent. Required if `enabled` is `true`. Optional otherwise.
        """
        pulumi.set(__self__, "enabled", enabled)
        if time_in_minutes is not None:
            pulumi.set(__self__, "time_in_minutes", time_in_minutes)
        if webhook_url is not None:
            pulumi.set(__self__, "webhook_url", webhook_url)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Whether to enable pre-shutdown notifications. Possible values are `true` and `false`. Defaults to `false`
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="timeInMinutes")
    def time_in_minutes(self) -> Optional[pulumi.Input[float]]:
        """
        Time in minutes between 15 and 120 before a shutdown event at which a notification will be sent. Defaults to `30`.
        """
        return pulumi.get(self, "time_in_minutes")

    @time_in_minutes.setter
    def time_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "time_in_minutes", value)

    @property
    @pulumi.getter(name="webhookUrl")
    def webhook_url(self) -> Optional[pulumi.Input[str]]:
        """
        The webhook URL to which the notification will be sent. Required if `enabled` is `true`. Optional otherwise.
        """
        return pulumi.get(self, "webhook_url")

    @webhook_url.setter
    def webhook_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webhook_url", value)


@pulumi.input_type
class LinuxVirtualMachineGalleryImageReferenceArgs:
    def __init__(__self__, *,
                 offer: pulumi.Input[str],
                 publisher: pulumi.Input[str],
                 sku: pulumi.Input[str],
                 version: pulumi.Input[str]):
        """
        :param pulumi.Input[str] offer: The Offer of the Gallery Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] publisher: The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Gallery Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] version: The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "offer", offer)
        pulumi.set(__self__, "publisher", publisher)
        pulumi.set(__self__, "sku", sku)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def offer(self) -> pulumi.Input[str]:
        """
        The Offer of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "offer")

    @offer.setter
    def offer(self, value: pulumi.Input[str]):
        pulumi.set(self, "offer", value)

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Input[str]:
        """
        The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: pulumi.Input[str]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input[str]:
        """
        The SKU of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class LinuxVirtualMachineInboundNatRuleArgs:
    def __init__(__self__, *,
                 backend_port: pulumi.Input[float],
                 protocol: pulumi.Input[str],
                 frontend_port: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[float] backend_port: The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protocol: The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        :param pulumi.Input[float] frontend_port: The frontend port associated with this Inbound NAT Rule.
        """
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "protocol", protocol)
        if frontend_port is not None:
            pulumi.set(__self__, "frontend_port", frontend_port)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> pulumi.Input[float]:
        """
        The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_port")

    @backend_port.setter
    def backend_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "backend_port", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> Optional[pulumi.Input[float]]:
        """
        The frontend port associated with this Inbound NAT Rule.
        """
        return pulumi.get(self, "frontend_port")

    @frontend_port.setter
    def frontend_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "frontend_port", value)


@pulumi.input_type
class ScheduleDailyRecurrenceArgs:
    def __init__(__self__, *,
                 time: pulumi.Input[str]):
        """
        :param pulumi.Input[str] time: The time each day when the schedule takes effect.
        """
        pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def time(self) -> pulumi.Input[str]:
        """
        The time each day when the schedule takes effect.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: pulumi.Input[str]):
        pulumi.set(self, "time", value)


@pulumi.input_type
class ScheduleHourlyRecurrenceArgs:
    def __init__(__self__, *,
                 minute: pulumi.Input[float]):
        pulumi.set(__self__, "minute", minute)

    @property
    @pulumi.getter
    def minute(self) -> pulumi.Input[float]:
        return pulumi.get(self, "minute")

    @minute.setter
    def minute(self, value: pulumi.Input[float]):
        pulumi.set(self, "minute", value)


@pulumi.input_type
class ScheduleNotificationSettingsArgs:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None,
                 time_in_minutes: Optional[pulumi.Input[float]] = None,
                 webhook_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] status: The status of the notification. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`
        :param pulumi.Input[float] time_in_minutes: Time in minutes before event at which notification will be sent.
        :param pulumi.Input[str] webhook_url: The webhook URL to which the notification will be sent.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if time_in_minutes is not None:
            pulumi.set(__self__, "time_in_minutes", time_in_minutes)
        if webhook_url is not None:
            pulumi.set(__self__, "webhook_url", webhook_url)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the notification. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="timeInMinutes")
    def time_in_minutes(self) -> Optional[pulumi.Input[float]]:
        """
        Time in minutes before event at which notification will be sent.
        """
        return pulumi.get(self, "time_in_minutes")

    @time_in_minutes.setter
    def time_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "time_in_minutes", value)

    @property
    @pulumi.getter(name="webhookUrl")
    def webhook_url(self) -> Optional[pulumi.Input[str]]:
        """
        The webhook URL to which the notification will be sent.
        """
        return pulumi.get(self, "webhook_url")

    @webhook_url.setter
    def webhook_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webhook_url", value)


@pulumi.input_type
class ScheduleWeeklyRecurrenceArgs:
    def __init__(__self__, *,
                 time: pulumi.Input[str],
                 week_days: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] time: The time when the schedule takes effect.
        :param pulumi.Input[List[pulumi.Input[str]]] week_days: A list of days that this schedule takes effect . Possible values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        pulumi.set(__self__, "time", time)
        if week_days is not None:
            pulumi.set(__self__, "week_days", week_days)

    @property
    @pulumi.getter
    def time(self) -> pulumi.Input[str]:
        """
        The time when the schedule takes effect.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: pulumi.Input[str]):
        pulumi.set(self, "time", value)

    @property
    @pulumi.getter(name="weekDays")
    def week_days(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        A list of days that this schedule takes effect . Possible values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        return pulumi.get(self, "week_days")

    @week_days.setter
    def week_days(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "week_days", value)


@pulumi.input_type
class VirtualNetworkSubnetArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 use_in_virtual_machine_creation: Optional[pulumi.Input[str]] = None,
                 use_public_ip_address: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        :param pulumi.Input[str] use_in_virtual_machine_creation: Can this subnet be used for creating Virtual Machines? Possible values are `Allow`, `Default` and `Deny`.
        :param pulumi.Input[str] use_public_ip_address: Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are `Allow`, `Default` and `Deny`.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if use_in_virtual_machine_creation is not None:
            pulumi.set(__self__, "use_in_virtual_machine_creation", use_in_virtual_machine_creation)
        if use_public_ip_address is not None:
            pulumi.set(__self__, "use_public_ip_address", use_public_ip_address)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="useInVirtualMachineCreation")
    def use_in_virtual_machine_creation(self) -> Optional[pulumi.Input[str]]:
        """
        Can this subnet be used for creating Virtual Machines? Possible values are `Allow`, `Default` and `Deny`.
        """
        return pulumi.get(self, "use_in_virtual_machine_creation")

    @use_in_virtual_machine_creation.setter
    def use_in_virtual_machine_creation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_in_virtual_machine_creation", value)

    @property
    @pulumi.getter(name="usePublicIpAddress")
    def use_public_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are `Allow`, `Default` and `Deny`.
        """
        return pulumi.get(self, "use_public_ip_address")

    @use_public_ip_address.setter
    def use_public_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_public_ip_address", value)


@pulumi.input_type
class WindowsVirtualMachineGalleryImageReferenceArgs:
    def __init__(__self__, *,
                 offer: pulumi.Input[str],
                 publisher: pulumi.Input[str],
                 sku: pulumi.Input[str],
                 version: pulumi.Input[str]):
        """
        :param pulumi.Input[str] offer: The Offer of the Gallery Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] publisher: The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Gallery Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] version: The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "offer", offer)
        pulumi.set(__self__, "publisher", publisher)
        pulumi.set(__self__, "sku", sku)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def offer(self) -> pulumi.Input[str]:
        """
        The Offer of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "offer")

    @offer.setter
    def offer(self, value: pulumi.Input[str]):
        pulumi.set(self, "offer", value)

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Input[str]:
        """
        The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: pulumi.Input[str]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input[str]:
        """
        The SKU of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class WindowsVirtualMachineInboundNatRuleArgs:
    def __init__(__self__, *,
                 backend_port: pulumi.Input[float],
                 protocol: pulumi.Input[str],
                 frontend_port: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[float] backend_port: The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protocol: The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        :param pulumi.Input[float] frontend_port: The frontend port associated with this Inbound NAT Rule.
        """
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "protocol", protocol)
        if frontend_port is not None:
            pulumi.set(__self__, "frontend_port", frontend_port)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> pulumi.Input[float]:
        """
        The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_port")

    @backend_port.setter
    def backend_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "backend_port", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> Optional[pulumi.Input[float]]:
        """
        The frontend port associated with this Inbound NAT Rule.
        """
        return pulumi.get(self, "frontend_port")

    @frontend_port.setter
    def frontend_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "frontend_port", value)


