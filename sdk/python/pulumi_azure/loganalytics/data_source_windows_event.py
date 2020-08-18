# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['DataSourceWindowsEvent']


class DataSourceWindowsEvent(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_log_name: Optional[pulumi.Input[str]] = None,
                 event_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Log Analytics Windows Event DataSource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018")
        example_data_source_windows_event = azure.loganalytics.DataSourceWindowsEvent("exampleDataSourceWindowsEvent",
            resource_group_name=example_resource_group.name,
            workspace_name=example_analytics_workspace.name,
            event_log_name="Application",
            event_types=["error"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_log_name: Specifies the name of the Windows Event Log to collect events from.
        :param pulumi.Input[List[pulumi.Input[str]]] event_types: Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
        :param pulumi.Input[str] name: The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
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

            if event_log_name is None:
                raise TypeError("Missing required property 'event_log_name'")
            __props__['event_log_name'] = event_log_name
            if event_types is None:
                raise TypeError("Missing required property 'event_types'")
            __props__['event_types'] = event_types
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
        super(DataSourceWindowsEvent, __self__).__init__(
            'azure:loganalytics/dataSourceWindowsEvent:DataSourceWindowsEvent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            event_log_name: Optional[pulumi.Input[str]] = None,
            event_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            workspace_name: Optional[pulumi.Input[str]] = None) -> 'DataSourceWindowsEvent':
        """
        Get an existing DataSourceWindowsEvent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_log_name: Specifies the name of the Windows Event Log to collect events from.
        :param pulumi.Input[List[pulumi.Input[str]]] event_types: Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
        :param pulumi.Input[str] name: The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["event_log_name"] = event_log_name
        __props__["event_types"] = event_types
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["workspace_name"] = workspace_name
        return DataSourceWindowsEvent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventLogName")
    def event_log_name(self) -> str:
        """
        Specifies the name of the Windows Event Log to collect events from.
        """
        return pulumi.get(self, "event_log_name")

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> List[str]:
        """
        Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
        """
        return pulumi.get(self, "event_types")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> str:
        """
        The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
        """
        return pulumi.get(self, "workspace_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

