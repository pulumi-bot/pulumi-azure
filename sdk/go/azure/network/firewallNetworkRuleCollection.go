// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Network Rule Collection within an Azure Firewall.
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
// 			Location: pulumi.String("North Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
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
// 				pulumi.String("10.0.1.0/24"),
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
// 		exampleFirewall, err := network.NewFirewall(ctx, "exampleFirewall", &network.FirewallArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.FirewallIpConfigurationArray{
// 				&network.FirewallIpConfigurationArgs{
// 					Name:              pulumi.String("configuration"),
// 					SubnetId:          exampleSubnet.ID(),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewFirewallNetworkRuleCollection(ctx, "exampleFirewallNetworkRuleCollection", &network.FirewallNetworkRuleCollectionArgs{
// 			AzureFirewallName: exampleFirewall.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Priority:          pulumi.Int(100),
// 			Action:            pulumi.String("Allow"),
// 			Rules: network.FirewallNetworkRuleCollectionRuleArray{
// 				&network.FirewallNetworkRuleCollectionRuleArgs{
// 					Name: pulumi.String("testrule"),
// 					SourceAddresses: pulumi.StringArray{
// 						pulumi.String("10.0.0.0/16"),
// 					},
// 					DestinationPorts: pulumi.StringArray{
// 						pulumi.String("53"),
// 					},
// 					DestinationAddresses: pulumi.StringArray{
// 						pulumi.String("8.8.8.8"),
// 						pulumi.String("8.8.4.4"),
// 					},
// 					Protocols: pulumi.StringArray{
// 						pulumi.String("TCP"),
// 						pulumi.String("UDP"),
// 					},
// 				},
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
// Azure Firewall Network Rule Collections can be imported using the `resource id`, e.g.
type FirewallNetworkRuleCollection struct {
	pulumi.CustomResourceState

	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringOutput `pulumi:"azureFirewallName"`
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules FirewallNetworkRuleCollectionRuleArrayOutput `pulumi:"rules"`
}

// NewFirewallNetworkRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewFirewallNetworkRuleCollection(ctx *pulumi.Context,
	name string, args *FirewallNetworkRuleCollectionArgs, opts ...pulumi.ResourceOption) (*FirewallNetworkRuleCollection, error) {
	if args == nil || args.Action == nil {
		return nil, errors.New("missing required argument 'Action'")
	}
	if args == nil || args.AzureFirewallName == nil {
		return nil, errors.New("missing required argument 'AzureFirewallName'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	if args == nil {
		args = &FirewallNetworkRuleCollectionArgs{}
	}
	var resource FirewallNetworkRuleCollection
	err := ctx.RegisterResource("azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallNetworkRuleCollection gets an existing FirewallNetworkRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallNetworkRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallNetworkRuleCollectionState, opts ...pulumi.ResourceOption) (*FirewallNetworkRuleCollection, error) {
	var resource FirewallNetworkRuleCollection
	err := ctx.ReadResource("azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallNetworkRuleCollection resources.
type firewallNetworkRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action *string `pulumi:"action"`
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName *string `pulumi:"azureFirewallName"`
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority *int `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules []FirewallNetworkRuleCollectionRule `pulumi:"rules"`
}

type FirewallNetworkRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringPtrInput
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringPtrInput
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntPtrInput
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `rule` blocks as defined below.
	Rules FirewallNetworkRuleCollectionRuleArrayInput
}

func (FirewallNetworkRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallNetworkRuleCollectionState)(nil)).Elem()
}

type firewallNetworkRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action string `pulumi:"action"`
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName string `pulumi:"azureFirewallName"`
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority int `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules []FirewallNetworkRuleCollectionRule `pulumi:"rules"`
}

// The set of arguments for constructing a FirewallNetworkRuleCollection resource.
type FirewallNetworkRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringInput
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringInput
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntInput
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `rule` blocks as defined below.
	Rules FirewallNetworkRuleCollectionRuleArrayInput
}

func (FirewallNetworkRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallNetworkRuleCollectionArgs)(nil)).Elem()
}
