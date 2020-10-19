// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Lighthouse Definition.
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
// 			Description:      pulumi.String("This is a lighthouse definition created via Terraform"),
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
	if args == nil || args.Authorizations == nil {
		return nil, errors.New("missing required argument 'Authorizations'")
	}
	if args == nil || args.ManagingTenantId == nil {
		return nil, errors.New("missing required argument 'ManagingTenantId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &DefinitionArgs{}
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

func (Definition) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil)).Elem()
}

func (i Definition) ToDefinitionOutput() DefinitionOutput {
	return i.ToDefinitionOutputWithContext(context.Background())
}

func (i Definition) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionOutput)
}

type DefinitionOutput struct {
	*pulumi.OutputState
}

func (DefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionOutput)(nil)).Elem()
}

func (o DefinitionOutput) ToDefinitionOutput() DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DefinitionOutput{})
}
