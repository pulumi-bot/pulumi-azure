# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Group']


class Group(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_management_group_id: Optional[pulumi.Input[str]] = None,
                 subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Management Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_subscription()
        example_parent = azure.management.Group("exampleParent",
            display_name="ParentGroup",
            subscription_ids=[current.subscription_id])
        example_child = azure.management.Group("exampleChild",
            display_name="ChildGroup",
            parent_management_group_id=example_parent.id,
            subscription_ids=[current.subscription_id])
        # other subscription IDs can go here
        ```

        ## Import

        Management Groups can be imported using the `management group resource id`, e.g.

        ```sh
         $ pulumi import azure:management/group:Group example /providers/Microsoft.Management/managementGroups/group1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
        :param pulumi.Input[str] group_id: The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        :param pulumi.Input[str] parent_management_group_id: The ID of the Parent Management Group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subscription_ids: A list of Subscription GUIDs which should be assigned to the Management Group.
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

            __props__['display_name'] = display_name
            if group_id is not None and not opts.urn:
                warnings.warn("""Deprecated in favour of `name`""", DeprecationWarning)
                pulumi.log.warn("group_id is deprecated: Deprecated in favour of `name`")
            __props__['group_id'] = group_id
            __props__['name'] = name
            __props__['parent_management_group_id'] = parent_management_group_id
            __props__['subscription_ids'] = subscription_ids
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:managementgroups/managementGroup:ManagementGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Group, __self__).__init__(
            'azure:management/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_management_group_id: Optional[pulumi.Input[str]] = None,
            subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
        :param pulumi.Input[str] group_id: The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        :param pulumi.Input[str] parent_management_group_id: The ID of the Parent Management Group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subscription_ids: A list of Subscription GUIDs which should be assigned to the Management Group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["display_name"] = display_name
        __props__["group_id"] = group_id
        __props__["name"] = name
        __props__["parent_management_group_id"] = parent_management_group_id
        __props__["subscription_ids"] = subscription_ids
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentManagementGroupId")
    def parent_management_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the Parent Management Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "parent_management_group_id")

    @property
    @pulumi.getter(name="subscriptionIds")
    def subscription_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of Subscription GUIDs which should be assigned to the Management Group.
        """
        return pulumi.get(self, "subscription_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

