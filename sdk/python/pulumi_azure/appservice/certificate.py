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
    expiration_date: pulumi.Output[str] = pulumi.property("expirationDate")
    """
    The expiration date for the certificate.
    """

    friendly_name: pulumi.Output[str] = pulumi.property("friendlyName")
    """
    The friendly name of the certificate.
    """

    host_names: pulumi.Output[List[str]] = pulumi.property("hostNames")
    """
    List of host names the certificate applies to.
    """

    hosting_environment_profile_id: pulumi.Output[Optional[str]] = pulumi.property("hostingEnvironmentProfileId")
    """
    Must be specified when the certificate is for an App Service Environment hosted App Service. Changing this forces a new resource to be created.
    """

    issue_date: pulumi.Output[str] = pulumi.property("issueDate")
    """
    The issue date for the certificate.
    """

    issuer: pulumi.Output[str] = pulumi.property("issuer")
    """
    The name of the certificate issuer.
    """

    key_vault_secret_id: pulumi.Output[Optional[str]] = pulumi.property("keyVaultSecretId")
    """
    The ID of the Key Vault secret. Changing this forces a new resource to be created.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Specifies the name of the certificate. Changing this forces a new resource to be created.
    """

    password: pulumi.Output[Optional[str]] = pulumi.property("password")
    """
    The password to access the certificate's private key. Changing this forces a new resource to be created.
    """

    pfx_blob: pulumi.Output[Optional[str]] = pulumi.property("pfxBlob")
    """
    The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
    """

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
    """

    subject_name: pulumi.Output[str] = pulumi.property("subjectName")
    """
    The subject name of the certificate.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")

    thumbprint: pulumi.Output[str] = pulumi.property("thumbprint")
    """
    The thumbprint for the certificate.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosting_environment_profile_id: Optional[pulumi.Input[str]] = None,
                 key_vault_secret_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pfx_blob: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an App Service certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hosting_environment_profile_id: Must be specified when the certificate is for an App Service Environment hosted App Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] key_vault_secret_id: The ID of the Key Vault secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] password: The password to access the certificate's private key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pfx_blob: The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
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

            __props__['hosting_environment_profile_id'] = hosting_environment_profile_id
            __props__['key_vault_secret_id'] = key_vault_secret_id
            __props__['location'] = location
            __props__['name'] = name
            __props__['password'] = password
            __props__['pfx_blob'] = pfx_blob
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['expiration_date'] = None
            __props__['friendly_name'] = None
            __props__['host_names'] = None
            __props__['issue_date'] = None
            __props__['issuer'] = None
            __props__['subject_name'] = None
            __props__['thumbprint'] = None
        super(Certificate, __self__).__init__(
            'azure:appservice/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            host_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            hosting_environment_profile_id: Optional[pulumi.Input[str]] = None,
            issue_date: Optional[pulumi.Input[str]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            key_vault_secret_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            pfx_blob: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subject_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_date: The expiration date for the certificate.
        :param pulumi.Input[str] friendly_name: The friendly name of the certificate.
        :param pulumi.Input[List[pulumi.Input[str]]] host_names: List of host names the certificate applies to.
        :param pulumi.Input[str] hosting_environment_profile_id: Must be specified when the certificate is for an App Service Environment hosted App Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] issue_date: The issue date for the certificate.
        :param pulumi.Input[str] issuer: The name of the certificate issuer.
        :param pulumi.Input[str] key_vault_secret_id: The ID of the Key Vault secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] password: The password to access the certificate's private key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pfx_blob: The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subject_name: The subject name of the certificate.
        :param pulumi.Input[str] thumbprint: The thumbprint for the certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["expiration_date"] = expiration_date
        __props__["friendly_name"] = friendly_name
        __props__["host_names"] = host_names
        __props__["hosting_environment_profile_id"] = hosting_environment_profile_id
        __props__["issue_date"] = issue_date
        __props__["issuer"] = issuer
        __props__["key_vault_secret_id"] = key_vault_secret_id
        __props__["location"] = location
        __props__["name"] = name
        __props__["password"] = password
        __props__["pfx_blob"] = pfx_blob
        __props__["resource_group_name"] = resource_group_name
        __props__["subject_name"] = subject_name
        __props__["tags"] = tags
        __props__["thumbprint"] = thumbprint
        return Certificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

