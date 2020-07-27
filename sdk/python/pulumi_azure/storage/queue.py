# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Queue']


class Queue(pulumi.CustomResource):
    metadata: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("metadata")
    """
    A mapping of MetaData which should be assigned to this Storage Queue.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
    """
    storage_account_name: pulumi.Output[str] = pulumi.output_property("storageAccountName")
    """
    Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, metadata: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, storage_account_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Queue within an Azure Storage Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_queue = azure.storage.Queue("exampleQueue", storage_account_name=example_account.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] metadata: A mapping of MetaData which should be assigned to this Storage Queue.
        :param pulumi.Input[str] name: The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
        :param pulumi.Input[str] storage_account_name: Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
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

            __props__['metadata'] = metadata
            __props__['name'] = name
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
        super(Queue, __self__).__init__(
            'azure:storage/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, metadata: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, storage_account_name: Optional[pulumi.Input[str]] = None) -> 'Queue':
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] metadata: A mapping of MetaData which should be assigned to this Storage Queue.
        :param pulumi.Input[str] name: The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
        :param pulumi.Input[str] storage_account_name: Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["storage_account_name"] = storage_account_name
        return Queue(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

