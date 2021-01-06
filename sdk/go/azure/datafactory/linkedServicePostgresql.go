// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between PostgreSQL and Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServicePostgresql(ctx, "exampleLinkedServicePostgresql", &datafactory.LinkedServicePostgresqlArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("Host=example;Port=5432;Database=example;UID=example;EncryptionMethod=0;Password=example"),
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
// Data Factory PostgreSQL Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServicePostgresql struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service PostgreSQL.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServicePostgresql registers a new resource with the given unique name, arguments, and options.
func NewLinkedServicePostgresql(ctx *pulumi.Context,
	name string, args *LinkedServicePostgresqlArgs, opts ...pulumi.ResourceOption) (*LinkedServicePostgresql, error) {
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
	var resource LinkedServicePostgresql
	err := ctx.RegisterResource("azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServicePostgresql gets an existing LinkedServicePostgresql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServicePostgresql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServicePostgresqlState, opts ...pulumi.ResourceOption) (*LinkedServicePostgresql, error) {
	var resource LinkedServicePostgresql
	err := ctx.ReadResource("azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServicePostgresql resources.
type linkedServicePostgresqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service PostgreSQL.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServicePostgresqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service PostgreSQL.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServicePostgresqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServicePostgresqlState)(nil)).Elem()
}

type linkedServicePostgresqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service PostgreSQL.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServicePostgresql resource.
type LinkedServicePostgresqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service PostgreSQL.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServicePostgresqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServicePostgresqlArgs)(nil)).Elem()
}

type LinkedServicePostgresqlInput interface {
	pulumi.Input

	ToLinkedServicePostgresqlOutput() LinkedServicePostgresqlOutput
	ToLinkedServicePostgresqlOutputWithContext(ctx context.Context) LinkedServicePostgresqlOutput
}

func (*LinkedServicePostgresql) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServicePostgresql)(nil))
}

func (i *LinkedServicePostgresql) ToLinkedServicePostgresqlOutput() LinkedServicePostgresqlOutput {
	return i.ToLinkedServicePostgresqlOutputWithContext(context.Background())
}

func (i *LinkedServicePostgresql) ToLinkedServicePostgresqlOutputWithContext(ctx context.Context) LinkedServicePostgresqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePostgresqlOutput)
}

func (i *LinkedServicePostgresql) ToLinkedServicePostgresqlPtrOutput() LinkedServicePostgresqlPtrOutput {
	return i.ToLinkedServicePostgresqlPtrOutputWithContext(context.Background())
}

func (i *LinkedServicePostgresql) ToLinkedServicePostgresqlPtrOutputWithContext(ctx context.Context) LinkedServicePostgresqlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePostgresqlPtrOutput)
}

type LinkedServicePostgresqlPtrInput interface {
	pulumi.Input

	ToLinkedServicePostgresqlPtrOutput() LinkedServicePostgresqlPtrOutput
	ToLinkedServicePostgresqlPtrOutputWithContext(ctx context.Context) LinkedServicePostgresqlPtrOutput
}

type linkedServicePostgresqlPtrType LinkedServicePostgresqlArgs

func (*linkedServicePostgresqlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServicePostgresql)(nil))
}

func (i *linkedServicePostgresqlPtrType) ToLinkedServicePostgresqlPtrOutput() LinkedServicePostgresqlPtrOutput {
	return i.ToLinkedServicePostgresqlPtrOutputWithContext(context.Background())
}

func (i *linkedServicePostgresqlPtrType) ToLinkedServicePostgresqlPtrOutputWithContext(ctx context.Context) LinkedServicePostgresqlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePostgresqlOutput).ToLinkedServicePostgresqlPtrOutput()
}

// LinkedServicePostgresqlArrayInput is an input type that accepts LinkedServicePostgresqlArray and LinkedServicePostgresqlArrayOutput values.
// You can construct a concrete instance of `LinkedServicePostgresqlArrayInput` via:
//
//          LinkedServicePostgresqlArray{ LinkedServicePostgresqlArgs{...} }
type LinkedServicePostgresqlArrayInput interface {
	pulumi.Input

	ToLinkedServicePostgresqlArrayOutput() LinkedServicePostgresqlArrayOutput
	ToLinkedServicePostgresqlArrayOutputWithContext(context.Context) LinkedServicePostgresqlArrayOutput
}

type LinkedServicePostgresqlArray []LinkedServicePostgresqlInput

func (LinkedServicePostgresqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServicePostgresql)(nil))
}

func (i LinkedServicePostgresqlArray) ToLinkedServicePostgresqlArrayOutput() LinkedServicePostgresqlArrayOutput {
	return i.ToLinkedServicePostgresqlArrayOutputWithContext(context.Background())
}

func (i LinkedServicePostgresqlArray) ToLinkedServicePostgresqlArrayOutputWithContext(ctx context.Context) LinkedServicePostgresqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePostgresqlArrayOutput)
}

// LinkedServicePostgresqlMapInput is an input type that accepts LinkedServicePostgresqlMap and LinkedServicePostgresqlMapOutput values.
// You can construct a concrete instance of `LinkedServicePostgresqlMapInput` via:
//
//          LinkedServicePostgresqlMap{ "key": LinkedServicePostgresqlArgs{...} }
type LinkedServicePostgresqlMapInput interface {
	pulumi.Input

	ToLinkedServicePostgresqlMapOutput() LinkedServicePostgresqlMapOutput
	ToLinkedServicePostgresqlMapOutputWithContext(context.Context) LinkedServicePostgresqlMapOutput
}

type LinkedServicePostgresqlMap map[string]LinkedServicePostgresqlInput

func (LinkedServicePostgresqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServicePostgresql)(nil))
}

func (i LinkedServicePostgresqlMap) ToLinkedServicePostgresqlMapOutput() LinkedServicePostgresqlMapOutput {
	return i.ToLinkedServicePostgresqlMapOutputWithContext(context.Background())
}

func (i LinkedServicePostgresqlMap) ToLinkedServicePostgresqlMapOutputWithContext(ctx context.Context) LinkedServicePostgresqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePostgresqlMapOutput)
}

type LinkedServicePostgresqlOutput struct {
	*pulumi.OutputState
}

func (LinkedServicePostgresqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServicePostgresql)(nil))
}

func (o LinkedServicePostgresqlOutput) ToLinkedServicePostgresqlOutput() LinkedServicePostgresqlOutput {
	return o
}

func (o LinkedServicePostgresqlOutput) ToLinkedServicePostgresqlOutputWithContext(ctx context.Context) LinkedServicePostgresqlOutput {
	return o
}

func (o LinkedServicePostgresqlOutput) ToLinkedServicePostgresqlPtrOutput() LinkedServicePostgresqlPtrOutput {
	return o.ToLinkedServicePostgresqlPtrOutputWithContext(context.Background())
}

func (o LinkedServicePostgresqlOutput) ToLinkedServicePostgresqlPtrOutputWithContext(ctx context.Context) LinkedServicePostgresqlPtrOutput {
	return o.ApplyT(func(v LinkedServicePostgresql) *LinkedServicePostgresql {
		return &v
	}).(LinkedServicePostgresqlPtrOutput)
}

type LinkedServicePostgresqlPtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServicePostgresqlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServicePostgresql)(nil))
}

func (o LinkedServicePostgresqlPtrOutput) ToLinkedServicePostgresqlPtrOutput() LinkedServicePostgresqlPtrOutput {
	return o
}

func (o LinkedServicePostgresqlPtrOutput) ToLinkedServicePostgresqlPtrOutputWithContext(ctx context.Context) LinkedServicePostgresqlPtrOutput {
	return o
}

type LinkedServicePostgresqlArrayOutput struct{ *pulumi.OutputState }

func (LinkedServicePostgresqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServicePostgresql)(nil))
}

func (o LinkedServicePostgresqlArrayOutput) ToLinkedServicePostgresqlArrayOutput() LinkedServicePostgresqlArrayOutput {
	return o
}

func (o LinkedServicePostgresqlArrayOutput) ToLinkedServicePostgresqlArrayOutputWithContext(ctx context.Context) LinkedServicePostgresqlArrayOutput {
	return o
}

func (o LinkedServicePostgresqlArrayOutput) Index(i pulumi.IntInput) LinkedServicePostgresqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServicePostgresql {
		return vs[0].([]LinkedServicePostgresql)[vs[1].(int)]
	}).(LinkedServicePostgresqlOutput)
}

type LinkedServicePostgresqlMapOutput struct{ *pulumi.OutputState }

func (LinkedServicePostgresqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServicePostgresql)(nil))
}

func (o LinkedServicePostgresqlMapOutput) ToLinkedServicePostgresqlMapOutput() LinkedServicePostgresqlMapOutput {
	return o
}

func (o LinkedServicePostgresqlMapOutput) ToLinkedServicePostgresqlMapOutputWithContext(ctx context.Context) LinkedServicePostgresqlMapOutput {
	return o
}

func (o LinkedServicePostgresqlMapOutput) MapIndex(k pulumi.StringInput) LinkedServicePostgresqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServicePostgresql {
		return vs[0].(map[string]LinkedServicePostgresql)[vs[1].(string)]
	}).(LinkedServicePostgresqlOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServicePostgresqlOutput{})
	pulumi.RegisterOutputType(LinkedServicePostgresqlPtrOutput{})
	pulumi.RegisterOutputType(LinkedServicePostgresqlArrayOutput{})
	pulumi.RegisterOutputType(LinkedServicePostgresqlMapOutput{})
}
