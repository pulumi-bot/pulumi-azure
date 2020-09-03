# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Probe']


class Probe(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interval_in_seconds: Optional[pulumi.Input[float]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number_of_probes: Optional[pulumi.Input[float]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 request_path: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a LoadBalancer Probe Resource.

        > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
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
        example_probe = azure.lb.Probe("exampleProbe",
            resource_group_name=example_resource_group.name,
            loadbalancer_id=example_load_balancer.id,
            port=22)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] interval_in_seconds: The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
        :param pulumi.Input[str] loadbalancer_id: The ID of the LoadBalancer in which to create the NAT Rule.
        :param pulumi.Input[str] name: Specifies the name of the Probe.
        :param pulumi.Input[float] number_of_probes: The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
        :param pulumi.Input[float] port: Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
        :param pulumi.Input[str] protocol: Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
        :param pulumi.Input[str] request_path: The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
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

            __props__['interval_in_seconds'] = interval_in_seconds
            if loadbalancer_id is None:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__['loadbalancer_id'] = loadbalancer_id
            __props__['name'] = name
            __props__['number_of_probes'] = number_of_probes
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            __props__['protocol'] = protocol
            __props__['request_path'] = request_path
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['load_balancer_rules'] = None
        super(Probe, __self__).__init__(
            'azure:lb/probe:Probe',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            interval_in_seconds: Optional[pulumi.Input[float]] = None,
            load_balancer_rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            loadbalancer_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            number_of_probes: Optional[pulumi.Input[float]] = None,
            port: Optional[pulumi.Input[float]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            request_path: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'Probe':
        """
        Get an existing Probe resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] interval_in_seconds: The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
        :param pulumi.Input[str] loadbalancer_id: The ID of the LoadBalancer in which to create the NAT Rule.
        :param pulumi.Input[str] name: Specifies the name of the Probe.
        :param pulumi.Input[float] number_of_probes: The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
        :param pulumi.Input[float] port: Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
        :param pulumi.Input[str] protocol: Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
        :param pulumi.Input[str] request_path: The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["interval_in_seconds"] = interval_in_seconds
        __props__["load_balancer_rules"] = load_balancer_rules
        __props__["loadbalancer_id"] = loadbalancer_id
        __props__["name"] = name
        __props__["number_of_probes"] = number_of_probes
        __props__["port"] = port
        __props__["protocol"] = protocol
        __props__["request_path"] = request_path
        __props__["resource_group_name"] = resource_group_name
        return Probe(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="intervalInSeconds")
    def interval_in_seconds(self) -> pulumi.Output[Optional[float]]:
        """
        The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
        """
        return pulumi.get(self, "interval_in_seconds")

    @property
    @pulumi.getter(name="loadBalancerRules")
    def load_balancer_rules(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "load_balancer_rules")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[str]:
        """
        The ID of the LoadBalancer in which to create the NAT Rule.
        """
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Probe.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfProbes")
    def number_of_probes(self) -> pulumi.Output[Optional[float]]:
        """
        The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
        """
        return pulumi.get(self, "number_of_probes")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[float]:
        """
        Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="requestPath")
    def request_path(self) -> pulumi.Output[Optional[str]]:
        """
        The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
        """
        return pulumi.get(self, "request_path")

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

