// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Batch pool.
//
// ## Import
//
// Batch Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:batch/pool:Pool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Batch/batchAccounts/myBatchAccount1/pools/myBatchPool1
// ```
type Pool struct {
	pulumi.CustomResourceState

	// Specifies the name of the Batch account in which the pool will be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// A `autoScale` block that describes the scale settings when using auto scale.
	AutoScale PoolAutoScalePtrOutput `pulumi:"autoScale"`
	// One or more `certificate` blocks that describe the certificates to be installed on each compute node in the pool.
	Certificates PoolCertificateArrayOutput `pulumi:"certificates"`
	// The container configuration used in the pool's VMs.
	ContainerConfiguration PoolContainerConfigurationPtrOutput `pulumi:"containerConfiguration"`
	// Specifies the display name of the Batch pool.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// A `fixedScale` block that describes the scale settings when using fixed scale.
	FixedScale PoolFixedScalePtrOutput `pulumi:"fixedScale"`
	// Specifies the maximum number of tasks that can run concurrently on a single compute node in the pool. Defaults to `1`. Changing this forces a new resource to be created.
	MaxTasksPerNode pulumi.IntPtrOutput `pulumi:"maxTasksPerNode"`
	// A map of custom batch pool metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Specifies the name of the Batch pool. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkConfiguration` block that describes the network configurations for the Batch pool.
	NetworkConfiguration PoolNetworkConfigurationPtrOutput `pulumi:"networkConfiguration"`
	// Specifies the Sku of the node agents that will be created in the Batch pool.
	NodeAgentSkuId pulumi.StringOutput `pulumi:"nodeAgentSkuId"`
	// The name of the resource group in which to create the Batch pool. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `startTask` block that describes the start task settings for the Batch pool.
	StartTask                  PoolStartTaskPtrOutput `pulumi:"startTask"`
	StopPendingResizeOperation pulumi.BoolPtrOutput   `pulumi:"stopPendingResizeOperation"`
	// A `storageImageReference` for the virtual machines that will compose the Batch pool.
	StorageImageReference PoolStorageImageReferenceOutput `pulumi:"storageImageReference"`
	// Specifies the size of the VM created in the Batch pool.
	VmSize pulumi.StringOutput `pulumi:"vmSize"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.NodeAgentSkuId == nil {
		return nil, errors.New("invalid value for required argument 'NodeAgentSkuId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageImageReference == nil {
		return nil, errors.New("invalid value for required argument 'StorageImageReference'")
	}
	if args.VmSize == nil {
		return nil, errors.New("invalid value for required argument 'VmSize'")
	}
	var resource Pool
	err := ctx.RegisterResource("azure:batch/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure:batch/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// Specifies the name of the Batch account in which the pool will be created. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// A `autoScale` block that describes the scale settings when using auto scale.
	AutoScale *PoolAutoScale `pulumi:"autoScale"`
	// One or more `certificate` blocks that describe the certificates to be installed on each compute node in the pool.
	Certificates []PoolCertificate `pulumi:"certificates"`
	// The container configuration used in the pool's VMs.
	ContainerConfiguration *PoolContainerConfiguration `pulumi:"containerConfiguration"`
	// Specifies the display name of the Batch pool.
	DisplayName *string `pulumi:"displayName"`
	// A `fixedScale` block that describes the scale settings when using fixed scale.
	FixedScale *PoolFixedScale `pulumi:"fixedScale"`
	// Specifies the maximum number of tasks that can run concurrently on a single compute node in the pool. Defaults to `1`. Changing this forces a new resource to be created.
	MaxTasksPerNode *int `pulumi:"maxTasksPerNode"`
	// A map of custom batch pool metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies the name of the Batch pool. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkConfiguration` block that describes the network configurations for the Batch pool.
	NetworkConfiguration *PoolNetworkConfiguration `pulumi:"networkConfiguration"`
	// Specifies the Sku of the node agents that will be created in the Batch pool.
	NodeAgentSkuId *string `pulumi:"nodeAgentSkuId"`
	// The name of the resource group in which to create the Batch pool. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `startTask` block that describes the start task settings for the Batch pool.
	StartTask                  *PoolStartTask `pulumi:"startTask"`
	StopPendingResizeOperation *bool          `pulumi:"stopPendingResizeOperation"`
	// A `storageImageReference` for the virtual machines that will compose the Batch pool.
	StorageImageReference *PoolStorageImageReference `pulumi:"storageImageReference"`
	// Specifies the size of the VM created in the Batch pool.
	VmSize *string `pulumi:"vmSize"`
}

type PoolState struct {
	// Specifies the name of the Batch account in which the pool will be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// A `autoScale` block that describes the scale settings when using auto scale.
	AutoScale PoolAutoScalePtrInput
	// One or more `certificate` blocks that describe the certificates to be installed on each compute node in the pool.
	Certificates PoolCertificateArrayInput
	// The container configuration used in the pool's VMs.
	ContainerConfiguration PoolContainerConfigurationPtrInput
	// Specifies the display name of the Batch pool.
	DisplayName pulumi.StringPtrInput
	// A `fixedScale` block that describes the scale settings when using fixed scale.
	FixedScale PoolFixedScalePtrInput
	// Specifies the maximum number of tasks that can run concurrently on a single compute node in the pool. Defaults to `1`. Changing this forces a new resource to be created.
	MaxTasksPerNode pulumi.IntPtrInput
	// A map of custom batch pool metadata.
	Metadata pulumi.StringMapInput
	// Specifies the name of the Batch pool. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkConfiguration` block that describes the network configurations for the Batch pool.
	NetworkConfiguration PoolNetworkConfigurationPtrInput
	// Specifies the Sku of the node agents that will be created in the Batch pool.
	NodeAgentSkuId pulumi.StringPtrInput
	// The name of the resource group in which to create the Batch pool. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `startTask` block that describes the start task settings for the Batch pool.
	StartTask                  PoolStartTaskPtrInput
	StopPendingResizeOperation pulumi.BoolPtrInput
	// A `storageImageReference` for the virtual machines that will compose the Batch pool.
	StorageImageReference PoolStorageImageReferencePtrInput
	// Specifies the size of the VM created in the Batch pool.
	VmSize pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Specifies the name of the Batch account in which the pool will be created. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// A `autoScale` block that describes the scale settings when using auto scale.
	AutoScale *PoolAutoScale `pulumi:"autoScale"`
	// One or more `certificate` blocks that describe the certificates to be installed on each compute node in the pool.
	Certificates []PoolCertificate `pulumi:"certificates"`
	// The container configuration used in the pool's VMs.
	ContainerConfiguration *PoolContainerConfiguration `pulumi:"containerConfiguration"`
	// Specifies the display name of the Batch pool.
	DisplayName *string `pulumi:"displayName"`
	// A `fixedScale` block that describes the scale settings when using fixed scale.
	FixedScale *PoolFixedScale `pulumi:"fixedScale"`
	// Specifies the maximum number of tasks that can run concurrently on a single compute node in the pool. Defaults to `1`. Changing this forces a new resource to be created.
	MaxTasksPerNode *int `pulumi:"maxTasksPerNode"`
	// A map of custom batch pool metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies the name of the Batch pool. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkConfiguration` block that describes the network configurations for the Batch pool.
	NetworkConfiguration *PoolNetworkConfiguration `pulumi:"networkConfiguration"`
	// Specifies the Sku of the node agents that will be created in the Batch pool.
	NodeAgentSkuId string `pulumi:"nodeAgentSkuId"`
	// The name of the resource group in which to create the Batch pool. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `startTask` block that describes the start task settings for the Batch pool.
	StartTask                  *PoolStartTask `pulumi:"startTask"`
	StopPendingResizeOperation *bool          `pulumi:"stopPendingResizeOperation"`
	// A `storageImageReference` for the virtual machines that will compose the Batch pool.
	StorageImageReference PoolStorageImageReference `pulumi:"storageImageReference"`
	// Specifies the size of the VM created in the Batch pool.
	VmSize string `pulumi:"vmSize"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Specifies the name of the Batch account in which the pool will be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// A `autoScale` block that describes the scale settings when using auto scale.
	AutoScale PoolAutoScalePtrInput
	// One or more `certificate` blocks that describe the certificates to be installed on each compute node in the pool.
	Certificates PoolCertificateArrayInput
	// The container configuration used in the pool's VMs.
	ContainerConfiguration PoolContainerConfigurationPtrInput
	// Specifies the display name of the Batch pool.
	DisplayName pulumi.StringPtrInput
	// A `fixedScale` block that describes the scale settings when using fixed scale.
	FixedScale PoolFixedScalePtrInput
	// Specifies the maximum number of tasks that can run concurrently on a single compute node in the pool. Defaults to `1`. Changing this forces a new resource to be created.
	MaxTasksPerNode pulumi.IntPtrInput
	// A map of custom batch pool metadata.
	Metadata pulumi.StringMapInput
	// Specifies the name of the Batch pool. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkConfiguration` block that describes the network configurations for the Batch pool.
	NetworkConfiguration PoolNetworkConfigurationPtrInput
	// Specifies the Sku of the node agents that will be created in the Batch pool.
	NodeAgentSkuId pulumi.StringInput
	// The name of the resource group in which to create the Batch pool. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `startTask` block that describes the start task settings for the Batch pool.
	StartTask                  PoolStartTaskPtrInput
	StopPendingResizeOperation pulumi.BoolPtrInput
	// A `storageImageReference` for the virtual machines that will compose the Batch pool.
	StorageImageReference PoolStorageImageReferenceInput
	// Specifies the size of the VM created in the Batch pool.
	VmSize pulumi.StringInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

func (i *Pool) ToPoolPtrOutput() PoolPtrOutput {
	return i.ToPoolPtrOutputWithContext(context.Background())
}

func (i *Pool) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPtrOutput)
}

type PoolPtrInput interface {
	pulumi.Input

	ToPoolPtrOutput() PoolPtrOutput
	ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput
}

type poolPtrType PoolArgs

func (*poolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil))
}

func (i *poolPtrType) ToPoolPtrOutput() PoolPtrOutput {
	return i.ToPoolPtrOutputWithContext(context.Background())
}

func (i *poolPtrType) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPtrOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//          PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Pool)(nil))
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//          PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Pool)(nil))
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct {
	*pulumi.OutputState
}

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func (o PoolOutput) ToPoolPtrOutput() PoolPtrOutput {
	return o.ToPoolPtrOutputWithContext(context.Background())
}

func (o PoolOutput) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return o.ApplyT(func(v Pool) *Pool {
		return &v
	}).(PoolPtrOutput)
}

type PoolPtrOutput struct {
	*pulumi.OutputState
}

func (PoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil))
}

func (o PoolPtrOutput) ToPoolPtrOutput() PoolPtrOutput {
	return o
}

func (o PoolPtrOutput) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return o
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pool)(nil))
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pool {
		return vs[0].([]Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pool)(nil))
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pool {
		return vs[0].(map[string]Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolPtrOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}
