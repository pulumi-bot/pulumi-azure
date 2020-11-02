// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a network security group that contains a list of network security rules.  Network security groups enable inbound or outbound traffic to be enabled or denied.
//
// > **NOTE on Network Security Groups and Network Security Rules:** This provider currently
// provides both a standalone Network Security Rule resource, and allows for Network Security Rules to be defined in-line within the Network Security Group resource.
// At this time you cannot use a Network Security Group with in-line Network Security Rules in conjunction with any Network Security Rule resources. Doing so will cause a conflict of rule settings and will overwrite rules.
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
// 		_, err = network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SecurityRules: network.NetworkSecurityGroupSecurityRuleArray{
// 				&network.NetworkSecurityGroupSecurityRuleArgs{
// 					Name:                     pulumi.String("test123"),
// 					Priority:                 pulumi.Int(100),
// 					Direction:                pulumi.String("Inbound"),
// 					Access:                   pulumi.String("Allow"),
// 					Protocol:                 pulumi.String("Tcp"),
// 					SourcePortRange:          pulumi.String("*"),
// 					DestinationPortRange:     pulumi.String("*"),
// 					SourceAddressPrefix:      pulumi.String("*"),
// 					DestinationAddressPrefix: pulumi.String("*"),
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
// Network Security Groups can be imported using the `resource id`, e.g.
type NetworkSecurityGroup struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the security rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A list of objects representing security rules, as defined below.
	SecurityRules NetworkSecurityGroupSecurityRuleArrayOutput `pulumi:"securityRules"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNetworkSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkSecurityGroupArgs{}
	}
	var resource NetworkSecurityGroup
	err := ctx.RegisterResource("azure:network/networkSecurityGroup:NetworkSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityGroup gets an existing NetworkSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	var resource NetworkSecurityGroup
	err := ctx.ReadResource("azure:network/networkSecurityGroup:NetworkSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityGroup resources.
type networkSecurityGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the security rule.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A list of objects representing security rules, as defined below.
	SecurityRules []NetworkSecurityGroupSecurityRule `pulumi:"securityRules"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type NetworkSecurityGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the security rule.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A list of objects representing security rules, as defined below.
	SecurityRules NetworkSecurityGroupSecurityRuleArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupState)(nil)).Elem()
}

type networkSecurityGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the security rule.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A list of objects representing security rules, as defined below.
	SecurityRules []NetworkSecurityGroupSecurityRule `pulumi:"securityRules"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkSecurityGroup resource.
type NetworkSecurityGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the security rule.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A list of objects representing security rules, as defined below.
	SecurityRules NetworkSecurityGroupSecurityRuleArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupArgs)(nil)).Elem()
}
