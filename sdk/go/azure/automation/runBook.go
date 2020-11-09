// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Automation Runbook.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/automation"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		exampleAccount, err := automation.NewAccount(ctx, "exampleAccount", &automation.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = automation.NewRunBook(ctx, "exampleRunBook", &automation.RunBookArgs{
// 			Location:              exampleResourceGroup.Location,
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			AutomationAccountName: exampleAccount.Name,
// 			LogVerbose:            pulumi.Bool(true),
// 			LogProgress:           pulumi.Bool(true),
// 			Description:           pulumi.String("This is an example runbook"),
// 			RunbookType:           pulumi.String("PowerShellWorkflow"),
// 			PublishContentLink: &automation.RunBookPublishContentLinkArgs{
// 				Uri: pulumi.String("https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RunBook struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The desired content of the runbook.
	Content pulumi.StringOutput `pulumi:"content"`
	// A description for this credential.
	Description  pulumi.StringPtrOutput        `pulumi:"description"`
	JobSchedules RunBookJobScheduleArrayOutput `pulumi:"jobSchedules"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Progress log option.
	LogProgress pulumi.BoolOutput `pulumi:"logProgress"`
	// Verbose log option.
	LogVerbose pulumi.BoolOutput `pulumi:"logVerbose"`
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The published runbook content link.
	PublishContentLink RunBookPublishContentLinkPtrOutput `pulumi:"publishContentLink"`
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType pulumi.StringOutput `pulumi:"runbookType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRunBook registers a new resource with the given unique name, arguments, and options.
func NewRunBook(ctx *pulumi.Context,
	name string, args *RunBookArgs, opts ...pulumi.ResourceOption) (*RunBook, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.LogProgress == nil {
		return nil, errors.New("missing required argument 'LogProgress'")
	}
	if args == nil || args.LogVerbose == nil {
		return nil, errors.New("missing required argument 'LogVerbose'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RunbookType == nil {
		return nil, errors.New("missing required argument 'RunbookType'")
	}
	if args == nil {
		args = &RunBookArgs{}
	}
	var resource RunBook
	err := ctx.RegisterResource("azure:automation/runBook:RunBook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRunBook gets an existing RunBook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRunBook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunBookState, opts ...pulumi.ResourceOption) (*RunBook, error) {
	var resource RunBook
	err := ctx.ReadResource("azure:automation/runBook:RunBook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RunBook resources.
type runBookState struct {
	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The desired content of the runbook.
	Content *string `pulumi:"content"`
	// A description for this credential.
	Description  *string              `pulumi:"description"`
	JobSchedules []RunBookJobSchedule `pulumi:"jobSchedules"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Progress log option.
	LogProgress *bool `pulumi:"logProgress"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The published runbook content link.
	PublishContentLink *RunBookPublishContentLink `pulumi:"publishContentLink"`
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType *string `pulumi:"runbookType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type RunBookState struct {
	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The desired content of the runbook.
	Content pulumi.StringPtrInput
	// A description for this credential.
	Description  pulumi.StringPtrInput
	JobSchedules RunBookJobScheduleArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Progress log option.
	LogProgress pulumi.BoolPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The published runbook content link.
	PublishContentLink RunBookPublishContentLinkPtrInput
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RunBookState) ElementType() reflect.Type {
	return reflect.TypeOf((*runBookState)(nil)).Elem()
}

type runBookArgs struct {
	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The desired content of the runbook.
	Content *string `pulumi:"content"`
	// A description for this credential.
	Description  *string              `pulumi:"description"`
	JobSchedules []RunBookJobSchedule `pulumi:"jobSchedules"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Progress log option.
	LogProgress bool `pulumi:"logProgress"`
	// Verbose log option.
	LogVerbose bool `pulumi:"logVerbose"`
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The published runbook content link.
	PublishContentLink *RunBookPublishContentLink `pulumi:"publishContentLink"`
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType string `pulumi:"runbookType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RunBook resource.
type RunBookArgs struct {
	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The desired content of the runbook.
	Content pulumi.StringPtrInput
	// A description for this credential.
	Description  pulumi.StringPtrInput
	JobSchedules RunBookJobScheduleArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Progress log option.
	LogProgress pulumi.BoolInput
	// Verbose log option.
	LogVerbose pulumi.BoolInput
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The published runbook content link.
	PublishContentLink RunBookPublishContentLinkPtrInput
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RunBookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runBookArgs)(nil)).Elem()
}

type RunBookInput interface {
	pulumi.Input

	ToRunBookOutput() RunBookOutput
	ToRunBookOutputWithContext(ctx context.Context) RunBookOutput
}

func (RunBook) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBook)(nil)).Elem()
}

func (i RunBook) ToRunBookOutput() RunBookOutput {
	return i.ToRunBookOutputWithContext(context.Background())
}

func (i RunBook) ToRunBookOutputWithContext(ctx context.Context) RunBookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookOutput)
}

type RunBookOutput struct {
	*pulumi.OutputState
}

func (RunBookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookOutput)(nil)).Elem()
}

func (o RunBookOutput) ToRunBookOutput() RunBookOutput {
	return o
}

func (o RunBookOutput) ToRunBookOutputWithContext(ctx context.Context) RunBookOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RunBookOutput{})
}
