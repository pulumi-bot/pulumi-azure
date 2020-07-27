# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['TopicAuthorizationRule']


class TopicAuthorizationRule(pulumi.CustomResource):
    listen: pulumi.Output[Optional[bool]] = pulumi.output_property("listen")
    """
    Grants listen access to this this Authorization Rule. Defaults to `false`.
    """
    manage: pulumi.Output[Optional[bool]] = pulumi.output_property("manage")
    """
    Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
    """
    namespace_name: pulumi.Output[str] = pulumi.output_property("namespaceName")
    """
    Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
    """
    primary_connection_string: pulumi.Output[str] = pulumi.output_property("primaryConnectionString")
    """
    The Primary Connection String for the ServiceBus Topic authorization Rule.
    """
    primary_key: pulumi.Output[str] = pulumi.output_property("primaryKey")
    """
    The Primary Key for the ServiceBus Topic authorization Rule.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
    """
    secondary_connection_string: pulumi.Output[str] = pulumi.output_property("secondaryConnectionString")
    """
    The Secondary Connection String for the ServiceBus Topic authorization Rule.
    """
    secondary_key: pulumi.Output[str] = pulumi.output_property("secondaryKey")
    """
    The Secondary Key for the ServiceBus Topic authorization Rule.
    """
    send: pulumi.Output[Optional[bool]] = pulumi.output_property("send")
    """
    Grants send access to this this Authorization Rule. Defaults to `false`.
    """
    topic_name: pulumi.Output[str] = pulumi.output_property("topicName")
    """
    Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, listen: Optional[pulumi.Input[bool]] = None, manage: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, namespace_name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, send: Optional[pulumi.Input[bool]] = None, topic_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a ServiceBus Topic authorization Rule within a ServiceBus Topic.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name)
        example_topic_authorization_rule = azure.servicebus.TopicAuthorizationRule("exampleTopicAuthorizationRule",
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            resource_group_name=example_resource_group.name,
            listen=True,
            send=False,
            manage=False)
        ```

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
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(TopicAuthorizationRule, __self__).__init__(
            'azure:servicebus/topicAuthorizationRule:TopicAuthorizationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, listen: Optional[pulumi.Input[bool]] = None, manage: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, namespace_name: Optional[pulumi.Input[str]] = None, primary_connection_string: Optional[pulumi.Input[str]] = None, primary_key: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, secondary_connection_string: Optional[pulumi.Input[str]] = None, secondary_key: Optional[pulumi.Input[str]] = None, send: Optional[pulumi.Input[bool]] = None, topic_name: Optional[pulumi.Input[str]] = None) -> 'TopicAuthorizationRule':
        """
        Get an existing TopicAuthorizationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

