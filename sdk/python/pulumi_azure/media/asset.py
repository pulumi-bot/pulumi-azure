# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['AssetArgs', 'Asset']

@pulumi.input_type
class AssetArgs:
    def __init__(__self__, *,
                 media_services_account_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 alternate_id: Optional[pulumi.Input[str]] = None,
                 container: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Asset resource.
        :param pulumi.Input[str] media_services_account_name: Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] alternate_id: The alternate ID of the Asset.
        :param pulumi.Input[str] container: The name of the asset blob container. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] description: The Asset description.
        :param pulumi.Input[str] name: The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] storage_account_name: The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
        """
        pulumi.set(__self__, "media_services_account_name", media_services_account_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if alternate_id is not None:
            pulumi.set(__self__, "alternate_id", alternate_id)
        if container is not None:
            pulumi.set(__self__, "container", container)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if storage_account_name is not None:
            pulumi.set(__self__, "storage_account_name", storage_account_name)

    @property
    @pulumi.getter(name="mediaServicesAccountName")
    def media_services_account_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "media_services_account_name")

    @media_services_account_name.setter
    def media_services_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "media_services_account_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="alternateId")
    def alternate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The alternate ID of the Asset.
        """
        return pulumi.get(self, "alternate_id")

    @alternate_id.setter
    def alternate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alternate_id", value)

    @property
    @pulumi.getter
    def container(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the asset blob container. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Asset description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_name", value)


class Asset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternate_id: Optional[pulumi.Input[str]] = None,
                 container: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 media_services_account_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Media Asset.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_service_account = azure.media.ServiceAccount("exampleServiceAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_accounts=[azure.media.ServiceAccountStorageAccountArgs(
                id=example_account.id,
                is_primary=True,
            )])
        example_asset = azure.media.Asset("exampleAsset",
            resource_group_name=example_resource_group.name,
            media_services_account_name=example_service_account.name,
            description="Asset description")
        ```

        ## Import

        Media Assets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:media/asset:Asset example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/assets/asset1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alternate_id: The alternate ID of the Asset.
        :param pulumi.Input[str] container: The name of the asset blob container. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] description: The Asset description.
        :param pulumi.Input[str] media_services_account_name: Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] name: The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] storage_account_name: The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Media Asset.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_service_account = azure.media.ServiceAccount("exampleServiceAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_accounts=[azure.media.ServiceAccountStorageAccountArgs(
                id=example_account.id,
                is_primary=True,
            )])
        example_asset = azure.media.Asset("exampleAsset",
            resource_group_name=example_resource_group.name,
            media_services_account_name=example_service_account.name,
            description="Asset description")
        ```

        ## Import

        Media Assets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:media/asset:Asset example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/assets/asset1
        ```

        :param str resource_name: The name of the resource.
        :param AssetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternate_id: Optional[pulumi.Input[str]] = None,
                 container: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 media_services_account_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['alternate_id'] = alternate_id
            __props__['container'] = container
            __props__['description'] = description
            if media_services_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'media_services_account_name'")
            __props__['media_services_account_name'] = media_services_account_name
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['storage_account_name'] = storage_account_name
        super(Asset, __self__).__init__(
            'azure:media/asset:Asset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alternate_id: Optional[pulumi.Input[str]] = None,
            container: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            media_services_account_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None) -> 'Asset':
        """
        Get an existing Asset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alternate_id: The alternate ID of the Asset.
        :param pulumi.Input[str] container: The name of the asset blob container. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] description: The Asset description.
        :param pulumi.Input[str] media_services_account_name: Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] name: The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
        :param pulumi.Input[str] storage_account_name: The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alternate_id"] = alternate_id
        __props__["container"] = container
        __props__["description"] = description
        __props__["media_services_account_name"] = media_services_account_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_account_name"] = storage_account_name
        return Asset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alternateId")
    def alternate_id(self) -> pulumi.Output[Optional[str]]:
        """
        The alternate ID of the Asset.
        """
        return pulumi.get(self, "alternate_id")

    @property
    @pulumi.getter
    def container(self) -> pulumi.Output[str]:
        """
        The name of the asset blob container. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "container")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Asset description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="mediaServicesAccountName")
    def media_services_account_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "media_services_account_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Output[str]:
        """
        The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
        """
        return pulumi.get(self, "storage_account_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

