// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type WorkspaceIdentity struct {
	// The (Client) ID of the Service Principal.
	PrincipalId *string `pulumi:"principalId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId *string `pulumi:"tenantId"`
	// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
	Type string `pulumi:"type"`
}

// WorkspaceIdentityInput is an input type that accepts WorkspaceIdentityArgs and WorkspaceIdentityOutput values.
// You can construct a concrete instance of `WorkspaceIdentityInput` via:
//
//          WorkspaceIdentityArgs{...}
type WorkspaceIdentityInput interface {
	pulumi.Input

	ToWorkspaceIdentityOutput() WorkspaceIdentityOutput
	ToWorkspaceIdentityOutputWithContext(context.Context) WorkspaceIdentityOutput
}

type WorkspaceIdentityArgs struct {
	// The (Client) ID of the Service Principal.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (WorkspaceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceIdentity)(nil)).Elem()
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityOutput() WorkspaceIdentityOutput {
	return i.ToWorkspaceIdentityOutputWithContext(context.Background())
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityOutputWithContext(ctx context.Context) WorkspaceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceIdentityOutput)
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return i.ToWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceIdentityOutput).ToWorkspaceIdentityPtrOutputWithContext(ctx)
}

// WorkspaceIdentityPtrInput is an input type that accepts WorkspaceIdentityArgs, WorkspaceIdentityPtr and WorkspaceIdentityPtrOutput values.
// You can construct a concrete instance of `WorkspaceIdentityPtrInput` via:
//
//          WorkspaceIdentityArgs{...}
//
//  or:
//
//          nil
type WorkspaceIdentityPtrInput interface {
	pulumi.Input

	ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput
	ToWorkspaceIdentityPtrOutputWithContext(context.Context) WorkspaceIdentityPtrOutput
}

type workspaceIdentityPtrType WorkspaceIdentityArgs

func WorkspaceIdentityPtr(v *WorkspaceIdentityArgs) WorkspaceIdentityPtrInput {
	return (*workspaceIdentityPtrType)(v)
}

func (*workspaceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceIdentity)(nil)).Elem()
}

func (i *workspaceIdentityPtrType) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return i.ToWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i *workspaceIdentityPtrType) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceIdentityPtrOutput)
}

type WorkspaceIdentityOutput struct{ *pulumi.OutputState }

func (WorkspaceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceIdentity)(nil)).Elem()
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityOutput() WorkspaceIdentityOutput {
	return o
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityOutputWithContext(ctx context.Context) WorkspaceIdentityOutput {
	return o
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return o.ToWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return o.ApplyT(func(v WorkspaceIdentity) *WorkspaceIdentity {
		return &v
	}).(WorkspaceIdentityPtrOutput)
}

// The (Client) ID of the Service Principal.
func (o WorkspaceIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The ID of the Tenant the Service Principal is assigned in.
func (o WorkspaceIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
func (o WorkspaceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type WorkspaceIdentityPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceIdentity)(nil)).Elem()
}

func (o WorkspaceIdentityPtrOutput) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return o
}

func (o WorkspaceIdentityPtrOutput) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return o
}

func (o WorkspaceIdentityPtrOutput) Elem() WorkspaceIdentityOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) WorkspaceIdentity { return *v }).(WorkspaceIdentityOutput)
}

// The (Client) ID of the Service Principal.
func (o WorkspaceIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The ID of the Tenant the Service Principal is assigned in.
func (o WorkspaceIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
func (o WorkspaceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type GetWorkspaceIdentity struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

// GetWorkspaceIdentityInput is an input type that accepts GetWorkspaceIdentityArgs and GetWorkspaceIdentityOutput values.
// You can construct a concrete instance of `GetWorkspaceIdentityInput` via:
//
//          GetWorkspaceIdentityArgs{...}
type GetWorkspaceIdentityInput interface {
	pulumi.Input

	ToGetWorkspaceIdentityOutput() GetWorkspaceIdentityOutput
	ToGetWorkspaceIdentityOutputWithContext(context.Context) GetWorkspaceIdentityOutput
}

type GetWorkspaceIdentityArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (GetWorkspaceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceIdentity)(nil)).Elem()
}

func (i GetWorkspaceIdentityArgs) ToGetWorkspaceIdentityOutput() GetWorkspaceIdentityOutput {
	return i.ToGetWorkspaceIdentityOutputWithContext(context.Background())
}

func (i GetWorkspaceIdentityArgs) ToGetWorkspaceIdentityOutputWithContext(ctx context.Context) GetWorkspaceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkspaceIdentityOutput)
}

// GetWorkspaceIdentityArrayInput is an input type that accepts GetWorkspaceIdentityArray and GetWorkspaceIdentityArrayOutput values.
// You can construct a concrete instance of `GetWorkspaceIdentityArrayInput` via:
//
//          GetWorkspaceIdentityArray{ GetWorkspaceIdentityArgs{...} }
type GetWorkspaceIdentityArrayInput interface {
	pulumi.Input

	ToGetWorkspaceIdentityArrayOutput() GetWorkspaceIdentityArrayOutput
	ToGetWorkspaceIdentityArrayOutputWithContext(context.Context) GetWorkspaceIdentityArrayOutput
}

type GetWorkspaceIdentityArray []GetWorkspaceIdentityInput

func (GetWorkspaceIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkspaceIdentity)(nil)).Elem()
}

func (i GetWorkspaceIdentityArray) ToGetWorkspaceIdentityArrayOutput() GetWorkspaceIdentityArrayOutput {
	return i.ToGetWorkspaceIdentityArrayOutputWithContext(context.Background())
}

func (i GetWorkspaceIdentityArray) ToGetWorkspaceIdentityArrayOutputWithContext(ctx context.Context) GetWorkspaceIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkspaceIdentityArrayOutput)
}

type GetWorkspaceIdentityOutput struct{ *pulumi.OutputState }

func (GetWorkspaceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceIdentity)(nil)).Elem()
}

func (o GetWorkspaceIdentityOutput) ToGetWorkspaceIdentityOutput() GetWorkspaceIdentityOutput {
	return o
}

func (o GetWorkspaceIdentityOutput) ToGetWorkspaceIdentityOutputWithContext(ctx context.Context) GetWorkspaceIdentityOutput {
	return o
}

func (o GetWorkspaceIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspaceIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o GetWorkspaceIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspaceIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o GetWorkspaceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspaceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type GetWorkspaceIdentityArrayOutput struct{ *pulumi.OutputState }

func (GetWorkspaceIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkspaceIdentity)(nil)).Elem()
}

func (o GetWorkspaceIdentityArrayOutput) ToGetWorkspaceIdentityArrayOutput() GetWorkspaceIdentityArrayOutput {
	return o
}

func (o GetWorkspaceIdentityArrayOutput) ToGetWorkspaceIdentityArrayOutputWithContext(ctx context.Context) GetWorkspaceIdentityArrayOutput {
	return o
}

func (o GetWorkspaceIdentityArrayOutput) Index(i pulumi.IntInput) GetWorkspaceIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetWorkspaceIdentity {
		return vs[0].([]GetWorkspaceIdentity)[vs[1].(int)]
	}).(GetWorkspaceIdentityOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceIdentityOutput{})
	pulumi.RegisterOutputType(WorkspaceIdentityPtrOutput{})
	pulumi.RegisterOutputType(GetWorkspaceIdentityOutput{})
	pulumi.RegisterOutputType(GetWorkspaceIdentityArrayOutput{})
}
