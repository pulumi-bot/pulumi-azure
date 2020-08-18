# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetAccountSASResult:
    """
    A collection of values returned by getAccountSAS.
    """
    def __init__(__self__, connection_string=None, expiry=None, https_only=None, id=None, permissions=None, resource_types=None, sas=None, services=None, signed_version=None, start=None):
        if connection_string and not isinstance(connection_string, str):
            raise TypeError("Expected argument 'connection_string' to be a str")
        __self__.connection_string = connection_string
        if expiry and not isinstance(expiry, str):
            raise TypeError("Expected argument 'expiry' to be a str")
        __self__.expiry = expiry
        if https_only and not isinstance(https_only, bool):
            raise TypeError("Expected argument 'https_only' to be a bool")
        __self__.https_only = https_only
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if permissions and not isinstance(permissions, dict):
            raise TypeError("Expected argument 'permissions' to be a dict")
        __self__.permissions = permissions
        if resource_types and not isinstance(resource_types, dict):
            raise TypeError("Expected argument 'resource_types' to be a dict")
        __self__.resource_types = resource_types
        if sas and not isinstance(sas, str):
            raise TypeError("Expected argument 'sas' to be a str")
        __self__.sas = sas
        """
        The computed Account Shared Access Signature (SAS).
        """
        if services and not isinstance(services, dict):
            raise TypeError("Expected argument 'services' to be a dict")
        __self__.services = services
        if signed_version and not isinstance(signed_version, str):
            raise TypeError("Expected argument 'signed_version' to be a str")
        __self__.signed_version = signed_version
        if start and not isinstance(start, str):
            raise TypeError("Expected argument 'start' to be a str")
        __self__.start = start


class AwaitableGetAccountSASResult(GetAccountSASResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountSASResult(
            connection_string=self.connection_string,
            expiry=self.expiry,
            https_only=self.https_only,
            id=self.id,
            permissions=self.permissions,
            resource_types=self.resource_types,
            sas=self.sas,
            services=self.services,
            signed_version=self.signed_version,
            start=self.start)


def get_account_sas(connection_string=None, expiry=None, https_only=None, permissions=None, resource_types=None, services=None, signed_version=None, start=None, opts=None):
    """
    Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account.

    Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.

    Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
    and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westus")
    example_account = azure.storage.Account("exampleAccount",
        resource_group_name=example_resource_group.name,
        location="westus",
        account_tier="Standard",
        account_replication_type="GRS",
        tags={
            "environment": "staging",
        })
    example_account_sas = example_account.primary_connection_string.apply(lambda primary_connection_string: azure.storage.get_account_sas(connection_string=primary_connection_string,
        https_only=True,
        signed_version="2017-07-29",
        resource_types={
            "service": True,
            "container": False,
            "object": False,
        },
        services={
            "blob": True,
            "queue": False,
            "table": False,
            "file": False,
        },
        start="2018-03-21",
        expiry="2020-03-21",
        permissions={
            "read": True,
            "write": True,
            "delete": False,
            "list": False,
            "add": True,
            "create": True,
            "update": False,
            "process": False,
        }))
    pulumi.export("sasUrlQueryString", example_account_sas.sas)
    ```


    :param str connection_string: The connection string for the storage account to which this SAS applies. Typically directly from the `primary_connection_string` attribute of a `storage.Account` resource.
    :param str expiry: The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
    :param bool https_only: Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
    :param dict permissions: A `permissions` block as defined below.
    :param dict resource_types: A `resource_types` block as defined below.
    :param dict services: A `services` block as defined below.
    :param str signed_version: Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to `2017-07-29`.
    :param str start: The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.

    The **permissions** object supports the following:

      * `add` (`bool`) - Should Add permissions be enabled for this SAS?
      * `create` (`bool`) - Should Create permissions be enabled for this SAS?
      * `delete` (`bool`) - Should Delete permissions be enabled for this SAS?
      * `list` (`bool`) - Should List permissions be enabled for this SAS?
      * `process` (`bool`) - Should Process permissions be enabled for this SAS?
      * `read` (`bool`) - Should Read permissions be enabled for this SAS?
      * `update` (`bool`) - Should Update permissions be enabled for this SAS?
      * `write` (`bool`) - Should Write permissions be enabled for this SAS?

    The **resource_types** object supports the following:

      * `container` (`bool`) - Should permission be granted to the container?
      * `object` (`bool`) - Should permission be granted only to a specific object?
      * `service` (`bool`) - Should permission be granted to the entire service?

    The **services** object supports the following:

      * `blob` (`bool`) - Should permission be granted to `blob` services within this storage account?
      * `file` (`bool`) - Should permission be granted to `file` services within this storage account?
      * `queue` (`bool`) - Should permission be granted to `queue` services within this storage account?
      * `table` (`bool`) - Should permission be granted to `table` services within this storage account?
    """
    __args__ = dict()
    __args__['connectionString'] = connection_string
    __args__['expiry'] = expiry
    __args__['httpsOnly'] = https_only
    __args__['permissions'] = permissions
    __args__['resourceTypes'] = resource_types
    __args__['services'] = services
    __args__['signedVersion'] = signed_version
    __args__['start'] = start
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:storage/getAccountSAS:getAccountSAS', __args__, opts=opts).value

    return AwaitableGetAccountSASResult(
        connection_string=__ret__.get('connectionString'),
        expiry=__ret__.get('expiry'),
        https_only=__ret__.get('httpsOnly'),
        id=__ret__.get('id'),
        permissions=__ret__.get('permissions'),
        resource_types=__ret__.get('resourceTypes'),
        sas=__ret__.get('sas'),
        services=__ret__.get('services'),
        signed_version=__ret__.get('signedVersion'),
        start=__ret__.get('start'))
