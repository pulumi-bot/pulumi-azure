// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Resource Provider Registrations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/resourceProviderRegistration:ResourceProviderRegistration example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights
// ```
type ResourceProviderRegistration struct {
	pulumi.CustomResourceState

	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewResourceProviderRegistration registers a new resource with the given unique name, arguments, and options.
func NewResourceProviderRegistration(ctx *pulumi.Context,
	name string, args *ResourceProviderRegistrationArgs, opts ...pulumi.ResourceOption) (*ResourceProviderRegistration, error) {
	if args == nil {
		args = &ResourceProviderRegistrationArgs{}
	}

	var resource ResourceProviderRegistration
	err := ctx.RegisterResource("azure:core/resourceProviderRegistration:ResourceProviderRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceProviderRegistration gets an existing ResourceProviderRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceProviderRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceProviderRegistrationState, opts ...pulumi.ResourceOption) (*ResourceProviderRegistration, error) {
	var resource ResourceProviderRegistration
	err := ctx.ReadResource("azure:core/resourceProviderRegistration:ResourceProviderRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceProviderRegistration resources.
type resourceProviderRegistrationState struct {
	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

type ResourceProviderRegistrationState struct {
	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (ResourceProviderRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceProviderRegistrationState)(nil)).Elem()
}

type resourceProviderRegistrationArgs struct {
	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ResourceProviderRegistration resource.
type ResourceProviderRegistrationArgs struct {
	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (ResourceProviderRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceProviderRegistrationArgs)(nil)).Elem()
}

type ResourceProviderRegistrationInput interface {
	pulumi.Input

	ToResourceProviderRegistrationOutput() ResourceProviderRegistrationOutput
	ToResourceProviderRegistrationOutputWithContext(ctx context.Context) ResourceProviderRegistrationOutput
}

func (*ResourceProviderRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProviderRegistration)(nil))
}

func (i *ResourceProviderRegistration) ToResourceProviderRegistrationOutput() ResourceProviderRegistrationOutput {
	return i.ToResourceProviderRegistrationOutputWithContext(context.Background())
}

func (i *ResourceProviderRegistration) ToResourceProviderRegistrationOutputWithContext(ctx context.Context) ResourceProviderRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceProviderRegistrationOutput)
}

type ResourceProviderRegistrationOutput struct {
	*pulumi.OutputState
}

func (ResourceProviderRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProviderRegistration)(nil))
}

func (o ResourceProviderRegistrationOutput) ToResourceProviderRegistrationOutput() ResourceProviderRegistrationOutput {
	return o
}

func (o ResourceProviderRegistrationOutput) ToResourceProviderRegistrationOutputWithContext(ctx context.Context) ResourceProviderRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceProviderRegistrationOutput{})
}
