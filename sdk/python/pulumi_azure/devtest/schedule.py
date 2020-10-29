# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Schedule']


class Schedule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_recurrence: Optional[pulumi.Input[pulumi.InputType['ScheduleDailyRecurrenceArgs']]] = None,
                 hourly_recurrence: Optional[pulumi.Input[pulumi.InputType['ScheduleHourlyRecurrenceArgs']]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_settings: Optional[pulumi.Input[pulumi.InputType['ScheduleNotificationSettingsArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 task_type: Optional[pulumi.Input[str]] = None,
                 time_zone_id: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['ScheduleWeeklyRecurrenceArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages automated startup and shutdown schedules for Azure Dev Test Lab.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lab_name: The name of the dev test lab. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The location where the schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the dev test lab schedule. Valid value for name depends on the `task_type`. For instance for task_type `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] task_type: The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
        :param pulumi.Input[str] time_zone_id: The time zone ID (e.g. Pacific Standard time).
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['daily_recurrence'] = daily_recurrence
            __props__['hourly_recurrence'] = hourly_recurrence
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            __props__['name'] = name
            if notification_settings is None:
                raise TypeError("Missing required property 'notification_settings'")
            __props__['notification_settings'] = notification_settings
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
            __props__['tags'] = tags
            if task_type is None:
                raise TypeError("Missing required property 'task_type'")
            __props__['task_type'] = task_type
            if time_zone_id is None:
                raise TypeError("Missing required property 'time_zone_id'")
            __props__['time_zone_id'] = time_zone_id
            __props__['weekly_recurrence'] = weekly_recurrence
        super(Schedule, __self__).__init__(
            'azure:devtest/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            daily_recurrence: Optional[pulumi.Input[pulumi.InputType['ScheduleDailyRecurrenceArgs']]] = None,
            hourly_recurrence: Optional[pulumi.Input[pulumi.InputType['ScheduleHourlyRecurrenceArgs']]] = None,
            lab_name: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_settings: Optional[pulumi.Input[pulumi.InputType['ScheduleNotificationSettingsArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            task_type: Optional[pulumi.Input[str]] = None,
            time_zone_id: Optional[pulumi.Input[str]] = None,
            weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['ScheduleWeeklyRecurrenceArgs']]] = None) -> 'Schedule':
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lab_name: The name of the dev test lab. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The location where the schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the dev test lab schedule. Valid value for name depends on the `task_type`. For instance for task_type `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] task_type: The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
        :param pulumi.Input[str] time_zone_id: The time zone ID (e.g. Pacific Standard time).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["daily_recurrence"] = daily_recurrence
        __props__["hourly_recurrence"] = hourly_recurrence
        __props__["lab_name"] = lab_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["notification_settings"] = notification_settings
        __props__["resource_group_name"] = resource_group_name
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["task_type"] = task_type
        __props__["time_zone_id"] = time_zone_id
        __props__["weekly_recurrence"] = weekly_recurrence
        return Schedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> pulumi.Output[Optional['outputs.ScheduleDailyRecurrence']]:
        return pulumi.get(self, "daily_recurrence")

    @property
    @pulumi.getter(name="hourlyRecurrence")
    def hourly_recurrence(self) -> pulumi.Output[Optional['outputs.ScheduleHourlyRecurrence']]:
        return pulumi.get(self, "hourly_recurrence")

    @property
    @pulumi.getter(name="labName")
    def lab_name(self) -> pulumi.Output[str]:
        """
        The name of the dev test lab. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location where the schedule is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the dev test lab schedule. Valid value for name depends on the `task_type`. For instance for task_type `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationSettings")
    def notification_settings(self) -> pulumi.Output['outputs.ScheduleNotificationSettings']:
        return pulumi.get(self, "notification_settings")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskType")
    def task_type(self) -> pulumi.Output[str]:
        """
        The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
        """
        return pulumi.get(self, "task_type")

    @property
    @pulumi.getter(name="timeZoneId")
    def time_zone_id(self) -> pulumi.Output[str]:
        """
        The time zone ID (e.g. Pacific Standard time).
        """
        return pulumi.get(self, "time_zone_id")

    @property
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> pulumi.Output[Optional['outputs.ScheduleWeeklyRecurrence']]:
        return pulumi.get(self, "weekly_recurrence")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

