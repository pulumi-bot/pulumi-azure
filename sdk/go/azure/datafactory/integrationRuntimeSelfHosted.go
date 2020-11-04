// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Data Factory Self-hosted Integration Runtime.
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
// 			Location: pulumi.String("eastus"),
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
// 		_, err = datafactory.NewIntegrationRuntimeSelfHosted(ctx, "exampleIntegrationRuntimeSelfHosted", &datafactory.IntegrationRuntimeSelfHostedArgs{
// 			ResourceGroupName: pulumi.String("example"),
// 			DataFactoryName:   pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IntegrationRuntimeSelfHosted struct {
	pulumi.CustomResourceState

	// The primary integration runtime authentication key.
	AuthKey1 pulumi.StringOutput `pulumi:"authKey1"`
	// The secondary integration runtime authentication key.
	AuthKey2 pulumi.StringOutput `pulumi:"authKey2"`
	// Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name which should be used for this Data Factory. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `rbacAuthorization` block as defined below.
	RbacAuthorizations IntegrationRuntimeSelfHostedRbacAuthorizationArrayOutput `pulumi:"rbacAuthorizations"`
	// The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIntegrationRuntimeSelfHosted registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRuntimeSelfHosted(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeSelfHostedArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntimeSelfHosted, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationRuntimeSelfHosted
	err := ctx.RegisterResource("azure:datafactory/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRuntimeSelfHosted gets an existing IntegrationRuntimeSelfHosted resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRuntimeSelfHosted(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeSelfHostedState, opts ...pulumi.ResourceOption) (*IntegrationRuntimeSelfHosted, error) {
	var resource IntegrationRuntimeSelfHosted
	err := ctx.ReadResource("azure:datafactory/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRuntimeSelfHosted resources.
type integrationRuntimeSelfHostedState struct {
	// The primary integration runtime authentication key.
	AuthKey1 *string `pulumi:"authKey1"`
	// The secondary integration runtime authentication key.
	AuthKey2 *string `pulumi:"authKey2"`
	// Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// The name which should be used for this Data Factory. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	Name *string `pulumi:"name"`
	// A `rbacAuthorization` block as defined below.
	RbacAuthorizations []IntegrationRuntimeSelfHostedRbacAuthorization `pulumi:"rbacAuthorizations"`
	// The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IntegrationRuntimeSelfHostedState struct {
	// The primary integration runtime authentication key.
	AuthKey1 pulumi.StringPtrInput
	// The secondary integration runtime authentication key.
	AuthKey2 pulumi.StringPtrInput
	// Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	DataFactoryName pulumi.StringPtrInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// The name which should be used for this Data Factory. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	Name pulumi.StringPtrInput
	// A `rbacAuthorization` block as defined below.
	RbacAuthorizations IntegrationRuntimeSelfHostedRbacAuthorizationArrayInput
	// The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IntegrationRuntimeSelfHostedState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeSelfHostedState)(nil)).Elem()
}

type integrationRuntimeSelfHostedArgs struct {
	// Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// The name which should be used for this Data Factory. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	Name *string `pulumi:"name"`
	// A `rbacAuthorization` block as defined below.
	RbacAuthorizations []IntegrationRuntimeSelfHostedRbacAuthorization `pulumi:"rbacAuthorizations"`
	// The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IntegrationRuntimeSelfHosted resource.
type IntegrationRuntimeSelfHostedArgs struct {
	// Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	DataFactoryName pulumi.StringInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// The name which should be used for this Data Factory. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	Name pulumi.StringPtrInput
	// A `rbacAuthorization` block as defined below.
	RbacAuthorizations IntegrationRuntimeSelfHostedRbacAuthorizationArrayInput
	// The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory Self-hosted Integration Runtime to be created.
	ResourceGroupName pulumi.StringInput
}

func (IntegrationRuntimeSelfHostedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeSelfHostedArgs)(nil)).Elem()
}
