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

__all__ = ['Database']


class Database(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collation: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 elastic_pool_name: Optional[pulumi.Input[str]] = None,
                 extended_auditing_policy: Optional[pulumi.Input[pulumi.InputType['DatabaseExtendedAuditingPolicyArgs']]] = None,
                 import_: Optional[pulumi.Input[pulumi.InputType['DatabaseImportArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_size_bytes: Optional[pulumi.Input[str]] = None,
                 max_size_gb: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_scale: Optional[pulumi.Input[bool]] = None,
                 requested_service_objective_id: Optional[pulumi.Input[str]] = None,
                 requested_service_objective_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 restore_point_in_time: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 source_database_deletion_date: Optional[pulumi.Input[str]] = None,
                 source_database_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 threat_detection_policy: Optional[pulumi.Input[pulumi.InputType['DatabaseThreatDetectionPolicyArgs']]] = None,
                 zone_redundant: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows you to manage an Azure SQL Database

        > **NOTE:** The Database Extended Auditing Policy Can be set inline here as well as with the mssql_database_extended_auditing_policy resource resource. You can only use one or the other and using both will cause a conflict.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_sql_server = azure.sql.SqlServer("exampleSqlServer",
            resource_group_name=example_resource_group.name,
            location="West US",
            version="12.0",
            administrator_login="4dm1n157r470r",
            administrator_login_password="4-v3ry-53cr37-p455w0rd",
            tags={
                "environment": "production",
            })
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_database = azure.sql.Database("exampleDatabase",
            resource_group_name=example_resource_group.name,
            location="West US",
            server_name=example_sql_server.name,
            extended_auditing_policy=azure.sql.DatabaseExtendedAuditingPolicyArgs(
                storage_endpoint=example_account.primary_blob_endpoint,
                storage_account_access_key=example_account.primary_access_key,
                storage_account_access_key_is_secondary=True,
                retention_in_days=6,
            ),
            tags={
                "environment": "production",
            })
        ```

        ## Import

        SQL Databases can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:sql/database:Database database1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/databases/database1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collation: The name of the collation. Applies only if `create_mode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_mode: Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
        :param pulumi.Input[str] edition: The edition of the database to be created. Applies only if `create_mode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] elastic_pool_name: The name of the elastic database pool.
        :param pulumi.Input[pulumi.InputType['DatabaseExtendedAuditingPolicyArgs']] extended_auditing_policy: A `extended_auditing_policy` block as defined below.
        :param pulumi.Input[pulumi.InputType['DatabaseImportArgs']] import_: A Database Import block as documented below. `create_mode` must be set to `Default`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] max_size_bytes: The maximum size that the database can grow to. Applies only if `create_mode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] name: The name of the database.
        :param pulumi.Input[bool] read_scale: Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
        :param pulumi.Input[str] requested_service_objective_id: A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
               .
        :param pulumi.Input[str] requested_service_objective_name: The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
        :param pulumi.Input[str] restore_point_in_time: The point in time for the restore. Only applies if `create_mode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
        :param pulumi.Input[str] server_name: The name of the SQL Server on which to create the database.
        :param pulumi.Input[str] source_database_deletion_date: The deletion date time of the source database. Only applies to deleted databases where `create_mode` is `PointInTimeRestore`.
        :param pulumi.Input[str] source_database_id: The URI of the source database if `create_mode` value is not `Default`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['DatabaseThreatDetectionPolicyArgs']] threat_detection_policy: Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        :param pulumi.Input[bool] zone_redundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
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

            __props__['collation'] = collation
            __props__['create_mode'] = create_mode
            __props__['edition'] = edition
            __props__['elastic_pool_name'] = elastic_pool_name
            if extended_auditing_policy is not None:
                warnings.warn("""the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.""", DeprecationWarning)
                pulumi.log.warn("extended_auditing_policy is deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.")
            __props__['extended_auditing_policy'] = extended_auditing_policy
            __props__['import_'] = import_
            __props__['location'] = location
            __props__['max_size_bytes'] = max_size_bytes
            __props__['max_size_gb'] = max_size_gb
            __props__['name'] = name
            __props__['read_scale'] = read_scale
            __props__['requested_service_objective_id'] = requested_service_objective_id
            __props__['requested_service_objective_name'] = requested_service_objective_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['restore_point_in_time'] = restore_point_in_time
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['source_database_deletion_date'] = source_database_deletion_date
            __props__['source_database_id'] = source_database_id
            __props__['tags'] = tags
            __props__['threat_detection_policy'] = threat_detection_policy
            __props__['zone_redundant'] = zone_redundant
            __props__['creation_date'] = None
            __props__['default_secondary_location'] = None
            __props__['encryption'] = None
        super(Database, __self__).__init__(
            'azure:sql/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            collation: Optional[pulumi.Input[str]] = None,
            create_mode: Optional[pulumi.Input[str]] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            default_secondary_location: Optional[pulumi.Input[str]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            elastic_pool_name: Optional[pulumi.Input[str]] = None,
            encryption: Optional[pulumi.Input[str]] = None,
            extended_auditing_policy: Optional[pulumi.Input[pulumi.InputType['DatabaseExtendedAuditingPolicyArgs']]] = None,
            import_: Optional[pulumi.Input[pulumi.InputType['DatabaseImportArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            max_size_bytes: Optional[pulumi.Input[str]] = None,
            max_size_gb: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            read_scale: Optional[pulumi.Input[bool]] = None,
            requested_service_objective_id: Optional[pulumi.Input[str]] = None,
            requested_service_objective_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            restore_point_in_time: Optional[pulumi.Input[str]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            source_database_deletion_date: Optional[pulumi.Input[str]] = None,
            source_database_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            threat_detection_policy: Optional[pulumi.Input[pulumi.InputType['DatabaseThreatDetectionPolicyArgs']]] = None,
            zone_redundant: Optional[pulumi.Input[bool]] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collation: The name of the collation. Applies only if `create_mode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_mode: Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
        :param pulumi.Input[str] creation_date: The creation date of the SQL Database.
        :param pulumi.Input[str] default_secondary_location: The default secondary location of the SQL Database.
        :param pulumi.Input[str] edition: The edition of the database to be created. Applies only if `create_mode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] elastic_pool_name: The name of the elastic database pool.
        :param pulumi.Input[pulumi.InputType['DatabaseExtendedAuditingPolicyArgs']] extended_auditing_policy: A `extended_auditing_policy` block as defined below.
        :param pulumi.Input[pulumi.InputType['DatabaseImportArgs']] import_: A Database Import block as documented below. `create_mode` must be set to `Default`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] max_size_bytes: The maximum size that the database can grow to. Applies only if `create_mode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] name: The name of the database.
        :param pulumi.Input[bool] read_scale: Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
        :param pulumi.Input[str] requested_service_objective_id: A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
               .
        :param pulumi.Input[str] requested_service_objective_name: The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
        :param pulumi.Input[str] restore_point_in_time: The point in time for the restore. Only applies if `create_mode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
        :param pulumi.Input[str] server_name: The name of the SQL Server on which to create the database.
        :param pulumi.Input[str] source_database_deletion_date: The deletion date time of the source database. Only applies to deleted databases where `create_mode` is `PointInTimeRestore`.
        :param pulumi.Input[str] source_database_id: The URI of the source database if `create_mode` value is not `Default`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['DatabaseThreatDetectionPolicyArgs']] threat_detection_policy: Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        :param pulumi.Input[bool] zone_redundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["collation"] = collation
        __props__["create_mode"] = create_mode
        __props__["creation_date"] = creation_date
        __props__["default_secondary_location"] = default_secondary_location
        __props__["edition"] = edition
        __props__["elastic_pool_name"] = elastic_pool_name
        __props__["encryption"] = encryption
        __props__["extended_auditing_policy"] = extended_auditing_policy
        __props__["import_"] = import_
        __props__["location"] = location
        __props__["max_size_bytes"] = max_size_bytes
        __props__["max_size_gb"] = max_size_gb
        __props__["name"] = name
        __props__["read_scale"] = read_scale
        __props__["requested_service_objective_id"] = requested_service_objective_id
        __props__["requested_service_objective_name"] = requested_service_objective_name
        __props__["resource_group_name"] = resource_group_name
        __props__["restore_point_in_time"] = restore_point_in_time
        __props__["server_name"] = server_name
        __props__["source_database_deletion_date"] = source_database_deletion_date
        __props__["source_database_id"] = source_database_id
        __props__["tags"] = tags
        __props__["threat_detection_policy"] = threat_detection_policy
        __props__["zone_redundant"] = zone_redundant
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def collation(self) -> pulumi.Output[str]:
        """
        The name of the collation. Applies only if `create_mode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
        """
        return pulumi.get(self, "create_mode")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        The creation date of the SQL Database.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="defaultSecondaryLocation")
    def default_secondary_location(self) -> pulumi.Output[str]:
        """
        The default secondary location of the SQL Database.
        """
        return pulumi.get(self, "default_secondary_location")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        The edition of the database to be created. Applies only if `create_mode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="elasticPoolName")
    def elastic_pool_name(self) -> pulumi.Output[str]:
        """
        The name of the elastic database pool.
        """
        return pulumi.get(self, "elastic_pool_name")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[str]:
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="extendedAuditingPolicy")
    def extended_auditing_policy(self) -> pulumi.Output['outputs.DatabaseExtendedAuditingPolicy']:
        """
        A `extended_auditing_policy` block as defined below.
        """
        return pulumi.get(self, "extended_auditing_policy")

    @property
    @pulumi.getter(name="import")
    def import_(self) -> pulumi.Output[Optional['outputs.DatabaseImport']]:
        """
        A Database Import block as documented below. `create_mode` must be set to `Default`.
        """
        return pulumi.get(self, "import_")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxSizeBytes")
    def max_size_bytes(self) -> pulumi.Output[str]:
        """
        The maximum size that the database can grow to. Applies only if `create_mode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        """
        return pulumi.get(self, "max_size_bytes")

    @property
    @pulumi.getter(name="maxSizeGb")
    def max_size_gb(self) -> pulumi.Output[str]:
        return pulumi.get(self, "max_size_gb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readScale")
    def read_scale(self) -> pulumi.Output[Optional[bool]]:
        """
        Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
        """
        return pulumi.get(self, "read_scale")

    @property
    @pulumi.getter(name="requestedServiceObjectiveId")
    def requested_service_objective_id(self) -> pulumi.Output[str]:
        """
        A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
        .
        """
        return pulumi.get(self, "requested_service_objective_id")

    @property
    @pulumi.getter(name="requestedServiceObjectiveName")
    def requested_service_objective_name(self) -> pulumi.Output[str]:
        """
        The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
        """
        return pulumi.get(self, "requested_service_objective_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="restorePointInTime")
    def restore_point_in_time(self) -> pulumi.Output[str]:
        """
        The point in time for the restore. Only applies if `create_mode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
        """
        return pulumi.get(self, "restore_point_in_time")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        The name of the SQL Server on which to create the database.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="sourceDatabaseDeletionDate")
    def source_database_deletion_date(self) -> pulumi.Output[str]:
        """
        The deletion date time of the source database. Only applies to deleted databases where `create_mode` is `PointInTimeRestore`.
        """
        return pulumi.get(self, "source_database_deletion_date")

    @property
    @pulumi.getter(name="sourceDatabaseId")
    def source_database_id(self) -> pulumi.Output[str]:
        """
        The URI of the source database if `create_mode` value is not `Default`.
        """
        return pulumi.get(self, "source_database_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="threatDetectionPolicy")
    def threat_detection_policy(self) -> pulumi.Output['outputs.DatabaseThreatDetectionPolicy']:
        """
        Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        """
        return pulumi.get(self, "threat_detection_policy")

    @property
    @pulumi.getter(name="zoneRedundant")
    def zone_redundant(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        """
        return pulumi.get(self, "zone_redundant")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

