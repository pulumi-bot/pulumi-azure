// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Traffic Manager Endpoint.
//
// Deprecated: azure.trafficmanager.Endpoint has been deprecated in favor of azure.network.TrafficManagerEndpoint
type Endpoint struct {
	pulumi.CustomResourceState

	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeaderArrayOutput `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      pulumi.StringOutput `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringOutput `pulumi:"endpointMonitorStatus"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringOutput `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayOutput `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints pulumi.IntPtrOutput `pulumi:"minChildEndpoints"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets EndpointSubnetArrayOutput `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringOutput `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &EndpointArgs{}
	}
	var resource Endpoint
	err := ctx.RegisterResource("azure:trafficmanager/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure:trafficmanager/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders []EndpointCustomHeader `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      *string `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string `pulumi:"endpointMonitorStatus"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings []string `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints *int `pulumi:"minChildEndpoints"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority *int `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName *string `pulumi:"profileName"`
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets []EndpointSubnet `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target *string `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type *string `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight *int `pulumi:"weight"`
}

type EndpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeaderArrayInput
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      pulumi.StringPtrInput
	EndpointMonitorStatus pulumi.StringPtrInput
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringPtrInput
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayInput
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints pulumi.IntPtrInput
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntPtrInput
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringPtrInput
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `subnet` blocks as defined below
	Subnets EndpointSubnetArrayInput
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringPtrInput
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringPtrInput
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringPtrInput
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders []EndpointCustomHeader `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation *string `pulumi:"endpointLocation"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings []string `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints *int `pulumi:"minChildEndpoints"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority *int `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets []EndpointSubnet `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target *string `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type string `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeaderArrayInput
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation pulumi.StringPtrInput
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringPtrInput
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayInput
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints pulumi.IntPtrInput
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntPtrInput
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringInput
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName pulumi.StringInput
	// One or more `subnet` blocks as defined below
	Subnets EndpointSubnetArrayInput
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringPtrInput
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringPtrInput
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringInput
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}
