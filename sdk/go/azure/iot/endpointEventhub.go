// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IotHub EventHub Endpoint
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
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
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAuthorizationRule, err := eventhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &eventhub.AuthorizationRuleArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
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
// 		_, err = iot.NewEndpointEventhub(ctx, "exampleEndpointEventhub", &iot.EndpointEventhubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			ConnectionString:  exampleAuthorizationRule.PrimaryConnectionString,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EndpointEventhub struct {
	pulumi.CustomResourceState

	// The connection string for the endpoint.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	IothubName       pulumi.StringOutput `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringOutput `pulumi:"name"`
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEndpointEventhub registers a new resource with the given unique name, arguments, and options.
func NewEndpointEventhub(ctx *pulumi.Context,
	name string, args *EndpointEventhubArgs, opts ...pulumi.ResourceOption) (*EndpointEventhub, error) {
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
	var resource EndpointEventhub
	err := ctx.RegisterResource("azure:iot/endpointEventhub:EndpointEventhub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointEventhub gets an existing EndpointEventhub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointEventhub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointEventhubState, opts ...pulumi.ResourceOption) (*EndpointEventhub, error) {
	var resource EndpointEventhub
	err := ctx.ReadResource("azure:iot/endpointEventhub:EndpointEventhub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointEventhub resources.
type endpointEventhubState struct {
	// The connection string for the endpoint.
	ConnectionString *string `pulumi:"connectionString"`
	IothubName       *string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              *string `pulumi:"name"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EndpointEventhubState struct {
	// The connection string for the endpoint.
	ConnectionString pulumi.StringPtrInput
	IothubName       pulumi.StringPtrInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
}

func (EndpointEventhubState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointEventhubState)(nil)).Elem()
}

type endpointEventhubArgs struct {
	// The connection string for the endpoint.
	ConnectionString string `pulumi:"connectionString"`
	IothubName       string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EndpointEventhub resource.
type EndpointEventhubArgs struct {
	// The connection string for the endpoint.
	ConnectionString pulumi.StringInput
	IothubName       pulumi.StringInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (EndpointEventhubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointEventhubArgs)(nil)).Elem()
}
