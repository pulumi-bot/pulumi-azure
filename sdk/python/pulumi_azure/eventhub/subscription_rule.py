# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables

warnings.warn("azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule", DeprecationWarning)


class SubscriptionRule(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
    """
    correlation_filter: pulumi.Output[dict]
    """
    A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.

      * `content_type` (`str`) - Content type of the message.
      * `correlationId` (`str`) - Identifier of the correlation.
      * `label` (`str`) - Application specific label.
      * `messageId` (`str`) - Identifier of the message.
      * `replyTo` (`str`) - Address of the queue to reply to.
      * `replyToSessionId` (`str`) - Session identifier to reply to.
      * `sessionId` (`str`) - Session identifier.
      * `to` (`str`) - Address to send to.
    """
    filter_type: pulumi.Output[str]
    """
    Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
    """
    namespace_name: pulumi.Output[str]
    """
    The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
    """
    sql_filter: pulumi.Output[str]
    """
    Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
    """
    subscription_name: pulumi.Output[str]
    """
    The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
    """
    topic_name: pulumi.Output[str]
    """
    The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
    """
    warnings.warn("azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule", DeprecationWarning)

    def __init__(__self__, resource_name, opts=None, action=None, correlation_filter=None, filter_type=None, name=None, namespace_name=None, resource_group_name=None, sql_filter=None, subscription_name=None, topic_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a ServiceBus Subscription Rule.

        ## Example Usage
        ### SQL Filter)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_subscription = azure.servicebus.Subscription("exampleSubscription",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            max_delivery_count=1)
        example_subscription_rule = azure.servicebus.SubscriptionRule("exampleSubscriptionRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            subscription_name=example_subscription.name,
            filter_type="SqlFilter",
            sql_filter="colour = 'red'")
        ```
        ### Correlation Filter)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_subscription = azure.servicebus.Subscription("exampleSubscription",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            max_delivery_count=1)
        example_subscription_rule = azure.servicebus.SubscriptionRule("exampleSubscriptionRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            subscription_name=example_subscription.name,
            filter_type="CorrelationFilter",
            correlation_filter={
                "correlationId": "high",
                "label": "red",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        :param pulumi.Input[dict] correlation_filter: A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        :param pulumi.Input[str] filter_type: Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_filter: Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        :param pulumi.Input[str] subscription_name: The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.

        The **correlation_filter** object supports the following:

          * `content_type` (`pulumi.Input[str]`) - Content type of the message.
          * `correlationId` (`pulumi.Input[str]`) - Identifier of the correlation.
          * `label` (`pulumi.Input[str]`) - Application specific label.
          * `messageId` (`pulumi.Input[str]`) - Identifier of the message.
          * `replyTo` (`pulumi.Input[str]`) - Address of the queue to reply to.
          * `replyToSessionId` (`pulumi.Input[str]`) - Session identifier to reply to.
          * `sessionId` (`pulumi.Input[str]`) - Session identifier.
          * `to` (`pulumi.Input[str]`) - Address to send to.
        """
        pulumi.log.warn("SubscriptionRule is deprecated: azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule")
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

            __props__['action'] = action
            __props__['correlation_filter'] = correlation_filter
            if filter_type is None:
                raise TypeError("Missing required property 'filter_type'")
            __props__['filter_type'] = filter_type
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sql_filter'] = sql_filter
            if subscription_name is None:
                raise TypeError("Missing required property 'subscription_name'")
            __props__['subscription_name'] = subscription_name
            if topic_name is None:
                raise TypeError("Missing required property 'topic_name'")
            __props__['topic_name'] = topic_name
        super(SubscriptionRule, __self__).__init__(
            'azure:eventhub/subscriptionRule:SubscriptionRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, action=None, correlation_filter=None, filter_type=None, name=None, namespace_name=None, resource_group_name=None, sql_filter=None, subscription_name=None, topic_name=None):
        """
        Get an existing SubscriptionRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        :param pulumi.Input[dict] correlation_filter: A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        :param pulumi.Input[str] filter_type: Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_filter: Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        :param pulumi.Input[str] subscription_name: The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.

        The **correlation_filter** object supports the following:

          * `content_type` (`pulumi.Input[str]`) - Content type of the message.
          * `correlationId` (`pulumi.Input[str]`) - Identifier of the correlation.
          * `label` (`pulumi.Input[str]`) - Application specific label.
          * `messageId` (`pulumi.Input[str]`) - Identifier of the message.
          * `replyTo` (`pulumi.Input[str]`) - Address of the queue to reply to.
          * `replyToSessionId` (`pulumi.Input[str]`) - Session identifier to reply to.
          * `sessionId` (`pulumi.Input[str]`) - Session identifier.
          * `to` (`pulumi.Input[str]`) - Address to send to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["correlation_filter"] = correlation_filter
        __props__["filter_type"] = filter_type
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["resource_group_name"] = resource_group_name
        __props__["sql_filter"] = sql_filter
        __props__["subscription_name"] = subscription_name
        __props__["topic_name"] = topic_name
        return SubscriptionRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
