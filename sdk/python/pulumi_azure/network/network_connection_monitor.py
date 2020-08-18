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

__all__ = ['NetworkConnectionMonitor']


class NetworkConnectionMonitor(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_start: Optional[pulumi.Input[bool]] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']]] = None,
                 interval_in_seconds: Optional[pulumi.Input[float]] = None,
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
        Configures a Network Connection Monitor to monitor communication between a Virtual Machine and an endpoint using a Network Watcher.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_start: Specifies whether the connection monitor will start automatically once created. Defaults to `true`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']] destination: A `destination` block as defined below.
        :param pulumi.Input[float] interval_in_seconds: Monitoring interval in seconds. Defaults to `60`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Network Connection Monitor. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Connection Monitor. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']] source: A `source` block as defined below.
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
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_start: Optional[pulumi.Input[bool]] = None,
            destination: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']]] = None,
            interval_in_seconds: Optional[pulumi.Input[float]] = None,
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
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_start: Specifies whether the connection monitor will start automatically once created. Defaults to `true`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']] destination: A `destination` block as defined below.
        :param pulumi.Input[float] interval_in_seconds: Monitoring interval in seconds. Defaults to `60`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Network Connection Monitor. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Connection Monitor. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']] source: A `source` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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
    def auto_start(self) -> Optional[bool]:
        """
        Specifies whether the connection monitor will start automatically once created. Defaults to `true`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "auto_start")

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.NetworkConnectionMonitorDestination':
        """
        A `destination` block as defined below.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="intervalInSeconds")
    def interval_in_seconds(self) -> Optional[float]:
        """
        Monitoring interval in seconds. Defaults to `60`.
        """
        return pulumi.get(self, "interval_in_seconds")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Network Connection Monitor. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkWatcherName")
    def network_watcher_name(self) -> str:
        """
        The name of the Network Watcher. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_watcher_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the Connection Monitor. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def source(self) -> 'outputs.NetworkConnectionMonitorSource':
        """
        A `source` block as defined below.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

