# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['StreamInputBlob']


class StreamInputBlob(pulumi.CustomResource):
    date_format: pulumi.Output[str] = pulumi.property("dateFormat")
    """
    The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the Stream Input Blob. Changing this forces a new resource to be created.
    """

    path_pattern: pulumi.Output[str] = pulumi.property("pathPattern")
    """
    The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
    """

    serialization: pulumi.Output['outputs.StreamInputBlobSerialization'] = pulumi.property("serialization")
    """
    A `serialization` block as defined below.
    """

    storage_account_key: pulumi.Output[str] = pulumi.property("storageAccountKey")
    """
    The Access Key which should be used to connect to this Storage Account.
    """

    storage_account_name: pulumi.Output[str] = pulumi.property("storageAccountName")
    """
    The name of the Storage Account.
    """

    storage_container_name: pulumi.Output[str] = pulumi.property("storageContainerName")
    """
    The name of the Container within the Storage Account.
    """

    stream_analytics_job_name: pulumi.Output[str] = pulumi.property("streamAnalyticsJobName")
    """
    The name of the Stream Analytics Job. Changing this forces a new resource to be created.
    """

    time_format: pulumi.Output[str] = pulumi.property("timeFormat")
    """
    The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 date_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path_pattern: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputBlobSerializationArgs']]] = None,
                 storage_account_key: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_container_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Stream Analytics Stream Input Blob.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=azurerm_resource_group["example"]["name"],
            location=azurerm_resource_group["example"]["location"],
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_stream_input_blob = azure.streamanalytics.StreamInputBlob("exampleStreamInputBlob",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            storage_account_name=example_account.name,
            storage_account_key=example_account.primary_access_key,
            storage_container_name=example_container.name,
            path_pattern="some-random-pattern",
            date_format="yyyy/MM/dd",
            time_format="HH",
            serialization={
                "type": "Json",
                "encoding": "UTF8",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Stream Input Blob. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['StreamInputBlobSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
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

            if date_format is None:
                raise TypeError("Missing required property 'date_format'")
            __props__['date_format'] = date_format
            __props__['name'] = name
            if path_pattern is None:
                raise TypeError("Missing required property 'path_pattern'")
            __props__['path_pattern'] = path_pattern
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if serialization is None:
                raise TypeError("Missing required property 'serialization'")
            __props__['serialization'] = serialization
            if storage_account_key is None:
                raise TypeError("Missing required property 'storage_account_key'")
            __props__['storage_account_key'] = storage_account_key
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if storage_container_name is None:
                raise TypeError("Missing required property 'storage_container_name'")
            __props__['storage_container_name'] = storage_container_name
            if stream_analytics_job_name is None:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
            if time_format is None:
                raise TypeError("Missing required property 'time_format'")
            __props__['time_format'] = time_format
        super(StreamInputBlob, __self__).__init__(
            'azure:streamanalytics/streamInputBlob:StreamInputBlob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            date_format: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path_pattern: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputBlobSerializationArgs']]] = None,
            storage_account_key: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            storage_container_name: Optional[pulumi.Input[str]] = None,
            stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
            time_format: Optional[pulumi.Input[str]] = None) -> 'StreamInputBlob':
        """
        Get an existing StreamInputBlob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Stream Input Blob. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['StreamInputBlobSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["date_format"] = date_format
        __props__["name"] = name
        __props__["path_pattern"] = path_pattern
        __props__["resource_group_name"] = resource_group_name
        __props__["serialization"] = serialization
        __props__["storage_account_key"] = storage_account_key
        __props__["storage_account_name"] = storage_account_name
        __props__["storage_container_name"] = storage_container_name
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        __props__["time_format"] = time_format
        return StreamInputBlob(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

