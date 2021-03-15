# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetCertificateDataResult',
    'AwaitableGetCertificateDataResult',
    'get_certificate_data',
]

@pulumi.output_type
class GetCertificateDataResult:
    """
    A collection of values returned by getCertificateData.
    """
    def __init__(__self__, expires=None, hex=None, id=None, key=None, key_vault_id=None, name=None, pem=None, tags=None, version=None):
        if expires and not isinstance(expires, str):
            raise TypeError("Expected argument 'expires' to be a str")
        pulumi.set(__self__, "expires", expires)
        if hex and not isinstance(hex, str):
            raise TypeError("Expected argument 'hex' to be a str")
        pulumi.set(__self__, "hex", hex)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if key_vault_id and not isinstance(key_vault_id, str):
            raise TypeError("Expected argument 'key_vault_id' to be a str")
        pulumi.set(__self__, "key_vault_id", key_vault_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pem and not isinstance(pem, str):
            raise TypeError("Expected argument 'pem' to be a str")
        pulumi.set(__self__, "pem", pem)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def expires(self) -> str:
        """
        Expiry date of certificate in RFC3339 format.
        """
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def hex(self) -> str:
        """
        The raw Key Vault Certificate data represented as a hexadecimal string.
        """
        return pulumi.get(self, "hex")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The Key Vault Certificate Key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> str:
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pem(self) -> str:
        """
        The Key Vault Certificate in PEM format.
        """
        return pulumi.get(self, "pem")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> str:
        return pulumi.get(self, "version")


class AwaitableGetCertificateDataResult(GetCertificateDataResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateDataResult(
            expires=self.expires,
            hex=self.hex,
            id=self.id,
            key=self.key,
            key_vault_id=self.key_vault_id,
            name=self.name,
            pem=self.pem,
            tags=self.tags,
            version=self.version)


def get_certificate_data(key_vault_id: Optional[str] = None,
                         name: Optional[str] = None,
                         version: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateDataResult:
    """
    Use this data source to access data stored in an existing Key Vault Certificate.

    > **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
    [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

    > **Note:** This data source uses the `GetSecret` function of the Azure API, to get the key of the certificate. Therefore you need secret/get permission

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_key_vault = azure.keyvault.get_key_vault(name="examplekv",
        resource_group_name="some-resource-group")
    example_certificate_data = azure.keyvault.get_certificate_data(name="secret-sauce",
        key_vault_id=example_key_vault.id)
    pulumi.export("examplePem", example_certificate_data.pem)
    ```


    :param str key_vault_id: Specifies the ID of the Key Vault instance where the Secret resides, available on the `keyvault.KeyVault` Data Source / Resource.
    :param str name: Specifies the name of the Key Vault Secret.
    :param str version: Specifies the version of the certificate to look up.  (Defaults to latest)
    """
    __args__ = dict()
    __args__['keyVaultId'] = key_vault_id
    __args__['name'] = name
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:keyvault/getCertificateData:getCertificateData', __args__, opts=opts, typ=GetCertificateDataResult).value

    return AwaitableGetCertificateDataResult(
        expires=__ret__.expires,
        hex=__ret__.hex,
        id=__ret__.id,
        key=__ret__.key,
        key_vault_id=__ret__.key_vault_id,
        name=__ret__.name,
        pem=__ret__.pem,
        tags=__ret__.tags,
        version=__ret__.version)
