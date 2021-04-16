# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PoolArgs', 'Pool']

@pulumi.input_type
class PoolArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 service_level: pulumi.Input[str],
                 size_in_tb: pulumi.Input[int],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Pool resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[int] size_in_tb: Provisioned size of the pool in TB. Value must be between `4` and `500`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "service_level", service_level)
        pulumi.set(__self__, "size_in_tb", size_in_tb)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> pulumi.Input[str]:
        """
        The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        """
        return pulumi.get(self, "service_level")

    @service_level.setter
    def service_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_level", value)

    @property
    @pulumi.getter(name="sizeInTb")
    def size_in_tb(self) -> pulumi.Input[int]:
        """
        Provisioned size of the pool in TB. Value must be between `4` and `500`.
        """
        return pulumi.get(self, "size_in_tb")

    @size_in_tb.setter
    def size_in_tb(self, value: pulumi.Input[int]):
        pulumi.set(self, "size_in_tb", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NetApp Pool. Changing this forces a new resource to be created.
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


@pulumi.input_type
class _PoolState:
    def __init__(__self__, *,
                 account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_level: Optional[pulumi.Input[str]] = None,
                 size_in_tb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Pool resources.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[int] size_in_tb: Provisioned size of the pool in TB. Value must be between `4` and `500`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if service_level is not None:
            pulumi.set(__self__, "service_level", service_level)
        if size_in_tb is not None:
            pulumi.set(__self__, "size_in_tb", size_in_tb)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NetApp Pool. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> Optional[pulumi.Input[str]]:
        """
        The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        """
        return pulumi.get(self, "service_level")

    @service_level.setter
    def service_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_level", value)

    @property
    @pulumi.getter(name="sizeInTb")
    def size_in_tb(self) -> Optional[pulumi.Input[int]]:
        """
        Provisioned size of the pool in TB. Value must be between `4` and `500`.
        """
        return pulumi.get(self, "size_in_tb")

    @size_in_tb.setter
    def size_in_tb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_tb", value)

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


class Pool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_level: Optional[pulumi.Input[str]] = None,
                 size_in_tb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a Pool within a NetApp Account.

        ## NetApp Pool Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.netapp.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pool = azure.netapp.Pool("examplePool",
            account_name=example_account.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            service_level="Premium",
            size_in_tb=4)
        ```

        ## Import

        NetApp Pool can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:netapp/pool:Pool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[int] size_in_tb: Provisioned size of the pool in TB. Value must be between `4` and `500`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Pool within a NetApp Account.

        ## NetApp Pool Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.netapp.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pool = azure.netapp.Pool("examplePool",
            account_name=example_account.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            service_level="Premium",
            size_in_tb=4)
        ```

        ## Import

        NetApp Pool can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:netapp/pool:Pool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1
        ```

        :param str resource_name: The name of the resource.
        :param PoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_level: Optional[pulumi.Input[str]] = None,
                 size_in_tb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = PoolArgs.__new__(PoolArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if service_level is None and not opts.urn:
                raise TypeError("Missing required property 'service_level'")
            __props__.__dict__["service_level"] = service_level
            if size_in_tb is None and not opts.urn:
                raise TypeError("Missing required property 'size_in_tb'")
            __props__.__dict__["size_in_tb"] = size_in_tb
            __props__.__dict__["tags"] = tags
        super(Pool, __self__).__init__(
            'azure:netapp/pool:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_level: Optional[pulumi.Input[str]] = None,
            size_in_tb: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[int] size_in_tb: Provisioned size of the pool in TB. Value must be between `4` and `500`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PoolState.__new__(_PoolState)

        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["service_level"] = service_level
        __props__.__dict__["size_in_tb"] = size_in_tb
        __props__.__dict__["tags"] = tags
        return Pool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the NetApp Pool. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> pulumi.Output[str]:
        """
        The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter(name="sizeInTb")
    def size_in_tb(self) -> pulumi.Output[int]:
        """
        Provisioned size of the pool in TB. Value must be between `4` and `500`.
        """
        return pulumi.get(self, "size_in_tb")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

