# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetIntegrationAccountResult',
    'AwaitableGetIntegrationAccountResult',
    'get_integration_account',
]


@pulumi.output_type
class _GetIntegrationAccountResult(dict):
    id: str = pulumi.property("id")
    location: str = pulumi.property("location")
    name: str = pulumi.property("name")
    resource_group_name: str = pulumi.property("resourceGroupName")
    sku_name: str = pulumi.property("skuName")
    tags: Mapping[str, str] = pulumi.property("tags")


class GetIntegrationAccountResult:
    """
    A collection of values returned by getIntegrationAccount.
    """
    def __init__(__self__, id=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None):
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
        The Azure Region where the Logic App Integration Account exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        __self__.sku_name = sku_name
        """
        The sku name of the Logic App Integration Account.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the Logic App Integration Account.
        """


class AwaitableGetIntegrationAccountResult(GetIntegrationAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationAccountResult(
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            sku_name=self.sku_name,
            tags=self.tags)


def get_integration_account(name: Optional[str] = None,
                            resource_group_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationAccountResult:
    """
    Use this data source to access information about an existing Logic App Integration Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.logicapps.get_integration_account(name="example-account",
        resource_group_name="example-resource-group")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of this Logic App Integration Account.
    :param str resource_group_name: The name of the Resource Group where the Logic App Integration Account exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:logicapps/getIntegrationAccount:getIntegrationAccount', __args__, opts=opts, typ=_GetIntegrationAccountResult).value

    return AwaitableGetIntegrationAccountResult(
        id=_utilities.get_dict_value(__ret__, 'id'),
        location=_utilities.get_dict_value(__ret__, 'location'),
        name=_utilities.get_dict_value(__ret__, 'name'),
        resource_group_name=_utilities.get_dict_value(__ret__, 'resourceGroupName'),
        sku_name=_utilities.get_dict_value(__ret__, 'skuName'),
        tags=_utilities.get_dict_value(__ret__, 'tags'))
