// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package purview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountIdentity struct {
	// The ID of the Principal (Client) in Azure Active Directory.
	PrincipalId *string `pulumi:"principalId"`
	// The ID of the Azure Active Directory Tenant.
	TenantId *string `pulumi:"tenantId"`
	// The type of Managed Identity assigned to this Purview Account.
	Type *string `pulumi:"type"`
}

// AccountIdentityInput is an input type that accepts AccountIdentityArgs and AccountIdentityOutput values.
// You can construct a concrete instance of `AccountIdentityInput` via:
//
//          AccountIdentityArgs{...}
type AccountIdentityInput interface {
	pulumi.Input

	ToAccountIdentityOutput() AccountIdentityOutput
	ToAccountIdentityOutputWithContext(context.Context) AccountIdentityOutput
}

type AccountIdentityArgs struct {
	// The ID of the Principal (Client) in Azure Active Directory.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The ID of the Azure Active Directory Tenant.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// The type of Managed Identity assigned to this Purview Account.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (AccountIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentity)(nil)).Elem()
}

func (i AccountIdentityArgs) ToAccountIdentityOutput() AccountIdentityOutput {
	return i.ToAccountIdentityOutputWithContext(context.Background())
}

func (i AccountIdentityArgs) ToAccountIdentityOutputWithContext(ctx context.Context) AccountIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityOutput)
}

// AccountIdentityArrayInput is an input type that accepts AccountIdentityArray and AccountIdentityArrayOutput values.
// You can construct a concrete instance of `AccountIdentityArrayInput` via:
//
//          AccountIdentityArray{ AccountIdentityArgs{...} }
type AccountIdentityArrayInput interface {
	pulumi.Input

	ToAccountIdentityArrayOutput() AccountIdentityArrayOutput
	ToAccountIdentityArrayOutputWithContext(context.Context) AccountIdentityArrayOutput
}

type AccountIdentityArray []AccountIdentityInput

func (AccountIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountIdentity)(nil)).Elem()
}

func (i AccountIdentityArray) ToAccountIdentityArrayOutput() AccountIdentityArrayOutput {
	return i.ToAccountIdentityArrayOutputWithContext(context.Background())
}

func (i AccountIdentityArray) ToAccountIdentityArrayOutputWithContext(ctx context.Context) AccountIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityArrayOutput)
}

type AccountIdentityOutput struct{ *pulumi.OutputState }

func (AccountIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentity)(nil)).Elem()
}

func (o AccountIdentityOutput) ToAccountIdentityOutput() AccountIdentityOutput {
	return o
}

func (o AccountIdentityOutput) ToAccountIdentityOutputWithContext(ctx context.Context) AccountIdentityOutput {
	return o
}

// The ID of the Principal (Client) in Azure Active Directory.
func (o AccountIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The ID of the Azure Active Directory Tenant.
func (o AccountIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of Managed Identity assigned to this Purview Account.
func (o AccountIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AccountIdentityArrayOutput struct{ *pulumi.OutputState }

func (AccountIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountIdentity)(nil)).Elem()
}

func (o AccountIdentityArrayOutput) ToAccountIdentityArrayOutput() AccountIdentityArrayOutput {
	return o
}

func (o AccountIdentityArrayOutput) ToAccountIdentityArrayOutputWithContext(ctx context.Context) AccountIdentityArrayOutput {
	return o
}

func (o AccountIdentityArrayOutput) Index(i pulumi.IntInput) AccountIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountIdentity {
		return vs[0].([]AccountIdentity)[vs[1].(int)]
	}).(AccountIdentityOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountIdentityOutput{})
	pulumi.RegisterOutputType(AccountIdentityArrayOutput{})
}
