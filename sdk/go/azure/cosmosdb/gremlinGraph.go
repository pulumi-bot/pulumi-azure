// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Gremlin Graph within a Cosmos DB Account.
type GremlinGraph struct {
	pulumi.CustomResourceState

	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings GremlinGraphAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
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
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
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
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings *GremlinGraphAutoscaleSettings `pulumi:"autoscaleSettings"`
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
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []GremlinGraphUniqueKey `pulumi:"uniqueKeys"`
}

type GremlinGraphState struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings GremlinGraphAutoscaleSettingsPtrInput
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
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
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
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings *GremlinGraphAutoscaleSettings `pulumi:"autoscaleSettings"`
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
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []GremlinGraphUniqueKey `pulumi:"uniqueKeys"`
}

// The set of arguments for constructing a GremlinGraph resource.
type GremlinGraphArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings GremlinGraphAutoscaleSettingsPtrInput
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
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayInput
}

func (GremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinGraphArgs)(nil)).Elem()
}
