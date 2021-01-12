// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServiceCor struct {
	// A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}

// ServiceCorInput is an input type that accepts ServiceCorArgs and ServiceCorOutput values.
// You can construct a concrete instance of `ServiceCorInput` via:
//
//          ServiceCorArgs{...}
type ServiceCorInput interface {
	pulumi.Input

	ToServiceCorOutput() ServiceCorOutput
	ToServiceCorOutputWithContext(context.Context) ServiceCorOutput
}

type ServiceCorArgs struct {
	// A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (ServiceCorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCor)(nil)).Elem()
}

func (i ServiceCorArgs) ToServiceCorOutput() ServiceCorOutput {
	return i.ToServiceCorOutputWithContext(context.Background())
}

func (i ServiceCorArgs) ToServiceCorOutputWithContext(ctx context.Context) ServiceCorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorOutput)
}

// ServiceCorArrayInput is an input type that accepts ServiceCorArray and ServiceCorArrayOutput values.
// You can construct a concrete instance of `ServiceCorArrayInput` via:
//
//          ServiceCorArray{ ServiceCorArgs{...} }
type ServiceCorArrayInput interface {
	pulumi.Input

	ToServiceCorArrayOutput() ServiceCorArrayOutput
	ToServiceCorArrayOutputWithContext(context.Context) ServiceCorArrayOutput
}

type ServiceCorArray []ServiceCorInput

func (ServiceCorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCor)(nil)).Elem()
}

func (i ServiceCorArray) ToServiceCorArrayOutput() ServiceCorArrayOutput {
	return i.ToServiceCorArrayOutputWithContext(context.Background())
}

func (i ServiceCorArray) ToServiceCorArrayOutputWithContext(ctx context.Context) ServiceCorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorArrayOutput)
}

type ServiceCorOutput struct{ *pulumi.OutputState }

func (ServiceCorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCor)(nil)).Elem()
}

func (o ServiceCorOutput) ToServiceCorOutput() ServiceCorOutput {
	return o
}

func (o ServiceCorOutput) ToServiceCorOutputWithContext(ctx context.Context) ServiceCorOutput {
	return o
}

// A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
func (o ServiceCorOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCor) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type ServiceCorArrayOutput struct{ *pulumi.OutputState }

func (ServiceCorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCor)(nil)).Elem()
}

func (o ServiceCorArrayOutput) ToServiceCorArrayOutput() ServiceCorArrayOutput {
	return o
}

func (o ServiceCorArrayOutput) ToServiceCorArrayOutputWithContext(ctx context.Context) ServiceCorArrayOutput {
	return o
}

func (o ServiceCorArrayOutput) Index(i pulumi.IntInput) ServiceCorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceCor {
		return vs[0].([]ServiceCor)[vs[1].(int)]
	}).(ServiceCorOutput)
}

type ServiceFeature struct {
	// The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
	Flag string `pulumi:"flag"`
	// A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
	Value string `pulumi:"value"`
}

// ServiceFeatureInput is an input type that accepts ServiceFeatureArgs and ServiceFeatureOutput values.
// You can construct a concrete instance of `ServiceFeatureInput` via:
//
//          ServiceFeatureArgs{...}
type ServiceFeatureInput interface {
	pulumi.Input

	ToServiceFeatureOutput() ServiceFeatureOutput
	ToServiceFeatureOutputWithContext(context.Context) ServiceFeatureOutput
}

type ServiceFeatureArgs struct {
	// The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
	Flag pulumi.StringInput `pulumi:"flag"`
	// A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ServiceFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceFeature)(nil)).Elem()
}

func (i ServiceFeatureArgs) ToServiceFeatureOutput() ServiceFeatureOutput {
	return i.ToServiceFeatureOutputWithContext(context.Background())
}

func (i ServiceFeatureArgs) ToServiceFeatureOutputWithContext(ctx context.Context) ServiceFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceFeatureOutput)
}

// ServiceFeatureArrayInput is an input type that accepts ServiceFeatureArray and ServiceFeatureArrayOutput values.
// You can construct a concrete instance of `ServiceFeatureArrayInput` via:
//
//          ServiceFeatureArray{ ServiceFeatureArgs{...} }
type ServiceFeatureArrayInput interface {
	pulumi.Input

	ToServiceFeatureArrayOutput() ServiceFeatureArrayOutput
	ToServiceFeatureArrayOutputWithContext(context.Context) ServiceFeatureArrayOutput
}

type ServiceFeatureArray []ServiceFeatureInput

func (ServiceFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceFeature)(nil)).Elem()
}

func (i ServiceFeatureArray) ToServiceFeatureArrayOutput() ServiceFeatureArrayOutput {
	return i.ToServiceFeatureArrayOutputWithContext(context.Background())
}

func (i ServiceFeatureArray) ToServiceFeatureArrayOutputWithContext(ctx context.Context) ServiceFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceFeatureArrayOutput)
}

type ServiceFeatureOutput struct{ *pulumi.OutputState }

func (ServiceFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceFeature)(nil)).Elem()
}

func (o ServiceFeatureOutput) ToServiceFeatureOutput() ServiceFeatureOutput {
	return o
}

func (o ServiceFeatureOutput) ToServiceFeatureOutputWithContext(ctx context.Context) ServiceFeatureOutput {
	return o
}

// The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
func (o ServiceFeatureOutput) Flag() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceFeature) string { return v.Flag }).(pulumi.StringOutput)
}

// A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
func (o ServiceFeatureOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceFeature) string { return v.Value }).(pulumi.StringOutput)
}

type ServiceFeatureArrayOutput struct{ *pulumi.OutputState }

func (ServiceFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceFeature)(nil)).Elem()
}

func (o ServiceFeatureArrayOutput) ToServiceFeatureArrayOutput() ServiceFeatureArrayOutput {
	return o
}

func (o ServiceFeatureArrayOutput) ToServiceFeatureArrayOutputWithContext(ctx context.Context) ServiceFeatureArrayOutput {
	return o
}

func (o ServiceFeatureArrayOutput) Index(i pulumi.IntInput) ServiceFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceFeature {
		return vs[0].([]ServiceFeature)[vs[1].(int)]
	}).(ServiceFeatureOutput)
}

type ServiceSku struct {
	// Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
	Capacity int `pulumi:"capacity"`
	// Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
	Name string `pulumi:"name"`
}

// ServiceSkuInput is an input type that accepts ServiceSkuArgs and ServiceSkuOutput values.
// You can construct a concrete instance of `ServiceSkuInput` via:
//
//          ServiceSkuArgs{...}
type ServiceSkuInput interface {
	pulumi.Input

	ToServiceSkuOutput() ServiceSkuOutput
	ToServiceSkuOutputWithContext(context.Context) ServiceSkuOutput
}

type ServiceSkuArgs struct {
	// Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
	Name pulumi.StringInput `pulumi:"name"`
}

func (ServiceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSku)(nil)).Elem()
}

func (i ServiceSkuArgs) ToServiceSkuOutput() ServiceSkuOutput {
	return i.ToServiceSkuOutputWithContext(context.Background())
}

func (i ServiceSkuArgs) ToServiceSkuOutputWithContext(ctx context.Context) ServiceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSkuOutput)
}

func (i ServiceSkuArgs) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return i.ToServiceSkuPtrOutputWithContext(context.Background())
}

func (i ServiceSkuArgs) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSkuOutput).ToServiceSkuPtrOutput()
}

// ServiceSkuPtrInput is an input type that accepts ServiceSkuArgs, ServiceSkuPtr and ServiceSkuPtrOutput values.
// You can construct a concrete instance of `ServiceSkuPtrInput` via:
//
//          ServiceSkuArgs{...}
//
//  or:
//
//          nil
type ServiceSkuPtrInput interface {
	pulumi.Input

	ToServiceSkuPtrOutput() ServiceSkuPtrOutput
	ToServiceSkuPtrOutputWithContext(context.Context) ServiceSkuPtrOutput
}

type serviceSkuPtrType ServiceSkuArgs

func ServiceSkuPtr(v *ServiceSkuArgs) ServiceSkuPtrInput {
	return (*serviceSkuPtrType)(v)
}

func (*serviceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSku)(nil)).Elem()
}

func (i *serviceSkuPtrType) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return i.ToServiceSkuPtrOutputWithContext(context.Background())
}

func (i *serviceSkuPtrType) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSkuOutput).ToServiceSkuPtrOutput()
}

type ServiceSkuOutput struct{ *pulumi.OutputState }

func (ServiceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSku)(nil)).Elem()
}

func (o ServiceSkuOutput) ToServiceSkuOutput() ServiceSkuOutput {
	return o
}

func (o ServiceSkuOutput) ToServiceSkuOutputWithContext(ctx context.Context) ServiceSkuOutput {
	return o
}

func (o ServiceSkuOutput) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return o.ToServiceSkuPtrOutputWithContext(context.Background())
}

func (o ServiceSkuOutput) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return o.ApplyT(func(v ServiceSku) *ServiceSku {
		return &v
	}).(ServiceSkuPtrOutput)
}

// Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
func (o ServiceSkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceSku) int { return v.Capacity }).(pulumi.IntOutput)
}

// Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
func (o ServiceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceSku) string { return v.Name }).(pulumi.StringOutput)
}

type ServiceSkuPtrOutput struct{ *pulumi.OutputState }

func (ServiceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSku)(nil)).Elem()
}

func (o ServiceSkuPtrOutput) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return o
}

func (o ServiceSkuPtrOutput) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return o
}

func (o ServiceSkuPtrOutput) Elem() ServiceSkuOutput {
	return o.ApplyT(func(v *ServiceSku) ServiceSku { return *v }).(ServiceSkuOutput)
}

// Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
func (o ServiceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

// Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
func (o ServiceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceCorOutput{})
	pulumi.RegisterOutputType(ServiceCorArrayOutput{})
	pulumi.RegisterOutputType(ServiceFeatureOutput{})
	pulumi.RegisterOutputType(ServiceFeatureArrayOutput{})
	pulumi.RegisterOutputType(ServiceSkuOutput{})
	pulumi.RegisterOutputType(ServiceSkuPtrOutput{})
}
