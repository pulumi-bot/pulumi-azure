// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Entity within a Table in an Azure Storage Account.
//
// ## Import
//
// Entities within a Table in an Azure Storage Account can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/tableEntity:TableEntity entity1 https://example.table.core.windows.net/table1(PartitionKey='samplepartition',RowKey='samplerow')
// ```
type TableEntity struct {
	pulumi.CustomResourceState

	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity pulumi.StringMapOutput `pulumi:"entity"`
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey pulumi.StringOutput `pulumi:"partitionKey"`
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey pulumi.StringOutput `pulumi:"rowKey"`
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// The name of the storage table in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	TableName pulumi.StringOutput `pulumi:"tableName"`
}

// NewTableEntity registers a new resource with the given unique name, arguments, and options.
func NewTableEntity(ctx *pulumi.Context,
	name string, args *TableEntityArgs, opts ...pulumi.ResourceOption) (*TableEntity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	if args.PartitionKey == nil {
		return nil, errors.New("invalid value for required argument 'PartitionKey'")
	}
	if args.RowKey == nil {
		return nil, errors.New("invalid value for required argument 'RowKey'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	var resource TableEntity
	err := ctx.RegisterResource("azure:storage/tableEntity:TableEntity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableEntity gets an existing TableEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableEntityState, opts ...pulumi.ResourceOption) (*TableEntity, error) {
	var resource TableEntity
	err := ctx.ReadResource("azure:storage/tableEntity:TableEntity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableEntity resources.
type tableEntityState struct {
	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity map[string]string `pulumi:"entity"`
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey *string `pulumi:"partitionKey"`
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey *string `pulumi:"rowKey"`
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The name of the storage table in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	TableName *string `pulumi:"tableName"`
}

type TableEntityState struct {
	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity pulumi.StringMapInput
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey pulumi.StringPtrInput
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
	// The name of the storage table in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	TableName pulumi.StringPtrInput
}

func (TableEntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableEntityState)(nil)).Elem()
}

type tableEntityArgs struct {
	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity map[string]string `pulumi:"entity"`
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey string `pulumi:"partitionKey"`
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey string `pulumi:"rowKey"`
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
	// The name of the storage table in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a TableEntity resource.
type TableEntityArgs struct {
	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity pulumi.StringMapInput
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey pulumi.StringInput
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey pulumi.StringInput
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
	// The name of the storage table in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	TableName pulumi.StringInput
}

func (TableEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableEntityArgs)(nil)).Elem()
}

type TableEntityInput interface {
	pulumi.Input

	ToTableEntityOutput() TableEntityOutput
	ToTableEntityOutputWithContext(ctx context.Context) TableEntityOutput
}

func (*TableEntity) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEntity)(nil))
}

func (i *TableEntity) ToTableEntityOutput() TableEntityOutput {
	return i.ToTableEntityOutputWithContext(context.Background())
}

func (i *TableEntity) ToTableEntityOutputWithContext(ctx context.Context) TableEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEntityOutput)
}

func (i *TableEntity) ToTableEntityPtrOutput() TableEntityPtrOutput {
	return i.ToTableEntityPtrOutputWithContext(context.Background())
}

func (i *TableEntity) ToTableEntityPtrOutputWithContext(ctx context.Context) TableEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEntityPtrOutput)
}

type TableEntityPtrInput interface {
	pulumi.Input

	ToTableEntityPtrOutput() TableEntityPtrOutput
	ToTableEntityPtrOutputWithContext(ctx context.Context) TableEntityPtrOutput
}

type tableEntityPtrType TableEntityArgs

func (*tableEntityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableEntity)(nil))
}

func (i *tableEntityPtrType) ToTableEntityPtrOutput() TableEntityPtrOutput {
	return i.ToTableEntityPtrOutputWithContext(context.Background())
}

func (i *tableEntityPtrType) ToTableEntityPtrOutputWithContext(ctx context.Context) TableEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEntityPtrOutput)
}

// TableEntityArrayInput is an input type that accepts TableEntityArray and TableEntityArrayOutput values.
// You can construct a concrete instance of `TableEntityArrayInput` via:
//
//          TableEntityArray{ TableEntityArgs{...} }
type TableEntityArrayInput interface {
	pulumi.Input

	ToTableEntityArrayOutput() TableEntityArrayOutput
	ToTableEntityArrayOutputWithContext(context.Context) TableEntityArrayOutput
}

type TableEntityArray []TableEntityInput

func (TableEntityArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TableEntity)(nil))
}

func (i TableEntityArray) ToTableEntityArrayOutput() TableEntityArrayOutput {
	return i.ToTableEntityArrayOutputWithContext(context.Background())
}

func (i TableEntityArray) ToTableEntityArrayOutputWithContext(ctx context.Context) TableEntityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEntityArrayOutput)
}

// TableEntityMapInput is an input type that accepts TableEntityMap and TableEntityMapOutput values.
// You can construct a concrete instance of `TableEntityMapInput` via:
//
//          TableEntityMap{ "key": TableEntityArgs{...} }
type TableEntityMapInput interface {
	pulumi.Input

	ToTableEntityMapOutput() TableEntityMapOutput
	ToTableEntityMapOutputWithContext(context.Context) TableEntityMapOutput
}

type TableEntityMap map[string]TableEntityInput

func (TableEntityMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TableEntity)(nil))
}

func (i TableEntityMap) ToTableEntityMapOutput() TableEntityMapOutput {
	return i.ToTableEntityMapOutputWithContext(context.Background())
}

func (i TableEntityMap) ToTableEntityMapOutputWithContext(ctx context.Context) TableEntityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEntityMapOutput)
}

type TableEntityOutput struct {
	*pulumi.OutputState
}

func (TableEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEntity)(nil))
}

func (o TableEntityOutput) ToTableEntityOutput() TableEntityOutput {
	return o
}

func (o TableEntityOutput) ToTableEntityOutputWithContext(ctx context.Context) TableEntityOutput {
	return o
}

func (o TableEntityOutput) ToTableEntityPtrOutput() TableEntityPtrOutput {
	return o.ToTableEntityPtrOutputWithContext(context.Background())
}

func (o TableEntityOutput) ToTableEntityPtrOutputWithContext(ctx context.Context) TableEntityPtrOutput {
	return o.ApplyT(func(v TableEntity) *TableEntity {
		return &v
	}).(TableEntityPtrOutput)
}

type TableEntityPtrOutput struct {
	*pulumi.OutputState
}

func (TableEntityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableEntity)(nil))
}

func (o TableEntityPtrOutput) ToTableEntityPtrOutput() TableEntityPtrOutput {
	return o
}

func (o TableEntityPtrOutput) ToTableEntityPtrOutputWithContext(ctx context.Context) TableEntityPtrOutput {
	return o
}

type TableEntityArrayOutput struct{ *pulumi.OutputState }

func (TableEntityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableEntity)(nil))
}

func (o TableEntityArrayOutput) ToTableEntityArrayOutput() TableEntityArrayOutput {
	return o
}

func (o TableEntityArrayOutput) ToTableEntityArrayOutputWithContext(ctx context.Context) TableEntityArrayOutput {
	return o
}

func (o TableEntityArrayOutput) Index(i pulumi.IntInput) TableEntityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TableEntity {
		return vs[0].([]TableEntity)[vs[1].(int)]
	}).(TableEntityOutput)
}

type TableEntityMapOutput struct{ *pulumi.OutputState }

func (TableEntityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TableEntity)(nil))
}

func (o TableEntityMapOutput) ToTableEntityMapOutput() TableEntityMapOutput {
	return o
}

func (o TableEntityMapOutput) ToTableEntityMapOutputWithContext(ctx context.Context) TableEntityMapOutput {
	return o
}

func (o TableEntityMapOutput) MapIndex(k pulumi.StringInput) TableEntityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TableEntity {
		return vs[0].(map[string]TableEntity)[vs[1].(string)]
	}).(TableEntityOutput)
}

func init() {
	pulumi.RegisterOutputType(TableEntityOutput{})
	pulumi.RegisterOutputType(TableEntityPtrOutput{})
	pulumi.RegisterOutputType(TableEntityArrayOutput{})
	pulumi.RegisterOutputType(TableEntityMapOutput{})
}
