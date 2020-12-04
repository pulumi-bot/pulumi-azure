# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VirtualHubIp']


class VirtualHubIp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_allocation_method: Optional[pulumi.Input[str]] = None,
                 public_ip_address_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 virtual_hub_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Virtual Hub IP.

        > **NOTE** Virtual Hub IP only supports Standard Virtual Hub without Virtual Wan.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_hub = azure.network.VirtualHub("exampleVirtualHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku="Standard")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Dynamic")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.5.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.5.1.0/24")
        example_virtual_hub_ip = azure.network.VirtualHubIp("exampleVirtualHubIp",
            virtual_hub_id=example_virtual_hub.id,
            private_ip_address="10.5.1.18",
            private_ip_allocation_method="Static",
            public_ip_address_id=example_public_ip.id,
            subnet_id=example_subnet.id)
        ```

        ## Import

        Virtual Hub IPs can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/virtualHubIp:VirtualHubIp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/ipConfigurations/ipConfig1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_ip_address: The private IP address of the IP configuration.
        :param pulumi.Input[str] private_ip_allocation_method: The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
        :param pulumi.Input[str] public_ip_address_id: The ID of the Public IP Address.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_hub_id: The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
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

            __props__['name'] = name
            __props__['private_ip_address'] = private_ip_address
            __props__['private_ip_allocation_method'] = private_ip_allocation_method
            __props__['public_ip_address_id'] = public_ip_address_id
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            if virtual_hub_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_hub_id'")
            __props__['virtual_hub_id'] = virtual_hub_id
        super(VirtualHubIp, __self__).__init__(
            'azure:network/virtualHubIp:VirtualHubIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_ip_address: Optional[pulumi.Input[str]] = None,
            private_ip_allocation_method: Optional[pulumi.Input[str]] = None,
            public_ip_address_id: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            virtual_hub_id: Optional[pulumi.Input[str]] = None) -> 'VirtualHubIp':
        """
        Get an existing VirtualHubIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_ip_address: The private IP address of the IP configuration.
        :param pulumi.Input[str] private_ip_allocation_method: The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
        :param pulumi.Input[str] public_ip_address_id: The ID of the Public IP Address.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_hub_id: The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["private_ip_address"] = private_ip_address
        __props__["private_ip_allocation_method"] = private_ip_allocation_method
        __props__["public_ip_address_id"] = public_ip_address_id
        __props__["subnet_id"] = subnet_id
        __props__["virtual_hub_id"] = virtual_hub_id
        return VirtualHubIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Virtual Hub IP. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> pulumi.Output[Optional[str]]:
        """
        The private IP address of the IP configuration.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="privateIpAllocationMethod")
    def private_ip_allocation_method(self) -> pulumi.Output[Optional[str]]:
        """
        The private IP address allocation method. Possible values are `Static` and `Dynamic` is allowed. Defaults to `Dynamic`.
        """
        return pulumi.get(self, "private_ip_allocation_method")

    @property
    @pulumi.getter(name="publicIpAddressId")
    def public_ip_address_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Public IP Address.
        """
        return pulumi.get(self, "public_ip_address_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="virtualHubId")
    def virtual_hub_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Hub within which this ip configuration should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_hub_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

