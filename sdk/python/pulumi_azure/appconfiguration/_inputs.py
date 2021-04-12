# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'ConfigurationStoreIdentityArgs',
    'ConfigurationStorePrimaryReadKeyArgs',
    'ConfigurationStorePrimaryWriteKeyArgs',
    'ConfigurationStoreSecondaryReadKeyArgs',
    'ConfigurationStoreSecondaryWriteKeyArgs',
]

@pulumi.input_type
class ConfigurationStoreIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Specifies the identity type of the App Configuration. At this time the only allowed value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The ID of the Principal (Client) in Azure Active Directory.
        :param pulumi.Input[str] tenant_id: The ID of the Azure Active Directory Tenant.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the identity type of the App Configuration. At this time the only allowed value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Principal (Client) in Azure Active Directory.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Azure Active Directory Tenant.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class ConfigurationStorePrimaryReadKeyArgs:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] connection_string: The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        :param pulumi.Input[str] id: The ID of the Access Key.
        :param pulumi.Input[str] secret: The Secret of the Access Key.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Access Key.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The Secret of the Access Key.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class ConfigurationStorePrimaryWriteKeyArgs:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] connection_string: The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        :param pulumi.Input[str] id: The ID of the Access Key.
        :param pulumi.Input[str] secret: The Secret of the Access Key.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Access Key.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The Secret of the Access Key.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class ConfigurationStoreSecondaryReadKeyArgs:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] connection_string: The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        :param pulumi.Input[str] id: The ID of the Access Key.
        :param pulumi.Input[str] secret: The Secret of the Access Key.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Access Key.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The Secret of the Access Key.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class ConfigurationStoreSecondaryWriteKeyArgs:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] connection_string: The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        :param pulumi.Input[str] id: The ID of the Access Key.
        :param pulumi.Input[str] secret: The Secret of the Access Key.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Access Key.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The Secret of the Access Key.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


