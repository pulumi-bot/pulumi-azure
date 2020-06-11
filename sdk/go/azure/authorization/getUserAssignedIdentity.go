// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing User Assigned Identity.
//
// ## Example Usage
//
// ### Reference An Existing)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := authorization.LookupUserAssignedIdentity(ctx, &authorization.LookupUserAssignedIdentityArgs{
// 			Name:              "name_of_user_assigned_identity",
// 			ResourceGroupName: "name_of_resource_group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("uaiClientId", example.ClientId)
// 		ctx.Export("uaiPrincipalId", example.PrincipalId)
// 		return nil
// 	})
// }
// ```
func LookupUserAssignedIdentity(ctx *pulumi.Context, args *LookupUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*LookupUserAssignedIdentityResult, error) {
	var rv LookupUserAssignedIdentityResult
	err := ctx.Invoke("azure:authorization/getUserAssignedIdentity:getUserAssignedIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserAssignedIdentity.
type LookupUserAssignedIdentityArgs struct {
	// The name of the User Assigned Identity.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the User Assigned Identity exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getUserAssignedIdentity.
type LookupUserAssignedIdentityResult struct {
	// The Client ID of the User Assigned Identity.
	ClientId string `pulumi:"clientId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure location where the User Assigned Identity exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The Service Principal ID of the User Assigned Identity.
	PrincipalId       string `pulumi:"principalId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the User Assigned Identity.
	Tags map[string]string `pulumi:"tags"`
}
