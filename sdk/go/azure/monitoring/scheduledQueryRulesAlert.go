// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AlertingAction Scheduled Query Rules resource within Azure Monitor.
//
// ## Import
//
// Scheduled Query Rule Alerts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/scheduledQueryRulesAlert:ScheduledQueryRulesAlert example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
// ```
type ScheduledQueryRulesAlert struct {
	pulumi.CustomResourceState

	// An `action` block as defined below.
	Action ScheduledQueryRulesAlertActionOutput `pulumi:"action"`
	// List of Resource IDs referred into query.
	AuthorizedResourceIds pulumi.StringArrayOutput `pulumi:"authorizedResourceIds"`
	// The resource URI over which log search query is to be run.
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
	Frequency pulumi.IntOutput    `pulumi:"frequency"`
	Location  pulumi.StringOutput `pulumi:"location"`
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Log search query.
	Query     pulumi.StringOutput    `pulumi:"query"`
	QueryType pulumi.StringPtrOutput `pulumi:"queryType"`
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
	Severity pulumi.IntPtrOutput    `pulumi:"severity"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	// Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
	Throttling pulumi.IntPtrOutput `pulumi:"throttling"`
	// Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
	TimeWindow pulumi.IntOutput `pulumi:"timeWindow"`
	// The condition that results in the alert rule being run.
	Trigger ScheduledQueryRulesAlertTriggerOutput `pulumi:"trigger"`
}

// NewScheduledQueryRulesAlert registers a new resource with the given unique name, arguments, and options.
func NewScheduledQueryRulesAlert(ctx *pulumi.Context,
	name string, args *ScheduledQueryRulesAlertArgs, opts ...pulumi.ResourceOption) (*ScheduledQueryRulesAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.DataSourceId == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceId'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TimeWindow == nil {
		return nil, errors.New("invalid value for required argument 'TimeWindow'")
	}
	if args.Trigger == nil {
		return nil, errors.New("invalid value for required argument 'Trigger'")
	}
	var resource ScheduledQueryRulesAlert
	err := ctx.RegisterResource("azure:monitoring/scheduledQueryRulesAlert:ScheduledQueryRulesAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledQueryRulesAlert gets an existing ScheduledQueryRulesAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledQueryRulesAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRulesAlertState, opts ...pulumi.ResourceOption) (*ScheduledQueryRulesAlert, error) {
	var resource ScheduledQueryRulesAlert
	err := ctx.ReadResource("azure:monitoring/scheduledQueryRulesAlert:ScheduledQueryRulesAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledQueryRulesAlert resources.
type scheduledQueryRulesAlertState struct {
	// An `action` block as defined below.
	Action *ScheduledQueryRulesAlertAction `pulumi:"action"`
	// List of Resource IDs referred into query.
	AuthorizedResourceIds []string `pulumi:"authorizedResourceIds"`
	// The resource URI over which log search query is to be run.
	DataSourceId *string `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
	Frequency *int    `pulumi:"frequency"`
	Location  *string `pulumi:"location"`
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Log search query.
	Query     *string `pulumi:"query"`
	QueryType *string `pulumi:"queryType"`
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
	Severity *int              `pulumi:"severity"`
	Tags     map[string]string `pulumi:"tags"`
	// Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
	Throttling *int `pulumi:"throttling"`
	// Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
	TimeWindow *int `pulumi:"timeWindow"`
	// The condition that results in the alert rule being run.
	Trigger *ScheduledQueryRulesAlertTrigger `pulumi:"trigger"`
}

type ScheduledQueryRulesAlertState struct {
	// An `action` block as defined below.
	Action ScheduledQueryRulesAlertActionPtrInput
	// List of Resource IDs referred into query.
	AuthorizedResourceIds pulumi.StringArrayInput
	// The resource URI over which log search query is to be run.
	DataSourceId pulumi.StringPtrInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled pulumi.BoolPtrInput
	// Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
	Frequency pulumi.IntPtrInput
	Location  pulumi.StringPtrInput
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Log search query.
	Query     pulumi.StringPtrInput
	QueryType pulumi.StringPtrInput
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName pulumi.StringPtrInput
	// Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
	Severity pulumi.IntPtrInput
	Tags     pulumi.StringMapInput
	// Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
	Throttling pulumi.IntPtrInput
	// Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
	TimeWindow pulumi.IntPtrInput
	// The condition that results in the alert rule being run.
	Trigger ScheduledQueryRulesAlertTriggerPtrInput
}

func (ScheduledQueryRulesAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRulesAlertState)(nil)).Elem()
}

type scheduledQueryRulesAlertArgs struct {
	// An `action` block as defined below.
	Action ScheduledQueryRulesAlertAction `pulumi:"action"`
	// List of Resource IDs referred into query.
	AuthorizedResourceIds []string `pulumi:"authorizedResourceIds"`
	// The resource URI over which log search query is to be run.
	DataSourceId string `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
	Frequency int     `pulumi:"frequency"`
	Location  *string `pulumi:"location"`
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Log search query.
	Query     string  `pulumi:"query"`
	QueryType *string `pulumi:"queryType"`
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
	Severity *int              `pulumi:"severity"`
	Tags     map[string]string `pulumi:"tags"`
	// Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
	Throttling *int `pulumi:"throttling"`
	// Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
	TimeWindow int `pulumi:"timeWindow"`
	// The condition that results in the alert rule being run.
	Trigger ScheduledQueryRulesAlertTrigger `pulumi:"trigger"`
}

// The set of arguments for constructing a ScheduledQueryRulesAlert resource.
type ScheduledQueryRulesAlertArgs struct {
	// An `action` block as defined below.
	Action ScheduledQueryRulesAlertActionInput
	// List of Resource IDs referred into query.
	AuthorizedResourceIds pulumi.StringArrayInput
	// The resource URI over which log search query is to be run.
	DataSourceId pulumi.StringInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled pulumi.BoolPtrInput
	// Frequency (in minutes) at which rule condition should be evaluated.  Values must be between 5 and 1440 (inclusive).
	Frequency pulumi.IntInput
	Location  pulumi.StringPtrInput
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Log search query.
	Query     pulumi.StringInput
	QueryType pulumi.StringPtrInput
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName pulumi.StringInput
	// Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
	Severity pulumi.IntPtrInput
	Tags     pulumi.StringMapInput
	// Time (in minutes) for which Alerts should be throttled or suppressed.  Values must be between 0 and 10000 (inclusive).
	Throttling pulumi.IntPtrInput
	// Time window for which data needs to be fetched for query (must be greater than or equal to `frequency`).  Values must be between 5 and 2880 (inclusive).
	TimeWindow pulumi.IntInput
	// The condition that results in the alert rule being run.
	Trigger ScheduledQueryRulesAlertTriggerInput
}

func (ScheduledQueryRulesAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRulesAlertArgs)(nil)).Elem()
}

type ScheduledQueryRulesAlertInput interface {
	pulumi.Input

	ToScheduledQueryRulesAlertOutput() ScheduledQueryRulesAlertOutput
	ToScheduledQueryRulesAlertOutputWithContext(ctx context.Context) ScheduledQueryRulesAlertOutput
}

func (*ScheduledQueryRulesAlert) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRulesAlert)(nil))
}

func (i *ScheduledQueryRulesAlert) ToScheduledQueryRulesAlertOutput() ScheduledQueryRulesAlertOutput {
	return i.ToScheduledQueryRulesAlertOutputWithContext(context.Background())
}

func (i *ScheduledQueryRulesAlert) ToScheduledQueryRulesAlertOutputWithContext(ctx context.Context) ScheduledQueryRulesAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesAlertOutput)
}

func (i *ScheduledQueryRulesAlert) ToScheduledQueryRulesAlertPtrOutput() ScheduledQueryRulesAlertPtrOutput {
	return i.ToScheduledQueryRulesAlertPtrOutputWithContext(context.Background())
}

func (i *ScheduledQueryRulesAlert) ToScheduledQueryRulesAlertPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesAlertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesAlertPtrOutput)
}

type ScheduledQueryRulesAlertPtrInput interface {
	pulumi.Input

	ToScheduledQueryRulesAlertPtrOutput() ScheduledQueryRulesAlertPtrOutput
	ToScheduledQueryRulesAlertPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesAlertPtrOutput
}

type ScheduledQueryRulesAlertOutput struct {
	*pulumi.OutputState
}

func (ScheduledQueryRulesAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRulesAlert)(nil))
}

func (o ScheduledQueryRulesAlertOutput) ToScheduledQueryRulesAlertOutput() ScheduledQueryRulesAlertOutput {
	return o
}

func (o ScheduledQueryRulesAlertOutput) ToScheduledQueryRulesAlertOutputWithContext(ctx context.Context) ScheduledQueryRulesAlertOutput {
	return o
}

type ScheduledQueryRulesAlertPtrOutput struct {
	*pulumi.OutputState
}

func (ScheduledQueryRulesAlertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRulesAlert)(nil))
}

func (o ScheduledQueryRulesAlertPtrOutput) ToScheduledQueryRulesAlertPtrOutput() ScheduledQueryRulesAlertPtrOutput {
	return o
}

func (o ScheduledQueryRulesAlertPtrOutput) ToScheduledQueryRulesAlertPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesAlertPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScheduledQueryRulesAlertOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRulesAlertPtrOutput{})
}
