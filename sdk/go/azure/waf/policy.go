// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Web Application Firewall Policy instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/waf"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US 2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = waf.NewPolicy(ctx, "examplePolicy", &waf.PolicyArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			CustomRules: waf.PolicyCustomRuleArray{
// 				&waf.PolicyCustomRuleArgs{
// 					Name:     pulumi.String("Rule1"),
// 					Priority: pulumi.Int(1),
// 					RuleType: pulumi.String("MatchRule"),
// 					MatchConditions: waf.PolicyCustomRuleMatchConditionArray{
// 						&waf.PolicyCustomRuleMatchConditionArgs{
// 							MatchVariables: waf.PolicyCustomRuleMatchConditionMatchVariableArray{
// 								&waf.PolicyCustomRuleMatchConditionMatchVariableArgs{
// 									VariableName: pulumi.String("RemoteAddr"),
// 								},
// 							},
// 							Operator:          pulumi.String("IPMatch"),
// 							NegationCondition: pulumi.Bool(false),
// 							MatchValues: pulumi.StringArray{
// 								pulumi.String("192.168.1.0/24"),
// 								pulumi.String("10.0.0.0/24"),
// 							},
// 						},
// 					},
// 					Action: pulumi.String("Block"),
// 				},
// 				&waf.PolicyCustomRuleArgs{
// 					Name:     pulumi.String("Rule2"),
// 					Priority: pulumi.Int(2),
// 					RuleType: pulumi.String("MatchRule"),
// 					MatchConditions: waf.PolicyCustomRuleMatchConditionArray{
// 						&waf.PolicyCustomRuleMatchConditionArgs{
// 							MatchVariables: waf.PolicyCustomRuleMatchConditionMatchVariableArray{
// 								&waf.PolicyCustomRuleMatchConditionMatchVariableArgs{
// 									VariableName: pulumi.String("RemoteAddr"),
// 								},
// 							},
// 							Operator:          pulumi.String("IPMatch"),
// 							NegationCondition: pulumi.Bool(false),
// 							MatchValues: pulumi.StringArray{
// 								pulumi.String("192.168.1.0/24"),
// 							},
// 						},
// 						&waf.PolicyCustomRuleMatchConditionArgs{
// 							MatchVariables: waf.PolicyCustomRuleMatchConditionMatchVariableArray{
// 								&waf.PolicyCustomRuleMatchConditionMatchVariableArgs{
// 									VariableName: pulumi.String("RequestHeaders"),
// 									Selector:     pulumi.String("UserAgent"),
// 								},
// 							},
// 							Operator:          pulumi.String("Contains"),
// 							NegationCondition: pulumi.Bool(false),
// 							MatchValues: pulumi.StringArray{
// 								pulumi.String("Windows"),
// 							},
// 						},
// 					},
// 					Action: pulumi.String("Block"),
// 				},
// 			},
// 			PolicySettings: &waf.PolicyPolicySettingsArgs{
// 				Enabled:                pulumi.Bool(true),
// 				Mode:                   pulumi.String("Prevention"),
// 				RequestBodyCheck:       pulumi.Bool(true),
// 				FileUploadLimitInMb:    pulumi.Int(100),
// 				MaxRequestBodySizeInKb: pulumi.Int(128),
// 			},
// 			ManagedRules: &waf.PolicyManagedRulesArgs{
// 				Exclusions: waf.PolicyManagedRulesExclusionArray{
// 					&waf.PolicyManagedRulesExclusionArgs{
// 						MatchVariable:         pulumi.String("RequestHeaderNames"),
// 						Selector:              pulumi.String("x-company-secret-header"),
// 						SelectorMatchOperator: pulumi.String("Equals"),
// 					},
// 					&waf.PolicyManagedRulesExclusionArgs{
// 						MatchVariable:         pulumi.String("RequestCookieNames"),
// 						Selector:              pulumi.String("too-tasty"),
// 						SelectorMatchOperator: pulumi.String("EndsWith"),
// 					},
// 				},
// 				ManagedRuleSets: waf.PolicyManagedRulesManagedRuleSetArray{
// 					&waf.PolicyManagedRulesManagedRuleSetArgs{
// 						Type:    pulumi.String("OWASP"),
// 						Version: pulumi.String("3.1"),
// 						RuleGroupOverrides: waf.PolicyManagedRulesManagedRuleSetRuleGroupOverrideArray{
// 							&waf.PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs{
// 								RuleGroupName: pulumi.String("REQUEST-920-PROTOCOL-ENFORCEMENT"),
// 								DisabledRules: pulumi.StringArray{
// 									pulumi.String("920300"),
// 									pulumi.String("920440"),
// 								},
// 							},
// 						},
// 					},
// 				},
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
// Web Application Firewall Policy can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:waf/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/example-wafpolicy
// ```
type Policy struct {
	pulumi.CustomResourceState

	// One or more `customRules` blocks as defined below.
	CustomRules PolicyCustomRuleArrayOutput `pulumi:"customRules"`
	// Resource location. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `managedRules` blocks as defined below.
	ManagedRules PolicyManagedRulesOutput `pulumi:"managedRules"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `policySettings` block as defined below.
	PolicySettings PolicyPolicySettingsPtrOutput `pulumi:"policySettings"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedRules == nil {
		return nil, errors.New("invalid value for required argument 'ManagedRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Policy
	err := ctx.RegisterResource("azure:waf/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure:waf/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// One or more `customRules` blocks as defined below.
	CustomRules []PolicyCustomRule `pulumi:"customRules"`
	// Resource location. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `managedRules` blocks as defined below.
	ManagedRules *PolicyManagedRules `pulumi:"managedRules"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `policySettings` block as defined below.
	PolicySettings *PolicyPolicySettings `pulumi:"policySettings"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

type PolicyState struct {
	// One or more `customRules` blocks as defined below.
	CustomRules PolicyCustomRuleArrayInput
	// Resource location. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `managedRules` blocks as defined below.
	ManagedRules PolicyManagedRulesPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `policySettings` block as defined below.
	PolicySettings PolicyPolicySettingsPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// One or more `customRules` blocks as defined below.
	CustomRules []PolicyCustomRule `pulumi:"customRules"`
	// Resource location. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `managedRules` blocks as defined below.
	ManagedRules PolicyManagedRules `pulumi:"managedRules"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `policySettings` block as defined below.
	PolicySettings *PolicyPolicySettings `pulumi:"policySettings"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// One or more `customRules` blocks as defined below.
	CustomRules PolicyCustomRuleArrayInput
	// Resource location. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `managedRules` blocks as defined below.
	ManagedRules PolicyManagedRulesInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `policySettings` block as defined below.
	PolicySettings PolicyPolicySettingsPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput
}

func (Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil)).Elem()
}

func (i Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i Policy) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i Policy) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyOutput)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

type PolicyPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
}
