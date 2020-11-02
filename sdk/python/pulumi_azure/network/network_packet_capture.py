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

__all__ = ['NetworkPacketCapture']


class NetworkPacketCapture(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkPacketCaptureFilterArgs']]]]] = None,
                 maximum_bytes_per_packet: Optional[pulumi.Input[int]] = None,
                 maximum_bytes_per_session: Optional[pulumi.Input[int]] = None,
                 maximum_capture_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_watcher_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_location: Optional[pulumi.Input[pulumi.InputType['NetworkPacketCaptureStorageLocationArgs']]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures Network Packet Capturing against a Virtual Machine using a Network Watcher.

        ## Import

        Packet Captures can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkPacketCaptureFilterArgs']]]] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['NetworkPacketCaptureStorageLocationArgs']] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
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

            __props__['filters'] = filters
            __props__['maximum_bytes_per_packet'] = maximum_bytes_per_packet
            __props__['maximum_bytes_per_session'] = maximum_bytes_per_session
            __props__['maximum_capture_duration'] = maximum_capture_duration
            __props__['name'] = name
            if network_watcher_name is None:
                raise TypeError("Missing required property 'network_watcher_name'")
            __props__['network_watcher_name'] = network_watcher_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_location is None:
                raise TypeError("Missing required property 'storage_location'")
            __props__['storage_location'] = storage_location
            if target_resource_id is None:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__['target_resource_id'] = target_resource_id
        super(NetworkPacketCapture, __self__).__init__(
            'azure:network/networkPacketCapture:NetworkPacketCapture',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkPacketCaptureFilterArgs']]]]] = None,
            maximum_bytes_per_packet: Optional[pulumi.Input[int]] = None,
            maximum_bytes_per_session: Optional[pulumi.Input[int]] = None,
            maximum_capture_duration: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_watcher_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_location: Optional[pulumi.Input[pulumi.InputType['NetworkPacketCaptureStorageLocationArgs']]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None) -> 'NetworkPacketCapture':
        """
        Get an existing NetworkPacketCapture resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkPacketCaptureFilterArgs']]]] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['NetworkPacketCaptureStorageLocationArgs']] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["filters"] = filters
        __props__["maximum_bytes_per_packet"] = maximum_bytes_per_packet
        __props__["maximum_bytes_per_session"] = maximum_bytes_per_session
        __props__["maximum_capture_duration"] = maximum_capture_duration
        __props__["name"] = name
        __props__["network_watcher_name"] = network_watcher_name
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_location"] = storage_location
        __props__["target_resource_id"] = target_resource_id
        return NetworkPacketCapture(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkPacketCaptureFilter']]]:
        """
        One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="maximumBytesPerPacket")
    def maximum_bytes_per_packet(self) -> pulumi.Output[Optional[int]]:
        """
        The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maximum_bytes_per_packet")

    @property
    @pulumi.getter(name="maximumBytesPerSession")
    def maximum_bytes_per_session(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maximum_bytes_per_session")

    @property
    @pulumi.getter(name="maximumCaptureDuration")
    def maximum_capture_duration(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maximum_capture_duration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkWatcherName")
    def network_watcher_name(self) -> pulumi.Output[str]:
        """
        The name of the Network Watcher. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_watcher_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="storageLocation")
    def storage_location(self) -> pulumi.Output['outputs.NetworkPacketCaptureStorageLocation']:
        """
        A `storage_location` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_location")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "target_resource_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

