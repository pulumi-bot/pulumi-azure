# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

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
    def __init__(__self__, collation=None, default_secondary_location=None, edition=None, elastic_pool_name=None, failover_group_id=None, id=None, location=None, name=None, read_scale=None, resource_group_name=None, server_name=None, tags=None):
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if default_secondary_location and not isinstance(default_secondary_location, str):
            raise TypeError("Expected argument 'default_secondary_location' to be a str")
        pulumi.set(__self__, "default_secondary_location", default_secondary_location)
        if edition and not isinstance(edition, str):
            raise TypeError("Expected argument 'edition' to be a str")
        pulumi.set(__self__, "edition", edition)
        if elastic_pool_name and not isinstance(elastic_pool_name, str):
            raise TypeError("Expected argument 'elastic_pool_name' to be a str")
        pulumi.set(__self__, "elastic_pool_name", elastic_pool_name)
        if failover_group_id and not isinstance(failover_group_id, str):
            raise TypeError("Expected argument 'failover_group_id' to be a str")
        pulumi.set(__self__, "failover_group_id", failover_group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if read_scale and not isinstance(read_scale, bool):
            raise TypeError("Expected argument 'read_scale' to be a bool")
        pulumi.set(__self__, "read_scale", read_scale)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if server_name and not isinstance(server_name, str):
            raise TypeError("Expected argument 'server_name' to be a str")
        pulumi.set(__self__, "server_name", server_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def collation(self) -> str:
        """
        The name of the collation.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="defaultSecondaryLocation")
    def default_secondary_location(self) -> str:
        """
        The default secondary location of the SQL Database.
        """
        return pulumi.get(self, "default_secondary_location")

    @property
    @pulumi.getter
    def edition(self) -> str:
        """
        The edition of the database.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="elasticPoolName")
    def elastic_pool_name(self) -> str:
        """
        The name of the elastic database pool the database belongs to.
        """
        return pulumi.get(self, "elastic_pool_name")

    @property
    @pulumi.getter(name="failoverGroupId")
    def failover_group_id(self) -> str:
        """
        The ID of the failover group the database belongs to.
        """
        return pulumi.get(self, "failover_group_id")

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
        The location of the Resource Group in which the SQL Server exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the database.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readScale")
    def read_scale(self) -> bool:
        """
        Indicate if read-only connections will be redirected to a high-available replica.
        """
        return pulumi.get(self, "read_scale")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which the database resides. This will always be the same resource group as the Database Server.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> str:
        """
        The name of the SQL Server on which to create the database.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            collation=self.collation,
            default_secondary_location=self.default_secondary_location,
            edition=self.edition,
            elastic_pool_name=self.elastic_pool_name,
            failover_group_id=self.failover_group_id,
            id=self.id,
            location=self.location,
            name=self.name,
            read_scale=self.read_scale,
            resource_group_name=self.resource_group_name,
            server_name=self.server_name,
            tags=self.tags)


def get_database(name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 server_name: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Use this data source to access information about an existing SQL Azure Database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.sql.get_database(name="example_db",
        server_name="example_db_server",
        resource_group_name="example-resources")
    pulumi.export("sqlDatabaseId", example.id)
    ```


    :param str name: The name of the SQL Database.
    :param str resource_group_name: Specifies the name of the Resource Group where the Azure SQL Database exists.
    :param str server_name: The name of the SQL Server.
    :param Mapping[str, str] tags: A mapping of tags assigned to the resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:sql/getDatabase:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        collation=__ret__.collation,
        default_secondary_location=__ret__.default_secondary_location,
        edition=__ret__.edition,
        elastic_pool_name=__ret__.elastic_pool_name,
        failover_group_id=__ret__.failover_group_id,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        read_scale=__ret__.read_scale,
        resource_group_name=__ret__.resource_group_name,
        server_name=__ret__.server_name,
        tags=__ret__.tags)
