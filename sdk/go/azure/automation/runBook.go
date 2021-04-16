// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//
// ## Import
//
// Automation Runbooks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/runBook:RunBook Get-AzureVMTutorial /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.LogProgress == nil {
		return nil, errors.New("invalid value for required argument 'LogProgress'")
	}
	if args.LogVerbose == nil {
		return nil, errors.New("invalid value for required argument 'LogVerbose'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RunbookType == nil {
		return nil, errors.New("invalid value for required argument 'RunbookType'")
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

func (*RunBook) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBook)(nil))
}

func (i *RunBook) ToRunBookOutput() RunBookOutput {
	return i.ToRunBookOutputWithContext(context.Background())
}

func (i *RunBook) ToRunBookOutputWithContext(ctx context.Context) RunBookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookOutput)
}

func (i *RunBook) ToRunBookPtrOutput() RunBookPtrOutput {
	return i.ToRunBookPtrOutputWithContext(context.Background())
}

func (i *RunBook) ToRunBookPtrOutputWithContext(ctx context.Context) RunBookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPtrOutput)
}

type RunBookPtrInput interface {
	pulumi.Input

	ToRunBookPtrOutput() RunBookPtrOutput
	ToRunBookPtrOutputWithContext(ctx context.Context) RunBookPtrOutput
}

type runBookPtrType RunBookArgs

func (*runBookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBook)(nil))
}

func (i *runBookPtrType) ToRunBookPtrOutput() RunBookPtrOutput {
	return i.ToRunBookPtrOutputWithContext(context.Background())
}

func (i *runBookPtrType) ToRunBookPtrOutputWithContext(ctx context.Context) RunBookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPtrOutput)
}

// RunBookArrayInput is an input type that accepts RunBookArray and RunBookArrayOutput values.
// You can construct a concrete instance of `RunBookArrayInput` via:
//
//          RunBookArray{ RunBookArgs{...} }
type RunBookArrayInput interface {
	pulumi.Input

	ToRunBookArrayOutput() RunBookArrayOutput
	ToRunBookArrayOutputWithContext(context.Context) RunBookArrayOutput
}

type RunBookArray []RunBookInput

func (RunBookArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RunBook)(nil))
}

func (i RunBookArray) ToRunBookArrayOutput() RunBookArrayOutput {
	return i.ToRunBookArrayOutputWithContext(context.Background())
}

func (i RunBookArray) ToRunBookArrayOutputWithContext(ctx context.Context) RunBookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookArrayOutput)
}

// RunBookMapInput is an input type that accepts RunBookMap and RunBookMapOutput values.
// You can construct a concrete instance of `RunBookMapInput` via:
//
//          RunBookMap{ "key": RunBookArgs{...} }
type RunBookMapInput interface {
	pulumi.Input

	ToRunBookMapOutput() RunBookMapOutput
	ToRunBookMapOutputWithContext(context.Context) RunBookMapOutput
}

type RunBookMap map[string]RunBookInput

func (RunBookMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RunBook)(nil))
}

func (i RunBookMap) ToRunBookMapOutput() RunBookMapOutput {
	return i.ToRunBookMapOutputWithContext(context.Background())
}

func (i RunBookMap) ToRunBookMapOutputWithContext(ctx context.Context) RunBookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookMapOutput)
}

type RunBookOutput struct {
	*pulumi.OutputState
}

func (RunBookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBook)(nil))
}

func (o RunBookOutput) ToRunBookOutput() RunBookOutput {
	return o
}

func (o RunBookOutput) ToRunBookOutputWithContext(ctx context.Context) RunBookOutput {
	return o
}

func (o RunBookOutput) ToRunBookPtrOutput() RunBookPtrOutput {
	return o.ToRunBookPtrOutputWithContext(context.Background())
}

func (o RunBookOutput) ToRunBookPtrOutputWithContext(ctx context.Context) RunBookPtrOutput {
	return o.ApplyT(func(v RunBook) *RunBook {
		return &v
	}).(RunBookPtrOutput)
}

type RunBookPtrOutput struct {
	*pulumi.OutputState
}

func (RunBookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBook)(nil))
}

func (o RunBookPtrOutput) ToRunBookPtrOutput() RunBookPtrOutput {
	return o
}

func (o RunBookPtrOutput) ToRunBookPtrOutputWithContext(ctx context.Context) RunBookPtrOutput {
	return o
}

type RunBookArrayOutput struct{ *pulumi.OutputState }

func (RunBookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunBook)(nil))
}

func (o RunBookArrayOutput) ToRunBookArrayOutput() RunBookArrayOutput {
	return o
}

func (o RunBookArrayOutput) ToRunBookArrayOutputWithContext(ctx context.Context) RunBookArrayOutput {
	return o
}

func (o RunBookArrayOutput) Index(i pulumi.IntInput) RunBookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RunBook {
		return vs[0].([]RunBook)[vs[1].(int)]
	}).(RunBookOutput)
}

type RunBookMapOutput struct{ *pulumi.OutputState }

func (RunBookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunBook)(nil))
}

func (o RunBookMapOutput) ToRunBookMapOutput() RunBookMapOutput {
	return o
}

func (o RunBookMapOutput) ToRunBookMapOutputWithContext(ctx context.Context) RunBookMapOutput {
	return o
}

func (o RunBookMapOutput) MapIndex(k pulumi.StringInput) RunBookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunBook {
		return vs[0].(map[string]RunBook)[vs[1].(string)]
	}).(RunBookOutput)
}

func init() {
	pulumi.RegisterOutputType(RunBookOutput{})
	pulumi.RegisterOutputType(RunBookPtrOutput{})
	pulumi.RegisterOutputType(RunBookArrayOutput{})
	pulumi.RegisterOutputType(RunBookMapOutput{})
}
