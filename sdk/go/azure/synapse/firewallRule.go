// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to Manages a Synapse Firewall Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/synapse"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("StorageV2"),
// 			IsHnsEnabled:           pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewWorkspace(ctx, "exampleWorkspace", &synapse.WorkspaceArgs{
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			Location:                        exampleResourceGroup.Location,
// 			StorageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.ID(),
// 			SqlAdministratorLogin:           pulumi.String("sqladminuser"),
// 			SqlAdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewFirewallRule(ctx, "exampleFirewallRule", &synapse.FirewallRuleArgs{
// 			SynapseWorkspaceId: pulumi.Any(azurerm_synapse_workspace.Test.Id),
// 			StartIpAddress:     pulumi.String("0.0.0.0"),
// 			EndIpAddress:       pulumi.String("255.255.255.255"),
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
// Synapse Firewall Rule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/firewallRule:FirewallRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.Synapse/workspaces/workspace1/firewallRules/rule1
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// The Name of the firewall rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	var resource FirewallRule
	err := ctx.RegisterResource("azure:synapse/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure:synapse/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// The Name of the firewall rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress *string `pulumi:"startIpAddress"`
	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
}

type FirewallRuleState struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress pulumi.StringPtrInput
	// The Name of the firewall rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress pulumi.StringPtrInput
	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The Name of the firewall rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress string `pulumi:"startIpAddress"`
	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress pulumi.StringInput
	// The Name of the firewall rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress pulumi.StringInput
	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRule)(nil))
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

func (i *FirewallRule) ToFirewallRulePtrOutput() FirewallRulePtrOutput {
	return i.ToFirewallRulePtrOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRulePtrOutputWithContext(ctx context.Context) FirewallRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulePtrOutput)
}

type FirewallRulePtrInput interface {
	pulumi.Input

	ToFirewallRulePtrOutput() FirewallRulePtrOutput
	ToFirewallRulePtrOutputWithContext(ctx context.Context) FirewallRulePtrOutput
}

type firewallRulePtrType FirewallRuleArgs

func (*firewallRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil))
}

func (i *firewallRulePtrType) ToFirewallRulePtrOutput() FirewallRulePtrOutput {
	return i.ToFirewallRulePtrOutputWithContext(context.Background())
}

func (i *firewallRulePtrType) ToFirewallRulePtrOutputWithContext(ctx context.Context) FirewallRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulePtrOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//          FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallRule)(nil))
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//          FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallRule)(nil))
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct {
	*pulumi.OutputState
}

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRule)(nil))
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRulePtrOutput() FirewallRulePtrOutput {
	return o.ToFirewallRulePtrOutputWithContext(context.Background())
}

func (o FirewallRuleOutput) ToFirewallRulePtrOutputWithContext(ctx context.Context) FirewallRulePtrOutput {
	return o.ApplyT(func(v FirewallRule) *FirewallRule {
		return &v
	}).(FirewallRulePtrOutput)
}

type FirewallRulePtrOutput struct {
	*pulumi.OutputState
}

func (FirewallRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil))
}

func (o FirewallRulePtrOutput) ToFirewallRulePtrOutput() FirewallRulePtrOutput {
	return o
}

func (o FirewallRulePtrOutput) ToFirewallRulePtrOutputWithContext(ctx context.Context) FirewallRulePtrOutput {
	return o
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRule)(nil))
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRule {
		return vs[0].([]FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallRule)(nil))
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallRule {
		return vs[0].(map[string]FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRulePtrOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}
