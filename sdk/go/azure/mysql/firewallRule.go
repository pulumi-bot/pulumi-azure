// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Firewall Rule for a MySQL Server.
//
// ## Example Usage
// ### Single IP Address)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mysql"
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
// 		exampleServer, err := mysql.NewServer(ctx, "exampleServer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewFirewallRule(ctx, "exampleFirewallRule", &mysql.FirewallRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			StartIpAddress:    pulumi.String("40.112.8.12"),
// 			EndIpAddress:      pulumi.String("40.112.8.12"),
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mysql"
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
// 		exampleServer, err := mysql.NewServer(ctx, "exampleServer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewFirewallRule(ctx, "exampleFirewallRule", &mysql.FirewallRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			StartIpAddress:    pulumi.String("40.112.0.0"),
// 			EndIpAddress:      pulumi.String("40.112.255.255"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Allow Access To Azure Services)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mysql"
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
// 		exampleServer, err := mysql.NewServer(ctx, "exampleServer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewFirewallRule(ctx, "exampleFirewallRule", &mysql.FirewallRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			StartIpAddress:    pulumi.String("0.0.0.0"),
// 			EndIpAddress:      pulumi.String("0.0.0.0"),
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
	// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil || args.EndIpAddress == nil {
		return nil, errors.New("missing required argument 'EndIpAddress'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.StartIpAddress == nil {
		return nil, errors.New("missing required argument 'StartIpAddress'")
	}
	if args == nil {
		args = &FirewallRuleArgs{}
	}
	var resource FirewallRule
	err := ctx.RegisterResource("azure:mysql/firewallRule:FirewallRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:mysql/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type FirewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringPtrInput
	// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
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
	// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringInput
	// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRule)(nil)).Elem()
}

func (i FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

type FirewallRuleOutput struct {
	*pulumi.OutputState
}

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRuleOutput)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
}
