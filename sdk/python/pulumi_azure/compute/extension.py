# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Extension(pulumi.CustomResource):
    auto_upgrade_minor_version: pulumi.Output[bool]
    """
    Specifies if the platform deploys
    the latest minor version update to the `type_handler_version` specified.
    """
    name: pulumi.Output[str]
    """
    The name of the virtual machine extension peering. Changing
    this forces a new resource to be created.
    """
    protected_settings: pulumi.Output[str]
    """
    The protected_settings passed to the
    extension, like settings, these are specified as a JSON object in a string.
    """
    publisher: pulumi.Output[str]
    """
    The publisher of the extension, available publishers
    can be found by using the Azure CLI.
    """
    settings: pulumi.Output[str]
    """
    The settings passed to the extension, these are
    specified as a JSON object in a string.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    type: pulumi.Output[str]
    """
    The type of extension, available types for a publisher can
    be found using the Azure CLI.
    """
    type_handler_version: pulumi.Output[str]
    """
    Specifies the version of the extension to
    use, available versions can be found using the Azure CLI.
    """
    virtual_machine_id: pulumi.Output[str]
    """
    The ID of the Virtual Machine. Changing this forces a new resource to be created
    """
    def __init__(__self__, resource_name, opts=None, auto_upgrade_minor_version=None, name=None, protected_settings=None, publisher=None, settings=None, tags=None, type=None, type_handler_version=None, virtual_machine_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Virtual Machine Extension to provide post deployment configuration
        and run automated tasks.

        > **NOTE:** Custom Script Extensions for Linux & Windows require that the `commandToExecute` returns a `0` exit code to be classified as successfully deployed. You can achieve this by appending `exit 0` to the end of your `commandToExecute`.

        > **NOTE:** Custom Script Extensions require that the Azure Virtual Machine Guest Agent is running on the Virtual Machine.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
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
            ip_configuration=[{
                "name": "testconfiguration1",
                "subnet_id": example_subnet.id,
                "privateIpAddressAllocation": "Dynamic",
            }])
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            tags={
                "environment": "staging",
            })
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_virtual_machine = azure.compute.VirtualMachine("exampleVirtualMachine",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            network_interface_ids=[example_network_interface.id],
            vm_size="Standard_F2",
            storage_image_reference={
                "publisher": "Canonical",
                "offer": "UbuntuServer",
                "sku": "16.04-LTS",
                "version": "latest",
            },
            storage_os_disk={
                "name": "myosdisk1",
                "vhdUri": pulumi.Output.all(example_account.primary_blob_endpoint, example_container.name).apply(lambda primary_blob_endpoint, name: f"{primary_blob_endpoint}{name}/myosdisk1.vhd"),
                "caching": "ReadWrite",
                "create_option": "FromImage",
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
        example_extension = azure.compute.Extension("exampleExtension",
            virtual_machine_id=example_virtual_machine.id,
            publisher="Microsoft.Azure.Extensions",
            type="CustomScript",
            type_handler_version="2.0",
            settings=\"\"\"	{
        		"commandToExecute": "hostname && uptime"
        	}
        \"\"\",
            tags={
                "environment": "Production",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade_minor_version: Specifies if the platform deploys
               the latest minor version update to the `type_handler_version` specified.
        :param pulumi.Input[str] name: The name of the virtual machine extension peering. Changing
               this forces a new resource to be created.
        :param pulumi.Input[str] protected_settings: The protected_settings passed to the
               extension, like settings, these are specified as a JSON object in a string.
        :param pulumi.Input[str] publisher: The publisher of the extension, available publishers
               can be found by using the Azure CLI.
        :param pulumi.Input[str] settings: The settings passed to the extension, these are
               specified as a JSON object in a string.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of extension, available types for a publisher can
               be found using the Azure CLI.
        :param pulumi.Input[str] type_handler_version: Specifies the version of the extension to
               use, available versions can be found using the Azure CLI.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created
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

            __props__['auto_upgrade_minor_version'] = auto_upgrade_minor_version
            __props__['name'] = name
            __props__['protected_settings'] = protected_settings
            if publisher is None:
                raise TypeError("Missing required property 'publisher'")
            __props__['publisher'] = publisher
            __props__['settings'] = settings
            __props__['tags'] = tags
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            if type_handler_version is None:
                raise TypeError("Missing required property 'type_handler_version'")
            __props__['type_handler_version'] = type_handler_version
            if virtual_machine_id is None:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__['virtual_machine_id'] = virtual_machine_id
        super(Extension, __self__).__init__(
            'azure:compute/extension:Extension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_upgrade_minor_version=None, name=None, protected_settings=None, publisher=None, settings=None, tags=None, type=None, type_handler_version=None, virtual_machine_id=None):
        """
        Get an existing Extension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade_minor_version: Specifies if the platform deploys
               the latest minor version update to the `type_handler_version` specified.
        :param pulumi.Input[str] name: The name of the virtual machine extension peering. Changing
               this forces a new resource to be created.
        :param pulumi.Input[str] protected_settings: The protected_settings passed to the
               extension, like settings, these are specified as a JSON object in a string.
        :param pulumi.Input[str] publisher: The publisher of the extension, available publishers
               can be found by using the Azure CLI.
        :param pulumi.Input[str] settings: The settings passed to the extension, these are
               specified as a JSON object in a string.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of extension, available types for a publisher can
               be found using the Azure CLI.
        :param pulumi.Input[str] type_handler_version: Specifies the version of the extension to
               use, available versions can be found using the Azure CLI.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_upgrade_minor_version"] = auto_upgrade_minor_version
        __props__["name"] = name
        __props__["protected_settings"] = protected_settings
        __props__["publisher"] = publisher
        __props__["settings"] = settings
        __props__["tags"] = tags
        __props__["type"] = type
        __props__["type_handler_version"] = type_handler_version
        __props__["virtual_machine_id"] = virtual_machine_id
        return Extension(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
