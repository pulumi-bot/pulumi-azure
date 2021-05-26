# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    A collection of values returned by getDatabase.
    """
    def __init__(__self__, collation=None, elastic_pool_id=None, id=None, license_type=None, max_size_gb=None, name=None, read_replica_count=None, read_scale=None, server_id=None, sku_name=None, storage_account_type=None, tags=None, zone_redundant=None):
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if elastic_pool_id and not isinstance(elastic_pool_id, str):
            raise TypeError("Expected argument 'elastic_pool_id' to be a str")
        pulumi.set(__self__, "elastic_pool_id", elastic_pool_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if license_type and not isinstance(license_type, str):
            raise TypeError("Expected argument 'license_type' to be a str")
        pulumi.set(__self__, "license_type", license_type)
        if max_size_gb and not isinstance(max_size_gb, int):
            raise TypeError("Expected argument 'max_size_gb' to be a int")
        pulumi.set(__self__, "max_size_gb", max_size_gb)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if read_replica_count and not isinstance(read_replica_count, int):
            raise TypeError("Expected argument 'read_replica_count' to be a int")
        pulumi.set(__self__, "read_replica_count", read_replica_count)
        if read_scale and not isinstance(read_scale, bool):
            raise TypeError("Expected argument 'read_scale' to be a bool")
        pulumi.set(__self__, "read_scale", read_scale)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        pulumi.set(__self__, "sku_name", sku_name)
        if storage_account_type and not isinstance(storage_account_type, str):
            raise TypeError("Expected argument 'storage_account_type' to be a str")
        pulumi.set(__self__, "storage_account_type", storage_account_type)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        pulumi.set(__self__, "zone_redundant", zone_redundant)

    @property
    @pulumi.getter
    def collation(self) -> str:
        """
        The collation of the database.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="elasticPoolId")
    def elastic_pool_id(self) -> str:
        """
        The id of the elastic pool containing this database.
        """
        return pulumi.get(self, "elastic_pool_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> str:
        """
        The license type to apply for this database.
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter(name="maxSizeGb")
    def max_size_gb(self) -> int:
        """
        The max size of the database in gigabytes.
        """
        return pulumi.get(self, "max_size_gb")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readReplicaCount")
    def read_replica_count(self) -> int:
        """
        The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed.
        """
        return pulumi.get(self, "read_replica_count")

    @property
    @pulumi.getter(name="readScale")
    def read_scale(self) -> bool:
        """
        If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica.
        """
        return pulumi.get(self, "read_scale")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> str:
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        """
        The name of the sku of the database.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="storageAccountType")
    def storage_account_type(self) -> str:
        """
        The storage account type used to store backups for this database.
        """
        return pulumi.get(self, "storage_account_type")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="zoneRedundant")
    def zone_redundant(self) -> bool:
        """
        Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        """
        return pulumi.get(self, "zone_redundant")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            collation=self.collation,
            elastic_pool_id=self.elastic_pool_id,
            id=self.id,
            license_type=self.license_type,
            max_size_gb=self.max_size_gb,
            name=self.name,
            read_replica_count=self.read_replica_count,
            read_scale=self.read_scale,
            server_id=self.server_id,
            sku_name=self.sku_name,
            storage_account_type=self.storage_account_type,
            tags=self.tags,
            zone_redundant=self.zone_redundant)


def get_database(name: Optional[str] = None,
                 server_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Use this data source to access information about an existing SQL database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.mssql.get_database(name="example-mssql-db",
        server_id="example-mssql-server-id")
    pulumi.export("databaseId", example.id)
    ```


    :param str name: The name of the Ms SQL Database.
    :param str server_id: The id of the Ms SQL Server on which to create the database.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serverId'] = server_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:mssql/getDatabase:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        collation=__ret__.collation,
        elastic_pool_id=__ret__.elastic_pool_id,
        id=__ret__.id,
        license_type=__ret__.license_type,
        max_size_gb=__ret__.max_size_gb,
        name=__ret__.name,
        read_replica_count=__ret__.read_replica_count,
        read_scale=__ret__.read_scale,
        server_id=__ret__.server_id,
        sku_name=__ret__.sku_name,
        storage_account_type=__ret__.storage_account_type,
        tags=__ret__.tags,
        zone_redundant=__ret__.zone_redundant)


@_utilities.lift_output_func(get_database)
def get_database_apply(name: Optional[pulumi.Input[str]] = None,
                       server_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    ...
