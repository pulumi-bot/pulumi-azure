# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Environment(pulumi.CustomResource):
    front_end_scale_factor: pulumi.Output[float]
    """
    Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
    """
    internal_load_balancing_mode: pulumi.Output[str]
    location: pulumi.Output[str]
    """
    The location where the App Service Environment exists.
    """
    name: pulumi.Output[str]
    """
    The name of the App Service Environment. Changing this forces a new resource to be created. 
    """
    pricing_tier: pulumi.Output[str]
    """
    Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
    """
    subnet_id: pulumi.Output[str]
    """
    The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource. Changing this forces a new resource to be created. 
    """
    def __init__(__self__, resource_name, opts=None, front_end_scale_factor=None, internal_load_balancing_mode=None, name=None, pricing_tier=None, resource_group_name=None, subnet_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
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
            front_end_scale_factor=10)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] front_end_scale_factor: Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
        :param pulumi.Input[str] name: The name of the App Service Environment. Changing this forces a new resource to be created. 
        :param pulumi.Input[str] pricing_tier: Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        :param pulumi.Input[str] subnet_id: The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created. 
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

            __props__['front_end_scale_factor'] = front_end_scale_factor
            __props__['internal_load_balancing_mode'] = internal_load_balancing_mode
            __props__['name'] = name
            __props__['pricing_tier'] = pricing_tier
            __props__['resource_group_name'] = resource_group_name
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['location'] = None
        super(Environment, __self__).__init__(
            'azure:appservice/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, front_end_scale_factor=None, internal_load_balancing_mode=None, location=None, name=None, pricing_tier=None, resource_group_name=None, subnet_id=None, tags=None):
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] front_end_scale_factor: Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
        :param pulumi.Input[str] location: The location where the App Service Environment exists.
        :param pulumi.Input[str] name: The name of the App Service Environment. Changing this forces a new resource to be created. 
        :param pulumi.Input[str] pricing_tier: Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        :param pulumi.Input[str] subnet_id: The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created. 
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["front_end_scale_factor"] = front_end_scale_factor
        __props__["internal_load_balancing_mode"] = internal_load_balancing_mode
        __props__["location"] = location
        __props__["name"] = name
        __props__["pricing_tier"] = pricing_tier
        __props__["resource_group_name"] = resource_group_name
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        return Environment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

