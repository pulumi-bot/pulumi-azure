# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['LinkedServicePostgresql']


class LinkedServicePostgresql(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linked Service (connection) between PostgreSQL and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_postgresql = azure.datafactory.LinkedServicePostgresql("exampleLinkedServicePostgresql",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            connection_string="Host=example;Port=5432;Database=example;UID=example;EncryptionMethod=0;Password=example")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[List[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] connection_string: The connection string in which to authenticate with PostgreSQL.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
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
            if connection_string is None:
                raise TypeError("Missing required property 'connection_string'")
            __props__['connection_string'] = connection_string
            if data_factory_name is None:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['integration_runtime_name'] = integration_runtime_name
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(LinkedServicePostgresql, __self__).__init__(
            'azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            integration_runtime_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'LinkedServicePostgresql':
        """
        Get an existing LinkedServicePostgresql resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[List[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] connection_string: The connection string in which to authenticate with PostgreSQL.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_properties"] = additional_properties
        __props__["annotations"] = annotations
        __props__["connection_string"] = connection_string
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["integration_runtime_name"] = integration_runtime_name
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        return LinkedServicePostgresql(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[Mapping[str, str]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter
    def annotations(self) -> Optional[List[str]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        The connection string in which to authenticate with PostgreSQL.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> str:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description for the Data Factory Linked Service PostgreSQL.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> Optional[str]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
        """
        return pulumi.get(self, "integration_runtime_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, str]]:
        """
        A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

