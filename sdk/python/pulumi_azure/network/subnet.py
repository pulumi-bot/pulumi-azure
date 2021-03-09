# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Subnet']


class Subnet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[str]] = None,
                 address_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 delegations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetDelegationArgs']]]]] = None,
                 enforce_private_link_endpoint_network_policies: Optional[pulumi.Input[bool]] = None,
                 enforce_private_link_service_network_policies: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_endpoint_policy_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_network_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.

        > **NOTE on Virtual Networks and Subnet's:** This provider currently
        provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
        At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"],
            delegations=[azure.network.SubnetDelegationArgs(
                name="delegation",
                service_delegation=azure.network.SubnetDelegationServiceDelegationArgs(
                    name="Microsoft.ContainerInstance/containerGroups",
                    actions=[
                        "Microsoft.Network/virtualNetworks/subnets/join/action",
                        "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
                    ],
                ),
            )])
        ```

        ## Import

        Subnets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/subnet:Subnet exampleSubnet /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The address prefix to use for the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_prefixes: The address prefixes to use for the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetDelegationArgs']]]] delegations: One or more `delegation` blocks as defined below.
        :param pulumi.Input[bool] enforce_private_link_endpoint_network_policies: Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
        :param pulumi.Input[bool] enforce_private_link_service_network_policies: Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforce_private_link_endpoint_network_policies`.
        :param pulumi.Input[str] name: The name of the subnet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_endpoint_policy_ids: The list of IDs of Service Endpoint Policies to associate with the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_endpoints: The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
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

            if address_prefix is not None and not opts.urn:
                warnings.warn("""Use the `address_prefixes` property instead.""", DeprecationWarning)
                pulumi.log.warn("""address_prefix is deprecated: Use the `address_prefixes` property instead.""")
            __props__['address_prefix'] = address_prefix
            __props__['address_prefixes'] = address_prefixes
            __props__['delegations'] = delegations
            __props__['enforce_private_link_endpoint_network_policies'] = enforce_private_link_endpoint_network_policies
            __props__['enforce_private_link_service_network_policies'] = enforce_private_link_service_network_policies
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_endpoint_policy_ids'] = service_endpoint_policy_ids
            __props__['service_endpoints'] = service_endpoints
            if virtual_network_name is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_network_name'")
            __props__['virtual_network_name'] = virtual_network_name
        super(Subnet, __self__).__init__(
            'azure:network/subnet:Subnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_prefix: Optional[pulumi.Input[str]] = None,
            address_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            delegations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetDelegationArgs']]]]] = None,
            enforce_private_link_endpoint_network_policies: Optional[pulumi.Input[bool]] = None,
            enforce_private_link_service_network_policies: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_endpoint_policy_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            service_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            virtual_network_name: Optional[pulumi.Input[str]] = None) -> 'Subnet':
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The address prefix to use for the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_prefixes: The address prefixes to use for the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetDelegationArgs']]]] delegations: One or more `delegation` blocks as defined below.
        :param pulumi.Input[bool] enforce_private_link_endpoint_network_policies: Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
        :param pulumi.Input[bool] enforce_private_link_service_network_policies: Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforce_private_link_endpoint_network_policies`.
        :param pulumi.Input[str] name: The name of the subnet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_endpoint_policy_ids: The list of IDs of Service Endpoint Policies to associate with the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_endpoints: The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_prefix"] = address_prefix
        __props__["address_prefixes"] = address_prefixes
        __props__["delegations"] = delegations
        __props__["enforce_private_link_endpoint_network_policies"] = enforce_private_link_endpoint_network_policies
        __props__["enforce_private_link_service_network_policies"] = enforce_private_link_service_network_policies
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["service_endpoint_policy_ids"] = service_endpoint_policy_ids
        __props__["service_endpoints"] = service_endpoints
        __props__["virtual_network_name"] = virtual_network_name
        return Subnet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> pulumi.Output[str]:
        """
        The address prefix to use for the subnet.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter(name="addressPrefixes")
    def address_prefixes(self) -> pulumi.Output[Sequence[str]]:
        """
        The address prefixes to use for the subnet.
        """
        return pulumi.get(self, "address_prefixes")

    @property
    @pulumi.getter
    def delegations(self) -> pulumi.Output[Optional[Sequence['outputs.SubnetDelegation']]]:
        """
        One or more `delegation` blocks as defined below.
        """
        return pulumi.get(self, "delegations")

    @property
    @pulumi.getter(name="enforcePrivateLinkEndpointNetworkPolicies")
    def enforce_private_link_endpoint_network_policies(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
        """
        return pulumi.get(self, "enforce_private_link_endpoint_network_policies")

    @property
    @pulumi.getter(name="enforcePrivateLinkServiceNetworkPolicies")
    def enforce_private_link_service_network_policies(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforce_private_link_endpoint_network_policies`.
        """
        return pulumi.get(self, "enforce_private_link_service_network_policies")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the subnet. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceEndpointPolicyIds")
    def service_endpoint_policy_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of IDs of Service Endpoint Policies to associate with the subnet.
        """
        return pulumi.get(self, "service_endpoint_policy_ids")

    @property
    @pulumi.getter(name="serviceEndpoints")
    def service_endpoints(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
        """
        return pulumi.get(self, "service_endpoints")

    @property
    @pulumi.getter(name="virtualNetworkName")
    def virtual_network_name(self) -> pulumi.Output[str]:
        """
        The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

