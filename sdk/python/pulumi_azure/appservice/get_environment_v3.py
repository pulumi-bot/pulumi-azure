# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetEnvironmentV3Result',
    'AwaitableGetEnvironmentV3Result',
    'get_environment_v3',
]

@pulumi.output_type
class GetEnvironmentV3Result:
    """
    A collection of values returned by getEnvironmentV3.
    """
    def __init__(__self__, cluster_settings=None, id=None, name=None, resource_group_name=None, subnet_id=None, tags=None):
        if cluster_settings and not isinstance(cluster_settings, list):
            raise TypeError("Expected argument 'cluster_settings' to be a list")
        pulumi.set(__self__, "cluster_settings", cluster_settings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterSettings")
    def cluster_settings(self) -> Sequence['outputs.GetEnvironmentV3ClusterSettingResult']:
        """
        A `cluster_setting` block as defined below.
        """
        return pulumi.get(self, "cluster_settings")

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
        """
        The name of the Cluster Setting.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of the v3 App Service Environment Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the v3 App Service Environment.
        """
        return pulumi.get(self, "tags")


class AwaitableGetEnvironmentV3Result(GetEnvironmentV3Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentV3Result(
            cluster_settings=self.cluster_settings,
            id=self.id,
            name=self.name,
            resource_group_name=self.resource_group_name,
            subnet_id=self.subnet_id,
            tags=self.tags)


def get_environment_v3(name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentV3Result:
    """
    Use this data source to access information about an existing 3rd Generation (v3) App Service Environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.appservice.get_environment_v3(name="example-ASE",
        resource_group_name="example-resource-group")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of this v3 App Service Environment.
    :param str resource_group_name: The name of the Resource Group where the v3 App Service Environment exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:appservice/getEnvironmentV3:getEnvironmentV3', __args__, opts=opts, typ=GetEnvironmentV3Result).value

    return AwaitableGetEnvironmentV3Result(
        cluster_settings=__ret__.cluster_settings,
        id=__ret__.id,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        subnet_id=__ret__.subnet_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_environment_v3)
def get_environment_v3_apply(name: Optional[pulumi.Input[str]] = None,
                             resource_group_name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentV3Result]:
    ...
