// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Logic App Integration Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/logicapps"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = logicapps.NewIntegrationAccount(ctx, "exampleIntegrationAccount", &logicapps.IntegrationAccountArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IntegrationAccount struct {
	pulumi.CustomResourceState

	// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Logic App Integration Account.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewIntegrationAccount registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccount(ctx *pulumi.Context,
	name string, args *IntegrationAccountArgs, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil {
		args = &IntegrationAccountArgs{}
	}
	var resource IntegrationAccount
	err := ctx.RegisterResource("azure:logicapps/integrationAccount:IntegrationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccount gets an existing IntegrationAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountState, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	var resource IntegrationAccount
	err := ctx.ReadResource("azure:logicapps/integrationAccount:IntegrationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccount resources.
type integrationAccountState struct {
	// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Logic App Integration Account.
	Tags map[string]string `pulumi:"tags"`
}

type IntegrationAccountState struct {
	// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Logic App Integration Account.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountState)(nil)).Elem()
}

type integrationAccountArgs struct {
	// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Logic App Integration Account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccount resource.
type IntegrationAccountArgs struct {
	// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
	ResourceGroupName pulumi.StringInput
	// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
	SkuName pulumi.StringInput
	// A mapping of tags which should be assigned to the Logic App Integration Account.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountArgs)(nil)).Elem()
}

type IntegrationAccountInput interface {
	pulumi.Input

	ToIntegrationAccountOutput() IntegrationAccountOutput
	ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput
}

func (IntegrationAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccount)(nil)).Elem()
}

func (i IntegrationAccount) ToIntegrationAccountOutput() IntegrationAccountOutput {
	return i.ToIntegrationAccountOutputWithContext(context.Background())
}

func (i IntegrationAccount) ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountOutput)
}

type IntegrationAccountOutput struct {
	*pulumi.OutputState
}

func (IntegrationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountOutput)(nil)).Elem()
}

func (o IntegrationAccountOutput) ToIntegrationAccountOutput() IntegrationAccountOutput {
	return o
}

func (o IntegrationAccountOutput) ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountOutput{})
}
