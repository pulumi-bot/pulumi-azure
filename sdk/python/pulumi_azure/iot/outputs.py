# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'IoTHubEndpoint',
    'IoTHubFallbackRoute',
    'IoTHubFileUpload',
    'IoTHubIpFilterRule',
    'IoTHubRoute',
    'IoTHubSharedAccessPolicy',
    'IoTHubSku',
    'IotHubDpsLinkedHub',
    'IotHubDpsSku',
    'TimeSeriesInsightsReferenceDataSetKeyProperty',
]

@pulumi.output_type
class IoTHubEndpoint(dict):
    batch_frequency_in_seconds: Optional[float] = pulumi.output_property("batchFrequencyInSeconds")
    """
    Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    connection_string: str = pulumi.output_property("connectionString")
    """
    The connection string for the endpoint.
    """
    container_name: Optional[str] = pulumi.output_property("containerName")
    """
    The name of storage container in the storage account. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    encoding: Optional[str] = pulumi.output_property("encoding")
    """
    Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    file_name_format: Optional[str] = pulumi.output_property("fileNameFormat")
    """
    File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    max_chunk_size_in_bytes: Optional[float] = pulumi.output_property("maxChunkSizeInBytes")
    """
    Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB). This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    name: str = pulumi.output_property("name")
    """
    The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
    """
    type: str = pulumi.output_property("type")
    """
    The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubFallbackRoute(dict):
    condition: Optional[str] = pulumi.output_property("condition")
    """
    The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
    """
    enabled: Optional[bool] = pulumi.output_property("enabled")
    """
    Used to specify whether the fallback route is enabled.
    """
    endpoint_names: Optional[List[str]] = pulumi.output_property("endpointNames")
    """
    The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
    """
    source: Optional[str] = pulumi.output_property("source")
    """
    The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubFileUpload(dict):
    connection_string: str = pulumi.output_property("connectionString")
    """
    The connection string for the Azure Storage account to which files are uploaded.
    """
    container_name: str = pulumi.output_property("containerName")
    """
    The name of the root container where you upload files. The container need not exist but should be creatable using the connection_string specified.
    """
    default_ttl: Optional[str] = pulumi.output_property("defaultTtl")
    """
    The period of time for which a file upload notification message is available to consume before it is expired by the IoT hub, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to 'PT1H' by default.
    """
    lock_duration: Optional[str] = pulumi.output_property("lockDuration")
    """
    The lock duration for the file upload notifications queue, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 5 and 300 seconds, and evaluates to 'PT1M' by default.
    """
    max_delivery_count: Optional[float] = pulumi.output_property("maxDeliveryCount")
    """
    The number of times the IoT hub attempts to deliver a file upload notification message. It evaluates to 10 by default.
    """
    notifications: Optional[bool] = pulumi.output_property("notifications")
    """
    Used to specify whether file notifications are sent to IoT Hub on upload. It evaluates to false by default.
    """
    sas_ttl: Optional[str] = pulumi.output_property("sasTtl")
    """
    The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 24 hours, and evaluates to 'PT1H' by default.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubIpFilterRule(dict):
    action: str = pulumi.output_property("action")
    """
    The desired action for requests captured by this rule. Possible values are  `Accept`, `Reject`
    """
    ip_mask: str = pulumi.output_property("ipMask")
    """
    The IP address range in CIDR notation for the rule.
    """
    name: str = pulumi.output_property("name")
    """
    The name of the filter.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubRoute(dict):
    condition: Optional[str] = pulumi.output_property("condition")
    """
    The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
    """
    enabled: bool = pulumi.output_property("enabled")
    """
    Used to specify whether a route is enabled.
    """
    endpoint_names: List[str] = pulumi.output_property("endpointNames")
    """
    The list of endpoints to which messages that satisfy the condition are routed.
    """
    name: str = pulumi.output_property("name")
    """
    The name of the route.
    """
    source: str = pulumi.output_property("source")
    """
    The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubSharedAccessPolicy(dict):
    key_name: Optional[str] = pulumi.output_property("keyName")
    """
    The name of the shared access policy.
    """
    permissions: Optional[str] = pulumi.output_property("permissions")
    """
    The permissions assigned to the shared access policy.
    """
    primary_key: Optional[str] = pulumi.output_property("primaryKey")
    """
    The primary key.
    """
    secondary_key: Optional[str] = pulumi.output_property("secondaryKey")
    """
    The secondary key.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubSku(dict):
    capacity: float = pulumi.output_property("capacity")
    """
    The number of provisioned IoT Hub units.
    """
    name: str = pulumi.output_property("name")
    """
    The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubDpsLinkedHub(dict):
    allocation_weight: Optional[float] = pulumi.output_property("allocationWeight")
    """
    The weight applied to the IoT Hub. Defaults to 0.
    """
    apply_allocation_policy: Optional[bool] = pulumi.output_property("applyAllocationPolicy")
    """
    Determines whether to apply allocation policies to the IoT Hub. Defaults to false.
    """
    connection_string: str = pulumi.output_property("connectionString")
    """
    The connection string to connect to the IoT Hub. Changing this forces a new resource.
    """
    hostname: Optional[str] = pulumi.output_property("hostname")
    """
    The IoT Hub hostname.
    """
    location: str = pulumi.output_property("location")
    """
    The location of the IoT hub. Changing this forces a new resource.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubDpsSku(dict):
    capacity: float = pulumi.output_property("capacity")
    """
    The number of provisioned IoT Device Provisioning Service units.
    """
    name: str = pulumi.output_property("name")
    """
    The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TimeSeriesInsightsReferenceDataSetKeyProperty(dict):
    name: str = pulumi.output_property("name")
    """
    The name of the key property.
    """
    type: str = pulumi.output_property("type")
    """
    The data type of the key property. Valid values include `Bool`, `DateTime`, `Double`, `String`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


