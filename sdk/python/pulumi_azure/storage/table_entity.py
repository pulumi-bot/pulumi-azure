# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['TableEntity']


class TableEntity(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 partition_key: Optional[pulumi.Input[str]] = None,
                 row_key: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Entity within a Table in an Azure Storage Account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] entity: A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        :param pulumi.Input[str] partition_key: The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] row_key: The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage table entity.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: The name of the storage table in which to create the storage table entity.
               Changing this forces a new resource to be created.
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

            if entity is None:
                raise TypeError("Missing required property 'entity'")
            __props__['entity'] = entity
            if partition_key is None:
                raise TypeError("Missing required property 'partition_key'")
            __props__['partition_key'] = partition_key
            if row_key is None:
                raise TypeError("Missing required property 'row_key'")
            __props__['row_key'] = row_key
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if table_name is None:
                raise TypeError("Missing required property 'table_name'")
            __props__['table_name'] = table_name
        super(TableEntity, __self__).__init__(
            'azure:storage/tableEntity:TableEntity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            entity: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            partition_key: Optional[pulumi.Input[str]] = None,
            row_key: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            table_name: Optional[pulumi.Input[str]] = None) -> 'TableEntity':
        """
        Get an existing TableEntity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] entity: A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        :param pulumi.Input[str] partition_key: The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] row_key: The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage table entity.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: The name of the storage table in which to create the storage table entity.
               Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["entity"] = entity
        __props__["partition_key"] = partition_key
        __props__["row_key"] = row_key
        __props__["storage_account_name"] = storage_account_name
        __props__["table_name"] = table_name
        return TableEntity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def entity(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        """
        return pulumi.get(self, "entity")

    @property
    @pulumi.getter(name="partitionKey")
    def partition_key(self) -> pulumi.Output[str]:
        """
        The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        """
        return pulumi.get(self, "partition_key")

    @property
    @pulumi.getter(name="rowKey")
    def row_key(self) -> pulumi.Output[str]:
        """
        The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        """
        return pulumi.get(self, "row_key")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Output[str]:
        """
        Specifies the storage account in which to create the storage table entity.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[str]:
        """
        The name of the storage table in which to create the storage table entity.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "table_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

