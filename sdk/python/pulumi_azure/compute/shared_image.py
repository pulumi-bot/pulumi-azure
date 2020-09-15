# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SharedImage']


class SharedImage(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 eula: Optional[pulumi.Input[str]] = None,
                 gallery_name: Optional[pulumi.Input[str]] = None,
                 hyper_v_generation: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[pulumi.InputType['SharedImageIdentifierArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_type: Optional[pulumi.Input[str]] = None,
                 privacy_statement_uri: Optional[pulumi.Input[str]] = None,
                 purchase_plan: Optional[pulumi.Input[pulumi.InputType['SharedImagePurchasePlanArgs']]] = None,
                 release_note_uri: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 specialized: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Shared Image within a Shared Image Gallery.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_shared_image_gallery = azure.compute.SharedImageGallery("exampleSharedImageGallery",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            description="Shared images and things.",
            tags={
                "Hello": "There",
                "World": "Example",
            })
        example_shared_image = azure.compute.SharedImage("exampleSharedImage",
            gallery_name=example_shared_image_gallery.name,
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            os_type="Linux",
            identifier=azure.compute.SharedImageIdentifierArgs(
                publisher="PublisherName",
                offer="OfferName",
                sku="ExampleSku",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of this Shared Image.
        :param pulumi.Input[str] eula: The End User Licence Agreement for the Shared Image.
        :param pulumi.Input[str] gallery_name: Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] hyper_v_generation: The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['SharedImageIdentifierArgs']] identifier: An `identifier` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Shared Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] privacy_statement_uri: The URI containing the Privacy Statement associated with this Shared Image.
        :param pulumi.Input[pulumi.InputType['SharedImagePurchasePlanArgs']] purchase_plan: A `purchase_plan` block as defined below.
        :param pulumi.Input[str] release_note_uri: The URI containing the Release Notes associated with this Shared Image.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] specialized: Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Shared Image.
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

            __props__['description'] = description
            __props__['eula'] = eula
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            __props__['hyper_v_generation'] = hyper_v_generation
            if identifier is None:
                raise TypeError("Missing required property 'identifier'")
            __props__['identifier'] = identifier
            __props__['location'] = location
            __props__['name'] = name
            if os_type is None:
                raise TypeError("Missing required property 'os_type'")
            __props__['os_type'] = os_type
            __props__['privacy_statement_uri'] = privacy_statement_uri
            __props__['purchase_plan'] = purchase_plan
            __props__['release_note_uri'] = release_note_uri
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['specialized'] = specialized
            __props__['tags'] = tags
        super(SharedImage, __self__).__init__(
            'azure:compute/sharedImage:SharedImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            eula: Optional[pulumi.Input[str]] = None,
            gallery_name: Optional[pulumi.Input[str]] = None,
            hyper_v_generation: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[pulumi.InputType['SharedImageIdentifierArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            os_type: Optional[pulumi.Input[str]] = None,
            privacy_statement_uri: Optional[pulumi.Input[str]] = None,
            purchase_plan: Optional[pulumi.Input[pulumi.InputType['SharedImagePurchasePlanArgs']]] = None,
            release_note_uri: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            specialized: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SharedImage':
        """
        Get an existing SharedImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of this Shared Image.
        :param pulumi.Input[str] eula: The End User Licence Agreement for the Shared Image.
        :param pulumi.Input[str] gallery_name: Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] hyper_v_generation: The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['SharedImageIdentifierArgs']] identifier: An `identifier` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Shared Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] privacy_statement_uri: The URI containing the Privacy Statement associated with this Shared Image.
        :param pulumi.Input[pulumi.InputType['SharedImagePurchasePlanArgs']] purchase_plan: A `purchase_plan` block as defined below.
        :param pulumi.Input[str] release_note_uri: The URI containing the Release Notes associated with this Shared Image.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] specialized: Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Shared Image.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["eula"] = eula
        __props__["gallery_name"] = gallery_name
        __props__["hyper_v_generation"] = hyper_v_generation
        __props__["identifier"] = identifier
        __props__["location"] = location
        __props__["name"] = name
        __props__["os_type"] = os_type
        __props__["privacy_statement_uri"] = privacy_statement_uri
        __props__["purchase_plan"] = purchase_plan
        __props__["release_note_uri"] = release_note_uri
        __props__["resource_group_name"] = resource_group_name
        __props__["specialized"] = specialized
        __props__["tags"] = tags
        return SharedImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of this Shared Image.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def eula(self) -> pulumi.Output[Optional[str]]:
        """
        The End User Licence Agreement for the Shared Image.
        """
        return pulumi.get(self, "eula")

    @property
    @pulumi.getter(name="galleryName")
    def gallery_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "gallery_name")

    @property
    @pulumi.getter(name="hyperVGeneration")
    def hyper_v_generation(self) -> pulumi.Output[Optional[str]]:
        """
        The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "hyper_v_generation")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output['outputs.SharedImageIdentifier']:
        """
        An `identifier` block as defined below.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Shared Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[str]:
        """
        The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="privacyStatementUri")
    def privacy_statement_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The URI containing the Privacy Statement associated with this Shared Image.
        """
        return pulumi.get(self, "privacy_statement_uri")

    @property
    @pulumi.getter(name="purchasePlan")
    def purchase_plan(self) -> pulumi.Output[Optional['outputs.SharedImagePurchasePlan']]:
        """
        A `purchase_plan` block as defined below.
        """
        return pulumi.get(self, "purchase_plan")

    @property
    @pulumi.getter(name="releaseNoteUri")
    def release_note_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The URI containing the Release Notes associated with this Shared Image.
        """
        return pulumi.get(self, "release_note_uri")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def specialized(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "specialized")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the Shared Image.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

