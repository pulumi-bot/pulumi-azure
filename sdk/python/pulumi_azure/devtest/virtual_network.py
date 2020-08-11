# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VirtualNetwork']


class VirtualNetwork(pulumi.CustomResource):
    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    A description for the Virtual Network.
    """

    lab_name: pulumi.Output[str] = pulumi.property("labName")
    """
    Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
    """

    subnet: pulumi.Output['outputs.VirtualNetworkSubnet'] = pulumi.property("subnet")
    """
    A `subnet` block as defined below.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    unique_identifier: pulumi.Output[str] = pulumi.property("uniqueIdentifier")
    """
    The unique immutable identifier of the Dev Test Virtual Network.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Virtual Network within a DevTest Lab.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_lab = azure.devtest.Lab("exampleLab",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "Sydney": "Australia",
            })
        example_virtual_network = azure.devtest.VirtualNetwork("exampleVirtualNetwork",
            lab_name=example_lab.name,
            resource_group_name=example_resource_group.name,
            subnet={
                "usePublicIpAddress": "Allow",
                "useInVirtualMachineCreation": "Allow",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the Virtual Network.
        :param pulumi.Input[str] lab_name: Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']] subnet: A `subnet` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['description'] = description
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['subnet'] = subnet
            __props__['tags'] = tags
            __props__['unique_identifier'] = None
        super(VirtualNetwork, __self__).__init__(
            'azure:devtest/virtualNetwork:VirtualNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            lab_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subnet: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            unique_identifier: Optional[pulumi.Input[str]] = None) -> 'VirtualNetwork':
        """
        Get an existing VirtualNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the Virtual Network.
        :param pulumi.Input[str] lab_name: Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkSubnetArgs']] subnet: A `subnet` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] unique_identifier: The unique immutable identifier of the Dev Test Virtual Network.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["lab_name"] = lab_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["subnet"] = subnet
        __props__["tags"] = tags
        __props__["unique_identifier"] = unique_identifier
        return VirtualNetwork(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

