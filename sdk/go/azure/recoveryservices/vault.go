// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Recovery Services Vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku:               pulumi.String("Standard"),
// 			SoftDeleteEnabled: pulumi.Bool(true),
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
// Recovery Services Vaults can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:recoveryservices/vault:Vault vault1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1
// ```
type Vault struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// Is soft delete enable for this Vault? Defaults to `true`.
	SoftDeleteEnabled pulumi.BoolPtrOutput `pulumi:"softDeleteEnabled"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Vault
	err := ctx.RegisterResource("azure:recoveryservices/vault:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("azure:recoveryservices/vault:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
	Sku *string `pulumi:"sku"`
	// Is soft delete enable for this Vault? Defaults to `true`.
	SoftDeleteEnabled *bool `pulumi:"softDeleteEnabled"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VaultState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
	Sku pulumi.StringPtrInput
	// Is soft delete enable for this Vault? Defaults to `true`.
	SoftDeleteEnabled pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
	Sku string `pulumi:"sku"`
	// Is soft delete enable for this Vault? Defaults to `true`.
	SoftDeleteEnabled *bool `pulumi:"softDeleteEnabled"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Sets the vault's SKU. Possible values include: `Standard`, `RS0`.
	Sku pulumi.StringInput
	// Is soft delete enable for this Vault? Defaults to `true`.
	SoftDeleteEnabled pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

type VaultPtrInput interface {
	pulumi.Input

	ToVaultPtrOutput() VaultPtrOutput
	ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput
}

func (Vault) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil)).Elem()
}

func (i Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

func (i Vault) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i Vault) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPtrOutput)
}

type VaultOutput struct {
	*pulumi.OutputState
}

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultOutput)(nil)).Elem()
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

type VaultPtrOutput struct {
	*pulumi.OutputState
}

func (VaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil)).Elem()
}

func (o VaultPtrOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o
}

func (o VaultPtrOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VaultOutput{})
	pulumi.RegisterOutputType(VaultPtrOutput{})
}
