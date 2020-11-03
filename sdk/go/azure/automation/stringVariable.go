// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a string variable in Azure Automation
type StringVariable struct {
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
	// The value of the Automation Variable as a `string`.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewStringVariable registers a new resource with the given unique name, arguments, and options.
func NewStringVariable(ctx *pulumi.Context,
	name string, args *StringVariableArgs, opts ...pulumi.ResourceOption) (*StringVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource StringVariable
	err := ctx.RegisterResource("azure:automation/stringVariable:StringVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStringVariable gets an existing StringVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStringVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StringVariableState, opts ...pulumi.ResourceOption) (*StringVariable, error) {
	var resource StringVariable
	err := ctx.ReadResource("azure:automation/stringVariable:StringVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StringVariable resources.
type stringVariableState struct {
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
	// The value of the Automation Variable as a `string`.
	Value *string `pulumi:"value"`
}

type StringVariableState struct {
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
	// The value of the Automation Variable as a `string`.
	Value pulumi.StringPtrInput
}

func (StringVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*stringVariableState)(nil)).Elem()
}

type stringVariableArgs struct {
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
	// The value of the Automation Variable as a `string`.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a StringVariable resource.
type StringVariableArgs struct {
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
	// The value of the Automation Variable as a `string`.
	Value pulumi.StringPtrInput
}

func (StringVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stringVariableArgs)(nil)).Elem()
}
