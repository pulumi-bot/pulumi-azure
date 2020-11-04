// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a SQL Container within a Cosmos DB Account.
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
// 		_, err := cosmosdb.NewSqlContainer(ctx, "example", &cosmosdb.SqlContainerArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_cosmosdb_account.Example.Resource_group_name),
// 			AccountName:       pulumi.Any(azurerm_cosmosdb_account.Example.Name),
// 			DatabaseName:      pulumi.Any(azurerm_cosmosdb_sql_database.Example.Name),
// 			PartitionKeyPath:  pulumi.String("/definition/id"),
// 			Throughput:        pulumi.Int(400),
// 			IndexingPolicy: &cosmosdb.SqlContainerIndexingPolicyArgs{
// 				IndexingMode: pulumi.String("Consistent"),
// 				IncludedPaths: cosmosdb.SqlContainerIndexingPolicyIncludedPathArray{
// 					&cosmosdb.SqlContainerIndexingPolicyIncludedPathArgs{
// 						Path: pulumi.String("/*"),
// 					},
// 					&cosmosdb.SqlContainerIndexingPolicyIncludedPathArgs{
// 						Path: pulumi.String("/included/?"),
// 					},
// 				},
// 				ExcludedPaths: cosmosdb.SqlContainerIndexingPolicyExcludedPathArray{
// 					&cosmosdb.SqlContainerIndexingPolicyExcludedPathArgs{
// 						Path: pulumi.String("/excluded/?"),
// 					},
// 				},
// 			},
// 			UniqueKeys: cosmosdb.SqlContainerUniqueKeyArray{
// 				&cosmosdb.SqlContainerUniqueKeyArgs{
// 					Paths: pulumi.StringArray{
// 						pulumi.String("/definition/idlong"),
// 						pulumi.String("/definition/idshort"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SqlContainer struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringOutput                    `pulumi:"accountName"`
	AutoscaleSettings SqlContainerAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntOutput `pulumi:"defaultTtl"`
	// An `indexingPolicy` block as defined below.
	IndexingPolicy SqlContainerIndexingPolicyOutput `pulumi:"indexingPolicy"`
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrOutput `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayOutput `pulumi:"uniqueKeys"`
}

// NewSqlContainer registers a new resource with the given unique name, arguments, and options.
func NewSqlContainer(ctx *pulumi.Context,
	name string, args *SqlContainerArgs, opts ...pulumi.ResourceOption) (*SqlContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SqlContainer
	err := ctx.RegisterResource("azure:cosmosdb/sqlContainer:SqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlContainer gets an existing SqlContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlContainerState, opts ...pulumi.ResourceOption) (*SqlContainer, error) {
	var resource SqlContainer
	err := ctx.ReadResource("azure:cosmosdb/sqlContainer:SqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlContainer resources.
type sqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName       *string                        `pulumi:"accountName"`
	AutoscaleSettings *SqlContainerAutoscaleSettings `pulumi:"autoscaleSettings"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// An `indexingPolicy` block as defined below.
	IndexingPolicy *SqlContainerIndexingPolicy `pulumi:"indexingPolicy"`
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []SqlContainerUniqueKey `pulumi:"uniqueKeys"`
}

type SqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringPtrInput
	AutoscaleSettings SqlContainerAutoscaleSettingsPtrInput
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntPtrInput
	// An `indexingPolicy` block as defined below.
	IndexingPolicy SqlContainerIndexingPolicyPtrInput
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayInput
}

func (SqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlContainerState)(nil)).Elem()
}

type sqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName       string                         `pulumi:"accountName"`
	AutoscaleSettings *SqlContainerAutoscaleSettings `pulumi:"autoscaleSettings"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// An `indexingPolicy` block as defined below.
	IndexingPolicy *SqlContainerIndexingPolicy `pulumi:"indexingPolicy"`
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []SqlContainerUniqueKey `pulumi:"uniqueKeys"`
}

// The set of arguments for constructing a SqlContainer resource.
type SqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringInput
	AutoscaleSettings SqlContainerAutoscaleSettingsPtrInput
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntPtrInput
	// An `indexingPolicy` block as defined below.
	IndexingPolicy SqlContainerIndexingPolicyPtrInput
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayInput
}

func (SqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlContainerArgs)(nil)).Elem()
}
