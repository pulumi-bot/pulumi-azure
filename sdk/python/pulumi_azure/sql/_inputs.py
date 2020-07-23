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
class DatabaseImportArgs:
    administrator_login: pulumi.Input[str] = pulumi.input_property("administratorLogin")
    """
    Specifies the name of the SQL administrator.
    """
    administrator_login_password: pulumi.Input[str] = pulumi.input_property("administratorLoginPassword")
    """
    Specifies the password of the SQL administrator.
    """
    authentication_type: pulumi.Input[str] = pulumi.input_property("authenticationType")
    """
    Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
    """
    storage_key: pulumi.Input[str] = pulumi.input_property("storageKey")
    """
    Specifies the access key for the storage account.
    """
    storage_key_type: pulumi.Input[str] = pulumi.input_property("storageKeyType")
    """
    Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
    """
    storage_uri: pulumi.Input[str] = pulumi.input_property("storageUri")
    """
    Specifies the blob URI of the .bacpac file.
    """
    operation_mode: Optional[pulumi.Input[str]] = pulumi.input_property("operationMode")
    """
    Specifies the type of import operation being performed. The only allowable value is `Import`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, administrator_login: pulumi.Input[str], administrator_login_password: pulumi.Input[str], authentication_type: pulumi.Input[str], storage_key: pulumi.Input[str], storage_key_type: pulumi.Input[str], storage_uri: pulumi.Input[str], operation_mode: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] administrator_login: Specifies the name of the SQL administrator.
        :param pulumi.Input[str] administrator_login_password: Specifies the password of the SQL administrator.
        :param pulumi.Input[str] authentication_type: Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
        :param pulumi.Input[str] storage_key: Specifies the access key for the storage account.
        :param pulumi.Input[str] storage_key_type: Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
        :param pulumi.Input[str] storage_uri: Specifies the blob URI of the .bacpac file.
        :param pulumi.Input[str] operation_mode: Specifies the type of import operation being performed. The only allowable value is `Import`.
        """
        __self__.administrator_login = administrator_login
        __self__.administrator_login_password = administrator_login_password
        __self__.authentication_type = authentication_type
        __self__.storage_key = storage_key
        __self__.storage_key_type = storage_key_type
        __self__.storage_uri = storage_uri
        __self__.operation_mode = operation_mode

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
class FailoverGroupPartnerServerArgs:
    id: pulumi.Input[str] = pulumi.input_property("id")
    """
    the SQL server ID
    """
    location: Optional[pulumi.Input[str]] = pulumi.input_property("location")
    """
    the location of the failover group.
    """
    role: Optional[pulumi.Input[str]] = pulumi.input_property("role")
    """
    local replication role of the failover group instance.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, id: pulumi.Input[str], location: Optional[pulumi.Input[str]] = None, role: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] id: the SQL server ID
        :param pulumi.Input[str] location: the location of the failover group.
        :param pulumi.Input[str] role: local replication role of the failover group instance.
        """
        __self__.id = id
        __self__.location = location
        __self__.role = role

@pulumi.input_type
class FailoverGroupReadWriteEndpointFailoverPolicyArgs:
    mode: pulumi.Input[str] = pulumi.input_property("mode")
    """
    the failover mode. Possible values are `Manual`, `Automatic`
    """
    grace_minutes: Optional[pulumi.Input[float]] = pulumi.input_property("graceMinutes")
    """
    Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, mode: pulumi.Input[str], grace_minutes: Optional[pulumi.Input[float]] = None) -> None:
        """
        :param pulumi.Input[str] mode: the failover mode. Possible values are `Manual`, `Automatic`
        :param pulumi.Input[float] grace_minutes: Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
        """
        __self__.mode = mode
        __self__.grace_minutes = grace_minutes

@pulumi.input_type
class FailoverGroupReadonlyEndpointFailoverPolicyArgs:
    mode: pulumi.Input[str] = pulumi.input_property("mode")
    """
    Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, mode: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] mode: Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
        """
        __self__.mode = mode

@pulumi.input_type
class SqlServerExtendedAuditingPolicyArgs:
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
class SqlServerIdentityArgs:
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
    The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, type: pulumi.Input[str], principal_id: Optional[pulumi.Input[str]] = None, tenant_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] type: Specifies the identity type of the Microsoft SQL Server. At this time the only allowed value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        :param pulumi.Input[str] tenant_id: The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
        """
        __self__.type = type
        __self__.principal_id = principal_id
        __self__.tenant_id = tenant_id
