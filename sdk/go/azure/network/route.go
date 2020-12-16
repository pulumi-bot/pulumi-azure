// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Route within a Route Table.
//
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource, and allows for Routes to be defined in-line within the Route Table resource.
// At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.
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
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleRouteTable, err := network.NewRouteTable(ctx, "exampleRouteTable", &network.RouteTableArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewRoute(ctx, "exampleRoute", &network.RouteArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			RouteTableName:    exampleRouteTable.Name,
// 			AddressPrefix:     pulumi.String("10.1.0.0/16"),
// 			NextHopType:       pulumi.String("vnetlocal"),
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
// Routes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/route:Route exampleRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1
// ```
type Route struct {
	pulumi.CustomResourceState

	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// The name of the route. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress pulumi.StringPtrOutput `pulumi:"nextHopInIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType pulumi.StringOutput `pulumi:"nextHopType"`
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName pulumi.StringOutput `pulumi:"routeTableName"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'AddressPrefix'")
	}
	if args.NextHopType == nil {
		return nil, errors.New("invalid value for required argument 'NextHopType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RouteTableName == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableName'")
	}
	var resource Route
	err := ctx.RegisterResource("azure:network/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azure:network/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The name of the route. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress *string `pulumi:"nextHopInIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType *string `pulumi:"nextHopType"`
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName *string `pulumi:"routeTableName"`
}

type RouteState struct {
	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix pulumi.StringPtrInput
	// The name of the route. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress pulumi.StringPtrInput
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType pulumi.StringPtrInput
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix string `pulumi:"addressPrefix"`
	// The name of the route. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress *string `pulumi:"nextHopInIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType string `pulumi:"nextHopType"`
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName string `pulumi:"routeTableName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix pulumi.StringInput
	// The name of the route. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress pulumi.StringPtrInput
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType pulumi.StringInput
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
