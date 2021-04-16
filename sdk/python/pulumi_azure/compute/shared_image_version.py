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

__all__ = ['SharedImageVersionArgs', 'SharedImageVersion']

@pulumi.input_type
class SharedImageVersionArgs:
    def __init__(__self__, *,
                 gallery_name: pulumi.Input[str],
                 image_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 target_regions: pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]],
                 exclude_from_latest: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_image_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_disk_snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SharedImageVersion resource.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] image_name: The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]] target_regions: One or more `target_region` blocks as documented below.
        :param pulumi.Input[bool] exclude_from_latest: Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        :param pulumi.Input[str] location: The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_image_id: The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_disk_snapshot_id: The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A collection of tags which should be applied to this resource.
        """
        pulumi.set(__self__, "gallery_name", gallery_name)
        pulumi.set(__self__, "image_name", image_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "target_regions", target_regions)
        if exclude_from_latest is not None:
            pulumi.set(__self__, "exclude_from_latest", exclude_from_latest)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed_image_id is not None:
            pulumi.set(__self__, "managed_image_id", managed_image_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os_disk_snapshot_id is not None:
            pulumi.set(__self__, "os_disk_snapshot_id", os_disk_snapshot_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="galleryName")
    def gallery_name(self) -> pulumi.Input[str]:
        """
        The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "gallery_name")

    @gallery_name.setter
    def gallery_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "gallery_name", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Input[str]:
        """
        The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="targetRegions")
    def target_regions(self) -> pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]]:
        """
        One or more `target_region` blocks as documented below.
        """
        return pulumi.get(self, "target_regions")

    @target_regions.setter
    def target_regions(self, value: pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]]):
        pulumi.set(self, "target_regions", value)

    @property
    @pulumi.getter(name="excludeFromLatest")
    def exclude_from_latest(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        """
        return pulumi.get(self, "exclude_from_latest")

    @exclude_from_latest.setter
    def exclude_from_latest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_from_latest", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managedImageId")
    def managed_image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "managed_image_id")

    @managed_image_id.setter
    def managed_image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_image_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osDiskSnapshotId")
    def os_disk_snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "os_disk_snapshot_id")

    @os_disk_snapshot_id.setter
    def os_disk_snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os_disk_snapshot_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A collection of tags which should be applied to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SharedImageVersionState:
    def __init__(__self__, *,
                 exclude_from_latest: Optional[pulumi.Input[bool]] = None,
                 gallery_name: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_image_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_disk_snapshot_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_regions: Optional[pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]]] = None):
        """
        Input properties used for looking up and filtering SharedImageVersion resources.
        :param pulumi.Input[bool] exclude_from_latest: Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] image_name: The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_image_id: The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_disk_snapshot_id: The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A collection of tags which should be applied to this resource.
        :param pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]] target_regions: One or more `target_region` blocks as documented below.
        """
        if exclude_from_latest is not None:
            pulumi.set(__self__, "exclude_from_latest", exclude_from_latest)
        if gallery_name is not None:
            pulumi.set(__self__, "gallery_name", gallery_name)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed_image_id is not None:
            pulumi.set(__self__, "managed_image_id", managed_image_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os_disk_snapshot_id is not None:
            pulumi.set(__self__, "os_disk_snapshot_id", os_disk_snapshot_id)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_regions is not None:
            pulumi.set(__self__, "target_regions", target_regions)

    @property
    @pulumi.getter(name="excludeFromLatest")
    def exclude_from_latest(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        """
        return pulumi.get(self, "exclude_from_latest")

    @exclude_from_latest.setter
    def exclude_from_latest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_from_latest", value)

    @property
    @pulumi.getter(name="galleryName")
    def gallery_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "gallery_name")

    @gallery_name.setter
    def gallery_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gallery_name", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managedImageId")
    def managed_image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "managed_image_id")

    @managed_image_id.setter
    def managed_image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_image_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osDiskSnapshotId")
    def os_disk_snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "os_disk_snapshot_id")

    @os_disk_snapshot_id.setter
    def os_disk_snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os_disk_snapshot_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A collection of tags which should be applied to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetRegions")
    def target_regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]]]:
        """
        One or more `target_region` blocks as documented below.
        """
        return pulumi.get(self, "target_regions")

    @target_regions.setter
    def target_regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SharedImageVersionTargetRegionArgs']]]]):
        pulumi.set(self, "target_regions", value)


class SharedImageVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
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
                 target_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]]] = None,
                 __props__=None):
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
            target_regions=[azure.compute.SharedImageVersionTargetRegionArgs(
                name=existing_shared_image.location,
                regional_replica_count=5,
                storage_account_type="Standard_LRS",
            )])
        ```

        ## Import

        Shared Image Versions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/sharedImageVersion:SharedImageVersion version /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
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
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]] target_regions: One or more `target_region` blocks as documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SharedImageVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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
            target_regions=[azure.compute.SharedImageVersionTargetRegionArgs(
                name=existing_shared_image.location,
                regional_replica_count=5,
                storage_account_type="Standard_LRS",
            )])
        ```

        ## Import

        Shared Image Versions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/sharedImageVersion:SharedImageVersion version /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
        ```

        :param str resource_name: The name of the resource.
        :param SharedImageVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SharedImageVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
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
                 target_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]]] = None,
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
            __props__ = SharedImageVersionArgs.__new__(SharedImageVersionArgs)

            __props__.__dict__["exclude_from_latest"] = exclude_from_latest
            if gallery_name is None and not opts.urn:
                raise TypeError("Missing required property 'gallery_name'")
            __props__.__dict__["gallery_name"] = gallery_name
            if image_name is None and not opts.urn:
                raise TypeError("Missing required property 'image_name'")
            __props__.__dict__["image_name"] = image_name
            __props__.__dict__["location"] = location
            __props__.__dict__["managed_image_id"] = managed_image_id
            __props__.__dict__["name"] = name
            __props__.__dict__["os_disk_snapshot_id"] = os_disk_snapshot_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if target_regions is None and not opts.urn:
                raise TypeError("Missing required property 'target_regions'")
            __props__.__dict__["target_regions"] = target_regions
        super(SharedImageVersion, __self__).__init__(
            'azure:compute/sharedImageVersion:SharedImageVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
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
            target_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]]] = None) -> 'SharedImageVersion':
        """
        Get an existing SharedImageVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SharedImageVersionTargetRegionArgs']]]] target_regions: One or more `target_region` blocks as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SharedImageVersionState.__new__(_SharedImageVersionState)

        __props__.__dict__["exclude_from_latest"] = exclude_from_latest
        __props__.__dict__["gallery_name"] = gallery_name
        __props__.__dict__["image_name"] = image_name
        __props__.__dict__["location"] = location
        __props__.__dict__["managed_image_id"] = managed_image_id
        __props__.__dict__["name"] = name
        __props__.__dict__["os_disk_snapshot_id"] = os_disk_snapshot_id
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["target_regions"] = target_regions
        return SharedImageVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="excludeFromLatest")
    def exclude_from_latest(self) -> pulumi.Output[Optional[bool]]:
        """
        Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        """
        return pulumi.get(self, "exclude_from_latest")

    @property
    @pulumi.getter(name="galleryName")
    def gallery_name(self) -> pulumi.Output[str]:
        """
        The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "gallery_name")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Output[str]:
        """
        The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedImageId")
    def managed_image_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "managed_image_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osDiskSnapshotId")
    def os_disk_snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "os_disk_snapshot_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A collection of tags which should be applied to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetRegions")
    def target_regions(self) -> pulumi.Output[Sequence['outputs.SharedImageVersionTargetRegion']]:
        """
        One or more `target_region` blocks as documented below.
        """
        return pulumi.get(self, "target_regions")

