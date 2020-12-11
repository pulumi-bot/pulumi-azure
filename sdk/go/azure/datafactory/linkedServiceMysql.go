// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between MySQL and Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServiceMysql(ctx, "exampleLinkedServiceMysql", &datafactory.LinkedServiceMysqlArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("Server=test;Port=3306;Database=test;User=test;SSLMode=1;UseSystemTrustStore=0;Password=test"),
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
// Data Factory MySql Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceMysql:LinkedServiceMysql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceMysql struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with MySQL.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service MySQL.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceMysql registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceMysql(ctx *pulumi.Context,
	name string, args *LinkedServiceMysqlArgs, opts ...pulumi.ResourceOption) (*LinkedServiceMysql, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceMysql
	err := ctx.RegisterResource("azure:datafactory/linkedServiceMysql:LinkedServiceMysql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceMysql gets an existing LinkedServiceMysql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceMysql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceMysqlState, opts ...pulumi.ResourceOption) (*LinkedServiceMysql, error) {
	var resource LinkedServiceMysql
	err := ctx.ReadResource("azure:datafactory/linkedServiceMysql:LinkedServiceMysql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceMysql resources.
type linkedServiceMysqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with MySQL.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service MySQL.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceMysqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with MySQL.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service MySQL.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceMysqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceMysqlState)(nil)).Elem()
}

type linkedServiceMysqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with MySQL.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service MySQL.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceMysql resource.
type LinkedServiceMysqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with MySQL.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service MySQL.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceMysqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceMysqlArgs)(nil)).Elem()
}

type LinkedServiceMysqlInput interface {
	pulumi.Input

	ToLinkedServiceMysqlOutput() LinkedServiceMysqlOutput
	ToLinkedServiceMysqlOutputWithContext(ctx context.Context) LinkedServiceMysqlOutput
}

func (*LinkedServiceMysql) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceMysql)(nil))
}

func (i *LinkedServiceMysql) ToLinkedServiceMysqlOutput() LinkedServiceMysqlOutput {
	return i.ToLinkedServiceMysqlOutputWithContext(context.Background())
}

func (i *LinkedServiceMysql) ToLinkedServiceMysqlOutputWithContext(ctx context.Context) LinkedServiceMysqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceMysqlOutput)
}

type LinkedServiceMysqlOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceMysqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceMysql)(nil))
}

func (o LinkedServiceMysqlOutput) ToLinkedServiceMysqlOutput() LinkedServiceMysqlOutput {
	return o
}

func (o LinkedServiceMysqlOutput) ToLinkedServiceMysqlOutputWithContext(ctx context.Context) LinkedServiceMysqlOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceMysqlOutput{})
}
