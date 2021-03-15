# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'CaaRecordRecord',
    'MxRecordRecord',
    'SrvRecordRecord',
    'TxtRecordRecord',
    'ZoneSoaRecord',
]

@pulumi.output_type
class CaaRecordRecord(dict):
    def __init__(__self__, *,
                 flags: int,
                 tag: str,
                 value: str):
        """
        :param int flags: Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
        :param str tag: A property tag, options are issue, issuewild and iodef.
        :param str value: A property value such as a registrar domain.
        """
        pulumi.set(__self__, "flags", flags)
        pulumi.set(__self__, "tag", tag)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def flags(self) -> int:
        """
        Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
        """
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter
    def tag(self) -> str:
        """
        A property tag, options are issue, issuewild and iodef.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        A property value such as a registrar domain.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MxRecordRecord(dict):
    def __init__(__self__, *,
                 exchange: str,
                 preference: str):
        """
        :param str exchange: The mail server responsible for the domain covered by the MX record.
        :param str preference: String representing the "preference” value of the MX records. Records with lower preference value take priority.
        """
        pulumi.set(__self__, "exchange", exchange)
        pulumi.set(__self__, "preference", preference)

    @property
    @pulumi.getter
    def exchange(self) -> str:
        """
        The mail server responsible for the domain covered by the MX record.
        """
        return pulumi.get(self, "exchange")

    @property
    @pulumi.getter
    def preference(self) -> str:
        """
        String representing the "preference” value of the MX records. Records with lower preference value take priority.
        """
        return pulumi.get(self, "preference")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SrvRecordRecord(dict):
    def __init__(__self__, *,
                 port: int,
                 priority: int,
                 target: str,
                 weight: int):
        """
        :param int port: Port the service is listening on.
        :param int priority: Priority of the SRV record.
        :param str target: FQDN of the service.
        :param int weight: Weight of the SRV record.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Port the service is listening on.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Priority of the SRV record.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        FQDN of the service.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def weight(self) -> int:
        """
        Weight of the SRV record.
        """
        return pulumi.get(self, "weight")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TxtRecordRecord(dict):
    def __init__(__self__, *,
                 value: str):
        """
        :param str value: The value of the record. Max length: 1024 characters
        """
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the record. Max length: 1024 characters
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ZoneSoaRecord(dict):
    def __init__(__self__, *,
                 email: str,
                 host_name: str,
                 expire_time: Optional[int] = None,
                 fqdn: Optional[str] = None,
                 minimum_ttl: Optional[int] = None,
                 refresh_time: Optional[int] = None,
                 retry_time: Optional[int] = None,
                 serial_number: Optional[int] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 ttl: Optional[int] = None):
        """
        :param str email: The email contact for the SOA record.
        :param str host_name: The domain name of the authoritative name server for the SOA record. Defaults to `ns1-03.azure-dns.com.`.
        :param int expire_time: The expire time for the SOA record. Defaults to `2419200`.
        :param int minimum_ttl: The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `300`.
        :param int refresh_time: The refresh time for the SOA record. Defaults to `3600`.
        :param int retry_time: The retry time for the SOA record. Defaults to `300`.
        :param int serial_number: The serial number for the SOA record. Defaults to `1`.
        :param Mapping[str, str] tags: A mapping of tags to assign to the Record Set.
        :param int ttl: The Time To Live of the SOA Record in seconds. Defaults to `3600`.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "host_name", host_name)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
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
    def email(self) -> str:
        """
        The email contact for the SOA record.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        The domain name of the authoritative name server for the SOA record. Defaults to `ns1-03.azure-dns.com.`.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[int]:
        """
        The expire time for the SOA record. Defaults to `2419200`.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[str]:
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="minimumTtl")
    def minimum_ttl(self) -> Optional[int]:
        """
        The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `300`.
        """
        return pulumi.get(self, "minimum_ttl")

    @property
    @pulumi.getter(name="refreshTime")
    def refresh_time(self) -> Optional[int]:
        """
        The refresh time for the SOA record. Defaults to `3600`.
        """
        return pulumi.get(self, "refresh_time")

    @property
    @pulumi.getter(name="retryTime")
    def retry_time(self) -> Optional[int]:
        """
        The retry time for the SOA record. Defaults to `300`.
        """
        return pulumi.get(self, "retry_time")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[int]:
        """
        The serial number for the SOA record. Defaults to `1`.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the Record Set.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def ttl(self) -> Optional[int]:
        """
        The Time To Live of the SOA Record in seconds. Defaults to `3600`.
        """
        return pulumi.get(self, "ttl")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


