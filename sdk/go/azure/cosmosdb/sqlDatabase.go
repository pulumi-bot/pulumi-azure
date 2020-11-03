// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a SQL Database within a Cosmos DB Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := cosmosdb.LookupAccount(ctx, &cosmosdb.LookupAccountArgs{
// 			Name:              "tfex-cosmosdb-account",
// 			ResourceGroupName: "tfex-cosmosdb-account-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewSqlDatabase(ctx, "exampleSqlDatabase", &cosmosdb.SqlDatabaseArgs{
// 			ResourceGroupName: pulumi.String(exampleAccount.ResourceGroupName),
// 			AccountName:       pulumi.String(exampleAccount.Name),
// 			Throughput:        pulumi.Int(400),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SqlDatabase struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringOutput                   `pulumi:"accountName"`
	AutoscaleSettings SqlDatabaseAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of SQL database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
}

// NewSqlDatabase registers a new resource with the given unique name, arguments, and options.
func NewSqlDatabase(ctx *pulumi.Context,
	name string, args *SqlDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SqlDatabase
	err := ctx.RegisterResource("azure:cosmosdb/sqlDatabase:SqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlDatabase gets an existing SqlDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDatabaseState, opts ...pulumi.ResourceOption) (*SqlDatabase, error) {
	var resource SqlDatabase
	err := ctx.ReadResource("azure:cosmosdb/sqlDatabase:SqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlDatabase resources.
type sqlDatabaseState struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName       *string                       `pulumi:"accountName"`
	AutoscaleSettings *SqlDatabaseAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of SQL database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

type SqlDatabaseState struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringPtrInput
	AutoscaleSettings SqlDatabaseAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of SQL database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (SqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseState)(nil)).Elem()
}

type sqlDatabaseArgs struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName       string                        `pulumi:"accountName"`
	AutoscaleSettings *SqlDatabaseAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of SQL database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

// The set of arguments for constructing a SqlDatabase resource.
type SqlDatabaseArgs struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringInput
	AutoscaleSettings SqlDatabaseAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of SQL database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (SqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseArgs)(nil)).Elem()
}
