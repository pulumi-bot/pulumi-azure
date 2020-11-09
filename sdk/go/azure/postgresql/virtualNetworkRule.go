// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a PostgreSQL Virtual Network Rule.
//
// > **NOTE:** PostgreSQL Virtual Network Rules [can only be used with SKU Tiers of `GeneralPurpose` or `MemoryOptimized`](https://docs.microsoft.com/en-us/azure/postgresql/concepts-data-access-and-security-vnet)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/postgresql"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.7.29.0/29"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		internal, err := network.NewSubnet(ctx, "internal", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.7.29.0/29"),
// 			},
// 			ServiceEndpoints: pulumi.StringArray{
// 				pulumi.String("Microsoft.Sql"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServer, err := postgresql.NewServer(ctx, "exampleServer", &postgresql.ServerArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("GP_Gen5_2"),
// 			StorageProfile: &postgresql.ServerStorageProfileArgs{
// 				StorageMb:           pulumi.Int(5120),
// 				BackupRetentionDays: pulumi.Int(7),
// 				GeoRedundantBackup:  pulumi.String("Disabled"),
// 			},
// 			AdministratorLogin:         pulumi.String("psqladminun"),
// 			AdministratorLoginPassword: pulumi.String("H@Sh1CoR3!"),
// 			Version:                    pulumi.String("9.5"),
// 			SslEnforcement:             pulumi.String("Enabled"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = postgresql.NewVirtualNetworkRule(ctx, "exampleVirtualNetworkRule", &postgresql.VirtualNetworkRuleArgs{
// 			ResourceGroupName:                exampleResourceGroup.Name,
// 			ServerName:                       exampleServer.Name,
// 			SubnetId:                         internal.ID(),
// 			IgnoreMissingVnetServiceEndpoint: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled? Defaults to `false`.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrOutput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the PostgreSQL virtual network rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The ID of the subnet that the PostgreSQL server will be connected to.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &VirtualNetworkRuleArgs{}
	}
	var resource VirtualNetworkRule
	err := ctx.RegisterResource("azure:postgresql/virtualNetworkRule:VirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkRule gets an existing VirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	var resource VirtualNetworkRule
	err := ctx.ReadResource("azure:postgresql/virtualNetworkRule:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type virtualNetworkRuleState struct {
	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled? Defaults to `false`.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the PostgreSQL virtual network rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// The ID of the subnet that the PostgreSQL server will be connected to.
	SubnetId *string `pulumi:"subnetId"`
}

type VirtualNetworkRuleState struct {
	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled? Defaults to `false`.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// The name of the PostgreSQL virtual network rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// The ID of the subnet that the PostgreSQL server will be connected to.
	SubnetId pulumi.StringPtrInput
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled? Defaults to `false`.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the PostgreSQL virtual network rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// The ID of the subnet that the PostgreSQL server will be connected to.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled? Defaults to `false`.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// The name of the PostgreSQL virtual network rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// The ID of the subnet that the PostgreSQL server will be connected to.
	SubnetId pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}

type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput
}

func (VirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRule) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRule) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleOutput struct {
	*pulumi.OutputState
}

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleOutput)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
}
