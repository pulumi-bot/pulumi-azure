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

__all__ = ['DatasetDelimitedText']


class DatasetDelimitedText(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 azure_blob_storage_location: Optional[pulumi.Input[pulumi.InputType['DatasetDelimitedTextAzureBlobStorageLocationArgs']]] = None,
                 column_delimiter: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 escape_character: Optional[pulumi.Input[str]] = None,
                 first_row_as_header: Optional[pulumi.Input[bool]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 http_server_location: Optional[pulumi.Input[pulumi.InputType['DatasetDelimitedTextHttpServerLocationArgs']]] = None,
                 linked_service_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 null_value: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 quote_character: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 row_delimiter: Optional[pulumi.Input[str]] = None,
                 schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetDelimitedTextSchemaColumnArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Delimited Text Dataset inside an Azure Data Factory.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Dataset.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Dataset.
        :param pulumi.Input[pulumi.InputType['DatasetDelimitedTextAzureBlobStorageLocationArgs']] azure_blob_storage_location: A `azure_blob_storage_location` block as defined below.
        :param pulumi.Input[str] column_delimiter: The column delimiter.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset.
        :param pulumi.Input[str] encoding: The encoding format for the file.
        :param pulumi.Input[str] escape_character: The escape character.
        :param pulumi.Input[bool] first_row_as_header: When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[pulumi.InputType['DatasetDelimitedTextHttpServerLocationArgs']] http_server_location: A `http_server_location` block as defined below.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] null_value: The null value string.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Dataset.
        :param pulumi.Input[str] quote_character: The quote character.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        :param pulumi.Input[str] row_delimiter: The row delimiter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetDelimitedTextSchemaColumnArgs']]]] schema_columns: A `schema_column` block as defined below.
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

            __props__['additional_properties'] = additional_properties
            __props__['annotations'] = annotations
            __props__['azure_blob_storage_location'] = azure_blob_storage_location
            __props__['column_delimiter'] = column_delimiter
            if data_factory_name is None:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['encoding'] = encoding
            __props__['escape_character'] = escape_character
            __props__['first_row_as_header'] = first_row_as_header
            __props__['folder'] = folder
            __props__['http_server_location'] = http_server_location
            if linked_service_name is None:
                raise TypeError("Missing required property 'linked_service_name'")
            __props__['linked_service_name'] = linked_service_name
            __props__['name'] = name
            __props__['null_value'] = null_value
            __props__['parameters'] = parameters
            __props__['quote_character'] = quote_character
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['row_delimiter'] = row_delimiter
            __props__['schema_columns'] = schema_columns
        super(DatasetDelimitedText, __self__).__init__(
            'azure:datafactory/datasetDelimitedText:DatasetDelimitedText',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            azure_blob_storage_location: Optional[pulumi.Input[pulumi.InputType['DatasetDelimitedTextAzureBlobStorageLocationArgs']]] = None,
            column_delimiter: Optional[pulumi.Input[str]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encoding: Optional[pulumi.Input[str]] = None,
            escape_character: Optional[pulumi.Input[str]] = None,
            first_row_as_header: Optional[pulumi.Input[bool]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            http_server_location: Optional[pulumi.Input[pulumi.InputType['DatasetDelimitedTextHttpServerLocationArgs']]] = None,
            linked_service_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            null_value: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            quote_character: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            row_delimiter: Optional[pulumi.Input[str]] = None,
            schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetDelimitedTextSchemaColumnArgs']]]]] = None) -> 'DatasetDelimitedText':
        """
        Get an existing DatasetDelimitedText resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Dataset.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Dataset.
        :param pulumi.Input[pulumi.InputType['DatasetDelimitedTextAzureBlobStorageLocationArgs']] azure_blob_storage_location: A `azure_blob_storage_location` block as defined below.
        :param pulumi.Input[str] column_delimiter: The column delimiter.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset.
        :param pulumi.Input[str] encoding: The encoding format for the file.
        :param pulumi.Input[str] escape_character: The escape character.
        :param pulumi.Input[bool] first_row_as_header: When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[pulumi.InputType['DatasetDelimitedTextHttpServerLocationArgs']] http_server_location: A `http_server_location` block as defined below.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] null_value: The null value string.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Dataset.
        :param pulumi.Input[str] quote_character: The quote character.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        :param pulumi.Input[str] row_delimiter: The row delimiter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetDelimitedTextSchemaColumnArgs']]]] schema_columns: A `schema_column` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_properties"] = additional_properties
        __props__["annotations"] = annotations
        __props__["azure_blob_storage_location"] = azure_blob_storage_location
        __props__["column_delimiter"] = column_delimiter
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["encoding"] = encoding
        __props__["escape_character"] = escape_character
        __props__["first_row_as_header"] = first_row_as_header
        __props__["folder"] = folder
        __props__["http_server_location"] = http_server_location
        __props__["linked_service_name"] = linked_service_name
        __props__["name"] = name
        __props__["null_value"] = null_value
        __props__["parameters"] = parameters
        __props__["quote_character"] = quote_character
        __props__["resource_group_name"] = resource_group_name
        __props__["row_delimiter"] = row_delimiter
        __props__["schema_columns"] = schema_columns
        return DatasetDelimitedText(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of additional properties to associate with the Data Factory Dataset.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Dataset.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="azureBlobStorageLocation")
    def azure_blob_storage_location(self) -> pulumi.Output[Optional['outputs.DatasetDelimitedTextAzureBlobStorageLocation']]:
        """
        A `azure_blob_storage_location` block as defined below.
        """
        return pulumi.get(self, "azure_blob_storage_location")

    @property
    @pulumi.getter(name="columnDelimiter")
    def column_delimiter(self) -> pulumi.Output[Optional[str]]:
        """
        The column delimiter.
        """
        return pulumi.get(self, "column_delimiter")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Output[str]:
        """
        The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the Data Factory Dataset.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encoding(self) -> pulumi.Output[Optional[str]]:
        """
        The encoding format for the file.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="escapeCharacter")
    def escape_character(self) -> pulumi.Output[Optional[str]]:
        """
        The escape character.
        """
        return pulumi.get(self, "escape_character")

    @property
    @pulumi.getter(name="firstRowAsHeader")
    def first_row_as_header(self) -> pulumi.Output[Optional[bool]]:
        """
        When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
        """
        return pulumi.get(self, "first_row_as_header")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[Optional[str]]:
        """
        The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="httpServerLocation")
    def http_server_location(self) -> pulumi.Output[Optional['outputs.DatasetDelimitedTextHttpServerLocation']]:
        """
        A `http_server_location` block as defined below.
        """
        return pulumi.get(self, "http_server_location")

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> pulumi.Output[str]:
        """
        The Data Factory Linked Service name in which to associate the Dataset with.
        """
        return pulumi.get(self, "linked_service_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nullValue")
    def null_value(self) -> pulumi.Output[Optional[str]]:
        """
        The null value string.
        """
        return pulumi.get(self, "null_value")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of parameters to associate with the Data Factory Dataset.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="quoteCharacter")
    def quote_character(self) -> pulumi.Output[Optional[str]]:
        """
        The quote character.
        """
        return pulumi.get(self, "quote_character")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="rowDelimiter")
    def row_delimiter(self) -> pulumi.Output[Optional[str]]:
        """
        The row delimiter.
        """
        return pulumi.get(self, "row_delimiter")

    @property
    @pulumi.getter(name="schemaColumns")
    def schema_columns(self) -> pulumi.Output[Optional[Sequence['outputs.DatasetDelimitedTextSchemaColumn']]]:
        """
        A `schema_column` block as defined below.
        """
        return pulumi.get(self, "schema_columns")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

