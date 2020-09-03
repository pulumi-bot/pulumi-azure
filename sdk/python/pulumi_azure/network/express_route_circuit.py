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

__all__ = ['ExpressRouteCircuit']


class ExpressRouteCircuit(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_classic_operations: Optional[pulumi.Input[bool]] = None,
                 bandwidth_in_mbps: Optional[pulumi.Input[float]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peering_location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_provider_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an ExpressRoute circuit.

        ## Example Usage

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
            sku=azure.network.ExpressRouteCircuitSkuArgs(
                tier="Standard",
                family="MeteredData",
            ),
            tags={
                "environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_classic_operations: Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
        :param pulumi.Input[float] bandwidth_in_mbps: The bandwidth in Mbps of the circuit being created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] peering_location: The name of the peering location and **not** the Azure resource location.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_provider_name: The name of the ExpressRoute Service Provider.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']] sku: A `sku` block for the ExpressRoute circuit as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['allow_classic_operations'] = allow_classic_operations
            if bandwidth_in_mbps is None:
                raise TypeError("Missing required property 'bandwidth_in_mbps'")
            __props__['bandwidth_in_mbps'] = bandwidth_in_mbps
            __props__['location'] = location
            __props__['name'] = name
            if peering_location is None:
                raise TypeError("Missing required property 'peering_location'")
            __props__['peering_location'] = peering_location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_provider_name is None:
                raise TypeError("Missing required property 'service_provider_name'")
            __props__['service_provider_name'] = service_provider_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['service_key'] = None
            __props__['service_provider_provisioning_state'] = None
        super(ExpressRouteCircuit, __self__).__init__(
            'azure:network/expressRouteCircuit:ExpressRouteCircuit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_classic_operations: Optional[pulumi.Input[bool]] = None,
            bandwidth_in_mbps: Optional[pulumi.Input[float]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            peering_location: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_key: Optional[pulumi.Input[str]] = None,
            service_provider_name: Optional[pulumi.Input[str]] = None,
            service_provider_provisioning_state: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ExpressRouteCircuit':
        """
        Get an existing ExpressRouteCircuit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_classic_operations: Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
        :param pulumi.Input[float] bandwidth_in_mbps: The bandwidth in Mbps of the circuit being created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] peering_location: The name of the peering location and **not** the Azure resource location.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_key: The string needed by the service provider to provision the ExpressRoute circuit.
        :param pulumi.Input[str] service_provider_name: The name of the ExpressRoute Service Provider.
        :param pulumi.Input[str] service_provider_provisioning_state: The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']] sku: A `sku` block for the ExpressRoute circuit as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_classic_operations"] = allow_classic_operations
        __props__["bandwidth_in_mbps"] = bandwidth_in_mbps
        __props__["location"] = location
        __props__["name"] = name
        __props__["peering_location"] = peering_location
        __props__["resource_group_name"] = resource_group_name
        __props__["service_key"] = service_key
        __props__["service_provider_name"] = service_provider_name
        __props__["service_provider_provisioning_state"] = service_provider_provisioning_state
        __props__["sku"] = sku
        __props__["tags"] = tags
        return ExpressRouteCircuit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowClassicOperations")
    def allow_classic_operations(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
        """
        return pulumi.get(self, "allow_classic_operations")

    @property
    @pulumi.getter(name="bandwidthInMbps")
    def bandwidth_in_mbps(self) -> pulumi.Output[float]:
        """
        The bandwidth in Mbps of the circuit being created.
        """
        return pulumi.get(self, "bandwidth_in_mbps")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringLocation")
    def peering_location(self) -> pulumi.Output[str]:
        """
        The name of the peering location and **not** the Azure resource location.
        """
        return pulumi.get(self, "peering_location")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceKey")
    def service_key(self) -> pulumi.Output[str]:
        """
        The string needed by the service provider to provision the ExpressRoute circuit.
        """
        return pulumi.get(self, "service_key")

    @property
    @pulumi.getter(name="serviceProviderName")
    def service_provider_name(self) -> pulumi.Output[str]:
        """
        The name of the ExpressRoute Service Provider.
        """
        return pulumi.get(self, "service_provider_name")

    @property
    @pulumi.getter(name="serviceProviderProvisioningState")
    def service_provider_provisioning_state(self) -> pulumi.Output[str]:
        """
        The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
        """
        return pulumi.get(self, "service_provider_provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.ExpressRouteCircuitSku']:
        """
        A `sku` block for the ExpressRoute circuit as documented below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

