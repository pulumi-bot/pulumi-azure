# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, id=None, location=None, name=None, resource_group_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure Region where the NetApp Account exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name)


def get_account(name=None, resource_group_name=None, opts=None):
    """
    Uses this data source to access information about an existing NetApp Account.

    ## NetApp Account Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.netapp.get_account(resource_group_name="acctestRG",
        name="acctestnetappaccount")
    pulumi.export("netappAccountId", example.id)
    ```


    :param str name: The name of the NetApp Account.
    :param str resource_group_name: The Name of the Resource Group where the NetApp Account exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:netapp/getAccount:getAccount', __args__, opts=opts).value

    return AwaitableGetAccountResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'))
