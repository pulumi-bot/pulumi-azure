# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['EventSubscription']


class EventSubscription(pulumi.CustomResource):
    advanced_filter: pulumi.Output[Optional['outputs.EventSubscriptionAdvancedFilter']] = pulumi.output_property("advancedFilter")
    """
    A `advanced_filter` block as defined below.
    """
    azure_function_endpoint: pulumi.Output[Optional['outputs.EventSubscriptionAzureFunctionEndpoint']] = pulumi.output_property("azureFunctionEndpoint")
    """
    An `azure_function_endpoint` block as defined below.
    """
    event_delivery_schema: pulumi.Output[Optional[str]] = pulumi.output_property("eventDeliverySchema")
    """
    Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
    """
    eventhub_endpoint: pulumi.Output['outputs.EventSubscriptionEventhubEndpoint'] = pulumi.output_property("eventhubEndpoint")
    """
    A `eventhub_endpoint` block as defined below.
    """
    eventhub_endpoint_id: pulumi.Output[str] = pulumi.output_property("eventhubEndpointId")
    """
    Specifies the id where the Event Hub is located.
    """
    expiration_time_utc: pulumi.Output[Optional[str]] = pulumi.output_property("expirationTimeUtc")
    """
    Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
    """
    hybrid_connection_endpoint: pulumi.Output['outputs.EventSubscriptionHybridConnectionEndpoint'] = pulumi.output_property("hybridConnectionEndpoint")
    """
    A `hybrid_connection_endpoint` block as defined below.
    """
    hybrid_connection_endpoint_id: pulumi.Output[str] = pulumi.output_property("hybridConnectionEndpointId")
    """
    Specifies the id where the Hybrid Connection is located.
    """
    included_event_types: pulumi.Output[List[str]] = pulumi.output_property("includedEventTypes")
    """
    A list of applicable event types that need to be part of the event subscription.
    """
    labels: pulumi.Output[Optional[List[str]]] = pulumi.output_property("labels")
    """
    A list of labels to assign to the event subscription.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
    """
    retry_policy: pulumi.Output['outputs.EventSubscriptionRetryPolicy'] = pulumi.output_property("retryPolicy")
    """
    A `retry_policy` block as defined below.
    """
    scope: pulumi.Output[str] = pulumi.output_property("scope")
    """
    Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
    """
    service_bus_queue_endpoint_id: pulumi.Output[Optional[str]] = pulumi.output_property("serviceBusQueueEndpointId")
    """
    Specifies the id where the Service Bus Queue is located.
    """
    service_bus_topic_endpoint_id: pulumi.Output[Optional[str]] = pulumi.output_property("serviceBusTopicEndpointId")
    """
    Specifies the id where the Service Bus Topic is located.
    """
    storage_blob_dead_letter_destination: pulumi.Output[Optional['outputs.EventSubscriptionStorageBlobDeadLetterDestination']] = pulumi.output_property("storageBlobDeadLetterDestination")
    """
    A `storage_blob_dead_letter_destination` block as defined below.
    """
    storage_queue_endpoint: pulumi.Output[Optional['outputs.EventSubscriptionStorageQueueEndpoint']] = pulumi.output_property("storageQueueEndpoint")
    """
    A `storage_queue_endpoint` block as defined below.
    """
    subject_filter: pulumi.Output[Optional['outputs.EventSubscriptionSubjectFilter']] = pulumi.output_property("subjectFilter")
    """
    A `subject_filter` block as defined below.
    """
    topic_name: pulumi.Output[str] = pulumi.output_property("topicName")
    """
    (Optional) Specifies the name of the topic to associate with the event subscription.
    """
    webhook_endpoint: pulumi.Output[Optional['outputs.EventSubscriptionWebhookEndpoint']] = pulumi.output_property("webhookEndpoint")
    """
    A `webhook_endpoint` block as defined below.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, advanced_filter: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionAdvancedFilterArgs']]] = None, azure_function_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionAzureFunctionEndpointArgs']]] = None, event_delivery_schema: Optional[pulumi.Input[str]] = None, eventhub_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionEventhubEndpointArgs']]] = None, eventhub_endpoint_id: Optional[pulumi.Input[str]] = None, expiration_time_utc: Optional[pulumi.Input[str]] = None, hybrid_connection_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionHybridConnectionEndpointArgs']]] = None, hybrid_connection_endpoint_id: Optional[pulumi.Input[str]] = None, included_event_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, labels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, retry_policy: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionRetryPolicyArgs']]] = None, scope: Optional[pulumi.Input[str]] = None, service_bus_queue_endpoint_id: Optional[pulumi.Input[str]] = None, service_bus_topic_endpoint_id: Optional[pulumi.Input[str]] = None, storage_blob_dead_letter_destination: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionStorageBlobDeadLetterDestinationArgs']]] = None, storage_queue_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionStorageQueueEndpointArgs']]] = None, subject_filter: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionSubjectFilterArgs']]] = None, topic_name: Optional[pulumi.Input[str]] = None, webhook_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionWebhookEndpointArgs']]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an EventGrid Event Subscription

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        default_resource_group = azure.core.ResourceGroup("defaultResourceGroup", location="West US 2")
        default_account = azure.storage.Account("defaultAccount",
            resource_group_name=default_resource_group.name,
            location=default_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            tags={
                "environment": "staging",
            })
        default_queue = azure.storage.Queue("defaultQueue", storage_account_name=default_account.name)
        default_event_subscription = azure.eventgrid.EventSubscription("defaultEventSubscription",
            scope=default_resource_group.id,
            storage_queue_endpoint={
                "storage_account_id": default_account.id,
                "queue_name": default_queue.name,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionAdvancedFilterArgs']] advanced_filter: A `advanced_filter` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionAzureFunctionEndpointArgs']] azure_function_endpoint: An `azure_function_endpoint` block as defined below.
        :param pulumi.Input[str] event_delivery_schema: Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionEventhubEndpointArgs']] eventhub_endpoint: A `eventhub_endpoint` block as defined below.
        :param pulumi.Input[str] eventhub_endpoint_id: Specifies the id where the Event Hub is located.
        :param pulumi.Input[str] expiration_time_utc: Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        :param pulumi.Input[pulumi.InputType['EventSubscriptionHybridConnectionEndpointArgs']] hybrid_connection_endpoint: A `hybrid_connection_endpoint` block as defined below.
        :param pulumi.Input[str] hybrid_connection_endpoint_id: Specifies the id where the Hybrid Connection is located.
        :param pulumi.Input[List[pulumi.Input[str]]] included_event_types: A list of applicable event types that need to be part of the event subscription.
        :param pulumi.Input[List[pulumi.Input[str]]] labels: A list of labels to assign to the event subscription.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionRetryPolicyArgs']] retry_policy: A `retry_policy` block as defined below.
        :param pulumi.Input[str] scope: Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_bus_queue_endpoint_id: Specifies the id where the Service Bus Queue is located.
        :param pulumi.Input[str] service_bus_topic_endpoint_id: Specifies the id where the Service Bus Topic is located.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionStorageBlobDeadLetterDestinationArgs']] storage_blob_dead_letter_destination: A `storage_blob_dead_letter_destination` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionStorageQueueEndpointArgs']] storage_queue_endpoint: A `storage_queue_endpoint` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionSubjectFilterArgs']] subject_filter: A `subject_filter` block as defined below.
        :param pulumi.Input[str] topic_name: (Optional) Specifies the name of the topic to associate with the event subscription.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionWebhookEndpointArgs']] webhook_endpoint: A `webhook_endpoint` block as defined below.
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

            __props__['advanced_filter'] = advanced_filter
            __props__['azure_function_endpoint'] = azure_function_endpoint
            __props__['event_delivery_schema'] = event_delivery_schema
            if eventhub_endpoint is not None:
                warnings.warn("Deprecated in favour of `eventhub_endpoint_id`", DeprecationWarning)
                pulumi.log.warn("eventhub_endpoint is deprecated: Deprecated in favour of `eventhub_endpoint_id`")
            __props__['eventhub_endpoint'] = eventhub_endpoint
            __props__['eventhub_endpoint_id'] = eventhub_endpoint_id
            __props__['expiration_time_utc'] = expiration_time_utc
            if hybrid_connection_endpoint is not None:
                warnings.warn("Deprecated in favour of `hybrid_connection_endpoint_id`", DeprecationWarning)
                pulumi.log.warn("hybrid_connection_endpoint is deprecated: Deprecated in favour of `hybrid_connection_endpoint_id`")
            __props__['hybrid_connection_endpoint'] = hybrid_connection_endpoint
            __props__['hybrid_connection_endpoint_id'] = hybrid_connection_endpoint_id
            __props__['included_event_types'] = included_event_types
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['retry_policy'] = retry_policy
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['service_bus_queue_endpoint_id'] = service_bus_queue_endpoint_id
            __props__['service_bus_topic_endpoint_id'] = service_bus_topic_endpoint_id
            __props__['storage_blob_dead_letter_destination'] = storage_blob_dead_letter_destination
            __props__['storage_queue_endpoint'] = storage_queue_endpoint
            __props__['subject_filter'] = subject_filter
            __props__['topic_name'] = topic_name
            __props__['webhook_endpoint'] = webhook_endpoint
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/eventSubscription:EventSubscription")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(EventSubscription, __self__).__init__(
            'azure:eventgrid/eventSubscription:EventSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, advanced_filter: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionAdvancedFilterArgs']]] = None, azure_function_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionAzureFunctionEndpointArgs']]] = None, event_delivery_schema: Optional[pulumi.Input[str]] = None, eventhub_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionEventhubEndpointArgs']]] = None, eventhub_endpoint_id: Optional[pulumi.Input[str]] = None, expiration_time_utc: Optional[pulumi.Input[str]] = None, hybrid_connection_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionHybridConnectionEndpointArgs']]] = None, hybrid_connection_endpoint_id: Optional[pulumi.Input[str]] = None, included_event_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, labels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, retry_policy: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionRetryPolicyArgs']]] = None, scope: Optional[pulumi.Input[str]] = None, service_bus_queue_endpoint_id: Optional[pulumi.Input[str]] = None, service_bus_topic_endpoint_id: Optional[pulumi.Input[str]] = None, storage_blob_dead_letter_destination: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionStorageBlobDeadLetterDestinationArgs']]] = None, storage_queue_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionStorageQueueEndpointArgs']]] = None, subject_filter: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionSubjectFilterArgs']]] = None, topic_name: Optional[pulumi.Input[str]] = None, webhook_endpoint: Optional[pulumi.Input[pulumi.InputType['EventSubscriptionWebhookEndpointArgs']]] = None) -> 'EventSubscription':
        """
        Get an existing EventSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionAdvancedFilterArgs']] advanced_filter: A `advanced_filter` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionAzureFunctionEndpointArgs']] azure_function_endpoint: An `azure_function_endpoint` block as defined below.
        :param pulumi.Input[str] event_delivery_schema: Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionEventhubEndpointArgs']] eventhub_endpoint: A `eventhub_endpoint` block as defined below.
        :param pulumi.Input[str] eventhub_endpoint_id: Specifies the id where the Event Hub is located.
        :param pulumi.Input[str] expiration_time_utc: Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        :param pulumi.Input[pulumi.InputType['EventSubscriptionHybridConnectionEndpointArgs']] hybrid_connection_endpoint: A `hybrid_connection_endpoint` block as defined below.
        :param pulumi.Input[str] hybrid_connection_endpoint_id: Specifies the id where the Hybrid Connection is located.
        :param pulumi.Input[List[pulumi.Input[str]]] included_event_types: A list of applicable event types that need to be part of the event subscription.
        :param pulumi.Input[List[pulumi.Input[str]]] labels: A list of labels to assign to the event subscription.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionRetryPolicyArgs']] retry_policy: A `retry_policy` block as defined below.
        :param pulumi.Input[str] scope: Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_bus_queue_endpoint_id: Specifies the id where the Service Bus Queue is located.
        :param pulumi.Input[str] service_bus_topic_endpoint_id: Specifies the id where the Service Bus Topic is located.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionStorageBlobDeadLetterDestinationArgs']] storage_blob_dead_letter_destination: A `storage_blob_dead_letter_destination` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionStorageQueueEndpointArgs']] storage_queue_endpoint: A `storage_queue_endpoint` block as defined below.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionSubjectFilterArgs']] subject_filter: A `subject_filter` block as defined below.
        :param pulumi.Input[str] topic_name: (Optional) Specifies the name of the topic to associate with the event subscription.
        :param pulumi.Input[pulumi.InputType['EventSubscriptionWebhookEndpointArgs']] webhook_endpoint: A `webhook_endpoint` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["advanced_filter"] = advanced_filter
        __props__["azure_function_endpoint"] = azure_function_endpoint
        __props__["event_delivery_schema"] = event_delivery_schema
        __props__["eventhub_endpoint"] = eventhub_endpoint
        __props__["eventhub_endpoint_id"] = eventhub_endpoint_id
        __props__["expiration_time_utc"] = expiration_time_utc
        __props__["hybrid_connection_endpoint"] = hybrid_connection_endpoint
        __props__["hybrid_connection_endpoint_id"] = hybrid_connection_endpoint_id
        __props__["included_event_types"] = included_event_types
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["retry_policy"] = retry_policy
        __props__["scope"] = scope
        __props__["service_bus_queue_endpoint_id"] = service_bus_queue_endpoint_id
        __props__["service_bus_topic_endpoint_id"] = service_bus_topic_endpoint_id
        __props__["storage_blob_dead_letter_destination"] = storage_blob_dead_letter_destination
        __props__["storage_queue_endpoint"] = storage_queue_endpoint
        __props__["subject_filter"] = subject_filter
        __props__["topic_name"] = topic_name
        __props__["webhook_endpoint"] = webhook_endpoint
        return EventSubscription(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

