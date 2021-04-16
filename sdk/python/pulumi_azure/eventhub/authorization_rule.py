# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthorizationRuleArgs', 'AuthorizationRule']

@pulumi.input_type
class AuthorizationRuleArgs:
    def __init__(__self__, *,
                 eventhub_name: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 listen: Optional[pulumi.Input[bool]] = None,
                 manage: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 send: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AuthorizationRule resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] send: Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        pulumi.set(__self__, "eventhub_name", eventhub_name)
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if listen is not None:
            pulumi.set(__self__, "listen", listen)
        if manage is not None:
            pulumi.set(__self__, "manage", manage)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if send is not None:
            pulumi.set(__self__, "send", send)

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
        The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def listen(self) -> Optional[pulumi.Input[bool]]:
        """
        Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        """
        return pulumi.get(self, "listen")

    @listen.setter
    def listen(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "listen", value)

    @property
    @pulumi.getter
    def manage(self) -> Optional[pulumi.Input[bool]]:
        """
        Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        """
        return pulumi.get(self, "manage")

    @manage.setter
    def manage(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def send(self) -> Optional[pulumi.Input[bool]]:
        """
        Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        return pulumi.get(self, "send")

    @send.setter
    def send(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send", value)


@pulumi.input_type
class _AuthorizationRuleState:
    def __init__(__self__, *,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 listen: Optional[pulumi.Input[bool]] = None,
                 manage: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 primary_connection_string: Optional[pulumi.Input[str]] = None,
                 primary_connection_string_alias: Optional[pulumi.Input[str]] = None,
                 primary_key: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 secondary_connection_string: Optional[pulumi.Input[str]] = None,
                 secondary_connection_string_alias: Optional[pulumi.Input[str]] = None,
                 secondary_key: Optional[pulumi.Input[str]] = None,
                 send: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AuthorizationRule resources.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The Primary Connection String for the Event Hubs authorization Rule.
        :param pulumi.Input[str] primary_connection_string_alias: The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] primary_key: The Primary Key for the Event Hubs authorization Rule.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string: The Secondary Connection String for the Event Hubs Authorization Rule.
        :param pulumi.Input[str] secondary_connection_string_alias: The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] secondary_key: The Secondary Key for the Event Hubs Authorization Rule.
        :param pulumi.Input[bool] send: Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        if eventhub_name is not None:
            pulumi.set(__self__, "eventhub_name", eventhub_name)
        if listen is not None:
            pulumi.set(__self__, "listen", listen)
        if manage is not None:
            pulumi.set(__self__, "manage", manage)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if primary_connection_string is not None:
            pulumi.set(__self__, "primary_connection_string", primary_connection_string)
        if primary_connection_string_alias is not None:
            pulumi.set(__self__, "primary_connection_string_alias", primary_connection_string_alias)
        if primary_key is not None:
            pulumi.set(__self__, "primary_key", primary_key)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_connection_string is not None:
            pulumi.set(__self__, "secondary_connection_string", secondary_connection_string)
        if secondary_connection_string_alias is not None:
            pulumi.set(__self__, "secondary_connection_string_alias", secondary_connection_string_alias)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)
        if send is not None:
            pulumi.set(__self__, "send", send)

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_name")

    @eventhub_name.setter
    def eventhub_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eventhub_name", value)

    @property
    @pulumi.getter
    def listen(self) -> Optional[pulumi.Input[bool]]:
        """
        Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        """
        return pulumi.get(self, "listen")

    @listen.setter
    def listen(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "listen", value)

    @property
    @pulumi.getter
    def manage(self) -> Optional[pulumi.Input[bool]]:
        """
        Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        """
        return pulumi.get(self, "manage")

    @manage.setter
    def manage(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Primary Connection String for the Event Hubs authorization Rule.
        """
        return pulumi.get(self, "primary_connection_string")

    @primary_connection_string.setter
    def primary_connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_connection_string", value)

    @property
    @pulumi.getter(name="primaryConnectionStringAlias")
    def primary_connection_string_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        """
        return pulumi.get(self, "primary_connection_string_alias")

    @primary_connection_string_alias.setter
    def primary_connection_string_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_connection_string_alias", value)

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Primary Key for the Event Hubs authorization Rule.
        """
        return pulumi.get(self, "primary_key")

    @primary_key.setter
    def primary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_key", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Connection String for the Event Hubs Authorization Rule.
        """
        return pulumi.get(self, "secondary_connection_string")

    @secondary_connection_string.setter
    def secondary_connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_connection_string", value)

    @property
    @pulumi.getter(name="secondaryConnectionStringAlias")
    def secondary_connection_string_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        """
        return pulumi.get(self, "secondary_connection_string_alias")

    @secondary_connection_string_alias.setter
    def secondary_connection_string_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_connection_string_alias", value)

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Key for the Event Hubs Authorization Rule.
        """
        return pulumi.get(self, "secondary_key")

    @secondary_key.setter
    def secondary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_key", value)

    @property
    @pulumi.getter
    def send(self) -> Optional[pulumi.Input[bool]]:
        """
        Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        return pulumi.get(self, "send")

    @send.setter
    def send(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send", value)


class AuthorizationRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 listen: Optional[pulumi.Input[bool]] = None,
                 manage: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 send: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages a Event Hubs authorization Rule within an Event Hub.

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
        example_authorization_rule = azure.eventhub.AuthorizationRule("exampleAuthorizationRule",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name,
            listen=True,
            send=False,
            manage=False)
        ```

        ## Import

        EventHub Authorization Rules can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventhub/authorizationRule:AuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/authorizationRules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] send: Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizationRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Event Hubs authorization Rule within an Event Hub.

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
        example_authorization_rule = azure.eventhub.AuthorizationRule("exampleAuthorizationRule",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name,
            listen=True,
            send=False,
            manage=False)
        ```

        ## Import

        EventHub Authorization Rules can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventhub/authorizationRule:AuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/authorizationRules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param AuthorizationRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizationRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 listen: Optional[pulumi.Input[bool]] = None,
                 manage: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 send: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizationRuleArgs.__new__(AuthorizationRuleArgs)

            if eventhub_name is None and not opts.urn:
                raise TypeError("Missing required property 'eventhub_name'")
            __props__.__dict__["eventhub_name"] = eventhub_name
            __props__.__dict__["listen"] = listen
            __props__.__dict__["manage"] = manage
            __props__.__dict__["name"] = name
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["send"] = send
            __props__.__dict__["primary_connection_string"] = None
            __props__.__dict__["primary_connection_string_alias"] = None
            __props__.__dict__["primary_key"] = None
            __props__.__dict__["secondary_connection_string"] = None
            __props__.__dict__["secondary_connection_string_alias"] = None
            __props__.__dict__["secondary_key"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AuthorizationRule, __self__).__init__(
            'azure:eventhub/authorizationRule:AuthorizationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eventhub_name: Optional[pulumi.Input[str]] = None,
            listen: Optional[pulumi.Input[bool]] = None,
            manage: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            primary_connection_string: Optional[pulumi.Input[str]] = None,
            primary_connection_string_alias: Optional[pulumi.Input[str]] = None,
            primary_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_connection_string: Optional[pulumi.Input[str]] = None,
            secondary_connection_string_alias: Optional[pulumi.Input[str]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            send: Optional[pulumi.Input[bool]] = None) -> 'AuthorizationRule':
        """
        Get an existing AuthorizationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The Primary Connection String for the Event Hubs authorization Rule.
        :param pulumi.Input[str] primary_connection_string_alias: The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] primary_key: The Primary Key for the Event Hubs authorization Rule.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string: The Secondary Connection String for the Event Hubs Authorization Rule.
        :param pulumi.Input[str] secondary_connection_string_alias: The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] secondary_key: The Secondary Key for the Event Hubs Authorization Rule.
        :param pulumi.Input[bool] send: Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthorizationRuleState.__new__(_AuthorizationRuleState)

        __props__.__dict__["eventhub_name"] = eventhub_name
        __props__.__dict__["listen"] = listen
        __props__.__dict__["manage"] = manage
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["primary_connection_string"] = primary_connection_string
        __props__.__dict__["primary_connection_string_alias"] = primary_connection_string_alias
        __props__.__dict__["primary_key"] = primary_key
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["secondary_connection_string"] = secondary_connection_string
        __props__.__dict__["secondary_connection_string_alias"] = secondary_connection_string_alias
        __props__.__dict__["secondary_key"] = secondary_key
        __props__.__dict__["send"] = send
        return AuthorizationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the EventHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_name")

    @property
    @pulumi.getter
    def listen(self) -> pulumi.Output[Optional[bool]]:
        """
        Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        """
        return pulumi.get(self, "listen")

    @property
    @pulumi.getter
    def manage(self) -> pulumi.Output[Optional[bool]]:
        """
        Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        """
        return pulumi.get(self, "manage")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
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
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> pulumi.Output[str]:
        """
        The Primary Connection String for the Event Hubs authorization Rule.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="primaryConnectionStringAlias")
    def primary_connection_string_alias(self) -> pulumi.Output[str]:
        """
        The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        """
        return pulumi.get(self, "primary_connection_string_alias")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> pulumi.Output[str]:
        """
        The Primary Key for the Event Hubs authorization Rule.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> pulumi.Output[str]:
        """
        The Secondary Connection String for the Event Hubs Authorization Rule.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="secondaryConnectionStringAlias")
    def secondary_connection_string_alias(self) -> pulumi.Output[str]:
        """
        The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        """
        return pulumi.get(self, "secondary_connection_string_alias")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[str]:
        """
        The Secondary Key for the Event Hubs Authorization Rule.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter
    def send(self) -> pulumi.Output[Optional[bool]]:
        """
        Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        return pulumi.get(self, "send")

