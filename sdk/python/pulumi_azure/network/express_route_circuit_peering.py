# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ExpressRouteCircuitPeering(pulumi.CustomResource):
    azure_asn: pulumi.Output[float]
    """
    The ASN used by Azure.
    """
    express_route_circuit_name: pulumi.Output[str]
    """
    The name of the ExpressRoute Circuit in which to create the Peering.
    """
    microsoft_peering_config: pulumi.Output[dict]
    """
    A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.

      * `advertisedPublicPrefixes` (`list`) - A list of Advertised Public Prefixes
      * `customerAsn` (`float`) - The CustomerASN of the peering
      * `routingRegistryName` (`str`) - The RoutingRegistryName of the configuration
    """
    peer_asn: pulumi.Output[float]
    """
    The Either a 16-bit or a 32-bit ASN. Can either be public or private..
    """
    peering_type: pulumi.Output[str]
    """
    The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
    """
    primary_azure_port: pulumi.Output[str]
    """
    The Primary Port used by Azure for this Peering.
    """
    primary_peer_address_prefix: pulumi.Output[str]
    """
    A `/30` subnet for the primary link.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the Express Route Circuit Peering. Changing this forces a new resource to be created.
    """
    secondary_azure_port: pulumi.Output[str]
    """
    The Secondary Port used by Azure for this Peering.
    """
    secondary_peer_address_prefix: pulumi.Output[str]
    """
    A `/30` subnet for the secondary link.
    """
    shared_key: pulumi.Output[str]
    """
    The shared key. Can be a maximum of 25 characters.
    """
    vlan_id: pulumi.Output[float]
    """
    A valid VLAN ID to establish this peering on.
    """
    def __init__(__self__, resource_name, opts=None, express_route_circuit_name=None, microsoft_peering_config=None, peer_asn=None, peering_type=None, primary_peer_address_prefix=None, resource_group_name=None, secondary_peer_address_prefix=None, shared_key=None, vlan_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an ExpressRoute Circuit Peering.

        ## Example Usage
        ### Creating A Microsoft Peering)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_express_route_circuit = azure.network.ExpressRouteCircuit("exampleExpressRouteCircuit",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            service_provider_name="Equinix",
            peering_location="Silicon Valley",
            bandwidth_in_mbps=50,
            sku={
                "tier": "Standard",
                "family": "MeteredData",
            },
            allow_classic_operations=False,
            tags={
                "environment": "Production",
            })
        example_express_route_circuit_peering = azure.network.ExpressRouteCircuitPeering("exampleExpressRouteCircuitPeering",
            peering_type="MicrosoftPeering",
            express_route_circuit_name=example_express_route_circuit.name,
            resource_group_name=example_resource_group.name,
            peer_asn=100,
            primary_peer_address_prefix="123.0.0.0/30",
            secondary_peer_address_prefix="123.0.0.4/30",
            vlan_id=300,
            microsoft_peering_config={
                "advertisedPublicPrefixes": ["123.1.0.0/24"],
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] express_route_circuit_name: The name of the ExpressRoute Circuit in which to create the Peering.
        :param pulumi.Input[dict] microsoft_peering_config: A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        :param pulumi.Input[float] peer_asn: The Either a 16-bit or a 32-bit ASN. Can either be public or private..
        :param pulumi.Input[str] peering_type: The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_peer_address_prefix: A `/30` subnet for the primary link.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_peer_address_prefix: A `/30` subnet for the secondary link.
        :param pulumi.Input[str] shared_key: The shared key. Can be a maximum of 25 characters.
        :param pulumi.Input[float] vlan_id: A valid VLAN ID to establish this peering on.

        The **microsoft_peering_config** object supports the following:

          * `advertisedPublicPrefixes` (`pulumi.Input[list]`) - A list of Advertised Public Prefixes
          * `customerAsn` (`pulumi.Input[float]`) - The CustomerASN of the peering
          * `routingRegistryName` (`pulumi.Input[str]`) - The RoutingRegistryName of the configuration
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

            if express_route_circuit_name is None:
                raise TypeError("Missing required property 'express_route_circuit_name'")
            __props__['express_route_circuit_name'] = express_route_circuit_name
            __props__['microsoft_peering_config'] = microsoft_peering_config
            __props__['peer_asn'] = peer_asn
            if peering_type is None:
                raise TypeError("Missing required property 'peering_type'")
            __props__['peering_type'] = peering_type
            if primary_peer_address_prefix is None:
                raise TypeError("Missing required property 'primary_peer_address_prefix'")
            __props__['primary_peer_address_prefix'] = primary_peer_address_prefix
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if secondary_peer_address_prefix is None:
                raise TypeError("Missing required property 'secondary_peer_address_prefix'")
            __props__['secondary_peer_address_prefix'] = secondary_peer_address_prefix
            __props__['shared_key'] = shared_key
            if vlan_id is None:
                raise TypeError("Missing required property 'vlan_id'")
            __props__['vlan_id'] = vlan_id
            __props__['azure_asn'] = None
            __props__['primary_azure_port'] = None
            __props__['secondary_azure_port'] = None
        super(ExpressRouteCircuitPeering, __self__).__init__(
            'azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, azure_asn=None, express_route_circuit_name=None, microsoft_peering_config=None, peer_asn=None, peering_type=None, primary_azure_port=None, primary_peer_address_prefix=None, resource_group_name=None, secondary_azure_port=None, secondary_peer_address_prefix=None, shared_key=None, vlan_id=None):
        """
        Get an existing ExpressRouteCircuitPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] azure_asn: The ASN used by Azure.
        :param pulumi.Input[str] express_route_circuit_name: The name of the ExpressRoute Circuit in which to create the Peering.
        :param pulumi.Input[dict] microsoft_peering_config: A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        :param pulumi.Input[float] peer_asn: The Either a 16-bit or a 32-bit ASN. Can either be public or private..
        :param pulumi.Input[str] peering_type: The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_azure_port: The Primary Port used by Azure for this Peering.
        :param pulumi.Input[str] primary_peer_address_prefix: A `/30` subnet for the primary link.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_azure_port: The Secondary Port used by Azure for this Peering.
        :param pulumi.Input[str] secondary_peer_address_prefix: A `/30` subnet for the secondary link.
        :param pulumi.Input[str] shared_key: The shared key. Can be a maximum of 25 characters.
        :param pulumi.Input[float] vlan_id: A valid VLAN ID to establish this peering on.

        The **microsoft_peering_config** object supports the following:

          * `advertisedPublicPrefixes` (`pulumi.Input[list]`) - A list of Advertised Public Prefixes
          * `customerAsn` (`pulumi.Input[float]`) - The CustomerASN of the peering
          * `routingRegistryName` (`pulumi.Input[str]`) - The RoutingRegistryName of the configuration
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["azure_asn"] = azure_asn
        __props__["express_route_circuit_name"] = express_route_circuit_name
        __props__["microsoft_peering_config"] = microsoft_peering_config
        __props__["peer_asn"] = peer_asn
        __props__["peering_type"] = peering_type
        __props__["primary_azure_port"] = primary_azure_port
        __props__["primary_peer_address_prefix"] = primary_peer_address_prefix
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_azure_port"] = secondary_azure_port
        __props__["secondary_peer_address_prefix"] = secondary_peer_address_prefix
        __props__["shared_key"] = shared_key
        __props__["vlan_id"] = vlan_id
        return ExpressRouteCircuitPeering(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
