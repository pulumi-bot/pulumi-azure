// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Load Balancer Outbound Rule.
//
// > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration and a Backend Address Pool Attached.
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
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
// 		exampleBackendAddressPool, err := lb.NewBackendAddressPool(ctx, "exampleBackendAddressPool", &lb.BackendAddressPoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			LoadbalancerId:    exampleLoadBalancer.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewOutboundRule(ctx, "exampleOutboundRule", &lb.OutboundRuleArgs{
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			LoadbalancerId:       exampleLoadBalancer.ID(),
// 			Protocol:             pulumi.String("Tcp"),
// 			BackendAddressPoolId: exampleBackendAddressPool.ID(),
// 			FrontendIpConfigurations: lb.OutboundRuleFrontendIpConfigurationArray{
// 				&lb.OutboundRuleFrontendIpConfigurationArgs{
// 					Name: pulumi.String("PublicIPAddress"),
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
//
// ## Import
//
// Load Balancer Outbound Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:lb/outboundRule:OutboundRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/outboundRules/rule1
// ```
type OutboundRule struct {
	pulumi.CustomResourceState

	// The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts pulumi.IntPtrOutput `pulumi:"allocatedOutboundPorts"`
	// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
	BackendAddressPoolId pulumi.StringOutput `pulumi:"backendAddressPoolId"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset pulumi.BoolPtrOutput `pulumi:"enableTcpReset"`
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations OutboundRuleFrontendIpConfigurationArrayOutput `pulumi:"frontendIpConfigurations"`
	// The timeout for the TCP idle connection
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewOutboundRule registers a new resource with the given unique name, arguments, and options.
func NewOutboundRule(ctx *pulumi.Context,
	name string, args *OutboundRuleArgs, opts ...pulumi.ResourceOption) (*OutboundRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendAddressPoolId == nil {
		return nil, errors.New("invalid value for required argument 'BackendAddressPoolId'")
	}
	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource OutboundRule
	err := ctx.RegisterResource("azure:lb/outboundRule:OutboundRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutboundRule gets an existing OutboundRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutboundRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutboundRuleState, opts ...pulumi.ResourceOption) (*OutboundRule, error) {
	var resource OutboundRule
	err := ctx.ReadResource("azure:lb/outboundRule:OutboundRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutboundRule resources.
type outboundRuleState struct {
	// The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts *int `pulumi:"allocatedOutboundPorts"`
	// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
	BackendAddressPoolId *string `pulumi:"backendAddressPoolId"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations []OutboundRuleFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// The timeout for the TCP idle connection
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type OutboundRuleState struct {
	// The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts pulumi.IntPtrInput
	// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
	BackendAddressPoolId pulumi.StringPtrInput
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset pulumi.BoolPtrInput
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations OutboundRuleFrontendIpConfigurationArrayInput
	// The timeout for the TCP idle connection
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (OutboundRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundRuleState)(nil)).Elem()
}

type outboundRuleArgs struct {
	// The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts *int `pulumi:"allocatedOutboundPorts"`
	// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
	BackendAddressPoolId string `pulumi:"backendAddressPoolId"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations []OutboundRuleFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// The timeout for the TCP idle connection
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a OutboundRule resource.
type OutboundRuleArgs struct {
	// The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts pulumi.IntPtrInput
	// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
	BackendAddressPoolId pulumi.StringInput
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset pulumi.BoolPtrInput
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations OutboundRuleFrontendIpConfigurationArrayInput
	// The timeout for the TCP idle connection
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol pulumi.StringInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (OutboundRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundRuleArgs)(nil)).Elem()
}

type OutboundRuleInput interface {
	pulumi.Input

	ToOutboundRuleOutput() OutboundRuleOutput
	ToOutboundRuleOutputWithContext(ctx context.Context) OutboundRuleOutput
}

func (*OutboundRule) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundRule)(nil))
}

func (i *OutboundRule) ToOutboundRuleOutput() OutboundRuleOutput {
	return i.ToOutboundRuleOutputWithContext(context.Background())
}

func (i *OutboundRule) ToOutboundRuleOutputWithContext(ctx context.Context) OutboundRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundRuleOutput)
}

func (i *OutboundRule) ToOutboundRulePtrOutput() OutboundRulePtrOutput {
	return i.ToOutboundRulePtrOutputWithContext(context.Background())
}

func (i *OutboundRule) ToOutboundRulePtrOutputWithContext(ctx context.Context) OutboundRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundRulePtrOutput)
}

type OutboundRulePtrInput interface {
	pulumi.Input

	ToOutboundRulePtrOutput() OutboundRulePtrOutput
	ToOutboundRulePtrOutputWithContext(ctx context.Context) OutboundRulePtrOutput
}

type outboundRulePtrType OutboundRuleArgs

func (*outboundRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundRule)(nil))
}

func (i *outboundRulePtrType) ToOutboundRulePtrOutput() OutboundRulePtrOutput {
	return i.ToOutboundRulePtrOutputWithContext(context.Background())
}

func (i *outboundRulePtrType) ToOutboundRulePtrOutputWithContext(ctx context.Context) OutboundRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundRulePtrOutput)
}

// OutboundRuleArrayInput is an input type that accepts OutboundRuleArray and OutboundRuleArrayOutput values.
// You can construct a concrete instance of `OutboundRuleArrayInput` via:
//
//          OutboundRuleArray{ OutboundRuleArgs{...} }
type OutboundRuleArrayInput interface {
	pulumi.Input

	ToOutboundRuleArrayOutput() OutboundRuleArrayOutput
	ToOutboundRuleArrayOutputWithContext(context.Context) OutboundRuleArrayOutput
}

type OutboundRuleArray []OutboundRuleInput

func (OutboundRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OutboundRule)(nil))
}

func (i OutboundRuleArray) ToOutboundRuleArrayOutput() OutboundRuleArrayOutput {
	return i.ToOutboundRuleArrayOutputWithContext(context.Background())
}

func (i OutboundRuleArray) ToOutboundRuleArrayOutputWithContext(ctx context.Context) OutboundRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundRuleArrayOutput)
}

// OutboundRuleMapInput is an input type that accepts OutboundRuleMap and OutboundRuleMapOutput values.
// You can construct a concrete instance of `OutboundRuleMapInput` via:
//
//          OutboundRuleMap{ "key": OutboundRuleArgs{...} }
type OutboundRuleMapInput interface {
	pulumi.Input

	ToOutboundRuleMapOutput() OutboundRuleMapOutput
	ToOutboundRuleMapOutputWithContext(context.Context) OutboundRuleMapOutput
}

type OutboundRuleMap map[string]OutboundRuleInput

func (OutboundRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OutboundRule)(nil))
}

func (i OutboundRuleMap) ToOutboundRuleMapOutput() OutboundRuleMapOutput {
	return i.ToOutboundRuleMapOutputWithContext(context.Background())
}

func (i OutboundRuleMap) ToOutboundRuleMapOutputWithContext(ctx context.Context) OutboundRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundRuleMapOutput)
}

type OutboundRuleOutput struct {
	*pulumi.OutputState
}

func (OutboundRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundRule)(nil))
}

func (o OutboundRuleOutput) ToOutboundRuleOutput() OutboundRuleOutput {
	return o
}

func (o OutboundRuleOutput) ToOutboundRuleOutputWithContext(ctx context.Context) OutboundRuleOutput {
	return o
}

func (o OutboundRuleOutput) ToOutboundRulePtrOutput() OutboundRulePtrOutput {
	return o.ToOutboundRulePtrOutputWithContext(context.Background())
}

func (o OutboundRuleOutput) ToOutboundRulePtrOutputWithContext(ctx context.Context) OutboundRulePtrOutput {
	return o.ApplyT(func(v OutboundRule) *OutboundRule {
		return &v
	}).(OutboundRulePtrOutput)
}

type OutboundRulePtrOutput struct {
	*pulumi.OutputState
}

func (OutboundRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundRule)(nil))
}

func (o OutboundRulePtrOutput) ToOutboundRulePtrOutput() OutboundRulePtrOutput {
	return o
}

func (o OutboundRulePtrOutput) ToOutboundRulePtrOutputWithContext(ctx context.Context) OutboundRulePtrOutput {
	return o
}

type OutboundRuleArrayOutput struct{ *pulumi.OutputState }

func (OutboundRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundRule)(nil))
}

func (o OutboundRuleArrayOutput) ToOutboundRuleArrayOutput() OutboundRuleArrayOutput {
	return o
}

func (o OutboundRuleArrayOutput) ToOutboundRuleArrayOutputWithContext(ctx context.Context) OutboundRuleArrayOutput {
	return o
}

func (o OutboundRuleArrayOutput) Index(i pulumi.IntInput) OutboundRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutboundRule {
		return vs[0].([]OutboundRule)[vs[1].(int)]
	}).(OutboundRuleOutput)
}

type OutboundRuleMapOutput struct{ *pulumi.OutputState }

func (OutboundRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutboundRule)(nil))
}

func (o OutboundRuleMapOutput) ToOutboundRuleMapOutput() OutboundRuleMapOutput {
	return o
}

func (o OutboundRuleMapOutput) ToOutboundRuleMapOutputWithContext(ctx context.Context) OutboundRuleMapOutput {
	return o
}

func (o OutboundRuleMapOutput) MapIndex(k pulumi.StringInput) OutboundRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutboundRule {
		return vs[0].(map[string]OutboundRule)[vs[1].(string)]
	}).(OutboundRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(OutboundRuleOutput{})
	pulumi.RegisterOutputType(OutboundRulePtrOutput{})
	pulumi.RegisterOutputType(OutboundRuleArrayOutput{})
	pulumi.RegisterOutputType(OutboundRuleMapOutput{})
}
