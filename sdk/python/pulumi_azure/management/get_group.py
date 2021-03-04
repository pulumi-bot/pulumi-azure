# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, display_name=None, group_id=None, id=None, name=None, parent_management_group_id=None, subscription_ids=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        if group_id is not None:
            warnings.warn("""Deprecated in favour of `name`""", DeprecationWarning)
            pulumi.log.warn("""group_id is deprecated: Deprecated in favour of `name`""")

        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_management_group_id and not isinstance(parent_management_group_id, str):
            raise TypeError("Expected argument 'parent_management_group_id' to be a str")
        pulumi.set(__self__, "parent_management_group_id", parent_management_group_id)
        if subscription_ids and not isinstance(subscription_ids, list):
            raise TypeError("Expected argument 'subscription_ids' to be a list")
        pulumi.set(__self__, "subscription_ids", subscription_ids)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentManagementGroupId")
    def parent_management_group_id(self) -> str:
        """
        The ID of any Parent Management Group.
        """
        return pulumi.get(self, "parent_management_group_id")

    @property
    @pulumi.getter(name="subscriptionIds")
    def subscription_ids(self) -> Sequence[str]:
        """
        A list of Subscription IDs which are assigned to the Management Group.
        """
        return pulumi.get(self, "subscription_ids")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            display_name=self.display_name,
            group_id=self.group_id,
            id=self.id,
            name=self.name,
            parent_management_group_id=self.parent_management_group_id,
            subscription_ids=self.subscription_ids)


def get_group(display_name: Optional[str] = None,
              group_id: Optional[str] = None,
              name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Use this data source to access information about an existing Management Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.management.get_group(name="00000000-0000-0000-0000-000000000000")
    pulumi.export("displayName", example.display_name)
    ```


    :param str display_name: Specifies the display name of this Management Group.
    :param str group_id: Specifies the name or UUID of this Management Group.
    :param str name: Specifies the name or UUID of this Management Group.
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['groupId'] = group_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:management/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        display_name=__ret__.display_name,
        group_id=__ret__.group_id,
        id=__ret__.id,
        name=__ret__.name,
        parent_management_group_id=__ret__.parent_management_group_id,
        subscription_ids=__ret__.subscription_ids)
