# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VirtualNetwork(pulumi.CustomResource):
    address_spaces: pulumi.Output[list]
    """
    The address space that is used the virtual
    network. You can supply more than one address space. Changing this forces
    a new resource to be created.
    """
    ddos_protection_plan: pulumi.Output[dict]
    """
    A `ddos_protection_plan` block as documented below.

      * `enable` (`bool`) - Enable/disable DDoS Protection Plan on Virtual Network.
      * `id` (`str`) - The Resource ID of DDoS Protection Plan.
    """
    dns_servers: pulumi.Output[list]
    """
    List of IP addresses of DNS servers
    """
    guid: pulumi.Output[str]
    """
    The GUID of the virtual network.
    """
    location: pulumi.Output[str]
    """
    The location/region where the virtual network is
    created. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the virtual network. Changing this forces a
    new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the virtual network.
    """
    subnets: pulumi.Output[list]
    """
    Can be specified multiple times to define multiple
    subnets. Each `subnet` block supports fields documented below.

      * `address_prefix` (`str`) - The address prefix to use for the subnet.
      * `id` (`str`) - The Resource ID of DDoS Protection Plan.
      * `name` (`str`) - The name of the virtual network. Changing this forces a
        new resource to be created.
      * `securityGroup` (`str`) - The Network Security Group to associate with
        the subnet. (Referenced by `id`, ie. `azurerm_network_security_group.example.id`)
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, address_spaces=None, ddos_protection_plan=None, dns_servers=None, location=None, name=None, resource_group_name=None, subnets=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a virtual network including any configured subnets. Each subnet can
        optionally be configured with a security group to be associated with the subnet.

        > **NOTE on Virtual Networks and Subnet's:** This provider currently
        provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
        At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_network_security_group = azure.network.NetworkSecurityGroup("exampleNetworkSecurityGroup",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_ddos_protection_plan = azure.network.DdosProtectionPlan("exampleDdosProtectionPlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/16"],
            dns_servers=[
                "10.0.0.4",
                "10.0.0.5",
            ],
            ddos_protection_plan={
                "id": example_ddos_protection_plan.id,
                "enable": True,
            },
            subnets=[
                {
                    "name": "subnet1",
                    "address_prefix": "10.0.1.0/24",
                },
                {
                    "name": "subnet2",
                    "address_prefix": "10.0.2.0/24",
                },
                {
                    "name": "subnet3",
                    "address_prefix": "10.0.3.0/24",
                    "securityGroup": example_network_security_group.id,
                },
            ],
            tags={
                "environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] address_spaces: The address space that is used the virtual
               network. You can supply more than one address space. Changing this forces
               a new resource to be created.
        :param pulumi.Input[dict] ddos_protection_plan: A `ddos_protection_plan` block as documented below.
        :param pulumi.Input[list] dns_servers: List of IP addresses of DNS servers
        :param pulumi.Input[str] location: The location/region where the virtual network is
               created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the virtual network. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the virtual network.
        :param pulumi.Input[list] subnets: Can be specified multiple times to define multiple
               subnets. Each `subnet` block supports fields documented below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **ddos_protection_plan** object supports the following:

          * `enable` (`pulumi.Input[bool]`) - Enable/disable DDoS Protection Plan on Virtual Network.
          * `id` (`pulumi.Input[str]`) - The Resource ID of DDoS Protection Plan.

        The **subnets** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - The address prefix to use for the subnet.
          * `id` (`pulumi.Input[str]`) - The Resource ID of DDoS Protection Plan.
          * `name` (`pulumi.Input[str]`) - The name of the virtual network. Changing this forces a
            new resource to be created.
          * `securityGroup` (`pulumi.Input[str]`) - The Network Security Group to associate with
            the subnet. (Referenced by `id`, ie. `azurerm_network_security_group.example.id`)
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

            if address_spaces is None:
                raise TypeError("Missing required property 'address_spaces'")
            __props__['address_spaces'] = address_spaces
            __props__['ddos_protection_plan'] = ddos_protection_plan
            __props__['dns_servers'] = dns_servers
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['subnets'] = subnets
            __props__['tags'] = tags
            __props__['guid'] = None
        super(VirtualNetwork, __self__).__init__(
            'azure:network/virtualNetwork:VirtualNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address_spaces=None, ddos_protection_plan=None, dns_servers=None, guid=None, location=None, name=None, resource_group_name=None, subnets=None, tags=None):
        """
        Get an existing VirtualNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] address_spaces: The address space that is used the virtual
               network. You can supply more than one address space. Changing this forces
               a new resource to be created.
        :param pulumi.Input[dict] ddos_protection_plan: A `ddos_protection_plan` block as documented below.
        :param pulumi.Input[list] dns_servers: List of IP addresses of DNS servers
        :param pulumi.Input[str] guid: The GUID of the virtual network.
        :param pulumi.Input[str] location: The location/region where the virtual network is
               created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the virtual network. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the virtual network.
        :param pulumi.Input[list] subnets: Can be specified multiple times to define multiple
               subnets. Each `subnet` block supports fields documented below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **ddos_protection_plan** object supports the following:

          * `enable` (`pulumi.Input[bool]`) - Enable/disable DDoS Protection Plan on Virtual Network.
          * `id` (`pulumi.Input[str]`) - The Resource ID of DDoS Protection Plan.

        The **subnets** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - The address prefix to use for the subnet.
          * `id` (`pulumi.Input[str]`) - The Resource ID of DDoS Protection Plan.
          * `name` (`pulumi.Input[str]`) - The name of the virtual network. Changing this forces a
            new resource to be created.
          * `securityGroup` (`pulumi.Input[str]`) - The Network Security Group to associate with
            the subnet. (Referenced by `id`, ie. `azurerm_network_security_group.example.id`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_spaces"] = address_spaces
        __props__["ddos_protection_plan"] = ddos_protection_plan
        __props__["dns_servers"] = dns_servers
        __props__["guid"] = guid
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["subnets"] = subnets
        __props__["tags"] = tags
        return VirtualNetwork(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
