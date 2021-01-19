// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
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
//
// ## Import
//
// Cosmos Cassandra KeySpace can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/cassandraKeyspace:CassandraKeyspace ks1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1
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

type CassandraKeyspaceInput interface {
	pulumi.Input

	ToCassandraKeyspaceOutput() CassandraKeyspaceOutput
	ToCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraKeyspaceOutput
}

func (*CassandraKeyspace) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspace)(nil))
}

func (i *CassandraKeyspace) ToCassandraKeyspaceOutput() CassandraKeyspaceOutput {
	return i.ToCassandraKeyspaceOutputWithContext(context.Background())
}

func (i *CassandraKeyspace) ToCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraKeyspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceOutput)
}

func (i *CassandraKeyspace) ToCassandraKeyspacePtrOutput() CassandraKeyspacePtrOutput {
	return i.ToCassandraKeyspacePtrOutputWithContext(context.Background())
}

func (i *CassandraKeyspace) ToCassandraKeyspacePtrOutputWithContext(ctx context.Context) CassandraKeyspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspacePtrOutput)
}

type CassandraKeyspacePtrInput interface {
	pulumi.Input

	ToCassandraKeyspacePtrOutput() CassandraKeyspacePtrOutput
	ToCassandraKeyspacePtrOutputWithContext(ctx context.Context) CassandraKeyspacePtrOutput
}

type cassandraKeyspacePtrType CassandraKeyspaceArgs

func (*cassandraKeyspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspace)(nil))
}

func (i *cassandraKeyspacePtrType) ToCassandraKeyspacePtrOutput() CassandraKeyspacePtrOutput {
	return i.ToCassandraKeyspacePtrOutputWithContext(context.Background())
}

func (i *cassandraKeyspacePtrType) ToCassandraKeyspacePtrOutputWithContext(ctx context.Context) CassandraKeyspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspacePtrOutput)
}

// CassandraKeyspaceArrayInput is an input type that accepts CassandraKeyspaceArray and CassandraKeyspaceArrayOutput values.
// You can construct a concrete instance of `CassandraKeyspaceArrayInput` via:
//
//          CassandraKeyspaceArray{ CassandraKeyspaceArgs{...} }
type CassandraKeyspaceArrayInput interface {
	pulumi.Input

	ToCassandraKeyspaceArrayOutput() CassandraKeyspaceArrayOutput
	ToCassandraKeyspaceArrayOutputWithContext(context.Context) CassandraKeyspaceArrayOutput
}

type CassandraKeyspaceArray []CassandraKeyspaceInput

func (CassandraKeyspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CassandraKeyspace)(nil))
}

func (i CassandraKeyspaceArray) ToCassandraKeyspaceArrayOutput() CassandraKeyspaceArrayOutput {
	return i.ToCassandraKeyspaceArrayOutputWithContext(context.Background())
}

func (i CassandraKeyspaceArray) ToCassandraKeyspaceArrayOutputWithContext(ctx context.Context) CassandraKeyspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceArrayOutput)
}

// CassandraKeyspaceMapInput is an input type that accepts CassandraKeyspaceMap and CassandraKeyspaceMapOutput values.
// You can construct a concrete instance of `CassandraKeyspaceMapInput` via:
//
//          CassandraKeyspaceMap{ "key": CassandraKeyspaceArgs{...} }
type CassandraKeyspaceMapInput interface {
	pulumi.Input

	ToCassandraKeyspaceMapOutput() CassandraKeyspaceMapOutput
	ToCassandraKeyspaceMapOutputWithContext(context.Context) CassandraKeyspaceMapOutput
}

type CassandraKeyspaceMap map[string]CassandraKeyspaceInput

func (CassandraKeyspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CassandraKeyspace)(nil))
}

func (i CassandraKeyspaceMap) ToCassandraKeyspaceMapOutput() CassandraKeyspaceMapOutput {
	return i.ToCassandraKeyspaceMapOutputWithContext(context.Background())
}

func (i CassandraKeyspaceMap) ToCassandraKeyspaceMapOutputWithContext(ctx context.Context) CassandraKeyspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceMapOutput)
}

type CassandraKeyspaceOutput struct {
	*pulumi.OutputState
}

func (CassandraKeyspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspace)(nil))
}

func (o CassandraKeyspaceOutput) ToCassandraKeyspaceOutput() CassandraKeyspaceOutput {
	return o
}

func (o CassandraKeyspaceOutput) ToCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraKeyspaceOutput {
	return o
}

func (o CassandraKeyspaceOutput) ToCassandraKeyspacePtrOutput() CassandraKeyspacePtrOutput {
	return o.ToCassandraKeyspacePtrOutputWithContext(context.Background())
}

func (o CassandraKeyspaceOutput) ToCassandraKeyspacePtrOutputWithContext(ctx context.Context) CassandraKeyspacePtrOutput {
	return o.ApplyT(func(v CassandraKeyspace) *CassandraKeyspace {
		return &v
	}).(CassandraKeyspacePtrOutput)
}

type CassandraKeyspacePtrOutput struct {
	*pulumi.OutputState
}

func (CassandraKeyspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspace)(nil))
}

func (o CassandraKeyspacePtrOutput) ToCassandraKeyspacePtrOutput() CassandraKeyspacePtrOutput {
	return o
}

func (o CassandraKeyspacePtrOutput) ToCassandraKeyspacePtrOutputWithContext(ctx context.Context) CassandraKeyspacePtrOutput {
	return o
}

type CassandraKeyspaceArrayOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraKeyspace)(nil))
}

func (o CassandraKeyspaceArrayOutput) ToCassandraKeyspaceArrayOutput() CassandraKeyspaceArrayOutput {
	return o
}

func (o CassandraKeyspaceArrayOutput) ToCassandraKeyspaceArrayOutputWithContext(ctx context.Context) CassandraKeyspaceArrayOutput {
	return o
}

func (o CassandraKeyspaceArrayOutput) Index(i pulumi.IntInput) CassandraKeyspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraKeyspace {
		return vs[0].([]CassandraKeyspace)[vs[1].(int)]
	}).(CassandraKeyspaceOutput)
}

type CassandraKeyspaceMapOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CassandraKeyspace)(nil))
}

func (o CassandraKeyspaceMapOutput) ToCassandraKeyspaceMapOutput() CassandraKeyspaceMapOutput {
	return o
}

func (o CassandraKeyspaceMapOutput) ToCassandraKeyspaceMapOutputWithContext(ctx context.Context) CassandraKeyspaceMapOutput {
	return o
}

func (o CassandraKeyspaceMapOutput) MapIndex(k pulumi.StringInput) CassandraKeyspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CassandraKeyspace {
		return vs[0].(map[string]CassandraKeyspace)[vs[1].(string)]
	}).(CassandraKeyspaceOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraKeyspaceOutput{})
	pulumi.RegisterOutputType(CassandraKeyspacePtrOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceArrayOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceMapOutput{})
}
