# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VpnSiteArgs', 'VpnSite']

@pulumi.input_type
class VpnSiteArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 virtual_wan_id: pulumi.Input[str],
                 address_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_model: Optional[pulumi.Input[str]] = None,
                 device_vendor: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VpnSite resource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] virtual_wan_id: The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_cidrs: Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        :param pulumi.Input[str] device_model: The model of the VPN device.
        :param pulumi.Input[str] device_vendor: The name of the VPN device vendor.
        :param pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]] links: One or more `link` blocks as defined below.
        :param pulumi.Input[str] location: The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] name: The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the VPN Site.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "virtual_wan_id", virtual_wan_id)
        if address_cidrs is not None:
            pulumi.set(__self__, "address_cidrs", address_cidrs)
        if device_model is not None:
            pulumi.set(__self__, "device_model", device_model)
        if device_vendor is not None:
            pulumi.set(__self__, "device_vendor", device_vendor)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="virtualWanId")
    def virtual_wan_id(self) -> pulumi.Input[str]:
        """
        The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "virtual_wan_id")

    @virtual_wan_id.setter
    def virtual_wan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_wan_id", value)

    @property
    @pulumi.getter(name="addressCidrs")
    def address_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        """
        return pulumi.get(self, "address_cidrs")

    @address_cidrs.setter
    def address_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_cidrs", value)

    @property
    @pulumi.getter(name="deviceModel")
    def device_model(self) -> Optional[pulumi.Input[str]]:
        """
        The model of the VPN device.
        """
        return pulumi.get(self, "device_model")

    @device_model.setter
    def device_model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_model", value)

    @property
    @pulumi.getter(name="deviceVendor")
    def device_vendor(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPN device vendor.
        """
        return pulumi.get(self, "device_vendor")

    @device_vendor.setter
    def device_vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_vendor", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]]]:
        """
        One or more `link` blocks as defined below.
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the VPN Site.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VpnSiteState:
    def __init__(__self__, *,
                 address_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_model: Optional[pulumi.Input[str]] = None,
                 device_vendor: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_wan_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpnSite resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_cidrs: Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        :param pulumi.Input[str] device_model: The model of the VPN device.
        :param pulumi.Input[str] device_vendor: The name of the VPN device vendor.
        :param pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]] links: One or more `link` blocks as defined below.
        :param pulumi.Input[str] location: The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] name: The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the VPN Site.
        :param pulumi.Input[str] virtual_wan_id: The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        """
        if address_cidrs is not None:
            pulumi.set(__self__, "address_cidrs", address_cidrs)
        if device_model is not None:
            pulumi.set(__self__, "device_model", device_model)
        if device_vendor is not None:
            pulumi.set(__self__, "device_vendor", device_vendor)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_wan_id is not None:
            pulumi.set(__self__, "virtual_wan_id", virtual_wan_id)

    @property
    @pulumi.getter(name="addressCidrs")
    def address_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        """
        return pulumi.get(self, "address_cidrs")

    @address_cidrs.setter
    def address_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_cidrs", value)

    @property
    @pulumi.getter(name="deviceModel")
    def device_model(self) -> Optional[pulumi.Input[str]]:
        """
        The model of the VPN device.
        """
        return pulumi.get(self, "device_model")

    @device_model.setter
    def device_model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_model", value)

    @property
    @pulumi.getter(name="deviceVendor")
    def device_vendor(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPN device vendor.
        """
        return pulumi.get(self, "device_vendor")

    @device_vendor.setter
    def device_vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_vendor", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]]]:
        """
        One or more `link` blocks as defined below.
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnSiteLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the VPN Site.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="virtualWanId")
    def virtual_wan_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "virtual_wan_id")

    @virtual_wan_id.setter
    def virtual_wan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_wan_id", value)


class VpnSite(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_model: Optional[pulumi.Input[str]] = None,
                 device_vendor: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnSiteLinkArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_wan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a VPN Site.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_wan = azure.network.VirtualWan("exampleVirtualWan",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_vpn_site = azure.network.VpnSite("exampleVpnSite",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            virtual_wan_id=example_virtual_wan.id,
            links=[azure.network.VpnSiteLinkArgs(
                name="link1",
                ip_address="10.0.0.1",
            )])
        ```

        ## Import

        VPN Sites can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/vpnSite:VpnSite example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnSites/site1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_cidrs: Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        :param pulumi.Input[str] device_model: The model of the VPN device.
        :param pulumi.Input[str] device_vendor: The name of the VPN device vendor.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnSiteLinkArgs']]]] links: One or more `link` blocks as defined below.
        :param pulumi.Input[str] location: The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] name: The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the VPN Site.
        :param pulumi.Input[str] virtual_wan_id: The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnSiteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a VPN Site.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_wan = azure.network.VirtualWan("exampleVirtualWan",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_vpn_site = azure.network.VpnSite("exampleVpnSite",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            virtual_wan_id=example_virtual_wan.id,
            links=[azure.network.VpnSiteLinkArgs(
                name="link1",
                ip_address="10.0.0.1",
            )])
        ```

        ## Import

        VPN Sites can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/vpnSite:VpnSite example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnSites/site1
        ```

        :param str resource_name: The name of the resource.
        :param VpnSiteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnSiteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_model: Optional[pulumi.Input[str]] = None,
                 device_vendor: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnSiteLinkArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_wan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpnSiteArgs.__new__(VpnSiteArgs)

            __props__.__dict__["address_cidrs"] = address_cidrs
            __props__.__dict__["device_model"] = device_model
            __props__.__dict__["device_vendor"] = device_vendor
            __props__.__dict__["links"] = links
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if virtual_wan_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_wan_id'")
            __props__.__dict__["virtual_wan_id"] = virtual_wan_id
        super(VpnSite, __self__).__init__(
            'azure:network/vpnSite:VpnSite',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            device_model: Optional[pulumi.Input[str]] = None,
            device_vendor: Optional[pulumi.Input[str]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnSiteLinkArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            virtual_wan_id: Optional[pulumi.Input[str]] = None) -> 'VpnSite':
        """
        Get an existing VpnSite resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_cidrs: Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        :param pulumi.Input[str] device_model: The model of the VPN device.
        :param pulumi.Input[str] device_vendor: The name of the VPN device vendor.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnSiteLinkArgs']]]] links: One or more `link` blocks as defined below.
        :param pulumi.Input[str] location: The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] name: The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the VPN Site.
        :param pulumi.Input[str] virtual_wan_id: The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnSiteState.__new__(_VpnSiteState)

        __props__.__dict__["address_cidrs"] = address_cidrs
        __props__.__dict__["device_model"] = device_model
        __props__.__dict__["device_vendor"] = device_vendor
        __props__.__dict__["links"] = links
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["virtual_wan_id"] = virtual_wan_id
        return VpnSite(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressCidrs")
    def address_cidrs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
        """
        return pulumi.get(self, "address_cidrs")

    @property
    @pulumi.getter(name="deviceModel")
    def device_model(self) -> pulumi.Output[Optional[str]]:
        """
        The model of the VPN device.
        """
        return pulumi.get(self, "device_model")

    @property
    @pulumi.getter(name="deviceVendor")
    def device_vendor(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the VPN device vendor.
        """
        return pulumi.get(self, "device_vendor")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Optional[Sequence['outputs.VpnSiteLink']]]:
        """
        One or more `link` blocks as defined below.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the VPN Site.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualWanId")
    def virtual_wan_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
        """
        return pulumi.get(self, "virtual_wan_id")

