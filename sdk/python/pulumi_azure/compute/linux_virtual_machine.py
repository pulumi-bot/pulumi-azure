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

__all__ = ['LinuxVirtualMachine']


class LinuxVirtualMachine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_capabilities: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdditionalCapabilitiesArgs']]] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdminSshKeyArgs']]]]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 allow_extension_operations: Optional[pulumi.Input[bool]] = None,
                 availability_set_id: Optional[pulumi.Input[str]] = None,
                 boot_diagnostics: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineBootDiagnosticsArgs']]] = None,
                 computer_name: Optional[pulumi.Input[str]] = None,
                 custom_data: Optional[pulumi.Input[str]] = None,
                 dedicated_host_id: Optional[pulumi.Input[str]] = None,
                 disable_password_authentication: Optional[pulumi.Input[bool]] = None,
                 eviction_policy: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_bid_price: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 os_disk: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineOsDiskArgs']]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachinePlanArgs']]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 provision_vm_agent: Optional[pulumi.Input[bool]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineSecretArgs']]]]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 source_image_id: Optional[pulumi.Input[str]] = None,
                 source_image_reference: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineSourceImageReferenceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_machine_scale_set_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linux Virtual Machine.

        ## Disclaimers

        > **Note** This provider will automatically remove the OS Disk by default - this behaviour can be configured using the `features` configuration within the Provider configuration block.

        > **Note** This resource does not support Unmanaged Disks. If you need to use Unmanaged Disks you can continue to use the `compute.VirtualMachine` resource instead.

        > **Note** This resource does not support attaching existing OS Disks. You can instead capture an image of the OS Disk or continue to use the `compute.VirtualMachine` resource instead.

        > In this release there's a known issue where the `public_ip_address` and `public_ip_addresses` fields may not be fully populated for Dynamic Public IP's.

        ## Example Usage

        This example provisions a basic Linux Virtual Machine on an internal network.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.2.0/24")
        example_network_interface = azure.network.NetworkInterface("exampleNetworkInterface",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.NetworkInterfaceIpConfigurationArgs(
                name="internal",
                subnet_id=example_subnet.id,
                private_ip_address_allocation="Dynamic",
            )])
        example_linux_virtual_machine = azure.compute.LinuxVirtualMachine("exampleLinuxVirtualMachine",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            size="Standard_F2",
            admin_username="adminuser",
            network_interface_ids=[example_network_interface.id],
            admin_ssh_keys=[azure.compute.LinuxVirtualMachineAdminSshKeyArgs(
                username="adminuser",
                public_key=(lambda path: open(path).read())("~/.ssh/id_rsa.pub"),
            )],
            os_disk=azure.compute.LinuxVirtualMachineOsDiskArgs(
                caching="ReadWrite",
                storage_account_type="Standard_LRS",
            ),
            source_image_reference=azure.compute.LinuxVirtualMachineSourceImageReferenceArgs(
                publisher="Canonical",
                offer="UbuntuServer",
                sku="16.04-LTS",
                version="latest",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdditionalCapabilitiesArgs']] additional_capabilities: A `additional_capabilities` block as defined below.
        :param pulumi.Input[str] admin_password: The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdminSshKeyArgs']]]] admin_ssh_keys: One or more `admin_ssh_key` blocks as defined below.
        :param pulumi.Input[str] admin_username: The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] allow_extension_operations: Should Extension Operations be allowed on this Virtual Machine?
        :param pulumi.Input[str] availability_set_id: Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineBootDiagnosticsArgs']] boot_diagnostics: A `boot_diagnostics` block as defined below.
        :param pulumi.Input[str] computer_name: Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computer_name`, then you must specify `computer_name`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] custom_data: The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] dedicated_host_id: The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] disable_password_authentication: Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eviction_policy: Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[float] max_bid_price: The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `eviction_policy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
        :param pulumi.Input[str] name: The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_interface_ids: . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineOsDiskArgs']] os_disk: A `os_disk` block as defined below.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachinePlanArgs']] plan: A `plan` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] priority: Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] provision_vm_agent: Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineSecretArgs']]]] secrets: One or more `secret` blocks as defined below.
        :param pulumi.Input[str] size: The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
        :param pulumi.Input[str] source_image_id: The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineSourceImageReferenceArgs']] source_image_reference: A `source_image_reference` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to this Virtual Machine.
        :param pulumi.Input[str] virtual_machine_scale_set_id: Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
        :param pulumi.Input[str] zone: The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
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
            __props__['admin_password'] = admin_password
            __props__['admin_ssh_keys'] = admin_ssh_keys
            if admin_username is None:
                raise TypeError("Missing required property 'admin_username'")
            __props__['admin_username'] = admin_username
            __props__['allow_extension_operations'] = allow_extension_operations
            __props__['availability_set_id'] = availability_set_id
            __props__['boot_diagnostics'] = boot_diagnostics
            __props__['computer_name'] = computer_name
            __props__['custom_data'] = custom_data
            __props__['dedicated_host_id'] = dedicated_host_id
            __props__['disable_password_authentication'] = disable_password_authentication
            __props__['eviction_policy'] = eviction_policy
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['max_bid_price'] = max_bid_price
            __props__['name'] = name
            if network_interface_ids is None:
                raise TypeError("Missing required property 'network_interface_ids'")
            __props__['network_interface_ids'] = network_interface_ids
            if os_disk is None:
                raise TypeError("Missing required property 'os_disk'")
            __props__['os_disk'] = os_disk
            __props__['plan'] = plan
            __props__['priority'] = priority
            __props__['provision_vm_agent'] = provision_vm_agent
            __props__['proximity_placement_group_id'] = proximity_placement_group_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['secrets'] = secrets
            if size is None:
                raise TypeError("Missing required property 'size'")
            __props__['size'] = size
            __props__['source_image_id'] = source_image_id
            __props__['source_image_reference'] = source_image_reference
            __props__['tags'] = tags
            __props__['virtual_machine_scale_set_id'] = virtual_machine_scale_set_id
            __props__['zone'] = zone
            __props__['private_ip_address'] = None
            __props__['private_ip_addresses'] = None
            __props__['public_ip_address'] = None
            __props__['public_ip_addresses'] = None
            __props__['virtual_machine_id'] = None
        super(LinuxVirtualMachine, __self__).__init__(
            'azure:compute/linuxVirtualMachine:LinuxVirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_capabilities: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdditionalCapabilitiesArgs']]] = None,
            admin_password: Optional[pulumi.Input[str]] = None,
            admin_ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdminSshKeyArgs']]]]] = None,
            admin_username: Optional[pulumi.Input[str]] = None,
            allow_extension_operations: Optional[pulumi.Input[bool]] = None,
            availability_set_id: Optional[pulumi.Input[str]] = None,
            boot_diagnostics: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineBootDiagnosticsArgs']]] = None,
            computer_name: Optional[pulumi.Input[str]] = None,
            custom_data: Optional[pulumi.Input[str]] = None,
            dedicated_host_id: Optional[pulumi.Input[str]] = None,
            disable_password_authentication: Optional[pulumi.Input[bool]] = None,
            eviction_policy: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineIdentityArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            max_bid_price: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            os_disk: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineOsDiskArgs']]] = None,
            plan: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachinePlanArgs']]] = None,
            priority: Optional[pulumi.Input[str]] = None,
            private_ip_address: Optional[pulumi.Input[str]] = None,
            private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            provision_vm_agent: Optional[pulumi.Input[bool]] = None,
            proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
            public_ip_address: Optional[pulumi.Input[str]] = None,
            public_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineSecretArgs']]]]] = None,
            size: Optional[pulumi.Input[str]] = None,
            source_image_id: Optional[pulumi.Input[str]] = None,
            source_image_reference: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineSourceImageReferenceArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None,
            virtual_machine_scale_set_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'LinuxVirtualMachine':
        """
        Get an existing LinuxVirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdditionalCapabilitiesArgs']] additional_capabilities: A `additional_capabilities` block as defined below.
        :param pulumi.Input[str] admin_password: The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineAdminSshKeyArgs']]]] admin_ssh_keys: One or more `admin_ssh_key` blocks as defined below.
        :param pulumi.Input[str] admin_username: The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] allow_extension_operations: Should Extension Operations be allowed on this Virtual Machine?
        :param pulumi.Input[str] availability_set_id: Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineBootDiagnosticsArgs']] boot_diagnostics: A `boot_diagnostics` block as defined below.
        :param pulumi.Input[str] computer_name: Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computer_name`, then you must specify `computer_name`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] custom_data: The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] dedicated_host_id: The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] disable_password_authentication: Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eviction_policy: Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[float] max_bid_price: The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `eviction_policy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
        :param pulumi.Input[str] name: The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_interface_ids: . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineOsDiskArgs']] os_disk: A `os_disk` block as defined below.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachinePlanArgs']] plan: A `plan` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] priority: Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_ip_address: The Primary Private IP Address assigned to this Virtual Machine.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_addresses: A list of Private IP Addresses assigned to this Virtual Machine.
        :param pulumi.Input[bool] provision_vm_agent: Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] public_ip_address: The Primary Public IP Address assigned to this Virtual Machine.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addresses: A list of the Public IP Addresses assigned to this Virtual Machine.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineSecretArgs']]]] secrets: One or more `secret` blocks as defined below.
        :param pulumi.Input[str] size: The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
        :param pulumi.Input[str] source_image_id: The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineSourceImageReferenceArgs']] source_image_reference: A `source_image_reference` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to this Virtual Machine.
        :param pulumi.Input[str] virtual_machine_id: A 128-bit identifier which uniquely identifies this Virtual Machine.
        :param pulumi.Input[str] virtual_machine_scale_set_id: Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
        :param pulumi.Input[str] zone: The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_capabilities"] = additional_capabilities
        __props__["admin_password"] = admin_password
        __props__["admin_ssh_keys"] = admin_ssh_keys
        __props__["admin_username"] = admin_username
        __props__["allow_extension_operations"] = allow_extension_operations
        __props__["availability_set_id"] = availability_set_id
        __props__["boot_diagnostics"] = boot_diagnostics
        __props__["computer_name"] = computer_name
        __props__["custom_data"] = custom_data
        __props__["dedicated_host_id"] = dedicated_host_id
        __props__["disable_password_authentication"] = disable_password_authentication
        __props__["eviction_policy"] = eviction_policy
        __props__["identity"] = identity
        __props__["location"] = location
        __props__["max_bid_price"] = max_bid_price
        __props__["name"] = name
        __props__["network_interface_ids"] = network_interface_ids
        __props__["os_disk"] = os_disk
        __props__["plan"] = plan
        __props__["priority"] = priority
        __props__["private_ip_address"] = private_ip_address
        __props__["private_ip_addresses"] = private_ip_addresses
        __props__["provision_vm_agent"] = provision_vm_agent
        __props__["proximity_placement_group_id"] = proximity_placement_group_id
        __props__["public_ip_address"] = public_ip_address
        __props__["public_ip_addresses"] = public_ip_addresses
        __props__["resource_group_name"] = resource_group_name
        __props__["secrets"] = secrets
        __props__["size"] = size
        __props__["source_image_id"] = source_image_id
        __props__["source_image_reference"] = source_image_reference
        __props__["tags"] = tags
        __props__["virtual_machine_id"] = virtual_machine_id
        __props__["virtual_machine_scale_set_id"] = virtual_machine_scale_set_id
        __props__["zone"] = zone
        return LinuxVirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalCapabilities")
    def additional_capabilities(self) -> pulumi.Output[Optional['outputs.LinuxVirtualMachineAdditionalCapabilities']]:
        """
        A `additional_capabilities` block as defined below.
        """
        return pulumi.get(self, "additional_capabilities")

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> pulumi.Output[Optional[str]]:
        """
        The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminSshKeys")
    def admin_ssh_keys(self) -> pulumi.Output[Optional[Sequence['outputs.LinuxVirtualMachineAdminSshKey']]]:
        """
        One or more `admin_ssh_key` blocks as defined below.
        """
        return pulumi.get(self, "admin_ssh_keys")

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> pulumi.Output[str]:
        """
        The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter(name="allowExtensionOperations")
    def allow_extension_operations(self) -> pulumi.Output[Optional[bool]]:
        """
        Should Extension Operations be allowed on this Virtual Machine?
        """
        return pulumi.get(self, "allow_extension_operations")

    @property
    @pulumi.getter(name="availabilitySetId")
    def availability_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "availability_set_id")

    @property
    @pulumi.getter(name="bootDiagnostics")
    def boot_diagnostics(self) -> pulumi.Output[Optional['outputs.LinuxVirtualMachineBootDiagnostics']]:
        """
        A `boot_diagnostics` block as defined below.
        """
        return pulumi.get(self, "boot_diagnostics")

    @property
    @pulumi.getter(name="computerName")
    def computer_name(self) -> pulumi.Output[str]:
        """
        Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computer_name`, then you must specify `computer_name`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "computer_name")

    @property
    @pulumi.getter(name="customData")
    def custom_data(self) -> pulumi.Output[Optional[str]]:
        """
        The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "custom_data")

    @property
    @pulumi.getter(name="dedicatedHostId")
    def dedicated_host_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "dedicated_host_id")

    @property
    @pulumi.getter(name="disablePasswordAuthentication")
    def disable_password_authentication(self) -> pulumi.Output[Optional[bool]]:
        """
        Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "disable_password_authentication")

    @property
    @pulumi.getter(name="evictionPolicy")
    def eviction_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eviction_policy")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.LinuxVirtualMachineIdentity']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxBidPrice")
    def max_bid_price(self) -> pulumi.Output[Optional[float]]:
        """
        The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `eviction_policy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
        """
        return pulumi.get(self, "max_bid_price")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaceIds")
    def network_interface_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
        """
        return pulumi.get(self, "network_interface_ids")

    @property
    @pulumi.getter(name="osDisk")
    def os_disk(self) -> pulumi.Output['outputs.LinuxVirtualMachineOsDisk']:
        """
        A `os_disk` block as defined below.
        """
        return pulumi.get(self, "os_disk")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[Optional['outputs.LinuxVirtualMachinePlan']]:
        """
        A `plan` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> pulumi.Output[str]:
        """
        The Primary Private IP Address assigned to this Virtual Machine.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="privateIpAddresses")
    def private_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of Private IP Addresses assigned to this Virtual Machine.
        """
        return pulumi.get(self, "private_ip_addresses")

    @property
    @pulumi.getter(name="provisionVmAgent")
    def provision_vm_agent(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "provision_vm_agent")

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> pulumi.Output[str]:
        """
        The Primary Public IP Address assigned to this Virtual Machine.
        """
        return pulumi.get(self, "public_ip_address")

    @property
    @pulumi.getter(name="publicIpAddresses")
    def public_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of the Public IP Addresses assigned to this Virtual Machine.
        """
        return pulumi.get(self, "public_ip_addresses")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def secrets(self) -> pulumi.Output[Optional[Sequence['outputs.LinuxVirtualMachineSecret']]]:
        """
        One or more `secret` blocks as defined below.
        """
        return pulumi.get(self, "secrets")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sourceImageId")
    def source_image_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_image_id")

    @property
    @pulumi.getter(name="sourceImageReference")
    def source_image_reference(self) -> pulumi.Output[Optional['outputs.LinuxVirtualMachineSourceImageReference']]:
        """
        A `source_image_reference` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_image_reference")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to this Virtual Machine.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[str]:
        """
        A 128-bit identifier which uniquely identifies this Virtual Machine.
        """
        return pulumi.get(self, "virtual_machine_id")

    @property
    @pulumi.getter(name="virtualMachineScaleSetId")
    def virtual_machine_scale_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_scale_set_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

