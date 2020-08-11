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

__all__ = ['Hub']


class Hub(pulumi.CustomResource):
    apns_credential: pulumi.Output[Optional['outputs.HubApnsCredential']] = pulumi.property("apnsCredential")
    """
    A `apns_credential` block as defined below.
    """

    gcm_credential: pulumi.Output[Optional['outputs.HubGcmCredential']] = pulumi.property("gcmCredential")
    """
    A `gcm_credential` block as defined below.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name to use for this Notification Hub. Changing this forces a new resource to be created.
    """

    namespace_name: pulumi.Output[str] = pulumi.property("namespaceName")
    """
    The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apns_credential: Optional[pulumi.Input[pulumi.InputType['HubApnsCredentialArgs']]] = None,
                 gcm_credential: Optional[pulumi.Input[pulumi.InputType['HubGcmCredentialArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Notification Hub within a Notification Hub Namespace.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="Australia East")
        example_namespace = azure.notificationhub.Namespace("exampleNamespace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            namespace_type="NotificationHub",
            sku_name="Free")
        example_hub = azure.notificationhub.Hub("exampleHub",
            namespace_name=example_namespace.name,
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['HubApnsCredentialArgs']] apns_credential: A `apns_credential` block as defined below.
        :param pulumi.Input[pulumi.InputType['HubGcmCredentialArgs']] gcm_credential: A `gcm_credential` block as defined below.
        :param pulumi.Input[str] location: The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Notification Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
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

            __props__['apns_credential'] = apns_credential
            __props__['gcm_credential'] = gcm_credential
            __props__['location'] = location
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Hub, __self__).__init__(
            'azure:notificationhub/hub:Hub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            apns_credential: Optional[pulumi.Input[pulumi.InputType['HubApnsCredentialArgs']]] = None,
            gcm_credential: Optional[pulumi.Input[pulumi.InputType['HubGcmCredentialArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Hub':
        """
        Get an existing Hub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['HubApnsCredentialArgs']] apns_credential: A `apns_credential` block as defined below.
        :param pulumi.Input[pulumi.InputType['HubGcmCredentialArgs']] gcm_credential: A `gcm_credential` block as defined below.
        :param pulumi.Input[str] location: The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name to use for this Notification Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["apns_credential"] = apns_credential
        __props__["gcm_credential"] = gcm_credential
        __props__["location"] = location
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Hub(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

