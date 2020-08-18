# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SharedImageVersion']


class SharedImageVersion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exclude_from_latest: Optional[pulumi.Input[bool]] = None,
                 gallery_name: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_image_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_disk_snapshot_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_regions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Version of a Shared Image within a Shared Image Gallery.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        existing_image = azure.compute.get_image(name="search-api",
            resource_group_name="packerimages")
        existing_shared_image = azure.compute.get_shared_image(name="existing-image",
            gallery_name="existing_gallery",
            resource_group_name="existing-resources")
        example = azure.compute.SharedImageVersion("example",
            gallery_name=existing_shared_image.gallery_name,
            image_name=existing_shared_image.name,
            resource_group_name=existing_shared_image.resource_group_name,
            location=existing_shared_image.location,
            managed_image_id=existing_image.id,
            target_regions=[{
                "name": existing_shared_image.location,
                "regionalReplicaCount": 5,
                "storage_account_type": "Standard_LRS",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] exclude_from_latest: Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] image_name: The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_image_id: The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_disk_snapshot_id: The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A collection of tags which should be applied to this resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]] target_regions: One or more `target_region` blocks as documented below.
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

            __props__['exclude_from_latest'] = exclude_from_latest
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            if image_name is None:
                raise TypeError("Missing required property 'image_name'")
            __props__['image_name'] = image_name
            __props__['location'] = location
            __props__['managed_image_id'] = managed_image_id
            __props__['name'] = name
            __props__['os_disk_snapshot_id'] = os_disk_snapshot_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if target_regions is None:
                raise TypeError("Missing required property 'target_regions'")
            __props__['target_regions'] = target_regions
        super(SharedImageVersion, __self__).__init__(
            'azure:compute/sharedImageVersion:SharedImageVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            exclude_from_latest: Optional[pulumi.Input[bool]] = None,
            gallery_name: Optional[pulumi.Input[str]] = None,
            image_name: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            managed_image_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            os_disk_snapshot_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_regions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]]] = None) -> 'SharedImageVersion':
        """
        Get an existing SharedImageVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] exclude_from_latest: Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] image_name: The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_image_id: The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_disk_snapshot_id: The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A collection of tags which should be applied to this resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]] target_regions: One or more `target_region` blocks as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["exclude_from_latest"] = exclude_from_latest
        __props__["gallery_name"] = gallery_name
        __props__["image_name"] = image_name
        __props__["location"] = location
        __props__["managed_image_id"] = managed_image_id
        __props__["name"] = name
        __props__["os_disk_snapshot_id"] = os_disk_snapshot_id
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["target_regions"] = target_regions
        return SharedImageVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="excludeFromLatest")
    def exclude_from_latest(self) -> Optional[bool]:
        """
        Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        """
        return pulumi.get(self, "exclude_from_latest")

    @property
    @pulumi.getter(name="galleryName")
    def gallery_name(self) -> str:
        """
        The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "gallery_name")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> str:
        """
        The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedImageId")
    def managed_image_id(self) -> Optional[str]:
        """
        The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "managed_image_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osDiskSnapshotId")
    def os_disk_snapshot_id(self) -> Optional[str]:
        """
        The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "os_disk_snapshot_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A collection of tags which should be applied to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetRegions")
    def target_regions(self) -> List['outputs.SharedImageVersionTargetRegion']:
        """
        One or more `target_region` blocks as documented below.
        """
        return pulumi.get(self, "target_regions")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

