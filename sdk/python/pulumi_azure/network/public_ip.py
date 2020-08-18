# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['PublicIp']


class PublicIp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_method: Optional[pulumi.Input[str]] = None,
                 domain_name_label: Optional[pulumi.Input[str]] = None,
                 idle_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_ip_prefix_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 reverse_fqdn: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Public IP Address.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            allocation_method="Static",
            tags={
                "environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_method: Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
        :param pulumi.Input[str] domain_name_label: Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
        :param pulumi.Input[float] idle_timeout_in_minutes: Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
        :param pulumi.Input[str] ip_version: The IP Version to use, IPv6 or IPv4.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Public IP resource . Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] public_ip_prefix_id: If specified then public IP address allocated will be provided from the public IP prefix resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the public ip.
        :param pulumi.Input[str] reverse_fqdn: A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
        :param pulumi.Input[str] sku: The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A collection containing the availability zone to allocate the Public IP in.
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

            if allocation_method is None:
                raise TypeError("Missing required property 'allocation_method'")
            __props__['allocation_method'] = allocation_method
            __props__['domain_name_label'] = domain_name_label
            __props__['idle_timeout_in_minutes'] = idle_timeout_in_minutes
            __props__['ip_version'] = ip_version
            __props__['location'] = location
            __props__['name'] = name
            __props__['public_ip_prefix_id'] = public_ip_prefix_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['reverse_fqdn'] = reverse_fqdn
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['fqdn'] = None
            __props__['ip_address'] = None
        super(PublicIp, __self__).__init__(
            'azure:network/publicIp:PublicIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_method: Optional[pulumi.Input[str]] = None,
            domain_name_label: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            idle_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            ip_version: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_ip_prefix_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            reverse_fqdn: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            zones: Optional[pulumi.Input[str]] = None) -> 'PublicIp':
        """
        Get an existing PublicIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_method: Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
        :param pulumi.Input[str] domain_name_label: Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
        :param pulumi.Input[str] fqdn: Fully qualified domain name of the A DNS record associated with the public IP. `domain_name_label` must be specified to get the `fqdn`. This is the concatenation of the `domain_name_label` and the regionalized DNS zone
        :param pulumi.Input[float] idle_timeout_in_minutes: Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
        :param pulumi.Input[str] ip_address: The IP address value that was allocated.
        :param pulumi.Input[str] ip_version: The IP Version to use, IPv6 or IPv4.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Public IP resource . Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] public_ip_prefix_id: If specified then public IP address allocated will be provided from the public IP prefix resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the public ip.
        :param pulumi.Input[str] reverse_fqdn: A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
        :param pulumi.Input[str] sku: The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A collection containing the availability zone to allocate the Public IP in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocation_method"] = allocation_method
        __props__["domain_name_label"] = domain_name_label
        __props__["fqdn"] = fqdn
        __props__["idle_timeout_in_minutes"] = idle_timeout_in_minutes
        __props__["ip_address"] = ip_address
        __props__["ip_version"] = ip_version
        __props__["location"] = location
        __props__["name"] = name
        __props__["public_ip_prefix_id"] = public_ip_prefix_id
        __props__["resource_group_name"] = resource_group_name
        __props__["reverse_fqdn"] = reverse_fqdn
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["zones"] = zones
        return PublicIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationMethod")
    def allocation_method(self) -> str:
        """
        Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
        """
        return pulumi.get(self, "allocation_method")

    @property
    @pulumi.getter(name="domainNameLabel")
    def domain_name_label(self) -> Optional[str]:
        """
        Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
        """
        return pulumi.get(self, "domain_name_label")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        Fully qualified domain name of the A DNS record associated with the public IP. `domain_name_label` must be specified to get the `fqdn`. This is the concatenation of the `domain_name_label` and the regionalized DNS zone
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="idleTimeoutInMinutes")
    def idle_timeout_in_minutes(self) -> Optional[float]:
        """
        Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
        """
        return pulumi.get(self, "idle_timeout_in_minutes")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The IP address value that was allocated.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[str]:
        """
        The IP Version to use, IPv6 or IPv4.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the Public IP resource . Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIpPrefixId")
    def public_ip_prefix_id(self) -> Optional[str]:
        """
        If specified then public IP address allocated will be provided from the public IP prefix resource.
        """
        return pulumi.get(self, "public_ip_prefix_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to
        create the public ip.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="reverseFqdn")
    def reverse_fqdn(self) -> Optional[str]:
        """
        A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
        """
        return pulumi.get(self, "reverse_fqdn")

    @property
    @pulumi.getter
    def sku(self) -> Optional[str]:
        """
        The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> Optional[str]:
        """
        A collection containing the availability zone to allocate the Public IP in.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

