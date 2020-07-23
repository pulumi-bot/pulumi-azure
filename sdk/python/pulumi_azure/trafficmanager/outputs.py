# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.output_type
class EndpointCustomHeader(dict):
    name: str = pulumi.output_property("name")
    """
    The name of the custom header.
    """
    value: str = pulumi.output_property("value")
    """
    The value of custom header. Applicable for Http and Https protocol.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EndpointSubnet(dict):
    first: str = pulumi.output_property("first")
    """
    The First IP....
    """
    last: Optional[str] = pulumi.output_property("last")
    """
    The Last IP...
    """
    scope: Optional[float] = pulumi.output_property("scope")
    """
    The Scope...
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProfileDnsConfig(dict):
    relative_name: str = pulumi.output_property("relativeName")
    """
    The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
    """
    ttl: float = pulumi.output_property("ttl")
    """
    The TTL value of the Profile used by Local DNS resolvers and clients.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProfileMonitorConfig(dict):
    custom_headers: Optional[List['outputs.ProfileMonitorConfigCustomHeader']] = pulumi.output_property("customHeaders")
    """
    One or more `custom_header` blocks as defined below.
    """
    expected_status_code_ranges: Optional[List[str]] = pulumi.output_property("expectedStatusCodeRanges")
    """
    A list of status code ranges in the format of `100-101`.
    """
    interval_in_seconds: Optional[float] = pulumi.output_property("intervalInSeconds")
    """
    The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
    """
    path: Optional[str] = pulumi.output_property("path")
    """
    The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
    """
    port: float = pulumi.output_property("port")
    """
    The port number used by the monitoring checks.
    """
    protocol: str = pulumi.output_property("protocol")
    """
    The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
    """
    timeout_in_seconds: Optional[float] = pulumi.output_property("timeoutInSeconds")
    """
    The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
    """
    tolerated_number_of_failures: Optional[float] = pulumi.output_property("toleratedNumberOfFailures")
    """
    The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProfileMonitorConfigCustomHeader(dict):
    name: str = pulumi.output_property("name")
    """
    The name of the custom header.
    """
    value: str = pulumi.output_property("value")
    """
    The value of custom header. Applicable for Http and Https protocol.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

