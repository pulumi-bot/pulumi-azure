// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a DateTime variable in Azure Automation
type DateTimeVariable struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewDateTimeVariable registers a new resource with the given unique name, arguments, and options.
func NewDateTimeVariable(ctx *pulumi.Context,
	name string, args *DateTimeVariableArgs, opts ...pulumi.ResourceOption) (*DateTimeVariable, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DateTimeVariableArgs{}
	}
	var resource DateTimeVariable
	err := ctx.RegisterResource("azure:automation/dateTimeVariable:DateTimeVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDateTimeVariable gets an existing DateTimeVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDateTimeVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DateTimeVariableState, opts ...pulumi.ResourceOption) (*DateTimeVariable, error) {
	var resource DateTimeVariable
	err := ctx.ReadResource("azure:automation/dateTimeVariable:DateTimeVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DateTimeVariable resources.
type dateTimeVariableState struct {
	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description *string `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted *bool `pulumi:"encrypted"`
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Value *string `pulumi:"value"`
}

type DateTimeVariableState struct {
	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The description of the Automation Variable.
	Description pulumi.StringPtrInput
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted pulumi.BoolPtrInput
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Value pulumi.StringPtrInput
}

func (DateTimeVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*dateTimeVariableState)(nil)).Elem()
}

type dateTimeVariableArgs struct {
	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description *string `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted *bool `pulumi:"encrypted"`
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a DateTimeVariable resource.
type DateTimeVariableArgs struct {
	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The description of the Automation Variable.
	Description pulumi.StringPtrInput
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted pulumi.BoolPtrInput
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Value pulumi.StringPtrInput
}

func (DateTimeVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dateTimeVariableArgs)(nil)).Elem()
}

type DateTimeVariableInput interface {
	pulumi.Input

	ToDateTimeVariableOutput() DateTimeVariableOutput
	ToDateTimeVariableOutputWithContext(ctx context.Context) DateTimeVariableOutput
}

func (DateTimeVariable) ElementType() reflect.Type {
	return reflect.TypeOf((*DateTimeVariable)(nil)).Elem()
}

func (i DateTimeVariable) ToDateTimeVariableOutput() DateTimeVariableOutput {
	return i.ToDateTimeVariableOutputWithContext(context.Background())
}

func (i DateTimeVariable) ToDateTimeVariableOutputWithContext(ctx context.Context) DateTimeVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateTimeVariableOutput)
}

type DateTimeVariableOutput struct {
	*pulumi.OutputState
}

func (DateTimeVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateTimeVariableOutput)(nil)).Elem()
}

func (o DateTimeVariableOutput) ToDateTimeVariableOutput() DateTimeVariableOutput {
	return o
}

func (o DateTimeVariableOutput) ToDateTimeVariableOutputWithContext(ctx context.Context) DateTimeVariableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DateTimeVariableOutput{})
}
