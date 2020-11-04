// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Data Lake Analytics Firewall Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datalake"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleStore, err := datalake.NewStore(ctx, "exampleStore", &datalake.StoreArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewAnalyticsAccount(ctx, "exampleAnalyticsAccount", &datalake.AnalyticsAccountArgs{
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			Location:                exampleResourceGroup.Location,
// 			DefaultStoreAccountName: exampleStore.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewAnalyticsFirewallRule(ctx, "exampleAnalyticsFirewallRule", &datalake.AnalyticsFirewallRuleArgs{
// 			AccountName:       pulumi.Any(azurerm_data_lake_analytics.Example.Name),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StartIpAddress:    pulumi.String("1.2.3.4"),
// 			EndIpAddress:      pulumi.String("2.3.4.5"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AnalyticsFirewallRule struct {
	pulumi.CustomResourceState

	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewAnalyticsFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsFirewallRule(ctx *pulumi.Context,
	name string, args *AnalyticsFirewallRuleArgs, opts ...pulumi.ResourceOption) (*AnalyticsFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	var resource AnalyticsFirewallRule
	err := ctx.RegisterResource("azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsFirewallRule gets an existing AnalyticsFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsFirewallRuleState, opts ...pulumi.ResourceOption) (*AnalyticsFirewallRule, error) {
	var resource AnalyticsFirewallRule
	err := ctx.ReadResource("azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsFirewallRule resources.
type analyticsFirewallRuleState struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName *string `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type AnalyticsFirewallRuleState struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName pulumi.StringPtrInput
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringPtrInput
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName pulumi.StringPtrInput
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringPtrInput
}

func (AnalyticsFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsFirewallRuleState)(nil)).Elem()
}

type analyticsFirewallRuleArgs struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName string `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a AnalyticsFirewallRule resource.
type AnalyticsFirewallRuleArgs struct {
	// Specifies the name of the Data Lake Analytics for which the Firewall Rule should take effect.
	AccountName pulumi.StringInput
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringInput
	// Specifies the name of the Data Lake Analytics. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Analytics.
	ResourceGroupName pulumi.StringInput
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringInput
}

func (AnalyticsFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsFirewallRuleArgs)(nil)).Elem()
}
