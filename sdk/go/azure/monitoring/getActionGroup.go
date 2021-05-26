// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access the properties of an Action Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := monitoring.LookupActionGroup(ctx, &monitoring.LookupActionGroupArgs{
// 			ResourceGroupName: "example-rg",
// 			Name:              "tfex-actiongroup",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("actionGroupId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupActionGroup(ctx *pulumi.Context, args *LookupActionGroupArgs, opts ...pulumi.InvokeOption) (*LookupActionGroupResult, error) {
	var rv LookupActionGroupResult
	err := ctx.Invoke("azure:monitoring/getActionGroup:getActionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActionGroup.
type LookupActionGroupArgs struct {
	// Specifies the name of the Action Group.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Action Group is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getActionGroup.
type LookupActionGroupResult struct {
	// One or more `armRoleReceiver` blocks as defined below.
	ArmRoleReceivers []GetActionGroupArmRoleReceiver `pulumi:"armRoleReceivers"`
	// One or more `automationRunbookReceiver` blocks as defined below.
	AutomationRunbookReceivers []GetActionGroupAutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	// One or more `azureAppPushReceiver` blocks as defined below.
	AzureAppPushReceivers []GetActionGroupAzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// One or more `azureFunctionReceiver` blocks as defined below.
	AzureFunctionReceivers []GetActionGroupAzureFunctionReceiver `pulumi:"azureFunctionReceivers"`
	// One or more `emailReceiver` blocks as defined below.
	EmailReceivers []GetActionGroupEmailReceiver `pulumi:"emailReceivers"`
	// Whether this action group is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// One or more `itsmReceiver` blocks as defined below.
	ItsmReceivers []GetActionGroupItsmReceiver `pulumi:"itsmReceivers"`
	// One or more `logicAppReceiver` blocks as defined below.
	LogicAppReceivers []GetActionGroupLogicAppReceiver `pulumi:"logicAppReceivers"`
	// The name of the webhook receiver.
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The short name of the action group.
	ShortName string `pulumi:"shortName"`
	// One or more `smsReceiver` blocks as defined below.
	SmsReceivers []GetActionGroupSmsReceiver `pulumi:"smsReceivers"`
	// One or more `voiceReceiver` blocks as defined below.
	VoiceReceivers []GetActionGroupVoiceReceiver `pulumi:"voiceReceivers"`
	// One or more `webhookReceiver` blocks as defined below.
	WebhookReceivers []GetActionGroupWebhookReceiver `pulumi:"webhookReceivers"`
}

func LookupActionGroupApply(ctx *pulumi.Context, args LookupActionGroupApplyInput, opts ...pulumi.InvokeOption) LookupActionGroupResultOutput {
	return args.ToLookupActionGroupApplyOutput().ApplyT(func(v LookupActionGroupArgs) (LookupActionGroupResult, error) {
		r, err := LookupActionGroup(ctx, &v, opts...)
		return *r, err

	}).(LookupActionGroupResultOutput)
}

// LookupActionGroupApplyInput is an input type that accepts LookupActionGroupApplyArgs and LookupActionGroupApplyOutput values.
// You can construct a concrete instance of `LookupActionGroupApplyInput` via:
//
//          LookupActionGroupApplyArgs{...}
type LookupActionGroupApplyInput interface {
	pulumi.Input

	ToLookupActionGroupApplyOutput() LookupActionGroupApplyOutput
	ToLookupActionGroupApplyOutputWithContext(context.Context) LookupActionGroupApplyOutput
}

// A collection of arguments for invoking getActionGroup.
type LookupActionGroupApplyArgs struct {
	// Specifies the name of the Action Group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group the Action Group is located in.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupActionGroupApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionGroupArgs)(nil)).Elem()
}

func (i LookupActionGroupApplyArgs) ToLookupActionGroupApplyOutput() LookupActionGroupApplyOutput {
	return i.ToLookupActionGroupApplyOutputWithContext(context.Background())
}

func (i LookupActionGroupApplyArgs) ToLookupActionGroupApplyOutputWithContext(ctx context.Context) LookupActionGroupApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupActionGroupApplyOutput)
}

// A collection of arguments for invoking getActionGroup.
type LookupActionGroupApplyOutput struct{ *pulumi.OutputState }

func (LookupActionGroupApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionGroupArgs)(nil)).Elem()
}

func (o LookupActionGroupApplyOutput) ToLookupActionGroupApplyOutput() LookupActionGroupApplyOutput {
	return o
}

func (o LookupActionGroupApplyOutput) ToLookupActionGroupApplyOutputWithContext(ctx context.Context) LookupActionGroupApplyOutput {
	return o
}

// Specifies the name of the Action Group.
func (o LookupActionGroupApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the name of the resource group the Action Group is located in.
func (o LookupActionGroupApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getActionGroup.
type LookupActionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupActionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionGroupResult)(nil)).Elem()
}

func (o LookupActionGroupResultOutput) ToLookupActionGroupResultOutput() LookupActionGroupResultOutput {
	return o
}

func (o LookupActionGroupResultOutput) ToLookupActionGroupResultOutputWithContext(ctx context.Context) LookupActionGroupResultOutput {
	return o
}

// One or more `armRoleReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) ArmRoleReceivers() GetActionGroupArmRoleReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupArmRoleReceiver { return v.ArmRoleReceivers }).(GetActionGroupArmRoleReceiverArrayOutput)
}

// One or more `automationRunbookReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) AutomationRunbookReceivers() GetActionGroupAutomationRunbookReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupAutomationRunbookReceiver {
		return v.AutomationRunbookReceivers
	}).(GetActionGroupAutomationRunbookReceiverArrayOutput)
}

// One or more `azureAppPushReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) AzureAppPushReceivers() GetActionGroupAzureAppPushReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupAzureAppPushReceiver { return v.AzureAppPushReceivers }).(GetActionGroupAzureAppPushReceiverArrayOutput)
}

// One or more `azureFunctionReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) AzureFunctionReceivers() GetActionGroupAzureFunctionReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupAzureFunctionReceiver { return v.AzureFunctionReceivers }).(GetActionGroupAzureFunctionReceiverArrayOutput)
}

// One or more `emailReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) EmailReceivers() GetActionGroupEmailReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupEmailReceiver { return v.EmailReceivers }).(GetActionGroupEmailReceiverArrayOutput)
}

// Whether this action group is enabled.
func (o LookupActionGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupActionGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupActionGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// One or more `itsmReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) ItsmReceivers() GetActionGroupItsmReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupItsmReceiver { return v.ItsmReceivers }).(GetActionGroupItsmReceiverArrayOutput)
}

// One or more `logicAppReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) LogicAppReceivers() GetActionGroupLogicAppReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupLogicAppReceiver { return v.LogicAppReceivers }).(GetActionGroupLogicAppReceiverArrayOutput)
}

// The name of the webhook receiver.
func (o LookupActionGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActionGroupResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The short name of the action group.
func (o LookupActionGroupResultOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.ShortName }).(pulumi.StringOutput)
}

// One or more `smsReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) SmsReceivers() GetActionGroupSmsReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupSmsReceiver { return v.SmsReceivers }).(GetActionGroupSmsReceiverArrayOutput)
}

// One or more `voiceReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) VoiceReceivers() GetActionGroupVoiceReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupVoiceReceiver { return v.VoiceReceivers }).(GetActionGroupVoiceReceiverArrayOutput)
}

// One or more `webhookReceiver` blocks as defined below.
func (o LookupActionGroupResultOutput) WebhookReceivers() GetActionGroupWebhookReceiverArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []GetActionGroupWebhookReceiver { return v.WebhookReceivers }).(GetActionGroupWebhookReceiverArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionGroupApplyOutput{})
	pulumi.RegisterOutputType(LookupActionGroupResultOutput{})
}
