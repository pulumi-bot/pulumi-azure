# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['IdentityProviderTwitterArgs', 'IdentityProviderTwitter']

@pulumi.input_type
class IdentityProviderTwitterArgs:
    def __init__(__self__, *,
                 api_key: pulumi.Input[str],
                 api_management_name: pulumi.Input[str],
                 api_secret_key: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a IdentityProviderTwitter resource.
        :param pulumi.Input[str] api_key: App Consumer API key for Twitter.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_secret_key: App Consumer API secret key for Twitter.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "api_key", api_key)
        pulumi.set(__self__, "api_management_name", api_management_name)
        pulumi.set(__self__, "api_secret_key", api_secret_key)
        pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Input[str]:
        """
        App Consumer API key for Twitter.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Input[str]:
        """
        The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="apiSecretKey")
    def api_secret_key(self) -> pulumi.Input[str]:
        """
        App Consumer API secret key for Twitter.
        """
        return pulumi.get(self, "api_secret_key")

    @api_secret_key.setter
    def api_secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_secret_key", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)


class IdentityProviderTwitter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_secret_key: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an API Management Twitter Identity Provider.

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
        example_identity_provider_twitter = azure.apimanagement.IdentityProviderTwitter("exampleIdentityProviderTwitter",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            api_key="00000000000000000000000000000000",
            api_secret_key="00000000000000000000000000000000")
        ```

        ## Import

        API Management Twitter Identity Provider can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/twitter
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: App Consumer API key for Twitter.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_secret_key: App Consumer API secret key for Twitter.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityProviderTwitterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an API Management Twitter Identity Provider.

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
        example_identity_provider_twitter = azure.apimanagement.IdentityProviderTwitter("exampleIdentityProviderTwitter",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            api_key="00000000000000000000000000000000",
            api_secret_key="00000000000000000000000000000000")
        ```

        ## Import

        API Management Twitter Identity Provider can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/twitter
        ```

        :param str resource_name: The name of the resource.
        :param IdentityProviderTwitterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityProviderTwitterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_secret_key: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if api_key is None and not opts.urn:
                raise TypeError("Missing required property 'api_key'")
            __props__['api_key'] = api_key
            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if api_secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'api_secret_key'")
            __props__['api_secret_key'] = api_secret_key
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(IdentityProviderTwitter, __self__).__init__(
            'azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            api_secret_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'IdentityProviderTwitter':
        """
        Get an existing IdentityProviderTwitter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: App Consumer API key for Twitter.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_secret_key: App Consumer API secret key for Twitter.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_key"] = api_key
        __props__["api_management_name"] = api_management_name
        __props__["api_secret_key"] = api_secret_key
        __props__["resource_group_name"] = resource_group_name
        return IdentityProviderTwitter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[str]:
        """
        App Consumer API key for Twitter.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="apiSecretKey")
    def api_secret_key(self) -> pulumi.Output[str]:
        """
        App Consumer API secret key for Twitter.
        """
        return pulumi.get(self, "api_secret_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

