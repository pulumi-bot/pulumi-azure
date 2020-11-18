// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Site Recovery Replication Fabric within a Recovery Services vault. Only Azure fabrics are supported at this time. Replication Fabrics serve as a container within an Azure region for other Site Recovery resources such as protection containers, protected items, network mappings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/siterecovery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.NewResourceGroup(ctx, "primary", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondary, err := core.NewResourceGroup(ctx, "secondary", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vault, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          secondary.Location,
// 			ResourceGroupName: secondary.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = siterecovery.NewFabric(ctx, "fabric", &siterecovery.FabricArgs{
// 			ResourceGroupName: secondary.Name,
// 			RecoveryVaultName: vault.Name,
// 			Location:          primary.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Fabric struct {
	pulumi.CustomResourceState

	// In what region should the fabric be located.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewFabric registers a new resource with the given unique name, arguments, and options.
func NewFabric(ctx *pulumi.Context,
	name string, args *FabricArgs, opts ...pulumi.ResourceOption) (*Fabric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Fabric
	err := ctx.RegisterResource("azure:siterecovery/fabric:Fabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFabric gets an existing Fabric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FabricState, opts ...pulumi.ResourceOption) (*Fabric, error) {
	var resource Fabric
	err := ctx.ReadResource("azure:siterecovery/fabric:Fabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fabric resources.
type fabricState struct {
	// In what region should the fabric be located.
	Location *string `pulumi:"location"`
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type FabricState struct {
	// In what region should the fabric be located.
	Location pulumi.StringPtrInput
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
}

func (FabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricState)(nil)).Elem()
}

type fabricArgs struct {
	// In what region should the fabric be located.
	Location *string `pulumi:"location"`
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Fabric resource.
type FabricArgs struct {
	// In what region should the fabric be located.
	Location pulumi.StringPtrInput
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
}

func (FabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricArgs)(nil)).Elem()
}

type FabricInput interface {
	pulumi.Input

	ToFabricOutput() FabricOutput
	ToFabricOutputWithContext(ctx context.Context) FabricOutput
}

func (Fabric) ElementType() reflect.Type {
	return reflect.TypeOf((*Fabric)(nil)).Elem()
}

func (i Fabric) ToFabricOutput() FabricOutput {
	return i.ToFabricOutputWithContext(context.Background())
}

func (i Fabric) ToFabricOutputWithContext(ctx context.Context) FabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricOutput)
}

type FabricOutput struct {
	*pulumi.OutputState
}

func (FabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FabricOutput)(nil)).Elem()
}

func (o FabricOutput) ToFabricOutput() FabricOutput {
	return o
}

func (o FabricOutput) ToFabricOutputWithContext(ctx context.Context) FabricOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FabricOutput{})
}
