// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Connection for a Virtual Hub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("172.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualWan(ctx, "test", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      pulumi.Any(azurerm_virtual_wan.Example.Id),
// 			AddressPrefix:     pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualHubConnection(ctx, "exampleVirtualHubConnection", &network.VirtualHubConnectionArgs{
// 			VirtualHubId:           exampleVirtualHub.ID(),
// 			RemoteVirtualNetworkId: exampleVirtualNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualHubConnection struct {
	pulumi.CustomResourceState

	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed pulumi.BoolPtrOutput `pulumi:"hubToVitualNetworkTrafficAllowed"`
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled pulumi.BoolPtrOutput `pulumi:"internetSecurityEnabled"`
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId pulumi.StringOutput `pulumi:"remoteVirtualNetworkId"`
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed pulumi.BoolPtrOutput `pulumi:"vitualNetworkToHubGatewaysTrafficAllowed"`
}

// NewVirtualHubConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubConnection(ctx *pulumi.Context,
	name string, args *VirtualHubConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualHubConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.RemoteVirtualNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'RemoteVirtualNetworkId'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource VirtualHubConnection
	err := ctx.RegisterResource("azure:network/virtualHubConnection:VirtualHubConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubConnection gets an existing VirtualHubConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubConnectionState, opts ...pulumi.ResourceOption) (*VirtualHubConnection, error) {
	var resource VirtualHubConnection
	err := ctx.ReadResource("azure:network/virtualHubConnection:VirtualHubConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubConnection resources.
type virtualHubConnectionState struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed *bool `pulumi:"hubToVitualNetworkTrafficAllowed"`
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled *bool `pulumi:"internetSecurityEnabled"`
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId *string `pulumi:"remoteVirtualNetworkId"`
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed *bool `pulumi:"vitualNetworkToHubGatewaysTrafficAllowed"`
}

type VirtualHubConnectionState struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed pulumi.BoolPtrInput
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled pulumi.BoolPtrInput
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId pulumi.StringPtrInput
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed pulumi.BoolPtrInput
}

func (VirtualHubConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubConnectionState)(nil)).Elem()
}

type virtualHubConnectionArgs struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed *bool `pulumi:"hubToVitualNetworkTrafficAllowed"`
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled *bool `pulumi:"internetSecurityEnabled"`
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId string `pulumi:"remoteVirtualNetworkId"`
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed *bool `pulumi:"vitualNetworkToHubGatewaysTrafficAllowed"`
}

// The set of arguments for constructing a VirtualHubConnection resource.
type VirtualHubConnectionArgs struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed pulumi.BoolPtrInput
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled pulumi.BoolPtrInput
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId pulumi.StringInput
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed pulumi.BoolPtrInput
}

func (VirtualHubConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubConnectionArgs)(nil)).Elem()
}
