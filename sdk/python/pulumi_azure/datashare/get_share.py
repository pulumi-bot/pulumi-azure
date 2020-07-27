# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetShareResult',
    'AwaitableGetShareResult',
    'get_share',
]


class GetShareResult:
    """
    A collection of values returned by getShare.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, account_id=None, description=None, id=None, kind=None, name=None, snapshot_schedules=None, terms=None) -> None:
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        __self__.account_id = account_id
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of the Data Share.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        __self__.kind = kind
        """
        The kind of the Data Share.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the snapshot schedule.
        """
        if snapshot_schedules and not isinstance(snapshot_schedules, list):
            raise TypeError("Expected argument 'snapshot_schedules' to be a list")
        __self__.snapshot_schedules = snapshot_schedules
        """
        A `snapshot_schedule` block as defined below.
        """
        if terms and not isinstance(terms, str):
            raise TypeError("Expected argument 'terms' to be a str")
        __self__.terms = terms
        """
        The terms of the Data Share.
        """


class AwaitableGetShareResult(GetShareResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetShareResult(
            account_id=self.account_id,
            description=self.description,
            id=self.id,
            kind=self.kind,
            name=self.name,
            snapshot_schedules=self.snapshot_schedules,
            terms=self.terms)


def get_share(account_id: Optional[str] = None, name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetShareResult:
    """
    Use this data source to access information about an existing Data Share.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_account = azure.datashare.get_account(name="example-account",
        resource_group_name="example-resource-group")
    example_share = azure.datashare.get_share(name="existing",
        account_id=data["azurerm_data_share_account"]["exmaple"]["id"])
    pulumi.export("id", example_share.id)
    ```


    :param str account_id: The ID of the Data Share account in which the Data Share is created.
    :param str name: The name of this Data Share.
    """
    __args__ = dict()
    __args__['accountId'] = account_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:datashare/getShare:getShare', __args__, opts=opts).value

    return AwaitableGetShareResult(
        account_id=__ret__.get('accountId'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        kind=__ret__.get('kind'),
        name=__ret__.get('name'),
        snapshot_schedules=__ret__.get('snapshotSchedules'),
        terms=__ret__.get('terms'))
