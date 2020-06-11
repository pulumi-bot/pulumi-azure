# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class NetworkPacketCapture(pulumi.CustomResource):
    filters: pulumi.Output[list]
    """
    One or more `filter` blocks as defined below. Changing this forces a new resource to be created.

      * `localIpAddress` (`str`) - The local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
      * `localPort` (`str`) - The local port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
      * `protocol` (`str`) - The Protocol to be filtered on. Possible values include `Any`, `TCP` and `UDP`. Changing this forces a new resource to be created.
      * `remoteIpAddress` (`str`) - The remote IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported.. Changing this forces a new resource to be created.
      * `remotePort` (`str`) - The remote port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
    """
    maximum_bytes_per_packet: pulumi.Output[float]
    """
    The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
    """
    maximum_bytes_per_session: pulumi.Output[float]
    """
    Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
    """
    maximum_capture_duration: pulumi.Output[float]
    """
    The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
    """
    network_watcher_name: pulumi.Output[str]
    """
    The name of the Network Watcher. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
    """
    storage_location: pulumi.Output[dict]
    """
    A `storage_location` block as defined below. Changing this forces a new resource to be created.

      * `filePath` (`str`) - A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with `/var/captures`.
      * `storage_account_id` (`str`) - The ID of the storage account to save the packet capture session
      * `storagePath` (`str`) - The URI of the storage path to save the packet capture.
    """
    target_resource_id: pulumi.Output[str]
    """
    The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, filters=None, maximum_bytes_per_packet=None, maximum_bytes_per_session=None, maximum_capture_duration=None, name=None, network_watcher_name=None, resource_group_name=None, storage_location=None, target_resource_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Configures Network Packet Capturing against a Virtual Machine using a Network Watcher.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_network_watcher = azure.network.NetworkWatcher("exampleNetworkWatcher",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.2.0/24")
        example_network_interface = azure.network.NetworkInterface("exampleNetworkInterface",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configuration=[{
                "name": "testconfiguration1",
                "subnet_id": example_subnet.id,
                "privateIpAddressAllocation": "Dynamic",
            }])
        example_virtual_machine = azure.compute.VirtualMachine("exampleVirtualMachine",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            network_interface_ids=[example_network_interface.id],
            vm_size="Standard_F2",
            storage_image_reference={
                "publisher": "Canonical",
                "offer": "UbuntuServer",
                "sku": "16.04-LTS",
                "version": "latest",
            },
            storage_os_disk={
                "name": "osdisk",
                "caching": "ReadWrite",
                "create_option": "FromImage",
                "managedDiskType": "Standard_LRS",
            },
            os_profile={
                "computer_name": "pctest-vm",
                "admin_username": "testadmin",
                "admin_password": "Password1234!",
            },
            os_profile_linux_config={
                "disable_password_authentication": False,
            })
        example_extension = azure.compute.Extension("exampleExtension",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            virtual_machine_name=example_virtual_machine.name,
            publisher="Microsoft.Azure.NetworkWatcher",
            type="NetworkWatcherAgentLinux",
            type_handler_version="1.4",
            auto_upgrade_minor_version=True)
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_network_packet_capture = azure.network.NetworkPacketCapture("exampleNetworkPacketCapture",
            network_watcher_name=example_network_watcher.name,
            resource_group_name=example_resource_group.name,
            target_resource_id=example_virtual_machine.id,
            storage_location={
                "storage_account_id": example_account.id,
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the Resource to capture packets from. Changing this forces a new resource to be created.

        The **filters** object supports the following:

          * `localIpAddress` (`pulumi.Input[str]`) - The local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
          * `localPort` (`pulumi.Input[str]`) - The local port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
          * `protocol` (`pulumi.Input[str]`) - The Protocol to be filtered on. Possible values include `Any`, `TCP` and `UDP`. Changing this forces a new resource to be created.
          * `remoteIpAddress` (`pulumi.Input[str]`) - The remote IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported.. Changing this forces a new resource to be created.
          * `remotePort` (`pulumi.Input[str]`) - The remote port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.

        The **storage_location** object supports the following:

          * `filePath` (`pulumi.Input[str]`) - A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with `/var/captures`.
          * `storage_account_id` (`pulumi.Input[str]`) - The ID of the storage account to save the packet capture session
          * `storagePath` (`pulumi.Input[str]`) - The URI of the storage path to save the packet capture.
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
    def get(resource_name, id, opts=None, filters=None, maximum_bytes_per_packet=None, maximum_bytes_per_session=None, maximum_capture_duration=None, name=None, network_watcher_name=None, resource_group_name=None, storage_location=None, target_resource_id=None):
        """
        Get an existing NetworkPacketCapture resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] filters: One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_bytes_per_packet: The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_bytes_per_session: Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
        :param pulumi.Input[float] maximum_capture_duration: The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_name: The name of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] storage_location: A `storage_location` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the Resource to capture packets from. Changing this forces a new resource to be created.

        The **filters** object supports the following:

          * `localIpAddress` (`pulumi.Input[str]`) - The local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
          * `localPort` (`pulumi.Input[str]`) - The local port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.
          * `protocol` (`pulumi.Input[str]`) - The Protocol to be filtered on. Possible values include `Any`, `TCP` and `UDP`. Changing this forces a new resource to be created.
          * `remoteIpAddress` (`pulumi.Input[str]`) - The remote IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported.. Changing this forces a new resource to be created.
          * `remotePort` (`pulumi.Input[str]`) - The remote port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Changing this forces a new resource to be created.

        The **storage_location** object supports the following:

          * `filePath` (`pulumi.Input[str]`) - A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with `/var/captures`.
          * `storage_account_id` (`pulumi.Input[str]`) - The ID of the storage account to save the packet capture session
          * `storagePath` (`pulumi.Input[str]`) - The URI of the storage path to save the packet capture.
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
