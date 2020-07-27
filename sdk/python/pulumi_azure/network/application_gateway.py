# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ApplicationGateway']


class ApplicationGateway(pulumi.CustomResource):
    authentication_certificates: pulumi.Output[Optional[List['outputs.ApplicationGatewayAuthenticationCertificate']]] = pulumi.output_property("authenticationCertificates")
    """
    One or more `authentication_certificate` blocks as defined below.
    """
    autoscale_configuration: pulumi.Output[Optional['outputs.ApplicationGatewayAutoscaleConfiguration']] = pulumi.output_property("autoscaleConfiguration")
    """
    A `autoscale_configuration` block as defined below.
    """
    backend_address_pools: pulumi.Output[List['outputs.ApplicationGatewayBackendAddressPool']] = pulumi.output_property("backendAddressPools")
    """
    One or more `backend_address_pool` blocks as defined below.
    """
    backend_http_settings: pulumi.Output[List['outputs.ApplicationGatewayBackendHttpSetting']] = pulumi.output_property("backendHttpSettings")
    """
    One or more `backend_http_settings` blocks as defined below.
    """
    custom_error_configurations: pulumi.Output[Optional[List['outputs.ApplicationGatewayCustomErrorConfiguration']]] = pulumi.output_property("customErrorConfigurations")
    """
    One or more `custom_error_configuration` blocks as defined below.
    """
    enable_http2: pulumi.Output[Optional[bool]] = pulumi.output_property("enableHttp2")
    """
    Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
    """
    firewall_policy_id: pulumi.Output[Optional[str]] = pulumi.output_property("firewallPolicyId")
    """
    The ID of the Web Application Firewall Policy.
    """
    frontend_ip_configurations: pulumi.Output[List['outputs.ApplicationGatewayFrontendIpConfiguration']] = pulumi.output_property("frontendIpConfigurations")
    """
    One or more `frontend_ip_configuration` blocks as defined below.
    """
    frontend_ports: pulumi.Output[List['outputs.ApplicationGatewayFrontendPort']] = pulumi.output_property("frontendPorts")
    """
    One or more `frontend_port` blocks as defined below.
    """
    gateway_ip_configurations: pulumi.Output[List['outputs.ApplicationGatewayGatewayIpConfiguration']] = pulumi.output_property("gatewayIpConfigurations")
    """
    One or more `gateway_ip_configuration` blocks as defined below.
    """
    http_listeners: pulumi.Output[List['outputs.ApplicationGatewayHttpListener']] = pulumi.output_property("httpListeners")
    """
    One or more `http_listener` blocks as defined below.
    """
    identity: pulumi.Output[Optional['outputs.ApplicationGatewayIdentity']] = pulumi.output_property("identity")
    """
    A `identity` block.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the Application Gateway. Changing this forces a new resource to be created.
    """
    probes: pulumi.Output[Optional[List['outputs.ApplicationGatewayProbe']]] = pulumi.output_property("probes")
    """
    One or more `probe` blocks as defined below.
    """
    redirect_configurations: pulumi.Output[Optional[List['outputs.ApplicationGatewayRedirectConfiguration']]] = pulumi.output_property("redirectConfigurations")
    """
    A `redirect_configuration` block as defined below.
    """
    request_routing_rules: pulumi.Output[List['outputs.ApplicationGatewayRequestRoutingRule']] = pulumi.output_property("requestRoutingRules")
    """
    One or more `request_routing_rule` blocks as defined below.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
    """
    rewrite_rule_sets: pulumi.Output[Optional[List['outputs.ApplicationGatewayRewriteRuleSet']]] = pulumi.output_property("rewriteRuleSets")
    """
    One or more `rewrite_rule_set` blocks as defined below. Only valid for v2 SKUs.
    """
    sku: pulumi.Output['outputs.ApplicationGatewaySku'] = pulumi.output_property("sku")
    """
    A `sku` block as defined below.
    """
    ssl_certificates: pulumi.Output[Optional[List['outputs.ApplicationGatewaySslCertificate']]] = pulumi.output_property("sslCertificates")
    """
    One or more `ssl_certificate` blocks as defined below.
    """
    ssl_policies: pulumi.Output[List['outputs.ApplicationGatewaySslPolicy']] = pulumi.output_property("sslPolicies")
    """
    a `ssl policy` block as defined below.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    trusted_root_certificates: pulumi.Output[Optional[List['outputs.ApplicationGatewayTrustedRootCertificate']]] = pulumi.output_property("trustedRootCertificates")
    """
    One or more `trusted_root_certificate` blocks as defined below.
    """
    url_path_maps: pulumi.Output[Optional[List['outputs.ApplicationGatewayUrlPathMap']]] = pulumi.output_property("urlPathMaps")
    """
    One or more `url_path_map` blocks as defined below.
    """
    waf_configuration: pulumi.Output[Optional['outputs.ApplicationGatewayWafConfiguration']] = pulumi.output_property("wafConfiguration")
    """
    A `waf_configuration` block as defined below.
    """
    zones: pulumi.Output[Optional[List[str]]] = pulumi.output_property("zones")
    """
    A collection of availability zones to spread the Application Gateway over.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, authentication_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayAuthenticationCertificateArgs']]]]] = None, autoscale_configuration: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewayAutoscaleConfigurationArgs']]] = None, backend_address_pools: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendAddressPoolArgs']]]]] = None, backend_http_settings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendHttpSettingArgs']]]]] = None, custom_error_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayCustomErrorConfigurationArgs']]]]] = None, enable_http2: Optional[pulumi.Input[bool]] = None, firewall_policy_id: Optional[pulumi.Input[str]] = None, frontend_ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendIpConfigurationArgs']]]]] = None, frontend_ports: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendPortArgs']]]]] = None, gateway_ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayGatewayIpConfigurationArgs']]]]] = None, http_listeners: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayHttpListenerArgs']]]]] = None, identity: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewayIdentityArgs']]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, probes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayProbeArgs']]]]] = None, redirect_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRedirectConfigurationArgs']]]]] = None, request_routing_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRequestRoutingRuleArgs']]]]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, rewrite_rule_sets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRewriteRuleSetArgs']]]]] = None, sku: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewaySkuArgs']]] = None, ssl_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslCertificateArgs']]]]] = None, ssl_policies: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslPolicyArgs']]]]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, trusted_root_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayTrustedRootCertificateArgs']]]]] = None, url_path_maps: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayUrlPathMapArgs']]]]] = None, waf_configuration: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewayWafConfigurationArgs']]] = None, zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an Application Gateway.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            address_spaces=["10.254.0.0/16"])
        frontend = azure.network.Subnet("frontend",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.254.0.0/24"])
        backend = azure.network.Subnet("backend",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.254.2.0/24"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            allocation_method="Dynamic")
        backend_address_pool_name = example_virtual_network.name.apply(lambda name: f"{name}-beap")
        frontend_port_name = example_virtual_network.name.apply(lambda name: f"{name}-feport")
        frontend_ip_configuration_name = example_virtual_network.name.apply(lambda name: f"{name}-feip")
        http_setting_name = example_virtual_network.name.apply(lambda name: f"{name}-be-htst")
        listener_name = example_virtual_network.name.apply(lambda name: f"{name}-httplstn")
        request_routing_rule_name = example_virtual_network.name.apply(lambda name: f"{name}-rqrt")
        redirect_configuration_name = example_virtual_network.name.apply(lambda name: f"{name}-rdrcfg")
        network = azure.network.ApplicationGateway("network",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku={
                "name": "Standard_Small",
                "tier": "Standard",
                "capacity": 2,
            },
            gateway_ip_configurations=[{
                "name": "my-gateway-ip-configuration",
                "subnet_id": frontend.id,
            }],
            frontend_ports=[{
                "name": frontend_port_name,
                "port": 80,
            }],
            frontend_ip_configurations=[{
                "name": frontend_ip_configuration_name,
                "public_ip_address_id": example_public_ip.id,
            }],
            backend_address_pools=[{
                "name": backend_address_pool_name,
            }],
            backend_http_settings=[{
                "name": http_setting_name,
                "cookieBasedAffinity": "Disabled",
                "path": "/path1/",
                "port": 80,
                "protocol": "Http",
                "requestTimeout": 60,
            }],
            http_listeners=[{
                "name": listener_name,
                "frontend_ip_configuration_name": frontend_ip_configuration_name,
                "frontendPortName": frontend_port_name,
                "protocol": "Http",
            }],
            request_routing_rules=[{
                "name": request_routing_rule_name,
                "ruleType": "Basic",
                "httpListenerName": listener_name,
                "backendAddressPoolName": backend_address_pool_name,
                "backendHttpSettingsName": http_setting_name,
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayAuthenticationCertificateArgs']]]] authentication_certificates: One or more `authentication_certificate` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewayAutoscaleConfigurationArgs']] autoscale_configuration: A `autoscale_configuration` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendAddressPoolArgs']]]] backend_address_pools: One or more `backend_address_pool` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendHttpSettingArgs']]]] backend_http_settings: One or more `backend_http_settings` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayCustomErrorConfigurationArgs']]]] custom_error_configurations: One or more `custom_error_configuration` blocks as defined below.
        :param pulumi.Input[bool] enable_http2: Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
        :param pulumi.Input[str] firewall_policy_id: The ID of the Web Application Firewall Policy.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendIpConfigurationArgs']]]] frontend_ip_configurations: One or more `frontend_ip_configuration` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendPortArgs']]]] frontend_ports: One or more `frontend_port` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayGatewayIpConfigurationArgs']]]] gateway_ip_configurations: One or more `gateway_ip_configuration` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayHttpListenerArgs']]]] http_listeners: One or more `http_listener` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewayIdentityArgs']] identity: A `identity` block.
        :param pulumi.Input[str] location: The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Application Gateway. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayProbeArgs']]]] probes: One or more `probe` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRedirectConfigurationArgs']]]] redirect_configurations: A `redirect_configuration` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRequestRoutingRuleArgs']]]] request_routing_rules: One or more `request_routing_rule` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRewriteRuleSetArgs']]]] rewrite_rule_sets: One or more `rewrite_rule_set` blocks as defined below. Only valid for v2 SKUs.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewaySkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslCertificateArgs']]]] ssl_certificates: One or more `ssl_certificate` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslPolicyArgs']]]] ssl_policies: a `ssl policy` block as defined below.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayTrustedRootCertificateArgs']]]] trusted_root_certificates: One or more `trusted_root_certificate` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayUrlPathMapArgs']]]] url_path_maps: One or more `url_path_map` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewayWafConfigurationArgs']] waf_configuration: A `waf_configuration` block as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] zones: A collection of availability zones to spread the Application Gateway over.
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

            __props__['authentication_certificates'] = authentication_certificates
            __props__['autoscale_configuration'] = autoscale_configuration
            if backend_address_pools is None:
                raise TypeError("Missing required property 'backend_address_pools'")
            __props__['backend_address_pools'] = backend_address_pools
            if backend_http_settings is None:
                raise TypeError("Missing required property 'backend_http_settings'")
            __props__['backend_http_settings'] = backend_http_settings
            __props__['custom_error_configurations'] = custom_error_configurations
            __props__['enable_http2'] = enable_http2
            __props__['firewall_policy_id'] = firewall_policy_id
            if frontend_ip_configurations is None:
                raise TypeError("Missing required property 'frontend_ip_configurations'")
            __props__['frontend_ip_configurations'] = frontend_ip_configurations
            if frontend_ports is None:
                raise TypeError("Missing required property 'frontend_ports'")
            __props__['frontend_ports'] = frontend_ports
            if gateway_ip_configurations is None:
                raise TypeError("Missing required property 'gateway_ip_configurations'")
            __props__['gateway_ip_configurations'] = gateway_ip_configurations
            if http_listeners is None:
                raise TypeError("Missing required property 'http_listeners'")
            __props__['http_listeners'] = http_listeners
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['name'] = name
            __props__['probes'] = probes
            __props__['redirect_configurations'] = redirect_configurations
            if request_routing_rules is None:
                raise TypeError("Missing required property 'request_routing_rules'")
            __props__['request_routing_rules'] = request_routing_rules
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['rewrite_rule_sets'] = rewrite_rule_sets
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['ssl_certificates'] = ssl_certificates
            __props__['ssl_policies'] = ssl_policies
            __props__['tags'] = tags
            __props__['trusted_root_certificates'] = trusted_root_certificates
            __props__['url_path_maps'] = url_path_maps
            __props__['waf_configuration'] = waf_configuration
            __props__['zones'] = zones
        super(ApplicationGateway, __self__).__init__(
            'azure:network/applicationGateway:ApplicationGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, authentication_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayAuthenticationCertificateArgs']]]]] = None, autoscale_configuration: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewayAutoscaleConfigurationArgs']]] = None, backend_address_pools: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendAddressPoolArgs']]]]] = None, backend_http_settings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendHttpSettingArgs']]]]] = None, custom_error_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayCustomErrorConfigurationArgs']]]]] = None, enable_http2: Optional[pulumi.Input[bool]] = None, firewall_policy_id: Optional[pulumi.Input[str]] = None, frontend_ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendIpConfigurationArgs']]]]] = None, frontend_ports: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendPortArgs']]]]] = None, gateway_ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayGatewayIpConfigurationArgs']]]]] = None, http_listeners: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayHttpListenerArgs']]]]] = None, identity: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewayIdentityArgs']]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, probes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayProbeArgs']]]]] = None, redirect_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRedirectConfigurationArgs']]]]] = None, request_routing_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRequestRoutingRuleArgs']]]]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, rewrite_rule_sets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRewriteRuleSetArgs']]]]] = None, sku: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewaySkuArgs']]] = None, ssl_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslCertificateArgs']]]]] = None, ssl_policies: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslPolicyArgs']]]]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, trusted_root_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayTrustedRootCertificateArgs']]]]] = None, url_path_maps: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayUrlPathMapArgs']]]]] = None, waf_configuration: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewayWafConfigurationArgs']]] = None, zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'ApplicationGateway':
        """
        Get an existing ApplicationGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayAuthenticationCertificateArgs']]]] authentication_certificates: One or more `authentication_certificate` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewayAutoscaleConfigurationArgs']] autoscale_configuration: A `autoscale_configuration` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendAddressPoolArgs']]]] backend_address_pools: One or more `backend_address_pool` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendHttpSettingArgs']]]] backend_http_settings: One or more `backend_http_settings` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayCustomErrorConfigurationArgs']]]] custom_error_configurations: One or more `custom_error_configuration` blocks as defined below.
        :param pulumi.Input[bool] enable_http2: Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
        :param pulumi.Input[str] firewall_policy_id: The ID of the Web Application Firewall Policy.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendIpConfigurationArgs']]]] frontend_ip_configurations: One or more `frontend_ip_configuration` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendPortArgs']]]] frontend_ports: One or more `frontend_port` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayGatewayIpConfigurationArgs']]]] gateway_ip_configurations: One or more `gateway_ip_configuration` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayHttpListenerArgs']]]] http_listeners: One or more `http_listener` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewayIdentityArgs']] identity: A `identity` block.
        :param pulumi.Input[str] location: The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Application Gateway. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayProbeArgs']]]] probes: One or more `probe` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRedirectConfigurationArgs']]]] redirect_configurations: A `redirect_configuration` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRequestRoutingRuleArgs']]]] request_routing_rules: One or more `request_routing_rule` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRewriteRuleSetArgs']]]] rewrite_rule_sets: One or more `rewrite_rule_set` blocks as defined below. Only valid for v2 SKUs.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewaySkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslCertificateArgs']]]] ssl_certificates: One or more `ssl_certificate` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslPolicyArgs']]]] ssl_policies: a `ssl policy` block as defined below.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayTrustedRootCertificateArgs']]]] trusted_root_certificates: One or more `trusted_root_certificate` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayUrlPathMapArgs']]]] url_path_maps: One or more `url_path_map` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['ApplicationGatewayWafConfigurationArgs']] waf_configuration: A `waf_configuration` block as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] zones: A collection of availability zones to spread the Application Gateway over.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authentication_certificates"] = authentication_certificates
        __props__["autoscale_configuration"] = autoscale_configuration
        __props__["backend_address_pools"] = backend_address_pools
        __props__["backend_http_settings"] = backend_http_settings
        __props__["custom_error_configurations"] = custom_error_configurations
        __props__["enable_http2"] = enable_http2
        __props__["firewall_policy_id"] = firewall_policy_id
        __props__["frontend_ip_configurations"] = frontend_ip_configurations
        __props__["frontend_ports"] = frontend_ports
        __props__["gateway_ip_configurations"] = gateway_ip_configurations
        __props__["http_listeners"] = http_listeners
        __props__["identity"] = identity
        __props__["location"] = location
        __props__["name"] = name
        __props__["probes"] = probes
        __props__["redirect_configurations"] = redirect_configurations
        __props__["request_routing_rules"] = request_routing_rules
        __props__["resource_group_name"] = resource_group_name
        __props__["rewrite_rule_sets"] = rewrite_rule_sets
        __props__["sku"] = sku
        __props__["ssl_certificates"] = ssl_certificates
        __props__["ssl_policies"] = ssl_policies
        __props__["tags"] = tags
        __props__["trusted_root_certificates"] = trusted_root_certificates
        __props__["url_path_maps"] = url_path_maps
        __props__["waf_configuration"] = waf_configuration
        __props__["zones"] = zones
        return ApplicationGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

