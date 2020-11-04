// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Load Balancer Rule.
//
// > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
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
// 		_, err = lb.NewRule(ctx, "exampleRule", &lb.RuleArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			LoadbalancerId:              exampleLoadBalancer.ID(),
// 			Protocol:                    pulumi.String("Tcp"),
// 			FrontendPort:                pulumi.Int(3389),
// 			BackendPort:                 pulumi.Int(3389),
// 			FrontendIpConfigurationName: pulumi.String("PublicIPAddress"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Rule struct {
	pulumi.CustomResourceState

	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId pulumi.StringOutput `pulumi:"backendAddressPoolId"`
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort pulumi.IntOutput `pulumi:"backendPort"`
	// Is snat enabled for this Load Balancer Rule? Default `false`.
	DisableOutboundSnat pulumi.BoolPtrOutput `pulumi:"disableOutboundSnat"`
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp pulumi.BoolPtrOutput `pulumi:"enableFloatingIp"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset            pulumi.BoolPtrOutput `pulumi:"enableTcpReset"`
	FrontendIpConfigurationId pulumi.StringOutput  `pulumi:"frontendIpConfigurationId"`
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName pulumi.StringOutput `pulumi:"frontendIpConfigurationName"`
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort pulumi.IntOutput `pulumi:"frontendPort"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes pulumi.IntOutput `pulumi:"idleTimeoutInMinutes"`
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution pulumi.StringOutput `pulumi:"loadDistribution"`
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the LB Rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId pulumi.StringOutput `pulumi:"probeId"`
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.BackendPort == nil {
		return nil, errors.New("invalid value for required argument 'BackendPort'")
	}
	if args.FrontendIpConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'FrontendIpConfigurationName'")
	}
	if args.FrontendPort == nil {
		return nil, errors.New("invalid value for required argument 'FrontendPort'")
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
	var resource Rule
	err := ctx.RegisterResource("azure:lb/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("azure:lb/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId *string `pulumi:"backendAddressPoolId"`
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort *int `pulumi:"backendPort"`
	// Is snat enabled for this Load Balancer Rule? Default `false`.
	DisableOutboundSnat *bool `pulumi:"disableOutboundSnat"`
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp *bool `pulumi:"enableFloatingIp"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset            *bool   `pulumi:"enableTcpReset"`
	FrontendIpConfigurationId *string `pulumi:"frontendIpConfigurationId"`
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName *string `pulumi:"frontendIpConfigurationName"`
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort *int `pulumi:"frontendPort"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution *string `pulumi:"loadDistribution"`
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the LB Rule.
	Name *string `pulumi:"name"`
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId *string `pulumi:"probeId"`
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type RuleState struct {
	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId pulumi.StringPtrInput
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort pulumi.IntPtrInput
	// Is snat enabled for this Load Balancer Rule? Default `false`.
	DisableOutboundSnat pulumi.BoolPtrInput
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp pulumi.BoolPtrInput
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset            pulumi.BoolPtrInput
	FrontendIpConfigurationId pulumi.StringPtrInput
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName pulumi.StringPtrInput
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort pulumi.IntPtrInput
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution pulumi.StringPtrInput
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the LB Rule.
	Name pulumi.StringPtrInput
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId *string `pulumi:"backendAddressPoolId"`
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort int `pulumi:"backendPort"`
	// Is snat enabled for this Load Balancer Rule? Default `false`.
	DisableOutboundSnat *bool `pulumi:"disableOutboundSnat"`
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp *bool `pulumi:"enableFloatingIp"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName string `pulumi:"frontendIpConfigurationName"`
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort int `pulumi:"frontendPort"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution *string `pulumi:"loadDistribution"`
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the LB Rule.
	Name *string `pulumi:"name"`
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId *string `pulumi:"probeId"`
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId pulumi.StringPtrInput
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort pulumi.IntInput
	// Is snat enabled for this Load Balancer Rule? Default `false`.
	DisableOutboundSnat pulumi.BoolPtrInput
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp pulumi.BoolPtrInput
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset pulumi.BoolPtrInput
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName pulumi.StringInput
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort pulumi.IntInput
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution pulumi.StringPtrInput
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the LB Rule.
	Name pulumi.StringPtrInput
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol pulumi.StringInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}
