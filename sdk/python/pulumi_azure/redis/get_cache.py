# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetCacheResult:
    """
    A collection of values returned by getCache.
    """
    def __init__(__self__, capacity=None, enable_non_ssl_port=None, family=None, hostname=None, id=None, location=None, minimum_tls_version=None, name=None, patch_schedules=None, port=None, primary_access_key=None, primary_connection_string=None, private_static_ip_address=None, redis_configurations=None, resource_group_name=None, secondary_access_key=None, secondary_connection_string=None, shard_count=None, sku_name=None, ssl_port=None, subnet_id=None, tags=None, zones=None):
        if capacity and not isinstance(capacity, float):
            raise TypeError("Expected argument 'capacity' to be a float")
        __self__.capacity = capacity
        """
        The size of the Redis Cache deployed.
        """
        if enable_non_ssl_port and not isinstance(enable_non_ssl_port, bool):
            raise TypeError("Expected argument 'enable_non_ssl_port' to be a bool")
        __self__.enable_non_ssl_port = enable_non_ssl_port
        """
        Whether the SSL port is enabled.
        """
        if family and not isinstance(family, str):
            raise TypeError("Expected argument 'family' to be a str")
        __self__.family = family
        """
        The SKU family/pricing group used. Possible values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
        """
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        __self__.hostname = hostname
        """
        The Hostname of the Redis Instance
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location of the Redis Cache.
        """
        if minimum_tls_version and not isinstance(minimum_tls_version, str):
            raise TypeError("Expected argument 'minimum_tls_version' to be a str")
        __self__.minimum_tls_version = minimum_tls_version
        """
        The minimum TLS version.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if patch_schedules and not isinstance(patch_schedules, list):
            raise TypeError("Expected argument 'patch_schedules' to be a list")
        __self__.patch_schedules = patch_schedules
        """
        A list of `patch_schedule` blocks as defined below - only available for Premium SKU's.
        """
        if port and not isinstance(port, float):
            raise TypeError("Expected argument 'port' to be a float")
        __self__.port = port
        """
        The non-SSL Port of the Redis Instance
        """
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        __self__.primary_access_key = primary_access_key
        """
        The Primary Access Key for the Redis Instance
        """
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        __self__.primary_connection_string = primary_connection_string
        """
        The primary connection string of the Redis Instance.
        """
        if private_static_ip_address and not isinstance(private_static_ip_address, str):
            raise TypeError("Expected argument 'private_static_ip_address' to be a str")
        __self__.private_static_ip_address = private_static_ip_address
        if redis_configurations and not isinstance(redis_configurations, list):
            raise TypeError("Expected argument 'redis_configurations' to be a list")
        __self__.redis_configurations = redis_configurations
        """
        A `redis_configuration` block as defined below.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        __self__.secondary_access_key = secondary_access_key
        """
        The Secondary Access Key for the Redis Instance
        """
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        __self__.secondary_connection_string = secondary_connection_string
        """
        The secondary connection string of the Redis Instance.
        """
        if shard_count and not isinstance(shard_count, float):
            raise TypeError("Expected argument 'shard_count' to be a float")
        __self__.shard_count = shard_count
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        __self__.sku_name = sku_name
        """
        The SKU of Redis used. Possible values are `Basic`, `Standard` and `Premium`.
        """
        if ssl_port and not isinstance(ssl_port, float):
            raise TypeError("Expected argument 'ssl_port' to be a float")
        __self__.ssl_port = ssl_port
        """
        The SSL Port of the Redis Instance
        """
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        __self__.subnet_id = subnet_id
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones


class AwaitableGetCacheResult(GetCacheResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCacheResult(
            capacity=self.capacity,
            enable_non_ssl_port=self.enable_non_ssl_port,
            family=self.family,
            hostname=self.hostname,
            id=self.id,
            location=self.location,
            minimum_tls_version=self.minimum_tls_version,
            name=self.name,
            patch_schedules=self.patch_schedules,
            port=self.port,
            primary_access_key=self.primary_access_key,
            primary_connection_string=self.primary_connection_string,
            private_static_ip_address=self.private_static_ip_address,
            redis_configurations=self.redis_configurations,
            resource_group_name=self.resource_group_name,
            secondary_access_key=self.secondary_access_key,
            secondary_connection_string=self.secondary_connection_string,
            shard_count=self.shard_count,
            sku_name=self.sku_name,
            ssl_port=self.ssl_port,
            subnet_id=self.subnet_id,
            tags=self.tags,
            zones=self.zones)


def get_cache(name=None, resource_group_name=None, zones=None, opts=None):
    """
    Use this data source to access information about an existing Redis Cache

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.redis.get_cache(name="myrediscache",
        resource_group_name="redis-cache")
    pulumi.export("primaryAccessKey", example.primary_access_key)
    pulumi.export("hostname", example.hostname)
    ```


    :param str name: The name of the Redis cache
    :param str resource_group_name: The name of the resource group the Redis cache instance is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['zones'] = zones
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:redis/getCache:getCache', __args__, opts=opts).value

    return AwaitableGetCacheResult(
        capacity=__ret__.get('capacity'),
        enable_non_ssl_port=__ret__.get('enableNonSslPort'),
        family=__ret__.get('family'),
        hostname=__ret__.get('hostname'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        minimum_tls_version=__ret__.get('minimumTlsVersion'),
        name=__ret__.get('name'),
        patch_schedules=__ret__.get('patchSchedules'),
        port=__ret__.get('port'),
        primary_access_key=__ret__.get('primaryAccessKey'),
        primary_connection_string=__ret__.get('primaryConnectionString'),
        private_static_ip_address=__ret__.get('privateStaticIpAddress'),
        redis_configurations=__ret__.get('redisConfigurations'),
        resource_group_name=__ret__.get('resourceGroupName'),
        secondary_access_key=__ret__.get('secondaryAccessKey'),
        secondary_connection_string=__ret__.get('secondaryConnectionString'),
        shard_count=__ret__.get('shardCount'),
        sku_name=__ret__.get('skuName'),
        ssl_port=__ret__.get('sslPort'),
        subnet_id=__ret__.get('subnetId'),
        tags=__ret__.get('tags'),
        zones=__ret__.get('zones'))
