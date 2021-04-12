# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['CustomDomainArgs', 'CustomDomain']

@pulumi.input_type
class CustomDomainArgs:
    def __init__(__self__, *,
                 api_management_id: pulumi.Input[str],
                 developer_portals: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainDeveloperPortalArgs']]]] = None,
                 managements: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainManagementArgs']]]] = None,
                 portals: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainPortalArgs']]]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainProxyArgs']]]] = None,
                 scms: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainScmArgs']]]] = None):
        """
        The set of arguments for constructing a CustomDomain resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDomainDeveloperPortalArgs']]] developer_portals: One or more `developer_portal` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDomainManagementArgs']]] managements: One or more `management` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDomainPortalArgs']]] portals: One or more `portal` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDomainProxyArgs']]] proxies: One or more `proxy` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDomainScmArgs']]] scms: One or more `scm` blocks as defined below.
        """
        pulumi.set(__self__, "api_management_id", api_management_id)
        if developer_portals is not None:
            pulumi.set(__self__, "developer_portals", developer_portals)
        if managements is not None:
            pulumi.set(__self__, "managements", managements)
        if portals is not None:
            pulumi.set(__self__, "portals", portals)
        if proxies is not None:
            pulumi.set(__self__, "proxies", proxies)
        if scms is not None:
            pulumi.set(__self__, "scms", scms)

    @property
    @pulumi.getter(name="apiManagementId")
    def api_management_id(self) -> pulumi.Input[str]:
        """
        The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
        """
        return pulumi.get(self, "api_management_id")

    @api_management_id.setter
    def api_management_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_id", value)

    @property
    @pulumi.getter(name="developerPortals")
    def developer_portals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainDeveloperPortalArgs']]]]:
        """
        One or more `developer_portal` blocks as defined below.
        """
        return pulumi.get(self, "developer_portals")

    @developer_portals.setter
    def developer_portals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainDeveloperPortalArgs']]]]):
        pulumi.set(self, "developer_portals", value)

    @property
    @pulumi.getter
    def managements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainManagementArgs']]]]:
        """
        One or more `management` blocks as defined below.
        """
        return pulumi.get(self, "managements")

    @managements.setter
    def managements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainManagementArgs']]]]):
        pulumi.set(self, "managements", value)

    @property
    @pulumi.getter
    def portals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainPortalArgs']]]]:
        """
        One or more `portal` blocks as defined below.
        """
        return pulumi.get(self, "portals")

    @portals.setter
    def portals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainPortalArgs']]]]):
        pulumi.set(self, "portals", value)

    @property
    @pulumi.getter
    def proxies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainProxyArgs']]]]:
        """
        One or more `proxy` blocks as defined below.
        """
        return pulumi.get(self, "proxies")

    @proxies.setter
    def proxies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainProxyArgs']]]]):
        pulumi.set(self, "proxies", value)

    @property
    @pulumi.getter
    def scms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainScmArgs']]]]:
        """
        One or more `scm` blocks as defined below.
        """
        return pulumi.get(self, "scms")

    @scms.setter
    def scms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDomainScmArgs']]]]):
        pulumi.set(self, "scms", value)


class CustomDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_id: Optional[pulumi.Input[str]] = None,
                 developer_portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainDeveloperPortalArgs']]]]] = None,
                 managements: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainManagementArgs']]]]] = None,
                 portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainPortalArgs']]]]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainProxyArgs']]]]] = None,
                 scms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainScmArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a API Management Custom Domain.

        ## Disclaimers

        > **Note:** It's possible to define Custom Domains both within the `apimanagement.Service` resource via the `hostname_configurations` block and by using this resource. However it's not possible to use both methods to manage Custom Domains within an API Management Service, since there will be conflicts.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=azurerm_resource_group["test"]["location"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            publisher_name="pub1",
            publisher_email="pub1@email.com",
            sku_name="Developer_1")
        example_certificate = azure.keyvault.Certificate("exampleCertificate",
            key_vault_id=azurerm_key_vault["test"]["id"],
            certificate_policy=azure.keyvault.CertificateCertificatePolicyArgs(
                issuer_parameters=azure.keyvault.CertificateCertificatePolicyIssuerParametersArgs(
                    name="Self",
                ),
                key_properties={
                    "exportable": True,
                    "key_size": 2048,
                    "key_type": "RSA",
                    "reuseKey": True,
                },
                lifetime_actions=[azure.keyvault.CertificateCertificatePolicyLifetimeActionArgs(
                    action=azure.keyvault.CertificateCertificatePolicyLifetimeActionActionArgs(
                        action_type="AutoRenew",
                    ),
                    trigger=azure.keyvault.CertificateCertificatePolicyLifetimeActionTriggerArgs(
                        days_before_expiry=30,
                    ),
                )],
                secret_properties=azure.keyvault.CertificateCertificatePolicySecretPropertiesArgs(
                    content_type="application/x-pkcs12",
                ),
                x509_certificate_properties=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesArgs(
                    key_usages=[
                        "cRLSign",
                        "dataEncipherment",
                        "digitalSignature",
                        "keyAgreement",
                        "keyCertSign",
                        "keyEncipherment",
                    ],
                    subject="CN=api.example.com",
                    validity_in_months=12,
                    subject_alternative_names=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs(
                        dns_names=[
                            "api.example.com",
                            "portal.example.com",
                        ],
                    ),
                ),
            ))
        example_custom_domain = azure.apimanagement.CustomDomain("exampleCustomDomain",
            api_management_id=example_service.id,
            proxies=[azure.apimanagement.CustomDomainProxyArgs(
                host_name="api.example.com",
                key_vault_id=azurerm_key_vault_certificate["test"]["secret_id"],
            )],
            developer_portals=[azure.apimanagement.CustomDomainDeveloperPortalArgs(
                host_name="portal.example.com",
                key_vault_id=azurerm_key_vault_certificate["test"]["secret_id"],
            )])
        ```

        ## Import

        API Management Custom Domains can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/customDomain:CustomDomain example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainDeveloperPortalArgs']]]] developer_portals: One or more `developer_portal` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainManagementArgs']]]] managements: One or more `management` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainPortalArgs']]]] portals: One or more `portal` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainProxyArgs']]]] proxies: One or more `proxy` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainScmArgs']]]] scms: One or more `scm` blocks as defined below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a API Management Custom Domain.

        ## Disclaimers

        > **Note:** It's possible to define Custom Domains both within the `apimanagement.Service` resource via the `hostname_configurations` block and by using this resource. However it's not possible to use both methods to manage Custom Domains within an API Management Service, since there will be conflicts.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=azurerm_resource_group["test"]["location"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            publisher_name="pub1",
            publisher_email="pub1@email.com",
            sku_name="Developer_1")
        example_certificate = azure.keyvault.Certificate("exampleCertificate",
            key_vault_id=azurerm_key_vault["test"]["id"],
            certificate_policy=azure.keyvault.CertificateCertificatePolicyArgs(
                issuer_parameters=azure.keyvault.CertificateCertificatePolicyIssuerParametersArgs(
                    name="Self",
                ),
                key_properties={
                    "exportable": True,
                    "key_size": 2048,
                    "key_type": "RSA",
                    "reuseKey": True,
                },
                lifetime_actions=[azure.keyvault.CertificateCertificatePolicyLifetimeActionArgs(
                    action=azure.keyvault.CertificateCertificatePolicyLifetimeActionActionArgs(
                        action_type="AutoRenew",
                    ),
                    trigger=azure.keyvault.CertificateCertificatePolicyLifetimeActionTriggerArgs(
                        days_before_expiry=30,
                    ),
                )],
                secret_properties=azure.keyvault.CertificateCertificatePolicySecretPropertiesArgs(
                    content_type="application/x-pkcs12",
                ),
                x509_certificate_properties=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesArgs(
                    key_usages=[
                        "cRLSign",
                        "dataEncipherment",
                        "digitalSignature",
                        "keyAgreement",
                        "keyCertSign",
                        "keyEncipherment",
                    ],
                    subject="CN=api.example.com",
                    validity_in_months=12,
                    subject_alternative_names=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs(
                        dns_names=[
                            "api.example.com",
                            "portal.example.com",
                        ],
                    ),
                ),
            ))
        example_custom_domain = azure.apimanagement.CustomDomain("exampleCustomDomain",
            api_management_id=example_service.id,
            proxies=[azure.apimanagement.CustomDomainProxyArgs(
                host_name="api.example.com",
                key_vault_id=azurerm_key_vault_certificate["test"]["secret_id"],
            )],
            developer_portals=[azure.apimanagement.CustomDomainDeveloperPortalArgs(
                host_name="portal.example.com",
                key_vault_id=azurerm_key_vault_certificate["test"]["secret_id"],
            )])
        ```

        ## Import

        API Management Custom Domains can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/customDomain:CustomDomain example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
        ```

        :param str resource_name: The name of the resource.
        :param CustomDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_id: Optional[pulumi.Input[str]] = None,
                 developer_portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainDeveloperPortalArgs']]]]] = None,
                 managements: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainManagementArgs']]]]] = None,
                 portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainPortalArgs']]]]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainProxyArgs']]]]] = None,
                 scms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainScmArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if api_management_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_id'")
            __props__['api_management_id'] = api_management_id
            __props__['developer_portals'] = developer_portals
            __props__['managements'] = managements
            __props__['portals'] = portals
            __props__['proxies'] = proxies
            __props__['scms'] = scms
        super(CustomDomain, __self__).__init__(
            'azure:apimanagement/customDomain:CustomDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_id: Optional[pulumi.Input[str]] = None,
            developer_portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainDeveloperPortalArgs']]]]] = None,
            managements: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainManagementArgs']]]]] = None,
            portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainPortalArgs']]]]] = None,
            proxies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainProxyArgs']]]]] = None,
            scms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainScmArgs']]]]] = None) -> 'CustomDomain':
        """
        Get an existing CustomDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainDeveloperPortalArgs']]]] developer_portals: One or more `developer_portal` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainManagementArgs']]]] managements: One or more `management` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainPortalArgs']]]] portals: One or more `portal` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainProxyArgs']]]] proxies: One or more `proxy` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomDomainScmArgs']]]] scms: One or more `scm` blocks as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_id"] = api_management_id
        __props__["developer_portals"] = developer_portals
        __props__["managements"] = managements
        __props__["portals"] = portals
        __props__["proxies"] = proxies
        __props__["scms"] = scms
        return CustomDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementId")
    def api_management_id(self) -> pulumi.Output[str]:
        """
        The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
        """
        return pulumi.get(self, "api_management_id")

    @property
    @pulumi.getter(name="developerPortals")
    def developer_portals(self) -> pulumi.Output[Optional[Sequence['outputs.CustomDomainDeveloperPortal']]]:
        """
        One or more `developer_portal` blocks as defined below.
        """
        return pulumi.get(self, "developer_portals")

    @property
    @pulumi.getter
    def managements(self) -> pulumi.Output[Optional[Sequence['outputs.CustomDomainManagement']]]:
        """
        One or more `management` blocks as defined below.
        """
        return pulumi.get(self, "managements")

    @property
    @pulumi.getter
    def portals(self) -> pulumi.Output[Optional[Sequence['outputs.CustomDomainPortal']]]:
        """
        One or more `portal` blocks as defined below.
        """
        return pulumi.get(self, "portals")

    @property
    @pulumi.getter
    def proxies(self) -> pulumi.Output[Optional[Sequence['outputs.CustomDomainProxy']]]:
        """
        One or more `proxy` blocks as defined below.
        """
        return pulumi.get(self, "proxies")

    @property
    @pulumi.getter
    def scms(self) -> pulumi.Output[Optional[Sequence['outputs.CustomDomainScm']]]:
        """
        One or more `scm` blocks as defined below.
        """
        return pulumi.get(self, "scms")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

