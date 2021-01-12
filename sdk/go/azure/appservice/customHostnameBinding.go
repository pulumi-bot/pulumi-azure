// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Hostname Binding within an App Service (or Function App).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
// 			Keepers: pulumi.Float64Map{
// 				"azi_id": pulumi.Float64(1),
// 			},
// 			ByteLength: pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAppService, err := appservice.NewAppService(ctx, "exampleAppService", &appservice.AppServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewCustomHostnameBinding(ctx, "exampleCustomHostnameBinding", &appservice.CustomHostnameBindingArgs{
// 			Hostname:          pulumi.String("www.mywebsite.com"),
// 			AppServiceName:    exampleAppService.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// App Service Custom Hostname Bindings can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/customHostnameBinding:CustomHostnameBinding mywebsite /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com
// ```
type CustomHostnameBinding struct {
	pulumi.CustomResourceState

	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState pulumi.StringOutput `pulumi:"sslState"`
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIp pulumi.StringOutput `pulumi:"virtualIp"`
}

// NewCustomHostnameBinding registers a new resource with the given unique name, arguments, and options.
func NewCustomHostnameBinding(ctx *pulumi.Context,
	name string, args *CustomHostnameBindingArgs, opts ...pulumi.ResourceOption) (*CustomHostnameBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceName == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceName'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CustomHostnameBinding
	err := ctx.RegisterResource("azure:appservice/customHostnameBinding:CustomHostnameBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomHostnameBinding gets an existing CustomHostnameBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomHostnameBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomHostnameBindingState, opts ...pulumi.ResourceOption) (*CustomHostnameBinding, error) {
	var resource CustomHostnameBinding
	err := ctx.ReadResource("azure:appservice/customHostnameBinding:CustomHostnameBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomHostnameBinding resources.
type customHostnameBindingState struct {
	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname *string `pulumi:"hostname"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState *string `pulumi:"sslState"`
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint *string `pulumi:"thumbprint"`
	// The virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIp *string `pulumi:"virtualIp"`
}

type CustomHostnameBindingState struct {
	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName pulumi.StringPtrInput
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrInput
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState pulumi.StringPtrInput
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint pulumi.StringPtrInput
	// The virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIp pulumi.StringPtrInput
}

func (CustomHostnameBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*customHostnameBindingState)(nil)).Elem()
}

type customHostnameBindingArgs struct {
	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname string `pulumi:"hostname"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState *string `pulumi:"sslState"`
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint *string `pulumi:"thumbprint"`
}

// The set of arguments for constructing a CustomHostnameBinding resource.
type CustomHostnameBindingArgs struct {
	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName pulumi.StringInput
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname pulumi.StringInput
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState pulumi.StringPtrInput
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint pulumi.StringPtrInput
}

func (CustomHostnameBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customHostnameBindingArgs)(nil)).Elem()
}

type CustomHostnameBindingInput interface {
	pulumi.Input

	ToCustomHostnameBindingOutput() CustomHostnameBindingOutput
	ToCustomHostnameBindingOutputWithContext(ctx context.Context) CustomHostnameBindingOutput
}

func (*CustomHostnameBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostnameBinding)(nil))
}

func (i *CustomHostnameBinding) ToCustomHostnameBindingOutput() CustomHostnameBindingOutput {
	return i.ToCustomHostnameBindingOutputWithContext(context.Background())
}

func (i *CustomHostnameBinding) ToCustomHostnameBindingOutputWithContext(ctx context.Context) CustomHostnameBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHostnameBindingOutput)
}

func (i *CustomHostnameBinding) ToCustomHostnameBindingPtrOutput() CustomHostnameBindingPtrOutput {
	return i.ToCustomHostnameBindingPtrOutputWithContext(context.Background())
}

func (i *CustomHostnameBinding) ToCustomHostnameBindingPtrOutputWithContext(ctx context.Context) CustomHostnameBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHostnameBindingPtrOutput)
}

type CustomHostnameBindingPtrInput interface {
	pulumi.Input

	ToCustomHostnameBindingPtrOutput() CustomHostnameBindingPtrOutput
	ToCustomHostnameBindingPtrOutputWithContext(ctx context.Context) CustomHostnameBindingPtrOutput
}

type customHostnameBindingPtrType CustomHostnameBindingArgs

func (*customHostnameBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomHostnameBinding)(nil))
}

func (i *customHostnameBindingPtrType) ToCustomHostnameBindingPtrOutput() CustomHostnameBindingPtrOutput {
	return i.ToCustomHostnameBindingPtrOutputWithContext(context.Background())
}

func (i *customHostnameBindingPtrType) ToCustomHostnameBindingPtrOutputWithContext(ctx context.Context) CustomHostnameBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHostnameBindingOutput).ToCustomHostnameBindingPtrOutput()
}

// CustomHostnameBindingArrayInput is an input type that accepts CustomHostnameBindingArray and CustomHostnameBindingArrayOutput values.
// You can construct a concrete instance of `CustomHostnameBindingArrayInput` via:
//
//          CustomHostnameBindingArray{ CustomHostnameBindingArgs{...} }
type CustomHostnameBindingArrayInput interface {
	pulumi.Input

	ToCustomHostnameBindingArrayOutput() CustomHostnameBindingArrayOutput
	ToCustomHostnameBindingArrayOutputWithContext(context.Context) CustomHostnameBindingArrayOutput
}

type CustomHostnameBindingArray []CustomHostnameBindingInput

func (CustomHostnameBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomHostnameBinding)(nil))
}

func (i CustomHostnameBindingArray) ToCustomHostnameBindingArrayOutput() CustomHostnameBindingArrayOutput {
	return i.ToCustomHostnameBindingArrayOutputWithContext(context.Background())
}

func (i CustomHostnameBindingArray) ToCustomHostnameBindingArrayOutputWithContext(ctx context.Context) CustomHostnameBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHostnameBindingArrayOutput)
}

// CustomHostnameBindingMapInput is an input type that accepts CustomHostnameBindingMap and CustomHostnameBindingMapOutput values.
// You can construct a concrete instance of `CustomHostnameBindingMapInput` via:
//
//          CustomHostnameBindingMap{ "key": CustomHostnameBindingArgs{...} }
type CustomHostnameBindingMapInput interface {
	pulumi.Input

	ToCustomHostnameBindingMapOutput() CustomHostnameBindingMapOutput
	ToCustomHostnameBindingMapOutputWithContext(context.Context) CustomHostnameBindingMapOutput
}

type CustomHostnameBindingMap map[string]CustomHostnameBindingInput

func (CustomHostnameBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomHostnameBinding)(nil))
}

func (i CustomHostnameBindingMap) ToCustomHostnameBindingMapOutput() CustomHostnameBindingMapOutput {
	return i.ToCustomHostnameBindingMapOutputWithContext(context.Background())
}

func (i CustomHostnameBindingMap) ToCustomHostnameBindingMapOutputWithContext(ctx context.Context) CustomHostnameBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHostnameBindingMapOutput)
}

type CustomHostnameBindingOutput struct {
	*pulumi.OutputState
}

func (CustomHostnameBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostnameBinding)(nil))
}

func (o CustomHostnameBindingOutput) ToCustomHostnameBindingOutput() CustomHostnameBindingOutput {
	return o
}

func (o CustomHostnameBindingOutput) ToCustomHostnameBindingOutputWithContext(ctx context.Context) CustomHostnameBindingOutput {
	return o
}

func (o CustomHostnameBindingOutput) ToCustomHostnameBindingPtrOutput() CustomHostnameBindingPtrOutput {
	return o.ToCustomHostnameBindingPtrOutputWithContext(context.Background())
}

func (o CustomHostnameBindingOutput) ToCustomHostnameBindingPtrOutputWithContext(ctx context.Context) CustomHostnameBindingPtrOutput {
	return o.ApplyT(func(v CustomHostnameBinding) *CustomHostnameBinding {
		return &v
	}).(CustomHostnameBindingPtrOutput)
}

type CustomHostnameBindingPtrOutput struct {
	*pulumi.OutputState
}

func (CustomHostnameBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomHostnameBinding)(nil))
}

func (o CustomHostnameBindingPtrOutput) ToCustomHostnameBindingPtrOutput() CustomHostnameBindingPtrOutput {
	return o
}

func (o CustomHostnameBindingPtrOutput) ToCustomHostnameBindingPtrOutputWithContext(ctx context.Context) CustomHostnameBindingPtrOutput {
	return o
}

type CustomHostnameBindingArrayOutput struct{ *pulumi.OutputState }

func (CustomHostnameBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomHostnameBinding)(nil))
}

func (o CustomHostnameBindingArrayOutput) ToCustomHostnameBindingArrayOutput() CustomHostnameBindingArrayOutput {
	return o
}

func (o CustomHostnameBindingArrayOutput) ToCustomHostnameBindingArrayOutputWithContext(ctx context.Context) CustomHostnameBindingArrayOutput {
	return o
}

func (o CustomHostnameBindingArrayOutput) Index(i pulumi.IntInput) CustomHostnameBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomHostnameBinding {
		return vs[0].([]CustomHostnameBinding)[vs[1].(int)]
	}).(CustomHostnameBindingOutput)
}

type CustomHostnameBindingMapOutput struct{ *pulumi.OutputState }

func (CustomHostnameBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomHostnameBinding)(nil))
}

func (o CustomHostnameBindingMapOutput) ToCustomHostnameBindingMapOutput() CustomHostnameBindingMapOutput {
	return o
}

func (o CustomHostnameBindingMapOutput) ToCustomHostnameBindingMapOutputWithContext(ctx context.Context) CustomHostnameBindingMapOutput {
	return o
}

func (o CustomHostnameBindingMapOutput) MapIndex(k pulumi.StringInput) CustomHostnameBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomHostnameBinding {
		return vs[0].(map[string]CustomHostnameBinding)[vs[1].(string)]
	}).(CustomHostnameBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomHostnameBindingOutput{})
	pulumi.RegisterOutputType(CustomHostnameBindingPtrOutput{})
	pulumi.RegisterOutputType(CustomHostnameBindingArrayOutput{})
	pulumi.RegisterOutputType(CustomHostnameBindingMapOutput{})
}
