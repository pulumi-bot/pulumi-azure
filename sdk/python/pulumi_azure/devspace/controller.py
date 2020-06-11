# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Controller(pulumi.CustomResource):
    data_plane_fqdn: pulumi.Output[str]
    """
    DNS name for accessing DataPlane services.
    """
    host_suffix: pulumi.Output[str]
    """
    The host suffix for the DevSpace Controller.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
    """
    sku_name: pulumi.Output[str]
    """
    Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    target_container_host_credentials_base64: pulumi.Output[str]
    """
    Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
    """
    target_container_host_resource_id: pulumi.Output[str]
    """
    The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, target_container_host_credentials_base64=None, target_container_host_resource_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a DevSpace Controller.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_container_host_credentials_base64: Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_container_host_resource_id: The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            if target_container_host_credentials_base64 is None:
                raise TypeError("Missing required property 'target_container_host_credentials_base64'")
            __props__['target_container_host_credentials_base64'] = target_container_host_credentials_base64
            if target_container_host_resource_id is None:
                raise TypeError("Missing required property 'target_container_host_resource_id'")
            __props__['target_container_host_resource_id'] = target_container_host_resource_id
            __props__['data_plane_fqdn'] = None
            __props__['host_suffix'] = None
        super(Controller, __self__).__init__(
            'azure:devspace/controller:Controller',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, data_plane_fqdn=None, host_suffix=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, target_container_host_credentials_base64=None, target_container_host_resource_id=None):
        """
        Get an existing Controller resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_plane_fqdn: DNS name for accessing DataPlane services.
        :param pulumi.Input[str] host_suffix: The host suffix for the DevSpace Controller.
        :param pulumi.Input[str] location: Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_container_host_credentials_base64: Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_container_host_resource_id: The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_plane_fqdn"] = data_plane_fqdn
        __props__["host_suffix"] = host_suffix
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["target_container_host_credentials_base64"] = target_container_host_credentials_base64
        __props__["target_container_host_resource_id"] = target_container_host_resource_id
        return Controller(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

