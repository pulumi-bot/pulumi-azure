# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetProviderResult',
    'AwaitableGetProviderResult',
    'get_provider',
]

@pulumi.output_type
class GetProviderResult:
    """
    A collection of values returned by getProvider.
    """
    def __init__(__self__, attestation_uri=None, id=None, location=None, name=None, resource_group_name=None, tags=None, trust_model=None):
        if attestation_uri and not isinstance(attestation_uri, str):
            raise TypeError("Expected argument 'attestation_uri' to be a str")
        pulumi.set(__self__, "attestation_uri", attestation_uri)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if trust_model and not isinstance(trust_model, str):
            raise TypeError("Expected argument 'trust_model' to be a str")
        pulumi.set(__self__, "trust_model", trust_model)

    @property
    @pulumi.getter(name="attestationUri")
    def attestation_uri(self) -> str:
        return pulumi.get(self, "attestation_uri")

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
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trustModel")
    def trust_model(self) -> str:
        return pulumi.get(self, "trust_model")


class AwaitableGetProviderResult(GetProviderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProviderResult(
            attestation_uri=self.attestation_uri,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            trust_model=self.trust_model)


def get_provider(name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProviderResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:attestation/getProvider:getProvider', __args__, opts=opts, typ=GetProviderResult).value

    return AwaitableGetProviderResult(
        attestation_uri=__ret__.attestation_uri,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags,
        trust_model=__ret__.trust_model)
