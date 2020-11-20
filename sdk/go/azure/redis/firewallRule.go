// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Firewall Rule associated with a Redis Cache.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/redis"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
// 			Keepers: pulumi.Float64Map{
// 				"azi_id": pulumi.Float64(1),
// 			},
// 			ByteLength: pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCache, err := redis.NewCache(ctx, "exampleCache", &redis.CacheArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Capacity:          pulumi.Int(1),
// 			Family:            pulumi.String("P"),
// 			SkuName:           pulumi.String("Premium"),
// 			EnableNonSslPort:  pulumi.Bool(false),
// 			RedisConfiguration: &redis.CacheRedisConfigurationArgs{
// 				Maxclients:        pulumi.Int(256),
// 				MaxmemoryReserved: pulumi.Int(2),
// 				MaxmemoryDelta:    pulumi.Int(2),
// 				MaxmemoryPolicy:   pulumi.String("allkeys-lru"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = redis.NewFirewallRule(ctx, "exampleFirewallRule", &redis.FirewallRuleArgs{
// 			RedisCacheName:    exampleCache.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StartIp:           pulumi.String("1.2.3.4"),
// 			EndIp:             pulumi.String("2.3.4.5"),
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
// Redis Firewall Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:redis/firewallRule:FirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1/firewallRules/rule1
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIp == nil {
		return nil, errors.New("invalid value for required argument 'EndIp'")
	}
	if args.RedisCacheName == nil {
		return nil, errors.New("invalid value for required argument 'RedisCacheName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIp == nil {
		return nil, errors.New("invalid value for required argument 'StartIp'")
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
