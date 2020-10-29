// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a MySQL Virtual Network Rule.
//
// > **NOTE:** MySQL Virtual Network Rules [can only be used with SKU Tiers of `GeneralPurpose` or `MemoryOptimized`](https://docs.microsoft.com/en-us/azure/mysql/concepts-data-access-and-security-vnet)
type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The ID of the subnet that the MySQL server will be connected to.
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
	err := ctx.RegisterResource("azure:mysql/virtualNetworkRule:VirtualNetworkRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:mysql/virtualNetworkRule:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type virtualNetworkRuleState struct {
	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// The ID of the subnet that the MySQL server will be connected to.
	SubnetId *string `pulumi:"subnetId"`
}

type VirtualNetworkRuleState struct {
	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// The ID of the subnet that the MySQL server will be connected to.
	SubnetId pulumi.StringPtrInput
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// The ID of the subnet that the MySQL server will be connected to.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// The ID of the subnet that the MySQL server will be connected to.
	SubnetId pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}
