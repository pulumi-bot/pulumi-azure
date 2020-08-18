# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetNetworkInterfaceResult:
    """
    A collection of values returned by getNetworkInterface.
    """
    def __init__(__self__, applied_dns_servers=None, dns_servers=None, enable_accelerated_networking=None, enable_ip_forwarding=None, id=None, internal_dns_name_label=None, ip_configurations=None, location=None, mac_address=None, name=None, network_security_group_id=None, private_ip_address=None, private_ip_addresses=None, resource_group_name=None, tags=None, virtual_machine_id=None):
        if applied_dns_servers and not isinstance(applied_dns_servers, list):
            raise TypeError("Expected argument 'applied_dns_servers' to be a list")
        __self__.applied_dns_servers = applied_dns_servers
        """
        List of DNS servers applied to the specified Network Interface.
        """
        if dns_servers and not isinstance(dns_servers, list):
            raise TypeError("Expected argument 'dns_servers' to be a list")
        __self__.dns_servers = dns_servers
        """
        The list of DNS servers used by the specified Network Interface.
        """
        if enable_accelerated_networking and not isinstance(enable_accelerated_networking, bool):
            raise TypeError("Expected argument 'enable_accelerated_networking' to be a bool")
        __self__.enable_accelerated_networking = enable_accelerated_networking
        """
        Indicates if accelerated networking is set on the specified Network Interface.
        """
        if enable_ip_forwarding and not isinstance(enable_ip_forwarding, bool):
            raise TypeError("Expected argument 'enable_ip_forwarding' to be a bool")
        __self__.enable_ip_forwarding = enable_ip_forwarding
        """
        Indicate if IP forwarding is set on the specified Network Interface.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if internal_dns_name_label and not isinstance(internal_dns_name_label, str):
            raise TypeError("Expected argument 'internal_dns_name_label' to be a str")
        __self__.internal_dns_name_label = internal_dns_name_label
        """
        The internal dns name label of the specified Network Interface.
        """
        if ip_configurations and not isinstance(ip_configurations, list):
            raise TypeError("Expected argument 'ip_configurations' to be a list")
        __self__.ip_configurations = ip_configurations
        """
        One or more `ip_configuration` blocks as defined below.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location of the specified Network Interface.
        """
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        __self__.mac_address = mac_address
        """
        The MAC address used by the specified Network Interface.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the IP Configuration.
        """
        if network_security_group_id and not isinstance(network_security_group_id, str):
            raise TypeError("Expected argument 'network_security_group_id' to be a str")
        __self__.network_security_group_id = network_security_group_id
        """
        The ID of the network security group associated to the specified Network Interface.
        """
        if private_ip_address and not isinstance(private_ip_address, str):
            raise TypeError("Expected argument 'private_ip_address' to be a str")
        __self__.private_ip_address = private_ip_address
        """
        The Private IP Address assigned to this Network Interface.
        """
        if private_ip_addresses and not isinstance(private_ip_addresses, list):
            raise TypeError("Expected argument 'private_ip_addresses' to be a list")
        __self__.private_ip_addresses = private_ip_addresses
        """
        The list of private ip addresses associates to the specified Network Interface.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        List the tags associated to the specified Network Interface.
        """
        if virtual_machine_id and not isinstance(virtual_machine_id, str):
            raise TypeError("Expected argument 'virtual_machine_id' to be a str")
        __self__.virtual_machine_id = virtual_machine_id
        """
        The ID of the virtual machine that the specified Network Interface is attached to.
        """


class AwaitableGetNetworkInterfaceResult(GetNetworkInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfaceResult(
            applied_dns_servers=self.applied_dns_servers,
            dns_servers=self.dns_servers,
            enable_accelerated_networking=self.enable_accelerated_networking,
            enable_ip_forwarding=self.enable_ip_forwarding,
            id=self.id,
            internal_dns_name_label=self.internal_dns_name_label,
            ip_configurations=self.ip_configurations,
            location=self.location,
            mac_address=self.mac_address,
            name=self.name,
            network_security_group_id=self.network_security_group_id,
            private_ip_address=self.private_ip_address,
            private_ip_addresses=self.private_ip_addresses,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            virtual_machine_id=self.virtual_machine_id)


def get_network_interface(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing Network Interface.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_network_interface(name="acctest-nic",
        resource_group_name="networking")
    pulumi.export("networkInterfaceId", example.id)
    ```


    :param str name: Specifies the name of the Network Interface.
    :param str resource_group_name: Specifies the name of the resource group the Network Interface is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getNetworkInterface:getNetworkInterface', __args__, opts=opts).value

    return AwaitableGetNetworkInterfaceResult(
        applied_dns_servers=__ret__.get('appliedDnsServers'),
        dns_servers=__ret__.get('dnsServers'),
        enable_accelerated_networking=__ret__.get('enableAcceleratedNetworking'),
        enable_ip_forwarding=__ret__.get('enableIpForwarding'),
        id=__ret__.get('id'),
        internal_dns_name_label=__ret__.get('internalDnsNameLabel'),
        ip_configurations=__ret__.get('ipConfigurations'),
        location=__ret__.get('location'),
        mac_address=__ret__.get('macAddress'),
        name=__ret__.get('name'),
        network_security_group_id=__ret__.get('networkSecurityGroupId'),
        private_ip_address=__ret__.get('privateIpAddress'),
        private_ip_addresses=__ret__.get('privateIpAddresses'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        virtual_machine_id=__ret__.get('virtualMachineId'))
