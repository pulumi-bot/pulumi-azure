# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetNamespaceAuthorizationRuleResult:
    """
    A collection of values returned by getNamespaceAuthorizationRule.
    """
    def __init__(__self__, id=None, listen=None, manage=None, name=None, namespace_name=None, primary_connection_string=None, primary_connection_string_alias=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_connection_string_alias=None, secondary_key=None, send=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if listen and not isinstance(listen, bool):
            raise TypeError("Expected argument 'listen' to be a bool")
        __self__.listen = listen
        """
        Does this Authorization Rule have permissions to Listen to the Event Hub?
        """
        if manage and not isinstance(manage, bool):
            raise TypeError("Expected argument 'manage' to be a bool")
        __self__.manage = manage
        """
        Does this Authorization Rule have permissions to Manage to the Event Hub?
        """
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
        The Primary Connection String for the Event Hubs authorization Rule.
        """
        if primary_connection_string_alias and not isinstance(primary_connection_string_alias, str):
            raise TypeError("Expected argument 'primary_connection_string_alias' to be a str")
        __self__.primary_connection_string_alias = primary_connection_string_alias
        """
        The alias of the Primary Connection String for the Event Hubs authorization Rule.
        """
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        __self__.primary_key = primary_key
        """
        The Primary Key for the Event Hubs authorization Rule.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        __self__.secondary_connection_string = secondary_connection_string
        """
        The Secondary Connection String for the Event Hubs authorization Rule.
        """
        if secondary_connection_string_alias and not isinstance(secondary_connection_string_alias, str):
            raise TypeError("Expected argument 'secondary_connection_string_alias' to be a str")
        __self__.secondary_connection_string_alias = secondary_connection_string_alias
        """
        The alias of the Secondary Connection String for the Event Hubs authorization Rule.
        """
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        __self__.secondary_key = secondary_key
        """
        The Secondary Key for the Event Hubs authorization Rule.
        """
        if send and not isinstance(send, bool):
            raise TypeError("Expected argument 'send' to be a bool")
        __self__.send = send
        """
        Does this Authorization Rule have permissions to Send to the Event Hub?
        """


class AwaitableGetNamespaceAuthorizationRuleResult(GetNamespaceAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNamespaceAuthorizationRuleResult(
            id=self.id,
            listen=self.listen,
            manage=self.manage,
            name=self.name,
            namespace_name=self.namespace_name,
            primary_connection_string=self.primary_connection_string,
            primary_connection_string_alias=self.primary_connection_string_alias,
            primary_key=self.primary_key,
            resource_group_name=self.resource_group_name,
            secondary_connection_string=self.secondary_connection_string,
            secondary_connection_string_alias=self.secondary_connection_string_alias,
            secondary_key=self.secondary_key,
            send=self.send)


def get_namespace_authorization_rule(name=None, namespace_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an Authorization Rule for an Event Hub Namespace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.eventhub.get_namespace_authorization_rule(name="navi",
        resource_group_name="example-resources",
        namespace_name="example-ns")
    pulumi.export("eventhubAuthorizationRuleId", data["azurem_eventhub_namespace_authorization_rule"]["example"]["id"])
    ```


    :param str name: The name of the EventHub Authorization Rule resource.
    :param str namespace_name: Specifies the name of the EventHub Namespace.
    :param str resource_group_name: The name of the resource group in which the EventHub Namespace exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:eventhub/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule', __args__, opts=opts).value

    return AwaitableGetNamespaceAuthorizationRuleResult(
        id=__ret__.get('id'),
        listen=__ret__.get('listen'),
        manage=__ret__.get('manage'),
        name=__ret__.get('name'),
        namespace_name=__ret__.get('namespaceName'),
        primary_connection_string=__ret__.get('primaryConnectionString'),
        primary_connection_string_alias=__ret__.get('primaryConnectionStringAlias'),
        primary_key=__ret__.get('primaryKey'),
        resource_group_name=__ret__.get('resourceGroupName'),
        secondary_connection_string=__ret__.get('secondaryConnectionString'),
        secondary_connection_string_alias=__ret__.get('secondaryConnectionStringAlias'),
        secondary_key=__ret__.get('secondaryKey'),
        send=__ret__.get('send'))
