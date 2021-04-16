# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DataExportRuleArgs', 'DataExportRule']

@pulumi.input_type
class DataExportRuleArgs:
    def __init__(__self__, *,
                 destination_resource_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 table_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 workspace_resource_id: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataExportRule resource.
        :param pulumi.Input[str] destination_resource_id: The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] table_names: A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        :param pulumi.Input[str] workspace_resource_id: The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[bool] enabled: Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        pulumi.set(__self__, "destination_resource_id", destination_resource_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "table_names", table_names)
        pulumi.set(__self__, "workspace_resource_id", workspace_resource_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="destinationResourceId")
    def destination_resource_id(self) -> pulumi.Input[str]:
        """
        The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        """
        return pulumi.get(self, "destination_resource_id")

    @destination_resource_id.setter
    def destination_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_resource_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="tableNames")
    def table_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        """
        return pulumi.get(self, "table_names")

    @table_names.setter
    def table_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "table_names", value)

    @property
    @pulumi.getter(name="workspaceResourceId")
    def workspace_resource_id(self) -> pulumi.Input[str]:
        """
        The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "workspace_resource_id")

    @workspace_resource_id.setter
    def workspace_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace_resource_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DataExportRuleState:
    def __init__(__self__, *,
                 destination_resource_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 export_rule_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 table_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workspace_resource_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataExportRule resources.
        :param pulumi.Input[str] destination_resource_id: The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        :param pulumi.Input[bool] enabled: Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        :param pulumi.Input[str] export_rule_id: The ID of the created Data Export Rule.
        :param pulumi.Input[str] name: The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] table_names: A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        :param pulumi.Input[str] workspace_resource_id: The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        if destination_resource_id is not None:
            pulumi.set(__self__, "destination_resource_id", destination_resource_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if export_rule_id is not None:
            pulumi.set(__self__, "export_rule_id", export_rule_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if table_names is not None:
            pulumi.set(__self__, "table_names", table_names)
        if workspace_resource_id is not None:
            pulumi.set(__self__, "workspace_resource_id", workspace_resource_id)

    @property
    @pulumi.getter(name="destinationResourceId")
    def destination_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        """
        return pulumi.get(self, "destination_resource_id")

    @destination_resource_id.setter
    def destination_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_resource_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="exportRuleId")
    def export_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the created Data Export Rule.
        """
        return pulumi.get(self, "export_rule_id")

    @export_rule_id.setter
    def export_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "export_rule_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="tableNames")
    def table_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        """
        return pulumi.get(self, "table_names")

    @table_names.setter
    def table_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "table_names", value)

    @property
    @pulumi.getter(name="workspaceResourceId")
    def workspace_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "workspace_resource_id")

    @workspace_resource_id.setter
    def workspace_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_resource_id", value)


class DataExportRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_resource_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 table_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workspace_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Log Analytics Data Export Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018",
            retention_in_days=30)
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_data_export_rule = azure.loganalytics.DataExportRule("exampleDataExportRule",
            resource_group_name=example_resource_group.name,
            workspace_resource_id=example_analytics_workspace.id,
            destination_resource_id=example_account.id,
            table_names=["Heartbeat"])
        ```

        ## Import

        Log Analytics Data Export Rule can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:loganalytics/dataExportRule:DataExportRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataExports/dataExport1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_resource_id: The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        :param pulumi.Input[bool] enabled: Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] table_names: A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        :param pulumi.Input[str] workspace_resource_id: The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataExportRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Log Analytics Data Export Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018",
            retention_in_days=30)
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_data_export_rule = azure.loganalytics.DataExportRule("exampleDataExportRule",
            resource_group_name=example_resource_group.name,
            workspace_resource_id=example_analytics_workspace.id,
            destination_resource_id=example_account.id,
            table_names=["Heartbeat"])
        ```

        ## Import

        Log Analytics Data Export Rule can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:loganalytics/dataExportRule:DataExportRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataExports/dataExport1
        ```

        :param str resource_name: The name of the resource.
        :param DataExportRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataExportRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_resource_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 table_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workspace_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataExportRuleArgs.__new__(DataExportRuleArgs)

            if destination_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'destination_resource_id'")
            __props__.__dict__["destination_resource_id"] = destination_resource_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if table_names is None and not opts.urn:
                raise TypeError("Missing required property 'table_names'")
            __props__.__dict__["table_names"] = table_names
            if workspace_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_resource_id'")
            __props__.__dict__["workspace_resource_id"] = workspace_resource_id
            __props__.__dict__["export_rule_id"] = None
        super(DataExportRule, __self__).__init__(
            'azure:loganalytics/dataExportRule:DataExportRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_resource_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            export_rule_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            table_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            workspace_resource_id: Optional[pulumi.Input[str]] = None) -> 'DataExportRule':
        """
        Get an existing DataExportRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_resource_id: The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        :param pulumi.Input[bool] enabled: Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        :param pulumi.Input[str] export_rule_id: The ID of the created Data Export Rule.
        :param pulumi.Input[str] name: The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] table_names: A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        :param pulumi.Input[str] workspace_resource_id: The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataExportRuleState.__new__(_DataExportRuleState)

        __props__.__dict__["destination_resource_id"] = destination_resource_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["export_rule_id"] = export_rule_id
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["table_names"] = table_names
        __props__.__dict__["workspace_resource_id"] = workspace_resource_id
        return DataExportRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationResourceId")
    def destination_resource_id(self) -> pulumi.Output[str]:
        """
        The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
        """
        return pulumi.get(self, "destination_resource_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is this Log Analytics Data Export Rule enabled? Possible values include `true` or `false`. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="exportRuleId")
    def export_rule_id(self) -> pulumi.Output[str]:
        """
        The ID of the created Data Export Rule.
        """
        return pulumi.get(self, "export_rule_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="tableNames")
    def table_names(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
        """
        return pulumi.get(self, "table_names")

    @property
    @pulumi.getter(name="workspaceResourceId")
    def workspace_resource_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
        """
        return pulumi.get(self, "workspace_resource_id")

