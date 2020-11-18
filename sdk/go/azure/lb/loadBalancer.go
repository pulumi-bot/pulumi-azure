// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Load Balancer Resource.
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
// 		_, err = lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
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
// 		return nil
// 	})
// }
// ```
type LoadBalancer struct {
	pulumi.CustomResourceState

	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations LoadBalancerFrontendIpConfigurationArrayOutput `pulumi:"frontendIpConfigurations"`
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Load Balancer.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LoadBalancer
	err := ctx.RegisterResource("azure:lb/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("azure:lb/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations []LoadBalancerFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Load Balancer.
	Name *string `pulumi:"name"`
	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type LoadBalancerState struct {
	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations LoadBalancerFrontendIpConfigurationArrayInput
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Load Balancer.
	Name pulumi.StringPtrInput
	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	PrivateIpAddress pulumi.StringPtrInput
	// The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
	PrivateIpAddresses pulumi.StringArrayInput
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations []LoadBalancerFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Load Balancer.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations LoadBalancerFrontendIpConfigurationArrayInput
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Load Balancer.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

type LoadBalancerOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerOutput)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerOutput{})
}
