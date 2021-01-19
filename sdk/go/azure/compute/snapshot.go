// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Disk Snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleManagedDisk, err := compute.NewManagedDisk(ctx, "exampleManagedDisk", &compute.ManagedDiskArgs{
// 			Location:           exampleResourceGroup.Location,
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			StorageAccountType: pulumi.String("Standard_LRS"),
// 			CreateOption:       pulumi.String("Empty"),
// 			DiskSizeGb:         pulumi.Int(10),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSnapshot(ctx, "exampleSnapshot", &compute.SnapshotArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			CreateOption:      pulumi.String("Copy"),
// 			SourceUri:         exampleManagedDisk.ID(),
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
// Snapshots can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/snapshot:Snapshot example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/snapshots/snapshot1
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption pulumi.StringOutput `pulumi:"createOption"`
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb         pulumi.IntOutput                    `pulumi:"diskSizeGb"`
	EncryptionSettings SnapshotEncryptionSettingsPtrOutput `pulumi:"encryptionSettings"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies a reference to an existing snapshot, when `createOption` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId pulumi.StringPtrOutput `pulumi:"sourceResourceId"`
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri pulumi.StringPtrOutput `pulumi:"sourceUri"`
	// Specifies the ID of an storage account. Used with `sourceUri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreateOption == nil {
		return nil, errors.New("invalid value for required argument 'CreateOption'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Snapshot
	err := ctx.RegisterResource("azure:compute/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure:compute/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption *string `pulumi:"createOption"`
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb         *int                        `pulumi:"diskSizeGb"`
	EncryptionSettings *SnapshotEncryptionSettings `pulumi:"encryptionSettings"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies a reference to an existing snapshot, when `createOption` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri *string `pulumi:"sourceUri"`
	// Specifies the ID of an storage account. Used with `sourceUri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SnapshotState struct {
	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption pulumi.StringPtrInput
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb         pulumi.IntPtrInput
	EncryptionSettings SnapshotEncryptionSettingsPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies a reference to an existing snapshot, when `createOption` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId pulumi.StringPtrInput
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri pulumi.StringPtrInput
	// Specifies the ID of an storage account. Used with `sourceUri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption string `pulumi:"createOption"`
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb         *int                        `pulumi:"diskSizeGb"`
	EncryptionSettings *SnapshotEncryptionSettings `pulumi:"encryptionSettings"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies a reference to an existing snapshot, when `createOption` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri *string `pulumi:"sourceUri"`
	// Specifies the ID of an storage account. Used with `sourceUri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption pulumi.StringInput
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb         pulumi.IntPtrInput
	EncryptionSettings SnapshotEncryptionSettingsPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies a reference to an existing snapshot, when `createOption` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId pulumi.StringPtrInput
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri pulumi.StringPtrInput
	// Specifies the ID of an storage account. Used with `sourceUri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

func (i *Snapshot) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return i.ToSnapshotPtrOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPtrOutput)
}

type SnapshotPtrInput interface {
	pulumi.Input

	ToSnapshotPtrOutput() SnapshotPtrOutput
	ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput
}

type snapshotPtrType SnapshotArgs

func (*snapshotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil))
}

func (i *snapshotPtrType) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return i.ToSnapshotPtrOutputWithContext(context.Background())
}

func (i *snapshotPtrType) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPtrOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//          SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Snapshot)(nil))
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//          SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Snapshot)(nil))
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct {
	*pulumi.OutputState
}

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return o.ToSnapshotPtrOutputWithContext(context.Background())
}

func (o SnapshotOutput) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return o.ApplyT(func(v Snapshot) *Snapshot {
		return &v
	}).(SnapshotPtrOutput)
}

type SnapshotPtrOutput struct {
	*pulumi.OutputState
}

func (SnapshotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil))
}

func (o SnapshotPtrOutput) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return o
}

func (o SnapshotPtrOutput) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return o
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Snapshot)(nil))
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Snapshot {
		return vs[0].([]Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Snapshot)(nil))
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Snapshot {
		return vs[0].(map[string]Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotPtrOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
