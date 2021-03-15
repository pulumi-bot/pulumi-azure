# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['NatPoolArgs', 'NatPool']

@pulumi.input_type
class NatPoolArgs:
    def __init__(__self__, *,
                 backend_port: pulumi.Input[int],
                 frontend_ip_configuration_name: pulumi.Input[str],
                 frontend_port_end: pulumi.Input[int],
                 frontend_port_start: pulumi.Input[int],
                 loadbalancer_id: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NatPool resource.
        :param pulumi.Input[int] backend_port: The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
        :param pulumi.Input[str] frontend_ip_configuration_name: The name of the frontend IP configuration exposing this rule.
        :param pulumi.Input[int] frontend_port_end: The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[int] frontend_port_start: The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the NAT pool.
        :param pulumi.Input[str] protocol: The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
        :param pulumi.Input[str] name: Specifies the name of the NAT pool.
        """
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "frontend_ip_configuration_name", frontend_ip_configuration_name)
        pulumi.set(__self__, "frontend_port_end", frontend_port_end)
        pulumi.set(__self__, "frontend_port_start", frontend_port_start)
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> pulumi.Input[int]:
        """
        The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
        """
        return pulumi.get(self, "backend_port")

    @backend_port.setter
    def backend_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "backend_port", value)

    @property
    @pulumi.getter(name="frontendIpConfigurationName")
    def frontend_ip_configuration_name(self) -> pulumi.Input[str]:
        """
        The name of the frontend IP configuration exposing this rule.
        """
        return pulumi.get(self, "frontend_ip_configuration_name")

    @frontend_ip_configuration_name.setter
    def frontend_ip_configuration_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "frontend_ip_configuration_name", value)

    @property
    @pulumi.getter(name="frontendPortEnd")
    def frontend_port_end(self) -> pulumi.Input[int]:
        """
        The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        """
        return pulumi.get(self, "frontend_port_end")

    @frontend_port_end.setter
    def frontend_port_end(self, value: pulumi.Input[int]):
        pulumi.set(self, "frontend_port_end", value)

    @property
    @pulumi.getter(name="frontendPortStart")
    def frontend_port_start(self) -> pulumi.Input[int]:
        """
        The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        """
        return pulumi.get(self, "frontend_port_start")

    @frontend_port_start.setter
    def frontend_port_start(self, value: pulumi.Input[int]):
        pulumi.set(self, "frontend_port_start", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Input[str]:
        """
        The ID of the Load Balancer in which to create the NAT pool.
        """
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the resource.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the NAT pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class NatPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NatPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Load Balancer NAT pool.

        > **NOTE:** This resource cannot be used with with virtual machines, instead use the `lb.NatRule` resource.

        > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location="West US",
            resource_group_name=example_resource_group.name,
            allocation_method="Static")
        example_load_balancer = azure.lb.LoadBalancer("exampleLoadBalancer",
            location="West US",
            resource_group_name=example_resource_group.name,
            frontend_ip_configurations=[azure.lb.LoadBalancerFrontendIpConfigurationArgs(
                name="PublicIPAddress",
                public_ip_address_id=example_public_ip.id,
            )])
        example_nat_pool = azure.lb.NatPool("exampleNatPool",
            resource_group_name=example_resource_group.name,
            loadbalancer_id=example_load_balancer.id,
            protocol="Tcp",
            frontend_port_start=80,
            frontend_port_end=81,
            backend_port=8080,
            frontend_ip_configuration_name="PublicIPAddress")
        ```

        ## Import

        Load Balancer NAT Pools can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:lb/natPool:NatPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1
        ```

        :param str resource_name: The name of the resource.
        :param NatPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_port: Optional[pulumi.Input[int]] = None,
                 frontend_ip_configuration_name: Optional[pulumi.Input[str]] = None,
                 frontend_port_end: Optional[pulumi.Input[int]] = None,
                 frontend_port_start: Optional[pulumi.Input[int]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Load Balancer NAT pool.

        > **NOTE:** This resource cannot be used with with virtual machines, instead use the `lb.NatRule` resource.

        > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location="West US",
            resource_group_name=example_resource_group.name,
            allocation_method="Static")
        example_load_balancer = azure.lb.LoadBalancer("exampleLoadBalancer",
            location="West US",
            resource_group_name=example_resource_group.name,
            frontend_ip_configurations=[azure.lb.LoadBalancerFrontendIpConfigurationArgs(
                name="PublicIPAddress",
                public_ip_address_id=example_public_ip.id,
            )])
        example_nat_pool = azure.lb.NatPool("exampleNatPool",
            resource_group_name=example_resource_group.name,
            loadbalancer_id=example_load_balancer.id,
            protocol="Tcp",
            frontend_port_start=80,
            frontend_port_end=81,
            backend_port=8080,
            frontend_ip_configuration_name="PublicIPAddress")
        ```

        ## Import

        Load Balancer NAT Pools can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:lb/natPool:NatPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] backend_port: The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
        :param pulumi.Input[str] frontend_ip_configuration_name: The name of the frontend IP configuration exposing this rule.
        :param pulumi.Input[int] frontend_port_end: The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[int] frontend_port_start: The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the NAT pool.
        :param pulumi.Input[str] name: Specifies the name of the NAT pool.
        :param pulumi.Input[str] protocol: The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NatPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_port: Optional[pulumi.Input[int]] = None,
                 frontend_ip_configuration_name: Optional[pulumi.Input[str]] = None,
                 frontend_port_end: Optional[pulumi.Input[int]] = None,
                 frontend_port_start: Optional[pulumi.Input[int]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
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

            if backend_port is None and not opts.urn:
                raise TypeError("Missing required property 'backend_port'")
            __props__['backend_port'] = backend_port
            if frontend_ip_configuration_name is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_ip_configuration_name'")
            __props__['frontend_ip_configuration_name'] = frontend_ip_configuration_name
            if frontend_port_end is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_port_end'")
            __props__['frontend_port_end'] = frontend_port_end
            if frontend_port_start is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_port_start'")
            __props__['frontend_port_start'] = frontend_port_start
            if loadbalancer_id is None and not opts.urn:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__['loadbalancer_id'] = loadbalancer_id
            __props__['name'] = name
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['frontend_ip_configuration_id'] = None
        super(NatPool, __self__).__init__(
            'azure:lb/natPool:NatPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_port: Optional[pulumi.Input[int]] = None,
            frontend_ip_configuration_id: Optional[pulumi.Input[str]] = None,
            frontend_ip_configuration_name: Optional[pulumi.Input[str]] = None,
            frontend_port_end: Optional[pulumi.Input[int]] = None,
            frontend_port_start: Optional[pulumi.Input[int]] = None,
            loadbalancer_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'NatPool':
        """
        Get an existing NatPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] backend_port: The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
        :param pulumi.Input[str] frontend_ip_configuration_name: The name of the frontend IP configuration exposing this rule.
        :param pulumi.Input[int] frontend_port_end: The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[int] frontend_port_start: The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the NAT pool.
        :param pulumi.Input[str] name: Specifies the name of the NAT pool.
        :param pulumi.Input[str] protocol: The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend_port"] = backend_port
        __props__["frontend_ip_configuration_id"] = frontend_ip_configuration_id
        __props__["frontend_ip_configuration_name"] = frontend_ip_configuration_name
        __props__["frontend_port_end"] = frontend_port_end
        __props__["frontend_port_start"] = frontend_port_start
        __props__["loadbalancer_id"] = loadbalancer_id
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["resource_group_name"] = resource_group_name
        return NatPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> pulumi.Output[int]:
        """
        The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
        """
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter(name="frontendIpConfigurationId")
    def frontend_ip_configuration_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "frontend_ip_configuration_id")

    @property
    @pulumi.getter(name="frontendIpConfigurationName")
    def frontend_ip_configuration_name(self) -> pulumi.Output[str]:
        """
        The name of the frontend IP configuration exposing this rule.
        """
        return pulumi.get(self, "frontend_ip_configuration_name")

    @property
    @pulumi.getter(name="frontendPortEnd")
    def frontend_port_end(self) -> pulumi.Output[int]:
        """
        The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        """
        return pulumi.get(self, "frontend_port_end")

    @property
    @pulumi.getter(name="frontendPortStart")
    def frontend_port_start(self) -> pulumi.Output[int]:
        """
        The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
        """
        return pulumi.get(self, "frontend_port_start")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[str]:
        """
        The ID of the Load Balancer in which to create the NAT pool.
        """
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the NAT pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the resource.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

