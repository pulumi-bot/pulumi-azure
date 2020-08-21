# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'PolicyFileShareBackupArgs',
    'PolicyFileShareRetentionDailyArgs',
    'PolicyVMBackupArgs',
    'PolicyVMRetentionDailyArgs',
    'PolicyVMRetentionMonthlyArgs',
    'PolicyVMRetentionWeeklyArgs',
    'PolicyVMRetentionYearlyArgs',
]

@pulumi.input_type
class PolicyFileShareBackupArgs:
    def __init__(__self__, *,
                 frequency: pulumi.Input[str],
                 time: pulumi.Input[str]):
        """
        :param pulumi.Input[str] frequency: Sets the backup frequency. Currently, only `Daily` is supported
        """
        pulumi.set(__self__, "frequency", frequency)
        pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Input[str]:
        """
        Sets the backup frequency. Currently, only `Daily` is supported
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: pulumi.Input[str]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def time(self) -> pulumi.Input[str]:
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: pulumi.Input[str]):
        pulumi.set(self, "time", value)


@pulumi.input_type
class PolicyFileShareRetentionDailyArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float]):
        """
        :param pulumi.Input[float] count: The number of daily backups to keep. Must be between `1` and `180` (inclusive)
        """
        pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        The number of daily backups to keep. Must be between `1` and `180` (inclusive)
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[float]):
        pulumi.set(self, "count", value)


@pulumi.input_type
class PolicyVMBackupArgs:
    def __init__(__self__, *,
                 frequency: pulumi.Input[str],
                 time: pulumi.Input[str],
                 weekdays: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] frequency: Sets the backup frequency. Must be either `Daily` or`Weekly`.
        :param pulumi.Input[str] time: The time of day to perform the backup in 24hour format.
        :param pulumi.Input[List[pulumi.Input[str]]] weekdays: The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        """
        pulumi.set(__self__, "frequency", frequency)
        pulumi.set(__self__, "time", time)
        if weekdays is not None:
            pulumi.set(__self__, "weekdays", weekdays)

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Input[str]:
        """
        Sets the backup frequency. Must be either `Daily` or`Weekly`.
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: pulumi.Input[str]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def time(self) -> pulumi.Input[str]:
        """
        The time of day to perform the backup in 24hour format.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: pulumi.Input[str]):
        pulumi.set(self, "time", value)

    @property
    @pulumi.getter
    def weekdays(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        """
        return pulumi.get(self, "weekdays")

    @weekdays.setter
    def weekdays(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "weekdays", value)


@pulumi.input_type
class PolicyVMRetentionDailyArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float]):
        """
        :param pulumi.Input[float] count: The number of yearly backups to keep. Must be between `1` and `9999`
        """
        pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        The number of yearly backups to keep. Must be between `1` and `9999`
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[float]):
        pulumi.set(self, "count", value)


@pulumi.input_type
class PolicyVMRetentionMonthlyArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float],
                 weekdays: pulumi.Input[List[pulumi.Input[str]]],
                 weeks: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[float] count: The number of yearly backups to keep. Must be between `1` and `9999`
        :param pulumi.Input[List[pulumi.Input[str]]] weekdays: The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        :param pulumi.Input[List[pulumi.Input[str]]] weeks: The weeks of the month to retain backups of. Must be one of `First`, `Second`, `Third`, `Fourth`, `Last`.
        """
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "weekdays", weekdays)
        pulumi.set(__self__, "weeks", weeks)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        The number of yearly backups to keep. Must be between `1` and `9999`
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[float]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def weekdays(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        """
        return pulumi.get(self, "weekdays")

    @weekdays.setter
    def weekdays(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "weekdays", value)

    @property
    @pulumi.getter
    def weeks(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The weeks of the month to retain backups of. Must be one of `First`, `Second`, `Third`, `Fourth`, `Last`.
        """
        return pulumi.get(self, "weeks")

    @weeks.setter
    def weeks(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "weeks", value)


@pulumi.input_type
class PolicyVMRetentionWeeklyArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float],
                 weekdays: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[float] count: The number of yearly backups to keep. Must be between `1` and `9999`
        :param pulumi.Input[List[pulumi.Input[str]]] weekdays: The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        """
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "weekdays", weekdays)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        The number of yearly backups to keep. Must be between `1` and `9999`
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[float]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def weekdays(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        """
        return pulumi.get(self, "weekdays")

    @weekdays.setter
    def weekdays(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "weekdays", value)


@pulumi.input_type
class PolicyVMRetentionYearlyArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float],
                 months: pulumi.Input[List[pulumi.Input[str]]],
                 weekdays: pulumi.Input[List[pulumi.Input[str]]],
                 weeks: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[float] count: The number of yearly backups to keep. Must be between `1` and `9999`
        :param pulumi.Input[List[pulumi.Input[str]]] months: The months of the year to retain backups of. Must be one of `January`, `February`, `March`, `April`, `May`, `June`, `July`, `Augest`, `September`, `October`, `November` and `December`.
        :param pulumi.Input[List[pulumi.Input[str]]] weekdays: The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        :param pulumi.Input[List[pulumi.Input[str]]] weeks: The weeks of the month to retain backups of. Must be one of `First`, `Second`, `Third`, `Fourth`, `Last`.
        """
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "months", months)
        pulumi.set(__self__, "weekdays", weekdays)
        pulumi.set(__self__, "weeks", weeks)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        The number of yearly backups to keep. Must be between `1` and `9999`
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[float]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def months(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The months of the year to retain backups of. Must be one of `January`, `February`, `March`, `April`, `May`, `June`, `July`, `Augest`, `September`, `October`, `November` and `December`.
        """
        return pulumi.get(self, "months")

    @months.setter
    def months(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "months", value)

    @property
    @pulumi.getter
    def weekdays(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        """
        return pulumi.get(self, "weekdays")

    @weekdays.setter
    def weekdays(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "weekdays", value)

    @property
    @pulumi.getter
    def weeks(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The weeks of the month to retain backups of. Must be one of `First`, `Second`, `Third`, `Fourth`, `Last`.
        """
        return pulumi.get(self, "weeks")

    @weeks.setter
    def weeks(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "weeks", value)


