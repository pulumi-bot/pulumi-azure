# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'CaaRecordRecordArgs',
    'MxRecordRecordArgs',
    'SrvRecordRecordArgs',
    'TxtRecordRecordArgs',
]

@pulumi.input_type
class CaaRecordRecordArgs:
    def __init__(__self__, *,
                 flags: pulumi.Input[float],
                 tag: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[float] flags: Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
        :param pulumi.Input[str] tag: A property tag, options are issue, issuewild and iodef.
        :param pulumi.Input[str] value: A property value such as a registrar domain.
        """
        pulumi.set(__self__, "flags", flags)
        pulumi.set(__self__, "tag", tag)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def flags(self) -> pulumi.Input[float]:
        """
        Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
        """
        return pulumi.get(self, "flags")

    @flags.setter
    def flags(self, value: pulumi.Input[float]):
        pulumi.set(self, "flags", value)

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Input[str]:
        """
        A property tag, options are issue, issuewild and iodef.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        A property value such as a registrar domain.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class MxRecordRecordArgs:
    def __init__(__self__, *,
                 exchange: pulumi.Input[str],
                 preference: pulumi.Input[str]):
        """
        :param pulumi.Input[str] exchange: The mail server responsible for the domain covered by the MX record.
        :param pulumi.Input[str] preference: String representing the "preference” value of the MX records. Records with lower preference value take priority.
        """
        pulumi.set(__self__, "exchange", exchange)
        pulumi.set(__self__, "preference", preference)

    @property
    @pulumi.getter
    def exchange(self) -> pulumi.Input[str]:
        """
        The mail server responsible for the domain covered by the MX record.
        """
        return pulumi.get(self, "exchange")

    @exchange.setter
    def exchange(self, value: pulumi.Input[str]):
        pulumi.set(self, "exchange", value)

    @property
    @pulumi.getter
    def preference(self) -> pulumi.Input[str]:
        """
        String representing the "preference” value of the MX records. Records with lower preference value take priority.
        """
        return pulumi.get(self, "preference")

    @preference.setter
    def preference(self, value: pulumi.Input[str]):
        pulumi.set(self, "preference", value)


@pulumi.input_type
class SrvRecordRecordArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[float],
                 priority: pulumi.Input[float],
                 target: pulumi.Input[str],
                 weight: pulumi.Input[float]):
        """
        :param pulumi.Input[float] port: Port the service is listening on.
        :param pulumi.Input[float] priority: Priority of the SRV record.
        :param pulumi.Input[str] target: FQDN of the service.
        :param pulumi.Input[float] weight: Weight of the SRV record.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[float]:
        """
        Port the service is listening on.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[float]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[float]:
        """
        Priority of the SRV record.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[float]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        """
        FQDN of the service.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[float]:
        """
        Weight of the SRV record.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[float]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class TxtRecordRecordArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] value: The value of the record. Max length: 1024 characters
        """
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the record. Max length: 1024 characters
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


