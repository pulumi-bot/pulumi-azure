# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ContainerStorageAccount(pulumi.CustomResource):
    recovery_vault_name: pulumi.Output[str]
    """
    The name of the vault where the storage account will be registered.
    """
    resource_group_name: pulumi.Output[str]
    """
    Name of the resource group where the vault is located.
    """
    storage_account_id: pulumi.Output[str]
    """
    Azure Resource ID of the storage account to be registered
    """
    def __init__(__self__, resource_name, opts=None, recovery_vault_name=None, resource_group_name=None, storage_account_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages registration of a storage account with Azure Backup. Storage accounts must be registered with an Azure Recovery Vault in order to backup file shares within the storage account. Registering a storage account with a vault creates what is known as a protection container within Azure Recovery Services. Once the container is created, Azure file shares within the storage account can be backed up using the `backup.ProtectedFileShare` resource.

        > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West US")
        vault = azure.recoveryservices.Vault("vault",
            location=rg.location,
            resource_group_name=rg.name,
            sku="Standard")
        sa = azure.storage.Account("sa",
            location=rg.location,
            resource_group_name=rg.name,
            account_tier="Standard",
            account_replication_type="LRS")
        container = azure.backup.ContainerStorageAccount("container",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            storage_account_id=sa.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault where the storage account will be registered.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault is located.
        :param pulumi.Input[str] storage_account_id: Azure Resource ID of the storage account to be registered
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_account_id is None:
                raise TypeError("Missing required property 'storage_account_id'")
            __props__['storage_account_id'] = storage_account_id
        super(ContainerStorageAccount, __self__).__init__(
            'azure:backup/containerStorageAccount:ContainerStorageAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, recovery_vault_name=None, resource_group_name=None, storage_account_id=None):
        """
        Get an existing ContainerStorageAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault where the storage account will be registered.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault is located.
        :param pulumi.Input[str] storage_account_id: Azure Resource ID of the storage account to be registered
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_account_id"] = storage_account_id
        return ContainerStorageAccount(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
