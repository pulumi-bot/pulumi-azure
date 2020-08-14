# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class DatasetJson(pulumi.CustomResource):
    additional_properties: pulumi.Output[dict]
    """
    A map of additional properties to associate with the Data Factory Dataset.
    """
    annotations: pulumi.Output[list]
    """
    List of tags that can be used for describing the Data Factory Dataset.
    """
    azure_blob_storage_location: pulumi.Output[dict]
    """
    A `azure_blob_storage_location` block as defined below.

      * `container` (`str`) - The container on the Azure Blob Storage Account hosting the file.
      * `filename` (`str`) - The filename of the file on the web server.
      * `path` (`str`) - The folder path to the file on the web server.
    """
    data_factory_name: pulumi.Output[str]
    """
    The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
    """
    description: pulumi.Output[str]
    """
    The description for the Data Factory Dataset.
    """
    encoding: pulumi.Output[str]
    """
    The encoding format for the file.
    """
    folder: pulumi.Output[str]
    """
    The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
    """
    http_server_location: pulumi.Output[dict]
    """
    A `http_server_location` block as defined below.

      * `filename` (`str`) - The filename of the file on the web server.
      * `path` (`str`) - The folder path to the file on the web server.
      * `relative_url` (`str`) - The base URL to the web server hosting the file.
    """
    linked_service_name: pulumi.Output[str]
    """
    The Data Factory Linked Service name in which to associate the Dataset with.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
    """
    parameters: pulumi.Output[dict]
    """
    A map of parameters to associate with the Data Factory Dataset.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
    """
    schema_columns: pulumi.Output[list]
    """
    A `schema_column` block as defined below.

      * `description` (`str`) - The description of the column.
      * `name` (`str`) - The name of the column.
      * `type` (`str`) - Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
    """
    def __init__(__self__, resource_name, opts=None, additional_properties=None, annotations=None, azure_blob_storage_location=None, data_factory_name=None, description=None, encoding=None, folder=None, http_server_location=None, linked_service_name=None, name=None, parameters=None, resource_group_name=None, schema_columns=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure JSON Dataset inside an Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_web = azure.datafactory.LinkedServiceWeb("exampleLinkedServiceWeb",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            authentication_type="Anonymous",
            url="https://www.bing.com")
        example_dataset_json = azure.datafactory.DatasetJson("exampleDatasetJson",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            linked_service_name=example_linked_service_web.name,
            http_server_location={
                "relative_url": "/fizz/buzz/",
                "path": "foo/bar/",
                "filename": "foo.txt",
            },
            encoding="UTF-8")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] additional_properties: A map of additional properties to associate with the Data Factory Dataset.
        :param pulumi.Input[list] annotations: List of tags that can be used for describing the Data Factory Dataset.
        :param pulumi.Input[dict] azure_blob_storage_location: A `azure_blob_storage_location` block as defined below.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset.
        :param pulumi.Input[str] encoding: The encoding format for the file.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[dict] http_server_location: A `http_server_location` block as defined below.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[dict] parameters: A map of parameters to associate with the Data Factory Dataset.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        :param pulumi.Input[list] schema_columns: A `schema_column` block as defined below.

        The **azure_blob_storage_location** object supports the following:

          * `container` (`pulumi.Input[str]`) - The container on the Azure Blob Storage Account hosting the file.
          * `filename` (`pulumi.Input[str]`) - The filename of the file on the web server.
          * `path` (`pulumi.Input[str]`) - The folder path to the file on the web server.

        The **http_server_location** object supports the following:

          * `filename` (`pulumi.Input[str]`) - The filename of the file on the web server.
          * `path` (`pulumi.Input[str]`) - The folder path to the file on the web server.
          * `relative_url` (`pulumi.Input[str]`) - The base URL to the web server hosting the file.

        The **schema_columns** object supports the following:

          * `description` (`pulumi.Input[str]`) - The description of the column.
          * `name` (`pulumi.Input[str]`) - The name of the column.
          * `type` (`pulumi.Input[str]`) - Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
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
            if data_factory_name is None:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['encoding'] = encoding
            __props__['folder'] = folder
            __props__['http_server_location'] = http_server_location
            if linked_service_name is None:
                raise TypeError("Missing required property 'linked_service_name'")
            __props__['linked_service_name'] = linked_service_name
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['schema_columns'] = schema_columns
        super(DatasetJson, __self__).__init__(
            'azure:datafactory/datasetJson:DatasetJson',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, additional_properties=None, annotations=None, azure_blob_storage_location=None, data_factory_name=None, description=None, encoding=None, folder=None, http_server_location=None, linked_service_name=None, name=None, parameters=None, resource_group_name=None, schema_columns=None):
        """
        Get an existing DatasetJson resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] additional_properties: A map of additional properties to associate with the Data Factory Dataset.
        :param pulumi.Input[list] annotations: List of tags that can be used for describing the Data Factory Dataset.
        :param pulumi.Input[dict] azure_blob_storage_location: A `azure_blob_storage_location` block as defined below.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset.
        :param pulumi.Input[str] encoding: The encoding format for the file.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[dict] http_server_location: A `http_server_location` block as defined below.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[dict] parameters: A map of parameters to associate with the Data Factory Dataset.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        :param pulumi.Input[list] schema_columns: A `schema_column` block as defined below.

        The **azure_blob_storage_location** object supports the following:

          * `container` (`pulumi.Input[str]`) - The container on the Azure Blob Storage Account hosting the file.
          * `filename` (`pulumi.Input[str]`) - The filename of the file on the web server.
          * `path` (`pulumi.Input[str]`) - The folder path to the file on the web server.

        The **http_server_location** object supports the following:

          * `filename` (`pulumi.Input[str]`) - The filename of the file on the web server.
          * `path` (`pulumi.Input[str]`) - The folder path to the file on the web server.
          * `relative_url` (`pulumi.Input[str]`) - The base URL to the web server hosting the file.

        The **schema_columns** object supports the following:

          * `description` (`pulumi.Input[str]`) - The description of the column.
          * `name` (`pulumi.Input[str]`) - The name of the column.
          * `type` (`pulumi.Input[str]`) - Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_properties"] = additional_properties
        __props__["annotations"] = annotations
        __props__["azure_blob_storage_location"] = azure_blob_storage_location
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["encoding"] = encoding
        __props__["folder"] = folder
        __props__["http_server_location"] = http_server_location
        __props__["linked_service_name"] = linked_service_name
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["schema_columns"] = schema_columns
        return DatasetJson(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
