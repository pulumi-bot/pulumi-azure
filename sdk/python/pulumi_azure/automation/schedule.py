# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Schedule(pulumi.CustomResource):
    automation_account_name: pulumi.Output[str]
    """
    The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
    """
    description: pulumi.Output[str]
    """
    A description for this Schedule.
    """
    expiry_time: pulumi.Output[str]
    """
    The end time of the schedule.
    """
    frequency: pulumi.Output[str]
    """
    The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
    """
    interval: pulumi.Output[float]
    """
    The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
    """
    month_days: pulumi.Output[list]
    """
    List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
    """
    monthly_occurrences: pulumi.Output[list]
    """
    List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthly_occurrence` block supports fields documented below.

      * `day` (`str`) - Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
      * `occurrence` (`float`) - Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Schedule. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
    """
    start_time: pulumi.Output[str]
    """
    Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
    """
    timezone: pulumi.Output[str]
    """
    The timezone of the start time. Defaults to `UTC`. For possible values see: https://msdn.microsoft.com/en-us/library/ms912391(v=winembedded.11).aspx
    """
    week_days: pulumi.Output[list]
    """
    List of days of the week that the job should execute on. Only valid when frequency is `Week`.
    """
    def __init__(__self__, resource_name, opts=None, automation_account_name=None, description=None, expiry_time=None, frequency=None, interval=None, month_days=None, monthly_occurrences=None, name=None, resource_group_name=None, start_time=None, timezone=None, week_days=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Automation Schedule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=[{
                "name": "Basic",
            }])
        example_schedule = azure.automation.Schedule("exampleSchedule",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            frequency="Week",
            interval=1,
            timezone="Central Europe Standard Time",
            start_time="2014-04-15T18:00:15+02:00",
            description="This is an example schedule",
            week_days=["Friday"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: A description for this Schedule.
        :param pulumi.Input[str] expiry_time: The end time of the schedule.
        :param pulumi.Input[str] frequency: The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
        :param pulumi.Input[float] interval: The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
        :param pulumi.Input[list] month_days: List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
        :param pulumi.Input[list] monthly_occurrences: List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthly_occurrence` block supports fields documented below.
        :param pulumi.Input[str] name: Specifies the name of the Schedule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_time: Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
        :param pulumi.Input[str] timezone: The timezone of the start time. Defaults to `UTC`. For possible values see: https://msdn.microsoft.com/en-us/library/ms912391(v=winembedded.11).aspx
        :param pulumi.Input[list] week_days: List of days of the week that the job should execute on. Only valid when frequency is `Week`.

        The **monthly_occurrences** object supports the following:

          * `day` (`pulumi.Input[str]`) - Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
          * `occurrence` (`pulumi.Input[float]`) - Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            __props__['description'] = description
            __props__['expiry_time'] = expiry_time
            if frequency is None:
                raise TypeError("Missing required property 'frequency'")
            __props__['frequency'] = frequency
            __props__['interval'] = interval
            __props__['month_days'] = month_days
            __props__['monthly_occurrences'] = monthly_occurrences
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['start_time'] = start_time
            __props__['timezone'] = timezone
            __props__['week_days'] = week_days
        super(Schedule, __self__).__init__(
            'azure:automation/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, automation_account_name=None, description=None, expiry_time=None, frequency=None, interval=None, month_days=None, monthly_occurrences=None, name=None, resource_group_name=None, start_time=None, timezone=None, week_days=None):
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: A description for this Schedule.
        :param pulumi.Input[str] expiry_time: The end time of the schedule.
        :param pulumi.Input[str] frequency: The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
        :param pulumi.Input[float] interval: The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
        :param pulumi.Input[list] month_days: List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
        :param pulumi.Input[list] monthly_occurrences: List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthly_occurrence` block supports fields documented below.
        :param pulumi.Input[str] name: Specifies the name of the Schedule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_time: Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
        :param pulumi.Input[str] timezone: The timezone of the start time. Defaults to `UTC`. For possible values see: https://msdn.microsoft.com/en-us/library/ms912391(v=winembedded.11).aspx
        :param pulumi.Input[list] week_days: List of days of the week that the job should execute on. Only valid when frequency is `Week`.

        The **monthly_occurrences** object supports the following:

          * `day` (`pulumi.Input[str]`) - Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
          * `occurrence` (`pulumi.Input[float]`) - Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["description"] = description
        __props__["expiry_time"] = expiry_time
        __props__["frequency"] = frequency
        __props__["interval"] = interval
        __props__["month_days"] = month_days
        __props__["monthly_occurrences"] = monthly_occurrences
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["start_time"] = start_time
        __props__["timezone"] = timezone
        __props__["week_days"] = week_days
        return Schedule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
