# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetCertificateOrderResult:
    """
    A collection of values returned by getCertificateOrder.
    """
    def __init__(__self__, app_service_certificate_not_renewable_reasons=None, auto_renew=None, certificates=None, csr=None, distinguished_name=None, domain_verification_token=None, expiration_time=None, id=None, intermediate_thumbprint=None, is_private_key_external=None, key_size=None, location=None, name=None, product_type=None, resource_group_name=None, root_thumbprint=None, signed_certificate_thumbprint=None, status=None, tags=None, validity_in_years=None):
        if app_service_certificate_not_renewable_reasons and not isinstance(app_service_certificate_not_renewable_reasons, list):
            raise TypeError("Expected argument 'app_service_certificate_not_renewable_reasons' to be a list")
        __self__.app_service_certificate_not_renewable_reasons = app_service_certificate_not_renewable_reasons
        """
        Reasons why App Service Certificate is not renewable at the current moment.
        """
        if auto_renew and not isinstance(auto_renew, bool):
            raise TypeError("Expected argument 'auto_renew' to be a bool")
        __self__.auto_renew = auto_renew
        """
        true if the certificate should be automatically renewed when it expires; otherwise, false.
        """
        if certificates and not isinstance(certificates, list):
            raise TypeError("Expected argument 'certificates' to be a list")
        __self__.certificates = certificates
        """
        State of the Key Vault secret. A `certificates` block as defined below.
        """
        if csr and not isinstance(csr, str):
            raise TypeError("Expected argument 'csr' to be a str")
        __self__.csr = csr
        """
        Last CSR that was created for this order.
        """
        if distinguished_name and not isinstance(distinguished_name, str):
            raise TypeError("Expected argument 'distinguished_name' to be a str")
        __self__.distinguished_name = distinguished_name
        """
        The Distinguished Name for the App Service Certificate Order.
        """
        if domain_verification_token and not isinstance(domain_verification_token, str):
            raise TypeError("Expected argument 'domain_verification_token' to be a str")
        __self__.domain_verification_token = domain_verification_token
        """
        Domain verification token.
        """
        if expiration_time and not isinstance(expiration_time, str):
            raise TypeError("Expected argument 'expiration_time' to be a str")
        __self__.expiration_time = expiration_time
        """
        Certificate expiration time.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if intermediate_thumbprint and not isinstance(intermediate_thumbprint, str):
            raise TypeError("Expected argument 'intermediate_thumbprint' to be a str")
        __self__.intermediate_thumbprint = intermediate_thumbprint
        """
        Certificate thumbprint intermediate certificate.
        """
        if is_private_key_external and not isinstance(is_private_key_external, bool):
            raise TypeError("Expected argument 'is_private_key_external' to be a bool")
        __self__.is_private_key_external = is_private_key_external
        """
        Whether the private key is external or not.
        """
        if key_size and not isinstance(key_size, float):
            raise TypeError("Expected argument 'key_size' to be a float")
        __self__.key_size = key_size
        """
        Certificate key size.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure location where the App Service exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if product_type and not isinstance(product_type, str):
            raise TypeError("Expected argument 'product_type' to be a str")
        __self__.product_type = product_type
        """
        Certificate product type, such as `Standard` or `WildCard`.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if root_thumbprint and not isinstance(root_thumbprint, str):
            raise TypeError("Expected argument 'root_thumbprint' to be a str")
        __self__.root_thumbprint = root_thumbprint
        """
        Certificate thumbprint for root certificate.
        """
        if signed_certificate_thumbprint and not isinstance(signed_certificate_thumbprint, str):
            raise TypeError("Expected argument 'signed_certificate_thumbprint' to be a str")
        __self__.signed_certificate_thumbprint = signed_certificate_thumbprint
        """
        Certificate thumbprint for signed certificate.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        Current order status.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        if validity_in_years and not isinstance(validity_in_years, float):
            raise TypeError("Expected argument 'validity_in_years' to be a float")
        __self__.validity_in_years = validity_in_years
        """
        Duration in years (must be between 1 and 3).
        """
class AwaitableGetCertificateOrderResult(GetCertificateOrderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateOrderResult(
            app_service_certificate_not_renewable_reasons=self.app_service_certificate_not_renewable_reasons,
            auto_renew=self.auto_renew,
            certificates=self.certificates,
            csr=self.csr,
            distinguished_name=self.distinguished_name,
            domain_verification_token=self.domain_verification_token,
            expiration_time=self.expiration_time,
            id=self.id,
            intermediate_thumbprint=self.intermediate_thumbprint,
            is_private_key_external=self.is_private_key_external,
            key_size=self.key_size,
            location=self.location,
            name=self.name,
            product_type=self.product_type,
            resource_group_name=self.resource_group_name,
            root_thumbprint=self.root_thumbprint,
            signed_certificate_thumbprint=self.signed_certificate_thumbprint,
            status=self.status,
            tags=self.tags,
            validity_in_years=self.validity_in_years)

def get_certificate_order(name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing App Service Certificate Order.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.appservice.get_certificate_order(name="example-cert-order",
        resource_group_name="example-resources")
    pulumi.export("certificateOrderId", example.id)
    ```


    :param str name: The name of the App Service.
    :param str resource_group_name: The Name of the Resource Group where the App Service exists.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:appservice/getCertificateOrder:getCertificateOrder', __args__, opts=opts).value

    return AwaitableGetCertificateOrderResult(
        app_service_certificate_not_renewable_reasons=__ret__.get('appServiceCertificateNotRenewableReasons'),
        auto_renew=__ret__.get('autoRenew'),
        certificates=__ret__.get('certificates'),
        csr=__ret__.get('csr'),
        distinguished_name=__ret__.get('distinguishedName'),
        domain_verification_token=__ret__.get('domainVerificationToken'),
        expiration_time=__ret__.get('expirationTime'),
        id=__ret__.get('id'),
        intermediate_thumbprint=__ret__.get('intermediateThumbprint'),
        is_private_key_external=__ret__.get('isPrivateKeyExternal'),
        key_size=__ret__.get('keySize'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        product_type=__ret__.get('productType'),
        resource_group_name=__ret__.get('resourceGroupName'),
        root_thumbprint=__ret__.get('rootThumbprint'),
        signed_certificate_thumbprint=__ret__.get('signedCertificateThumbprint'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        validity_in_years=__ret__.get('validityInYears'))
