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

__all__ = ['SrvRecordArgs', 'SrvRecord']

@pulumi.input_type
class SrvRecordArgs:
    def __init__(__self__, *,
                 records: pulumi.Input[Sequence[pulumi.Input['SrvRecordRecordArgs']]],
                 resource_group_name: pulumi.Input[str],
                 ttl: pulumi.Input[int],
                 zone_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SrvRecord resource.
        :param pulumi.Input[Sequence[pulumi.Input['SrvRecordRecordArgs']]] records: A list of values that make up the SRV record. Each `record` block supports fields documented below.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[int] ttl: The Time To Live (TTL) of the DNS record in seconds.
        :param pulumi.Input[str] zone_name: Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the DNS SRV Record.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "records", records)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "ttl", ttl)
        pulumi.set(__self__, "zone_name", zone_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def records(self) -> pulumi.Input[Sequence[pulumi.Input['SrvRecordRecordArgs']]]:
        """
        A list of values that make up the SRV record. Each `record` block supports fields documented below.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: pulumi.Input[Sequence[pulumi.Input['SrvRecordRecordArgs']]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[int]:
        """
        The Time To Live (TTL) of the DNS record in seconds.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[int]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Input[str]:
        """
        Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the DNS SRV Record.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class SrvRecord(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SrvRecordRecordArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Enables you to manage DNS SRV Records within Azure DNS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_zone = azure.dns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        example_srv_record = azure.dns.SrvRecord("exampleSrvRecord",
            zone_name=example_zone.name,
            resource_group_name=example_resource_group.name,
            ttl=300,
            records=[azure.dns.SrvRecordRecordArgs(
                priority=1,
                weight=5,
                port=8080,
                target="target1.contoso.com",
            )],
            tags={
                "Environment": "Production",
            })
        ```

        ## Import

        SRV records can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dns/srvRecord:SrvRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/SRV/myrecord1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the DNS SRV Record.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SrvRecordRecordArgs']]]] records: A list of values that make up the SRV record. Each `record` block supports fields documented below.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[int] ttl: The Time To Live (TTL) of the DNS record in seconds.
        :param pulumi.Input[str] zone_name: Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SrvRecordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enables you to manage DNS SRV Records within Azure DNS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_zone = azure.dns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        example_srv_record = azure.dns.SrvRecord("exampleSrvRecord",
            zone_name=example_zone.name,
            resource_group_name=example_resource_group.name,
            ttl=300,
            records=[azure.dns.SrvRecordRecordArgs(
                priority=1,
                weight=5,
                port=8080,
                target="target1.contoso.com",
            )],
            tags={
                "Environment": "Production",
            })
        ```

        ## Import

        SRV records can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dns/srvRecord:SrvRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/SRV/myrecord1
        ```

        :param str resource_name: The name of the resource.
        :param SrvRecordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SrvRecordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SrvRecordRecordArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None,
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

            __props__['name'] = name
            if records is None and not opts.urn:
                raise TypeError("Missing required property 'records'")
            __props__['records'] = records
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if ttl is None and not opts.urn:
                raise TypeError("Missing required property 'ttl'")
            __props__['ttl'] = ttl
            if zone_name is None and not opts.urn:
                raise TypeError("Missing required property 'zone_name'")
            __props__['zone_name'] = zone_name
            __props__['fqdn'] = None
        super(SrvRecord, __self__).__init__(
            'azure:dns/srvRecord:SrvRecord',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SrvRecordRecordArgs']]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            zone_name: Optional[pulumi.Input[str]] = None) -> 'SrvRecord':
        """
        Get an existing SrvRecord resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fqdn: The FQDN of the DNS SRV Record.
        :param pulumi.Input[str] name: The name of the DNS SRV Record.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SrvRecordRecordArgs']]]] records: A list of values that make up the SRV record. Each `record` block supports fields documented below.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[int] ttl: The Time To Live (TTL) of the DNS record in seconds.
        :param pulumi.Input[str] zone_name: Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["fqdn"] = fqdn
        __props__["name"] = name
        __props__["records"] = records
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["ttl"] = ttl
        __props__["zone_name"] = zone_name
        return SrvRecord(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The FQDN of the DNS SRV Record.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the DNS SRV Record.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def records(self) -> pulumi.Output[Sequence['outputs.SrvRecordRecord']]:
        """
        A list of values that make up the SRV record. Each `record` block supports fields documented below.
        """
        return pulumi.get(self, "records")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        The Time To Live (TTL) of the DNS record in seconds.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Output[str]:
        """
        Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zone_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

