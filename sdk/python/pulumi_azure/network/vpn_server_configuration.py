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

__all__ = ['VpnServerConfiguration']


class VpnServerConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_active_directory_authentications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationAzureActiveDirectoryAuthenticationArgs']]]]] = None,
                 client_revoked_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRevokedCertificateArgs']]]]] = None,
                 client_root_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRootCertificateArgs']]]]] = None,
                 ipsec_policy: Optional[pulumi.Input[pulumi.InputType['VpnServerConfigurationIpsecPolicyArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 radius_server: Optional[pulumi.Input[pulumi.InputType['VpnServerConfigurationRadiusServerArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_authentication_types: Optional[pulumi.Input[str]] = None,
                 vpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a VPN Server Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West Europe")
        test = azure.network.VpnServerConfiguration("test",
            resource_group_name=example.name,
            location=example.location,
            vpn_authentication_types=["Certificate"],
            client_root_certificates=[azure.network.VpnServerConfigurationClientRootCertificateArgs(
                name="DigiCert-Federated-ID-Root-CA",
                public_cert_data=\"\"\"MIIDuzCCAqOgAwIBAgIQCHTZWCM+IlfFIRXIvyKSrjANBgkqhkiG9w0BAQsFADBn
        MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
        d3cuZGlnaWNlcnQuY29tMSYwJAYDVQQDEx1EaWdpQ2VydCBGZWRlcmF0ZWQgSUQg
        Um9vdCBDQTAeFw0xMzAxMTUxMjAwMDBaFw0zMzAxMTUxMjAwMDBaMGcxCzAJBgNV
        BAYTAlVTMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdp
        Y2VydC5jb20xJjAkBgNVBAMTHURpZ2lDZXJ0IEZlZGVyYXRlZCBJRCBSb290IENB
        MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvAEB4pcCqnNNOWE6Ur5j
        QPUH+1y1F9KdHTRSza6k5iDlXq1kGS1qAkuKtw9JsiNRrjltmFnzMZRBbX8Tlfl8
        zAhBmb6dDduDGED01kBsTkgywYPxXVTKec0WxYEEF0oMn4wSYNl0lt2eJAKHXjNf
        GTwiibdP8CUR2ghSM2sUTI8Nt1Omfc4SMHhGhYD64uJMbX98THQ/4LMGuYegou+d
        GTiahfHtjn7AboSEknwAMJHCh5RlYZZ6B1O4QbKJ+34Q0eKgnI3X6Vc9u0zf6DH8
        Dk+4zQDYRRTqTnVO3VT8jzqDlCRuNtq6YvryOWN74/dq8LQhUnXHvFyrsdMaE1X2
        DwIDAQABo2MwYTAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNV
        HQ4EFgQUGRdkFnbGt1EWjKwbUne+5OaZvRYwHwYDVR0jBBgwFoAUGRdkFnbGt1EW
        jKwbUne+5OaZvRYwDQYJKoZIhvcNAQELBQADggEBAHcqsHkrjpESqfuVTRiptJfP
        9JbdtWqRTmOf6uJi2c8YVqI6XlKXsD8C1dUUaaHKLUJzvKiazibVuBwMIT84AyqR
        QELn3e0BtgEymEygMU569b01ZPxoFSnNXc7qDZBDef8WfqAV/sxkTi8L9BkmFYfL
        uGLOhRJOFprPdoDIUBB+tmCl3oDcBy3vnUeOEioz8zAkprcb3GHwHAK+vHmmfgcn
        WsfMLH4JCLa/tRYL+Rw/N3ybCkDp00s0WUZ+AoDywSl0Q/ZEnNY0MsFiw6LyIdbq
        M/s/1JRtO3bDSzD9TazRVzn2oBqzSa8VgIo5C1nOnoAKJTlsClJKvIhnRlaLQqk=
        \"\"\",
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationAzureActiveDirectoryAuthenticationArgs']]]] azure_active_directory_authentications: A `azure_active_directory_authentication` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRevokedCertificateArgs']]]] client_revoked_certificates: One or more `client_revoked_certificate` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRootCertificateArgs']]]] client_root_certificates: One or more `client_root_certificate` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['VpnServerConfigurationIpsecPolicyArgs']] ipsec_policy: A `ipsec_policy` block as defined below.
        :param pulumi.Input[str] location: The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VpnServerConfigurationRadiusServerArgs']] radius_server: A `radius_server` block as defined below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpn_authentication_types: A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpn_protocols: A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
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

            __props__['azure_active_directory_authentications'] = azure_active_directory_authentications
            __props__['client_revoked_certificates'] = client_revoked_certificates
            __props__['client_root_certificates'] = client_root_certificates
            __props__['ipsec_policy'] = ipsec_policy
            __props__['location'] = location
            __props__['name'] = name
            __props__['radius_server'] = radius_server
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if vpn_authentication_types is None:
                raise TypeError("Missing required property 'vpn_authentication_types'")
            __props__['vpn_authentication_types'] = vpn_authentication_types
            __props__['vpn_protocols'] = vpn_protocols
        super(VpnServerConfiguration, __self__).__init__(
            'azure:network/vpnServerConfiguration:VpnServerConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            azure_active_directory_authentications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationAzureActiveDirectoryAuthenticationArgs']]]]] = None,
            client_revoked_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRevokedCertificateArgs']]]]] = None,
            client_root_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRootCertificateArgs']]]]] = None,
            ipsec_policy: Optional[pulumi.Input[pulumi.InputType['VpnServerConfigurationIpsecPolicyArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            radius_server: Optional[pulumi.Input[pulumi.InputType['VpnServerConfigurationRadiusServerArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpn_authentication_types: Optional[pulumi.Input[str]] = None,
            vpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'VpnServerConfiguration':
        """
        Get an existing VpnServerConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationAzureActiveDirectoryAuthenticationArgs']]]] azure_active_directory_authentications: A `azure_active_directory_authentication` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRevokedCertificateArgs']]]] client_revoked_certificates: One or more `client_revoked_certificate` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnServerConfigurationClientRootCertificateArgs']]]] client_root_certificates: One or more `client_root_certificate` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['VpnServerConfigurationIpsecPolicyArgs']] ipsec_policy: A `ipsec_policy` block as defined below.
        :param pulumi.Input[str] location: The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VpnServerConfigurationRadiusServerArgs']] radius_server: A `radius_server` block as defined below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpn_authentication_types: A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpn_protocols: A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["azure_active_directory_authentications"] = azure_active_directory_authentications
        __props__["client_revoked_certificates"] = client_revoked_certificates
        __props__["client_root_certificates"] = client_root_certificates
        __props__["ipsec_policy"] = ipsec_policy
        __props__["location"] = location
        __props__["name"] = name
        __props__["radius_server"] = radius_server
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["vpn_authentication_types"] = vpn_authentication_types
        __props__["vpn_protocols"] = vpn_protocols
        return VpnServerConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureActiveDirectoryAuthentications")
    def azure_active_directory_authentications(self) -> pulumi.Output[Optional[Sequence['outputs.VpnServerConfigurationAzureActiveDirectoryAuthentication']]]:
        """
        A `azure_active_directory_authentication` block as defined below.
        """
        return pulumi.get(self, "azure_active_directory_authentications")

    @property
    @pulumi.getter(name="clientRevokedCertificates")
    def client_revoked_certificates(self) -> pulumi.Output[Optional[Sequence['outputs.VpnServerConfigurationClientRevokedCertificate']]]:
        """
        One or more `client_revoked_certificate` blocks as defined below.
        """
        return pulumi.get(self, "client_revoked_certificates")

    @property
    @pulumi.getter(name="clientRootCertificates")
    def client_root_certificates(self) -> pulumi.Output[Optional[Sequence['outputs.VpnServerConfigurationClientRootCertificate']]]:
        """
        One or more `client_root_certificate` blocks as defined below.
        """
        return pulumi.get(self, "client_root_certificates")

    @property
    @pulumi.getter(name="ipsecPolicy")
    def ipsec_policy(self) -> pulumi.Output[Optional['outputs.VpnServerConfigurationIpsecPolicy']]:
        """
        A `ipsec_policy` block as defined below.
        """
        return pulumi.get(self, "ipsec_policy")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="radiusServer")
    def radius_server(self) -> pulumi.Output[Optional['outputs.VpnServerConfigurationRadiusServer']]:
        """
        A `radius_server` block as defined below.
        """
        return pulumi.get(self, "radius_server")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpnAuthenticationTypes")
    def vpn_authentication_types(self) -> pulumi.Output[str]:
        """
        A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
        """
        return pulumi.get(self, "vpn_authentication_types")

    @property
    @pulumi.getter(name="vpnProtocols")
    def vpn_protocols(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
        """
        return pulumi.get(self, "vpn_protocols")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

