// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Storage Sync Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleSync, err := storage.NewSync(ctx, "exampleSync", &storage.SyncArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewSyncGroup(ctx, "exampleSyncGroup", &storage.SyncGroupArgs{
// 			StorageSyncId: exampleSync.ID(),
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
// Storage Sync Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/syncGroup:SyncGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/group1
// ```
type SyncGroup struct {
	pulumi.CustomResourceState

	// The name which should be used for this Storage Sync Group. Changing this forces a new Storage Sync Group to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the Storage Sync where this Storage Sync Group is. Changing this forces a new Storage Sync Group to be created.
	StorageSyncId pulumi.StringOutput `pulumi:"storageSyncId"`
}

// NewSyncGroup registers a new resource with the given unique name, arguments, and options.
func NewSyncGroup(ctx *pulumi.Context,
	name string, args *SyncGroupArgs, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageSyncId == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncId'")
	}
	var resource SyncGroup
	err := ctx.RegisterResource("azure:storage/syncGroup:SyncGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncGroup gets an existing SyncGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncGroupState, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	var resource SyncGroup
	err := ctx.ReadResource("azure:storage/syncGroup:SyncGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncGroup resources.
type syncGroupState struct {
	// The name which should be used for this Storage Sync Group. Changing this forces a new Storage Sync Group to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Storage Sync where this Storage Sync Group is. Changing this forces a new Storage Sync Group to be created.
	StorageSyncId *string `pulumi:"storageSyncId"`
}

type SyncGroupState struct {
	// The name which should be used for this Storage Sync Group. Changing this forces a new Storage Sync Group to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Storage Sync where this Storage Sync Group is. Changing this forces a new Storage Sync Group to be created.
	StorageSyncId pulumi.StringPtrInput
}

func (SyncGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupState)(nil)).Elem()
}

type syncGroupArgs struct {
	// The name which should be used for this Storage Sync Group. Changing this forces a new Storage Sync Group to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Storage Sync where this Storage Sync Group is. Changing this forces a new Storage Sync Group to be created.
	StorageSyncId string `pulumi:"storageSyncId"`
}

// The set of arguments for constructing a SyncGroup resource.
type SyncGroupArgs struct {
	// The name which should be used for this Storage Sync Group. Changing this forces a new Storage Sync Group to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Storage Sync where this Storage Sync Group is. Changing this forces a new Storage Sync Group to be created.
	StorageSyncId pulumi.StringInput
}

func (SyncGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupArgs)(nil)).Elem()
}

type SyncGroupInput interface {
	pulumi.Input

	ToSyncGroupOutput() SyncGroupOutput
	ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput
}

func (*SyncGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroup)(nil))
}

func (i *SyncGroup) ToSyncGroupOutput() SyncGroupOutput {
	return i.ToSyncGroupOutputWithContext(context.Background())
}

func (i *SyncGroup) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupOutput)
}

func (i *SyncGroup) ToSyncGroupPtrOutput() SyncGroupPtrOutput {
	return i.ToSyncGroupPtrOutputWithContext(context.Background())
}

func (i *SyncGroup) ToSyncGroupPtrOutputWithContext(ctx context.Context) SyncGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupPtrOutput)
}

type SyncGroupPtrInput interface {
	pulumi.Input

	ToSyncGroupPtrOutput() SyncGroupPtrOutput
	ToSyncGroupPtrOutputWithContext(ctx context.Context) SyncGroupPtrOutput
}

type SyncGroupOutput struct {
	*pulumi.OutputState
}

func (SyncGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroup)(nil))
}

func (o SyncGroupOutput) ToSyncGroupOutput() SyncGroupOutput {
	return o
}

func (o SyncGroupOutput) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return o
}

type SyncGroupPtrOutput struct {
	*pulumi.OutputState
}

func (SyncGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil))
}

func (o SyncGroupPtrOutput) ToSyncGroupPtrOutput() SyncGroupPtrOutput {
	return o
}

func (o SyncGroupPtrOutput) ToSyncGroupPtrOutputWithContext(ctx context.Context) SyncGroupPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SyncGroupOutput{})
	pulumi.RegisterOutputType(SyncGroupPtrOutput{})
}
