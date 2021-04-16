// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a site recovery network mapping on Azure. A network mapping decides how to translate connected netwroks when a VM is migrated from one region to another.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/siterecovery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primaryResourceGroup, err := core.NewResourceGroup(ctx, "primaryResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryResourceGroup, err := core.NewResourceGroup(ctx, "secondaryResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vault, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          secondaryResourceGroup.Location,
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryFabric, err := siterecovery.NewFabric(ctx, "primaryFabric", &siterecovery.FabricArgs{
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			RecoveryVaultName: vault.Name,
// 			Location:          primaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = siterecovery.NewFabric(ctx, "secondaryFabric", &siterecovery.FabricArgs{
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			RecoveryVaultName: vault.Name,
// 			Location:          secondaryResourceGroup.Location,
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			primaryFabric,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		primaryVirtualNetwork, err := network.NewVirtualNetwork(ctx, "primaryVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: primaryResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("192.168.1.0/24"),
// 			},
// 			Location: primaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryVirtualNetwork, err := network.NewVirtualNetwork(ctx, "secondaryVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: secondaryResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("192.168.2.0/24"),
// 			},
// 			Location: secondaryResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = siterecovery.NewNetworkMapping(ctx, "recovery_mapping", &siterecovery.NetworkMappingArgs{
// 			ResourceGroupName:        secondaryResourceGroup.Name,
// 			RecoveryVaultName:        vault.Name,
// 			SourceRecoveryFabricName: pulumi.String("primary-fabric"),
// 			TargetRecoveryFabricName: pulumi.String("secondary-fabric"),
// 			SourceNetworkId:          primaryVirtualNetwork.ID(),
// 			TargetNetworkId:          secondaryVirtualNetwork.ID(),
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
// Site Recovery Network Mapping can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:siterecovery/networkMapping:NetworkMapping mymapping /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/primary-fabric-name/replicationNetworks/azureNetwork/replicationNetworkMappings/mapping-name
// ```
type NetworkMapping struct {
	pulumi.CustomResourceState

	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The id of the primary network.
	SourceNetworkId pulumi.StringOutput `pulumi:"sourceNetworkId"`
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName pulumi.StringOutput `pulumi:"sourceRecoveryFabricName"`
	// The id of the recovery network.
	TargetNetworkId pulumi.StringOutput `pulumi:"targetNetworkId"`
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName pulumi.StringOutput `pulumi:"targetRecoveryFabricName"`
}

// NewNetworkMapping registers a new resource with the given unique name, arguments, and options.
func NewNetworkMapping(ctx *pulumi.Context,
	name string, args *NetworkMappingArgs, opts ...pulumi.ResourceOption) (*NetworkMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'SourceNetworkId'")
	}
	if args.SourceRecoveryFabricName == nil {
		return nil, errors.New("invalid value for required argument 'SourceRecoveryFabricName'")
	}
	if args.TargetNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'TargetNetworkId'")
	}
	if args.TargetRecoveryFabricName == nil {
		return nil, errors.New("invalid value for required argument 'TargetRecoveryFabricName'")
	}
	var resource NetworkMapping
	err := ctx.RegisterResource("azure:siterecovery/networkMapping:NetworkMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkMapping gets an existing NetworkMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkMappingState, opts ...pulumi.ResourceOption) (*NetworkMapping, error) {
	var resource NetworkMapping
	err := ctx.ReadResource("azure:siterecovery/networkMapping:NetworkMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkMapping resources.
type networkMappingState struct {
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The id of the primary network.
	SourceNetworkId *string `pulumi:"sourceNetworkId"`
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName *string `pulumi:"sourceRecoveryFabricName"`
	// The id of the recovery network.
	TargetNetworkId *string `pulumi:"targetNetworkId"`
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName *string `pulumi:"targetRecoveryFabricName"`
}

type NetworkMappingState struct {
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
	// The id of the primary network.
	SourceNetworkId pulumi.StringPtrInput
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName pulumi.StringPtrInput
	// The id of the recovery network.
	TargetNetworkId pulumi.StringPtrInput
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName pulumi.StringPtrInput
}

func (NetworkMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkMappingState)(nil)).Elem()
}

type networkMappingArgs struct {
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The id of the primary network.
	SourceNetworkId string `pulumi:"sourceNetworkId"`
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName string `pulumi:"sourceRecoveryFabricName"`
	// The id of the recovery network.
	TargetNetworkId string `pulumi:"targetNetworkId"`
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName string `pulumi:"targetRecoveryFabricName"`
}

// The set of arguments for constructing a NetworkMapping resource.
type NetworkMappingArgs struct {
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
	// The id of the primary network.
	SourceNetworkId pulumi.StringInput
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName pulumi.StringInput
	// The id of the recovery network.
	TargetNetworkId pulumi.StringInput
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName pulumi.StringInput
}

func (NetworkMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkMappingArgs)(nil)).Elem()
}

type NetworkMappingInput interface {
	pulumi.Input

	ToNetworkMappingOutput() NetworkMappingOutput
	ToNetworkMappingOutputWithContext(ctx context.Context) NetworkMappingOutput
}

func (*NetworkMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkMapping)(nil))
}

func (i *NetworkMapping) ToNetworkMappingOutput() NetworkMappingOutput {
	return i.ToNetworkMappingOutputWithContext(context.Background())
}

func (i *NetworkMapping) ToNetworkMappingOutputWithContext(ctx context.Context) NetworkMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingOutput)
}

func (i *NetworkMapping) ToNetworkMappingPtrOutput() NetworkMappingPtrOutput {
	return i.ToNetworkMappingPtrOutputWithContext(context.Background())
}

func (i *NetworkMapping) ToNetworkMappingPtrOutputWithContext(ctx context.Context) NetworkMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingPtrOutput)
}

type NetworkMappingPtrInput interface {
	pulumi.Input

	ToNetworkMappingPtrOutput() NetworkMappingPtrOutput
	ToNetworkMappingPtrOutputWithContext(ctx context.Context) NetworkMappingPtrOutput
}

type networkMappingPtrType NetworkMappingArgs

func (*networkMappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkMapping)(nil))
}

func (i *networkMappingPtrType) ToNetworkMappingPtrOutput() NetworkMappingPtrOutput {
	return i.ToNetworkMappingPtrOutputWithContext(context.Background())
}

func (i *networkMappingPtrType) ToNetworkMappingPtrOutputWithContext(ctx context.Context) NetworkMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingPtrOutput)
}

// NetworkMappingArrayInput is an input type that accepts NetworkMappingArray and NetworkMappingArrayOutput values.
// You can construct a concrete instance of `NetworkMappingArrayInput` via:
//
//          NetworkMappingArray{ NetworkMappingArgs{...} }
type NetworkMappingArrayInput interface {
	pulumi.Input

	ToNetworkMappingArrayOutput() NetworkMappingArrayOutput
	ToNetworkMappingArrayOutputWithContext(context.Context) NetworkMappingArrayOutput
}

type NetworkMappingArray []NetworkMappingInput

func (NetworkMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkMapping)(nil))
}

func (i NetworkMappingArray) ToNetworkMappingArrayOutput() NetworkMappingArrayOutput {
	return i.ToNetworkMappingArrayOutputWithContext(context.Background())
}

func (i NetworkMappingArray) ToNetworkMappingArrayOutputWithContext(ctx context.Context) NetworkMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingArrayOutput)
}

// NetworkMappingMapInput is an input type that accepts NetworkMappingMap and NetworkMappingMapOutput values.
// You can construct a concrete instance of `NetworkMappingMapInput` via:
//
//          NetworkMappingMap{ "key": NetworkMappingArgs{...} }
type NetworkMappingMapInput interface {
	pulumi.Input

	ToNetworkMappingMapOutput() NetworkMappingMapOutput
	ToNetworkMappingMapOutputWithContext(context.Context) NetworkMappingMapOutput
}

type NetworkMappingMap map[string]NetworkMappingInput

func (NetworkMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkMapping)(nil))
}

func (i NetworkMappingMap) ToNetworkMappingMapOutput() NetworkMappingMapOutput {
	return i.ToNetworkMappingMapOutputWithContext(context.Background())
}

func (i NetworkMappingMap) ToNetworkMappingMapOutputWithContext(ctx context.Context) NetworkMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingMapOutput)
}

type NetworkMappingOutput struct {
	*pulumi.OutputState
}

func (NetworkMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkMapping)(nil))
}

func (o NetworkMappingOutput) ToNetworkMappingOutput() NetworkMappingOutput {
	return o
}

func (o NetworkMappingOutput) ToNetworkMappingOutputWithContext(ctx context.Context) NetworkMappingOutput {
	return o
}

func (o NetworkMappingOutput) ToNetworkMappingPtrOutput() NetworkMappingPtrOutput {
	return o.ToNetworkMappingPtrOutputWithContext(context.Background())
}

func (o NetworkMappingOutput) ToNetworkMappingPtrOutputWithContext(ctx context.Context) NetworkMappingPtrOutput {
	return o.ApplyT(func(v NetworkMapping) *NetworkMapping {
		return &v
	}).(NetworkMappingPtrOutput)
}

type NetworkMappingPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkMappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkMapping)(nil))
}

func (o NetworkMappingPtrOutput) ToNetworkMappingPtrOutput() NetworkMappingPtrOutput {
	return o
}

func (o NetworkMappingPtrOutput) ToNetworkMappingPtrOutputWithContext(ctx context.Context) NetworkMappingPtrOutput {
	return o
}

type NetworkMappingArrayOutput struct{ *pulumi.OutputState }

func (NetworkMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkMapping)(nil))
}

func (o NetworkMappingArrayOutput) ToNetworkMappingArrayOutput() NetworkMappingArrayOutput {
	return o
}

func (o NetworkMappingArrayOutput) ToNetworkMappingArrayOutputWithContext(ctx context.Context) NetworkMappingArrayOutput {
	return o
}

func (o NetworkMappingArrayOutput) Index(i pulumi.IntInput) NetworkMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkMapping {
		return vs[0].([]NetworkMapping)[vs[1].(int)]
	}).(NetworkMappingOutput)
}

type NetworkMappingMapOutput struct{ *pulumi.OutputState }

func (NetworkMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkMapping)(nil))
}

func (o NetworkMappingMapOutput) ToNetworkMappingMapOutput() NetworkMappingMapOutput {
	return o
}

func (o NetworkMappingMapOutput) ToNetworkMappingMapOutputWithContext(ctx context.Context) NetworkMappingMapOutput {
	return o
}

func (o NetworkMappingMapOutput) MapIndex(k pulumi.StringInput) NetworkMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkMapping {
		return vs[0].(map[string]NetworkMapping)[vs[1].(string)]
	}).(NetworkMappingOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkMappingOutput{})
	pulumi.RegisterOutputType(NetworkMappingPtrOutput{})
	pulumi.RegisterOutputType(NetworkMappingArrayOutput{})
	pulumi.RegisterOutputType(NetworkMappingMapOutput{})
}
