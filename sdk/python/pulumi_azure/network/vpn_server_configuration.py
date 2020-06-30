# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VpnServerConfiguration(pulumi.CustomResource):
    azure_active_directory_authentications: pulumi.Output[list]
    """
    A `azure_active_directory_authentication` block as defined below.

      * `audience` (`str`) - The Audience which should be used for authentication.
      * `issuer` (`str`) - The Issuer which should be used for authentication.
      * `tenant` (`str`) - The Tenant which should be used for authentication.
    """
    client_revoked_certificates: pulumi.Output[list]
    """
    One or more `client_revoked_certificate` blocks as defined below.

      * `name` (`str`) - A name used to uniquely identify this certificate.
      * `thumbprint` (`str`) - The Thumbprint of the Certificate.
    """
    client_root_certificates: pulumi.Output[list]
    """
    One or more `client_root_certificate` blocks as defined below.

      * `name` (`str`) - A name used to uniquely identify this certificate.
      * `publicCertData` (`str`) - The Public Key Data associated with the Certificate.
    """
    ipsec_policy: pulumi.Output[dict]
    """
    A `ipsec_policy` block as defined below.

      * `dhGroup` (`str`) - The DH Group, used in IKE Phase 1. Possible values include `DHGroup1`, `DHGroup2`, `DHGroup14`, `DHGroup24`, `DHGroup2048`, `ECP256`, `ECP384` and `None`.
      * `ikeEncryption` (`str`) - The IKE encryption algorithm, used for IKE Phase 2. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128` and `GCMAES256`.
      * `ikeIntegrity` (`str`) - The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include `GCMAES128`, `GCMAES256`, `MD5`, `SHA1`, `SHA256` and `SHA384`.
      * `ipsecEncryption` (`str`) - The IPSec encryption algorithm, used for IKE phase 1. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128`, `GCMAES192`, `GCMAES256` and `None`.
      * `ipsecIntegrity` (`str`) - The IPSec integrity algorithm, used for IKE phase 1. Possible values include `GCMAES128`, `GCMAES192`, `GCMAES256`, `MD5`, `SHA1` and `SHA256`.
      * `pfsGroup` (`str`) - The Pfs Group, used in IKE Phase 2. Possible values include `ECP256`, `ECP384`, `PFS1`, `PFS2`, `PFS14`, `PFS24`, `PFS2048`, `PFSMM` and `None`.
      * `saDataSizeKilobytes` (`float`) - The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
      * `saLifetimeSeconds` (`float`) - The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.
    """
    location: pulumi.Output[str]
    """
    The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
    """
    radius_server: pulumi.Output[dict]
    """
    A `radius_server` block as defined below.

      * `address` (`str`) - The Address of the Radius Server.
      * `client_root_certificates` (`list`) - One or more `client_root_certificate` blocks as defined above.
        * `name` (`str`) - A name used to uniquely identify this certificate.
        * `thumbprint` (`str`) - The Thumbprint of the Certificate.

      * `secret` (`str`) - The Secret used to communicate with the Radius Server.
      * `serverRootCertificates` (`list`) - One or more `server_root_certificate` blocks as defined below.
        * `name` (`str`) - A name used to uniquely identify this certificate.
        * `publicCertData` (`str`) - The Public Key Data associated with the Certificate.
    """
    resource_group_name: pulumi.Output[str]
    """
    The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpn_authentication_types: pulumi.Output[str]
    """
    A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
    """
    vpn_protocols: pulumi.Output[list]
    """
    A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
    """
    def __init__(__self__, resource_name, opts=None, azure_active_directory_authentications=None, client_revoked_certificates=None, client_root_certificates=None, ipsec_policy=None, location=None, name=None, radius_server=None, resource_group_name=None, tags=None, vpn_authentication_types=None, vpn_protocols=None, __props__=None, __name__=None, __opts__=None):
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
            client_root_certificate=[{
                "name": "DigiCert-Federated-ID-Root-CA",
                "publicCertData": \"\"\"MIIDuzCCAqOgAwIBAgIQCHTZWCM+IlfFIRXIvyKSrjANBgkqhkiG9w0BAQsFADBn
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
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] azure_active_directory_authentications: A `azure_active_directory_authentication` block as defined below.
        :param pulumi.Input[list] client_revoked_certificates: One or more `client_revoked_certificate` blocks as defined below.
        :param pulumi.Input[list] client_root_certificates: One or more `client_root_certificate` blocks as defined below.
        :param pulumi.Input[dict] ipsec_policy: A `ipsec_policy` block as defined below.
        :param pulumi.Input[str] location: The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] radius_server: A `radius_server` block as defined below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpn_authentication_types: A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
        :param pulumi.Input[list] vpn_protocols: A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.

        The **azure_active_directory_authentications** object supports the following:

          * `audience` (`pulumi.Input[str]`) - The Audience which should be used for authentication.
          * `issuer` (`pulumi.Input[str]`) - The Issuer which should be used for authentication.
          * `tenant` (`pulumi.Input[str]`) - The Tenant which should be used for authentication.

        The **client_revoked_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.

        The **client_root_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
          * `publicCertData` (`pulumi.Input[str]`) - The Public Key Data associated with the Certificate.

        The **ipsec_policy** object supports the following:

          * `dhGroup` (`pulumi.Input[str]`) - The DH Group, used in IKE Phase 1. Possible values include `DHGroup1`, `DHGroup2`, `DHGroup14`, `DHGroup24`, `DHGroup2048`, `ECP256`, `ECP384` and `None`.
          * `ikeEncryption` (`pulumi.Input[str]`) - The IKE encryption algorithm, used for IKE Phase 2. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128` and `GCMAES256`.
          * `ikeIntegrity` (`pulumi.Input[str]`) - The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include `GCMAES128`, `GCMAES256`, `MD5`, `SHA1`, `SHA256` and `SHA384`.
          * `ipsecEncryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm, used for IKE phase 1. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128`, `GCMAES192`, `GCMAES256` and `None`.
          * `ipsecIntegrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm, used for IKE phase 1. Possible values include `GCMAES128`, `GCMAES192`, `GCMAES256`, `MD5`, `SHA1` and `SHA256`.
          * `pfsGroup` (`pulumi.Input[str]`) - The Pfs Group, used in IKE Phase 2. Possible values include `ECP256`, `ECP384`, `PFS1`, `PFS2`, `PFS14`, `PFS24`, `PFS2048`, `PFSMM` and `None`.
          * `saDataSizeKilobytes` (`pulumi.Input[float]`) - The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
          * `saLifetimeSeconds` (`pulumi.Input[float]`) - The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.

        The **radius_server** object supports the following:

          * `address` (`pulumi.Input[str]`) - The Address of the Radius Server.
          * `client_root_certificates` (`pulumi.Input[list]`) - One or more `client_root_certificate` blocks as defined above.
            * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
            * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.

          * `secret` (`pulumi.Input[str]`) - The Secret used to communicate with the Radius Server.
          * `serverRootCertificates` (`pulumi.Input[list]`) - One or more `server_root_certificate` blocks as defined below.
            * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
            * `publicCertData` (`pulumi.Input[str]`) - The Public Key Data associated with the Certificate.
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
    def get(resource_name, id, opts=None, azure_active_directory_authentications=None, client_revoked_certificates=None, client_root_certificates=None, ipsec_policy=None, location=None, name=None, radius_server=None, resource_group_name=None, tags=None, vpn_authentication_types=None, vpn_protocols=None):
        """
        Get an existing VpnServerConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] azure_active_directory_authentications: A `azure_active_directory_authentication` block as defined below.
        :param pulumi.Input[list] client_revoked_certificates: One or more `client_revoked_certificate` blocks as defined below.
        :param pulumi.Input[list] client_root_certificates: One or more `client_root_certificate` blocks as defined below.
        :param pulumi.Input[dict] ipsec_policy: A `ipsec_policy` block as defined below.
        :param pulumi.Input[str] location: The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] radius_server: A `radius_server` block as defined below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpn_authentication_types: A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
        :param pulumi.Input[list] vpn_protocols: A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.

        The **azure_active_directory_authentications** object supports the following:

          * `audience` (`pulumi.Input[str]`) - The Audience which should be used for authentication.
          * `issuer` (`pulumi.Input[str]`) - The Issuer which should be used for authentication.
          * `tenant` (`pulumi.Input[str]`) - The Tenant which should be used for authentication.

        The **client_revoked_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.

        The **client_root_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
          * `publicCertData` (`pulumi.Input[str]`) - The Public Key Data associated with the Certificate.

        The **ipsec_policy** object supports the following:

          * `dhGroup` (`pulumi.Input[str]`) - The DH Group, used in IKE Phase 1. Possible values include `DHGroup1`, `DHGroup2`, `DHGroup14`, `DHGroup24`, `DHGroup2048`, `ECP256`, `ECP384` and `None`.
          * `ikeEncryption` (`pulumi.Input[str]`) - The IKE encryption algorithm, used for IKE Phase 2. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128` and `GCMAES256`.
          * `ikeIntegrity` (`pulumi.Input[str]`) - The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include `GCMAES128`, `GCMAES256`, `MD5`, `SHA1`, `SHA256` and `SHA384`.
          * `ipsecEncryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm, used for IKE phase 1. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128`, `GCMAES192`, `GCMAES256` and `None`.
          * `ipsecIntegrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm, used for IKE phase 1. Possible values include `GCMAES128`, `GCMAES192`, `GCMAES256`, `MD5`, `SHA1` and `SHA256`.
          * `pfsGroup` (`pulumi.Input[str]`) - The Pfs Group, used in IKE Phase 2. Possible values include `ECP256`, `ECP384`, `PFS1`, `PFS2`, `PFS14`, `PFS24`, `PFS2048`, `PFSMM` and `None`.
          * `saDataSizeKilobytes` (`pulumi.Input[float]`) - The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
          * `saLifetimeSeconds` (`pulumi.Input[float]`) - The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.

        The **radius_server** object supports the following:

          * `address` (`pulumi.Input[str]`) - The Address of the Radius Server.
          * `client_root_certificates` (`pulumi.Input[list]`) - One or more `client_root_certificate` blocks as defined above.
            * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
            * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.

          * `secret` (`pulumi.Input[str]`) - The Secret used to communicate with the Radius Server.
          * `serverRootCertificates` (`pulumi.Input[list]`) - One or more `server_root_certificate` blocks as defined below.
            * `name` (`pulumi.Input[str]`) - A name used to uniquely identify this certificate.
            * `publicCertData` (`pulumi.Input[str]`) - The Public Key Data associated with the Certificate.
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
