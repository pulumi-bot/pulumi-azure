// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Automation Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/automation"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := automation.LookupAccount(ctx, "azure:automation:getAccount", &automation.LookupAccountArgs{
// 			Name:              "example-account",
// 			ResourceGroupName: "example-resources",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("automationAccountId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure:automation/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// The name of the Automation Account.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Automation Account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAccount.
type LookupAccountResult struct {
	// The Endpoint for this Auomation Account.
	Endpoint string `pulumi:"endpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The Primary Access Key for the Automation Account.
	PrimaryKey        string `pulumi:"primaryKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary Access Key for the Automation Account.
	SecondaryKey string `pulumi:"secondaryKey"`
}
