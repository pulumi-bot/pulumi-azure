# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'EndpointCustomHeaderArgs',
    'EndpointSubnetArgs',
    'ProfileDnsConfigArgs',
    'ProfileMonitorConfigArgs',
    'ProfileMonitorConfigCustomHeaderArgs',
]

@pulumi.input_type
class EndpointCustomHeaderArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the custom header.
    """
    value: pulumi.Input[str] = pulumi.input_property("value")
    """
    The value of custom header. Applicable for Http and Https protocol.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], value: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] name: The name of the custom header.
        :param pulumi.Input[str] value: The value of custom header. Applicable for Http and Https protocol.
        """
        __self__.name = name
        __self__.value = value

@pulumi.input_type
class EndpointSubnetArgs:
    first: pulumi.Input[str] = pulumi.input_property("first")
    """
    The First IP....
    """
    last: Optional[pulumi.Input[str]] = pulumi.input_property("last")
    """
    The Last IP...
    """
    scope: Optional[pulumi.Input[float]] = pulumi.input_property("scope")
    """
    The Scope...
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, first: pulumi.Input[str], last: Optional[pulumi.Input[str]] = None, scope: Optional[pulumi.Input[float]] = None) -> None:
        """
        :param pulumi.Input[str] first: The First IP....
        :param pulumi.Input[str] last: The Last IP...
        :param pulumi.Input[float] scope: The Scope...
        """
        __self__.first = first
        __self__.last = last
        __self__.scope = scope

@pulumi.input_type
class ProfileDnsConfigArgs:
    relative_name: pulumi.Input[str] = pulumi.input_property("relativeName")
    """
    The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
    """
    ttl: pulumi.Input[float] = pulumi.input_property("ttl")
    """
    The TTL value of the Profile used by Local DNS resolvers and clients.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, relative_name: pulumi.Input[str], ttl: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[str] relative_name: The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
        :param pulumi.Input[float] ttl: The TTL value of the Profile used by Local DNS resolvers and clients.
        """
        __self__.relative_name = relative_name
        __self__.ttl = ttl

@pulumi.input_type
class ProfileMonitorConfigArgs:
    port: pulumi.Input[float] = pulumi.input_property("port")
    """
    The port number used by the monitoring checks.
    """
    protocol: pulumi.Input[str] = pulumi.input_property("protocol")
    """
    The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
    """
    custom_headers: Optional[pulumi.Input[List[pulumi.Input['ProfileMonitorConfigCustomHeaderArgs']]]] = pulumi.input_property("customHeaders")
    """
    One or more `custom_header` blocks as defined below.
    """
    expected_status_code_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("expectedStatusCodeRanges")
    """
    A list of status code ranges in the format of `100-101`.
    """
    interval_in_seconds: Optional[pulumi.Input[float]] = pulumi.input_property("intervalInSeconds")
    """
    The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
    """
    path: Optional[pulumi.Input[str]] = pulumi.input_property("path")
    """
    The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
    """
    timeout_in_seconds: Optional[pulumi.Input[float]] = pulumi.input_property("timeoutInSeconds")
    """
    The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
    """
    tolerated_number_of_failures: Optional[pulumi.Input[float]] = pulumi.input_property("toleratedNumberOfFailures")
    """
    The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, port: pulumi.Input[float], protocol: pulumi.Input[str], custom_headers: Optional[pulumi.Input[List[pulumi.Input['ProfileMonitorConfigCustomHeaderArgs']]]] = None, expected_status_code_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, interval_in_seconds: Optional[pulumi.Input[float]] = None, path: Optional[pulumi.Input[str]] = None, timeout_in_seconds: Optional[pulumi.Input[float]] = None, tolerated_number_of_failures: Optional[pulumi.Input[float]] = None) -> None:
        """
        :param pulumi.Input[float] port: The port number used by the monitoring checks.
        :param pulumi.Input[str] protocol: The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
        :param pulumi.Input[List[pulumi.Input['ProfileMonitorConfigCustomHeaderArgs']]] custom_headers: One or more `custom_header` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] expected_status_code_ranges: A list of status code ranges in the format of `100-101`.
        :param pulumi.Input[float] interval_in_seconds: The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
        :param pulumi.Input[str] path: The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
        :param pulumi.Input[float] timeout_in_seconds: The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
        :param pulumi.Input[float] tolerated_number_of_failures: The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
        """
        __self__.port = port
        __self__.protocol = protocol
        __self__.custom_headers = custom_headers
        __self__.expected_status_code_ranges = expected_status_code_ranges
        __self__.interval_in_seconds = interval_in_seconds
        __self__.path = path
        __self__.timeout_in_seconds = timeout_in_seconds
        __self__.tolerated_number_of_failures = tolerated_number_of_failures

@pulumi.input_type
class ProfileMonitorConfigCustomHeaderArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the custom header.
    """
    value: pulumi.Input[str] = pulumi.input_property("value")
    """
    The value of custom header. Applicable for Http and Https protocol.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], value: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] name: The name of the custom header.
        :param pulumi.Input[str] value: The value of custom header. Applicable for Http and Https protocol.
        """
        __self__.name = name
        __self__.value = value

