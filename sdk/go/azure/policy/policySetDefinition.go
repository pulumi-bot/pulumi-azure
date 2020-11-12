// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a policy set definition.
//
// > **NOTE:**  Policy set definitions (also known as policy initiatives) do not take effect until they are assigned to a scope using a Policy Set Assignment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/policy"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := policy.NewPolicySetDefinition(ctx, "example", &policy.PolicySetDefinitionArgs{
// 			DisplayName: pulumi.String("Test Policy Set"),
// 			Parameters:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "    {\n", "        \"allowedLocations\": {\n", "            \"type\": \"Array\",\n", "            \"metadata\": {\n", "                \"description\": \"The list of allowed locations for resources.\",\n", "                \"displayName\": \"Allowed locations\",\n", "                \"strongType\": \"location\"\n", "            }\n", "        }\n", "    }\n", "\n")),
// 			PolicyDefinitionReferences: policy.PolicySetDefinitionPolicyDefinitionReferenceArray{
// 				&policy.PolicySetDefinitionPolicyDefinitionReferenceArgs{
// 					ParameterValues:    pulumi.String(fmt.Sprintf("%v%v%v%v", "    {\n", "      \"listOfAllowedLocations\": {\"value\": \"[parameters('allowedLocations')]\"}\n", "    }\n", "    \n")),
// 					PolicyDefinitionId: pulumi.String("/providers/Microsoft.Authorization/policyDefinitions/e765b5de-1225-4ba3-bd56-1ac6695af988"),
// 				},
// 			},
// 			PolicyType: pulumi.String("Custom"),
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
// Policy Set Definitions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:policy/policySetDefinition:PolicySetDefinition example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/testPolicySet
// ```
//
//  or
//
// ```sh
//  $ pulumi import azure:policy/policySetDefinition:PolicySetDefinition example /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.Authorization/policySetDefinitions/testPolicySet
// ```
type PolicySetDefinition struct {
	pulumi.CustomResourceState

	// The description of the policy set definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId pulumi.StringOutput `pulumi:"managementGroupId"`
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	ManagementGroupName pulumi.StringOutput `pulumi:"managementGroupName"`
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// One or more `policyDefinitionReference` blocks as defined below.
	PolicyDefinitionReferences PolicySetDefinitionPolicyDefinitionReferenceArrayOutput `pulumi:"policyDefinitionReferences"`
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
	//
	// Deprecated: Deprecated in favor of `policy_definition_reference`
	PolicyDefinitions pulumi.StringOutput `pulumi:"policyDefinitions"`
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
}

// NewPolicySetDefinition registers a new resource with the given unique name, arguments, and options.
func NewPolicySetDefinition(ctx *pulumi.Context,
	name string, args *PolicySetDefinitionArgs, opts ...pulumi.ResourceOption) (*PolicySetDefinition, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.PolicyType == nil {
		return nil, errors.New("missing required argument 'PolicyType'")
	}
	if args == nil {
		args = &PolicySetDefinitionArgs{}
	}
	var resource PolicySetDefinition
	err := ctx.RegisterResource("azure:policy/policySetDefinition:PolicySetDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicySetDefinition gets an existing PolicySetDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicySetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetDefinitionState, opts ...pulumi.ResourceOption) (*PolicySetDefinition, error) {
	var resource PolicySetDefinition
	err := ctx.ReadResource("azure:policy/policySetDefinition:PolicySetDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicySetDefinition resources.
type policySetDefinitionState struct {
	// The description of the policy set definition.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	ManagementGroupName *string `pulumi:"managementGroupName"`
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// One or more `policyDefinitionReference` blocks as defined below.
	PolicyDefinitionReferences []PolicySetDefinitionPolicyDefinitionReference `pulumi:"policyDefinitionReferences"`
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
	//
	// Deprecated: Deprecated in favor of `policy_definition_reference`
	PolicyDefinitions *string `pulumi:"policyDefinitions"`
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType *string `pulumi:"policyType"`
}

type PolicySetDefinitionState struct {
	// The description of the policy set definition.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrInput
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId pulumi.StringPtrInput
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	ManagementGroupName pulumi.StringPtrInput
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata pulumi.StringPtrInput
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// One or more `policyDefinitionReference` blocks as defined below.
	PolicyDefinitionReferences PolicySetDefinitionPolicyDefinitionReferenceArrayInput
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
	//
	// Deprecated: Deprecated in favor of `policy_definition_reference`
	PolicyDefinitions pulumi.StringPtrInput
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringPtrInput
}

func (PolicySetDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionState)(nil)).Elem()
}

type policySetDefinitionArgs struct {
	// The description of the policy set definition.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName string `pulumi:"displayName"`
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	ManagementGroupName *string `pulumi:"managementGroupName"`
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// One or more `policyDefinitionReference` blocks as defined below.
	PolicyDefinitionReferences []PolicySetDefinitionPolicyDefinitionReference `pulumi:"policyDefinitionReferences"`
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
	//
	// Deprecated: Deprecated in favor of `policy_definition_reference`
	PolicyDefinitions *string `pulumi:"policyDefinitions"`
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType string `pulumi:"policyType"`
}

// The set of arguments for constructing a PolicySetDefinition resource.
type PolicySetDefinitionArgs struct {
	// The description of the policy set definition.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringInput
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId pulumi.StringPtrInput
	// The name of the Management Group where this policy set definition should be defined. Changing this forces a new resource to be created.
	ManagementGroupName pulumi.StringPtrInput
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata pulumi.StringPtrInput
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// One or more `policyDefinitionReference` blocks as defined below.
	PolicyDefinitionReferences PolicySetDefinitionPolicyDefinitionReferenceArrayInput
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions.
	//
	// Deprecated: Deprecated in favor of `policy_definition_reference`
	PolicyDefinitions pulumi.StringPtrInput
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringInput
}

func (PolicySetDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionArgs)(nil)).Elem()
}

type PolicySetDefinitionInput interface {
	pulumi.Input

	ToPolicySetDefinitionOutput() PolicySetDefinitionOutput
	ToPolicySetDefinitionOutputWithContext(ctx context.Context) PolicySetDefinitionOutput
}

func (PolicySetDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySetDefinition)(nil)).Elem()
}

func (i PolicySetDefinition) ToPolicySetDefinitionOutput() PolicySetDefinitionOutput {
	return i.ToPolicySetDefinitionOutputWithContext(context.Background())
}

func (i PolicySetDefinition) ToPolicySetDefinitionOutputWithContext(ctx context.Context) PolicySetDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetDefinitionOutput)
}

type PolicySetDefinitionOutput struct {
	*pulumi.OutputState
}

func (PolicySetDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySetDefinitionOutput)(nil)).Elem()
}

func (o PolicySetDefinitionOutput) ToPolicySetDefinitionOutput() PolicySetDefinitionOutput {
	return o
}

func (o PolicySetDefinitionOutput) ToPolicySetDefinitionOutputWithContext(ctx context.Context) PolicySetDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicySetDefinitionOutput{})
}
