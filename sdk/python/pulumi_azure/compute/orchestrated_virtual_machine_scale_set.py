# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class OrchestratedVirtualMachineScaleSet(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
    """
    platform_fault_domain_count: pulumi.Output[float]
    """
    Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
    """
    single_placement_group: pulumi.Output[bool]
    """
    Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
    """
    unique_id: pulumi.Output[str]
    """
    The Unique ID for the Orchestrated Virtual Machine Scale Set.
    """
    zones: pulumi.Output[str]
    """
    A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, platform_fault_domain_count=None, resource_group_name=None, single_placement_group=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Orchestrated Virtual Machine Scale Set.

        > **Note:** Orchestrated Virtual Machine Scale Sets are in Public Preview - [more details can be found in the Azure Documentation](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/orchestration-modes).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_orchestrated_virtual_machine_scale_set = azure.compute.OrchestratedVirtualMachineScaleSet("exampleOrchestratedVirtualMachineScaleSet",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            platform_fault_domain_count=5,
            single_placement_group=True,
            zones=["1"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[float] platform_fault_domain_count: Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] single_placement_group: Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
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
            opts.version = utilities.get_version()
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
            if single_placement_group is None:
                raise TypeError("Missing required property 'single_placement_group'")
            __props__['single_placement_group'] = single_placement_group
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['unique_id'] = None
        super(OrchestratedVirtualMachineScaleSet, __self__).__init__(
            'azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, name=None, platform_fault_domain_count=None, resource_group_name=None, single_placement_group=None, tags=None, unique_id=None, zones=None):
        """
        Get an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[float] platform_fault_domain_count: Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] single_placement_group: Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] unique_id: The Unique ID for the Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["platform_fault_domain_count"] = platform_fault_domain_count
        __props__["resource_group_name"] = resource_group_name
        __props__["single_placement_group"] = single_placement_group
        __props__["tags"] = tags
        __props__["unique_id"] = unique_id
        __props__["zones"] = zones
        return OrchestratedVirtualMachineScaleSet(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
