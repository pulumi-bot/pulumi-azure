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

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 storage_accounts: pulumi.Input[Sequence[pulumi.Input['AccountStorageAccountArgs']]],
                 identity: Optional[pulumi.Input['AccountIdentityArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 storage_authentication_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['AccountStorageAccountArgs']]] storage_accounts: One or more `storage_account` blocks as defined below.
        :param pulumi.Input['AccountIdentityArgs'] identity: An `identity` block is documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_authentication_type: Specifies the storage authentication type. 
               Possible value is  `ManagedIdentity` or `System`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags assigned to the resource.
               ---
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "storage_accounts", storage_accounts)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if storage_authentication_type is not None:
            pulumi.set(__self__, "storage_authentication_type", storage_authentication_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="storageAccounts")
    def storage_accounts(self) -> pulumi.Input[Sequence[pulumi.Input['AccountStorageAccountArgs']]]:
        """
        One or more `storage_account` blocks as defined below.
        """
        return pulumi.get(self, "storage_accounts")

    @storage_accounts.setter
    def storage_accounts(self, value: pulumi.Input[Sequence[pulumi.Input['AccountStorageAccountArgs']]]):
        pulumi.set(self, "storage_accounts", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['AccountIdentityArgs']]:
        """
        An `identity` block is documented below.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['AccountIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="storageAuthenticationType")
    def storage_authentication_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the storage authentication type. 
        Possible value is  `ManagedIdentity` or `System`.
        """
        return pulumi.get(self, "storage_authentication_type")

    @storage_authentication_type.setter
    def storage_authentication_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_authentication_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags assigned to the resource.
        ---
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""azure.mediaservices.Account has been deprecated in favor of azure.media.ServiceAccount""", DeprecationWarning)


class Account(pulumi.CustomResource):
    warnings.warn("""azure.mediaservices.Account has been deprecated in favor of azure.media.ServiceAccount""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['AccountIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountStorageAccountArgs']]]]] = None,
                 storage_authentication_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Media Services Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_service_account = azure.media.ServiceAccount("exampleServiceAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_accounts=[azure.media.ServiceAccountStorageAccountArgs(
                id=example_account.id,
                is_primary=True,
            )])
        ```

        ## Import

        Media Services Accounts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mediaservices/account:Account account /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/mediaservices/account1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AccountIdentityArgs']] identity: An `identity` block is documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountStorageAccountArgs']]]] storage_accounts: One or more `storage_account` blocks as defined below.
        :param pulumi.Input[str] storage_authentication_type: Specifies the storage authentication type. 
               Possible value is  `ManagedIdentity` or `System`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags assigned to the resource.
               ---
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Media Services Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_service_account = azure.media.ServiceAccount("exampleServiceAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_accounts=[azure.media.ServiceAccountStorageAccountArgs(
                id=example_account.id,
                is_primary=True,
            )])
        ```

        ## Import

        Media Services Accounts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mediaservices/account:Account account /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/mediaservices/account1
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['AccountIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountStorageAccountArgs']]]]] = None,
                 storage_authentication_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        pulumi.log.warn("""Account is deprecated: azure.mediaservices.Account has been deprecated in favor of azure.media.ServiceAccount""")
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

            __props__['identity'] = identity
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_accounts is None and not opts.urn:
                raise TypeError("Missing required property 'storage_accounts'")
            __props__['storage_accounts'] = storage_accounts
            __props__['storage_authentication_type'] = storage_authentication_type
            __props__['tags'] = tags
        super(Account, __self__).__init__(
            'azure:mediaservices/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['AccountIdentityArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountStorageAccountArgs']]]]] = None,
            storage_authentication_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AccountIdentityArgs']] identity: An `identity` block is documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountStorageAccountArgs']]]] storage_accounts: One or more `storage_account` blocks as defined below.
        :param pulumi.Input[str] storage_authentication_type: Specifies the storage authentication type. 
               Possible value is  `ManagedIdentity` or `System`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags assigned to the resource.
               ---
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["identity"] = identity
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_accounts"] = storage_accounts
        __props__["storage_authentication_type"] = storage_authentication_type
        __props__["tags"] = tags
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output['outputs.AccountIdentity']:
        """
        An `identity` block is documented below.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="storageAccounts")
    def storage_accounts(self) -> pulumi.Output[Sequence['outputs.AccountStorageAccount']]:
        """
        One or more `storage_account` blocks as defined below.
        """
        return pulumi.get(self, "storage_accounts")

    @property
    @pulumi.getter(name="storageAuthenticationType")
    def storage_authentication_type(self) -> pulumi.Output[str]:
        """
        Specifies the storage authentication type. 
        Possible value is  `ManagedIdentity` or `System`.
        """
        return pulumi.get(self, "storage_authentication_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags assigned to the resource.
        ---
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

