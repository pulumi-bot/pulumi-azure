# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
]


class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, access_tier=None, account_kind=None, account_replication_type=None, account_tier=None, allow_blob_public_access=None, custom_domains=None, enable_https_traffic_only=None, id=None, is_hns_enabled=None, location=None, min_tls_version=None, name=None, primary_access_key=None, primary_blob_connection_string=None, primary_blob_endpoint=None, primary_blob_host=None, primary_connection_string=None, primary_dfs_endpoint=None, primary_dfs_host=None, primary_file_endpoint=None, primary_file_host=None, primary_location=None, primary_queue_endpoint=None, primary_queue_host=None, primary_table_endpoint=None, primary_table_host=None, primary_web_endpoint=None, primary_web_host=None, resource_group_name=None, secondary_access_key=None, secondary_blob_connection_string=None, secondary_blob_endpoint=None, secondary_blob_host=None, secondary_connection_string=None, secondary_dfs_endpoint=None, secondary_dfs_host=None, secondary_file_endpoint=None, secondary_file_host=None, secondary_location=None, secondary_queue_endpoint=None, secondary_queue_host=None, secondary_table_endpoint=None, secondary_table_host=None, secondary_web_endpoint=None, secondary_web_host=None, tags=None):
        if access_tier and not isinstance(access_tier, str):
            raise TypeError("Expected argument 'access_tier' to be a str")
        __self__.access_tier = access_tier
        """
        The access tier for `BlobStorage` accounts.
        """
        if account_kind and not isinstance(account_kind, str):
            raise TypeError("Expected argument 'account_kind' to be a str")
        __self__.account_kind = account_kind
        """
        The Kind of account.
        """
        if account_replication_type and not isinstance(account_replication_type, str):
            raise TypeError("Expected argument 'account_replication_type' to be a str")
        __self__.account_replication_type = account_replication_type
        """
        The type of replication used for this storage account.
        """
        if account_tier and not isinstance(account_tier, str):
            raise TypeError("Expected argument 'account_tier' to be a str")
        __self__.account_tier = account_tier
        """
        The Tier of this storage account.
        """
        if allow_blob_public_access and not isinstance(allow_blob_public_access, bool):
            raise TypeError("Expected argument 'allow_blob_public_access' to be a bool")
        __self__.allow_blob_public_access = allow_blob_public_access
        """
        Is public access allowed to all blobs or containers in the storage account?
        """
        if custom_domains and not isinstance(custom_domains, list):
            raise TypeError("Expected argument 'custom_domains' to be a list")
        __self__.custom_domains = custom_domains
        """
        A `custom_domain` block as documented below.
        """
        if enable_https_traffic_only and not isinstance(enable_https_traffic_only, bool):
            raise TypeError("Expected argument 'enable_https_traffic_only' to be a bool")
        __self__.enable_https_traffic_only = enable_https_traffic_only
        """
        Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
        for more information.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if is_hns_enabled and not isinstance(is_hns_enabled, bool):
            raise TypeError("Expected argument 'is_hns_enabled' to be a bool")
        __self__.is_hns_enabled = is_hns_enabled
        """
        Is Hierarchical Namespace enabled?
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure location where the Storage Account exists
        """
        if min_tls_version and not isinstance(min_tls_version, str):
            raise TypeError("Expected argument 'min_tls_version' to be a str")
        __self__.min_tls_version = min_tls_version
        """
        The minimum supported TLS version for this storage account.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The Custom Domain Name used for the Storage Account.
        """
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        __self__.primary_access_key = primary_access_key
        """
        The primary access key for the Storage Account.
        """
        if primary_blob_connection_string and not isinstance(primary_blob_connection_string, str):
            raise TypeError("Expected argument 'primary_blob_connection_string' to be a str")
        __self__.primary_blob_connection_string = primary_blob_connection_string
        """
        The connection string associated with the primary blob location
        """
        if primary_blob_endpoint and not isinstance(primary_blob_endpoint, str):
            raise TypeError("Expected argument 'primary_blob_endpoint' to be a str")
        __self__.primary_blob_endpoint = primary_blob_endpoint
        """
        The endpoint URL for blob storage in the primary location.
        """
        if primary_blob_host and not isinstance(primary_blob_host, str):
            raise TypeError("Expected argument 'primary_blob_host' to be a str")
        __self__.primary_blob_host = primary_blob_host
        """
        The hostname with port if applicable for blob storage in the primary location.
        """
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        __self__.primary_connection_string = primary_connection_string
        """
        The connection string associated with the primary location
        """
        if primary_dfs_endpoint and not isinstance(primary_dfs_endpoint, str):
            raise TypeError("Expected argument 'primary_dfs_endpoint' to be a str")
        __self__.primary_dfs_endpoint = primary_dfs_endpoint
        """
        The endpoint URL for DFS storage in the primary location.
        """
        if primary_dfs_host and not isinstance(primary_dfs_host, str):
            raise TypeError("Expected argument 'primary_dfs_host' to be a str")
        __self__.primary_dfs_host = primary_dfs_host
        """
        The hostname with port if applicable for DFS storage in the primary location.
        """
        if primary_file_endpoint and not isinstance(primary_file_endpoint, str):
            raise TypeError("Expected argument 'primary_file_endpoint' to be a str")
        __self__.primary_file_endpoint = primary_file_endpoint
        """
        The endpoint URL for file storage in the primary location.
        """
        if primary_file_host and not isinstance(primary_file_host, str):
            raise TypeError("Expected argument 'primary_file_host' to be a str")
        __self__.primary_file_host = primary_file_host
        """
        The hostname with port if applicable for file storage in the primary location.
        """
        if primary_location and not isinstance(primary_location, str):
            raise TypeError("Expected argument 'primary_location' to be a str")
        __self__.primary_location = primary_location
        """
        The primary location of the Storage Account.
        """
        if primary_queue_endpoint and not isinstance(primary_queue_endpoint, str):
            raise TypeError("Expected argument 'primary_queue_endpoint' to be a str")
        __self__.primary_queue_endpoint = primary_queue_endpoint
        """
        The endpoint URL for queue storage in the primary location.
        """
        if primary_queue_host and not isinstance(primary_queue_host, str):
            raise TypeError("Expected argument 'primary_queue_host' to be a str")
        __self__.primary_queue_host = primary_queue_host
        """
        The hostname with port if applicable for queue storage in the primary location.
        """
        if primary_table_endpoint and not isinstance(primary_table_endpoint, str):
            raise TypeError("Expected argument 'primary_table_endpoint' to be a str")
        __self__.primary_table_endpoint = primary_table_endpoint
        """
        The endpoint URL for table storage in the primary location.
        """
        if primary_table_host and not isinstance(primary_table_host, str):
            raise TypeError("Expected argument 'primary_table_host' to be a str")
        __self__.primary_table_host = primary_table_host
        """
        The hostname with port if applicable for table storage in the primary location.
        """
        if primary_web_endpoint and not isinstance(primary_web_endpoint, str):
            raise TypeError("Expected argument 'primary_web_endpoint' to be a str")
        __self__.primary_web_endpoint = primary_web_endpoint
        """
        The endpoint URL for web storage in the primary location.
        """
        if primary_web_host and not isinstance(primary_web_host, str):
            raise TypeError("Expected argument 'primary_web_host' to be a str")
        __self__.primary_web_host = primary_web_host
        """
        The hostname with port if applicable for web storage in the primary location.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        __self__.secondary_access_key = secondary_access_key
        """
        The secondary access key for the Storage Account.
        """
        if secondary_blob_connection_string and not isinstance(secondary_blob_connection_string, str):
            raise TypeError("Expected argument 'secondary_blob_connection_string' to be a str")
        __self__.secondary_blob_connection_string = secondary_blob_connection_string
        """
        The connection string associated with the secondary blob location
        """
        if secondary_blob_endpoint and not isinstance(secondary_blob_endpoint, str):
            raise TypeError("Expected argument 'secondary_blob_endpoint' to be a str")
        __self__.secondary_blob_endpoint = secondary_blob_endpoint
        """
        The endpoint URL for blob storage in the secondary location.
        """
        if secondary_blob_host and not isinstance(secondary_blob_host, str):
            raise TypeError("Expected argument 'secondary_blob_host' to be a str")
        __self__.secondary_blob_host = secondary_blob_host
        """
        The hostname with port if applicable for blob storage in the secondary location.
        """
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        __self__.secondary_connection_string = secondary_connection_string
        """
        The connection string associated with the secondary location
        """
        if secondary_dfs_endpoint and not isinstance(secondary_dfs_endpoint, str):
            raise TypeError("Expected argument 'secondary_dfs_endpoint' to be a str")
        __self__.secondary_dfs_endpoint = secondary_dfs_endpoint
        """
        The endpoint URL for DFS storage in the secondary location.
        """
        if secondary_dfs_host and not isinstance(secondary_dfs_host, str):
            raise TypeError("Expected argument 'secondary_dfs_host' to be a str")
        __self__.secondary_dfs_host = secondary_dfs_host
        """
        The hostname with port if applicable for DFS storage in the secondary location.
        """
        if secondary_file_endpoint and not isinstance(secondary_file_endpoint, str):
            raise TypeError("Expected argument 'secondary_file_endpoint' to be a str")
        __self__.secondary_file_endpoint = secondary_file_endpoint
        """
        The endpoint URL for file storage in the secondary location.
        """
        if secondary_file_host and not isinstance(secondary_file_host, str):
            raise TypeError("Expected argument 'secondary_file_host' to be a str")
        __self__.secondary_file_host = secondary_file_host
        """
        The hostname with port if applicable for file storage in the secondary location.
        """
        if secondary_location and not isinstance(secondary_location, str):
            raise TypeError("Expected argument 'secondary_location' to be a str")
        __self__.secondary_location = secondary_location
        """
        The secondary location of the Storage Account.
        """
        if secondary_queue_endpoint and not isinstance(secondary_queue_endpoint, str):
            raise TypeError("Expected argument 'secondary_queue_endpoint' to be a str")
        __self__.secondary_queue_endpoint = secondary_queue_endpoint
        """
        The endpoint URL for queue storage in the secondary location.
        """
        if secondary_queue_host and not isinstance(secondary_queue_host, str):
            raise TypeError("Expected argument 'secondary_queue_host' to be a str")
        __self__.secondary_queue_host = secondary_queue_host
        """
        The hostname with port if applicable for queue storage in the secondary location.
        """
        if secondary_table_endpoint and not isinstance(secondary_table_endpoint, str):
            raise TypeError("Expected argument 'secondary_table_endpoint' to be a str")
        __self__.secondary_table_endpoint = secondary_table_endpoint
        """
        The endpoint URL for table storage in the secondary location.
        """
        if secondary_table_host and not isinstance(secondary_table_host, str):
            raise TypeError("Expected argument 'secondary_table_host' to be a str")
        __self__.secondary_table_host = secondary_table_host
        """
        The hostname with port if applicable for table storage in the secondary location.
        """
        if secondary_web_endpoint and not isinstance(secondary_web_endpoint, str):
            raise TypeError("Expected argument 'secondary_web_endpoint' to be a str")
        __self__.secondary_web_endpoint = secondary_web_endpoint
        """
        The endpoint URL for web storage in the secondary location.
        """
        if secondary_web_host and not isinstance(secondary_web_host, str):
            raise TypeError("Expected argument 'secondary_web_host' to be a str")
        __self__.secondary_web_host = secondary_web_host
        """
        The hostname with port if applicable for web storage in the secondary location.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assigned to the resource.
        """


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            access_tier=self.access_tier,
            account_kind=self.account_kind,
            account_replication_type=self.account_replication_type,
            account_tier=self.account_tier,
            allow_blob_public_access=self.allow_blob_public_access,
            custom_domains=self.custom_domains,
            enable_https_traffic_only=self.enable_https_traffic_only,
            id=self.id,
            is_hns_enabled=self.is_hns_enabled,
            location=self.location,
            min_tls_version=self.min_tls_version,
            name=self.name,
            primary_access_key=self.primary_access_key,
            primary_blob_connection_string=self.primary_blob_connection_string,
            primary_blob_endpoint=self.primary_blob_endpoint,
            primary_blob_host=self.primary_blob_host,
            primary_connection_string=self.primary_connection_string,
            primary_dfs_endpoint=self.primary_dfs_endpoint,
            primary_dfs_host=self.primary_dfs_host,
            primary_file_endpoint=self.primary_file_endpoint,
            primary_file_host=self.primary_file_host,
            primary_location=self.primary_location,
            primary_queue_endpoint=self.primary_queue_endpoint,
            primary_queue_host=self.primary_queue_host,
            primary_table_endpoint=self.primary_table_endpoint,
            primary_table_host=self.primary_table_host,
            primary_web_endpoint=self.primary_web_endpoint,
            primary_web_host=self.primary_web_host,
            resource_group_name=self.resource_group_name,
            secondary_access_key=self.secondary_access_key,
            secondary_blob_connection_string=self.secondary_blob_connection_string,
            secondary_blob_endpoint=self.secondary_blob_endpoint,
            secondary_blob_host=self.secondary_blob_host,
            secondary_connection_string=self.secondary_connection_string,
            secondary_dfs_endpoint=self.secondary_dfs_endpoint,
            secondary_dfs_host=self.secondary_dfs_host,
            secondary_file_endpoint=self.secondary_file_endpoint,
            secondary_file_host=self.secondary_file_host,
            secondary_location=self.secondary_location,
            secondary_queue_endpoint=self.secondary_queue_endpoint,
            secondary_queue_host=self.secondary_queue_host,
            secondary_table_endpoint=self.secondary_table_endpoint,
            secondary_table_host=self.secondary_table_host,
            secondary_web_endpoint=self.secondary_web_endpoint,
            secondary_web_host=self.secondary_web_host,
            tags=self.tags)


def get_account(min_tls_version: Optional[str] = None,
                name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Use this data source to access information about an existing Storage Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.storage.get_account(name="packerimages",
        resource_group_name="packer-storage")
    pulumi.export("storageAccountTier", example.account_tier)
    ```


    :param str min_tls_version: The minimum supported TLS version for this storage account.
    :param str name: Specifies the name of the Storage Account
    :param str resource_group_name: Specifies the name of the resource group the Storage Account is located in.
    """
    __args__ = dict()
    __args__['minTlsVersion'] = min_tls_version
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:storage/getAccount:getAccount', __args__, opts=opts).value

    return AwaitableGetAccountResult(
        access_tier=__ret__.get('accessTier'),
        account_kind=__ret__.get('accountKind'),
        account_replication_type=__ret__.get('accountReplicationType'),
        account_tier=__ret__.get('accountTier'),
        allow_blob_public_access=__ret__.get('allowBlobPublicAccess'),
        custom_domains=__ret__.get('customDomains'),
        enable_https_traffic_only=__ret__.get('enableHttpsTrafficOnly'),
        id=__ret__.get('id'),
        is_hns_enabled=__ret__.get('isHnsEnabled'),
        location=__ret__.get('location'),
        min_tls_version=__ret__.get('minTlsVersion'),
        name=__ret__.get('name'),
        primary_access_key=__ret__.get('primaryAccessKey'),
        primary_blob_connection_string=__ret__.get('primaryBlobConnectionString'),
        primary_blob_endpoint=__ret__.get('primaryBlobEndpoint'),
        primary_blob_host=__ret__.get('primaryBlobHost'),
        primary_connection_string=__ret__.get('primaryConnectionString'),
        primary_dfs_endpoint=__ret__.get('primaryDfsEndpoint'),
        primary_dfs_host=__ret__.get('primaryDfsHost'),
        primary_file_endpoint=__ret__.get('primaryFileEndpoint'),
        primary_file_host=__ret__.get('primaryFileHost'),
        primary_location=__ret__.get('primaryLocation'),
        primary_queue_endpoint=__ret__.get('primaryQueueEndpoint'),
        primary_queue_host=__ret__.get('primaryQueueHost'),
        primary_table_endpoint=__ret__.get('primaryTableEndpoint'),
        primary_table_host=__ret__.get('primaryTableHost'),
        primary_web_endpoint=__ret__.get('primaryWebEndpoint'),
        primary_web_host=__ret__.get('primaryWebHost'),
        resource_group_name=__ret__.get('resourceGroupName'),
        secondary_access_key=__ret__.get('secondaryAccessKey'),
        secondary_blob_connection_string=__ret__.get('secondaryBlobConnectionString'),
        secondary_blob_endpoint=__ret__.get('secondaryBlobEndpoint'),
        secondary_blob_host=__ret__.get('secondaryBlobHost'),
        secondary_connection_string=__ret__.get('secondaryConnectionString'),
        secondary_dfs_endpoint=__ret__.get('secondaryDfsEndpoint'),
        secondary_dfs_host=__ret__.get('secondaryDfsHost'),
        secondary_file_endpoint=__ret__.get('secondaryFileEndpoint'),
        secondary_file_host=__ret__.get('secondaryFileHost'),
        secondary_location=__ret__.get('secondaryLocation'),
        secondary_queue_endpoint=__ret__.get('secondaryQueueEndpoint'),
        secondary_queue_host=__ret__.get('secondaryQueueHost'),
        secondary_table_endpoint=__ret__.get('secondaryTableEndpoint'),
        secondary_table_host=__ret__.get('secondaryTableHost'),
        secondary_web_endpoint=__ret__.get('secondaryWebEndpoint'),
        secondary_web_host=__ret__.get('secondaryWebHost'),
        tags=__ret__.get('tags'))
