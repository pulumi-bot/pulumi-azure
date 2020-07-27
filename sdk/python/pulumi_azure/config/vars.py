# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'auxiliary_tenant_ids',
    'client_certificate_password',
    'client_certificate_path',
    'client_id',
    'client_secret',
    'disable_correlation_request_id',
    'disable_terraform_partner_id',
    'environment',
    'features',
    'location',
    'metadata_host',
    'metadata_url',
    'msi_endpoint',
    'partner_id',
    'skip_credentials_validation',
    'skip_provider_registration',
    'storage_use_azuread',
    'subscription_id',
    'tenant_id',
    'use_msi',
]

__config__ = pulumi.Config('azure')

auxiliary_tenant_ids = __config__.get('auxiliaryTenantIds')

client_certificate_password = __config__.get('clientCertificatePassword')
"""
The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
Certificate
"""

client_certificate_path = __config__.get('clientCertificatePath')
"""
The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
Principal using a Client Certificate.
"""

client_id = __config__.get('clientId')
"""
The Client ID which should be used.
"""

client_secret = __config__.get('clientSecret')
"""
The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
"""

disable_correlation_request_id = __config__.get('disableCorrelationRequestId')
"""
This will disable the x-ms-correlation-request-id header.
"""

disable_terraform_partner_id = __config__.get('disableTerraformPartnerId')
"""
This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
"""

environment = __config__.get('environment') or (_utilities.get_env('AZURE_ENVIRONMENT', 'ARM_ENVIRONMENT') or 'public')
"""
The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
public.
"""

features = __config__.get('features')

location = __config__.get('location') or _utilities.get_env('ARM_LOCATION')

metadata_host = __config__.get('metadataHost')
"""
The Hostname which should be used for the Azure Metadata Service.
"""

metadata_url = __config__.get('metadataUrl') or (_utilities.get_env('ARM_METADATA_URL') or '')
"""
Deprecated - replaced by `metadata_host`.
"""

msi_endpoint = __config__.get('msiEndpoint')
"""
The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
automatically.
"""

partner_id = __config__.get('partnerId')
"""
A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
"""

skip_credentials_validation = __config__.get('skipCredentialsValidation')
"""
This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
"""

skip_provider_registration = __config__.get('skipProviderRegistration') or (_utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)
"""
Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
registered?
"""

storage_use_azuread = __config__.get('storageUseAzuread') or (_utilities.get_env_bool('ARM_STORAGE_USE_AZUREAD') or False)
"""
Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
"""

subscription_id = __config__.get('subscriptionId') or (_utilities.get_env('ARM_SUBSCRIPTION_ID') or '')
"""
The Subscription ID which should be used.
"""

tenant_id = __config__.get('tenantId')
"""
The Tenant ID which should be used.
"""

use_msi = __config__.get('useMsi')
"""
Allowed Managed Service Identity be used for Authentication.
"""

