# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetTrafficManagerResult',
    'AwaitableGetTrafficManagerResult',
    'get_traffic_manager',
]


class GetTrafficManagerResult:
    """
    A collection of values returned by getTrafficManager.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, id=None, name=None) -> None:
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name


class AwaitableGetTrafficManagerResult(GetTrafficManagerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrafficManagerResult(
            id=self.id,
            name=self.name)


def get_traffic_manager(name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrafficManagerResult:
    """
    Use this data source to access the ID of a specified Traffic Manager Geographical Location within the Geographical Hierarchy.

    ## Example Usage
    ### World)

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_traffic_manager(name="World")
    pulumi.export("locationCode", example.id)
    ```


    :param str name: Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getTrafficManager:getTrafficManager', __args__, opts=opts).value

    return AwaitableGetTrafficManagerResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'))
