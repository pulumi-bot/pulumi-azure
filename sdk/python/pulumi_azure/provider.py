# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from ._inputs import *

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auxiliary_tenant_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_correlation_request_id: Optional[pulumi.Input[bool]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 features: Optional[pulumi.Input[pulumi.InputType['ProviderFeaturesArgs']]] = None,
                 metadata_host: Optional[pulumi.Input[str]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_provider_registration: Optional[pulumi.Input[bool]] = None,
                 storage_use_azuread: Optional[pulumi.Input[bool]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the azurerm package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_certificate_password: The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
               Certificate
        :param pulumi.Input[str] client_certificate_path: The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
               Principal using a Client Certificate.
        :param pulumi.Input[str] client_id: The Client ID which should be used.
        :param pulumi.Input[str] client_secret: The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        :param pulumi.Input[bool] disable_correlation_request_id: This will disable the x-ms-correlation-request-id header.
        :param pulumi.Input[bool] disable_terraform_partner_id: This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        :param pulumi.Input[str] environment: The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
               public.
        :param pulumi.Input[str] metadata_host: The Hostname which should be used for the Azure Metadata Service.
        :param pulumi.Input[str] metadata_url: Deprecated - replaced by `metadata_host`.
        :param pulumi.Input[str] msi_endpoint: The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
               automatically.
        :param pulumi.Input[str] partner_id: A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        :param pulumi.Input[bool] skip_credentials_validation: This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
        :param pulumi.Input[bool] skip_provider_registration: Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
               registered?
        :param pulumi.Input[bool] storage_use_azuread: Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
        :param pulumi.Input[str] subscription_id: The Subscription ID which should be used.
        :param pulumi.Input[str] tenant_id: The Tenant ID which should be used.
        :param pulumi.Input[bool] use_msi: Allowed Managed Service Identity be used for Authentication.
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

            __props__['auxiliary_tenant_ids'] = pulumi.Output.from_input(auxiliary_tenant_ids).apply(pulumi.runtime.to_json) if auxiliary_tenant_ids is not None else None
            __props__['client_certificate_password'] = client_certificate_password
            __props__['client_certificate_path'] = client_certificate_path
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['disable_correlation_request_id'] = pulumi.Output.from_input(disable_correlation_request_id).apply(pulumi.runtime.to_json) if disable_correlation_request_id is not None else None
            __props__['disable_terraform_partner_id'] = pulumi.Output.from_input(disable_terraform_partner_id).apply(pulumi.runtime.to_json) if disable_terraform_partner_id is not None else None
            if environment is None:
                environment = (_utilities.get_env('AZURE_ENVIRONMENT', 'ARM_ENVIRONMENT') or 'public')
            __props__['environment'] = environment
            __props__['features'] = pulumi.Output.from_input(features).apply(pulumi.runtime.to_json) if features is not None else None
            if metadata_host is None:
                metadata_host = (_utilities.get_env('ARM_METADATA_HOSTNAME') or '')
            __props__['metadata_host'] = metadata_host
            if metadata_url is None:
                metadata_url = (_utilities.get_env('ARM_METADATA_URL') or '')
            if metadata_url is not None and not opts.urn:
                warnings.warn("""use `metadata_host` instead""", DeprecationWarning)
                pulumi.log.warn("metadata_url is deprecated: use `metadata_host` instead")
            __props__['metadata_url'] = metadata_url
            __props__['msi_endpoint'] = msi_endpoint
            __props__['partner_id'] = partner_id
            __props__['skip_credentials_validation'] = pulumi.Output.from_input(skip_credentials_validation).apply(pulumi.runtime.to_json) if skip_credentials_validation is not None else None
            if skip_provider_registration is None:
                skip_provider_registration = (_utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)
            __props__['skip_provider_registration'] = pulumi.Output.from_input(skip_provider_registration).apply(pulumi.runtime.to_json) if skip_provider_registration is not None else None
            if storage_use_azuread is None:
                storage_use_azuread = (_utilities.get_env_bool('ARM_STORAGE_USE_AZUREAD') or False)
            __props__['storage_use_azuread'] = pulumi.Output.from_input(storage_use_azuread).apply(pulumi.runtime.to_json) if storage_use_azuread is not None else None
            if subscription_id is None:
                subscription_id = (_utilities.get_env('ARM_SUBSCRIPTION_ID') or '')
            __props__['subscription_id'] = subscription_id
            __props__['tenant_id'] = tenant_id
            __props__['use_msi'] = pulumi.Output.from_input(use_msi).apply(pulumi.runtime.to_json) if use_msi is not None else None
        super(Provider, __self__).__init__(
            'azure',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="auxiliaryTenantIds")
    def auxiliary_tenant_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "auxiliary_tenant_ids")

    @property
    @pulumi.getter(name="clientCertificatePassword")
    def client_certificate_password(self) -> pulumi.Output[Optional[str]]:
        """
        The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return pulumi.get(self, "client_certificate_password")

    @property
    @pulumi.getter(name="clientCertificatePath")
    def client_certificate_path(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return pulumi.get(self, "client_certificate_path")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Client ID which should be used.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="disableCorrelationRequestId")
    def disable_correlation_request_id(self) -> pulumi.Output[Optional[bool]]:
        """
        This will disable the x-ms-correlation-request-id header.
        """
        return pulumi.get(self, "disable_correlation_request_id")

    @property
    @pulumi.getter(name="disableTerraformPartnerId")
    def disable_terraform_partner_id(self) -> pulumi.Output[Optional[bool]]:
        """
        This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        """
        return pulumi.get(self, "disable_terraform_partner_id")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[str]]:
        """
        The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
        public.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def features(self) -> pulumi.Output[Optional['outputs.ProviderFeatures']]:
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="metadataHost")
    def metadata_host(self) -> pulumi.Output[Optional[str]]:
        """
        The Hostname which should be used for the Azure Metadata Service.
        """
        return pulumi.get(self, "metadata_host")

    @property
    @pulumi.getter(name="metadataUrl")
    def metadata_url(self) -> pulumi.Output[Optional[str]]:
        """
        Deprecated - replaced by `metadata_host`.
        """
        return pulumi.get(self, "metadata_url")

    @property
    @pulumi.getter(name="msiEndpoint")
    def msi_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
        automatically.
        """
        return pulumi.get(self, "msi_endpoint")

    @property
    @pulumi.getter(name="partnerId")
    def partner_id(self) -> pulumi.Output[Optional[str]]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return pulumi.get(self, "partner_id")

    @property
    @pulumi.getter(name="skipCredentialsValidation")
    def skip_credentials_validation(self) -> pulumi.Output[Optional[bool]]:
        """
        This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
        """
        return pulumi.get(self, "skip_credentials_validation")

    @property
    @pulumi.getter(name="skipProviderRegistration")
    def skip_provider_registration(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
        registered?
        """
        return pulumi.get(self, "skip_provider_registration")

    @property
    @pulumi.getter(name="storageUseAzuread")
    def storage_use_azuread(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
        """
        return pulumi.get(self, "storage_use_azuread")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Subscription ID which should be used.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Tenant ID which should be used.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="useMsi")
    def use_msi(self) -> pulumi.Output[Optional[bool]]:
        """
        Allowed Managed Service Identity be used for Authentication.
        """
        return pulumi.get(self, "use_msi")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

