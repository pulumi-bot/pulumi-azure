# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ChannelTeamsArgs', 'ChannelTeams']

@pulumi.input_type
class ChannelTeamsArgs:
    def __init__(__self__, *,
                 bot_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 calling_web_hook: Optional[pulumi.Input[str]] = None,
                 enable_calling: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ChannelTeams resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        :param pulumi.Input[str] calling_web_hook: Specifies the webhook for Microsoft Teams channel calls.
        :param pulumi.Input[bool] enable_calling: Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "bot_name", bot_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if calling_web_hook is not None:
            pulumi.set(__self__, "calling_web_hook", calling_web_hook)
        if enable_calling is not None:
            pulumi.set(__self__, "enable_calling", enable_calling)
        if location is not None:
            pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter(name="botName")
    def bot_name(self) -> pulumi.Input[str]:
        """
        The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "bot_name")

    @bot_name.setter
    def bot_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bot_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="callingWebHook")
    def calling_web_hook(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the webhook for Microsoft Teams channel calls.
        """
        return pulumi.get(self, "calling_web_hook")

    @calling_web_hook.setter
    def calling_web_hook(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "calling_web_hook", value)

    @property
    @pulumi.getter(name="enableCalling")
    def enable_calling(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        """
        return pulumi.get(self, "enable_calling")

    @enable_calling.setter
    def enable_calling(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_calling", value)

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


class ChannelTeams(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bot_name: Optional[pulumi.Input[str]] = None,
                 calling_web_hook: Optional[pulumi.Input[str]] = None,
                 enable_calling: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a MS Teams integration for a Bot Channel

        > **Note** A bot can only have a single MS Teams Channel associated with it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_channels_registration = azure.bot.ChannelsRegistration("exampleChannelsRegistration",
            location="global",
            resource_group_name=example_resource_group.name,
            sku="F0",
            microsoft_app_id=current.client_id)
        example_channel_teams = azure.bot.ChannelTeams("exampleChannelTeams",
            bot_name=example_channels_registration.name,
            location=example_channels_registration.location,
            resource_group_name=example_resource_group.name)
        ```

        ## Import

        The Microsoft Teams Integration for a Bot Channel can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:bot/channelTeams:ChannelTeams example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/MsTeamsChannel
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] calling_web_hook: Specifies the webhook for Microsoft Teams channel calls.
        :param pulumi.Input[bool] enable_calling: Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ChannelTeamsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a MS Teams integration for a Bot Channel

        > **Note** A bot can only have a single MS Teams Channel associated with it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_channels_registration = azure.bot.ChannelsRegistration("exampleChannelsRegistration",
            location="global",
            resource_group_name=example_resource_group.name,
            sku="F0",
            microsoft_app_id=current.client_id)
        example_channel_teams = azure.bot.ChannelTeams("exampleChannelTeams",
            bot_name=example_channels_registration.name,
            location=example_channels_registration.location,
            resource_group_name=example_resource_group.name)
        ```

        ## Import

        The Microsoft Teams Integration for a Bot Channel can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:bot/channelTeams:ChannelTeams example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/MsTeamsChannel
        ```

        :param str resource_name: The name of the resource.
        :param ChannelTeamsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ChannelTeamsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bot_name: Optional[pulumi.Input[str]] = None,
                 calling_web_hook: Optional[pulumi.Input[str]] = None,
                 enable_calling: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
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

            if bot_name is None and not opts.urn:
                raise TypeError("Missing required property 'bot_name'")
            __props__['bot_name'] = bot_name
            __props__['calling_web_hook'] = calling_web_hook
            __props__['enable_calling'] = enable_calling
            __props__['location'] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(ChannelTeams, __self__).__init__(
            'azure:bot/channelTeams:ChannelTeams',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bot_name: Optional[pulumi.Input[str]] = None,
            calling_web_hook: Optional[pulumi.Input[str]] = None,
            enable_calling: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'ChannelTeams':
        """
        Get an existing ChannelTeams resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] calling_web_hook: Specifies the webhook for Microsoft Teams channel calls.
        :param pulumi.Input[bool] enable_calling: Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bot_name"] = bot_name
        __props__["calling_web_hook"] = calling_web_hook
        __props__["enable_calling"] = enable_calling
        __props__["location"] = location
        __props__["resource_group_name"] = resource_group_name
        return ChannelTeams(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="botName")
    def bot_name(self) -> pulumi.Output[str]:
        """
        The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "bot_name")

    @property
    @pulumi.getter(name="callingWebHook")
    def calling_web_hook(self) -> pulumi.Output[str]:
        """
        Specifies the webhook for Microsoft Teams channel calls.
        """
        return pulumi.get(self, "calling_web_hook")

    @property
    @pulumi.getter(name="enableCalling")
    def enable_calling(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        """
        return pulumi.get(self, "enable_calling")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

