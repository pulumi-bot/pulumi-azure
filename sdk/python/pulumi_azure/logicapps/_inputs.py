# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'ActionHttpRunAfterArgs',
    'TriggerRecurrenceScheduleArgs',
]

@pulumi.input_type
class ActionHttpRunAfterArgs:
    def __init__(__self__, *,
                 action_name: pulumi.Input[str],
                 action_result: pulumi.Input[str]):
        """
        :param pulumi.Input[str] action_name: Specifies the name of the precedent HTTP Action.
        :param pulumi.Input[str] action_result: Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
        """
        pulumi.set(__self__, "action_name", action_name)
        pulumi.set(__self__, "action_result", action_result)

    @property
    @pulumi.getter(name="actionName")
    def action_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the precedent HTTP Action.
        """
        return pulumi.get(self, "action_name")

    @action_name.setter
    def action_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "action_name", value)

    @property
    @pulumi.getter(name="actionResult")
    def action_result(self) -> pulumi.Input[str]:
        """
        Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
        """
        return pulumi.get(self, "action_result")

    @action_result.setter
    def action_result(self, value: pulumi.Input[str]):
        pulumi.set(self, "action_result", value)


@pulumi.input_type
class TriggerRecurrenceScheduleArgs:
    def __init__(__self__, *,
                 at_these_hours: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 at_these_minutes: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 on_these_days: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] at_these_hours: Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] at_these_minutes: Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] on_these_days: Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
        """
        if at_these_hours is not None:
            pulumi.set(__self__, "at_these_hours", at_these_hours)
        if at_these_minutes is not None:
            pulumi.set(__self__, "at_these_minutes", at_these_minutes)
        if on_these_days is not None:
            pulumi.set(__self__, "on_these_days", on_these_days)

    @property
    @pulumi.getter(name="atTheseHours")
    def at_these_hours(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
        """
        return pulumi.get(self, "at_these_hours")

    @at_these_hours.setter
    def at_these_hours(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "at_these_hours", value)

    @property
    @pulumi.getter(name="atTheseMinutes")
    def at_these_minutes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
        """
        return pulumi.get(self, "at_these_minutes")

    @at_these_minutes.setter
    def at_these_minutes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "at_these_minutes", value)

    @property
    @pulumi.getter(name="onTheseDays")
    def on_these_days(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
        """
        return pulumi.get(self, "on_these_days")

    @on_these_days.setter
    def on_these_days(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "on_these_days", value)


