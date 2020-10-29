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

__all__ = ['Certificate']


class Certificate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[pulumi.InputType['CertificateCertificateArgs']]] = None,
                 certificate_policy: Optional[pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Key Vault Certificate.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CertificateCertificateArgs']] certificate: A `certificate` block as defined below, used to Import an existing certificate.
        :param pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']] certificate_policy: A `certificate_policy` block as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Certificate should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
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

            __props__['certificate'] = certificate
            if certificate_policy is None:
                raise TypeError("Missing required property 'certificate_policy'")
            __props__['certificate_policy'] = certificate_policy
            if key_vault_id is None:
                raise TypeError("Missing required property 'key_vault_id'")
            __props__['key_vault_id'] = key_vault_id
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['certificate_attributes'] = None
            __props__['certificate_data'] = None
            __props__['secret_id'] = None
            __props__['thumbprint'] = None
            __props__['version'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:keyvault/certifiate:Certifiate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Certificate, __self__).__init__(
            'azure:keyvault/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[pulumi.InputType['CertificateCertificateArgs']]] = None,
            certificate_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateCertificateAttributeArgs']]]]] = None,
            certificate_data: Optional[pulumi.Input[str]] = None,
            certificate_policy: Optional[pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CertificateCertificateArgs']] certificate: A `certificate` block as defined below, used to Import an existing certificate.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateCertificateAttributeArgs']]]] certificate_attributes: A `certificate_attribute` block as defined below.
        :param pulumi.Input[str] certificate_data: The raw Key Vault Certificate data represented as a hexadecimal string.
        :param pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']] certificate_policy: A `certificate_policy` block as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Certificate should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secret_id: The ID of the associated Key Vault Secret.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] thumbprint: The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        :param pulumi.Input[str] version: The current version of the Key Vault Certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate"] = certificate
        __props__["certificate_attributes"] = certificate_attributes
        __props__["certificate_data"] = certificate_data
        __props__["certificate_policy"] = certificate_policy
        __props__["key_vault_id"] = key_vault_id
        __props__["name"] = name
        __props__["secret_id"] = secret_id
        __props__["tags"] = tags
        __props__["thumbprint"] = thumbprint
        __props__["version"] = version
        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional['outputs.CertificateCertificate']]:
        """
        A `certificate` block as defined below, used to Import an existing certificate.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateAttributes")
    def certificate_attributes(self) -> pulumi.Output[Sequence['outputs.CertificateCertificateAttribute']]:
        """
        A `certificate_attribute` block as defined below.
        """
        return pulumi.get(self, "certificate_attributes")

    @property
    @pulumi.getter(name="certificateData")
    def certificate_data(self) -> pulumi.Output[str]:
        """
        The raw Key Vault Certificate data represented as a hexadecimal string.
        """
        return pulumi.get(self, "certificate_data")

    @property
    @pulumi.getter(name="certificatePolicy")
    def certificate_policy(self) -> pulumi.Output['outputs.CertificateCertificatePolicy']:
        """
        A `certificate_policy` block as defined below.
        """
        return pulumi.get(self, "certificate_policy")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the Key Vault where the Certificate should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        The ID of the associated Key Vault Secret.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Output[str]:
        """
        The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The current version of the Key Vault Certificate.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

