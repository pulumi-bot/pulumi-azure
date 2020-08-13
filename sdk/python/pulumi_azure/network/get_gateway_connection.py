# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetGatewayConnectionResult',
    'AwaitableGetGatewayConnectionResult',
    'get_gateway_connection',
]


@pulumi.output_type
class _GetGatewayConnectionResult(dict):
    authorization_key: str = pulumi.property("authorizationKey")
    connection_protocol: str = pulumi.property("connectionProtocol")
    egress_bytes_transferred: float = pulumi.property("egressBytesTransferred")
    enable_bgp: bool = pulumi.property("enableBgp")
    express_route_circuit_id: str = pulumi.property("expressRouteCircuitId")
    express_route_gateway_bypass: bool = pulumi.property("expressRouteGatewayBypass")
    id: str = pulumi.property("id")
    ingress_bytes_transferred: float = pulumi.property("ingressBytesTransferred")
    ipsec_policies: List['outputs.GetGatewayConnectionIpsecPolicyResult'] = pulumi.property("ipsecPolicies")
    local_network_gateway_id: str = pulumi.property("localNetworkGatewayId")
    location: str = pulumi.property("location")
    name: str = pulumi.property("name")
    peer_virtual_network_gateway_id: str = pulumi.property("peerVirtualNetworkGatewayId")
    resource_group_name: str = pulumi.property("resourceGroupName")
    resource_guid: str = pulumi.property("resourceGuid")
    routing_weight: float = pulumi.property("routingWeight")
    shared_key: str = pulumi.property("sharedKey")
    tags: Mapping[str, str] = pulumi.property("tags")
    traffic_selector_policy: 'outputs.GetGatewayConnectionTrafficSelectorPolicyResult' = pulumi.property("trafficSelectorPolicy")
    type: str = pulumi.property("type")
    use_policy_based_traffic_selectors: bool = pulumi.property("usePolicyBasedTrafficSelectors")
    virtual_network_gateway_id: str = pulumi.property("virtualNetworkGatewayId")


class GetGatewayConnectionResult:
    """
    A collection of values returned by getGatewayConnection.
    """
    def __init__(__self__, authorization_key=None, connection_protocol=None, egress_bytes_transferred=None, enable_bgp=None, express_route_circuit_id=None, express_route_gateway_bypass=None, id=None, ingress_bytes_transferred=None, ipsec_policies=None, local_network_gateway_id=None, location=None, name=None, peer_virtual_network_gateway_id=None, resource_group_name=None, resource_guid=None, routing_weight=None, shared_key=None, tags=None, traffic_selector_policy=None, type=None, use_policy_based_traffic_selectors=None, virtual_network_gateway_id=None):
        if authorization_key and not isinstance(authorization_key, str):
            raise TypeError("Expected argument 'authorization_key' to be a str")
        __self__.authorization_key = authorization_key
        """
        The authorization key associated with the
        Express Route Circuit. This field is present only if the type is an
        ExpressRoute connection.
        """
        if connection_protocol and not isinstance(connection_protocol, str):
            raise TypeError("Expected argument 'connection_protocol' to be a str")
        __self__.connection_protocol = connection_protocol
        if egress_bytes_transferred and not isinstance(egress_bytes_transferred, float):
            raise TypeError("Expected argument 'egress_bytes_transferred' to be a float")
        __self__.egress_bytes_transferred = egress_bytes_transferred
        if enable_bgp and not isinstance(enable_bgp, bool):
            raise TypeError("Expected argument 'enable_bgp' to be a bool")
        __self__.enable_bgp = enable_bgp
        """
        If `true`, BGP (Border Gateway Protocol) is enabled
        for this connection.
        """
        if express_route_circuit_id and not isinstance(express_route_circuit_id, str):
            raise TypeError("Expected argument 'express_route_circuit_id' to be a str")
        __self__.express_route_circuit_id = express_route_circuit_id
        """
        The ID of the Express Route Circuit
        (i.e. when `type` is `ExpressRoute`).
        """
        if express_route_gateway_bypass and not isinstance(express_route_gateway_bypass, bool):
            raise TypeError("Expected argument 'express_route_gateway_bypass' to be a bool")
        __self__.express_route_gateway_bypass = express_route_gateway_bypass
        """
        If `true`, data packets will bypass ExpressRoute Gateway for data forwarding. This is only valid for ExpressRoute connections.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ingress_bytes_transferred and not isinstance(ingress_bytes_transferred, float):
            raise TypeError("Expected argument 'ingress_bytes_transferred' to be a float")
        __self__.ingress_bytes_transferred = ingress_bytes_transferred
        if ipsec_policies and not isinstance(ipsec_policies, list):
            raise TypeError("Expected argument 'ipsec_policies' to be a list")
        __self__.ipsec_policies = ipsec_policies
        if local_network_gateway_id and not isinstance(local_network_gateway_id, str):
            raise TypeError("Expected argument 'local_network_gateway_id' to be a str")
        __self__.local_network_gateway_id = local_network_gateway_id
        """
        The ID of the local network gateway
        when a Site-to-Site connection (i.e. when `type` is `IPsec`).
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location/region where the connection is
        located.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if peer_virtual_network_gateway_id and not isinstance(peer_virtual_network_gateway_id, str):
            raise TypeError("Expected argument 'peer_virtual_network_gateway_id' to be a str")
        __self__.peer_virtual_network_gateway_id = peer_virtual_network_gateway_id
        """
        The ID of the peer virtual
        network gateway when a VNet-to-VNet connection (i.e. when `type`
        is `Vnet2Vnet`).
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        __self__.resource_guid = resource_guid
        if routing_weight and not isinstance(routing_weight, float):
            raise TypeError("Expected argument 'routing_weight' to be a float")
        __self__.routing_weight = routing_weight
        """
        The routing weight.
        """
        if shared_key and not isinstance(shared_key, str):
            raise TypeError("Expected argument 'shared_key' to be a str")
        __self__.shared_key = shared_key
        """
        The shared IPSec key.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        if traffic_selector_policy and not isinstance(traffic_selector_policy, dict):
            raise TypeError("Expected argument 'traffic_selector_policy' to be a dict")
        __self__.traffic_selector_policy = traffic_selector_policy
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of connection. Valid options are `IPsec`
        (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
        """
        if use_policy_based_traffic_selectors and not isinstance(use_policy_based_traffic_selectors, bool):
            raise TypeError("Expected argument 'use_policy_based_traffic_selectors' to be a bool")
        __self__.use_policy_based_traffic_selectors = use_policy_based_traffic_selectors
        """
        If `true`, policy-based traffic
        selectors are enabled for this connection. Enabling policy-based traffic
        selectors requires an `ipsec_policy` block.
        """
        if virtual_network_gateway_id and not isinstance(virtual_network_gateway_id, str):
            raise TypeError("Expected argument 'virtual_network_gateway_id' to be a str")
        __self__.virtual_network_gateway_id = virtual_network_gateway_id
        """
        The ID of the Virtual Network Gateway
        in which the connection is created.
        """


class AwaitableGetGatewayConnectionResult(GetGatewayConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayConnectionResult(
            authorization_key=self.authorization_key,
            connection_protocol=self.connection_protocol,
            egress_bytes_transferred=self.egress_bytes_transferred,
            enable_bgp=self.enable_bgp,
            express_route_circuit_id=self.express_route_circuit_id,
            express_route_gateway_bypass=self.express_route_gateway_bypass,
            id=self.id,
            ingress_bytes_transferred=self.ingress_bytes_transferred,
            ipsec_policies=self.ipsec_policies,
            local_network_gateway_id=self.local_network_gateway_id,
            location=self.location,
            name=self.name,
            peer_virtual_network_gateway_id=self.peer_virtual_network_gateway_id,
            resource_group_name=self.resource_group_name,
            resource_guid=self.resource_guid,
            routing_weight=self.routing_weight,
            shared_key=self.shared_key,
            tags=self.tags,
            traffic_selector_policy=self.traffic_selector_policy,
            type=self.type,
            use_policy_based_traffic_selectors=self.use_policy_based_traffic_selectors,
            virtual_network_gateway_id=self.virtual_network_gateway_id)


def get_gateway_connection(name: Optional[str] = None,
                           resource_group_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayConnectionResult:
    """
    Use this data source to access information about an existing Virtual Network Gateway Connection.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_gateway_connection(name="production",
        resource_group_name="networking")
    pulumi.export("virtualNetworkGatewayConnectionId", example.id)
    ```


    :param str name: Specifies the name of the Virtual Network Gateway Connection.
    :param str resource_group_name: Specifies the name of the resource group the Virtual Network Gateway Connection is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getGatewayConnection:getGatewayConnection', __args__, opts=opts, typ=_GetGatewayConnectionResult).value

    return AwaitableGetGatewayConnectionResult(
        authorization_key=_utilities.get_dict_value(__ret__, 'authorizationKey'),
        connection_protocol=_utilities.get_dict_value(__ret__, 'connectionProtocol'),
        egress_bytes_transferred=_utilities.get_dict_value(__ret__, 'egressBytesTransferred'),
        enable_bgp=_utilities.get_dict_value(__ret__, 'enableBgp'),
        express_route_circuit_id=_utilities.get_dict_value(__ret__, 'expressRouteCircuitId'),
        express_route_gateway_bypass=_utilities.get_dict_value(__ret__, 'expressRouteGatewayBypass'),
        id=_utilities.get_dict_value(__ret__, 'id'),
        ingress_bytes_transferred=_utilities.get_dict_value(__ret__, 'ingressBytesTransferred'),
        ipsec_policies=_utilities.get_dict_value(__ret__, 'ipsecPolicies'),
        local_network_gateway_id=_utilities.get_dict_value(__ret__, 'localNetworkGatewayId'),
        location=_utilities.get_dict_value(__ret__, 'location'),
        name=_utilities.get_dict_value(__ret__, 'name'),
        peer_virtual_network_gateway_id=_utilities.get_dict_value(__ret__, 'peerVirtualNetworkGatewayId'),
        resource_group_name=_utilities.get_dict_value(__ret__, 'resourceGroupName'),
        resource_guid=_utilities.get_dict_value(__ret__, 'resourceGuid'),
        routing_weight=_utilities.get_dict_value(__ret__, 'routingWeight'),
        shared_key=_utilities.get_dict_value(__ret__, 'sharedKey'),
        tags=_utilities.get_dict_value(__ret__, 'tags'),
        traffic_selector_policy=_utilities.get_dict_value(__ret__, 'trafficSelectorPolicy'),
        type=_utilities.get_dict_value(__ret__, 'type'),
        use_policy_based_traffic_selectors=_utilities.get_dict_value(__ret__, 'usePolicyBasedTrafficSelectors'),
        virtual_network_gateway_id=_utilities.get_dict_value(__ret__, 'virtualNetworkGatewayId'))
