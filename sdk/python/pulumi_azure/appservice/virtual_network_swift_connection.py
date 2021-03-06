# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VirtualNetworkSwiftConnection(pulumi.CustomResource):
    app_service_id: pulumi.Output[str]
    """
    The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
    """
    subnet_id: pulumi.Output[str]
    """
    The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
    """
    def __init__(__self__, resource_name, opts=None, app_service_id=None, subnet_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an App Service Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        test_resource_group = azure.core.ResourceGroup("testResourceGroup", location="uksouth")
        test_virtual_network = azure.network.VirtualNetwork("testVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=test_resource_group.location,
            resource_group_name=test_resource_group.name)
        test1 = azure.network.Subnet("test1",
            resource_group_name=test_resource_group.name,
            virtual_network_name=test_virtual_network.name,
            address_prefix="10.0.1.0/24",
            delegation=[{
                "name": "acctestdelegation",
                "service_delegation": {
                    "name": "Microsoft.Web/serverFarms",
                    "actions": ["Microsoft.Network/virtualNetworks/subnets/action"],
                },
            }])
        test_plan = azure.appservice.Plan("testPlan",
            location=test_resource_group.location,
            resource_group_name=test_resource_group.name,
            sku={
                "tier": "Standard",
                "size": "S1",
            })
        test_app_service = azure.appservice.AppService("testAppService",
            location=test_resource_group.location,
            resource_group_name=test_resource_group.name,
            app_service_plan_id=test_plan.id)
        test_virtual_network_swift_connection = azure.appservice.VirtualNetworkSwiftConnection("testVirtualNetworkSwiftConnection",
            app_service_id=test_app_service.id,
            subnet_id=test1.id)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_id: The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
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

            if app_service_id is None:
                raise TypeError("Missing required property 'app_service_id'")
            __props__['app_service_id'] = app_service_id
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
        super(VirtualNetworkSwiftConnection, __self__).__init__(
            'azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_service_id=None, subnet_id=None):
        """
        Get an existing VirtualNetworkSwiftConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_id: The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_service_id"] = app_service_id
        __props__["subnet_id"] = subnet_id
        return VirtualNetworkSwiftConnection(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

