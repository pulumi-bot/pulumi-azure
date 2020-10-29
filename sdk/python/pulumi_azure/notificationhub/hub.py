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

__all__ = ['Hub']


class Hub(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
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
            id: pulumi.Input[str],
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
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

    @property
    @pulumi.getter(name="apnsCredential")
    def apns_credential(self) -> pulumi.Output[Optional['outputs.HubApnsCredential']]:
        """
        A `apns_credential` block as defined below.
        """
        return pulumi.get(self, "apns_credential")

    @property
    @pulumi.getter(name="gcmCredential")
    def gcm_credential(self) -> pulumi.Output[Optional['outputs.HubGcmCredential']]:
        """
        A `gcm_credential` block as defined below.
        """
        return pulumi.get(self, "gcm_credential")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name to use for this Notification Hub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

