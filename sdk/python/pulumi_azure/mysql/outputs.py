# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ServerStorageProfile',
    'ServerThreatDetectionPolicy',
]

@pulumi.output_type
class ServerStorageProfile(dict):
    auto_grow: Optional[str] = pulumi.output_property("autoGrow")
    backup_retention_days: Optional[float] = pulumi.output_property("backupRetentionDays")
    """
    Backup retention days for the server, supported values are between `7` and `35` days.
    """
    geo_redundant_backup: Optional[str] = pulumi.output_property("geoRedundantBackup")
    storage_mb: Optional[float] = pulumi.output_property("storageMb")
    """
    Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerThreatDetectionPolicy(dict):
    disabled_alerts: Optional[List[str]] = pulumi.output_property("disabledAlerts")
    """
    Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
    """
    email_account_admins: Optional[bool] = pulumi.output_property("emailAccountAdmins")
    """
    Should the account administrators be emailed when this alert is triggered?
    """
    email_addresses: Optional[List[str]] = pulumi.output_property("emailAddresses")
    """
    A list of email addresses which alerts should be sent to.
    """
    enabled: Optional[bool] = pulumi.output_property("enabled")
    """
    Is the policy enabled?
    """
    retention_days: Optional[float] = pulumi.output_property("retentionDays")
    """
    Specifies the number of days to keep in the Threat Detection audit logs.
    """
    storage_account_access_key: Optional[str] = pulumi.output_property("storageAccountAccessKey")
    """
    Specifies the identifier key of the Threat Detection audit storage account.
    """
    storage_endpoint: Optional[str] = pulumi.output_property("storageEndpoint")
    """
    Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


