# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'NamespaceNetworkRuleSetNetworkRule',
    'SubscriptionRuleCorrelationFilter',
]

@pulumi.output_type
class NamespaceNetworkRuleSetNetworkRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "subnetId":
            suggest = "subnet_id"
        elif key == "ignoreMissingVnetServiceEndpoint":
            suggest = "ignore_missing_vnet_service_endpoint"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NamespaceNetworkRuleSetNetworkRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NamespaceNetworkRuleSetNetworkRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NamespaceNetworkRuleSetNetworkRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 subnet_id: str,
                 ignore_missing_vnet_service_endpoint: Optional[bool] = None):
        """
        :param str subnet_id: The Subnet ID which should be able to access this ServiceBus Namespace.
        :param bool ignore_missing_vnet_service_endpoint: Should the ServiceBus Namespace Network Rule Set ignore missing Virtual Network Service Endpoint option in the Subnet? Defaults to `false`.
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        if ignore_missing_vnet_service_endpoint is not None:
            pulumi.set(__self__, "ignore_missing_vnet_service_endpoint", ignore_missing_vnet_service_endpoint)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The Subnet ID which should be able to access this ServiceBus Namespace.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="ignoreMissingVnetServiceEndpoint")
    def ignore_missing_vnet_service_endpoint(self) -> Optional[bool]:
        """
        Should the ServiceBus Namespace Network Rule Set ignore missing Virtual Network Service Endpoint option in the Subnet? Defaults to `false`.
        """
        return pulumi.get(self, "ignore_missing_vnet_service_endpoint")


@pulumi.output_type
class SubscriptionRuleCorrelationFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "contentType":
            suggest = "content_type"
        elif key == "correlationId":
            suggest = "correlation_id"
        elif key == "messageId":
            suggest = "message_id"
        elif key == "replyTo":
            suggest = "reply_to"
        elif key == "replyToSessionId":
            suggest = "reply_to_session_id"
        elif key == "sessionId":
            suggest = "session_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionRuleCorrelationFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionRuleCorrelationFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionRuleCorrelationFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 content_type: Optional[str] = None,
                 correlation_id: Optional[str] = None,
                 label: Optional[str] = None,
                 message_id: Optional[str] = None,
                 properties: Optional[Mapping[str, str]] = None,
                 reply_to: Optional[str] = None,
                 reply_to_session_id: Optional[str] = None,
                 session_id: Optional[str] = None,
                 to: Optional[str] = None):
        """
        :param str content_type: Content type of the message.
        :param str correlation_id: Identifier of the correlation.
        :param str label: Application specific label.
        :param str message_id: Identifier of the message.
        :param Mapping[str, str] properties: A list of user defined properties to be included in the filter. Specified as a map of name/value pairs.
        :param str reply_to: Address of the queue to reply to.
        :param str reply_to_session_id: Session identifier to reply to.
        :param str session_id: Session identifier.
        :param str to: Address to send to.
        """
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if correlation_id is not None:
            pulumi.set(__self__, "correlation_id", correlation_id)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if message_id is not None:
            pulumi.set(__self__, "message_id", message_id)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if reply_to is not None:
            pulumi.set(__self__, "reply_to", reply_to)
        if reply_to_session_id is not None:
            pulumi.set(__self__, "reply_to_session_id", reply_to_session_id)
        if session_id is not None:
            pulumi.set(__self__, "session_id", session_id)
        if to is not None:
            pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        """
        Content type of the message.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="correlationId")
    def correlation_id(self) -> Optional[str]:
        """
        Identifier of the correlation.
        """
        return pulumi.get(self, "correlation_id")

    @property
    @pulumi.getter
    def label(self) -> Optional[str]:
        """
        Application specific label.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="messageId")
    def message_id(self) -> Optional[str]:
        """
        Identifier of the message.
        """
        return pulumi.get(self, "message_id")

    @property
    @pulumi.getter
    def properties(self) -> Optional[Mapping[str, str]]:
        """
        A list of user defined properties to be included in the filter. Specified as a map of name/value pairs.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="replyTo")
    def reply_to(self) -> Optional[str]:
        """
        Address of the queue to reply to.
        """
        return pulumi.get(self, "reply_to")

    @property
    @pulumi.getter(name="replyToSessionId")
    def reply_to_session_id(self) -> Optional[str]:
        """
        Session identifier to reply to.
        """
        return pulumi.get(self, "reply_to_session_id")

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> Optional[str]:
        """
        Session identifier.
        """
        return pulumi.get(self, "session_id")

    @property
    @pulumi.getter
    def to(self) -> Optional[str]:
        """
        Address to send to.
        """
        return pulumi.get(self, "to")


