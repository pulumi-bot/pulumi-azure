// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure File Share Backup Policy within a Recovery Services vault.
//
// > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/backup"
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
// 		vault, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewPolicyFileShare(ctx, "policy", &backup.PolicyFileShareArgs{
// 			ResourceGroupName: rg.Name,
// 			RecoveryVaultName: vault.Name,
// 			Timezone:          pulumi.String("UTC"),
// 			Backup: &backup.PolicyFileShareBackupArgs{
// 				Frequency: pulumi.String("Daily"),
// 				Time:      pulumi.String("23:00"),
// 			},
// 			RetentionDaily: &backup.PolicyFileShareRetentionDailyArgs{
// 				Count: pulumi.Int(10),
// 			},
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
// Azure File Share Backup Policies can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:backup/policyFileShare:PolicyFileShare policy1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
// ```
type PolicyFileShare struct {
	pulumi.CustomResourceState

	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackupOutput `pulumi:"backup"`
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDailyOutput `pulumi:"retentionDaily"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewPolicyFileShare registers a new resource with the given unique name, arguments, and options.
func NewPolicyFileShare(ctx *pulumi.Context,
	name string, args *PolicyFileShareArgs, opts ...pulumi.ResourceOption) (*PolicyFileShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backup == nil {
		return nil, errors.New("invalid value for required argument 'Backup'")
	}
	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RetentionDaily == nil {
		return nil, errors.New("invalid value for required argument 'RetentionDaily'")
	}
	var resource PolicyFileShare
	err := ctx.RegisterResource("azure:backup/policyFileShare:PolicyFileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyFileShare gets an existing PolicyFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyFileShareState, opts ...pulumi.ResourceOption) (*PolicyFileShare, error) {
	var resource PolicyFileShare
	err := ctx.ReadResource("azure:backup/policyFileShare:PolicyFileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyFileShare resources.
type policyFileShareState struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup *PolicyFileShareBackup `pulumi:"backup"`
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily *PolicyFileShareRetentionDaily `pulumi:"retentionDaily"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone *string `pulumi:"timezone"`
}

type PolicyFileShareState struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackupPtrInput
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringPtrInput
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDailyPtrInput
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrInput
}

func (PolicyFileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFileShareState)(nil)).Elem()
}

type policyFileShareArgs struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackup `pulumi:"backup"`
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDaily `pulumi:"retentionDaily"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a PolicyFileShare resource.
type PolicyFileShareArgs struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackupInput
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringInput
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDailyInput
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrInput
}

func (PolicyFileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFileShareArgs)(nil)).Elem()
}

type PolicyFileShareInput interface {
	pulumi.Input

	ToPolicyFileShareOutput() PolicyFileShareOutput
	ToPolicyFileShareOutputWithContext(ctx context.Context) PolicyFileShareOutput
}

func (PolicyFileShare) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyFileShare)(nil)).Elem()
}

func (i PolicyFileShare) ToPolicyFileShareOutput() PolicyFileShareOutput {
	return i.ToPolicyFileShareOutputWithContext(context.Background())
}

func (i PolicyFileShare) ToPolicyFileShareOutputWithContext(ctx context.Context) PolicyFileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyFileShareOutput)
}

type PolicyFileShareOutput struct {
	*pulumi.OutputState
}

func (PolicyFileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyFileShareOutput)(nil)).Elem()
}

func (o PolicyFileShareOutput) ToPolicyFileShareOutput() PolicyFileShareOutput {
	return o
}

func (o PolicyFileShareOutput) ToPolicyFileShareOutputWithContext(ctx context.Context) PolicyFileShareOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyFileShareOutput{})
}
