// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Automation Schedule.
type Schedule struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// A description for this Schedule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The end time of the schedule.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
	MonthDays pulumi.IntArrayOutput `pulumi:"monthDays"`
	// List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthlyOccurrence` block supports fields documented below.
	MonthlyOccurrences ScheduleMonthlyOccurrenceArrayOutput `pulumi:"monthlyOccurrences"`
	// Specifies the name of the Schedule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The timezone of the start time. Defaults to `UTC`. For possible values see: https://s2.automation.ext.azure.com/api/Orchestrator/TimeZones?_=1594792230258
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// List of days of the week that the job should execute on. Only valid when frequency is `Week`.
	WeekDays pulumi.StringArrayOutput `pulumi:"weekDays"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Schedule
	err := ctx.RegisterResource("azure:automation/schedule:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure:automation/schedule:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
	// The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// A description for this Schedule.
	Description *string `pulumi:"description"`
	// The end time of the schedule.
	ExpiryTime *string `pulumi:"expiryTime"`
	// The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
	Frequency *string `pulumi:"frequency"`
	// The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
	Interval *int `pulumi:"interval"`
	// List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
	MonthDays []int `pulumi:"monthDays"`
	// List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthlyOccurrence` block supports fields documented below.
	MonthlyOccurrences []ScheduleMonthlyOccurrence `pulumi:"monthlyOccurrences"`
	// Specifies the name of the Schedule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
	StartTime *string `pulumi:"startTime"`
	// The timezone of the start time. Defaults to `UTC`. For possible values see: https://s2.automation.ext.azure.com/api/Orchestrator/TimeZones?_=1594792230258
	Timezone *string `pulumi:"timezone"`
	// List of days of the week that the job should execute on. Only valid when frequency is `Week`.
	WeekDays []string `pulumi:"weekDays"`
}

type ScheduleState struct {
	// The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// A description for this Schedule.
	Description pulumi.StringPtrInput
	// The end time of the schedule.
	ExpiryTime pulumi.StringPtrInput
	// The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
	Frequency pulumi.StringPtrInput
	// The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
	Interval pulumi.IntPtrInput
	// List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
	MonthDays pulumi.IntArrayInput
	// List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthlyOccurrence` block supports fields documented below.
	MonthlyOccurrences ScheduleMonthlyOccurrenceArrayInput
	// Specifies the name of the Schedule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
	StartTime pulumi.StringPtrInput
	// The timezone of the start time. Defaults to `UTC`. For possible values see: https://s2.automation.ext.azure.com/api/Orchestrator/TimeZones?_=1594792230258
	Timezone pulumi.StringPtrInput
	// List of days of the week that the job should execute on. Only valid when frequency is `Week`.
	WeekDays pulumi.StringArrayInput
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	// The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// A description for this Schedule.
	Description *string `pulumi:"description"`
	// The end time of the schedule.
	ExpiryTime *string `pulumi:"expiryTime"`
	// The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
	Frequency string `pulumi:"frequency"`
	// The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
	Interval *int `pulumi:"interval"`
	// List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
	MonthDays []int `pulumi:"monthDays"`
	// List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthlyOccurrence` block supports fields documented below.
	MonthlyOccurrences []ScheduleMonthlyOccurrence `pulumi:"monthlyOccurrences"`
	// Specifies the name of the Schedule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
	StartTime *string `pulumi:"startTime"`
	// The timezone of the start time. Defaults to `UTC`. For possible values see: https://s2.automation.ext.azure.com/api/Orchestrator/TimeZones?_=1594792230258
	Timezone *string `pulumi:"timezone"`
	// List of days of the week that the job should execute on. Only valid when frequency is `Week`.
	WeekDays []string `pulumi:"weekDays"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// The name of the automation account in which the Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// A description for this Schedule.
	Description pulumi.StringPtrInput
	// The end time of the schedule.
	ExpiryTime pulumi.StringPtrInput
	// The frequency of the schedule. - can be either `OneTime`, `Day`, `Hour`, `Week`, or `Month`.
	Frequency pulumi.StringInput
	// The number of `frequency`s between runs. Only valid when frequency is `Day`, `Hour`, `Week`, or `Month` and defaults to `1`.
	Interval pulumi.IntPtrInput
	// List of days of the month that the job should execute on. Must be between `1` and `31`. `-1` for last day of the month. Only valid when frequency is `Month`.
	MonthDays pulumi.IntArrayInput
	// List of occurrences of days within a month. Only valid when frequency is `Month`. The `monthlyOccurrence` block supports fields documented below.
	MonthlyOccurrences ScheduleMonthlyOccurrenceArrayInput
	// Specifies the name of the Schedule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Start time of the schedule. Must be at least five minutes in the future. Defaults to seven minutes in the future from the time the resource is created.
	StartTime pulumi.StringPtrInput
	// The timezone of the start time. Defaults to `UTC`. For possible values see: https://s2.automation.ext.azure.com/api/Orchestrator/TimeZones?_=1594792230258
	Timezone pulumi.StringPtrInput
	// List of days of the week that the job should execute on. Only valid when frequency is `Week`.
	WeekDays pulumi.StringArrayInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}
