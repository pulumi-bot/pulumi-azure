# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ElasticPool']


class ElasticPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_dtu_max: Optional[pulumi.Input[float]] = None,
                 db_dtu_min: Optional[pulumi.Input[float]] = None,
                 dtu: Optional[pulumi.Input[float]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_size: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows you to manage an Azure SQL Elastic Pool.

        > **NOTE:** -  This version of the `Elasticpool` resource is being **deprecated** and should no longer be used. Please use the mssql.ElasticPool version instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_sql_server = azure.sql.SqlServer("exampleSqlServer",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            version="12.0",
            administrator_login="4dm1n157r470r",
            administrator_login_password="4-v3ry-53cr37-p455w0rd")
        example_elastic_pool = azure.sql.ElasticPool("exampleElasticPool",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            server_name=example_sql_server.name,
            edition="Basic",
            dtu=50,
            db_dtu_min=0,
            db_dtu_max=5,
            pool_size=5000)
        ```

        > **NOTE on `sql.ElasticPool`:** -  The values of `edition`, `dtu`, and `pool_size` must be consistent with the [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). Any inconsistent argument configuration will be rejected.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] db_dtu_max: The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
        :param pulumi.Input[float] db_dtu_min: The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
        :param pulumi.Input[float] dtu: The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
        :param pulumi.Input[str] edition: The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        :param pulumi.Input[float] pool_size: The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
        :param pulumi.Input[str] server_name: The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['db_dtu_max'] = db_dtu_max
            __props__['db_dtu_min'] = db_dtu_min
            if dtu is None:
                raise TypeError("Missing required property 'dtu'")
            __props__['dtu'] = dtu
            if edition is None:
                raise TypeError("Missing required property 'edition'")
            __props__['edition'] = edition
            __props__['location'] = location
            __props__['name'] = name
            __props__['pool_size'] = pool_size
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['tags'] = tags
            __props__['creation_date'] = None
        super(ElasticPool, __self__).__init__(
            'azure:sql/elasticPool:ElasticPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            db_dtu_max: Optional[pulumi.Input[float]] = None,
            db_dtu_min: Optional[pulumi.Input[float]] = None,
            dtu: Optional[pulumi.Input[float]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_size: Optional[pulumi.Input[float]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ElasticPool':
        """
        Get an existing ElasticPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_date: The creation date of the SQL Elastic Pool.
        :param pulumi.Input[float] db_dtu_max: The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
        :param pulumi.Input[float] db_dtu_min: The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
        :param pulumi.Input[float] dtu: The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
        :param pulumi.Input[str] edition: The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        :param pulumi.Input[float] pool_size: The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
        :param pulumi.Input[str] server_name: The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_date"] = creation_date
        __props__["db_dtu_max"] = db_dtu_max
        __props__["db_dtu_min"] = db_dtu_min
        __props__["dtu"] = dtu
        __props__["edition"] = edition
        __props__["location"] = location
        __props__["name"] = name
        __props__["pool_size"] = pool_size
        __props__["resource_group_name"] = resource_group_name
        __props__["server_name"] = server_name
        __props__["tags"] = tags
        return ElasticPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        The creation date of the SQL Elastic Pool.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="dbDtuMax")
    def db_dtu_max(self) -> pulumi.Output[float]:
        """
        The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
        """
        return pulumi.get(self, "db_dtu_max")

    @property
    @pulumi.getter(name="dbDtuMin")
    def db_dtu_min(self) -> pulumi.Output[float]:
        """
        The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
        """
        return pulumi.get(self, "db_dtu_min")

    @property
    @pulumi.getter
    def dtu(self) -> pulumi.Output[float]:
        """
        The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
        """
        return pulumi.get(self, "dtu")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolSize")
    def pool_size(self) -> pulumi.Output[float]:
        """
        The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
        """
        return pulumi.get(self, "pool_size")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

