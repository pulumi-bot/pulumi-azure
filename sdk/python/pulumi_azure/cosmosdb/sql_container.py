# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SqlContainer(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
    """
    database_name: pulumi.Output[str]
    """
    The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
    """
    default_ttl: pulumi.Output[float]
    """
    The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
    """
    partition_key_path: pulumi.Output[str]
    """
    Define a partition key. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
    """
    throughput: pulumi.Output[float]
    """
    The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
    """
    unique_keys: pulumi.Output[list]
    """
    One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.

      * `paths` (`list`) - A list of paths to use for this unique key.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, database_name=None, default_ttl=None, name=None, partition_key_path=None, resource_group_name=None, throughput=None, unique_keys=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a SQL Container within a Cosmos DB Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.cosmosdb.SqlContainer("example",
            resource_group_name=azurerm_cosmosdb_account["example"]["resource_group_name"],
            account_name=azurerm_cosmosdb_account["example"]["name"],
            database_name=azurerm_cosmosdb_sql_database["example"]["name"],
            partition_key_path="/definition/id",
            throughput=400,
            unique_key=[{
                "paths": [
                    "/definition/idlong",
                    "/definition/idshort",
                ],
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[float] default_ttl: The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partition_key_path: Define a partition key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
        :param pulumi.Input[float] throughput: The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
        :param pulumi.Input[list] unique_keys: One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.

        The **unique_keys** object supports the following:

          * `paths` (`pulumi.Input[list]`) - A list of paths to use for this unique key.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['default_ttl'] = default_ttl
            __props__['name'] = name
            __props__['partition_key_path'] = partition_key_path
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['throughput'] = throughput
            __props__['unique_keys'] = unique_keys
        super(SqlContainer, __self__).__init__(
            'azure:cosmosdb/sqlContainer:SqlContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, database_name=None, default_ttl=None, name=None, partition_key_path=None, resource_group_name=None, throughput=None, unique_keys=None):
        """
        Get an existing SqlContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[float] default_ttl: The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partition_key_path: Define a partition key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
        :param pulumi.Input[float] throughput: The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
        :param pulumi.Input[list] unique_keys: One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.

        The **unique_keys** object supports the following:

          * `paths` (`pulumi.Input[list]`) - A list of paths to use for this unique key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["database_name"] = database_name
        __props__["default_ttl"] = default_ttl
        __props__["name"] = name
        __props__["partition_key_path"] = partition_key_path
        __props__["resource_group_name"] = resource_group_name
        __props__["throughput"] = throughput
        __props__["unique_keys"] = unique_keys
        return SqlContainer(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
