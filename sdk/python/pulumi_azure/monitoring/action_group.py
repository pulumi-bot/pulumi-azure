# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ActionGroup']


class ActionGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arm_role_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupArmRoleReceiverArgs']]]]] = None,
                 automation_runbook_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAutomationRunbookReceiverArgs']]]]] = None,
                 azure_app_push_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureAppPushReceiverArgs']]]]] = None,
                 azure_function_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureFunctionReceiverArgs']]]]] = None,
                 email_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupEmailReceiverArgs']]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 itsm_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupItsmReceiverArgs']]]]] = None,
                 logic_app_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupLogicAppReceiverArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 sms_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupSmsReceiverArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 voice_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupVoiceReceiverArgs']]]]] = None,
                 webhook_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupWebhookReceiverArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Action Group within Azure Monitor.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_action_group = azure.monitoring.ActionGroup("exampleActionGroup",
            resource_group_name=example_resource_group.name,
            short_name="p0action",
            arm_role_receivers=[{
                "name": "armroleaction",
                "roleId": "de139f84-1756-47ae-9be6-808fbbe84772",
                "useCommonAlertSchema": True,
            }],
            automation_runbook_receivers=[{
                "name": "action_name_1",
                "automationAccountId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001",
                "runbook_name": "my runbook",
                "webhookResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001/webhooks/webhook_alert",
                "isGlobalRunbook": True,
                "service_uri": "https://s13events.azure-automation.net/webhooks?token=randomtoken",
                "useCommonAlertSchema": True,
            }],
            azure_app_push_receivers=[{
                "name": "pushtoadmin",
                "email_address": "admin@contoso.com",
            }],
            azure_function_receivers=[{
                "name": "funcaction",
                "functionAppResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-funcapp/providers/Microsoft.Web/sites/funcapp",
                "functionName": "myfunc",
                "httpTriggerUrl": "https://example.com/trigger",
                "useCommonAlertSchema": True,
            }],
            email_receivers=[
                {
                    "name": "sendtoadmin",
                    "email_address": "admin@contoso.com",
                },
                {
                    "name": "sendtodevops",
                    "email_address": "devops@contoso.com",
                    "useCommonAlertSchema": True,
                },
            ],
            itsm_receivers=[{
                "name": "createorupdateticket",
                "workspace_id": "6eee3a18-aac3-40e4-b98e-1f309f329816",
                "connectionId": "53de6956-42b4-41ba-be3c-b154cdf17b13",
                "ticketConfiguration": "{}",
                "region": "southcentralus",
            }],
            logic_app_receivers=[{
                "name": "logicappaction",
                "resource_id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-logicapp/providers/Microsoft.Logic/workflows/logicapp",
                "callbackUrl": "https://logicapptriggerurl/...",
                "useCommonAlertSchema": True,
            }],
            sms_receivers=[{
                "name": "oncallmsg",
                "countryCode": "1",
                "phoneNumber": "1231231234",
            }],
            voice_receivers=[{
                "name": "remotesupport",
                "countryCode": "86",
                "phoneNumber": "13888888888",
            }],
            webhook_receivers=[{
                "name": "callmyapiaswell",
                "service_uri": "http://example.com/alert",
                "useCommonAlertSchema": True,
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupArmRoleReceiverArgs']]]] arm_role_receivers: One or more `arm_role_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAutomationRunbookReceiverArgs']]]] automation_runbook_receivers: One or more `automation_runbook_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureAppPushReceiverArgs']]]] azure_app_push_receivers: One or more `azure_app_push_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureFunctionReceiverArgs']]]] azure_function_receivers: One or more `azure_function_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupEmailReceiverArgs']]]] email_receivers: One or more `email_receiver` blocks as defined below.
        :param pulumi.Input[bool] enabled: Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupItsmReceiverArgs']]]] itsm_receivers: One or more `itsm_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupLogicAppReceiverArgs']]]] logic_app_receivers: One or more `logic_app_receiver` blocks as defined below.
        :param pulumi.Input[str] name: The name of the Action Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Action Group instance.
        :param pulumi.Input[str] short_name: The short name of the action group. This will be used in SMS messages.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupSmsReceiverArgs']]]] sms_receivers: One or more `sms_receiver` blocks as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupVoiceReceiverArgs']]]] voice_receivers: One or more `voice_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupWebhookReceiverArgs']]]] webhook_receivers: One or more `webhook_receiver` blocks as defined below.
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

            __props__['arm_role_receivers'] = arm_role_receivers
            __props__['automation_runbook_receivers'] = automation_runbook_receivers
            __props__['azure_app_push_receivers'] = azure_app_push_receivers
            __props__['azure_function_receivers'] = azure_function_receivers
            __props__['email_receivers'] = email_receivers
            __props__['enabled'] = enabled
            __props__['itsm_receivers'] = itsm_receivers
            __props__['logic_app_receivers'] = logic_app_receivers
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if short_name is None:
                raise TypeError("Missing required property 'short_name'")
            __props__['short_name'] = short_name
            __props__['sms_receivers'] = sms_receivers
            __props__['tags'] = tags
            __props__['voice_receivers'] = voice_receivers
            __props__['webhook_receivers'] = webhook_receivers
        super(ActionGroup, __self__).__init__(
            'azure:monitoring/actionGroup:ActionGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arm_role_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupArmRoleReceiverArgs']]]]] = None,
            automation_runbook_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAutomationRunbookReceiverArgs']]]]] = None,
            azure_app_push_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureAppPushReceiverArgs']]]]] = None,
            azure_function_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureFunctionReceiverArgs']]]]] = None,
            email_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupEmailReceiverArgs']]]]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            itsm_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupItsmReceiverArgs']]]]] = None,
            logic_app_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupLogicAppReceiverArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None,
            sms_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupSmsReceiverArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            voice_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupVoiceReceiverArgs']]]]] = None,
            webhook_receivers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupWebhookReceiverArgs']]]]] = None) -> 'ActionGroup':
        """
        Get an existing ActionGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupArmRoleReceiverArgs']]]] arm_role_receivers: One or more `arm_role_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAutomationRunbookReceiverArgs']]]] automation_runbook_receivers: One or more `automation_runbook_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureAppPushReceiverArgs']]]] azure_app_push_receivers: One or more `azure_app_push_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupAzureFunctionReceiverArgs']]]] azure_function_receivers: One or more `azure_function_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupEmailReceiverArgs']]]] email_receivers: One or more `email_receiver` blocks as defined below.
        :param pulumi.Input[bool] enabled: Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupItsmReceiverArgs']]]] itsm_receivers: One or more `itsm_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupLogicAppReceiverArgs']]]] logic_app_receivers: One or more `logic_app_receiver` blocks as defined below.
        :param pulumi.Input[str] name: The name of the Action Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Action Group instance.
        :param pulumi.Input[str] short_name: The short name of the action group. This will be used in SMS messages.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupSmsReceiverArgs']]]] sms_receivers: One or more `sms_receiver` blocks as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupVoiceReceiverArgs']]]] voice_receivers: One or more `voice_receiver` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionGroupWebhookReceiverArgs']]]] webhook_receivers: One or more `webhook_receiver` blocks as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arm_role_receivers"] = arm_role_receivers
        __props__["automation_runbook_receivers"] = automation_runbook_receivers
        __props__["azure_app_push_receivers"] = azure_app_push_receivers
        __props__["azure_function_receivers"] = azure_function_receivers
        __props__["email_receivers"] = email_receivers
        __props__["enabled"] = enabled
        __props__["itsm_receivers"] = itsm_receivers
        __props__["logic_app_receivers"] = logic_app_receivers
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["short_name"] = short_name
        __props__["sms_receivers"] = sms_receivers
        __props__["tags"] = tags
        __props__["voice_receivers"] = voice_receivers
        __props__["webhook_receivers"] = webhook_receivers
        return ActionGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="armRoleReceivers")
    def arm_role_receivers(self) -> Optional[List['outputs.ActionGroupArmRoleReceiver']]:
        """
        One or more `arm_role_receiver` blocks as defined below.
        """
        return pulumi.get(self, "arm_role_receivers")

    @property
    @pulumi.getter(name="automationRunbookReceivers")
    def automation_runbook_receivers(self) -> Optional[List['outputs.ActionGroupAutomationRunbookReceiver']]:
        """
        One or more `automation_runbook_receiver` blocks as defined below.
        """
        return pulumi.get(self, "automation_runbook_receivers")

    @property
    @pulumi.getter(name="azureAppPushReceivers")
    def azure_app_push_receivers(self) -> Optional[List['outputs.ActionGroupAzureAppPushReceiver']]:
        """
        One or more `azure_app_push_receiver` blocks as defined below.
        """
        return pulumi.get(self, "azure_app_push_receivers")

    @property
    @pulumi.getter(name="azureFunctionReceivers")
    def azure_function_receivers(self) -> Optional[List['outputs.ActionGroupAzureFunctionReceiver']]:
        """
        One or more `azure_function_receiver` blocks as defined below.
        """
        return pulumi.get(self, "azure_function_receivers")

    @property
    @pulumi.getter(name="emailReceivers")
    def email_receivers(self) -> Optional[List['outputs.ActionGroupEmailReceiver']]:
        """
        One or more `email_receiver` blocks as defined below.
        """
        return pulumi.get(self, "email_receivers")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="itsmReceivers")
    def itsm_receivers(self) -> Optional[List['outputs.ActionGroupItsmReceiver']]:
        """
        One or more `itsm_receiver` blocks as defined below.
        """
        return pulumi.get(self, "itsm_receivers")

    @property
    @pulumi.getter(name="logicAppReceivers")
    def logic_app_receivers(self) -> Optional[List['outputs.ActionGroupLogicAppReceiver']]:
        """
        One or more `logic_app_receiver` blocks as defined below.
        """
        return pulumi.get(self, "logic_app_receivers")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Action Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the Action Group instance.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> str:
        """
        The short name of the action group. This will be used in SMS messages.
        """
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter(name="smsReceivers")
    def sms_receivers(self) -> Optional[List['outputs.ActionGroupSmsReceiver']]:
        """
        One or more `sms_receiver` blocks as defined below.
        """
        return pulumi.get(self, "sms_receivers")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="voiceReceivers")
    def voice_receivers(self) -> Optional[List['outputs.ActionGroupVoiceReceiver']]:
        """
        One or more `voice_receiver` blocks as defined below.
        """
        return pulumi.get(self, "voice_receivers")

    @property
    @pulumi.getter(name="webhookReceivers")
    def webhook_receivers(self) -> Optional[List['outputs.ActionGroupWebhookReceiver']]:
        """
        One or more `webhook_receiver` blocks as defined below.
        """
        return pulumi.get(self, "webhook_receivers")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

