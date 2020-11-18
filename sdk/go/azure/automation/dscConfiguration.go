// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Automation DSC Configuration.
type DscConfiguration struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringOutput `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location pulumi.StringOutput `pulumi:"location"`
	// Verbose log option.
	LogVerbose pulumi.BoolPtrOutput `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	State             pulumi.StringOutput `pulumi:"state"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDscConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscConfiguration(ctx *pulumi.Context,
	name string, args *DscConfigurationArgs, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ContentEmbedded == nil {
		return nil, errors.New("invalid value for required argument 'ContentEmbedded'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DscConfiguration
	err := ctx.RegisterResource("azure:automation/dscConfiguration:DscConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscConfiguration gets an existing DscConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscConfigurationState, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	var resource DscConfiguration
	err := ctx.ReadResource("azure:automation/dscConfiguration:DscConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscConfiguration resources.
type dscConfigurationState struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded *string `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description *string `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location *string `pulumi:"location"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	State             *string `pulumi:"state"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type DscConfigurationState struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringPtrInput
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrInput
	// Must be the same location as the Automation Account.
	Location pulumi.StringPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	State             pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DscConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationState)(nil)).Elem()
}

type dscConfigurationArgs struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded string `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description *string `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location *string `pulumi:"location"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DscConfiguration resource.
type DscConfigurationArgs struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringInput
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrInput
	// Must be the same location as the Automation Account.
	Location pulumi.StringPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DscConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationArgs)(nil)).Elem()
}

type DscConfigurationInput interface {
	pulumi.Input

	ToDscConfigurationOutput() DscConfigurationOutput
	ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput
}

func (DscConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfiguration)(nil)).Elem()
}

func (i DscConfiguration) ToDscConfigurationOutput() DscConfigurationOutput {
	return i.ToDscConfigurationOutputWithContext(context.Background())
}

func (i DscConfiguration) ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationOutput)
}

type DscConfigurationOutput struct {
	*pulumi.OutputState
}

func (DscConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationOutput)(nil)).Elem()
}

func (o DscConfigurationOutput) ToDscConfigurationOutput() DscConfigurationOutput {
	return o
}

func (o DscConfigurationOutput) ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DscConfigurationOutput{})
}
