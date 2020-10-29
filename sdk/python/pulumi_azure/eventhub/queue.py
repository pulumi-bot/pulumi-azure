# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Queue']

warnings.warn("azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue", DeprecationWarning)


class Queue(pulumi.CustomResource):
    warnings.warn("azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_delete_on_idle: Optional[pulumi.Input[str]] = None,
                 dead_lettering_on_message_expiration: Optional[pulumi.Input[bool]] = None,
                 default_message_ttl: Optional[pulumi.Input[str]] = None,
                 duplicate_detection_history_time_window: Optional[pulumi.Input[str]] = None,
                 enable_batched_operations: Optional[pulumi.Input[bool]] = None,
                 enable_express: Optional[pulumi.Input[bool]] = None,
                 enable_partitioning: Optional[pulumi.Input[bool]] = None,
                 forward_dead_lettered_messages_to: Optional[pulumi.Input[str]] = None,
                 forward_to: Optional[pulumi.Input[str]] = None,
                 lock_duration: Optional[pulumi.Input[str]] = None,
                 max_delivery_count: Optional[pulumi.Input[int]] = None,
                 max_size_in_megabytes: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 requires_duplicate_detection: Optional[pulumi.Input[bool]] = None,
                 requires_session: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a ServiceBus Queue.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[bool] dead_lettering_on_message_expiration: Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
        :param pulumi.Input[str] forward_dead_lettered_messages_to: The name of a Queue or Topic to automatically forward dead lettered messages to.
        :param pulumi.Input[str] forward_to: The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
        :param pulumi.Input[str] lock_duration: The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
        :param pulumi.Input[int] max_delivery_count: Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
        :param pulumi.Input[int] max_size_in_megabytes: Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[bool] requires_session: Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
        """
        pulumi.log.warn("Queue is deprecated: azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue")
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

            __props__['auto_delete_on_idle'] = auto_delete_on_idle
            __props__['dead_lettering_on_message_expiration'] = dead_lettering_on_message_expiration
            __props__['default_message_ttl'] = default_message_ttl
            __props__['duplicate_detection_history_time_window'] = duplicate_detection_history_time_window
            __props__['enable_batched_operations'] = enable_batched_operations
            __props__['enable_express'] = enable_express
            __props__['enable_partitioning'] = enable_partitioning
            __props__['forward_dead_lettered_messages_to'] = forward_dead_lettered_messages_to
            __props__['forward_to'] = forward_to
            __props__['lock_duration'] = lock_duration
            __props__['max_delivery_count'] = max_delivery_count
            __props__['max_size_in_megabytes'] = max_size_in_megabytes
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['requires_duplicate_detection'] = requires_duplicate_detection
            __props__['requires_session'] = requires_session
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
        super(Queue, __self__).__init__(
            'azure:eventhub/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_delete_on_idle: Optional[pulumi.Input[str]] = None,
            dead_lettering_on_message_expiration: Optional[pulumi.Input[bool]] = None,
            default_message_ttl: Optional[pulumi.Input[str]] = None,
            duplicate_detection_history_time_window: Optional[pulumi.Input[str]] = None,
            enable_batched_operations: Optional[pulumi.Input[bool]] = None,
            enable_express: Optional[pulumi.Input[bool]] = None,
            enable_partitioning: Optional[pulumi.Input[bool]] = None,
            forward_dead_lettered_messages_to: Optional[pulumi.Input[str]] = None,
            forward_to: Optional[pulumi.Input[str]] = None,
            lock_duration: Optional[pulumi.Input[str]] = None,
            max_delivery_count: Optional[pulumi.Input[int]] = None,
            max_size_in_megabytes: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            requires_duplicate_detection: Optional[pulumi.Input[bool]] = None,
            requires_session: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Queue':
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[bool] dead_lettering_on_message_expiration: Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
        :param pulumi.Input[str] forward_dead_lettered_messages_to: The name of a Queue or Topic to automatically forward dead lettered messages to.
        :param pulumi.Input[str] forward_to: The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
        :param pulumi.Input[str] lock_duration: The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
        :param pulumi.Input[int] max_delivery_count: Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
        :param pulumi.Input[int] max_size_in_megabytes: Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[bool] requires_session: Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_delete_on_idle"] = auto_delete_on_idle
        __props__["dead_lettering_on_message_expiration"] = dead_lettering_on_message_expiration
        __props__["default_message_ttl"] = default_message_ttl
        __props__["duplicate_detection_history_time_window"] = duplicate_detection_history_time_window
        __props__["enable_batched_operations"] = enable_batched_operations
        __props__["enable_express"] = enable_express
        __props__["enable_partitioning"] = enable_partitioning
        __props__["forward_dead_lettered_messages_to"] = forward_dead_lettered_messages_to
        __props__["forward_to"] = forward_to
        __props__["lock_duration"] = lock_duration
        __props__["max_delivery_count"] = max_delivery_count
        __props__["max_size_in_megabytes"] = max_size_in_megabytes
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["requires_duplicate_detection"] = requires_duplicate_detection
        __props__["requires_session"] = requires_session
        __props__["resource_group_name"] = resource_group_name
        __props__["status"] = status
        return Queue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDeleteOnIdle")
    def auto_delete_on_idle(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
        """
        return pulumi.get(self, "auto_delete_on_idle")

    @property
    @pulumi.getter(name="deadLetteringOnMessageExpiration")
    def dead_lettering_on_message_expiration(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        """
        return pulumi.get(self, "dead_lettering_on_message_expiration")

    @property
    @pulumi.getter(name="defaultMessageTtl")
    def default_message_ttl(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
        """
        return pulumi.get(self, "default_message_ttl")

    @property
    @pulumi.getter(name="duplicateDetectionHistoryTimeWindow")
    def duplicate_detection_history_time_window(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
        """
        return pulumi.get(self, "duplicate_detection_history_time_window")

    @property
    @pulumi.getter(name="enableBatchedOperations")
    def enable_batched_operations(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enable_batched_operations")

    @property
    @pulumi.getter(name="enableExpress")
    def enable_express(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
        """
        return pulumi.get(self, "enable_express")

    @property
    @pulumi.getter(name="enablePartitioning")
    def enable_partitioning(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
        """
        return pulumi.get(self, "enable_partitioning")

    @property
    @pulumi.getter(name="forwardDeadLetteredMessagesTo")
    def forward_dead_lettered_messages_to(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a Queue or Topic to automatically forward dead lettered messages to.
        """
        return pulumi.get(self, "forward_dead_lettered_messages_to")

    @property
    @pulumi.getter(name="forwardTo")
    def forward_to(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
        """
        return pulumi.get(self, "forward_to")

    @property
    @pulumi.getter(name="lockDuration")
    def lock_duration(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
        """
        return pulumi.get(self, "lock_duration")

    @property
    @pulumi.getter(name="maxDeliveryCount")
    def max_delivery_count(self) -> pulumi.Output[Optional[int]]:
        """
        Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
        """
        return pulumi.get(self, "max_delivery_count")

    @property
    @pulumi.getter(name="maxSizeInMegabytes")
    def max_size_in_megabytes(self) -> pulumi.Output[int]:
        """
        Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
        """
        return pulumi.get(self, "max_size_in_megabytes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="requiresDuplicateDetection")
    def requires_duplicate_detection(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
        """
        return pulumi.get(self, "requires_duplicate_detection")

    @property
    @pulumi.getter(name="requiresSession")
    def requires_session(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
        """
        return pulumi.get(self, "requires_session")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

