// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a maintenance assignment to virtual machine.
//
// ## Import
//
// Maintenance Assignment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:maintenance/assignmentVirtualMachine:AssignmentVirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/microsoft.compute/virtualMachines/vm1/providers/Microsoft.Maintenance/configurationAssignments/assign1
// ```
type AssignmentVirtualMachine struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId pulumi.StringOutput `pulumi:"maintenanceConfigurationId"`
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewAssignmentVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewAssignmentVirtualMachine(ctx *pulumi.Context,
	name string, args *AssignmentVirtualMachineArgs, opts ...pulumi.ResourceOption) (*AssignmentVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaintenanceConfigurationId == nil {
		return nil, errors.New("invalid value for required argument 'MaintenanceConfigurationId'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	var resource AssignmentVirtualMachine
	err := ctx.RegisterResource("azure:maintenance/assignmentVirtualMachine:AssignmentVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignmentVirtualMachine gets an existing AssignmentVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignmentVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentVirtualMachineState, opts ...pulumi.ResourceOption) (*AssignmentVirtualMachine, error) {
	var resource AssignmentVirtualMachine
	err := ctx.ReadResource("azure:maintenance/assignmentVirtualMachine:AssignmentVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssignmentVirtualMachine resources.
type assignmentVirtualMachineState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type AssignmentVirtualMachineState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringPtrInput
}

func (AssignmentVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentVirtualMachineState)(nil)).Elem()
}

type assignmentVirtualMachineArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId string `pulumi:"maintenanceConfigurationId"`
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a AssignmentVirtualMachine resource.
type AssignmentVirtualMachineArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId pulumi.StringInput
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringInput
}

func (AssignmentVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentVirtualMachineArgs)(nil)).Elem()
}

type AssignmentVirtualMachineInput interface {
	pulumi.Input

	ToAssignmentVirtualMachineOutput() AssignmentVirtualMachineOutput
	ToAssignmentVirtualMachineOutputWithContext(ctx context.Context) AssignmentVirtualMachineOutput
}

type AssignmentVirtualMachinePtrInput interface {
	pulumi.Input

	ToAssignmentVirtualMachinePtrOutput() AssignmentVirtualMachinePtrOutput
	ToAssignmentVirtualMachinePtrOutputWithContext(ctx context.Context) AssignmentVirtualMachinePtrOutput
}

func (AssignmentVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentVirtualMachine)(nil)).Elem()
}

func (i AssignmentVirtualMachine) ToAssignmentVirtualMachineOutput() AssignmentVirtualMachineOutput {
	return i.ToAssignmentVirtualMachineOutputWithContext(context.Background())
}

func (i AssignmentVirtualMachine) ToAssignmentVirtualMachineOutputWithContext(ctx context.Context) AssignmentVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentVirtualMachineOutput)
}

func (i AssignmentVirtualMachine) ToAssignmentVirtualMachinePtrOutput() AssignmentVirtualMachinePtrOutput {
	return i.ToAssignmentVirtualMachinePtrOutputWithContext(context.Background())
}

func (i AssignmentVirtualMachine) ToAssignmentVirtualMachinePtrOutputWithContext(ctx context.Context) AssignmentVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentVirtualMachinePtrOutput)
}

type AssignmentVirtualMachineOutput struct {
	*pulumi.OutputState
}

func (AssignmentVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentVirtualMachineOutput)(nil)).Elem()
}

func (o AssignmentVirtualMachineOutput) ToAssignmentVirtualMachineOutput() AssignmentVirtualMachineOutput {
	return o
}

func (o AssignmentVirtualMachineOutput) ToAssignmentVirtualMachineOutputWithContext(ctx context.Context) AssignmentVirtualMachineOutput {
	return o
}

type AssignmentVirtualMachinePtrOutput struct {
	*pulumi.OutputState
}

func (AssignmentVirtualMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentVirtualMachine)(nil)).Elem()
}

func (o AssignmentVirtualMachinePtrOutput) ToAssignmentVirtualMachinePtrOutput() AssignmentVirtualMachinePtrOutput {
	return o
}

func (o AssignmentVirtualMachinePtrOutput) ToAssignmentVirtualMachinePtrOutputWithContext(ctx context.Context) AssignmentVirtualMachinePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssignmentVirtualMachineOutput{})
	pulumi.RegisterOutputType(AssignmentVirtualMachinePtrOutput{})
}
