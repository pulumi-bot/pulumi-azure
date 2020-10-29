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

__all__ = ['SqlContainer']


class SqlContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 autoscale_settings: Optional[pulumi.Input[pulumi.InputType['SqlContainerAutoscaleSettingsArgs']]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 default_ttl: Optional[pulumi.Input[int]] = None,
                 indexing_policy: Optional[pulumi.Input[pulumi.InputType['SqlContainerIndexingPolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_key_path: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 throughput: Optional[pulumi.Input[int]] = None,
                 unique_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlContainerUniqueKeyArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a SQL Container within a Cosmos DB Account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[int] default_ttl: The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
        :param pulumi.Input[pulumi.InputType['SqlContainerIndexingPolicyArgs']] indexing_policy: An `indexing_policy` block as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partition_key_path: Define a partition key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] throughput: The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlContainerUniqueKeyArgs']]]] unique_keys: One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.
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
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['default_ttl'] = default_ttl
            __props__['indexing_policy'] = indexing_policy
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            autoscale_settings: Optional[pulumi.Input[pulumi.InputType['SqlContainerAutoscaleSettingsArgs']]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            default_ttl: Optional[pulumi.Input[int]] = None,
            indexing_policy: Optional[pulumi.Input[pulumi.InputType['SqlContainerIndexingPolicyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition_key_path: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            throughput: Optional[pulumi.Input[int]] = None,
            unique_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlContainerUniqueKeyArgs']]]]] = None) -> 'SqlContainer':
        """
        Get an existing SqlContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
        :param pulumi.Input[int] default_ttl: The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
        :param pulumi.Input[pulumi.InputType['SqlContainerIndexingPolicyArgs']] indexing_policy: An `indexing_policy` block as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partition_key_path: Define a partition key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
        :param pulumi.Input[int] throughput: The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlContainerUniqueKeyArgs']]]] unique_keys: One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["autoscale_settings"] = autoscale_settings
        __props__["database_name"] = database_name
        __props__["default_ttl"] = default_ttl
        __props__["indexing_policy"] = indexing_policy
        __props__["name"] = name
        __props__["partition_key_path"] = partition_key_path
        __props__["resource_group_name"] = resource_group_name
        __props__["throughput"] = throughput
        __props__["unique_keys"] = unique_keys
        return SqlContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="autoscaleSettings")
    def autoscale_settings(self) -> pulumi.Output[Optional['outputs.SqlContainerAutoscaleSettings']]:
        return pulumi.get(self, "autoscale_settings")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="defaultTtl")
    def default_ttl(self) -> pulumi.Output[int]:
        """
        The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
        """
        return pulumi.get(self, "default_ttl")

    @property
    @pulumi.getter(name="indexingPolicy")
    def indexing_policy(self) -> pulumi.Output['outputs.SqlContainerIndexingPolicy']:
        """
        An `indexing_policy` block as defined below.
        """
        return pulumi.get(self, "indexing_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
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
        The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def throughput(self) -> pulumi.Output[int]:
        """
        The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
        """
        return pulumi.get(self, "throughput")

    @property
    @pulumi.getter(name="uniqueKeys")
    def unique_keys(self) -> pulumi.Output[Optional[Sequence['outputs.SqlContainerUniqueKey']]]:
        """
        One or more `unique_key` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "unique_keys")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

