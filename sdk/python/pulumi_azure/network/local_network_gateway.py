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

__all__ = ['LocalNetworkGateway']


class LocalNetworkGateway(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_spaces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bgp_settings: Optional[pulumi.Input[pulumi.InputType['LocalNetworkGatewayBgpSettingsArgs']]] = None,
                 gateway_address: Optional[pulumi.Input[str]] = None,
                 gateway_fqdn: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a local network gateway connection over which specific connections can be configured.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West US")
        home = azure.network.LocalNetworkGateway("home",
            resource_group_name=example.name,
            location=example.location,
            gateway_address="12.13.14.15",
            address_spaces=["10.0.0.0/16"])
        ```

        ## Import

        Local Network Gateways can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/localNetworkGateway:LocalNetworkGateway lng1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/localNetworkGateways/lng1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_spaces: The list of string CIDRs representing the
               address spaces the gateway exposes.
        :param pulumi.Input[pulumi.InputType['LocalNetworkGatewayBgpSettingsArgs']] bgp_settings: A `bgp_settings` block as defined below containing the
               Local Network Gateway's BGP speaker settings.
        :param pulumi.Input[str] gateway_address: The gateway IP address to connect with.
        :param pulumi.Input[str] gateway_fqdn: The gateway FQDN to connect with.
        :param pulumi.Input[str] location: The location/region where the local network gateway is
               created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the local network gateway. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the local network gateway.
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

            if address_spaces is None:
                raise TypeError("Missing required property 'address_spaces'")
            __props__['address_spaces'] = address_spaces
            __props__['bgp_settings'] = bgp_settings
            __props__['gateway_address'] = gateway_address
            __props__['gateway_fqdn'] = gateway_fqdn
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(LocalNetworkGateway, __self__).__init__(
            'azure:network/localNetworkGateway:LocalNetworkGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_spaces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bgp_settings: Optional[pulumi.Input[pulumi.InputType['LocalNetworkGatewayBgpSettingsArgs']]] = None,
            gateway_address: Optional[pulumi.Input[str]] = None,
            gateway_fqdn: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'LocalNetworkGateway':
        """
        Get an existing LocalNetworkGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_spaces: The list of string CIDRs representing the
               address spaces the gateway exposes.
        :param pulumi.Input[pulumi.InputType['LocalNetworkGatewayBgpSettingsArgs']] bgp_settings: A `bgp_settings` block as defined below containing the
               Local Network Gateway's BGP speaker settings.
        :param pulumi.Input[str] gateway_address: The gateway IP address to connect with.
        :param pulumi.Input[str] gateway_fqdn: The gateway FQDN to connect with.
        :param pulumi.Input[str] location: The location/region where the local network gateway is
               created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the local network gateway. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the local network gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_spaces"] = address_spaces
        __props__["bgp_settings"] = bgp_settings
        __props__["gateway_address"] = gateway_address
        __props__["gateway_fqdn"] = gateway_fqdn
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return LocalNetworkGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressSpaces")
    def address_spaces(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of string CIDRs representing the
        address spaces the gateway exposes.
        """
        return pulumi.get(self, "address_spaces")

    @property
    @pulumi.getter(name="bgpSettings")
    def bgp_settings(self) -> pulumi.Output[Optional['outputs.LocalNetworkGatewayBgpSettings']]:
        """
        A `bgp_settings` block as defined below containing the
        Local Network Gateway's BGP speaker settings.
        """
        return pulumi.get(self, "bgp_settings")

    @property
    @pulumi.getter(name="gatewayAddress")
    def gateway_address(self) -> pulumi.Output[Optional[str]]:
        """
        The gateway IP address to connect with.
        """
        return pulumi.get(self, "gateway_address")

    @property
    @pulumi.getter(name="gatewayFqdn")
    def gateway_fqdn(self) -> pulumi.Output[Optional[str]]:
        """
        The gateway FQDN to connect with.
        """
        return pulumi.get(self, "gateway_fqdn")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location/region where the local network gateway is
        created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the local network gateway. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to
        create the local network gateway.
        """
        return pulumi.get(self, "resource_group_name")

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

