// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a user assigned identity.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("eastus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewUserAssignedIdentity(ctx, "exampleUserAssignedIdentity", &authorization.UserAssignedIdentityArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
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
// User Assigned Identities can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:msi/userAssignedIdentity:UserAssignedIdentity exampleIdentity /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acceptanceTestResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity
// ```
//
// Deprecated: azure.msi.UserAssignedIdentity has been deprecated in favor of azure.authorization.UserAssignedIdentity
type UserAssignedIdentity struct {
	pulumi.CustomResourceState

	// Client ID associated with the user assigned identity.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The location/region where the user assigned identity is
	// created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service Principal ID associated with the user assigned identity.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewUserAssignedIdentity registers a new resource with the given unique name, arguments, and options.
func NewUserAssignedIdentity(ctx *pulumi.Context,
	name string, args *UserAssignedIdentityArgs, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource UserAssignedIdentity
	err := ctx.RegisterResource("azure:msi/userAssignedIdentity:UserAssignedIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAssignedIdentity gets an existing UserAssignedIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAssignedIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAssignedIdentityState, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	var resource UserAssignedIdentity
	err := ctx.ReadResource("azure:msi/userAssignedIdentity:UserAssignedIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAssignedIdentity resources.
type userAssignedIdentityState struct {
	// Client ID associated with the user assigned identity.
	ClientId *string `pulumi:"clientId"`
	// The location/region where the user assigned identity is
	// created.
	Location *string `pulumi:"location"`
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name *string `pulumi:"name"`
	// Service Principal ID associated with the user assigned identity.
	PrincipalId *string `pulumi:"principalId"`
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type UserAssignedIdentityState struct {
	// Client ID associated with the user assigned identity.
	ClientId pulumi.StringPtrInput
	// The location/region where the user assigned identity is
	// created.
	Location pulumi.StringPtrInput
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name pulumi.StringPtrInput
	// Service Principal ID associated with the user assigned identity.
	PrincipalId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (UserAssignedIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityState)(nil)).Elem()
}

type userAssignedIdentityArgs struct {
	// The location/region where the user assigned identity is
	// created.
	Location *string `pulumi:"location"`
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a UserAssignedIdentity resource.
type UserAssignedIdentityArgs struct {
	// The location/region where the user assigned identity is
	// created.
	Location pulumi.StringPtrInput
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityArgs)(nil)).Elem()
}

type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput
}

type UserAssignedIdentityPtrInput interface {
	pulumi.Input

	ToUserAssignedIdentityPtrOutput() UserAssignedIdentityPtrOutput
	ToUserAssignedIdentityPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPtrOutput
}

func (UserAssignedIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentity) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i UserAssignedIdentity) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}

func (i UserAssignedIdentity) ToUserAssignedIdentityPtrOutput() UserAssignedIdentityPtrOutput {
	return i.ToUserAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i UserAssignedIdentity) ToUserAssignedIdentityPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPtrOutput)
}

type UserAssignedIdentityOutput struct {
	*pulumi.OutputState
}

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityOutput)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

type UserAssignedIdentityPtrOutput struct {
	*pulumi.OutputState
}

func (UserAssignedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityPtrOutput) ToUserAssignedIdentityPtrOutput() UserAssignedIdentityPtrOutput {
	return o
}

func (o UserAssignedIdentityPtrOutput) ToUserAssignedIdentityPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPtrOutput{})
}
