# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Environment']


class Environment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_user_ip_cidrs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 front_end_scale_factor: Optional[pulumi.Input[float]] = None,
                 internal_load_balancing_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pricing_tier: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_whitelisted_ip_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an App Service Environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westeurope")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/16"])
        ase = azure.network.Subnet("ase",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.1.0/24")
        gateway = azure.network.Subnet("gateway",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.2.0/24")
        example_environment = azure.appservice.Environment("exampleEnvironment",
            subnet_id=ase.id,
            pricing_tier="I2",
            front_end_scale_factor=10,
            internal_load_balancing_mode="Web, Publishing",
            allowed_user_ip_cidrs=[
                "11.22.33.44/32",
                "55.66.77.0/24",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_user_ip_cidrs: Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
        :param pulumi.Input[float] front_end_scale_factor: Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
        :param pulumi.Input[str] internal_load_balancing_mode: Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
        :param pulumi.Input[str] name: The name of the App Service Environment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pricing_tier: Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        :param pulumi.Input[str] subnet_id: The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
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

            __props__['allowed_user_ip_cidrs'] = allowed_user_ip_cidrs
            __props__['front_end_scale_factor'] = front_end_scale_factor
            __props__['internal_load_balancing_mode'] = internal_load_balancing_mode
            __props__['name'] = name
            __props__['pricing_tier'] = pricing_tier
            __props__['resource_group_name'] = resource_group_name
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            if user_whitelisted_ip_ranges is not None:
                warnings.warn("this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format", DeprecationWarning)
                pulumi.log.warn("user_whitelisted_ip_ranges is deprecated: this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format")
            __props__['user_whitelisted_ip_ranges'] = user_whitelisted_ip_ranges
            __props__['location'] = None
        super(Environment, __self__).__init__(
            'azure:appservice/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_user_ip_cidrs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            front_end_scale_factor: Optional[pulumi.Input[float]] = None,
            internal_load_balancing_mode: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pricing_tier: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            user_whitelisted_ip_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_user_ip_cidrs: Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
        :param pulumi.Input[float] front_end_scale_factor: Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
        :param pulumi.Input[str] internal_load_balancing_mode: Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
        :param pulumi.Input[str] location: The location where the App Service Environment exists.
        :param pulumi.Input[str] name: The name of the App Service Environment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pricing_tier: Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        :param pulumi.Input[str] subnet_id: The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_user_ip_cidrs"] = allowed_user_ip_cidrs
        __props__["front_end_scale_factor"] = front_end_scale_factor
        __props__["internal_load_balancing_mode"] = internal_load_balancing_mode
        __props__["location"] = location
        __props__["name"] = name
        __props__["pricing_tier"] = pricing_tier
        __props__["resource_group_name"] = resource_group_name
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["user_whitelisted_ip_ranges"] = user_whitelisted_ip_ranges
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedUserIpCidrs")
    def allowed_user_ip_cidrs(self) -> List[str]:
        """
        Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
        """
        return pulumi.get(self, "allowed_user_ip_cidrs")

    @property
    @pulumi.getter(name="frontEndScaleFactor")
    def front_end_scale_factor(self) -> Optional[float]:
        """
        Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
        """
        return pulumi.get(self, "front_end_scale_factor")

    @property
    @pulumi.getter(name="internalLoadBalancingMode")
    def internal_load_balancing_mode(self) -> Optional[str]:
        """
        Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
        """
        return pulumi.get(self, "internal_load_balancing_mode")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location where the App Service Environment exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the App Service Environment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pricingTier")
    def pricing_tier(self) -> Optional[str]:
        """
        Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
        """
        return pulumi.get(self, "pricing_tier")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userWhitelistedIpRanges")
    def user_whitelisted_ip_ranges(self) -> List[str]:
        return pulumi.get(self, "user_whitelisted_ip_ranges")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

