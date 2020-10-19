// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual Hub within a Virtual WAN.
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
// 		exampleVirtualWan, err := network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualHub struct {
	pulumi.CustomResourceState

	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `route` blocks as defined below.
	Routes VirtualHubRouteArrayOutput `pulumi:"routes"`
	// A mapping of tags to assign to the Virtual Hub.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId pulumi.StringOutput `pulumi:"virtualWanId"`
}

// NewVirtualHub registers a new resource with the given unique name, arguments, and options.
func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil || args.AddressPrefix == nil {
		return nil, errors.New("missing required argument 'AddressPrefix'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualWanId == nil {
		return nil, errors.New("missing required argument 'VirtualWanId'")
	}
	if args == nil {
		args = &VirtualHubArgs{}
	}
	var resource VirtualHub
	err := ctx.RegisterResource("azure:network/virtualHub:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHub gets an existing VirtualHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azure:network/virtualHub:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHub resources.
type virtualHubState struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `route` blocks as defined below.
	Routes []VirtualHubRoute `pulumi:"routes"`
	// A mapping of tags to assign to the Virtual Hub.
	Tags map[string]string `pulumi:"tags"`
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId *string `pulumi:"virtualWanId"`
}

type VirtualHubState struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
	AddressPrefix pulumi.StringPtrInput
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `route` blocks as defined below.
	Routes VirtualHubRouteArrayInput
	// A mapping of tags to assign to the Virtual Hub.
	Tags pulumi.StringMapInput
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId pulumi.StringPtrInput
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
	AddressPrefix string `pulumi:"addressPrefix"`
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `route` blocks as defined below.
	Routes []VirtualHubRoute `pulumi:"routes"`
	// A mapping of tags to assign to the Virtual Hub.
	Tags map[string]string `pulumi:"tags"`
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId string `pulumi:"virtualWanId"`
}

// The set of arguments for constructing a VirtualHub resource.
type VirtualHubArgs struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
	AddressPrefix pulumi.StringInput
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `route` blocks as defined below.
	Routes VirtualHubRouteArrayInput
	// A mapping of tags to assign to the Virtual Hub.
	Tags pulumi.StringMapInput
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId pulumi.StringInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}

type VirtualHubInput interface {
	pulumi.Input

	ToVirtualHubOutput() VirtualHubOutput
	ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput
}

func (VirtualHub) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHub)(nil)).Elem()
}

func (i VirtualHub) ToVirtualHubOutput() VirtualHubOutput {
	return i.ToVirtualHubOutputWithContext(context.Background())
}

func (i VirtualHub) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubOutput)
}

type VirtualHubOutput struct {
	*pulumi.OutputState
}

func (VirtualHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubOutput)(nil)).Elem()
}

func (o VirtualHubOutput) ToVirtualHubOutput() VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualHubOutput{})
}
