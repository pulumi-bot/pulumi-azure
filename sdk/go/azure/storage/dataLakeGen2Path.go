// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Data Lake Gen2 Path in a File System within an Azure Storage Account.
//
// > **NOTE:** This Resource requires using Azure Active Directory to connect to Azure Storage, which in turn requires the `Storage` specific roles - which are not granted by default.
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("StorageV2"),
// 			IsHnsEnabled:           pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewDataLakeGen2Path(ctx, "exampleDataLakeGen2Path", &storage.DataLakeGen2PathArgs{
// 			Path:             pulumi.String("example"),
// 			FilesystemName:   exampleDataLakeGen2Filesystem.Name,
// 			StorageAccountId: exampleAccount.ID(),
// 			Resource:         pulumi.String("directory"),
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
// Data Lake Gen2 Paths can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/dataLakeGen2Path:DataLakeGen2Path example https://account1.dfs.core.windows.net/fileSystem1/path
// ```
type DataLakeGen2Path struct {
	pulumi.CustomResourceState

	// One or more `ace` blocks as defined below to specify the entries for the ACL for the path.
	Aces DataLakeGen2PathAceArrayOutput `pulumi:"aces"`
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	FilesystemName pulumi.StringOutput `pulumi:"filesystemName"`
	// Specifies the Object ID of the Azure Active Directory Group to make the owning group.
	Group pulumi.StringOutput `pulumi:"group"`
	// Specifies the Object ID of the Azure Active Directory User to make the owning user.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path pulumi.StringOutput `pulumi:"path"`
	// Specifies the type for path to create. Currently only `directory` is supported.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
}

// NewDataLakeGen2Path registers a new resource with the given unique name, arguments, and options.
func NewDataLakeGen2Path(ctx *pulumi.Context,
	name string, args *DataLakeGen2PathArgs, opts ...pulumi.ResourceOption) (*DataLakeGen2Path, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilesystemName == nil {
		return nil, errors.New("invalid value for required argument 'FilesystemName'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource DataLakeGen2Path
	err := ctx.RegisterResource("azure:storage/dataLakeGen2Path:DataLakeGen2Path", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLakeGen2Path gets an existing DataLakeGen2Path resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLakeGen2Path(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLakeGen2PathState, opts ...pulumi.ResourceOption) (*DataLakeGen2Path, error) {
	var resource DataLakeGen2Path
	err := ctx.ReadResource("azure:storage/dataLakeGen2Path:DataLakeGen2Path", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLakeGen2Path resources.
type dataLakeGen2PathState struct {
	// One or more `ace` blocks as defined below to specify the entries for the ACL for the path.
	Aces []DataLakeGen2PathAce `pulumi:"aces"`
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	FilesystemName *string `pulumi:"filesystemName"`
	// Specifies the Object ID of the Azure Active Directory Group to make the owning group.
	Group *string `pulumi:"group"`
	// Specifies the Object ID of the Azure Active Directory User to make the owning user.
	Owner *string `pulumi:"owner"`
	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path *string `pulumi:"path"`
	// Specifies the type for path to create. Currently only `directory` is supported.
	Resource *string `pulumi:"resource"`
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type DataLakeGen2PathState struct {
	// One or more `ace` blocks as defined below to specify the entries for the ACL for the path.
	Aces DataLakeGen2PathAceArrayInput
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	FilesystemName pulumi.StringPtrInput
	// Specifies the Object ID of the Azure Active Directory Group to make the owning group.
	Group pulumi.StringPtrInput
	// Specifies the Object ID of the Azure Active Directory User to make the owning user.
	Owner pulumi.StringPtrInput
	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path pulumi.StringPtrInput
	// Specifies the type for path to create. Currently only `directory` is supported.
	Resource pulumi.StringPtrInput
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
}

func (DataLakeGen2PathState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeGen2PathState)(nil)).Elem()
}

type dataLakeGen2PathArgs struct {
	// One or more `ace` blocks as defined below to specify the entries for the ACL for the path.
	Aces []DataLakeGen2PathAce `pulumi:"aces"`
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	FilesystemName string `pulumi:"filesystemName"`
	// Specifies the Object ID of the Azure Active Directory Group to make the owning group.
	Group *string `pulumi:"group"`
	// Specifies the Object ID of the Azure Active Directory User to make the owning user.
	Owner *string `pulumi:"owner"`
	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path string `pulumi:"path"`
	// Specifies the type for path to create. Currently only `directory` is supported.
	Resource string `pulumi:"resource"`
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a DataLakeGen2Path resource.
type DataLakeGen2PathArgs struct {
	// One or more `ace` blocks as defined below to specify the entries for the ACL for the path.
	Aces DataLakeGen2PathAceArrayInput
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	FilesystemName pulumi.StringInput
	// Specifies the Object ID of the Azure Active Directory Group to make the owning group.
	Group pulumi.StringPtrInput
	// Specifies the Object ID of the Azure Active Directory User to make the owning user.
	Owner pulumi.StringPtrInput
	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path pulumi.StringInput
	// Specifies the type for path to create. Currently only `directory` is supported.
	Resource pulumi.StringInput
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringInput
}

func (DataLakeGen2PathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeGen2PathArgs)(nil)).Elem()
}

type DataLakeGen2PathInput interface {
	pulumi.Input

	ToDataLakeGen2PathOutput() DataLakeGen2PathOutput
	ToDataLakeGen2PathOutputWithContext(ctx context.Context) DataLakeGen2PathOutput
}

func (*DataLakeGen2Path) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeGen2Path)(nil))
}

func (i *DataLakeGen2Path) ToDataLakeGen2PathOutput() DataLakeGen2PathOutput {
	return i.ToDataLakeGen2PathOutputWithContext(context.Background())
}

func (i *DataLakeGen2Path) ToDataLakeGen2PathOutputWithContext(ctx context.Context) DataLakeGen2PathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeGen2PathOutput)
}

func (i *DataLakeGen2Path) ToDataLakeGen2PathPtrOutput() DataLakeGen2PathPtrOutput {
	return i.ToDataLakeGen2PathPtrOutputWithContext(context.Background())
}

func (i *DataLakeGen2Path) ToDataLakeGen2PathPtrOutputWithContext(ctx context.Context) DataLakeGen2PathPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeGen2PathPtrOutput)
}

type DataLakeGen2PathPtrInput interface {
	pulumi.Input

	ToDataLakeGen2PathPtrOutput() DataLakeGen2PathPtrOutput
	ToDataLakeGen2PathPtrOutputWithContext(ctx context.Context) DataLakeGen2PathPtrOutput
}

type dataLakeGen2PathPtrType DataLakeGen2PathArgs

func (*dataLakeGen2PathPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeGen2Path)(nil))
}

func (i *dataLakeGen2PathPtrType) ToDataLakeGen2PathPtrOutput() DataLakeGen2PathPtrOutput {
	return i.ToDataLakeGen2PathPtrOutputWithContext(context.Background())
}

func (i *dataLakeGen2PathPtrType) ToDataLakeGen2PathPtrOutputWithContext(ctx context.Context) DataLakeGen2PathPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeGen2PathPtrOutput)
}

// DataLakeGen2PathArrayInput is an input type that accepts DataLakeGen2PathArray and DataLakeGen2PathArrayOutput values.
// You can construct a concrete instance of `DataLakeGen2PathArrayInput` via:
//
//          DataLakeGen2PathArray{ DataLakeGen2PathArgs{...} }
type DataLakeGen2PathArrayInput interface {
	pulumi.Input

	ToDataLakeGen2PathArrayOutput() DataLakeGen2PathArrayOutput
	ToDataLakeGen2PathArrayOutputWithContext(context.Context) DataLakeGen2PathArrayOutput
}

type DataLakeGen2PathArray []DataLakeGen2PathInput

func (DataLakeGen2PathArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DataLakeGen2Path)(nil))
}

func (i DataLakeGen2PathArray) ToDataLakeGen2PathArrayOutput() DataLakeGen2PathArrayOutput {
	return i.ToDataLakeGen2PathArrayOutputWithContext(context.Background())
}

func (i DataLakeGen2PathArray) ToDataLakeGen2PathArrayOutputWithContext(ctx context.Context) DataLakeGen2PathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeGen2PathArrayOutput)
}

// DataLakeGen2PathMapInput is an input type that accepts DataLakeGen2PathMap and DataLakeGen2PathMapOutput values.
// You can construct a concrete instance of `DataLakeGen2PathMapInput` via:
//
//          DataLakeGen2PathMap{ "key": DataLakeGen2PathArgs{...} }
type DataLakeGen2PathMapInput interface {
	pulumi.Input

	ToDataLakeGen2PathMapOutput() DataLakeGen2PathMapOutput
	ToDataLakeGen2PathMapOutputWithContext(context.Context) DataLakeGen2PathMapOutput
}

type DataLakeGen2PathMap map[string]DataLakeGen2PathInput

func (DataLakeGen2PathMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DataLakeGen2Path)(nil))
}

func (i DataLakeGen2PathMap) ToDataLakeGen2PathMapOutput() DataLakeGen2PathMapOutput {
	return i.ToDataLakeGen2PathMapOutputWithContext(context.Background())
}

func (i DataLakeGen2PathMap) ToDataLakeGen2PathMapOutputWithContext(ctx context.Context) DataLakeGen2PathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeGen2PathMapOutput)
}

type DataLakeGen2PathOutput struct {
	*pulumi.OutputState
}

func (DataLakeGen2PathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeGen2Path)(nil))
}

func (o DataLakeGen2PathOutput) ToDataLakeGen2PathOutput() DataLakeGen2PathOutput {
	return o
}

func (o DataLakeGen2PathOutput) ToDataLakeGen2PathOutputWithContext(ctx context.Context) DataLakeGen2PathOutput {
	return o
}

func (o DataLakeGen2PathOutput) ToDataLakeGen2PathPtrOutput() DataLakeGen2PathPtrOutput {
	return o.ToDataLakeGen2PathPtrOutputWithContext(context.Background())
}

func (o DataLakeGen2PathOutput) ToDataLakeGen2PathPtrOutputWithContext(ctx context.Context) DataLakeGen2PathPtrOutput {
	return o.ApplyT(func(v DataLakeGen2Path) *DataLakeGen2Path {
		return &v
	}).(DataLakeGen2PathPtrOutput)
}

type DataLakeGen2PathPtrOutput struct {
	*pulumi.OutputState
}

func (DataLakeGen2PathPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeGen2Path)(nil))
}

func (o DataLakeGen2PathPtrOutput) ToDataLakeGen2PathPtrOutput() DataLakeGen2PathPtrOutput {
	return o
}

func (o DataLakeGen2PathPtrOutput) ToDataLakeGen2PathPtrOutputWithContext(ctx context.Context) DataLakeGen2PathPtrOutput {
	return o
}

type DataLakeGen2PathArrayOutput struct{ *pulumi.OutputState }

func (DataLakeGen2PathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataLakeGen2Path)(nil))
}

func (o DataLakeGen2PathArrayOutput) ToDataLakeGen2PathArrayOutput() DataLakeGen2PathArrayOutput {
	return o
}

func (o DataLakeGen2PathArrayOutput) ToDataLakeGen2PathArrayOutputWithContext(ctx context.Context) DataLakeGen2PathArrayOutput {
	return o
}

func (o DataLakeGen2PathArrayOutput) Index(i pulumi.IntInput) DataLakeGen2PathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataLakeGen2Path {
		return vs[0].([]DataLakeGen2Path)[vs[1].(int)]
	}).(DataLakeGen2PathOutput)
}

type DataLakeGen2PathMapOutput struct{ *pulumi.OutputState }

func (DataLakeGen2PathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataLakeGen2Path)(nil))
}

func (o DataLakeGen2PathMapOutput) ToDataLakeGen2PathMapOutput() DataLakeGen2PathMapOutput {
	return o
}

func (o DataLakeGen2PathMapOutput) ToDataLakeGen2PathMapOutputWithContext(ctx context.Context) DataLakeGen2PathMapOutput {
	return o
}

func (o DataLakeGen2PathMapOutput) MapIndex(k pulumi.StringInput) DataLakeGen2PathOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataLakeGen2Path {
		return vs[0].(map[string]DataLakeGen2Path)[vs[1].(string)]
	}).(DataLakeGen2PathOutput)
}

func init() {
	pulumi.RegisterOutputType(DataLakeGen2PathOutput{})
	pulumi.RegisterOutputType(DataLakeGen2PathPtrOutput{})
	pulumi.RegisterOutputType(DataLakeGen2PathArrayOutput{})
	pulumi.RegisterOutputType(DataLakeGen2PathMapOutput{})
}
