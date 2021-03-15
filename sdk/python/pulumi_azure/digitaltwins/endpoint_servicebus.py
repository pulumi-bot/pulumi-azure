# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['EndpointServicebusArgs', 'EndpointServicebus']

@pulumi.input_type
class EndpointServicebusArgs:
    def __init__(__self__, *,
                 digital_twins_id: pulumi.Input[str],
                 servicebus_primary_connection_string: pulumi.Input[str],
                 servicebus_secondary_connection_string: pulumi.Input[str],
                 dead_letter_storage_secret: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndpointServicebus resource.
        :param pulumi.Input[str] digital_twins_id: The ID of the Digital Twins Instance. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        :param pulumi.Input[str] servicebus_primary_connection_string: The primary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission. .
        :param pulumi.Input[str] servicebus_secondary_connection_string: The secondary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission.
        :param pulumi.Input[str] dead_letter_storage_secret: The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
        :param pulumi.Input[str] name: The name which should be used for this Digital Twins Service Bus Endpoint. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        """
        pulumi.set(__self__, "digital_twins_id", digital_twins_id)
        pulumi.set(__self__, "servicebus_primary_connection_string", servicebus_primary_connection_string)
        pulumi.set(__self__, "servicebus_secondary_connection_string", servicebus_secondary_connection_string)
        if dead_letter_storage_secret is not None:
            pulumi.set(__self__, "dead_letter_storage_secret", dead_letter_storage_secret)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="digitalTwinsId")
    def digital_twins_id(self) -> pulumi.Input[str]:
        """
        The ID of the Digital Twins Instance. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        """
        return pulumi.get(self, "digital_twins_id")

    @digital_twins_id.setter
    def digital_twins_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "digital_twins_id", value)

    @property
    @pulumi.getter(name="servicebusPrimaryConnectionString")
    def servicebus_primary_connection_string(self) -> pulumi.Input[str]:
        """
        The primary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission. .
        """
        return pulumi.get(self, "servicebus_primary_connection_string")

    @servicebus_primary_connection_string.setter
    def servicebus_primary_connection_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "servicebus_primary_connection_string", value)

    @property
    @pulumi.getter(name="servicebusSecondaryConnectionString")
    def servicebus_secondary_connection_string(self) -> pulumi.Input[str]:
        """
        The secondary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission.
        """
        return pulumi.get(self, "servicebus_secondary_connection_string")

    @servicebus_secondary_connection_string.setter
    def servicebus_secondary_connection_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "servicebus_secondary_connection_string", value)

    @property
    @pulumi.getter(name="deadLetterStorageSecret")
    def dead_letter_storage_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
        """
        return pulumi.get(self, "dead_letter_storage_secret")

    @dead_letter_storage_secret.setter
    def dead_letter_storage_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dead_letter_storage_secret", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Digital Twins Service Bus Endpoint. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class EndpointServicebus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointServicebusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Digital Twins Service Bus Endpoint.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_instance = azure.digitaltwins.Instance("exampleInstance",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard")
        example_topic = azure.servicebus.Topic("exampleTopic",
            namespace_name=example_namespace.name,
            resource_group_name=example_resource_group.name)
        example_topic_authorization_rule = azure.servicebus.TopicAuthorizationRule("exampleTopicAuthorizationRule",
            namespace_name=example_namespace.name,
            resource_group_name=example_resource_group.name,
            topic_name=example_topic.name,
            listen=False,
            send=True,
            manage=False)
        example_endpoint_servicebus = azure.digitaltwins.EndpointServicebus("exampleEndpointServicebus",
            digital_twins_id=example_instance.id,
            servicebus_primary_connection_string=example_topic_authorization_rule.primary_connection_string,
            servicebus_secondary_connection_string=example_topic_authorization_rule.secondary_connection_string)
        ```

        ## Import

        Digital Twins Service Bus Endpoints can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:digitaltwins/endpointServicebus:EndpointServicebus example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
        ```

        :param str resource_name: The name of the resource.
        :param EndpointServicebusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dead_letter_storage_secret: Optional[pulumi.Input[str]] = None,
                 digital_twins_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servicebus_primary_connection_string: Optional[pulumi.Input[str]] = None,
                 servicebus_secondary_connection_string: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Digital Twins Service Bus Endpoint.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_instance = azure.digitaltwins.Instance("exampleInstance",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard")
        example_topic = azure.servicebus.Topic("exampleTopic",
            namespace_name=example_namespace.name,
            resource_group_name=example_resource_group.name)
        example_topic_authorization_rule = azure.servicebus.TopicAuthorizationRule("exampleTopicAuthorizationRule",
            namespace_name=example_namespace.name,
            resource_group_name=example_resource_group.name,
            topic_name=example_topic.name,
            listen=False,
            send=True,
            manage=False)
        example_endpoint_servicebus = azure.digitaltwins.EndpointServicebus("exampleEndpointServicebus",
            digital_twins_id=example_instance.id,
            servicebus_primary_connection_string=example_topic_authorization_rule.primary_connection_string,
            servicebus_secondary_connection_string=example_topic_authorization_rule.secondary_connection_string)
        ```

        ## Import

        Digital Twins Service Bus Endpoints can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:digitaltwins/endpointServicebus:EndpointServicebus example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dead_letter_storage_secret: The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
        :param pulumi.Input[str] digital_twins_id: The ID of the Digital Twins Instance. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        :param pulumi.Input[str] name: The name which should be used for this Digital Twins Service Bus Endpoint. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        :param pulumi.Input[str] servicebus_primary_connection_string: The primary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission. .
        :param pulumi.Input[str] servicebus_secondary_connection_string: The secondary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointServicebusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dead_letter_storage_secret: Optional[pulumi.Input[str]] = None,
                 digital_twins_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servicebus_primary_connection_string: Optional[pulumi.Input[str]] = None,
                 servicebus_secondary_connection_string: Optional[pulumi.Input[str]] = None,
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

            __props__['dead_letter_storage_secret'] = dead_letter_storage_secret
            if digital_twins_id is None and not opts.urn:
                raise TypeError("Missing required property 'digital_twins_id'")
            __props__['digital_twins_id'] = digital_twins_id
            __props__['name'] = name
            if servicebus_primary_connection_string is None and not opts.urn:
                raise TypeError("Missing required property 'servicebus_primary_connection_string'")
            __props__['servicebus_primary_connection_string'] = servicebus_primary_connection_string
            if servicebus_secondary_connection_string is None and not opts.urn:
                raise TypeError("Missing required property 'servicebus_secondary_connection_string'")
            __props__['servicebus_secondary_connection_string'] = servicebus_secondary_connection_string
        super(EndpointServicebus, __self__).__init__(
            'azure:digitaltwins/endpointServicebus:EndpointServicebus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dead_letter_storage_secret: Optional[pulumi.Input[str]] = None,
            digital_twins_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            servicebus_primary_connection_string: Optional[pulumi.Input[str]] = None,
            servicebus_secondary_connection_string: Optional[pulumi.Input[str]] = None) -> 'EndpointServicebus':
        """
        Get an existing EndpointServicebus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dead_letter_storage_secret: The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
        :param pulumi.Input[str] digital_twins_id: The ID of the Digital Twins Instance. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        :param pulumi.Input[str] name: The name which should be used for this Digital Twins Service Bus Endpoint. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        :param pulumi.Input[str] servicebus_primary_connection_string: The primary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission. .
        :param pulumi.Input[str] servicebus_secondary_connection_string: The secondary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dead_letter_storage_secret"] = dead_letter_storage_secret
        __props__["digital_twins_id"] = digital_twins_id
        __props__["name"] = name
        __props__["servicebus_primary_connection_string"] = servicebus_primary_connection_string
        __props__["servicebus_secondary_connection_string"] = servicebus_secondary_connection_string
        return EndpointServicebus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deadLetterStorageSecret")
    def dead_letter_storage_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
        """
        return pulumi.get(self, "dead_letter_storage_secret")

    @property
    @pulumi.getter(name="digitalTwinsId")
    def digital_twins_id(self) -> pulumi.Output[str]:
        """
        The ID of the Digital Twins Instance. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        """
        return pulumi.get(self, "digital_twins_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Digital Twins Service Bus Endpoint. Changing this forces a new Digital Twins Service Bus Endpoint to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="servicebusPrimaryConnectionString")
    def servicebus_primary_connection_string(self) -> pulumi.Output[str]:
        """
        The primary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission. .
        """
        return pulumi.get(self, "servicebus_primary_connection_string")

    @property
    @pulumi.getter(name="servicebusSecondaryConnectionString")
    def servicebus_secondary_connection_string(self) -> pulumi.Output[str]:
        """
        The secondary connection string of the Service Bus Topic Authorization Rule with a minimum of `send` permission.
        """
        return pulumi.get(self, "servicebus_secondary_connection_string")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

