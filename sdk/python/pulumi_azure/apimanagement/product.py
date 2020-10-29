# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Product']


class Product(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 approval_required: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 published: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subscription_required: Optional[pulumi.Input[bool]] = None,
                 subscriptions_limit: Optional[pulumi.Input[int]] = None,
                 terms: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an API Management Product.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] approval_required: Do subscribers need to be approved prior to being able to use the Product?
        :param pulumi.Input[str] description: A description of this Product, which may include HTML formatting tags.
        :param pulumi.Input[str] display_name: The Display Name for this API Management Product.
        :param pulumi.Input[str] product_id: The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] published: Is this Product Published?
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] subscription_required: Is a Subscription required to access API's included in this Product?
        :param pulumi.Input[int] subscriptions_limit: The number of subscriptions a user can have to this Product at the same time.
        :param pulumi.Input[str] terms: The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
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
            __props__['approval_required'] = approval_required
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if product_id is None:
                raise TypeError("Missing required property 'product_id'")
            __props__['product_id'] = product_id
            if published is None:
                raise TypeError("Missing required property 'published'")
            __props__['published'] = published
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if subscription_required is None:
                raise TypeError("Missing required property 'subscription_required'")
            __props__['subscription_required'] = subscription_required
            __props__['subscriptions_limit'] = subscriptions_limit
            __props__['terms'] = terms
        super(Product, __self__).__init__(
            'azure:apimanagement/product:Product',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            approval_required: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            product_id: Optional[pulumi.Input[str]] = None,
            published: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subscription_required: Optional[pulumi.Input[bool]] = None,
            subscriptions_limit: Optional[pulumi.Input[int]] = None,
            terms: Optional[pulumi.Input[str]] = None) -> 'Product':
        """
        Get an existing Product resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] approval_required: Do subscribers need to be approved prior to being able to use the Product?
        :param pulumi.Input[str] description: A description of this Product, which may include HTML formatting tags.
        :param pulumi.Input[str] display_name: The Display Name for this API Management Product.
        :param pulumi.Input[str] product_id: The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] published: Is this Product Published?
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] subscription_required: Is a Subscription required to access API's included in this Product?
        :param pulumi.Input[int] subscriptions_limit: The number of subscriptions a user can have to this Product at the same time.
        :param pulumi.Input[str] terms: The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["approval_required"] = approval_required
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["product_id"] = product_id
        __props__["published"] = published
        __props__["resource_group_name"] = resource_group_name
        __props__["subscription_required"] = subscription_required
        __props__["subscriptions_limit"] = subscriptions_limit
        __props__["terms"] = terms
        return Product(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="approvalRequired")
    def approval_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Do subscribers need to be approved prior to being able to use the Product?
        """
        return pulumi.get(self, "approval_required")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of this Product, which may include HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The Display Name for this API Management Product.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> pulumi.Output[str]:
        """
        The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter
    def published(self) -> pulumi.Output[bool]:
        """
        Is this Product Published?
        """
        return pulumi.get(self, "published")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="subscriptionRequired")
    def subscription_required(self) -> pulumi.Output[bool]:
        """
        Is a Subscription required to access API's included in this Product?
        """
        return pulumi.get(self, "subscription_required")

    @property
    @pulumi.getter(name="subscriptionsLimit")
    def subscriptions_limit(self) -> pulumi.Output[Optional[int]]:
        """
        The number of subscriptions a user can have to this Product at the same time.
        """
        return pulumi.get(self, "subscriptions_limit")

    @property
    @pulumi.getter
    def terms(self) -> pulumi.Output[Optional[str]]:
        """
        The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
        """
        return pulumi.get(self, "terms")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

