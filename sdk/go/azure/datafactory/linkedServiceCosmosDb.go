// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.
//
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Import
//
// Data Factory Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceCosmosDb struct {
	pulumi.CustomResourceState

	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint pulumi.StringPtrOutput `pulumi:"accountEndpoint"`
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey pulumi.StringPtrOutput `pulumi:"accountKey"`
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString pulumi.StringPtrOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The name of the database. Required if `connectionString` is unspecified.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceCosmosDb registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceCosmosDb(ctx *pulumi.Context,
	name string, args *LinkedServiceCosmosDbArgs, opts ...pulumi.ResourceOption) (*LinkedServiceCosmosDb, error) {
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LinkedServiceCosmosDbArgs{}
	}
	var resource LinkedServiceCosmosDb
	err := ctx.RegisterResource("azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceCosmosDb gets an existing LinkedServiceCosmosDb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceCosmosDb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceCosmosDbState, opts ...pulumi.ResourceOption) (*LinkedServiceCosmosDb, error) {
	var resource LinkedServiceCosmosDb
	err := ctx.ReadResource("azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceCosmosDb resources.
type linkedServiceCosmosDbState struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint *string `pulumi:"accountEndpoint"`
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey *string `pulumi:"accountKey"`
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The name of the database. Required if `connectionString` is unspecified.
	Database *string `pulumi:"database"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceCosmosDbState struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint pulumi.StringPtrInput
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey pulumi.StringPtrInput
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The name of the database. Required if `connectionString` is unspecified.
	Database pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceCosmosDbState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceCosmosDbState)(nil)).Elem()
}

type linkedServiceCosmosDbArgs struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint *string `pulumi:"accountEndpoint"`
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey *string `pulumi:"accountKey"`
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The name of the database. Required if `connectionString` is unspecified.
	Database *string `pulumi:"database"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceCosmosDb resource.
type LinkedServiceCosmosDbArgs struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint pulumi.StringPtrInput
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey pulumi.StringPtrInput
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The name of the database. Required if `connectionString` is unspecified.
	Database pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceCosmosDbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceCosmosDbArgs)(nil)).Elem()
}

type LinkedServiceCosmosDbInput interface {
	pulumi.Input

	ToLinkedServiceCosmosDbOutput() LinkedServiceCosmosDbOutput
	ToLinkedServiceCosmosDbOutputWithContext(ctx context.Context) LinkedServiceCosmosDbOutput
}

func (LinkedServiceCosmosDb) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceCosmosDb)(nil)).Elem()
}

func (i LinkedServiceCosmosDb) ToLinkedServiceCosmosDbOutput() LinkedServiceCosmosDbOutput {
	return i.ToLinkedServiceCosmosDbOutputWithContext(context.Background())
}

func (i LinkedServiceCosmosDb) ToLinkedServiceCosmosDbOutputWithContext(ctx context.Context) LinkedServiceCosmosDbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceCosmosDbOutput)
}

type LinkedServiceCosmosDbOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceCosmosDbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceCosmosDbOutput)(nil)).Elem()
}

func (o LinkedServiceCosmosDbOutput) ToLinkedServiceCosmosDbOutput() LinkedServiceCosmosDbOutput {
	return o
}

func (o LinkedServiceCosmosDbOutput) ToLinkedServiceCosmosDbOutputWithContext(ctx context.Context) LinkedServiceCosmosDbOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceCosmosDbOutput{})
}
