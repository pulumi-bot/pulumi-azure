# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Server']


class Server(pulumi.CustomResource):
    administrator_login: pulumi.Output[str] = pulumi.property("administratorLogin")
    """
    The Administrator Login for the MySQL Server. Required when `create_mode` is `Default`. Changing this forces a new resource to be created.
    """

    administrator_login_password: pulumi.Output[Optional[str]] = pulumi.property("administratorLoginPassword")
    """
    The Password associated with the `administrator_login` for the MySQL Server. Required when `create_mode` is `Default`.
    """

    auto_grow_enabled: pulumi.Output[bool] = pulumi.property("autoGrowEnabled")
    """
    Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
    """

    backup_retention_days: pulumi.Output[float] = pulumi.property("backupRetentionDays")
    """
    Backup retention days for the server, supported values are between `7` and `35` days.
    """

    create_mode: pulumi.Output[Optional[str]] = pulumi.property("createMode")
    """
    The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
    """

    creation_source_server_id: pulumi.Output[Optional[str]] = pulumi.property("creationSourceServerId")
    """
    For creation modes other than `Default`, the source server ID to use.
    """

    fqdn: pulumi.Output[str] = pulumi.property("fqdn")
    """
    The FQDN of the MySQL Server.
    """

    geo_redundant_backup_enabled: pulumi.Output[bool] = pulumi.property("geoRedundantBackupEnabled")
    """
    Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
    """

    infrastructure_encryption_enabled: pulumi.Output[Optional[bool]] = pulumi.property("infrastructureEncryptionEnabled")
    """
    Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
    """

    public_network_access_enabled: pulumi.Output[Optional[bool]] = pulumi.property("publicNetworkAccessEnabled")
    """
    Whether or not public network access is allowed for this server. Defaults to `true`.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
    """

    restore_point_in_time: pulumi.Output[Optional[str]] = pulumi.property("restorePointInTime")
    """
    When `create_mode` is `PointInTimeRestore`, specifies the point in time to restore from `creation_source_server_id`.
    """

    sku_name: pulumi.Output[str] = pulumi.property("skuName")
    """
    Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
    """

    ssl_enforcement: pulumi.Output[str] = pulumi.property("sslEnforcement")

    ssl_enforcement_enabled: pulumi.Output[Optional[bool]] = pulumi.property("sslEnforcementEnabled")
    """
    Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
    """

    ssl_minimal_tls_version_enforced: pulumi.Output[Optional[str]] = pulumi.property("sslMinimalTlsVersionEnforced")
    """
    The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
    """

    storage_mb: pulumi.Output[float] = pulumi.property("storageMb")
    """
    Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
    """

    storage_profile: pulumi.Output['outputs.ServerStorageProfile'] = pulumi.property("storageProfile")

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    threat_detection_policy: pulumi.Output[Optional['outputs.ServerThreatDetectionPolicy']] = pulumi.property("threatDetectionPolicy")
    """
    Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threat_detection_policy` block supports fields documented below.
    """

    version: pulumi.Output[str] = pulumi.property("version")
    """
    Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 auto_grow_enabled: Optional[pulumi.Input[bool]] = None,
                 backup_retention_days: Optional[pulumi.Input[float]] = None,
                 create_mode: Optional[pulumi.Input[str]] = None,
                 creation_source_server_id: Optional[pulumi.Input[str]] = None,
                 geo_redundant_backup_enabled: Optional[pulumi.Input[bool]] = None,
                 infrastructure_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 restore_point_in_time: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 ssl_enforcement: Optional[pulumi.Input[str]] = None,
                 ssl_enforcement_enabled: Optional[pulumi.Input[bool]] = None,
                 ssl_minimal_tls_version_enforced: Optional[pulumi.Input[str]] = None,
                 storage_mb: Optional[pulumi.Input[float]] = None,
                 storage_profile: Optional[pulumi.Input[pulumi.InputType['ServerStorageProfileArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 threat_detection_policy: Optional[pulumi.Input[pulumi.InputType['ServerThreatDetectionPolicyArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a MySQL Server.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_server = azure.mysql.Server("exampleServer",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            administrator_login="mysqladminun",
            administrator_login_password="H@Sh1CoR3!",
            sku_name="B_Gen5_2",
            storage_mb=5120,
            version="5.7",
            auto_grow_enabled=True,
            backup_retention_days=7,
            geo_redundant_backup_enabled=True,
            infrastructure_encryption_enabled=True,
            public_network_access_enabled=False,
            ssl_enforcement_enabled=True,
            ssl_minimal_tls_version_enforced="TLS1_2")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_login: The Administrator Login for the MySQL Server. Required when `create_mode` is `Default`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] administrator_login_password: The Password associated with the `administrator_login` for the MySQL Server. Required when `create_mode` is `Default`.
        :param pulumi.Input[bool] auto_grow_enabled: Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
        :param pulumi.Input[float] backup_retention_days: Backup retention days for the server, supported values are between `7` and `35` days.
        :param pulumi.Input[str] create_mode: The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
        :param pulumi.Input[str] creation_source_server_id: For creation modes other than `Default`, the source server ID to use.
        :param pulumi.Input[bool] geo_redundant_backup_enabled: Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
        :param pulumi.Input[bool] infrastructure_encryption_enabled: Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] restore_point_in_time: When `create_mode` is `PointInTimeRestore`, specifies the point in time to restore from `creation_source_server_id`.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
        :param pulumi.Input[bool] ssl_enforcement_enabled: Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
        :param pulumi.Input[str] ssl_minimal_tls_version_enforced: The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
        :param pulumi.Input[float] storage_mb: Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['ServerThreatDetectionPolicyArgs']] threat_detection_policy: Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threat_detection_policy` block supports fields documented below.
        :param pulumi.Input[str] version: Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
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

            __props__['administrator_login'] = administrator_login
            __props__['administrator_login_password'] = administrator_login_password
            __props__['auto_grow_enabled'] = auto_grow_enabled
            __props__['backup_retention_days'] = backup_retention_days
            __props__['create_mode'] = create_mode
            __props__['creation_source_server_id'] = creation_source_server_id
            __props__['geo_redundant_backup_enabled'] = geo_redundant_backup_enabled
            __props__['infrastructure_encryption_enabled'] = infrastructure_encryption_enabled
            __props__['location'] = location
            __props__['name'] = name
            __props__['public_network_access_enabled'] = public_network_access_enabled
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['restore_point_in_time'] = restore_point_in_time
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            if ssl_enforcement is not None:
                warnings.warn("this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.", DeprecationWarning)
                pulumi.log.warn("ssl_enforcement is deprecated: this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.")
            __props__['ssl_enforcement'] = ssl_enforcement
            __props__['ssl_enforcement_enabled'] = ssl_enforcement_enabled
            __props__['ssl_minimal_tls_version_enforced'] = ssl_minimal_tls_version_enforced
            __props__['storage_mb'] = storage_mb
            if storage_profile is not None:
                warnings.warn("all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.", DeprecationWarning)
                pulumi.log.warn("storage_profile is deprecated: all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.")
            __props__['storage_profile'] = storage_profile
            __props__['tags'] = tags
            __props__['threat_detection_policy'] = threat_detection_policy
            if version is None:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
            __props__['fqdn'] = None
        super(Server, __self__).__init__(
            'azure:mysql/server:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            administrator_login: Optional[pulumi.Input[str]] = None,
            administrator_login_password: Optional[pulumi.Input[str]] = None,
            auto_grow_enabled: Optional[pulumi.Input[bool]] = None,
            backup_retention_days: Optional[pulumi.Input[float]] = None,
            create_mode: Optional[pulumi.Input[str]] = None,
            creation_source_server_id: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            geo_redundant_backup_enabled: Optional[pulumi.Input[bool]] = None,
            infrastructure_encryption_enabled: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            restore_point_in_time: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            ssl_enforcement: Optional[pulumi.Input[str]] = None,
            ssl_enforcement_enabled: Optional[pulumi.Input[bool]] = None,
            ssl_minimal_tls_version_enforced: Optional[pulumi.Input[str]] = None,
            storage_mb: Optional[pulumi.Input[float]] = None,
            storage_profile: Optional[pulumi.Input[pulumi.InputType['ServerStorageProfileArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            threat_detection_policy: Optional[pulumi.Input[pulumi.InputType['ServerThreatDetectionPolicyArgs']]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Server':
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_login: The Administrator Login for the MySQL Server. Required when `create_mode` is `Default`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] administrator_login_password: The Password associated with the `administrator_login` for the MySQL Server. Required when `create_mode` is `Default`.
        :param pulumi.Input[bool] auto_grow_enabled: Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
        :param pulumi.Input[float] backup_retention_days: Backup retention days for the server, supported values are between `7` and `35` days.
        :param pulumi.Input[str] create_mode: The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
        :param pulumi.Input[str] creation_source_server_id: For creation modes other than `Default`, the source server ID to use.
        :param pulumi.Input[str] fqdn: The FQDN of the MySQL Server.
        :param pulumi.Input[bool] geo_redundant_backup_enabled: Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
        :param pulumi.Input[bool] infrastructure_encryption_enabled: Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] restore_point_in_time: When `create_mode` is `PointInTimeRestore`, specifies the point in time to restore from `creation_source_server_id`.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
        :param pulumi.Input[bool] ssl_enforcement_enabled: Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
        :param pulumi.Input[str] ssl_minimal_tls_version_enforced: The minimum TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
        :param pulumi.Input[float] storage_mb: Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['ServerThreatDetectionPolicyArgs']] threat_detection_policy: Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threat_detection_policy` block supports fields documented below.
        :param pulumi.Input[str] version: Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["administrator_login"] = administrator_login
        __props__["administrator_login_password"] = administrator_login_password
        __props__["auto_grow_enabled"] = auto_grow_enabled
        __props__["backup_retention_days"] = backup_retention_days
        __props__["create_mode"] = create_mode
        __props__["creation_source_server_id"] = creation_source_server_id
        __props__["fqdn"] = fqdn
        __props__["geo_redundant_backup_enabled"] = geo_redundant_backup_enabled
        __props__["infrastructure_encryption_enabled"] = infrastructure_encryption_enabled
        __props__["location"] = location
        __props__["name"] = name
        __props__["public_network_access_enabled"] = public_network_access_enabled
        __props__["resource_group_name"] = resource_group_name
        __props__["restore_point_in_time"] = restore_point_in_time
        __props__["sku_name"] = sku_name
        __props__["ssl_enforcement"] = ssl_enforcement
        __props__["ssl_enforcement_enabled"] = ssl_enforcement_enabled
        __props__["ssl_minimal_tls_version_enforced"] = ssl_minimal_tls_version_enforced
        __props__["storage_mb"] = storage_mb
        __props__["storage_profile"] = storage_profile
        __props__["tags"] = tags
        __props__["threat_detection_policy"] = threat_detection_policy
        __props__["version"] = version
        return Server(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

