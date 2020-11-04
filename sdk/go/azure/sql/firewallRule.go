// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to manage an Azure SQL Firewall Rule
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sql"
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
// 		exampleSqlServer, err := sql.NewSqlServer(ctx, "exampleSqlServer", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   pulumi.String("West US"),
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("4dm1n157r470r"),
// 			AdministratorLoginPassword: pulumi.String("4-v3ry-53cr37-p455w0rd"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewFirewallRule(ctx, "exampleFirewallRule", &sql.FirewallRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleSqlServer.Name,
// 			StartIpAddress:    pulumi.String("10.0.17.62"),
// 			EndIpAddress:      pulumi.String("10.0.17.62"),
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

	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// The name of the firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the sql server.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the Firewall Rule.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The starting IP address to allow through the firewall for this rule.
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
	err := ctx.RegisterResource("azure:sql/firewallRule:FirewallRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:sql/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the sql server.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the Firewall Rule.
	ServerName *string `pulumi:"serverName"`
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type FirewallRuleState struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress pulumi.StringPtrInput
	// The name of the firewall rule.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the sql server.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the SQL Server on which to create the Firewall Rule.
	ServerName pulumi.StringPtrInput
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the sql server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the Firewall Rule.
	ServerName string `pulumi:"serverName"`
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The ending IP address to allow through the firewall for this rule.
	EndIpAddress pulumi.StringInput
	// The name of the firewall rule.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the sql server.
	ResourceGroupName pulumi.StringInput
	// The name of the SQL Server on which to create the Firewall Rule.
	ServerName pulumi.StringInput
	// The starting IP address to allow through the firewall for this rule.
	StartIpAddress pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}
