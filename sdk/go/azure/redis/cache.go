// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Redis Cache.
//
// ## Default Redis Configuration Values
//
// | Redis Value                     | Basic        | Standard     | Premium      |
// | ------------------------------- | ------------ | ------------ | ------------ |
// | enableAuthentication           | true         | true         | true         |
// | maxmemoryReserved              | 2            | 50           | 200          |
// | maxfragmentationmemoryReserved | 2            | 50           | 200          |
// | maxmemoryDelta                 | 2            | 50           | 200          |
// | maxmemoryPolicy                | volatile-lru | volatile-lru | volatile-lru |
//
// > **NOTE:** The `maxmemoryReserved`, `maxmemoryDelta` and `maxfragmentationmemory-reserved` settings are only available for Standard and Premium caches. More details are available in the Relevant Links section below._
//
// ***
//
// A `patchSchedule` block supports the following:
//
// * `dayOfWeek` (Required) the Weekday name - possible values include `Monday`, `Tuesday`, `Wednesday` etc.
//
// * `startHourUtc` - (Optional) the Start Hour for maintenance in UTC - possible values range from `0 - 23`.
//
// > **Note:** The Patch Window lasts for `5` hours from the `startHourUtc`.
//
// ## Relevant Links
//
//  - [Azure Redis Cache: SKU specific configuration limitations](https://azure.microsoft.com/en-us/documentation/articles/cache-configure/#advanced-settings)
//  - [Redis: Available Configuration Settings](http://redis.io/topics/config)
type Cache struct {
	pulumi.CustomResourceState

	// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// Enable the non-SSL port (6379) - disabled by default.
	EnableNonSslPort pulumi.BoolPtrOutput `pulumi:"enableNonSslPort"`
	// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
	Family pulumi.StringOutput `pulumi:"family"`
	// The Hostname of the Redis Instance
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The location of the resource group.
	Location pulumi.StringOutput `pulumi:"location"`
	// The minimum TLS version.  Defaults to `1.0`.
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the Redis instance. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
	PatchSchedules CachePatchScheduleArrayOutput `pulumi:"patchSchedules"`
	// The non-SSL Port of the Redis Instance
	Port pulumi.IntOutput `pulumi:"port"`
	// The Primary Access Key for the Redis Instance
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The primary connection string of the Redis Instance.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
	PrivateStaticIpAddress pulumi.StringOutput `pulumi:"privateStaticIpAddress"`
	// A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
	RedisConfiguration CacheRedisConfigurationOutput `pulumi:"redisConfiguration"`
	// The name of the resource group in which to
	// create the Redis instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Access Key for the Redis Instance
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The secondary connection string of the Redis Instance.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
	ShardCount pulumi.IntPtrOutput `pulumi:"shardCount"`
	// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// The SSL Port of the Redis Instance
	SslPort pulumi.IntOutput `pulumi:"sslPort"`
	// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewCache registers a new resource with the given unique name, arguments, and options.
func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil || args.Capacity == nil {
		return nil, errors.New("missing required argument 'Capacity'")
	}
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil {
		args = &CacheArgs{}
	}
	var resource Cache
	err := ctx.RegisterResource("azure:redis/cache:Cache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCache gets an existing Cache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheState, opts ...pulumi.ResourceOption) (*Cache, error) {
	var resource Cache
	err := ctx.ReadResource("azure:redis/cache:Cache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cache resources.
type cacheState struct {
	// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
	Capacity *int `pulumi:"capacity"`
	// Enable the non-SSL port (6379) - disabled by default.
	EnableNonSslPort *bool `pulumi:"enableNonSslPort"`
	// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
	Family *string `pulumi:"family"`
	// The Hostname of the Redis Instance
	Hostname *string `pulumi:"hostname"`
	// The location of the resource group.
	Location *string `pulumi:"location"`
	// The minimum TLS version.  Defaults to `1.0`.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the Redis instance. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
	PatchSchedules []CachePatchSchedule `pulumi:"patchSchedules"`
	// The non-SSL Port of the Redis Instance
	Port *int `pulumi:"port"`
	// The Primary Access Key for the Redis Instance
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The primary connection string of the Redis Instance.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
	PrivateStaticIpAddress *string `pulumi:"privateStaticIpAddress"`
	// A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
	RedisConfiguration *CacheRedisConfiguration `pulumi:"redisConfiguration"`
	// The name of the resource group in which to
	// create the Redis instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Access Key for the Redis Instance
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The secondary connection string of the Redis Instance.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
	ShardCount *int `pulumi:"shardCount"`
	// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
	SkuName *string `pulumi:"skuName"`
	// The SSL Port of the Redis Instance
	SslPort *int `pulumi:"sslPort"`
	// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
	Zones *string `pulumi:"zones"`
}

type CacheState struct {
	// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
	Capacity pulumi.IntPtrInput
	// Enable the non-SSL port (6379) - disabled by default.
	EnableNonSslPort pulumi.BoolPtrInput
	// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
	Family pulumi.StringPtrInput
	// The Hostname of the Redis Instance
	Hostname pulumi.StringPtrInput
	// The location of the resource group.
	Location pulumi.StringPtrInput
	// The minimum TLS version.  Defaults to `1.0`.
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the Redis instance. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
	PatchSchedules CachePatchScheduleArrayInput
	// The non-SSL Port of the Redis Instance
	Port pulumi.IntPtrInput
	// The Primary Access Key for the Redis Instance
	PrimaryAccessKey pulumi.StringPtrInput
	// The primary connection string of the Redis Instance.
	PrimaryConnectionString pulumi.StringPtrInput
	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
	PrivateStaticIpAddress pulumi.StringPtrInput
	// A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
	RedisConfiguration CacheRedisConfigurationPtrInput
	// The name of the resource group in which to
	// create the Redis instance.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Access Key for the Redis Instance
	SecondaryAccessKey pulumi.StringPtrInput
	// The secondary connection string of the Redis Instance.
	SecondaryConnectionString pulumi.StringPtrInput
	// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
	ShardCount pulumi.IntPtrInput
	// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
	SkuName pulumi.StringPtrInput
	// The SSL Port of the Redis Instance
	SslPort pulumi.IntPtrInput
	// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
	Zones pulumi.StringPtrInput
}

func (CacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheState)(nil)).Elem()
}

type cacheArgs struct {
	// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
	Capacity int `pulumi:"capacity"`
	// Enable the non-SSL port (6379) - disabled by default.
	EnableNonSslPort *bool `pulumi:"enableNonSslPort"`
	// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
	Family string `pulumi:"family"`
	// The location of the resource group.
	Location *string `pulumi:"location"`
	// The minimum TLS version.  Defaults to `1.0`.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the Redis instance. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
	PatchSchedules []CachePatchSchedule `pulumi:"patchSchedules"`
	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
	PrivateStaticIpAddress *string `pulumi:"privateStaticIpAddress"`
	// A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
	RedisConfiguration *CacheRedisConfiguration `pulumi:"redisConfiguration"`
	// The name of the resource group in which to
	// create the Redis instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
	ShardCount *int `pulumi:"shardCount"`
	// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
	SkuName string `pulumi:"skuName"`
	// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a Cache resource.
type CacheArgs struct {
	// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
	Capacity pulumi.IntInput
	// Enable the non-SSL port (6379) - disabled by default.
	EnableNonSslPort pulumi.BoolPtrInput
	// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
	Family pulumi.StringInput
	// The location of the resource group.
	Location pulumi.StringPtrInput
	// The minimum TLS version.  Defaults to `1.0`.
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the Redis instance. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
	PatchSchedules CachePatchScheduleArrayInput
	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
	PrivateStaticIpAddress pulumi.StringPtrInput
	// A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
	RedisConfiguration CacheRedisConfigurationPtrInput
	// The name of the resource group in which to
	// create the Redis instance.
	ResourceGroupName pulumi.StringInput
	// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
	ShardCount pulumi.IntPtrInput
	// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
	SkuName pulumi.StringInput
	// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
	Zones pulumi.StringPtrInput
}

func (CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheArgs)(nil)).Elem()
}
