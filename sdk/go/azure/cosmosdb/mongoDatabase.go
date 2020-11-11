// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Mongo Database within a Cosmos DB Account.
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
// 		_, err = cosmosdb.NewMongoDatabase(ctx, "exampleMongoDatabase", &cosmosdb.MongoDatabaseArgs{
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
type MongoDatabase struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringOutput                     `pulumi:"accountName"`
	AutoscaleSettings MongoDatabaseAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB Mongo Database. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
}

// NewMongoDatabase registers a new resource with the given unique name, arguments, and options.
func NewMongoDatabase(ctx *pulumi.Context,
	name string, args *MongoDatabaseArgs, opts ...pulumi.ResourceOption) (*MongoDatabase, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MongoDatabaseArgs{}
	}
	var resource MongoDatabase
	err := ctx.RegisterResource("azure:cosmosdb/mongoDatabase:MongoDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoDatabase gets an existing MongoDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDatabaseState, opts ...pulumi.ResourceOption) (*MongoDatabase, error) {
	var resource MongoDatabase
	err := ctx.ReadResource("azure:cosmosdb/mongoDatabase:MongoDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoDatabase resources.
type mongoDatabaseState struct {
	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	AccountName       *string                         `pulumi:"accountName"`
	AutoscaleSettings *MongoDatabaseAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB Mongo Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

type MongoDatabaseState struct {
	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringPtrInput
	AutoscaleSettings MongoDatabaseAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB Mongo Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (MongoDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDatabaseState)(nil)).Elem()
}

type mongoDatabaseArgs struct {
	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	AccountName       string                          `pulumi:"accountName"`
	AutoscaleSettings *MongoDatabaseAutoscaleSettings `pulumi:"autoscaleSettings"`
	// Specifies the name of the Cosmos DB Mongo Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

// The set of arguments for constructing a MongoDatabase resource.
type MongoDatabaseArgs struct {
	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringInput
	AutoscaleSettings MongoDatabaseAutoscaleSettingsPtrInput
	// Specifies the name of the Cosmos DB Mongo Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (MongoDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDatabaseArgs)(nil)).Elem()
}

type MongoDatabaseInput interface {
	pulumi.Input

	ToMongoDatabaseOutput() MongoDatabaseOutput
	ToMongoDatabaseOutputWithContext(ctx context.Context) MongoDatabaseOutput
}

func (MongoDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDatabase)(nil)).Elem()
}

func (i MongoDatabase) ToMongoDatabaseOutput() MongoDatabaseOutput {
	return i.ToMongoDatabaseOutputWithContext(context.Background())
}

func (i MongoDatabase) ToMongoDatabaseOutputWithContext(ctx context.Context) MongoDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDatabaseOutput)
}

type MongoDatabaseOutput struct {
	*pulumi.OutputState
}

func (MongoDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDatabaseOutput)(nil)).Elem()
}

func (o MongoDatabaseOutput) ToMongoDatabaseOutput() MongoDatabaseOutput {
	return o
}

func (o MongoDatabaseOutput) ToMongoDatabaseOutputWithContext(ctx context.Context) MongoDatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MongoDatabaseOutput{})
}
