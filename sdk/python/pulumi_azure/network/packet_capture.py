# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PacketCaptureArgs', 'PacketCapture']

@pulumi.input_type
class PacketCaptureArgs:
    def __init__(__self__, *,
                 network_watcher_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 storage_location: pulumi.Input['PacketCaptureStorageLocationArgs'],
                 target_resource_id: pulumi.Input[str],
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['PacketCaptureFilterArgs']]]] = None,
                 maximum_bytes_per_packet: Optional[pulumi.Input[int]] = None,
                 maximum_bytes_per_session: Optional[pulumi.Input[int]] = None,
                 maximum_capture_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PacketCapture resource.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input['PacketCaptureStorageLocationArgs'] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['PacketCaptureFilterArgs']]] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Packet Capture. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "network_watcher_name", network_watcher_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "storage_location", storage_location)
        pulumi.set(__self__, "target_resource_id", target_resource_id)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if maximum_bytes_per_packet is not None:
            pulumi.set(__self__, "maximum_bytes_per_packet", maximum_bytes_per_packet)
        if maximum_bytes_per_session is not None:
            pulumi.set(__self__, "maximum_bytes_per_session", maximum_bytes_per_session)
        if maximum_capture_duration is not None:
            pulumi.set(__self__, "maximum_capture_duration", maximum_capture_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="networkWatcherName")
    def network_watcher_name(self) -> pulumi.Input[str]:
        """
        The name of the Network Watcher. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_watcher_name")

    @network_watcher_name.setter
    def network_watcher_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_watcher_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="storageLocation")
    def storage_location(self) -> pulumi.Input['PacketCaptureStorageLocationArgs']:
        """
        A `storage_location` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_location")

    @storage_location.setter
    def storage_location(self, value: pulumi.Input['PacketCaptureStorageLocationArgs']):
        pulumi.set(self, "storage_location", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_id", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PacketCaptureFilterArgs']]]]:
        """
        One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PacketCaptureFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter(name="maximumBytesPerPacket")
    def maximum_bytes_per_packet(self) -> Optional[pulumi.Input[int]]:
        """
        The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maximum_bytes_per_packet")

    @maximum_bytes_per_packet.setter
    def maximum_bytes_per_packet(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_bytes_per_packet", value)

    @property
    @pulumi.getter(name="maximumBytesPerSession")
    def maximum_bytes_per_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maximum_bytes_per_session")

    @maximum_bytes_per_session.setter
    def maximum_bytes_per_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_bytes_per_session", value)

    @property
    @pulumi.getter(name="maximumCaptureDuration")
    def maximum_capture_duration(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maximum_capture_duration")

    @maximum_capture_duration.setter
    def maximum_capture_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_capture_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to use for this Packet Capture. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class PacketCapture(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PacketCaptureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configures Packet Capturing against a Virtual Machine using a Network Watcher.

        > **NOTE:** This resource has been deprecated in favour of the `network.NetworkConnectionMonitor` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, and information on migrating across can be found in this guide.

        ## Import

        Packet Captures can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/packetCapture:PacketCapture capture1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
        ```

        :param str resource_name: The name of the resource.
        :param PacketCaptureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PacketCaptureFilterArgs']]]]] = None,
                 maximum_bytes_per_packet: Optional[pulumi.Input[int]] = None,
                 maximum_bytes_per_session: Optional[pulumi.Input[int]] = None,
                 maximum_capture_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_watcher_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_location: Optional[pulumi.Input[pulumi.InputType['PacketCaptureStorageLocationArgs']]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures Packet Capturing against a Virtual Machine using a Network Watcher.

        > **NOTE:** This resource has been deprecated in favour of the `network.NetworkConnectionMonitor` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, and information on migrating across can be found in this guide.

        ## Import

        Packet Captures can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/packetCapture:PacketCapture capture1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PacketCaptureFilterArgs']]]] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Packet Capture. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PacketCaptureStorageLocationArgs']] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PacketCaptureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PacketCaptureFilterArgs']]]]] = None,
                 maximum_bytes_per_packet: Optional[pulumi.Input[int]] = None,
                 maximum_bytes_per_session: Optional[pulumi.Input[int]] = None,
                 maximum_capture_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_watcher_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_location: Optional[pulumi.Input[pulumi.InputType['PacketCaptureStorageLocationArgs']]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if network_watcher_name is None and not opts.urn:
                raise TypeError("Missing required property 'network_watcher_name'")
            __props__['network_watcher_name'] = network_watcher_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_location is None and not opts.urn:
                raise TypeError("Missing required property 'storage_location'")
            __props__['storage_location'] = storage_location
            if target_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__['target_resource_id'] = target_resource_id
        super(PacketCapture, __self__).__init__(
            'azure:network/packetCapture:PacketCapture',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PacketCaptureFilterArgs']]]]] = None,
            maximum_bytes_per_packet: Optional[pulumi.Input[int]] = None,
            maximum_bytes_per_session: Optional[pulumi.Input[int]] = None,
            maximum_capture_duration: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_watcher_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_location: Optional[pulumi.Input[pulumi.InputType['PacketCaptureStorageLocationArgs']]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None) -> 'PacketCapture':
        """
        Get an existing PacketCapture resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PacketCaptureFilterArgs']]]] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[int] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Packet Capture. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PacketCaptureStorageLocationArgs']] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
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
        return PacketCapture(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence['outputs.PacketCaptureFilter']]]:
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
        The name to use for this Packet Capture. Changing this forces a new resource to be created.
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
    def storage_location(self) -> pulumi.Output['outputs.PacketCaptureStorageLocation']:
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

