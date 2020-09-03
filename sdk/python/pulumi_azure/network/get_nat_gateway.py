# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetNatGatewayResult',
    'AwaitableGetNatGatewayResult',
    'get_nat_gateway',
]

@pulumi.output_type
class GetNatGatewayResult:
    """
    A collection of values returned by getNatGateway.
    """
    def __init__(__self__, id=None, idle_timeout_in_minutes=None, location=None, name=None, public_ip_address_ids=None, public_ip_prefix_ids=None, resource_group_name=None, resource_guid=None, sku_name=None, tags=None, zones=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idle_timeout_in_minutes and not isinstance(idle_timeout_in_minutes, float):
            raise TypeError("Expected argument 'idle_timeout_in_minutes' to be a float")
        pulumi.set(__self__, "idle_timeout_in_minutes", idle_timeout_in_minutes)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if public_ip_address_ids and not isinstance(public_ip_address_ids, list):
            raise TypeError("Expected argument 'public_ip_address_ids' to be a list")
        pulumi.set(__self__, "public_ip_address_ids", public_ip_address_ids)
        if public_ip_prefix_ids and not isinstance(public_ip_prefix_ids, list):
            raise TypeError("Expected argument 'public_ip_prefix_ids' to be a list")
        pulumi.set(__self__, "public_ip_prefix_ids", public_ip_prefix_ids)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        pulumi.set(__self__, "resource_guid", resource_guid)
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        pulumi.set(__self__, "sku_name", sku_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idleTimeoutInMinutes")
    def idle_timeout_in_minutes(self) -> float:
        """
        The idle timeout in minutes which is used for the NAT Gateway.
        """
        return pulumi.get(self, "idle_timeout_in_minutes")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location where the NAT Gateway exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIpAddressIds")
    def public_ip_address_ids(self) -> Sequence[str]:
        """
        A list of existing Public IP Address resource IDs which the NAT Gateway is using.
        """
        return pulumi.get(self, "public_ip_address_ids")

    @property
    @pulumi.getter(name="publicIpPrefixIds")
    def public_ip_prefix_ids(self) -> Sequence[str]:
        """
        A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
        """
        return pulumi.get(self, "public_ip_prefix_ids")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> str:
        """
        The Resource GUID of the NAT Gateway.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        """
        The SKU used by the NAT Gateway.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[str]:
        """
        A list of Availability Zones which the NAT Gateway exists in.
        """
        return pulumi.get(self, "zones")


class AwaitableGetNatGatewayResult(GetNatGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNatGatewayResult(
            id=self.id,
            idle_timeout_in_minutes=self.idle_timeout_in_minutes,
            location=self.location,
            name=self.name,
            public_ip_address_ids=self.public_ip_address_ids,
            public_ip_prefix_ids=self.public_ip_prefix_ids,
            resource_group_name=self.resource_group_name,
            resource_guid=self.resource_guid,
            sku_name=self.sku_name,
            tags=self.tags,
            zones=self.zones)


def get_nat_gateway(name: Optional[str] = None,
                    public_ip_address_ids: Optional[Sequence[str]] = None,
                    public_ip_prefix_ids: Optional[Sequence[str]] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNatGatewayResult:
    """
    Use this data source to access information about an existing NAT Gateway.


    :param str name: Specifies the Name of the NAT Gateway.
    :param Sequence[str] public_ip_address_ids: A list of existing Public IP Address resource IDs which the NAT Gateway is using.
    :param Sequence[str] public_ip_prefix_ids: A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
    :param str resource_group_name: Specifies the name of the Resource Group where the NAT Gateway exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['publicIpAddressIds'] = public_ip_address_ids
    __args__['publicIpPrefixIds'] = public_ip_prefix_ids
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getNatGateway:getNatGateway', __args__, opts=opts, typ=GetNatGatewayResult).value

    return AwaitableGetNatGatewayResult(
        id=__ret__.id,
        idle_timeout_in_minutes=__ret__.idle_timeout_in_minutes,
        location=__ret__.location,
        name=__ret__.name,
        public_ip_address_ids=__ret__.public_ip_address_ids,
        public_ip_prefix_ids=__ret__.public_ip_prefix_ids,
        resource_group_name=__ret__.resource_group_name,
        resource_guid=__ret__.resource_guid,
        sku_name=__ret__.sku_name,
        tags=__ret__.tags,
        zones=__ret__.zones)
