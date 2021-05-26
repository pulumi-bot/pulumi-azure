// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Sentinel Alert Rule Template.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sentinel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Create incidents based on Azure Security Center for IoT alerts"
// 		example, err := sentinel.GetAlertRuleTemplate(ctx, &sentinel.GetAlertRuleTemplateArgs{
// 			LogAnalyticsWorkspaceId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1",
// 			DisplayName:             &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func GetAlertRuleTemplate(ctx *pulumi.Context, args *GetAlertRuleTemplateArgs, opts ...pulumi.InvokeOption) (*GetAlertRuleTemplateResult, error) {
	var rv GetAlertRuleTemplateResult
	err := ctx.Invoke("azure:sentinel/getAlertRuleTemplate:getAlertRuleTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlertRuleTemplate.
type GetAlertRuleTemplateArgs struct {
	// The display name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the Log Analytics Workspace.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAlertRuleTemplate.
type GetAlertRuleTemplateResult struct {
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string `pulumi:"id"`
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	Name                    string `pulumi:"name"`
	// A `scheduledTemplate` block as defined below. This only applies to Sentinel Scheduled Alert Rule Template.
	ScheduledTemplates []GetAlertRuleTemplateScheduledTemplate `pulumi:"scheduledTemplates"`
	// A `securityIncidentTemplate` block as defined below. This only applies to Sentinel MS Security Incident Alert Rule Template.
	SecurityIncidentTemplates []GetAlertRuleTemplateSecurityIncidentTemplate `pulumi:"securityIncidentTemplates"`
}

func GetAlertRuleTemplateApply(ctx *pulumi.Context, args GetAlertRuleTemplateApplyInput, opts ...pulumi.InvokeOption) GetAlertRuleTemplateResultOutput {
	return args.ToGetAlertRuleTemplateApplyOutput().ApplyT(func(v GetAlertRuleTemplateArgs) (GetAlertRuleTemplateResult, error) {
		r, err := GetAlertRuleTemplate(ctx, &v, opts...)
		return *r, err

	}).(GetAlertRuleTemplateResultOutput)
}

// GetAlertRuleTemplateApplyInput is an input type that accepts GetAlertRuleTemplateApplyArgs and GetAlertRuleTemplateApplyOutput values.
// You can construct a concrete instance of `GetAlertRuleTemplateApplyInput` via:
//
//          GetAlertRuleTemplateApplyArgs{...}
type GetAlertRuleTemplateApplyInput interface {
	pulumi.Input

	ToGetAlertRuleTemplateApplyOutput() GetAlertRuleTemplateApplyOutput
	ToGetAlertRuleTemplateApplyOutputWithContext(context.Context) GetAlertRuleTemplateApplyOutput
}

// A collection of arguments for invoking getAlertRuleTemplate.
type GetAlertRuleTemplateApplyArgs struct {
	// The display name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The ID of the Log Analytics Workspace.
	LogAnalyticsWorkspaceId pulumi.StringInput `pulumi:"logAnalyticsWorkspaceId"`
	// The name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetAlertRuleTemplateApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlertRuleTemplateArgs)(nil)).Elem()
}

func (i GetAlertRuleTemplateApplyArgs) ToGetAlertRuleTemplateApplyOutput() GetAlertRuleTemplateApplyOutput {
	return i.ToGetAlertRuleTemplateApplyOutputWithContext(context.Background())
}

func (i GetAlertRuleTemplateApplyArgs) ToGetAlertRuleTemplateApplyOutputWithContext(ctx context.Context) GetAlertRuleTemplateApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAlertRuleTemplateApplyOutput)
}

// A collection of arguments for invoking getAlertRuleTemplate.
type GetAlertRuleTemplateApplyOutput struct{ *pulumi.OutputState }

func (GetAlertRuleTemplateApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlertRuleTemplateArgs)(nil)).Elem()
}

func (o GetAlertRuleTemplateApplyOutput) ToGetAlertRuleTemplateApplyOutput() GetAlertRuleTemplateApplyOutput {
	return o
}

func (o GetAlertRuleTemplateApplyOutput) ToGetAlertRuleTemplateApplyOutputWithContext(ctx context.Context) GetAlertRuleTemplateApplyOutput {
	return o
}

// The display name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
func (o GetAlertRuleTemplateApplyOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateArgs) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the Log Analytics Workspace.
func (o GetAlertRuleTemplateApplyOutput) LogAnalyticsWorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateArgs) string { return v.LogAnalyticsWorkspaceId }).(pulumi.StringOutput)
}

// The name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
func (o GetAlertRuleTemplateApplyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateArgs) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getAlertRuleTemplate.
type GetAlertRuleTemplateResultOutput struct{ *pulumi.OutputState }

func (GetAlertRuleTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlertRuleTemplateResult)(nil)).Elem()
}

func (o GetAlertRuleTemplateResultOutput) ToGetAlertRuleTemplateResultOutput() GetAlertRuleTemplateResultOutput {
	return o
}

func (o GetAlertRuleTemplateResultOutput) ToGetAlertRuleTemplateResultOutputWithContext(ctx context.Context) GetAlertRuleTemplateResultOutput {
	return o
}

func (o GetAlertRuleTemplateResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlertRuleTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAlertRuleTemplateResultOutput) LogAnalyticsWorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateResult) string { return v.LogAnalyticsWorkspaceId }).(pulumi.StringOutput)
}

func (o GetAlertRuleTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// A `scheduledTemplate` block as defined below. This only applies to Sentinel Scheduled Alert Rule Template.
func (o GetAlertRuleTemplateResultOutput) ScheduledTemplates() GetAlertRuleTemplateScheduledTemplateArrayOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateResult) []GetAlertRuleTemplateScheduledTemplate {
		return v.ScheduledTemplates
	}).(GetAlertRuleTemplateScheduledTemplateArrayOutput)
}

// A `securityIncidentTemplate` block as defined below. This only applies to Sentinel MS Security Incident Alert Rule Template.
func (o GetAlertRuleTemplateResultOutput) SecurityIncidentTemplates() GetAlertRuleTemplateSecurityIncidentTemplateArrayOutput {
	return o.ApplyT(func(v GetAlertRuleTemplateResult) []GetAlertRuleTemplateSecurityIncidentTemplate {
		return v.SecurityIncidentTemplates
	}).(GetAlertRuleTemplateSecurityIncidentTemplateArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlertRuleTemplateApplyOutput{})
	pulumi.RegisterOutputType(GetAlertRuleTemplateResultOutput{})
}
