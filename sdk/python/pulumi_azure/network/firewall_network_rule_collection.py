# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FirewallNetworkRuleCollectionArgs', 'FirewallNetworkRuleCollection']

@pulumi.input_type
class FirewallNetworkRuleCollectionArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 azure_firewall_name: pulumi.Input[str],
                 priority: pulumi.Input[int],
                 resource_group_name: pulumi.Input[str],
                 rules: pulumi.Input[Sequence[pulumi.Input['FirewallNetworkRuleCollectionRuleArgs']]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallNetworkRuleCollection resource.
        :param pulumi.Input[str] action: Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        :param pulumi.Input[str] azure_firewall_name: Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] priority: Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallNetworkRuleCollectionRuleArgs']]] rules: One or more `rule` blocks as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "azure_firewall_name", azure_firewall_name)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "rules", rules)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="azureFirewallName")
    def azure_firewall_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "azure_firewall_name")

    @azure_firewall_name.setter
    def azure_firewall_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "azure_firewall_name", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallNetworkRuleCollectionRuleArgs']]]:
        """
        One or more `rule` blocks as defined below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallNetworkRuleCollectionRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class FirewallNetworkRuleCollection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallNetworkRuleCollectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Network Rule Collection within an Azure Firewall.

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
            address_prefixes=["10.0.1.0/24"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_firewall = azure.network.Firewall("exampleFirewall",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.FirewallIpConfigurationArgs(
                name="configuration",
                subnet_id=example_subnet.id,
                public_ip_address_id=example_public_ip.id,
            )])
        example_firewall_network_rule_collection = azure.network.FirewallNetworkRuleCollection("exampleFirewallNetworkRuleCollection",
            azure_firewall_name=example_firewall.name,
            resource_group_name=example_resource_group.name,
            priority=100,
            action="Allow",
            rules=[azure.network.FirewallNetworkRuleCollectionRuleArgs(
                name="testrule",
                source_addresses=["10.0.0.0/16"],
                destination_ports=["53"],
                destination_addresses=[
                    "8.8.8.8",
                    "8.8.4.4",
                ],
                protocols=[
                    "TCP",
                    "UDP",
                ],
            )])
        ```

        ## Import

        Azure Firewall Network Rule Collections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/networkRuleCollections/mycollection
        ```

        :param str resource_name: The name of the resource.
        :param FirewallNetworkRuleCollectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 azure_firewall_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallNetworkRuleCollectionRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Network Rule Collection within an Azure Firewall.

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
            address_prefixes=["10.0.1.0/24"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_firewall = azure.network.Firewall("exampleFirewall",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.FirewallIpConfigurationArgs(
                name="configuration",
                subnet_id=example_subnet.id,
                public_ip_address_id=example_public_ip.id,
            )])
        example_firewall_network_rule_collection = azure.network.FirewallNetworkRuleCollection("exampleFirewallNetworkRuleCollection",
            azure_firewall_name=example_firewall.name,
            resource_group_name=example_resource_group.name,
            priority=100,
            action="Allow",
            rules=[azure.network.FirewallNetworkRuleCollectionRuleArgs(
                name="testrule",
                source_addresses=["10.0.0.0/16"],
                destination_ports=["53"],
                destination_addresses=[
                    "8.8.8.8",
                    "8.8.4.4",
                ],
                protocols=[
                    "TCP",
                    "UDP",
                ],
            )])
        ```

        ## Import

        Azure Firewall Network Rule Collections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/networkRuleCollections/mycollection
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        :param pulumi.Input[str] azure_firewall_name: Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[int] priority: Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallNetworkRuleCollectionRuleArgs']]]] rules: One or more `rule` blocks as defined below.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallNetworkRuleCollectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 azure_firewall_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallNetworkRuleCollectionRuleArgs']]]]] = None,
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

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            if azure_firewall_name is None and not opts.urn:
                raise TypeError("Missing required property 'azure_firewall_name'")
            __props__['azure_firewall_name'] = azure_firewall_name
            __props__['name'] = name
            if priority is None and not opts.urn:
                raise TypeError("Missing required property 'priority'")
            __props__['priority'] = priority
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__['rules'] = rules
        super(FirewallNetworkRuleCollection, __self__).__init__(
            'azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            azure_firewall_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallNetworkRuleCollectionRuleArgs']]]]] = None) -> 'FirewallNetworkRuleCollection':
        """
        Get an existing FirewallNetworkRuleCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        :param pulumi.Input[str] azure_firewall_name: Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[int] priority: Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallNetworkRuleCollectionRuleArgs']]]] rules: One or more `rule` blocks as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["azure_firewall_name"] = azure_firewall_name
        __props__["name"] = name
        __props__["priority"] = priority
        __props__["resource_group_name"] = resource_group_name
        __props__["rules"] = rules
        return FirewallNetworkRuleCollection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="azureFirewallName")
    def azure_firewall_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "azure_firewall_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.FirewallNetworkRuleCollectionRule']]:
        """
        One or more `rule` blocks as defined below.
        """
        return pulumi.get(self, "rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

