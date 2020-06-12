# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Group(pulumi.CustomResource):
    api_management_name: pulumi.Output[str]
    """
    The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
    """
    description: pulumi.Output[str]
    """
    The description of this API Management Group.
    """
    display_name: pulumi.Output[str]
    """
    The display name of this API Management Group.
    """
    external_id: pulumi.Output[str]
    """
    The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
    """
    name: pulumi.Output[str]
    """
    The name of the API Management Group. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
    """
    type: pulumi.Output[str]
    """
    The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
    """
    def __init__(__self__, resource_name, opts=None, api_management_name=None, description=None, display_name=None, external_id=None, name=None, resource_group_name=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an API Management Group.


        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="pub1",
            publisher_email="pub1@email.com",
            sku_name="Developer_1")
        example_group = azure.apimanagement.Group("exampleGroup",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            display_name="Example Group",
            description="This is an example API management group.")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of this API Management Group.
        :param pulumi.Input[str] display_name: The display name of this API Management Group.
        :param pulumi.Input[str] external_id: The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
        :param pulumi.Input[str] name: The name of the API Management Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
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

            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['external_id'] = external_id
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = type
        super(Group, __self__).__init__(
            'azure:apimanagement/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_management_name=None, description=None, display_name=None, external_id=None, name=None, resource_group_name=None, type=None):
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of this API Management Group.
        :param pulumi.Input[str] display_name: The display name of this API Management Group.
        :param pulumi.Input[str] external_id: The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
        :param pulumi.Input[str] name: The name of the API Management Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["external_id"] = external_id
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["type"] = type
        return Group(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
