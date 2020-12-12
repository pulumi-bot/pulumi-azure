// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Recurrence Trigger within a Logic App Workflow
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/logicapps"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkflow, err := logicapps.NewWorkflow(ctx, "exampleWorkflow", &logicapps.WorkflowArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logicapps.NewTriggerRecurrence(ctx, "exampleTriggerRecurrence", &logicapps.TriggerRecurrenceArgs{
// 			LogicAppId: exampleWorkflow.ID(),
// 			Frequency:  pulumi.String("Day"),
// 			Interval:   pulumi.Int(1),
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
// Logic App Recurrence Triggers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/triggerRecurrence:TriggerRecurrence daily /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/daily
// ```
type TriggerRecurrence struct {
	pulumi.CustomResourceState

	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// Specifies the time zone for this trigger.  Supported time zone options are listed [here](https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values)
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewTriggerRecurrence registers a new resource with the given unique name, arguments, and options.
func NewTriggerRecurrence(ctx *pulumi.Context,
	name string, args *TriggerRecurrenceArgs, opts ...pulumi.ResourceOption) (*TriggerRecurrence, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Interval == nil {
		return nil, errors.New("invalid value for required argument 'Interval'")
	}
	if args.LogicAppId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppId'")
	}
	var resource TriggerRecurrence
	err := ctx.RegisterResource("azure:logicapps/triggerRecurrence:TriggerRecurrence", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerRecurrence gets an existing TriggerRecurrence resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerRecurrence(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerRecurrenceState, opts ...pulumi.ResourceOption) (*TriggerRecurrence, error) {
	var resource TriggerRecurrence
	err := ctx.ReadResource("azure:logicapps/triggerRecurrence:TriggerRecurrence", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerRecurrence resources.
type triggerRecurrenceState struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency *string `pulumi:"frequency"`
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval *int `pulumi:"interval"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime *string `pulumi:"startTime"`
	// Specifies the time zone for this trigger.  Supported time zone options are listed [here](https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values)
	TimeZone *string `pulumi:"timeZone"`
}

type TriggerRecurrenceState struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency pulumi.StringPtrInput
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval pulumi.IntPtrInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime pulumi.StringPtrInput
	// Specifies the time zone for this trigger.  Supported time zone options are listed [here](https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values)
	TimeZone pulumi.StringPtrInput
}

func (TriggerRecurrenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerRecurrenceState)(nil)).Elem()
}

type triggerRecurrenceArgs struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency string `pulumi:"frequency"`
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval int `pulumi:"interval"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime *string `pulumi:"startTime"`
	// Specifies the time zone for this trigger.  Supported time zone options are listed [here](https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values)
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a TriggerRecurrence resource.
type TriggerRecurrenceArgs struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency pulumi.StringInput
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval pulumi.IntInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime pulumi.StringPtrInput
	// Specifies the time zone for this trigger.  Supported time zone options are listed [here](https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values)
	TimeZone pulumi.StringPtrInput
}

func (TriggerRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerRecurrenceArgs)(nil)).Elem()
}

type TriggerRecurrenceInput interface {
	pulumi.Input

	ToTriggerRecurrenceOutput() TriggerRecurrenceOutput
	ToTriggerRecurrenceOutputWithContext(ctx context.Context) TriggerRecurrenceOutput
}

func (*TriggerRecurrence) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerRecurrence)(nil))
}

func (i *TriggerRecurrence) ToTriggerRecurrenceOutput() TriggerRecurrenceOutput {
	return i.ToTriggerRecurrenceOutputWithContext(context.Background())
}

func (i *TriggerRecurrence) ToTriggerRecurrenceOutputWithContext(ctx context.Context) TriggerRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerRecurrenceOutput)
}

func (i *TriggerRecurrence) ToTriggerRecurrencePtrOutput() TriggerRecurrencePtrOutput {
	return i.ToTriggerRecurrencePtrOutputWithContext(context.Background())
}

func (i *TriggerRecurrence) ToTriggerRecurrencePtrOutputWithContext(ctx context.Context) TriggerRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerRecurrencePtrOutput)
}

type TriggerRecurrencePtrInput interface {
	pulumi.Input

	ToTriggerRecurrencePtrOutput() TriggerRecurrencePtrOutput
	ToTriggerRecurrencePtrOutputWithContext(ctx context.Context) TriggerRecurrencePtrOutput
}

type TriggerRecurrenceOutput struct {
	*pulumi.OutputState
}

func (TriggerRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerRecurrence)(nil))
}

func (o TriggerRecurrenceOutput) ToTriggerRecurrenceOutput() TriggerRecurrenceOutput {
	return o
}

func (o TriggerRecurrenceOutput) ToTriggerRecurrenceOutputWithContext(ctx context.Context) TriggerRecurrenceOutput {
	return o
}

type TriggerRecurrencePtrOutput struct {
	*pulumi.OutputState
}

func (TriggerRecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerRecurrence)(nil))
}

func (o TriggerRecurrencePtrOutput) ToTriggerRecurrencePtrOutput() TriggerRecurrencePtrOutput {
	return o
}

func (o TriggerRecurrencePtrOutput) ToTriggerRecurrencePtrOutputWithContext(ctx context.Context) TriggerRecurrencePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TriggerRecurrenceOutput{})
	pulumi.RegisterOutputType(TriggerRecurrencePtrOutput{})
}
