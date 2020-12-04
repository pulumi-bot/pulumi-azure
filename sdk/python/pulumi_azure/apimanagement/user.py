# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['User']


class User(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 confirmation: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an API Management User.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="My Company",
            publisher_email="company@exmaple.com",
            sku_name="Developer_1")
        example_user = azure.apimanagement.User("exampleUser",
            user_id="5931a75ae4bbd512288c680b",
            api_management_name=example_service.name,
            resource_group_name=example_resource_group.name,
            first_name="Example",
            last_name="User",
            email="user@example.com",
            state="active")
        ```

        ## Import

        API Management Users can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/user:User example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/users/abc123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] confirmation: The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] email: The email address associated with this user.
        :param pulumi.Input[str] first_name: The first name for this user.
        :param pulumi.Input[str] last_name: The last name for this user.
        :param pulumi.Input[str] note: A note about this user.
        :param pulumi.Input[str] password: The password associated with this user.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] state: The state of this user. Possible values are `active`, `blocked` and `pending`.
        :param pulumi.Input[str] user_id: The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
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

            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            __props__['confirmation'] = confirmation
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            if first_name is None and not opts.urn:
                raise TypeError("Missing required property 'first_name'")
            __props__['first_name'] = first_name
            if last_name is None and not opts.urn:
                raise TypeError("Missing required property 'last_name'")
            __props__['last_name'] = last_name
            __props__['note'] = note
            __props__['password'] = password
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['state'] = state
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
        super(User, __self__).__init__(
            'azure:apimanagement/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            confirmation: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            first_name: Optional[pulumi.Input[str]] = None,
            last_name: Optional[pulumi.Input[str]] = None,
            note: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] confirmation: The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] email: The email address associated with this user.
        :param pulumi.Input[str] first_name: The first name for this user.
        :param pulumi.Input[str] last_name: The last name for this user.
        :param pulumi.Input[str] note: A note about this user.
        :param pulumi.Input[str] password: The password associated with this user.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] state: The state of this user. Possible values are `active`, `blocked` and `pending`.
        :param pulumi.Input[str] user_id: The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["confirmation"] = confirmation
        __props__["email"] = email
        __props__["first_name"] = first_name
        __props__["last_name"] = last_name
        __props__["note"] = note
        __props__["password"] = password
        __props__["resource_group_name"] = resource_group_name
        __props__["state"] = state
        __props__["user_id"] = user_id
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter
    def confirmation(self) -> pulumi.Output[Optional[str]]:
        """
        The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "confirmation")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email address associated with this user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[str]:
        """
        The first name for this user.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[str]:
        """
        The last name for this user.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter
    def note(self) -> pulumi.Output[Optional[str]]:
        """
        A note about this user.
        """
        return pulumi.get(self, "note")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password associated with this user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of this user. Possible values are `active`, `blocked` and `pending`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "user_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

