# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ZipBlob']

warnings.warn("""ZipBlob resource is deprecated in the 2.0 version of the provider. Use Blob resource instead.""", DeprecationWarning)


class ZipBlob(pulumi.CustomResource):
    warnings.warn("""ZipBlob resource is deprecated in the 2.0 version of the provider. Use Blob resource instead.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_tier: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[pulumi.Archive]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parallelism: Optional[pulumi.Input[int]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 source_content: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_container_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ZipBlob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        pulumi.log.warn("ZipBlob is deprecated: ZipBlob resource is deprecated in the 2.0 version of the provider. Use Blob resource instead.")
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
            __props__['content'] = content
            __props__['content_type'] = content_type
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['parallelism'] = parallelism
            __props__['size'] = size
            __props__['source_content'] = source_content
            __props__['source_uri'] = source_uri
            if storage_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if storage_container_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_container_name'")
            __props__['storage_container_name'] = storage_container_name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['url'] = None
        super(ZipBlob, __self__).__init__(
            'azure:storage/zipBlob:ZipBlob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_tier: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[pulumi.Archive]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parallelism: Optional[pulumi.Input[int]] = None,
            size: Optional[pulumi.Input[int]] = None,
            source_content: Optional[pulumi.Input[str]] = None,
            source_uri: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            storage_container_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ZipBlob':
        """
        Get an existing ZipBlob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_tier"] = access_tier
        __props__["content"] = content
        __props__["content_type"] = content_type
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["parallelism"] = parallelism
        __props__["size"] = size
        __props__["source_content"] = source_content
        __props__["source_uri"] = source_uri
        __props__["storage_account_name"] = storage_account_name
        __props__["storage_container_name"] = storage_container_name
        __props__["type"] = type
        __props__["url"] = url
        return ZipBlob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessTier")
    def access_tier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access_tier")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[pulumi.Archive]]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parallelism(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "parallelism")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sourceContent")
    def source_content(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "source_content")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "storage_account_name")

    @property
    @pulumi.getter(name="storageContainerName")
    def storage_container_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "storage_container_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

