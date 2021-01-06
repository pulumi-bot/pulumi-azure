// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Customer Managed Key for a Kusto Cluster.
//
// ## Import
//
// Customer Managed Keys for a Kusto Cluster can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/clusterCustomerManagedKey:ClusterCustomerManagedKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
// ```
type ClusterCustomerManagedKey struct {
	pulumi.CustomResourceState

	// The ID of the Kusto Cluster. Changing this forces a new resource to be created.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name of Key Vault Key.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// The version of Key Vault Key.
	KeyVersion pulumi.StringOutput `pulumi:"keyVersion"`
}

// NewClusterCustomerManagedKey registers a new resource with the given unique name, arguments, and options.
func NewClusterCustomerManagedKey(ctx *pulumi.Context,
	name string, args *ClusterCustomerManagedKeyArgs, opts ...pulumi.ResourceOption) (*ClusterCustomerManagedKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.KeyVersion == nil {
		return nil, errors.New("invalid value for required argument 'KeyVersion'")
	}
	var resource ClusterCustomerManagedKey
	err := ctx.RegisterResource("azure:kusto/clusterCustomerManagedKey:ClusterCustomerManagedKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterCustomerManagedKey gets an existing ClusterCustomerManagedKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterCustomerManagedKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterCustomerManagedKeyState, opts ...pulumi.ResourceOption) (*ClusterCustomerManagedKey, error) {
	var resource ClusterCustomerManagedKey
	err := ctx.ReadResource("azure:kusto/clusterCustomerManagedKey:ClusterCustomerManagedKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterCustomerManagedKey resources.
type clusterCustomerManagedKeyState struct {
	// The ID of the Kusto Cluster. Changing this forces a new resource to be created.
	ClusterId *string `pulumi:"clusterId"`
	// The name of Key Vault Key.
	KeyName *string `pulumi:"keyName"`
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// The version of Key Vault Key.
	KeyVersion *string `pulumi:"keyVersion"`
}

type ClusterCustomerManagedKeyState struct {
	// The ID of the Kusto Cluster. Changing this forces a new resource to be created.
	ClusterId pulumi.StringPtrInput
	// The name of Key Vault Key.
	KeyName pulumi.StringPtrInput
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringPtrInput
	// The version of Key Vault Key.
	KeyVersion pulumi.StringPtrInput
}

func (ClusterCustomerManagedKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterCustomerManagedKeyState)(nil)).Elem()
}

type clusterCustomerManagedKeyArgs struct {
	// The ID of the Kusto Cluster. Changing this forces a new resource to be created.
	ClusterId string `pulumi:"clusterId"`
	// The name of Key Vault Key.
	KeyName string `pulumi:"keyName"`
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The version of Key Vault Key.
	KeyVersion string `pulumi:"keyVersion"`
}

// The set of arguments for constructing a ClusterCustomerManagedKey resource.
type ClusterCustomerManagedKeyArgs struct {
	// The ID of the Kusto Cluster. Changing this forces a new resource to be created.
	ClusterId pulumi.StringInput
	// The name of Key Vault Key.
	KeyName pulumi.StringInput
	// The ID of the Key Vault. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringInput
	// The version of Key Vault Key.
	KeyVersion pulumi.StringInput
}

func (ClusterCustomerManagedKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterCustomerManagedKeyArgs)(nil)).Elem()
}

type ClusterCustomerManagedKeyInput interface {
	pulumi.Input

	ToClusterCustomerManagedKeyOutput() ClusterCustomerManagedKeyOutput
	ToClusterCustomerManagedKeyOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyOutput
}

func (*ClusterCustomerManagedKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCustomerManagedKey)(nil))
}

func (i *ClusterCustomerManagedKey) ToClusterCustomerManagedKeyOutput() ClusterCustomerManagedKeyOutput {
	return i.ToClusterCustomerManagedKeyOutputWithContext(context.Background())
}

func (i *ClusterCustomerManagedKey) ToClusterCustomerManagedKeyOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCustomerManagedKeyOutput)
}

func (i *ClusterCustomerManagedKey) ToClusterCustomerManagedKeyPtrOutput() ClusterCustomerManagedKeyPtrOutput {
	return i.ToClusterCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (i *ClusterCustomerManagedKey) ToClusterCustomerManagedKeyPtrOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCustomerManagedKeyPtrOutput)
}

type ClusterCustomerManagedKeyPtrInput interface {
	pulumi.Input

	ToClusterCustomerManagedKeyPtrOutput() ClusterCustomerManagedKeyPtrOutput
	ToClusterCustomerManagedKeyPtrOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyPtrOutput
}

type clusterCustomerManagedKeyPtrType ClusterCustomerManagedKeyArgs

func (*clusterCustomerManagedKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCustomerManagedKey)(nil))
}

func (i *clusterCustomerManagedKeyPtrType) ToClusterCustomerManagedKeyPtrOutput() ClusterCustomerManagedKeyPtrOutput {
	return i.ToClusterCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (i *clusterCustomerManagedKeyPtrType) ToClusterCustomerManagedKeyPtrOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCustomerManagedKeyOutput).ToClusterCustomerManagedKeyPtrOutput()
}

// ClusterCustomerManagedKeyArrayInput is an input type that accepts ClusterCustomerManagedKeyArray and ClusterCustomerManagedKeyArrayOutput values.
// You can construct a concrete instance of `ClusterCustomerManagedKeyArrayInput` via:
//
//          ClusterCustomerManagedKeyArray{ ClusterCustomerManagedKeyArgs{...} }
type ClusterCustomerManagedKeyArrayInput interface {
	pulumi.Input

	ToClusterCustomerManagedKeyArrayOutput() ClusterCustomerManagedKeyArrayOutput
	ToClusterCustomerManagedKeyArrayOutputWithContext(context.Context) ClusterCustomerManagedKeyArrayOutput
}

type ClusterCustomerManagedKeyArray []ClusterCustomerManagedKeyInput

func (ClusterCustomerManagedKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterCustomerManagedKey)(nil))
}

func (i ClusterCustomerManagedKeyArray) ToClusterCustomerManagedKeyArrayOutput() ClusterCustomerManagedKeyArrayOutput {
	return i.ToClusterCustomerManagedKeyArrayOutputWithContext(context.Background())
}

func (i ClusterCustomerManagedKeyArray) ToClusterCustomerManagedKeyArrayOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCustomerManagedKeyArrayOutput)
}

// ClusterCustomerManagedKeyMapInput is an input type that accepts ClusterCustomerManagedKeyMap and ClusterCustomerManagedKeyMapOutput values.
// You can construct a concrete instance of `ClusterCustomerManagedKeyMapInput` via:
//
//          ClusterCustomerManagedKeyMap{ "key": ClusterCustomerManagedKeyArgs{...} }
type ClusterCustomerManagedKeyMapInput interface {
	pulumi.Input

	ToClusterCustomerManagedKeyMapOutput() ClusterCustomerManagedKeyMapOutput
	ToClusterCustomerManagedKeyMapOutputWithContext(context.Context) ClusterCustomerManagedKeyMapOutput
}

type ClusterCustomerManagedKeyMap map[string]ClusterCustomerManagedKeyInput

func (ClusterCustomerManagedKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterCustomerManagedKey)(nil))
}

func (i ClusterCustomerManagedKeyMap) ToClusterCustomerManagedKeyMapOutput() ClusterCustomerManagedKeyMapOutput {
	return i.ToClusterCustomerManagedKeyMapOutputWithContext(context.Background())
}

func (i ClusterCustomerManagedKeyMap) ToClusterCustomerManagedKeyMapOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCustomerManagedKeyMapOutput)
}

type ClusterCustomerManagedKeyOutput struct {
	*pulumi.OutputState
}

func (ClusterCustomerManagedKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCustomerManagedKey)(nil))
}

func (o ClusterCustomerManagedKeyOutput) ToClusterCustomerManagedKeyOutput() ClusterCustomerManagedKeyOutput {
	return o
}

func (o ClusterCustomerManagedKeyOutput) ToClusterCustomerManagedKeyOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyOutput {
	return o
}

func (o ClusterCustomerManagedKeyOutput) ToClusterCustomerManagedKeyPtrOutput() ClusterCustomerManagedKeyPtrOutput {
	return o.ToClusterCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (o ClusterCustomerManagedKeyOutput) ToClusterCustomerManagedKeyPtrOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyPtrOutput {
	return o.ApplyT(func(v ClusterCustomerManagedKey) *ClusterCustomerManagedKey {
		return &v
	}).(ClusterCustomerManagedKeyPtrOutput)
}

type ClusterCustomerManagedKeyPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterCustomerManagedKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCustomerManagedKey)(nil))
}

func (o ClusterCustomerManagedKeyPtrOutput) ToClusterCustomerManagedKeyPtrOutput() ClusterCustomerManagedKeyPtrOutput {
	return o
}

func (o ClusterCustomerManagedKeyPtrOutput) ToClusterCustomerManagedKeyPtrOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyPtrOutput {
	return o
}

type ClusterCustomerManagedKeyArrayOutput struct{ *pulumi.OutputState }

func (ClusterCustomerManagedKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterCustomerManagedKey)(nil))
}

func (o ClusterCustomerManagedKeyArrayOutput) ToClusterCustomerManagedKeyArrayOutput() ClusterCustomerManagedKeyArrayOutput {
	return o
}

func (o ClusterCustomerManagedKeyArrayOutput) ToClusterCustomerManagedKeyArrayOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyArrayOutput {
	return o
}

func (o ClusterCustomerManagedKeyArrayOutput) Index(i pulumi.IntInput) ClusterCustomerManagedKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterCustomerManagedKey {
		return vs[0].([]ClusterCustomerManagedKey)[vs[1].(int)]
	}).(ClusterCustomerManagedKeyOutput)
}

type ClusterCustomerManagedKeyMapOutput struct{ *pulumi.OutputState }

func (ClusterCustomerManagedKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterCustomerManagedKey)(nil))
}

func (o ClusterCustomerManagedKeyMapOutput) ToClusterCustomerManagedKeyMapOutput() ClusterCustomerManagedKeyMapOutput {
	return o
}

func (o ClusterCustomerManagedKeyMapOutput) ToClusterCustomerManagedKeyMapOutputWithContext(ctx context.Context) ClusterCustomerManagedKeyMapOutput {
	return o
}

func (o ClusterCustomerManagedKeyMapOutput) MapIndex(k pulumi.StringInput) ClusterCustomerManagedKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterCustomerManagedKey {
		return vs[0].(map[string]ClusterCustomerManagedKey)[vs[1].(string)]
	}).(ClusterCustomerManagedKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterCustomerManagedKeyOutput{})
	pulumi.RegisterOutputType(ClusterCustomerManagedKeyPtrOutput{})
	pulumi.RegisterOutputType(ClusterCustomerManagedKeyArrayOutput{})
	pulumi.RegisterOutputType(ClusterCustomerManagedKeyMapOutput{})
}
