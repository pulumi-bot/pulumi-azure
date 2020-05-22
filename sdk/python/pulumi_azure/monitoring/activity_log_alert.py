# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ActivityLogAlert(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    One or more `action` blocks as defined below.

      * `actionGroupId` (`str`) - The ID of the Action Group can be sourced from the `monitoring.ActionGroup` resource.
      * `webhookProperties` (`dict`) - The map of custom string properties to include with the post operation. These data are appended to the webhook payload.
    """
    criteria: pulumi.Output[dict]
    """
    A `criteria` block as defined below.

      * `caller` (`str`) - The email address or Azure Active Directory identifier of the user who performed the operation.
      * `category` (`str`) - The category of the operation. Possible values are `Administrative`, `Autoscale`, `Policy`, `Recommendation`, `ResourceHealth`, `Security` and `ServiceHealth`.
      * `level` (`str`) - The severity level of the event. Possible values are `Verbose`, `Informational`, `Warning`, `Error`, and `Critical`.
      * `operationName` (`str`) - The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: `<resourceProvider>/<resourceType>/<operation>`.
      * `resourceGroup` (`str`) - The name of resource group monitored by the activity log alert.
      * `resource_id` (`str`) - The specific resource monitored by the activity log alert. It should be within one of the `scopes`.
      * `resourceProvider` (`str`) - The name of the resource provider monitored by the activity log alert.
      * `resourceType` (`str`) - The resource type monitored by the activity log alert.
      * `status` (`str`) - The status of the event. For example, `Started`, `Failed`, or `Succeeded`.
      * `subStatus` (`str`) - The sub status of the event.
    """
    description: pulumi.Output[str]
    """
    The description of this activity log alert.
    """
    enabled: pulumi.Output[bool]
    """
    Should this Activity Log Alert be enabled? Defaults to `true`.
    """
    name: pulumi.Output[str]
    """
    The name of the activity log alert. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the activity log alert instance.
    """
    scopes: pulumi.Output[list]
    """
    The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, actions=None, criteria=None, description=None, enabled=None, name=None, resource_group_name=None, scopes=None, tags=None, __props__=None, __name__=None, __opts__=None):
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
            webhook_receiver=[{
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
            action=[{
                "actionGroupId": main_action_group.id,
                "webhookProperties": {
                    "from": "source",
                },
            }])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: One or more `action` blocks as defined below.
        :param pulumi.Input[dict] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] description: The description of this activity log alert.
        :param pulumi.Input[bool] enabled: Should this Activity Log Alert be enabled? Defaults to `true`.
        :param pulumi.Input[str] name: The name of the activity log alert. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the activity log alert instance.
        :param pulumi.Input[list] scopes: The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **actions** object supports the following:

          * `actionGroupId` (`pulumi.Input[str]`) - The ID of the Action Group can be sourced from the `monitoring.ActionGroup` resource.
          * `webhookProperties` (`pulumi.Input[dict]`) - The map of custom string properties to include with the post operation. These data are appended to the webhook payload.

        The **criteria** object supports the following:

          * `caller` (`pulumi.Input[str]`) - The email address or Azure Active Directory identifier of the user who performed the operation.
          * `category` (`pulumi.Input[str]`) - The category of the operation. Possible values are `Administrative`, `Autoscale`, `Policy`, `Recommendation`, `ResourceHealth`, `Security` and `ServiceHealth`.
          * `level` (`pulumi.Input[str]`) - The severity level of the event. Possible values are `Verbose`, `Informational`, `Warning`, `Error`, and `Critical`.
          * `operationName` (`pulumi.Input[str]`) - The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: `<resourceProvider>/<resourceType>/<operation>`.
          * `resourceGroup` (`pulumi.Input[str]`) - The name of resource group monitored by the activity log alert.
          * `resource_id` (`pulumi.Input[str]`) - The specific resource monitored by the activity log alert. It should be within one of the `scopes`.
          * `resourceProvider` (`pulumi.Input[str]`) - The name of the resource provider monitored by the activity log alert.
          * `resourceType` (`pulumi.Input[str]`) - The resource type monitored by the activity log alert.
          * `status` (`pulumi.Input[str]`) - The status of the event. For example, `Started`, `Failed`, or `Succeeded`.
          * `subStatus` (`pulumi.Input[str]`) - The sub status of the event.
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
    def get(resource_name, id, opts=None, actions=None, criteria=None, description=None, enabled=None, name=None, resource_group_name=None, scopes=None, tags=None):
        """
        Get an existing ActivityLogAlert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: One or more `action` blocks as defined below.
        :param pulumi.Input[dict] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] description: The description of this activity log alert.
        :param pulumi.Input[bool] enabled: Should this Activity Log Alert be enabled? Defaults to `true`.
        :param pulumi.Input[str] name: The name of the activity log alert. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the activity log alert instance.
        :param pulumi.Input[list] scopes: The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **actions** object supports the following:

          * `actionGroupId` (`pulumi.Input[str]`) - The ID of the Action Group can be sourced from the `monitoring.ActionGroup` resource.
          * `webhookProperties` (`pulumi.Input[dict]`) - The map of custom string properties to include with the post operation. These data are appended to the webhook payload.

        The **criteria** object supports the following:

          * `caller` (`pulumi.Input[str]`) - The email address or Azure Active Directory identifier of the user who performed the operation.
          * `category` (`pulumi.Input[str]`) - The category of the operation. Possible values are `Administrative`, `Autoscale`, `Policy`, `Recommendation`, `ResourceHealth`, `Security` and `ServiceHealth`.
          * `level` (`pulumi.Input[str]`) - The severity level of the event. Possible values are `Verbose`, `Informational`, `Warning`, `Error`, and `Critical`.
          * `operationName` (`pulumi.Input[str]`) - The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: `<resourceProvider>/<resourceType>/<operation>`.
          * `resourceGroup` (`pulumi.Input[str]`) - The name of resource group monitored by the activity log alert.
          * `resource_id` (`pulumi.Input[str]`) - The specific resource monitored by the activity log alert. It should be within one of the `scopes`.
          * `resourceProvider` (`pulumi.Input[str]`) - The name of the resource provider monitored by the activity log alert.
          * `resourceType` (`pulumi.Input[str]`) - The resource type monitored by the activity log alert.
          * `status` (`pulumi.Input[str]`) - The status of the event. For example, `Started`, `Failed`, or `Succeeded`.
          * `subStatus` (`pulumi.Input[str]`) - The sub status of the event.
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

