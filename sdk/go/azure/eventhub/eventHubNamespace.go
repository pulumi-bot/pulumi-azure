// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EventHub Namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
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
// 		_, err = eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Capacity:          pulumi.Int(2),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// EventHub Namespaces can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/eventHubNamespace:EventHubNamespace namespace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1
// ```
type EventHubNamespace struct {
	pulumi.CustomResourceState

	// Is Auto Inflate enabled for the EventHub Namespace?
	AutoInflateEnabled pulumi.BoolPtrOutput `pulumi:"autoInflateEnabled"`
	// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Default capacity has a maximum of `20`, but can be increased in blocks of 20 on a committed purchase basis.
	Capacity pulumi.IntPtrOutput `pulumi:"capacity"`
	// Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
	DedicatedClusterId pulumi.StringPtrOutput `pulumi:"dedicatedClusterId"`
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString pulumi.StringOutput `pulumi:"defaultPrimaryConnectionString"`
	// The alias of the primary connection string for the authorization
	// rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
	DefaultPrimaryConnectionStringAlias pulumi.StringOutput `pulumi:"defaultPrimaryConnectionStringAlias"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey pulumi.StringOutput `pulumi:"defaultPrimaryKey"`
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString pulumi.StringOutput `pulumi:"defaultSecondaryConnectionString"`
	// The alias of the secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
	DefaultSecondaryConnectionStringAlias pulumi.StringOutput `pulumi:"defaultSecondaryConnectionStringAlias"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey pulumi.StringOutput `pulumi:"defaultSecondaryKey"`
	// An `identity` block as defined below.
	Identity EventHubNamespaceIdentityPtrOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
	MaximumThroughputUnits pulumi.IntOutput `pulumi:"maximumThroughputUnits"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkRulesets` block as defined below.
	NetworkRulesets EventHubNamespaceNetworkRulesetsOutput `pulumi:"networkRulesets"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Defines which tier to use. Valid options are `Basic` and `Standard`.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewEventHubNamespace registers a new resource with the given unique name, arguments, and options.
func NewEventHubNamespace(ctx *pulumi.Context,
	name string, args *EventHubNamespaceArgs, opts ...pulumi.ResourceOption) (*EventHubNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource EventHubNamespace
	err := ctx.RegisterResource("azure:eventhub/eventHubNamespace:EventHubNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubNamespace gets an existing EventHubNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubNamespaceState, opts ...pulumi.ResourceOption) (*EventHubNamespace, error) {
	var resource EventHubNamespace
	err := ctx.ReadResource("azure:eventhub/eventHubNamespace:EventHubNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubNamespace resources.
type eventHubNamespaceState struct {
	// Is Auto Inflate enabled for the EventHub Namespace?
	AutoInflateEnabled *bool `pulumi:"autoInflateEnabled"`
	// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Default capacity has a maximum of `20`, but can be increased in blocks of 20 on a committed purchase basis.
	Capacity *int `pulumi:"capacity"`
	// Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
	DedicatedClusterId *string `pulumi:"dedicatedClusterId"`
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString *string `pulumi:"defaultPrimaryConnectionString"`
	// The alias of the primary connection string for the authorization
	// rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
	DefaultPrimaryConnectionStringAlias *string `pulumi:"defaultPrimaryConnectionStringAlias"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey *string `pulumi:"defaultPrimaryKey"`
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString *string `pulumi:"defaultSecondaryConnectionString"`
	// The alias of the secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
	DefaultSecondaryConnectionStringAlias *string `pulumi:"defaultSecondaryConnectionStringAlias"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey *string `pulumi:"defaultSecondaryKey"`
	// An `identity` block as defined below.
	Identity *EventHubNamespaceIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
	MaximumThroughputUnits *int `pulumi:"maximumThroughputUnits"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkRulesets` block as defined below.
	NetworkRulesets *EventHubNamespaceNetworkRulesets `pulumi:"networkRulesets"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Defines which tier to use. Valid options are `Basic` and `Standard`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type EventHubNamespaceState struct {
	// Is Auto Inflate enabled for the EventHub Namespace?
	AutoInflateEnabled pulumi.BoolPtrInput
	// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Default capacity has a maximum of `20`, but can be increased in blocks of 20 on a committed purchase basis.
	Capacity pulumi.IntPtrInput
	// Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
	DedicatedClusterId pulumi.StringPtrInput
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString pulumi.StringPtrInput
	// The alias of the primary connection string for the authorization
	// rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
	DefaultPrimaryConnectionStringAlias pulumi.StringPtrInput
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey pulumi.StringPtrInput
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString pulumi.StringPtrInput
	// The alias of the secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
	DefaultSecondaryConnectionStringAlias pulumi.StringPtrInput
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity EventHubNamespaceIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
	MaximumThroughputUnits pulumi.IntPtrInput
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkRulesets` block as defined below.
	NetworkRulesets EventHubNamespaceNetworkRulesetsPtrInput
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Defines which tier to use. Valid options are `Basic` and `Standard`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput
}

func (EventHubNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubNamespaceState)(nil)).Elem()
}

type eventHubNamespaceArgs struct {
	// Is Auto Inflate enabled for the EventHub Namespace?
	AutoInflateEnabled *bool `pulumi:"autoInflateEnabled"`
	// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Default capacity has a maximum of `20`, but can be increased in blocks of 20 on a committed purchase basis.
	Capacity *int `pulumi:"capacity"`
	// Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
	DedicatedClusterId *string `pulumi:"dedicatedClusterId"`
	// An `identity` block as defined below.
	Identity *EventHubNamespaceIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
	MaximumThroughputUnits *int `pulumi:"maximumThroughputUnits"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkRulesets` block as defined below.
	NetworkRulesets *EventHubNamespaceNetworkRulesets `pulumi:"networkRulesets"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines which tier to use. Valid options are `Basic` and `Standard`.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a EventHubNamespace resource.
type EventHubNamespaceArgs struct {
	// Is Auto Inflate enabled for the EventHub Namespace?
	AutoInflateEnabled pulumi.BoolPtrInput
	// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Default capacity has a maximum of `20`, but can be increased in blocks of 20 on a committed purchase basis.
	Capacity pulumi.IntPtrInput
	// Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
	DedicatedClusterId pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity EventHubNamespaceIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
	MaximumThroughputUnits pulumi.IntPtrInput
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkRulesets` block as defined below.
	NetworkRulesets EventHubNamespaceNetworkRulesetsPtrInput
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Defines which tier to use. Valid options are `Basic` and `Standard`.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput
}

func (EventHubNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubNamespaceArgs)(nil)).Elem()
}

type EventHubNamespaceInput interface {
	pulumi.Input

	ToEventHubNamespaceOutput() EventHubNamespaceOutput
	ToEventHubNamespaceOutputWithContext(ctx context.Context) EventHubNamespaceOutput
}

func (*EventHubNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubNamespace)(nil))
}

func (i *EventHubNamespace) ToEventHubNamespaceOutput() EventHubNamespaceOutput {
	return i.ToEventHubNamespaceOutputWithContext(context.Background())
}

func (i *EventHubNamespace) ToEventHubNamespaceOutputWithContext(ctx context.Context) EventHubNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubNamespaceOutput)
}

func (i *EventHubNamespace) ToEventHubNamespacePtrOutput() EventHubNamespacePtrOutput {
	return i.ToEventHubNamespacePtrOutputWithContext(context.Background())
}

func (i *EventHubNamespace) ToEventHubNamespacePtrOutputWithContext(ctx context.Context) EventHubNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubNamespacePtrOutput)
}

type EventHubNamespacePtrInput interface {
	pulumi.Input

	ToEventHubNamespacePtrOutput() EventHubNamespacePtrOutput
	ToEventHubNamespacePtrOutputWithContext(ctx context.Context) EventHubNamespacePtrOutput
}

type eventHubNamespacePtrType EventHubNamespaceArgs

func (*eventHubNamespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubNamespace)(nil))
}

func (i *eventHubNamespacePtrType) ToEventHubNamespacePtrOutput() EventHubNamespacePtrOutput {
	return i.ToEventHubNamespacePtrOutputWithContext(context.Background())
}

func (i *eventHubNamespacePtrType) ToEventHubNamespacePtrOutputWithContext(ctx context.Context) EventHubNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubNamespacePtrOutput)
}

// EventHubNamespaceArrayInput is an input type that accepts EventHubNamespaceArray and EventHubNamespaceArrayOutput values.
// You can construct a concrete instance of `EventHubNamespaceArrayInput` via:
//
//          EventHubNamespaceArray{ EventHubNamespaceArgs{...} }
type EventHubNamespaceArrayInput interface {
	pulumi.Input

	ToEventHubNamespaceArrayOutput() EventHubNamespaceArrayOutput
	ToEventHubNamespaceArrayOutputWithContext(context.Context) EventHubNamespaceArrayOutput
}

type EventHubNamespaceArray []EventHubNamespaceInput

func (EventHubNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EventHubNamespace)(nil))
}

func (i EventHubNamespaceArray) ToEventHubNamespaceArrayOutput() EventHubNamespaceArrayOutput {
	return i.ToEventHubNamespaceArrayOutputWithContext(context.Background())
}

func (i EventHubNamespaceArray) ToEventHubNamespaceArrayOutputWithContext(ctx context.Context) EventHubNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubNamespaceArrayOutput)
}

// EventHubNamespaceMapInput is an input type that accepts EventHubNamespaceMap and EventHubNamespaceMapOutput values.
// You can construct a concrete instance of `EventHubNamespaceMapInput` via:
//
//          EventHubNamespaceMap{ "key": EventHubNamespaceArgs{...} }
type EventHubNamespaceMapInput interface {
	pulumi.Input

	ToEventHubNamespaceMapOutput() EventHubNamespaceMapOutput
	ToEventHubNamespaceMapOutputWithContext(context.Context) EventHubNamespaceMapOutput
}

type EventHubNamespaceMap map[string]EventHubNamespaceInput

func (EventHubNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EventHubNamespace)(nil))
}

func (i EventHubNamespaceMap) ToEventHubNamespaceMapOutput() EventHubNamespaceMapOutput {
	return i.ToEventHubNamespaceMapOutputWithContext(context.Background())
}

func (i EventHubNamespaceMap) ToEventHubNamespaceMapOutputWithContext(ctx context.Context) EventHubNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubNamespaceMapOutput)
}

type EventHubNamespaceOutput struct {
	*pulumi.OutputState
}

func (EventHubNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubNamespace)(nil))
}

func (o EventHubNamespaceOutput) ToEventHubNamespaceOutput() EventHubNamespaceOutput {
	return o
}

func (o EventHubNamespaceOutput) ToEventHubNamespaceOutputWithContext(ctx context.Context) EventHubNamespaceOutput {
	return o
}

func (o EventHubNamespaceOutput) ToEventHubNamespacePtrOutput() EventHubNamespacePtrOutput {
	return o.ToEventHubNamespacePtrOutputWithContext(context.Background())
}

func (o EventHubNamespaceOutput) ToEventHubNamespacePtrOutputWithContext(ctx context.Context) EventHubNamespacePtrOutput {
	return o.ApplyT(func(v EventHubNamespace) *EventHubNamespace {
		return &v
	}).(EventHubNamespacePtrOutput)
}

type EventHubNamespacePtrOutput struct {
	*pulumi.OutputState
}

func (EventHubNamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubNamespace)(nil))
}

func (o EventHubNamespacePtrOutput) ToEventHubNamespacePtrOutput() EventHubNamespacePtrOutput {
	return o
}

func (o EventHubNamespacePtrOutput) ToEventHubNamespacePtrOutputWithContext(ctx context.Context) EventHubNamespacePtrOutput {
	return o
}

type EventHubNamespaceArrayOutput struct{ *pulumi.OutputState }

func (EventHubNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHubNamespace)(nil))
}

func (o EventHubNamespaceArrayOutput) ToEventHubNamespaceArrayOutput() EventHubNamespaceArrayOutput {
	return o
}

func (o EventHubNamespaceArrayOutput) ToEventHubNamespaceArrayOutputWithContext(ctx context.Context) EventHubNamespaceArrayOutput {
	return o
}

func (o EventHubNamespaceArrayOutput) Index(i pulumi.IntInput) EventHubNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHubNamespace {
		return vs[0].([]EventHubNamespace)[vs[1].(int)]
	}).(EventHubNamespaceOutput)
}

type EventHubNamespaceMapOutput struct{ *pulumi.OutputState }

func (EventHubNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubNamespace)(nil))
}

func (o EventHubNamespaceMapOutput) ToEventHubNamespaceMapOutput() EventHubNamespaceMapOutput {
	return o
}

func (o EventHubNamespaceMapOutput) ToEventHubNamespaceMapOutputWithContext(ctx context.Context) EventHubNamespaceMapOutput {
	return o
}

func (o EventHubNamespaceMapOutput) MapIndex(k pulumi.StringInput) EventHubNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHubNamespace {
		return vs[0].(map[string]EventHubNamespace)[vs[1].(string)]
	}).(EventHubNamespaceOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubNamespaceOutput{})
	pulumi.RegisterOutputType(EventHubNamespacePtrOutput{})
	pulumi.RegisterOutputType(EventHubNamespaceArrayOutput{})
	pulumi.RegisterOutputType(EventHubNamespaceMapOutput{})
}
