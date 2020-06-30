# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ChannelSlack(pulumi.CustomResource):
    bot_name: pulumi.Output[str]
    """
    The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
    """
    client_id: pulumi.Output[str]
    """
    The Client ID that will be used to authenticate with Slack.
    """
    client_secret: pulumi.Output[str]
    """
    The Client Secret that will be used to authenticate with Slack.
    """
    landing_page_url: pulumi.Output[str]
    """
    The Slack Landing Page URL.
    """
    location: pulumi.Output[str]
    """
    The supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
    """
    verification_token: pulumi.Output[str]
    """
    The Verification Token that will be used to authenticate with Slack.
    """
    def __init__(__self__, resource_name, opts=None, bot_name=None, client_id=None, client_secret=None, landing_page_url=None, location=None, resource_group_name=None, verification_token=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Slack integration for a Bot Channel

        > **Note** A bot can only have a single Slack Channel associated with it.

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
        example_channel_slack = azure.bot.ChannelSlack("exampleChannelSlack",
            bot_name=example_channels_registration.name,
            location=example_channels_registration.location,
            resource_group_name=example_resource_group.name,
            client_id="exampleId",
            client_secret="exampleSecret",
            verification_token="exampleVerificationToken")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID that will be used to authenticate with Slack.
        :param pulumi.Input[str] client_secret: The Client Secret that will be used to authenticate with Slack.
        :param pulumi.Input[str] landing_page_url: The Slack Landing Page URL.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        :param pulumi.Input[str] verification_token: The Verification Token that will be used to authenticate with Slack.
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

            if bot_name is None:
                raise TypeError("Missing required property 'bot_name'")
            __props__['bot_name'] = bot_name
            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['landing_page_url'] = landing_page_url
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if verification_token is None:
                raise TypeError("Missing required property 'verification_token'")
            __props__['verification_token'] = verification_token
        super(ChannelSlack, __self__).__init__(
            'azure:bot/channelSlack:ChannelSlack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bot_name=None, client_id=None, client_secret=None, landing_page_url=None, location=None, resource_group_name=None, verification_token=None):
        """
        Get an existing ChannelSlack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID that will be used to authenticate with Slack.
        :param pulumi.Input[str] client_secret: The Client Secret that will be used to authenticate with Slack.
        :param pulumi.Input[str] landing_page_url: The Slack Landing Page URL.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        :param pulumi.Input[str] verification_token: The Verification Token that will be used to authenticate with Slack.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bot_name"] = bot_name
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["landing_page_url"] = landing_page_url
        __props__["location"] = location
        __props__["resource_group_name"] = resource_group_name
        __props__["verification_token"] = verification_token
        return ChannelSlack(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
