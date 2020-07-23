# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.input_type
class DatabaseExtendedAuditingPolicyArgs:
    storage_account_access_key: pulumi.Input[str] = pulumi.input_property("storageAccountAccessKey")
    """
    Specifies the access key to use for the auditing storage account.
    """
    storage_endpoint: pulumi.Input[str] = pulumi.input_property("storageEndpoint")
    """
    Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
    """
    retention_in_days: Optional[pulumi.Input[float]] = pulumi.input_property("retentionInDays")
    """
    Specifies the number of days to retain logs for in the storage account.
    """
    storage_account_access_key_is_secondary: Optional[pulumi.Input[bool]] = pulumi.input_property("storageAccountAccessKeyIsSecondary")
    """
    Specifies whether `storage_account_access_key` value is the storage's secondary key.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, storage_account_access_key: pulumi.Input[str], storage_endpoint: pulumi.Input[str], retention_in_days: Optional[pulumi.Input[float]] = None, storage_account_access_key_is_secondary: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[str] storage_account_access_key: Specifies the access key to use for the auditing storage account.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
        :param pulumi.Input[float] retention_in_days: Specifies the number of days to retain logs for in the storage account.
        :param pulumi.Input[bool] storage_account_access_key_is_secondary: Specifies whether `storage_account_access_key` value is the storage's secondary key.
        """
        __self__.storage_account_access_key = storage_account_access_key
        __self__.storage_endpoint = storage_endpoint
        __self__.retention_in_days = retention_in_days
        __self__.storage_account_access_key_is_secondary = storage_account_access_key_is_secondary

@pulumi.input_type
class DatabaseThreatDetectionPolicyArgs:
    disabled_alerts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("disabledAlerts")
    """
    Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
    """
    email_account_admins: Optional[pulumi.Input[str]] = pulumi.input_property("emailAccountAdmins")
    """
    Should the account administrators be emailed when this alert is triggered?
    """
    email_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("emailAddresses")
    """
    A list of email addresses which alerts should be sent to.
    """
    retention_days: Optional[pulumi.Input[float]] = pulumi.input_property("retentionDays")
    """
    Specifies the number of days to keep in the Threat Detection audit logs.
    """
    state: Optional[pulumi.Input[str]] = pulumi.input_property("state")
    """
    The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
    """
    storage_account_access_key: Optional[pulumi.Input[str]] = pulumi.input_property("storageAccountAccessKey")
    """
    Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
    """
    storage_endpoint: Optional[pulumi.Input[str]] = pulumi.input_property("storageEndpoint")
    """
    Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
    """
    use_server_default: Optional[pulumi.Input[str]] = pulumi.input_property("useServerDefault")
    """
    Should the default server policy be used? Defaults to `Disabled`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, disabled_alerts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, email_account_admins: Optional[pulumi.Input[str]] = None, email_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, retention_days: Optional[pulumi.Input[float]] = None, state: Optional[pulumi.Input[str]] = None, storage_account_access_key: Optional[pulumi.Input[str]] = None, storage_endpoint: Optional[pulumi.Input[str]] = None, use_server_default: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] disabled_alerts: Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
        :param pulumi.Input[str] email_account_admins: Should the account administrators be emailed when this alert is triggered?
        :param pulumi.Input[List[pulumi.Input[str]]] email_addresses: A list of email addresses which alerts should be sent to.
        :param pulumi.Input[float] retention_days: Specifies the number of days to keep in the Threat Detection audit logs.
        :param pulumi.Input[str] state: The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        :param pulumi.Input[str] use_server_default: Should the default server policy be used? Defaults to `Disabled`.
        """
        __self__.disabled_alerts = disabled_alerts
        __self__.email_account_admins = email_account_admins
        __self__.email_addresses = email_addresses
        __self__.retention_days = retention_days
        __self__.state = state
        __self__.storage_account_access_key = storage_account_access_key
        __self__.storage_endpoint = storage_endpoint
        __self__.use_server_default = use_server_default

@pulumi.input_type
class DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArgs:
    results: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("results")
    """
    A list representing a result of the baseline.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, results: pulumi.Input[List[pulumi.Input[str]]]) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] results: A list representing a result of the baseline.
        """
        __self__.results = results

@pulumi.input_type
class ElasticPoolPerDatabaseSettingsArgs:
    max_capacity: pulumi.Input[float] = pulumi.input_property("maxCapacity")
    """
    The maximum capacity any one database can consume.
    """
    min_capacity: pulumi.Input[float] = pulumi.input_property("minCapacity")
    """
    The minimum capacity all databases are guaranteed.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, max_capacity: pulumi.Input[float], min_capacity: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[float] max_capacity: The maximum capacity any one database can consume.
        :param pulumi.Input[float] min_capacity: The minimum capacity all databases are guaranteed.
        """
        __self__.max_capacity = max_capacity
        __self__.min_capacity = min_capacity

@pulumi.input_type
class ElasticPoolSkuArgs:
    capacity: pulumi.Input[float] = pulumi.input_property("capacity")
    """
    The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
    """
    tier: pulumi.Input[str] = pulumi.input_property("tier")
    """
    The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
    """
    family: Optional[pulumi.Input[str]] = pulumi.input_property("family")
    """
    The `family` of hardware `Gen4` or `Gen5`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, capacity: pulumi.Input[float], name: pulumi.Input[str], tier: pulumi.Input[str], family: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[float] capacity: The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
        :param pulumi.Input[str] name: Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
        :param pulumi.Input[str] tier: The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
        :param pulumi.Input[str] family: The `family` of hardware `Gen4` or `Gen5`.
        """
        __self__.capacity = capacity
        __self__.name = name
        __self__.tier = tier
        __self__.family = family

@pulumi.input_type
class ServerAzureadAdministratorArgs:
    login_username: pulumi.Input[str] = pulumi.input_property("loginUsername")
    """
    (Required)  The login username of the Azure AD Administrator of this SQL Server.
    """
    object_id: pulumi.Input[str] = pulumi.input_property("objectId")
    """
    (Required) The object id of the Azure AD Administrator of this SQL Server.
    """
    tenant_id: Optional[pulumi.Input[str]] = pulumi.input_property("tenantId")
    """
    (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, login_username: pulumi.Input[str], object_id: pulumi.Input[str], tenant_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] login_username: (Required)  The login username of the Azure AD Administrator of this SQL Server.
        :param pulumi.Input[str] object_id: (Required) The object id of the Azure AD Administrator of this SQL Server.
        :param pulumi.Input[str] tenant_id: (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
        """
        __self__.login_username = login_username
        __self__.object_id = object_id
        __self__.tenant_id = tenant_id

@pulumi.input_type
class ServerExtendedAuditingPolicyArgs:
    storage_account_access_key: pulumi.Input[str] = pulumi.input_property("storageAccountAccessKey")
    """
    (Required)  Specifies the access key to use for the auditing storage account.
    """
    storage_endpoint: pulumi.Input[str] = pulumi.input_property("storageEndpoint")
    """
    (Required) Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
    """
    retention_in_days: Optional[pulumi.Input[float]] = pulumi.input_property("retentionInDays")
    """
    (Optional) Specifies the number of days to retain logs for in the storage account.
    """
    storage_account_access_key_is_secondary: Optional[pulumi.Input[bool]] = pulumi.input_property("storageAccountAccessKeyIsSecondary")
    """
    (Optional) Specifies whether `storage_account_access_key` value is the storage's secondary key.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, storage_account_access_key: pulumi.Input[str], storage_endpoint: pulumi.Input[str], retention_in_days: Optional[pulumi.Input[float]] = None, storage_account_access_key_is_secondary: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[str] storage_account_access_key: (Required)  Specifies the access key to use for the auditing storage account.
        :param pulumi.Input[str] storage_endpoint: (Required) Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
        :param pulumi.Input[float] retention_in_days: (Optional) Specifies the number of days to retain logs for in the storage account.
        :param pulumi.Input[bool] storage_account_access_key_is_secondary: (Optional) Specifies whether `storage_account_access_key` value is the storage's secondary key.
        """
        __self__.storage_account_access_key = storage_account_access_key
        __self__.storage_endpoint = storage_endpoint
        __self__.retention_in_days = retention_in_days
        __self__.storage_account_access_key_is_secondary = storage_account_access_key_is_secondary

@pulumi.input_type
class ServerIdentityArgs:
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    Specifies the identity type of the Microsoft SQL Server. At this time the only allowed value is `SystemAssigned`.
    """
    principal_id: Optional[pulumi.Input[str]] = pulumi.input_property("principalId")
    """
    The Principal ID for the Service Principal associated with the Identity of this SQL Server.
    """
    tenant_id: Optional[pulumi.Input[str]] = pulumi.input_property("tenantId")
    """
    (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, type: pulumi.Input[str], principal_id: Optional[pulumi.Input[str]] = None, tenant_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] type: Specifies the identity type of the Microsoft SQL Server. At this time the only allowed value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        :param pulumi.Input[str] tenant_id: (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
        """
        __self__.type = type
        __self__.principal_id = principal_id
        __self__.tenant_id = tenant_id

@pulumi.input_type
class ServerVulnerabilityAssessmentRecurringScansArgs:
    email_subscription_admins: Optional[pulumi.Input[bool]] = pulumi.input_property("emailSubscriptionAdmins")
    """
    Boolean flag which specifies if the schedule scan notification will be sent to the subscription administrators. Defaults to `false`.
    """
    emails: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("emails")
    """
    Specifies an array of e-mail addresses to which the scan notification is sent.
    """
    enabled: Optional[pulumi.Input[bool]] = pulumi.input_property("enabled")
    """
    Boolean flag which specifies if recurring scans is enabled or disabled. Defaults to `false`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, email_subscription_admins: Optional[pulumi.Input[bool]] = None, emails: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, enabled: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[bool] email_subscription_admins: Boolean flag which specifies if the schedule scan notification will be sent to the subscription administrators. Defaults to `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] emails: Specifies an array of e-mail addresses to which the scan notification is sent.
        :param pulumi.Input[bool] enabled: Boolean flag which specifies if recurring scans is enabled or disabled. Defaults to `false`.
        """
        __self__.email_subscription_admins = email_subscription_admins
        __self__.emails = emails
        __self__.enabled = enabled

@pulumi.input_type
class VirtualMachineAutoPatchingArgs:
    day_of_week: pulumi.Input[str] = pulumi.input_property("dayOfWeek")
    """
    The day of week to apply the patch on.
    """
    maintenance_window_duration_in_minutes: pulumi.Input[float] = pulumi.input_property("maintenanceWindowDurationInMinutes")
    """
    The size of the Maintenance Window in minutes.
    """
    maintenance_window_starting_hour: pulumi.Input[float] = pulumi.input_property("maintenanceWindowStartingHour")
    """
    The Hour, in the Virtual Machine Time-Zone when the patching maintenance window should begin.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, day_of_week: pulumi.Input[str], maintenance_window_duration_in_minutes: pulumi.Input[float], maintenance_window_starting_hour: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[str] day_of_week: The day of week to apply the patch on.
        :param pulumi.Input[float] maintenance_window_duration_in_minutes: The size of the Maintenance Window in minutes.
        :param pulumi.Input[float] maintenance_window_starting_hour: The Hour, in the Virtual Machine Time-Zone when the patching maintenance window should begin.
        """
        __self__.day_of_week = day_of_week
        __self__.maintenance_window_duration_in_minutes = maintenance_window_duration_in_minutes
        __self__.maintenance_window_starting_hour = maintenance_window_starting_hour

@pulumi.input_type
class VirtualMachineKeyVaultCredentialArgs:
    key_vault_url: pulumi.Input[str] = pulumi.input_property("keyVaultUrl")
    """
    The azure Key Vault url. Changing this forces a new resource to be created.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The credential name.
    """
    service_principal_name: pulumi.Input[str] = pulumi.input_property("servicePrincipalName")
    """
    The service principal name to access key vault. Changing this forces a new resource to be created.
    """
    service_principal_secret: pulumi.Input[str] = pulumi.input_property("servicePrincipalSecret")
    """
    The service principal name secret to access key vault. Changing this forces a new resource to be created.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key_vault_url: pulumi.Input[str], name: pulumi.Input[str], service_principal_name: pulumi.Input[str], service_principal_secret: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] key_vault_url: The azure Key Vault url. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The credential name.
        :param pulumi.Input[str] service_principal_name: The service principal name to access key vault. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_principal_secret: The service principal name secret to access key vault. Changing this forces a new resource to be created.
        """
        __self__.key_vault_url = key_vault_url
        __self__.name = name
        __self__.service_principal_name = service_principal_name
        __self__.service_principal_secret = service_principal_secret
