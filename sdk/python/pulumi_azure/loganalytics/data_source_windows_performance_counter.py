# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DataSourceWindowsPerformanceCounter(pulumi.CustomResource):
    counter_name: pulumi.Output[str]
    """
    The friendly name of the performance counter.
    """
    instance_name: pulumi.Output[str]
    """
    The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
    """
    interval_seconds: pulumi.Output[float]
    """
    The time of sample interval in seconds. Supports values between 10 and 2147483647.
    """
    name: pulumi.Output[str]
    """
    The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
    """
    object_name: pulumi.Output[str]
    """
    The object name of the Log Analytics Windows Performance Counter DataSource.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
    """
    workspace_name: pulumi.Output[str]
    """
    The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
    """
    def __init__(__self__, resource_name, opts=None, counter_name=None, instance_name=None, interval_seconds=None, name=None, object_name=None, resource_group_name=None, workspace_name=None, __props__=None, __name__=None, __opts__=None):
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


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] counter_name: The friendly name of the performance counter.
        :param pulumi.Input[str] instance_name: The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        :param pulumi.Input[float] interval_seconds: The time of sample interval in seconds. Supports values between 10 and 2147483647.
        :param pulumi.Input[str] name: The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] object_name: The object name of the Log Analytics Windows Performance Counter DataSource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
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

            if counter_name is None:
                raise TypeError("Missing required property 'counter_name'")
            __props__['counter_name'] = counter_name
            if instance_name is None:
                raise TypeError("Missing required property 'instance_name'")
            __props__['instance_name'] = instance_name
            if interval_seconds is None:
                raise TypeError("Missing required property 'interval_seconds'")
            __props__['interval_seconds'] = interval_seconds
            __props__['name'] = name
            if object_name is None:
                raise TypeError("Missing required property 'object_name'")
            __props__['object_name'] = object_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
        super(DataSourceWindowsPerformanceCounter, __self__).__init__(
            'azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, counter_name=None, instance_name=None, interval_seconds=None, name=None, object_name=None, resource_group_name=None, workspace_name=None):
        """
        Get an existing DataSourceWindowsPerformanceCounter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] counter_name: The friendly name of the performance counter.
        :param pulumi.Input[str] instance_name: The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
        :param pulumi.Input[float] interval_seconds: The time of sample interval in seconds. Supports values between 10 and 2147483647.
        :param pulumi.Input[str] name: The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] object_name: The object name of the Log Analytics Windows Performance Counter DataSource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        :param pulumi.Input[str] workspace_name: The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["counter_name"] = counter_name
        __props__["instance_name"] = instance_name
        __props__["interval_seconds"] = interval_seconds
        __props__["name"] = name
        __props__["object_name"] = object_name
        __props__["resource_group_name"] = resource_group_name
        __props__["workspace_name"] = workspace_name
        return DataSourceWindowsPerformanceCounter(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

