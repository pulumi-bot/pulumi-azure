# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class NatGateway(pulumi.CustomResource):
    idle_timeout_in_minutes: pulumi.Output[float]
    """
    The idle timeout which should be used in minutes. Defaults to `4`.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the NAT Gateway. Changing this forces a new resource to be created.
    """
    public_ip_address_ids: pulumi.Output[list]
    """
    A list of Public IP Address ID's which should be associated with the NAT Gateway resource.
    """
    public_ip_prefix_ids: pulumi.Output[list]
    """
    A list of Public IP Prefix ID's which should be associated with the NAT Gateway resource.
    """
    resource_group_name: pulumi.Output[str]
    """
    Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
    """
    resource_guid: pulumi.Output[str]
    """
    The resource GUID property of the NAT Gateway.
    """
    sku_name: pulumi.Output[str]
    """
    The SKU which should be used. At this time the only supported value is `Standard`. Defaults to `Standard`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
    """
    zones: pulumi.Output[list]
    """
    A list of availability zones where the NAT Gateway should be provisioned. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, idle_timeout_in_minutes=None, location=None, name=None, public_ip_address_ids=None, public_ip_prefix_ids=None, resource_group_name=None, sku_name=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Azure NAT Gateway.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="eastus2")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard",
            zones=["1"])
        example_public_ip_prefix = azure.network.PublicIpPrefix("examplePublicIpPrefix",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            prefix_length=30,
            zones=["1"])
        example_nat_gateway = azure.network.NatGateway("exampleNatGateway",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            public_ip_address_ids=[example_public_ip.id],
            public_ip_prefix_ids=[example_public_ip_prefix.id],
            sku_name="Standard",
            idle_timeout_in_minutes=10,
            zones=["1"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] idle_timeout_in_minutes: The idle timeout which should be used in minutes. Defaults to `4`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the NAT Gateway. Changing this forces a new resource to be created.
        :param pulumi.Input[list] public_ip_address_ids: A list of Public IP Address ID's which should be associated with the NAT Gateway resource.
        :param pulumi.Input[list] public_ip_prefix_ids: A list of Public IP Prefix ID's which should be associated with the NAT Gateway resource.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: The SKU which should be used. At this time the only supported value is `Standard`. Defaults to `Standard`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[list] zones: A list of availability zones where the NAT Gateway should be provisioned. Changing this forces a new resource to be created.
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

            __props__['idle_timeout_in_minutes'] = idle_timeout_in_minutes
            __props__['location'] = location
            __props__['name'] = name
            if public_ip_address_ids is not None:
                warnings.warn("Inline Public IP Address ID Deprecations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_association` resource. This field will be removed in the next major version of the Azure Provider.", DeprecationWarning)
                pulumi.log.warn("public_ip_address_ids is deprecated: Inline Public IP Address ID Deprecations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_association` resource. This field will be removed in the next major version of the Azure Provider.")
            __props__['public_ip_address_ids'] = public_ip_address_ids
            __props__['public_ip_prefix_ids'] = public_ip_prefix_ids
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['resource_guid'] = None
        super(NatGateway, __self__).__init__(
            'azure:network/natGateway:NatGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, idle_timeout_in_minutes=None, location=None, name=None, public_ip_address_ids=None, public_ip_prefix_ids=None, resource_group_name=None, resource_guid=None, sku_name=None, tags=None, zones=None):
        """
        Get an existing NatGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] idle_timeout_in_minutes: The idle timeout which should be used in minutes. Defaults to `4`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the NAT Gateway. Changing this forces a new resource to be created.
        :param pulumi.Input[list] public_ip_address_ids: A list of Public IP Address ID's which should be associated with the NAT Gateway resource.
        :param pulumi.Input[list] public_ip_prefix_ids: A list of Public IP Prefix ID's which should be associated with the NAT Gateway resource.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the NAT Gateway.
        :param pulumi.Input[str] sku_name: The SKU which should be used. At this time the only supported value is `Standard`. Defaults to `Standard`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[list] zones: A list of availability zones where the NAT Gateway should be provisioned. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["idle_timeout_in_minutes"] = idle_timeout_in_minutes
        __props__["location"] = location
        __props__["name"] = name
        __props__["public_ip_address_ids"] = public_ip_address_ids
        __props__["public_ip_prefix_ids"] = public_ip_prefix_ids
        __props__["resource_group_name"] = resource_group_name
        __props__["resource_guid"] = resource_guid
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["zones"] = zones
        return NatGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
