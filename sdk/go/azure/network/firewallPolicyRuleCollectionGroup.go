// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Firewall Policy Rule Collection Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 					Rules: network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleArray{
// 						&network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleArgs{
// 							Name: pulumi.String("app_rule_collection1_rule1"),
// 							Protocols: network.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolArray{
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
// 					Rules: network.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleArray{
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
// 					Rules: network.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleArray{
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

func (FirewallPolicyRuleCollectionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionGroup)(nil)).Elem()
}

func (i FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return i.ToFirewallPolicyRuleCollectionGroupOutputWithContext(context.Background())
}

func (i FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupOutput)
}

type FirewallPolicyRuleCollectionGroupOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyRuleCollectionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionGroupOutput)(nil)).Elem()
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupOutput{})
}
