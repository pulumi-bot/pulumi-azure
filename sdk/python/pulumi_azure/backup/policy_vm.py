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

__all__ = ['PolicyVM']


class PolicyVM(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup: Optional[pulumi.Input[pulumi.InputType['PolicyVMBackupArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recovery_vault_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_daily: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionDailyArgs']]] = None,
                 retention_monthly: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionMonthlyArgs']]] = None,
                 retention_weekly: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionWeeklyArgs']]] = None,
                 retention_yearly: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionYearlyArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Backup VM Backup Policy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyVMBackupArgs']] backup: Configures the Policy backup frequency, times & days as documented in the `backup` block below.
        :param pulumi.Input[str] name: Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionDailyArgs']] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionMonthlyArgs']] retention_monthly: Configures the policy monthly retention as documented in the `retention_monthly` block below.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionWeeklyArgs']] retention_weekly: Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionYearlyArgs']] retention_yearly: Configures the policy yearly retention as documented in the `retention_yearly` block below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`
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
            __props__['retention_daily'] = retention_daily
            __props__['retention_monthly'] = retention_monthly
            __props__['retention_weekly'] = retention_weekly
            __props__['retention_yearly'] = retention_yearly
            __props__['tags'] = tags
            __props__['timezone'] = timezone
        super(PolicyVM, __self__).__init__(
            'azure:backup/policyVM:PolicyVM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup: Optional[pulumi.Input[pulumi.InputType['PolicyVMBackupArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            recovery_vault_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            retention_daily: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionDailyArgs']]] = None,
            retention_monthly: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionMonthlyArgs']]] = None,
            retention_weekly: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionWeeklyArgs']]] = None,
            retention_yearly: Optional[pulumi.Input[pulumi.InputType['PolicyVMRetentionYearlyArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            timezone: Optional[pulumi.Input[str]] = None) -> 'PolicyVM':
        """
        Get an existing PolicyVM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyVMBackupArgs']] backup: Configures the Policy backup frequency, times & days as documented in the `backup` block below.
        :param pulumi.Input[str] name: Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionDailyArgs']] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionMonthlyArgs']] retention_monthly: Configures the policy monthly retention as documented in the `retention_monthly` block below.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionWeeklyArgs']] retention_weekly: Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        :param pulumi.Input[pulumi.InputType['PolicyVMRetentionYearlyArgs']] retention_yearly: Configures the policy yearly retention as documented in the `retention_yearly` block below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup"] = backup
        __props__["name"] = name
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["retention_daily"] = retention_daily
        __props__["retention_monthly"] = retention_monthly
        __props__["retention_weekly"] = retention_weekly
        __props__["retention_yearly"] = retention_yearly
        __props__["tags"] = tags
        __props__["timezone"] = timezone
        return PolicyVM(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backup(self) -> pulumi.Output['outputs.PolicyVMBackup']:
        """
        Configures the Policy backup frequency, times & days as documented in the `backup` block below.
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
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
    def retention_daily(self) -> pulumi.Output[Optional['outputs.PolicyVMRetentionDaily']]:
        """
        Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        """
        return pulumi.get(self, "retention_daily")

    @property
    @pulumi.getter(name="retentionMonthly")
    def retention_monthly(self) -> pulumi.Output[Optional['outputs.PolicyVMRetentionMonthly']]:
        """
        Configures the policy monthly retention as documented in the `retention_monthly` block below.
        """
        return pulumi.get(self, "retention_monthly")

    @property
    @pulumi.getter(name="retentionWeekly")
    def retention_weekly(self) -> pulumi.Output[Optional['outputs.PolicyVMRetentionWeekly']]:
        """
        Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        """
        return pulumi.get(self, "retention_weekly")

    @property
    @pulumi.getter(name="retentionYearly")
    def retention_yearly(self) -> pulumi.Output[Optional['outputs.PolicyVMRetentionYearly']]:
        """
        Configures the policy yearly retention as documented in the `retention_yearly` block below.
        """
        return pulumi.get(self, "retention_yearly")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the timezone. Defaults to `UTC`
        """
        return pulumi.get(self, "timezone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

