# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Subscription']


class Subscription(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_tracing: Optional[pulumi.Input[bool]] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 primary_key: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 secondary_key: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Subscription within a API Management Service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_tracing: Determines whether tracing can be enabled.  Defaults to `true`.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name of this Subscription.
        :param pulumi.Input[str] product_id: The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] state: The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
        :param pulumi.Input[str] subscription_id: An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
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

            __props__['allow_tracing'] = allow_tracing
            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['primary_key'] = primary_key
            __props__['product_id'] = product_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['secondary_key'] = secondary_key
            __props__['state'] = state
            __props__['subscription_id'] = subscription_id
            if user_id is None:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
        super(Subscription, __self__).__init__(
            'azure:apimanagement/subscription:Subscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_tracing: Optional[pulumi.Input[bool]] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            primary_key: Optional[pulumi.Input[str]] = None,
            product_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            subscription_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'Subscription':
        """
        Get an existing Subscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_tracing: Determines whether tracing can be enabled.  Defaults to `true`.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name of this Subscription.
        :param pulumi.Input[str] product_id: The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] state: The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
        :param pulumi.Input[str] subscription_id: An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_id: The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_tracing"] = allow_tracing
        __props__["api_management_name"] = api_management_name
        __props__["display_name"] = display_name
        __props__["primary_key"] = primary_key
        __props__["product_id"] = product_id
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_key"] = secondary_key
        __props__["state"] = state
        __props__["subscription_id"] = subscription_id
        __props__["user_id"] = user_id
        return Subscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowTracing")
    def allow_tracing(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether tracing can be enabled.  Defaults to `true`.
        """
        return pulumi.get(self, "allow_tracing")

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of this Subscription.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[str]:
        """
        An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "user_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

