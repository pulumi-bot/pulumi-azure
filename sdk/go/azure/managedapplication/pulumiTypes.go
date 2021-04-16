// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managedapplication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationPlan struct {
	// Specifies the name of the plan from the marketplace.
	Name string `pulumi:"name"`
	// Specifies the product of the plan from the marketplace.
	Product string `pulumi:"product"`
	// Specifies the promotion code to use with the plan.
	PromotionCode *string `pulumi:"promotionCode"`
	// Specifies the publisher of the plan.
	Publisher string `pulumi:"publisher"`
	// Specifies the version of the plan from the marketplace.
	Version string `pulumi:"version"`
}

// ApplicationPlanInput is an input type that accepts ApplicationPlanArgs and ApplicationPlanOutput values.
// You can construct a concrete instance of `ApplicationPlanInput` via:
//
//          ApplicationPlanArgs{...}
type ApplicationPlanInput interface {
	pulumi.Input

	ToApplicationPlanOutput() ApplicationPlanOutput
	ToApplicationPlanOutputWithContext(context.Context) ApplicationPlanOutput
}

type ApplicationPlanArgs struct {
	// Specifies the name of the plan from the marketplace.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the product of the plan from the marketplace.
	Product pulumi.StringInput `pulumi:"product"`
	// Specifies the promotion code to use with the plan.
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	// Specifies the publisher of the plan.
	Publisher pulumi.StringInput `pulumi:"publisher"`
	// Specifies the version of the plan from the marketplace.
	Version pulumi.StringInput `pulumi:"version"`
}

func (ApplicationPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPlan)(nil)).Elem()
}

func (i ApplicationPlanArgs) ToApplicationPlanOutput() ApplicationPlanOutput {
	return i.ToApplicationPlanOutputWithContext(context.Background())
}

func (i ApplicationPlanArgs) ToApplicationPlanOutputWithContext(ctx context.Context) ApplicationPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPlanOutput)
}

func (i ApplicationPlanArgs) ToApplicationPlanPtrOutput() ApplicationPlanPtrOutput {
	return i.ToApplicationPlanPtrOutputWithContext(context.Background())
}

func (i ApplicationPlanArgs) ToApplicationPlanPtrOutputWithContext(ctx context.Context) ApplicationPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPlanOutput).ToApplicationPlanPtrOutputWithContext(ctx)
}

// ApplicationPlanPtrInput is an input type that accepts ApplicationPlanArgs, ApplicationPlanPtr and ApplicationPlanPtrOutput values.
// You can construct a concrete instance of `ApplicationPlanPtrInput` via:
//
//          ApplicationPlanArgs{...}
//
//  or:
//
//          nil
type ApplicationPlanPtrInput interface {
	pulumi.Input

	ToApplicationPlanPtrOutput() ApplicationPlanPtrOutput
	ToApplicationPlanPtrOutputWithContext(context.Context) ApplicationPlanPtrOutput
}

type applicationPlanPtrType ApplicationPlanArgs

func ApplicationPlanPtr(v *ApplicationPlanArgs) ApplicationPlanPtrInput {
	return (*applicationPlanPtrType)(v)
}

func (*applicationPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPlan)(nil)).Elem()
}

func (i *applicationPlanPtrType) ToApplicationPlanPtrOutput() ApplicationPlanPtrOutput {
	return i.ToApplicationPlanPtrOutputWithContext(context.Background())
}

func (i *applicationPlanPtrType) ToApplicationPlanPtrOutputWithContext(ctx context.Context) ApplicationPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPlanPtrOutput)
}

type ApplicationPlanOutput struct{ *pulumi.OutputState }

func (ApplicationPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPlan)(nil)).Elem()
}

func (o ApplicationPlanOutput) ToApplicationPlanOutput() ApplicationPlanOutput {
	return o
}

func (o ApplicationPlanOutput) ToApplicationPlanOutputWithContext(ctx context.Context) ApplicationPlanOutput {
	return o
}

func (o ApplicationPlanOutput) ToApplicationPlanPtrOutput() ApplicationPlanPtrOutput {
	return o.ToApplicationPlanPtrOutputWithContext(context.Background())
}

func (o ApplicationPlanOutput) ToApplicationPlanPtrOutputWithContext(ctx context.Context) ApplicationPlanPtrOutput {
	return o.ApplyT(func(v ApplicationPlan) *ApplicationPlan {
		return &v
	}).(ApplicationPlanPtrOutput)
}

// Specifies the name of the plan from the marketplace.
func (o ApplicationPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPlan) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the product of the plan from the marketplace.
func (o ApplicationPlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPlan) string { return v.Product }).(pulumi.StringOutput)
}

// Specifies the promotion code to use with the plan.
func (o ApplicationPlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

// Specifies the publisher of the plan.
func (o ApplicationPlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPlan) string { return v.Publisher }).(pulumi.StringOutput)
}

// Specifies the version of the plan from the marketplace.
func (o ApplicationPlanOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPlan) string { return v.Version }).(pulumi.StringOutput)
}

type ApplicationPlanPtrOutput struct{ *pulumi.OutputState }

func (ApplicationPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPlan)(nil)).Elem()
}

func (o ApplicationPlanPtrOutput) ToApplicationPlanPtrOutput() ApplicationPlanPtrOutput {
	return o
}

func (o ApplicationPlanPtrOutput) ToApplicationPlanPtrOutputWithContext(ctx context.Context) ApplicationPlanPtrOutput {
	return o
}

func (o ApplicationPlanPtrOutput) Elem() ApplicationPlanOutput {
	return o.ApplyT(func(v *ApplicationPlan) ApplicationPlan { return *v }).(ApplicationPlanOutput)
}

// Specifies the name of the plan from the marketplace.
func (o ApplicationPlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Specifies the product of the plan from the marketplace.
func (o ApplicationPlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

// Specifies the promotion code to use with the plan.
func (o ApplicationPlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

// Specifies the publisher of the plan.
func (o ApplicationPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

// Specifies the version of the plan from the marketplace.
func (o ApplicationPlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type DefinitionAuthorization struct {
	// Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
}

// DefinitionAuthorizationInput is an input type that accepts DefinitionAuthorizationArgs and DefinitionAuthorizationOutput values.
// You can construct a concrete instance of `DefinitionAuthorizationInput` via:
//
//          DefinitionAuthorizationArgs{...}
type DefinitionAuthorizationInput interface {
	pulumi.Input

	ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput
	ToDefinitionAuthorizationOutputWithContext(context.Context) DefinitionAuthorizationOutput
}

type DefinitionAuthorizationArgs struct {
	// Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
	// Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
	ServicePrincipalId pulumi.StringInput `pulumi:"servicePrincipalId"`
}

func (DefinitionAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionAuthorization)(nil)).Elem()
}

func (i DefinitionAuthorizationArgs) ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput {
	return i.ToDefinitionAuthorizationOutputWithContext(context.Background())
}

func (i DefinitionAuthorizationArgs) ToDefinitionAuthorizationOutputWithContext(ctx context.Context) DefinitionAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionAuthorizationOutput)
}

// DefinitionAuthorizationArrayInput is an input type that accepts DefinitionAuthorizationArray and DefinitionAuthorizationArrayOutput values.
// You can construct a concrete instance of `DefinitionAuthorizationArrayInput` via:
//
//          DefinitionAuthorizationArray{ DefinitionAuthorizationArgs{...} }
type DefinitionAuthorizationArrayInput interface {
	pulumi.Input

	ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput
	ToDefinitionAuthorizationArrayOutputWithContext(context.Context) DefinitionAuthorizationArrayOutput
}

type DefinitionAuthorizationArray []DefinitionAuthorizationInput

func (DefinitionAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefinitionAuthorization)(nil)).Elem()
}

func (i DefinitionAuthorizationArray) ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput {
	return i.ToDefinitionAuthorizationArrayOutputWithContext(context.Background())
}

func (i DefinitionAuthorizationArray) ToDefinitionAuthorizationArrayOutputWithContext(ctx context.Context) DefinitionAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionAuthorizationArrayOutput)
}

type DefinitionAuthorizationOutput struct{ *pulumi.OutputState }

func (DefinitionAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionAuthorization)(nil)).Elem()
}

func (o DefinitionAuthorizationOutput) ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput {
	return o
}

func (o DefinitionAuthorizationOutput) ToDefinitionAuthorizationOutputWithContext(ctx context.Context) DefinitionAuthorizationOutput {
	return o
}

// Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
func (o DefinitionAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

// Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
func (o DefinitionAuthorizationOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionAuthorization) string { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

type DefinitionAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (DefinitionAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefinitionAuthorization)(nil)).Elem()
}

func (o DefinitionAuthorizationArrayOutput) ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput {
	return o
}

func (o DefinitionAuthorizationArrayOutput) ToDefinitionAuthorizationArrayOutputWithContext(ctx context.Context) DefinitionAuthorizationArrayOutput {
	return o
}

func (o DefinitionAuthorizationArrayOutput) Index(i pulumi.IntInput) DefinitionAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DefinitionAuthorization {
		return vs[0].([]DefinitionAuthorization)[vs[1].(int)]
	}).(DefinitionAuthorizationOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationPlanOutput{})
	pulumi.RegisterOutputType(ApplicationPlanPtrOutput{})
	pulumi.RegisterOutputType(DefinitionAuthorizationOutput{})
	pulumi.RegisterOutputType(DefinitionAuthorizationArrayOutput{})
}
