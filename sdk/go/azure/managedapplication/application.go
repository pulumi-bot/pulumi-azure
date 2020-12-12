// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managedapplication

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Managed Application.
//
// ## Import
//
// Managed Application can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:managedapplication/application:Application example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applications/app1
// ```
type Application struct {
	pulumi.CustomResourceState

	// The application definition ID to deploy.
	ApplicationDefinitionId pulumi.StringPtrOutput `pulumi:"applicationDefinitionId"`
	// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
	ManagedResourceGroupName pulumi.StringOutput `pulumi:"managedResourceGroupName"`
	// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name and value pairs that define the managed application outputs.
	Outputs pulumi.StringMapOutput `pulumi:"outputs"`
	// A mapping of name and value pairs to pass to the managed application as parameters.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// One `plan` block as defined below.
	Plan ApplicationPlanPtrOutput `pulumi:"plan"`
	// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ManagedResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedResourceGroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Application
	err := ctx.RegisterResource("azure:managedapplication/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure:managedapplication/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The application definition ID to deploy.
	ApplicationDefinitionId *string `pulumi:"applicationDefinitionId"`
	// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name and value pairs that define the managed application outputs.
	Outputs map[string]string `pulumi:"outputs"`
	// A mapping of name and value pairs to pass to the managed application as parameters.
	Parameters map[string]string `pulumi:"parameters"`
	// One `plan` block as defined below.
	Plan *ApplicationPlan `pulumi:"plan"`
	// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ApplicationState struct {
	// The application definition ID to deploy.
	ApplicationDefinitionId pulumi.StringPtrInput
	// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
	ManagedResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name and value pairs that define the managed application outputs.
	Outputs pulumi.StringMapInput
	// A mapping of name and value pairs to pass to the managed application as parameters.
	Parameters pulumi.StringMapInput
	// One `plan` block as defined below.
	Plan ApplicationPlanPtrInput
	// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The application definition ID to deploy.
	ApplicationDefinitionId *string `pulumi:"applicationDefinitionId"`
	// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
	Kind string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
	ManagedResourceGroupName string `pulumi:"managedResourceGroupName"`
	// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A mapping of name and value pairs to pass to the managed application as parameters.
	Parameters map[string]string `pulumi:"parameters"`
	// One `plan` block as defined below.
	Plan *ApplicationPlan `pulumi:"plan"`
	// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The application definition ID to deploy.
	ApplicationDefinitionId pulumi.StringPtrInput
	// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
	Kind pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
	ManagedResourceGroupName pulumi.StringInput
	// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A mapping of name and value pairs to pass to the managed application as parameters.
	Parameters pulumi.StringMapInput
	// One `plan` block as defined below.
	Plan ApplicationPlanPtrInput
	// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

func (i *Application) ToApplicationPtrOutput() ApplicationPtrOutput {
	return i.ToApplicationPtrOutputWithContext(context.Background())
}

func (i *Application) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPtrOutput)
}

type ApplicationPtrInput interface {
	pulumi.Input

	ToApplicationPtrOutput() ApplicationPtrOutput
	ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput
}

type ApplicationOutput struct {
	*pulumi.OutputState
}

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

type ApplicationPtrOutput struct {
	*pulumi.OutputState
}

func (ApplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil))
}

func (o ApplicationPtrOutput) ToApplicationPtrOutput() ApplicationPtrOutput {
	return o
}

func (o ApplicationPtrOutput) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationPtrOutput{})
}
