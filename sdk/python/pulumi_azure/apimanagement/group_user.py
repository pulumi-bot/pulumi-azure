# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['GroupUser']


class GroupUser(pulumi.CustomResource):
    api_management_name: pulumi.Output[str] = pulumi.output_property("apiManagementName")
    """
    The name of the API Management Service. Changing this forces a new resource to be created.
    """
    group_name: pulumi.Output[str] = pulumi.output_property("groupName")
    """
    The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
    """
    user_id: pulumi.Output[str] = pulumi.output_property("userId")
    """
    The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, api_management_name: Optional[pulumi.Input[str]] = None, group_name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, user_id: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an API Management User Assignment to a Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_user = azure.apimanagement.get_user(user_id="my-user",
            api_management_name="example-apim",
            resource_group_name="search-service")
        example_group_user = azure.apimanagement.GroupUser("exampleGroupUser",
            user_id=example_user.id,
            group_name="example-group",
            resource_group_name=example_user.resource_group_name,
            api_management_name=example_user.api_management_name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] group_name: The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
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

            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if group_name is None:
                raise TypeError("Missing required property 'group_name'")
            __props__['group_name'] = group_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if user_id is None:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
        super(GroupUser, __self__).__init__(
            'azure:apimanagement/groupUser:GroupUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, api_management_name: Optional[pulumi.Input[str]] = None, group_name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, user_id: Optional[pulumi.Input[str]] = None) -> 'GroupUser':
        """
        Get an existing GroupUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] group_name: The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["group_name"] = group_name
        __props__["resource_group_name"] = resource_group_name
        __props__["user_id"] = user_id
        return GroupUser(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

