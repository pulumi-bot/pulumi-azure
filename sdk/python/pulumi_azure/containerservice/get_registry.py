# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetRegistryResult',
    'AwaitableGetRegistryResult',
    'get_registry',
]

@pulumi.output_type
class GetRegistryResult:
    """
    A collection of values returned by getRegistry.
    """
    def __init__(__self__, admin_enabled=None, admin_password=None, admin_username=None, id=None, location=None, login_server=None, name=None, resource_group_name=None, sku=None, storage_account_id=None, tags=None):
        if admin_enabled and not isinstance(admin_enabled, bool):
            raise TypeError("Expected argument 'admin_enabled' to be a bool")
        pulumi.set(__self__, "admin_enabled", admin_enabled)
        if admin_password and not isinstance(admin_password, str):
            raise TypeError("Expected argument 'admin_password' to be a str")
        pulumi.set(__self__, "admin_password", admin_password)
        if admin_username and not isinstance(admin_username, str):
            raise TypeError("Expected argument 'admin_username' to be a str")
        pulumi.set(__self__, "admin_username", admin_username)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if login_server and not isinstance(login_server, str):
            raise TypeError("Expected argument 'login_server' to be a str")
        pulumi.set(__self__, "login_server", login_server)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        pulumi.set(__self__, "sku", sku)
        if storage_account_id and not isinstance(storage_account_id, str):
            raise TypeError("Expected argument 'storage_account_id' to be a str")
        pulumi.set(__self__, "storage_account_id", storage_account_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="adminEnabled")
    def admin_enabled(self) -> bool:
        """
        Is the Administrator account enabled for this Container Registry.
        """
        return pulumi.get(self, "admin_enabled")

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> str:
        """
        The Password associated with the Container Registry Admin account - if the admin account is enabled.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The Username associated with the Container Registry Admin account - if the admin account is enabled.
        """
        return pulumi.get(self, "admin_username")

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
        The Azure Region in which this Container Registry exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="loginServer")
    def login_server(self) -> str:
        """
        The URL that can be used to log into the container registry.
        """
        return pulumi.get(self, "login_server")

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
    def sku(self) -> str:
        """
        The SKU of this Container Registry, such as `Basic`.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> str:
        """
        The ID of the Storage Account used for this Container Registry. This is only returned for `Classic` SKU's.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the Container Registry.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRegistryResult(GetRegistryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryResult(
            admin_enabled=self.admin_enabled,
            admin_password=self.admin_password,
            admin_username=self.admin_username,
            id=self.id,
            location=self.location,
            login_server=self.login_server,
            name=self.name,
            resource_group_name=self.resource_group_name,
            sku=self.sku,
            storage_account_id=self.storage_account_id,
            tags=self.tags)


def get_registry(name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryResult:
    """
    Use this data source to access information about an existing Container Registry.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.containerservice.get_registry(name="testacr",
        resource_group_name="test")
    pulumi.export("loginServer", example.login_server)
    ```


    :param str name: The name of the Container Registry.
    :param str resource_group_name: The Name of the Resource Group where this Container Registry exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:containerservice/getRegistry:getRegistry', __args__, opts=opts, typ=GetRegistryResult).value

    return AwaitableGetRegistryResult(
        admin_enabled=__ret__.admin_enabled,
        admin_password=__ret__.admin_password,
        admin_username=__ret__.admin_username,
        id=__ret__.id,
        location=__ret__.location,
        login_server=__ret__.login_server,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        sku=__ret__.sku,
        storage_account_id=__ret__.storage_account_id,
        tags=__ret__.tags)
