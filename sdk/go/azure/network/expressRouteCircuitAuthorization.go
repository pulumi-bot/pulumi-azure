// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an ExpressRoute Circuit Authorization.
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
// 		exampleExpressRouteCircuit, err := network.NewExpressRouteCircuit(ctx, "exampleExpressRouteCircuit", &network.ExpressRouteCircuitArgs{
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			Location:            exampleResourceGroup.Location,
// 			ServiceProviderName: pulumi.String("Equinix"),
// 			PeeringLocation:     pulumi.String("Silicon Valley"),
// 			BandwidthInMbps:     pulumi.Int(50),
// 			Sku: &network.ExpressRouteCircuitSkuArgs{
// 				Tier:   pulumi.String("Standard"),
// 				Family: pulumi.String("MeteredData"),
// 			},
// 			AllowClassicOperations: pulumi.Bool(false),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewExpressRouteCircuitAuthorization(ctx, "exampleExpressRouteCircuitAuthorization", &network.ExpressRouteCircuitAuthorizationArgs{
// 			ExpressRouteCircuitName: exampleExpressRouteCircuit.Name,
// 			ResourceGroupName:       exampleResourceGroup.Name,
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
// ExpressRoute Circuit Authorizations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/expressRouteCircuitAuthorization:ExpressRouteCircuitAuthorization auth1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute/authorizations/auth1
// ```
type ExpressRouteCircuitAuthorization struct {
	pulumi.CustomResourceState

	// The Authorization Key.
	AuthorizationKey pulumi.StringOutput `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus pulumi.StringOutput `pulumi:"authorizationUseStatus"`
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName pulumi.StringOutput `pulumi:"expressRouteCircuitName"`
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewExpressRouteCircuitAuthorization registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpressRouteCircuitName == nil {
		return nil, errors.New("invalid value for required argument 'ExpressRouteCircuitName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ExpressRouteCircuitAuthorization
	err := ctx.RegisterResource("azure:network/expressRouteCircuitAuthorization:ExpressRouteCircuitAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitAuthorization gets an existing ExpressRouteCircuitAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	var resource ExpressRouteCircuitAuthorization
	err := ctx.ReadResource("azure:network/expressRouteCircuitAuthorization:ExpressRouteCircuitAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitAuthorization resources.
type expressRouteCircuitAuthorizationState struct {
	// The Authorization Key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName *string `pulumi:"expressRouteCircuitName"`
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ExpressRouteCircuitAuthorizationState struct {
	// The Authorization Key.
	AuthorizationKey pulumi.StringPtrInput
	// The authorization use status.
	AuthorizationUseStatus pulumi.StringPtrInput
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName pulumi.StringPtrInput
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ExpressRouteCircuitAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationState)(nil)).Elem()
}

type expressRouteCircuitAuthorizationArgs struct {
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName string `pulumi:"expressRouteCircuitName"`
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitAuthorization resource.
type ExpressRouteCircuitAuthorizationArgs struct {
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName pulumi.StringInput
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationArgs)(nil)).Elem()
}

type ExpressRouteCircuitAuthorizationInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationOutput() ExpressRouteCircuitAuthorizationOutput
	ToExpressRouteCircuitAuthorizationOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationOutput
}

func (*ExpressRouteCircuitAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorization)(nil))
}

func (i *ExpressRouteCircuitAuthorization) ToExpressRouteCircuitAuthorizationOutput() ExpressRouteCircuitAuthorizationOutput {
	return i.ToExpressRouteCircuitAuthorizationOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitAuthorization) ToExpressRouteCircuitAuthorizationOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationOutput)
}

type ExpressRouteCircuitAuthorizationOutput struct {
	*pulumi.OutputState
}

func (ExpressRouteCircuitAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorization)(nil))
}

func (o ExpressRouteCircuitAuthorizationOutput) ToExpressRouteCircuitAuthorizationOutput() ExpressRouteCircuitAuthorizationOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationOutput) ToExpressRouteCircuitAuthorizationOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationOutput{})
}
