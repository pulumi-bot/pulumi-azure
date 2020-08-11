# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Account(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    An `identity` block as defined below.

      * `principal_id` (`str`) - The Principal ID for the Service Principal associated with the Identity of this Data Share Account.
      * `tenant_id` (`str`) - The Tenant ID for the Service Principal associated with the Identity of this Data Share Account.
      * `type` (`str`) - Specifies the identity type of the Data Share Account. At this time the only allowed value is `SystemAssigned`.
    """
    location: pulumi.Output[str]
    """
    The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
    """
    name: pulumi.Output[str]
    """
    The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags which should be assigned to the Data Share Account.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Data Share Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.datashare.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            identity={
                "type": "SystemAssigned",
            },
            tags={
                "foo": "bar",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Data Share Account.

        The **identity** object supports the following:

          * `principal_id` (`pulumi.Input[str]`) - The Principal ID for the Service Principal associated with the Identity of this Data Share Account.
          * `tenant_id` (`pulumi.Input[str]`) - The Tenant ID for the Service Principal associated with the Identity of this Data Share Account.
          * `type` (`pulumi.Input[str]`) - Specifies the identity type of the Data Share Account. At this time the only allowed value is `SystemAssigned`.
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

            if identity is None:
                raise TypeError("Missing required property 'identity'")
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Account, __self__).__init__(
            'azure:datashare/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, identity=None, location=None, name=None, resource_group_name=None, tags=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Data Share Account.

        The **identity** object supports the following:

          * `principal_id` (`pulumi.Input[str]`) - The Principal ID for the Service Principal associated with the Identity of this Data Share Account.
          * `tenant_id` (`pulumi.Input[str]`) - The Tenant ID for the Service Principal associated with the Identity of this Data Share Account.
          * `type` (`pulumi.Input[str]`) - Specifies the identity type of the Data Share Account. At this time the only allowed value is `SystemAssigned`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["identity"] = identity
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Account(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
