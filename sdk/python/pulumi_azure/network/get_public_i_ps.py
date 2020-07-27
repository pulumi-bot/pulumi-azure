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
    'GetPublicIPsResult',
    'AwaitableGetPublicIPsResult',
    'get_public_i_ps',
]


class GetPublicIPsResult:
    """
    A collection of values returned by getPublicIPs.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, allocation_type=None, attached=None, id=None, name_prefix=None, public_ips=None, resource_group_name=None) -> None:
        if allocation_type and not isinstance(allocation_type, str):
            raise TypeError("Expected argument 'allocation_type' to be a str")
        __self__.allocation_type = allocation_type
        if attached and not isinstance(attached, bool):
            raise TypeError("Expected argument 'attached' to be a bool")
        __self__.attached = attached
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        __self__.name_prefix = name_prefix
        if public_ips and not isinstance(public_ips, list):
            raise TypeError("Expected argument 'public_ips' to be a list")
        __self__.public_ips = public_ips
        """
        A List of `public_ips` blocks as defined below filtered by the criteria above.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name


class AwaitableGetPublicIPsResult(GetPublicIPsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublicIPsResult(
            allocation_type=self.allocation_type,
            attached=self.attached,
            id=self.id,
            name_prefix=self.name_prefix,
            public_ips=self.public_ips,
            resource_group_name=self.resource_group_name)


def get_public_i_ps(allocation_type: Optional[str] = None, attached: Optional[bool] = None, name_prefix: Optional[str] = None, resource_group_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublicIPsResult:
    """
    Use this data source to access information about a set of existing Public IP Addresses.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_public_i_ps(attached=False,
        resource_group_name="pip-test")
    ```


    :param str allocation_type: The Allocation Type for the Public IP Address. Possible values include `Static` or `Dynamic`.
    :param bool attached: Filter to include IP Addresses which are attached to a device, such as a VM/LB (`true`) or unattached (`false`).
    :param str name_prefix: A prefix match used for the IP Addresses `name` field, case sensitive.
    :param str resource_group_name: Specifies the name of the resource group.
    """
    __args__ = dict()
    __args__['allocationType'] = allocation_type
    __args__['attached'] = attached
    __args__['namePrefix'] = name_prefix
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getPublicIPs:getPublicIPs', __args__, opts=opts).value

    return AwaitableGetPublicIPsResult(
        allocation_type=__ret__.get('allocationType'),
        attached=__ret__.get('attached'),
        id=__ret__.get('id'),
        name_prefix=__ret__.get('namePrefix'),
        public_ips=__ret__.get('publicIps'),
        resource_group_name=__ret__.get('resourceGroupName'))
