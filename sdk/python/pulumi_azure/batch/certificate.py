# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Certificate']


class Certificate(pulumi.CustomResource):
    account_name: pulumi.Output[str] = pulumi.output_property("accountName")
    """
    Specifies the name of the Batch account. Changing this forces a new resource to be created.
    """
    certificate: pulumi.Output[str] = pulumi.output_property("certificate")
    """
    The base64-encoded contents of the certificate.
    """
    format: pulumi.Output[str] = pulumi.output_property("format")
    """
    The format of the certificate. Possible values are `Cer` or `Pfx`.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The generated name of the certificate.
    """
    password: pulumi.Output[Optional[str]] = pulumi.output_property("password")
    """
    The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
    """
    public_data: pulumi.Output[str] = pulumi.output_property("publicData")
    """
    The public key of the certificate.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
    """
    thumbprint: pulumi.Output[str] = pulumi.output_property("thumbprint")
    """
    The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
    """
    thumbprint_algorithm: pulumi.Output[str] = pulumi.output_property("thumbprintAlgorithm")
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, account_name: Optional[pulumi.Input[str]] = None, certificate: Optional[pulumi.Input[str]] = None, format: Optional[pulumi.Input[str]] = None, password: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, thumbprint: Optional[pulumi.Input[str]] = None, thumbprint_algorithm: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a certificate in an Azure Batch account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] certificate: The base64-encoded contents of the certificate.
        :param pulumi.Input[str] format: The format of the certificate. Possible values are `Cer` or `Pfx`.
        :param pulumi.Input[str] password: The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] thumbprint: The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if certificate is None:
                raise TypeError("Missing required property 'certificate'")
            __props__['certificate'] = certificate
            if format is None:
                raise TypeError("Missing required property 'format'")
            __props__['format'] = format
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if thumbprint is None:
                raise TypeError("Missing required property 'thumbprint'")
            __props__['thumbprint'] = thumbprint
            if thumbprint_algorithm is None:
                raise TypeError("Missing required property 'thumbprint_algorithm'")
            __props__['thumbprint_algorithm'] = thumbprint_algorithm
            __props__['name'] = None
            __props__['public_data'] = None
        super(Certificate, __self__).__init__(
            'azure:batch/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, account_name: Optional[pulumi.Input[str]] = None, certificate: Optional[pulumi.Input[str]] = None, format: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, password: Optional[pulumi.Input[str]] = None, public_data: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, thumbprint: Optional[pulumi.Input[str]] = None, thumbprint_algorithm: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] certificate: The base64-encoded contents of the certificate.
        :param pulumi.Input[str] format: The format of the certificate. Possible values are `Cer` or `Pfx`.
        :param pulumi.Input[str] name: The generated name of the certificate.
        :param pulumi.Input[str] password: The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
        :param pulumi.Input[str] public_data: The public key of the certificate.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] thumbprint: The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["certificate"] = certificate
        __props__["format"] = format
        __props__["name"] = name
        __props__["password"] = password
        __props__["public_data"] = public_data
        __props__["resource_group_name"] = resource_group_name
        __props__["thumbprint"] = thumbprint
        __props__["thumbprint_algorithm"] = thumbprint_algorithm
        return Certificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

