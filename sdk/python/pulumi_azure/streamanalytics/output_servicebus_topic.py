# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['OutputServicebusTopic']


class OutputServicebusTopic(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['OutputServicebusTopicSerializationArgs']]] = None,
                 servicebus_namespace: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_key: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Stream Analytics Output to a ServiceBus Topic.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard")
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_output_servicebus_topic = azure.streamanalytics.OutputServicebusTopic("exampleOutputServicebusTopic",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            topic_name=example_topic.name,
            servicebus_namespace=example_namespace.name,
            shared_access_policy_key=example_namespace.default_primary_key,
            shared_access_policy_name="RootManageSharedAccessKey",
            serialization=azure.streamanalytics.OutputServicebusTopicSerializationArgs(
                format="Avro",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['OutputServicebusTopicSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the Service Bus Topic.
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

            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if serialization is None:
                raise TypeError("Missing required property 'serialization'")
            __props__['serialization'] = serialization
            if servicebus_namespace is None:
                raise TypeError("Missing required property 'servicebus_namespace'")
            __props__['servicebus_namespace'] = servicebus_namespace
            if shared_access_policy_key is None:
                raise TypeError("Missing required property 'shared_access_policy_key'")
            __props__['shared_access_policy_key'] = shared_access_policy_key
            if shared_access_policy_name is None:
                raise TypeError("Missing required property 'shared_access_policy_name'")
            __props__['shared_access_policy_name'] = shared_access_policy_name
            if stream_analytics_job_name is None:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
            if topic_name is None:
                raise TypeError("Missing required property 'topic_name'")
            __props__['topic_name'] = topic_name
        super(OutputServicebusTopic, __self__).__init__(
            'azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            serialization: Optional[pulumi.Input[pulumi.InputType['OutputServicebusTopicSerializationArgs']]] = None,
            servicebus_namespace: Optional[pulumi.Input[str]] = None,
            shared_access_policy_key: Optional[pulumi.Input[str]] = None,
            shared_access_policy_name: Optional[pulumi.Input[str]] = None,
            stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'OutputServicebusTopic':
        """
        Get an existing OutputServicebusTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['OutputServicebusTopicSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the Service Bus Topic.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["serialization"] = serialization
        __props__["servicebus_namespace"] = servicebus_namespace
        __props__["shared_access_policy_key"] = shared_access_policy_key
        __props__["shared_access_policy_name"] = shared_access_policy_name
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        __props__["topic_name"] = topic_name
        return OutputServicebusTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Output. Changing this forces a new resource to be created.
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
    def serialization(self) -> pulumi.Output['outputs.OutputServicebusTopicSerialization']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @property
    @pulumi.getter(name="servicebusNamespace")
    def servicebus_namespace(self) -> pulumi.Output[str]:
        """
        The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
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

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        The name of the Service Bus Topic.
        """
        return pulumi.get(self, "topic_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

