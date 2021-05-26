// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Role Definition.
func LookupRoleDefinition(ctx *pulumi.Context, args *LookupRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRoleDefinitionResult, error) {
	var rv LookupRoleDefinitionResult
	err := ctx.Invoke("azure:authorization/getRoleDefinition:getRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoleDefinition.
type LookupRoleDefinitionArgs struct {
	// Specifies the Name of either a built-in or custom Role Definition.
	Name *string `pulumi:"name"`
	// Specifies the ID of the Role Definition as a UUID/GUID.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// Specifies the Scope at which the Custom Role Definition exists.
	Scope *string `pulumi:"scope"`
}

// A collection of values returned by getRoleDefinition.
type LookupRoleDefinitionResult struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// the Description of the built-in Role.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// a `permissions` block as documented below.
	Permissions      []GetRoleDefinitionPermission `pulumi:"permissions"`
	RoleDefinitionId string                        `pulumi:"roleDefinitionId"`
	Scope            *string                       `pulumi:"scope"`
	// the Type of the Role.
	Type string `pulumi:"type"`
}

func LookupRoleDefinitionApply(ctx *pulumi.Context, args LookupRoleDefinitionApplyInput, opts ...pulumi.InvokeOption) LookupRoleDefinitionResultOutput {
	return args.ToLookupRoleDefinitionApplyOutput().ApplyT(func(v LookupRoleDefinitionArgs) (LookupRoleDefinitionResult, error) {
		r, err := LookupRoleDefinition(ctx, &v, opts...)
		return *r, err

	}).(LookupRoleDefinitionResultOutput)
}

// LookupRoleDefinitionApplyInput is an input type that accepts LookupRoleDefinitionApplyArgs and LookupRoleDefinitionApplyOutput values.
// You can construct a concrete instance of `LookupRoleDefinitionApplyInput` via:
//
//          LookupRoleDefinitionApplyArgs{...}
type LookupRoleDefinitionApplyInput interface {
	pulumi.Input

	ToLookupRoleDefinitionApplyOutput() LookupRoleDefinitionApplyOutput
	ToLookupRoleDefinitionApplyOutputWithContext(context.Context) LookupRoleDefinitionApplyOutput
}

// A collection of arguments for invoking getRoleDefinition.
type LookupRoleDefinitionApplyArgs struct {
	// Specifies the Name of either a built-in or custom Role Definition.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the ID of the Role Definition as a UUID/GUID.
	RoleDefinitionId pulumi.StringPtrInput `pulumi:"roleDefinitionId"`
	// Specifies the Scope at which the Custom Role Definition exists.
	Scope pulumi.StringPtrInput `pulumi:"scope"`
}

func (LookupRoleDefinitionApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionArgs)(nil)).Elem()
}

func (i LookupRoleDefinitionApplyArgs) ToLookupRoleDefinitionApplyOutput() LookupRoleDefinitionApplyOutput {
	return i.ToLookupRoleDefinitionApplyOutputWithContext(context.Background())
}

func (i LookupRoleDefinitionApplyArgs) ToLookupRoleDefinitionApplyOutputWithContext(ctx context.Context) LookupRoleDefinitionApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupRoleDefinitionApplyOutput)
}

// A collection of arguments for invoking getRoleDefinition.
type LookupRoleDefinitionApplyOutput struct{ *pulumi.OutputState }

func (LookupRoleDefinitionApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionArgs)(nil)).Elem()
}

func (o LookupRoleDefinitionApplyOutput) ToLookupRoleDefinitionApplyOutput() LookupRoleDefinitionApplyOutput {
	return o
}

func (o LookupRoleDefinitionApplyOutput) ToLookupRoleDefinitionApplyOutputWithContext(ctx context.Context) LookupRoleDefinitionApplyOutput {
	return o
}

// Specifies the Name of either a built-in or custom Role Definition.
func (o LookupRoleDefinitionApplyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionArgs) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Specifies the ID of the Role Definition as a UUID/GUID.
func (o LookupRoleDefinitionApplyOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionArgs) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// Specifies the Scope at which the Custom Role Definition exists.
func (o LookupRoleDefinitionApplyOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionArgs) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getRoleDefinition.
type LookupRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionResult)(nil)).Elem()
}

func (o LookupRoleDefinitionResultOutput) ToLookupRoleDefinitionResultOutput() LookupRoleDefinitionResultOutput {
	return o
}

func (o LookupRoleDefinitionResultOutput) ToLookupRoleDefinitionResultOutputWithContext(ctx context.Context) LookupRoleDefinitionResultOutput {
	return o
}

// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
func (o LookupRoleDefinitionResultOutput) AssignableScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) []string { return v.AssignableScopes }).(pulumi.StringArrayOutput)
}

// the Description of the built-in Role.
func (o LookupRoleDefinitionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// a `permissions` block as documented below.
func (o LookupRoleDefinitionResultOutput) Permissions() GetRoleDefinitionPermissionArrayOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) []GetRoleDefinitionPermission { return v.Permissions }).(GetRoleDefinitionPermissionArrayOutput)
}

func (o LookupRoleDefinitionResultOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o LookupRoleDefinitionResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// the Type of the Role.
func (o LookupRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleDefinitionApplyOutput{})
	pulumi.RegisterOutputType(LookupRoleDefinitionResultOutput{})
}
