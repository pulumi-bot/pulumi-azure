# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class VirtualHub(pulumi.CustomResource):
    address_prefix: pulumi.Output[str]
    """
    The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the Virtual Hub. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
    """
    routes: pulumi.Output[list]
    """
    One or more `route` blocks as defined below.

      * `address_prefixes` (`list`) - A list of Address Prefixes.
      * `nextHopIpAddress` (`str`) - The IP Address that Packets should be forwarded to as the Next Hop.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the Virtual Hub.
    """
    virtual_wan_id: pulumi.Output[str]
    """
    The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, address_prefix=None, location=None, name=None, resource_group_name=None, routes=None, tags=None, virtual_wan_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Virtual Hub within a Virtual WAN.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_wan = azure.network.VirtualWan("exampleVirtualWan",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_virtual_hub = azure.network.VirtualHub("exampleVirtualHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            virtual_wan_id=example_virtual_wan.id,
            address_prefix="10.0.1.0/24")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Virtual Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[list] routes: One or more `route` blocks as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Virtual Hub.
        :param pulumi.Input[str] virtual_wan_id: The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.

        The **routes** object supports the following:

          * `address_prefixes` (`pulumi.Input[list]`) - A list of Address Prefixes.
          * `nextHopIpAddress` (`pulumi.Input[str]`) - The IP Address that Packets should be forwarded to as the Next Hop.
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

            if address_prefix is None:
                raise TypeError("Missing required property 'address_prefix'")
            __props__['address_prefix'] = address_prefix
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['routes'] = routes
            __props__['tags'] = tags
            if virtual_wan_id is None:
                raise TypeError("Missing required property 'virtual_wan_id'")
            __props__['virtual_wan_id'] = virtual_wan_id
        super(VirtualHub, __self__).__init__(
            'azure:network/virtualHub:VirtualHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address_prefix=None, location=None, name=None, resource_group_name=None, routes=None, tags=None, virtual_wan_id=None):
        """
        Get an existing VirtualHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Virtual Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[list] routes: One or more `route` blocks as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Virtual Hub.
        :param pulumi.Input[str] virtual_wan_id: The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.

        The **routes** object supports the following:

          * `address_prefixes` (`pulumi.Input[list]`) - A list of Address Prefixes.
          * `nextHopIpAddress` (`pulumi.Input[str]`) - The IP Address that Packets should be forwarded to as the Next Hop.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_prefix"] = address_prefix
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["routes"] = routes
        __props__["tags"] = tags
        __props__["virtual_wan_id"] = virtual_wan_id
        return VirtualHub(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
