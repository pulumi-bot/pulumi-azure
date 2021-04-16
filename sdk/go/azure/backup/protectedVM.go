// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Azure Backup for an Azure VM
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		exampleVault, err := recoveryservices.NewVault(ctx, "exampleVault", &recoveryservices.VaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePolicyVM, err := backup.NewPolicyVM(ctx, "examplePolicyVM", &backup.PolicyVMArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			RecoveryVaultName: exampleVault.Name,
// 			Backup: &backup.PolicyVMBackupArgs{
// 				Frequency: pulumi.String("Daily"),
// 				Time:      pulumi.String("23:00"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewProtectedVM(ctx, "vm1", &backup.ProtectedVMArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			RecoveryVaultName: exampleVault.Name,
// 			SourceVmId:        pulumi.Any(azurerm_virtual_machine.Example.Id),
// 			BackupPolicyId:    examplePolicyVM.ID(),
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
// Recovery Services Protected VMs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:backup/protectedVM:ProtectedVM item1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/iaasvmcontainer;iaasvmcontainerv2;group1;vm1/protectedItems/vm;iaasvmcontainerv2;group1;vm1"
// ```
//
//  Note the ID requires quoting as there are semicolons
type ProtectedVM struct {
	pulumi.CustomResourceState

	// Specifies the id of the backup policy to use.
	BackupPolicyId pulumi.StringOutput `pulumi:"backupPolicyId"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVmId pulumi.StringOutput `pulumi:"sourceVmId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewProtectedVM registers a new resource with the given unique name, arguments, and options.
func NewProtectedVM(ctx *pulumi.Context,
	name string, args *ProtectedVMArgs, opts ...pulumi.ResourceOption) (*ProtectedVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'BackupPolicyId'")
	}
	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceVmId == nil {
		return nil, errors.New("invalid value for required argument 'SourceVmId'")
	}
	var resource ProtectedVM
	err := ctx.RegisterResource("azure:backup/protectedVM:ProtectedVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectedVM gets an existing ProtectedVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectedVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectedVMState, opts ...pulumi.ResourceOption) (*ProtectedVM, error) {
	var resource ProtectedVM
	err := ctx.ReadResource("azure:backup/protectedVM:ProtectedVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectedVM resources.
type protectedVMState struct {
	// Specifies the id of the backup policy to use.
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVmId *string `pulumi:"sourceVmId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ProtectedVMState struct {
	// Specifies the id of the backup policy to use.
	BackupPolicyId pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringPtrInput
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVmId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ProtectedVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedVMState)(nil)).Elem()
}

type protectedVMArgs struct {
	// Specifies the id of the backup policy to use.
	BackupPolicyId string `pulumi:"backupPolicyId"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVmId string `pulumi:"sourceVmId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ProtectedVM resource.
type ProtectedVMArgs struct {
	// Specifies the id of the backup policy to use.
	BackupPolicyId pulumi.StringInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringInput
	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVmId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ProtectedVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedVMArgs)(nil)).Elem()
}

type ProtectedVMInput interface {
	pulumi.Input

	ToProtectedVMOutput() ProtectedVMOutput
	ToProtectedVMOutputWithContext(ctx context.Context) ProtectedVMOutput
}

func (*ProtectedVM) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedVM)(nil))
}

func (i *ProtectedVM) ToProtectedVMOutput() ProtectedVMOutput {
	return i.ToProtectedVMOutputWithContext(context.Background())
}

func (i *ProtectedVM) ToProtectedVMOutputWithContext(ctx context.Context) ProtectedVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedVMOutput)
}

func (i *ProtectedVM) ToProtectedVMPtrOutput() ProtectedVMPtrOutput {
	return i.ToProtectedVMPtrOutputWithContext(context.Background())
}

func (i *ProtectedVM) ToProtectedVMPtrOutputWithContext(ctx context.Context) ProtectedVMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedVMPtrOutput)
}

type ProtectedVMPtrInput interface {
	pulumi.Input

	ToProtectedVMPtrOutput() ProtectedVMPtrOutput
	ToProtectedVMPtrOutputWithContext(ctx context.Context) ProtectedVMPtrOutput
}

type protectedVMPtrType ProtectedVMArgs

func (*protectedVMPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedVM)(nil))
}

func (i *protectedVMPtrType) ToProtectedVMPtrOutput() ProtectedVMPtrOutput {
	return i.ToProtectedVMPtrOutputWithContext(context.Background())
}

func (i *protectedVMPtrType) ToProtectedVMPtrOutputWithContext(ctx context.Context) ProtectedVMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedVMPtrOutput)
}

// ProtectedVMArrayInput is an input type that accepts ProtectedVMArray and ProtectedVMArrayOutput values.
// You can construct a concrete instance of `ProtectedVMArrayInput` via:
//
//          ProtectedVMArray{ ProtectedVMArgs{...} }
type ProtectedVMArrayInput interface {
	pulumi.Input

	ToProtectedVMArrayOutput() ProtectedVMArrayOutput
	ToProtectedVMArrayOutputWithContext(context.Context) ProtectedVMArrayOutput
}

type ProtectedVMArray []ProtectedVMInput

func (ProtectedVMArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ProtectedVM)(nil))
}

func (i ProtectedVMArray) ToProtectedVMArrayOutput() ProtectedVMArrayOutput {
	return i.ToProtectedVMArrayOutputWithContext(context.Background())
}

func (i ProtectedVMArray) ToProtectedVMArrayOutputWithContext(ctx context.Context) ProtectedVMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedVMArrayOutput)
}

// ProtectedVMMapInput is an input type that accepts ProtectedVMMap and ProtectedVMMapOutput values.
// You can construct a concrete instance of `ProtectedVMMapInput` via:
//
//          ProtectedVMMap{ "key": ProtectedVMArgs{...} }
type ProtectedVMMapInput interface {
	pulumi.Input

	ToProtectedVMMapOutput() ProtectedVMMapOutput
	ToProtectedVMMapOutputWithContext(context.Context) ProtectedVMMapOutput
}

type ProtectedVMMap map[string]ProtectedVMInput

func (ProtectedVMMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ProtectedVM)(nil))
}

func (i ProtectedVMMap) ToProtectedVMMapOutput() ProtectedVMMapOutput {
	return i.ToProtectedVMMapOutputWithContext(context.Background())
}

func (i ProtectedVMMap) ToProtectedVMMapOutputWithContext(ctx context.Context) ProtectedVMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedVMMapOutput)
}

type ProtectedVMOutput struct {
	*pulumi.OutputState
}

func (ProtectedVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedVM)(nil))
}

func (o ProtectedVMOutput) ToProtectedVMOutput() ProtectedVMOutput {
	return o
}

func (o ProtectedVMOutput) ToProtectedVMOutputWithContext(ctx context.Context) ProtectedVMOutput {
	return o
}

func (o ProtectedVMOutput) ToProtectedVMPtrOutput() ProtectedVMPtrOutput {
	return o.ToProtectedVMPtrOutputWithContext(context.Background())
}

func (o ProtectedVMOutput) ToProtectedVMPtrOutputWithContext(ctx context.Context) ProtectedVMPtrOutput {
	return o.ApplyT(func(v ProtectedVM) *ProtectedVM {
		return &v
	}).(ProtectedVMPtrOutput)
}

type ProtectedVMPtrOutput struct {
	*pulumi.OutputState
}

func (ProtectedVMPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedVM)(nil))
}

func (o ProtectedVMPtrOutput) ToProtectedVMPtrOutput() ProtectedVMPtrOutput {
	return o
}

func (o ProtectedVMPtrOutput) ToProtectedVMPtrOutputWithContext(ctx context.Context) ProtectedVMPtrOutput {
	return o
}

type ProtectedVMArrayOutput struct{ *pulumi.OutputState }

func (ProtectedVMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProtectedVM)(nil))
}

func (o ProtectedVMArrayOutput) ToProtectedVMArrayOutput() ProtectedVMArrayOutput {
	return o
}

func (o ProtectedVMArrayOutput) ToProtectedVMArrayOutputWithContext(ctx context.Context) ProtectedVMArrayOutput {
	return o
}

func (o ProtectedVMArrayOutput) Index(i pulumi.IntInput) ProtectedVMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProtectedVM {
		return vs[0].([]ProtectedVM)[vs[1].(int)]
	}).(ProtectedVMOutput)
}

type ProtectedVMMapOutput struct{ *pulumi.OutputState }

func (ProtectedVMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProtectedVM)(nil))
}

func (o ProtectedVMMapOutput) ToProtectedVMMapOutput() ProtectedVMMapOutput {
	return o
}

func (o ProtectedVMMapOutput) ToProtectedVMMapOutputWithContext(ctx context.Context) ProtectedVMMapOutput {
	return o
}

func (o ProtectedVMMapOutput) MapIndex(k pulumi.StringInput) ProtectedVMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProtectedVM {
		return vs[0].(map[string]ProtectedVM)[vs[1].(string)]
	}).(ProtectedVMOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectedVMOutput{})
	pulumi.RegisterOutputType(ProtectedVMPtrOutput{})
	pulumi.RegisterOutputType(ProtectedVMArrayOutput{})
	pulumi.RegisterOutputType(ProtectedVMMapOutput{})
}
