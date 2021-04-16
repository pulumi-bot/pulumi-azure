// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Table within an Azure Storage Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewTable(ctx, "exampleTable", &storage.TableArgs{
// 			StorageAccountName: exampleAccount.Name,
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
// Table's within a Storage Account can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/table:Table table1 "https://example.table.core.windows.net/Tables('replace-with-table-name')"
// ```
type Table struct {
	pulumi.CustomResourceState

	// One or more `acl` blocks as defined below.
	Acls TableAclArrayOutput `pulumi:"acls"`
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	var resource Table
	err := ctx.RegisterResource("azure:storage/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("azure:storage/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// One or more `acl` blocks as defined below.
	Acls []TableAcl `pulumi:"acls"`
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name *string `pulumi:"name"`
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type TableState struct {
	// One or more `acl` blocks as defined below.
	Acls TableAclArrayInput
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// One or more `acl` blocks as defined below.
	Acls []TableAcl `pulumi:"acls"`
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name *string `pulumi:"name"`
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// One or more `acl` blocks as defined below.
	Acls TableAclArrayInput
	// The name of the storage table. Must be unique within the storage account the table is located.
	Name pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage table.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil))
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

func (i *Table) ToTablePtrOutput() TablePtrOutput {
	return i.ToTablePtrOutputWithContext(context.Background())
}

func (i *Table) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablePtrOutput)
}

type TablePtrInput interface {
	pulumi.Input

	ToTablePtrOutput() TablePtrOutput
	ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput
}

type tablePtrType TableArgs

func (*tablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil))
}

func (i *tablePtrType) ToTablePtrOutput() TablePtrOutput {
	return i.ToTablePtrOutputWithContext(context.Background())
}

func (i *tablePtrType) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablePtrOutput)
}

// TableArrayInput is an input type that accepts TableArray and TableArrayOutput values.
// You can construct a concrete instance of `TableArrayInput` via:
//
//          TableArray{ TableArgs{...} }
type TableArrayInput interface {
	pulumi.Input

	ToTableArrayOutput() TableArrayOutput
	ToTableArrayOutputWithContext(context.Context) TableArrayOutput
}

type TableArray []TableInput

func (TableArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Table)(nil))
}

func (i TableArray) ToTableArrayOutput() TableArrayOutput {
	return i.ToTableArrayOutputWithContext(context.Background())
}

func (i TableArray) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableArrayOutput)
}

// TableMapInput is an input type that accepts TableMap and TableMapOutput values.
// You can construct a concrete instance of `TableMapInput` via:
//
//          TableMap{ "key": TableArgs{...} }
type TableMapInput interface {
	pulumi.Input

	ToTableMapOutput() TableMapOutput
	ToTableMapOutputWithContext(context.Context) TableMapOutput
}

type TableMap map[string]TableInput

func (TableMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Table)(nil))
}

func (i TableMap) ToTableMapOutput() TableMapOutput {
	return i.ToTableMapOutputWithContext(context.Background())
}

func (i TableMap) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableMapOutput)
}

type TableOutput struct {
	*pulumi.OutputState
}

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil))
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ToTablePtrOutput() TablePtrOutput {
	return o.ToTablePtrOutputWithContext(context.Background())
}

func (o TableOutput) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return o.ApplyT(func(v Table) *Table {
		return &v
	}).(TablePtrOutput)
}

type TablePtrOutput struct {
	*pulumi.OutputState
}

func (TablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil))
}

func (o TablePtrOutput) ToTablePtrOutput() TablePtrOutput {
	return o
}

func (o TablePtrOutput) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return o
}

type TableArrayOutput struct{ *pulumi.OutputState }

func (TableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Table)(nil))
}

func (o TableArrayOutput) ToTableArrayOutput() TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return o
}

func (o TableArrayOutput) Index(i pulumi.IntInput) TableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Table {
		return vs[0].([]Table)[vs[1].(int)]
	}).(TableOutput)
}

type TableMapOutput struct{ *pulumi.OutputState }

func (TableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Table)(nil))
}

func (o TableMapOutput) ToTableMapOutput() TableMapOutput {
	return o
}

func (o TableMapOutput) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return o
}

func (o TableMapOutput) MapIndex(k pulumi.StringInput) TableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Table {
		return vs[0].(map[string]Table)[vs[1].(string)]
	}).(TableOutput)
}

func init() {
	pulumi.RegisterOutputType(TableOutput{})
	pulumi.RegisterOutputType(TablePtrOutput{})
	pulumi.RegisterOutputType(TableArrayOutput{})
	pulumi.RegisterOutputType(TableMapOutput{})
}
