# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['TopicAuthorizationRule']

warnings.warn("azure.eventhub.TopicAuthorizationRule has been deprecated in favor of azure.servicebus.TopicAuthorizationRule", DeprecationWarning)


class TopicAuthorizationRule(pulumi.CustomResource):
    warnings.warn("azure.eventhub.TopicAuthorizationRule has been deprecated in favor of azure.servicebus.TopicAuthorizationRule", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 listen: Optional[pulumi.Input[bool]] = None,
                 manage: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 send: Optional[pulumi.Input[bool]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a ServiceBus Topic authorization Rule within a ServiceBus Topic.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] listen: Grants listen access to this this Authorization Rule. Defaults to `false`.
        :param pulumi.Input[bool] manage: Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] send: Grants send access to this this Authorization Rule. Defaults to `false`.
        :param pulumi.Input[str] topic_name: Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
        """
        pulumi.log.warn("TopicAuthorizationRule is deprecated: azure.eventhub.TopicAuthorizationRule has been deprecated in favor of azure.servicebus.TopicAuthorizationRule")
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

            __props__['listen'] = listen
            __props__['manage'] = manage
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['send'] = send
            if topic_name is None:
                raise TypeError("Missing required property 'topic_name'")
            __props__['topic_name'] = topic_name
            __props__['primary_connection_string'] = None
            __props__['primary_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['secondary_key'] = None
        super(TopicAuthorizationRule, __self__).__init__(
            'azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            listen: Optional[pulumi.Input[bool]] = None,
            manage: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            primary_connection_string: Optional[pulumi.Input[str]] = None,
            primary_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_connection_string: Optional[pulumi.Input[str]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            send: Optional[pulumi.Input[bool]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'TopicAuthorizationRule':
        """
        Get an existing TopicAuthorizationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] listen: Grants listen access to this this Authorization Rule. Defaults to `false`.
        :param pulumi.Input[bool] manage: Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The Primary Connection String for the ServiceBus Topic authorization Rule.
        :param pulumi.Input[str] primary_key: The Primary Key for the ServiceBus Topic authorization Rule.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string: The Secondary Connection String for the ServiceBus Topic authorization Rule.
        :param pulumi.Input[str] secondary_key: The Secondary Key for the ServiceBus Topic authorization Rule.
        :param pulumi.Input[bool] send: Grants send access to this this Authorization Rule. Defaults to `false`.
        :param pulumi.Input[str] topic_name: Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["listen"] = listen
        __props__["manage"] = manage
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["primary_connection_string"] = primary_connection_string
        __props__["primary_key"] = primary_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["secondary_key"] = secondary_key
        __props__["send"] = send
        __props__["topic_name"] = topic_name
        return TopicAuthorizationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def listen(self) -> pulumi.Output[Optional[bool]]:
        """
        Grants listen access to this this Authorization Rule. Defaults to `false`.
        """
        return pulumi.get(self, "listen")

    @property
    @pulumi.getter
    def manage(self) -> pulumi.Output[Optional[bool]]:
        """
        Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        """
        return pulumi.get(self, "manage")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> pulumi.Output[str]:
        """
        The Primary Connection String for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> pulumi.Output[str]:
        """
        The Primary Key for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> pulumi.Output[str]:
        """
        The Secondary Connection String for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[str]:
        """
        The Secondary Key for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter
    def send(self) -> pulumi.Output[Optional[bool]]:
        """
        Grants send access to this this Authorization Rule. Defaults to `false`.
        """
        return pulumi.get(self, "send")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "topic_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

