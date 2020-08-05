# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetServiceEndpointConnectionsResult:
    """
    A collection of values returned by getServiceEndpointConnections.
    """
    def __init__(__self__, id=None, location=None, private_endpoint_connections=None, resource_group_name=None, service_id=None, service_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if private_endpoint_connections and not isinstance(private_endpoint_connections, list):
            raise TypeError("Expected argument 'private_endpoint_connections' to be a list")
        __self__.private_endpoint_connections = private_endpoint_connections
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        __self__.service_id = service_id
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        __self__.service_name = service_name
        """
        The name of the private link service.
        """


class AwaitableGetServiceEndpointConnectionsResult(GetServiceEndpointConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceEndpointConnectionsResult(
            id=self.id,
            location=self.location,
            private_endpoint_connections=self.private_endpoint_connections,
            resource_group_name=self.resource_group_name,
            service_id=self.service_id,
            service_name=self.service_name)


def get_service_endpoint_connections(resource_group_name=None, service_id=None, opts=None):
    """
    Use this data source to access endpoint connection information about an existing Private Link Service.

    > **NOTE** Private Link is currently in Public Preview.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.privatelink.get_service_endpoint_connections(service_id=azurerm_private_link_service["example"]["id"],
        resource_group_name=azurerm_resource_group["example"]["name"])
    pulumi.export("privateEndpointStatus", example.private_endpoint_connections[0]["status"])
    ```


    :param str resource_group_name: The name of the resource group in which the private link service resides.
    :param str service_id: The resource ID of the private link service.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceId'] = service_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:privatelink/getServiceEndpointConnections:getServiceEndpointConnections', __args__, opts=opts).value

    return AwaitableGetServiceEndpointConnectionsResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        private_endpoint_connections=__ret__.get('privateEndpointConnections'),
        resource_group_name=__ret__.get('resourceGroupName'),
        service_id=__ret__.get('serviceId'),
        service_name=__ret__.get('serviceName'))
