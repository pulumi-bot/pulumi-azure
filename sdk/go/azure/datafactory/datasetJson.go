// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure JSON Dataset inside an Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLinkedServiceWeb, err := datafactory.NewLinkedServiceWeb(ctx, "exampleLinkedServiceWeb", &datafactory.LinkedServiceWebArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryName:    exampleFactory.Name,
// 			AuthenticationType: pulumi.String("Anonymous"),
// 			Url:                pulumi.String("https://www.bing.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetJson(ctx, "exampleDatasetJson", &datafactory.DatasetJsonArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			LinkedServiceName: exampleLinkedServiceWeb.Name,
// 			HttpServerLocation: &datafactory.DatasetJsonHttpServerLocationArgs{
// 				RelativeUrl: pulumi.String("/fizz/buzz/"),
// 				Path:        pulumi.String("foo/bar/"),
// 				Filename:    pulumi.String("foo.txt"),
// 			},
// 			Encoding: pulumi.String("UTF-8"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DatasetJson struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetJsonAzureBlobStorageLocationPtrOutput `pulumi:"azureBlobStorageLocation"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encoding format for the file.
	Encoding pulumi.StringPtrOutput `pulumi:"encoding"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetJsonHttpServerLocationPtrOutput `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetJsonSchemaColumnArrayOutput `pulumi:"schemaColumns"`
}

// NewDatasetJson registers a new resource with the given unique name, arguments, and options.
func NewDatasetJson(ctx *pulumi.Context,
	name string, args *DatasetJsonArgs, opts ...pulumi.ResourceOption) (*DatasetJson, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.LinkedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'LinkedServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DatasetJson
	err := ctx.RegisterResource("azure:datafactory/datasetJson:DatasetJson", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetJson gets an existing DatasetJson resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetJson(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetJsonState, opts ...pulumi.ResourceOption) (*DatasetJson, error) {
	var resource DatasetJson
	err := ctx.ReadResource("azure:datafactory/datasetJson:DatasetJson", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetJson resources.
type datasetJsonState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation *DatasetJsonAzureBlobStorageLocation `pulumi:"azureBlobStorageLocation"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The encoding format for the file.
	Encoding *string `pulumi:"encoding"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation *DatasetJsonHttpServerLocation `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetJsonSchemaColumn `pulumi:"schemaColumns"`
}

type DatasetJsonState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetJsonAzureBlobStorageLocationPtrInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The encoding format for the file.
	Encoding pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetJsonHttpServerLocationPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetJsonSchemaColumnArrayInput
}

func (DatasetJsonState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetJsonState)(nil)).Elem()
}

type datasetJsonArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation *DatasetJsonAzureBlobStorageLocation `pulumi:"azureBlobStorageLocation"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The encoding format for the file.
	Encoding *string `pulumi:"encoding"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation *DatasetJsonHttpServerLocation `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetJsonSchemaColumn `pulumi:"schemaColumns"`
}

// The set of arguments for constructing a DatasetJson resource.
type DatasetJsonArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetJsonAzureBlobStorageLocationPtrInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The encoding format for the file.
	Encoding pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetJsonHttpServerLocationPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetJsonSchemaColumnArrayInput
}

func (DatasetJsonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetJsonArgs)(nil)).Elem()
}
