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

__all__ = ['CertificateIssuer']


class CertificateIssuer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateIssuerAdminArgs']]]]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Key Vault Certificate Issuer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        current = azure.core.get_client_config()
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="standard",
            tenant_id=current.tenant_id)
        example_certificate_issuer = azure.keyvault.CertificateIssuer("exampleCertificateIssuer",
            org_id="ExampleOrgName",
            key_vault_id=example_key_vault.id,
            provider_name="DigiCert",
            account_id="0000",
            password="example-password")
        ```

        ## Import

        Key Vault Certificate Issuers can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:keyvault/certificateIssuer:CertificateIssuer example "https://key-vault-name.vault.azure.net/certificates/issuers/example"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account number with the third-party Certificate Issuer.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateIssuerAdminArgs']]]] admins: One or more `admin` blocks as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault in which to create the Certificate Issuer.
        :param pulumi.Input[str] name: The name which should be used for this Key Vault Certificate Issuer. Changing this forces a new Key Vault Certificate Issuer to be created.
        :param pulumi.Input[str] org_id: The ID of the organization as provided to the issuer.
        :param pulumi.Input[str] password: The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
        :param pulumi.Input[str] provider_name: The name of the third-party Certificate Issuer. Possible values are: `DigiCert`, `GlobalSign`, `OneCertV2-PrivateCA`, `OneCertV2-PublicCA` and `SslAdminV2`.
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

            __props__['account_id'] = account_id
            __props__['admins'] = admins
            if key_vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_vault_id'")
            __props__['key_vault_id'] = key_vault_id
            __props__['name'] = name
            __props__['org_id'] = org_id
            __props__['password'] = password
            if provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'provider_name'")
            __props__['provider_name'] = provider_name
        super(CertificateIssuer, __self__).__init__(
            'azure:keyvault/certificateIssuer:CertificateIssuer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            admins: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateIssuerAdminArgs']]]]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            provider_name: Optional[pulumi.Input[str]] = None) -> 'CertificateIssuer':
        """
        Get an existing CertificateIssuer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account number with the third-party Certificate Issuer.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateIssuerAdminArgs']]]] admins: One or more `admin` blocks as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault in which to create the Certificate Issuer.
        :param pulumi.Input[str] name: The name which should be used for this Key Vault Certificate Issuer. Changing this forces a new Key Vault Certificate Issuer to be created.
        :param pulumi.Input[str] org_id: The ID of the organization as provided to the issuer.
        :param pulumi.Input[str] password: The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
        :param pulumi.Input[str] provider_name: The name of the third-party Certificate Issuer. Possible values are: `DigiCert`, `GlobalSign`, `OneCertV2-PrivateCA`, `OneCertV2-PublicCA` and `SslAdminV2`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["admins"] = admins
        __props__["key_vault_id"] = key_vault_id
        __props__["name"] = name
        __props__["org_id"] = org_id
        __props__["password"] = password
        __props__["provider_name"] = provider_name
        return CertificateIssuer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The account number with the third-party Certificate Issuer.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def admins(self) -> pulumi.Output[Optional[Sequence['outputs.CertificateIssuerAdmin']]]:
        """
        One or more `admin` blocks as defined below.
        """
        return pulumi.get(self, "admins")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the Key Vault in which to create the Certificate Issuer.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Key Vault Certificate Issuer. Changing this forces a new Key Vault Certificate Issuer to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the organization as provided to the issuer.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[str]:
        """
        The name of the third-party Certificate Issuer. Possible values are: `DigiCert`, `GlobalSign`, `OneCertV2-PrivateCA`, `OneCertV2-PublicCA` and `SslAdminV2`.
        """
        return pulumi.get(self, "provider_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

