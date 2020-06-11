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
    dsc_primary_access_key: pulumi.Output[str]
    """
    The Primary Access Key for the DSC Endpoint associated with this Automation Account.
    """
    dsc_secondary_access_key: pulumi.Output[str]
    """
    The Secondary Access Key for the DSC Endpoint associated with this Automation Account.
    """
    dsc_server_endpoint: pulumi.Output[str]
    """
    The DSC Server Endpoint associated with this Automation Account.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Automation Account. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Automation Account is created. Changing this forces a new resource to be created.
    """
    sku_name: pulumi.Output[str]
    """
    The SKU name of the account - only `Basic` is supported at this time.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Automation Account.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Basic",
            tags={
                "environment": "development",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Automation Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Automation Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: The SKU name of the account - only `Basic` is supported at this time.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['dsc_primary_access_key'] = None
            __props__['dsc_secondary_access_key'] = None
            __props__['dsc_server_endpoint'] = None
        super(Account, __self__).__init__(
            'azure:automation/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dsc_primary_access_key=None, dsc_secondary_access_key=None, dsc_server_endpoint=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dsc_primary_access_key: The Primary Access Key for the DSC Endpoint associated with this Automation Account.
        :param pulumi.Input[str] dsc_secondary_access_key: The Secondary Access Key for the DSC Endpoint associated with this Automation Account.
        :param pulumi.Input[str] dsc_server_endpoint: The DSC Server Endpoint associated with this Automation Account.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Automation Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Automation Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: The SKU name of the account - only `Basic` is supported at this time.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dsc_primary_access_key"] = dsc_primary_access_key
        __props__["dsc_secondary_access_key"] = dsc_secondary_access_key
        __props__["dsc_server_endpoint"] = dsc_server_endpoint
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        return Account(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

