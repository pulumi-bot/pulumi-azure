// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Spring Cloud Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appplatform"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appplatform.NewSpringCloudService(ctx, "exampleSpringCloudService", &appplatform.SpringCloudServiceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("S0"),
// 			ConfigServerGitSetting: &appplatform.SpringCloudServiceConfigServerGitSettingArgs{
// 				Uri:   pulumi.String("https://github.com/Azure-Samples/piggymetrics"),
// 				Label: pulumi.String("config"),
// 				SearchPaths: pulumi.StringArray{
// 					pulumi.String("dir1"),
// 					pulumi.String("dir2"),
// 				},
// 			},
// 			Trace: &appplatform.SpringCloudServiceTraceArgs{
// 				InstrumentationKey: exampleInsights.InstrumentationKey,
// 				SampleRate:         pulumi.Float64(10),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Env": pulumi.String("staging"),
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
// Spring Cloud services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appplatform/springCloudService:SpringCloudService example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/Spring/spring1
// ```
type SpringCloudService struct {
	pulumi.CustomResourceState

	// A `configServerGitSetting` block as defined below.
	ConfigServerGitSetting SpringCloudServiceConfigServerGitSettingPtrOutput `pulumi:"configServerGitSetting"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `network` block as defined below. Changing this forces a new resource to be created.
	Network SpringCloudServiceNetworkPtrOutput `pulumi:"network"`
	// A list of the outbound Public IP Addresses used by this Spring Cloud Service.
	OutboundPublicIpAddresses pulumi.StringArrayOutput `pulumi:"outboundPublicIpAddresses"`
	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this Spring Cloud Service. Possible values are `B0` and `S0`. Defaults to `S0`.
	SkuName pulumi.StringPtrOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `trace` block as defined below.
	Trace SpringCloudServiceTracePtrOutput `pulumi:"trace"`
}

// NewSpringCloudService registers a new resource with the given unique name, arguments, and options.
func NewSpringCloudService(ctx *pulumi.Context,
	name string, args *SpringCloudServiceArgs, opts ...pulumi.ResourceOption) (*SpringCloudService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SpringCloudService
	err := ctx.RegisterResource("azure:appplatform/springCloudService:SpringCloudService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringCloudService gets an existing SpringCloudService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringCloudService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringCloudServiceState, opts ...pulumi.ResourceOption) (*SpringCloudService, error) {
	var resource SpringCloudService
	err := ctx.ReadResource("azure:appplatform/springCloudService:SpringCloudService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpringCloudService resources.
type springCloudServiceState struct {
	// A `configServerGitSetting` block as defined below.
	ConfigServerGitSetting *SpringCloudServiceConfigServerGitSetting `pulumi:"configServerGitSetting"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `network` block as defined below. Changing this forces a new resource to be created.
	Network *SpringCloudServiceNetwork `pulumi:"network"`
	// A list of the outbound Public IP Addresses used by this Spring Cloud Service.
	OutboundPublicIpAddresses []string `pulumi:"outboundPublicIpAddresses"`
	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this Spring Cloud Service. Possible values are `B0` and `S0`. Defaults to `S0`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `trace` block as defined below.
	Trace *SpringCloudServiceTrace `pulumi:"trace"`
}

type SpringCloudServiceState struct {
	// A `configServerGitSetting` block as defined below.
	ConfigServerGitSetting SpringCloudServiceConfigServerGitSettingPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `network` block as defined below. Changing this forces a new resource to be created.
	Network SpringCloudServiceNetworkPtrInput
	// A list of the outbound Public IP Addresses used by this Spring Cloud Service.
	OutboundPublicIpAddresses pulumi.StringArrayInput
	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the SKU Name for this Spring Cloud Service. Possible values are `B0` and `S0`. Defaults to `S0`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `trace` block as defined below.
	Trace SpringCloudServiceTracePtrInput
}

func (SpringCloudServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudServiceState)(nil)).Elem()
}

type springCloudServiceArgs struct {
	// A `configServerGitSetting` block as defined below.
	ConfigServerGitSetting *SpringCloudServiceConfigServerGitSetting `pulumi:"configServerGitSetting"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `network` block as defined below. Changing this forces a new resource to be created.
	Network *SpringCloudServiceNetwork `pulumi:"network"`
	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this Spring Cloud Service. Possible values are `B0` and `S0`. Defaults to `S0`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `trace` block as defined below.
	Trace *SpringCloudServiceTrace `pulumi:"trace"`
}

// The set of arguments for constructing a SpringCloudService resource.
type SpringCloudServiceArgs struct {
	// A `configServerGitSetting` block as defined below.
	ConfigServerGitSetting SpringCloudServiceConfigServerGitSettingPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `network` block as defined below. Changing this forces a new resource to be created.
	Network SpringCloudServiceNetworkPtrInput
	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for this Spring Cloud Service. Possible values are `B0` and `S0`. Defaults to `S0`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `trace` block as defined below.
	Trace SpringCloudServiceTracePtrInput
}

func (SpringCloudServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudServiceArgs)(nil)).Elem()
}

type SpringCloudServiceInput interface {
	pulumi.Input

	ToSpringCloudServiceOutput() SpringCloudServiceOutput
	ToSpringCloudServiceOutputWithContext(ctx context.Context) SpringCloudServiceOutput
}

func (*SpringCloudService) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudService)(nil))
}

func (i *SpringCloudService) ToSpringCloudServiceOutput() SpringCloudServiceOutput {
	return i.ToSpringCloudServiceOutputWithContext(context.Background())
}

func (i *SpringCloudService) ToSpringCloudServiceOutputWithContext(ctx context.Context) SpringCloudServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudServiceOutput)
}

func (i *SpringCloudService) ToSpringCloudServicePtrOutput() SpringCloudServicePtrOutput {
	return i.ToSpringCloudServicePtrOutputWithContext(context.Background())
}

func (i *SpringCloudService) ToSpringCloudServicePtrOutputWithContext(ctx context.Context) SpringCloudServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudServicePtrOutput)
}

type SpringCloudServicePtrInput interface {
	pulumi.Input

	ToSpringCloudServicePtrOutput() SpringCloudServicePtrOutput
	ToSpringCloudServicePtrOutputWithContext(ctx context.Context) SpringCloudServicePtrOutput
}

type springCloudServicePtrType SpringCloudServiceArgs

func (*springCloudServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudService)(nil))
}

func (i *springCloudServicePtrType) ToSpringCloudServicePtrOutput() SpringCloudServicePtrOutput {
	return i.ToSpringCloudServicePtrOutputWithContext(context.Background())
}

func (i *springCloudServicePtrType) ToSpringCloudServicePtrOutputWithContext(ctx context.Context) SpringCloudServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudServicePtrOutput)
}

// SpringCloudServiceArrayInput is an input type that accepts SpringCloudServiceArray and SpringCloudServiceArrayOutput values.
// You can construct a concrete instance of `SpringCloudServiceArrayInput` via:
//
//          SpringCloudServiceArray{ SpringCloudServiceArgs{...} }
type SpringCloudServiceArrayInput interface {
	pulumi.Input

	ToSpringCloudServiceArrayOutput() SpringCloudServiceArrayOutput
	ToSpringCloudServiceArrayOutputWithContext(context.Context) SpringCloudServiceArrayOutput
}

type SpringCloudServiceArray []SpringCloudServiceInput

func (SpringCloudServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SpringCloudService)(nil))
}

func (i SpringCloudServiceArray) ToSpringCloudServiceArrayOutput() SpringCloudServiceArrayOutput {
	return i.ToSpringCloudServiceArrayOutputWithContext(context.Background())
}

func (i SpringCloudServiceArray) ToSpringCloudServiceArrayOutputWithContext(ctx context.Context) SpringCloudServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudServiceArrayOutput)
}

// SpringCloudServiceMapInput is an input type that accepts SpringCloudServiceMap and SpringCloudServiceMapOutput values.
// You can construct a concrete instance of `SpringCloudServiceMapInput` via:
//
//          SpringCloudServiceMap{ "key": SpringCloudServiceArgs{...} }
type SpringCloudServiceMapInput interface {
	pulumi.Input

	ToSpringCloudServiceMapOutput() SpringCloudServiceMapOutput
	ToSpringCloudServiceMapOutputWithContext(context.Context) SpringCloudServiceMapOutput
}

type SpringCloudServiceMap map[string]SpringCloudServiceInput

func (SpringCloudServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SpringCloudService)(nil))
}

func (i SpringCloudServiceMap) ToSpringCloudServiceMapOutput() SpringCloudServiceMapOutput {
	return i.ToSpringCloudServiceMapOutputWithContext(context.Background())
}

func (i SpringCloudServiceMap) ToSpringCloudServiceMapOutputWithContext(ctx context.Context) SpringCloudServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudServiceMapOutput)
}

type SpringCloudServiceOutput struct {
	*pulumi.OutputState
}

func (SpringCloudServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudService)(nil))
}

func (o SpringCloudServiceOutput) ToSpringCloudServiceOutput() SpringCloudServiceOutput {
	return o
}

func (o SpringCloudServiceOutput) ToSpringCloudServiceOutputWithContext(ctx context.Context) SpringCloudServiceOutput {
	return o
}

func (o SpringCloudServiceOutput) ToSpringCloudServicePtrOutput() SpringCloudServicePtrOutput {
	return o.ToSpringCloudServicePtrOutputWithContext(context.Background())
}

func (o SpringCloudServiceOutput) ToSpringCloudServicePtrOutputWithContext(ctx context.Context) SpringCloudServicePtrOutput {
	return o.ApplyT(func(v SpringCloudService) *SpringCloudService {
		return &v
	}).(SpringCloudServicePtrOutput)
}

type SpringCloudServicePtrOutput struct {
	*pulumi.OutputState
}

func (SpringCloudServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudService)(nil))
}

func (o SpringCloudServicePtrOutput) ToSpringCloudServicePtrOutput() SpringCloudServicePtrOutput {
	return o
}

func (o SpringCloudServicePtrOutput) ToSpringCloudServicePtrOutputWithContext(ctx context.Context) SpringCloudServicePtrOutput {
	return o
}

type SpringCloudServiceArrayOutput struct{ *pulumi.OutputState }

func (SpringCloudServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpringCloudService)(nil))
}

func (o SpringCloudServiceArrayOutput) ToSpringCloudServiceArrayOutput() SpringCloudServiceArrayOutput {
	return o
}

func (o SpringCloudServiceArrayOutput) ToSpringCloudServiceArrayOutputWithContext(ctx context.Context) SpringCloudServiceArrayOutput {
	return o
}

func (o SpringCloudServiceArrayOutput) Index(i pulumi.IntInput) SpringCloudServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpringCloudService {
		return vs[0].([]SpringCloudService)[vs[1].(int)]
	}).(SpringCloudServiceOutput)
}

type SpringCloudServiceMapOutput struct{ *pulumi.OutputState }

func (SpringCloudServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SpringCloudService)(nil))
}

func (o SpringCloudServiceMapOutput) ToSpringCloudServiceMapOutput() SpringCloudServiceMapOutput {
	return o
}

func (o SpringCloudServiceMapOutput) ToSpringCloudServiceMapOutputWithContext(ctx context.Context) SpringCloudServiceMapOutput {
	return o
}

func (o SpringCloudServiceMapOutput) MapIndex(k pulumi.StringInput) SpringCloudServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SpringCloudService {
		return vs[0].(map[string]SpringCloudService)[vs[1].(string)]
	}).(SpringCloudServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(SpringCloudServiceOutput{})
	pulumi.RegisterOutputType(SpringCloudServicePtrOutput{})
	pulumi.RegisterOutputType(SpringCloudServiceArrayOutput{})
	pulumi.RegisterOutputType(SpringCloudServiceMapOutput{})
}
