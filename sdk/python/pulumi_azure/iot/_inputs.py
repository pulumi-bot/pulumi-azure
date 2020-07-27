# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'IoTHubEndpointArgs',
    'IoTHubFallbackRouteArgs',
    'IoTHubFileUploadArgs',
    'IoTHubIpFilterRuleArgs',
    'IoTHubRouteArgs',
    'IoTHubSharedAccessPolicyArgs',
    'IoTHubSkuArgs',
    'IotHubDpsLinkedHubArgs',
    'IotHubDpsSkuArgs',
    'TimeSeriesInsightsReferenceDataSetKeyPropertyArgs',
]

@pulumi.input_type
class IoTHubEndpointArgs:
    connection_string: pulumi.Input[str] = pulumi.input_property("connectionString")
    """
    The connection string for the endpoint.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
    """
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
    """
    batch_frequency_in_seconds: Optional[pulumi.Input[float]] = pulumi.input_property("batchFrequencyInSeconds")
    """
    Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    container_name: Optional[pulumi.Input[str]] = pulumi.input_property("containerName")
    """
    The name of storage container in the storage account. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    encoding: Optional[pulumi.Input[str]] = pulumi.input_property("encoding")
    """
    Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    file_name_format: Optional[pulumi.Input[str]] = pulumi.input_property("fileNameFormat")
    """
    File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """
    max_chunk_size_in_bytes: Optional[pulumi.Input[float]] = pulumi.input_property("maxChunkSizeInBytes")
    """
    Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB). This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, connection_string: pulumi.Input[str], name: pulumi.Input[str], type: pulumi.Input[str], batch_frequency_in_seconds: Optional[pulumi.Input[float]] = None, container_name: Optional[pulumi.Input[str]] = None, encoding: Optional[pulumi.Input[str]] = None, file_name_format: Optional[pulumi.Input[str]] = None, max_chunk_size_in_bytes: Optional[pulumi.Input[float]] = None) -> None:
        """
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        :param pulumi.Input[str] type: The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
        :param pulumi.Input[float] batch_frequency_in_seconds: Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param pulumi.Input[str] container_name: The name of storage container in the storage account. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param pulumi.Input[str] encoding: Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param pulumi.Input[str] file_name_format: File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param pulumi.Input[float] max_chunk_size_in_bytes: Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB). This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        """
        __self__.connection_string = connection_string
        __self__.name = name
        __self__.type = type
        __self__.batch_frequency_in_seconds = batch_frequency_in_seconds
        __self__.container_name = container_name
        __self__.encoding = encoding
        __self__.file_name_format = file_name_format
        __self__.max_chunk_size_in_bytes = max_chunk_size_in_bytes

@pulumi.input_type
class IoTHubFallbackRouteArgs:
    condition: Optional[pulumi.Input[str]] = pulumi.input_property("condition")
    """
    The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
    """
    enabled: Optional[pulumi.Input[bool]] = pulumi.input_property("enabled")
    """
    Used to specify whether the fallback route is enabled.
    """
    endpoint_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("endpointNames")
    """
    The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
    """
    source: Optional[pulumi.Input[str]] = pulumi.input_property("source")
    """
    The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, condition: Optional[pulumi.Input[str]] = None, enabled: Optional[pulumi.Input[bool]] = None, endpoint_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, source: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] condition: The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        :param pulumi.Input[bool] enabled: Used to specify whether the fallback route is enabled.
        :param pulumi.Input[List[pulumi.Input[str]]] endpoint_names: The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
        :param pulumi.Input[str] source: The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
        """
        __self__.condition = condition
        __self__.enabled = enabled
        __self__.endpoint_names = endpoint_names
        __self__.source = source

@pulumi.input_type
class IoTHubFileUploadArgs:
    connection_string: pulumi.Input[str] = pulumi.input_property("connectionString")
    """
    The connection string for the Azure Storage account to which files are uploaded.
    """
    container_name: pulumi.Input[str] = pulumi.input_property("containerName")
    """
    The name of the root container where you upload files. The container need not exist but should be creatable using the connection_string specified.
    """
    default_ttl: Optional[pulumi.Input[str]] = pulumi.input_property("defaultTtl")
    """
    The period of time for which a file upload notification message is available to consume before it is expired by the IoT hub, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to 'PT1H' by default.
    """
    lock_duration: Optional[pulumi.Input[str]] = pulumi.input_property("lockDuration")
    """
    The lock duration for the file upload notifications queue, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 5 and 300 seconds, and evaluates to 'PT1M' by default.
    """
    max_delivery_count: Optional[pulumi.Input[float]] = pulumi.input_property("maxDeliveryCount")
    """
    The number of times the IoT hub attempts to deliver a file upload notification message. It evaluates to 10 by default.
    """
    notifications: Optional[pulumi.Input[bool]] = pulumi.input_property("notifications")
    """
    Used to specify whether file notifications are sent to IoT Hub on upload. It evaluates to false by default.
    """
    sas_ttl: Optional[pulumi.Input[str]] = pulumi.input_property("sasTtl")
    """
    The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 24 hours, and evaluates to 'PT1H' by default.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, connection_string: pulumi.Input[str], container_name: pulumi.Input[str], default_ttl: Optional[pulumi.Input[str]] = None, lock_duration: Optional[pulumi.Input[str]] = None, max_delivery_count: Optional[pulumi.Input[float]] = None, notifications: Optional[pulumi.Input[bool]] = None, sas_ttl: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] connection_string: The connection string for the Azure Storage account to which files are uploaded.
        :param pulumi.Input[str] container_name: The name of the root container where you upload files. The container need not exist but should be creatable using the connection_string specified.
        :param pulumi.Input[str] default_ttl: The period of time for which a file upload notification message is available to consume before it is expired by the IoT hub, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to 'PT1H' by default.
        :param pulumi.Input[str] lock_duration: The lock duration for the file upload notifications queue, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 5 and 300 seconds, and evaluates to 'PT1M' by default.
        :param pulumi.Input[float] max_delivery_count: The number of times the IoT hub attempts to deliver a file upload notification message. It evaluates to 10 by default.
        :param pulumi.Input[bool] notifications: Used to specify whether file notifications are sent to IoT Hub on upload. It evaluates to false by default.
        :param pulumi.Input[str] sas_ttl: The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 24 hours, and evaluates to 'PT1H' by default.
        """
        __self__.connection_string = connection_string
        __self__.container_name = container_name
        __self__.default_ttl = default_ttl
        __self__.lock_duration = lock_duration
        __self__.max_delivery_count = max_delivery_count
        __self__.notifications = notifications
        __self__.sas_ttl = sas_ttl

@pulumi.input_type
class IoTHubIpFilterRuleArgs:
    action: pulumi.Input[str] = pulumi.input_property("action")
    """
    The desired action for requests captured by this rule. Possible values are  `Accept`, `Reject`
    """
    ip_mask: pulumi.Input[str] = pulumi.input_property("ipMask")
    """
    The IP address range in CIDR notation for the rule.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the filter.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, action: pulumi.Input[str], ip_mask: pulumi.Input[str], name: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] action: The desired action for requests captured by this rule. Possible values are  `Accept`, `Reject`
        :param pulumi.Input[str] ip_mask: The IP address range in CIDR notation for the rule.
        :param pulumi.Input[str] name: The name of the filter.
        """
        __self__.action = action
        __self__.ip_mask = ip_mask
        __self__.name = name

@pulumi.input_type
class IoTHubRouteArgs:
    enabled: pulumi.Input[bool] = pulumi.input_property("enabled")
    """
    Used to specify whether a route is enabled.
    """
    endpoint_names: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("endpointNames")
    """
    The list of endpoints to which messages that satisfy the condition are routed.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the route.
    """
    source: pulumi.Input[str] = pulumi.input_property("source")
    """
    The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
    """
    condition: Optional[pulumi.Input[str]] = pulumi.input_property("condition")
    """
    The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, enabled: pulumi.Input[bool], endpoint_names: pulumi.Input[List[pulumi.Input[str]]], name: pulumi.Input[str], source: pulumi.Input[str], condition: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[bool] enabled: Used to specify whether a route is enabled.
        :param pulumi.Input[List[pulumi.Input[str]]] endpoint_names: The list of endpoints to which messages that satisfy the condition are routed.
        :param pulumi.Input[str] name: The name of the route.
        :param pulumi.Input[str] source: The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
        :param pulumi.Input[str] condition: The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        """
        __self__.enabled = enabled
        __self__.endpoint_names = endpoint_names
        __self__.name = name
        __self__.source = source
        __self__.condition = condition

@pulumi.input_type
class IoTHubSharedAccessPolicyArgs:
    key_name: Optional[pulumi.Input[str]] = pulumi.input_property("keyName")
    """
    The name of the shared access policy.
    """
    permissions: Optional[pulumi.Input[str]] = pulumi.input_property("permissions")
    """
    The permissions assigned to the shared access policy.
    """
    primary_key: Optional[pulumi.Input[str]] = pulumi.input_property("primaryKey")
    """
    The primary key.
    """
    secondary_key: Optional[pulumi.Input[str]] = pulumi.input_property("secondaryKey")
    """
    The secondary key.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key_name: Optional[pulumi.Input[str]] = None, permissions: Optional[pulumi.Input[str]] = None, primary_key: Optional[pulumi.Input[str]] = None, secondary_key: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] key_name: The name of the shared access policy.
        :param pulumi.Input[str] permissions: The permissions assigned to the shared access policy.
        :param pulumi.Input[str] primary_key: The primary key.
        :param pulumi.Input[str] secondary_key: The secondary key.
        """
        __self__.key_name = key_name
        __self__.permissions = permissions
        __self__.primary_key = primary_key
        __self__.secondary_key = secondary_key

@pulumi.input_type
class IoTHubSkuArgs:
    capacity: pulumi.Input[float] = pulumi.input_property("capacity")
    """
    The number of provisioned IoT Hub units.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, capacity: pulumi.Input[float], name: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[float] capacity: The number of provisioned IoT Hub units.
        :param pulumi.Input[str] name: The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
        """
        __self__.capacity = capacity
        __self__.name = name

@pulumi.input_type
class IotHubDpsLinkedHubArgs:
    connection_string: pulumi.Input[str] = pulumi.input_property("connectionString")
    """
    The connection string to connect to the IoT Hub. Changing this forces a new resource.
    """
    location: pulumi.Input[str] = pulumi.input_property("location")
    """
    The location of the IoT hub. Changing this forces a new resource.
    """
    allocation_weight: Optional[pulumi.Input[float]] = pulumi.input_property("allocationWeight")
    """
    The weight applied to the IoT Hub. Defaults to 0.
    """
    apply_allocation_policy: Optional[pulumi.Input[bool]] = pulumi.input_property("applyAllocationPolicy")
    """
    Determines whether to apply allocation policies to the IoT Hub. Defaults to false.
    """
    hostname: Optional[pulumi.Input[str]] = pulumi.input_property("hostname")
    """
    The IoT Hub hostname.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, connection_string: pulumi.Input[str], location: pulumi.Input[str], allocation_weight: Optional[pulumi.Input[float]] = None, apply_allocation_policy: Optional[pulumi.Input[bool]] = None, hostname: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] connection_string: The connection string to connect to the IoT Hub. Changing this forces a new resource.
        :param pulumi.Input[str] location: The location of the IoT hub. Changing this forces a new resource.
        :param pulumi.Input[float] allocation_weight: The weight applied to the IoT Hub. Defaults to 0.
        :param pulumi.Input[bool] apply_allocation_policy: Determines whether to apply allocation policies to the IoT Hub. Defaults to false.
        :param pulumi.Input[str] hostname: The IoT Hub hostname.
        """
        __self__.connection_string = connection_string
        __self__.location = location
        __self__.allocation_weight = allocation_weight
        __self__.apply_allocation_policy = apply_allocation_policy
        __self__.hostname = hostname

@pulumi.input_type
class IotHubDpsSkuArgs:
    capacity: pulumi.Input[float] = pulumi.input_property("capacity")
    """
    The number of provisioned IoT Device Provisioning Service units.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, capacity: pulumi.Input[float], name: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[float] capacity: The number of provisioned IoT Device Provisioning Service units.
        :param pulumi.Input[str] name: The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
        """
        __self__.capacity = capacity
        __self__.name = name

@pulumi.input_type
class TimeSeriesInsightsReferenceDataSetKeyPropertyArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the key property.
    """
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    The data type of the key property. Valid values include `Bool`, `DateTime`, `Double`, `String`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], type: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] name: The name of the key property.
        :param pulumi.Input[str] type: The data type of the key property. Valid values include `Bool`, `DateTime`, `Double`, `String`.
        """
        __self__.name = name
        __self__.type = type

