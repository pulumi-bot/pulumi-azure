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
    'GetMariaDbServerResult',
    'AwaitableGetMariaDbServerResult',
    'get_maria_db_server',
]

@pulumi.output_type
class GetMariaDbServerResult:
    """
    A collection of values returned by getMariaDbServer.
    """
    def __init__(__self__, administrator_login=None, fqdn=None, id=None, location=None, name=None, resource_group_name=None, sku_name=None, ssl_enforcement=None, storage_profiles=None, tags=None, version=None):
        if administrator_login and not isinstance(administrator_login, str):
            raise TypeError("Expected argument 'administrator_login' to be a str")
        pulumi.set(__self__, "administrator_login", administrator_login)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
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
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        pulumi.set(__self__, "sku_name", sku_name)
        if ssl_enforcement and not isinstance(ssl_enforcement, str):
            raise TypeError("Expected argument 'ssl_enforcement' to be a str")
        pulumi.set(__self__, "ssl_enforcement", ssl_enforcement)
        if storage_profiles and not isinstance(storage_profiles, list):
            raise TypeError("Expected argument 'storage_profiles' to be a list")
        pulumi.set(__self__, "storage_profiles", storage_profiles)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> str:
        """
        The Administrator Login for the MariaDB Server.
        """
        return pulumi.get(self, "administrator_login")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        The FQDN of the MariaDB Server.
        """
        return pulumi.get(self, "fqdn")

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
        The Azure location where the resource exists.
        """
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
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        """
        The SKU Name for this MariaDB Server.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="sslEnforcement")
    def ssl_enforcement(self) -> str:
        """
        The SSL being enforced on connections.
        """
        return pulumi.get(self, "ssl_enforcement")

    @property
    @pulumi.getter(name="storageProfiles")
    def storage_profiles(self) -> Sequence['outputs.GetMariaDbServerStorageProfileResult']:
        """
        A `storage_profile` block as defined below.
        """
        return pulumi.get(self, "storage_profiles")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the resource.
        ---
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of MariaDB being used.
        """
        return pulumi.get(self, "version")


class AwaitableGetMariaDbServerResult(GetMariaDbServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMariaDbServerResult(
            administrator_login=self.administrator_login,
            fqdn=self.fqdn,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            sku_name=self.sku_name,
            ssl_enforcement=self.ssl_enforcement,
            storage_profiles=self.storage_profiles,
            tags=self.tags,
            version=self.version)


def get_maria_db_server(name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMariaDbServerResult:
    """
    Use this data source to access information about an existing MariaDB Server.


    :param str name: The name of the MariaDB Server to retrieve information about.
    :param str resource_group_name: The name of the resource group where the MariaDB Server exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:mariadb/getMariaDbServer:getMariaDbServer', __args__, opts=opts, typ=GetMariaDbServerResult).value

    return AwaitableGetMariaDbServerResult(
        administrator_login=__ret__.administrator_login,
        fqdn=__ret__.fqdn,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        sku_name=__ret__.sku_name,
        ssl_enforcement=__ret__.ssl_enforcement,
        storage_profiles=__ret__.storage_profiles,
        tags=__ret__.tags,
        version=__ret__.version)
