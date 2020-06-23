# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Account(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the Azure Maps Account. Changing this forces a new resource to be created.
    """
    primary_access_key: pulumi.Output[str]
    """
    The primary key used to authenticate and authorize access to the Maps REST APIs.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
    """
    secondary_access_key: pulumi.Output[str]
    """
    The secondary key used to authenticate and authorize access to the Maps REST APIs.
    """
    sku_name: pulumi.Output[str]
    """
    The sku of the Azure Maps Account. Possible values are `S0` and `S1`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the Azure Maps Account.
    """
    x_ms_client_id: pulumi.Output[str]
    """
    A unique identifier for the Maps Account.
    """
    def __init__(__self__, resource_name, opts=None, name=None, resource_group_name=None, sku_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Maps Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.maps.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            sku_name="S1",
            tags={
                "environment": "Test",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Azure Maps Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: The sku of the Azure Maps Account. Possible values are `S0` and `S1`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Azure Maps Account.
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
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['primary_access_key'] = None
            __props__['secondary_access_key'] = None
            __props__['x_ms_client_id'] = None
        super(Account, __self__).__init__(
            'azure:maps/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, primary_access_key=None, resource_group_name=None, secondary_access_key=None, sku_name=None, tags=None, x_ms_client_id=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Azure Maps Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The primary key used to authenticate and authorize access to the Maps REST APIs.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The secondary key used to authenticate and authorize access to the Maps REST APIs.
        :param pulumi.Input[str] sku_name: The sku of the Azure Maps Account. Possible values are `S0` and `S1`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Azure Maps Account.
        :param pulumi.Input[str] x_ms_client_id: A unique identifier for the Maps Account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["primary_access_key"] = primary_access_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["x_ms_client_id"] = x_ms_client_id
        return Account(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
