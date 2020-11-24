// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a LoadBalancer Probe Resource.
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
// 		_, err = lb.NewProbe(ctx, "exampleProbe", &lb.ProbeArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			LoadbalancerId:    exampleLoadBalancer.ID(),
// 			Port:              pulumi.Int(22),
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
// Load Balancer Probes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:lb/probe:Probe example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1
// ```
type Probe struct {
	pulumi.CustomResourceState

	// The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
	IntervalInSeconds pulumi.IntPtrOutput      `pulumi:"intervalInSeconds"`
	LoadBalancerRules pulumi.StringArrayOutput `pulumi:"loadBalancerRules"`
	// The ID of the LoadBalancer in which to create the NAT Rule.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the Probe.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
	NumberOfProbes pulumi.IntPtrOutput `pulumi:"numberOfProbes"`
	// Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
	Port pulumi.IntOutput `pulumi:"port"`
	// Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
	RequestPath pulumi.StringPtrOutput `pulumi:"requestPath"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProbe registers a new resource with the given unique name, arguments, and options.
func NewProbe(ctx *pulumi.Context,
	name string, args *ProbeArgs, opts ...pulumi.ResourceOption) (*Probe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Probe
	err := ctx.RegisterResource("azure:lb/probe:Probe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProbe gets an existing Probe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProbe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProbeState, opts ...pulumi.ResourceOption) (*Probe, error) {
	var resource Probe
	err := ctx.ReadResource("azure:lb/probe:Probe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Probe resources.
type probeState struct {
	// The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
	IntervalInSeconds *int     `pulumi:"intervalInSeconds"`
	LoadBalancerRules []string `pulumi:"loadBalancerRules"`
	// The ID of the LoadBalancer in which to create the NAT Rule.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the Probe.
	Name *string `pulumi:"name"`
	// The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
	NumberOfProbes *int `pulumi:"numberOfProbes"`
	// Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
	Port *int `pulumi:"port"`
	// Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
	Protocol *string `pulumi:"protocol"`
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
	RequestPath *string `pulumi:"requestPath"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ProbeState struct {
	// The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
	IntervalInSeconds pulumi.IntPtrInput
	LoadBalancerRules pulumi.StringArrayInput
	// The ID of the LoadBalancer in which to create the NAT Rule.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the Probe.
	Name pulumi.StringPtrInput
	// The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
	NumberOfProbes pulumi.IntPtrInput
	// Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
	Port pulumi.IntPtrInput
	// Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
	Protocol pulumi.StringPtrInput
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
	RequestPath pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringPtrInput
}

func (ProbeState) ElementType() reflect.Type {
	return reflect.TypeOf((*probeState)(nil)).Elem()
}

type probeArgs struct {
	// The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The ID of the LoadBalancer in which to create the NAT Rule.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the Probe.
	Name *string `pulumi:"name"`
	// The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
	NumberOfProbes *int `pulumi:"numberOfProbes"`
	// Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
	Port int `pulumi:"port"`
	// Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
	Protocol *string `pulumi:"protocol"`
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
	RequestPath *string `pulumi:"requestPath"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Probe resource.
type ProbeArgs struct {
	// The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
	IntervalInSeconds pulumi.IntPtrInput
	// The ID of the LoadBalancer in which to create the NAT Rule.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the Probe.
	Name pulumi.StringPtrInput
	// The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
	NumberOfProbes pulumi.IntPtrInput
	// Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
	Port pulumi.IntInput
	// Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
	Protocol pulumi.StringPtrInput
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to `Http` or `Https`. Otherwise, it is not allowed.
	RequestPath pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
}

func (ProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*probeArgs)(nil)).Elem()
}

type ProbeInput interface {
	pulumi.Input

	ToProbeOutput() ProbeOutput
	ToProbeOutputWithContext(ctx context.Context) ProbeOutput
}

func (Probe) ElementType() reflect.Type {
	return reflect.TypeOf((*Probe)(nil)).Elem()
}

func (i Probe) ToProbeOutput() ProbeOutput {
	return i.ToProbeOutputWithContext(context.Background())
}

func (i Probe) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeOutput)
}

type ProbeOutput struct {
	*pulumi.OutputState
}

func (ProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeOutput)(nil)).Elem()
}

func (o ProbeOutput) ToProbeOutput() ProbeOutput {
	return o
}

func (o ProbeOutput) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProbeOutput{})
}
