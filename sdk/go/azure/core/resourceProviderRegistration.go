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

func (i *ResourceProviderRegistration) ToResourceProviderRegistrationPtrOutput() ResourceProviderRegistrationPtrOutput {
	return i.ToResourceProviderRegistrationPtrOutputWithContext(context.Background())
}

func (i *ResourceProviderRegistration) ToResourceProviderRegistrationPtrOutputWithContext(ctx context.Context) ResourceProviderRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceProviderRegistrationPtrOutput)
}

type ResourceProviderRegistrationPtrInput interface {
	pulumi.Input

	ToResourceProviderRegistrationPtrOutput() ResourceProviderRegistrationPtrOutput
	ToResourceProviderRegistrationPtrOutputWithContext(ctx context.Context) ResourceProviderRegistrationPtrOutput
}

type resourceProviderRegistrationPtrType ResourceProviderRegistrationArgs

func (*resourceProviderRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceProviderRegistration)(nil))
}

func (i *resourceProviderRegistrationPtrType) ToResourceProviderRegistrationPtrOutput() ResourceProviderRegistrationPtrOutput {
	return i.ToResourceProviderRegistrationPtrOutputWithContext(context.Background())
}

func (i *resourceProviderRegistrationPtrType) ToResourceProviderRegistrationPtrOutputWithContext(ctx context.Context) ResourceProviderRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceProviderRegistrationPtrOutput)
}

// ResourceProviderRegistrationArrayInput is an input type that accepts ResourceProviderRegistrationArray and ResourceProviderRegistrationArrayOutput values.
// You can construct a concrete instance of `ResourceProviderRegistrationArrayInput` via:
//
//          ResourceProviderRegistrationArray{ ResourceProviderRegistrationArgs{...} }
type ResourceProviderRegistrationArrayInput interface {
	pulumi.Input

	ToResourceProviderRegistrationArrayOutput() ResourceProviderRegistrationArrayOutput
	ToResourceProviderRegistrationArrayOutputWithContext(context.Context) ResourceProviderRegistrationArrayOutput
}

type ResourceProviderRegistrationArray []ResourceProviderRegistrationInput

func (ResourceProviderRegistrationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResourceProviderRegistration)(nil))
}

func (i ResourceProviderRegistrationArray) ToResourceProviderRegistrationArrayOutput() ResourceProviderRegistrationArrayOutput {
	return i.ToResourceProviderRegistrationArrayOutputWithContext(context.Background())
}

func (i ResourceProviderRegistrationArray) ToResourceProviderRegistrationArrayOutputWithContext(ctx context.Context) ResourceProviderRegistrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceProviderRegistrationArrayOutput)
}

// ResourceProviderRegistrationMapInput is an input type that accepts ResourceProviderRegistrationMap and ResourceProviderRegistrationMapOutput values.
// You can construct a concrete instance of `ResourceProviderRegistrationMapInput` via:
//
//          ResourceProviderRegistrationMap{ "key": ResourceProviderRegistrationArgs{...} }
type ResourceProviderRegistrationMapInput interface {
	pulumi.Input

	ToResourceProviderRegistrationMapOutput() ResourceProviderRegistrationMapOutput
	ToResourceProviderRegistrationMapOutputWithContext(context.Context) ResourceProviderRegistrationMapOutput
}

type ResourceProviderRegistrationMap map[string]ResourceProviderRegistrationInput

func (ResourceProviderRegistrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResourceProviderRegistration)(nil))
}

func (i ResourceProviderRegistrationMap) ToResourceProviderRegistrationMapOutput() ResourceProviderRegistrationMapOutput {
	return i.ToResourceProviderRegistrationMapOutputWithContext(context.Background())
}

func (i ResourceProviderRegistrationMap) ToResourceProviderRegistrationMapOutputWithContext(ctx context.Context) ResourceProviderRegistrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceProviderRegistrationMapOutput)
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

func (o ResourceProviderRegistrationOutput) ToResourceProviderRegistrationPtrOutput() ResourceProviderRegistrationPtrOutput {
	return o.ToResourceProviderRegistrationPtrOutputWithContext(context.Background())
}

func (o ResourceProviderRegistrationOutput) ToResourceProviderRegistrationPtrOutputWithContext(ctx context.Context) ResourceProviderRegistrationPtrOutput {
	return o.ApplyT(func(v ResourceProviderRegistration) *ResourceProviderRegistration {
		return &v
	}).(ResourceProviderRegistrationPtrOutput)
}

type ResourceProviderRegistrationPtrOutput struct {
	*pulumi.OutputState
}

func (ResourceProviderRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceProviderRegistration)(nil))
}

func (o ResourceProviderRegistrationPtrOutput) ToResourceProviderRegistrationPtrOutput() ResourceProviderRegistrationPtrOutput {
	return o
}

func (o ResourceProviderRegistrationPtrOutput) ToResourceProviderRegistrationPtrOutputWithContext(ctx context.Context) ResourceProviderRegistrationPtrOutput {
	return o
}

type ResourceProviderRegistrationArrayOutput struct{ *pulumi.OutputState }

func (ResourceProviderRegistrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceProviderRegistration)(nil))
}

func (o ResourceProviderRegistrationArrayOutput) ToResourceProviderRegistrationArrayOutput() ResourceProviderRegistrationArrayOutput {
	return o
}

func (o ResourceProviderRegistrationArrayOutput) ToResourceProviderRegistrationArrayOutputWithContext(ctx context.Context) ResourceProviderRegistrationArrayOutput {
	return o
}

func (o ResourceProviderRegistrationArrayOutput) Index(i pulumi.IntInput) ResourceProviderRegistrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceProviderRegistration {
		return vs[0].([]ResourceProviderRegistration)[vs[1].(int)]
	}).(ResourceProviderRegistrationOutput)
}

type ResourceProviderRegistrationMapOutput struct{ *pulumi.OutputState }

func (ResourceProviderRegistrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceProviderRegistration)(nil))
}

func (o ResourceProviderRegistrationMapOutput) ToResourceProviderRegistrationMapOutput() ResourceProviderRegistrationMapOutput {
	return o
}

func (o ResourceProviderRegistrationMapOutput) ToResourceProviderRegistrationMapOutputWithContext(ctx context.Context) ResourceProviderRegistrationMapOutput {
	return o
}

func (o ResourceProviderRegistrationMapOutput) MapIndex(k pulumi.StringInput) ResourceProviderRegistrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceProviderRegistration {
		return vs[0].(map[string]ResourceProviderRegistration)[vs[1].(string)]
	}).(ResourceProviderRegistrationOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceProviderRegistrationOutput{})
	pulumi.RegisterOutputType(ResourceProviderRegistrationPtrOutput{})
	pulumi.RegisterOutputType(ResourceProviderRegistrationArrayOutput{})
	pulumi.RegisterOutputType(ResourceProviderRegistrationMapOutput{})
}
