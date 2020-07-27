# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetPlatformImageResult',
    'AwaitableGetPlatformImageResult',
    'get_platform_image',
]


class GetPlatformImageResult:
    """
    A collection of values returned by getPlatformImage.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, id=None, location=None, offer=None, publisher=None, sku=None, version=None) -> None:
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if offer and not isinstance(offer, str):
            raise TypeError("Expected argument 'offer' to be a str")
        __self__.offer = offer
        if publisher and not isinstance(publisher, str):
            raise TypeError("Expected argument 'publisher' to be a str")
        __self__.publisher = publisher
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        __self__.sku = sku
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        __self__.version = version


class AwaitableGetPlatformImageResult(GetPlatformImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlatformImageResult(
            id=self.id,
            location=self.location,
            offer=self.offer,
            publisher=self.publisher,
            sku=self.sku,
            version=self.version)


def get_platform_image(location: Optional[str] = None, offer: Optional[str] = None, publisher: Optional[str] = None, sku: Optional[str] = None, version: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlatformImageResult:
    """
    Use this data source to access information about a Platform Image.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.compute.get_platform_image(location="West Europe",
        publisher="Canonical",
        offer="UbuntuServer",
        sku="16.04-LTS")
    pulumi.export("id", example.id)
    ```


    :param str location: Specifies the Location to pull information about this Platform Image from.
    :param str offer: Specifies the Offer associated with the Platform Image.
    :param str publisher: Specifies the Publisher associated with the Platform Image.
    :param str sku: Specifies the SKU of the Platform Image.
    :param str version: The version of the Platform Image.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['offer'] = offer
    __args__['publisher'] = publisher
    __args__['sku'] = sku
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getPlatformImage:getPlatformImage', __args__, opts=opts).value

    return AwaitableGetPlatformImageResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        offer=__ret__.get('offer'),
        publisher=__ret__.get('publisher'),
        sku=__ret__.get('sku'),
        version=__ret__.get('version'))
