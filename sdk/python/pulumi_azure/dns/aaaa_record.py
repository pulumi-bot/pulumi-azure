# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AaaaRecord(pulumi.CustomResource):
    fqdn: pulumi.Output[str]
    """
    The FQDN of the DNS AAAA Record.
    """
    name: pulumi.Output[str]
    """
    The name of the DNS AAAA Record.
    """
    records: pulumi.Output[list]
    """
    List of IPv4 Addresses. Conflicts with `target_resource_id`.
    """
    resource_group_name: pulumi.Output[str]
    """
    Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    target_resource_id: pulumi.Output[str]
    """
    The Azure resource id of the target object. Conflicts with `records`
    """
    ttl: pulumi.Output[float]
    zone_name: pulumi.Output[str]
    """
    Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, name=None, records=None, resource_group_name=None, tags=None, target_resource_id=None, ttl=None, zone_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Enables you to manage DNS AAAA Records within Azure DNS.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_zone = azure.dns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        example_aaaa_record = azure.dns.AaaaRecord("exampleAaaaRecord",
            zone_name=example_zone.name,
            resource_group_name=example_resource_group.name,
            ttl=300)
        ```

        ## Example Usage (Alias Record)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_zone = azure.dns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Dynamic",
            ip_version="IPv6")
        example_aaaa_record = azure.dns.AaaaRecord("exampleAaaaRecord",
            zone_name=example_zone.name,
            resource_group_name=example_resource_group.name,
            ttl=300,
            target_resource_id=example_public_ip.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the DNS AAAA Record.
        :param pulumi.Input[list] records: List of IPv4 Addresses. Conflicts with `target_resource_id`.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_resource_id: The Azure resource id of the target object. Conflicts with `records`
        :param pulumi.Input[str] zone_name: Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
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
            __props__['records'] = records
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['target_resource_id'] = target_resource_id
            if ttl is None:
                raise TypeError("Missing required property 'ttl'")
            __props__['ttl'] = ttl
            if zone_name is None:
                raise TypeError("Missing required property 'zone_name'")
            __props__['zone_name'] = zone_name
            __props__['fqdn'] = None
        super(AaaaRecord, __self__).__init__(
            'azure:dns/aaaaRecord:AaaaRecord',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, fqdn=None, name=None, records=None, resource_group_name=None, tags=None, target_resource_id=None, ttl=None, zone_name=None):
        """
        Get an existing AaaaRecord resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fqdn: The FQDN of the DNS AAAA Record.
        :param pulumi.Input[str] name: The name of the DNS AAAA Record.
        :param pulumi.Input[list] records: List of IPv4 Addresses. Conflicts with `target_resource_id`.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_resource_id: The Azure resource id of the target object. Conflicts with `records`
        :param pulumi.Input[str] zone_name: Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["fqdn"] = fqdn
        __props__["name"] = name
        __props__["records"] = records
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["target_resource_id"] = target_resource_id
        __props__["ttl"] = ttl
        __props__["zone_name"] = zone_name
        return AaaaRecord(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
