# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupUserArgs', 'GroupUser']

@pulumi.input_type
class GroupUserArgs:
    def __init__(__self__, *,
                 api_management_name: pulumi.Input[str],
                 group_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 user_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a GroupUser resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] group_name: The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "api_management_name", api_management_name)
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Input[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _GroupUserState:
    def __init__(__self__, *,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupUser resources.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] group_name: The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        if api_management_name is not None:
            pulumi.set(__self__, "api_management_name", api_management_name)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class GroupUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
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

        ## Import

        API Management Group Users can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/groupUser:GroupUser example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/groups/groupId/users/user123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] group_name: The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        API Management Group Users can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/groupUser:GroupUser example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/groups/groupId/users/user123
        ```

        :param str resource_name: The name of the resource.
        :param GroupUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupUserArgs.__new__(GroupUserArgs)

            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__.__dict__["api_management_name"] = api_management_name
            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(GroupUser, __self__).__init__(
            'azure:apimanagement/groupUser:GroupUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'GroupUser':
        """
        Get an existing GroupUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] group_name: The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupUserState.__new__(_GroupUserState)

        __props__.__dict__["api_management_name"] = api_management_name
        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["user_id"] = user_id
        return GroupUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "user_id")

