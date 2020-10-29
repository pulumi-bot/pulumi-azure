# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['EndpointStorageContainer']


class EndpointStorageContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_frequency_in_seconds: Optional[pulumi.Input[int]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 container_name: Optional[pulumi.Input[str]] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 file_name_format: Optional[pulumi.Input[str]] = None,
                 iothub_name: Optional[pulumi.Input[str]] = None,
                 max_chunk_size_in_bytes: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an IotHub Storage Container Endpoint

        > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] batch_frequency_in_seconds: Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] container_name: The name of storage container in the storage account.
               *
        :param pulumi.Input[str] encoding: Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'.
        :param pulumi.Input[str] file_name_format: File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered.
        :param pulumi.Input[str] iothub_name: The name of the IoTHub to which this Storage Container Endpoint belongs. Changing this forces a new resource to be created.
        :param pulumi.Input[int] max_chunk_size_in_bytes: Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
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

            __props__['batch_frequency_in_seconds'] = batch_frequency_in_seconds
            if connection_string is None:
                raise TypeError("Missing required property 'connection_string'")
            __props__['connection_string'] = connection_string
            if container_name is None:
                raise TypeError("Missing required property 'container_name'")
            __props__['container_name'] = container_name
            __props__['encoding'] = encoding
            __props__['file_name_format'] = file_name_format
            if iothub_name is None:
                raise TypeError("Missing required property 'iothub_name'")
            __props__['iothub_name'] = iothub_name
            __props__['max_chunk_size_in_bytes'] = max_chunk_size_in_bytes
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(EndpointStorageContainer, __self__).__init__(
            'azure:iot/endpointStorageContainer:EndpointStorageContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            batch_frequency_in_seconds: Optional[pulumi.Input[int]] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            container_name: Optional[pulumi.Input[str]] = None,
            encoding: Optional[pulumi.Input[str]] = None,
            file_name_format: Optional[pulumi.Input[str]] = None,
            iothub_name: Optional[pulumi.Input[str]] = None,
            max_chunk_size_in_bytes: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'EndpointStorageContainer':
        """
        Get an existing EndpointStorageContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] batch_frequency_in_seconds: Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] container_name: The name of storage container in the storage account.
               *
        :param pulumi.Input[str] encoding: Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'.
        :param pulumi.Input[str] file_name_format: File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered.
        :param pulumi.Input[str] iothub_name: The name of the IoTHub to which this Storage Container Endpoint belongs. Changing this forces a new resource to be created.
        :param pulumi.Input[int] max_chunk_size_in_bytes: Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["batch_frequency_in_seconds"] = batch_frequency_in_seconds
        __props__["connection_string"] = connection_string
        __props__["container_name"] = container_name
        __props__["encoding"] = encoding
        __props__["file_name_format"] = file_name_format
        __props__["iothub_name"] = iothub_name
        __props__["max_chunk_size_in_bytes"] = max_chunk_size_in_bytes
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        return EndpointStorageContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="batchFrequencyInSeconds")
    def batch_frequency_in_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        """
        return pulumi.get(self, "batch_frequency_in_seconds")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        The connection string for the endpoint.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Output[str]:
        """
        The name of storage container in the storage account.
        *
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter
    def encoding(self) -> pulumi.Output[Optional[str]]:
        """
        Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fileNameFormat")
    def file_name_format(self) -> pulumi.Output[Optional[str]]:
        """
        File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered.
        """
        return pulumi.get(self, "file_name_format")

    @property
    @pulumi.getter(name="iothubName")
    def iothub_name(self) -> pulumi.Output[str]:
        """
        The name of the IoTHub to which this Storage Container Endpoint belongs. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "iothub_name")

    @property
    @pulumi.getter(name="maxChunkSizeInBytes")
    def max_chunk_size_in_bytes(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        """
        return pulumi.get(self, "max_chunk_size_in_bytes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

