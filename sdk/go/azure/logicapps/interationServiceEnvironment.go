// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages private and isolated Logic App instances within an Azure virtual network.
type InterationServiceEnvironment struct {
	pulumi.CustomResourceState

	// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
	AccessEndpointType pulumi.StringOutput `pulumi:"accessEndpointType"`
	// The list of access endpoint ip addresses of connector.
	ConnectorEndpointIpAddresses pulumi.StringArrayOutput `pulumi:"connectorEndpointIpAddresses"`
	// The list of outgoing ip addresses of connector.
	ConnectorOutboundIpAddresses pulumi.StringArrayOutput `pulumi:"connectorOutboundIpAddresses"`
	// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
	SkuName pulumi.StringPtrOutput `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Integration Service Environment.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
	VirtualNetworkSubnetIds pulumi.StringArrayOutput `pulumi:"virtualNetworkSubnetIds"`
	// The list of access endpoint ip addresses of workflow.
	WorkflowEndpointIpAddresses pulumi.StringArrayOutput `pulumi:"workflowEndpointIpAddresses"`
	// The list of outgoing ip addresses of workflow.
	WorkflowOutboundIpAddresses pulumi.StringArrayOutput `pulumi:"workflowOutboundIpAddresses"`
}

// NewInterationServiceEnvironment registers a new resource with the given unique name, arguments, and options.
func NewInterationServiceEnvironment(ctx *pulumi.Context,
	name string, args *InterationServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*InterationServiceEnvironment, error) {
	if args == nil || args.AccessEndpointType == nil {
		return nil, errors.New("missing required argument 'AccessEndpointType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkSubnetIds == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkSubnetIds'")
	}
	if args == nil {
		args = &InterationServiceEnvironmentArgs{}
	}
	var resource InterationServiceEnvironment
	err := ctx.RegisterResource("azure:logicapps/interationServiceEnvironment:InterationServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterationServiceEnvironment gets an existing InterationServiceEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterationServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterationServiceEnvironmentState, opts ...pulumi.ResourceOption) (*InterationServiceEnvironment, error) {
	var resource InterationServiceEnvironment
	err := ctx.ReadResource("azure:logicapps/interationServiceEnvironment:InterationServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterationServiceEnvironment resources.
type interationServiceEnvironmentState struct {
	// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
	AccessEndpointType *string `pulumi:"accessEndpointType"`
	// The list of access endpoint ip addresses of connector.
	ConnectorEndpointIpAddresses []string `pulumi:"connectorEndpointIpAddresses"`
	// The list of outgoing ip addresses of connector.
	ConnectorOutboundIpAddresses []string `pulumi:"connectorOutboundIpAddresses"`
	// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	Location *string `pulumi:"location"`
	// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Integration Service Environment.
	Tags map[string]string `pulumi:"tags"`
	// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
	VirtualNetworkSubnetIds []string `pulumi:"virtualNetworkSubnetIds"`
	// The list of access endpoint ip addresses of workflow.
	WorkflowEndpointIpAddresses []string `pulumi:"workflowEndpointIpAddresses"`
	// The list of outgoing ip addresses of workflow.
	WorkflowOutboundIpAddresses []string `pulumi:"workflowOutboundIpAddresses"`
}

type InterationServiceEnvironmentState struct {
	// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
	AccessEndpointType pulumi.StringPtrInput
	// The list of access endpoint ip addresses of connector.
	ConnectorEndpointIpAddresses pulumi.StringArrayInput
	// The list of outgoing ip addresses of connector.
	ConnectorOutboundIpAddresses pulumi.StringArrayInput
	// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	Location pulumi.StringPtrInput
	// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
	SkuName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Integration Service Environment.
	Tags pulumi.StringMapInput
	// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
	VirtualNetworkSubnetIds pulumi.StringArrayInput
	// The list of access endpoint ip addresses of workflow.
	WorkflowEndpointIpAddresses pulumi.StringArrayInput
	// The list of outgoing ip addresses of workflow.
	WorkflowOutboundIpAddresses pulumi.StringArrayInput
}

func (InterationServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*interationServiceEnvironmentState)(nil)).Elem()
}

type interationServiceEnvironmentArgs struct {
	// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
	AccessEndpointType string `pulumi:"accessEndpointType"`
	// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	Location *string `pulumi:"location"`
	// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Integration Service Environment.
	Tags map[string]string `pulumi:"tags"`
	// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
	VirtualNetworkSubnetIds []string `pulumi:"virtualNetworkSubnetIds"`
}

// The set of arguments for constructing a InterationServiceEnvironment resource.
type InterationServiceEnvironmentArgs struct {
	// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
	AccessEndpointType pulumi.StringInput
	// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	Location pulumi.StringPtrInput
	// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
	ResourceGroupName pulumi.StringInput
	// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
	SkuName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Integration Service Environment.
	Tags pulumi.StringMapInput
	// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
	VirtualNetworkSubnetIds pulumi.StringArrayInput
}

func (InterationServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interationServiceEnvironmentArgs)(nil)).Elem()
}
