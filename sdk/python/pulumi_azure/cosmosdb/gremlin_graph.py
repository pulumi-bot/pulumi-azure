# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GremlinGraph']


class GremlinGraph(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 autoscale_settings: Optional[pulumi.Input[pulumi.InputType['GremlinGraphAutoscaleSettingsArgs']]] = None,
                 conflict_resolution_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphConflictResolutionPolicyArgs']]]]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 index_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphIndexPolicyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_key_path: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 throughput: Optional[pulumi.Input[int]] = None,
                 unique_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphUniqueKeyArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Gremlin Graph within a Cosmos DB Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_account = azure.cosmosdb.get_account(name="tfex-cosmosdb-account",
            resource_group_name="tfex-cosmosdb-account-rg")
        example_gremlin_database = azure.cosmosdb.GremlinDatabase("exampleGremlinDatabase",
            resource_group_name=example_account.resource_group_name,
            account_name=example_account.name)
        example_gremlin_graph = azure.cosmosdb.GremlinGraph("exampleGremlinGraph",
            resource_group_name=azurerm_cosmosdb_account["example"]["resource_group_name"],
            account_name=azurerm_cosmosdb_account["example"]["name"],
            database_name=example_gremlin_database.name,
            partition_key_path="/Example",
            throughput=400,
            index_policies=[azure.cosmosdb.GremlinGraphIndexPolicyArgs(
                automatic=True,
                indexing_mode="Consistent",
                included_paths=["/*"],
                excluded_paths=["/\"_etag\"/?"],
            )],
            conflict_resolution_policies=[azure.cosmosdb.GremlinGraphConflictResolutionPolicyArgs(
                mode="LastWriterWins",
                conflict_resolution_path="/_ts",
            )],
            unique_keys=[azure.cosmosdb.GremlinGraphUniqueKeyArgs(
                paths=[
                    "/definition/id1",
                    "/definition/id2",
                ],
            )])
        ```

        > **NOTE:** The CosmosDB Account needs to have the `EnableGremlin` capability enabled to use this resource - which can be done by adding this to the `capabilities` list within the `cosmosdb.Account` resource.

        ## Import

        Cosmos Gremlin Graphs can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['GremlinGraphAutoscaleSettingsArgs']] autoscale_settings: An `autoscale_settings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partition_key_path` to be set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphConflictResolutionPolicyArgs']]]] conflict_resolution_policies: The conflict resolution policy for the graph. One or more `conflict_resolution_policy` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphIndexPolicyArgs']]]] index_policies: The configuration of the indexing policy. One or more `index_policy` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partition_key_path: Define a partition key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] throughput: The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphUniqueKeyArgs']]]] unique_keys: One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['autoscale_settings'] = autoscale_settings
            if conflict_resolution_policies is None:
                raise TypeError("Missing required property 'conflict_resolution_policies'")
            __props__['conflict_resolution_policies'] = conflict_resolution_policies
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            if index_policies is None:
                raise TypeError("Missing required property 'index_policies'")
            __props__['index_policies'] = index_policies
            __props__['name'] = name
            __props__['partition_key_path'] = partition_key_path
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['throughput'] = throughput
            __props__['unique_keys'] = unique_keys
        super(GremlinGraph, __self__).__init__(
            'azure:cosmosdb/gremlinGraph:GremlinGraph',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            autoscale_settings: Optional[pulumi.Input[pulumi.InputType['GremlinGraphAutoscaleSettingsArgs']]] = None,
            conflict_resolution_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphConflictResolutionPolicyArgs']]]]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            index_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphIndexPolicyArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition_key_path: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            throughput: Optional[pulumi.Input[int]] = None,
            unique_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphUniqueKeyArgs']]]]] = None) -> 'GremlinGraph':
        """
        Get an existing GremlinGraph resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['GremlinGraphAutoscaleSettingsArgs']] autoscale_settings: An `autoscale_settings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partition_key_path` to be set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphConflictResolutionPolicyArgs']]]] conflict_resolution_policies: The conflict resolution policy for the graph. One or more `conflict_resolution_policy` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphIndexPolicyArgs']]]] index_policies: The configuration of the indexing policy. One or more `index_policy` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partition_key_path: Define a partition key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] throughput: The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GremlinGraphUniqueKeyArgs']]]] unique_keys: One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["autoscale_settings"] = autoscale_settings
        __props__["conflict_resolution_policies"] = conflict_resolution_policies
        __props__["database_name"] = database_name
        __props__["index_policies"] = index_policies
        __props__["name"] = name
        __props__["partition_key_path"] = partition_key_path
        __props__["resource_group_name"] = resource_group_name
        __props__["throughput"] = throughput
        __props__["unique_keys"] = unique_keys
        return GremlinGraph(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="autoscaleSettings")
    def autoscale_settings(self) -> pulumi.Output[Optional['outputs.GremlinGraphAutoscaleSettings']]:
        """
        An `autoscale_settings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partition_key_path` to be set.
        """
        return pulumi.get(self, "autoscale_settings")

    @property
    @pulumi.getter(name="conflictResolutionPolicies")
    def conflict_resolution_policies(self) -> pulumi.Output[Sequence['outputs.GremlinGraphConflictResolutionPolicy']]:
        """
        The conflict resolution policy for the graph. One or more `conflict_resolution_policy` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "conflict_resolution_policies")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="indexPolicies")
    def index_policies(self) -> pulumi.Output[Sequence['outputs.GremlinGraphIndexPolicy']]:
        """
        The configuration of the indexing policy. One or more `index_policy` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "index_policies")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionKeyPath")
    def partition_key_path(self) -> pulumi.Output[Optional[str]]:
        """
        Define a partition key. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "partition_key_path")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def throughput(self) -> pulumi.Output[int]:
        """
        The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
        """
        return pulumi.get(self, "throughput")

    @property
    @pulumi.getter(name="uniqueKeys")
    def unique_keys(self) -> pulumi.Output[Optional[Sequence['outputs.GremlinGraphUniqueKey']]]:
        """
        One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "unique_keys")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

