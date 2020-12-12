// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Security Partner Provider which could be associated to virtual hub.
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
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.2.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpnGateway, err := network.NewVpnGateway(ctx, "exampleVpnGateway", &network.VpnGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			VirtualHubId:      exampleVirtualHub.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSecurityPartnerProvider(ctx, "exampleSecurityPartnerProvider", &network.SecurityPartnerProviderArgs{
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			Location:             exampleResourceGroup.Location,
// 			VirtualHubId:         exampleVirtualHub.ID(),
// 			SecurityProviderName: pulumi.String("IBoss"),
// 			Tags: pulumi.StringMap{
// 				"ENV": pulumi.String("Prod"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleVpnGateway,
// 		}))
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
// Security Partner Providers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/securityPartnerProvider:SecurityPartnerProvider example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/securityPartnerProviders/securityPartnerProvider1
// ```
type SecurityPartnerProvider struct {
	pulumi.CustomResourceState

	// The Azure Region where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Security Partner Provider. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The security provider name. Possible values are `ZScaler`, `IBoss` and `Checkpoint` is allowed. Changing this forces a new resource to be created.
	SecurityProviderName pulumi.StringOutput `pulumi:"securityProviderName"`
	// A mapping of tags which should be assigned to the Security Partner Provider.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Hub within which this Security Partner Provider should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrOutput `pulumi:"virtualHubId"`
}

// NewSecurityPartnerProvider registers a new resource with the given unique name, arguments, and options.
func NewSecurityPartnerProvider(ctx *pulumi.Context,
	name string, args *SecurityPartnerProviderArgs, opts ...pulumi.ResourceOption) (*SecurityPartnerProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecurityProviderName == nil {
		return nil, errors.New("invalid value for required argument 'SecurityProviderName'")
	}
	var resource SecurityPartnerProvider
	err := ctx.RegisterResource("azure:network/securityPartnerProvider:SecurityPartnerProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPartnerProvider gets an existing SecurityPartnerProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPartnerProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPartnerProviderState, opts ...pulumi.ResourceOption) (*SecurityPartnerProvider, error) {
	var resource SecurityPartnerProvider
	err := ctx.ReadResource("azure:network/securityPartnerProvider:SecurityPartnerProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPartnerProvider resources.
type securityPartnerProviderState struct {
	// The Azure Region where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Security Partner Provider. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The security provider name. Possible values are `ZScaler`, `IBoss` and `Checkpoint` is allowed. Changing this forces a new resource to be created.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// A mapping of tags which should be assigned to the Security Partner Provider.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Hub within which this Security Partner Provider should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type SecurityPartnerProviderState struct {
	// The Azure Region where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Security Partner Provider. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The security provider name. Possible values are `ZScaler`, `IBoss` and `Checkpoint` is allowed. Changing this forces a new resource to be created.
	SecurityProviderName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Security Partner Provider.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Hub within which this Security Partner Provider should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (SecurityPartnerProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPartnerProviderState)(nil)).Elem()
}

type securityPartnerProviderArgs struct {
	// The Azure Region where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Security Partner Provider. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security provider name. Possible values are `ZScaler`, `IBoss` and `Checkpoint` is allowed. Changing this forces a new resource to be created.
	SecurityProviderName string `pulumi:"securityProviderName"`
	// A mapping of tags which should be assigned to the Security Partner Provider.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Hub within which this Security Partner Provider should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a SecurityPartnerProvider resource.
type SecurityPartnerProviderArgs struct {
	// The Azure Region where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Security Partner Provider. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Security Partner Provider should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The security provider name. Possible values are `ZScaler`, `IBoss` and `Checkpoint` is allowed. Changing this forces a new resource to be created.
	SecurityProviderName pulumi.StringInput
	// A mapping of tags which should be assigned to the Security Partner Provider.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Hub within which this Security Partner Provider should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (SecurityPartnerProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPartnerProviderArgs)(nil)).Elem()
}

type SecurityPartnerProviderInput interface {
	pulumi.Input

	ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput
	ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput
}

func (*SecurityPartnerProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPartnerProvider)(nil))
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput {
	return i.ToSecurityPartnerProviderOutputWithContext(context.Background())
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPartnerProviderOutput)
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderPtrOutput() SecurityPartnerProviderPtrOutput {
	return i.ToSecurityPartnerProviderPtrOutputWithContext(context.Background())
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderPtrOutputWithContext(ctx context.Context) SecurityPartnerProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPartnerProviderPtrOutput)
}

type SecurityPartnerProviderPtrInput interface {
	pulumi.Input

	ToSecurityPartnerProviderPtrOutput() SecurityPartnerProviderPtrOutput
	ToSecurityPartnerProviderPtrOutputWithContext(ctx context.Context) SecurityPartnerProviderPtrOutput
}

type SecurityPartnerProviderOutput struct {
	*pulumi.OutputState
}

func (SecurityPartnerProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPartnerProvider)(nil))
}

func (o SecurityPartnerProviderOutput) ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput {
	return o
}

func (o SecurityPartnerProviderOutput) ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput {
	return o
}

type SecurityPartnerProviderPtrOutput struct {
	*pulumi.OutputState
}

func (SecurityPartnerProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPartnerProvider)(nil))
}

func (o SecurityPartnerProviderPtrOutput) ToSecurityPartnerProviderPtrOutput() SecurityPartnerProviderPtrOutput {
	return o
}

func (o SecurityPartnerProviderPtrOutput) ToSecurityPartnerProviderPtrOutputWithContext(ctx context.Context) SecurityPartnerProviderPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecurityPartnerProviderOutput{})
	pulumi.RegisterOutputType(SecurityPartnerProviderPtrOutput{})
}
