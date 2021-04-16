# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DataSourceWindowsPerformanceCounterArgs', 'DataSourceWindowsPerformanceCounter']

@pulumi.input_type
class DataSourceWindowsPerformanceCounterArgs:
    def __init__(__self__, *,
                 counter_name: pulumi.Input[str],
                 instance_name: pulumi.Input[str],
                 interval_seconds: pulumi.Input[int],
                 object_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 workspace_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataSourceWindowsPerformanceCounter resource.
        :param pulumi.Input[str] counter_name: The friendly name of the performance counter.
        :param pulumi.Input[str] instance_name: The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        :param pulumi.Input[int] interval_seconds: The time of sample interval in seconds. Supports values between 10 and 2147483647.
        :param pulumi.Input[str] object_name: The object name of the Log Analytics Windows Performance Counter DataSource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        pulumi.set(__self__, "counter_name", counter_name)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "interval_seconds", interval_seconds)
        pulumi.set(__self__, "object_name", object_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="counterName")
    def counter_name(self) -> pulumi.Input[str]:
        """
        The friendly name of the performance counter.
        """
        return pulumi.get(self, "counter_name")

    @counter_name.setter
    def counter_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "counter_name", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="intervalSeconds")
    def interval_seconds(self) -> pulumi.Input[int]:
        """
        The time of sample interval in seconds. Supports values between 10 and 2147483647.
        """
        return pulumi.get(self, "interval_seconds")

    @interval_seconds.setter
    def interval_seconds(self, value: pulumi.Input[int]):
        pulumi.set(self, "interval_seconds", value)

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> pulumi.Input[str]:
        """
        The object name of the Log Analytics Windows Performance Counter DataSource.
        """
        return pulumi.get(self, "object_name")

    @object_name.setter
    def object_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[str]:
        """
        The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DataSourceWindowsPerformanceCounterState:
    def __init__(__self__, *,
                 counter_name: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 interval_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataSourceWindowsPerformanceCounter resources.
        :param pulumi.Input[str] counter_name: The friendly name of the performance counter.
        :param pulumi.Input[str] instance_name: The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        :param pulumi.Input[int] interval_seconds: The time of sample interval in seconds. Supports values between 10 and 2147483647.
        :param pulumi.Input[str] name: The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] object_name: The object name of the Log Analytics Windows Performance Counter DataSource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        if counter_name is not None:
            pulumi.set(__self__, "counter_name", counter_name)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if interval_seconds is not None:
            pulumi.set(__self__, "interval_seconds", interval_seconds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_name is not None:
            pulumi.set(__self__, "object_name", object_name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if workspace_name is not None:
            pulumi.set(__self__, "workspace_name", workspace_name)

    @property
    @pulumi.getter(name="counterName")
    def counter_name(self) -> Optional[pulumi.Input[str]]:
        """
        The friendly name of the performance counter.
        """
        return pulumi.get(self, "counter_name")

    @counter_name.setter
    def counter_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "counter_name", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="intervalSeconds")
    def interval_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The time of sample interval in seconds. Supports values between 10 and 2147483647.
        """
        return pulumi.get(self, "interval_seconds")

    @interval_seconds.setter
    def interval_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval_seconds", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> Optional[pulumi.Input[str]]:
        """
        The object name of the Log Analytics Windows Performance Counter DataSource.
        """
        return pulumi.get(self, "object_name")

    @object_name.setter
    def object_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_name", value)


class DataSourceWindowsPerformanceCounter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 counter_name: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 interval_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Log Analytics (formally Operational Insights) Windows Performance Counter DataSource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018")
        example_data_source_windows_performance_counter = azure.loganalytics.DataSourceWindowsPerformanceCounter("exampleDataSourceWindowsPerformanceCounter",
            resource_group_name=example_resource_group.name,
            workspace_name=example_analytics_workspace.name,
            object_name="CPU",
            instance_name="*",
            counter_name="CPU",
            interval_seconds=10)
        ```

        ## Import

        Log Analytics Windows Performance Counter DataSources can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/datasources/datasource1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] counter_name: The friendly name of the performance counter.
        :param pulumi.Input[str] instance_name: The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        :param pulumi.Input[int] interval_seconds: The time of sample interval in seconds. Supports values between 10 and 2147483647.
        :param pulumi.Input[str] name: The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] object_name: The object name of the Log Analytics Windows Performance Counter DataSource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourceWindowsPerformanceCounterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Log Analytics (formally Operational Insights) Windows Performance Counter DataSource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018")
        example_data_source_windows_performance_counter = azure.loganalytics.DataSourceWindowsPerformanceCounter("exampleDataSourceWindowsPerformanceCounter",
            resource_group_name=example_resource_group.name,
            workspace_name=example_analytics_workspace.name,
            object_name="CPU",
            instance_name="*",
            counter_name="CPU",
            interval_seconds=10)
        ```

        ## Import

        Log Analytics Windows Performance Counter DataSources can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/datasources/datasource1
        ```

        :param str resource_name: The name of the resource.
        :param DataSourceWindowsPerformanceCounterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceWindowsPerformanceCounterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 counter_name: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 interval_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = DataSourceWindowsPerformanceCounterArgs.__new__(DataSourceWindowsPerformanceCounterArgs)

            if counter_name is None and not opts.urn:
                raise TypeError("Missing required property 'counter_name'")
            __props__.__dict__["counter_name"] = counter_name
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            if interval_seconds is None and not opts.urn:
                raise TypeError("Missing required property 'interval_seconds'")
            __props__.__dict__["interval_seconds"] = interval_seconds
            __props__.__dict__["name"] = name
            if object_name is None and not opts.urn:
                raise TypeError("Missing required property 'object_name'")
            __props__.__dict__["object_name"] = object_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
        super(DataSourceWindowsPerformanceCounter, __self__).__init__(
            'azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            counter_name: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            interval_seconds: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            workspace_name: Optional[pulumi.Input[str]] = None) -> 'DataSourceWindowsPerformanceCounter':
        """
        Get an existing DataSourceWindowsPerformanceCounter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] counter_name: The friendly name of the performance counter.
        :param pulumi.Input[str] instance_name: The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        :param pulumi.Input[int] interval_seconds: The time of sample interval in seconds. Supports values between 10 and 2147483647.
        :param pulumi.Input[str] name: The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] object_name: The object name of the Log Analytics Windows Performance Counter DataSource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourceWindowsPerformanceCounterState.__new__(_DataSourceWindowsPerformanceCounterState)

        __props__.__dict__["counter_name"] = counter_name
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["interval_seconds"] = interval_seconds
        __props__.__dict__["name"] = name
        __props__.__dict__["object_name"] = object_name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["workspace_name"] = workspace_name
        return DataSourceWindowsPerformanceCounter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="counterName")
    def counter_name(self) -> pulumi.Output[str]:
        """
        The friendly name of the performance counter.
        """
        return pulumi.get(self, "counter_name")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="intervalSeconds")
    def interval_seconds(self) -> pulumi.Output[int]:
        """
        The time of sample interval in seconds. Supports values between 10 and 2147483647.
        """
        return pulumi.get(self, "interval_seconds")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> pulumi.Output[str]:
        """
        The object name of the Log Analytics Windows Performance Counter DataSource.
        """
        return pulumi.get(self, "object_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Output[str]:
        """
        The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        return pulumi.get(self, "workspace_name")

