# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Account(pulumi.CustomResource):
    active_directory: pulumi.Output[dict]
    """
    A `active_directory` block as defined below.

      * `dns_servers` (`list`) - A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
      * `domain` (`str`) - The name of the Active Directory domain.
      * `organizationalUnit` (`str`) - The Organizational Unit (OU) within the Active Directory Domain.
      * `password` (`str`) - The password associated with the `username`.
      * `smbServerName` (`str`) - The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
      * `username` (`str`) - The Username of Active Directory Domain Administrator.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the NetApp Account. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, active_directory=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a NetApp Account.

        > **NOTE:** Azure allows only one active directory can be joined to a single subscription at a time for NetApp Account.

        ## NetApp Account Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.netapp.Account("exampleAccount",
            active_directory={
                "dns_servers": ["1.2.3.4"],
                "domain": "westcentralus.com",
                "organizationalUnit": "OU=FirstLevel",
                "password": "aduserpwd",
                "smbServerName": "SMBSERVER",
                "username": "aduser",
            },
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] active_directory: A `active_directory` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **active_directory** object supports the following:

          * `dns_servers` (`pulumi.Input[list]`) - A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
          * `domain` (`pulumi.Input[str]`) - The name of the Active Directory domain.
          * `organizationalUnit` (`pulumi.Input[str]`) - The Organizational Unit (OU) within the Active Directory Domain.
          * `password` (`pulumi.Input[str]`) - The password associated with the `username`.
          * `smbServerName` (`pulumi.Input[str]`) - The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
          * `username` (`pulumi.Input[str]`) - The Username of Active Directory Domain Administrator.
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

            __props__['active_directory'] = active_directory
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Account, __self__).__init__(
            'azure:netapp/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, active_directory=None, location=None, name=None, resource_group_name=None, tags=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] active_directory: A `active_directory` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **active_directory** object supports the following:

          * `dns_servers` (`pulumi.Input[list]`) - A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
          * `domain` (`pulumi.Input[str]`) - The name of the Active Directory domain.
          * `organizationalUnit` (`pulumi.Input[str]`) - The Organizational Unit (OU) within the Active Directory Domain.
          * `password` (`pulumi.Input[str]`) - The password associated with the `username`.
          * `smbServerName` (`pulumi.Input[str]`) - The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
          * `username` (`pulumi.Input[str]`) - The Username of Active Directory Domain Administrator.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["active_directory"] = active_directory
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Account(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

