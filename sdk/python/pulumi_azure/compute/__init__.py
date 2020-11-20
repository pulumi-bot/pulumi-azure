# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .availability_set import *
from .bastion_host import *
from .data_disk_attachment import *
from .dedicated_host import *
from .dedicated_host_group import *
from .disk_encryption_set import *
from .extension import *
from .get_availability_set import *
from .get_dedicated_host import *
from .get_dedicated_host_group import *
from .get_disk_encryption_set import *
from .get_image import *
from .get_images import *
from .get_managed_disk import *
from .get_platform_image import *
from .get_shared_image import *
from .get_shared_image_gallery import *
from .get_shared_image_version import *
from .get_shared_image_versions import *
from .get_snapshot import *
from .get_virtual_machine import *
from .get_virtual_machine_scale_set import *
from .image import *
from .linux_virtual_machine import *
from .linux_virtual_machine_scale_set import *
from .managed_disk import *
from .orchestrated_virtual_machine_scale_set import *
from .scale_set import *
from .shared_image import *
from .shared_image_gallery import *
from .shared_image_version import *
from .snapshot import *
from .virtual_machine import *
from .virtual_machine_scale_set_extension import *
from .windows_virtual_machine import *
from .windows_virtual_machine_scale_set import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "azure:compute/availabilitySet:AvailabilitySet":
            return AvailabilitySet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/bastionHost:BastionHost":
            return BastionHost(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/dataDiskAttachment:DataDiskAttachment":
            return DataDiskAttachment(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/dedicatedHost:DedicatedHost":
            return DedicatedHost(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/dedicatedHostGroup:DedicatedHostGroup":
            return DedicatedHostGroup(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/diskEncryptionSet:DiskEncryptionSet":
            return DiskEncryptionSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/extension:Extension":
            return Extension(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/image:Image":
            return Image(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/linuxVirtualMachine:LinuxVirtualMachine":
            return LinuxVirtualMachine(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/linuxVirtualMachineScaleSet:LinuxVirtualMachineScaleSet":
            return LinuxVirtualMachineScaleSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/managedDisk:ManagedDisk":
            return ManagedDisk(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet":
            return OrchestratedVirtualMachineScaleSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/scaleSet:ScaleSet":
            return ScaleSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/sharedImage:SharedImage":
            return SharedImage(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/sharedImageGallery:SharedImageGallery":
            return SharedImageGallery(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/sharedImageVersion:SharedImageVersion":
            return SharedImageVersion(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/snapshot:Snapshot":
            return Snapshot(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/virtualMachine:VirtualMachine":
            return VirtualMachine(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension":
            return VirtualMachineScaleSetExtension(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/windowsVirtualMachine:WindowsVirtualMachine":
            return WindowsVirtualMachine(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:compute/windowsVirtualMachineScaleSet:WindowsVirtualMachineScaleSet":
            return WindowsVirtualMachineScaleSet(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("azure", "compute/availabilitySet", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/bastionHost", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/dataDiskAttachment", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/dedicatedHost", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/dedicatedHostGroup", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/diskEncryptionSet", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/extension", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/image", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/linuxVirtualMachine", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/linuxVirtualMachineScaleSet", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/managedDisk", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/orchestratedVirtualMachineScaleSet", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/scaleSet", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/sharedImage", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/sharedImageGallery", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/sharedImageVersion", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/snapshot", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/virtualMachine", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/virtualMachineScaleSetExtension", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/windowsVirtualMachine", _module_instance)
pulumi.runtime.register_resource_module("azure", "compute/windowsVirtualMachineScaleSet", _module_instance)
