# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ServerIpv4FirewallRuleArgs',
]

@pulumi.input_type
class ServerIpv4FirewallRuleArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 range_end: pulumi.Input[str],
                 range_start: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Specifies the name of the firewall rule.
        :param pulumi.Input[str] range_end: End of the firewall rule range as IPv4 address.
        :param pulumi.Input[str] range_start: Start of the firewall rule range as IPv4 address.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "range_end", range_end)
        pulumi.set(__self__, "range_start", range_start)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the firewall rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rangeEnd")
    def range_end(self) -> pulumi.Input[str]:
        """
        End of the firewall rule range as IPv4 address.
        """
        return pulumi.get(self, "range_end")

    @range_end.setter
    def range_end(self, value: pulumi.Input[str]):
        pulumi.set(self, "range_end", value)

    @property
    @pulumi.getter(name="rangeStart")
    def range_start(self) -> pulumi.Input[str]:
        """
        Start of the firewall rule range as IPv4 address.
        """
        return pulumi.get(self, "range_start")

    @range_start.setter
    def range_start(self, value: pulumi.Input[str]):
        pulumi.set(self, "range_start", value)


