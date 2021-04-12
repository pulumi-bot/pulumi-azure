# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'LinkServiceNatIpConfigurationArgs',
    'MxRecordRecordArgs',
    'SRVRecordRecordArgs',
    'TxtRecordRecordArgs',
    'ZoneSoaRecordArgs',
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
        pulumi.set(__self__, "subnet_id", subnet_id)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if private_ip_address_version is not None:
            pulumi.set(__self__, "private_ip_address_version", private_ip_address_version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def primary(self) -> pulumi.Input[bool]:
        """
        Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: pulumi.Input[bool]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the Subnet which should be used for the Private Link Service.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a Private Static IP Address for this IP Configuration.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="privateIpAddressVersion")
    def private_ip_address_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
        """
        return pulumi.get(self, "private_ip_address_version")

    @private_ip_address_version.setter
    def private_ip_address_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address_version", value)


@pulumi.input_type
class MxRecordRecordArgs:
    def __init__(__self__, *,
                 exchange: pulumi.Input[str],
                 preference: pulumi.Input[int]):
        """
        :param pulumi.Input[str] exchange: The FQDN of the exchange to MX record points to.
        :param pulumi.Input[int] preference: The preference of the MX record.
        """
        pulumi.set(__self__, "exchange", exchange)
        pulumi.set(__self__, "preference", preference)

    @property
    @pulumi.getter
    def exchange(self) -> pulumi.Input[str]:
        """
        The FQDN of the exchange to MX record points to.
        """
        return pulumi.get(self, "exchange")

    @exchange.setter
    def exchange(self, value: pulumi.Input[str]):
        pulumi.set(self, "exchange", value)

    @property
    @pulumi.getter
    def preference(self) -> pulumi.Input[int]:
        """
        The preference of the MX record.
        """
        return pulumi.get(self, "preference")

    @preference.setter
    def preference(self, value: pulumi.Input[int]):
        pulumi.set(self, "preference", value)


@pulumi.input_type
class SRVRecordRecordArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[int],
                 priority: pulumi.Input[int],
                 target: pulumi.Input[str],
                 weight: pulumi.Input[int]):
        """
        :param pulumi.Input[int] port: The Port the service is listening on.
        :param pulumi.Input[int] priority: The priority of the SRV record.
        :param pulumi.Input[str] target: The FQDN of the service.
        :param pulumi.Input[int] weight: The Weight of the SRV record.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The Port the service is listening on.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        The priority of the SRV record.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        """
        The FQDN of the service.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[int]:
        """
        The Weight of the SRV record.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[int]):
        pulumi.set(self, "weight", value)


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
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ZoneSoaRecordArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 expire_time: Optional[pulumi.Input[int]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 minimum_ttl: Optional[pulumi.Input[int]] = None,
                 refresh_time: Optional[pulumi.Input[int]] = None,
                 retry_time: Optional[pulumi.Input[int]] = None,
                 serial_number: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] email: The email contact for the SOA record.
        :param pulumi.Input[int] expire_time: The expire time for the SOA record. Defaults to `2419200`.
        :param pulumi.Input[str] fqdn: The fully qualified domain name of the Record Set.
        :param pulumi.Input[str] host_name: The domain name of the authoritative name server for the SOA record.
        :param pulumi.Input[int] minimum_ttl: The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `10`.
        :param pulumi.Input[int] refresh_time: The refresh time for the SOA record. Defaults to `3600`.
        :param pulumi.Input[int] retry_time: The retry time for the SOA record. Defaults to `300`.
        :param pulumi.Input[int] serial_number: The serial number for the SOA record.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Record Set.
        :param pulumi.Input[int] ttl: The Time To Live of the SOA Record in seconds. Defaults to `3600`.
        """
        pulumi.set(__self__, "email", email)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if minimum_ttl is not None:
            pulumi.set(__self__, "minimum_ttl", minimum_ttl)
        if refresh_time is not None:
            pulumi.set(__self__, "refresh_time", refresh_time)
        if retry_time is not None:
            pulumi.set(__self__, "retry_time", retry_time)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The email contact for the SOA record.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[int]]:
        """
        The expire time for the SOA record. Defaults to `2419200`.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified domain name of the Record Set.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name of the authoritative name server for the SOA record.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="minimumTtl")
    def minimum_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `10`.
        """
        return pulumi.get(self, "minimum_ttl")

    @minimum_ttl.setter
    def minimum_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_ttl", value)

    @property
    @pulumi.getter(name="refreshTime")
    def refresh_time(self) -> Optional[pulumi.Input[int]]:
        """
        The refresh time for the SOA record. Defaults to `3600`.
        """
        return pulumi.get(self, "refresh_time")

    @refresh_time.setter
    def refresh_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "refresh_time", value)

    @property
    @pulumi.getter(name="retryTime")
    def retry_time(self) -> Optional[pulumi.Input[int]]:
        """
        The retry time for the SOA record. Defaults to `300`.
        """
        return pulumi.get(self, "retry_time")

    @retry_time.setter
    def retry_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_time", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input[int]]:
        """
        The serial number for the SOA record.
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the Record Set.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The Time To Live of the SOA Record in seconds. Defaults to `3600`.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)


