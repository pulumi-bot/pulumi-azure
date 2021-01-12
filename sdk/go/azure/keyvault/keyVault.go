// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Key Vault's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:keyvault/keyVault:KeyVault example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/vault1
// ```
type KeyVault struct {
	pulumi.CustomResourceState

	// A list of up to 16 objects describing access policies, as described below.
	AccessPolicies KeyVaultAccessPolicyArrayOutput `pulumi:"accessPolicies"`
	// One or more `contact` block as defined below.
	Contacts KeyVaultContactArrayOutput `pulumi:"contacts"`
	// Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
	EnableRbacAuthorization pulumi.BoolPtrOutput `pulumi:"enableRbacAuthorization"`
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment pulumi.BoolPtrOutput `pulumi:"enabledForDeployment"`
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption pulumi.BoolPtrOutput `pulumi:"enabledForDiskEncryption"`
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment pulumi.BoolPtrOutput `pulumi:"enabledForTemplateDeployment"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls KeyVaultNetworkAclsOutput `pulumi:"networkAcls"`
	// Is Purge Protection enabled for this Key Vault? Defaults to `false`.
	PurgeProtectionEnabled pulumi.BoolPtrOutput `pulumi:"purgeProtectionEnabled"`
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// Deprecated: Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
	SoftDeleteEnabled pulumi.BoolOutput `pulumi:"softDeleteEnabled"`
	// The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
	SoftDeleteRetentionDays pulumi.IntPtrOutput `pulumi:"softDeleteRetentionDays"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultUri pulumi.StringOutput `pulumi:"vaultUri"`
}

// NewKeyVault registers a new resource with the given unique name, arguments, and options.
func NewKeyVault(ctx *pulumi.Context,
	name string, args *KeyVaultArgs, opts ...pulumi.ResourceOption) (*KeyVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource KeyVault
	err := ctx.RegisterResource("azure:keyvault/keyVault:KeyVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyVault gets an existing KeyVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyVaultState, opts ...pulumi.ResourceOption) (*KeyVault, error) {
	var resource KeyVault
	err := ctx.ReadResource("azure:keyvault/keyVault:KeyVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyVault resources.
type keyVaultState struct {
	// A list of up to 16 objects describing access policies, as described below.
	AccessPolicies []KeyVaultAccessPolicy `pulumi:"accessPolicies"`
	// One or more `contact` block as defined below.
	Contacts []KeyVaultContact `pulumi:"contacts"`
	// Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
	EnableRbacAuthorization *bool `pulumi:"enableRbacAuthorization"`
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment *bool `pulumi:"enabledForDeployment"`
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption *bool `pulumi:"enabledForDiskEncryption"`
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment *bool `pulumi:"enabledForTemplateDeployment"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls *KeyVaultNetworkAcls `pulumi:"networkAcls"`
	// Is Purge Protection enabled for this Key Vault? Defaults to `false`.
	PurgeProtectionEnabled *bool `pulumi:"purgeProtectionEnabled"`
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName *string `pulumi:"skuName"`
	// Deprecated: Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
	SoftDeleteEnabled *bool `pulumi:"softDeleteEnabled"`
	// The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
	SoftDeleteRetentionDays *int `pulumi:"softDeleteRetentionDays"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId *string `pulumi:"tenantId"`
	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultUri *string `pulumi:"vaultUri"`
}

type KeyVaultState struct {
	// A list of up to 16 objects describing access policies, as described below.
	AccessPolicies KeyVaultAccessPolicyArrayInput
	// One or more `contact` block as defined below.
	Contacts KeyVaultContactArrayInput
	// Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
	EnableRbacAuthorization pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkAcls` block as defined below.
	NetworkAcls KeyVaultNetworkAclsPtrInput
	// Is Purge Protection enabled for this Key Vault? Defaults to `false`.
	PurgeProtectionEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName pulumi.StringPtrInput
	// Deprecated: Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
	SoftDeleteEnabled pulumi.BoolPtrInput
	// The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
	SoftDeleteRetentionDays pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringPtrInput
	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultUri pulumi.StringPtrInput
}

func (KeyVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyVaultState)(nil)).Elem()
}

type keyVaultArgs struct {
	// A list of up to 16 objects describing access policies, as described below.
	AccessPolicies []KeyVaultAccessPolicy `pulumi:"accessPolicies"`
	// One or more `contact` block as defined below.
	Contacts []KeyVaultContact `pulumi:"contacts"`
	// Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
	EnableRbacAuthorization *bool `pulumi:"enableRbacAuthorization"`
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment *bool `pulumi:"enabledForDeployment"`
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption *bool `pulumi:"enabledForDiskEncryption"`
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment *bool `pulumi:"enabledForTemplateDeployment"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls *KeyVaultNetworkAcls `pulumi:"networkAcls"`
	// Is Purge Protection enabled for this Key Vault? Defaults to `false`.
	PurgeProtectionEnabled *bool `pulumi:"purgeProtectionEnabled"`
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName string `pulumi:"skuName"`
	// Deprecated: Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
	SoftDeleteEnabled *bool `pulumi:"softDeleteEnabled"`
	// The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
	SoftDeleteRetentionDays *int `pulumi:"softDeleteRetentionDays"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a KeyVault resource.
type KeyVaultArgs struct {
	// A list of up to 16 objects describing access policies, as described below.
	AccessPolicies KeyVaultAccessPolicyArrayInput
	// One or more `contact` block as defined below.
	Contacts KeyVaultContactArrayInput
	// Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
	EnableRbacAuthorization pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkAcls` block as defined below.
	NetworkAcls KeyVaultNetworkAclsPtrInput
	// Is Purge Protection enabled for this Key Vault? Defaults to `false`.
	PurgeProtectionEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName pulumi.StringInput
	// Deprecated: Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
	SoftDeleteEnabled pulumi.BoolPtrInput
	// The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
	SoftDeleteRetentionDays pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringInput
}

func (KeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyVaultArgs)(nil)).Elem()
}

type KeyVaultInput interface {
	pulumi.Input

	ToKeyVaultOutput() KeyVaultOutput
	ToKeyVaultOutputWithContext(ctx context.Context) KeyVaultOutput
}

func (*KeyVault) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVault)(nil))
}

func (i *KeyVault) ToKeyVaultOutput() KeyVaultOutput {
	return i.ToKeyVaultOutputWithContext(context.Background())
}

func (i *KeyVault) ToKeyVaultOutputWithContext(ctx context.Context) KeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultOutput)
}

func (i *KeyVault) ToKeyVaultPtrOutput() KeyVaultPtrOutput {
	return i.ToKeyVaultPtrOutputWithContext(context.Background())
}

func (i *KeyVault) ToKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPtrOutput)
}

type KeyVaultPtrInput interface {
	pulumi.Input

	ToKeyVaultPtrOutput() KeyVaultPtrOutput
	ToKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultPtrOutput
}

type keyVaultPtrType KeyVaultArgs

func (*keyVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVault)(nil))
}

func (i *keyVaultPtrType) ToKeyVaultPtrOutput() KeyVaultPtrOutput {
	return i.ToKeyVaultPtrOutputWithContext(context.Background())
}

func (i *keyVaultPtrType) ToKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultOutput).ToKeyVaultPtrOutput()
}

// KeyVaultArrayInput is an input type that accepts KeyVaultArray and KeyVaultArrayOutput values.
// You can construct a concrete instance of `KeyVaultArrayInput` via:
//
//          KeyVaultArray{ KeyVaultArgs{...} }
type KeyVaultArrayInput interface {
	pulumi.Input

	ToKeyVaultArrayOutput() KeyVaultArrayOutput
	ToKeyVaultArrayOutputWithContext(context.Context) KeyVaultArrayOutput
}

type KeyVaultArray []KeyVaultInput

func (KeyVaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVault)(nil))
}

func (i KeyVaultArray) ToKeyVaultArrayOutput() KeyVaultArrayOutput {
	return i.ToKeyVaultArrayOutputWithContext(context.Background())
}

func (i KeyVaultArray) ToKeyVaultArrayOutputWithContext(ctx context.Context) KeyVaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultArrayOutput)
}

// KeyVaultMapInput is an input type that accepts KeyVaultMap and KeyVaultMapOutput values.
// You can construct a concrete instance of `KeyVaultMapInput` via:
//
//          KeyVaultMap{ "key": KeyVaultArgs{...} }
type KeyVaultMapInput interface {
	pulumi.Input

	ToKeyVaultMapOutput() KeyVaultMapOutput
	ToKeyVaultMapOutputWithContext(context.Context) KeyVaultMapOutput
}

type KeyVaultMap map[string]KeyVaultInput

func (KeyVaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyVault)(nil))
}

func (i KeyVaultMap) ToKeyVaultMapOutput() KeyVaultMapOutput {
	return i.ToKeyVaultMapOutputWithContext(context.Background())
}

func (i KeyVaultMap) ToKeyVaultMapOutputWithContext(ctx context.Context) KeyVaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultMapOutput)
}

type KeyVaultOutput struct {
	*pulumi.OutputState
}

func (KeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVault)(nil))
}

func (o KeyVaultOutput) ToKeyVaultOutput() KeyVaultOutput {
	return o
}

func (o KeyVaultOutput) ToKeyVaultOutputWithContext(ctx context.Context) KeyVaultOutput {
	return o
}

func (o KeyVaultOutput) ToKeyVaultPtrOutput() KeyVaultPtrOutput {
	return o.ToKeyVaultPtrOutputWithContext(context.Background())
}

func (o KeyVaultOutput) ToKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultPtrOutput {
	return o.ApplyT(func(v KeyVault) *KeyVault {
		return &v
	}).(KeyVaultPtrOutput)
}

type KeyVaultPtrOutput struct {
	*pulumi.OutputState
}

func (KeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVault)(nil))
}

func (o KeyVaultPtrOutput) ToKeyVaultPtrOutput() KeyVaultPtrOutput {
	return o
}

func (o KeyVaultPtrOutput) ToKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultPtrOutput {
	return o
}

type KeyVaultArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVault)(nil))
}

func (o KeyVaultArrayOutput) ToKeyVaultArrayOutput() KeyVaultArrayOutput {
	return o
}

func (o KeyVaultArrayOutput) ToKeyVaultArrayOutputWithContext(ctx context.Context) KeyVaultArrayOutput {
	return o
}

func (o KeyVaultArrayOutput) Index(i pulumi.IntInput) KeyVaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVault {
		return vs[0].([]KeyVault)[vs[1].(int)]
	}).(KeyVaultOutput)
}

type KeyVaultMapOutput struct{ *pulumi.OutputState }

func (KeyVaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyVault)(nil))
}

func (o KeyVaultMapOutput) ToKeyVaultMapOutput() KeyVaultMapOutput {
	return o
}

func (o KeyVaultMapOutput) ToKeyVaultMapOutputWithContext(ctx context.Context) KeyVaultMapOutput {
	return o
}

func (o KeyVaultMapOutput) MapIndex(k pulumi.StringInput) KeyVaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyVault {
		return vs[0].(map[string]KeyVault)[vs[1].(string)]
	}).(KeyVaultOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultMapOutput{})
}
