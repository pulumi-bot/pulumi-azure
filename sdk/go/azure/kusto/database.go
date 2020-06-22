// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Database
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/kusto"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		cluster, err := kusto.NewCluster(ctx, "cluster", &kusto.ClusterArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Standard_D13_v2"),
// 				Capacity: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewDatabase(ctx, "database", &kusto.DatabaseArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			ClusterName:       cluster.Name,
// 			HotCachePeriod:    pulumi.String("P7D"),
// 			SoftDeletePeriod:  pulumi.String("P31D"),
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

	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod pulumi.StringPtrOutput `pulumi:"hotCachePeriod"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The size of the database in bytes.
	Size pulumi.Float64Output `pulumi:"size"`
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod pulumi.StringPtrOutput `pulumi:"softDeletePeriod"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("azure:kusto/database:Database", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:kusto/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The size of the database in bytes.
	Size *float64 `pulumi:"size"`
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod *string `pulumi:"softDeletePeriod"`
}

type DatabaseState struct {
	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod pulumi.StringPtrInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The size of the database in bytes.
	Size pulumi.Float64PtrInput
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod *string `pulumi:"softDeletePeriod"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod pulumi.StringPtrInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}
