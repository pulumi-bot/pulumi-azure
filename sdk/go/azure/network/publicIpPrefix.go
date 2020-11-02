// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Public IP Prefix.
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
// 		_, err = network.NewPublicIpPrefix(ctx, "examplePublicIpPrefix", &network.PublicIpPrefixArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PrefixLength:      pulumi.Int(31),
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
// Public IP Prefixes can be imported using the `resource id`, e.g.
type PublicIpPrefix struct {
	pulumi.CustomResourceState

	// The IP address prefix value that was allocated.
	IpPrefix pulumi.StringOutput `pulumi:"ipPrefix"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Public IP Prefix in.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewPublicIpPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicIpPrefix(ctx *pulumi.Context,
	name string, args *PublicIpPrefixArgs, opts ...pulumi.ResourceOption) (*PublicIpPrefix, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PublicIpPrefixArgs{}
	}
	var resource PublicIpPrefix
	err := ctx.RegisterResource("azure:network/publicIpPrefix:PublicIpPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIpPrefix gets an existing PublicIpPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIpPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpPrefixState, opts ...pulumi.ResourceOption) (*PublicIpPrefix, error) {
	var resource PublicIpPrefix
	err := ctx.ReadResource("azure:network/publicIpPrefix:PublicIpPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIpPrefix resources.
type publicIpPrefixState struct {
	// The IP address prefix value that was allocated.
	IpPrefix *string `pulumi:"ipPrefix"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength *int `pulumi:"prefixLength"`
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Public IP Prefix in.
	Zones *string `pulumi:"zones"`
}

type PublicIpPrefixState struct {
	// The IP address prefix value that was allocated.
	IpPrefix pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength pulumi.IntPtrInput
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Public IP Prefix in.
	Zones pulumi.StringPtrInput
}

func (PublicIpPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpPrefixState)(nil)).Elem()
}

type publicIpPrefixArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength *int `pulumi:"prefixLength"`
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Public IP Prefix in.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIpPrefix resource.
type PublicIpPrefixArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength pulumi.IntPtrInput
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Public IP Prefix in.
	Zones pulumi.StringPtrInput
}

func (PublicIpPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpPrefixArgs)(nil)).Elem()
}
