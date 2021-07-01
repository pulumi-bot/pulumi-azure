// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Data Share Data Lake Gen2 Dataset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datashare"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
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
// 		exampleAccount, err := datashare.NewAccount(ctx, "exampleAccount", &datashare.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Identity: &datashare.AccountIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleShare, err := datashare.NewShare(ctx, "exampleShare", &datashare.ShareArgs{
// 			AccountId: exampleAccount.ID(),
// 			Kind:      pulumi.String("CopyBased"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewAccount(ctx, "exampleStorage_accountAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountKind:            pulumi.String("BlobStorage"),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleStorage / accountAccount.Id,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              exampleStorage / accountAccount.Id,
// 			RoleDefinitionName: pulumi.String("Storage Blob Data Reader"),
// 			PrincipalId: exampleServicePrincipal.ApplyT(func(exampleServicePrincipal GetServicePrincipalResult) (string, error) {
// 				return exampleServicePrincipal.ObjectId, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datashare.NewDatasetDataLakeGen2(ctx, "exampleDatasetDataLakeGen2", &datashare.DatasetDataLakeGen2Args{
// 			ShareId:          exampleShare.ID(),
// 			StorageAccountId: exampleStorage / accountAccount.Id,
// 			FileSystemName:   exampleDataLakeGen2Filesystem.Name,
// 			FilePath:         pulumi.String("myfile.txt"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAssignment,
// 		}))
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
// Data Share Data Lake Gen2 Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
// ```
type DatasetDataLakeGen2 struct {
	pulumi.CustomResourceState

	// The name of the Data Share Dataset.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath pulumi.StringPtrOutput `pulumi:"filePath"`
	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FileSystemName pulumi.StringOutput `pulumi:"fileSystemName"`
	// The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath pulumi.StringPtrOutput `pulumi:"folderPath"`
	// The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	ShareId pulumi.StringOutput `pulumi:"shareId"`
	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
}

// NewDatasetDataLakeGen2 registers a new resource with the given unique name, arguments, and options.
func NewDatasetDataLakeGen2(ctx *pulumi.Context,
	name string, args *DatasetDataLakeGen2Args, opts ...pulumi.ResourceOption) (*DatasetDataLakeGen2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemName == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemName'")
	}
	if args.ShareId == nil {
		return nil, errors.New("invalid value for required argument 'ShareId'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource DatasetDataLakeGen2
	err := ctx.RegisterResource("azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetDataLakeGen2 gets an existing DatasetDataLakeGen2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetDataLakeGen2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetDataLakeGen2State, opts ...pulumi.ResourceOption) (*DatasetDataLakeGen2, error) {
	var resource DatasetDataLakeGen2
	err := ctx.ReadResource("azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetDataLakeGen2 resources.
type datasetDataLakeGen2State struct {
	// The name of the Data Share Dataset.
	DisplayName *string `pulumi:"displayName"`
	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath *string `pulumi:"filePath"`
	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FileSystemName *string `pulumi:"fileSystemName"`
	// The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath *string `pulumi:"folderPath"`
	// The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	ShareId *string `pulumi:"shareId"`
	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type DatasetDataLakeGen2State struct {
	// The name of the Data Share Dataset.
	DisplayName pulumi.StringPtrInput
	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath pulumi.StringPtrInput
	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FileSystemName pulumi.StringPtrInput
	// The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath pulumi.StringPtrInput
	// The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	ShareId pulumi.StringPtrInput
	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	StorageAccountId pulumi.StringPtrInput
}

func (DatasetDataLakeGen2State) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDataLakeGen2State)(nil)).Elem()
}

type datasetDataLakeGen2Args struct {
	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath *string `pulumi:"filePath"`
	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FileSystemName string `pulumi:"fileSystemName"`
	// The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath *string `pulumi:"folderPath"`
	// The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	ShareId string `pulumi:"shareId"`
	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a DatasetDataLakeGen2 resource.
type DatasetDataLakeGen2Args struct {
	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath pulumi.StringPtrInput
	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FileSystemName pulumi.StringInput
	// The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath pulumi.StringPtrInput
	// The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	ShareId pulumi.StringInput
	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	StorageAccountId pulumi.StringInput
}

func (DatasetDataLakeGen2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDataLakeGen2Args)(nil)).Elem()
}

type DatasetDataLakeGen2Input interface {
	pulumi.Input

	ToDatasetDataLakeGen2Output() DatasetDataLakeGen2Output
	ToDatasetDataLakeGen2OutputWithContext(ctx context.Context) DatasetDataLakeGen2Output
}

func (*DatasetDataLakeGen2) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetDataLakeGen2)(nil))
}

func (i *DatasetDataLakeGen2) ToDatasetDataLakeGen2Output() DatasetDataLakeGen2Output {
	return i.ToDatasetDataLakeGen2OutputWithContext(context.Background())
}

func (i *DatasetDataLakeGen2) ToDatasetDataLakeGen2OutputWithContext(ctx context.Context) DatasetDataLakeGen2Output {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDataLakeGen2Output)
}

func (i *DatasetDataLakeGen2) ToDatasetDataLakeGen2PtrOutput() DatasetDataLakeGen2PtrOutput {
	return i.ToDatasetDataLakeGen2PtrOutputWithContext(context.Background())
}

func (i *DatasetDataLakeGen2) ToDatasetDataLakeGen2PtrOutputWithContext(ctx context.Context) DatasetDataLakeGen2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDataLakeGen2PtrOutput)
}

type DatasetDataLakeGen2PtrInput interface {
	pulumi.Input

	ToDatasetDataLakeGen2PtrOutput() DatasetDataLakeGen2PtrOutput
	ToDatasetDataLakeGen2PtrOutputWithContext(ctx context.Context) DatasetDataLakeGen2PtrOutput
}

type datasetDataLakeGen2PtrType DatasetDataLakeGen2Args

func (*datasetDataLakeGen2PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetDataLakeGen2)(nil))
}

func (i *datasetDataLakeGen2PtrType) ToDatasetDataLakeGen2PtrOutput() DatasetDataLakeGen2PtrOutput {
	return i.ToDatasetDataLakeGen2PtrOutputWithContext(context.Background())
}

func (i *datasetDataLakeGen2PtrType) ToDatasetDataLakeGen2PtrOutputWithContext(ctx context.Context) DatasetDataLakeGen2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDataLakeGen2PtrOutput)
}

// DatasetDataLakeGen2ArrayInput is an input type that accepts DatasetDataLakeGen2Array and DatasetDataLakeGen2ArrayOutput values.
// You can construct a concrete instance of `DatasetDataLakeGen2ArrayInput` via:
//
//          DatasetDataLakeGen2Array{ DatasetDataLakeGen2Args{...} }
type DatasetDataLakeGen2ArrayInput interface {
	pulumi.Input

	ToDatasetDataLakeGen2ArrayOutput() DatasetDataLakeGen2ArrayOutput
	ToDatasetDataLakeGen2ArrayOutputWithContext(context.Context) DatasetDataLakeGen2ArrayOutput
}

type DatasetDataLakeGen2Array []DatasetDataLakeGen2Input

func (DatasetDataLakeGen2Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatasetDataLakeGen2)(nil))
}

func (i DatasetDataLakeGen2Array) ToDatasetDataLakeGen2ArrayOutput() DatasetDataLakeGen2ArrayOutput {
	return i.ToDatasetDataLakeGen2ArrayOutputWithContext(context.Background())
}

func (i DatasetDataLakeGen2Array) ToDatasetDataLakeGen2ArrayOutputWithContext(ctx context.Context) DatasetDataLakeGen2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDataLakeGen2ArrayOutput)
}

// DatasetDataLakeGen2MapInput is an input type that accepts DatasetDataLakeGen2Map and DatasetDataLakeGen2MapOutput values.
// You can construct a concrete instance of `DatasetDataLakeGen2MapInput` via:
//
//          DatasetDataLakeGen2Map{ "key": DatasetDataLakeGen2Args{...} }
type DatasetDataLakeGen2MapInput interface {
	pulumi.Input

	ToDatasetDataLakeGen2MapOutput() DatasetDataLakeGen2MapOutput
	ToDatasetDataLakeGen2MapOutputWithContext(context.Context) DatasetDataLakeGen2MapOutput
}

type DatasetDataLakeGen2Map map[string]DatasetDataLakeGen2Input

func (DatasetDataLakeGen2Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatasetDataLakeGen2)(nil))
}

func (i DatasetDataLakeGen2Map) ToDatasetDataLakeGen2MapOutput() DatasetDataLakeGen2MapOutput {
	return i.ToDatasetDataLakeGen2MapOutputWithContext(context.Background())
}

func (i DatasetDataLakeGen2Map) ToDatasetDataLakeGen2MapOutputWithContext(ctx context.Context) DatasetDataLakeGen2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDataLakeGen2MapOutput)
}

type DatasetDataLakeGen2Output struct {
	*pulumi.OutputState
}

func (DatasetDataLakeGen2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetDataLakeGen2)(nil))
}

func (o DatasetDataLakeGen2Output) ToDatasetDataLakeGen2Output() DatasetDataLakeGen2Output {
	return o
}

func (o DatasetDataLakeGen2Output) ToDatasetDataLakeGen2OutputWithContext(ctx context.Context) DatasetDataLakeGen2Output {
	return o
}

func (o DatasetDataLakeGen2Output) ToDatasetDataLakeGen2PtrOutput() DatasetDataLakeGen2PtrOutput {
	return o.ToDatasetDataLakeGen2PtrOutputWithContext(context.Background())
}

func (o DatasetDataLakeGen2Output) ToDatasetDataLakeGen2PtrOutputWithContext(ctx context.Context) DatasetDataLakeGen2PtrOutput {
	return o.ApplyT(func(v DatasetDataLakeGen2) *DatasetDataLakeGen2 {
		return &v
	}).(DatasetDataLakeGen2PtrOutput)
}

type DatasetDataLakeGen2PtrOutput struct {
	*pulumi.OutputState
}

func (DatasetDataLakeGen2PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetDataLakeGen2)(nil))
}

func (o DatasetDataLakeGen2PtrOutput) ToDatasetDataLakeGen2PtrOutput() DatasetDataLakeGen2PtrOutput {
	return o
}

func (o DatasetDataLakeGen2PtrOutput) ToDatasetDataLakeGen2PtrOutputWithContext(ctx context.Context) DatasetDataLakeGen2PtrOutput {
	return o
}

type DatasetDataLakeGen2ArrayOutput struct{ *pulumi.OutputState }

func (DatasetDataLakeGen2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetDataLakeGen2)(nil))
}

func (o DatasetDataLakeGen2ArrayOutput) ToDatasetDataLakeGen2ArrayOutput() DatasetDataLakeGen2ArrayOutput {
	return o
}

func (o DatasetDataLakeGen2ArrayOutput) ToDatasetDataLakeGen2ArrayOutputWithContext(ctx context.Context) DatasetDataLakeGen2ArrayOutput {
	return o
}

func (o DatasetDataLakeGen2ArrayOutput) Index(i pulumi.IntInput) DatasetDataLakeGen2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetDataLakeGen2 {
		return vs[0].([]DatasetDataLakeGen2)[vs[1].(int)]
	}).(DatasetDataLakeGen2Output)
}

type DatasetDataLakeGen2MapOutput struct{ *pulumi.OutputState }

func (DatasetDataLakeGen2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatasetDataLakeGen2)(nil))
}

func (o DatasetDataLakeGen2MapOutput) ToDatasetDataLakeGen2MapOutput() DatasetDataLakeGen2MapOutput {
	return o
}

func (o DatasetDataLakeGen2MapOutput) ToDatasetDataLakeGen2MapOutputWithContext(ctx context.Context) DatasetDataLakeGen2MapOutput {
	return o
}

func (o DatasetDataLakeGen2MapOutput) MapIndex(k pulumi.StringInput) DatasetDataLakeGen2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatasetDataLakeGen2 {
		return vs[0].(map[string]DatasetDataLakeGen2)[vs[1].(string)]
	}).(DatasetDataLakeGen2Output)
}

func init() {
	pulumi.RegisterOutputType(DatasetDataLakeGen2Output{})
	pulumi.RegisterOutputType(DatasetDataLakeGen2PtrOutput{})
	pulumi.RegisterOutputType(DatasetDataLakeGen2ArrayOutput{})
	pulumi.RegisterOutputType(DatasetDataLakeGen2MapOutput{})
}
