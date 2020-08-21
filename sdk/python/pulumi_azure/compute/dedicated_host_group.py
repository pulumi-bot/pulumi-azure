# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['DedicatedHostGroup']


class DedicatedHostGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manage a Dedicated Host Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_dedicated_host_group = azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            platform_fault_domain_count=1)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        :param pulumi.Input[float] platform_fault_domain_count: The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
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
            if platform_fault_domain_count is None:
                raise TypeError("Missing required property 'platform_fault_domain_count'")
            __props__['platform_fault_domain_count'] = platform_fault_domain_count
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['zones'] = zones
        super(DedicatedHostGroup, __self__).__init__(
            'azure:compute/dedicatedHostGroup:DedicatedHostGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_fault_domain_count: Optional[pulumi.Input[float]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            zones: Optional[pulumi.Input[str]] = None) -> 'DedicatedHostGroup':
        """
        Get an existing DedicatedHostGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        :param pulumi.Input[float] platform_fault_domain_count: The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["platform_fault_domain_count"] = platform_fault_domain_count
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["zones"] = zones
        return DedicatedHostGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> float:
        """
        The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> Optional[str]:
        """
        A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

