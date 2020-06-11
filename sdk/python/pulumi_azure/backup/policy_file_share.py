# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PolicyFileShare(pulumi.CustomResource):
    backup: pulumi.Output[dict]
    """
    Configures the Policy backup frequency and times as documented in the `backup` block below.

      * `frequency` (`str`) - Sets the backup frequency. Currently, only `Daily` is supported
      * `time` (`str`)
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the policy. Changing this forces a new resource to be created.
    """
    recovery_vault_name: pulumi.Output[str]
    """
    Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
    """
    retention_daily: pulumi.Output[dict]
    """
    Configures the policy daily retention as documented in the `retention_daily` block below.

      * `count` (`float`) - The number of daily backups to keep. Must be between `1` and `180` (inclusive)
    """
    timezone: pulumi.Output[str]
    """
    Specifies the timezone. Defaults to `UTC`
    """
    def __init__(__self__, resource_name, opts=None, backup=None, name=None, recovery_vault_name=None, resource_group_name=None, retention_daily=None, timezone=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure File Share Backup Policy within a Recovery Services vault.

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
        policy = azure.backup.PolicyFileShare("policy",
            resource_group_name=rg.name,
            recovery_vault_name=vault.name,
            timezone="UTC",
            backup={
                "frequency": "Daily",
                "time": "23:00",
            },
            retention_daily={
                "count": 10,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] backup: Configures the Policy backup frequency and times as documented in the `backup` block below.
        :param pulumi.Input[str] name: Specifies the name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`

        The **backup** object supports the following:

          * `frequency` (`pulumi.Input[str]`) - Sets the backup frequency. Currently, only `Daily` is supported
          * `time` (`pulumi.Input[str]`)

        The **retention_daily** object supports the following:

          * `count` (`pulumi.Input[float]`) - The number of daily backups to keep. Must be between `1` and `180` (inclusive)
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

            if backup is None:
                raise TypeError("Missing required property 'backup'")
            __props__['backup'] = backup
            __props__['name'] = name
            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if retention_daily is None:
                raise TypeError("Missing required property 'retention_daily'")
            __props__['retention_daily'] = retention_daily
            __props__['timezone'] = timezone
        super(PolicyFileShare, __self__).__init__(
            'azure:backup/policyFileShare:PolicyFileShare',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backup=None, name=None, recovery_vault_name=None, resource_group_name=None, retention_daily=None, timezone=None):
        """
        Get an existing PolicyFileShare resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] backup: Configures the Policy backup frequency and times as documented in the `backup` block below.
        :param pulumi.Input[str] name: Specifies the name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`

        The **backup** object supports the following:

          * `frequency` (`pulumi.Input[str]`) - Sets the backup frequency. Currently, only `Daily` is supported
          * `time` (`pulumi.Input[str]`)

        The **retention_daily** object supports the following:

          * `count` (`pulumi.Input[float]`) - The number of daily backups to keep. Must be between `1` and `180` (inclusive)
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

