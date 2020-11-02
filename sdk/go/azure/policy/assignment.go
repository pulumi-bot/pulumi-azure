// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Configures the specified Policy Definition at the specified Scope. Also, Policy Set Definitions are supported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/policy"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleDefinition, err := policy.NewDefinition(ctx, "exampleDefinition", &policy.DefinitionArgs{
// 			PolicyType:  pulumi.String("Custom"),
// 			Mode:        pulumi.String("All"),
// 			DisplayName: pulumi.String("my-policy-definition"),
// 			PolicyRule: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "	{\n", "    \"if\": {\n", "      \"not\": {\n", "        \"field\": \"location\",\n", "        \"in\": \"[parameters('allowedLocations')]\"\n", "      }\n", "    },\n", "    \"then\": {\n", "      \"effect\": \"audit\"\n", "    }\n", "  }\n")),
// 			Parameters: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v", "	{\n", "    \"allowedLocations\": {\n", "      \"type\": \"Array\",\n", "      \"metadata\": {\n", "        \"description\": \"The list of allowed locations for resources.\",\n", "        \"displayName\": \"Allowed locations\",\n", "        \"strongType\": \"location\"\n", "      }\n", "    }\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = policy.NewAssignment(ctx, "exampleAssignment", &policy.AssignmentArgs{
// 			Scope:              exampleResourceGroup.ID(),
// 			PolicyDefinitionId: exampleDefinition.ID(),
// 			Description:        pulumi.String("Policy Assignment created via an Acceptance Test"),
// 			DisplayName:        pulumi.String("My Example Policy Assignment"),
// 			Metadata:           pulumi.String(fmt.Sprintf("%v%v%v", "    {\n", "    \"category\": \"General\"\n", "    }\n")),
// 			Parameters:         pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", "  \"allowedLocations\": {\n", "    \"value\": [ \"West Europe\" ]\n", "  }\n", "}\n")),
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
// Policy Assignments can be imported using the `policy name`, e.g.
type Assignment struct {
	pulumi.CustomResourceState

	// A description to use for this Policy Assignment. Changing this forces a new resource to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
	EnforcementMode pulumi.BoolPtrOutput `pulumi:"enforcementMode"`
	// An `identity` block.
	Identity AssignmentIdentityOutput `pulumi:"identity"`
	// The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name of the Policy Assignment. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
	NotScopes pulumi.StringArrayOutput `pulumi:"notScopes"`
	// Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The ID of the Policy Definition to be applied at the specified Scope.
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil || args.PolicyDefinitionId == nil {
		return nil, errors.New("missing required argument 'PolicyDefinitionId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &AssignmentArgs{}
	}
	var resource Assignment
	err := ctx.RegisterResource("azure:policy/assignment:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure:policy/assignment:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
	// A description to use for this Policy Assignment. Changing this forces a new resource to be created.
	Description *string `pulumi:"description"`
	// A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
	EnforcementMode *bool `pulumi:"enforcementMode"`
	// An `identity` block.
	Identity *AssignmentIdentity `pulumi:"identity"`
	// The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
	Metadata *string `pulumi:"metadata"`
	// The name of the Policy Assignment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
	NotScopes []string `pulumi:"notScopes"`
	// Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
	Parameters *string `pulumi:"parameters"`
	// The ID of the Policy Definition to be applied at the specified Scope.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
}

type AssignmentState struct {
	// A description to use for this Policy Assignment. Changing this forces a new resource to be created.
	Description pulumi.StringPtrInput
	// A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
	EnforcementMode pulumi.BoolPtrInput
	// An `identity` block.
	Identity AssignmentIdentityPtrInput
	// The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
	Metadata pulumi.StringPtrInput
	// The name of the Policy Assignment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
	NotScopes pulumi.StringArrayInput
	// Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
	Parameters pulumi.StringPtrInput
	// The ID of the Policy Definition to be applied at the specified Scope.
	PolicyDefinitionId pulumi.StringPtrInput
	// The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	// A description to use for this Policy Assignment. Changing this forces a new resource to be created.
	Description *string `pulumi:"description"`
	// A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
	EnforcementMode *bool `pulumi:"enforcementMode"`
	// An `identity` block.
	Identity *AssignmentIdentity `pulumi:"identity"`
	// The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
	Metadata *string `pulumi:"metadata"`
	// The name of the Policy Assignment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
	NotScopes []string `pulumi:"notScopes"`
	// Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
	Parameters *string `pulumi:"parameters"`
	// The ID of the Policy Definition to be applied at the specified Scope.
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// A description to use for this Policy Assignment. Changing this forces a new resource to be created.
	Description pulumi.StringPtrInput
	// A friendly display name to use for this Policy Assignment. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// Can be set to 'true' or 'false' to control whether the assignment is enforced (true) or not (false). Default is 'true'.
	EnforcementMode pulumi.BoolPtrInput
	// An `identity` block.
	Identity AssignmentIdentityPtrInput
	// The Azure location where this policy assignment should exist. This is required when an Identity is assigned. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The metadata for the policy assignment. This is a JSON string representing additional metadata that should be stored with the policy assignment.
	Metadata pulumi.StringPtrInput
	// The name of the Policy Assignment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of the Policy Assignment's excluded scopes. The list must contain Resource IDs (such as Subscriptions e.g. `/subscriptions/00000000-0000-0000-000000000000` or Resource Groups e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`).
	NotScopes pulumi.StringArrayInput
	// Parameters for the policy definition. This field is a JSON string that maps to the Parameters field from the Policy Definition. Changing this forces a new resource to be created.
	Parameters pulumi.StringPtrInput
	// The ID of the Policy Definition to be applied at the specified Scope.
	PolicyDefinitionId pulumi.StringInput
	// The Scope at which the Policy Assignment should be applied, which must be a Resource ID (such as Subscription e.g. `/subscriptions/00000000-0000-0000-000000000000` or a Resource Group e.g.`/subscriptions/00000000-0000-0000-000000000000/resourceGroups/myResourceGroup`). Changing this forces a new resource to be created.
	Scope pulumi.StringInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}
