# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Zone(pulumi.CustomResource):
    max_number_of_record_sets: pulumi.Output[float]
    """
    The maximum number of record sets that can be created in this Private DNS zone.
    """
    max_number_of_virtual_network_links: pulumi.Output[float]
    """
    The maximum number of virtual networks that can be linked to this Private DNS zone.
    """
    max_number_of_virtual_network_links_with_registration: pulumi.Output[float]
    """
    The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
    """
    name: pulumi.Output[str]
    """
    The name of the Private DNS Zone. Must be a valid domain name.
    """
    number_of_record_sets: pulumi.Output[float]
    """
    The current number of record sets in this Private DNS zone.
    """
    resource_group_name: pulumi.Output[str]
    """
    Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Enables you to manage Private DNS zones within Azure DNS. These zones are hosted on Azure's name servers.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_zone = azure.privatedns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Private DNS Zone. Must be a valid domain name.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['max_number_of_record_sets'] = None
            __props__['max_number_of_virtual_network_links'] = None
            __props__['max_number_of_virtual_network_links_with_registration'] = None
            __props__['number_of_record_sets'] = None
        super(Zone, __self__).__init__(
            'azure:privatedns/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, max_number_of_record_sets=None, max_number_of_virtual_network_links=None, max_number_of_virtual_network_links_with_registration=None, name=None, number_of_record_sets=None, resource_group_name=None, tags=None):
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] max_number_of_record_sets: The maximum number of record sets that can be created in this Private DNS zone.
        :param pulumi.Input[float] max_number_of_virtual_network_links: The maximum number of virtual networks that can be linked to this Private DNS zone.
        :param pulumi.Input[float] max_number_of_virtual_network_links_with_registration: The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
        :param pulumi.Input[str] name: The name of the Private DNS Zone. Must be a valid domain name.
        :param pulumi.Input[float] number_of_record_sets: The current number of record sets in this Private DNS zone.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["max_number_of_record_sets"] = max_number_of_record_sets
        __props__["max_number_of_virtual_network_links"] = max_number_of_virtual_network_links
        __props__["max_number_of_virtual_network_links_with_registration"] = max_number_of_virtual_network_links_with_registration
        __props__["name"] = name
        __props__["number_of_record_sets"] = number_of_record_sets
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Zone(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
