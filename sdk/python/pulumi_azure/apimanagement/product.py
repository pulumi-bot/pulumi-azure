# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Product(pulumi.CustomResource):
    api_management_name: pulumi.Output[str]
    """
    The name of the API Management Service. Changing this forces a new resource to be created.
    """
    approval_required: pulumi.Output[bool]
    """
    Do subscribers need to be approved prior to being able to use the Product?
    """
    description: pulumi.Output[str]
    """
    A description of this Product, which may include HTML formatting tags.
    """
    display_name: pulumi.Output[str]
    """
    The Display Name for this API Management Product.
    """
    product_id: pulumi.Output[str]
    """
    The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
    """
    published: pulumi.Output[bool]
    """
    Is this Product Published?
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
    """
    subscription_required: pulumi.Output[bool]
    """
    Is a Subscription required to access API's included in this Product?
    """
    subscriptions_limit: pulumi.Output[float]
    """
    The number of subscriptions a user can have to this Product at the same time.
    """
    terms: pulumi.Output[str]
    """
    The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
    """
    def __init__(__self__, resource_name, opts=None, api_management_name=None, approval_required=None, description=None, display_name=None, product_id=None, published=None, resource_group_name=None, subscription_required=None, subscriptions_limit=None, terms=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an API Management Product.

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
        example_product = azure.apimanagement.Product("exampleProduct",
            product_id="test-product",
            api_management_name=example_service.name,
            resource_group_name=example_resource_group.name,
            display_name="Test Product",
            subscription_required=True,
            approval_required=True,
            published=True)
        ```


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
        :param pulumi.Input[float] subscriptions_limit: The number of subscriptions a user can have to this Product at the same time.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, api_management_name=None, approval_required=None, description=None, display_name=None, product_id=None, published=None, resource_group_name=None, subscription_required=None, subscriptions_limit=None, terms=None):
        """
        Get an existing Product resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] approval_required: Do subscribers need to be approved prior to being able to use the Product?
        :param pulumi.Input[str] description: A description of this Product, which may include HTML formatting tags.
        :param pulumi.Input[str] display_name: The Display Name for this API Management Product.
        :param pulumi.Input[str] product_id: The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] published: Is this Product Published?
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] subscription_required: Is a Subscription required to access API's included in this Product?
        :param pulumi.Input[float] subscriptions_limit: The number of subscriptions a user can have to this Product at the same time.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

