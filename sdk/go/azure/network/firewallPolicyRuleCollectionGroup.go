// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Firewall Policy Rule Collection Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
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
// 		exampleFirewallPolicy, err := network.NewFirewallPolicy(ctx, "exampleFirewallPolicy", &network.FirewallPolicyArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewFirewallPolicyRuleCollectionGroup(ctx, "exampleFirewallPolicyRuleCollectionGroup", &network.FirewallPolicyRuleCollectionGroupArgs{
// 			FirewallPolicyId: exampleFirewallPolicy.ID(),
// 			Priority:         pulumi.Int(500),
// 			ApplicationRuleCollections: network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionArray{
// 				&network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionArgs{
// 					Name:     pulumi.String("app_rule_collection1"),
// 					Priority: pulumi.Int(500),
// 					Action:   pulumi.String("Deny"),
// 					Rules: []network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleArgs{
// 						&network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleArgs{
// 							Name: pulumi.String("app_rule_collection1_rule1"),
// 							Protocols: []network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolArgs{
// 								&network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolArgs{
// 									Type: pulumi.String("Http"),
// 									Port: pulumi.Int(80),
// 								},
// 								&network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolArgs{
// 									Type: pulumi.String("Https"),
// 									Port: pulumi.Int(443),
// 								},
// 							},
// 							SourceAddresses: pulumi.StringArray{
// 								pulumi.String("10.0.0.1"),
// 							},
// 							DestinationFqdns: pulumi.StringArray{
// 								pulumi.String(".microsoft.com"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			NetworkRuleCollections: network.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionArray{
// 				&network.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionArgs{
// 					Name:     pulumi.String("network_rule_collection1"),
// 					Priority: pulumi.Int(400),
// 					Action:   pulumi.String("Deny"),
// 					Rules: []network.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleArgs{
// 						&network.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleArgs{
// 							Name: pulumi.String("network_rule_collection1_rule1"),
// 							Protocols: pulumi.StringArray{
// 								pulumi.String("TCP"),
// 								pulumi.String("UDP"),
// 							},
// 							SourceAddresses: pulumi.StringArray{
// 								pulumi.String("10.0.0.1"),
// 							},
// 							DestinationAddresses: pulumi.StringArray{
// 								pulumi.String("192.168.1.1"),
// 								pulumi.String("192.168.1.2"),
// 							},
// 							DestinationPorts: pulumi.StringArray{
// 								pulumi.String("80"),
// 								pulumi.String("1000-2000"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			NatRuleCollections: network.FirewallPolicyRuleCollectionGroupNatRuleCollectionArray{
// 				&network.FirewallPolicyRuleCollectionGroupNatRuleCollectionArgs{
// 					Name:     pulumi.String("nat_rule_collection1"),
// 					Priority: pulumi.Int(300),
// 					Action:   pulumi.String("Dnat"),
// 					Rules: []network.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleArgs{
// 						&network.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleArgs{
// 							Name: pulumi.String("nat_rule_collection1_rule1"),
// 							Protocols: pulumi.StringArray{
// 								pulumi.String("TCP"),
// 								pulumi.String("UDP"),
// 							},
// 							SourceAddresses: pulumi.StringArray{
// 								pulumi.String("10.0.0.1"),
// 								pulumi.String("10.0.0.2"),
// 							},
// 							DestinationAddress: pulumi.String("192.168.1.1"),
// 							DestinationPorts: pulumi.StringArray{
// 								pulumi.String("80"),
// 								pulumi.String("1000-2000"),
// 							},
// 							TranslatedAddress: pulumi.String("192.168.0.1"),
// 							TranslatedPort:    pulumi.Int(8080),
// 						},
// 					},
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
// Firewall Policy Rule Collection Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/firewallPolicyRuleCollectionGroup:FirewallPolicyRuleCollectionGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/firewallPolicies/policy1/ruleCollectionGroups/gruop1
// ```
type FirewallPolicyRuleCollectionGroup struct {
	pulumi.CustomResourceState

	// One or more `applicationRuleCollection` blocks as defined below.
	ApplicationRuleCollections FirewallPolicyRuleCollectionGroupApplicationRuleCollectionArrayOutput `pulumi:"applicationRuleCollections"`
	// The ID of the Firewall Policy where the Firewall Policy Rule Collection Group should exist. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	FirewallPolicyId pulumi.StringOutput `pulumi:"firewallPolicyId"`
	// The name which should be used for this Firewall Policy Rule Collection Group. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `natRuleCollection` blocks as defined below.
	NatRuleCollections FirewallPolicyRuleCollectionGroupNatRuleCollectionArrayOutput `pulumi:"natRuleCollections"`
	// One or more `networkRuleCollection` blocks as defined below.
	NetworkRuleCollections FirewallPolicyRuleCollectionGroupNetworkRuleCollectionArrayOutput `pulumi:"networkRuleCollections"`
	// The priority of the Firewall Policy Rule Collection Group. The range is 100-65000.
	Priority pulumi.IntOutput `pulumi:"priority"`
}

// NewFirewallPolicyRuleCollectionGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicyId'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.RegisterResource("azure:network/firewallPolicyRuleCollectionGroup:FirewallPolicyRuleCollectionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyRuleCollectionGroup gets an existing FirewallPolicyRuleCollectionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleCollectionGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.ReadResource("azure:network/firewallPolicyRuleCollectionGroup:FirewallPolicyRuleCollectionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyRuleCollectionGroup resources.
type firewallPolicyRuleCollectionGroupState struct {
	// One or more `applicationRuleCollection` blocks as defined below.
	ApplicationRuleCollections []FirewallPolicyRuleCollectionGroupApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	// The ID of the Firewall Policy where the Firewall Policy Rule Collection Group should exist. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	FirewallPolicyId *string `pulumi:"firewallPolicyId"`
	// The name which should be used for this Firewall Policy Rule Collection Group. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	Name *string `pulumi:"name"`
	// One or more `natRuleCollection` blocks as defined below.
	NatRuleCollections []FirewallPolicyRuleCollectionGroupNatRuleCollection `pulumi:"natRuleCollections"`
	// One or more `networkRuleCollection` blocks as defined below.
	NetworkRuleCollections []FirewallPolicyRuleCollectionGroupNetworkRuleCollection `pulumi:"networkRuleCollections"`
	// The priority of the Firewall Policy Rule Collection Group. The range is 100-65000.
	Priority *int `pulumi:"priority"`
}

type FirewallPolicyRuleCollectionGroupState struct {
	// One or more `applicationRuleCollection` blocks as defined below.
	ApplicationRuleCollections FirewallPolicyRuleCollectionGroupApplicationRuleCollectionArrayInput
	// The ID of the Firewall Policy where the Firewall Policy Rule Collection Group should exist. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	FirewallPolicyId pulumi.StringPtrInput
	// The name which should be used for this Firewall Policy Rule Collection Group. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	Name pulumi.StringPtrInput
	// One or more `natRuleCollection` blocks as defined below.
	NatRuleCollections FirewallPolicyRuleCollectionGroupNatRuleCollectionArrayInput
	// One or more `networkRuleCollection` blocks as defined below.
	NetworkRuleCollections FirewallPolicyRuleCollectionGroupNetworkRuleCollectionArrayInput
	// The priority of the Firewall Policy Rule Collection Group. The range is 100-65000.
	Priority pulumi.IntPtrInput
}

func (FirewallPolicyRuleCollectionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupState)(nil)).Elem()
}

type firewallPolicyRuleCollectionGroupArgs struct {
	// One or more `applicationRuleCollection` blocks as defined below.
	ApplicationRuleCollections []FirewallPolicyRuleCollectionGroupApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	// The ID of the Firewall Policy where the Firewall Policy Rule Collection Group should exist. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	FirewallPolicyId string `pulumi:"firewallPolicyId"`
	// The name which should be used for this Firewall Policy Rule Collection Group. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	Name *string `pulumi:"name"`
	// One or more `natRuleCollection` blocks as defined below.
	NatRuleCollections []FirewallPolicyRuleCollectionGroupNatRuleCollection `pulumi:"natRuleCollections"`
	// One or more `networkRuleCollection` blocks as defined below.
	NetworkRuleCollections []FirewallPolicyRuleCollectionGroupNetworkRuleCollection `pulumi:"networkRuleCollections"`
	// The priority of the Firewall Policy Rule Collection Group. The range is 100-65000.
	Priority int `pulumi:"priority"`
}

// The set of arguments for constructing a FirewallPolicyRuleCollectionGroup resource.
type FirewallPolicyRuleCollectionGroupArgs struct {
	// One or more `applicationRuleCollection` blocks as defined below.
	ApplicationRuleCollections FirewallPolicyRuleCollectionGroupApplicationRuleCollectionArrayInput
	// The ID of the Firewall Policy where the Firewall Policy Rule Collection Group should exist. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	FirewallPolicyId pulumi.StringInput
	// The name which should be used for this Firewall Policy Rule Collection Group. Changing this forces a new Firewall Policy Rule Collection Group to be created.
	Name pulumi.StringPtrInput
	// One or more `natRuleCollection` blocks as defined below.
	NatRuleCollections FirewallPolicyRuleCollectionGroupNatRuleCollectionArrayInput
	// One or more `networkRuleCollection` blocks as defined below.
	NetworkRuleCollections FirewallPolicyRuleCollectionGroupNetworkRuleCollectionArrayInput
	// The priority of the Firewall Policy Rule Collection Group. The range is 100-65000.
	Priority pulumi.IntInput
}

func (FirewallPolicyRuleCollectionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupArgs)(nil)).Elem()
}

type FirewallPolicyRuleCollectionGroupInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput
	ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput
}

func (*FirewallPolicyRuleCollectionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionGroup)(nil))
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return i.ToFirewallPolicyRuleCollectionGroupOutputWithContext(context.Background())
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupOutput)
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupPtrOutput() FirewallPolicyRuleCollectionGroupPtrOutput {
	return i.ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(context.Background())
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupPtrOutput)
}

type FirewallPolicyRuleCollectionGroupPtrInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupPtrOutput() FirewallPolicyRuleCollectionGroupPtrOutput
	ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupPtrOutput
}

type firewallPolicyRuleCollectionGroupPtrType FirewallPolicyRuleCollectionGroupArgs

func (*firewallPolicyRuleCollectionGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionGroup)(nil))
}

func (i *firewallPolicyRuleCollectionGroupPtrType) ToFirewallPolicyRuleCollectionGroupPtrOutput() FirewallPolicyRuleCollectionGroupPtrOutput {
	return i.ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(context.Background())
}

func (i *firewallPolicyRuleCollectionGroupPtrType) ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupPtrOutput)
}

// FirewallPolicyRuleCollectionGroupArrayInput is an input type that accepts FirewallPolicyRuleCollectionGroupArray and FirewallPolicyRuleCollectionGroupArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyRuleCollectionGroupArrayInput` via:
//
//          FirewallPolicyRuleCollectionGroupArray{ FirewallPolicyRuleCollectionGroupArgs{...} }
type FirewallPolicyRuleCollectionGroupArrayInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupArrayOutput() FirewallPolicyRuleCollectionGroupArrayOutput
	ToFirewallPolicyRuleCollectionGroupArrayOutputWithContext(context.Context) FirewallPolicyRuleCollectionGroupArrayOutput
}

type FirewallPolicyRuleCollectionGroupArray []FirewallPolicyRuleCollectionGroupInput

func (FirewallPolicyRuleCollectionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallPolicyRuleCollectionGroup)(nil))
}

func (i FirewallPolicyRuleCollectionGroupArray) ToFirewallPolicyRuleCollectionGroupArrayOutput() FirewallPolicyRuleCollectionGroupArrayOutput {
	return i.ToFirewallPolicyRuleCollectionGroupArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyRuleCollectionGroupArray) ToFirewallPolicyRuleCollectionGroupArrayOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupArrayOutput)
}

// FirewallPolicyRuleCollectionGroupMapInput is an input type that accepts FirewallPolicyRuleCollectionGroupMap and FirewallPolicyRuleCollectionGroupMapOutput values.
// You can construct a concrete instance of `FirewallPolicyRuleCollectionGroupMapInput` via:
//
//          FirewallPolicyRuleCollectionGroupMap{ "key": FirewallPolicyRuleCollectionGroupArgs{...} }
type FirewallPolicyRuleCollectionGroupMapInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupMapOutput() FirewallPolicyRuleCollectionGroupMapOutput
	ToFirewallPolicyRuleCollectionGroupMapOutputWithContext(context.Context) FirewallPolicyRuleCollectionGroupMapOutput
}

type FirewallPolicyRuleCollectionGroupMap map[string]FirewallPolicyRuleCollectionGroupInput

func (FirewallPolicyRuleCollectionGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallPolicyRuleCollectionGroup)(nil))
}

func (i FirewallPolicyRuleCollectionGroupMap) ToFirewallPolicyRuleCollectionGroupMapOutput() FirewallPolicyRuleCollectionGroupMapOutput {
	return i.ToFirewallPolicyRuleCollectionGroupMapOutputWithContext(context.Background())
}

func (i FirewallPolicyRuleCollectionGroupMap) ToFirewallPolicyRuleCollectionGroupMapOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupMapOutput)
}

type FirewallPolicyRuleCollectionGroupOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyRuleCollectionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionGroup)(nil))
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupPtrOutput() FirewallPolicyRuleCollectionGroupPtrOutput {
	return o.ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupPtrOutput {
	return o.ApplyT(func(v FirewallPolicyRuleCollectionGroup) *FirewallPolicyRuleCollectionGroup {
		return &v
	}).(FirewallPolicyRuleCollectionGroupPtrOutput)
}

type FirewallPolicyRuleCollectionGroupPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyRuleCollectionGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionGroup)(nil))
}

func (o FirewallPolicyRuleCollectionGroupPtrOutput) ToFirewallPolicyRuleCollectionGroupPtrOutput() FirewallPolicyRuleCollectionGroupPtrOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupPtrOutput) ToFirewallPolicyRuleCollectionGroupPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupPtrOutput {
	return o
}

type FirewallPolicyRuleCollectionGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleCollectionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPolicyRuleCollectionGroup)(nil))
}

func (o FirewallPolicyRuleCollectionGroupArrayOutput) ToFirewallPolicyRuleCollectionGroupArrayOutput() FirewallPolicyRuleCollectionGroupArrayOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupArrayOutput) ToFirewallPolicyRuleCollectionGroupArrayOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupArrayOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupArrayOutput) Index(i pulumi.IntInput) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallPolicyRuleCollectionGroup {
		return vs[0].([]FirewallPolicyRuleCollectionGroup)[vs[1].(int)]
	}).(FirewallPolicyRuleCollectionGroupOutput)
}

type FirewallPolicyRuleCollectionGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleCollectionGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallPolicyRuleCollectionGroup)(nil))
}

func (o FirewallPolicyRuleCollectionGroupMapOutput) ToFirewallPolicyRuleCollectionGroupMapOutput() FirewallPolicyRuleCollectionGroupMapOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupMapOutput) ToFirewallPolicyRuleCollectionGroupMapOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupMapOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallPolicyRuleCollectionGroup {
		return vs[0].(map[string]FirewallPolicyRuleCollectionGroup)[vs[1].(string)]
	}).(FirewallPolicyRuleCollectionGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupMapOutput{})
}
