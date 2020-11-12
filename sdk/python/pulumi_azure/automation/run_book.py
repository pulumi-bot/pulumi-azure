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

__all__ = ['RunBook']


class RunBook(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 job_schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunBookJobScheduleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_progress: Optional[pulumi.Input[bool]] = None,
                 log_verbose: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 publish_content_link: Optional[pulumi.Input[pulumi.InputType['RunBookPublishContentLinkArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 runbook_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Automation Runbook.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Basic")
        example_run_book = azure.automation.RunBook("exampleRunBook",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            log_verbose=True,
            log_progress=True,
            description="This is an example runbook",
            runbook_type="PowerShellWorkflow",
            publish_content_link=azure.automation.RunBookPublishContentLinkArgs(
                uri="https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1",
            ))
        ```

        ## Import

        Automation Runbooks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:automation/runBook:RunBook Get-AzureVMTutorial /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content: The desired content of the runbook.
        :param pulumi.Input[str] description: A description for this credential.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] log_progress: Progress log option.
        :param pulumi.Input[bool] log_verbose: Verbose log option.
        :param pulumi.Input[str] name: Specifies the name of the Runbook. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['RunBookPublishContentLinkArgs']] publish_content_link: The published runbook content link.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runbook_type: The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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
            __props__['content'] = content
            __props__['description'] = description
            __props__['job_schedules'] = job_schedules
            __props__['location'] = location
            if log_progress is None:
                raise TypeError("Missing required property 'log_progress'")
            __props__['log_progress'] = log_progress
            if log_verbose is None:
                raise TypeError("Missing required property 'log_verbose'")
            __props__['log_verbose'] = log_verbose
            __props__['name'] = name
            __props__['publish_content_link'] = publish_content_link
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if runbook_type is None:
                raise TypeError("Missing required property 'runbook_type'")
            __props__['runbook_type'] = runbook_type
            __props__['tags'] = tags
        super(RunBook, __self__).__init__(
            'azure:automation/runBook:RunBook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            automation_account_name: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            job_schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunBookJobScheduleArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            log_progress: Optional[pulumi.Input[bool]] = None,
            log_verbose: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            publish_content_link: Optional[pulumi.Input[pulumi.InputType['RunBookPublishContentLinkArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            runbook_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'RunBook':
        """
        Get an existing RunBook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content: The desired content of the runbook.
        :param pulumi.Input[str] description: A description for this credential.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] log_progress: Progress log option.
        :param pulumi.Input[bool] log_verbose: Verbose log option.
        :param pulumi.Input[str] name: Specifies the name of the Runbook. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['RunBookPublishContentLinkArgs']] publish_content_link: The published runbook content link.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runbook_type: The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["content"] = content
        __props__["description"] = description
        __props__["job_schedules"] = job_schedules
        __props__["location"] = location
        __props__["log_progress"] = log_progress
        __props__["log_verbose"] = log_verbose
        __props__["name"] = name
        __props__["publish_content_link"] = publish_content_link
        __props__["resource_group_name"] = resource_group_name
        __props__["runbook_type"] = runbook_type
        __props__["tags"] = tags
        return RunBook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> pulumi.Output[str]:
        """
        The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "automation_account_name")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The desired content of the runbook.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for this credential.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="jobSchedules")
    def job_schedules(self) -> pulumi.Output[Sequence['outputs.RunBookJobSchedule']]:
        return pulumi.get(self, "job_schedules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logProgress")
    def log_progress(self) -> pulumi.Output[bool]:
        """
        Progress log option.
        """
        return pulumi.get(self, "log_progress")

    @property
    @pulumi.getter(name="logVerbose")
    def log_verbose(self) -> pulumi.Output[bool]:
        """
        Verbose log option.
        """
        return pulumi.get(self, "log_verbose")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Runbook. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publishContentLink")
    def publish_content_link(self) -> pulumi.Output[Optional['outputs.RunBookPublishContentLink']]:
        """
        The published runbook content link.
        """
        return pulumi.get(self, "publish_content_link")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="runbookType")
    def runbook_type(self) -> pulumi.Output[str]:
        """
        The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        """
        return pulumi.get(self, "runbook_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

