# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['JobSchedule']


class JobSchedule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 job_schedule_id: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 run_on: Optional[pulumi.Input[str]] = None,
                 runbook_name: Optional[pulumi.Input[str]] = None,
                 schedule_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Links an Automation Runbook and Schedule.

        ## Example Usage

        This is an example of just the Job Schedule.

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.automation.JobSchedule("example",
            automation_account_name="tf-automation-account",
            parameters={
                "resourcegroup": "tf-rgr-vm",
                "vmname": "TF-VM-01",
            },
            resource_group_name="tf-rgr-automation",
            runbook_name="Get-VirtualMachine",
            schedule_name="hour")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] job_schedule_id: The UUID identifying the Automation Job Schedule.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] run_on: Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runbook_name: The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
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

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            __props__['job_schedule_id'] = job_schedule_id
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['run_on'] = run_on
            if runbook_name is None:
                raise TypeError("Missing required property 'runbook_name'")
            __props__['runbook_name'] = runbook_name
            if schedule_name is None:
                raise TypeError("Missing required property 'schedule_name'")
            __props__['schedule_name'] = schedule_name
        super(JobSchedule, __self__).__init__(
            'azure:automation/jobSchedule:JobSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            automation_account_name: Optional[pulumi.Input[str]] = None,
            job_schedule_id: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            run_on: Optional[pulumi.Input[str]] = None,
            runbook_name: Optional[pulumi.Input[str]] = None,
            schedule_name: Optional[pulumi.Input[str]] = None) -> 'JobSchedule':
        """
        Get an existing JobSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] job_schedule_id: The UUID identifying the Automation Job Schedule.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] run_on: Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runbook_name: The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["job_schedule_id"] = job_schedule_id
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["run_on"] = run_on
        __props__["runbook_name"] = runbook_name
        __props__["schedule_name"] = schedule_name
        return JobSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> str:
        """
        The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "automation_account_name")

    @property
    @pulumi.getter(name="jobScheduleId")
    def job_schedule_id(self) -> str:
        """
        The UUID identifying the Automation Job Schedule.
        """
        return pulumi.get(self, "job_schedule_id")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, str]]:
        """
        A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="runOn")
    def run_on(self) -> Optional[str]:
        """
        Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "run_on")

    @property
    @pulumi.getter(name="runbookName")
    def runbook_name(self) -> str:
        """
        The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "runbook_name")

    @property
    @pulumi.getter(name="scheduleName")
    def schedule_name(self) -> str:
        return pulumi.get(self, "schedule_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

