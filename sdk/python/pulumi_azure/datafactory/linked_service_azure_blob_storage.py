# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['LinkedServiceAzureBlobStorageArgs', 'LinkedServiceAzureBlobStorage']

@pulumi.input_type
class LinkedServiceAzureBlobStorageArgs:
    def __init__(__self__, *,
                 data_factory_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 sas_uri: Optional[pulumi.Input[str]] = None,
                 service_endpoint: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 service_principal_key: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_managed_identity: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a LinkedServiceAzureBlobStorage resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] connection_string: The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] sas_uri: The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        :param pulumi.Input[str] service_endpoint: The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        :param pulumi.Input[str] service_principal_id: The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        :param pulumi.Input[str] service_principal_key: The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        :param pulumi.Input[str] tenant_id: The tenant id or name in which to authenticate against the Azure Blob Storage account.
        :param pulumi.Input[bool] use_managed_identity: Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        """
        pulumi.set(__self__, "data_factory_name", data_factory_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if integration_runtime_name is not None:
            pulumi.set(__self__, "integration_runtime_name", integration_runtime_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if sas_uri is not None:
            pulumi.set(__self__, "sas_uri", sas_uri)
        if service_endpoint is not None:
            pulumi.set(__self__, "service_endpoint", service_endpoint)
        if service_principal_id is not None:
            pulumi.set(__self__, "service_principal_id", service_principal_id)
        if service_principal_key is not None:
            pulumi.set(__self__, "service_principal_key", service_principal_key)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if use_managed_identity is not None:
            pulumi.set(__self__, "use_managed_identity", use_managed_identity)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Input[str]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Data Factory Linked Service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> Optional[pulumi.Input[str]]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "integration_runtime_name")

    @integration_runtime_name.setter
    def integration_runtime_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_runtime_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of parameters to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="sasUri")
    def sas_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        """
        return pulumi.get(self, "sas_uri")

    @sas_uri.setter
    def sas_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sas_uri", value)

    @property
    @pulumi.getter(name="serviceEndpoint")
    def service_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        """
        return pulumi.get(self, "service_endpoint")

    @service_endpoint.setter
    def service_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="servicePrincipalKey")
    def service_principal_key(self) -> Optional[pulumi.Input[str]]:
        """
        The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        """
        return pulumi.get(self, "service_principal_key")

    @service_principal_key.setter
    def service_principal_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_key", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant id or name in which to authenticate against the Azure Blob Storage account.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="useManagedIdentity")
    def use_managed_identity(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        """
        return pulumi.get(self, "use_managed_identity")

    @use_managed_identity.setter
    def use_managed_identity(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_managed_identity", value)


class LinkedServiceAzureBlobStorage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sas_uri: Optional[pulumi.Input[str]] = None,
                 service_endpoint: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 service_principal_key: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_managed_identity: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linked Service (connection) between an Azure Blob Storage Account and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = example_resource_group.name.apply(lambda name: azure.storage.get_account(name="storageaccountname",
            resource_group_name=name))
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_azure_blob_storage = azure.datafactory.LinkedServiceAzureBlobStorage("exampleLinkedServiceAzureBlobStorage",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            connection_string=example_account.primary_connection_string)
        ```

        ## Import

        Data Factory Linked Service's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] connection_string: The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] sas_uri: The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        :param pulumi.Input[str] service_endpoint: The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        :param pulumi.Input[str] service_principal_id: The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        :param pulumi.Input[str] service_principal_key: The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        :param pulumi.Input[str] tenant_id: The tenant id or name in which to authenticate against the Azure Blob Storage account.
        :param pulumi.Input[bool] use_managed_identity: Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkedServiceAzureBlobStorageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Linked Service (connection) between an Azure Blob Storage Account and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = example_resource_group.name.apply(lambda name: azure.storage.get_account(name="storageaccountname",
            resource_group_name=name))
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_azure_blob_storage = azure.datafactory.LinkedServiceAzureBlobStorage("exampleLinkedServiceAzureBlobStorage",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            connection_string=example_account.primary_connection_string)
        ```

        ## Import

        Data Factory Linked Service's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
        ```

        :param str resource_name: The name of the resource.
        :param LinkedServiceAzureBlobStorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkedServiceAzureBlobStorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sas_uri: Optional[pulumi.Input[str]] = None,
                 service_endpoint: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 service_principal_key: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_managed_identity: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__['connection_string'] = connection_string
            if data_factory_name is None and not opts.urn:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['integration_runtime_name'] = integration_runtime_name
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sas_uri'] = sas_uri
            __props__['service_endpoint'] = service_endpoint
            __props__['service_principal_id'] = service_principal_id
            __props__['service_principal_key'] = service_principal_key
            __props__['tenant_id'] = tenant_id
            __props__['use_managed_identity'] = use_managed_identity
        super(LinkedServiceAzureBlobStorage, __self__).__init__(
            'azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            integration_runtime_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sas_uri: Optional[pulumi.Input[str]] = None,
            service_endpoint: Optional[pulumi.Input[str]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            service_principal_key: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            use_managed_identity: Optional[pulumi.Input[bool]] = None) -> 'LinkedServiceAzureBlobStorage':
        """
        Get an existing LinkedServiceAzureBlobStorage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] connection_string: The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] sas_uri: The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        :param pulumi.Input[str] service_endpoint: The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        :param pulumi.Input[str] service_principal_id: The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        :param pulumi.Input[str] service_principal_key: The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        :param pulumi.Input[str] tenant_id: The tenant id or name in which to authenticate against the Azure Blob Storage account.
        :param pulumi.Input[bool] use_managed_identity: Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
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
        __props__["sas_uri"] = sas_uri
        __props__["service_endpoint"] = service_endpoint
        __props__["service_principal_id"] = service_principal_id
        __props__["service_principal_key"] = service_principal_key
        __props__["tenant_id"] = tenant_id
        __props__["use_managed_identity"] = use_managed_identity
        return LinkedServiceAzureBlobStorage(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[Optional[str]]:
        """
        The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        """
        return pulumi.get(self, "connection_string")

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
        Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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
    @pulumi.getter(name="sasUri")
    def sas_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        """
        return pulumi.get(self, "sas_uri")

    @property
    @pulumi.getter(name="serviceEndpoint")
    def service_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        """
        return pulumi.get(self, "service_endpoint")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[Optional[str]]:
        """
        The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="servicePrincipalKey")
    def service_principal_key(self) -> pulumi.Output[Optional[str]]:
        """
        The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        """
        return pulumi.get(self, "service_principal_key")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The tenant id or name in which to authenticate against the Azure Blob Storage account.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="useManagedIdentity")
    def use_managed_identity(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        """
        return pulumi.get(self, "use_managed_identity")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

