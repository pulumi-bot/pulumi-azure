# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDomainTopicResult',
    'AwaitableGetDomainTopicResult',
    'get_domain_topic',
]

@pulumi.output_type
class GetDomainTopicResult:
    """
    A collection of values returned by getDomainTopic.
    """
    def __init__(__self__, domain_name=None, id=None, name=None, resource_group_name=None):
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        """
        The EventGrid Domain Topic Domain name.
        """
        return pulumi.get(self, "domain_name")

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
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")


class AwaitableGetDomainTopicResult(GetDomainTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainTopicResult(
            domain_name=self.domain_name,
            id=self.id,
            name=self.name,
            resource_group_name=self.resource_group_name)


def get_domain_topic(domain_name: Optional[str] = None,
                     name: Optional[str] = None,
                     resource_group_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainTopicResult:
    """
    Use this data source to access information about an existing EventGrid Domain Topic


    :param str domain_name: The name of the EventGrid Domain Topic domain.
    :param str name: The name of the EventGrid Domain Topic resource.
    :param str resource_group_name: The name of the resource group in which the EventGrid Domain Topic exists.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:eventgrid/getDomainTopic:getDomainTopic', __args__, opts=opts, typ=GetDomainTopicResult).value

    return AwaitableGetDomainTopicResult(
        domain_name=__ret__.domain_name,
        id=__ret__.id,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name)


@_utilities.lift_output_func(get_domain_topic)
def get_domain_topic_apply(domain_name: Optional[pulumi.Input[str]] = None,
                           name: Optional[pulumi.Input[str]] = None,
                           resource_group_name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainTopicResult]:
    ...
