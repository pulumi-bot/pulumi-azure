// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managedapplication

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses this data source to access information about an existing Managed Application Definition.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/managedapplication"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := managedapplication.LookupDefinition(ctx, &managedapplication.LookupDefinitionArgs{
// 			Name:              "example-managedappdef",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupDefinition(ctx *pulumi.Context, args *LookupDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupDefinitionResult, error) {
	var rv LookupDefinitionResult
	err := ctx.Invoke("azure:managedapplication/getDefinition:getDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefinition.
type LookupDefinitionArgs struct {
	// Specifies the name of the Managed Application Definition.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where this Managed Application Definition exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getDefinition.
type LookupDefinitionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}
