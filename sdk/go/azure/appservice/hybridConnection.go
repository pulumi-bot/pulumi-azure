// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service Hybrid Connection for an existing App Service, Relay and Service Bus.
//
// ## Example Usage
//
// This example provisions an App Service, a Relay Hybrid Connection, and a Service Bus using their outputs to create the App Service Hybrid Connection.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/relay"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleNamespace, err := relay.NewNamespace(ctx, "exampleNamespace", &relay.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleHybridConnection, err := relay.NewHybridConnection(ctx, "exampleHybridConnection", &relay.HybridConnectionArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			RelayNamespaceName: exampleNamespace.Name,
// 			UserMetadata:       pulumi.String("examplemetadata"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewHybridConnection(ctx, "exampleAppservice_hybridConnectionHybridConnection", &appservice.HybridConnectionArgs{
// 			AppServiceName:    exampleAppService.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			RelayId:           exampleHybridConnection.ID(),
// 			Hostname:          pulumi.String("testhostname.example"),
// 			Port:              pulumi.Int(8080),
// 			SendKeyName:       pulumi.String("exampleSharedAccessKey"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type HybridConnection struct {
	pulumi.CustomResourceState

	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The hostname of the endpoint.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The name of the Relay Namespace.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The port of the endpoint.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the Service Bus Relay. Changing this forces a new resource to be created.
	RelayId   pulumi.StringOutput `pulumi:"relayId"`
	RelayName pulumi.StringOutput `pulumi:"relayName"`
	// The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
	SendKeyName pulumi.StringPtrOutput `pulumi:"sendKeyName"`
	// The value of the Service Bus Primary Access key.
	SendKeyValue pulumi.StringOutput `pulumi:"sendKeyValue"`
	// The name of the Service Bus namespace.
	ServiceBusNamespace pulumi.StringOutput `pulumi:"serviceBusNamespace"`
	// The suffix for the service bus endpoint.
	ServiceBusSuffix pulumi.StringOutput `pulumi:"serviceBusSuffix"`
}

// NewHybridConnection registers a new resource with the given unique name, arguments, and options.
func NewHybridConnection(ctx *pulumi.Context,
	name string, args *HybridConnectionArgs, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AppServiceName == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceName'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.RelayId == nil {
		return nil, errors.New("invalid value for required argument 'RelayId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource HybridConnection
	err := ctx.RegisterResource("azure:appservice/hybridConnection:HybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridConnection gets an existing HybridConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridConnectionState, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	var resource HybridConnection
	err := ctx.ReadResource("azure:appservice/hybridConnection:HybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridConnection resources.
type hybridConnectionState struct {
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// The hostname of the endpoint.
	Hostname *string `pulumi:"hostname"`
	// The name of the Relay Namespace.
	NamespaceName *string `pulumi:"namespaceName"`
	// The port of the endpoint.
	Port *int `pulumi:"port"`
	// The ID of the Service Bus Relay. Changing this forces a new resource to be created.
	RelayId   *string `pulumi:"relayId"`
	RelayName *string `pulumi:"relayName"`
	// The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
	SendKeyName *string `pulumi:"sendKeyName"`
	// The value of the Service Bus Primary Access key.
	SendKeyValue *string `pulumi:"sendKeyValue"`
	// The name of the Service Bus namespace.
	ServiceBusNamespace *string `pulumi:"serviceBusNamespace"`
	// The suffix for the service bus endpoint.
	ServiceBusSuffix *string `pulumi:"serviceBusSuffix"`
}

type HybridConnectionState struct {
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	AppServiceName pulumi.StringPtrInput
	// The hostname of the endpoint.
	Hostname pulumi.StringPtrInput
	// The name of the Relay Namespace.
	NamespaceName pulumi.StringPtrInput
	// The port of the endpoint.
	Port pulumi.IntPtrInput
	// The ID of the Service Bus Relay. Changing this forces a new resource to be created.
	RelayId   pulumi.StringPtrInput
	RelayName pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
	SendKeyName pulumi.StringPtrInput
	// The value of the Service Bus Primary Access key.
	SendKeyValue pulumi.StringPtrInput
	// The name of the Service Bus namespace.
	ServiceBusNamespace pulumi.StringPtrInput
	// The suffix for the service bus endpoint.
	ServiceBusSuffix pulumi.StringPtrInput
}

func (HybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionState)(nil)).Elem()
}

type hybridConnectionArgs struct {
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// The hostname of the endpoint.
	Hostname string `pulumi:"hostname"`
	// The port of the endpoint.
	Port int `pulumi:"port"`
	// The ID of the Service Bus Relay. Changing this forces a new resource to be created.
	RelayId string `pulumi:"relayId"`
	// The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
	SendKeyName *string `pulumi:"sendKeyName"`
}

// The set of arguments for constructing a HybridConnection resource.
type HybridConnectionArgs struct {
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	AppServiceName pulumi.StringInput
	// The hostname of the endpoint.
	Hostname pulumi.StringInput
	// The port of the endpoint.
	Port pulumi.IntInput
	// The ID of the Service Bus Relay. Changing this forces a new resource to be created.
	RelayId pulumi.StringInput
	// The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
	SendKeyName pulumi.StringPtrInput
}

func (HybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionArgs)(nil)).Elem()
}
