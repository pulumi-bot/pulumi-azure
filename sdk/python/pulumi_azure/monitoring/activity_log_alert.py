# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ActivityLogAlert']


class ActivityLogAlert(pulumi.CustomResource):
    actions: pulumi.Output[Optional[List['outputs.ActivityLogAlertAction']]] = pulumi.output_property("actions")
    """
    One or more `action` blocks as defined below.
    """
    criteria: pulumi.Output['outputs.ActivityLogAlertCriteria'] = pulumi.output_property("criteria")
    """
    A `criteria` block as defined below.
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    The description of this activity log alert.
    """
    enabled: pulumi.Output[Optional[bool]] = pulumi.output_property("enabled")
    """
    Should this Activity Log Alert be enabled? Defaults to `true`.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the activity log alert. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the activity log alert instance.
    """
    scopes: pulumi.Output[List[str]] = pulumi.output_property("scopes")
    """
    The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActivityLogAlertActionArgs']]]]] = None, criteria: Optional[pulumi.Input[pulumi.InputType['ActivityLogAlertCriteriaArgs']]] = None, description: Optional[pulumi.Input[str]] = None, enabled: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an Activity Log Alert within Azure Monitor.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        main_resource_group = azure.core.ResourceGroup("mainResourceGroup", location="West US")
        main_action_group = azure.monitoring.ActionGroup("mainActionGroup",
            resource_group_name=main_resource_group.name,
            short_name="p0action",
            webhook_receivers=[{
                "name": "callmyapi",
                "service_uri": "http://example.com/alert",
            }])
        to_monitor = azure.storage.Account("toMonitor",
            resource_group_name=main_resource_group.name,
            location=main_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        main_activity_log_alert = azure.monitoring.ActivityLogAlert("mainActivityLogAlert",
            resource_group_name=main_resource_group.name,
            scopes=[main_resource_group.id],
            description="This alert will monitor a specific storage account updates.",
            criteria={
                "resource_id": to_monitor.id,
                "operationName": "Microsoft.Storage/storageAccounts/write",
                "category": "Recommendation",
            },
            actions=[{
                "action_group_id": main_action_group.id,
                "webhookProperties": {
                    "from": "source",
                },
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActivityLogAlertActionArgs']]]] actions: One or more `action` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ActivityLogAlertCriteriaArgs']] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] description: The description of this activity log alert.
        :param pulumi.Input[bool] enabled: Should this Activity Log Alert be enabled? Defaults to `true`.
        :param pulumi.Input[str] name: The name of the activity log alert. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the activity log alert instance.
        :param pulumi.Input[List[pulumi.Input[str]]] scopes: The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['actions'] = actions
            if criteria is None:
                raise TypeError("Missing required property 'criteria'")
            __props__['criteria'] = criteria
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if scopes is None:
                raise TypeError("Missing required property 'scopes'")
            __props__['scopes'] = scopes
            __props__['tags'] = tags
        super(ActivityLogAlert, __self__).__init__(
            'azure:monitoring/activityLogAlert:ActivityLogAlert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActivityLogAlertActionArgs']]]]] = None, criteria: Optional[pulumi.Input[pulumi.InputType['ActivityLogAlertCriteriaArgs']]] = None, description: Optional[pulumi.Input[str]] = None, enabled: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None) -> 'ActivityLogAlert':
        """
        Get an existing ActivityLogAlert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActivityLogAlertActionArgs']]]] actions: One or more `action` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ActivityLogAlertCriteriaArgs']] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] description: The description of this activity log alert.
        :param pulumi.Input[bool] enabled: Should this Activity Log Alert be enabled? Defaults to `true`.
        :param pulumi.Input[str] name: The name of the activity log alert. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the activity log alert instance.
        :param pulumi.Input[List[pulumi.Input[str]]] scopes: The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions"] = actions
        __props__["criteria"] = criteria
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["scopes"] = scopes
        __props__["tags"] = tags
        return ActivityLogAlert(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

