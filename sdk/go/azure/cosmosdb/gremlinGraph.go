// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Gremlin Graph within a Cosmos DB Account.
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
// 		exampleAccount, err := cosmosdb.LookupAccount(ctx, "azure:cosmosdb:getAccount", &cosmosdb.LookupAccountArgs{
// 			Name:              "tfex-cosmosdb-account",
// 			ResourceGroupName: "tfex-cosmosdb-account-rg",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleGremlinDatabase, err := cosmosdb.NewGremlinDatabase(ctx, "exampleGremlinDatabase", &cosmosdb.GremlinDatabaseArgs{
// 			ResourceGroupName: pulumi.String(exampleAccount.ResourceGroupName),
// 			AccountName:       pulumi.String(exampleAccount.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewGremlinGraph(ctx, "exampleGremlinGraph", &cosmosdb.GremlinGraphArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_cosmosdb_account.Example.Resource_group_name),
// 			AccountName:       pulumi.Any(azurerm_cosmosdb_account.Example.Name),
// 			DatabaseName:      exampleGremlinDatabase.Name,
// 			PartitionKeyPath:  pulumi.String("/Example"),
// 			Throughput:        pulumi.Int(400),
// 			IndexPolicies: cosmosdb.GremlinGraphIndexPolicyArray{
// 				&cosmosdb.GremlinGraphIndexPolicyArgs{
// 					Automatic:    pulumi.Bool(true),
// 					IndexingMode: pulumi.String("Consistent"),
// 					IncludedPaths: pulumi.StringArray{
// 						pulumi.String("/*"),
// 					},
// 					ExcludedPaths: pulumi.StringArray{
// 						pulumi.String("/\"_etag\"/?"),
// 					},
// 				},
// 			},
// 			ConflictResolutionPolicies: cosmosdb.GremlinGraphConflictResolutionPolicyArray{
// 				&cosmosdb.GremlinGraphConflictResolutionPolicyArgs{
// 					Mode:                   pulumi.String("LastWriterWins"),
// 					ConflictResolutionPath: pulumi.String("/_ts"),
// 				},
// 			},
// 			UniqueKeys: cosmosdb.GremlinGraphUniqueKeyArray{
// 				&cosmosdb.GremlinGraphUniqueKeyArgs{
// 					Paths: pulumi.StringArray{
// 						pulumi.String("/definition/id1"),
// 						pulumi.String("/definition/id2"),
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
type GremlinGraph struct {
	pulumi.CustomResourceState

	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
	ConflictResolutionPolicies GremlinGraphConflictResolutionPolicyArrayOutput `pulumi:"conflictResolutionPolicies"`
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
	IndexPolicies GremlinGraphIndexPolicyArrayOutput `pulumi:"indexPolicies"`
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrOutput `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of the Gremlin database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayOutput `pulumi:"uniqueKeys"`
}

// NewGremlinGraph registers a new resource with the given unique name, arguments, and options.
func NewGremlinGraph(ctx *pulumi.Context,
	name string, args *GremlinGraphArgs, opts ...pulumi.ResourceOption) (*GremlinGraph, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ConflictResolutionPolicies == nil {
		return nil, errors.New("missing required argument 'ConflictResolutionPolicies'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.IndexPolicies == nil {
		return nil, errors.New("missing required argument 'IndexPolicies'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GremlinGraphArgs{}
	}
	var resource GremlinGraph
	err := ctx.RegisterResource("azure:cosmosdb/gremlinGraph:GremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGremlinGraph gets an existing GremlinGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinGraphState, opts ...pulumi.ResourceOption) (*GremlinGraph, error) {
	var resource GremlinGraph
	err := ctx.ReadResource("azure:cosmosdb/gremlinGraph:GremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GremlinGraph resources.
type gremlinGraphState struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
	ConflictResolutionPolicies []GremlinGraphConflictResolutionPolicy `pulumi:"conflictResolutionPolicies"`
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
	IndexPolicies []GremlinGraphIndexPolicy `pulumi:"indexPolicies"`
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of the Gremlin database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []GremlinGraphUniqueKey `pulumi:"uniqueKeys"`
}

type GremlinGraphState struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
	ConflictResolutionPolicies GremlinGraphConflictResolutionPolicyArrayInput
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
	IndexPolicies GremlinGraphIndexPolicyArrayInput
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of the Gremlin database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayInput
}

func (GremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinGraphState)(nil)).Elem()
}

type gremlinGraphArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
	ConflictResolutionPolicies []GremlinGraphConflictResolutionPolicy `pulumi:"conflictResolutionPolicies"`
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
	IndexPolicies []GremlinGraphIndexPolicy `pulumi:"indexPolicies"`
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of the Gremlin database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []GremlinGraphUniqueKey `pulumi:"uniqueKeys"`
}

// The set of arguments for constructing a GremlinGraph resource.
type GremlinGraphArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
	ConflictResolutionPolicies GremlinGraphConflictResolutionPolicyArrayInput
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
	IndexPolicies GremlinGraphIndexPolicyArrayInput
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of the Gremlin database (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayInput
}

func (GremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinGraphArgs)(nil)).Elem()
}
