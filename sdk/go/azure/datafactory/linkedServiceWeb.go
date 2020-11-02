// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between a Web Server and Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServiceWeb(ctx, "exampleLinkedServiceWeb", &datafactory.LinkedServiceWebArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryName:    exampleFactory.Name,
// 			AuthenticationType: pulumi.String("Anonymous"),
// 			Url:                pulumi.String("http://www.bing.com"),
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
type LinkedServiceWeb struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	Password   pulumi.StringPtrOutput `pulumi:"password"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The URL of the web service endpoint (e.g. http://www.microsoft.com).
	Url      pulumi.StringOutput    `pulumi:"url"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewLinkedServiceWeb registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceWeb(ctx *pulumi.Context,
	name string, args *LinkedServiceWebArgs, opts ...pulumi.ResourceOption) (*LinkedServiceWeb, error) {
	if args == nil || args.AuthenticationType == nil {
		return nil, errors.New("missing required argument 'AuthenticationType'")
	}
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &LinkedServiceWebArgs{}
	}
	var resource LinkedServiceWeb
	err := ctx.RegisterResource("azure:datafactory/linkedServiceWeb:LinkedServiceWeb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceWeb gets an existing LinkedServiceWeb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceWeb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceWebState, opts ...pulumi.ResourceOption) (*LinkedServiceWeb, error) {
	var resource LinkedServiceWeb
	err := ctx.ReadResource("azure:datafactory/linkedServiceWeb:LinkedServiceWeb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceWeb resources.
type linkedServiceWebState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType *string `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	Password   *string           `pulumi:"password"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The URL of the web service endpoint (e.g. http://www.microsoft.com).
	Url      *string `pulumi:"url"`
	Username *string `pulumi:"username"`
}

type LinkedServiceWebState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	Password   pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The URL of the web service endpoint (e.g. http://www.microsoft.com).
	Url      pulumi.StringPtrInput
	Username pulumi.StringPtrInput
}

func (LinkedServiceWebState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceWebState)(nil)).Elem()
}

type linkedServiceWebArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType string `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	Password   *string           `pulumi:"password"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The URL of the web service endpoint (e.g. http://www.microsoft.com).
	Url      string  `pulumi:"url"`
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a LinkedServiceWeb resource.
type LinkedServiceWebArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	Password   pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The URL of the web service endpoint (e.g. http://www.microsoft.com).
	Url      pulumi.StringInput
	Username pulumi.StringPtrInput
}

func (LinkedServiceWebArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceWebArgs)(nil)).Elem()
}
