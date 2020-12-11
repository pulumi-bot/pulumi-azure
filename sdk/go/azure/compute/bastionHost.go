// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Bastion Host.
//
// > **Note:** Bastion Hosts are a preview feature in Azure, and therefore are only supported in a select number of regions. [Read more](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq).
//
// ## Example Usage
//
// This example deploys an Azure Bastion Host Instance to a target virtual network.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
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
// 				pulumi.String("192.168.1.0/24"),
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
// 				pulumi.String("192.168.1.224/27"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewBastionHost(ctx, "exampleBastionHost", &compute.BastionHostArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfiguration: &compute.BastionHostIpConfigurationArgs{
// 				Name:              pulumi.String("configuration"),
// 				SubnetId:          exampleSubnet.ID(),
// 				PublicIpAddressId: examplePublicIp.ID(),
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
// Bastion Hosts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/bastionHost:BastionHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
// ```
type BastionHost struct {
	pulumi.CustomResourceState

	// The FQDN for the Bastion Host.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// A `ipConfiguration` block as defined below.
	IpConfiguration BastionHostIpConfigurationPtrOutput `pulumi:"ipConfiguration"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewBastionHost registers a new resource with the given unique name, arguments, and options.
func NewBastionHost(ctx *pulumi.Context,
	name string, args *BastionHostArgs, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource BastionHost
	err := ctx.RegisterResource("azure:compute/bastionHost:BastionHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBastionHost gets an existing BastionHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBastionHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BastionHostState, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	var resource BastionHost
	err := ctx.ReadResource("azure:compute/bastionHost:BastionHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BastionHost resources.
type bastionHostState struct {
	// The FQDN for the Bastion Host.
	DnsName *string `pulumi:"dnsName"`
	// A `ipConfiguration` block as defined below.
	IpConfiguration *BastionHostIpConfiguration `pulumi:"ipConfiguration"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type BastionHostState struct {
	// The FQDN for the Bastion Host.
	DnsName pulumi.StringPtrInput
	// A `ipConfiguration` block as defined below.
	IpConfiguration BastionHostIpConfigurationPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (BastionHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostState)(nil)).Elem()
}

type bastionHostArgs struct {
	// A `ipConfiguration` block as defined below.
	IpConfiguration *BastionHostIpConfiguration `pulumi:"ipConfiguration"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BastionHost resource.
type BastionHostArgs struct {
	// A `ipConfiguration` block as defined below.
	IpConfiguration BastionHostIpConfigurationPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (BastionHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostArgs)(nil)).Elem()
}

type BastionHostInput interface {
	pulumi.Input

	ToBastionHostOutput() BastionHostOutput
	ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput
}

type BastionHostPtrInput interface {
	pulumi.Input

	ToBastionHostPtrOutput() BastionHostPtrOutput
	ToBastionHostPtrOutputWithContext(ctx context.Context) BastionHostPtrOutput
}

func (BastionHost) ElementType() reflect.Type {
	return reflect.TypeOf((*BastionHost)(nil)).Elem()
}

func (i BastionHost) ToBastionHostOutput() BastionHostOutput {
	return i.ToBastionHostOutputWithContext(context.Background())
}

func (i BastionHost) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostOutput)
}

func (i BastionHost) ToBastionHostPtrOutput() BastionHostPtrOutput {
	return i.ToBastionHostPtrOutputWithContext(context.Background())
}

func (i BastionHost) ToBastionHostPtrOutputWithContext(ctx context.Context) BastionHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostPtrOutput)
}

type BastionHostOutput struct {
	*pulumi.OutputState
}

func (BastionHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BastionHostOutput)(nil)).Elem()
}

func (o BastionHostOutput) ToBastionHostOutput() BastionHostOutput {
	return o
}

func (o BastionHostOutput) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return o
}

type BastionHostPtrOutput struct {
	*pulumi.OutputState
}

func (BastionHostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (o BastionHostPtrOutput) ToBastionHostPtrOutput() BastionHostPtrOutput {
	return o
}

func (o BastionHostPtrOutput) ToBastionHostPtrOutputWithContext(ctx context.Context) BastionHostPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BastionHostOutput{})
	pulumi.RegisterOutputType(BastionHostPtrOutput{})
}
