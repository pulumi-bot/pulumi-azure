# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetEndpointConnectionResult',
    'AwaitableGetEndpointConnectionResult',
    'get_endpoint_connection',
]

@pulumi.output_type
class GetEndpointConnectionResult:
    """
    A collection of values returned by getEndpointConnection.
    """
    def __init__(__self__, id=None, location=None, name=None, private_service_connections=None, resource_group_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_service_connections and not isinstance(private_service_connections, list):
            raise TypeError("Expected argument 'private_service_connections' to be a list")
        pulumi.set(__self__, "private_service_connections", private_service_connections)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The supported Azure location where the resource exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the private endpoint.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateServiceConnections")
    def private_service_connections(self) -> Sequence['outputs.GetEndpointConnectionPrivateServiceConnectionResult']:
        return pulumi.get(self, "private_service_connections")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")


class AwaitableGetEndpointConnectionResult(GetEndpointConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointConnectionResult(
            id=self.id,
            location=self.location,
            name=self.name,
            private_service_connections=self.private_service_connections,
            resource_group_name=self.resource_group_name)


def get_endpoint_connection(name: Optional[str] = None,
                            resource_group_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointConnectionResult:
    """
    Use this data source to access the connection status information about an existing Private Endpoint Connection.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.privatelink.get_endpoint_connection(name="example-private-endpoint",
        resource_group_name="example-rg")
    pulumi.export("privateEndpointStatus", example.private_service_connections[0].status)
    ```


    :param str name: Specifies the Name of the private endpoint.
    :param str resource_group_name: Specifies the Name of the Resource Group within which the private endpoint exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:privatelink/getEndpointConnection:getEndpointConnection', __args__, opts=opts, typ=GetEndpointConnectionResult).value

    return AwaitableGetEndpointConnectionResult(
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        private_service_connections=__ret__.private_service_connections,
        resource_group_name=__ret__.resource_group_name)
