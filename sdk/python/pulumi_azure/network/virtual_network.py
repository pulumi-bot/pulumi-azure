# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VirtualNetworkArgs', 'VirtualNetwork']

@pulumi.input_type
class VirtualNetworkArgs:
    def __init__(__self__, *,
                 address_spaces: pulumi.Input[Sequence[pulumi.Input[str]]],
                 resource_group_name: pulumi.Input[str],
                 bgp_community: Optional[pulumi.Input[str]] = None,
                 ddos_protection_plan: Optional[pulumi.Input['VirtualNetworkDdosProtectionPlanArgs']] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_protection_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a VirtualNetwork resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_spaces: The address space that is used the virtual network. You can supply more than one address space.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the virtual network.
        :param pulumi.Input[str] bgp_community: The BGP community attribute in format `<as-number>:<community-value>`.
        :param pulumi.Input['VirtualNetworkDdosProtectionPlanArgs'] ddos_protection_plan: A `ddos_protection_plan` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: List of IP addresses of DNS servers
        :param pulumi.Input[str] location: The location/region where the virtual network is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the virtual network. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]] subnets: Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] vm_protection_enabled: Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        pulumi.set(__self__, "address_spaces", address_spaces)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if bgp_community is not None:
            pulumi.set(__self__, "bgp_community", bgp_community)
        if ddos_protection_plan is not None:
            pulumi.set(__self__, "ddos_protection_plan", ddos_protection_plan)
        if dns_servers is not None:
            pulumi.set(__self__, "dns_servers", dns_servers)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnets is not None:
            pulumi.set(__self__, "subnets", subnets)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vm_protection_enabled is not None:
            pulumi.set(__self__, "vm_protection_enabled", vm_protection_enabled)

    @property
    @pulumi.getter(name="addressSpaces")
    def address_spaces(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The address space that is used the virtual network. You can supply more than one address space.
        """
        return pulumi.get(self, "address_spaces")

    @address_spaces.setter
    def address_spaces(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "address_spaces", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the virtual network.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="bgpCommunity")
    def bgp_community(self) -> Optional[pulumi.Input[str]]:
        """
        The BGP community attribute in format `<as-number>:<community-value>`.
        """
        return pulumi.get(self, "bgp_community")

    @bgp_community.setter
    def bgp_community(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_community", value)

    @property
    @pulumi.getter(name="ddosProtectionPlan")
    def ddos_protection_plan(self) -> Optional[pulumi.Input['VirtualNetworkDdosProtectionPlanArgs']]:
        """
        A `ddos_protection_plan` block as documented below.
        """
        return pulumi.get(self, "ddos_protection_plan")

    @ddos_protection_plan.setter
    def ddos_protection_plan(self, value: Optional[pulumi.Input['VirtualNetworkDdosProtectionPlanArgs']]):
        pulumi.set(self, "ddos_protection_plan", value)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses of DNS servers
        """
        return pulumi.get(self, "dns_servers")

    @dns_servers.setter
    def dns_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_servers", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location/region where the virtual network is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the virtual network. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]]]:
        """
        Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        """
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]]]):
        pulumi.set(self, "subnets", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vmProtectionEnabled")
    def vm_protection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        return pulumi.get(self, "vm_protection_enabled")

    @vm_protection_enabled.setter
    def vm_protection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vm_protection_enabled", value)


@pulumi.input_type
class _VirtualNetworkState:
    def __init__(__self__, *,
                 address_spaces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bgp_community: Optional[pulumi.Input[str]] = None,
                 ddos_protection_plan: Optional[pulumi.Input['VirtualNetworkDdosProtectionPlanArgs']] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 guid: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_protection_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering VirtualNetwork resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_spaces: The address space that is used the virtual network. You can supply more than one address space.
        :param pulumi.Input[str] bgp_community: The BGP community attribute in format `<as-number>:<community-value>`.
        :param pulumi.Input['VirtualNetworkDdosProtectionPlanArgs'] ddos_protection_plan: A `ddos_protection_plan` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: List of IP addresses of DNS servers
        :param pulumi.Input[str] guid: The GUID of the virtual network.
        :param pulumi.Input[str] location: The location/region where the virtual network is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the virtual network. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the virtual network.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]] subnets: Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] vm_protection_enabled: Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        if address_spaces is not None:
            pulumi.set(__self__, "address_spaces", address_spaces)
        if bgp_community is not None:
            pulumi.set(__self__, "bgp_community", bgp_community)
        if ddos_protection_plan is not None:
            pulumi.set(__self__, "ddos_protection_plan", ddos_protection_plan)
        if dns_servers is not None:
            pulumi.set(__self__, "dns_servers", dns_servers)
        if guid is not None:
            pulumi.set(__self__, "guid", guid)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if subnets is not None:
            pulumi.set(__self__, "subnets", subnets)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vm_protection_enabled is not None:
            pulumi.set(__self__, "vm_protection_enabled", vm_protection_enabled)

    @property
    @pulumi.getter(name="addressSpaces")
    def address_spaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The address space that is used the virtual network. You can supply more than one address space.
        """
        return pulumi.get(self, "address_spaces")

    @address_spaces.setter
    def address_spaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_spaces", value)

    @property
    @pulumi.getter(name="bgpCommunity")
    def bgp_community(self) -> Optional[pulumi.Input[str]]:
        """
        The BGP community attribute in format `<as-number>:<community-value>`.
        """
        return pulumi.get(self, "bgp_community")

    @bgp_community.setter
    def bgp_community(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_community", value)

    @property
    @pulumi.getter(name="ddosProtectionPlan")
    def ddos_protection_plan(self) -> Optional[pulumi.Input['VirtualNetworkDdosProtectionPlanArgs']]:
        """
        A `ddos_protection_plan` block as documented below.
        """
        return pulumi.get(self, "ddos_protection_plan")

    @ddos_protection_plan.setter
    def ddos_protection_plan(self, value: Optional[pulumi.Input['VirtualNetworkDdosProtectionPlanArgs']]):
        pulumi.set(self, "ddos_protection_plan", value)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses of DNS servers
        """
        return pulumi.get(self, "dns_servers")

    @dns_servers.setter
    def dns_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_servers", value)

    @property
    @pulumi.getter
    def guid(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the virtual network.
        """
        return pulumi.get(self, "guid")

    @guid.setter
    def guid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "guid", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location/region where the virtual network is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the virtual network. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the virtual network.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]]]:
        """
        Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        """
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNetworkSubnetArgs']]]]):
        pulumi.set(self, "subnets", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vmProtectionEnabled")
    def vm_protection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        return pulumi.get(self, "vm_protection_enabled")

    @vm_protection_enabled.setter
    def vm_protection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vm_protection_enabled", value)


class VirtualNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_spaces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bgp_community: Optional[pulumi.Input[str]] = None,
                 ddos_protection_plan: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkDdosProtectionPlanArgs']]] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_protection_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
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

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
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
            ddos_protection_plan=azure.network.VirtualNetworkDdosProtectionPlanArgs(
                id=example_ddos_protection_plan.id,
                enable=True,
            ),
            subnets=[
                azure.network.VirtualNetworkSubnetArgs(
                    name="subnet1",
                    address_prefix="10.0.1.0/24",
                ),
                azure.network.VirtualNetworkSubnetArgs(
                    name="subnet2",
                    address_prefix="10.0.2.0/24",
                ),
                azure.network.VirtualNetworkSubnetArgs(
                    name="subnet3",
                    address_prefix="10.0.3.0/24",
                    security_group=example_network_security_group.id,
                ),
            ],
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        Virtual Networks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/virtualNetwork:VirtualNetwork exampleNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_spaces: The address space that is used the virtual network. You can supply more than one address space.
        :param pulumi.Input[str] bgp_community: The BGP community attribute in format `<as-number>:<community-value>`.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkDdosProtectionPlanArgs']] ddos_protection_plan: A `ddos_protection_plan` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: List of IP addresses of DNS servers
        :param pulumi.Input[str] location: The location/region where the virtual network is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the virtual network. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the virtual network.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]]] subnets: Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] vm_protection_enabled: Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
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
            ddos_protection_plan=azure.network.VirtualNetworkDdosProtectionPlanArgs(
                id=example_ddos_protection_plan.id,
                enable=True,
            ),
            subnets=[
                azure.network.VirtualNetworkSubnetArgs(
                    name="subnet1",
                    address_prefix="10.0.1.0/24",
                ),
                azure.network.VirtualNetworkSubnetArgs(
                    name="subnet2",
                    address_prefix="10.0.2.0/24",
                ),
                azure.network.VirtualNetworkSubnetArgs(
                    name="subnet3",
                    address_prefix="10.0.3.0/24",
                    security_group=example_network_security_group.id,
                ),
            ],
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        Virtual Networks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/virtualNetwork:VirtualNetwork exampleNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
        ```

        :param str resource_name: The name of the resource.
        :param VirtualNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_spaces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bgp_community: Optional[pulumi.Input[str]] = None,
                 ddos_protection_plan: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkDdosProtectionPlanArgs']]] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_protection_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualNetworkArgs.__new__(VirtualNetworkArgs)

            if address_spaces is None and not opts.urn:
                raise TypeError("Missing required property 'address_spaces'")
            __props__.__dict__["address_spaces"] = address_spaces
            __props__.__dict__["bgp_community"] = bgp_community
            __props__.__dict__["ddos_protection_plan"] = ddos_protection_plan
            __props__.__dict__["dns_servers"] = dns_servers
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["subnets"] = subnets
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vm_protection_enabled"] = vm_protection_enabled
            __props__.__dict__["guid"] = None
        super(VirtualNetwork, __self__).__init__(
            'azure:network/virtualNetwork:VirtualNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_spaces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bgp_community: Optional[pulumi.Input[str]] = None,
            ddos_protection_plan: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkDdosProtectionPlanArgs']]] = None,
            dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            guid: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vm_protection_enabled: Optional[pulumi.Input[bool]] = None) -> 'VirtualNetwork':
        """
        Get an existing VirtualNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_spaces: The address space that is used the virtual network. You can supply more than one address space.
        :param pulumi.Input[str] bgp_community: The BGP community attribute in format `<as-number>:<community-value>`.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkDdosProtectionPlanArgs']] ddos_protection_plan: A `ddos_protection_plan` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: List of IP addresses of DNS servers
        :param pulumi.Input[str] guid: The GUID of the virtual network.
        :param pulumi.Input[str] location: The location/region where the virtual network is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the virtual network. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the virtual network.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]]] subnets: Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] vm_protection_enabled: Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualNetworkState.__new__(_VirtualNetworkState)

        __props__.__dict__["address_spaces"] = address_spaces
        __props__.__dict__["bgp_community"] = bgp_community
        __props__.__dict__["ddos_protection_plan"] = ddos_protection_plan
        __props__.__dict__["dns_servers"] = dns_servers
        __props__.__dict__["guid"] = guid
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["subnets"] = subnets
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vm_protection_enabled"] = vm_protection_enabled
        return VirtualNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressSpaces")
    def address_spaces(self) -> pulumi.Output[Sequence[str]]:
        """
        The address space that is used the virtual network. You can supply more than one address space.
        """
        return pulumi.get(self, "address_spaces")

    @property
    @pulumi.getter(name="bgpCommunity")
    def bgp_community(self) -> pulumi.Output[Optional[str]]:
        """
        The BGP community attribute in format `<as-number>:<community-value>`.
        """
        return pulumi.get(self, "bgp_community")

    @property
    @pulumi.getter(name="ddosProtectionPlan")
    def ddos_protection_plan(self) -> pulumi.Output[Optional['outputs.VirtualNetworkDdosProtectionPlan']]:
        """
        A `ddos_protection_plan` block as documented below.
        """
        return pulumi.get(self, "ddos_protection_plan")

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of IP addresses of DNS servers
        """
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter
    def guid(self) -> pulumi.Output[str]:
        """
        The GUID of the virtual network.
        """
        return pulumi.get(self, "guid")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location/region where the virtual network is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the virtual network. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the virtual network.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Output[Sequence['outputs.VirtualNetworkSubnet']]:
        """
        Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vmProtectionEnabled")
    def vm_protection_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
        """
        return pulumi.get(self, "vm_protection_enabled")

