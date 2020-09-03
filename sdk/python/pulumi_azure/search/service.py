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

__all__ = ['Service']


class Service(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[float]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 replica_count: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Search Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.search.Service("exampleService",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku="standard")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
        :param pulumi.Input[float] partition_count: The number of partitions which should be created.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this resource. Defaults to `true`.
        :param pulumi.Input[float] replica_count: The number of replica's which should be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
        :param pulumi.Input[str] sku: The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Search Service.
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

            __props__['location'] = location
            __props__['name'] = name
            __props__['partition_count'] = partition_count
            __props__['public_network_access_enabled'] = public_network_access_enabled
            __props__['replica_count'] = replica_count
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['primary_key'] = None
            __props__['query_keys'] = None
            __props__['secondary_key'] = None
        super(Service, __self__).__init__(
            'azure:search/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition_count: Optional[pulumi.Input[float]] = None,
            primary_key: Optional[pulumi.Input[str]] = None,
            public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
            query_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceQueryKeyArgs']]]]] = None,
            replica_count: Optional[pulumi.Input[float]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
        :param pulumi.Input[float] partition_count: The number of partitions which should be created.
        :param pulumi.Input[str] primary_key: The Primary Key used for Search Service Administration.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this resource. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceQueryKeyArgs']]]] query_keys: A `query_keys` block as defined below.
        :param pulumi.Input[float] replica_count: The number of replica's which should be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
        :param pulumi.Input[str] secondary_key: The Secondary Key used for Search Service Administration.
        :param pulumi.Input[str] sku: The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Search Service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["partition_count"] = partition_count
        __props__["primary_key"] = primary_key
        __props__["public_network_access_enabled"] = public_network_access_enabled
        __props__["query_keys"] = query_keys
        __props__["replica_count"] = replica_count
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_key"] = secondary_key
        __props__["sku"] = sku
        __props__["tags"] = tags
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> pulumi.Output[float]:
        """
        The number of partitions which should be created.
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> pulumi.Output[str]:
        """
        The Primary Key used for Search Service Administration.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not public network access is allowed for this resource. Defaults to `true`.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @property
    @pulumi.getter(name="queryKeys")
    def query_keys(self) -> pulumi.Output[Sequence['outputs.ServiceQueryKey']]:
        """
        A `query_keys` block as defined below.
        """
        return pulumi.get(self, "query_keys")

    @property
    @pulumi.getter(name="replicaCount")
    def replica_count(self) -> pulumi.Output[float]:
        """
        The number of replica's which should be created.
        """
        return pulumi.get(self, "replica_count")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[str]:
        """
        The Secondary Key used for Search Service Administration.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[str]:
        """
        The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Search Service.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

