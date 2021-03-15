# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['RouteArgs', 'Route']

@pulumi.input_type
class RouteArgs:
    def __init__(__self__, *,
                 address_prefix: pulumi.Input[str],
                 next_hop_type: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 route_table_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 next_hop_in_ip_address: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Route resource.
        :param pulumi.Input[str] address_prefix: The destination CIDR to which the route applies, such as `10.1.0.0/16`
        :param pulumi.Input[str] next_hop_type: The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_table_name: The name of the route table within which create the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] next_hop_in_ip_address: Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        """
        pulumi.set(__self__, "address_prefix", address_prefix)
        pulumi.set(__self__, "next_hop_type", next_hop_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "route_table_name", route_table_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if next_hop_in_ip_address is not None:
            pulumi.set(__self__, "next_hop_in_ip_address", next_hop_in_ip_address)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> pulumi.Input[str]:
        """
        The destination CIDR to which the route applies, such as `10.1.0.0/16`
        """
        return pulumi.get(self, "address_prefix")

    @address_prefix.setter
    def address_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_prefix", value)

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> pulumi.Input[str]:
        """
        The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        """
        return pulumi.get(self, "next_hop_type")

    @next_hop_type.setter
    def next_hop_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_hop_type", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> pulumi.Input[str]:
        """
        The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "route_table_name")

    @route_table_name.setter
    def route_table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_table_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nextHopInIpAddress")
    def next_hop_in_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        """
        return pulumi.get(self, "next_hop_in_ip_address")

    @next_hop_in_ip_address.setter
    def next_hop_in_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop_in_ip_address", value)


class Route(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Route within a Route Table.

        > **NOTE on Route Tables and Routes:** This provider currently
        provides both a standalone Route resource, and allows for Routes to be defined in-line within the Route Table resource.
        At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_route_table = azure.network.RouteTable("exampleRouteTable",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_route = azure.network.Route("exampleRoute",
            resource_group_name=example_resource_group.name,
            route_table_name=example_route_table.name,
            address_prefix="10.1.0.0/16",
            next_hop_type="vnetlocal")
        ```

        ## Import

        Routes can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/route:Route exampleRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1
        ```

        :param str resource_name: The name of the resource.
        :param RouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 next_hop_in_ip_address: Optional[pulumi.Input[str]] = None,
                 next_hop_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Route within a Route Table.

        > **NOTE on Route Tables and Routes:** This provider currently
        provides both a standalone Route resource, and allows for Routes to be defined in-line within the Route Table resource.
        At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_route_table = azure.network.RouteTable("exampleRouteTable",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_route = azure.network.Route("exampleRoute",
            resource_group_name=example_resource_group.name,
            route_table_name=example_route_table.name,
            address_prefix="10.1.0.0/16",
            next_hop_type="vnetlocal")
        ```

        ## Import

        Routes can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/route:Route exampleRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The destination CIDR to which the route applies, such as `10.1.0.0/16`
        :param pulumi.Input[str] name: The name of the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] next_hop_in_ip_address: Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        :param pulumi.Input[str] next_hop_type: The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_table_name: The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 next_hop_in_ip_address: Optional[pulumi.Input[str]] = None,
                 next_hop_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if address_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'address_prefix'")
            __props__['address_prefix'] = address_prefix
            __props__['name'] = name
            __props__['next_hop_in_ip_address'] = next_hop_in_ip_address
            if next_hop_type is None and not opts.urn:
                raise TypeError("Missing required property 'next_hop_type'")
            __props__['next_hop_type'] = next_hop_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if route_table_name is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_name'")
            __props__['route_table_name'] = route_table_name
        super(Route, __self__).__init__(
            'azure:network/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_prefix: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            next_hop_in_ip_address: Optional[pulumi.Input[str]] = None,
            next_hop_type: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            route_table_name: Optional[pulumi.Input[str]] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The destination CIDR to which the route applies, such as `10.1.0.0/16`
        :param pulumi.Input[str] name: The name of the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] next_hop_in_ip_address: Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        :param pulumi.Input[str] next_hop_type: The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_table_name: The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_prefix"] = address_prefix
        __props__["name"] = name
        __props__["next_hop_in_ip_address"] = next_hop_in_ip_address
        __props__["next_hop_type"] = next_hop_type
        __props__["resource_group_name"] = resource_group_name
        __props__["route_table_name"] = route_table_name
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> pulumi.Output[str]:
        """
        The destination CIDR to which the route applies, such as `10.1.0.0/16`
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the route. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextHopInIpAddress")
    def next_hop_in_ip_address(self) -> pulumi.Output[Optional[str]]:
        """
        Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        """
        return pulumi.get(self, "next_hop_in_ip_address")

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> pulumi.Output[str]:
        """
        The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        """
        return pulumi.get(self, "next_hop_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> pulumi.Output[str]:
        """
        The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "route_table_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

