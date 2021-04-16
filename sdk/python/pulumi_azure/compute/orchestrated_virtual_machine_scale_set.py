# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrchestratedVirtualMachineScaleSetArgs', 'OrchestratedVirtualMachineScaleSet']

@pulumi.input_type
class OrchestratedVirtualMachineScaleSetArgs:
    def __init__(__self__, *,
                 platform_fault_domain_count: pulumi.Input[int],
                 resource_group_name: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 single_placement_group: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrchestratedVirtualMachineScaleSet resource.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] single_placement_group: Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if proximity_placement_group_id is not None:
            pulumi.set(__self__, "proximity_placement_group_id", proximity_placement_group_id)
        if single_placement_group is not None:
            pulumi.set(__self__, "single_placement_group", single_placement_group)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Input[int]:
        """
        Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "platform_fault_domain_count", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @proximity_placement_group_id.setter
    def proximity_placement_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proximity_placement_group_id", value)

    @property
    @pulumi.getter(name="singlePlacementGroup")
    def single_placement_group(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        """
        return pulumi.get(self, "single_placement_group")

    @single_placement_group.setter
    def single_placement_group(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "single_placement_group", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[str]]:
        """
        A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zones", value)


@pulumi.input_type
class _OrchestratedVirtualMachineScaleSetState:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 single_placement_group: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 unique_id: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrchestratedVirtualMachineScaleSet resources.
        :param pulumi.Input[str] location: The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] single_placement_group: Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] unique_id: The Unique ID for the Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_fault_domain_count is not None:
            pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        if proximity_placement_group_id is not None:
            pulumi.set(__self__, "proximity_placement_group_id", proximity_placement_group_id)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if single_placement_group is not None:
            pulumi.set(__self__, "single_placement_group", single_placement_group)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if unique_id is not None:
            pulumi.set(__self__, "unique_id", unique_id)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_fault_domain_count", value)

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @proximity_placement_group_id.setter
    def proximity_placement_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proximity_placement_group_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="singlePlacementGroup")
    def single_placement_group(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        """
        return pulumi.get(self, "single_placement_group")

    @single_placement_group.setter
    def single_placement_group(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "single_placement_group", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Unique ID for the Orchestrated Virtual Machine Scale Set.
        """
        return pulumi.get(self, "unique_id")

    @unique_id.setter
    def unique_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unique_id", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[str]]:
        """
        A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zones", value)


class OrchestratedVirtualMachineScaleSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 single_placement_group: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Orchestrated Virtual Machine Scale Set.

        > **Note:** Orchestrated Virtual Machine Scale Sets are in Public Preview and it may receive breaking changes - [more details can be found in the Azure Documentation](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/orchestration-modes).

        > **Note:** Azure is planning to deprecate the `single_placement_group` attribute in the Orchestrated Virtual Machine Scale Set starting from api-version `2019-12-01` and there will be a breaking change in the Orchestrated Virtual Machine Scale Set.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_orchestrated_virtual_machine_scale_set = azure.compute.OrchestratedVirtualMachineScaleSet("exampleOrchestratedVirtualMachineScaleSet",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            platform_fault_domain_count=1,
            zones=["1"])
        ```

        ## Import

        An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/Microsoft.Compute/virtualMachineScaleSets/scaleset1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] single_placement_group: Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrchestratedVirtualMachineScaleSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Orchestrated Virtual Machine Scale Set.

        > **Note:** Orchestrated Virtual Machine Scale Sets are in Public Preview and it may receive breaking changes - [more details can be found in the Azure Documentation](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/orchestration-modes).

        > **Note:** Azure is planning to deprecate the `single_placement_group` attribute in the Orchestrated Virtual Machine Scale Set starting from api-version `2019-12-01` and there will be a breaking change in the Orchestrated Virtual Machine Scale Set.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_orchestrated_virtual_machine_scale_set = azure.compute.OrchestratedVirtualMachineScaleSet("exampleOrchestratedVirtualMachineScaleSet",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            platform_fault_domain_count=1,
            zones=["1"])
        ```

        ## Import

        An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/Microsoft.Compute/virtualMachineScaleSets/scaleset1
        ```

        :param str resource_name: The name of the resource.
        :param OrchestratedVirtualMachineScaleSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrchestratedVirtualMachineScaleSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 single_placement_group: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrchestratedVirtualMachineScaleSetArgs.__new__(OrchestratedVirtualMachineScaleSetArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if platform_fault_domain_count is None and not opts.urn:
                raise TypeError("Missing required property 'platform_fault_domain_count'")
            __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
            __props__.__dict__["proximity_placement_group_id"] = proximity_placement_group_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["single_placement_group"] = single_placement_group
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zones"] = zones
            __props__.__dict__["unique_id"] = None
        super(OrchestratedVirtualMachineScaleSet, __self__).__init__(
            'azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
            proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            single_placement_group: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            unique_id: Optional[pulumi.Input[str]] = None,
            zones: Optional[pulumi.Input[str]] = None) -> 'OrchestratedVirtualMachineScaleSet':
        """
        Get an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] single_placement_group: Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] unique_id: The Unique ID for the Orchestrated Virtual Machine Scale Set.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrchestratedVirtualMachineScaleSetState.__new__(_OrchestratedVirtualMachineScaleSetState)

        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
        __props__.__dict__["proximity_placement_group_id"] = proximity_placement_group_id
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["single_placement_group"] = single_placement_group
        __props__.__dict__["tags"] = tags
        __props__.__dict__["unique_id"] = unique_id
        __props__.__dict__["zones"] = zones
        return OrchestratedVirtualMachineScaleSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Output[int]:
        """
        Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="singlePlacementGroup")
    def single_placement_group(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        """
        return pulumi.get(self, "single_placement_group")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> pulumi.Output[str]:
        """
        The Unique ID for the Orchestrated Virtual Machine Scale Set.
        """
        return pulumi.get(self, "unique_id")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[str]]:
        """
        A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

