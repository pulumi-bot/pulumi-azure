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

__all__ = ['PolicyFileShareArgs', 'PolicyFileShare']

@pulumi.input_type
class PolicyFileShareArgs:
    def __init__(__self__, *,
                 backup: pulumi.Input['PolicyFileShareBackupArgs'],
                 recovery_vault_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 retention_daily: pulumi.Input['PolicyFileShareRetentionDailyArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PolicyFileShare resource.
        :param pulumi.Input['PolicyFileShareBackupArgs'] backup: Configures the Policy backup frequency and times as documented in the `backup` block below.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input['PolicyFileShareRetentionDailyArgs'] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below.
        :param pulumi.Input[str] name: Specifies the name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] timezone: Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
        """
        pulumi.set(__self__, "backup", backup)
        pulumi.set(__self__, "recovery_vault_name", recovery_vault_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "retention_daily", retention_daily)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter
    def backup(self) -> pulumi.Input['PolicyFileShareBackupArgs']:
        """
        Configures the Policy backup frequency and times as documented in the `backup` block below.
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: pulumi.Input['PolicyFileShareBackupArgs']):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter(name="recoveryVaultName")
    def recovery_vault_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "recovery_vault_name")

    @recovery_vault_name.setter
    def recovery_vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "recovery_vault_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="retentionDaily")
    def retention_daily(self) -> pulumi.Input['PolicyFileShareRetentionDailyArgs']:
        """
        Configures the policy daily retention as documented in the `retention_daily` block below.
        """
        return pulumi.get(self, "retention_daily")

    @retention_daily.setter
    def retention_daily(self, value: pulumi.Input['PolicyFileShareRetentionDailyArgs']):
        pulumi.set(self, "retention_daily", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the policy. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


class PolicyFileShare(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyFileShareArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure File Share Backup Policy within a Recovery Services vault.

        > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        vault = azure.recoveryservices.Vault("vault",
            location=rg.location,
            resource_group_name=rg.name,
            sku="Standard")
        policy = azure.backup.PolicyFileShare("policy",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            timezone="UTC",
            backup=azure.backup.PolicyFileShareBackupArgs(
                frequency="Daily",
                time="23:00",
            ),
            retention_daily=azure.backup.PolicyFileShareRetentionDailyArgs(
                count=10,
            ))
        ```

        ## Import

        Azure File Share Backup Policies can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:backup/policyFileShare:PolicyFileShare policy1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
        ```

        :param str resource_name: The name of the resource.
        :param PolicyFileShareArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup: Optional[pulumi.Input[pulumi.InputType['PolicyFileShareBackupArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_daily: Optional[pulumi.Input[pulumi.InputType['PolicyFileShareRetentionDailyArgs']]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure File Share Backup Policy within a Recovery Services vault.

        > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        vault = azure.recoveryservices.Vault("vault",
            location=rg.location,
            resource_group_name=rg.name,
            sku="Standard")
        policy = azure.backup.PolicyFileShare("policy",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            timezone="UTC",
            backup=azure.backup.PolicyFileShareBackupArgs(
                frequency="Daily",
                time="23:00",
            ),
            retention_daily=azure.backup.PolicyFileShareRetentionDailyArgs(
                count=10,
            ))
        ```

        ## Import

        Azure File Share Backup Policies can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:backup/policyFileShare:PolicyFileShare policy1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyFileShareBackupArgs']] backup: Configures the Policy backup frequency and times as documented in the `backup` block below.
        :param pulumi.Input[str] name: Specifies the name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyFileShareRetentionDailyArgs']] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below.
        :param pulumi.Input[str] timezone: Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyFileShareArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup: Optional[pulumi.Input[pulumi.InputType['PolicyFileShareBackupArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_daily: Optional[pulumi.Input[pulumi.InputType['PolicyFileShareRetentionDailyArgs']]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
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

            if backup is None and not opts.urn:
                raise TypeError("Missing required property 'backup'")
            __props__['backup'] = backup
            __props__['name'] = name
            if recovery_vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if retention_daily is None and not opts.urn:
                raise TypeError("Missing required property 'retention_daily'")
            __props__['retention_daily'] = retention_daily
            __props__['timezone'] = timezone
        super(PolicyFileShare, __self__).__init__(
            'azure:backup/policyFileShare:PolicyFileShare',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup: Optional[pulumi.Input[pulumi.InputType['PolicyFileShareBackupArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            recovery_vault_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            retention_daily: Optional[pulumi.Input[pulumi.InputType['PolicyFileShareRetentionDailyArgs']]] = None,
            timezone: Optional[pulumi.Input[str]] = None) -> 'PolicyFileShare':
        """
        Get an existing PolicyFileShare resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyFileShareBackupArgs']] backup: Configures the Policy backup frequency and times as documented in the `backup` block below.
        :param pulumi.Input[str] name: Specifies the name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyFileShareRetentionDailyArgs']] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below.
        :param pulumi.Input[str] timezone: Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup"] = backup
        __props__["name"] = name
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["retention_daily"] = retention_daily
        __props__["timezone"] = timezone
        return PolicyFileShare(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backup(self) -> pulumi.Output['outputs.PolicyFileShareBackup']:
        """
        Configures the Policy backup frequency and times as documented in the `backup` block below.
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the policy. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

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
        The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="retentionDaily")
    def retention_daily(self) -> pulumi.Output['outputs.PolicyFileShareRetentionDaily']:
        """
        Configures the policy daily retention as documented in the `retention_daily` block below.
        """
        return pulumi.get(self, "retention_daily")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
        """
        return pulumi.get(self, "timezone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

