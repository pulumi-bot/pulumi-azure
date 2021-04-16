// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Policy Set Definition.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/policy"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Policy Set Definition Example"
// 		example, err := policy.LookupPolicySetDefinition(ctx, &policy.LookupPolicySetDefinitionArgs{
// 			DisplayName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupPolicySetDefinition(ctx *pulumi.Context, args *LookupPolicySetDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionResult, error) {
	var rv LookupPolicySetDefinitionResult
	err := ctx.Invoke("azure:policy/getPolicySetDefinition:getPolicySetDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicySetDefinition.
type LookupPolicySetDefinitionArgs struct {
	// Specifies the display name of the Policy Set Definition. Conflicts with `name`.
	DisplayName *string `pulumi:"displayName"`
	// Only retrieve Policy Set Definitions from this Management Group.
	ManagementGroupName *string `pulumi:"managementGroupName"`
	// Specifies the name of the Policy Set Definition. Conflicts with `displayName`.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getPolicySetDefinition.
type LookupPolicySetDefinitionResult struct {
	// The description of this policy definition group.
	Description string `pulumi:"description"`
	// The display name of this policy definition group.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string  `pulumi:"id"`
	ManagementGroupName *string `pulumi:"managementGroupName"`
	// Any Metadata defined in the Policy Set Definition.
	Metadata string `pulumi:"metadata"`
	// The name of this policy definition group.
	Name string `pulumi:"name"`
	// The mapping of the parameter values for the referenced policy rule. The keys are the parameter names.
	Parameters string `pulumi:"parameters"`
	// One or more `policyDefinitionGroup` blocks as defined below.
	PolicyDefinitionGroups []GetPolicySetDefinitionPolicyDefinitionGroup `pulumi:"policyDefinitionGroups"`
	// One or more `policyDefinitionReference` blocks as defined below.
	PolicyDefinitionReferences []GetPolicySetDefinitionPolicyDefinitionReference `pulumi:"policyDefinitionReferences"`
	// The policy definitions contained within the policy set definition.
	PolicyDefinitions string `pulumi:"policyDefinitions"`
	// The Type of the Policy Set Definition.
	PolicyType string `pulumi:"policyType"`
}
