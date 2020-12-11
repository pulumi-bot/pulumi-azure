// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Monitor Action Rule which type is suppression.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/monitoring"
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
// 		_, err = monitoring.NewActionRuleSuppression(ctx, "exampleActionRuleSuppression", &monitoring.ActionRuleSuppressionArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Scope: &monitoring.ActionRuleSuppressionScopeArgs{
// 				Type: pulumi.String("ResourceGroup"),
// 				ResourceIds: pulumi.StringArray{
// 					exampleResourceGroup.ID(),
// 				},
// 			},
// 			Suppression: &monitoring.ActionRuleSuppressionSuppressionArgs{
// 				RecurrenceType: pulumi.String("Weekly"),
// 				Schedule: &monitoring.ActionRuleSuppressionSuppressionScheduleArgs{
// 					StartDateUtc: pulumi.String("2019-01-01T01:02:03Z"),
// 					EndDateUtc:   pulumi.String("2019-01-03T15:02:07Z"),
// 					RecurrenceWeeklies: pulumi.StringArray{
// 						pulumi.String("Sunday"),
// 						pulumi.String("Monday"),
// 						pulumi.String("Friday"),
// 						pulumi.String("Saturday"),
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
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
// Monitor Action Rule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/actionRuleSuppression:ActionRuleSuppression example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
// ```
type ActionRuleSuppression struct {
	pulumi.CustomResourceState

	// A `condition` block as defined below.
	Condition ActionRuleSuppressionConditionPtrOutput `pulumi:"condition"`
	// Specifies a description for the Action Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `scope` block as defined below.
	Scope ActionRuleSuppressionScopePtrOutput `pulumi:"scope"`
	// A `suppression` block as defined below.
	Suppression ActionRuleSuppressionSuppressionOutput `pulumi:"suppression"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewActionRuleSuppression registers a new resource with the given unique name, arguments, and options.
func NewActionRuleSuppression(ctx *pulumi.Context,
	name string, args *ActionRuleSuppressionArgs, opts ...pulumi.ResourceOption) (*ActionRuleSuppression, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Suppression == nil {
		return nil, errors.New("invalid value for required argument 'Suppression'")
	}
	var resource ActionRuleSuppression
	err := ctx.RegisterResource("azure:monitoring/actionRuleSuppression:ActionRuleSuppression", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionRuleSuppression gets an existing ActionRuleSuppression resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionRuleSuppression(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionRuleSuppressionState, opts ...pulumi.ResourceOption) (*ActionRuleSuppression, error) {
	var resource ActionRuleSuppression
	err := ctx.ReadResource("azure:monitoring/actionRuleSuppression:ActionRuleSuppression", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionRuleSuppression resources.
type actionRuleSuppressionState struct {
	// A `condition` block as defined below.
	Condition *ActionRuleSuppressionCondition `pulumi:"condition"`
	// Specifies a description for the Action Rule.
	Description *string `pulumi:"description"`
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `scope` block as defined below.
	Scope *ActionRuleSuppressionScope `pulumi:"scope"`
	// A `suppression` block as defined below.
	Suppression *ActionRuleSuppressionSuppression `pulumi:"suppression"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ActionRuleSuppressionState struct {
	// A `condition` block as defined below.
	Condition ActionRuleSuppressionConditionPtrInput
	// Specifies a description for the Action Rule.
	Description pulumi.StringPtrInput
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `scope` block as defined below.
	Scope ActionRuleSuppressionScopePtrInput
	// A `suppression` block as defined below.
	Suppression ActionRuleSuppressionSuppressionPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActionRuleSuppressionState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleSuppressionState)(nil)).Elem()
}

type actionRuleSuppressionArgs struct {
	// A `condition` block as defined below.
	Condition *ActionRuleSuppressionCondition `pulumi:"condition"`
	// Specifies a description for the Action Rule.
	Description *string `pulumi:"description"`
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `scope` block as defined below.
	Scope *ActionRuleSuppressionScope `pulumi:"scope"`
	// A `suppression` block as defined below.
	Suppression ActionRuleSuppressionSuppression `pulumi:"suppression"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ActionRuleSuppression resource.
type ActionRuleSuppressionArgs struct {
	// A `condition` block as defined below.
	Condition ActionRuleSuppressionConditionPtrInput
	// Specifies a description for the Action Rule.
	Description pulumi.StringPtrInput
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `scope` block as defined below.
	Scope ActionRuleSuppressionScopePtrInput
	// A `suppression` block as defined below.
	Suppression ActionRuleSuppressionSuppressionInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActionRuleSuppressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleSuppressionArgs)(nil)).Elem()
}

type ActionRuleSuppressionInput interface {
	pulumi.Input

	ToActionRuleSuppressionOutput() ActionRuleSuppressionOutput
	ToActionRuleSuppressionOutputWithContext(ctx context.Context) ActionRuleSuppressionOutput
}

type ActionRuleSuppressionPtrInput interface {
	pulumi.Input

	ToActionRuleSuppressionPtrOutput() ActionRuleSuppressionPtrOutput
	ToActionRuleSuppressionPtrOutputWithContext(ctx context.Context) ActionRuleSuppressionPtrOutput
}

func (ActionRuleSuppression) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleSuppression)(nil)).Elem()
}

func (i ActionRuleSuppression) ToActionRuleSuppressionOutput() ActionRuleSuppressionOutput {
	return i.ToActionRuleSuppressionOutputWithContext(context.Background())
}

func (i ActionRuleSuppression) ToActionRuleSuppressionOutputWithContext(ctx context.Context) ActionRuleSuppressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleSuppressionOutput)
}

func (i ActionRuleSuppression) ToActionRuleSuppressionPtrOutput() ActionRuleSuppressionPtrOutput {
	return i.ToActionRuleSuppressionPtrOutputWithContext(context.Background())
}

func (i ActionRuleSuppression) ToActionRuleSuppressionPtrOutputWithContext(ctx context.Context) ActionRuleSuppressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleSuppressionPtrOutput)
}

type ActionRuleSuppressionOutput struct {
	*pulumi.OutputState
}

func (ActionRuleSuppressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleSuppressionOutput)(nil)).Elem()
}

func (o ActionRuleSuppressionOutput) ToActionRuleSuppressionOutput() ActionRuleSuppressionOutput {
	return o
}

func (o ActionRuleSuppressionOutput) ToActionRuleSuppressionOutputWithContext(ctx context.Context) ActionRuleSuppressionOutput {
	return o
}

type ActionRuleSuppressionPtrOutput struct {
	*pulumi.OutputState
}

func (ActionRuleSuppressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleSuppression)(nil)).Elem()
}

func (o ActionRuleSuppressionPtrOutput) ToActionRuleSuppressionPtrOutput() ActionRuleSuppressionPtrOutput {
	return o
}

func (o ActionRuleSuppressionPtrOutput) ToActionRuleSuppressionPtrOutputWithContext(ctx context.Context) ActionRuleSuppressionPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActionRuleSuppressionOutput{})
	pulumi.RegisterOutputType(ActionRuleSuppressionPtrOutput{})
}
