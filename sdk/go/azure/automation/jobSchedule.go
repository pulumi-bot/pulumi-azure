// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Links an Automation Runbook and Schedule.
//
// ## Example Usage
//
// This is an example of just the Job Schedule.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/automation"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := automation.NewJobSchedule(ctx, "example", &automation.JobScheduleArgs{
// 			AutomationAccountName: pulumi.String("tf-automation-account"),
// 			Parameters: pulumi.StringMap{
// 				"resourcegroup": pulumi.String("tf-rgr-vm"),
// 				"vmname":        pulumi.String("TF-VM-01"),
// 			},
// 			ResourceGroupName: pulumi.String("tf-rgr-automation"),
// 			RunbookName:       pulumi.String("Get-VirtualMachine"),
// 			ScheduleName:      pulumi.String("hour"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type JobSchedule struct {
	pulumi.CustomResourceState

	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId pulumi.StringOutput `pulumi:"jobScheduleId"`
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn pulumi.StringPtrOutput `pulumi:"runOn"`
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  pulumi.StringOutput `pulumi:"runbookName"`
	ScheduleName pulumi.StringOutput `pulumi:"scheduleName"`
}

// NewJobSchedule registers a new resource with the given unique name, arguments, and options.
func NewJobSchedule(ctx *pulumi.Context,
	name string, args *JobScheduleArgs, opts ...pulumi.ResourceOption) (*JobSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RunbookName == nil {
		return nil, errors.New("invalid value for required argument 'RunbookName'")
	}
	if args.ScheduleName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleName'")
	}
	var resource JobSchedule
	err := ctx.RegisterResource("azure:automation/jobSchedule:JobSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobSchedule gets an existing JobSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobScheduleState, opts ...pulumi.ResourceOption) (*JobSchedule, error) {
	var resource JobSchedule
	err := ctx.ReadResource("azure:automation/jobSchedule:JobSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobSchedule resources.
type jobScheduleState struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId *string `pulumi:"jobScheduleId"`
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn *string `pulumi:"runOn"`
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  *string `pulumi:"runbookName"`
	ScheduleName *string `pulumi:"scheduleName"`
}

type JobScheduleState struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId pulumi.StringPtrInput
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn pulumi.StringPtrInput
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  pulumi.StringPtrInput
	ScheduleName pulumi.StringPtrInput
}

func (JobScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobScheduleState)(nil)).Elem()
}

type jobScheduleArgs struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId *string `pulumi:"jobScheduleId"`
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn *string `pulumi:"runOn"`
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  string `pulumi:"runbookName"`
	ScheduleName string `pulumi:"scheduleName"`
}

// The set of arguments for constructing a JobSchedule resource.
type JobScheduleArgs struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId pulumi.StringPtrInput
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn pulumi.StringPtrInput
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  pulumi.StringInput
	ScheduleName pulumi.StringInput
}

func (JobScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobScheduleArgs)(nil)).Elem()
}
