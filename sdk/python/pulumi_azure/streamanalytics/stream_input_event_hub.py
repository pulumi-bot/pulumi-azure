# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['StreamInputEventHubArgs', 'StreamInputEventHub']

@pulumi.input_type
class StreamInputEventHubArgs:
    def __init__(__self__, *,
                 eventhub_consumer_group_name: pulumi.Input[str],
                 eventhub_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 serialization: pulumi.Input['StreamInputEventHubSerializationArgs'],
                 servicebus_namespace: pulumi.Input[str],
                 shared_access_policy_key: pulumi.Input[str],
                 shared_access_policy_name: pulumi.Input[str],
                 stream_analytics_job_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StreamInputEventHub resource.
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] eventhub_name: The name of the Event Hub.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input['StreamInputEventHubSerializationArgs'] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Stream Input EventHub. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "eventhub_consumer_group_name", eventhub_consumer_group_name)
        pulumi.set(__self__, "eventhub_name", eventhub_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "serialization", serialization)
        pulumi.set(__self__, "servicebus_namespace", servicebus_namespace)
        pulumi.set(__self__, "shared_access_policy_key", shared_access_policy_key)
        pulumi.set(__self__, "shared_access_policy_name", shared_access_policy_name)
        pulumi.set(__self__, "stream_analytics_job_name", stream_analytics_job_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="eventhubConsumerGroupName")
    def eventhub_consumer_group_name(self) -> pulumi.Input[str]:
        """
        The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        """
        return pulumi.get(self, "eventhub_consumer_group_name")

    @eventhub_consumer_group_name.setter
    def eventhub_consumer_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "eventhub_consumer_group_name", value)

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> pulumi.Input[str]:
        """
        The name of the Event Hub.
        """
        return pulumi.get(self, "eventhub_name")

    @eventhub_name.setter
    def eventhub_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "eventhub_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def serialization(self) -> pulumi.Input['StreamInputEventHubSerializationArgs']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @serialization.setter
    def serialization(self, value: pulumi.Input['StreamInputEventHubSerializationArgs']):
        pulumi.set(self, "serialization", value)

    @property
    @pulumi.getter(name="servicebusNamespace")
    def servicebus_namespace(self) -> pulumi.Input[str]:
        """
        The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "servicebus_namespace")

    @servicebus_namespace.setter
    def servicebus_namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "servicebus_namespace", value)

    @property
    @pulumi.getter(name="sharedAccessPolicyKey")
    def shared_access_policy_key(self) -> pulumi.Input[str]:
        """
        The shared access policy key for the specified shared access policy.
        """
        return pulumi.get(self, "shared_access_policy_key")

    @shared_access_policy_key.setter
    def shared_access_policy_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "shared_access_policy_key", value)

    @property
    @pulumi.getter(name="sharedAccessPolicyName")
    def shared_access_policy_name(self) -> pulumi.Input[str]:
        """
        The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "shared_access_policy_name")

    @shared_access_policy_name.setter
    def shared_access_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "shared_access_policy_name", value)

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Input[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @stream_analytics_job_name.setter
    def stream_analytics_job_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_analytics_job_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Input EventHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class StreamInputEventHub(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamInputEventHubArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Stream Analytics Stream Input EventHub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            capacity=1)
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            namespace_name=example_event_hub_namespace.name,
            resource_group_name=example_resource_group.name,
            partition_count=2,
            message_retention=1)
        example_consumer_group = azure.eventhub.ConsumerGroup("exampleConsumerGroup",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name)
        example_stream_input_event_hub = azure.streamanalytics.StreamInputEventHub("exampleStreamInputEventHub",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            eventhub_consumer_group_name=example_consumer_group.name,
            eventhub_name=example_event_hub.name,
            servicebus_namespace=example_event_hub_namespace.name,
            shared_access_policy_key=example_event_hub_namespace.default_primary_key,
            shared_access_policy_name="RootManageSharedAccessKey",
            serialization=azure.streamanalytics.StreamInputEventHubSerializationArgs(
                type="Json",
                encoding="UTF8",
            ))
        ```

        ## Import

        Stream Analytics Stream Input EventHub's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/streamInputEventHub:StreamInputEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
        ```

        :param str resource_name: The name of the resource.
        :param StreamInputEventHubArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputEventHubSerializationArgs']]] = None,
                 servicebus_namespace: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_key: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Stream Analytics Stream Input EventHub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            capacity=1)
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            namespace_name=example_event_hub_namespace.name,
            resource_group_name=example_resource_group.name,
            partition_count=2,
            message_retention=1)
        example_consumer_group = azure.eventhub.ConsumerGroup("exampleConsumerGroup",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name)
        example_stream_input_event_hub = azure.streamanalytics.StreamInputEventHub("exampleStreamInputEventHub",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            eventhub_consumer_group_name=example_consumer_group.name,
            eventhub_name=example_event_hub.name,
            servicebus_namespace=example_event_hub_namespace.name,
            shared_access_policy_key=example_event_hub_namespace.default_primary_key,
            shared_access_policy_name="RootManageSharedAccessKey",
            serialization=azure.streamanalytics.StreamInputEventHubSerializationArgs(
                type="Json",
                encoding="UTF8",
            ))
        ```

        ## Import

        Stream Analytics Stream Input EventHub's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/streamInputEventHub:StreamInputEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] eventhub_name: The name of the Event Hub.
        :param pulumi.Input[str] name: The name of the Stream Input EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['StreamInputEventHubSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamInputEventHubArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputEventHubSerializationArgs']]] = None,
                 servicebus_namespace: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_key: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
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

            if eventhub_consumer_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'eventhub_consumer_group_name'")
            __props__['eventhub_consumer_group_name'] = eventhub_consumer_group_name
            if eventhub_name is None and not opts.urn:
                raise TypeError("Missing required property 'eventhub_name'")
            __props__['eventhub_name'] = eventhub_name
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if serialization is None and not opts.urn:
                raise TypeError("Missing required property 'serialization'")
            __props__['serialization'] = serialization
            if servicebus_namespace is None and not opts.urn:
                raise TypeError("Missing required property 'servicebus_namespace'")
            __props__['servicebus_namespace'] = servicebus_namespace
            if shared_access_policy_key is None and not opts.urn:
                raise TypeError("Missing required property 'shared_access_policy_key'")
            __props__['shared_access_policy_key'] = shared_access_policy_key
            if shared_access_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'shared_access_policy_name'")
            __props__['shared_access_policy_name'] = shared_access_policy_name
            if stream_analytics_job_name is None and not opts.urn:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
        super(StreamInputEventHub, __self__).__init__(
            'azure:streamanalytics/streamInputEventHub:StreamInputEventHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
            eventhub_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputEventHubSerializationArgs']]] = None,
            servicebus_namespace: Optional[pulumi.Input[str]] = None,
            shared_access_policy_key: Optional[pulumi.Input[str]] = None,
            shared_access_policy_name: Optional[pulumi.Input[str]] = None,
            stream_analytics_job_name: Optional[pulumi.Input[str]] = None) -> 'StreamInputEventHub':
        """
        Get an existing StreamInputEventHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] eventhub_name: The name of the Event Hub.
        :param pulumi.Input[str] name: The name of the Stream Input EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['StreamInputEventHubSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["eventhub_consumer_group_name"] = eventhub_consumer_group_name
        __props__["eventhub_name"] = eventhub_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["serialization"] = serialization
        __props__["servicebus_namespace"] = servicebus_namespace
        __props__["shared_access_policy_key"] = shared_access_policy_key
        __props__["shared_access_policy_name"] = shared_access_policy_name
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        return StreamInputEventHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventhubConsumerGroupName")
    def eventhub_consumer_group_name(self) -> pulumi.Output[str]:
        """
        The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        """
        return pulumi.get(self, "eventhub_consumer_group_name")

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> pulumi.Output[str]:
        """
        The name of the Event Hub.
        """
        return pulumi.get(self, "eventhub_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Input EventHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def serialization(self) -> pulumi.Output['outputs.StreamInputEventHubSerialization']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @property
    @pulumi.getter(name="servicebusNamespace")
    def servicebus_namespace(self) -> pulumi.Output[str]:
        """
        The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "servicebus_namespace")

    @property
    @pulumi.getter(name="sharedAccessPolicyKey")
    def shared_access_policy_key(self) -> pulumi.Output[str]:
        """
        The shared access policy key for the specified shared access policy.
        """
        return pulumi.get(self, "shared_access_policy_key")

    @property
    @pulumi.getter(name="sharedAccessPolicyName")
    def shared_access_policy_name(self) -> pulumi.Output[str]:
        """
        The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "shared_access_policy_name")

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

