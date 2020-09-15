# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetAccountBlobContainerSASResult',
    'AwaitableGetAccountBlobContainerSASResult',
    'get_account_blob_container_sas',
]

@pulumi.output_type
class GetAccountBlobContainerSASResult:
    """
    A collection of values returned by getAccountBlobContainerSAS.
    """
    def __init__(__self__, cache_control=None, connection_string=None, container_name=None, content_disposition=None, content_encoding=None, content_language=None, content_type=None, expiry=None, https_only=None, id=None, ip_address=None, permissions=None, sas=None, start=None):
        if cache_control and not isinstance(cache_control, str):
            raise TypeError("Expected argument 'cache_control' to be a str")
        pulumi.set(__self__, "cache_control", cache_control)
        if connection_string and not isinstance(connection_string, str):
            raise TypeError("Expected argument 'connection_string' to be a str")
        pulumi.set(__self__, "connection_string", connection_string)
        if container_name and not isinstance(container_name, str):
            raise TypeError("Expected argument 'container_name' to be a str")
        pulumi.set(__self__, "container_name", container_name)
        if content_disposition and not isinstance(content_disposition, str):
            raise TypeError("Expected argument 'content_disposition' to be a str")
        pulumi.set(__self__, "content_disposition", content_disposition)
        if content_encoding and not isinstance(content_encoding, str):
            raise TypeError("Expected argument 'content_encoding' to be a str")
        pulumi.set(__self__, "content_encoding", content_encoding)
        if content_language and not isinstance(content_language, str):
            raise TypeError("Expected argument 'content_language' to be a str")
        pulumi.set(__self__, "content_language", content_language)
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        pulumi.set(__self__, "content_type", content_type)
        if expiry and not isinstance(expiry, str):
            raise TypeError("Expected argument 'expiry' to be a str")
        pulumi.set(__self__, "expiry", expiry)
        if https_only and not isinstance(https_only, bool):
            raise TypeError("Expected argument 'https_only' to be a bool")
        pulumi.set(__self__, "https_only", https_only)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if permissions and not isinstance(permissions, dict):
            raise TypeError("Expected argument 'permissions' to be a dict")
        pulumi.set(__self__, "permissions", permissions)
        if sas and not isinstance(sas, str):
            raise TypeError("Expected argument 'sas' to be a str")
        pulumi.set(__self__, "sas", sas)
        if start and not isinstance(start, str):
            raise TypeError("Expected argument 'start' to be a str")
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter(name="cacheControl")
    def cache_control(self) -> Optional[str]:
        return pulumi.get(self, "cache_control")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> str:
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter(name="contentDisposition")
    def content_disposition(self) -> Optional[str]:
        return pulumi.get(self, "content_disposition")

    @property
    @pulumi.getter(name="contentEncoding")
    def content_encoding(self) -> Optional[str]:
        return pulumi.get(self, "content_encoding")

    @property
    @pulumi.getter(name="contentLanguage")
    def content_language(self) -> Optional[str]:
        return pulumi.get(self, "content_language")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def expiry(self) -> str:
        return pulumi.get(self, "expiry")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[bool]:
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def permissions(self) -> 'outputs.GetAccountBlobContainerSASPermissionsResult':
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def sas(self) -> str:
        """
        The computed Blob Container Shared Access Signature (SAS).
        """
        return pulumi.get(self, "sas")

    @property
    @pulumi.getter
    def start(self) -> str:
        return pulumi.get(self, "start")


class AwaitableGetAccountBlobContainerSASResult(GetAccountBlobContainerSASResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountBlobContainerSASResult(
            cache_control=self.cache_control,
            connection_string=self.connection_string,
            container_name=self.container_name,
            content_disposition=self.content_disposition,
            content_encoding=self.content_encoding,
            content_language=self.content_language,
            content_type=self.content_type,
            expiry=self.expiry,
            https_only=self.https_only,
            id=self.id,
            ip_address=self.ip_address,
            permissions=self.permissions,
            sas=self.sas,
            start=self.start)


def get_account_blob_container_sas(cache_control: Optional[str] = None,
                                   connection_string: Optional[str] = None,
                                   container_name: Optional[str] = None,
                                   content_disposition: Optional[str] = None,
                                   content_encoding: Optional[str] = None,
                                   content_language: Optional[str] = None,
                                   content_type: Optional[str] = None,
                                   expiry: Optional[str] = None,
                                   https_only: Optional[bool] = None,
                                   ip_address: Optional[str] = None,
                                   permissions: Optional[pulumi.InputType['GetAccountBlobContainerSASPermissionsArgs']] = None,
                                   start: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountBlobContainerSASResult:
    """
    Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.

    Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account Blob Container.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    rg = azure.core.ResourceGroup("rg", location="westus")
    storage = azure.storage.Account("storage",
        resource_group_name=rg.name,
        location=rg.location,
        account_tier="Standard",
        account_replication_type="LRS")
    container = azure.storage.Container("container",
        storage_account_name=storage.name,
        container_access_type="private")
    example = pulumi.Output.all(storage.primary_connection_string, container.name).apply(lambda primary_connection_string, name: azure.storage.get_account_blob_container_sas(connection_string=primary_connection_string,
        container_name=name,
        https_only=True,
        ip_address="168.1.5.65",
        start="2018-03-21",
        expiry="2018-03-21",
        permissions=azure.storage.GetAccountBlobContainerSASPermissionsArgs(
            read=True,
            add=True,
            create=False,
            write=False,
            delete=True,
            list=True,
        ),
        cache_control="max-age=5",
        content_disposition="inline",
        content_encoding="deflate",
        content_language="en-US",
        content_type="application/json"))
    pulumi.export("sasUrlQueryString", example.sas)
    ```


    :param str cache_control: The `Cache-Control` response header that is sent when this SAS token is used.
    :param str connection_string: The connection string for the storage account to which this SAS applies. Typically directly from the `primary_connection_string` attribute of an `storage.Account` resource.
    :param str container_name: Name of the container.
    :param str content_disposition: The `Content-Disposition` response header that is sent when this SAS token is used.
    :param str content_encoding: The `Content-Encoding` response header that is sent when this SAS token is used.
    :param str content_language: The `Content-Language` response header that is sent when this SAS token is used.
    :param str content_type: The `Content-Type` response header that is sent when this SAS token is used.
    :param str expiry: The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
    :param bool https_only: Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
    :param str ip_address: Single ipv4 address or range (connected with a dash) of ipv4 addresses.
    :param pulumi.InputType['GetAccountBlobContainerSASPermissionsArgs'] permissions: A `permissions` block as defined below.
    :param str start: The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
    """
    __args__ = dict()
    __args__['cacheControl'] = cache_control
    __args__['connectionString'] = connection_string
    __args__['containerName'] = container_name
    __args__['contentDisposition'] = content_disposition
    __args__['contentEncoding'] = content_encoding
    __args__['contentLanguage'] = content_language
    __args__['contentType'] = content_type
    __args__['expiry'] = expiry
    __args__['httpsOnly'] = https_only
    __args__['ipAddress'] = ip_address
    __args__['permissions'] = permissions
    __args__['start'] = start
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:storage/getAccountBlobContainerSAS:getAccountBlobContainerSAS', __args__, opts=opts, typ=GetAccountBlobContainerSASResult).value

    return AwaitableGetAccountBlobContainerSASResult(
        cache_control=__ret__.cache_control,
        connection_string=__ret__.connection_string,
        container_name=__ret__.container_name,
        content_disposition=__ret__.content_disposition,
        content_encoding=__ret__.content_encoding,
        content_language=__ret__.content_language,
        content_type=__ret__.content_type,
        expiry=__ret__.expiry,
        https_only=__ret__.https_only,
        id=__ret__.id,
        ip_address=__ret__.ip_address,
        permissions=__ret__.permissions,
        sas=__ret__.sas,
        start=__ret__.start)
