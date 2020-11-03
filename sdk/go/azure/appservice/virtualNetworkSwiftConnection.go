// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).
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
// 		_, err = appservice.NewVirtualNetworkSwiftConnection(ctx, "exampleVirtualNetworkSwiftConnection", &appservice.VirtualNetworkSwiftConnectionArgs{
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
type VirtualNetworkSwiftConnection struct {
	pulumi.CustomResourceState

	// The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringOutput `pulumi:"appServiceId"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewVirtualNetworkSwiftConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkSwiftConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkSwiftConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AppServiceId == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource VirtualNetworkSwiftConnection
	err := ctx.RegisterResource("azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkSwiftConnection gets an existing VirtualNetworkSwiftConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkSwiftConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkSwiftConnection, error) {
	var resource VirtualNetworkSwiftConnection
	err := ctx.ReadResource("azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkSwiftConnection resources.
type virtualNetworkSwiftConnectionState struct {
	// The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId *string `pulumi:"appServiceId"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId *string `pulumi:"subnetId"`
}

type VirtualNetworkSwiftConnectionState struct {
	// The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringPtrInput
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringPtrInput
}

func (VirtualNetworkSwiftConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkSwiftConnectionState)(nil)).Elem()
}

type virtualNetworkSwiftConnectionArgs struct {
	// The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId string `pulumi:"appServiceId"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VirtualNetworkSwiftConnection resource.
type VirtualNetworkSwiftConnectionArgs struct {
	// The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringInput
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringInput
}

func (VirtualNetworkSwiftConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkSwiftConnectionArgs)(nil)).Elem()
}
