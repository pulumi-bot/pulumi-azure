// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Disk Encryption Set.
//
// > **NOTE:** At this time the Key Vault used to store the Active Key for this Disk Encryption Set must have both Soft Delete & Purge Protection enabled - which are not yet supported by this provider.
//
// ## Import
//
// Disk Encryption Sets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/diskEncryptionSet:DiskEncryptionSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/diskEncryptionSets/encryptionSet1
// ```
type DiskEncryptionSet struct {
	pulumi.CustomResourceState

	// A `identity` block defined below.
	Identity DiskEncryptionSetIdentityOutput `pulumi:"identity"`
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId pulumi.StringOutput `pulumi:"keyVaultKeyId"`
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDiskEncryptionSet registers a new resource with the given unique name, arguments, and options.
func NewDiskEncryptionSet(ctx *pulumi.Context,
	name string, args *DiskEncryptionSetArgs, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.KeyVaultKeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultKeyId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DiskEncryptionSet
	err := ctx.RegisterResource("azure:compute/diskEncryptionSet:DiskEncryptionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskEncryptionSet gets an existing DiskEncryptionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskEncryptionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskEncryptionSetState, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	var resource DiskEncryptionSet
	err := ctx.ReadResource("azure:compute/diskEncryptionSet:DiskEncryptionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskEncryptionSet resources.
type diskEncryptionSetState struct {
	// A `identity` block defined below.
	Identity *DiskEncryptionSetIdentity `pulumi:"identity"`
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId *string `pulumi:"keyVaultKeyId"`
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags map[string]string `pulumi:"tags"`
}

type DiskEncryptionSetState struct {
	// A `identity` block defined below.
	Identity DiskEncryptionSetIdentityPtrInput
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId pulumi.StringPtrInput
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags pulumi.StringMapInput
}

func (DiskEncryptionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetState)(nil)).Elem()
}

type diskEncryptionSetArgs struct {
	// A `identity` block defined below.
	Identity DiskEncryptionSetIdentity `pulumi:"identity"`
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId string `pulumi:"keyVaultKeyId"`
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DiskEncryptionSet resource.
type DiskEncryptionSetArgs struct {
	// A `identity` block defined below.
	Identity DiskEncryptionSetIdentityInput
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId pulumi.StringInput
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags pulumi.StringMapInput
}

func (DiskEncryptionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetArgs)(nil)).Elem()
}

type DiskEncryptionSetInput interface {
	pulumi.Input

	ToDiskEncryptionSetOutput() DiskEncryptionSetOutput
	ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput
}

type DiskEncryptionSetPtrInput interface {
	pulumi.Input

	ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput
	ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput
}

func (DiskEncryptionSet) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSet)(nil)).Elem()
}

func (i DiskEncryptionSet) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return i.ToDiskEncryptionSetOutputWithContext(context.Background())
}

func (i DiskEncryptionSet) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetOutput)
}

func (i DiskEncryptionSet) ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput {
	return i.ToDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionSet) ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetPtrOutput)
}

type DiskEncryptionSetOutput struct {
	*pulumi.OutputState
}

func (DiskEncryptionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetOutput)(nil)).Elem()
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return o
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return o
}

type DiskEncryptionSetPtrOutput struct {
	*pulumi.OutputState
}

func (DiskEncryptionSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil)).Elem()
}

func (o DiskEncryptionSetPtrOutput) ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput {
	return o
}

func (o DiskEncryptionSetPtrOutput) ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DiskEncryptionSetOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetPtrOutput{})
}
