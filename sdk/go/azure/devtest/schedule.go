// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages automated startup and shutdown schedules for Azure Dev Test Lab.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/devtest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLab, err := devtest.NewLab(ctx, "exampleLab", &devtest.LabArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = devtest.NewSchedule(ctx, "exampleSchedule", &devtest.ScheduleArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			LabName:           exampleLab.Name,
// 			WeeklyRecurrence: &devtest.ScheduleWeeklyRecurrenceArgs{
// 				Time: pulumi.String("1100"),
// 				WeekDays: pulumi.StringArray{
// 					pulumi.String("Monday"),
// 					pulumi.String("Tuesday"),
// 				},
// 			},
// 			TimeZoneId:           pulumi.String("Pacific Standard Time"),
// 			TaskType:             pulumi.String("LabVmsStartupTask"),
// 			NotificationSettings: nil,
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// DevTest Schedule's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:devtest/schedule:Schedule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DevTestLab/labs/myDevTestLab/schedules/labvmautostart
// ```
type Schedule struct {
	pulumi.CustomResourceState

	DailyRecurrence  ScheduleDailyRecurrencePtrOutput  `pulumi:"dailyRecurrence"`
	HourlyRecurrence ScheduleHourlyRecurrencePtrOutput `pulumi:"hourlyRecurrence"`
	// The name of the dev test lab. Changing this forces a new resource to be created.
	LabName pulumi.StringOutput `pulumi:"labName"`
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
	Name                 pulumi.StringOutput                `pulumi:"name"`
	NotificationSettings ScheduleNotificationSettingsOutput `pulumi:"notificationSettings"`
	// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// The time zone ID (e.g. Pacific Standard time).
	TimeZoneId       pulumi.StringOutput               `pulumi:"timeZoneId"`
	WeeklyRecurrence ScheduleWeeklyRecurrencePtrOutput `pulumi:"weeklyRecurrence"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.NotificationSettings == nil {
		return nil, errors.New("invalid value for required argument 'NotificationSettings'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	if args.TimeZoneId == nil {
		return nil, errors.New("invalid value for required argument 'TimeZoneId'")
	}
	var resource Schedule
	err := ctx.RegisterResource("azure:devtest/schedule:Schedule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:devtest/schedule:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
	DailyRecurrence  *ScheduleDailyRecurrence  `pulumi:"dailyRecurrence"`
	HourlyRecurrence *ScheduleHourlyRecurrence `pulumi:"hourlyRecurrence"`
	// The name of the dev test lab. Changing this forces a new resource to be created.
	LabName *string `pulumi:"labName"`
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
	Name                 *string                       `pulumi:"name"`
	NotificationSettings *ScheduleNotificationSettings `pulumi:"notificationSettings"`
	// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
	TaskType *string `pulumi:"taskType"`
	// The time zone ID (e.g. Pacific Standard time).
	TimeZoneId       *string                   `pulumi:"timeZoneId"`
	WeeklyRecurrence *ScheduleWeeklyRecurrence `pulumi:"weeklyRecurrence"`
}

type ScheduleState struct {
	DailyRecurrence  ScheduleDailyRecurrencePtrInput
	HourlyRecurrence ScheduleHourlyRecurrencePtrInput
	// The name of the dev test lab. Changing this forces a new resource to be created.
	LabName pulumi.StringPtrInput
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
	Name                 pulumi.StringPtrInput
	NotificationSettings ScheduleNotificationSettingsPtrInput
	// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
	TaskType pulumi.StringPtrInput
	// The time zone ID (e.g. Pacific Standard time).
	TimeZoneId       pulumi.StringPtrInput
	WeeklyRecurrence ScheduleWeeklyRecurrencePtrInput
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	DailyRecurrence  *ScheduleDailyRecurrence  `pulumi:"dailyRecurrence"`
	HourlyRecurrence *ScheduleHourlyRecurrence `pulumi:"hourlyRecurrence"`
	// The name of the dev test lab. Changing this forces a new resource to be created.
	LabName string `pulumi:"labName"`
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
	Name                 *string                      `pulumi:"name"`
	NotificationSettings ScheduleNotificationSettings `pulumi:"notificationSettings"`
	// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
	TaskType string `pulumi:"taskType"`
	// The time zone ID (e.g. Pacific Standard time).
	TimeZoneId       string                    `pulumi:"timeZoneId"`
	WeeklyRecurrence *ScheduleWeeklyRecurrence `pulumi:"weeklyRecurrence"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	DailyRecurrence  ScheduleDailyRecurrencePtrInput
	HourlyRecurrence ScheduleHourlyRecurrencePtrInput
	// The name of the dev test lab. Changing this forces a new resource to be created.
	LabName pulumi.StringInput
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
	Name                 pulumi.StringPtrInput
	NotificationSettings ScheduleNotificationSettingsInput
	// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
	TaskType pulumi.StringInput
	// The time zone ID (e.g. Pacific Standard time).
	TimeZoneId       pulumi.StringInput
	WeeklyRecurrence ScheduleWeeklyRecurrencePtrInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}

type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput
}

func (Schedule) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

type ScheduleOutput struct {
	*pulumi.OutputState
}

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleOutput)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
