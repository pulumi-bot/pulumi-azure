# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'HubApnsCredential',
    'HubGcmCredential',
    'GetHubApnsCredentialResult',
    'GetHubGcmCredentialResult',
    'GetNamespaceSkuResult',
]

@pulumi.output_type
class HubApnsCredential(dict):
    def __init__(__self__, *,
                 application_mode: str,
                 bundle_id: str,
                 key_id: str,
                 team_id: str,
                 token: str):
        """
        :param str application_mode: The Application Mode which defines which server the APNS Messages should be sent to. Possible values are `Production` and `Sandbox`.
        :param str bundle_id: The Bundle ID of the iOS/macOS application to send push notifications for, such as `com.org.example`.
        :param str key_id: The Apple Push Notifications Service (APNS) Key.
        :param str team_id: The ID of the team the Token.
        :param str token: The Push Token associated with the Apple Developer Account. This is the contents of the `key` downloaded from [the Apple Developer Portal](https://developer.apple.com/account/ios/authkey/) between the `-----BEGIN PRIVATE KEY-----` and `-----END PRIVATE KEY-----` blocks.
        """
        pulumi.set(__self__, "application_mode", application_mode)
        pulumi.set(__self__, "bundle_id", bundle_id)
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "team_id", team_id)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="applicationMode")
    def application_mode(self) -> str:
        """
        The Application Mode which defines which server the APNS Messages should be sent to. Possible values are `Production` and `Sandbox`.
        """
        return pulumi.get(self, "application_mode")

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> str:
        """
        The Bundle ID of the iOS/macOS application to send push notifications for, such as `com.org.example`.
        """
        return pulumi.get(self, "bundle_id")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        The Apple Push Notifications Service (APNS) Key.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> str:
        """
        The ID of the team the Token.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The Push Token associated with the Apple Developer Account. This is the contents of the `key` downloaded from [the Apple Developer Portal](https://developer.apple.com/account/ios/authkey/) between the `-----BEGIN PRIVATE KEY-----` and `-----END PRIVATE KEY-----` blocks.
        """
        return pulumi.get(self, "token")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HubGcmCredential(dict):
    def __init__(__self__, *,
                 api_key: str):
        """
        :param str api_key: The API Key associated with the Google Cloud Messaging service.
        """
        pulumi.set(__self__, "api_key", api_key)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        The API Key associated with the Google Cloud Messaging service.
        """
        return pulumi.get(self, "api_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetHubApnsCredentialResult(dict):
    def __init__(__self__, *,
                 application_mode: str,
                 bundle_id: str,
                 key_id: str,
                 team_id: str,
                 token: str):
        """
        :param str application_mode: The Application Mode which defines which server the APNS Messages should be sent to. Possible values are `Production` and `Sandbox`.
        :param str bundle_id: The Bundle ID of the iOS/macOS application to send push notifications for, such as `com.org.example`.
        :param str key_id: The Apple Push Notifications Service (APNS) Key.
        :param str team_id: The ID of the team the Token.
        :param str token: The Push Token associated with the Apple Developer Account.
        """
        pulumi.set(__self__, "application_mode", application_mode)
        pulumi.set(__self__, "bundle_id", bundle_id)
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "team_id", team_id)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="applicationMode")
    def application_mode(self) -> str:
        """
        The Application Mode which defines which server the APNS Messages should be sent to. Possible values are `Production` and `Sandbox`.
        """
        return pulumi.get(self, "application_mode")

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> str:
        """
        The Bundle ID of the iOS/macOS application to send push notifications for, such as `com.org.example`.
        """
        return pulumi.get(self, "bundle_id")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        The Apple Push Notifications Service (APNS) Key.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> str:
        """
        The ID of the team the Token.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The Push Token associated with the Apple Developer Account.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class GetHubGcmCredentialResult(dict):
    def __init__(__self__, *,
                 api_key: str):
        """
        :param str api_key: The API Key associated with the Google Cloud Messaging service.
        """
        pulumi.set(__self__, "api_key", api_key)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        The API Key associated with the Google Cloud Messaging service.
        """
        return pulumi.get(self, "api_key")


@pulumi.output_type
class GetNamespaceSkuResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Specifies the Name of the Notification Hub Namespace.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the Name of the Notification Hub Namespace.
        """
        return pulumi.get(self, "name")


