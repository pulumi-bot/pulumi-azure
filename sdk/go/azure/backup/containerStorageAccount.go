// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages registration of a storage account with Azure Backup. Storage accounts must be registered with an Azure Recovery Vault in order to backup file shares within the storage account. Registering a storage account with a vault creates what is known as a protection container within Azure Recovery Services. Once the container is created, Azure file shares within the storage account can be backed up using the `backup.ProtectedFileShare` resource.
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		sa, err := storage.NewAccount(ctx, "sa", &storage.AccountArgs{
// 			Location:               rg.Location,
// 			ResourceGroupName:      rg.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewContainerStorageAccount(ctx, "container", &backup.ContainerStorageAccountArgs{
// 			ResourceGroupName: rg.Name,
// 			RecoveryVaultName: vault.Name,
// 			StorageAccountId:  sa.ID(),
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
// Backup Storage Account Containers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:backup/containerStorageAccount:ContainerStorageAccount mycontainer "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/backupFabrics/Azure/protectionContainers/StorageContainer;storage;storage-rg-name;storage-account"
// ```
//
//  Note the ID requires quoting as there are semicolons
type ContainerStorageAccount struct {
	pulumi.CustomResourceState

	// The name of the vault where the storage account will be registered.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Storage Account to be registered
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
}

// NewContainerStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewContainerStorageAccount(ctx *pulumi.Context,
	name string, args *ContainerStorageAccountArgs, opts ...pulumi.ResourceOption) (*ContainerStorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource ContainerStorageAccount
	err := ctx.RegisterResource("azure:backup/containerStorageAccount:ContainerStorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerStorageAccount gets an existing ContainerStorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerStorageAccountState, opts ...pulumi.ResourceOption) (*ContainerStorageAccount, error) {
	var resource ContainerStorageAccount
	err := ctx.ReadResource("azure:backup/containerStorageAccount:ContainerStorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerStorageAccount resources.
type containerStorageAccountState struct {
	// The name of the vault where the storage account will be registered.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Storage Account to be registered
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type ContainerStorageAccountState struct {
	// The name of the vault where the storage account will be registered.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault is located.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Storage Account to be registered
	StorageAccountId pulumi.StringPtrInput
}

func (ContainerStorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerStorageAccountState)(nil)).Elem()
}

type containerStorageAccountArgs struct {
	// The name of the vault where the storage account will be registered.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Storage Account to be registered
	StorageAccountId string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a ContainerStorageAccount resource.
type ContainerStorageAccountArgs struct {
	// The name of the vault where the storage account will be registered.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault is located.
	ResourceGroupName pulumi.StringInput
	// The ID of the Storage Account to be registered
	StorageAccountId pulumi.StringInput
}

func (ContainerStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerStorageAccountArgs)(nil)).Elem()
}

type ContainerStorageAccountInput interface {
	pulumi.Input

	ToContainerStorageAccountOutput() ContainerStorageAccountOutput
	ToContainerStorageAccountOutputWithContext(ctx context.Context) ContainerStorageAccountOutput
}

func (*ContainerStorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerStorageAccount)(nil))
}

func (i *ContainerStorageAccount) ToContainerStorageAccountOutput() ContainerStorageAccountOutput {
	return i.ToContainerStorageAccountOutputWithContext(context.Background())
}

func (i *ContainerStorageAccount) ToContainerStorageAccountOutputWithContext(ctx context.Context) ContainerStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerStorageAccountOutput)
}

func (i *ContainerStorageAccount) ToContainerStorageAccountPtrOutput() ContainerStorageAccountPtrOutput {
	return i.ToContainerStorageAccountPtrOutputWithContext(context.Background())
}

func (i *ContainerStorageAccount) ToContainerStorageAccountPtrOutputWithContext(ctx context.Context) ContainerStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerStorageAccountPtrOutput)
}

type ContainerStorageAccountPtrInput interface {
	pulumi.Input

	ToContainerStorageAccountPtrOutput() ContainerStorageAccountPtrOutput
	ToContainerStorageAccountPtrOutputWithContext(ctx context.Context) ContainerStorageAccountPtrOutput
}

type containerStorageAccountPtrType ContainerStorageAccountArgs

func (*containerStorageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerStorageAccount)(nil))
}

func (i *containerStorageAccountPtrType) ToContainerStorageAccountPtrOutput() ContainerStorageAccountPtrOutput {
	return i.ToContainerStorageAccountPtrOutputWithContext(context.Background())
}

func (i *containerStorageAccountPtrType) ToContainerStorageAccountPtrOutputWithContext(ctx context.Context) ContainerStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerStorageAccountOutput).ToContainerStorageAccountPtrOutput()
}

// ContainerStorageAccountArrayInput is an input type that accepts ContainerStorageAccountArray and ContainerStorageAccountArrayOutput values.
// You can construct a concrete instance of `ContainerStorageAccountArrayInput` via:
//
//          ContainerStorageAccountArray{ ContainerStorageAccountArgs{...} }
type ContainerStorageAccountArrayInput interface {
	pulumi.Input

	ToContainerStorageAccountArrayOutput() ContainerStorageAccountArrayOutput
	ToContainerStorageAccountArrayOutputWithContext(context.Context) ContainerStorageAccountArrayOutput
}

type ContainerStorageAccountArray []ContainerStorageAccountInput

func (ContainerStorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerStorageAccount)(nil))
}

func (i ContainerStorageAccountArray) ToContainerStorageAccountArrayOutput() ContainerStorageAccountArrayOutput {
	return i.ToContainerStorageAccountArrayOutputWithContext(context.Background())
}

func (i ContainerStorageAccountArray) ToContainerStorageAccountArrayOutputWithContext(ctx context.Context) ContainerStorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerStorageAccountArrayOutput)
}

// ContainerStorageAccountMapInput is an input type that accepts ContainerStorageAccountMap and ContainerStorageAccountMapOutput values.
// You can construct a concrete instance of `ContainerStorageAccountMapInput` via:
//
//          ContainerStorageAccountMap{ "key": ContainerStorageAccountArgs{...} }
type ContainerStorageAccountMapInput interface {
	pulumi.Input

	ToContainerStorageAccountMapOutput() ContainerStorageAccountMapOutput
	ToContainerStorageAccountMapOutputWithContext(context.Context) ContainerStorageAccountMapOutput
}

type ContainerStorageAccountMap map[string]ContainerStorageAccountInput

func (ContainerStorageAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerStorageAccount)(nil))
}

func (i ContainerStorageAccountMap) ToContainerStorageAccountMapOutput() ContainerStorageAccountMapOutput {
	return i.ToContainerStorageAccountMapOutputWithContext(context.Background())
}

func (i ContainerStorageAccountMap) ToContainerStorageAccountMapOutputWithContext(ctx context.Context) ContainerStorageAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerStorageAccountMapOutput)
}

type ContainerStorageAccountOutput struct {
	*pulumi.OutputState
}

func (ContainerStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerStorageAccount)(nil))
}

func (o ContainerStorageAccountOutput) ToContainerStorageAccountOutput() ContainerStorageAccountOutput {
	return o
}

func (o ContainerStorageAccountOutput) ToContainerStorageAccountOutputWithContext(ctx context.Context) ContainerStorageAccountOutput {
	return o
}

func (o ContainerStorageAccountOutput) ToContainerStorageAccountPtrOutput() ContainerStorageAccountPtrOutput {
	return o.ToContainerStorageAccountPtrOutputWithContext(context.Background())
}

func (o ContainerStorageAccountOutput) ToContainerStorageAccountPtrOutputWithContext(ctx context.Context) ContainerStorageAccountPtrOutput {
	return o.ApplyT(func(v ContainerStorageAccount) *ContainerStorageAccount {
		return &v
	}).(ContainerStorageAccountPtrOutput)
}

type ContainerStorageAccountPtrOutput struct {
	*pulumi.OutputState
}

func (ContainerStorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerStorageAccount)(nil))
}

func (o ContainerStorageAccountPtrOutput) ToContainerStorageAccountPtrOutput() ContainerStorageAccountPtrOutput {
	return o
}

func (o ContainerStorageAccountPtrOutput) ToContainerStorageAccountPtrOutputWithContext(ctx context.Context) ContainerStorageAccountPtrOutput {
	return o
}

type ContainerStorageAccountArrayOutput struct{ *pulumi.OutputState }

func (ContainerStorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerStorageAccount)(nil))
}

func (o ContainerStorageAccountArrayOutput) ToContainerStorageAccountArrayOutput() ContainerStorageAccountArrayOutput {
	return o
}

func (o ContainerStorageAccountArrayOutput) ToContainerStorageAccountArrayOutputWithContext(ctx context.Context) ContainerStorageAccountArrayOutput {
	return o
}

func (o ContainerStorageAccountArrayOutput) Index(i pulumi.IntInput) ContainerStorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerStorageAccount {
		return vs[0].([]ContainerStorageAccount)[vs[1].(int)]
	}).(ContainerStorageAccountOutput)
}

type ContainerStorageAccountMapOutput struct{ *pulumi.OutputState }

func (ContainerStorageAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerStorageAccount)(nil))
}

func (o ContainerStorageAccountMapOutput) ToContainerStorageAccountMapOutput() ContainerStorageAccountMapOutput {
	return o
}

func (o ContainerStorageAccountMapOutput) ToContainerStorageAccountMapOutputWithContext(ctx context.Context) ContainerStorageAccountMapOutput {
	return o
}

func (o ContainerStorageAccountMapOutput) MapIndex(k pulumi.StringInput) ContainerStorageAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ContainerStorageAccount {
		return vs[0].(map[string]ContainerStorageAccount)[vs[1].(string)]
	}).(ContainerStorageAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerStorageAccountOutput{})
	pulumi.RegisterOutputType(ContainerStorageAccountPtrOutput{})
	pulumi.RegisterOutputType(ContainerStorageAccountArrayOutput{})
	pulumi.RegisterOutputType(ContainerStorageAccountMapOutput{})
}
