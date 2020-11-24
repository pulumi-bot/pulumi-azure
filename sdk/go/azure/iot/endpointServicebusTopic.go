// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IotHub ServiceBus Topic Endpoint
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
// 		exampleTopic, err := servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NamespaceName:     exampleNamespace.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopicAuthorizationRule, err := servicebus.NewTopicAuthorizationRule(ctx, "exampleTopicAuthorizationRule", &servicebus.TopicAuthorizationRuleArgs{
// 			NamespaceName:     exampleNamespace.Name,
// 			TopicName:         exampleTopic.Name,
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
// 		_, err = iot.NewEndpointServicebusTopic(ctx, "exampleEndpointServicebusTopic", &iot.EndpointServicebusTopicArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			ConnectionString:  exampleTopicAuthorizationRule.PrimaryConnectionString,
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
// IoTHub ServiceBus Topic Endpoint can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/endpointServicebusTopic:EndpointServicebusTopic servicebus_topic1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/servicebustopic_endpoint1
// ```
type EndpointServicebusTopic struct {
	pulumi.CustomResourceState

	// The connection string for the endpoint.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	IothubName       pulumi.StringOutput `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringOutput `pulumi:"name"`
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEndpointServicebusTopic registers a new resource with the given unique name, arguments, and options.
func NewEndpointServicebusTopic(ctx *pulumi.Context,
	name string, args *EndpointServicebusTopicArgs, opts ...pulumi.ResourceOption) (*EndpointServicebusTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.IothubName == nil {
		return nil, errors.New("invalid value for required argument 'IothubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EndpointServicebusTopic
	err := ctx.RegisterResource("azure:iot/endpointServicebusTopic:EndpointServicebusTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointServicebusTopic gets an existing EndpointServicebusTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointServicebusTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointServicebusTopicState, opts ...pulumi.ResourceOption) (*EndpointServicebusTopic, error) {
	var resource EndpointServicebusTopic
	err := ctx.ReadResource("azure:iot/endpointServicebusTopic:EndpointServicebusTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointServicebusTopic resources.
type endpointServicebusTopicState struct {
	// The connection string for the endpoint.
	ConnectionString *string `pulumi:"connectionString"`
	IothubName       *string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              *string `pulumi:"name"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EndpointServicebusTopicState struct {
	// The connection string for the endpoint.
	ConnectionString pulumi.StringPtrInput
	IothubName       pulumi.StringPtrInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
}

func (EndpointServicebusTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointServicebusTopicState)(nil)).Elem()
}

type endpointServicebusTopicArgs struct {
	// The connection string for the endpoint.
	ConnectionString string `pulumi:"connectionString"`
	IothubName       string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EndpointServicebusTopic resource.
type EndpointServicebusTopicArgs struct {
	// The connection string for the endpoint.
	ConnectionString pulumi.StringInput
	IothubName       pulumi.StringInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (EndpointServicebusTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointServicebusTopicArgs)(nil)).Elem()
}

type EndpointServicebusTopicInput interface {
	pulumi.Input

	ToEndpointServicebusTopicOutput() EndpointServicebusTopicOutput
	ToEndpointServicebusTopicOutputWithContext(ctx context.Context) EndpointServicebusTopicOutput
}

func (EndpointServicebusTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointServicebusTopic)(nil)).Elem()
}

func (i EndpointServicebusTopic) ToEndpointServicebusTopicOutput() EndpointServicebusTopicOutput {
	return i.ToEndpointServicebusTopicOutputWithContext(context.Background())
}

func (i EndpointServicebusTopic) ToEndpointServicebusTopicOutputWithContext(ctx context.Context) EndpointServicebusTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointServicebusTopicOutput)
}

type EndpointServicebusTopicOutput struct {
	*pulumi.OutputState
}

func (EndpointServicebusTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointServicebusTopicOutput)(nil)).Elem()
}

func (o EndpointServicebusTopicOutput) ToEndpointServicebusTopicOutput() EndpointServicebusTopicOutput {
	return o
}

func (o EndpointServicebusTopicOutput) ToEndpointServicebusTopicOutputWithContext(ctx context.Context) EndpointServicebusTopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EndpointServicebusTopicOutput{})
}
