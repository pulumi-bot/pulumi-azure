// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil)).Elem()
}

func (i Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct {
	*pulumi.OutputState
}

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceOutput)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
