// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Key Vault Access Policy.
//
// > **NOTE:** It's possible to define Key Vault Access Policies both within the `keyvault.KeyVault` resource via the `accessPolicy` block and by using the `keyvault.AccessPolicy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
//
// > **NOTE:** Azure permits a maximum of 1024 Access Policies per Key Vault - [more information can be found in this document](https://docs.microsoft.com/en-us/azure/key-vault/key-vault-secure-your-key-vault#data-plane-access-control).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.String(current.TenantId),
// 			SkuName:           pulumi.String("premium"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.NewAccessPolicy(ctx, "exampleAccessPolicy", &keyvault.AccessPolicyArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			TenantId:   pulumi.String(current.TenantId),
// 			ObjectId:   pulumi.String(current.ObjectId),
// 			KeyPermissions: pulumi.StringArray{
// 				pulumi.String("Get"),
// 			},
// 			SecretPermissions: pulumi.StringArray{
// 				pulumi.String("Get"),
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
// Key Vault Access Policies can be imported using the Resource ID of the Key Vault, plus some additional metadata. If both an `object_id` and `application_id` are specified, then the Access Policy can be imported using the following code
//
// ```sh
//  $ pulumi import azure:keyvault/accessPolicy:AccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111/applicationId/22222222-2222-2222-2222-222222222222
// ```
//
//  where `11111111-1111-1111-1111-111111111111` is the `object_id` and `22222222-2222-2222-2222-222222222222` is the `application_id`. --- Access Policies with an `object_id` but no `application_id` can be imported using the following command
//
// ```sh
//  $ pulumi import azure:keyvault/accessPolicy:AccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111
// ```
//
//  where `11111111-1111-1111-1111-111111111111` is the `object_id`.
type AccessPolicy struct {
	pulumi.CustomResourceState

	// The object ID of an Application in Azure Active Directory.
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
	CertificatePermissions pulumi.StringArrayOutput `pulumi:"certificatePermissions"`
	// List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
	KeyPermissions pulumi.StringArrayOutput `pulumi:"keyPermissions"`
	// Specifies the id of the Key Vault resource. Changing this
	// forces a new resource to be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// The object ID of a user, service principal or security
	// group in the Azure Active Directory tenant for the vault. The object ID must
	// be unique for the list of access policies. Changing this forces a new resource
	// to be created.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
	SecretPermissions pulumi.StringArrayOutput `pulumi:"secretPermissions"`
	// List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
	StoragePermissions pulumi.StringArrayOutput `pulumi:"storagePermissions"`
	// The Azure Active Directory tenant ID that should be used
	// for authenticating requests to the key vault. Changing this forces a new resource
	// to be created.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource AccessPolicy
	err := ctx.RegisterResource("azure:keyvault/accessPolicy:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicy gets an existing AccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("azure:keyvault/accessPolicy:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicy resources.
type accessPolicyState struct {
	// The object ID of an Application in Azure Active Directory.
	ApplicationId *string `pulumi:"applicationId"`
	// List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
	CertificatePermissions []string `pulumi:"certificatePermissions"`
	// List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
	KeyPermissions []string `pulumi:"keyPermissions"`
	// Specifies the id of the Key Vault resource. Changing this
	// forces a new resource to be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// The object ID of a user, service principal or security
	// group in the Azure Active Directory tenant for the vault. The object ID must
	// be unique for the list of access policies. Changing this forces a new resource
	// to be created.
	ObjectId *string `pulumi:"objectId"`
	// List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
	SecretPermissions []string `pulumi:"secretPermissions"`
	// List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
	StoragePermissions []string `pulumi:"storagePermissions"`
	// The Azure Active Directory tenant ID that should be used
	// for authenticating requests to the key vault. Changing this forces a new resource
	// to be created.
	TenantId *string `pulumi:"tenantId"`
}

type AccessPolicyState struct {
	// The object ID of an Application in Azure Active Directory.
	ApplicationId pulumi.StringPtrInput
	// List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
	CertificatePermissions pulumi.StringArrayInput
	// List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
	KeyPermissions pulumi.StringArrayInput
	// Specifies the id of the Key Vault resource. Changing this
	// forces a new resource to be created.
	KeyVaultId pulumi.StringPtrInput
	// The object ID of a user, service principal or security
	// group in the Azure Active Directory tenant for the vault. The object ID must
	// be unique for the list of access policies. Changing this forces a new resource
	// to be created.
	ObjectId pulumi.StringPtrInput
	// List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
	SecretPermissions pulumi.StringArrayInput
	// List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
	StoragePermissions pulumi.StringArrayInput
	// The Azure Active Directory tenant ID that should be used
	// for authenticating requests to the key vault. Changing this forces a new resource
	// to be created.
	TenantId pulumi.StringPtrInput
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	// The object ID of an Application in Azure Active Directory.
	ApplicationId *string `pulumi:"applicationId"`
	// List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
	CertificatePermissions []string `pulumi:"certificatePermissions"`
	// List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
	KeyPermissions []string `pulumi:"keyPermissions"`
	// Specifies the id of the Key Vault resource. Changing this
	// forces a new resource to be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The object ID of a user, service principal or security
	// group in the Azure Active Directory tenant for the vault. The object ID must
	// be unique for the list of access policies. Changing this forces a new resource
	// to be created.
	ObjectId string `pulumi:"objectId"`
	// List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
	SecretPermissions []string `pulumi:"secretPermissions"`
	// List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
	StoragePermissions []string `pulumi:"storagePermissions"`
	// The Azure Active Directory tenant ID that should be used
	// for authenticating requests to the key vault. Changing this forces a new resource
	// to be created.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AccessPolicy resource.
type AccessPolicyArgs struct {
	// The object ID of an Application in Azure Active Directory.
	ApplicationId pulumi.StringPtrInput
	// List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
	CertificatePermissions pulumi.StringArrayInput
	// List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
	KeyPermissions pulumi.StringArrayInput
	// Specifies the id of the Key Vault resource. Changing this
	// forces a new resource to be created.
	KeyVaultId pulumi.StringInput
	// The object ID of a user, service principal or security
	// group in the Azure Active Directory tenant for the vault. The object ID must
	// be unique for the list of access policies. Changing this forces a new resource
	// to be created.
	ObjectId pulumi.StringInput
	// List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
	SecretPermissions pulumi.StringArrayInput
	// List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
	StoragePermissions pulumi.StringArrayInput
	// The Azure Active Directory tenant ID that should be used
	// for authenticating requests to the key vault. Changing this forces a new resource
	// to be created.
	TenantId pulumi.StringInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}

type AccessPolicyInput interface {
	pulumi.Input

	ToAccessPolicyOutput() AccessPolicyOutput
	ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput
}

func (*AccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicy)(nil))
}

func (i *AccessPolicy) ToAccessPolicyOutput() AccessPolicyOutput {
	return i.ToAccessPolicyOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyOutput)
}

func (i *AccessPolicy) ToAccessPolicyPtrOutput() AccessPolicyPtrOutput {
	return i.ToAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyPtrOutputWithContext(ctx context.Context) AccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyPtrOutput)
}

type AccessPolicyPtrInput interface {
	pulumi.Input

	ToAccessPolicyPtrOutput() AccessPolicyPtrOutput
	ToAccessPolicyPtrOutputWithContext(ctx context.Context) AccessPolicyPtrOutput
}

type accessPolicyPtrType AccessPolicyArgs

func (*accessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil))
}

func (i *accessPolicyPtrType) ToAccessPolicyPtrOutput() AccessPolicyPtrOutput {
	return i.ToAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *accessPolicyPtrType) ToAccessPolicyPtrOutputWithContext(ctx context.Context) AccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyPtrOutput)
}

// AccessPolicyArrayInput is an input type that accepts AccessPolicyArray and AccessPolicyArrayOutput values.
// You can construct a concrete instance of `AccessPolicyArrayInput` via:
//
//          AccessPolicyArray{ AccessPolicyArgs{...} }
type AccessPolicyArrayInput interface {
	pulumi.Input

	ToAccessPolicyArrayOutput() AccessPolicyArrayOutput
	ToAccessPolicyArrayOutputWithContext(context.Context) AccessPolicyArrayOutput
}

type AccessPolicyArray []AccessPolicyInput

func (AccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AccessPolicy)(nil))
}

func (i AccessPolicyArray) ToAccessPolicyArrayOutput() AccessPolicyArrayOutput {
	return i.ToAccessPolicyArrayOutputWithContext(context.Background())
}

func (i AccessPolicyArray) ToAccessPolicyArrayOutputWithContext(ctx context.Context) AccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyArrayOutput)
}

// AccessPolicyMapInput is an input type that accepts AccessPolicyMap and AccessPolicyMapOutput values.
// You can construct a concrete instance of `AccessPolicyMapInput` via:
//
//          AccessPolicyMap{ "key": AccessPolicyArgs{...} }
type AccessPolicyMapInput interface {
	pulumi.Input

	ToAccessPolicyMapOutput() AccessPolicyMapOutput
	ToAccessPolicyMapOutputWithContext(context.Context) AccessPolicyMapOutput
}

type AccessPolicyMap map[string]AccessPolicyInput

func (AccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AccessPolicy)(nil))
}

func (i AccessPolicyMap) ToAccessPolicyMapOutput() AccessPolicyMapOutput {
	return i.ToAccessPolicyMapOutputWithContext(context.Background())
}

func (i AccessPolicyMap) ToAccessPolicyMapOutputWithContext(ctx context.Context) AccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyMapOutput)
}

type AccessPolicyOutput struct {
	*pulumi.OutputState
}

func (AccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicy)(nil))
}

func (o AccessPolicyOutput) ToAccessPolicyOutput() AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyPtrOutput() AccessPolicyPtrOutput {
	return o.ToAccessPolicyPtrOutputWithContext(context.Background())
}

func (o AccessPolicyOutput) ToAccessPolicyPtrOutputWithContext(ctx context.Context) AccessPolicyPtrOutput {
	return o.ApplyT(func(v AccessPolicy) *AccessPolicy {
		return &v
	}).(AccessPolicyPtrOutput)
}

type AccessPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (AccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil))
}

func (o AccessPolicyPtrOutput) ToAccessPolicyPtrOutput() AccessPolicyPtrOutput {
	return o
}

func (o AccessPolicyPtrOutput) ToAccessPolicyPtrOutputWithContext(ctx context.Context) AccessPolicyPtrOutput {
	return o
}

type AccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicy)(nil))
}

func (o AccessPolicyArrayOutput) ToAccessPolicyArrayOutput() AccessPolicyArrayOutput {
	return o
}

func (o AccessPolicyArrayOutput) ToAccessPolicyArrayOutputWithContext(ctx context.Context) AccessPolicyArrayOutput {
	return o
}

func (o AccessPolicyArrayOutput) Index(i pulumi.IntInput) AccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPolicy {
		return vs[0].([]AccessPolicy)[vs[1].(int)]
	}).(AccessPolicyOutput)
}

type AccessPolicyMapOutput struct{ *pulumi.OutputState }

func (AccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccessPolicy)(nil))
}

func (o AccessPolicyMapOutput) ToAccessPolicyMapOutput() AccessPolicyMapOutput {
	return o
}

func (o AccessPolicyMapOutput) ToAccessPolicyMapOutputWithContext(ctx context.Context) AccessPolicyMapOutput {
	return o
}

func (o AccessPolicyMapOutput) MapIndex(k pulumi.StringInput) AccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccessPolicy {
		return vs[0].(map[string]AccessPolicy)[vs[1].(string)]
	}).(AccessPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyOutput{})
	pulumi.RegisterOutputType(AccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(AccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyMapOutput{})
}
