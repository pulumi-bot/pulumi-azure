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

__all__ = ['VirtualMachine']


class VirtualMachine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_capabilities: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAdditionalCapabilitiesArgs']]] = None,
                 availability_set_id: Optional[pulumi.Input[str]] = None,
                 boot_diagnostics: Optional[pulumi.Input[pulumi.InputType['VirtualMachineBootDiagnosticsArgs']]] = None,
                 delete_data_disks_on_termination: Optional[pulumi.Input[bool]] = None,
                 delete_os_disk_on_termination: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']]] = None,
                 license_type: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 os_profile: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']]] = None,
                 os_profile_linux_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']]] = None,
                 os_profile_secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]]] = None,
                 os_profile_windows_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']]] = None,
                 primary_network_interface_id: Optional[pulumi.Input[str]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_data_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]]] = None,
                 storage_image_reference: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']]] = None,
                 storage_os_disk: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Virtual Machine.

        ## Disclaimers

        > **Note:** The `compute.VirtualMachine` resource has been superseded by the `compute.LinuxVirtualMachine` and `compute.WindowsVirtualMachine` resources. The existing `compute.VirtualMachine` resource will continue to be available throughout the 2.x releases however is in a feature-frozen state to maintain compatibility - new functionality will instead be added to the `compute.LinuxVirtualMachine` and `compute.WindowsVirtualMachine` resources.

        > **Note:** Data Disks can be attached either directly on the `compute.VirtualMachine` resource, or using the `compute.DataDiskAttachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAdditionalCapabilitiesArgs']] additional_capabilities: A `additional_capabilities` block.
        :param pulumi.Input[str] availability_set_id: The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualMachineBootDiagnosticsArgs']] boot_diagnostics: A `boot_diagnostics` block.
        :param pulumi.Input[bool] delete_data_disks_on_termination: Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        :param pulumi.Input[bool] delete_os_disk_on_termination: Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']] identity: A `identity` block.
        :param pulumi.Input[str] license_type: Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        :param pulumi.Input[str] location: Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_interface_ids: A list of Network Interface ID's which should be associated with the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']] os_profile: An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']] os_profile_linux_config: A `os_profile_linux_config` block.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]] os_profile_secrets: One or more `os_profile_secrets` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']] os_profile_windows_config: A `os_profile_windows_config` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']] plan: A `plan` block.
        :param pulumi.Input[str] primary_network_interface_id: The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]] storage_data_disks: One or more `storage_data_disk` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']] storage_image_reference: A `storage_image_reference` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']] storage_os_disk: A `storage_os_disk` block.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Virtual Machine.
        :param pulumi.Input[str] vm_size: Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
        :param pulumi.Input[str] zones: A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
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

            __props__['additional_capabilities'] = additional_capabilities
            __props__['availability_set_id'] = availability_set_id
            __props__['boot_diagnostics'] = boot_diagnostics
            __props__['delete_data_disks_on_termination'] = delete_data_disks_on_termination
            __props__['delete_os_disk_on_termination'] = delete_os_disk_on_termination
            __props__['identity'] = identity
            __props__['license_type'] = license_type
            __props__['location'] = location
            __props__['name'] = name
            if network_interface_ids is None:
                raise TypeError("Missing required property 'network_interface_ids'")
            __props__['network_interface_ids'] = network_interface_ids
            __props__['os_profile'] = os_profile
            __props__['os_profile_linux_config'] = os_profile_linux_config
            __props__['os_profile_secrets'] = os_profile_secrets
            __props__['os_profile_windows_config'] = os_profile_windows_config
            __props__['plan'] = plan
            __props__['primary_network_interface_id'] = primary_network_interface_id
            __props__['proximity_placement_group_id'] = proximity_placement_group_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['storage_data_disks'] = storage_data_disks
            __props__['storage_image_reference'] = storage_image_reference
            if storage_os_disk is None:
                raise TypeError("Missing required property 'storage_os_disk'")
            __props__['storage_os_disk'] = storage_os_disk
            __props__['tags'] = tags
            if vm_size is None:
                raise TypeError("Missing required property 'vm_size'")
            __props__['vm_size'] = vm_size
            __props__['zones'] = zones
        super(VirtualMachine, __self__).__init__(
            'azure:compute/virtualMachine:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_capabilities: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAdditionalCapabilitiesArgs']]] = None,
            availability_set_id: Optional[pulumi.Input[str]] = None,
            boot_diagnostics: Optional[pulumi.Input[pulumi.InputType['VirtualMachineBootDiagnosticsArgs']]] = None,
            delete_data_disks_on_termination: Optional[pulumi.Input[bool]] = None,
            delete_os_disk_on_termination: Optional[pulumi.Input[bool]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']]] = None,
            license_type: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            os_profile: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']]] = None,
            os_profile_linux_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']]] = None,
            os_profile_secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]]] = None,
            os_profile_windows_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']]] = None,
            plan: Optional[pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']]] = None,
            primary_network_interface_id: Optional[pulumi.Input[str]] = None,
            proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_data_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]]] = None,
            storage_image_reference: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']]] = None,
            storage_os_disk: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vm_size: Optional[pulumi.Input[str]] = None,
            zones: Optional[pulumi.Input[str]] = None) -> 'VirtualMachine':
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAdditionalCapabilitiesArgs']] additional_capabilities: A `additional_capabilities` block.
        :param pulumi.Input[str] availability_set_id: The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualMachineBootDiagnosticsArgs']] boot_diagnostics: A `boot_diagnostics` block.
        :param pulumi.Input[bool] delete_data_disks_on_termination: Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        :param pulumi.Input[bool] delete_os_disk_on_termination: Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']] identity: A `identity` block.
        :param pulumi.Input[str] license_type: Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        :param pulumi.Input[str] location: Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_interface_ids: A list of Network Interface ID's which should be associated with the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']] os_profile: An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']] os_profile_linux_config: A `os_profile_linux_config` block.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]] os_profile_secrets: One or more `os_profile_secrets` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']] os_profile_windows_config: A `os_profile_windows_config` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']] plan: A `plan` block.
        :param pulumi.Input[str] primary_network_interface_id: The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]] storage_data_disks: One or more `storage_data_disk` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']] storage_image_reference: A `storage_image_reference` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']] storage_os_disk: A `storage_os_disk` block.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Virtual Machine.
        :param pulumi.Input[str] vm_size: Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
        :param pulumi.Input[str] zones: A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_capabilities"] = additional_capabilities
        __props__["availability_set_id"] = availability_set_id
        __props__["boot_diagnostics"] = boot_diagnostics
        __props__["delete_data_disks_on_termination"] = delete_data_disks_on_termination
        __props__["delete_os_disk_on_termination"] = delete_os_disk_on_termination
        __props__["identity"] = identity
        __props__["license_type"] = license_type
        __props__["location"] = location
        __props__["name"] = name
        __props__["network_interface_ids"] = network_interface_ids
        __props__["os_profile"] = os_profile
        __props__["os_profile_linux_config"] = os_profile_linux_config
        __props__["os_profile_secrets"] = os_profile_secrets
        __props__["os_profile_windows_config"] = os_profile_windows_config
        __props__["plan"] = plan
        __props__["primary_network_interface_id"] = primary_network_interface_id
        __props__["proximity_placement_group_id"] = proximity_placement_group_id
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_data_disks"] = storage_data_disks
        __props__["storage_image_reference"] = storage_image_reference
        __props__["storage_os_disk"] = storage_os_disk
        __props__["tags"] = tags
        __props__["vm_size"] = vm_size
        __props__["zones"] = zones
        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalCapabilities")
    def additional_capabilities(self) -> pulumi.Output[Optional['outputs.VirtualMachineAdditionalCapabilities']]:
        """
        A `additional_capabilities` block.
        """
        return pulumi.get(self, "additional_capabilities")

    @property
    @pulumi.getter(name="availabilitySetId")
    def availability_set_id(self) -> pulumi.Output[str]:
        """
        The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "availability_set_id")

    @property
    @pulumi.getter(name="bootDiagnostics")
    def boot_diagnostics(self) -> pulumi.Output[Optional['outputs.VirtualMachineBootDiagnostics']]:
        """
        A `boot_diagnostics` block.
        """
        return pulumi.get(self, "boot_diagnostics")

    @property
    @pulumi.getter(name="deleteDataDisksOnTermination")
    def delete_data_disks_on_termination(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        """
        return pulumi.get(self, "delete_data_disks_on_termination")

    @property
    @pulumi.getter(name="deleteOsDiskOnTermination")
    def delete_os_disk_on_termination(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        """
        return pulumi.get(self, "delete_os_disk_on_termination")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output['outputs.VirtualMachineIdentity']:
        """
        A `identity` block.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> pulumi.Output[str]:
        """
        Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaceIds")
    def network_interface_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of Network Interface ID's which should be associated with the Virtual Machine.
        """
        return pulumi.get(self, "network_interface_ids")

    @property
    @pulumi.getter(name="osProfile")
    def os_profile(self) -> pulumi.Output[Optional['outputs.VirtualMachineOsProfile']]:
        """
        An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        """
        return pulumi.get(self, "os_profile")

    @property
    @pulumi.getter(name="osProfileLinuxConfig")
    def os_profile_linux_config(self) -> pulumi.Output[Optional['outputs.VirtualMachineOsProfileLinuxConfig']]:
        """
        A `os_profile_linux_config` block.
        """
        return pulumi.get(self, "os_profile_linux_config")

    @property
    @pulumi.getter(name="osProfileSecrets")
    def os_profile_secrets(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualMachineOsProfileSecret']]]:
        """
        One or more `os_profile_secrets` blocks.
        """
        return pulumi.get(self, "os_profile_secrets")

    @property
    @pulumi.getter(name="osProfileWindowsConfig")
    def os_profile_windows_config(self) -> pulumi.Output[Optional['outputs.VirtualMachineOsProfileWindowsConfig']]:
        """
        A `os_profile_windows_config` block.
        """
        return pulumi.get(self, "os_profile_windows_config")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[Optional['outputs.VirtualMachinePlan']]:
        """
        A `plan` block.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="primaryNetworkInterfaceId")
    def primary_network_interface_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        """
        return pulumi.get(self, "primary_network_interface_id")

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="storageDataDisks")
    def storage_data_disks(self) -> pulumi.Output[Sequence['outputs.VirtualMachineStorageDataDisk']]:
        """
        One or more `storage_data_disk` blocks.
        """
        return pulumi.get(self, "storage_data_disks")

    @property
    @pulumi.getter(name="storageImageReference")
    def storage_image_reference(self) -> pulumi.Output['outputs.VirtualMachineStorageImageReference']:
        """
        A `storage_image_reference` block.
        """
        return pulumi.get(self, "storage_image_reference")

    @property
    @pulumi.getter(name="storageOsDisk")
    def storage_os_disk(self) -> pulumi.Output['outputs.VirtualMachineStorageOsDisk']:
        """
        A `storage_os_disk` block.
        """
        return pulumi.get(self, "storage_os_disk")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the Virtual Machine.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[str]:
        """
        Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[str]]:
        """
        A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

