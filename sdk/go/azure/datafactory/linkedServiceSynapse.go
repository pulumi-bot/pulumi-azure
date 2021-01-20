// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between Synapse and Azure Data Factory.
//
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
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
// 		_, err = datafactory.NewLinkedServiceSynapse(ctx, "exampleLinkedServiceSynapse", &datafactory.LinkedServiceSynapseArgs{
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
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
// 		_, err = datafactory.NewLinkedServiceSynapse(ctx, "exampleLinkedServiceSynapse", &datafactory.LinkedServiceSynapseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;"),
// 			KeyVaultPassword: &datafactory.LinkedServiceSynapseKeyVaultPasswordArgs{
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
// Data Factory Synapse Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceSynapse struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with the Synapse.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Synapse.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSynapseKeyVaultPasswordPtrOutput `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceSynapse registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSynapse(ctx *pulumi.Context,
	name string, args *LinkedServiceSynapseArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSynapse, error) {
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
	var resource LinkedServiceSynapse
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSynapse gets an existing LinkedServiceSynapse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSynapse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSynapseState, opts ...pulumi.ResourceOption) (*LinkedServiceSynapse, error) {
	var resource LinkedServiceSynapse
	err := ctx.ReadResource("azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSynapse resources.
type linkedServiceSynapseState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the Synapse.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Synapse.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSynapseKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceSynapseState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the Synapse.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service Synapse.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSynapseKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceSynapseState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSynapseState)(nil)).Elem()
}

type linkedServiceSynapseArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the Synapse.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Synapse.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSynapseKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceSynapse resource.
type LinkedServiceSynapseArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the Synapse.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service Synapse.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSynapseKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceSynapseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSynapseArgs)(nil)).Elem()
}

type LinkedServiceSynapseInput interface {
	pulumi.Input

	ToLinkedServiceSynapseOutput() LinkedServiceSynapseOutput
	ToLinkedServiceSynapseOutputWithContext(ctx context.Context) LinkedServiceSynapseOutput
}

func (*LinkedServiceSynapse) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSynapse)(nil))
}

func (i *LinkedServiceSynapse) ToLinkedServiceSynapseOutput() LinkedServiceSynapseOutput {
	return i.ToLinkedServiceSynapseOutputWithContext(context.Background())
}

func (i *LinkedServiceSynapse) ToLinkedServiceSynapseOutputWithContext(ctx context.Context) LinkedServiceSynapseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSynapseOutput)
}

type LinkedServiceSynapseOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceSynapseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSynapse)(nil))
}

func (o LinkedServiceSynapseOutput) ToLinkedServiceSynapseOutput() LinkedServiceSynapseOutput {
	return o
}

func (o LinkedServiceSynapseOutput) ToLinkedServiceSynapseOutputWithContext(ctx context.Context) LinkedServiceSynapseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceSynapseOutput{})
}
