# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetVirtualNetworkResult',
    'AwaitableGetVirtualNetworkResult',
    'get_virtual_network',
]


class GetVirtualNetworkResult:
    """
    A collection of values returned by getVirtualNetwork.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, address_spaces=None, dns_servers=None, guid=None, id=None, location=None, name=None, resource_group_name=None, subnets=None, vnet_peerings=None) -> None:
        if address_spaces and not isinstance(address_spaces, list):
            raise TypeError("Expected argument 'address_spaces' to be a list")
        __self__.address_spaces = address_spaces
        """
        The list of address spaces used by the virtual network.
        """
        if dns_servers and not isinstance(dns_servers, list):
            raise TypeError("Expected argument 'dns_servers' to be a list")
        __self__.dns_servers = dns_servers
        """
        The list of DNS servers used by the virtual network.
        """
        if guid and not isinstance(guid, str):
            raise TypeError("Expected argument 'guid' to be a str")
        __self__.guid = guid
        """
        The GUID of the virtual network.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Location of the virtual network.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        __self__.subnets = subnets
        """
        The list of name of the subnets that are attached to this virtual network.
        """
        if vnet_peerings and not isinstance(vnet_peerings, dict):
            raise TypeError("Expected argument 'vnet_peerings' to be a dict")
        __self__.vnet_peerings = vnet_peerings
        """
        A mapping of name - virtual network id of the virtual network peerings.
        """


class AwaitableGetVirtualNetworkResult(GetVirtualNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualNetworkResult(
            address_spaces=self.address_spaces,
            dns_servers=self.dns_servers,
            guid=self.guid,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            subnets=self.subnets,
            vnet_peerings=self.vnet_peerings)


def get_virtual_network(name: Optional[str] = None, resource_group_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualNetworkResult:
    """
    Use this data source to access information about an existing Virtual Network.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_virtual_network(name="production",
        resource_group_name="networking")
    pulumi.export("virtualNetworkId", example.id)
    ```


    :param str name: Specifies the name of the Virtual Network.
    :param str resource_group_name: Specifies the name of the resource group the Virtual Network is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getVirtualNetwork:getVirtualNetwork', __args__, opts=opts).value

    return AwaitableGetVirtualNetworkResult(
        address_spaces=__ret__.get('addressSpaces'),
        dns_servers=__ret__.get('dnsServers'),
        guid=__ret__.get('guid'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        subnets=__ret__.get('subnets'),
        vnet_peerings=__ret__.get('vnetPeerings'))
