// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.
//
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
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
// 		_, err = datafactory.NewLinkedServiceAzureFunction(ctx, "exampleLinkedServiceAzureFunction", &datafactory.LinkedServiceAzureFunctionArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			Url: exampleFunctionApp.ApplyT(func(exampleFunctionApp appservice.LookupFunctionAppResult) (string, error) {
// 				return exampleFunctionApp.DefaultHostname, nil
// 			}).(pulumi.StringOutput),
// 			Key: pulumi.String("foo"),
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
// Data Factory Linked Service's can be imported using the `resource id`, e.g.
type LinkedServiceAzureFunction struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// The system key of the Azure Function.
	Key pulumi.StringOutput `pulumi:"key"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The url of the Azure Function.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewLinkedServiceAzureFunction registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceAzureFunction(ctx *pulumi.Context,
	name string, args *LinkedServiceAzureFunctionArgs, opts ...pulumi.ResourceOption) (*LinkedServiceAzureFunction, error) {
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &LinkedServiceAzureFunctionArgs{}
	}
	var resource LinkedServiceAzureFunction
	err := ctx.RegisterResource("azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceAzureFunction gets an existing LinkedServiceAzureFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceAzureFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceAzureFunctionState, opts ...pulumi.ResourceOption) (*LinkedServiceAzureFunction, error) {
	var resource LinkedServiceAzureFunction
	err := ctx.ReadResource("azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceAzureFunction resources.
type linkedServiceAzureFunctionState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// The system key of the Azure Function.
	Key *string `pulumi:"key"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The url of the Azure Function.
	Url *string `pulumi:"url"`
}

type LinkedServiceAzureFunctionState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// The system key of the Azure Function.
	Key pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The url of the Azure Function.
	Url pulumi.StringPtrInput
}

func (LinkedServiceAzureFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureFunctionState)(nil)).Elem()
}

type linkedServiceAzureFunctionArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// The system key of the Azure Function.
	Key string `pulumi:"key"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The url of the Azure Function.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a LinkedServiceAzureFunction resource.
type LinkedServiceAzureFunctionArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// The system key of the Azure Function.
	Key pulumi.StringInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The url of the Azure Function.
	Url pulumi.StringInput
}

func (LinkedServiceAzureFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureFunctionArgs)(nil)).Elem()
}
