# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ChannelDirectLine(pulumi.CustomResource):
    bot_name: pulumi.Output[str]
    location: pulumi.Output[str]
    resource_group_name: pulumi.Output[str]
    sites: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, bot_name=None, location=None, resource_group_name=None, sites=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Directline integration for a Bot Channel

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **sites** object supports the following:

          * `enabled` (`pulumi.Input[bool]`)
          * `enhancedAuthenticationEnabled` (`pulumi.Input[bool]`)
          * `id` (`pulumi.Input[str]`)
          * `key` (`pulumi.Input[str]`)
          * `key2` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)
          * `trustedOrigins` (`pulumi.Input[list]`)
          * `v1Allowed` (`pulumi.Input[bool]`)
          * `v3Allowed` (`pulumi.Input[bool]`)
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
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sites is None:
                raise TypeError("Missing required property 'sites'")
            __props__['sites'] = sites
        super(ChannelDirectLine, __self__).__init__(
            'azure:bot/channelDirectLine:ChannelDirectLine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bot_name=None, location=None, resource_group_name=None, sites=None):
        """
        Get an existing ChannelDirectLine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **sites** object supports the following:

          * `enabled` (`pulumi.Input[bool]`)
          * `enhancedAuthenticationEnabled` (`pulumi.Input[bool]`)
          * `id` (`pulumi.Input[str]`)
          * `key` (`pulumi.Input[str]`)
          * `key2` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)
          * `trustedOrigins` (`pulumi.Input[list]`)
          * `v1Allowed` (`pulumi.Input[bool]`)
          * `v3Allowed` (`pulumi.Input[bool]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bot_name"] = bot_name
        __props__["location"] = location
        __props__["resource_group_name"] = resource_group_name
        __props__["sites"] = sites
        return ChannelDirectLine(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

