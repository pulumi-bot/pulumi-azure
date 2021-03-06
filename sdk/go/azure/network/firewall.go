// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Firewall.
type Firewall struct {
	pulumi.CustomResourceState

	// A `ipConfiguration` block as documented below.
	IpConfigurations FirewallIpConfigurationArrayOutput `pulumi:"ipConfigurations"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the availability zones in which the Azure Firewall should be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil || args.IpConfigurations == nil {
		return nil, errors.New("missing required argument 'IpConfigurations'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FirewallArgs{}
	}
	var resource Firewall
	err := ctx.RegisterResource("azure:network/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("azure:network/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// A `ipConfiguration` block as documented below.
	IpConfigurations []FirewallIpConfiguration `pulumi:"ipConfigurations"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the availability zones in which the Azure Firewall should be created.
	Zones []string `pulumi:"zones"`
}

type FirewallState struct {
	// A `ipConfiguration` block as documented below.
	IpConfigurations FirewallIpConfigurationArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the availability zones in which the Azure Firewall should be created.
	Zones pulumi.StringArrayInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// A `ipConfiguration` block as documented below.
	IpConfigurations []FirewallIpConfiguration `pulumi:"ipConfigurations"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the availability zones in which the Azure Firewall should be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// A `ipConfiguration` block as documented below.
	IpConfigurations FirewallIpConfigurationArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the availability zones in which the Azure Firewall should be created.
	Zones pulumi.StringArrayInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}
