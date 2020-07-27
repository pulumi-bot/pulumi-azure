# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['NetworkSecurityRule']


class NetworkSecurityRule(pulumi.CustomResource):
    access: pulumi.Output[str] = pulumi.output_property("access")
    """
    Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    A description for this rule. Restricted to 140 characters.
    """
    destination_address_prefix: pulumi.Output[Optional[str]] = pulumi.output_property("destinationAddressPrefix")
    """
    CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destination_address_prefixes` is not specified.
    """
    destination_address_prefixes: pulumi.Output[Optional[List[str]]] = pulumi.output_property("destinationAddressPrefixes")
    """
    List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
    """
    destination_application_security_group_ids: pulumi.Output[Optional[str]] = pulumi.output_property("destinationApplicationSecurityGroupIds")
    """
    A List of destination Application Security Group ID's
    """
    destination_port_range: pulumi.Output[Optional[str]] = pulumi.output_property("destinationPortRange")
    """
    Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
    """
    destination_port_ranges: pulumi.Output[Optional[List[str]]] = pulumi.output_property("destinationPortRanges")
    """
    List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
    """
    direction: pulumi.Output[str] = pulumi.output_property("direction")
    """
    The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
    """
    network_security_group_name: pulumi.Output[str] = pulumi.output_property("networkSecurityGroupName")
    """
    The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
    """
    priority: pulumi.Output[float] = pulumi.output_property("priority")
    """
    Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
    """
    protocol: pulumi.Output[str] = pulumi.output_property("protocol")
    """
    Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, or `*` (which matches all).
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
    """
    source_address_prefix: pulumi.Output[Optional[str]] = pulumi.output_property("sourceAddressPrefix")
    """
    CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
    """
    source_address_prefixes: pulumi.Output[Optional[List[str]]] = pulumi.output_property("sourceAddressPrefixes")
    """
    List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
    """
    source_application_security_group_ids: pulumi.Output[Optional[str]] = pulumi.output_property("sourceApplicationSecurityGroupIds")
    """
    A List of source Application Security Group ID's
    """
    source_port_range: pulumi.Output[Optional[str]] = pulumi.output_property("sourcePortRange")
    """
    Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
    """
    source_port_ranges: pulumi.Output[Optional[List[str]]] = pulumi.output_property("sourcePortRanges")
    """
    List of source ports or port ranges. This is required if `source_port_range` is not specified.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, access: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, destination_address_prefix: Optional[pulumi.Input[str]] = None, destination_address_prefixes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, destination_application_security_group_ids: Optional[pulumi.Input[str]] = None, destination_port_range: Optional[pulumi.Input[str]] = None, destination_port_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, direction: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, network_security_group_name: Optional[pulumi.Input[str]] = None, priority: Optional[pulumi.Input[float]] = None, protocol: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, source_address_prefix: Optional[pulumi.Input[str]] = None, source_address_prefixes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, source_application_security_group_ids: Optional[pulumi.Input[str]] = None, source_port_range: Optional[pulumi.Input[str]] = None, source_port_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Network Security Rule.

        > **NOTE on Network Security Groups and Network Security Rules:** This provider currently
        provides both a standalone Network Security Rule resource, and allows for Network Security Rules to be defined in-line within the Network Security Group resource.
        At this time you cannot use a Network Security Group with in-line Network Security Rules in conjunction with any Network Security Rule resources. Doing so will cause a conflict of rule settings and will overwrite rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_network_security_group = azure.network.NetworkSecurityGroup("exampleNetworkSecurityGroup",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_network_security_rule = azure.network.NetworkSecurityRule("exampleNetworkSecurityRule",
            priority=100,
            direction="Outbound",
            access="Allow",
            protocol="Tcp",
            source_port_range="*",
            destination_port_range="*",
            source_address_prefix="*",
            destination_address_prefix="*",
            resource_group_name=example_resource_group.name,
            network_security_group_name=example_network_security_group.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access: Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
        :param pulumi.Input[str] description: A description for this rule. Restricted to 140 characters.
        :param pulumi.Input[str] destination_address_prefix: CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destination_address_prefixes` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_address_prefixes: List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
        :param pulumi.Input[str] destination_application_security_group_ids: A List of destination Application Security Group ID's
        :param pulumi.Input[str] destination_port_range: Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_port_ranges: List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
        :param pulumi.Input[str] direction: The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
        :param pulumi.Input[str] name: The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_security_group_name: The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
        :param pulumi.Input[float] priority: Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        :param pulumi.Input[str] protocol: Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, or `*` (which matches all).
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_address_prefix: CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] source_address_prefixes: List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
        :param pulumi.Input[str] source_application_security_group_ids: A List of source Application Security Group ID's
        :param pulumi.Input[str] source_port_range: Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] source_port_ranges: List of source ports or port ranges. This is required if `source_port_range` is not specified.
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

            if access is None:
                raise TypeError("Missing required property 'access'")
            __props__['access'] = access
            __props__['description'] = description
            __props__['destination_address_prefix'] = destination_address_prefix
            __props__['destination_address_prefixes'] = destination_address_prefixes
            __props__['destination_application_security_group_ids'] = destination_application_security_group_ids
            __props__['destination_port_range'] = destination_port_range
            __props__['destination_port_ranges'] = destination_port_ranges
            if direction is None:
                raise TypeError("Missing required property 'direction'")
            __props__['direction'] = direction
            __props__['name'] = name
            if network_security_group_name is None:
                raise TypeError("Missing required property 'network_security_group_name'")
            __props__['network_security_group_name'] = network_security_group_name
            if priority is None:
                raise TypeError("Missing required property 'priority'")
            __props__['priority'] = priority
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['source_address_prefix'] = source_address_prefix
            __props__['source_address_prefixes'] = source_address_prefixes
            __props__['source_application_security_group_ids'] = source_application_security_group_ids
            __props__['source_port_range'] = source_port_range
            __props__['source_port_ranges'] = source_port_ranges
        super(NetworkSecurityRule, __self__).__init__(
            'azure:network/networkSecurityRule:NetworkSecurityRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, access: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, destination_address_prefix: Optional[pulumi.Input[str]] = None, destination_address_prefixes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, destination_application_security_group_ids: Optional[pulumi.Input[str]] = None, destination_port_range: Optional[pulumi.Input[str]] = None, destination_port_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, direction: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, network_security_group_name: Optional[pulumi.Input[str]] = None, priority: Optional[pulumi.Input[float]] = None, protocol: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, source_address_prefix: Optional[pulumi.Input[str]] = None, source_address_prefixes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, source_application_security_group_ids: Optional[pulumi.Input[str]] = None, source_port_range: Optional[pulumi.Input[str]] = None, source_port_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'NetworkSecurityRule':
        """
        Get an existing NetworkSecurityRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access: Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
        :param pulumi.Input[str] description: A description for this rule. Restricted to 140 characters.
        :param pulumi.Input[str] destination_address_prefix: CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destination_address_prefixes` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_address_prefixes: List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
        :param pulumi.Input[str] destination_application_security_group_ids: A List of destination Application Security Group ID's
        :param pulumi.Input[str] destination_port_range: Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_port_ranges: List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
        :param pulumi.Input[str] direction: The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
        :param pulumi.Input[str] name: The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_security_group_name: The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
        :param pulumi.Input[float] priority: Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        :param pulumi.Input[str] protocol: Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, or `*` (which matches all).
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_address_prefix: CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] source_address_prefixes: List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
        :param pulumi.Input[str] source_application_security_group_ids: A List of source Application Security Group ID's
        :param pulumi.Input[str] source_port_range: Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
        :param pulumi.Input[List[pulumi.Input[str]]] source_port_ranges: List of source ports or port ranges. This is required if `source_port_range` is not specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access"] = access
        __props__["description"] = description
        __props__["destination_address_prefix"] = destination_address_prefix
        __props__["destination_address_prefixes"] = destination_address_prefixes
        __props__["destination_application_security_group_ids"] = destination_application_security_group_ids
        __props__["destination_port_range"] = destination_port_range
        __props__["destination_port_ranges"] = destination_port_ranges
        __props__["direction"] = direction
        __props__["name"] = name
        __props__["network_security_group_name"] = network_security_group_name
        __props__["priority"] = priority
        __props__["protocol"] = protocol
        __props__["resource_group_name"] = resource_group_name
        __props__["source_address_prefix"] = source_address_prefix
        __props__["source_address_prefixes"] = source_address_prefixes
        __props__["source_application_security_group_ids"] = source_application_security_group_ids
        __props__["source_port_range"] = source_port_range
        __props__["source_port_ranges"] = source_port_ranges
        return NetworkSecurityRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

