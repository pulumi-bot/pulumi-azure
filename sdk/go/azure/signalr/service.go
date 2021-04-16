// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure SignalR service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/signalr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = signalr.NewService(ctx, "exampleService", &signalr.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &signalr.ServiceSkuArgs{
// 				Name:     pulumi.String("Free_F1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 			Cors: signalr.ServiceCorArray{
// 				&signalr.ServiceCorArgs{
// 					AllowedOrigins: pulumi.StringArray{
// 						pulumi.String("http://www.example.com"),
// 					},
// 				},
// 			},
// 			Features: signalr.ServiceFeatureArray{
// 				&signalr.ServiceFeatureArgs{
// 					Flag:  pulumi.String("ServiceMode"),
// 					Value: pulumi.String("Default"),
// 				},
// 			},
// 			UpstreamEndpoints: signalr.ServiceUpstreamEndpointArray{
// 				&signalr.ServiceUpstreamEndpointArgs{
// 					CategoryPatterns: pulumi.StringArray{
// 						pulumi.String("connections"),
// 						pulumi.String("messages"),
// 					},
// 					EventPatterns: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					HubPatterns: pulumi.StringArray{
// 						pulumi.String("hub1"),
// 					},
// 					UrlTemplate: pulumi.String("http://foo.com"),
// 				},
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
// SignalR services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:signalr/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/SignalR/tfex-signalr
// ```
type Service struct {
	pulumi.CustomResourceState

	// A `cors` block as documented below.
	Cors ServiceCorArrayOutput `pulumi:"cors"`
	// A `features` block as documented below.
	Features ServiceFeatureArrayOutput `pulumi:"features"`
	// The FQDN of the SignalR service.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The publicly accessible IP of the SignalR service.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the SignalR service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary access key for the SignalR service.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The primary connection string for the SignalR service.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for browser/client use.
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary access key for the SignalR service.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The secondary connection string for the SignalR service.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for customer server side use.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// A `sku` block as documented below.
	Sku ServiceSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoints ServiceUpstreamEndpointArrayOutput `pulumi:"upstreamEndpoints"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Service
	err := ctx.RegisterResource("azure:signalr/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:signalr/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// A `cors` block as documented below.
	Cors []ServiceCor `pulumi:"cors"`
	// A `features` block as documented below.
	Features []ServiceFeature `pulumi:"features"`
	// The FQDN of the SignalR service.
	Hostname *string `pulumi:"hostname"`
	// The publicly accessible IP of the SignalR service.
	IpAddress *string `pulumi:"ipAddress"`
	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the SignalR service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary access key for the SignalR service.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The primary connection string for the SignalR service.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for browser/client use.
	PublicPort *int `pulumi:"publicPort"`
	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary access key for the SignalR service.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The secondary connection string for the SignalR service.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for customer server side use.
	ServerPort *int `pulumi:"serverPort"`
	// A `sku` block as documented below.
	Sku *ServiceSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoints []ServiceUpstreamEndpoint `pulumi:"upstreamEndpoints"`
}

type ServiceState struct {
	// A `cors` block as documented below.
	Cors ServiceCorArrayInput
	// A `features` block as documented below.
	Features ServiceFeatureArrayInput
	// The FQDN of the SignalR service.
	Hostname pulumi.StringPtrInput
	// The publicly accessible IP of the SignalR service.
	IpAddress pulumi.StringPtrInput
	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the SignalR service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary access key for the SignalR service.
	PrimaryAccessKey pulumi.StringPtrInput
	// The primary connection string for the SignalR service.
	PrimaryConnectionString pulumi.StringPtrInput
	// The publicly accessible port of the SignalR service which is designed for browser/client use.
	PublicPort pulumi.IntPtrInput
	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary access key for the SignalR service.
	SecondaryAccessKey pulumi.StringPtrInput
	// The secondary connection string for the SignalR service.
	SecondaryConnectionString pulumi.StringPtrInput
	// The publicly accessible port of the SignalR service which is designed for customer server side use.
	ServerPort pulumi.IntPtrInput
	// A `sku` block as documented below.
	Sku ServiceSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoints ServiceUpstreamEndpointArrayInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// A `cors` block as documented below.
	Cors []ServiceCor `pulumi:"cors"`
	// A `features` block as documented below.
	Features []ServiceFeature `pulumi:"features"`
	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the SignalR service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as documented below.
	Sku ServiceSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoints []ServiceUpstreamEndpoint `pulumi:"upstreamEndpoints"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// A `cors` block as documented below.
	Cors ServiceCorArrayInput
	// A `features` block as documented below.
	Features ServiceFeatureArrayInput
	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the SignalR service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `sku` block as documented below.
	Sku ServiceSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoints ServiceUpstreamEndpointArrayInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

func (i *Service) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *Service) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

type ServicePtrInput interface {
	pulumi.Input

	ToServicePtrOutput() ServicePtrOutput
	ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput
}

type servicePtrType ServiceArgs

func (*servicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (i *servicePtrType) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *servicePtrType) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Service)(nil))
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Service)(nil))
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct {
	*pulumi.OutputState
}

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) ToServicePtrOutput() ServicePtrOutput {
	return o.ToServicePtrOutputWithContext(context.Background())
}

func (o ServiceOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o.ApplyT(func(v Service) *Service {
		return &v
	}).(ServicePtrOutput)
}

type ServicePtrOutput struct {
	*pulumi.OutputState
}

func (ServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (o ServicePtrOutput) ToServicePtrOutput() ServicePtrOutput {
	return o
}

func (o ServicePtrOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Service)(nil))
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Service {
		return vs[0].([]Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Service)(nil))
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Service {
		return vs[0].(map[string]Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServicePtrOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
