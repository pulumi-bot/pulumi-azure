# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['TopicArgs', 'Topic']

@pulumi.input_type
class TopicArgs:
    def __init__(__self__, *,
                 namespace_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 auto_delete_on_idle: Optional[pulumi.Input[str]] = None,
                 default_message_ttl: Optional[pulumi.Input[str]] = None,
                 duplicate_detection_history_time_window: Optional[pulumi.Input[str]] = None,
                 enable_batched_operations: Optional[pulumi.Input[bool]] = None,
                 enable_express: Optional[pulumi.Input[bool]] = None,
                 enable_partitioning: Optional[pulumi.Input[bool]] = None,
                 max_size_in_megabytes: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 requires_duplicate_detection: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 support_ordering: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Topic resource.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create
               this topic in. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the
               Topic is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of TTL of messages sent to this topic if no
               TTL value is set on the message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which
               duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls if server-side
               batched operations are enabled. Defaults to false.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities
               are enabled. An express topic holds a message in memory temporarily before writing
               it to persistent storage. Defaults to false.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable
               the topic to be partitioned across multiple message brokers. Defaults to false.
               Changing this forces a new resource to be created.
        :param pulumi.Input[int] max_size_in_megabytes: Integer value which controls the size of
               memory allocated for the topic. For supported values see the "Queue/topic size"
               section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Topic resource. Changing this forces a
               new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether
               the Topic requires duplicate detection. Defaults to false. Changing this forces
               a new resource to be created.
        :param pulumi.Input[str] status: The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        :param pulumi.Input[bool] support_ordering: Boolean flag which controls whether the Topic
               supports ordering. Defaults to false.
        """
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if auto_delete_on_idle is not None:
            pulumi.set(__self__, "auto_delete_on_idle", auto_delete_on_idle)
        if default_message_ttl is not None:
            pulumi.set(__self__, "default_message_ttl", default_message_ttl)
        if duplicate_detection_history_time_window is not None:
            pulumi.set(__self__, "duplicate_detection_history_time_window", duplicate_detection_history_time_window)
        if enable_batched_operations is not None:
            pulumi.set(__self__, "enable_batched_operations", enable_batched_operations)
        if enable_express is not None:
            pulumi.set(__self__, "enable_express", enable_express)
        if enable_partitioning is not None:
            pulumi.set(__self__, "enable_partitioning", enable_partitioning)
        if max_size_in_megabytes is not None:
            pulumi.set(__self__, "max_size_in_megabytes", max_size_in_megabytes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if requires_duplicate_detection is not None:
            pulumi.set(__self__, "requires_duplicate_detection", requires_duplicate_detection)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if support_ordering is not None:
            pulumi.set(__self__, "support_ordering", support_ordering)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        The name of the ServiceBus Namespace to create
        this topic in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to
        create the namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="autoDeleteOnIdle")
    def auto_delete_on_idle(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO 8601 timespan duration of the idle interval after which the
        Topic is automatically deleted, minimum of 5 minutes.
        """
        return pulumi.get(self, "auto_delete_on_idle")

    @auto_delete_on_idle.setter
    def auto_delete_on_idle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_delete_on_idle", value)

    @property
    @pulumi.getter(name="defaultMessageTtl")
    def default_message_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO 8601 timespan duration of TTL of messages sent to this topic if no
        TTL value is set on the message itself.
        """
        return pulumi.get(self, "default_message_ttl")

    @default_message_ttl.setter
    def default_message_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_message_ttl", value)

    @property
    @pulumi.getter(name="duplicateDetectionHistoryTimeWindow")
    def duplicate_detection_history_time_window(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO 8601 timespan duration during which
        duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        """
        return pulumi.get(self, "duplicate_detection_history_time_window")

    @duplicate_detection_history_time_window.setter
    def duplicate_detection_history_time_window(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duplicate_detection_history_time_window", value)

    @property
    @pulumi.getter(name="enableBatchedOperations")
    def enable_batched_operations(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag which controls if server-side
        batched operations are enabled. Defaults to false.
        """
        return pulumi.get(self, "enable_batched_operations")

    @enable_batched_operations.setter
    def enable_batched_operations(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_batched_operations", value)

    @property
    @pulumi.getter(name="enableExpress")
    def enable_express(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag which controls whether Express Entities
        are enabled. An express topic holds a message in memory temporarily before writing
        it to persistent storage. Defaults to false.
        """
        return pulumi.get(self, "enable_express")

    @enable_express.setter
    def enable_express(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_express", value)

    @property
    @pulumi.getter(name="enablePartitioning")
    def enable_partitioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag which controls whether to enable
        the topic to be partitioned across multiple message brokers. Defaults to false.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "enable_partitioning")

    @enable_partitioning.setter
    def enable_partitioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_partitioning", value)

    @property
    @pulumi.getter(name="maxSizeInMegabytes")
    def max_size_in_megabytes(self) -> Optional[pulumi.Input[int]]:
        """
        Integer value which controls the size of
        memory allocated for the topic. For supported values see the "Queue/topic size"
        section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        """
        return pulumi.get(self, "max_size_in_megabytes")

    @max_size_in_megabytes.setter
    def max_size_in_megabytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_size_in_megabytes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the ServiceBus Topic resource. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requiresDuplicateDetection")
    def requires_duplicate_detection(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag which controls whether
        the Topic requires duplicate detection. Defaults to false. Changing this forces
        a new resource to be created.
        """
        return pulumi.get(self, "requires_duplicate_detection")

    @requires_duplicate_detection.setter
    def requires_duplicate_detection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requires_duplicate_detection", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="supportOrdering")
    def support_ordering(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag which controls whether the Topic
        supports ordering. Defaults to false.
        """
        return pulumi.get(self, "support_ordering")

    @support_ordering.setter
    def support_ordering(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "support_ordering", value)


class Topic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_delete_on_idle: Optional[pulumi.Input[str]] = None,
                 default_message_ttl: Optional[pulumi.Input[str]] = None,
                 duplicate_detection_history_time_window: Optional[pulumi.Input[str]] = None,
                 enable_batched_operations: Optional[pulumi.Input[bool]] = None,
                 enable_express: Optional[pulumi.Input[bool]] = None,
                 enable_partitioning: Optional[pulumi.Input[bool]] = None,
                 max_size_in_megabytes: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 requires_duplicate_detection: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 support_ordering: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a ServiceBus Topic.

        **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.

        ## Example Usage

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
        ```

        ## Import

        Service Bus Topics can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:servicebus/topic:Topic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the
               Topic is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of TTL of messages sent to this topic if no
               TTL value is set on the message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which
               duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls if server-side
               batched operations are enabled. Defaults to false.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities
               are enabled. An express topic holds a message in memory temporarily before writing
               it to persistent storage. Defaults to false.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable
               the topic to be partitioned across multiple message brokers. Defaults to false.
               Changing this forces a new resource to be created.
        :param pulumi.Input[int] max_size_in_megabytes: Integer value which controls the size of
               memory allocated for the topic. For supported values see the "Queue/topic size"
               section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Topic resource. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create
               this topic in. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether
               the Topic requires duplicate detection. Defaults to false. Changing this forces
               a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        :param pulumi.Input[bool] support_ordering: Boolean flag which controls whether the Topic
               supports ordering. Defaults to false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a ServiceBus Topic.

        **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.

        ## Example Usage

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
        ```

        ## Import

        Service Bus Topics can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:servicebus/topic:Topic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1
        ```

        :param str resource_name: The name of the resource.
        :param TopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_delete_on_idle: Optional[pulumi.Input[str]] = None,
                 default_message_ttl: Optional[pulumi.Input[str]] = None,
                 duplicate_detection_history_time_window: Optional[pulumi.Input[str]] = None,
                 enable_batched_operations: Optional[pulumi.Input[bool]] = None,
                 enable_express: Optional[pulumi.Input[bool]] = None,
                 enable_partitioning: Optional[pulumi.Input[bool]] = None,
                 max_size_in_megabytes: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 requires_duplicate_detection: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 support_ordering: Optional[pulumi.Input[bool]] = None,
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

            __props__['auto_delete_on_idle'] = auto_delete_on_idle
            __props__['default_message_ttl'] = default_message_ttl
            __props__['duplicate_detection_history_time_window'] = duplicate_detection_history_time_window
            __props__['enable_batched_operations'] = enable_batched_operations
            __props__['enable_express'] = enable_express
            __props__['enable_partitioning'] = enable_partitioning
            __props__['max_size_in_megabytes'] = max_size_in_megabytes
            __props__['name'] = name
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['requires_duplicate_detection'] = requires_duplicate_detection
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
            __props__['support_ordering'] = support_ordering
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/topic:Topic")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Topic, __self__).__init__(
            'azure:servicebus/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_delete_on_idle: Optional[pulumi.Input[str]] = None,
            default_message_ttl: Optional[pulumi.Input[str]] = None,
            duplicate_detection_history_time_window: Optional[pulumi.Input[str]] = None,
            enable_batched_operations: Optional[pulumi.Input[bool]] = None,
            enable_express: Optional[pulumi.Input[bool]] = None,
            enable_partitioning: Optional[pulumi.Input[bool]] = None,
            max_size_in_megabytes: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            requires_duplicate_detection: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            support_ordering: Optional[pulumi.Input[bool]] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the
               Topic is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of TTL of messages sent to this topic if no
               TTL value is set on the message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which
               duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls if server-side
               batched operations are enabled. Defaults to false.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities
               are enabled. An express topic holds a message in memory temporarily before writing
               it to persistent storage. Defaults to false.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable
               the topic to be partitioned across multiple message brokers. Defaults to false.
               Changing this forces a new resource to be created.
        :param pulumi.Input[int] max_size_in_megabytes: Integer value which controls the size of
               memory allocated for the topic. For supported values see the "Queue/topic size"
               section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Topic resource. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create
               this topic in. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether
               the Topic requires duplicate detection. Defaults to false. Changing this forces
               a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        :param pulumi.Input[bool] support_ordering: Boolean flag which controls whether the Topic
               supports ordering. Defaults to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_delete_on_idle"] = auto_delete_on_idle
        __props__["default_message_ttl"] = default_message_ttl
        __props__["duplicate_detection_history_time_window"] = duplicate_detection_history_time_window
        __props__["enable_batched_operations"] = enable_batched_operations
        __props__["enable_express"] = enable_express
        __props__["enable_partitioning"] = enable_partitioning
        __props__["max_size_in_megabytes"] = max_size_in_megabytes
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["requires_duplicate_detection"] = requires_duplicate_detection
        __props__["resource_group_name"] = resource_group_name
        __props__["status"] = status
        __props__["support_ordering"] = support_ordering
        return Topic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDeleteOnIdle")
    def auto_delete_on_idle(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration of the idle interval after which the
        Topic is automatically deleted, minimum of 5 minutes.
        """
        return pulumi.get(self, "auto_delete_on_idle")

    @property
    @pulumi.getter(name="defaultMessageTtl")
    def default_message_ttl(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration of TTL of messages sent to this topic if no
        TTL value is set on the message itself.
        """
        return pulumi.get(self, "default_message_ttl")

    @property
    @pulumi.getter(name="duplicateDetectionHistoryTimeWindow")
    def duplicate_detection_history_time_window(self) -> pulumi.Output[str]:
        """
        The ISO 8601 timespan duration during which
        duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        """
        return pulumi.get(self, "duplicate_detection_history_time_window")

    @property
    @pulumi.getter(name="enableBatchedOperations")
    def enable_batched_operations(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls if server-side
        batched operations are enabled. Defaults to false.
        """
        return pulumi.get(self, "enable_batched_operations")

    @property
    @pulumi.getter(name="enableExpress")
    def enable_express(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether Express Entities
        are enabled. An express topic holds a message in memory temporarily before writing
        it to persistent storage. Defaults to false.
        """
        return pulumi.get(self, "enable_express")

    @property
    @pulumi.getter(name="enablePartitioning")
    def enable_partitioning(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether to enable
        the topic to be partitioned across multiple message brokers. Defaults to false.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "enable_partitioning")

    @property
    @pulumi.getter(name="maxSizeInMegabytes")
    def max_size_in_megabytes(self) -> pulumi.Output[int]:
        """
        Integer value which controls the size of
        memory allocated for the topic. For supported values see the "Queue/topic size"
        section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        """
        return pulumi.get(self, "max_size_in_megabytes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the ServiceBus Topic resource. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the ServiceBus Namespace to create
        this topic in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="requiresDuplicateDetection")
    def requires_duplicate_detection(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether
        the Topic requires duplicate detection. Defaults to false. Changing this forces
        a new resource to be created.
        """
        return pulumi.get(self, "requires_duplicate_detection")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to
        create the namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="supportOrdering")
    def support_ordering(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag which controls whether the Topic
        supports ordering. Defaults to false.
        """
        return pulumi.get(self, "support_ordering")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

