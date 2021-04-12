# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ConsumerGroupArgs', 'ConsumerGroup']

@pulumi.input_type
class ConsumerGroupArgs:
    def __init__(__self__, *,
                 eventhub_name: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConsumerGroup resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: Specifies the user metadata.
        """
        pulumi.set(__self__, "eventhub_name", eventhub_name)
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if user_metadata is not None:
            pulumi.set(__self__, "user_metadata", user_metadata)

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the EventHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_name")

    @eventhub_name.setter
    def eventhub_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "eventhub_name", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user metadata.
        """
        return pulumi.get(self, "user_metadata")

    @user_metadata.setter
    def user_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_metadata", value)


class ConsumerGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location="West US",
            resource_group_name=example_resource_group.name,
            sku="Basic",
            capacity=2,
            tags={
                "environment": "Production",
            })
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            namespace_name=example_event_hub_namespace.name,
            resource_group_name=example_resource_group.name,
            partition_count=2,
            message_retention=2)
        example_consumer_group = azure.eventhub.ConsumerGroup("exampleConsumerGroup",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name,
            user_metadata="some-meta-data")
        ```

        ## Import

        EventHub Consumer Groups can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventhub/consumerGroup:ConsumerGroup consumerGroup1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/consumergroups/consumerGroup1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: Specifies the user metadata.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConsumerGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location="West US",
            resource_group_name=example_resource_group.name,
            sku="Basic",
            capacity=2,
            tags={
                "environment": "Production",
            })
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            namespace_name=example_event_hub_namespace.name,
            resource_group_name=example_resource_group.name,
            partition_count=2,
            message_retention=2)
        example_consumer_group = azure.eventhub.ConsumerGroup("exampleConsumerGroup",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name,
            user_metadata="some-meta-data")
        ```

        ## Import

        EventHub Consumer Groups can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventhub/consumerGroup:ConsumerGroup consumerGroup1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/consumergroups/consumerGroup1
        ```

        :param str resource_name: The name of the resource.
        :param ConsumerGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConsumerGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None,
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

            if eventhub_name is None and not opts.urn:
                raise TypeError("Missing required property 'eventhub_name'")
            __props__['eventhub_name'] = eventhub_name
            __props__['name'] = name
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['user_metadata'] = user_metadata
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ConsumerGroup, __self__).__init__(
            'azure:eventhub/consumerGroup:ConsumerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eventhub_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            user_metadata: Optional[pulumi.Input[str]] = None) -> 'ConsumerGroup':
        """
        Get an existing ConsumerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: Specifies the user metadata.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["eventhub_name"] = eventhub_name
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["resource_group_name"] = resource_group_name
        __props__["user_metadata"] = user_metadata
        return ConsumerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the EventHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the user metadata.
        """
        return pulumi.get(self, "user_metadata")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

