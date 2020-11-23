// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service Slot's Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("uksouth"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("example-delegation"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.Web/serverFarms"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/action"),
// 						},
// 					},
// 				},
// 			},
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
// 		_, err = appservice.NewSlot(ctx, "example_staging", &appservice.SlotArgs{
// 			AppServiceName:    exampleAppService.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewSlotVirtualNetworkSwiftConnection(ctx, "exampleSlotVirtualNetworkSwiftConnection", &appservice.SlotVirtualNetworkSwiftConnectionArgs{
// 			SlotName:     example_staging.Name,
// 			AppServiceId: exampleAppService.ID(),
// 			SubnetId:     exampleSubnet.ID(),
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
// App Service Slot Virtual Network Associations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection myassociation /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/stageing/networkconfig/virtualNetwork
// ```
type SlotVirtualNetworkSwiftConnection struct {
	pulumi.CustomResourceState

	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringOutput `pulumi:"appServiceId"`
	// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
	SlotName pulumi.StringOutput `pulumi:"slotName"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSlotVirtualNetworkSwiftConnection registers a new resource with the given unique name, arguments, and options.
func NewSlotVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, args *SlotVirtualNetworkSwiftConnectionArgs, opts ...pulumi.ResourceOption) (*SlotVirtualNetworkSwiftConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceId == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceId'")
	}
	if args.SlotName == nil {
		return nil, errors.New("invalid value for required argument 'SlotName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource SlotVirtualNetworkSwiftConnection
	err := ctx.RegisterResource("azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlotVirtualNetworkSwiftConnection gets an existing SlotVirtualNetworkSwiftConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlotVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SlotVirtualNetworkSwiftConnectionState, opts ...pulumi.ResourceOption) (*SlotVirtualNetworkSwiftConnection, error) {
	var resource SlotVirtualNetworkSwiftConnection
	err := ctx.ReadResource("azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SlotVirtualNetworkSwiftConnection resources.
type slotVirtualNetworkSwiftConnectionState struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId *string `pulumi:"appServiceId"`
	// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
	SlotName *string `pulumi:"slotName"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId *string `pulumi:"subnetId"`
}

type SlotVirtualNetworkSwiftConnectionState struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringPtrInput
	// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
	SlotName pulumi.StringPtrInput
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringPtrInput
}

func (SlotVirtualNetworkSwiftConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*slotVirtualNetworkSwiftConnectionState)(nil)).Elem()
}

type slotVirtualNetworkSwiftConnectionArgs struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId string `pulumi:"appServiceId"`
	// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
	SlotName string `pulumi:"slotName"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SlotVirtualNetworkSwiftConnection resource.
type SlotVirtualNetworkSwiftConnectionArgs struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringInput
	// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
	SlotName pulumi.StringInput
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringInput
}

func (SlotVirtualNetworkSwiftConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*slotVirtualNetworkSwiftConnectionArgs)(nil)).Elem()
}

type SlotVirtualNetworkSwiftConnectionInput interface {
	pulumi.Input

	ToSlotVirtualNetworkSwiftConnectionOutput() SlotVirtualNetworkSwiftConnectionOutput
	ToSlotVirtualNetworkSwiftConnectionOutputWithContext(ctx context.Context) SlotVirtualNetworkSwiftConnectionOutput
}

func (SlotVirtualNetworkSwiftConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotVirtualNetworkSwiftConnection)(nil)).Elem()
}

func (i SlotVirtualNetworkSwiftConnection) ToSlotVirtualNetworkSwiftConnectionOutput() SlotVirtualNetworkSwiftConnectionOutput {
	return i.ToSlotVirtualNetworkSwiftConnectionOutputWithContext(context.Background())
}

func (i SlotVirtualNetworkSwiftConnection) ToSlotVirtualNetworkSwiftConnectionOutputWithContext(ctx context.Context) SlotVirtualNetworkSwiftConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotVirtualNetworkSwiftConnectionOutput)
}

type SlotVirtualNetworkSwiftConnectionOutput struct {
	*pulumi.OutputState
}

func (SlotVirtualNetworkSwiftConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotVirtualNetworkSwiftConnectionOutput)(nil)).Elem()
}

func (o SlotVirtualNetworkSwiftConnectionOutput) ToSlotVirtualNetworkSwiftConnectionOutput() SlotVirtualNetworkSwiftConnectionOutput {
	return o
}

func (o SlotVirtualNetworkSwiftConnectionOutput) ToSlotVirtualNetworkSwiftConnectionOutputWithContext(ctx context.Context) SlotVirtualNetworkSwiftConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SlotVirtualNetworkSwiftConnectionOutput{})
}
