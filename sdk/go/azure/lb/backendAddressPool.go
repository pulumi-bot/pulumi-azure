// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Load Balancer Backend Address Pool.
//
// > **NOTE:** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/lb"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FrontendIpConfigurations: lb.LoadBalancerFrontendIpConfigurationArray{
// 				&lb.LoadBalancerFrontendIpConfigurationArgs{
// 					Name:              pulumi.String("PublicIPAddress"),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewBackendAddressPool(ctx, "exampleBackendAddressPool", &lb.BackendAddressPoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			LoadbalancerId:    exampleLoadBalancer.ID(),
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
// Load Balancer Backend Address Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:lb/backendAddressPool:BackendAddressPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1
// ```
type BackendAddressPool struct {
	pulumi.CustomResourceState

	// The Backend IP Configurations associated with this Backend Address Pool.
	BackendIpConfigurations pulumi.StringArrayOutput `pulumi:"backendIpConfigurations"`
	// The Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules pulumi.StringArrayOutput `pulumi:"loadBalancingRules"`
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewBackendAddressPool registers a new resource with the given unique name, arguments, and options.
func NewBackendAddressPool(ctx *pulumi.Context,
	name string, args *BackendAddressPoolArgs, opts ...pulumi.ResourceOption) (*BackendAddressPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource BackendAddressPool
	err := ctx.RegisterResource("azure:lb/backendAddressPool:BackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendAddressPool gets an existing BackendAddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendAddressPoolState, opts ...pulumi.ResourceOption) (*BackendAddressPool, error) {
	var resource BackendAddressPool
	err := ctx.ReadResource("azure:lb/backendAddressPool:BackendAddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendAddressPool resources.
type backendAddressPoolState struct {
	// The Backend IP Configurations associated with this Backend Address Pool.
	BackendIpConfigurations []string `pulumi:"backendIpConfigurations"`
	// The Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules []string `pulumi:"loadBalancingRules"`
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type BackendAddressPoolState struct {
	// The Backend IP Configurations associated with this Backend Address Pool.
	BackendIpConfigurations pulumi.StringArrayInput
	// The Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules pulumi.StringArrayInput
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the Backend Address Pool.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringPtrInput
}

func (BackendAddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendAddressPoolState)(nil)).Elem()
}

type backendAddressPoolArgs struct {
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a BackendAddressPool resource.
type BackendAddressPoolArgs struct {
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the Backend Address Pool.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
}

func (BackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendAddressPoolArgs)(nil)).Elem()
}

type BackendAddressPoolInput interface {
	pulumi.Input

	ToBackendAddressPoolOutput() BackendAddressPoolOutput
	ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput
}

func (*BackendAddressPool) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPool)(nil))
}

func (i *BackendAddressPool) ToBackendAddressPoolOutput() BackendAddressPoolOutput {
	return i.ToBackendAddressPoolOutputWithContext(context.Background())
}

func (i *BackendAddressPool) ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolOutput)
}

func (i *BackendAddressPool) ToBackendAddressPoolPtrOutput() BackendAddressPoolPtrOutput {
	return i.ToBackendAddressPoolPtrOutputWithContext(context.Background())
}

func (i *BackendAddressPool) ToBackendAddressPoolPtrOutputWithContext(ctx context.Context) BackendAddressPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolPtrOutput)
}

type BackendAddressPoolPtrInput interface {
	pulumi.Input

	ToBackendAddressPoolPtrOutput() BackendAddressPoolPtrOutput
	ToBackendAddressPoolPtrOutputWithContext(ctx context.Context) BackendAddressPoolPtrOutput
}

type backendAddressPoolPtrType BackendAddressPoolArgs

func (*backendAddressPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAddressPool)(nil))
}

func (i *backendAddressPoolPtrType) ToBackendAddressPoolPtrOutput() BackendAddressPoolPtrOutput {
	return i.ToBackendAddressPoolPtrOutputWithContext(context.Background())
}

func (i *backendAddressPoolPtrType) ToBackendAddressPoolPtrOutputWithContext(ctx context.Context) BackendAddressPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolPtrOutput)
}

// BackendAddressPoolArrayInput is an input type that accepts BackendAddressPoolArray and BackendAddressPoolArrayOutput values.
// You can construct a concrete instance of `BackendAddressPoolArrayInput` via:
//
//          BackendAddressPoolArray{ BackendAddressPoolArgs{...} }
type BackendAddressPoolArrayInput interface {
	pulumi.Input

	ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput
	ToBackendAddressPoolArrayOutputWithContext(context.Context) BackendAddressPoolArrayOutput
}

type BackendAddressPoolArray []BackendAddressPoolInput

func (BackendAddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BackendAddressPool)(nil))
}

func (i BackendAddressPoolArray) ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput {
	return i.ToBackendAddressPoolArrayOutputWithContext(context.Background())
}

func (i BackendAddressPoolArray) ToBackendAddressPoolArrayOutputWithContext(ctx context.Context) BackendAddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolArrayOutput)
}

// BackendAddressPoolMapInput is an input type that accepts BackendAddressPoolMap and BackendAddressPoolMapOutput values.
// You can construct a concrete instance of `BackendAddressPoolMapInput` via:
//
//          BackendAddressPoolMap{ "key": BackendAddressPoolArgs{...} }
type BackendAddressPoolMapInput interface {
	pulumi.Input

	ToBackendAddressPoolMapOutput() BackendAddressPoolMapOutput
	ToBackendAddressPoolMapOutputWithContext(context.Context) BackendAddressPoolMapOutput
}

type BackendAddressPoolMap map[string]BackendAddressPoolInput

func (BackendAddressPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BackendAddressPool)(nil))
}

func (i BackendAddressPoolMap) ToBackendAddressPoolMapOutput() BackendAddressPoolMapOutput {
	return i.ToBackendAddressPoolMapOutputWithContext(context.Background())
}

func (i BackendAddressPoolMap) ToBackendAddressPoolMapOutputWithContext(ctx context.Context) BackendAddressPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolMapOutput)
}

type BackendAddressPoolOutput struct {
	*pulumi.OutputState
}

func (BackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPool)(nil))
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolOutput() BackendAddressPoolOutput {
	return o
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput {
	return o
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolPtrOutput() BackendAddressPoolPtrOutput {
	return o.ToBackendAddressPoolPtrOutputWithContext(context.Background())
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolPtrOutputWithContext(ctx context.Context) BackendAddressPoolPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *BackendAddressPool {
		return &v
	}).(BackendAddressPoolPtrOutput)
}

type BackendAddressPoolPtrOutput struct {
	*pulumi.OutputState
}

func (BackendAddressPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAddressPool)(nil))
}

func (o BackendAddressPoolPtrOutput) ToBackendAddressPoolPtrOutput() BackendAddressPoolPtrOutput {
	return o
}

func (o BackendAddressPoolPtrOutput) ToBackendAddressPoolPtrOutputWithContext(ctx context.Context) BackendAddressPoolPtrOutput {
	return o
}

type BackendAddressPoolArrayOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPool)(nil))
}

func (o BackendAddressPoolArrayOutput) ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput {
	return o
}

func (o BackendAddressPoolArrayOutput) ToBackendAddressPoolArrayOutputWithContext(ctx context.Context) BackendAddressPoolArrayOutput {
	return o
}

func (o BackendAddressPoolArrayOutput) Index(i pulumi.IntInput) BackendAddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendAddressPool {
		return vs[0].([]BackendAddressPool)[vs[1].(int)]
	}).(BackendAddressPoolOutput)
}

type BackendAddressPoolMapOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackendAddressPool)(nil))
}

func (o BackendAddressPoolMapOutput) ToBackendAddressPoolMapOutput() BackendAddressPoolMapOutput {
	return o
}

func (o BackendAddressPoolMapOutput) ToBackendAddressPoolMapOutputWithContext(ctx context.Context) BackendAddressPoolMapOutput {
	return o
}

func (o BackendAddressPoolMapOutput) MapIndex(k pulumi.StringInput) BackendAddressPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackendAddressPool {
		return vs[0].(map[string]BackendAddressPool)[vs[1].(string)]
	}).(BackendAddressPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendAddressPoolOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolPtrOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolArrayOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolMapOutput{})
}
