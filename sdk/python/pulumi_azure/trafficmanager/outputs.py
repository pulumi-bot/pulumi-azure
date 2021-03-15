# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'EndpointCustomHeader',
    'EndpointSubnet',
    'ProfileDnsConfig',
    'ProfileMonitorConfig',
    'ProfileMonitorConfigCustomHeader',
]

@pulumi.output_type
class EndpointCustomHeader(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        :param str name: The name of the custom header.
        :param str value: The value of custom header. Applicable for Http and Https protocol.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the custom header.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of custom header. Applicable for Http and Https protocol.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EndpointSubnet(dict):
    def __init__(__self__, *,
                 first: str,
                 last: Optional[str] = None,
                 scope: Optional[int] = None):
        """
        :param str first: The First IP....
        :param str last: The Last IP...
        :param int scope: The Scope...
        """
        pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def first(self) -> str:
        """
        The First IP....
        """
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[str]:
        """
        The Last IP...
        """
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def scope(self) -> Optional[int]:
        """
        The Scope...
        """
        return pulumi.get(self, "scope")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProfileDnsConfig(dict):
    def __init__(__self__, *,
                 relative_name: str,
                 ttl: int):
        """
        :param str relative_name: The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
        :param int ttl: The TTL value of the Profile used by Local DNS resolvers and clients.
        """
        pulumi.set(__self__, "relative_name", relative_name)
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="relativeName")
    def relative_name(self) -> str:
        """
        The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "relative_name")

    @property
    @pulumi.getter
    def ttl(self) -> int:
        """
        The TTL value of the Profile used by Local DNS resolvers and clients.
        """
        return pulumi.get(self, "ttl")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProfileMonitorConfig(dict):
    def __init__(__self__, *,
                 port: int,
                 protocol: str,
                 custom_headers: Optional[Sequence['outputs.ProfileMonitorConfigCustomHeader']] = None,
                 expected_status_code_ranges: Optional[Sequence[str]] = None,
                 interval_in_seconds: Optional[int] = None,
                 path: Optional[str] = None,
                 timeout_in_seconds: Optional[int] = None,
                 tolerated_number_of_failures: Optional[int] = None):
        """
        :param int port: The port number used by the monitoring checks.
        :param str protocol: The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
        :param Sequence['ProfileMonitorConfigCustomHeaderArgs'] custom_headers: One or more `custom_header` blocks as defined below.
        :param Sequence[str] expected_status_code_ranges: A list of status code ranges in the format of `100-101`.
        :param int interval_in_seconds: The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
        :param str path: The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
        :param int timeout_in_seconds: The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
        :param int tolerated_number_of_failures: The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        if custom_headers is not None:
            pulumi.set(__self__, "custom_headers", custom_headers)
        if expected_status_code_ranges is not None:
            pulumi.set(__self__, "expected_status_code_ranges", expected_status_code_ranges)
        if interval_in_seconds is not None:
            pulumi.set(__self__, "interval_in_seconds", interval_in_seconds)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if timeout_in_seconds is not None:
            pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)
        if tolerated_number_of_failures is not None:
            pulumi.set(__self__, "tolerated_number_of_failures", tolerated_number_of_failures)

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port number used by the monitoring checks.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="customHeaders")
    def custom_headers(self) -> Optional[Sequence['outputs.ProfileMonitorConfigCustomHeader']]:
        """
        One or more `custom_header` blocks as defined below.
        """
        return pulumi.get(self, "custom_headers")

    @property
    @pulumi.getter(name="expectedStatusCodeRanges")
    def expected_status_code_ranges(self) -> Optional[Sequence[str]]:
        """
        A list of status code ranges in the format of `100-101`.
        """
        return pulumi.get(self, "expected_status_code_ranges")

    @property
    @pulumi.getter(name="intervalInSeconds")
    def interval_in_seconds(self) -> Optional[int]:
        """
        The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
        """
        return pulumi.get(self, "interval_in_seconds")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[int]:
        """
        The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
        """
        return pulumi.get(self, "timeout_in_seconds")

    @property
    @pulumi.getter(name="toleratedNumberOfFailures")
    def tolerated_number_of_failures(self) -> Optional[int]:
        """
        The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
        """
        return pulumi.get(self, "tolerated_number_of_failures")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProfileMonitorConfigCustomHeader(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        :param str name: The name of the custom header.
        :param str value: The value of custom header. Applicable for Http and Https protocol.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the custom header.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of custom header. Applicable for Http and Https protocol.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


