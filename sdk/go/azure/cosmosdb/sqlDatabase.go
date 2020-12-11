// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
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
//
// ## Import
//
// Cosmos SQL Database can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/sqlDatabase:SqlDatabase db1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1
// ```
type SqlDatabase struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	AutoscaleSettings SqlDatabaseAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	Throughput        pulumi.IntOutput    `pulumi:"throughput"`
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
	AccountName *string `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	AutoscaleSettings *SqlDatabaseAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	Throughput        *int    `pulumi:"throughput"`
}

type SqlDatabaseState struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	AutoscaleSettings SqlDatabaseAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	Throughput        pulumi.IntPtrInput
}

func (SqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseState)(nil)).Elem()
}

type sqlDatabaseArgs struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	AutoscaleSettings *SqlDatabaseAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Throughput        *int   `pulumi:"throughput"`
}

// The set of arguments for constructing a SqlDatabase resource.
type SqlDatabaseArgs struct {
	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	AutoscaleSettings SqlDatabaseAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	Throughput        pulumi.IntPtrInput
}

func (SqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseArgs)(nil)).Elem()
}

type SqlDatabaseInput interface {
	pulumi.Input

	ToSqlDatabaseOutput() SqlDatabaseOutput
	ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput
}

type SqlDatabasePtrInput interface {
	pulumi.Input

	ToSqlDatabasePtrOutput() SqlDatabasePtrOutput
	ToSqlDatabasePtrOutputWithContext(ctx context.Context) SqlDatabasePtrOutput
}

func (SqlDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabase)(nil)).Elem()
}

func (i SqlDatabase) ToSqlDatabaseOutput() SqlDatabaseOutput {
	return i.ToSqlDatabaseOutputWithContext(context.Background())
}

func (i SqlDatabase) ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseOutput)
}

func (i SqlDatabase) ToSqlDatabasePtrOutput() SqlDatabasePtrOutput {
	return i.ToSqlDatabasePtrOutputWithContext(context.Background())
}

func (i SqlDatabase) ToSqlDatabasePtrOutputWithContext(ctx context.Context) SqlDatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabasePtrOutput)
}

type SqlDatabaseOutput struct {
	*pulumi.OutputState
}

func (SqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseOutput)(nil)).Elem()
}

func (o SqlDatabaseOutput) ToSqlDatabaseOutput() SqlDatabaseOutput {
	return o
}

func (o SqlDatabaseOutput) ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput {
	return o
}

type SqlDatabasePtrOutput struct {
	*pulumi.OutputState
}

func (SqlDatabasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabase)(nil)).Elem()
}

func (o SqlDatabasePtrOutput) ToSqlDatabasePtrOutput() SqlDatabasePtrOutput {
	return o
}

func (o SqlDatabasePtrOutput) ToSqlDatabasePtrOutputWithContext(ctx context.Context) SqlDatabasePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlDatabaseOutput{})
	pulumi.RegisterOutputType(SqlDatabasePtrOutput{})
}
