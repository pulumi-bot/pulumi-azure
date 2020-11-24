# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .authorization_rule import *
from .cluster import *
from .consumer_group import *
from .domain import *
from .event_grid_topic import *
from .event_hub import *
from .event_hub_authorization_rule import *
from .event_hub_consumer_group import *
from .event_hub_namespace import *
from .event_hub_namespace_authorization_rule import *
from .event_subscription import *
from .eventhub_namespace_disaster_recovery_config import *
from .get_authorization_rule import *
from .get_consume_group import *
from .get_event_hub import *
from .get_eventhub_namespace import *
from .get_namespace import *
from .get_namespace_authorization_rule import *
from .get_service_bus_namespace import *
from .namespace import *
from .namespace_authorization_rule import *
from .queue import *
from .queue_authorization_rule import *
from .subscription import *
from .subscription_rule import *
from .topic import *
from .topic_authorization_rule import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:eventhub/authorizationRule:AuthorizationRule":
                return AuthorizationRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/consumerGroup:ConsumerGroup":
                return ConsumerGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/domain:Domain":
                return Domain(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventGridTopic:EventGridTopic":
                return EventGridTopic(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventHub:EventHub":
                return EventHub(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule":
                return EventHubAuthorizationRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup":
                return EventHubConsumerGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventHubNamespace:EventHubNamespace":
                return EventHubNamespace(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventHubNamespaceAuthorizationRule:EventHubNamespaceAuthorizationRule":
                return EventHubNamespaceAuthorizationRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventSubscription:EventSubscription":
                return EventSubscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/eventhubNamespaceDisasterRecoveryConfig:EventhubNamespaceDisasterRecoveryConfig":
                return EventhubNamespaceDisasterRecoveryConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/namespace:Namespace":
                return Namespace(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/namespaceAuthorizationRule:NamespaceAuthorizationRule":
                return NamespaceAuthorizationRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/queue:Queue":
                return Queue(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/queueAuthorizationRule:QueueAuthorizationRule":
                return QueueAuthorizationRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/subscription:Subscription":
                return Subscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/subscriptionRule:SubscriptionRule":
                return SubscriptionRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/topic:Topic":
                return Topic(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule":
                return TopicAuthorizationRule(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "eventhub/authorizationRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/cluster", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/consumerGroup", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/domain", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventGridTopic", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventHub", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventHubAuthorizationRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventHubConsumerGroup", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventHubNamespace", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventHubNamespaceAuthorizationRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventSubscription", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/eventhubNamespaceDisasterRecoveryConfig", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/namespace", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/namespaceAuthorizationRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/queue", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/queueAuthorizationRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/subscription", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/subscriptionRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/topic", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventhub/topicAuthorizationRule", _module_instance)

_register_module()
