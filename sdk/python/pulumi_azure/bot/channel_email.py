# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ChannelEmail']


class ChannelEmail(pulumi.CustomResource):
    bot_name: pulumi.Output[str] = pulumi.property("botName")
    """
    The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
    """

    email_address: pulumi.Output[str] = pulumi.property("emailAddress")
    """
    The email address that the Bot will authenticate with.
    """

    email_password: pulumi.Output[str] = pulumi.property("emailPassword")
    """
    The email password that the Bot will authenticate with.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    The supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bot_name: Optional[pulumi.Input[str]] = None,
                 email_address: Optional[pulumi.Input[str]] = None,
                 email_password: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Email integration for a Bot Channel

        > **Note** A bot can only have a single Email Channel associated with it.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] email_address: The email address that the Bot will authenticate with.
        :param pulumi.Input[str] email_password: The email password that the Bot will authenticate with.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
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
            if email_address is None:
                raise TypeError("Missing required property 'email_address'")
            __props__['email_address'] = email_address
            if email_password is None:
                raise TypeError("Missing required property 'email_password'")
            __props__['email_password'] = email_password
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(ChannelEmail, __self__).__init__(
            'azure:bot/channelEmail:ChannelEmail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            bot_name: Optional[pulumi.Input[str]] = None,
            email_address: Optional[pulumi.Input[str]] = None,
            email_password: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'ChannelEmail':
        """
        Get an existing ChannelEmail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] email_address: The email address that the Bot will authenticate with.
        :param pulumi.Input[str] email_password: The email password that the Bot will authenticate with.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bot_name"] = bot_name
        __props__["email_address"] = email_address
        __props__["email_password"] = email_password
        __props__["location"] = location
        __props__["resource_group_name"] = resource_group_name
        return ChannelEmail(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

