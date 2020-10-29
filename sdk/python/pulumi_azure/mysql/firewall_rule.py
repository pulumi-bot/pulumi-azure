# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['FirewallRule']


class FirewallRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_ip_address: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 start_ip_address: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Firewall Rule for a MySQL Server.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_ip_address: Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_ip_address: Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
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

            if end_ip_address is None:
                raise TypeError("Missing required property 'end_ip_address'")
            __props__['end_ip_address'] = end_ip_address
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            if start_ip_address is None:
                raise TypeError("Missing required property 'start_ip_address'")
            __props__['start_ip_address'] = start_ip_address
        super(FirewallRule, __self__).__init__(
            'azure:mysql/firewallRule:FirewallRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            end_ip_address: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            start_ip_address: Optional[pulumi.Input[str]] = None) -> 'FirewallRule':
        """
        Get an existing FirewallRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_ip_address: Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_ip_address: Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["end_ip_address"] = end_ip_address
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["server_name"] = server_name
        __props__["start_ip_address"] = start_ip_address
        return FirewallRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endIpAddress")
    def end_ip_address(self) -> pulumi.Output[str]:
        """
        Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "end_ip_address")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="startIpAddress")
    def start_ip_address(self) -> pulumi.Output[str]:
        """
        Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "start_ip_address")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

