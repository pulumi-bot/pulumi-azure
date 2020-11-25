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

__all__ = ['VpnGatewayConnection']


class VpnGatewayConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internet_security_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_vpn_site_id: Optional[pulumi.Input[str]] = None,
                 routings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpn_links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a VPN Gateway Connection.

        ## Import

        VPN Gateway Connections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/vpnGatewayConnection:VpnGatewayConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/conn1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internet_security_enabled: Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        :param pulumi.Input[str] name: The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[str] remote_vpn_site_id: The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]] routings: A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]] vpn_links: One or more `vpn_link` blocks as defined below.
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

            __props__['internet_security_enabled'] = internet_security_enabled
            __props__['name'] = name
            if remote_vpn_site_id is None:
                raise TypeError("Missing required property 'remote_vpn_site_id'")
            __props__['remote_vpn_site_id'] = remote_vpn_site_id
            __props__['routings'] = routings
            if vpn_gateway_id is None:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__['vpn_gateway_id'] = vpn_gateway_id
            if vpn_links is None:
                raise TypeError("Missing required property 'vpn_links'")
            __props__['vpn_links'] = vpn_links
        super(VpnGatewayConnection, __self__).__init__(
            'azure:network/vpnGatewayConnection:VpnGatewayConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            internet_security_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote_vpn_site_id: Optional[pulumi.Input[str]] = None,
            routings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None,
            vpn_links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]]] = None) -> 'VpnGatewayConnection':
        """
        Get an existing VpnGatewayConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internet_security_enabled: Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        :param pulumi.Input[str] name: The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[str] remote_vpn_site_id: The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]] routings: A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]] vpn_links: One or more `vpn_link` blocks as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["internet_security_enabled"] = internet_security_enabled
        __props__["name"] = name
        __props__["remote_vpn_site_id"] = remote_vpn_site_id
        __props__["routings"] = routings
        __props__["vpn_gateway_id"] = vpn_gateway_id
        __props__["vpn_links"] = vpn_links
        return VpnGatewayConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="internetSecurityEnabled")
    def internet_security_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        """
        return pulumi.get(self, "internet_security_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteVpnSiteId")
    def remote_vpn_site_id(self) -> pulumi.Output[str]:
        """
        The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "remote_vpn_site_id")

    @property
    @pulumi.getter
    def routings(self) -> pulumi.Output[Sequence['outputs.VpnGatewayConnectionRouting']]:
        """
        A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        """
        return pulumi.get(self, "routings")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @property
    @pulumi.getter(name="vpnLinks")
    def vpn_links(self) -> pulumi.Output[Sequence['outputs.VpnGatewayConnectionVpnLink']]:
        """
        One or more `vpn_link` blocks as defined below.
        """
        return pulumi.get(self, "vpn_links")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

