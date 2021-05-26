// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Automation String Variable.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/automation"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := automation.LookupStringVariable(ctx, &automation.LookupStringVariableArgs{
// 			Name:                  "tfex-example-var",
// 			ResourceGroupName:     "tfex-example-rg",
// 			AutomationAccountName: "tfex-example-account",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("variableId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupStringVariable(ctx *pulumi.Context, args *LookupStringVariableArgs, opts ...pulumi.InvokeOption) (*LookupStringVariableResult, error) {
	var rv LookupStringVariableResult
	err := ctx.Invoke("azure:automation/getStringVariable:getStringVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStringVariable.
type LookupStringVariableArgs struct {
	// The name of the automation account in which the Automation Variable exists.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of the Automation Variable.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the automation account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getStringVariable.
type LookupStringVariableResult struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description string `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted bool `pulumi:"encrypted"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The value of the Automation Variable as a `string`.
	Value string `pulumi:"value"`
}

func LookupStringVariableApply(ctx *pulumi.Context, args LookupStringVariableApplyInput, opts ...pulumi.InvokeOption) LookupStringVariableResultOutput {
	return args.ToLookupStringVariableApplyOutput().ApplyT(func(v LookupStringVariableArgs) (LookupStringVariableResult, error) {
		r, err := LookupStringVariable(ctx, &v, opts...)
		return *r, err

	}).(LookupStringVariableResultOutput)
}

// LookupStringVariableApplyInput is an input type that accepts LookupStringVariableApplyArgs and LookupStringVariableApplyOutput values.
// You can construct a concrete instance of `LookupStringVariableApplyInput` via:
//
//          LookupStringVariableApplyArgs{...}
type LookupStringVariableApplyInput interface {
	pulumi.Input

	ToLookupStringVariableApplyOutput() LookupStringVariableApplyOutput
	ToLookupStringVariableApplyOutputWithContext(context.Context) LookupStringVariableApplyOutput
}

// A collection of arguments for invoking getStringVariable.
type LookupStringVariableApplyArgs struct {
	// The name of the automation account in which the Automation Variable exists.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The name of the Automation Variable.
	Name pulumi.StringInput `pulumi:"name"`
	// The Name of the Resource Group where the automation account exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStringVariableApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStringVariableArgs)(nil)).Elem()
}

func (i LookupStringVariableApplyArgs) ToLookupStringVariableApplyOutput() LookupStringVariableApplyOutput {
	return i.ToLookupStringVariableApplyOutputWithContext(context.Background())
}

func (i LookupStringVariableApplyArgs) ToLookupStringVariableApplyOutputWithContext(ctx context.Context) LookupStringVariableApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupStringVariableApplyOutput)
}

// A collection of arguments for invoking getStringVariable.
type LookupStringVariableApplyOutput struct{ *pulumi.OutputState }

func (LookupStringVariableApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStringVariableArgs)(nil)).Elem()
}

func (o LookupStringVariableApplyOutput) ToLookupStringVariableApplyOutput() LookupStringVariableApplyOutput {
	return o
}

func (o LookupStringVariableApplyOutput) ToLookupStringVariableApplyOutputWithContext(ctx context.Context) LookupStringVariableApplyOutput {
	return o
}

// The name of the automation account in which the Automation Variable exists.
func (o LookupStringVariableApplyOutput) AutomationAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableArgs) string { return v.AutomationAccountName }).(pulumi.StringOutput)
}

// The name of the Automation Variable.
func (o LookupStringVariableApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The Name of the Resource Group where the automation account exists.
func (o LookupStringVariableApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getStringVariable.
type LookupStringVariableResultOutput struct{ *pulumi.OutputState }

func (LookupStringVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStringVariableResult)(nil)).Elem()
}

func (o LookupStringVariableResultOutput) ToLookupStringVariableResultOutput() LookupStringVariableResultOutput {
	return o
}

func (o LookupStringVariableResultOutput) ToLookupStringVariableResultOutputWithContext(ctx context.Context) LookupStringVariableResultOutput {
	return o
}

func (o LookupStringVariableResultOutput) AutomationAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableResult) string { return v.AutomationAccountName }).(pulumi.StringOutput)
}

// The description of the Automation Variable.
func (o LookupStringVariableResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableResult) string { return v.Description }).(pulumi.StringOutput)
}

// Specifies if the Automation Variable is encrypted. Defaults to `false`.
func (o LookupStringVariableResultOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStringVariableResult) bool { return v.Encrypted }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStringVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStringVariableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStringVariableResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The value of the Automation Variable as a `string`.
func (o LookupStringVariableResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStringVariableResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStringVariableApplyOutput{})
	pulumi.RegisterOutputType(LookupStringVariableResultOutput{})
}
