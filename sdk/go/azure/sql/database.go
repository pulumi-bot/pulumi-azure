// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to manage an Azure SQL Database
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sql"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSqlServer, err := sql.NewSqlServer(ctx, "exampleSqlServer", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   pulumi.String("West US"),
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("4dm1n157r470r"),
// 			AdministratorLoginPassword: pulumi.String("4-v3ry-53cr37-p455w0rd"),
// 			Tags: pulumi.Map{
// 				"environment": pulumi.String("production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewDatabase(ctx, "exampleDatabase", &sql.DatabaseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          pulumi.String("West US"),
// 			ServerName:        exampleSqlServer.Name,
// 			ExtendedAuditingPolicy: &sql.DatabaseExtendedAuditingPolicyArgs{
// 				StorageEndpoint:                    exampleAccount.PrimaryBlobEndpoint,
// 				StorageAccountAccessKey:            exampleAccount.PrimaryAccessKey,
// 				StorageAccountAccessKeyIsSecondary: pulumi.Bool(true),
// 				RetentionInDays:                    pulumi.Int(6),
// 			},
// 			Tags: pulumi.Map{
// 				"environment": pulumi.String("production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Database struct {
	pulumi.CustomResourceState

	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringOutput `pulumi:"collation"`
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// The creation date of the SQL Database.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition pulumi.StringOutput `pulumi:"edition"`
	// The name of the elastic database pool.
	ElasticPoolName pulumi.StringOutput `pulumi:"elasticPoolName"`
	Encryption      pulumi.StringOutput `pulumi:"encryption"`
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy DatabaseExtendedAuditingPolicyPtrOutput `pulumi:"extendedAuditingPolicy"`
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import DatabaseImportPtrOutput `pulumi:"import"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes pulumi.StringOutput `pulumi:"maxSizeBytes"`
	MaxSizeGb    pulumi.StringOutput `pulumi:"maxSizeGb"`
	// The name of the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale pulumi.BoolPtrOutput `pulumi:"readScale"`
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId pulumi.StringOutput `pulumi:"requestedServiceObjectiveId"`
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ``` shell az sql db list-editions -l westus --edition Standard -o table  ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName pulumi.StringOutput `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime pulumi.StringOutput `pulumi:"restorePointInTime"`
	// The name of the SQL Server on which to create the database.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate pulumi.StringOutput `pulumi:"sourceDatabaseDeletionDate"`
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId pulumi.StringOutput `pulumi:"sourceDatabaseId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy DatabaseThreatDetectionPolicyOutput `pulumi:"threatDetectionPolicy"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("azure:sql/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure:sql/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode *string `pulumi:"createMode"`
	// The creation date of the SQL Database.
	CreationDate *string `pulumi:"creationDate"`
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation *string `pulumi:"defaultSecondaryLocation"`
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition *string `pulumi:"edition"`
	// The name of the elastic database pool.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	Encryption      *string `pulumi:"encryption"`
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy *DatabaseExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import *DatabaseImport `pulumi:"import"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes *string `pulumi:"maxSizeBytes"`
	MaxSizeGb    *string `pulumi:"maxSizeGb"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale *bool `pulumi:"readScale"`
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId *string `pulumi:"requestedServiceObjectiveId"`
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ``` shell az sql db list-editions -l westus --edition Standard -o table  ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the SQL Server on which to create the database.
	ServerName *string `pulumi:"serverName"`
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy *DatabaseThreatDetectionPolicy `pulumi:"threatDetectionPolicy"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type DatabaseState struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode pulumi.StringPtrInput
	// The creation date of the SQL Database.
	CreationDate pulumi.StringPtrInput
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation pulumi.StringPtrInput
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition pulumi.StringPtrInput
	// The name of the elastic database pool.
	ElasticPoolName pulumi.StringPtrInput
	Encryption      pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy DatabaseExtendedAuditingPolicyPtrInput
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import DatabaseImportPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes pulumi.StringPtrInput
	MaxSizeGb    pulumi.StringPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale pulumi.BoolPtrInput
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId pulumi.StringPtrInput
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ``` shell az sql db list-editions -l westus --edition Standard -o table  ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName pulumi.StringPtrInput
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName pulumi.StringPtrInput
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime pulumi.StringPtrInput
	// The name of the SQL Server on which to create the database.
	ServerName pulumi.StringPtrInput
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy DatabaseThreatDetectionPolicyPtrInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode *string `pulumi:"createMode"`
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition *string `pulumi:"edition"`
	// The name of the elastic database pool.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy *DatabaseExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import *DatabaseImport `pulumi:"import"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes *string `pulumi:"maxSizeBytes"`
	MaxSizeGb    *string `pulumi:"maxSizeGb"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale *bool `pulumi:"readScale"`
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId *string `pulumi:"requestedServiceObjectiveId"`
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ``` shell az sql db list-editions -l westus --edition Standard -o table  ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the SQL Server on which to create the database.
	ServerName string `pulumi:"serverName"`
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy *DatabaseThreatDetectionPolicy `pulumi:"threatDetectionPolicy"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode pulumi.StringPtrInput
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition pulumi.StringPtrInput
	// The name of the elastic database pool.
	ElasticPoolName pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy DatabaseExtendedAuditingPolicyPtrInput
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import DatabaseImportPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes pulumi.StringPtrInput
	MaxSizeGb    pulumi.StringPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale pulumi.BoolPtrInput
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId pulumi.StringPtrInput
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ``` shell az sql db list-editions -l westus --edition Standard -o table  ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName pulumi.StringPtrInput
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName pulumi.StringInput
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime pulumi.StringPtrInput
	// The name of the SQL Server on which to create the database.
	ServerName pulumi.StringInput
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy DatabaseThreatDetectionPolicyPtrInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}
