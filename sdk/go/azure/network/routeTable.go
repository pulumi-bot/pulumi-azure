// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Route Table
//
// > **NOTE on Route Tables and Routes:** There is both a standalone `route` resource, and allows for Routes to be defined in-line within the `routeTable` resource.
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
// 		_, err = network.NewRouteTable(ctx, "exampleRouteTable", &network.RouteTableArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			DisableBgpRoutePropagation: pulumi.Bool(false),
// 			Routes: network.RouteTableRouteArray{
// 				&network.RouteTableRouteArgs{
// 					Name:          pulumi.String("route1"),
// 					AddressPrefix: pulumi.String("10.1.0.0/16"),
// 					NextHopType:   pulumi.String("vnetlocal"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// Route Tables can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/routeTable:RouteTable example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1
// ```
type RouteTable struct {
	pulumi.CustomResourceState

	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation pulumi.BoolPtrOutput `pulumi:"disableBgpRoutePropagation"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the route.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A list of objects representing routes. Each object accepts the arguments documented below.
	Routes RouteTableRouteArrayOutput `pulumi:"routes"`
	// The collection of Subnets associated with this route table.
	Subnets pulumi.StringArrayOutput `pulumi:"subnets"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource RouteTable
	err := ctx.RegisterResource("azure:network/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("azure:network/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `pulumi:"disableBgpRoutePropagation"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the route.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A list of objects representing routes. Each object accepts the arguments documented below.
	Routes []RouteTableRoute `pulumi:"routes"`
	// The collection of Subnets associated with this route table.
	Subnets []string `pulumi:"subnets"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type RouteTableState struct {
	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the route.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A list of objects representing routes. Each object accepts the arguments documented below.
	Routes RouteTableRouteArrayInput
	// The collection of Subnets associated with this route table.
	Subnets pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `pulumi:"disableBgpRoutePropagation"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the route.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A list of objects representing routes. Each object accepts the arguments documented below.
	Routes []RouteTableRoute `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the route.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A list of objects representing routes. Each object accepts the arguments documented below.
	Routes RouteTableRouteArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTable)(nil)).Elem()
}

func (i RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

type RouteTableOutput struct {
	*pulumi.OutputState
}

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableOutput)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteTableOutput{})
}
