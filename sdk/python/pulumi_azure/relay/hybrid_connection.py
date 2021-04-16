# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HybridConnectionArgs', 'HybridConnection']

@pulumi.input_type
class HybridConnectionArgs:
    def __init__(__self__, *,
                 relay_namespace_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 requires_client_authorization: Optional[pulumi.Input[bool]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HybridConnection resource.
        :param pulumi.Input[str] relay_namespace_name: The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_client_authorization: Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        pulumi.set(__self__, "relay_namespace_name", relay_namespace_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if requires_client_authorization is not None:
            pulumi.set(__self__, "requires_client_authorization", requires_client_authorization)
        if user_metadata is not None:
            pulumi.set(__self__, "user_metadata", user_metadata)

    @property
    @pulumi.getter(name="relayNamespaceName")
    def relay_namespace_name(self) -> pulumi.Input[str]:
        """
        The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "relay_namespace_name")

    @relay_namespace_name.setter
    def relay_namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "relay_namespace_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requiresClientAuthorization")
    def requires_client_authorization(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "requires_client_authorization")

    @requires_client_authorization.setter
    def requires_client_authorization(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requires_client_authorization", value)

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        return pulumi.get(self, "user_metadata")

    @user_metadata.setter
    def user_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_metadata", value)


@pulumi.input_type
class _HybridConnectionState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 relay_namespace_name: Optional[pulumi.Input[str]] = None,
                 requires_client_authorization: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HybridConnection resources.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] relay_namespace_name: The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_client_authorization: Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if relay_namespace_name is not None:
            pulumi.set(__self__, "relay_namespace_name", relay_namespace_name)
        if requires_client_authorization is not None:
            pulumi.set(__self__, "requires_client_authorization", requires_client_authorization)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if user_metadata is not None:
            pulumi.set(__self__, "user_metadata", user_metadata)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="relayNamespaceName")
    def relay_namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "relay_namespace_name")

    @relay_namespace_name.setter
    def relay_namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "relay_namespace_name", value)

    @property
    @pulumi.getter(name="requiresClientAuthorization")
    def requires_client_authorization(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "requires_client_authorization")

    @requires_client_authorization.setter
    def requires_client_authorization(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requires_client_authorization", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        return pulumi.get(self, "user_metadata")

    @user_metadata.setter
    def user_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_metadata", value)


class HybridConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 relay_namespace_name: Optional[pulumi.Input[str]] = None,
                 requires_client_authorization: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Azure Relay Hybrid Connection.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.relay.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Standard",
            tags={
                "source": "managed",
            })
        example_hybrid_connection = azure.relay.HybridConnection("exampleHybridConnection",
            resource_group_name=example_resource_group.name,
            relay_namespace_name=example_namespace.name,
            requires_client_authorization=False,
            user_metadata="testmetadata")
        ```

        ## Import

        Relay Hybrid Connection's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:relay/hybridConnection:HybridConnection relay1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1/hybridConnections/hconn1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] relay_namespace_name: The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_client_authorization: Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HybridConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure Relay Hybrid Connection.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.relay.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Standard",
            tags={
                "source": "managed",
            })
        example_hybrid_connection = azure.relay.HybridConnection("exampleHybridConnection",
            resource_group_name=example_resource_group.name,
            relay_namespace_name=example_namespace.name,
            requires_client_authorization=False,
            user_metadata="testmetadata")
        ```

        ## Import

        Relay Hybrid Connection's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:relay/hybridConnection:HybridConnection relay1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1/hybridConnections/hconn1
        ```

        :param str resource_name: The name of the resource.
        :param HybridConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HybridConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 relay_namespace_name: Optional[pulumi.Input[str]] = None,
                 requires_client_authorization: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_metadata: Optional[pulumi.Input[str]] = None,
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
            __props__ = HybridConnectionArgs.__new__(HybridConnectionArgs)

            __props__.__dict__["name"] = name
            if relay_namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'relay_namespace_name'")
            __props__.__dict__["relay_namespace_name"] = relay_namespace_name
            __props__.__dict__["requires_client_authorization"] = requires_client_authorization
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["user_metadata"] = user_metadata
        super(HybridConnection, __self__).__init__(
            'azure:relay/hybridConnection:HybridConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            relay_namespace_name: Optional[pulumi.Input[str]] = None,
            requires_client_authorization: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            user_metadata: Optional[pulumi.Input[str]] = None) -> 'HybridConnection':
        """
        Get an existing HybridConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] relay_namespace_name: The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_client_authorization: Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_metadata: The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HybridConnectionState.__new__(_HybridConnectionState)

        __props__.__dict__["name"] = name
        __props__.__dict__["relay_namespace_name"] = relay_namespace_name
        __props__.__dict__["requires_client_authorization"] = requires_client_authorization
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["user_metadata"] = user_metadata
        return HybridConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="relayNamespaceName")
    def relay_namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "relay_namespace_name")

    @property
    @pulumi.getter(name="requiresClientAuthorization")
    def requires_client_authorization(self) -> pulumi.Output[Optional[bool]]:
        """
        Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "requires_client_authorization")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> pulumi.Output[Optional[str]]:
        """
        The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        """
        return pulumi.get(self, "user_metadata")

