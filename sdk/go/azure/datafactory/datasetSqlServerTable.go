// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a SQL Server Table Dataset inside a Azure Data Factory.
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
// 		exampleLinkedServiceSqlServer, err := datafactory.NewLinkedServiceSqlServer(ctx, "exampleLinkedServiceSqlServer", &datafactory.LinkedServiceSqlServerArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;Password=test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetSqlServerTable(ctx, "exampleDatasetSqlServerTable", &datafactory.DatasetSqlServerTableArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			LinkedServiceName: exampleLinkedServiceSqlServer.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DatasetSqlServerTable struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset SQL Server Table.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset SQL Server Table.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset SQL Server Table.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset SQL Server Table. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset SQL Server Table.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset SQL Server Table. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetSqlServerTableSchemaColumnArrayOutput `pulumi:"schemaColumns"`
	// The table name of the Data Factory Dataset SQL Server Table.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewDatasetSqlServerTable registers a new resource with the given unique name, arguments, and options.
func NewDatasetSqlServerTable(ctx *pulumi.Context,
	name string, args *DatasetSqlServerTableArgs, opts ...pulumi.ResourceOption) (*DatasetSqlServerTable, error) {
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
	var resource DatasetSqlServerTable
	err := ctx.RegisterResource("azure:datafactory/datasetSqlServerTable:DatasetSqlServerTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetSqlServerTable gets an existing DatasetSqlServerTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetSqlServerTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetSqlServerTableState, opts ...pulumi.ResourceOption) (*DatasetSqlServerTable, error) {
	var resource DatasetSqlServerTable
	err := ctx.ReadResource("azure:datafactory/datasetSqlServerTable:DatasetSqlServerTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetSqlServerTable resources.
type datasetSqlServerTableState struct {
	// A map of additional properties to associate with the Data Factory Dataset SQL Server Table.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset SQL Server Table.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset SQL Server Table.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset SQL Server Table. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset SQL Server Table.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset SQL Server Table. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetSqlServerTableSchemaColumn `pulumi:"schemaColumns"`
	// The table name of the Data Factory Dataset SQL Server Table.
	TableName *string `pulumi:"tableName"`
}

type DatasetSqlServerTableState struct {
	// A map of additional properties to associate with the Data Factory Dataset SQL Server Table.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset SQL Server Table.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset SQL Server Table.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset SQL Server Table. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset SQL Server Table.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset SQL Server Table. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetSqlServerTableSchemaColumnArrayInput
	// The table name of the Data Factory Dataset SQL Server Table.
	TableName pulumi.StringPtrInput
}

func (DatasetSqlServerTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetSqlServerTableState)(nil)).Elem()
}

type datasetSqlServerTableArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset SQL Server Table.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset SQL Server Table.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset SQL Server Table.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset SQL Server Table. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset SQL Server Table.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset SQL Server Table. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetSqlServerTableSchemaColumn `pulumi:"schemaColumns"`
	// The table name of the Data Factory Dataset SQL Server Table.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a DatasetSqlServerTable resource.
type DatasetSqlServerTableArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset SQL Server Table.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset SQL Server Table.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Dataset SQL Server Table.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset SQL Server Table. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset SQL Server Table.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset SQL Server Table. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetSqlServerTableSchemaColumnArrayInput
	// The table name of the Data Factory Dataset SQL Server Table.
	TableName pulumi.StringPtrInput
}

func (DatasetSqlServerTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetSqlServerTableArgs)(nil)).Elem()
}

type DatasetSqlServerTableInput interface {
	pulumi.Input

	ToDatasetSqlServerTableOutput() DatasetSqlServerTableOutput
	ToDatasetSqlServerTableOutputWithContext(ctx context.Context) DatasetSqlServerTableOutput
}

func (DatasetSqlServerTable) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetSqlServerTable)(nil)).Elem()
}

func (i DatasetSqlServerTable) ToDatasetSqlServerTableOutput() DatasetSqlServerTableOutput {
	return i.ToDatasetSqlServerTableOutputWithContext(context.Background())
}

func (i DatasetSqlServerTable) ToDatasetSqlServerTableOutputWithContext(ctx context.Context) DatasetSqlServerTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetSqlServerTableOutput)
}

type DatasetSqlServerTableOutput struct {
	*pulumi.OutputState
}

func (DatasetSqlServerTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetSqlServerTableOutput)(nil)).Elem()
}

func (o DatasetSqlServerTableOutput) ToDatasetSqlServerTableOutput() DatasetSqlServerTableOutput {
	return o
}

func (o DatasetSqlServerTableOutput) ToDatasetSqlServerTableOutputWithContext(ctx context.Context) DatasetSqlServerTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetSqlServerTableOutput{})
}
