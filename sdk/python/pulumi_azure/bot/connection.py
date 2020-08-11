# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Connection']


class Connection(pulumi.CustomResource):
    bot_name: pulumi.Output[str] = pulumi.property("botName")
    """
    The name of the Bot Resource this connection will be associated with. Changing this forces a new resource to be created.
    """

    client_id: pulumi.Output[str] = pulumi.property("clientId")
    """
    The Client ID that will be used to authenticate with the service provider.
    """

    client_secret: pulumi.Output[str] = pulumi.property("clientSecret")
    """
    The Client Secret that will be used to authenticate with the service provider.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    The supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Specifies the name of the Bot Connection. Changing this forces a new resource to be created. Must be globally unique.
    """

    parameters: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("parameters")
    """
    A map of additional parameters to apply to the connection.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the resource group in which to create the Bot Connection. Changing this forces a new resource to be created.
    """

    scopes: pulumi.Output[Optional[str]] = pulumi.property("scopes")
    """
    The Scopes at which the connection should be applied.
    """

    service_provider_name: pulumi.Output[str] = pulumi.property("serviceProviderName")
    """
    The name of the service provider that will be associated with this connection. Changing this forces a new resource to be created.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bot_name: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[str]] = None,
                 service_provider_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Bot Connection.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_channels_registration = azure.bot.ChannelsRegistration("exampleChannelsRegistration",
            location="global",
            resource_group_name=example_resource_group.name,
            sku="F0",
            microsoft_app_id=current.client_id)
        example_connection = azure.bot.Connection("exampleConnection",
            bot_name=example_channels_registration.name,
            location=example_channels_registration.location,
            resource_group_name=example_resource_group.name,
            service_provider_name="box",
            client_id="exampleId",
            client_secret="exampleSecret")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this connection will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID that will be used to authenticate with the service provider.
        :param pulumi.Input[str] client_secret: The Client Secret that will be used to authenticate with the service provider.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Bot Connection. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of additional parameters to apply to the connection.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scopes: The Scopes at which the connection should be applied.
        :param pulumi.Input[str] service_provider_name: The name of the service provider that will be associated with this connection. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            if bot_name is None:
                raise TypeError("Missing required property 'bot_name'")
            __props__['bot_name'] = bot_name
            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['location'] = location
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scopes'] = scopes
            if service_provider_name is None:
                raise TypeError("Missing required property 'service_provider_name'")
            __props__['service_provider_name'] = service_provider_name
            __props__['tags'] = tags
        super(Connection, __self__).__init__(
            'azure:bot/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            bot_name: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[str]] = None,
            service_provider_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this connection will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID that will be used to authenticate with the service provider.
        :param pulumi.Input[str] client_secret: The Client Secret that will be used to authenticate with the service provider.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Bot Connection. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of additional parameters to apply to the connection.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scopes: The Scopes at which the connection should be applied.
        :param pulumi.Input[str] service_provider_name: The name of the service provider that will be associated with this connection. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bot_name"] = bot_name
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["location"] = location
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["scopes"] = scopes
        __props__["service_provider_name"] = service_provider_name
        __props__["tags"] = tags
        return Connection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

