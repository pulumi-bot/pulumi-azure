// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Key Vault Secret.
type Secret struct {
	pulumi.CustomResourceState

	// Specifies the content type for the Key Vault Secret.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// The ID of the Key Vault where the Secret should be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate pulumi.StringPtrOutput `pulumi:"notBeforeDate"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the value of the Key Vault Secret.
	Value pulumi.StringOutput `pulumi:"value"`
	// The current version of the Key Vault Secret.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource Secret
	err := ctx.RegisterResource("azure:keyvault/secret:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("azure:keyvault/secret:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	// Specifies the content type for the Key Vault Secret.
	ContentType *string `pulumi:"contentType"`
	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate *string `pulumi:"expirationDate"`
	// The ID of the Key Vault where the Secret should be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate *string `pulumi:"notBeforeDate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the value of the Key Vault Secret.
	Value *string `pulumi:"value"`
	// The current version of the Key Vault Secret.
	Version *string `pulumi:"version"`
}

type SecretState struct {
	// Specifies the content type for the Key Vault Secret.
	ContentType pulumi.StringPtrInput
	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate pulumi.StringPtrInput
	// The ID of the Key Vault where the Secret should be created.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the value of the Key Vault Secret.
	Value pulumi.StringPtrInput
	// The current version of the Key Vault Secret.
	Version pulumi.StringPtrInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// Specifies the content type for the Key Vault Secret.
	ContentType *string `pulumi:"contentType"`
	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate *string `pulumi:"expirationDate"`
	// The ID of the Key Vault where the Secret should be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate *string `pulumi:"notBeforeDate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the value of the Key Vault Secret.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// Specifies the content type for the Key Vault Secret.
	ContentType pulumi.StringPtrInput
	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate pulumi.StringPtrInput
	// The ID of the Key Vault where the Secret should be created.
	KeyVaultId pulumi.StringInput
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the value of the Key Vault Secret.
	Value pulumi.StringInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}
