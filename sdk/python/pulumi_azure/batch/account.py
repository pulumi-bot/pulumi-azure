# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Account']


class Account(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_reference: Optional[pulumi.Input[pulumi.InputType['AccountKeyVaultReferenceArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_allocation_mode: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Batch account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westeurope")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_batch_account_account = azure.batch.Account("exampleBatch/accountAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            pool_allocation_mode="BatchService",
            storage_account_id=example_account.id,
            tags={
                "env": "test",
            })
        ```

        ## Import

        Batch Account can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:batch/account:Account example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Batch/batchAccounts/account1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AccountKeyVaultReferenceArgs']] key_vault_reference: A `key_vault_reference` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the `UserSubscription` pool allocation mode.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_allocation_mode: Specifies the mode to use for pool allocation. Possible values are `BatchService` or `UserSubscription`. Defaults to `BatchService`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the storage account to use for the Batch account. If not specified, Azure Batch will manage the storage.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['key_vault_reference'] = key_vault_reference
            __props__['location'] = location
            __props__['name'] = name
            __props__['pool_allocation_mode'] = pool_allocation_mode
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['storage_account_id'] = storage_account_id
            __props__['tags'] = tags
            __props__['account_endpoint'] = None
            __props__['primary_access_key'] = None
            __props__['secondary_access_key'] = None
        super(Account, __self__).__init__(
            'azure:batch/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_endpoint: Optional[pulumi.Input[str]] = None,
            key_vault_reference: Optional[pulumi.Input[pulumi.InputType['AccountKeyVaultReferenceArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_allocation_mode: Optional[pulumi.Input[str]] = None,
            primary_access_key: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_access_key: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_endpoint: The account endpoint used to interact with the Batch service.
        :param pulumi.Input[pulumi.InputType['AccountKeyVaultReferenceArgs']] key_vault_reference: A `key_vault_reference` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the `UserSubscription` pool allocation mode.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_allocation_mode: Specifies the mode to use for pool allocation. Possible values are `BatchService` or `UserSubscription`. Defaults to `BatchService`.
        :param pulumi.Input[str] primary_access_key: The Batch account primary access key.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The Batch account secondary access key.
        :param pulumi.Input[str] storage_account_id: Specifies the storage account to use for the Batch account. If not specified, Azure Batch will manage the storage.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_endpoint"] = account_endpoint
        __props__["key_vault_reference"] = key_vault_reference
        __props__["location"] = location
        __props__["name"] = name
        __props__["pool_allocation_mode"] = pool_allocation_mode
        __props__["primary_access_key"] = primary_access_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["storage_account_id"] = storage_account_id
        __props__["tags"] = tags
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountEndpoint")
    def account_endpoint(self) -> pulumi.Output[str]:
        """
        The account endpoint used to interact with the Batch service.
        """
        return pulumi.get(self, "account_endpoint")

    @property
    @pulumi.getter(name="keyVaultReference")
    def key_vault_reference(self) -> pulumi.Output[Optional['outputs.AccountKeyVaultReference']]:
        """
        A `key_vault_reference` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the `UserSubscription` pool allocation mode.
        """
        return pulumi.get(self, "key_vault_reference")

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
        Specifies the name of the Batch account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolAllocationMode")
    def pool_allocation_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the mode to use for pool allocation. Possible values are `BatchService` or `UserSubscription`. Defaults to `BatchService`.
        """
        return pulumi.get(self, "pool_allocation_mode")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> pulumi.Output[str]:
        """
        The Batch account primary access key.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> pulumi.Output[str]:
        """
        The Batch account secondary access key.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[str]:
        """
        Specifies the storage account to use for the Batch account. If not specified, Azure Batch will manage the storage.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

