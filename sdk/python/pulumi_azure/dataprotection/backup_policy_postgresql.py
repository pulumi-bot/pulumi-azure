# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BackupPolicyPostgresqlArgs', 'BackupPolicyPostgresql']

@pulumi.input_type
class BackupPolicyPostgresqlArgs:
    def __init__(__self__, *,
                 backup_repeating_time_intervals: pulumi.Input[Sequence[pulumi.Input[str]]],
                 default_retention_duration: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 vault_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 retention_rules: Optional[pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]]] = None):
        """
        The set of arguments for constructing a BackupPolicyPostgresql resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_repeating_time_intervals: Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] default_retention_duration: The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] vault_name: The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] name: The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]] retention_rules: One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        pulumi.set(__self__, "backup_repeating_time_intervals", backup_repeating_time_intervals)
        pulumi.set(__self__, "default_retention_duration", default_retention_duration)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "vault_name", vault_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_rules is not None:
            pulumi.set(__self__, "retention_rules", retention_rules)

    @property
    @pulumi.getter(name="backupRepeatingTimeIntervals")
    def backup_repeating_time_intervals(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "backup_repeating_time_intervals")

    @backup_repeating_time_intervals.setter
    def backup_repeating_time_intervals(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "backup_repeating_time_intervals", value)

    @property
    @pulumi.getter(name="defaultRetentionDuration")
    def default_retention_duration(self) -> pulumi.Input[str]:
        """
        The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "default_retention_duration")

    @default_retention_duration.setter
    def default_retention_duration(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_retention_duration", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> pulumi.Input[str]:
        """
        The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "vault_name")

    @vault_name.setter
    def vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vault_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionRules")
    def retention_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]]]:
        """
        One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "retention_rules")

    @retention_rules.setter
    def retention_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]]]):
        pulumi.set(self, "retention_rules", value)


@pulumi.input_type
class _BackupPolicyPostgresqlState:
    def __init__(__self__, *,
                 backup_repeating_time_intervals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_retention_duration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_rules: Optional[pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BackupPolicyPostgresql resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_repeating_time_intervals: Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] default_retention_duration: The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] name: The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]] retention_rules: One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] vault_name: The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        if backup_repeating_time_intervals is not None:
            pulumi.set(__self__, "backup_repeating_time_intervals", backup_repeating_time_intervals)
        if default_retention_duration is not None:
            pulumi.set(__self__, "default_retention_duration", default_retention_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if retention_rules is not None:
            pulumi.set(__self__, "retention_rules", retention_rules)
        if vault_name is not None:
            pulumi.set(__self__, "vault_name", vault_name)

    @property
    @pulumi.getter(name="backupRepeatingTimeIntervals")
    def backup_repeating_time_intervals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "backup_repeating_time_intervals")

    @backup_repeating_time_intervals.setter
    def backup_repeating_time_intervals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backup_repeating_time_intervals", value)

    @property
    @pulumi.getter(name="defaultRetentionDuration")
    def default_retention_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "default_retention_duration")

    @default_retention_duration.setter
    def default_retention_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_retention_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="retentionRules")
    def retention_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]]]:
        """
        One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "retention_rules")

    @retention_rules.setter
    def retention_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackupPolicyPostgresqlRetentionRuleArgs']]]]):
        pulumi.set(self, "retention_rules", value)

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "vault_name")

    @vault_name.setter
    def vault_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_name", value)


class BackupPolicyPostgresql(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_repeating_time_intervals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_retention_duration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupPolicyPostgresqlRetentionRuleArgs']]]]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Backup Policy to back up PostgreSQL.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        example_backup_vault = azure.dataprotection.BackupVault("exampleBackupVault",
            resource_group_name=rg.name,
            location=rg.location,
            datastore_type="VaultStore",
            redundancy="LocallyRedundant")
        example_backup_policy_postgresql = azure.dataprotection.BackupPolicyPostgresql("exampleBackupPolicyPostgresql",
            resource_group_name=rg.name,
            vault_name=example_backup_vault.name,
            backup_repeating_time_intervals=["R/2021-05-23T02:30:00+00:00/P1W"],
            default_retention_duration="P4M",
            retention_rules=[
                azure.dataprotection.BackupPolicyPostgresqlRetentionRuleArgs(
                    name="weekly",
                    duration="P6M",
                    priority=20,
                    criteria=azure.dataprotection.BackupPolicyPostgresqlRetentionRuleCriteriaArgs(
                        absolute_criteria="FirstOfWeek",
                    ),
                ),
                azure.dataprotection.BackupPolicyPostgresqlRetentionRuleArgs(
                    name="thursday",
                    duration="P1W",
                    priority=25,
                    criteria=azure.dataprotection.BackupPolicyPostgresqlRetentionRuleCriteriaArgs(
                        days_of_weeks=["Thursday"],
                        scheduled_backup_times=["2021-05-23T02:30:00Z"],
                    ),
                ),
                azure.dataprotection.BackupPolicyPostgresqlRetentionRuleArgs(
                    name="monthly",
                    duration="P1D",
                    priority=15,
                    criteria=azure.dataprotection.BackupPolicyPostgresqlRetentionRuleCriteriaArgs(
                        weeks_of_months=[
                            "First",
                            "Last",
                        ],
                        days_of_weeks=["Tuesday"],
                        scheduled_backup_times=["2021-05-23T02:30:00Z"],
                    ),
                ),
            ])
        ```

        ## Import

        Backup Policy PostgreSQLs can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dataprotection/backupPolicyPostgresql:BackupPolicyPostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_repeating_time_intervals: Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] default_retention_duration: The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] name: The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupPolicyPostgresqlRetentionRuleArgs']]]] retention_rules: One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] vault_name: The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupPolicyPostgresqlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Backup Policy to back up PostgreSQL.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        example_backup_vault = azure.dataprotection.BackupVault("exampleBackupVault",
            resource_group_name=rg.name,
            location=rg.location,
            datastore_type="VaultStore",
            redundancy="LocallyRedundant")
        example_backup_policy_postgresql = azure.dataprotection.BackupPolicyPostgresql("exampleBackupPolicyPostgresql",
            resource_group_name=rg.name,
            vault_name=example_backup_vault.name,
            backup_repeating_time_intervals=["R/2021-05-23T02:30:00+00:00/P1W"],
            default_retention_duration="P4M",
            retention_rules=[
                azure.dataprotection.BackupPolicyPostgresqlRetentionRuleArgs(
                    name="weekly",
                    duration="P6M",
                    priority=20,
                    criteria=azure.dataprotection.BackupPolicyPostgresqlRetentionRuleCriteriaArgs(
                        absolute_criteria="FirstOfWeek",
                    ),
                ),
                azure.dataprotection.BackupPolicyPostgresqlRetentionRuleArgs(
                    name="thursday",
                    duration="P1W",
                    priority=25,
                    criteria=azure.dataprotection.BackupPolicyPostgresqlRetentionRuleCriteriaArgs(
                        days_of_weeks=["Thursday"],
                        scheduled_backup_times=["2021-05-23T02:30:00Z"],
                    ),
                ),
                azure.dataprotection.BackupPolicyPostgresqlRetentionRuleArgs(
                    name="monthly",
                    duration="P1D",
                    priority=15,
                    criteria=azure.dataprotection.BackupPolicyPostgresqlRetentionRuleCriteriaArgs(
                        weeks_of_months=[
                            "First",
                            "Last",
                        ],
                        days_of_weeks=["Tuesday"],
                        scheduled_backup_times=["2021-05-23T02:30:00Z"],
                    ),
                ),
            ])
        ```

        ## Import

        Backup Policy PostgreSQLs can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dataprotection/backupPolicyPostgresql:BackupPolicyPostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
        ```

        :param str resource_name: The name of the resource.
        :param BackupPolicyPostgresqlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupPolicyPostgresqlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_repeating_time_intervals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_retention_duration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupPolicyPostgresqlRetentionRuleArgs']]]]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = BackupPolicyPostgresqlArgs.__new__(BackupPolicyPostgresqlArgs)

            if backup_repeating_time_intervals is None and not opts.urn:
                raise TypeError("Missing required property 'backup_repeating_time_intervals'")
            __props__.__dict__["backup_repeating_time_intervals"] = backup_repeating_time_intervals
            if default_retention_duration is None and not opts.urn:
                raise TypeError("Missing required property 'default_retention_duration'")
            __props__.__dict__["default_retention_duration"] = default_retention_duration
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["retention_rules"] = retention_rules
            if vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'vault_name'")
            __props__.__dict__["vault_name"] = vault_name
        super(BackupPolicyPostgresql, __self__).__init__(
            'azure:dataprotection/backupPolicyPostgresql:BackupPolicyPostgresql',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_repeating_time_intervals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            default_retention_duration: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            retention_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupPolicyPostgresqlRetentionRuleArgs']]]]] = None,
            vault_name: Optional[pulumi.Input[str]] = None) -> 'BackupPolicyPostgresql':
        """
        Get an existing BackupPolicyPostgresql resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_repeating_time_intervals: Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] default_retention_duration: The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] name: The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupPolicyPostgresqlRetentionRuleArgs']]]] retention_rules: One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param pulumi.Input[str] vault_name: The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupPolicyPostgresqlState.__new__(_BackupPolicyPostgresqlState)

        __props__.__dict__["backup_repeating_time_intervals"] = backup_repeating_time_intervals
        __props__.__dict__["default_retention_duration"] = default_retention_duration
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["retention_rules"] = retention_rules
        __props__.__dict__["vault_name"] = vault_name
        return BackupPolicyPostgresql(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupRepeatingTimeIntervals")
    def backup_repeating_time_intervals(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies a list of repeating time interval. It supports weekly back. It should follow `ISO 8601` repeating time interval. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "backup_repeating_time_intervals")

    @property
    @pulumi.getter(name="defaultRetentionDuration")
    def default_retention_duration(self) -> pulumi.Output[str]:
        """
        The duration of default retention rule. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "default_retention_duration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Backup Policy PostgreSQL. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="retentionRules")
    def retention_rules(self) -> pulumi.Output[Optional[Sequence['outputs.BackupPolicyPostgresqlRetentionRule']]]:
        """
        One or more `retention_rule` blocks as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "retention_rules")

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> pulumi.Output[str]:
        """
        The name of the Backup Vault where the Backup Policy PostgreSQL should exist. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "vault_name")
