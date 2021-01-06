// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to manage an Azure SQL Elastic Pool via the `v3.0` API which allows for `vCore` and `DTU` based configurations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mssql"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("westeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSqlServer, err := sql.NewSqlServer(ctx, "exampleSqlServer", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("4dm1n157r470r"),
// 			AdministratorLoginPassword: pulumi.String("4-v3ry-53cr37-p455w0rd"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewElasticPool(ctx, "exampleElasticPool", &mssql.ElasticPoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ServerName:        exampleSqlServer.Name,
// 			LicenseType:       pulumi.String("LicenseIncluded"),
// 			MaxSizeGb:         pulumi.Float64(756),
// 			Sku: &mssql.ElasticPoolSkuArgs{
// 				Name:     pulumi.String("GP_Gen5"),
// 				Tier:     pulumi.String("GeneralPurpose"),
// 				Family:   pulumi.String("Gen5"),
// 				Capacity: pulumi.Int(4),
// 			},
// 			PerDatabaseSettings: &mssql.ElasticPoolPerDatabaseSettingsArgs{
// 				MinCapacity: pulumi.Float64(0.25),
// 				MaxCapacity: pulumi.Float64(4),
// 			},
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
// SQL Elastic Pool can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mssql/elasticPool:ElasticPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/myelasticpoolname
// ```
type ElasticPool struct {
	pulumi.CustomResourceState

	// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes pulumi.IntOutput `pulumi:"maxSizeBytes"`
	// The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
	MaxSizeGb pulumi.Float64Output `pulumi:"maxSizeGb"`
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `perDatabaseSettings` block as defined below.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsOutput `pulumi:"perDatabaseSettings"`
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// A `sku` block as defined below.
	Sku ElasticPoolSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PerDatabaseSettings == nil {
		return nil, errors.New("invalid value for required argument 'PerDatabaseSettings'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource ElasticPool
	err := ctx.RegisterResource("azure:mssql/elasticPool:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticPool gets an existing ElasticPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure:mssql/elasticPool:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticPool resources.
type elasticPoolState struct {
	// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	// The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
	MaxSizeGb *float64 `pulumi:"maxSizeGb"`
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `perDatabaseSettings` block as defined below.
	PerDatabaseSettings *ElasticPoolPerDatabaseSettings `pulumi:"perDatabaseSettings"`
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// A `sku` block as defined below.
	Sku *ElasticPoolSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type ElasticPoolState struct {
	// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes pulumi.IntPtrInput
	// The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
	MaxSizeGb pulumi.Float64PtrInput
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `perDatabaseSettings` block as defined below.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsPtrInput
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// A `sku` block as defined below.
	Sku ElasticPoolSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	// The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
	MaxSizeGb *float64 `pulumi:"maxSizeGb"`
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `perDatabaseSettings` block as defined below.
	PerDatabaseSettings ElasticPoolPerDatabaseSettings `pulumi:"perDatabaseSettings"`
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// A `sku` block as defined below.
	Sku ElasticPoolSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes pulumi.IntPtrInput
	// The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
	MaxSizeGb pulumi.Float64PtrInput
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `perDatabaseSettings` block as defined below.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsInput
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName pulumi.StringInput
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// A `sku` block as defined below.
	Sku ElasticPoolSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ElasticPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolArgs)(nil)).Elem()
}

type ElasticPoolInput interface {
	pulumi.Input

	ToElasticPoolOutput() ElasticPoolOutput
	ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput
}

func (*ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

func (i *ElasticPool) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return i.ToElasticPoolPtrOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPtrOutput)
}

type ElasticPoolPtrInput interface {
	pulumi.Input

	ToElasticPoolPtrOutput() ElasticPoolPtrOutput
	ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput
}

type elasticPoolPtrType ElasticPoolArgs

func (*elasticPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil))
}

func (i *elasticPoolPtrType) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return i.ToElasticPoolPtrOutputWithContext(context.Background())
}

func (i *elasticPoolPtrType) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput).ToElasticPoolPtrOutput()
}

// ElasticPoolArrayInput is an input type that accepts ElasticPoolArray and ElasticPoolArrayOutput values.
// You can construct a concrete instance of `ElasticPoolArrayInput` via:
//
//          ElasticPoolArray{ ElasticPoolArgs{...} }
type ElasticPoolArrayInput interface {
	pulumi.Input

	ToElasticPoolArrayOutput() ElasticPoolArrayOutput
	ToElasticPoolArrayOutputWithContext(context.Context) ElasticPoolArrayOutput
}

type ElasticPoolArray []ElasticPoolInput

func (ElasticPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ElasticPool)(nil))
}

func (i ElasticPoolArray) ToElasticPoolArrayOutput() ElasticPoolArrayOutput {
	return i.ToElasticPoolArrayOutputWithContext(context.Background())
}

func (i ElasticPoolArray) ToElasticPoolArrayOutputWithContext(ctx context.Context) ElasticPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolArrayOutput)
}

// ElasticPoolMapInput is an input type that accepts ElasticPoolMap and ElasticPoolMapOutput values.
// You can construct a concrete instance of `ElasticPoolMapInput` via:
//
//          ElasticPoolMap{ "key": ElasticPoolArgs{...} }
type ElasticPoolMapInput interface {
	pulumi.Input

	ToElasticPoolMapOutput() ElasticPoolMapOutput
	ToElasticPoolMapOutputWithContext(context.Context) ElasticPoolMapOutput
}

type ElasticPoolMap map[string]ElasticPoolInput

func (ElasticPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ElasticPool)(nil))
}

func (i ElasticPoolMap) ToElasticPoolMapOutput() ElasticPoolMapOutput {
	return i.ToElasticPoolMapOutputWithContext(context.Background())
}

func (i ElasticPoolMap) ToElasticPoolMapOutputWithContext(ctx context.Context) ElasticPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolMapOutput)
}

type ElasticPoolOutput struct {
	*pulumi.OutputState
}

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return o.ToElasticPoolPtrOutputWithContext(context.Background())
}

func (o ElasticPoolOutput) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return o.ApplyT(func(v ElasticPool) *ElasticPool {
		return &v
	}).(ElasticPoolPtrOutput)
}

type ElasticPoolPtrOutput struct {
	*pulumi.OutputState
}

func (ElasticPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil))
}

func (o ElasticPoolPtrOutput) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return o
}

func (o ElasticPoolPtrOutput) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return o
}

type ElasticPoolArrayOutput struct{ *pulumi.OutputState }

func (ElasticPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ElasticPool)(nil))
}

func (o ElasticPoolArrayOutput) ToElasticPoolArrayOutput() ElasticPoolArrayOutput {
	return o
}

func (o ElasticPoolArrayOutput) ToElasticPoolArrayOutputWithContext(ctx context.Context) ElasticPoolArrayOutput {
	return o
}

func (o ElasticPoolArrayOutput) Index(i pulumi.IntInput) ElasticPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ElasticPool {
		return vs[0].([]ElasticPool)[vs[1].(int)]
	}).(ElasticPoolOutput)
}

type ElasticPoolMapOutput struct{ *pulumi.OutputState }

func (ElasticPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ElasticPool)(nil))
}

func (o ElasticPoolMapOutput) ToElasticPoolMapOutput() ElasticPoolMapOutput {
	return o
}

func (o ElasticPoolMapOutput) ToElasticPoolMapOutputWithContext(ctx context.Context) ElasticPoolMapOutput {
	return o
}

func (o ElasticPoolMapOutput) MapIndex(k pulumi.StringInput) ElasticPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ElasticPool {
		return vs[0].(map[string]ElasticPool)[vs[1].(string)]
	}).(ElasticPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticPoolOutput{})
	pulumi.RegisterOutputType(ElasticPoolPtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolArrayOutput{})
	pulumi.RegisterOutputType(ElasticPoolMapOutput{})
}
