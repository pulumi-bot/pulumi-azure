// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Firewall Rule associated with a Redis Cache.
type FirewallRule struct {
	pulumi.CustomResourceState

	// The highest IP address included in the range.
	EndIp pulumi.StringOutput `pulumi:"endIp"`
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName pulumi.StringOutput `pulumi:"redisCacheName"`
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The lowest IP address included in the range
	StartIp pulumi.StringOutput `pulumi:"startIp"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil || args.EndIp == nil {
		return nil, errors.New("missing required argument 'EndIp'")
	}
	if args == nil || args.RedisCacheName == nil {
		return nil, errors.New("missing required argument 'RedisCacheName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StartIp == nil {
		return nil, errors.New("missing required argument 'StartIp'")
	}
	if args == nil {
		args = &FirewallRuleArgs{}
	}
	var resource FirewallRule
	err := ctx.RegisterResource("azure:redis/firewallRule:FirewallRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:redis/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The highest IP address included in the range.
	EndIp *string `pulumi:"endIp"`
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName *string `pulumi:"redisCacheName"`
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The lowest IP address included in the range
	StartIp *string `pulumi:"startIp"`
}

type FirewallRuleState struct {
	// The highest IP address included in the range.
	EndIp pulumi.StringPtrInput
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName pulumi.StringPtrInput
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName pulumi.StringPtrInput
	// The lowest IP address included in the range
	StartIp pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The highest IP address included in the range.
	EndIp string `pulumi:"endIp"`
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName string `pulumi:"redisCacheName"`
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The lowest IP address included in the range
	StartIp string `pulumi:"startIp"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The highest IP address included in the range.
	EndIp pulumi.StringInput
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName pulumi.StringInput
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName pulumi.StringInput
	// The lowest IP address included in the range
	StartIp pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}
