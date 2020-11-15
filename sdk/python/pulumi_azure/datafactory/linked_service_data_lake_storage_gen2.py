# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['LinkedServiceDataLakeStorageGen2']


class LinkedServiceDataLakeStorageGen2(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 service_principal_key: Optional[pulumi.Input[str]] = None,
                 tenant: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 use_managed_identity: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linked Service (connection) between Data Lake Storage Gen2 and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        current = azure.core.get_client_config()
        example_linked_service_data_lake_storage_gen2 = azure.datafactory.LinkedServiceDataLakeStorageGen2("exampleLinkedServiceDataLakeStorageGen2",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            service_principal_id=current.client_id,
            service_principal_key="exampleKey",
            tenant="11111111-1111-1111-1111-111111111111",
            url="https://datalakestoragegen2")
        ```

        ## Import

        Data Factory Data Lake Storage Gen2 Linked Services can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] service_principal_id: The service principal id in which to authenticate against the Azure Data Lake Storage Gen2 account. Required if `use_managed_identity` is true.
        :param pulumi.Input[str] service_principal_key: The service principal key in which to authenticate against the Azure Data Lake Storage Gen2 account.  Required if `use_managed_identity` is true.
        :param pulumi.Input[str] tenant: (Required) The tenant id or name in which to authenticate against the Azure Data Lake Storage Gen2 account.
        :param pulumi.Input[str] url: The endpoint for the Azure Data Lake Storage Gen2 service.
        :param pulumi.Input[bool] use_managed_identity: Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `service_principal_id` and `service_principal_key`
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
            __props__['service_principal_id'] = service_principal_id
            __props__['service_principal_key'] = service_principal_key
            __props__['tenant'] = tenant
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['use_managed_identity'] = use_managed_identity
        super(LinkedServiceDataLakeStorageGen2, __self__).__init__(
            'azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            integration_runtime_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            service_principal_key: Optional[pulumi.Input[str]] = None,
            tenant: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            use_managed_identity: Optional[pulumi.Input[bool]] = None) -> 'LinkedServiceDataLakeStorageGen2':
        """
        Get an existing LinkedServiceDataLakeStorageGen2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] service_principal_id: The service principal id in which to authenticate against the Azure Data Lake Storage Gen2 account. Required if `use_managed_identity` is true.
        :param pulumi.Input[str] service_principal_key: The service principal key in which to authenticate against the Azure Data Lake Storage Gen2 account.  Required if `use_managed_identity` is true.
        :param pulumi.Input[str] tenant: (Required) The tenant id or name in which to authenticate against the Azure Data Lake Storage Gen2 account.
        :param pulumi.Input[str] url: The endpoint for the Azure Data Lake Storage Gen2 service.
        :param pulumi.Input[bool] use_managed_identity: Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `service_principal_id` and `service_principal_key`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_properties"] = additional_properties
        __props__["annotations"] = annotations
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["integration_runtime_name"] = integration_runtime_name
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["service_principal_id"] = service_principal_id
        __props__["service_principal_key"] = service_principal_key
        __props__["tenant"] = tenant
        __props__["url"] = url
        __props__["use_managed_identity"] = use_managed_identity
        return LinkedServiceDataLakeStorageGen2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Output[str]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the Data Factory Linked Service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> pulumi.Output[Optional[str]]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "integration_runtime_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of parameters to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[Optional[str]]:
        """
        The service principal id in which to authenticate against the Azure Data Lake Storage Gen2 account. Required if `use_managed_identity` is true.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="servicePrincipalKey")
    def service_principal_key(self) -> pulumi.Output[Optional[str]]:
        """
        The service principal key in which to authenticate against the Azure Data Lake Storage Gen2 account.  Required if `use_managed_identity` is true.
        """
        return pulumi.get(self, "service_principal_key")

    @property
    @pulumi.getter
    def tenant(self) -> pulumi.Output[Optional[str]]:
        """
        (Required) The tenant id or name in which to authenticate against the Azure Data Lake Storage Gen2 account.
        """
        return pulumi.get(self, "tenant")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The endpoint for the Azure Data Lake Storage Gen2 service.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="useManagedIdentity")
    def use_managed_identity(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `service_principal_id` and `service_principal_key`
        """
        return pulumi.get(self, "use_managed_identity")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

