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

__all__ = ['EventHub']


class EventHub(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capture_description: Optional[pulumi.Input[pulumi.InputType['EventHubCaptureDescriptionArgs']]] = None,
                 message_retention: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Event Hubs as a nested resource within a Event Hubs namespace.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            capacity=1,
            tags={
                "environment": "Production",
            })
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            namespace_name=example_event_hub_namespace.name,
            resource_group_name=example_resource_group.name,
            partition_count=2,
            message_retention=1)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventHubCaptureDescriptionArgs']] capture_description: A `capture_description` block as defined below.
        :param pulumi.Input[float] message_retention: Specifies the number of days to retain the events for this Event Hub.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[float] partition_count: Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
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

            __props__['capture_description'] = capture_description
            if message_retention is None:
                raise TypeError("Missing required property 'message_retention'")
            __props__['message_retention'] = message_retention
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if partition_count is None:
                raise TypeError("Missing required property 'partition_count'")
            __props__['partition_count'] = partition_count
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['partition_ids'] = None
        super(EventHub, __self__).__init__(
            'azure:eventhub/eventHub:EventHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            capture_description: Optional[pulumi.Input[pulumi.InputType['EventHubCaptureDescriptionArgs']]] = None,
            message_retention: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            partition_count: Optional[pulumi.Input[float]] = None,
            partition_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'EventHub':
        """
        Get an existing EventHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventHubCaptureDescriptionArgs']] capture_description: A `capture_description` block as defined below.
        :param pulumi.Input[float] message_retention: Specifies the number of days to retain the events for this Event Hub.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[float] partition_count: Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[str]]] partition_ids: The identifiers for partitions created for Event Hubs.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["capture_description"] = capture_description
        __props__["message_retention"] = message_retention
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["partition_count"] = partition_count
        __props__["partition_ids"] = partition_ids
        __props__["resource_group_name"] = resource_group_name
        return EventHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="captureDescription")
    def capture_description(self) -> Optional['outputs.EventHubCaptureDescription']:
        """
        A `capture_description` block as defined below.
        """
        return pulumi.get(self, "capture_description")

    @property
    @pulumi.getter(name="messageRetention")
    def message_retention(self) -> float:
        """
        Specifies the number of days to retain the events for this Event Hub.
        """
        return pulumi.get(self, "message_retention")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> str:
        """
        Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> float:
        """
        Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="partitionIds")
    def partition_ids(self) -> List[str]:
        """
        The identifiers for partitions created for Event Hubs.
        """
        return pulumi.get(self, "partition_ids")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

