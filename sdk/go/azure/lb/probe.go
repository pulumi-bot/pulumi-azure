// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a LoadBalancer Probe Resource.
//
// > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
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
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.
	RequestPath pulumi.StringPtrOutput `pulumi:"requestPath"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProbe registers a new resource with the given unique name, arguments, and options.
func NewProbe(ctx *pulumi.Context,
	name string, args *ProbeArgs, opts ...pulumi.ResourceOption) (*Probe, error) {
	if args == nil || args.LoadbalancerId == nil {
		return nil, errors.New("missing required argument 'LoadbalancerId'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ProbeArgs{}
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
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.
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
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.
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
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.
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
	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.
	RequestPath pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
}

func (ProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*probeArgs)(nil)).Elem()
}
