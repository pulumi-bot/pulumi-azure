# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['IdentityProviderFacebook']


class IdentityProviderFacebook(pulumi.CustomResource):
    api_management_name: pulumi.Output[str] = pulumi.output_property("apiManagementName")
    """
    The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
    """
    app_id: pulumi.Output[str] = pulumi.output_property("appId")
    """
    App ID for Facebook.
    """
    app_secret: pulumi.Output[str] = pulumi.output_property("appSecret")
    """
    App Secret for Facebook.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, api_management_name: Optional[pulumi.Input[str]] = None, app_id: Optional[pulumi.Input[str]] = None, app_secret: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an API Management Facebook Identity Provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="My Company",
            publisher_email="company@mycompany.io",
            sku_name="Developer_1")
        example_identity_provider_facebook = azure.apimanagement.IdentityProviderFacebook("exampleIdentityProviderFacebook",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            app_id="00000000000000000000000000000000",
            app_secret="00000000000000000000000000000000")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_id: App ID for Facebook.
        :param pulumi.Input[str] app_secret: App Secret for Facebook.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
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
            if app_id is None:
                raise TypeError("Missing required property 'app_id'")
            __props__['app_id'] = app_id
            if app_secret is None:
                raise TypeError("Missing required property 'app_secret'")
            __props__['app_secret'] = app_secret
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(IdentityProviderFacebook, __self__).__init__(
            'azure:apimanagement/identityProviderFacebook:IdentityProviderFacebook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, api_management_name: Optional[pulumi.Input[str]] = None, app_id: Optional[pulumi.Input[str]] = None, app_secret: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None) -> 'IdentityProviderFacebook':
        """
        Get an existing IdentityProviderFacebook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_id: App ID for Facebook.
        :param pulumi.Input[str] app_secret: App Secret for Facebook.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["app_id"] = app_id
        __props__["app_secret"] = app_secret
        __props__["resource_group_name"] = resource_group_name
        return IdentityProviderFacebook(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

