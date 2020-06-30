# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class IntegrationAccount(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
    """
    name: pulumi.Output[str]
    """
    The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
    """
    sku_name: pulumi.Output[str]
    """
    The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags which should be assigned to the Logic App Integration Account.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Logic App Integration Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_integration_account = azure.logicapps.IntegrationAccount("exampleIntegrationAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku_name="Standard",
            tags={
                "foo": "bar",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        :param pulumi.Input[str] name: The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        :param pulumi.Input[str] sku_name: The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Logic App Integration Account.
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
        super(IntegrationAccount, __self__).__init__(
            'azure:logicapps/integrationAccount:IntegrationAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None):
        """
        Get an existing IntegrationAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        :param pulumi.Input[str] name: The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        :param pulumi.Input[str] sku_name: The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Logic App Integration Account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        return IntegrationAccount(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
