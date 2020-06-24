# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class CertificateOrder(pulumi.CustomResource):
    app_service_certificate_not_renewable_reasons: pulumi.Output[list]
    """
    Reasons why App Service Certificate is not renewable at the current moment.
    """
    auto_renew: pulumi.Output[bool]
    """
    true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
    """
    certificates: pulumi.Output[list]
    """
    State of the Key Vault secret. A `certificates` block as defined below.

      * `certificateName` (`str`) - The name of the App Service Certificate.
      * `key_vault_id` (`str`) - Key Vault resource Id.
      * `keyVaultSecretName` (`str`) - Key Vault secret name.
      * `provisioningState` (`str`) - Status of the Key Vault secret.
    """
    csr: pulumi.Output[str]
    """
    Last CSR that was created for this order.
    """
    distinguished_name: pulumi.Output[str]
    """
    The Distinguished Name for the App Service Certificate Order.
    """
    domain_verification_token: pulumi.Output[str]
    """
    Domain verification token.
    """
    expiration_time: pulumi.Output[str]
    """
    Certificate expiration time.
    """
    intermediate_thumbprint: pulumi.Output[str]
    """
    Certificate thumbprint intermediate certificate.
    """
    is_private_key_external: pulumi.Output[bool]
    """
    Whether the private key is external or not.
    """
    key_size: pulumi.Output[float]
    """
    Certificate key size.  Defaults to 2048.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the certificate. Changing this forces a new resource to be created.
    """
    product_type: pulumi.Output[str]
    """
    Certificate product type, such as `Standard` or `WildCard`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
    """
    root_thumbprint: pulumi.Output[str]
    """
    Certificate thumbprint for root certificate.
    """
    signed_certificate_thumbprint: pulumi.Output[str]
    """
    Certificate thumbprint for signed certificate.
    """
    status: pulumi.Output[str]
    """
    Current order status.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    validity_in_years: pulumi.Output[float]
    """
    Duration in years (must be between `1` and `3`).  Defaults to `1`.
    """
    def __init__(__self__, resource_name, opts=None, auto_renew=None, csr=None, distinguished_name=None, key_size=None, location=None, name=None, product_type=None, resource_group_name=None, tags=None, validity_in_years=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an App Service Certificate Order.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_certificate_order = azure.appservice.CertificateOrder("exampleCertificateOrder",
            resource_group_name=example_resource_group.name,
            location="global",
            distinguished_name="CN=example.com",
            product_type="Standard")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
        :param pulumi.Input[str] csr: Last CSR that was created for this order.
        :param pulumi.Input[str] distinguished_name: The Distinguished Name for the App Service Certificate Order.
        :param pulumi.Input[float] key_size: Certificate key size.  Defaults to 2048.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
        :param pulumi.Input[str] name: Specifies the name of the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] product_type: Certificate product type, such as `Standard` or `WildCard`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[float] validity_in_years: Duration in years (must be between `1` and `3`).  Defaults to `1`.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_renew'] = auto_renew
            __props__['csr'] = csr
            __props__['distinguished_name'] = distinguished_name
            __props__['key_size'] = key_size
            __props__['location'] = location
            __props__['name'] = name
            __props__['product_type'] = product_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['validity_in_years'] = validity_in_years
            __props__['app_service_certificate_not_renewable_reasons'] = None
            __props__['certificates'] = None
            __props__['domain_verification_token'] = None
            __props__['expiration_time'] = None
            __props__['intermediate_thumbprint'] = None
            __props__['is_private_key_external'] = None
            __props__['root_thumbprint'] = None
            __props__['signed_certificate_thumbprint'] = None
            __props__['status'] = None
        super(CertificateOrder, __self__).__init__(
            'azure:appservice/certificateOrder:CertificateOrder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_service_certificate_not_renewable_reasons=None, auto_renew=None, certificates=None, csr=None, distinguished_name=None, domain_verification_token=None, expiration_time=None, intermediate_thumbprint=None, is_private_key_external=None, key_size=None, location=None, name=None, product_type=None, resource_group_name=None, root_thumbprint=None, signed_certificate_thumbprint=None, status=None, tags=None, validity_in_years=None):
        """
        Get an existing CertificateOrder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_service_certificate_not_renewable_reasons: Reasons why App Service Certificate is not renewable at the current moment.
        :param pulumi.Input[bool] auto_renew: true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
        :param pulumi.Input[list] certificates: State of the Key Vault secret. A `certificates` block as defined below.
        :param pulumi.Input[str] csr: Last CSR that was created for this order.
        :param pulumi.Input[str] distinguished_name: The Distinguished Name for the App Service Certificate Order.
        :param pulumi.Input[str] domain_verification_token: Domain verification token.
        :param pulumi.Input[str] expiration_time: Certificate expiration time.
        :param pulumi.Input[str] intermediate_thumbprint: Certificate thumbprint intermediate certificate.
        :param pulumi.Input[bool] is_private_key_external: Whether the private key is external or not.
        :param pulumi.Input[float] key_size: Certificate key size.  Defaults to 2048.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
        :param pulumi.Input[str] name: Specifies the name of the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] product_type: Certificate product type, such as `Standard` or `WildCard`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] root_thumbprint: Certificate thumbprint for root certificate.
        :param pulumi.Input[str] signed_certificate_thumbprint: Certificate thumbprint for signed certificate.
        :param pulumi.Input[str] status: Current order status.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[float] validity_in_years: Duration in years (must be between `1` and `3`).  Defaults to `1`.

        The **certificates** object supports the following:

          * `certificateName` (`pulumi.Input[str]`) - The name of the App Service Certificate.
          * `key_vault_id` (`pulumi.Input[str]`) - Key Vault resource Id.
          * `keyVaultSecretName` (`pulumi.Input[str]`) - Key Vault secret name.
          * `provisioningState` (`pulumi.Input[str]`) - Status of the Key Vault secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_service_certificate_not_renewable_reasons"] = app_service_certificate_not_renewable_reasons
        __props__["auto_renew"] = auto_renew
        __props__["certificates"] = certificates
        __props__["csr"] = csr
        __props__["distinguished_name"] = distinguished_name
        __props__["domain_verification_token"] = domain_verification_token
        __props__["expiration_time"] = expiration_time
        __props__["intermediate_thumbprint"] = intermediate_thumbprint
        __props__["is_private_key_external"] = is_private_key_external
        __props__["key_size"] = key_size
        __props__["location"] = location
        __props__["name"] = name
        __props__["product_type"] = product_type
        __props__["resource_group_name"] = resource_group_name
        __props__["root_thumbprint"] = root_thumbprint
        __props__["signed_certificate_thumbprint"] = signed_certificate_thumbprint
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["validity_in_years"] = validity_in_years
        return CertificateOrder(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
