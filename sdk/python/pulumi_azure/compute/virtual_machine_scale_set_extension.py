# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VirtualMachineScaleSetExtension']


class VirtualMachineScaleSetExtension(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_upgrade_minor_version: Optional[pulumi.Input[bool]] = None,
                 force_update_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protected_settings: Optional[pulumi.Input[str]] = None,
                 provision_after_extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 publisher: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 type_handler_version: Optional[pulumi.Input[str]] = None,
                 virtual_machine_scale_set_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Extension for a Virtual Machine Scale Set.

        > **NOTE:** This resource is not intended to be used with the `compute.ScaleSet` resource - instead it's intended for this to be used with the `compute.LinuxVirtualMachineScaleSet` and `compute.WindowsVirtualMachineScaleSet` resources.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_azure as azure

        example_linux_virtual_machine_scale_set = azure.compute.LinuxVirtualMachineScaleSet("exampleLinuxVirtualMachineScaleSet")
        #...
        example_virtual_machine_scale_set_extension = azure.compute.VirtualMachineScaleSetExtension("exampleVirtualMachineScaleSetExtension",
            virtual_machine_scale_set_id=example_linux_virtual_machine_scale_set.id,
            publisher="Microsoft.Azure.Extensions",
            type="CustomScript",
            type_handler_version="2.0",
            settings=json.dumps({
                "commandToExecute": "echo $HOSTNAME",
            }))
        ```

        ## Import

        Virtual Machine Scale Set Extensions can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade_minor_version: Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        :param pulumi.Input[str] force_update_tag: A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        :param pulumi.Input[str] name: The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protected_settings: A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] provision_after_extensions: An ordered list of Extension names which this should be provisioned after.
        :param pulumi.Input[str] publisher: Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] settings: A JSON String which specifies Settings for the Extension.
        :param pulumi.Input[str] type: Specifies the Type of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type_handler_version: Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        :param pulumi.Input[str] virtual_machine_scale_set_id: The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
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

            __props__['auto_upgrade_minor_version'] = auto_upgrade_minor_version
            __props__['force_update_tag'] = force_update_tag
            __props__['name'] = name
            __props__['protected_settings'] = protected_settings
            __props__['provision_after_extensions'] = provision_after_extensions
            if publisher is None:
                raise TypeError("Missing required property 'publisher'")
            __props__['publisher'] = publisher
            __props__['settings'] = settings
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            if type_handler_version is None:
                raise TypeError("Missing required property 'type_handler_version'")
            __props__['type_handler_version'] = type_handler_version
            if virtual_machine_scale_set_id is None:
                raise TypeError("Missing required property 'virtual_machine_scale_set_id'")
            __props__['virtual_machine_scale_set_id'] = virtual_machine_scale_set_id
        super(VirtualMachineScaleSetExtension, __self__).__init__(
            'azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_upgrade_minor_version: Optional[pulumi.Input[bool]] = None,
            force_update_tag: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protected_settings: Optional[pulumi.Input[str]] = None,
            provision_after_extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            publisher: Optional[pulumi.Input[str]] = None,
            settings: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            type_handler_version: Optional[pulumi.Input[str]] = None,
            virtual_machine_scale_set_id: Optional[pulumi.Input[str]] = None) -> 'VirtualMachineScaleSetExtension':
        """
        Get an existing VirtualMachineScaleSetExtension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade_minor_version: Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        :param pulumi.Input[str] force_update_tag: A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        :param pulumi.Input[str] name: The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protected_settings: A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] provision_after_extensions: An ordered list of Extension names which this should be provisioned after.
        :param pulumi.Input[str] publisher: Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] settings: A JSON String which specifies Settings for the Extension.
        :param pulumi.Input[str] type: Specifies the Type of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type_handler_version: Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        :param pulumi.Input[str] virtual_machine_scale_set_id: The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_upgrade_minor_version"] = auto_upgrade_minor_version
        __props__["force_update_tag"] = force_update_tag
        __props__["name"] = name
        __props__["protected_settings"] = protected_settings
        __props__["provision_after_extensions"] = provision_after_extensions
        __props__["publisher"] = publisher
        __props__["settings"] = settings
        __props__["type"] = type
        __props__["type_handler_version"] = type_handler_version
        __props__["virtual_machine_scale_set_id"] = virtual_machine_scale_set_id
        return VirtualMachineScaleSetExtension(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoUpgradeMinorVersion")
    def auto_upgrade_minor_version(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        """
        return pulumi.get(self, "auto_upgrade_minor_version")

    @property
    @pulumi.getter(name="forceUpdateTag")
    def force_update_tag(self) -> pulumi.Output[Optional[str]]:
        """
        A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        """
        return pulumi.get(self, "force_update_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectedSettings")
    def protected_settings(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        """
        return pulumi.get(self, "protected_settings")

    @property
    @pulumi.getter(name="provisionAfterExtensions")
    def provision_after_extensions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An ordered list of Extension names which this should be provisioned after.
        """
        return pulumi.get(self, "provision_after_extensions")

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Output[str]:
        """
        Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON String which specifies Settings for the Extension.
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the Type of the Extension. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="typeHandlerVersion")
    def type_handler_version(self) -> pulumi.Output[str]:
        """
        Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        """
        return pulumi.get(self, "type_handler_version")

    @property
    @pulumi.getter(name="virtualMachineScaleSetId")
    def virtual_machine_scale_set_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_scale_set_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

