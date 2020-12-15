// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Event Hubs authorization Rule within an Event Hub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
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
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Basic"),
// 			Capacity:          pulumi.Int(2),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &eventhub.AuthorizationRuleArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Listen:            pulumi.Bool(true),
// 			Send:              pulumi.Bool(false),
// 			Manage:            pulumi.Bool(false),
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
// EventHub Authorization Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/authorizationRule:AuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/authorizationRules/rule1
// ```
type AuthorizationRule struct {
	pulumi.CustomResourceState

	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
	PrimaryConnectionStringAlias pulumi.StringOutput `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
	SecondaryConnectionStringAlias pulumi.StringOutput `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationRule(ctx *pulumi.Context,
	name string, args *AuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*AuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventhubName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/authorizationRule:AuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationRule gets an existing AuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationRuleState, opts ...pulumi.ResourceOption) (*AuthorizationRule, error) {
	var resource AuthorizationRule
	err := ctx.ReadResource("azure:eventhub/authorizationRule:AuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationRule resources.
type authorizationRuleState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName *string `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
	PrimaryConnectionStringAlias *string `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
	SecondaryConnectionStringAlias *string `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type AuthorizationRuleState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
	PrimaryConnectionStringAlias pulumi.StringPtrInput
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
	SecondaryConnectionStringAlias pulumi.StringPtrInput
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (AuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRuleState)(nil)).Elem()
}

type authorizationRuleArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName string `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a AuthorizationRule resource.
type AuthorizationRuleArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringInput
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (AuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRuleArgs)(nil)).Elem()
}

type AuthorizationRuleInput interface {
	pulumi.Input

	ToAuthorizationRuleOutput() AuthorizationRuleOutput
	ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput
}

func (*AuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationRule)(nil))
}

func (i *AuthorizationRule) ToAuthorizationRuleOutput() AuthorizationRuleOutput {
	return i.ToAuthorizationRuleOutputWithContext(context.Background())
}

func (i *AuthorizationRule) ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleOutput)
}

type AuthorizationRuleOutput struct {
	*pulumi.OutputState
}

func (AuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationRule)(nil))
}

func (o AuthorizationRuleOutput) ToAuthorizationRuleOutput() AuthorizationRuleOutput {
	return o
}

func (o AuthorizationRuleOutput) ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AuthorizationRuleOutput{})
}
