// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Automation Bool Variable.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/automation"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := automation.LookupBoolVariable(ctx, &automation.LookupBoolVariableArgs{
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
func LookupBoolVariable(ctx *pulumi.Context, args *LookupBoolVariableArgs, opts ...pulumi.InvokeOption) (*LookupBoolVariableResult, error) {
	var rv LookupBoolVariableResult
	err := ctx.Invoke("azure:automation/getBoolVariable:getBoolVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBoolVariable.
type LookupBoolVariableArgs struct {
	// The name of the automation account in which the Automation Variable exists.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of the Automation Variable.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the automation account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getBoolVariable.
type LookupBoolVariableResult struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description string `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted bool `pulumi:"encrypted"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The value of the Automation Variable as a `boolean`.
	Value bool `pulumi:"value"`
}
