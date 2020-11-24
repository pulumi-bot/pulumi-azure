// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securitycenter

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a resources Advanced Threat Protection setting.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/securitycenter"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:               pulumi.Any(azurerm_resource_group.Example.Location),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securitycenter.NewAdvancedThreatProtection(ctx, "exampleAdvancedThreatProtection", &securitycenter.AdvancedThreatProtectionArgs{
// 			TargetResourceId: exampleAccount.ID(),
// 			Enabled:          pulumi.Bool(true),
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
// Advanced Threat Protection can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleaccount/providers/Microsoft.Security/advancedThreatProtectionSettings/default
// ```
type AdvancedThreatProtection struct {
	pulumi.CustomResourceState

	// Should Advanced Threat Protection be enabled on this resource?
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewAdvancedThreatProtection registers a new resource with the given unique name, arguments, and options.
func NewAdvancedThreatProtection(ctx *pulumi.Context,
	name string, args *AdvancedThreatProtectionArgs, opts ...pulumi.ResourceOption) (*AdvancedThreatProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	var resource AdvancedThreatProtection
	err := ctx.RegisterResource("azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdvancedThreatProtection gets an existing AdvancedThreatProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdvancedThreatProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdvancedThreatProtectionState, opts ...pulumi.ResourceOption) (*AdvancedThreatProtection, error) {
	var resource AdvancedThreatProtection
	err := ctx.ReadResource("azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdvancedThreatProtection resources.
type advancedThreatProtectionState struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type AdvancedThreatProtectionState struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled pulumi.BoolPtrInput
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (AdvancedThreatProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*advancedThreatProtectionState)(nil)).Elem()
}

type advancedThreatProtectionArgs struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled bool `pulumi:"enabled"`
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a AdvancedThreatProtection resource.
type AdvancedThreatProtectionArgs struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled pulumi.BoolInput
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (AdvancedThreatProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*advancedThreatProtectionArgs)(nil)).Elem()
}

type AdvancedThreatProtectionInput interface {
	pulumi.Input

	ToAdvancedThreatProtectionOutput() AdvancedThreatProtectionOutput
	ToAdvancedThreatProtectionOutputWithContext(ctx context.Context) AdvancedThreatProtectionOutput
}

func (AdvancedThreatProtection) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedThreatProtection)(nil)).Elem()
}

func (i AdvancedThreatProtection) ToAdvancedThreatProtectionOutput() AdvancedThreatProtectionOutput {
	return i.ToAdvancedThreatProtectionOutputWithContext(context.Background())
}

func (i AdvancedThreatProtection) ToAdvancedThreatProtectionOutputWithContext(ctx context.Context) AdvancedThreatProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedThreatProtectionOutput)
}

type AdvancedThreatProtectionOutput struct {
	*pulumi.OutputState
}

func (AdvancedThreatProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedThreatProtectionOutput)(nil)).Elem()
}

func (o AdvancedThreatProtectionOutput) ToAdvancedThreatProtectionOutput() AdvancedThreatProtectionOutput {
	return o
}

func (o AdvancedThreatProtectionOutput) ToAdvancedThreatProtectionOutputWithContext(ctx context.Context) AdvancedThreatProtectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AdvancedThreatProtectionOutput{})
}
