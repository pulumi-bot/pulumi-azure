# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ShareArgs', 'Share']

@pulumi.input_type
class ShareArgs:
    def __init__(__self__, *,
                 storage_account_name: pulumi.Input[str],
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Share resource.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the share.
               Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]] acls: One or more `acl` blocks as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A mapping of MetaData for this File Share.
        :param pulumi.Input[str] name: The name of the share. Must be unique within the storage account where the share is located.
        :param pulumi.Input[int] quota: The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        """
        pulumi.set(__self__, "storage_account_name", storage_account_name)
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if quota is not None:
            pulumi.set(__self__, "quota", quota)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Input[str]:
        """
        Specifies the storage account in which to create the share.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]]]:
        """
        One or more `acl` blocks as defined below.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of MetaData for this File Share.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the share. Must be unique within the storage account where the share is located.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def quota(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        """
        return pulumi.get(self, "quota")

    @quota.setter
    def quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "quota", value)


@pulumi.input_type
class _ShareState:
    def __init__(__self__, *,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[int]] = None,
                 resource_manager_id: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Share resources.
        :param pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]] acls: One or more `acl` blocks as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A mapping of MetaData for this File Share.
        :param pulumi.Input[str] name: The name of the share. Must be unique within the storage account where the share is located.
        :param pulumi.Input[int] quota: The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        :param pulumi.Input[str] resource_manager_id: The Resource Manager ID of this File Share.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the share.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] url: The URL of the File Share
        """
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if quota is not None:
            pulumi.set(__self__, "quota", quota)
        if resource_manager_id is not None:
            pulumi.set(__self__, "resource_manager_id", resource_manager_id)
        if storage_account_name is not None:
            pulumi.set(__self__, "storage_account_name", storage_account_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]]]:
        """
        One or more `acl` blocks as defined below.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ShareAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of MetaData for this File Share.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the share. Must be unique within the storage account where the share is located.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def quota(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        """
        return pulumi.get(self, "quota")

    @quota.setter
    def quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "quota", value)

    @property
    @pulumi.getter(name="resourceManagerId")
    def resource_manager_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Resource Manager ID of this File Share.
        """
        return pulumi.get(self, "resource_manager_id")

    @resource_manager_id.setter
    def resource_manager_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_manager_id", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the storage account in which to create the share.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the File Share
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Share(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShareAclArgs']]]]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[int]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a File Share within Azure Storage.

        > **Note:** The storage share supports two storage tiers: premium and standard. Standard file shares are created in general purpose (GPv1 or GPv2) storage accounts and premium file shares are created in FileStorage storage accounts. For further information, refer to the section "What storage tiers are supported in Azure Files?" of [documentation](https://docs.microsoft.com/en-us/azure/storage/files/storage-files-faq#general).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_share = azure.storage.Share("exampleShare",
            storage_account_name=example_account.name,
            quota=50,
            acls=[azure.storage.ShareAclArgs(
                id="MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI",
                access_policies=[azure.storage.ShareAclAccessPolicyArgs(
                    permissions="rwdl",
                    start="2019-07-02T09:38:21.0000000Z",
                    expiry="2019-07-02T10:38:21.0000000Z",
                )],
            )])
        ```

        ## Import

        Storage Shares can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:storage/share:Share exampleShare https://account1.file.core.windows.net/share1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShareAclArgs']]]] acls: One or more `acl` blocks as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A mapping of MetaData for this File Share.
        :param pulumi.Input[str] name: The name of the share. Must be unique within the storage account where the share is located.
        :param pulumi.Input[int] quota: The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the share.
               Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ShareArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a File Share within Azure Storage.

        > **Note:** The storage share supports two storage tiers: premium and standard. Standard file shares are created in general purpose (GPv1 or GPv2) storage accounts and premium file shares are created in FileStorage storage accounts. For further information, refer to the section "What storage tiers are supported in Azure Files?" of [documentation](https://docs.microsoft.com/en-us/azure/storage/files/storage-files-faq#general).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_share = azure.storage.Share("exampleShare",
            storage_account_name=example_account.name,
            quota=50,
            acls=[azure.storage.ShareAclArgs(
                id="MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI",
                access_policies=[azure.storage.ShareAclAccessPolicyArgs(
                    permissions="rwdl",
                    start="2019-07-02T09:38:21.0000000Z",
                    expiry="2019-07-02T10:38:21.0000000Z",
                )],
            )])
        ```

        ## Import

        Storage Shares can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:storage/share:Share exampleShare https://account1.file.core.windows.net/share1
        ```

        :param str resource_name: The name of the resource.
        :param ShareArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShareArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShareAclArgs']]]]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[int]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ShareArgs.__new__(ShareArgs)

            __props__.__dict__["acls"] = acls
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["quota"] = quota
            if storage_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__.__dict__["storage_account_name"] = storage_account_name
            __props__.__dict__["resource_manager_id"] = None
            __props__.__dict__["url"] = None
        super(Share, __self__).__init__(
            'azure:storage/share:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShareAclArgs']]]]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            quota: Optional[pulumi.Input[int]] = None,
            resource_manager_id: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Share':
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShareAclArgs']]]] acls: One or more `acl` blocks as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A mapping of MetaData for this File Share.
        :param pulumi.Input[str] name: The name of the share. Must be unique within the storage account where the share is located.
        :param pulumi.Input[int] quota: The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        :param pulumi.Input[str] resource_manager_id: The Resource Manager ID of this File Share.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the share.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] url: The URL of the File Share
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ShareState.__new__(_ShareState)

        __props__.__dict__["acls"] = acls
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["quota"] = quota
        __props__.__dict__["resource_manager_id"] = resource_manager_id
        __props__.__dict__["storage_account_name"] = storage_account_name
        __props__.__dict__["url"] = url
        return Share(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Output[Optional[Sequence['outputs.ShareAcl']]]:
        """
        One or more `acl` blocks as defined below.
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A mapping of MetaData for this File Share.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the share. Must be unique within the storage account where the share is located.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def quota(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
        """
        return pulumi.get(self, "quota")

    @property
    @pulumi.getter(name="resourceManagerId")
    def resource_manager_id(self) -> pulumi.Output[str]:
        """
        The Resource Manager ID of this File Share.
        """
        return pulumi.get(self, "resource_manager_id")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Output[str]:
        """
        Specifies the storage account in which to create the share.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the File Share
        """
        return pulumi.get(self, "url")

