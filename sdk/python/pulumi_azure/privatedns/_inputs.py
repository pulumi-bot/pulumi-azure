# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'LinkServiceNatIpConfigurationArgs',
    'MxRecordRecordArgs',
    'SRVRecordRecordArgs',
    'TxtRecordRecordArgs',
]

@pulumi.input_type
class LinkServiceNatIpConfigurationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 primary: pulumi.Input[bool],
                 subnet_id: pulumi.Input[str],
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_address_version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] primary: Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: Specifies the ID of the Subnet which should be used for the Private Link Service.
        :param pulumi.Input[str] private_ip_address: Specifies a Private Static IP Address for this IP Configuration.
        :param pulumi.Input[str] private_ip_address_version: The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "primary", primary)
        pulumi.set(__self__, "subnetId", subnet_id)
        pulumi.set(__self__, "privateIpAddress", private_ip_address)
        pulumi.set(__self__, "privateIpAddressVersion", private_ip_address_version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def primary(self) -> pulumi.Input[bool]:
        """
        Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
        """
        ...

    @primary.setter
    def primary(self, value: pulumi.Input[bool]):
        ...

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the Subnet which should be used for the Private Link Service.
        """
        ...

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a Private Static IP Address for this IP Configuration.
        """
        ...

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="privateIpAddressVersion")
    def private_ip_address_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
        """
        ...

    @private_ip_address_version.setter
    def private_ip_address_version(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class MxRecordRecordArgs:
    def __init__(__self__, *,
                 exchange: pulumi.Input[str],
                 preference: pulumi.Input[float]):
        """
        :param pulumi.Input[str] exchange: The FQDN of the exchange to MX record points to.
        :param pulumi.Input[float] preference: The preference of the MX record.
        """
        pulumi.set(__self__, "exchange", exchange)
        pulumi.set(__self__, "preference", preference)

    @property
    @pulumi.getter
    def exchange(self) -> pulumi.Input[str]:
        """
        The FQDN of the exchange to MX record points to.
        """
        ...

    @exchange.setter
    def exchange(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def preference(self) -> pulumi.Input[float]:
        """
        The preference of the MX record.
        """
        ...

    @preference.setter
    def preference(self, value: pulumi.Input[float]):
        ...


@pulumi.input_type
class SRVRecordRecordArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[float],
                 priority: pulumi.Input[float],
                 target: pulumi.Input[str],
                 weight: pulumi.Input[float]):
        """
        :param pulumi.Input[float] port: The Port the service is listening on.
        :param pulumi.Input[float] priority: The priority of the SRV record.
        :param pulumi.Input[str] target: The FQDN of the service.
        :param pulumi.Input[float] weight: The Weight of the SRV record.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[float]:
        """
        The Port the service is listening on.
        """
        ...

    @port.setter
    def port(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[float]:
        """
        The priority of the SRV record.
        """
        ...

    @priority.setter
    def priority(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        """
        The FQDN of the service.
        """
        ...

    @target.setter
    def target(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[float]:
        """
        The Weight of the SRV record.
        """
        ...

    @weight.setter
    def weight(self, value: pulumi.Input[float]):
        ...


@pulumi.input_type
class TxtRecordRecordArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] value: The value of the TXT record. Max length: 1024 characters
        """
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the TXT record. Max length: 1024 characters
        """
        ...

    @value.setter
    def value(self, value: pulumi.Input[str]):
        ...


