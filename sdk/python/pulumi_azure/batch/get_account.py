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
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
]

@pulumi.output_type
class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, account_endpoint=None, id=None, key_vault_references=None, location=None, name=None, pool_allocation_mode=None, primary_access_key=None, resource_group_name=None, secondary_access_key=None, storage_account_id=None, tags=None):
        if account_endpoint and not isinstance(account_endpoint, str):
            raise TypeError("Expected argument 'account_endpoint' to be a str")
        pulumi.set(__self__, "account_endpoint", account_endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_vault_references and not isinstance(key_vault_references, list):
            raise TypeError("Expected argument 'key_vault_references' to be a list")
        pulumi.set(__self__, "key_vault_references", key_vault_references)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pool_allocation_mode and not isinstance(pool_allocation_mode, str):
            raise TypeError("Expected argument 'pool_allocation_mode' to be a str")
        pulumi.set(__self__, "pool_allocation_mode", pool_allocation_mode)
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        pulumi.set(__self__, "primary_access_key", primary_access_key)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        pulumi.set(__self__, "secondary_access_key", secondary_access_key)
        if storage_account_id and not isinstance(storage_account_id, str):
            raise TypeError("Expected argument 'storage_account_id' to be a str")
        pulumi.set(__self__, "storage_account_id", storage_account_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountEndpoint")
    def account_endpoint(self) -> str:
        """
        The account endpoint used to interact with the Batch service.
        """
        return pulumi.get(self, "account_endpoint")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyVaultReferences")
    def key_vault_references(self) -> Sequence['outputs.GetAccountKeyVaultReferenceResult']:
        """
        The `key_vault_reference` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the `UserSubscription` pool allocation mode.
        """
        return pulumi.get(self, "key_vault_references")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure Region in which this Batch account exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The Batch account name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolAllocationMode")
    def pool_allocation_mode(self) -> str:
        """
        The pool allocation mode configured for this Batch account.
        """
        return pulumi.get(self, "pool_allocation_mode")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> str:
        """
        The Batch account primary access key.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> str:
        """
        The Batch account secondary access key.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> str:
        """
        The ID of the Storage Account used for this Batch account.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the Batch account.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            account_endpoint=self.account_endpoint,
            id=self.id,
            key_vault_references=self.key_vault_references,
            location=self.location,
            name=self.name,
            pool_allocation_mode=self.pool_allocation_mode,
            primary_access_key=self.primary_access_key,
            resource_group_name=self.resource_group_name,
            secondary_access_key=self.secondary_access_key,
            storage_account_id=self.storage_account_id,
            tags=self.tags)


def get_account(name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Use this data source to access information about an existing Batch Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.batch.get_account(name="testbatchaccount",
        resource_group_name="test")
    pulumi.export("poolAllocationMode", example.pool_allocation_mode)
    ```


    :param str name: The name of the Batch account.
    :param str resource_group_name: The Name of the Resource Group where this Batch account exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:batch/getAccount:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        account_endpoint=__ret__.account_endpoint,
        id=__ret__.id,
        key_vault_references=__ret__.key_vault_references,
        location=__ret__.location,
        name=__ret__.name,
        pool_allocation_mode=__ret__.pool_allocation_mode,
        primary_access_key=__ret__.primary_access_key,
        resource_group_name=__ret__.resource_group_name,
        secondary_access_key=__ret__.secondary_access_key,
        storage_account_id=__ret__.storage_account_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_account)
def get_account_apply(name: Optional[pulumi.Input[str]] = None,
                      resource_group_name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    ...
