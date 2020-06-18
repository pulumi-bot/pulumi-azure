# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class CustomProvider(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    Any number of `action` block as defined below. One of `resource_type` or `action` must be specified.

      * `endpoint` (`str`) - Specifies the endpoint of the action.
      * `name` (`str`) - Specifies the name of the action.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Custom Provider.
    """
    resource_types: pulumi.Output[list]
    """
    Any number of `resource_type` block as defined below. One of `resource_type` or `action` must be specified.

      * `endpoint` (`str`) - Specifies the endpoint of the route definition.
      * `name` (`str`) - Specifies the name of the route definition.
      * `routingType` (`str`) - The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    validations: pulumi.Output[list]
    """
    Any number of `validation` block as defined below.

      * `specification` (`str`) - The endpoint where the validation specification is located.
    """
    def __init__(__self__, resource_name, opts=None, actions=None, location=None, name=None, resource_group_name=None, resource_types=None, tags=None, validations=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Custom Provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_custom_provider = azure.core.CustomProvider("exampleCustomProvider",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            resource_type=[{
                "name": "dEf1",
                "endpoint": "https://testendpoint.com/",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: Any number of `action` block as defined below. One of `resource_type` or `action` must be specified.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Custom Provider.
        :param pulumi.Input[list] resource_types: Any number of `resource_type` block as defined below. One of `resource_type` or `action` must be specified.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[list] validations: Any number of `validation` block as defined below.

        The **actions** object supports the following:

          * `endpoint` (`pulumi.Input[str]`) - Specifies the endpoint of the action.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the action.

        The **resource_types** object supports the following:

          * `endpoint` (`pulumi.Input[str]`) - Specifies the endpoint of the route definition.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the route definition.
          * `routingType` (`pulumi.Input[str]`) - The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.

        The **validations** object supports the following:

          * `specification` (`pulumi.Input[str]`) - The endpoint where the validation specification is located.
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

            __props__['actions'] = actions
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_types'] = resource_types
            __props__['tags'] = tags
            __props__['validations'] = validations
        super(CustomProvider, __self__).__init__(
            'azure:core/customProvider:CustomProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions=None, location=None, name=None, resource_group_name=None, resource_types=None, tags=None, validations=None):
        """
        Get an existing CustomProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: Any number of `action` block as defined below. One of `resource_type` or `action` must be specified.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Custom Provider.
        :param pulumi.Input[list] resource_types: Any number of `resource_type` block as defined below. One of `resource_type` or `action` must be specified.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[list] validations: Any number of `validation` block as defined below.

        The **actions** object supports the following:

          * `endpoint` (`pulumi.Input[str]`) - Specifies the endpoint of the action.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the action.

        The **resource_types** object supports the following:

          * `endpoint` (`pulumi.Input[str]`) - Specifies the endpoint of the route definition.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the route definition.
          * `routingType` (`pulumi.Input[str]`) - The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.

        The **validations** object supports the following:

          * `specification` (`pulumi.Input[str]`) - The endpoint where the validation specification is located.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions"] = actions
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["resource_types"] = resource_types
        __props__["tags"] = tags
        __props__["validations"] = validations
        return CustomProvider(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
