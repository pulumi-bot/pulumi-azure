# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetAppServiceEnvironmentResult',
    'AwaitableGetAppServiceEnvironmentResult',
    'get_app_service_environment',
]

@pulumi.output_type
class GetAppServiceEnvironmentResult:
    """
    A collection of values returned by getAppServiceEnvironment.
    """
    def __init__(__self__, front_end_scale_factor=None, id=None, location=None, name=None, pricing_tier=None, resource_group_name=None, tags=None):
        if front_end_scale_factor and not isinstance(front_end_scale_factor, float):
            raise TypeError("Expected argument 'front_end_scale_factor' to be a float")
        pulumi.set(__self__, "front_end_scale_factor", front_end_scale_factor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pricing_tier and not isinstance(pricing_tier, str):
            raise TypeError("Expected argument 'pricing_tier' to be a str")
        pulumi.set(__self__, "pricing_tier", pricing_tier)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="frontEndScaleFactor")
    def front_end_scale_factor(self) -> float:
        """
        The number of app instances per App Service Environment Front End
        """
        return pulumi.get(self, "front_end_scale_factor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure location where the App Service Environment exists
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pricingTier")
    def pricing_tier(self) -> str:
        """
        The Pricing Tier (Isolated SKU) of the App Service Environment.
        """
        return pulumi.get(self, "pricing_tier")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAppServiceEnvironmentResult(GetAppServiceEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppServiceEnvironmentResult(
            front_end_scale_factor=self.front_end_scale_factor,
            id=self.id,
            location=self.location,
            name=self.name,
            pricing_tier=self.pricing_tier,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_app_service_environment(name: Optional[str] = None,
                                resource_group_name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppServiceEnvironmentResult:
    """
    Use this data source to access information about an existing App Service Environment

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.appservice.get_app_service_environment(name="example-ase",
        resource_group_name="example-rg")
    pulumi.export("appServiceEnvironmentId", data["azurerm_app_service_environment"]["id"])
    ```


    :param str name: The name of the App Service Environment.
    :param str resource_group_name: The Name of the Resource Group where the App Service Environment exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment', __args__, opts=opts, typ=GetAppServiceEnvironmentResult).value

    return AwaitableGetAppServiceEnvironmentResult(
        front_end_scale_factor=__ret__.front_end_scale_factor,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        pricing_tier=__ret__.pricing_tier,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags)
