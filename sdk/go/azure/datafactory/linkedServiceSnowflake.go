// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between Snowflake and Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceSnowflake(ctx, "exampleLinkedServiceSnowflake", &datafactory.LinkedServiceSnowflakeArgs{
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
// ### With Password In Key Vault
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.String(current.TenantId),
// 			SkuName:           pulumi.String("standard"),
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
// 		exampleLinkedServiceKeyVault, err := datafactory.NewLinkedServiceKeyVault(ctx, "exampleLinkedServiceKeyVault", &datafactory.LinkedServiceKeyVaultArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			KeyVaultId:        exampleKeyVault.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceSnowflake(ctx, "exampleLinkedServiceSnowflake", &datafactory.LinkedServiceSnowflakeArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("jdbc:snowflake://account.region.snowflakecomputing.com/?user=user&db=db&warehouse=wh"),
// 			KeyVaultPassword: &datafactory.LinkedServiceSnowflakeKeyVaultPasswordArgs{
// 				LinkedServiceName: exampleLinkedServiceKeyVault.Name,
// 				SecretName:        pulumi.String("secret"),
// 			},
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
// Data Factory Snowflake Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceSnowflake:LinkedServiceSnowflake example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceSnowflake struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with Snowflake.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Snowflake password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSnowflakeKeyVaultPasswordPtrOutput `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceSnowflake registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSnowflake(ctx *pulumi.Context,
	name string, args *LinkedServiceSnowflakeArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSnowflake, error) {
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
	var resource LinkedServiceSnowflake
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSnowflake:LinkedServiceSnowflake", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSnowflake gets an existing LinkedServiceSnowflake resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSnowflake(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSnowflakeState, opts ...pulumi.ResourceOption) (*LinkedServiceSnowflake, error) {
	var resource LinkedServiceSnowflake
	err := ctx.ReadResource("azure:datafactory/linkedServiceSnowflake:LinkedServiceSnowflake", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSnowflake resources.
type linkedServiceSnowflakeState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with Snowflake.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Snowflake password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSnowflakeKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceSnowflakeState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with Snowflake.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store Snowflake password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSnowflakeKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceSnowflakeState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSnowflakeState)(nil)).Elem()
}

type linkedServiceSnowflakeArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with Snowflake.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Snowflake password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSnowflakeKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceSnowflake resource.
type LinkedServiceSnowflakeArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with Snowflake.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store Snowflake password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSnowflakeKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceSnowflakeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSnowflakeArgs)(nil)).Elem()
}

type LinkedServiceSnowflakeInput interface {
	pulumi.Input

	ToLinkedServiceSnowflakeOutput() LinkedServiceSnowflakeOutput
	ToLinkedServiceSnowflakeOutputWithContext(ctx context.Context) LinkedServiceSnowflakeOutput
}

func (*LinkedServiceSnowflake) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSnowflake)(nil))
}

func (i *LinkedServiceSnowflake) ToLinkedServiceSnowflakeOutput() LinkedServiceSnowflakeOutput {
	return i.ToLinkedServiceSnowflakeOutputWithContext(context.Background())
}

func (i *LinkedServiceSnowflake) ToLinkedServiceSnowflakeOutputWithContext(ctx context.Context) LinkedServiceSnowflakeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSnowflakeOutput)
}

func (i *LinkedServiceSnowflake) ToLinkedServiceSnowflakePtrOutput() LinkedServiceSnowflakePtrOutput {
	return i.ToLinkedServiceSnowflakePtrOutputWithContext(context.Background())
}

func (i *LinkedServiceSnowflake) ToLinkedServiceSnowflakePtrOutputWithContext(ctx context.Context) LinkedServiceSnowflakePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSnowflakePtrOutput)
}

type LinkedServiceSnowflakePtrInput interface {
	pulumi.Input

	ToLinkedServiceSnowflakePtrOutput() LinkedServiceSnowflakePtrOutput
	ToLinkedServiceSnowflakePtrOutputWithContext(ctx context.Context) LinkedServiceSnowflakePtrOutput
}

type linkedServiceSnowflakePtrType LinkedServiceSnowflakeArgs

func (*linkedServiceSnowflakePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSnowflake)(nil))
}

func (i *linkedServiceSnowflakePtrType) ToLinkedServiceSnowflakePtrOutput() LinkedServiceSnowflakePtrOutput {
	return i.ToLinkedServiceSnowflakePtrOutputWithContext(context.Background())
}

func (i *linkedServiceSnowflakePtrType) ToLinkedServiceSnowflakePtrOutputWithContext(ctx context.Context) LinkedServiceSnowflakePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSnowflakePtrOutput)
}

// LinkedServiceSnowflakeArrayInput is an input type that accepts LinkedServiceSnowflakeArray and LinkedServiceSnowflakeArrayOutput values.
// You can construct a concrete instance of `LinkedServiceSnowflakeArrayInput` via:
//
//          LinkedServiceSnowflakeArray{ LinkedServiceSnowflakeArgs{...} }
type LinkedServiceSnowflakeArrayInput interface {
	pulumi.Input

	ToLinkedServiceSnowflakeArrayOutput() LinkedServiceSnowflakeArrayOutput
	ToLinkedServiceSnowflakeArrayOutputWithContext(context.Context) LinkedServiceSnowflakeArrayOutput
}

type LinkedServiceSnowflakeArray []LinkedServiceSnowflakeInput

func (LinkedServiceSnowflakeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkedServiceSnowflake)(nil))
}

func (i LinkedServiceSnowflakeArray) ToLinkedServiceSnowflakeArrayOutput() LinkedServiceSnowflakeArrayOutput {
	return i.ToLinkedServiceSnowflakeArrayOutputWithContext(context.Background())
}

func (i LinkedServiceSnowflakeArray) ToLinkedServiceSnowflakeArrayOutputWithContext(ctx context.Context) LinkedServiceSnowflakeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSnowflakeArrayOutput)
}

// LinkedServiceSnowflakeMapInput is an input type that accepts LinkedServiceSnowflakeMap and LinkedServiceSnowflakeMapOutput values.
// You can construct a concrete instance of `LinkedServiceSnowflakeMapInput` via:
//
//          LinkedServiceSnowflakeMap{ "key": LinkedServiceSnowflakeArgs{...} }
type LinkedServiceSnowflakeMapInput interface {
	pulumi.Input

	ToLinkedServiceSnowflakeMapOutput() LinkedServiceSnowflakeMapOutput
	ToLinkedServiceSnowflakeMapOutputWithContext(context.Context) LinkedServiceSnowflakeMapOutput
}

type LinkedServiceSnowflakeMap map[string]LinkedServiceSnowflakeInput

func (LinkedServiceSnowflakeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkedServiceSnowflake)(nil))
}

func (i LinkedServiceSnowflakeMap) ToLinkedServiceSnowflakeMapOutput() LinkedServiceSnowflakeMapOutput {
	return i.ToLinkedServiceSnowflakeMapOutputWithContext(context.Background())
}

func (i LinkedServiceSnowflakeMap) ToLinkedServiceSnowflakeMapOutputWithContext(ctx context.Context) LinkedServiceSnowflakeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSnowflakeMapOutput)
}

type LinkedServiceSnowflakeOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceSnowflakeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSnowflake)(nil))
}

func (o LinkedServiceSnowflakeOutput) ToLinkedServiceSnowflakeOutput() LinkedServiceSnowflakeOutput {
	return o
}

func (o LinkedServiceSnowflakeOutput) ToLinkedServiceSnowflakeOutputWithContext(ctx context.Context) LinkedServiceSnowflakeOutput {
	return o
}

func (o LinkedServiceSnowflakeOutput) ToLinkedServiceSnowflakePtrOutput() LinkedServiceSnowflakePtrOutput {
	return o.ToLinkedServiceSnowflakePtrOutputWithContext(context.Background())
}

func (o LinkedServiceSnowflakeOutput) ToLinkedServiceSnowflakePtrOutputWithContext(ctx context.Context) LinkedServiceSnowflakePtrOutput {
	return o.ApplyT(func(v LinkedServiceSnowflake) *LinkedServiceSnowflake {
		return &v
	}).(LinkedServiceSnowflakePtrOutput)
}

type LinkedServiceSnowflakePtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceSnowflakePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSnowflake)(nil))
}

func (o LinkedServiceSnowflakePtrOutput) ToLinkedServiceSnowflakePtrOutput() LinkedServiceSnowflakePtrOutput {
	return o
}

func (o LinkedServiceSnowflakePtrOutput) ToLinkedServiceSnowflakePtrOutputWithContext(ctx context.Context) LinkedServiceSnowflakePtrOutput {
	return o
}

type LinkedServiceSnowflakeArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceSnowflakeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceSnowflake)(nil))
}

func (o LinkedServiceSnowflakeArrayOutput) ToLinkedServiceSnowflakeArrayOutput() LinkedServiceSnowflakeArrayOutput {
	return o
}

func (o LinkedServiceSnowflakeArrayOutput) ToLinkedServiceSnowflakeArrayOutputWithContext(ctx context.Context) LinkedServiceSnowflakeArrayOutput {
	return o
}

func (o LinkedServiceSnowflakeArrayOutput) Index(i pulumi.IntInput) LinkedServiceSnowflakeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceSnowflake {
		return vs[0].([]LinkedServiceSnowflake)[vs[1].(int)]
	}).(LinkedServiceSnowflakeOutput)
}

type LinkedServiceSnowflakeMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceSnowflakeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceSnowflake)(nil))
}

func (o LinkedServiceSnowflakeMapOutput) ToLinkedServiceSnowflakeMapOutput() LinkedServiceSnowflakeMapOutput {
	return o
}

func (o LinkedServiceSnowflakeMapOutput) ToLinkedServiceSnowflakeMapOutputWithContext(ctx context.Context) LinkedServiceSnowflakeMapOutput {
	return o
}

func (o LinkedServiceSnowflakeMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceSnowflakeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceSnowflake {
		return vs[0].(map[string]LinkedServiceSnowflake)[vs[1].(string)]
	}).(LinkedServiceSnowflakeOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceSnowflakeOutput{})
	pulumi.RegisterOutputType(LinkedServiceSnowflakePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceSnowflakeArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceSnowflakeMapOutput{})
}
