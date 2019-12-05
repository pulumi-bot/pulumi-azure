// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub Route
// 
// > **NOTE:** Routes can be defined either directly on the `iot.IoTHub` resource, or using the `iot.Route` resourcs - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iothub_route.html.markdown.
type Route struct {
	s *pulumi.ResourceState
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOpt) (*Route, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.EndpointNames == nil {
		return nil, errors.New("missing required argument 'EndpointNames'")
	}
	if args == nil || args.IothubName == nil {
		return nil, errors.New("missing required argument 'IothubName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["enabled"] = nil
		inputs["endpointNames"] = nil
		inputs["iothubName"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["source"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["enabled"] = args.Enabled
		inputs["endpointNames"] = args.EndpointNames
		inputs["iothubName"] = args.IothubName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["source"] = args.Source
	}
	s, err := ctx.RegisterResource("azure:iot/route:Route", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Route{s: s}, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteState, opts ...pulumi.ResourceOpt) (*Route, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["enabled"] = state.Enabled
		inputs["endpointNames"] = state.EndpointNames
		inputs["iothubName"] = state.IothubName
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["source"] = state.Source
	}
	s, err := ctx.ReadResource("azure:iot/route:Route", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Route{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Route) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Route) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
func (r *Route) Condition() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["condition"])
}

// Specifies whether a route is enabled.
func (r *Route) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
func (r *Route) EndpointNames() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["endpointNames"])
}

// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
func (r *Route) IothubName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["iothubName"])
}

// The name of the route.
func (r *Route) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
func (r *Route) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
func (r *Route) Source() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["source"])
}

// Input properties used for looking up and filtering Route resources.
type RouteState struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition interface{}
	// Specifies whether a route is enabled.
	Enabled interface{}
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames interface{}
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName interface{}
	// The name of the route.
	Name interface{}
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source interface{}
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition interface{}
	// Specifies whether a route is enabled.
	Enabled interface{}
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames interface{}
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName interface{}
	// The name of the route.
	Name interface{}
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source interface{}
}