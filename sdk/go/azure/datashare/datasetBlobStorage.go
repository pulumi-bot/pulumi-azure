// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Data Share Blob Storage Dataset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datashare"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi-azuread/sdk/v2/go/azuread"
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
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("RAGRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  pulumi.String(exampleStorage / accountAccount.Name),
// 			ContainerAccessType: pulumi.String("container"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              pulumi.String(exampleStorage / accountAccount.Id),
// 			RoleDefinitionName: pulumi.String("Storage Blob Data Reader"),
// 			PrincipalId: exampleServicePrincipal.ApplyT(func(exampleServicePrincipal azuread.LookupServicePrincipalResult) (string, error) {
// 				return exampleServicePrincipal.ObjectId, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datashare.NewDatasetBlobStorage(ctx, "exampleDatasetBlobStorage", &datashare.DatasetBlobStorageArgs{
// 			DataShareId:   exampleShare.ID(),
// 			ContainerName: exampleContainer.Name,
// 			StorageAccount: &datashare.DatasetBlobStorageStorageAccountArgs{
// 				Name:              pulumi.String(exampleStorage / accountAccount.Name),
// 				ResourceGroupName: pulumi.String(exampleStorage / accountAccount.ResourceGroupName),
// 				SubscriptionId:    pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			},
// 			FilePath: pulumi.String("myfile.txt"),
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
type DatasetBlobStorage struct {
	pulumi.CustomResourceState

	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	DataShareId pulumi.StringOutput `pulumi:"dataShareId"`
	// The name of the Data Share Dataset.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath pulumi.StringPtrOutput `pulumi:"filePath"`
	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath pulumi.StringPtrOutput `pulumi:"folderPath"`
	// The name which should be used for this Data Share Blob Storage Dataset. Changing this forces a new Data Share Blob Storage Dataset to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `storageAccount` block as defined below.
	StorageAccount DatasetBlobStorageStorageAccountOutput `pulumi:"storageAccount"`
}

// NewDatasetBlobStorage registers a new resource with the given unique name, arguments, and options.
func NewDatasetBlobStorage(ctx *pulumi.Context,
	name string, args *DatasetBlobStorageArgs, opts ...pulumi.ResourceOption) (*DatasetBlobStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DataShareId == nil {
		return nil, errors.New("invalid value for required argument 'DataShareId'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	var resource DatasetBlobStorage
	err := ctx.RegisterResource("azure:datashare/datasetBlobStorage:DatasetBlobStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetBlobStorage gets an existing DatasetBlobStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetBlobStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetBlobStorageState, opts ...pulumi.ResourceOption) (*DatasetBlobStorage, error) {
	var resource DatasetBlobStorage
	err := ctx.ReadResource("azure:datashare/datasetBlobStorage:DatasetBlobStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetBlobStorage resources.
type datasetBlobStorageState struct {
	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ContainerName *string `pulumi:"containerName"`
	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	DataShareId *string `pulumi:"dataShareId"`
	// The name of the Data Share Dataset.
	DisplayName *string `pulumi:"displayName"`
	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath *string `pulumi:"filePath"`
	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath *string `pulumi:"folderPath"`
	// The name which should be used for this Data Share Blob Storage Dataset. Changing this forces a new Data Share Blob Storage Dataset to be created.
	Name *string `pulumi:"name"`
	// A `storageAccount` block as defined below.
	StorageAccount *DatasetBlobStorageStorageAccount `pulumi:"storageAccount"`
}

type DatasetBlobStorageState struct {
	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ContainerName pulumi.StringPtrInput
	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	DataShareId pulumi.StringPtrInput
	// The name of the Data Share Dataset.
	DisplayName pulumi.StringPtrInput
	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath pulumi.StringPtrInput
	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath pulumi.StringPtrInput
	// The name which should be used for this Data Share Blob Storage Dataset. Changing this forces a new Data Share Blob Storage Dataset to be created.
	Name pulumi.StringPtrInput
	// A `storageAccount` block as defined below.
	StorageAccount DatasetBlobStorageStorageAccountPtrInput
}

func (DatasetBlobStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetBlobStorageState)(nil)).Elem()
}

type datasetBlobStorageArgs struct {
	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ContainerName string `pulumi:"containerName"`
	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	DataShareId string `pulumi:"dataShareId"`
	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath *string `pulumi:"filePath"`
	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath *string `pulumi:"folderPath"`
	// The name which should be used for this Data Share Blob Storage Dataset. Changing this forces a new Data Share Blob Storage Dataset to be created.
	Name *string `pulumi:"name"`
	// A `storageAccount` block as defined below.
	StorageAccount DatasetBlobStorageStorageAccount `pulumi:"storageAccount"`
}

// The set of arguments for constructing a DatasetBlobStorage resource.
type DatasetBlobStorageArgs struct {
	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ContainerName pulumi.StringInput
	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	DataShareId pulumi.StringInput
	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath pulumi.StringPtrInput
	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath pulumi.StringPtrInput
	// The name which should be used for this Data Share Blob Storage Dataset. Changing this forces a new Data Share Blob Storage Dataset to be created.
	Name pulumi.StringPtrInput
	// A `storageAccount` block as defined below.
	StorageAccount DatasetBlobStorageStorageAccountInput
}

func (DatasetBlobStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetBlobStorageArgs)(nil)).Elem()
}

type DatasetBlobStorageInput interface {
	pulumi.Input

	ToDatasetBlobStorageOutput() DatasetBlobStorageOutput
	ToDatasetBlobStorageOutputWithContext(ctx context.Context) DatasetBlobStorageOutput
}

func (DatasetBlobStorage) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetBlobStorage)(nil)).Elem()
}

func (i DatasetBlobStorage) ToDatasetBlobStorageOutput() DatasetBlobStorageOutput {
	return i.ToDatasetBlobStorageOutputWithContext(context.Background())
}

func (i DatasetBlobStorage) ToDatasetBlobStorageOutputWithContext(ctx context.Context) DatasetBlobStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetBlobStorageOutput)
}

type DatasetBlobStorageOutput struct {
	*pulumi.OutputState
}

func (DatasetBlobStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetBlobStorageOutput)(nil)).Elem()
}

func (o DatasetBlobStorageOutput) ToDatasetBlobStorageOutput() DatasetBlobStorageOutput {
	return o
}

func (o DatasetBlobStorageOutput) ToDatasetBlobStorageOutputWithContext(ctx context.Context) DatasetBlobStorageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetBlobStorageOutput{})
}
