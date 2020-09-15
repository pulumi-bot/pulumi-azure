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

__all__ = ['Endpoint']


class Endpoint(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_types_to_compresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 delivery_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointDeliveryRuleArgs']]]]] = None,
                 geo_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGeoFilterArgs']]]]] = None,
                 global_delivery_rule: Optional[pulumi.Input[pulumi.InputType['EndpointGlobalDeliveryRuleArgs']]] = None,
                 is_compression_enabled: Optional[pulumi.Input[bool]] = None,
                 is_http_allowed: Optional[pulumi.Input[bool]] = None,
                 is_https_allowed: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 optimization_type: Optional[pulumi.Input[str]] = None,
                 origin_host_header: Optional[pulumi.Input[str]] = None,
                 origin_path: Optional[pulumi.Input[str]] = None,
                 origins: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointOriginArgs']]]]] = None,
                 probe_path: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 querystring_caching_behaviour: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A CDN Endpoint is the entity within a CDN Profile containing configuration information regarding caching behaviours and origins. The CDN Endpoint is exposed using the URL format <endpointname>.azureedge.net.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_profile = azure.cdn.Profile("exampleProfile",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard_Verizon")
        example_endpoint = azure.cdn.Endpoint("exampleEndpoint",
            profile_name=example_profile.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            origins=[azure.cdn.EndpointOriginArgs(
                name="example",
                host_name="www.contoso.com",
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] content_types_to_compresses: An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointDeliveryRuleArgs']]]] delivery_rules: Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `delivery_rule` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGeoFilterArgs']]]] geo_filters: A set of Geo Filters for this CDN Endpoint. Each `geo_filter` block supports fields documented below.
        :param pulumi.Input[pulumi.InputType['EndpointGlobalDeliveryRuleArgs']] global_delivery_rule: Actions that are valid for all resources regardless of any conditions. A `global_delivery_rule` block as defined below.
        :param pulumi.Input[bool] is_compression_enabled: Indicates whether compression is to be enabled. Defaults to false.
        :param pulumi.Input[bool] is_http_allowed: Defaults to `true`.
        :param pulumi.Input[bool] is_https_allowed: Defaults to `true`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        :param pulumi.Input[str] optimization_type: What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
        :param pulumi.Input[str] origin_host_header: The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
        :param pulumi.Input[str] origin_path: The path used at for origin requests.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointOriginArgs']]]] origins: The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
        :param pulumi.Input[str] probe_path: the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `origin_path`.
        :param pulumi.Input[str] profile_name: The CDN Profile to which to attach the CDN Endpoint.
        :param pulumi.Input[str] querystring_caching_behaviour: Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. Defaults to `IgnoreQueryString`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the CDN Endpoint.
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

            __props__['content_types_to_compresses'] = content_types_to_compresses
            __props__['delivery_rules'] = delivery_rules
            __props__['geo_filters'] = geo_filters
            __props__['global_delivery_rule'] = global_delivery_rule
            __props__['is_compression_enabled'] = is_compression_enabled
            __props__['is_http_allowed'] = is_http_allowed
            __props__['is_https_allowed'] = is_https_allowed
            __props__['location'] = location
            __props__['name'] = name
            __props__['optimization_type'] = optimization_type
            __props__['origin_host_header'] = origin_host_header
            __props__['origin_path'] = origin_path
            if origins is None:
                raise TypeError("Missing required property 'origins'")
            __props__['origins'] = origins
            __props__['probe_path'] = probe_path
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['querystring_caching_behaviour'] = querystring_caching_behaviour
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['host_name'] = None
        super(Endpoint, __self__).__init__(
            'azure:cdn/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            content_types_to_compresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            delivery_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointDeliveryRuleArgs']]]]] = None,
            geo_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGeoFilterArgs']]]]] = None,
            global_delivery_rule: Optional[pulumi.Input[pulumi.InputType['EndpointGlobalDeliveryRuleArgs']]] = None,
            host_name: Optional[pulumi.Input[str]] = None,
            is_compression_enabled: Optional[pulumi.Input[bool]] = None,
            is_http_allowed: Optional[pulumi.Input[bool]] = None,
            is_https_allowed: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            optimization_type: Optional[pulumi.Input[str]] = None,
            origin_host_header: Optional[pulumi.Input[str]] = None,
            origin_path: Optional[pulumi.Input[str]] = None,
            origins: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointOriginArgs']]]]] = None,
            probe_path: Optional[pulumi.Input[str]] = None,
            profile_name: Optional[pulumi.Input[str]] = None,
            querystring_caching_behaviour: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] content_types_to_compresses: An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointDeliveryRuleArgs']]]] delivery_rules: Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `delivery_rule` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGeoFilterArgs']]]] geo_filters: A set of Geo Filters for this CDN Endpoint. Each `geo_filter` block supports fields documented below.
        :param pulumi.Input[pulumi.InputType['EndpointGlobalDeliveryRuleArgs']] global_delivery_rule: Actions that are valid for all resources regardless of any conditions. A `global_delivery_rule` block as defined below.
        :param pulumi.Input[str] host_name: A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] is_compression_enabled: Indicates whether compression is to be enabled. Defaults to false.
        :param pulumi.Input[bool] is_http_allowed: Defaults to `true`.
        :param pulumi.Input[bool] is_https_allowed: Defaults to `true`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        :param pulumi.Input[str] optimization_type: What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
        :param pulumi.Input[str] origin_host_header: The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
        :param pulumi.Input[str] origin_path: The path used at for origin requests.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointOriginArgs']]]] origins: The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
        :param pulumi.Input[str] probe_path: the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `origin_path`.
        :param pulumi.Input[str] profile_name: The CDN Profile to which to attach the CDN Endpoint.
        :param pulumi.Input[str] querystring_caching_behaviour: Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. Defaults to `IgnoreQueryString`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the CDN Endpoint.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_types_to_compresses"] = content_types_to_compresses
        __props__["delivery_rules"] = delivery_rules
        __props__["geo_filters"] = geo_filters
        __props__["global_delivery_rule"] = global_delivery_rule
        __props__["host_name"] = host_name
        __props__["is_compression_enabled"] = is_compression_enabled
        __props__["is_http_allowed"] = is_http_allowed
        __props__["is_https_allowed"] = is_https_allowed
        __props__["location"] = location
        __props__["name"] = name
        __props__["optimization_type"] = optimization_type
        __props__["origin_host_header"] = origin_host_header
        __props__["origin_path"] = origin_path
        __props__["origins"] = origins
        __props__["probe_path"] = probe_path
        __props__["profile_name"] = profile_name
        __props__["querystring_caching_behaviour"] = querystring_caching_behaviour
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentTypesToCompresses")
    def content_types_to_compresses(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
        """
        return pulumi.get(self, "content_types_to_compresses")

    @property
    @pulumi.getter(name="deliveryRules")
    def delivery_rules(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointDeliveryRule']]]:
        """
        Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `delivery_rule` blocks as defined below.
        """
        return pulumi.get(self, "delivery_rules")

    @property
    @pulumi.getter(name="geoFilters")
    def geo_filters(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointGeoFilter']]]:
        """
        A set of Geo Filters for this CDN Endpoint. Each `geo_filter` block supports fields documented below.
        """
        return pulumi.get(self, "geo_filters")

    @property
    @pulumi.getter(name="globalDeliveryRule")
    def global_delivery_rule(self) -> pulumi.Output[Optional['outputs.EndpointGlobalDeliveryRule']]:
        """
        Actions that are valid for all resources regardless of any conditions. A `global_delivery_rule` block as defined below.
        """
        return pulumi.get(self, "global_delivery_rule")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Output[str]:
        """
        A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="isCompressionEnabled")
    def is_compression_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether compression is to be enabled. Defaults to false.
        """
        return pulumi.get(self, "is_compression_enabled")

    @property
    @pulumi.getter(name="isHttpAllowed")
    def is_http_allowed(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "is_http_allowed")

    @property
    @pulumi.getter(name="isHttpsAllowed")
    def is_https_allowed(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "is_https_allowed")

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
        The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="optimizationType")
    def optimization_type(self) -> pulumi.Output[Optional[str]]:
        """
        What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
        """
        return pulumi.get(self, "optimization_type")

    @property
    @pulumi.getter(name="originHostHeader")
    def origin_host_header(self) -> pulumi.Output[Optional[str]]:
        """
        The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
        """
        return pulumi.get(self, "origin_host_header")

    @property
    @pulumi.getter(name="originPath")
    def origin_path(self) -> pulumi.Output[str]:
        """
        The path used at for origin requests.
        """
        return pulumi.get(self, "origin_path")

    @property
    @pulumi.getter
    def origins(self) -> pulumi.Output[Sequence['outputs.EndpointOrigin']]:
        """
        The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
        """
        return pulumi.get(self, "origins")

    @property
    @pulumi.getter(name="probePath")
    def probe_path(self) -> pulumi.Output[str]:
        """
        the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `origin_path`.
        """
        return pulumi.get(self, "probe_path")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Output[str]:
        """
        The CDN Profile to which to attach the CDN Endpoint.
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="querystringCachingBehaviour")
    def querystring_caching_behaviour(self) -> pulumi.Output[Optional[str]]:
        """
        Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. Defaults to `IgnoreQueryString`.
        """
        return pulumi.get(self, "querystring_caching_behaviour")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the CDN Endpoint.
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

