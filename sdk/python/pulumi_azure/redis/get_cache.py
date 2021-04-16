# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetCacheResult',
    'AwaitableGetCacheResult',
    'get_cache',
]

@pulumi.output_type
class GetCacheResult:
    """
    A collection of values returned by getCache.
    """
    def __init__(__self__, capacity=None, enable_non_ssl_port=None, family=None, hostname=None, id=None, location=None, minimum_tls_version=None, name=None, patch_schedules=None, port=None, primary_access_key=None, primary_connection_string=None, private_static_ip_address=None, redis_configurations=None, resource_group_name=None, secondary_access_key=None, secondary_connection_string=None, shard_count=None, sku_name=None, ssl_port=None, subnet_id=None, tags=None, zones=None):
        if capacity and not isinstance(capacity, int):
            raise TypeError("Expected argument 'capacity' to be a int")
        pulumi.set(__self__, "capacity", capacity)
        if enable_non_ssl_port and not isinstance(enable_non_ssl_port, bool):
            raise TypeError("Expected argument 'enable_non_ssl_port' to be a bool")
        pulumi.set(__self__, "enable_non_ssl_port", enable_non_ssl_port)
        if family and not isinstance(family, str):
            raise TypeError("Expected argument 'family' to be a str")
        pulumi.set(__self__, "family", family)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if minimum_tls_version and not isinstance(minimum_tls_version, str):
            raise TypeError("Expected argument 'minimum_tls_version' to be a str")
        pulumi.set(__self__, "minimum_tls_version", minimum_tls_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if patch_schedules and not isinstance(patch_schedules, list):
            raise TypeError("Expected argument 'patch_schedules' to be a list")
        pulumi.set(__self__, "patch_schedules", patch_schedules)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        pulumi.set(__self__, "primary_access_key", primary_access_key)
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        pulumi.set(__self__, "primary_connection_string", primary_connection_string)
        if private_static_ip_address and not isinstance(private_static_ip_address, str):
            raise TypeError("Expected argument 'private_static_ip_address' to be a str")
        pulumi.set(__self__, "private_static_ip_address", private_static_ip_address)
        if redis_configurations and not isinstance(redis_configurations, list):
            raise TypeError("Expected argument 'redis_configurations' to be a list")
        pulumi.set(__self__, "redis_configurations", redis_configurations)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        pulumi.set(__self__, "secondary_access_key", secondary_access_key)
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        pulumi.set(__self__, "secondary_connection_string", secondary_connection_string)
        if shard_count and not isinstance(shard_count, int):
            raise TypeError("Expected argument 'shard_count' to be a int")
        pulumi.set(__self__, "shard_count", shard_count)
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        pulumi.set(__self__, "sku_name", sku_name)
        if ssl_port and not isinstance(ssl_port, int):
            raise TypeError("Expected argument 'ssl_port' to be a int")
        pulumi.set(__self__, "ssl_port", ssl_port)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def capacity(self) -> int:
        """
        The size of the Redis Cache deployed.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="enableNonSslPort")
    def enable_non_ssl_port(self) -> bool:
        """
        Whether the SSL port is enabled.
        """
        return pulumi.get(self, "enable_non_ssl_port")

    @property
    @pulumi.getter
    def family(self) -> str:
        """
        The SKU family/pricing group used. Possible values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        The Hostname of the Redis Instance
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the Redis Cache.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="minimumTlsVersion")
    def minimum_tls_version(self) -> str:
        """
        The minimum TLS version.
        """
        return pulumi.get(self, "minimum_tls_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="patchSchedules")
    def patch_schedules(self) -> Sequence['outputs.GetCachePatchScheduleResult']:
        """
        A list of `patch_schedule` blocks as defined below.
        """
        return pulumi.get(self, "patch_schedules")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The non-SSL Port of the Redis Instance
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> str:
        """
        The Primary Access Key for the Redis Instance
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> str:
        """
        The primary connection string of the Redis Instance.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="privateStaticIpAddress")
    def private_static_ip_address(self) -> str:
        return pulumi.get(self, "private_static_ip_address")

    @property
    @pulumi.getter(name="redisConfigurations")
    def redis_configurations(self) -> Sequence['outputs.GetCacheRedisConfigurationResult']:
        """
        A `redis_configuration` block as defined below.
        """
        return pulumi.get(self, "redis_configurations")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> str:
        """
        The Secondary Access Key for the Redis Instance
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> str:
        """
        The secondary connection string of the Redis Instance.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> int:
        return pulumi.get(self, "shard_count")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        """
        The SKU of Redis used. Possible values are `Basic`, `Standard` and `Premium`.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="sslPort")
    def ssl_port(self) -> int:
        """
        The SSL Port of the Redis Instance
        """
        return pulumi.get(self, "ssl_port")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[str]:
        return pulumi.get(self, "zones")


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


def get_cache(name: Optional[str] = None,
              resource_group_name: Optional[str] = None,
              zones: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCacheResult:
    """
    Use this data source to access information about an existing Redis Cache

    ## Example Usage

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
    __ret__ = pulumi.runtime.invoke('azure:redis/getCache:getCache', __args__, opts=opts, typ=GetCacheResult).value

    return AwaitableGetCacheResult(
        capacity=__ret__.capacity,
        enable_non_ssl_port=__ret__.enable_non_ssl_port,
        family=__ret__.family,
        hostname=__ret__.hostname,
        id=__ret__.id,
        location=__ret__.location,
        minimum_tls_version=__ret__.minimum_tls_version,
        name=__ret__.name,
        patch_schedules=__ret__.patch_schedules,
        port=__ret__.port,
        primary_access_key=__ret__.primary_access_key,
        primary_connection_string=__ret__.primary_connection_string,
        private_static_ip_address=__ret__.private_static_ip_address,
        redis_configurations=__ret__.redis_configurations,
        resource_group_name=__ret__.resource_group_name,
        secondary_access_key=__ret__.secondary_access_key,
        secondary_connection_string=__ret__.secondary_connection_string,
        shard_count=__ret__.shard_count,
        sku_name=__ret__.sku_name,
        ssl_port=__ret__.ssl_port,
        subnet_id=__ret__.subnet_id,
        tags=__ret__.tags,
        zones=__ret__.zones)
