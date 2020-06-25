// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AnalyticsSolutionPlan struct {
	Name *string `pulumi:"name"`
	// The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
	Product string `pulumi:"product"`
	// A promotion code to be used with the solution.
	PromotionCode *string `pulumi:"promotionCode"`
	// The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
	Publisher string `pulumi:"publisher"`
}

// AnalyticsSolutionPlanInput is an input type that accepts AnalyticsSolutionPlanArgs and AnalyticsSolutionPlanOutput values.
// You can construct a concrete instance of `AnalyticsSolutionPlanInput` via:
//
//          AnalyticsSolutionPlanArgs{...}
type AnalyticsSolutionPlanInput interface {
	pulumi.Input

	ToAnalyticsSolutionPlanOutput() AnalyticsSolutionPlanOutput
	ToAnalyticsSolutionPlanOutputWithContext(context.Context) AnalyticsSolutionPlanOutput
}

type AnalyticsSolutionPlanArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
	Product pulumi.StringInput `pulumi:"product"`
	// A promotion code to be used with the solution.
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	// The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
	Publisher pulumi.StringInput `pulumi:"publisher"`
}

func (AnalyticsSolutionPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsSolutionPlan)(nil)).Elem()
}

func (i AnalyticsSolutionPlanArgs) ToAnalyticsSolutionPlanOutput() AnalyticsSolutionPlanOutput {
	return i.ToAnalyticsSolutionPlanOutputWithContext(context.Background())
}

func (i AnalyticsSolutionPlanArgs) ToAnalyticsSolutionPlanOutputWithContext(ctx context.Context) AnalyticsSolutionPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsSolutionPlanOutput)
}

func (i AnalyticsSolutionPlanArgs) ToAnalyticsSolutionPlanPtrOutput() AnalyticsSolutionPlanPtrOutput {
	return i.ToAnalyticsSolutionPlanPtrOutputWithContext(context.Background())
}

func (i AnalyticsSolutionPlanArgs) ToAnalyticsSolutionPlanPtrOutputWithContext(ctx context.Context) AnalyticsSolutionPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsSolutionPlanOutput).ToAnalyticsSolutionPlanPtrOutputWithContext(ctx)
}

// AnalyticsSolutionPlanPtrInput is an input type that accepts AnalyticsSolutionPlanArgs, AnalyticsSolutionPlanPtr and AnalyticsSolutionPlanPtrOutput values.
// You can construct a concrete instance of `AnalyticsSolutionPlanPtrInput` via:
//
//          AnalyticsSolutionPlanArgs{...}
//
//  or:
//
//          nil
type AnalyticsSolutionPlanPtrInput interface {
	pulumi.Input

	ToAnalyticsSolutionPlanPtrOutput() AnalyticsSolutionPlanPtrOutput
	ToAnalyticsSolutionPlanPtrOutputWithContext(context.Context) AnalyticsSolutionPlanPtrOutput
}

type analyticsSolutionPlanPtrType AnalyticsSolutionPlanArgs

func AnalyticsSolutionPlanPtr(v *AnalyticsSolutionPlanArgs) AnalyticsSolutionPlanPtrInput {
	return (*analyticsSolutionPlanPtrType)(v)
}

func (*analyticsSolutionPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsSolutionPlan)(nil)).Elem()
}

func (i *analyticsSolutionPlanPtrType) ToAnalyticsSolutionPlanPtrOutput() AnalyticsSolutionPlanPtrOutput {
	return i.ToAnalyticsSolutionPlanPtrOutputWithContext(context.Background())
}

func (i *analyticsSolutionPlanPtrType) ToAnalyticsSolutionPlanPtrOutputWithContext(ctx context.Context) AnalyticsSolutionPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsSolutionPlanPtrOutput)
}

type AnalyticsSolutionPlanOutput struct{ *pulumi.OutputState }

func (AnalyticsSolutionPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsSolutionPlan)(nil)).Elem()
}

func (o AnalyticsSolutionPlanOutput) ToAnalyticsSolutionPlanOutput() AnalyticsSolutionPlanOutput {
	return o
}

func (o AnalyticsSolutionPlanOutput) ToAnalyticsSolutionPlanOutputWithContext(ctx context.Context) AnalyticsSolutionPlanOutput {
	return o
}

func (o AnalyticsSolutionPlanOutput) ToAnalyticsSolutionPlanPtrOutput() AnalyticsSolutionPlanPtrOutput {
	return o.ToAnalyticsSolutionPlanPtrOutputWithContext(context.Background())
}

func (o AnalyticsSolutionPlanOutput) ToAnalyticsSolutionPlanPtrOutputWithContext(ctx context.Context) AnalyticsSolutionPlanPtrOutput {
	return o.ApplyT(func(v AnalyticsSolutionPlan) *AnalyticsSolutionPlan {
		return &v
	}).(AnalyticsSolutionPlanPtrOutput)
}
func (o AnalyticsSolutionPlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AnalyticsSolutionPlan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
func (o AnalyticsSolutionPlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v AnalyticsSolutionPlan) string { return v.Product }).(pulumi.StringOutput)
}

// A promotion code to be used with the solution.
func (o AnalyticsSolutionPlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AnalyticsSolutionPlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

// The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
func (o AnalyticsSolutionPlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v AnalyticsSolutionPlan) string { return v.Publisher }).(pulumi.StringOutput)
}

type AnalyticsSolutionPlanPtrOutput struct{ *pulumi.OutputState }

func (AnalyticsSolutionPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsSolutionPlan)(nil)).Elem()
}

func (o AnalyticsSolutionPlanPtrOutput) ToAnalyticsSolutionPlanPtrOutput() AnalyticsSolutionPlanPtrOutput {
	return o
}

func (o AnalyticsSolutionPlanPtrOutput) ToAnalyticsSolutionPlanPtrOutputWithContext(ctx context.Context) AnalyticsSolutionPlanPtrOutput {
	return o
}

func (o AnalyticsSolutionPlanPtrOutput) Elem() AnalyticsSolutionPlanOutput {
	return o.ApplyT(func(v *AnalyticsSolutionPlan) AnalyticsSolutionPlan { return *v }).(AnalyticsSolutionPlanOutput)
}

func (o AnalyticsSolutionPlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsSolutionPlan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
func (o AnalyticsSolutionPlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsSolutionPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

// A promotion code to be used with the solution.
func (o AnalyticsSolutionPlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsSolutionPlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

// The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
func (o AnalyticsSolutionPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsSolutionPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsSolutionPlanOutput{})
	pulumi.RegisterOutputType(AnalyticsSolutionPlanPtrOutput{})
}
