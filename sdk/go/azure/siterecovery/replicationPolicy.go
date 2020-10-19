// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Site Recovery replication policy within a recovery vault. Replication policies define the frequency at which recovery points are created and how long they are stored.
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
// 		_, err = siterecovery.NewReplicationPolicy(ctx, "policy", &siterecovery.ReplicationPolicyArgs{
// 			ResourceGroupName:                               secondary.Name,
// 			RecoveryVaultName:                               vault.Name,
// 			RecoveryPointRetentionInMinutes:                 pulumi.Int(24 * 60),
// 			ApplicationConsistentSnapshotFrequencyInMinutes: pulumi.Int(4 * 60),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ReplicationPolicy struct {
	pulumi.CustomResourceState

	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	ApplicationConsistentSnapshotFrequencyInMinutes pulumi.IntOutput `pulumi:"applicationConsistentSnapshotFrequencyInMinutes"`
	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`
	// The duration in minutes for which the recovery points need to be stored.
	RecoveryPointRetentionInMinutes pulumi.IntOutput `pulumi:"recoveryPointRetentionInMinutes"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewReplicationPolicy registers a new resource with the given unique name, arguments, and options.
func NewReplicationPolicy(ctx *pulumi.Context,
	name string, args *ReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	if args == nil || args.ApplicationConsistentSnapshotFrequencyInMinutes == nil {
		return nil, errors.New("missing required argument 'ApplicationConsistentSnapshotFrequencyInMinutes'")
	}
	if args == nil || args.RecoveryPointRetentionInMinutes == nil {
		return nil, errors.New("missing required argument 'RecoveryPointRetentionInMinutes'")
	}
	if args == nil || args.RecoveryVaultName == nil {
		return nil, errors.New("missing required argument 'RecoveryVaultName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ReplicationPolicyArgs{}
	}
	var resource ReplicationPolicy
	err := ctx.RegisterResource("azure:siterecovery/replicationPolicy:ReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationPolicy gets an existing ReplicationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationPolicyState, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	var resource ReplicationPolicy
	err := ctx.ReadResource("azure:siterecovery/replicationPolicy:ReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationPolicy resources.
type replicationPolicyState struct {
	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	ApplicationConsistentSnapshotFrequencyInMinutes *int `pulumi:"applicationConsistentSnapshotFrequencyInMinutes"`
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// The duration in minutes for which the recovery points need to be stored.
	RecoveryPointRetentionInMinutes *int `pulumi:"recoveryPointRetentionInMinutes"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ReplicationPolicyState struct {
	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	ApplicationConsistentSnapshotFrequencyInMinutes pulumi.IntPtrInput
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// The duration in minutes for which the recovery points need to be stored.
	RecoveryPointRetentionInMinutes pulumi.IntPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
}

func (ReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyState)(nil)).Elem()
}

type replicationPolicyArgs struct {
	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	ApplicationConsistentSnapshotFrequencyInMinutes int `pulumi:"applicationConsistentSnapshotFrequencyInMinutes"`
	// The name of the network mapping.
	Name *string `pulumi:"name"`
	// The duration in minutes for which the recovery points need to be stored.
	RecoveryPointRetentionInMinutes int `pulumi:"recoveryPointRetentionInMinutes"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ReplicationPolicy resource.
type ReplicationPolicyArgs struct {
	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	ApplicationConsistentSnapshotFrequencyInMinutes pulumi.IntInput
	// The name of the network mapping.
	Name pulumi.StringPtrInput
	// The duration in minutes for which the recovery points need to be stored.
	RecoveryPointRetentionInMinutes pulumi.IntInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
}

func (ReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyArgs)(nil)).Elem()
}

type ReplicationPolicyInput interface {
	pulumi.Input

	ToReplicationPolicyOutput() ReplicationPolicyOutput
	ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput
}

func (ReplicationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationPolicy)(nil)).Elem()
}

func (i ReplicationPolicy) ToReplicationPolicyOutput() ReplicationPolicyOutput {
	return i.ToReplicationPolicyOutputWithContext(context.Background())
}

func (i ReplicationPolicy) ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationPolicyOutput)
}

type ReplicationPolicyOutput struct {
	*pulumi.OutputState
}

func (ReplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationPolicyOutput)(nil)).Elem()
}

func (o ReplicationPolicyOutput) ToReplicationPolicyOutput() ReplicationPolicyOutput {
	return o
}

func (o ReplicationPolicyOutput) ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationPolicyOutput{})
}
