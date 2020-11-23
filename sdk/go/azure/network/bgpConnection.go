// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Bgp Connection for a Virtual Hub.
//
// ## Import
//
// Virtual Hub Bgp Connections can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/bgpConnection:BgpConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/connection1
// ```
type BgpConnection struct {
	pulumi.CustomResourceState

	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn pulumi.IntOutput `pulumi:"peerAsn"`
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
}

// NewBgpConnection registers a new resource with the given unique name, arguments, and options.
func NewBgpConnection(ctx *pulumi.Context,
	name string, args *BgpConnectionArgs, opts ...pulumi.ResourceOption) (*BgpConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerAsn == nil {
		return nil, errors.New("invalid value for required argument 'PeerAsn'")
	}
	if args.PeerIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerIp'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource BgpConnection
	err := ctx.RegisterResource("azure:network/bgpConnection:BgpConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpConnection gets an existing BgpConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpConnectionState, opts ...pulumi.ResourceOption) (*BgpConnection, error) {
	var resource BgpConnection
	err := ctx.ReadResource("azure:network/bgpConnection:BgpConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpConnection resources.
type bgpConnectionState struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn *int `pulumi:"peerAsn"`
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp *string `pulumi:"peerIp"`
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type BgpConnectionState struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn pulumi.IntPtrInput
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp pulumi.StringPtrInput
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (BgpConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpConnectionState)(nil)).Elem()
}

type bgpConnectionArgs struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn int `pulumi:"peerAsn"`
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp string `pulumi:"peerIp"`
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a BgpConnection resource.
type BgpConnectionArgs struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn pulumi.IntInput
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp pulumi.StringInput
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
}

func (BgpConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpConnectionArgs)(nil)).Elem()
}

type BgpConnectionInput interface {
	pulumi.Input

	ToBgpConnectionOutput() BgpConnectionOutput
	ToBgpConnectionOutputWithContext(ctx context.Context) BgpConnectionOutput
}

func (BgpConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpConnection)(nil)).Elem()
}

func (i BgpConnection) ToBgpConnectionOutput() BgpConnectionOutput {
	return i.ToBgpConnectionOutputWithContext(context.Background())
}

func (i BgpConnection) ToBgpConnectionOutputWithContext(ctx context.Context) BgpConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpConnectionOutput)
}

type BgpConnectionOutput struct {
	*pulumi.OutputState
}

func (BgpConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpConnectionOutput)(nil)).Elem()
}

func (o BgpConnectionOutput) ToBgpConnectionOutput() BgpConnectionOutput {
	return o
}

func (o BgpConnectionOutput) ToBgpConnectionOutputWithContext(ctx context.Context) BgpConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BgpConnectionOutput{})
}