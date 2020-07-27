# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VirtualMachine']


class VirtualMachine(pulumi.CustomResource):
    additional_capabilities: pulumi.Output[Optional['outputs.VirtualMachineAdditionalCapabilities']] = pulumi.output_property("additionalCapabilities")
    """
    A `additional_capabilities` block.
    """
    availability_set_id: pulumi.Output[str] = pulumi.output_property("availabilitySetId")
    """
    The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
    """
    boot_diagnostics: pulumi.Output[Optional['outputs.VirtualMachineBootDiagnostics']] = pulumi.output_property("bootDiagnostics")
    """
    A `boot_diagnostics` block.
    """
    delete_data_disks_on_termination: pulumi.Output[Optional[bool]] = pulumi.output_property("deleteDataDisksOnTermination")
    """
    Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
    """
    delete_os_disk_on_termination: pulumi.Output[Optional[bool]] = pulumi.output_property("deleteOsDiskOnTermination")
    """
    Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
    """
    identity: pulumi.Output['outputs.VirtualMachineIdentity'] = pulumi.output_property("identity")
    """
    A `identity` block.
    """
    license_type: pulumi.Output[str] = pulumi.output_property("licenseType")
    """
    Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
    """
    network_interface_ids: pulumi.Output[List[str]] = pulumi.output_property("networkInterfaceIds")
    """
    A list of Network Interface ID's which should be associated with the Virtual Machine.
    """
    os_profile: pulumi.Output[Optional['outputs.VirtualMachineOsProfile']] = pulumi.output_property("osProfile")
    """
    An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
    """
    os_profile_linux_config: pulumi.Output[Optional['outputs.VirtualMachineOsProfileLinuxConfig']] = pulumi.output_property("osProfileLinuxConfig")
    """
    A `os_profile_linux_config` block.
    """
    os_profile_secrets: pulumi.Output[Optional[List['outputs.VirtualMachineOsProfileSecret']]] = pulumi.output_property("osProfileSecrets")
    """
    One or more `os_profile_secrets` blocks.
    """
    os_profile_windows_config: pulumi.Output[Optional['outputs.VirtualMachineOsProfileWindowsConfig']] = pulumi.output_property("osProfileWindowsConfig")
    """
    A `os_profile_windows_config` block.
    """
    plan: pulumi.Output[Optional['outputs.VirtualMachinePlan']] = pulumi.output_property("plan")
    """
    A `plan` block.
    """
    primary_network_interface_id: pulumi.Output[Optional[str]] = pulumi.output_property("primaryNetworkInterfaceId")
    """
    The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
    """
    proximity_placement_group_id: pulumi.Output[Optional[str]] = pulumi.output_property("proximityPlacementGroupId")
    """
    The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
    """
    storage_data_disks: pulumi.Output[List['outputs.VirtualMachineStorageDataDisk']] = pulumi.output_property("storageDataDisks")
    """
    One or more `storage_data_disk` blocks.
    """
    storage_image_reference: pulumi.Output['outputs.VirtualMachineStorageImageReference'] = pulumi.output_property("storageImageReference")
    """
    A `storage_image_reference` block.
    """
    storage_os_disk: pulumi.Output['outputs.VirtualMachineStorageOsDisk'] = pulumi.output_property("storageOsDisk")
    """
    A `storage_os_disk` block.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the Virtual Machine.
    """
    vm_size: pulumi.Output[str] = pulumi.output_property("vmSize")
    """
    Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
    """
    zones: pulumi.Output[Optional[str]] = pulumi.output_property("zones")
    """
    A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, additional_capabilities: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAdditionalCapabilitiesArgs']]] = None, availability_set_id: Optional[pulumi.Input[str]] = None, boot_diagnostics: Optional[pulumi.Input[pulumi.InputType['VirtualMachineBootDiagnosticsArgs']]] = None, delete_data_disks_on_termination: Optional[pulumi.Input[bool]] = None, delete_os_disk_on_termination: Optional[pulumi.Input[bool]] = None, identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']]] = None, license_type: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, network_interface_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, os_profile: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']]] = None, os_profile_linux_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']]] = None, os_profile_secrets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]]] = None, os_profile_windows_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']]] = None, plan: Optional[pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']]] = None, primary_network_interface_id: Optional[pulumi.Input[str]] = None, proximity_placement_group_id: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, storage_data_disks: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]]] = None, storage_image_reference: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']]] = None, storage_os_disk: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, vm_size: Optional[pulumi.Input[str]] = None, zones: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Virtual Machine.

        ## Disclaimers

        > **Note:** The `compute.VirtualMachine` resource has been superseded by the `compute.LinuxVirtualMachine` and `compute.WindowsVirtualMachine` resources. The existing `compute.VirtualMachine` resource will continue to be available throughout the 2.x releases however is in a feature-frozen state to maintain compatibility - new functionality will instead be added to the `compute.LinuxVirtualMachine` and `compute.WindowsVirtualMachine` resources.

        > **Note:** Data Disks can be attached either directly on the `compute.VirtualMachine` resource, or using the `compute.DataDiskAttachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.

        ## Example Usage
        ### From An Azure Platform Image)

        This example provisions a Virtual Machine with Managed Disks.

        ```python
        import pulumi
        import pulumi_azure as azure

        config = pulumi.Config()
        prefix = config.get("prefix")
        if prefix is None:
            prefix = "tfvmex"
        main_resource_group = azure.core.ResourceGroup("mainResourceGroup", location="West US 2")
        main_virtual_network = azure.network.VirtualNetwork("mainVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=main_resource_group.location,
            resource_group_name=main_resource_group.name)
        internal = azure.network.Subnet("internal",
            resource_group_name=main_resource_group.name,
            virtual_network_name=main_virtual_network.name,
            address_prefix="10.0.2.0/24")
        main_network_interface = azure.network.NetworkInterface("mainNetworkInterface",
            location=main_resource_group.location,
            resource_group_name=main_resource_group.name,
            ip_configurations=[{
                "name": "testconfiguration1",
                "subnet_id": internal.id,
                "privateIpAddressAllocation": "Dynamic",
            }])
        main_virtual_machine = azure.compute.VirtualMachine("mainVirtualMachine",
            location=main_resource_group.location,
            resource_group_name=main_resource_group.name,
            network_interface_ids=[main_network_interface.id],
            vm_size="Standard_DS1_v2",
            storage_image_reference={
                "publisher": "Canonical",
                "offer": "UbuntuServer",
                "sku": "16.04-LTS",
                "version": "latest",
            },
            storage_os_disk={
                "name": "myosdisk1",
                "caching": "ReadWrite",
                "create_option": "FromImage",
                "managedDiskType": "Standard_LRS",
            },
            os_profile={
                "computer_name": "hostname",
                "admin_username": "testadmin",
                "admin_password": "Password1234!",
            },
            os_profile_linux_config={
                "disable_password_authentication": False,
            },
            tags={
                "environment": "staging",
            })
        ```

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
        :param pulumi.Input[List[pulumi.Input[str]]] network_interface_ids: A list of Network Interface ID's which should be associated with the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']] os_profile: An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']] os_profile_linux_config: A `os_profile_linux_config` block.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]] os_profile_secrets: One or more `os_profile_secrets` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']] os_profile_windows_config: A `os_profile_windows_config` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']] plan: A `plan` block.
        :param pulumi.Input[str] primary_network_interface_id: The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]] storage_data_disks: One or more `storage_data_disk` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']] storage_image_reference: A `storage_image_reference` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']] storage_os_disk: A `storage_os_disk` block.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Virtual Machine.
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
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, additional_capabilities: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAdditionalCapabilitiesArgs']]] = None, availability_set_id: Optional[pulumi.Input[str]] = None, boot_diagnostics: Optional[pulumi.Input[pulumi.InputType['VirtualMachineBootDiagnosticsArgs']]] = None, delete_data_disks_on_termination: Optional[pulumi.Input[bool]] = None, delete_os_disk_on_termination: Optional[pulumi.Input[bool]] = None, identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']]] = None, license_type: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, network_interface_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, os_profile: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']]] = None, os_profile_linux_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']]] = None, os_profile_secrets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]]] = None, os_profile_windows_config: Optional[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']]] = None, plan: Optional[pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']]] = None, primary_network_interface_id: Optional[pulumi.Input[str]] = None, proximity_placement_group_id: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, storage_data_disks: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]]] = None, storage_image_reference: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']]] = None, storage_os_disk: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, vm_size: Optional[pulumi.Input[str]] = None, zones: Optional[pulumi.Input[str]] = None) -> 'VirtualMachine':
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[List[pulumi.Input[str]]] network_interface_ids: A list of Network Interface ID's which should be associated with the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileArgs']] os_profile: An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileLinuxConfigArgs']] os_profile_linux_config: A `os_profile_linux_config` block.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineOsProfileSecretArgs']]]] os_profile_secrets: One or more `os_profile_secrets` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineOsProfileWindowsConfigArgs']] os_profile_windows_config: A `os_profile_windows_config` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachinePlanArgs']] plan: A `plan` block.
        :param pulumi.Input[str] primary_network_interface_id: The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualMachineStorageDataDiskArgs']]]] storage_data_disks: One or more `storage_data_disk` blocks.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageImageReferenceArgs']] storage_image_reference: A `storage_image_reference` block.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageOsDiskArgs']] storage_os_disk: A `storage_os_disk` block.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Virtual Machine.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

