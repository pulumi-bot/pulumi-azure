# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class BackendAddressPool(pulumi.CustomResource):
    backend_ip_configurations: pulumi.Output[list]
    """
    The Backend IP Configurations associated with this Backend Address Pool.
    """
    load_balancing_rules: pulumi.Output[list]
    """
    The Load Balancing Rules associated with this Backend Address Pool.
    """
    loadbalancer_id: pulumi.Output[str]
    """
    The ID of the Load Balancer in which to create the Backend Address Pool.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Backend Address Pool.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the resource.
    """
    def __init__(__self__, resource_name, opts=None, loadbalancer_id=None, name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Load Balancer Backend Address Pool.

        > **NOTE:** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location="West US",
            resource_group_name=example_resource_group.name,
            allocation_method="Static")
        example_load_balancer = azure.lb.LoadBalancer("exampleLoadBalancer",
            location="West US",
            resource_group_name=example_resource_group.name,
            frontend_ip_configuration=[{
                "name": "PublicIPAddress",
                "public_ip_address_id": example_public_ip.id,
            }])
        example_backend_address_pool = azure.lb.BackendAddressPool("exampleBackendAddressPool",
            resource_group_name=example_resource_group.name,
            loadbalancer_id=example_load_balancer.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the Backend Address Pool.
        :param pulumi.Input[str] name: Specifies the name of the Backend Address Pool.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
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

            if loadbalancer_id is None:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__['loadbalancer_id'] = loadbalancer_id
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['backend_ip_configurations'] = None
            __props__['load_balancing_rules'] = None
        super(BackendAddressPool, __self__).__init__(
            'azure:lb/backendAddressPool:BackendAddressPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend_ip_configurations=None, load_balancing_rules=None, loadbalancer_id=None, name=None, resource_group_name=None):
        """
        Get an existing BackendAddressPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] backend_ip_configurations: The Backend IP Configurations associated with this Backend Address Pool.
        :param pulumi.Input[list] load_balancing_rules: The Load Balancing Rules associated with this Backend Address Pool.
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the Backend Address Pool.
        :param pulumi.Input[str] name: Specifies the name of the Backend Address Pool.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend_ip_configurations"] = backend_ip_configurations
        __props__["load_balancing_rules"] = load_balancing_rules
        __props__["loadbalancer_id"] = loadbalancer_id
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        return BackendAddressPool(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
