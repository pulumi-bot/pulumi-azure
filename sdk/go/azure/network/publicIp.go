// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Public IP Address.
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
// 		_, err = network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			AllocationMethod:  pulumi.String("Static"),
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
type PublicIp struct {
	pulumi.CustomResourceState

	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod pulumi.StringOutput `pulumi:"allocationMethod"`
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel pulumi.StringPtrOutput `pulumi:"domainNameLabel"`
	// Fully qualified domain name of the A DNS record associated with the public IP. `domainNameLabel` must be specified to get the `fqdn`. This is the concatenation of the `domainNameLabel` and the regionalized DNS zone
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// The IP address value that was allocated.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The IP Version to use, IPv6 or IPv4.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId pulumi.StringPtrOutput `pulumi:"publicIpPrefixId"`
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn pulumi.StringPtrOutput `pulumi:"reverseFqdn"`
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Public IP in.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewPublicIp registers a new resource with the given unique name, arguments, and options.
func NewPublicIp(ctx *pulumi.Context,
	name string, args *PublicIpArgs, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AllocationMethod == nil {
		return nil, errors.New("invalid value for required argument 'AllocationMethod'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource PublicIp
	err := ctx.RegisterResource("azure:network/publicIp:PublicIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIp gets an existing PublicIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpState, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	var resource PublicIp
	err := ctx.ReadResource("azure:network/publicIp:PublicIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIp resources.
type publicIpState struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod *string `pulumi:"allocationMethod"`
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	// Fully qualified domain name of the A DNS record associated with the public IP. `domainNameLabel` must be specified to get the `fqdn`. This is the concatenation of the `domainNameLabel` and the regionalized DNS zone
	Fqdn *string `pulumi:"fqdn"`
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The IP address value that was allocated.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP Version to use, IPv6 or IPv4.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId *string `pulumi:"publicIpPrefixId"`
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `pulumi:"reverseFqdn"`
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Public IP in.
	Zones *string `pulumi:"zones"`
}

type PublicIpState struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod pulumi.StringPtrInput
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel pulumi.StringPtrInput
	// Fully qualified domain name of the A DNS record associated with the public IP. `domainNameLabel` must be specified to get the `fqdn`. This is the concatenation of the `domainNameLabel` and the regionalized DNS zone
	Fqdn pulumi.StringPtrInput
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The IP address value that was allocated.
	IpAddress pulumi.StringPtrInput
	// The IP Version to use, IPv6 or IPv4.
	IpVersion pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName pulumi.StringPtrInput
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn pulumi.StringPtrInput
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Public IP in.
	Zones pulumi.StringPtrInput
}

func (PublicIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpState)(nil)).Elem()
}

type publicIpArgs struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod string `pulumi:"allocationMethod"`
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The IP Version to use, IPv6 or IPv4.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId *string `pulumi:"publicIpPrefixId"`
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `pulumi:"reverseFqdn"`
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Public IP in.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIp resource.
type PublicIpArgs struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod pulumi.StringInput
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel pulumi.StringPtrInput
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The IP Version to use, IPv6 or IPv4.
	IpVersion pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName pulumi.StringInput
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn pulumi.StringPtrInput
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Public IP in.
	Zones pulumi.StringPtrInput
}

func (PublicIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpArgs)(nil)).Elem()
}
