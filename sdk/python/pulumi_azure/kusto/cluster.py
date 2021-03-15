# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 sku: pulumi.Input['ClusterSkuArgs'],
                 double_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 enable_disk_encryption: Optional[pulumi.Input[bool]] = None,
                 enable_purge: Optional[pulumi.Input[bool]] = None,
                 enable_streaming_ingest: Optional[pulumi.Input[bool]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input['ClusterIdentityArgs']] = None,
                 language_extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 optimized_auto_scale: Optional[pulumi.Input['ClusterOptimizedAutoScaleArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trusted_external_tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_network_configuration: Optional[pulumi.Input['ClusterVirtualNetworkConfigurationArgs']] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input['ClusterSkuArgs'] sku: A `sku` block as defined below.
        :param pulumi.Input[bool] double_encryption_enabled: Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_disk_encryption: Specifies if the cluster's disks are encrypted.
        :param pulumi.Input[bool] enable_purge: Specifies if the purge operations are enabled.
        :param pulumi.Input[bool] enable_streaming_ingest: Specifies if the streaming ingest is enabled.
        :param pulumi.Input[str] engine: . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
        :param pulumi.Input['ClusterIdentityArgs'] identity: An identity block.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] language_extensions: An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        :param pulumi.Input[str] location: The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input['ClusterOptimizedAutoScaleArgs'] optimized_auto_scale: An `optimized_auto_scale` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_external_tenants: Specifies a list of tenant IDs that are trusted by the cluster.
        :param pulumi.Input['ClusterVirtualNetworkConfigurationArgs'] virtual_network_configuration: A `virtual_network_configuration` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if double_encryption_enabled is not None:
            pulumi.set(__self__, "double_encryption_enabled", double_encryption_enabled)
        if enable_disk_encryption is not None:
            pulumi.set(__self__, "enable_disk_encryption", enable_disk_encryption)
        if enable_purge is not None:
            pulumi.set(__self__, "enable_purge", enable_purge)
        if enable_streaming_ingest is not None:
            pulumi.set(__self__, "enable_streaming_ingest", enable_streaming_ingest)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if language_extensions is not None:
            pulumi.set(__self__, "language_extensions", language_extensions)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if optimized_auto_scale is not None:
            pulumi.set(__self__, "optimized_auto_scale", optimized_auto_scale)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trusted_external_tenants is not None:
            pulumi.set(__self__, "trusted_external_tenants", trusted_external_tenants)
        if virtual_network_configuration is not None:
            pulumi.set(__self__, "virtual_network_configuration", virtual_network_configuration)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input['ClusterSkuArgs']:
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['ClusterSkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="doubleEncryptionEnabled")
    def double_encryption_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "double_encryption_enabled")

    @double_encryption_enabled.setter
    def double_encryption_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "double_encryption_enabled", value)

    @property
    @pulumi.getter(name="enableDiskEncryption")
    def enable_disk_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the cluster's disks are encrypted.
        """
        return pulumi.get(self, "enable_disk_encryption")

    @enable_disk_encryption.setter
    def enable_disk_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_disk_encryption", value)

    @property
    @pulumi.getter(name="enablePurge")
    def enable_purge(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the purge operations are enabled.
        """
        return pulumi.get(self, "enable_purge")

    @enable_purge.setter
    def enable_purge(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_purge", value)

    @property
    @pulumi.getter(name="enableStreamingIngest")
    def enable_streaming_ingest(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the streaming ingest is enabled.
        """
        return pulumi.get(self, "enable_streaming_ingest")

    @enable_streaming_ingest.setter
    def enable_streaming_ingest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_streaming_ingest", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ClusterIdentityArgs']]:
        """
        An identity block.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ClusterIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="languageExtensions")
    def language_extensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        """
        return pulumi.get(self, "language_extensions")

    @language_extensions.setter
    def language_extensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "language_extensions", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="optimizedAutoScale")
    def optimized_auto_scale(self) -> Optional[pulumi.Input['ClusterOptimizedAutoScaleArgs']]:
        """
        An `optimized_auto_scale` block as defined below.
        """
        return pulumi.get(self, "optimized_auto_scale")

    @optimized_auto_scale.setter
    def optimized_auto_scale(self, value: Optional[pulumi.Input['ClusterOptimizedAutoScaleArgs']]):
        pulumi.set(self, "optimized_auto_scale", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trustedExternalTenants")
    def trusted_external_tenants(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of tenant IDs that are trusted by the cluster.
        """
        return pulumi.get(self, "trusted_external_tenants")

    @trusted_external_tenants.setter
    def trusted_external_tenants(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "trusted_external_tenants", value)

    @property
    @pulumi.getter(name="virtualNetworkConfiguration")
    def virtual_network_configuration(self) -> Optional[pulumi.Input['ClusterVirtualNetworkConfigurationArgs']]:
        """
        A `virtual_network_configuration` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_configuration")

    @virtual_network_configuration.setter
    def virtual_network_configuration(self, value: Optional[pulumi.Input['ClusterVirtualNetworkConfigurationArgs']]):
        pulumi.set(self, "virtual_network_configuration", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "zones", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Cluster

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        example = azure.kusto.Cluster("example",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ),
            tags={
                "Environment": "Production",
            })
        ```

        ## Import

        Kusto Clusters can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:kusto/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 double_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 enable_disk_encryption: Optional[pulumi.Input[bool]] = None,
                 enable_purge: Optional[pulumi.Input[bool]] = None,
                 enable_streaming_ingest: Optional[pulumi.Input[bool]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['ClusterIdentityArgs']]] = None,
                 language_extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 optimized_auto_scale: Optional[pulumi.Input[pulumi.InputType['ClusterOptimizedAutoScaleArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ClusterSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trusted_external_tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_network_configuration: Optional[pulumi.Input[pulumi.InputType['ClusterVirtualNetworkConfigurationArgs']]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Cluster

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        example = azure.kusto.Cluster("example",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ),
            tags={
                "Environment": "Production",
            })
        ```

        ## Import

        Kusto Clusters can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:kusto/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] double_encryption_enabled: Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_disk_encryption: Specifies if the cluster's disks are encrypted.
        :param pulumi.Input[bool] enable_purge: Specifies if the purge operations are enabled.
        :param pulumi.Input[bool] enable_streaming_ingest: Specifies if the streaming ingest is enabled.
        :param pulumi.Input[str] engine: . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
        :param pulumi.Input[pulumi.InputType['ClusterIdentityArgs']] identity: An identity block.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] language_extensions: An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        :param pulumi.Input[str] location: The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ClusterOptimizedAutoScaleArgs']] optimized_auto_scale: An `optimized_auto_scale` block as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ClusterSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_external_tenants: Specifies a list of tenant IDs that are trusted by the cluster.
        :param pulumi.Input[pulumi.InputType['ClusterVirtualNetworkConfigurationArgs']] virtual_network_configuration: A `virtual_network_configuration` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 double_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 enable_disk_encryption: Optional[pulumi.Input[bool]] = None,
                 enable_purge: Optional[pulumi.Input[bool]] = None,
                 enable_streaming_ingest: Optional[pulumi.Input[bool]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['ClusterIdentityArgs']]] = None,
                 language_extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 optimized_auto_scale: Optional[pulumi.Input[pulumi.InputType['ClusterOptimizedAutoScaleArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ClusterSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trusted_external_tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_network_configuration: Optional[pulumi.Input[pulumi.InputType['ClusterVirtualNetworkConfigurationArgs']]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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

            __props__['double_encryption_enabled'] = double_encryption_enabled
            __props__['enable_disk_encryption'] = enable_disk_encryption
            __props__['enable_purge'] = enable_purge
            __props__['enable_streaming_ingest'] = enable_streaming_ingest
            __props__['engine'] = engine
            __props__['identity'] = identity
            __props__['language_extensions'] = language_extensions
            __props__['location'] = location
            __props__['name'] = name
            __props__['optimized_auto_scale'] = optimized_auto_scale
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['trusted_external_tenants'] = trusted_external_tenants
            __props__['virtual_network_configuration'] = virtual_network_configuration
            __props__['zones'] = zones
            __props__['data_ingestion_uri'] = None
            __props__['uri'] = None
        super(Cluster, __self__).__init__(
            'azure:kusto/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_ingestion_uri: Optional[pulumi.Input[str]] = None,
            double_encryption_enabled: Optional[pulumi.Input[bool]] = None,
            enable_disk_encryption: Optional[pulumi.Input[bool]] = None,
            enable_purge: Optional[pulumi.Input[bool]] = None,
            enable_streaming_ingest: Optional[pulumi.Input[bool]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['ClusterIdentityArgs']]] = None,
            language_extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            optimized_auto_scale: Optional[pulumi.Input[pulumi.InputType['ClusterOptimizedAutoScaleArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[pulumi.InputType['ClusterSkuArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            trusted_external_tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            uri: Optional[pulumi.Input[str]] = None,
            virtual_network_configuration: Optional[pulumi.Input[pulumi.InputType['ClusterVirtualNetworkConfigurationArgs']]] = None,
            zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_ingestion_uri: The Kusto Cluster URI to be used for data ingestion.
        :param pulumi.Input[bool] double_encryption_enabled: Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_disk_encryption: Specifies if the cluster's disks are encrypted.
        :param pulumi.Input[bool] enable_purge: Specifies if the purge operations are enabled.
        :param pulumi.Input[bool] enable_streaming_ingest: Specifies if the streaming ingest is enabled.
        :param pulumi.Input[str] engine: . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
        :param pulumi.Input[pulumi.InputType['ClusterIdentityArgs']] identity: An identity block.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] language_extensions: An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        :param pulumi.Input[str] location: The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ClusterOptimizedAutoScaleArgs']] optimized_auto_scale: An `optimized_auto_scale` block as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ClusterSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_external_tenants: Specifies a list of tenant IDs that are trusted by the cluster.
        :param pulumi.Input[str] uri: The FQDN of the Azure Kusto Cluster.
        :param pulumi.Input[pulumi.InputType['ClusterVirtualNetworkConfigurationArgs']] virtual_network_configuration: A `virtual_network_configuration` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_ingestion_uri"] = data_ingestion_uri
        __props__["double_encryption_enabled"] = double_encryption_enabled
        __props__["enable_disk_encryption"] = enable_disk_encryption
        __props__["enable_purge"] = enable_purge
        __props__["enable_streaming_ingest"] = enable_streaming_ingest
        __props__["engine"] = engine
        __props__["identity"] = identity
        __props__["language_extensions"] = language_extensions
        __props__["location"] = location
        __props__["name"] = name
        __props__["optimized_auto_scale"] = optimized_auto_scale
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["trusted_external_tenants"] = trusted_external_tenants
        __props__["uri"] = uri
        __props__["virtual_network_configuration"] = virtual_network_configuration
        __props__["zones"] = zones
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataIngestionUri")
    def data_ingestion_uri(self) -> pulumi.Output[str]:
        """
        The Kusto Cluster URI to be used for data ingestion.
        """
        return pulumi.get(self, "data_ingestion_uri")

    @property
    @pulumi.getter(name="doubleEncryptionEnabled")
    def double_encryption_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "double_encryption_enabled")

    @property
    @pulumi.getter(name="enableDiskEncryption")
    def enable_disk_encryption(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the cluster's disks are encrypted.
        """
        return pulumi.get(self, "enable_disk_encryption")

    @property
    @pulumi.getter(name="enablePurge")
    def enable_purge(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the purge operations are enabled.
        """
        return pulumi.get(self, "enable_purge")

    @property
    @pulumi.getter(name="enableStreamingIngest")
    def enable_streaming_ingest(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the streaming ingest is enabled.
        """
        return pulumi.get(self, "enable_streaming_ingest")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[Optional[str]]:
        """
        . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output['outputs.ClusterIdentity']:
        """
        An identity block.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="languageExtensions")
    def language_extensions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        """
        return pulumi.get(self, "language_extensions")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="optimizedAutoScale")
    def optimized_auto_scale(self) -> pulumi.Output[Optional['outputs.ClusterOptimizedAutoScale']]:
        """
        An `optimized_auto_scale` block as defined below.
        """
        return pulumi.get(self, "optimized_auto_scale")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.ClusterSku']:
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trustedExternalTenants")
    def trusted_external_tenants(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies a list of tenant IDs that are trusted by the cluster.
        """
        return pulumi.get(self, "trusted_external_tenants")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        The FQDN of the Azure Kusto Cluster.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter(name="virtualNetworkConfiguration")
    def virtual_network_configuration(self) -> pulumi.Output[Optional['outputs.ClusterVirtualNetworkConfiguration']]:
        """
        A `virtual_network_configuration` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_configuration")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

