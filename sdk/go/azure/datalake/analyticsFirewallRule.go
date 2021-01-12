// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Data Lake Analytics Firewall Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datalake"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleStore, err := datalake.NewStore(ctx, "exampleStore", &datalake.StoreArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewAnalyticsAccount(ctx, "exampleAnalyticsAccount", &datalake.AnalyticsAccountArgs{
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			Location:                exampleResourceGroup.Location,
// 			DefaultStoreAccountName: exampleStore.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewAnalyticsFirewallRule(ctx, "exampleAnalyticsFirewallRule", &datalake.AnalyticsFirewallRuleArgs{
// 			AccountName:       pulumi.Any(azurerm_data_lake_analytics.Example.Name),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StartIpAddress:    pulumi.String("1.2.3.4"),
// 			EndIpAddress:      pulumi.String("2.3.4.5"),
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
// Data Lake Analytics Firewall Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DataLakeAnalytics/accounts/mydatalakeaccount/firewallRules/rule1
// ```
type AnalyticsFirewallRule struct {
	pulumi.CustomResourceState

	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewAnalyticsFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsFirewallRule(ctx *pulumi.Context,
	name string, args *AnalyticsFirewallRuleArgs, opts ...pulumi.ResourceOption) (*AnalyticsFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	var resource AnalyticsFirewallRule
	err := ctx.RegisterResource("azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsFirewallRule gets an existing AnalyticsFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsFirewallRuleState, opts ...pulumi.ResourceOption) (*AnalyticsFirewallRule, error) {
	var resource AnalyticsFirewallRule
	err := ctx.ReadResource("azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsFirewallRule resources.
type analyticsFirewallRuleState struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName *string `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type AnalyticsFirewallRuleState struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName pulumi.StringPtrInput
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringPtrInput
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName pulumi.StringPtrInput
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringPtrInput
}

func (AnalyticsFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsFirewallRuleState)(nil)).Elem()
}

type analyticsFirewallRuleArgs struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName string `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a AnalyticsFirewallRule resource.
type AnalyticsFirewallRuleArgs struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName pulumi.StringInput
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringInput
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName pulumi.StringInput
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringInput
}

func (AnalyticsFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsFirewallRuleArgs)(nil)).Elem()
}

type AnalyticsFirewallRuleInput interface {
	pulumi.Input

	ToAnalyticsFirewallRuleOutput() AnalyticsFirewallRuleOutput
	ToAnalyticsFirewallRuleOutputWithContext(ctx context.Context) AnalyticsFirewallRuleOutput
}

func (*AnalyticsFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsFirewallRule)(nil))
}

func (i *AnalyticsFirewallRule) ToAnalyticsFirewallRuleOutput() AnalyticsFirewallRuleOutput {
	return i.ToAnalyticsFirewallRuleOutputWithContext(context.Background())
}

func (i *AnalyticsFirewallRule) ToAnalyticsFirewallRuleOutputWithContext(ctx context.Context) AnalyticsFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsFirewallRuleOutput)
}

func (i *AnalyticsFirewallRule) ToAnalyticsFirewallRulePtrOutput() AnalyticsFirewallRulePtrOutput {
	return i.ToAnalyticsFirewallRulePtrOutputWithContext(context.Background())
}

func (i *AnalyticsFirewallRule) ToAnalyticsFirewallRulePtrOutputWithContext(ctx context.Context) AnalyticsFirewallRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsFirewallRulePtrOutput)
}

type AnalyticsFirewallRulePtrInput interface {
	pulumi.Input

	ToAnalyticsFirewallRulePtrOutput() AnalyticsFirewallRulePtrOutput
	ToAnalyticsFirewallRulePtrOutputWithContext(ctx context.Context) AnalyticsFirewallRulePtrOutput
}

type analyticsFirewallRulePtrType AnalyticsFirewallRuleArgs

func (*analyticsFirewallRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsFirewallRule)(nil))
}

func (i *analyticsFirewallRulePtrType) ToAnalyticsFirewallRulePtrOutput() AnalyticsFirewallRulePtrOutput {
	return i.ToAnalyticsFirewallRulePtrOutputWithContext(context.Background())
}

func (i *analyticsFirewallRulePtrType) ToAnalyticsFirewallRulePtrOutputWithContext(ctx context.Context) AnalyticsFirewallRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsFirewallRuleOutput).ToAnalyticsFirewallRulePtrOutput()
}

// AnalyticsFirewallRuleArrayInput is an input type that accepts AnalyticsFirewallRuleArray and AnalyticsFirewallRuleArrayOutput values.
// You can construct a concrete instance of `AnalyticsFirewallRuleArrayInput` via:
//
//          AnalyticsFirewallRuleArray{ AnalyticsFirewallRuleArgs{...} }
type AnalyticsFirewallRuleArrayInput interface {
	pulumi.Input

	ToAnalyticsFirewallRuleArrayOutput() AnalyticsFirewallRuleArrayOutput
	ToAnalyticsFirewallRuleArrayOutputWithContext(context.Context) AnalyticsFirewallRuleArrayOutput
}

type AnalyticsFirewallRuleArray []AnalyticsFirewallRuleInput

func (AnalyticsFirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsFirewallRule)(nil))
}

func (i AnalyticsFirewallRuleArray) ToAnalyticsFirewallRuleArrayOutput() AnalyticsFirewallRuleArrayOutput {
	return i.ToAnalyticsFirewallRuleArrayOutputWithContext(context.Background())
}

func (i AnalyticsFirewallRuleArray) ToAnalyticsFirewallRuleArrayOutputWithContext(ctx context.Context) AnalyticsFirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsFirewallRuleArrayOutput)
}

// AnalyticsFirewallRuleMapInput is an input type that accepts AnalyticsFirewallRuleMap and AnalyticsFirewallRuleMapOutput values.
// You can construct a concrete instance of `AnalyticsFirewallRuleMapInput` via:
//
//          AnalyticsFirewallRuleMap{ "key": AnalyticsFirewallRuleArgs{...} }
type AnalyticsFirewallRuleMapInput interface {
	pulumi.Input

	ToAnalyticsFirewallRuleMapOutput() AnalyticsFirewallRuleMapOutput
	ToAnalyticsFirewallRuleMapOutputWithContext(context.Context) AnalyticsFirewallRuleMapOutput
}

type AnalyticsFirewallRuleMap map[string]AnalyticsFirewallRuleInput

func (AnalyticsFirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AnalyticsFirewallRule)(nil))
}

func (i AnalyticsFirewallRuleMap) ToAnalyticsFirewallRuleMapOutput() AnalyticsFirewallRuleMapOutput {
	return i.ToAnalyticsFirewallRuleMapOutputWithContext(context.Background())
}

func (i AnalyticsFirewallRuleMap) ToAnalyticsFirewallRuleMapOutputWithContext(ctx context.Context) AnalyticsFirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsFirewallRuleMapOutput)
}

type AnalyticsFirewallRuleOutput struct {
	*pulumi.OutputState
}

func (AnalyticsFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsFirewallRule)(nil))
}

func (o AnalyticsFirewallRuleOutput) ToAnalyticsFirewallRuleOutput() AnalyticsFirewallRuleOutput {
	return o
}

func (o AnalyticsFirewallRuleOutput) ToAnalyticsFirewallRuleOutputWithContext(ctx context.Context) AnalyticsFirewallRuleOutput {
	return o
}

func (o AnalyticsFirewallRuleOutput) ToAnalyticsFirewallRulePtrOutput() AnalyticsFirewallRulePtrOutput {
	return o.ToAnalyticsFirewallRulePtrOutputWithContext(context.Background())
}

func (o AnalyticsFirewallRuleOutput) ToAnalyticsFirewallRulePtrOutputWithContext(ctx context.Context) AnalyticsFirewallRulePtrOutput {
	return o.ApplyT(func(v AnalyticsFirewallRule) *AnalyticsFirewallRule {
		return &v
	}).(AnalyticsFirewallRulePtrOutput)
}

type AnalyticsFirewallRulePtrOutput struct {
	*pulumi.OutputState
}

func (AnalyticsFirewallRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsFirewallRule)(nil))
}

func (o AnalyticsFirewallRulePtrOutput) ToAnalyticsFirewallRulePtrOutput() AnalyticsFirewallRulePtrOutput {
	return o
}

func (o AnalyticsFirewallRulePtrOutput) ToAnalyticsFirewallRulePtrOutputWithContext(ctx context.Context) AnalyticsFirewallRulePtrOutput {
	return o
}

type AnalyticsFirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (AnalyticsFirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsFirewallRule)(nil))
}

func (o AnalyticsFirewallRuleArrayOutput) ToAnalyticsFirewallRuleArrayOutput() AnalyticsFirewallRuleArrayOutput {
	return o
}

func (o AnalyticsFirewallRuleArrayOutput) ToAnalyticsFirewallRuleArrayOutputWithContext(ctx context.Context) AnalyticsFirewallRuleArrayOutput {
	return o
}

func (o AnalyticsFirewallRuleArrayOutput) Index(i pulumi.IntInput) AnalyticsFirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AnalyticsFirewallRule {
		return vs[0].([]AnalyticsFirewallRule)[vs[1].(int)]
	}).(AnalyticsFirewallRuleOutput)
}

type AnalyticsFirewallRuleMapOutput struct{ *pulumi.OutputState }

func (AnalyticsFirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AnalyticsFirewallRule)(nil))
}

func (o AnalyticsFirewallRuleMapOutput) ToAnalyticsFirewallRuleMapOutput() AnalyticsFirewallRuleMapOutput {
	return o
}

func (o AnalyticsFirewallRuleMapOutput) ToAnalyticsFirewallRuleMapOutputWithContext(ctx context.Context) AnalyticsFirewallRuleMapOutput {
	return o
}

func (o AnalyticsFirewallRuleMapOutput) MapIndex(k pulumi.StringInput) AnalyticsFirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AnalyticsFirewallRule {
		return vs[0].(map[string]AnalyticsFirewallRule)[vs[1].(string)]
	}).(AnalyticsFirewallRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsFirewallRuleOutput{})
	pulumi.RegisterOutputType(AnalyticsFirewallRulePtrOutput{})
	pulumi.RegisterOutputType(AnalyticsFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(AnalyticsFirewallRuleMapOutput{})
}
