// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IotHub ServiceBus Queue Endpoint
//
// > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicebus"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleQueue, err := servicebus.NewQueue(ctx, "exampleQueue", &servicebus.QueueArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleQueueAuthorizationRule, err := servicebus.NewQueueAuthorizationRule(ctx, "exampleQueueAuthorizationRule", &servicebus.QueueAuthorizationRuleArgs{
// 			NamespaceName:     exampleNamespace.Name,
// 			QueueName:         exampleQueue.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Listen:            pulumi.Bool(false),
// 			Send:              pulumi.Bool(true),
// 			Manage:            pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("B1"),
// 				Tier:     pulumi.String("Basic"),
// 				Capacity: pulumi.Int(1),
// 			},
// 			Tags: pulumi.StringMap{
// 				"purpose": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewEndpointServicebusQueue(ctx, "exampleEndpointServicebusQueue", &iot.EndpointServicebusQueueArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			ConnectionString:  exampleQueueAuthorizationRule.PrimaryConnectionString,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EndpointServicebusQueue struct {
	pulumi.CustomResourceState

	// The connection string for the endpoint.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	IothubName       pulumi.StringOutput `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringOutput `pulumi:"name"`
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEndpointServicebusQueue registers a new resource with the given unique name, arguments, and options.
func NewEndpointServicebusQueue(ctx *pulumi.Context,
	name string, args *EndpointServicebusQueueArgs, opts ...pulumi.ResourceOption) (*EndpointServicebusQueue, error) {
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.IothubName == nil {
		return nil, errors.New("missing required argument 'IothubName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EndpointServicebusQueueArgs{}
	}
	var resource EndpointServicebusQueue
	err := ctx.RegisterResource("azure:iot/endpointServicebusQueue:EndpointServicebusQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointServicebusQueue gets an existing EndpointServicebusQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointServicebusQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointServicebusQueueState, opts ...pulumi.ResourceOption) (*EndpointServicebusQueue, error) {
	var resource EndpointServicebusQueue
	err := ctx.ReadResource("azure:iot/endpointServicebusQueue:EndpointServicebusQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointServicebusQueue resources.
type endpointServicebusQueueState struct {
	// The connection string for the endpoint.
	ConnectionString *string `pulumi:"connectionString"`
	IothubName       *string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              *string `pulumi:"name"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EndpointServicebusQueueState struct {
	// The connection string for the endpoint.
	ConnectionString pulumi.StringPtrInput
	IothubName       pulumi.StringPtrInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
}

func (EndpointServicebusQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointServicebusQueueState)(nil)).Elem()
}

type endpointServicebusQueueArgs struct {
	// The connection string for the endpoint.
	ConnectionString string `pulumi:"connectionString"`
	IothubName       string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EndpointServicebusQueue resource.
type EndpointServicebusQueueArgs struct {
	// The connection string for the endpoint.
	ConnectionString pulumi.StringInput
	IothubName       pulumi.StringInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (EndpointServicebusQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointServicebusQueueArgs)(nil)).Elem()
}
