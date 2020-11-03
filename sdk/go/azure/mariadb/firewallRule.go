// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Firewall Rule for a MariaDB Server
//
// ## Example Usage
// ### Single IP Address)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mariadb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mariadb.NewFirewallRule(ctx, "example", &mariadb.FirewallRuleArgs{
// 			EndIpAddress:      pulumi.String("40.112.8.12"),
// 			ResourceGroupName: pulumi.String("test-rg"),
// 			ServerName:        pulumi.String("test-server"),
// 			StartIpAddress:    pulumi.String("40.112.8.12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### IP Range)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mariadb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mariadb.NewFirewallRule(ctx, "example", &mariadb.FirewallRuleArgs{
// 			EndIpAddress:      pulumi.String("40.112.255.255"),
// 			ResourceGroupName: pulumi.String("test-rg"),
// 			ServerName:        pulumi.String("test-server"),
// 			StartIpAddress:    pulumi.String("40.112.0.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	var resource FirewallRule
	err := ctx.RegisterResource("azure:mariadb/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure:mariadb/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type FirewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringPtrInput
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringInput
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}
