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

__all__ = ['LinuxVirtualMachineArgs', 'LinuxVirtualMachine']

@pulumi.input_type
class LinuxVirtualMachineArgs:
    def __init__(__self__, *,
                 gallery_image_reference: pulumi.Input['LinuxVirtualMachineGalleryImageReferenceArgs'],
                 lab_name: pulumi.Input[str],
                 lab_subnet_name: pulumi.Input[str],
                 lab_virtual_network_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 size: pulumi.Input[str],
                 storage_type: pulumi.Input[str],
                 username: pulumi.Input[str],
                 allow_claim: Optional[pulumi.Input[bool]] = None,
                 disallow_public_ip_address: Optional[pulumi.Input[bool]] = None,
                 inbound_nat_rules: Optional[pulumi.Input[Sequence[pulumi.Input['LinuxVirtualMachineInboundNatRuleArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 ssh_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a LinuxVirtualMachine resource.
        :param pulumi.Input['LinuxVirtualMachineGalleryImageReferenceArgs'] gallery_image_reference: A `gallery_image_reference` block as defined below.
        :param pulumi.Input[str] lab_name: Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_subnet_name: The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_virtual_network_id: The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] size: The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_type: The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
        :param pulumi.Input[str] username: The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] allow_claim: Can this Virtual Machine be claimed by users? Defaults to `true`.
        :param pulumi.Input[bool] disallow_public_ip_address: Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['LinuxVirtualMachineInboundNatRuleArgs']]] inbound_nat_rules: One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notes: Any notes about the Virtual Machine.
        :param pulumi.Input[str] password: The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ssh_key: The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "gallery_image_reference", gallery_image_reference)
        pulumi.set(__self__, "lab_name", lab_name)
        pulumi.set(__self__, "lab_subnet_name", lab_subnet_name)
        pulumi.set(__self__, "lab_virtual_network_id", lab_virtual_network_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "storage_type", storage_type)
        pulumi.set(__self__, "username", username)
        if allow_claim is not None:
            pulumi.set(__self__, "allow_claim", allow_claim)
        if disallow_public_ip_address is not None:
            pulumi.set(__self__, "disallow_public_ip_address", disallow_public_ip_address)
        if inbound_nat_rules is not None:
            pulumi.set(__self__, "inbound_nat_rules", inbound_nat_rules)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if ssh_key is not None:
            pulumi.set(__self__, "ssh_key", ssh_key)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="galleryImageReference")
    def gallery_image_reference(self) -> pulumi.Input['LinuxVirtualMachineGalleryImageReferenceArgs']:
        """
        A `gallery_image_reference` block as defined below.
        """
        return pulumi.get(self, "gallery_image_reference")

    @gallery_image_reference.setter
    def gallery_image_reference(self, value: pulumi.Input['LinuxVirtualMachineGalleryImageReferenceArgs']):
        pulumi.set(self, "gallery_image_reference", value)

    @property
    @pulumi.getter(name="labName")
    def lab_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_name")

    @lab_name.setter
    def lab_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "lab_name", value)

    @property
    @pulumi.getter(name="labSubnetName")
    def lab_subnet_name(self) -> pulumi.Input[str]:
        """
        The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_subnet_name")

    @lab_subnet_name.setter
    def lab_subnet_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "lab_subnet_name", value)

    @property
    @pulumi.getter(name="labVirtualNetworkId")
    def lab_virtual_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_virtual_network_id")

    @lab_virtual_network_id.setter
    def lab_virtual_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "lab_virtual_network_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[str]:
        """
        The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[str]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Input[str]:
        """
        The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="allowClaim")
    def allow_claim(self) -> Optional[pulumi.Input[bool]]:
        """
        Can this Virtual Machine be claimed by users? Defaults to `true`.
        """
        return pulumi.get(self, "allow_claim")

    @allow_claim.setter
    def allow_claim(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_claim", value)

    @property
    @pulumi.getter(name="disallowPublicIpAddress")
    def disallow_public_ip_address(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "disallow_public_ip_address")

    @disallow_public_ip_address.setter
    def disallow_public_ip_address(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disallow_public_ip_address", value)

    @property
    @pulumi.getter(name="inboundNatRules")
    def inbound_nat_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LinuxVirtualMachineInboundNatRuleArgs']]]]:
        """
        One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "inbound_nat_rules")

    @inbound_nat_rules.setter
    def inbound_nat_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LinuxVirtualMachineInboundNatRuleArgs']]]]):
        pulumi.set(self, "inbound_nat_rules", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        Any notes about the Virtual Machine.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="sshKey")
    def ssh_key(self) -> Optional[pulumi.Input[str]]:
        """
        The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ssh_key")

    @ssh_key.setter
    def ssh_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_key", value)

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


class LinuxVirtualMachine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_claim: Optional[pulumi.Input[bool]] = None,
                 disallow_public_ip_address: Optional[pulumi.Input[bool]] = None,
                 gallery_image_reference: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineGalleryImageReferenceArgs']]] = None,
                 inbound_nat_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineInboundNatRuleArgs']]]]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 lab_subnet_name: Optional[pulumi.Input[str]] = None,
                 lab_virtual_network_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 ssh_key: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linux Virtual Machine within a Dev Test Lab.

        ## Import

        Dev Test Linux Virtual Machines can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:devtest/linuxVirtualMachine:LinuxVirtualMachine machine1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_claim: Can this Virtual Machine be claimed by users? Defaults to `true`.
        :param pulumi.Input[bool] disallow_public_ip_address: Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineGalleryImageReferenceArgs']] gallery_image_reference: A `gallery_image_reference` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineInboundNatRuleArgs']]]] inbound_nat_rules: One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_name: Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_subnet_name: The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_virtual_network_id: The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notes: Any notes about the Virtual Machine.
        :param pulumi.Input[str] password: The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] size: The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ssh_key: The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_type: The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] username: The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinuxVirtualMachineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Linux Virtual Machine within a Dev Test Lab.

        ## Import

        Dev Test Linux Virtual Machines can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:devtest/linuxVirtualMachine:LinuxVirtualMachine machine1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
        ```

        :param str resource_name: The name of the resource.
        :param LinuxVirtualMachineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinuxVirtualMachineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_claim: Optional[pulumi.Input[bool]] = None,
                 disallow_public_ip_address: Optional[pulumi.Input[bool]] = None,
                 gallery_image_reference: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineGalleryImageReferenceArgs']]] = None,
                 inbound_nat_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineInboundNatRuleArgs']]]]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 lab_subnet_name: Optional[pulumi.Input[str]] = None,
                 lab_virtual_network_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 ssh_key: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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

            __props__['allow_claim'] = allow_claim
            __props__['disallow_public_ip_address'] = disallow_public_ip_address
            if gallery_image_reference is None and not opts.urn:
                raise TypeError("Missing required property 'gallery_image_reference'")
            __props__['gallery_image_reference'] = gallery_image_reference
            __props__['inbound_nat_rules'] = inbound_nat_rules
            if lab_name is None and not opts.urn:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            if lab_subnet_name is None and not opts.urn:
                raise TypeError("Missing required property 'lab_subnet_name'")
            __props__['lab_subnet_name'] = lab_subnet_name
            if lab_virtual_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'lab_virtual_network_id'")
            __props__['lab_virtual_network_id'] = lab_virtual_network_id
            __props__['location'] = location
            __props__['name'] = name
            __props__['notes'] = notes
            __props__['password'] = password
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__['size'] = size
            __props__['ssh_key'] = ssh_key
            if storage_type is None and not opts.urn:
                raise TypeError("Missing required property 'storage_type'")
            __props__['storage_type'] = storage_type
            __props__['tags'] = tags
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
            __props__['fqdn'] = None
            __props__['unique_identifier'] = None
        super(LinuxVirtualMachine, __self__).__init__(
            'azure:devtest/linuxVirtualMachine:LinuxVirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_claim: Optional[pulumi.Input[bool]] = None,
            disallow_public_ip_address: Optional[pulumi.Input[bool]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            gallery_image_reference: Optional[pulumi.Input[pulumi.InputType['LinuxVirtualMachineGalleryImageReferenceArgs']]] = None,
            inbound_nat_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineInboundNatRuleArgs']]]]] = None,
            lab_name: Optional[pulumi.Input[str]] = None,
            lab_subnet_name: Optional[pulumi.Input[str]] = None,
            lab_virtual_network_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[str]] = None,
            ssh_key: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            unique_identifier: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'LinuxVirtualMachine':
        """
        Get an existing LinuxVirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_claim: Can this Virtual Machine be claimed by users? Defaults to `true`.
        :param pulumi.Input[bool] disallow_public_ip_address: Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
        :param pulumi.Input[str] fqdn: The FQDN of the Virtual Machine.
        :param pulumi.Input[pulumi.InputType['LinuxVirtualMachineGalleryImageReferenceArgs']] gallery_image_reference: A `gallery_image_reference` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LinuxVirtualMachineInboundNatRuleArgs']]]] inbound_nat_rules: One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_name: Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_subnet_name: The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] lab_virtual_network_id: The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notes: Any notes about the Virtual Machine.
        :param pulumi.Input[str] password: The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] size: The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ssh_key: The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_type: The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] unique_identifier: The unique immutable identifier of the Virtual Machine.
        :param pulumi.Input[str] username: The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_claim"] = allow_claim
        __props__["disallow_public_ip_address"] = disallow_public_ip_address
        __props__["fqdn"] = fqdn
        __props__["gallery_image_reference"] = gallery_image_reference
        __props__["inbound_nat_rules"] = inbound_nat_rules
        __props__["lab_name"] = lab_name
        __props__["lab_subnet_name"] = lab_subnet_name
        __props__["lab_virtual_network_id"] = lab_virtual_network_id
        __props__["location"] = location
        __props__["name"] = name
        __props__["notes"] = notes
        __props__["password"] = password
        __props__["resource_group_name"] = resource_group_name
        __props__["size"] = size
        __props__["ssh_key"] = ssh_key
        __props__["storage_type"] = storage_type
        __props__["tags"] = tags
        __props__["unique_identifier"] = unique_identifier
        __props__["username"] = username
        return LinuxVirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowClaim")
    def allow_claim(self) -> pulumi.Output[Optional[bool]]:
        """
        Can this Virtual Machine be claimed by users? Defaults to `true`.
        """
        return pulumi.get(self, "allow_claim")

    @property
    @pulumi.getter(name="disallowPublicIpAddress")
    def disallow_public_ip_address(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "disallow_public_ip_address")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The FQDN of the Virtual Machine.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="galleryImageReference")
    def gallery_image_reference(self) -> pulumi.Output['outputs.LinuxVirtualMachineGalleryImageReference']:
        """
        A `gallery_image_reference` block as defined below.
        """
        return pulumi.get(self, "gallery_image_reference")

    @property
    @pulumi.getter(name="inboundNatRules")
    def inbound_nat_rules(self) -> pulumi.Output[Optional[Sequence['outputs.LinuxVirtualMachineInboundNatRule']]]:
        """
        One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "inbound_nat_rules")

    @property
    @pulumi.getter(name="labName")
    def lab_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_name")

    @property
    @pulumi.getter(name="labSubnetName")
    def lab_subnet_name(self) -> pulumi.Output[str]:
        """
        The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_subnet_name")

    @property
    @pulumi.getter(name="labVirtualNetworkId")
    def lab_virtual_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "lab_virtual_network_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        Any notes about the Virtual Machine.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sshKey")
    def ssh_key(self) -> pulumi.Output[Optional[str]]:
        """
        The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ssh_key")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[str]:
        """
        The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[str]:
        """
        The unique immutable identifier of the Virtual Machine.
        """
        return pulumi.get(self, "unique_identifier")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "username")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

