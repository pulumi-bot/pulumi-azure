# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Certificate']


class Certificate(pulumi.CustomResource):
    api_management_name: pulumi.Output[str] = pulumi.property("apiManagementName")
    """
    The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
    """

    data: pulumi.Output[str] = pulumi.property("data")
    """
    The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
    """

    expiration: pulumi.Output[str] = pulumi.property("expiration")
    """
    The Expiration Date of this Certificate, formatted as an RFC3339 string.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the API Management Certificate. Changing this forces a new resource to be created.
    """

    password: pulumi.Output[Optional[str]] = pulumi.property("password")
    """
    The password used for this certificate. Changing this forces a new resource to be created.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
    """

    subject: pulumi.Output[str] = pulumi.property("subject")
    """
    The Subject of this Certificate.
    """

    thumbprint: pulumi.Output[str] = pulumi.property("thumbprint")
    """
    The Thumbprint of this Certificate.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Certificate within an API Management Service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] data: The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the API Management Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] password: The password used for this certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
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

            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if data is None:
                raise TypeError("Missing required property 'data'")
            __props__['data'] = data
            __props__['name'] = name
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['expiration'] = None
            __props__['subject'] = None
            __props__['thumbprint'] = None
        super(Certificate, __self__).__init__(
            'azure:apimanagement/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            data: Optional[pulumi.Input[str]] = None,
            expiration: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subject: Optional[pulumi.Input[str]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] data: The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
        :param pulumi.Input[str] expiration: The Expiration Date of this Certificate, formatted as an RFC3339 string.
        :param pulumi.Input[str] name: The name of the API Management Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] password: The password used for this certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subject: The Subject of this Certificate.
        :param pulumi.Input[str] thumbprint: The Thumbprint of this Certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["data"] = data
        __props__["expiration"] = expiration
        __props__["name"] = name
        __props__["password"] = password
        __props__["resource_group_name"] = resource_group_name
        __props__["subject"] = subject
        __props__["thumbprint"] = thumbprint
        return Certificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

