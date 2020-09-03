# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Lab']


class Lab(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Dev Test Lab.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_lab = azure.devtest.Lab("exampleLab",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "Sydney": "Australia",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_type: The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['storage_type'] = storage_type
            __props__['tags'] = tags
            __props__['artifacts_storage_account_id'] = None
            __props__['default_premium_storage_account_id'] = None
            __props__['default_storage_account_id'] = None
            __props__['key_vault_id'] = None
            __props__['premium_data_disk_storage_account_id'] = None
            __props__['unique_identifier'] = None
        super(Lab, __self__).__init__(
            'azure:devtest/lab:Lab',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            artifacts_storage_account_id: Optional[pulumi.Input[str]] = None,
            default_premium_storage_account_id: Optional[pulumi.Input[str]] = None,
            default_storage_account_id: Optional[pulumi.Input[str]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            premium_data_disk_storage_account_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            unique_identifier: Optional[pulumi.Input[str]] = None) -> 'Lab':
        """
        Get an existing Lab resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] artifacts_storage_account_id: The ID of the Storage Account used for Artifact Storage.
        :param pulumi.Input[str] default_premium_storage_account_id: The ID of the Default Premium Storage Account for this Dev Test Lab.
        :param pulumi.Input[str] default_storage_account_id: The ID of the Default Storage Account for this Dev Test Lab.
        :param pulumi.Input[str] key_vault_id: The ID of the Key used for this Dev Test Lab.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
        :param pulumi.Input[str] premium_data_disk_storage_account_id: The ID of the Storage Account used for Storage of Premium Data Disk.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_type: The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] unique_identifier: The unique immutable identifier of the Dev Test Lab.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["artifacts_storage_account_id"] = artifacts_storage_account_id
        __props__["default_premium_storage_account_id"] = default_premium_storage_account_id
        __props__["default_storage_account_id"] = default_storage_account_id
        __props__["key_vault_id"] = key_vault_id
        __props__["location"] = location
        __props__["name"] = name
        __props__["premium_data_disk_storage_account_id"] = premium_data_disk_storage_account_id
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_type"] = storage_type
        __props__["tags"] = tags
        __props__["unique_identifier"] = unique_identifier
        return Lab(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="artifactsStorageAccountId")
    def artifacts_storage_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the Storage Account used for Artifact Storage.
        """
        return pulumi.get(self, "artifacts_storage_account_id")

    @property
    @pulumi.getter(name="defaultPremiumStorageAccountId")
    def default_premium_storage_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the Default Premium Storage Account for this Dev Test Lab.
        """
        return pulumi.get(self, "default_premium_storage_account_id")

    @property
    @pulumi.getter(name="defaultStorageAccountId")
    def default_storage_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the Default Storage Account for this Dev Test Lab.
        """
        return pulumi.get(self, "default_storage_account_id")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the Key used for this Dev Test Lab.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="premiumDataDiskStorageAccountId")
    def premium_data_disk_storage_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the Storage Account used for Storage of Premium Data Disk.
        """
        return pulumi.get(self, "premium_data_disk_storage_account_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[str]:
        """
        The unique immutable identifier of the Dev Test Lab.
        """
        return pulumi.get(self, "unique_identifier")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

