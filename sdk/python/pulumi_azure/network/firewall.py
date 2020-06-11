# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Firewall(pulumi.CustomResource):
    ip_configurations: pulumi.Output[list]
    """
    A `ip_configuration` block as documented below.

      * `name` (`str`) - Specifies the name of the IP Configuration.
      * `private_ip_address` (`str`) - The private IP address of the Azure Firewall.
      * `public_ip_address_id` (`str`) - The Resource ID of the Public IP Address associated with the firewall.
      * `subnet_id` (`str`) - Reference to the subnet associated with the IP Configuration.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Firewall. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    zones: pulumi.Output[list]
    """
    Specifies the availability zones in which the Azure Firewall should be created.
    """
    def __init__(__self__, resource_name, opts=None, ip_configurations=None, location=None, name=None, resource_group_name=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Firewall.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="North Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.1.0/24")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_firewall = azure.network.Firewall("exampleFirewall",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configuration=[{
                "name": "configuration",
                "subnet_id": example_subnet.id,
                "public_ip_address_id": example_public_ip.id,
            }])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] ip_configurations: A `ip_configuration` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[list] zones: Specifies the availability zones in which the Azure Firewall should be created.

        The **ip_configurations** object supports the following:

          * `name` (`pulumi.Input[str]`) - Specifies the name of the IP Configuration.
          * `private_ip_address` (`pulumi.Input[str]`) - The private IP address of the Azure Firewall.
          * `public_ip_address_id` (`pulumi.Input[str]`) - The Resource ID of the Public IP Address associated with the firewall.
          * `subnet_id` (`pulumi.Input[str]`) - Reference to the subnet associated with the IP Configuration.
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

            if ip_configurations is None:
                raise TypeError("Missing required property 'ip_configurations'")
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['zones'] = zones
        super(Firewall, __self__).__init__(
            'azure:network/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ip_configurations=None, location=None, name=None, resource_group_name=None, tags=None, zones=None):
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] ip_configurations: A `ip_configuration` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[list] zones: Specifies the availability zones in which the Azure Firewall should be created.

        The **ip_configurations** object supports the following:

          * `name` (`pulumi.Input[str]`) - Specifies the name of the IP Configuration.
          * `private_ip_address` (`pulumi.Input[str]`) - The private IP address of the Azure Firewall.
          * `public_ip_address_id` (`pulumi.Input[str]`) - The Resource ID of the Public IP Address associated with the firewall.
          * `subnet_id` (`pulumi.Input[str]`) - Reference to the subnet associated with the IP Configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ip_configurations"] = ip_configurations
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["zones"] = zones
        return Firewall(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
