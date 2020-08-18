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

__all__ = ['CacheNfsTarget']


class CacheNfsTarget(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_junctions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CacheNfsTargetNamespaceJunctionArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 target_host_name: Optional[pulumi.Input[str]] = None,
                 usage_model: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a NFS Target within a HPC Cache.

        > **NOTE:**: By request of the service team the provider no longer automatically registering the `Microsoft.StorageCache` Resource Provider for this resource. To register it you can run `az provider register --namespace 'Microsoft.StorageCache'`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cache_name: The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the HPC Cache NFS Target. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CacheNfsTargetNamespaceJunctionArgs']]]] namespace_junctions: Can be specified multiple times to define multiple `namespace_junction`. Each `namespace_juntion` block supports fields documented below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_host_name: The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] usage_model: The type of usage of the HPC Cache NFS Target.
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

            if cache_name is None:
                raise TypeError("Missing required property 'cache_name'")
            __props__['cache_name'] = cache_name
            __props__['name'] = name
            if namespace_junctions is None:
                raise TypeError("Missing required property 'namespace_junctions'")
            __props__['namespace_junctions'] = namespace_junctions
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if target_host_name is None:
                raise TypeError("Missing required property 'target_host_name'")
            __props__['target_host_name'] = target_host_name
            if usage_model is None:
                raise TypeError("Missing required property 'usage_model'")
            __props__['usage_model'] = usage_model
        super(CacheNfsTarget, __self__).__init__(
            'azure:hpc/cacheNfsTarget:CacheNfsTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            cache_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_junctions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CacheNfsTargetNamespaceJunctionArgs']]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            target_host_name: Optional[pulumi.Input[str]] = None,
            usage_model: Optional[pulumi.Input[str]] = None) -> 'CacheNfsTarget':
        """
        Get an existing CacheNfsTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cache_name: The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the HPC Cache NFS Target. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CacheNfsTargetNamespaceJunctionArgs']]]] namespace_junctions: Can be specified multiple times to define multiple `namespace_junction`. Each `namespace_juntion` block supports fields documented below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_host_name: The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] usage_model: The type of usage of the HPC Cache NFS Target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cache_name"] = cache_name
        __props__["name"] = name
        __props__["namespace_junctions"] = namespace_junctions
        __props__["resource_group_name"] = resource_group_name
        __props__["target_host_name"] = target_host_name
        __props__["usage_model"] = usage_model
        return CacheNfsTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cacheName")
    def cache_name(self) -> str:
        """
        The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cache_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the HPC Cache NFS Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceJunctions")
    def namespace_junctions(self) -> List['outputs.CacheNfsTargetNamespaceJunction']:
        """
        Can be specified multiple times to define multiple `namespace_junction`. Each `namespace_juntion` block supports fields documented below.
        """
        return pulumi.get(self, "namespace_junctions")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="targetHostName")
    def target_host_name(self) -> str:
        """
        The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "target_host_name")

    @property
    @pulumi.getter(name="usageModel")
    def usage_model(self) -> str:
        """
        The type of usage of the HPC Cache NFS Target.
        """
        return pulumi.get(self, "usage_model")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

