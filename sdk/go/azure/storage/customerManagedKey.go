// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Customer Managed Key for a Storage Account.
//
// ## Import
//
// Customer Managed Keys for a Storage Account can be imported using the `resource id` of the Storage Account, e.g.
//
// ```sh
//  $ pulumi import azure:storage/customerManagedKey:CustomerManagedKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
// ```
type CustomerManagedKey struct {
	pulumi.CustomResourceState

	// The name of Key Vault Key.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
	KeyVersion pulumi.StringPtrOutput `pulumi:"keyVersion"`
	// The ID of the Storage Account. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
}

// NewCustomerManagedKey registers a new resource with the given unique name, arguments, and options.
func NewCustomerManagedKey(ctx *pulumi.Context,
	name string, args *CustomerManagedKeyArgs, opts ...pulumi.ResourceOption) (*CustomerManagedKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource CustomerManagedKey
	err := ctx.RegisterResource("azure:storage/customerManagedKey:CustomerManagedKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerManagedKey gets an existing CustomerManagedKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerManagedKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerManagedKeyState, opts ...pulumi.ResourceOption) (*CustomerManagedKey, error) {
	var resource CustomerManagedKey
	err := ctx.ReadResource("azure:storage/customerManagedKey:CustomerManagedKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerManagedKey resources.
type customerManagedKeyState struct {
	// The name of Key Vault Key.
	KeyName *string `pulumi:"keyName"`
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
	KeyVersion *string `pulumi:"keyVersion"`
	// The ID of the Storage Account. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type CustomerManagedKeyState struct {
	// The name of Key Vault Key.
	KeyName pulumi.StringPtrInput
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringPtrInput
	// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
	KeyVersion pulumi.StringPtrInput
	// The ID of the Storage Account. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
}

func (CustomerManagedKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerManagedKeyState)(nil)).Elem()
}

type customerManagedKeyArgs struct {
	// The name of Key Vault Key.
	KeyName string `pulumi:"keyName"`
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
	KeyVersion *string `pulumi:"keyVersion"`
	// The ID of the Storage Account. Changing this forces a new resource to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a CustomerManagedKey resource.
type CustomerManagedKeyArgs struct {
	// The name of Key Vault Key.
	KeyName pulumi.StringInput
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringInput
	// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
	KeyVersion pulumi.StringPtrInput
	// The ID of the Storage Account. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringInput
}

func (CustomerManagedKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerManagedKeyArgs)(nil)).Elem()
}

type CustomerManagedKeyInput interface {
	pulumi.Input

	ToCustomerManagedKeyOutput() CustomerManagedKeyOutput
	ToCustomerManagedKeyOutputWithContext(ctx context.Context) CustomerManagedKeyOutput
}

type CustomerManagedKeyPtrInput interface {
	pulumi.Input

	ToCustomerManagedKeyPtrOutput() CustomerManagedKeyPtrOutput
	ToCustomerManagedKeyPtrOutputWithContext(ctx context.Context) CustomerManagedKeyPtrOutput
}

func (CustomerManagedKey) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKey)(nil)).Elem()
}

func (i CustomerManagedKey) ToCustomerManagedKeyOutput() CustomerManagedKeyOutput {
	return i.ToCustomerManagedKeyOutputWithContext(context.Background())
}

func (i CustomerManagedKey) ToCustomerManagedKeyOutputWithContext(ctx context.Context) CustomerManagedKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyOutput)
}

func (i CustomerManagedKey) ToCustomerManagedKeyPtrOutput() CustomerManagedKeyPtrOutput {
	return i.ToCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (i CustomerManagedKey) ToCustomerManagedKeyPtrOutputWithContext(ctx context.Context) CustomerManagedKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyPtrOutput)
}

type CustomerManagedKeyOutput struct {
	*pulumi.OutputState
}

func (CustomerManagedKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyOutput)(nil)).Elem()
}

func (o CustomerManagedKeyOutput) ToCustomerManagedKeyOutput() CustomerManagedKeyOutput {
	return o
}

func (o CustomerManagedKeyOutput) ToCustomerManagedKeyOutputWithContext(ctx context.Context) CustomerManagedKeyOutput {
	return o
}

type CustomerManagedKeyPtrOutput struct {
	*pulumi.OutputState
}

func (CustomerManagedKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKey)(nil)).Elem()
}

func (o CustomerManagedKeyPtrOutput) ToCustomerManagedKeyPtrOutput() CustomerManagedKeyPtrOutput {
	return o
}

func (o CustomerManagedKeyPtrOutput) ToCustomerManagedKeyPtrOutputWithContext(ctx context.Context) CustomerManagedKeyPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomerManagedKeyOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyPtrOutput{})
}
