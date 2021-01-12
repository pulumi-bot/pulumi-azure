// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a [Lighthouse](https://docs.microsoft.com/en-us/azure/lighthouse) Definition.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/lighthouse"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "b24988ac-6180-42a0-ab88-20f7382dd24c"
// 		contributor, err := authorization.LookupRoleDefinition(ctx, &authorization.LookupRoleDefinitionArgs{
// 			RoleDefinitionId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lighthouse.NewDefinition(ctx, "example", &lighthouse.DefinitionArgs{
// 			Description:      pulumi.String("This is a lighthouse definition created IaC"),
// 			ManagingTenantId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Authorizations: lighthouse.DefinitionAuthorizationArray{
// 				&lighthouse.DefinitionAuthorizationArgs{
// 					PrincipalId:      pulumi.String("00000000-0000-0000-0000-000000000000"),
// 					RoleDefinitionId: pulumi.String(contributor.RoleDefinitionId),
// 				},
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
// Lighthouse Definitions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:lighthouse/definition:Definition example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationDefinitions/00000000-0000-0000-0000-000000000000
// ```
type Definition struct {
	pulumi.CustomResourceState

	// An authorization block as defined below.
	Authorizations DefinitionAuthorizationArrayOutput `pulumi:"authorizations"`
	// A description of the Lighthouse Definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
	LighthouseDefinitionId pulumi.StringOutput `pulumi:"lighthouseDefinitionId"`
	// The ID of the managing tenant.
	ManagingTenantId pulumi.StringOutput `pulumi:"managingTenantId"`
	// The name of the Lighthouse Definition.
	Name  pulumi.StringOutput `pulumi:"name"`
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOption) (*Definition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorizations == nil {
		return nil, errors.New("invalid value for required argument 'Authorizations'")
	}
	if args.ManagingTenantId == nil {
		return nil, errors.New("invalid value for required argument 'ManagingTenantId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource Definition
	err := ctx.RegisterResource("azure:lighthouse/definition:Definition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefinition gets an existing Definition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefinitionState, opts ...pulumi.ResourceOption) (*Definition, error) {
	var resource Definition
	err := ctx.ReadResource("azure:lighthouse/definition:Definition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Definition resources.
type definitionState struct {
	// An authorization block as defined below.
	Authorizations []DefinitionAuthorization `pulumi:"authorizations"`
	// A description of the Lighthouse Definition.
	Description *string `pulumi:"description"`
	// A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
	LighthouseDefinitionId *string `pulumi:"lighthouseDefinitionId"`
	// The ID of the managing tenant.
	ManagingTenantId *string `pulumi:"managingTenantId"`
	// The name of the Lighthouse Definition.
	Name  *string `pulumi:"name"`
	Scope *string `pulumi:"scope"`
}

type DefinitionState struct {
	// An authorization block as defined below.
	Authorizations DefinitionAuthorizationArrayInput
	// A description of the Lighthouse Definition.
	Description pulumi.StringPtrInput
	// A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
	LighthouseDefinitionId pulumi.StringPtrInput
	// The ID of the managing tenant.
	ManagingTenantId pulumi.StringPtrInput
	// The name of the Lighthouse Definition.
	Name  pulumi.StringPtrInput
	Scope pulumi.StringPtrInput
}

func (DefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionState)(nil)).Elem()
}

type definitionArgs struct {
	// An authorization block as defined below.
	Authorizations []DefinitionAuthorization `pulumi:"authorizations"`
	// A description of the Lighthouse Definition.
	Description *string `pulumi:"description"`
	// A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
	LighthouseDefinitionId *string `pulumi:"lighthouseDefinitionId"`
	// The ID of the managing tenant.
	ManagingTenantId string `pulumi:"managingTenantId"`
	// The name of the Lighthouse Definition.
	Name  *string `pulumi:"name"`
	Scope string  `pulumi:"scope"`
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// An authorization block as defined below.
	Authorizations DefinitionAuthorizationArrayInput
	// A description of the Lighthouse Definition.
	Description pulumi.StringPtrInput
	// A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
	LighthouseDefinitionId pulumi.StringPtrInput
	// The ID of the managing tenant.
	ManagingTenantId pulumi.StringInput
	// The name of the Lighthouse Definition.
	Name  pulumi.StringPtrInput
	Scope pulumi.StringInput
}

func (DefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionArgs)(nil)).Elem()
}

type DefinitionInput interface {
	pulumi.Input

	ToDefinitionOutput() DefinitionOutput
	ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput
}

func (*Definition) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil))
}

func (i *Definition) ToDefinitionOutput() DefinitionOutput {
	return i.ToDefinitionOutputWithContext(context.Background())
}

func (i *Definition) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionOutput)
}

func (i *Definition) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return i.ToDefinitionPtrOutputWithContext(context.Background())
}

func (i *Definition) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPtrOutput)
}

type DefinitionPtrInput interface {
	pulumi.Input

	ToDefinitionPtrOutput() DefinitionPtrOutput
	ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput
}

type definitionPtrType DefinitionArgs

func (*definitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Definition)(nil))
}

func (i *definitionPtrType) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return i.ToDefinitionPtrOutputWithContext(context.Background())
}

func (i *definitionPtrType) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionOutput).ToDefinitionPtrOutput()
}

// DefinitionArrayInput is an input type that accepts DefinitionArray and DefinitionArrayOutput values.
// You can construct a concrete instance of `DefinitionArrayInput` via:
//
//          DefinitionArray{ DefinitionArgs{...} }
type DefinitionArrayInput interface {
	pulumi.Input

	ToDefinitionArrayOutput() DefinitionArrayOutput
	ToDefinitionArrayOutputWithContext(context.Context) DefinitionArrayOutput
}

type DefinitionArray []DefinitionInput

func (DefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Definition)(nil))
}

func (i DefinitionArray) ToDefinitionArrayOutput() DefinitionArrayOutput {
	return i.ToDefinitionArrayOutputWithContext(context.Background())
}

func (i DefinitionArray) ToDefinitionArrayOutputWithContext(ctx context.Context) DefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionArrayOutput)
}

// DefinitionMapInput is an input type that accepts DefinitionMap and DefinitionMapOutput values.
// You can construct a concrete instance of `DefinitionMapInput` via:
//
//          DefinitionMap{ "key": DefinitionArgs{...} }
type DefinitionMapInput interface {
	pulumi.Input

	ToDefinitionMapOutput() DefinitionMapOutput
	ToDefinitionMapOutputWithContext(context.Context) DefinitionMapOutput
}

type DefinitionMap map[string]DefinitionInput

func (DefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Definition)(nil))
}

func (i DefinitionMap) ToDefinitionMapOutput() DefinitionMapOutput {
	return i.ToDefinitionMapOutputWithContext(context.Background())
}

func (i DefinitionMap) ToDefinitionMapOutputWithContext(ctx context.Context) DefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionMapOutput)
}

type DefinitionOutput struct {
	*pulumi.OutputState
}

func (DefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil))
}

func (o DefinitionOutput) ToDefinitionOutput() DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return o.ToDefinitionPtrOutputWithContext(context.Background())
}

func (o DefinitionOutput) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return o.ApplyT(func(v Definition) *Definition {
		return &v
	}).(DefinitionPtrOutput)
}

type DefinitionPtrOutput struct {
	*pulumi.OutputState
}

func (DefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Definition)(nil))
}

func (o DefinitionPtrOutput) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return o
}

func (o DefinitionPtrOutput) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return o
}

type DefinitionArrayOutput struct{ *pulumi.OutputState }

func (DefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Definition)(nil))
}

func (o DefinitionArrayOutput) ToDefinitionArrayOutput() DefinitionArrayOutput {
	return o
}

func (o DefinitionArrayOutput) ToDefinitionArrayOutputWithContext(ctx context.Context) DefinitionArrayOutput {
	return o
}

func (o DefinitionArrayOutput) Index(i pulumi.IntInput) DefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Definition {
		return vs[0].([]Definition)[vs[1].(int)]
	}).(DefinitionOutput)
}

type DefinitionMapOutput struct{ *pulumi.OutputState }

func (DefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Definition)(nil))
}

func (o DefinitionMapOutput) ToDefinitionMapOutput() DefinitionMapOutput {
	return o
}

func (o DefinitionMapOutput) ToDefinitionMapOutputWithContext(ctx context.Context) DefinitionMapOutput {
	return o
}

func (o DefinitionMapOutput) MapIndex(k pulumi.StringInput) DefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Definition {
		return vs[0].(map[string]Definition)[vs[1].(string)]
	}).(DefinitionOutput)
}

func init() {
	pulumi.RegisterOutputType(DefinitionOutput{})
	pulumi.RegisterOutputType(DefinitionPtrOutput{})
	pulumi.RegisterOutputType(DefinitionArrayOutput{})
	pulumi.RegisterOutputType(DefinitionMapOutput{})
}
