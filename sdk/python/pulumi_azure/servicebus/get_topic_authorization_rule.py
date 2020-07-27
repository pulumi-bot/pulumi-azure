# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetTopicAuthorizationRuleResult',
    'AwaitableGetTopicAuthorizationRuleResult',
    'get_topic_authorization_rule',
]


class GetTopicAuthorizationRuleResult:
    """
    A collection of values returned by getTopicAuthorizationRule.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, id=None, listen=None, manage=None, name=None, namespace_name=None, primary_connection_string=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_key=None, send=None, topic_name=None) -> None:
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if listen and not isinstance(listen, bool):
            raise TypeError("Expected argument 'listen' to be a bool")
        __self__.listen = listen
        if manage and not isinstance(manage, bool):
            raise TypeError("Expected argument 'manage' to be a bool")
        __self__.manage = manage
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        __self__.namespace_name = namespace_name
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        __self__.primary_connection_string = primary_connection_string
        """
        The Primary Connection String for the ServiceBus Topic authorization Rule.
        """
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        __self__.primary_key = primary_key
        """
        The Primary Key for the ServiceBus Topic authorization Rule.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        __self__.secondary_connection_string = secondary_connection_string
        """
        The Secondary Connection String for the ServiceBus Topic authorization Rule.
        """
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        __self__.secondary_key = secondary_key
        """
        The Secondary Key for the ServiceBus Topic authorization Rule.
        """
        if send and not isinstance(send, bool):
            raise TypeError("Expected argument 'send' to be a bool")
        __self__.send = send
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        __self__.topic_name = topic_name


class AwaitableGetTopicAuthorizationRuleResult(GetTopicAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicAuthorizationRuleResult(
            id=self.id,
            listen=self.listen,
            manage=self.manage,
            name=self.name,
            namespace_name=self.namespace_name,
            primary_connection_string=self.primary_connection_string,
            primary_key=self.primary_key,
            resource_group_name=self.resource_group_name,
            secondary_connection_string=self.secondary_connection_string,
            secondary_key=self.secondary_key,
            send=self.send,
            topic_name=self.topic_name)


def get_topic_authorization_rule(name: Optional[str] = None, namespace_name: Optional[str] = None, resource_group_name: Optional[str] = None, topic_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicAuthorizationRuleResult:
    """
    Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.servicebus.get_topic_authorization_rule(name="example-tfex_name",
        namespace_name="example-namespace",
        resource_group_name="example-resources",
        topic_name="example-servicebus_topic")
    pulumi.export("servicebusAuthorizationRuleId", data["azurem_servicebus_topic_authorization_rule"]["example"]["id"])
    ```


    :param str name: The name of the ServiceBus Topic Authorization Rule resource.
    :param str namespace_name: The name of the ServiceBus Namespace.
    :param str resource_group_name: The name of the resource group in which the ServiceBus Namespace exists.
    :param str topic_name: The name of the ServiceBus Topic.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['topicName'] = topic_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:servicebus/getTopicAuthorizationRule:getTopicAuthorizationRule', __args__, opts=opts).value

    return AwaitableGetTopicAuthorizationRuleResult(
        id=__ret__.get('id'),
        listen=__ret__.get('listen'),
        manage=__ret__.get('manage'),
        name=__ret__.get('name'),
        namespace_name=__ret__.get('namespaceName'),
        primary_connection_string=__ret__.get('primaryConnectionString'),
        primary_key=__ret__.get('primaryKey'),
        resource_group_name=__ret__.get('resourceGroupName'),
        secondary_connection_string=__ret__.get('secondaryConnectionString'),
        secondary_key=__ret__.get('secondaryKey'),
        send=__ret__.get('send'),
        topic_name=__ret__.get('topicName'))
