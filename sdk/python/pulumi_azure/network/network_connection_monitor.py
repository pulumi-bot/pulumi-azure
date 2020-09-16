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

__all__ = ['NetworkConnectionMonitor']


class NetworkConnectionMonitor(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_start: Optional[pulumi.Input[bool]] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']]] = None,
                 interval_in_seconds: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_watcher_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Network Connection Monitor.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_network_watcher = azure.network.NetworkWatcher("exampleNetworkWatcher",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        src_virtual_machine = azure.compute.get_virtual_machine(name="example-vm",
            resource_group_name=example_resource_group.name)
        src_extension = azure.compute.Extension("srcExtension",
            virtual_machine_id=src_virtual_machine.id,
            publisher="Microsoft.Azure.NetworkWatcher",
            type="NetworkWatcherAgentLinux",
            type_handler_version="1.4",
            auto_upgrade_minor_version=True)
        example_network_connection_monitor = azure.network.NetworkConnectionMonitor("exampleNetworkConnectionMonitor",
            network_watcher_name=example_network_watcher.name,
            resource_group_name=example_resource_group.name,
            location=example_network_watcher.location,
            auto_start=False,
            interval_in_seconds=30,
            source=azure.network.NetworkConnectionMonitorSourceArgs(
                virtual_machine_id=src_virtual_machine.id,
                port=20020,
            ),
            destination=azure.network.NetworkConnectionMonitorDestinationArgs(
                address="mycompany.io",
                port=443,
            ),
            tags={
                "foo": "bar",
            },
            opts=ResourceOptions(depends_on=[src_extension]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_start: Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']] destination: A `destination` block as defined below.
        :param pulumi.Input[int] interval_in_seconds: Monitoring interval in seconds.
        :param pulumi.Input[str] location: The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[str] name: The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']] source: A `source` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Network Connection Monitor.
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

            __props__['auto_start'] = auto_start
            if destination is None:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['interval_in_seconds'] = interval_in_seconds
            __props__['location'] = location
            __props__['name'] = name
            if network_watcher_name is None:
                raise TypeError("Missing required property 'network_watcher_name'")
            __props__['network_watcher_name'] = network_watcher_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source is None:
                raise TypeError("Missing required property 'source'")
            __props__['source'] = source
            __props__['tags'] = tags
        super(NetworkConnectionMonitor, __self__).__init__(
            'azure:network/networkConnectionMonitor:NetworkConnectionMonitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_start: Optional[pulumi.Input[bool]] = None,
            destination: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']]] = None,
            interval_in_seconds: Optional[pulumi.Input[int]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_watcher_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'NetworkConnectionMonitor':
        """
        Get an existing NetworkConnectionMonitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_start: Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']] destination: A `destination` block as defined below.
        :param pulumi.Input[int] interval_in_seconds: Monitoring interval in seconds.
        :param pulumi.Input[str] location: The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[str] name: The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']] source: A `source` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Network Connection Monitor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_start"] = auto_start
        __props__["destination"] = destination
        __props__["interval_in_seconds"] = interval_in_seconds
        __props__["location"] = location
        __props__["name"] = name
        __props__["network_watcher_name"] = network_watcher_name
        __props__["resource_group_name"] = resource_group_name
        __props__["source"] = source
        __props__["tags"] = tags
        return NetworkConnectionMonitor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoStart")
    def auto_start(self) -> pulumi.Output[Optional[bool]]:
        """
        Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
        """
        return pulumi.get(self, "auto_start")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output['outputs.NetworkConnectionMonitorDestination']:
        """
        A `destination` block as defined below.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="intervalInSeconds")
    def interval_in_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        Monitoring interval in seconds.
        """
        return pulumi.get(self, "interval_in_seconds")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkWatcherName")
    def network_watcher_name(self) -> pulumi.Output[str]:
        """
        The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
        """
        return pulumi.get(self, "network_watcher_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output['outputs.NetworkConnectionMonitorSource']:
        """
        A `source` block as defined below.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Network Connection Monitor.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

