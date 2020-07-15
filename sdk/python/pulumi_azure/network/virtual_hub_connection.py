# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class VirtualHubConnection(pulumi.CustomResource):
    hub_to_vitual_network_traffic_allowed: pulumi.Output[bool]
    """
    Is the Virtual Hub traffic allowed to transit via the Remote Virtual Network? Changing this forces a new resource to be created.
    """
    internet_security_enabled: pulumi.Output[bool]
    """
    Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
    """
    remote_virtual_network_id: pulumi.Output[str]
    """
    The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
    """
    virtual_hub_id: pulumi.Output[str]
    """
    The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
    """
    vitual_network_to_hub_gateways_traffic_allowed: pulumi.Output[bool]
    """
    Is Remote Virtual Network traffic allowed to transit the Hub's Virtual Network Gateway's? Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, hub_to_vitual_network_traffic_allowed=None, internet_security_enabled=None, name=None, remote_virtual_network_id=None, virtual_hub_id=None, vitual_network_to_hub_gateways_traffic_allowed=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Connection for a Virtual Hub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["172.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        test = azure.network.VirtualWan("test",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_virtual_hub = azure.network.VirtualHub("exampleVirtualHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            virtual_wan_id=azurerm_virtual_wan["example"]["id"],
            address_prefix="10.0.1.0/24")
        example_virtual_hub_connection = azure.network.VirtualHubConnection("exampleVirtualHubConnection",
            virtual_hub_id=example_virtual_hub.id,
            remote_virtual_network_id=example_virtual_network.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] hub_to_vitual_network_traffic_allowed: Is the Virtual Hub traffic allowed to transit via the Remote Virtual Network? Changing this forces a new resource to be created.
        :param pulumi.Input[bool] internet_security_enabled: Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] remote_virtual_network_id: The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_hub_id: The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] vitual_network_to_hub_gateways_traffic_allowed: Is Remote Virtual Network traffic allowed to transit the Hub's Virtual Network Gateway's? Changing this forces a new resource to be created.
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

            __props__['hub_to_vitual_network_traffic_allowed'] = hub_to_vitual_network_traffic_allowed
            __props__['internet_security_enabled'] = internet_security_enabled
            __props__['name'] = name
            if remote_virtual_network_id is None:
                raise TypeError("Missing required property 'remote_virtual_network_id'")
            __props__['remote_virtual_network_id'] = remote_virtual_network_id
            if virtual_hub_id is None:
                raise TypeError("Missing required property 'virtual_hub_id'")
            __props__['virtual_hub_id'] = virtual_hub_id
            __props__['vitual_network_to_hub_gateways_traffic_allowed'] = vitual_network_to_hub_gateways_traffic_allowed
        super(VirtualHubConnection, __self__).__init__(
            'azure:network/virtualHubConnection:VirtualHubConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, hub_to_vitual_network_traffic_allowed=None, internet_security_enabled=None, name=None, remote_virtual_network_id=None, virtual_hub_id=None, vitual_network_to_hub_gateways_traffic_allowed=None):
        """
        Get an existing VirtualHubConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] hub_to_vitual_network_traffic_allowed: Is the Virtual Hub traffic allowed to transit via the Remote Virtual Network? Changing this forces a new resource to be created.
        :param pulumi.Input[bool] internet_security_enabled: Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] remote_virtual_network_id: The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_hub_id: The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] vitual_network_to_hub_gateways_traffic_allowed: Is Remote Virtual Network traffic allowed to transit the Hub's Virtual Network Gateway's? Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["hub_to_vitual_network_traffic_allowed"] = hub_to_vitual_network_traffic_allowed
        __props__["internet_security_enabled"] = internet_security_enabled
        __props__["name"] = name
        __props__["remote_virtual_network_id"] = remote_virtual_network_id
        __props__["virtual_hub_id"] = virtual_hub_id
        __props__["vitual_network_to_hub_gateways_traffic_allowed"] = vitual_network_to_hub_gateways_traffic_allowed
        return VirtualHubConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
