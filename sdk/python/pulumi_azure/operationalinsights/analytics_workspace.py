# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AnalyticsWorkspace']


class AnalyticsWorkspace(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_quota_gb: Optional[pulumi.Input[float]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Log Analytics (formally Operational Insights) Workspace.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="East US")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018",
            retention_in_days=30)
        ```

        ## Import

        Log Analytics Workspaces can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] daily_quota_gb: The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] retention_in_days: The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
        :param pulumi.Input[str] sku: Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
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

            __props__['daily_quota_gb'] = daily_quota_gb
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retention_in_days'] = retention_in_days
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['portal_url'] = None
            __props__['primary_shared_key'] = None
            __props__['secondary_shared_key'] = None
            __props__['workspace_id'] = None
        super(AnalyticsWorkspace, __self__).__init__(
            'azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            daily_quota_gb: Optional[pulumi.Input[float]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            portal_url: Optional[pulumi.Input[str]] = None,
            primary_shared_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            retention_in_days: Optional[pulumi.Input[int]] = None,
            secondary_shared_key: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'AnalyticsWorkspace':
        """
        Get an existing AnalyticsWorkspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] daily_quota_gb: The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_shared_key: The Primary shared key for the Log Analytics Workspace.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] retention_in_days: The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
        :param pulumi.Input[str] secondary_shared_key: The Secondary shared key for the Log Analytics Workspace.
        :param pulumi.Input[str] sku: Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workspace_id: The Workspace (or Customer) ID for the Log Analytics Workspace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["daily_quota_gb"] = daily_quota_gb
        __props__["location"] = location
        __props__["name"] = name
        __props__["portal_url"] = portal_url
        __props__["primary_shared_key"] = primary_shared_key
        __props__["resource_group_name"] = resource_group_name
        __props__["retention_in_days"] = retention_in_days
        __props__["secondary_shared_key"] = secondary_shared_key
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["workspace_id"] = workspace_id
        return AnalyticsWorkspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dailyQuotaGb")
    def daily_quota_gb(self) -> pulumi.Output[Optional[float]]:
        """
        The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
        """
        return pulumi.get(self, "daily_quota_gb")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalUrl")
    def portal_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portal_url")

    @property
    @pulumi.getter(name="primarySharedKey")
    def primary_shared_key(self) -> pulumi.Output[str]:
        """
        The Primary shared key for the Log Analytics Workspace.
        """
        return pulumi.get(self, "primary_shared_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> pulumi.Output[int]:
        """
        The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
        """
        return pulumi.get(self, "retention_in_days")

    @property
    @pulumi.getter(name="secondarySharedKey")
    def secondary_shared_key(self) -> pulumi.Output[str]:
        """
        The Secondary shared key for the Log Analytics Workspace.
        """
        return pulumi.get(self, "secondary_shared_key")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        """
        The Workspace (or Customer) ID for the Log Analytics Workspace.
        """
        return pulumi.get(self, "workspace_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

