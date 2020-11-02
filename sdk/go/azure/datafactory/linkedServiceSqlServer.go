// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between a SQL Server and Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServiceSqlServer(ctx, "exampleLinkedServiceSqlServer", &datafactory.LinkedServiceSqlServerArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;Password=test"),
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
// Data Factory SQL Server Linked Service's can be imported using the `resource id`, e.g.
type LinkedServiceSqlServer struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceSqlServer registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, args *LinkedServiceSqlServerArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSqlServer, error) {
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LinkedServiceSqlServerArgs{}
	}
	var resource LinkedServiceSqlServer
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSqlServer gets an existing LinkedServiceSqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSqlServerState, opts ...pulumi.ResourceOption) (*LinkedServiceSqlServer, error) {
	var resource LinkedServiceSqlServer
	err := ctx.ReadResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSqlServer resources.
type linkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceSqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSqlServerState)(nil)).Elem()
}

type linkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceSqlServer resource.
type LinkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceSqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSqlServerArgs)(nil)).Elem()
}
