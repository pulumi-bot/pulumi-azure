# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class MongoCollection(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    database_name: pulumi.Output[str]
    """
    The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
    """
    default_ttl_seconds: pulumi.Output[float]
    """
    The default Time To Live in seconds. If the value is `0` items are not automatically expired.
    """
    indices: pulumi.Output[list]
    """
    One or more `index` blocks as defined below.

      * `keys` (`list`) - Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
      * `unique` (`bool`) - Is the index unique or not? Defaults to `false`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
    """
    shard_key: pulumi.Output[str]
    """
    The name of the key to partition on for sharding. There must not be any other unique index keys.
    """
    system_indexes: pulumi.Output[list]
    """
    One or more `system_indexes` blocks as defined below.

      * `keys` (`list`) - Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
      * `unique` (`bool`) - Is the index unique or not? Defaults to `false`.
    """
    throughput: pulumi.Output[float]
    """
    The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, database_name=None, default_ttl_seconds=None, indices=None, name=None, resource_group_name=None, shard_key=None, throughput=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Mongo Collection within a Cosmos DB Account.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_account = azure.cosmosdb.get_account(name="tfex-cosmosdb-account",
            resource_group_name="tfex-cosmosdb-account-rg")
        example_mongo_database = azure.cosmosdb.MongoDatabase("exampleMongoDatabase",
            resource_group_name=example_account.resource_group_name,
            account_name=example_account.name)
        example_mongo_collection = azure.cosmosdb.MongoCollection("exampleMongoCollection",
            resource_group_name=example_account.resource_group_name,
            account_name=example_account.name,
            database_name=example_mongo_database.name,
            default_ttl_seconds="777",
            shard_key="uniqueKey",
            throughput=400)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[float] default_ttl_seconds: The default Time To Live in seconds. If the value is `0` items are not automatically expired.
        :param pulumi.Input[list] indices: One or more `index` blocks as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] shard_key: The name of the key to partition on for sharding. There must not be any other unique index keys.
        :param pulumi.Input[float] throughput: The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.

        The **indices** object supports the following:

          * `keys` (`pulumi.Input[list]`) - Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
          * `unique` (`pulumi.Input[bool]`) - Is the index unique or not? Defaults to `false`.
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
            __props__['default_ttl_seconds'] = default_ttl_seconds
            __props__['indices'] = indices
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['shard_key'] = shard_key
            __props__['throughput'] = throughput
            __props__['system_indexes'] = None
        super(MongoCollection, __self__).__init__(
            'azure:cosmosdb/mongoCollection:MongoCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, database_name=None, default_ttl_seconds=None, indices=None, name=None, resource_group_name=None, shard_key=None, system_indexes=None, throughput=None):
        """
        Get an existing MongoCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[float] default_ttl_seconds: The default Time To Live in seconds. If the value is `0` items are not automatically expired.
        :param pulumi.Input[list] indices: One or more `index` blocks as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] shard_key: The name of the key to partition on for sharding. There must not be any other unique index keys.
        :param pulumi.Input[list] system_indexes: One or more `system_indexes` blocks as defined below.
        :param pulumi.Input[float] throughput: The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.

        The **indices** object supports the following:

          * `keys` (`pulumi.Input[list]`) - Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
          * `unique` (`pulumi.Input[bool]`) - Is the index unique or not? Defaults to `false`.

        The **system_indexes** object supports the following:

          * `keys` (`pulumi.Input[list]`) - Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
          * `unique` (`pulumi.Input[bool]`) - Is the index unique or not? Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["database_name"] = database_name
        __props__["default_ttl_seconds"] = default_ttl_seconds
        __props__["indices"] = indices
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["shard_key"] = shard_key
        __props__["system_indexes"] = system_indexes
        __props__["throughput"] = throughput
        return MongoCollection(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

