# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WebAppArgs', 'WebApp']

@pulumi.input_type
class WebAppArgs:
    def __init__(__self__, *,
                 microsoft_app_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 sku: pulumi.Input[str],
                 developer_app_insights_api_key: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_application_id: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 luis_app_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 luis_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a WebApp resource.
        :param pulumi.Input[str] microsoft_app_id: The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights API Key to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights Application ID to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_key: The Application Insights Key to associate with the Web App Bot.
        :param pulumi.Input[str] display_name: The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        :param pulumi.Input[str] endpoint: The Web App Bot endpoint.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] luis_app_ids: A list of LUIS App IDs to associate with the Web App Bot.
        :param pulumi.Input[str] luis_key: The LUIS key to associate with the Web App Bot.
        :param pulumi.Input[str] name: Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "microsoft_app_id", microsoft_app_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if developer_app_insights_api_key is not None:
            pulumi.set(__self__, "developer_app_insights_api_key", developer_app_insights_api_key)
        if developer_app_insights_application_id is not None:
            pulumi.set(__self__, "developer_app_insights_application_id", developer_app_insights_application_id)
        if developer_app_insights_key is not None:
            pulumi.set(__self__, "developer_app_insights_key", developer_app_insights_key)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if luis_app_ids is not None:
            pulumi.set(__self__, "luis_app_ids", luis_app_ids)
        if luis_key is not None:
            pulumi.set(__self__, "luis_key", luis_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="microsoftAppId")
    def microsoft_app_id(self) -> pulumi.Input[str]:
        """
        The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "microsoft_app_id")

    @microsoft_app_id.setter
    def microsoft_app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "microsoft_app_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input[str]:
        """
        The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="developerAppInsightsApiKey")
    def developer_app_insights_api_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights API Key to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_api_key")

    @developer_app_insights_api_key.setter
    def developer_app_insights_api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_api_key", value)

    @property
    @pulumi.getter(name="developerAppInsightsApplicationId")
    def developer_app_insights_application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights Application ID to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_application_id")

    @developer_app_insights_application_id.setter
    def developer_app_insights_application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_application_id", value)

    @property
    @pulumi.getter(name="developerAppInsightsKey")
    def developer_app_insights_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights Key to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_key")

    @developer_app_insights_key.setter
    def developer_app_insights_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_key", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The Web App Bot endpoint.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="luisAppIds")
    def luis_app_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of LUIS App IDs to associate with the Web App Bot.
        """
        return pulumi.get(self, "luis_app_ids")

    @luis_app_ids.setter
    def luis_app_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "luis_app_ids", value)

    @property
    @pulumi.getter(name="luisKey")
    def luis_key(self) -> Optional[pulumi.Input[str]]:
        """
        The LUIS key to associate with the Web App Bot.
        """
        return pulumi.get(self, "luis_key")

    @luis_key.setter
    def luis_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "luis_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _WebAppState:
    def __init__(__self__, *,
                 developer_app_insights_api_key: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_application_id: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 luis_app_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 luis_key: Optional[pulumi.Input[str]] = None,
                 microsoft_app_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering WebApp resources.
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights API Key to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights Application ID to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_key: The Application Insights Key to associate with the Web App Bot.
        :param pulumi.Input[str] display_name: The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        :param pulumi.Input[str] endpoint: The Web App Bot endpoint.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] luis_app_ids: A list of LUIS App IDs to associate with the Web App Bot.
        :param pulumi.Input[str] luis_key: The LUIS key to associate with the Web App Bot.
        :param pulumi.Input[str] microsoft_app_id: The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if developer_app_insights_api_key is not None:
            pulumi.set(__self__, "developer_app_insights_api_key", developer_app_insights_api_key)
        if developer_app_insights_application_id is not None:
            pulumi.set(__self__, "developer_app_insights_application_id", developer_app_insights_application_id)
        if developer_app_insights_key is not None:
            pulumi.set(__self__, "developer_app_insights_key", developer_app_insights_key)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if luis_app_ids is not None:
            pulumi.set(__self__, "luis_app_ids", luis_app_ids)
        if luis_key is not None:
            pulumi.set(__self__, "luis_key", luis_key)
        if microsoft_app_id is not None:
            pulumi.set(__self__, "microsoft_app_id", microsoft_app_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="developerAppInsightsApiKey")
    def developer_app_insights_api_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights API Key to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_api_key")

    @developer_app_insights_api_key.setter
    def developer_app_insights_api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_api_key", value)

    @property
    @pulumi.getter(name="developerAppInsightsApplicationId")
    def developer_app_insights_application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights Application ID to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_application_id")

    @developer_app_insights_application_id.setter
    def developer_app_insights_application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_application_id", value)

    @property
    @pulumi.getter(name="developerAppInsightsKey")
    def developer_app_insights_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights Key to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_key")

    @developer_app_insights_key.setter
    def developer_app_insights_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_key", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The Web App Bot endpoint.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="luisAppIds")
    def luis_app_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of LUIS App IDs to associate with the Web App Bot.
        """
        return pulumi.get(self, "luis_app_ids")

    @luis_app_ids.setter
    def luis_app_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "luis_app_ids", value)

    @property
    @pulumi.getter(name="luisKey")
    def luis_key(self) -> Optional[pulumi.Input[str]]:
        """
        The LUIS key to associate with the Web App Bot.
        """
        return pulumi.get(self, "luis_key")

    @luis_key.setter
    def luis_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "luis_key", value)

    @property
    @pulumi.getter(name="microsoftAppId")
    def microsoft_app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "microsoft_app_id")

    @microsoft_app_id.setter
    def microsoft_app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microsoft_app_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class WebApp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 developer_app_insights_api_key: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_application_id: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 luis_app_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 luis_key: Optional[pulumi.Input[str]] = None,
                 microsoft_app_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a Bot Web App.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_web_app = azure.bot.WebApp("exampleWebApp",
            location="global",
            resource_group_name=example_resource_group.name,
            sku="F0",
            microsoft_app_id=current.client_id)
        ```

        ## Import

        Bot Web App's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:bot/webApp:WebApp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights API Key to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights Application ID to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_key: The Application Insights Key to associate with the Web App Bot.
        :param pulumi.Input[str] display_name: The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        :param pulumi.Input[str] endpoint: The Web App Bot endpoint.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] luis_app_ids: A list of LUIS App IDs to associate with the Web App Bot.
        :param pulumi.Input[str] luis_key: The LUIS key to associate with the Web App Bot.
        :param pulumi.Input[str] microsoft_app_id: The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAppArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Bot Web App.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_web_app = azure.bot.WebApp("exampleWebApp",
            location="global",
            resource_group_name=example_resource_group.name,
            sku="F0",
            microsoft_app_id=current.client_id)
        ```

        ## Import

        Bot Web App's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:bot/webApp:WebApp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
        ```

        :param str resource_name: The name of the resource.
        :param WebAppArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAppArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 developer_app_insights_api_key: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_application_id: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 luis_app_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 luis_key: Optional[pulumi.Input[str]] = None,
                 microsoft_app_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = WebAppArgs.__new__(WebAppArgs)

            __props__.__dict__["developer_app_insights_api_key"] = developer_app_insights_api_key
            __props__.__dict__["developer_app_insights_application_id"] = developer_app_insights_application_id
            __props__.__dict__["developer_app_insights_key"] = developer_app_insights_key
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["location"] = location
            __props__.__dict__["luis_app_ids"] = luis_app_ids
            __props__.__dict__["luis_key"] = luis_key
            if microsoft_app_id is None and not opts.urn:
                raise TypeError("Missing required property 'microsoft_app_id'")
            __props__.__dict__["microsoft_app_id"] = microsoft_app_id
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
        super(WebApp, __self__).__init__(
            'azure:bot/webApp:WebApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            developer_app_insights_api_key: Optional[pulumi.Input[str]] = None,
            developer_app_insights_application_id: Optional[pulumi.Input[str]] = None,
            developer_app_insights_key: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            luis_app_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            luis_key: Optional[pulumi.Input[str]] = None,
            microsoft_app_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'WebApp':
        """
        Get an existing WebApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights API Key to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights Application ID to associate with the Web App Bot.
        :param pulumi.Input[str] developer_app_insights_key: The Application Insights Key to associate with the Web App Bot.
        :param pulumi.Input[str] display_name: The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        :param pulumi.Input[str] endpoint: The Web App Bot endpoint.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] luis_app_ids: A list of LUIS App IDs to associate with the Web App Bot.
        :param pulumi.Input[str] luis_key: The LUIS key to associate with the Web App Bot.
        :param pulumi.Input[str] microsoft_app_id: The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebAppState.__new__(_WebAppState)

        __props__.__dict__["developer_app_insights_api_key"] = developer_app_insights_api_key
        __props__.__dict__["developer_app_insights_application_id"] = developer_app_insights_application_id
        __props__.__dict__["developer_app_insights_key"] = developer_app_insights_key
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["location"] = location
        __props__.__dict__["luis_app_ids"] = luis_app_ids
        __props__.__dict__["luis_key"] = luis_key
        __props__.__dict__["microsoft_app_id"] = microsoft_app_id
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["sku"] = sku
        __props__.__dict__["tags"] = tags
        return WebApp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="developerAppInsightsApiKey")
    def developer_app_insights_api_key(self) -> pulumi.Output[str]:
        """
        The Application Insights API Key to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_api_key")

    @property
    @pulumi.getter(name="developerAppInsightsApplicationId")
    def developer_app_insights_application_id(self) -> pulumi.Output[str]:
        """
        The Application Insights Application ID to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_application_id")

    @property
    @pulumi.getter(name="developerAppInsightsKey")
    def developer_app_insights_key(self) -> pulumi.Output[str]:
        """
        The Application Insights Key to associate with the Web App Bot.
        """
        return pulumi.get(self, "developer_app_insights_key")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The Web App Bot endpoint.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="luisAppIds")
    def luis_app_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of LUIS App IDs to associate with the Web App Bot.
        """
        return pulumi.get(self, "luis_app_ids")

    @property
    @pulumi.getter(name="luisKey")
    def luis_key(self) -> pulumi.Output[Optional[str]]:
        """
        The LUIS key to associate with the Web App Bot.
        """
        return pulumi.get(self, "luis_key")

    @property
    @pulumi.getter(name="microsoftAppId")
    def microsoft_app_id(self) -> pulumi.Output[str]:
        """
        The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "microsoft_app_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[str]:
        """
        The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

