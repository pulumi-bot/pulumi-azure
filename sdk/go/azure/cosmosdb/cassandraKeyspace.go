// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Cassandra KeySpace within a Cosmos DB Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "tflex-cosmosdb-account-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := cosmosdb.NewAccount(ctx, "exampleAccount", &cosmosdb.AccountArgs{
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			OfferType:         pulumi.String("Standard"),
// 			Capabilities: cosmosdb.AccountCapabilityArray{
// 				&cosmosdb.AccountCapabilityArgs{
// 					Name: pulumi.String("EnableCassandra"),
// 				},
// 			},
// 			ConsistencyPolicy: &cosmosdb.AccountConsistencyPolicyArgs{
// 				ConsistencyLevel: pulumi.String("Strong"),
// 			},
// 			GeoLocations: cosmosdb.AccountGeoLocationArray{
// 				&cosmosdb.AccountGeoLocationArgs{
// 					Location:         pulumi.String("West US"),
// 					FailoverPriority: pulumi.Int(0),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewCassandraKeyspace(ctx, "exampleCassandraKeyspace", &cosmosdb.CassandraKeyspaceArgs{
// 			ResourceGroupName: exampleAccount.ResourceGroupName,
// 			AccountName:       exampleAccount.Name,
// 			Throughput:        pulumi.Int(400),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CassandraKeyspace struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	AutoscaleSettings CassandraKeyspaceAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB Cassandra KeySpace. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
}

// NewCassandraKeyspace registers a new resource with the given unique name, arguments, and options.
func NewCassandraKeyspace(ctx *pulumi.Context,
	name string, args *CassandraKeyspaceArgs, opts ...pulumi.ResourceOption) (*CassandraKeyspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CassandraKeyspace
	err := ctx.RegisterResource("azure:cosmosdb/cassandraKeyspace:CassandraKeyspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraKeyspace gets an existing CassandraKeyspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraKeyspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraKeyspaceState, opts ...pulumi.ResourceOption) (*CassandraKeyspace, error) {
	var resource CassandraKeyspace
	err := ctx.ReadResource("azure:cosmosdb/cassandraKeyspace:CassandraKeyspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraKeyspace resources.
type cassandraKeyspaceState struct {
	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	AutoscaleSettings *CassandraKeyspaceAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB Cassandra KeySpace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

type CassandraKeyspaceState struct {
	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	AutoscaleSettings CassandraKeyspaceAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB Cassandra KeySpace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (CassandraKeyspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraKeyspaceState)(nil)).Elem()
}

type cassandraKeyspaceArgs struct {
	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	AutoscaleSettings *CassandraKeyspaceAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB Cassandra KeySpace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

// The set of arguments for constructing a CassandraKeyspace resource.
type CassandraKeyspaceArgs struct {
	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	AutoscaleSettings CassandraKeyspaceAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB Cassandra KeySpace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (CassandraKeyspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraKeyspaceArgs)(nil)).Elem()
}
