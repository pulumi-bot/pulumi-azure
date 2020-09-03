# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetSubscriptionsResult',
    'AwaitableGetSubscriptionsResult',
    'get_subscriptions',
]

@pulumi.output_type
class GetSubscriptionsResult:
    """
    A collection of values returned by getSubscriptions.
    """
    def __init__(__self__, display_name_contains=None, display_name_prefix=None, id=None, subscriptions=None):
        if display_name_contains and not isinstance(display_name_contains, str):
            raise TypeError("Expected argument 'display_name_contains' to be a str")
        pulumi.set(__self__, "display_name_contains", display_name_contains)
        if display_name_prefix and not isinstance(display_name_prefix, str):
            raise TypeError("Expected argument 'display_name_prefix' to be a str")
        pulumi.set(__self__, "display_name_prefix", display_name_prefix)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if subscriptions and not isinstance(subscriptions, list):
            raise TypeError("Expected argument 'subscriptions' to be a list")
        pulumi.set(__self__, "subscriptions", subscriptions)

    @property
    @pulumi.getter(name="displayNameContains")
    def display_name_contains(self) -> Optional[str]:
        return pulumi.get(self, "display_name_contains")

    @property
    @pulumi.getter(name="displayNamePrefix")
    def display_name_prefix(self) -> Optional[str]:
        return pulumi.get(self, "display_name_prefix")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def subscriptions(self) -> Sequence['outputs.GetSubscriptionsSubscriptionResult']:
        """
        One or more `subscription` blocks as defined below.
        """
        return pulumi.get(self, "subscriptions")


class AwaitableGetSubscriptionsResult(GetSubscriptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubscriptionsResult(
            display_name_contains=self.display_name_contains,
            display_name_prefix=self.display_name_prefix,
            id=self.id,
            subscriptions=self.subscriptions)


def get_subscriptions(display_name_contains: Optional[str] = None,
                      display_name_prefix: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubscriptionsResult:
    """
    Use this data source to access information about all the Subscriptions currently available.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    available = azure.core.get_subscriptions()
    pulumi.export("availableSubscriptions", available.subscriptions)
    pulumi.export("firstAvailableSubscriptionDisplayName", available.subscriptions[0].display_name)
    ```


    :param str display_name_contains: A case-insensitive value which must be contained within the `display_name` field, used to filter the results
    :param str display_name_prefix: A case-insensitive prefix which can be used to filter on the `display_name` field
    """
    __args__ = dict()
    __args__['displayNameContains'] = display_name_contains
    __args__['displayNamePrefix'] = display_name_prefix
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:core/getSubscriptions:getSubscriptions', __args__, opts=opts, typ=GetSubscriptionsResult).value

    return AwaitableGetSubscriptionsResult(
        display_name_contains=__ret__.display_name_contains,
        display_name_prefix=__ret__.display_name_prefix,
        id=__ret__.id,
        subscriptions=__ret__.subscriptions)
