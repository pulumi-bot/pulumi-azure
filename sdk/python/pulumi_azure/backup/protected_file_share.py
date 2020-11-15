# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ProtectedFileShare']


class ProtectedFileShare(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_policy_id: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_file_share_name: Optional[pulumi.Input[str]] = None,
                 source_storage_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Backup Protected File Share to enable backups for file shares within an Azure Storage Account

        > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)

        > **NOTE** Azure Backup for Azure File Shares does not support Soft Delete at this time. Deleting this resource will also delete all associated backup data. Please exercise caution. Consider using [`protect`](https://www.pulumi.com/docs/intro/concepts/programming-model/#protect) to guard against accidental deletion.

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
        example_share = azure.storage.Share("exampleShare", storage_account_name=sa.name)
        protection_container = azure.backup.ContainerStorageAccount("protection-container",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            storage_account_id=sa.id)
        example_policy_file_share = azure.backup.PolicyFileShare("examplePolicyFileShare",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            backup=azure.backup.PolicyFileShareBackupArgs(
                frequency="Daily",
                time="23:00",
            ),
            retention_daily=azure.backup.PolicyFileShareRetentionDailyArgs(
                count=10,
            ))
        share1 = azure.backup.ProtectedFileShare("share1",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            source_storage_account_id=protection_container.storage_account_id,
            source_file_share_name=example_share.name,
            backup_policy_id=example_policy_file_share.id)
        ```

        ## Import

        Azure Backup Protected File Shares can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:backup/protectedFileShare:ProtectedFileShare item1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/StorageContainer;storage;group2;example-storage-account/protectedItems/AzureFileShare;example-share"
        ```

         Note the ID requires quoting as there are semicolons

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_policy_id: Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_file_share_name: Specifies the name of the file share to backup. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_storage_account_id: Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
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

            if backup_policy_id is None:
                raise TypeError("Missing required property 'backup_policy_id'")
            __props__['backup_policy_id'] = backup_policy_id
            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_file_share_name is None:
                raise TypeError("Missing required property 'source_file_share_name'")
            __props__['source_file_share_name'] = source_file_share_name
            if source_storage_account_id is None:
                raise TypeError("Missing required property 'source_storage_account_id'")
            __props__['source_storage_account_id'] = source_storage_account_id
        super(ProtectedFileShare, __self__).__init__(
            'azure:backup/protectedFileShare:ProtectedFileShare',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_policy_id: Optional[pulumi.Input[str]] = None,
            recovery_vault_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            source_file_share_name: Optional[pulumi.Input[str]] = None,
            source_storage_account_id: Optional[pulumi.Input[str]] = None) -> 'ProtectedFileShare':
        """
        Get an existing ProtectedFileShare resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_policy_id: Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_file_share_name: Specifies the name of the file share to backup. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_storage_account_id: Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup_policy_id"] = backup_policy_id
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["source_file_share_name"] = source_file_share_name
        __props__["source_storage_account_id"] = source_storage_account_id
        return ProtectedFileShare(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPolicyId")
    def backup_policy_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
        """
        return pulumi.get(self, "backup_policy_id")

    @property
    @pulumi.getter(name="recoveryVaultName")
    def recovery_vault_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "recovery_vault_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sourceFileShareName")
    def source_file_share_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the file share to backup. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_file_share_name")

    @property
    @pulumi.getter(name="sourceStorageAccountId")
    def source_storage_account_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_storage_account_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

