// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Data Share Data Lake Gen1 Dataset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datalake"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datashare"
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
// 		exampleStore, err := datalake.NewStore(ctx, "exampleStore", &datalake.StoreArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			FirewallState:     pulumi.String("Disabled"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewStoreFile(ctx, "exampleStoreFile", &datalake.StoreFileArgs{
// 			AccountName:    exampleStore.Name,
// 			LocalFilePath:  pulumi.String("./example/myfile.txt"),
// 			RemoteFilePath: pulumi.String("/example/myfile.txt"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              exampleStore.ID(),
// 			RoleDefinitionName: pulumi.String("Owner"),
// 			PrincipalId: exampleServicePrincipal.ApplyT(func(exampleServicePrincipal azuread.LookupServicePrincipalResult) (string, error) {
// 				return exampleServicePrincipal.ObjectId, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datashare.NewDatasetDataLakeGen1(ctx, "exampleDatasetDataLakeGen1", &datashare.DatasetDataLakeGen1Args{
// 			DataShareId:     exampleShare.ID(),
// 			DataLakeStoreId: exampleStore.ID(),
// 			FileName:        pulumi.String("myfile.txt"),
// 			FolderPath:      pulumi.String("example"),
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
type DatasetDataLakeGen1 struct {
	pulumi.CustomResourceState

	// The resource ID of the Data Lake Store to be shared with the receiver.
	DataLakeStoreId pulumi.StringOutput `pulumi:"dataLakeStoreId"`
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	DataShareId pulumi.StringOutput `pulumi:"dataShareId"`
	// The displayed name of the Data Share Dataset.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FileName pulumi.StringPtrOutput `pulumi:"fileName"`
	// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FolderPath pulumi.StringOutput `pulumi:"folderPath"`
	// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDatasetDataLakeGen1 registers a new resource with the given unique name, arguments, and options.
func NewDatasetDataLakeGen1(ctx *pulumi.Context,
	name string, args *DatasetDataLakeGen1Args, opts ...pulumi.ResourceOption) (*DatasetDataLakeGen1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DataLakeStoreId == nil {
		return nil, errors.New("invalid value for required argument 'DataLakeStoreId'")
	}
	if args.DataShareId == nil {
		return nil, errors.New("invalid value for required argument 'DataShareId'")
	}
	if args.FolderPath == nil {
		return nil, errors.New("invalid value for required argument 'FolderPath'")
	}
	var resource DatasetDataLakeGen1
	err := ctx.RegisterResource("azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetDataLakeGen1 gets an existing DatasetDataLakeGen1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetDataLakeGen1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetDataLakeGen1State, opts ...pulumi.ResourceOption) (*DatasetDataLakeGen1, error) {
	var resource DatasetDataLakeGen1
	err := ctx.ReadResource("azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetDataLakeGen1 resources.
type datasetDataLakeGen1State struct {
	// The resource ID of the Data Lake Store to be shared with the receiver.
	DataLakeStoreId *string `pulumi:"dataLakeStoreId"`
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	DataShareId *string `pulumi:"dataShareId"`
	// The displayed name of the Data Share Dataset.
	DisplayName *string `pulumi:"displayName"`
	// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FileName *string `pulumi:"fileName"`
	// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FolderPath *string `pulumi:"folderPath"`
	// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	Name *string `pulumi:"name"`
}

type DatasetDataLakeGen1State struct {
	// The resource ID of the Data Lake Store to be shared with the receiver.
	DataLakeStoreId pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	DataShareId pulumi.StringPtrInput
	// The displayed name of the Data Share Dataset.
	DisplayName pulumi.StringPtrInput
	// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FileName pulumi.StringPtrInput
	// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FolderPath pulumi.StringPtrInput
	// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	Name pulumi.StringPtrInput
}

func (DatasetDataLakeGen1State) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDataLakeGen1State)(nil)).Elem()
}

type datasetDataLakeGen1Args struct {
	// The resource ID of the Data Lake Store to be shared with the receiver.
	DataLakeStoreId string `pulumi:"dataLakeStoreId"`
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	DataShareId string `pulumi:"dataShareId"`
	// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FileName *string `pulumi:"fileName"`
	// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FolderPath string `pulumi:"folderPath"`
	// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a DatasetDataLakeGen1 resource.
type DatasetDataLakeGen1Args struct {
	// The resource ID of the Data Lake Store to be shared with the receiver.
	DataLakeStoreId pulumi.StringInput
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	DataShareId pulumi.StringInput
	// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FileName pulumi.StringPtrInput
	// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	FolderPath pulumi.StringInput
	// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
	Name pulumi.StringPtrInput
}

func (DatasetDataLakeGen1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDataLakeGen1Args)(nil)).Elem()
}
