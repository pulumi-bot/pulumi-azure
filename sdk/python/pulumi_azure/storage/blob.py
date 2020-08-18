# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Blob']


class Blob(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_tier: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parallelism: Optional[pulumi.Input[float]] = None,
                 size: Optional[pulumi.Input[float]] = None,
                 source: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
                 source_content: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_container_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Blob within a Storage Container.

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
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_blob = azure.storage.Blob("exampleBlob",
            storage_account_name=example_account.name,
            storage_container_name=example_container.name,
            type="Block",
            source=pulumi.FileAsset("some-local-file.zip"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_tier: The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
        :param pulumi.Input[str] content_type: The content type of the storage blob. Cannot be defined if `source_uri` is defined. Defaults to `application/octet-stream`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A map of custom blob metadata.
        :param pulumi.Input[str] name: The name of the storage blob. Must be unique within the storage container the blob is located.
        :param pulumi.Input[float] parallelism: The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
        :param pulumi.Input[float] size: Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] source: An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `source_content` or `source_uri` is specified.
        :param pulumi.Input[str] source_content: The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `source_uri` is specified.
        :param pulumi.Input[str] source_uri: The URI of an existing blob, or a file in the Azure File service, to use as the source contents
               for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `source_content` is specified.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage container.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_name: The name of the storage container in which this blob should be created.
        :param pulumi.Input[str] type: The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
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

            __props__['access_tier'] = access_tier
            __props__['content_type'] = content_type
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['parallelism'] = parallelism
            __props__['size'] = size
            __props__['source'] = source
            __props__['source_content'] = source_content
            __props__['source_uri'] = source_uri
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if storage_container_name is None:
                raise TypeError("Missing required property 'storage_container_name'")
            __props__['storage_container_name'] = storage_container_name
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['url'] = None
        super(Blob, __self__).__init__(
            'azure:storage/blob:Blob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_tier: Optional[pulumi.Input[str]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parallelism: Optional[pulumi.Input[float]] = None,
            size: Optional[pulumi.Input[float]] = None,
            source: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
            source_content: Optional[pulumi.Input[str]] = None,
            source_uri: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            storage_container_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Blob':
        """
        Get an existing Blob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_tier: The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
        :param pulumi.Input[str] content_type: The content type of the storage blob. Cannot be defined if `source_uri` is defined. Defaults to `application/octet-stream`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A map of custom blob metadata.
        :param pulumi.Input[str] name: The name of the storage blob. Must be unique within the storage container the blob is located.
        :param pulumi.Input[float] parallelism: The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
        :param pulumi.Input[float] size: Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] source: An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `source_content` or `source_uri` is specified.
        :param pulumi.Input[str] source_content: The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `source_uri` is specified.
        :param pulumi.Input[str] source_uri: The URI of an existing blob, or a file in the Azure File service, to use as the source contents
               for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `source_content` is specified.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage container.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_name: The name of the storage container in which this blob should be created.
        :param pulumi.Input[str] type: The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] url: The URL of the blob
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_tier"] = access_tier
        __props__["content_type"] = content_type
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["parallelism"] = parallelism
        __props__["size"] = size
        __props__["source"] = source
        __props__["source_content"] = source_content
        __props__["source_uri"] = source_uri
        __props__["storage_account_name"] = storage_account_name
        __props__["storage_container_name"] = storage_container_name
        __props__["type"] = type
        __props__["url"] = url
        return Blob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessTier")
    def access_tier(self) -> str:
        """
        The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
        """
        return pulumi.get(self, "access_tier")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        """
        The content type of the storage blob. Cannot be defined if `source_uri` is defined. Defaults to `application/octet-stream`.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, str]:
        """
        A map of custom blob metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the storage blob. Must be unique within the storage container the blob is located.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parallelism(self) -> Optional[float]:
        """
        The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
        """
        return pulumi.get(self, "parallelism")

    @property
    @pulumi.getter
    def size(self) -> Optional[float]:
        """
        Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def source(self) -> Optional[Union[pulumi.Asset, pulumi.Archive]]:
        """
        An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `source_content` or `source_uri` is specified.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceContent")
    def source_content(self) -> Optional[str]:
        """
        The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `source_uri` is specified.
        """
        return pulumi.get(self, "source_content")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> Optional[str]:
        """
        The URI of an existing blob, or a file in the Azure File service, to use as the source contents
        for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `source_content` is specified.
        """
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> str:
        """
        Specifies the storage account in which to create the storage container.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @property
    @pulumi.getter(name="storageContainerName")
    def storage_container_name(self) -> str:
        """
        The name of the storage container in which this blob should be created.
        """
        return pulumi.get(self, "storage_container_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL of the blob
        """
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

