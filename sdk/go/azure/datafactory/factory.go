// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Data Factory (Version 2).
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
// 		_, err = datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Factory struct {
	pulumi.CustomResourceState

	// A `githubConfiguration` block as defined below.
	GithubConfiguration FactoryGithubConfigurationPtrOutput `pulumi:"githubConfiguration"`
	// An `identity` block as defined below.
	Identity FactoryIdentityOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Data Factory.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `vstsConfiguration` block as defined below.
	VstsConfiguration FactoryVstsConfigurationPtrOutput `pulumi:"vstsConfiguration"`
}

// NewFactory registers a new resource with the given unique name, arguments, and options.
func NewFactory(ctx *pulumi.Context,
	name string, args *FactoryArgs, opts ...pulumi.ResourceOption) (*Factory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Factory
	err := ctx.RegisterResource("azure:datafactory/factory:Factory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFactory gets an existing Factory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFactory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FactoryState, opts ...pulumi.ResourceOption) (*Factory, error) {
	var resource Factory
	err := ctx.ReadResource("azure:datafactory/factory:Factory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Factory resources.
type factoryState struct {
	// A `githubConfiguration` block as defined below.
	GithubConfiguration *FactoryGithubConfiguration `pulumi:"githubConfiguration"`
	// An `identity` block as defined below.
	Identity *FactoryIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Factory.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `vstsConfiguration` block as defined below.
	VstsConfiguration *FactoryVstsConfiguration `pulumi:"vstsConfiguration"`
}

type FactoryState struct {
	// A `githubConfiguration` block as defined below.
	GithubConfiguration FactoryGithubConfigurationPtrInput
	// An `identity` block as defined below.
	Identity FactoryIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `vstsConfiguration` block as defined below.
	VstsConfiguration FactoryVstsConfigurationPtrInput
}

func (FactoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryState)(nil)).Elem()
}

type factoryArgs struct {
	// A `githubConfiguration` block as defined below.
	GithubConfiguration *FactoryGithubConfiguration `pulumi:"githubConfiguration"`
	// An `identity` block as defined below.
	Identity *FactoryIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Factory.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `vstsConfiguration` block as defined below.
	VstsConfiguration *FactoryVstsConfiguration `pulumi:"vstsConfiguration"`
}

// The set of arguments for constructing a Factory resource.
type FactoryArgs struct {
	// A `githubConfiguration` block as defined below.
	GithubConfiguration FactoryGithubConfigurationPtrInput
	// An `identity` block as defined below.
	Identity FactoryIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `vstsConfiguration` block as defined below.
	VstsConfiguration FactoryVstsConfigurationPtrInput
}

func (FactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryArgs)(nil)).Elem()
}
