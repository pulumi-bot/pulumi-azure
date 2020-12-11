// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Policy Remediation at the specified Scope.
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
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDefinition, err := policy.NewDefinition(ctx, "exampleDefinition", &policy.DefinitionArgs{
// 			PolicyType:  pulumi.String("Custom"),
// 			Mode:        pulumi.String("All"),
// 			DisplayName: pulumi.String("my-policy-definition"),
// 			PolicyRule:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "    {\n", "    \"if\": {\n", "      \"not\": {\n", "        \"field\": \"location\",\n", "        \"in\": \"[parameters('allowedLocations')]\"\n", "      }\n", "    },\n", "    \"then\": {\n", "      \"effect\": \"audit\"\n", "    }\n", "  }\n")),
// 			Parameters:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v", "    {\n", "    \"allowedLocations\": {\n", "      \"type\": \"Array\",\n", "      \"metadata\": {\n", "        \"description\": \"The list of allowed locations for resources.\",\n", "        \"displayName\": \"Allowed locations\",\n", "        \"strongType\": \"location\"\n", "      }\n", "    }\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := policy.NewAssignment(ctx, "exampleAssignment", &policy.AssignmentArgs{
// 			Scope:              exampleResourceGroup.ID(),
// 			PolicyDefinitionId: exampleDefinition.ID(),
// 			Description:        pulumi.String("Policy Assignment created via an Acceptance Test"),
// 			DisplayName:        pulumi.String("My Example Policy Assignment"),
// 			Parameters:         pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", "  \"allowedLocations\": {\n", "    \"value\": [ \"West Europe\" ]\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = policy.NewRemediation(ctx, "exampleRemediation", &policy.RemediationArgs{
// 			Scope:              exampleAssignment.Scope,
// 			PolicyAssignmentId: exampleAssignment.ID(),
// 			LocationFilters: pulumi.StringArray{
// 				pulumi.String("West Europe"),
// 			},
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
// Policy Remediations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights/remediations/remediation1
// ```
//
//  or
//
// ```sh
//  $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PolicyInsights/remediations/remediation1
// ```
//
//  or
//
// ```sh
//  $ pulumi import azure:policy/remediation:Remediation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.PolicyInsights/remediations/remediation1
// ```
//
//  or
//
// ```sh
//  $ pulumi import azure:policy/remediation:Remediation example /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.PolicyInsights/remediations/remediation1
// ```
type Remediation struct {
	pulumi.CustomResourceState

	// A list of the resource locations that will be remediated.
	LocationFilters pulumi.StringArrayOutput `pulumi:"locationFilters"`
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentId pulumi.StringOutput `pulumi:"policyAssignmentId"`
	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrOutput `pulumi:"policyDefinitionReferenceId"`
	// The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
	ResourceDiscoveryMode pulumi.StringPtrOutput `pulumi:"resourceDiscoveryMode"`
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewRemediation registers a new resource with the given unique name, arguments, and options.
func NewRemediation(ctx *pulumi.Context,
	name string, args *RemediationArgs, opts ...pulumi.ResourceOption) (*Remediation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyAssignmentId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource Remediation
	err := ctx.RegisterResource("azure:policy/remediation:Remediation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediation gets an existing Remediation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationState, opts ...pulumi.ResourceOption) (*Remediation, error) {
	var resource Remediation
	err := ctx.ReadResource("azure:policy/remediation:Remediation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Remediation resources.
type remediationState struct {
	// A list of the resource locations that will be remediated.
	LocationFilters []string `pulumi:"locationFilters"`
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
	ResourceDiscoveryMode *string `pulumi:"resourceDiscoveryMode"`
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope *string `pulumi:"scope"`
}

type RemediationState struct {
	// A list of the resource locations that will be remediated.
	LocationFilters pulumi.StringArrayInput
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrInput
	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
	ResourceDiscoveryMode pulumi.StringPtrInput
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope pulumi.StringPtrInput
}

func (RemediationState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationState)(nil)).Elem()
}

type remediationArgs struct {
	// A list of the resource locations that will be remediated.
	LocationFilters []string `pulumi:"locationFilters"`
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentId string `pulumi:"policyAssignmentId"`
	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
	ResourceDiscoveryMode *string `pulumi:"resourceDiscoveryMode"`
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a Remediation resource.
type RemediationArgs struct {
	// A list of the resource locations that will be remediated.
	LocationFilters pulumi.StringArrayInput
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentId pulumi.StringInput
	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The way that resources to remediate are discovered. Possible values are `ExistingNonCompliant`, `ReEvaluateCompliance`. Defaults to `ExistingNonCompliant`.
	ResourceDiscoveryMode pulumi.StringPtrInput
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope pulumi.StringInput
}

func (RemediationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationArgs)(nil)).Elem()
}

type RemediationInput interface {
	pulumi.Input

	ToRemediationOutput() RemediationOutput
	ToRemediationOutputWithContext(ctx context.Context) RemediationOutput
}

type RemediationPtrInput interface {
	pulumi.Input

	ToRemediationPtrOutput() RemediationPtrOutput
	ToRemediationPtrOutputWithContext(ctx context.Context) RemediationPtrOutput
}

func (Remediation) ElementType() reflect.Type {
	return reflect.TypeOf((*Remediation)(nil)).Elem()
}

func (i Remediation) ToRemediationOutput() RemediationOutput {
	return i.ToRemediationOutputWithContext(context.Background())
}

func (i Remediation) ToRemediationOutputWithContext(ctx context.Context) RemediationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationOutput)
}

func (i Remediation) ToRemediationPtrOutput() RemediationPtrOutput {
	return i.ToRemediationPtrOutputWithContext(context.Background())
}

func (i Remediation) ToRemediationPtrOutputWithContext(ctx context.Context) RemediationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPtrOutput)
}

type RemediationOutput struct {
	*pulumi.OutputState
}

func (RemediationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationOutput)(nil)).Elem()
}

func (o RemediationOutput) ToRemediationOutput() RemediationOutput {
	return o
}

func (o RemediationOutput) ToRemediationOutputWithContext(ctx context.Context) RemediationOutput {
	return o
}

type RemediationPtrOutput struct {
	*pulumi.OutputState
}

func (RemediationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Remediation)(nil)).Elem()
}

func (o RemediationPtrOutput) ToRemediationPtrOutput() RemediationPtrOutput {
	return o
}

func (o RemediationPtrOutput) ToRemediationPtrOutputWithContext(ctx context.Context) RemediationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RemediationOutput{})
	pulumi.RegisterOutputType(RemediationPtrOutput{})
}
