// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServiceIdentity struct {
	// The (Client) ID of the Service Principal.
	PrincipalId *string `pulumi:"principalId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId *string `pulumi:"tenantId"`
	// The Type of Identity which should be used for the Search Service. At this time the only possible value is `SystemAssigned`.
	Type string `pulumi:"type"`
}

// ServiceIdentityInput is an input type that accepts ServiceIdentityArgs and ServiceIdentityOutput values.
// You can construct a concrete instance of `ServiceIdentityInput` via:
//
//          ServiceIdentityArgs{...}
type ServiceIdentityInput interface {
	pulumi.Input

	ToServiceIdentityOutput() ServiceIdentityOutput
	ToServiceIdentityOutputWithContext(context.Context) ServiceIdentityOutput
}

type ServiceIdentityArgs struct {
	// The (Client) ID of the Service Principal.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// The Type of Identity which should be used for the Search Service. At this time the only possible value is `SystemAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIdentity)(nil)).Elem()
}

func (i ServiceIdentityArgs) ToServiceIdentityOutput() ServiceIdentityOutput {
	return i.ToServiceIdentityOutputWithContext(context.Background())
}

func (i ServiceIdentityArgs) ToServiceIdentityOutputWithContext(ctx context.Context) ServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIdentityOutput)
}

func (i ServiceIdentityArgs) ToServiceIdentityPtrOutput() ServiceIdentityPtrOutput {
	return i.ToServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ServiceIdentityArgs) ToServiceIdentityPtrOutputWithContext(ctx context.Context) ServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIdentityOutput).ToServiceIdentityPtrOutput()
}

// ServiceIdentityPtrInput is an input type that accepts ServiceIdentityArgs, ServiceIdentityPtr and ServiceIdentityPtrOutput values.
// You can construct a concrete instance of `ServiceIdentityPtrInput` via:
//
//          ServiceIdentityArgs{...}
//
//  or:
//
//          nil
type ServiceIdentityPtrInput interface {
	pulumi.Input

	ToServiceIdentityPtrOutput() ServiceIdentityPtrOutput
	ToServiceIdentityPtrOutputWithContext(context.Context) ServiceIdentityPtrOutput
}

type serviceIdentityPtrType ServiceIdentityArgs

func ServiceIdentityPtr(v *ServiceIdentityArgs) ServiceIdentityPtrInput {
	return (*serviceIdentityPtrType)(v)
}

func (*serviceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIdentity)(nil)).Elem()
}

func (i *serviceIdentityPtrType) ToServiceIdentityPtrOutput() ServiceIdentityPtrOutput {
	return i.ToServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *serviceIdentityPtrType) ToServiceIdentityPtrOutputWithContext(ctx context.Context) ServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIdentityOutput).ToServiceIdentityPtrOutput()
}

type ServiceIdentityOutput struct{ *pulumi.OutputState }

func (ServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIdentity)(nil)).Elem()
}

func (o ServiceIdentityOutput) ToServiceIdentityOutput() ServiceIdentityOutput {
	return o
}

func (o ServiceIdentityOutput) ToServiceIdentityOutputWithContext(ctx context.Context) ServiceIdentityOutput {
	return o
}

func (o ServiceIdentityOutput) ToServiceIdentityPtrOutput() ServiceIdentityPtrOutput {
	return o.ToServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ServiceIdentityOutput) ToServiceIdentityPtrOutputWithContext(ctx context.Context) ServiceIdentityPtrOutput {
	return o.ApplyT(func(v ServiceIdentity) *ServiceIdentity {
		return &v
	}).(ServiceIdentityPtrOutput)
}

// The (Client) ID of the Service Principal.
func (o ServiceIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The ID of the Tenant the Service Principal is assigned in.
func (o ServiceIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for the Search Service. At this time the only possible value is `SystemAssigned`.
func (o ServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIdentity)(nil)).Elem()
}

func (o ServiceIdentityPtrOutput) ToServiceIdentityPtrOutput() ServiceIdentityPtrOutput {
	return o
}

func (o ServiceIdentityPtrOutput) ToServiceIdentityPtrOutputWithContext(ctx context.Context) ServiceIdentityPtrOutput {
	return o
}

func (o ServiceIdentityPtrOutput) Elem() ServiceIdentityOutput {
	return o.ApplyT(func(v *ServiceIdentity) ServiceIdentity { return *v }).(ServiceIdentityOutput)
}

// The (Client) ID of the Service Principal.
func (o ServiceIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The ID of the Tenant the Service Principal is assigned in.
func (o ServiceIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for the Search Service. At this time the only possible value is `SystemAssigned`.
func (o ServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ServiceQueryKey struct {
	// The value of this Query Key.
	Key *string `pulumi:"key"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name *string `pulumi:"name"`
}

// ServiceQueryKeyInput is an input type that accepts ServiceQueryKeyArgs and ServiceQueryKeyOutput values.
// You can construct a concrete instance of `ServiceQueryKeyInput` via:
//
//          ServiceQueryKeyArgs{...}
type ServiceQueryKeyInput interface {
	pulumi.Input

	ToServiceQueryKeyOutput() ServiceQueryKeyOutput
	ToServiceQueryKeyOutputWithContext(context.Context) ServiceQueryKeyOutput
}

type ServiceQueryKeyArgs struct {
	// The value of this Query Key.
	Key pulumi.StringPtrInput `pulumi:"key"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ServiceQueryKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQueryKey)(nil)).Elem()
}

func (i ServiceQueryKeyArgs) ToServiceQueryKeyOutput() ServiceQueryKeyOutput {
	return i.ToServiceQueryKeyOutputWithContext(context.Background())
}

func (i ServiceQueryKeyArgs) ToServiceQueryKeyOutputWithContext(ctx context.Context) ServiceQueryKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQueryKeyOutput)
}

// ServiceQueryKeyArrayInput is an input type that accepts ServiceQueryKeyArray and ServiceQueryKeyArrayOutput values.
// You can construct a concrete instance of `ServiceQueryKeyArrayInput` via:
//
//          ServiceQueryKeyArray{ ServiceQueryKeyArgs{...} }
type ServiceQueryKeyArrayInput interface {
	pulumi.Input

	ToServiceQueryKeyArrayOutput() ServiceQueryKeyArrayOutput
	ToServiceQueryKeyArrayOutputWithContext(context.Context) ServiceQueryKeyArrayOutput
}

type ServiceQueryKeyArray []ServiceQueryKeyInput

func (ServiceQueryKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQueryKey)(nil)).Elem()
}

func (i ServiceQueryKeyArray) ToServiceQueryKeyArrayOutput() ServiceQueryKeyArrayOutput {
	return i.ToServiceQueryKeyArrayOutputWithContext(context.Background())
}

func (i ServiceQueryKeyArray) ToServiceQueryKeyArrayOutputWithContext(ctx context.Context) ServiceQueryKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQueryKeyArrayOutput)
}

type ServiceQueryKeyOutput struct{ *pulumi.OutputState }

func (ServiceQueryKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQueryKey)(nil)).Elem()
}

func (o ServiceQueryKeyOutput) ToServiceQueryKeyOutput() ServiceQueryKeyOutput {
	return o
}

func (o ServiceQueryKeyOutput) ToServiceQueryKeyOutputWithContext(ctx context.Context) ServiceQueryKeyOutput {
	return o
}

// The value of this Query Key.
func (o ServiceQueryKeyOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQueryKey) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
func (o ServiceQueryKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQueryKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ServiceQueryKeyArrayOutput struct{ *pulumi.OutputState }

func (ServiceQueryKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQueryKey)(nil)).Elem()
}

func (o ServiceQueryKeyArrayOutput) ToServiceQueryKeyArrayOutput() ServiceQueryKeyArrayOutput {
	return o
}

func (o ServiceQueryKeyArrayOutput) ToServiceQueryKeyArrayOutputWithContext(ctx context.Context) ServiceQueryKeyArrayOutput {
	return o
}

func (o ServiceQueryKeyArrayOutput) Index(i pulumi.IntInput) ServiceQueryKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceQueryKey {
		return vs[0].([]ServiceQueryKey)[vs[1].(int)]
	}).(ServiceQueryKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceIdentityOutput{})
	pulumi.RegisterOutputType(ServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServiceQueryKeyOutput{})
	pulumi.RegisterOutputType(ServiceQueryKeyArrayOutput{})
}
