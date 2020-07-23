# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.input_type
class DomainInputMappingDefaultValuesArgs:
    data_version: Optional[pulumi.Input[str]] = pulumi.input_property("dataVersion")
    """
    Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    event_type: Optional[pulumi.Input[str]] = pulumi.input_property("eventType")
    """
    Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    subject: Optional[pulumi.Input[str]] = pulumi.input_property("subject")
    """
    Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, data_version: Optional[pulumi.Input[str]] = None, event_type: Optional[pulumi.Input[str]] = None, subject: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] data_version: Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] event_type: Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subject: Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        """
        __self__.data_version = data_version
        __self__.event_type = event_type
        __self__.subject = subject

@pulumi.input_type
class DomainInputMappingFieldsArgs:
    data_version: Optional[pulumi.Input[str]] = pulumi.input_property("dataVersion")
    """
    Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    event_time: Optional[pulumi.Input[str]] = pulumi.input_property("eventTime")
    """
    Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    event_type: Optional[pulumi.Input[str]] = pulumi.input_property("eventType")
    """
    Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    id: Optional[pulumi.Input[str]] = pulumi.input_property("id")
    """
    Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    subject: Optional[pulumi.Input[str]] = pulumi.input_property("subject")
    """
    Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    topic: Optional[pulumi.Input[str]] = pulumi.input_property("topic")
    """
    Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, data_version: Optional[pulumi.Input[str]] = None, event_time: Optional[pulumi.Input[str]] = None, event_type: Optional[pulumi.Input[str]] = None, id: Optional[pulumi.Input[str]] = None, subject: Optional[pulumi.Input[str]] = None, topic: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] data_version: Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] event_time: Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] event_type: Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] id: Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subject: Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic: Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        """
        __self__.data_version = data_version
        __self__.event_time = event_time
        __self__.event_type = event_type
        __self__.id = id
        __self__.subject = subject
        __self__.topic = topic

@pulumi.input_type
class EventSubscriptionAdvancedFilterArgs:
    bool_equals: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterBoolEqualArgs']]]] = pulumi.input_property("boolEquals")
    """
    Compares a value of an event using a single boolean value.
    """
    number_greater_than_or_equals: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberGreaterThanOrEqualArgs']]]] = pulumi.input_property("numberGreaterThanOrEquals")
    """
    Compares a value of an event using a single floating point number.
    """
    number_greater_thans: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberGreaterThanArgs']]]] = pulumi.input_property("numberGreaterThans")
    """
    Compares a value of an event using a single floating point number.
    """
    number_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberInArgs']]]] = pulumi.input_property("numberIns")
    """
    Compares a value of an event using multiple floating point numbers.
    """
    number_less_than_or_equals: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberLessThanOrEqualArgs']]]] = pulumi.input_property("numberLessThanOrEquals")
    """
    Compares a value of an event using a single floating point number.
    """
    number_less_thans: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberLessThanArgs']]]] = pulumi.input_property("numberLessThans")
    """
    Compares a value of an event using a single floating point number.
    """
    number_not_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberNotInArgs']]]] = pulumi.input_property("numberNotIns")
    """
    Compares a value of an event using multiple floating point numbers.
    """
    string_begins_withs: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringBeginsWithArgs']]]] = pulumi.input_property("stringBeginsWiths")
    """
    Compares a value of an event using multiple string values.
    """
    string_contains: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringContainArgs']]]] = pulumi.input_property("stringContains")
    """
    Compares a value of an event using multiple string values.
    """
    string_ends_withs: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringEndsWithArgs']]]] = pulumi.input_property("stringEndsWiths")
    """
    Compares a value of an event using multiple string values.
    """
    string_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringInArgs']]]] = pulumi.input_property("stringIns")
    """
    Compares a value of an event using multiple string values.
    """
    string_not_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringNotInArgs']]]] = pulumi.input_property("stringNotIns")
    """
    Compares a value of an event using multiple string values.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, bool_equals: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterBoolEqualArgs']]]] = None, number_greater_than_or_equals: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberGreaterThanOrEqualArgs']]]] = None, number_greater_thans: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberGreaterThanArgs']]]] = None, number_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberInArgs']]]] = None, number_less_than_or_equals: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberLessThanOrEqualArgs']]]] = None, number_less_thans: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberLessThanArgs']]]] = None, number_not_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberNotInArgs']]]] = None, string_begins_withs: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringBeginsWithArgs']]]] = None, string_contains: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringContainArgs']]]] = None, string_ends_withs: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringEndsWithArgs']]]] = None, string_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringInArgs']]]] = None, string_not_ins: Optional[pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringNotInArgs']]]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterBoolEqualArgs']]] bool_equals: Compares a value of an event using a single boolean value.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberGreaterThanOrEqualArgs']]] number_greater_than_or_equals: Compares a value of an event using a single floating point number.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberGreaterThanArgs']]] number_greater_thans: Compares a value of an event using a single floating point number.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberInArgs']]] number_ins: Compares a value of an event using multiple floating point numbers.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberLessThanOrEqualArgs']]] number_less_than_or_equals: Compares a value of an event using a single floating point number.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberLessThanArgs']]] number_less_thans: Compares a value of an event using a single floating point number.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterNumberNotInArgs']]] number_not_ins: Compares a value of an event using multiple floating point numbers.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringBeginsWithArgs']]] string_begins_withs: Compares a value of an event using multiple string values.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringContainArgs']]] string_contains: Compares a value of an event using multiple string values.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringEndsWithArgs']]] string_ends_withs: Compares a value of an event using multiple string values.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringInArgs']]] string_ins: Compares a value of an event using multiple string values.
        :param pulumi.Input[List[pulumi.Input['EventSubscriptionAdvancedFilterStringNotInArgs']]] string_not_ins: Compares a value of an event using multiple string values.
        """
        __self__.bool_equals = bool_equals
        __self__.number_greater_than_or_equals = number_greater_than_or_equals
        __self__.number_greater_thans = number_greater_thans
        __self__.number_ins = number_ins
        __self__.number_less_than_or_equals = number_less_than_or_equals
        __self__.number_less_thans = number_less_thans
        __self__.number_not_ins = number_not_ins
        __self__.string_begins_withs = string_begins_withs
        __self__.string_contains = string_contains
        __self__.string_ends_withs = string_ends_withs
        __self__.string_ins = string_ins
        __self__.string_not_ins = string_not_ins

@pulumi.input_type
class EventSubscriptionAdvancedFilterBoolEqualArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    value: pulumi.Input[bool] = pulumi.input_property("value")
    """
    Specifies a single value to compare to when using a single value operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], value: pulumi.Input[bool]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[bool] value: Specifies a single value to compare to when using a single value operator.
        """
        __self__.key = key
        __self__.value = value

@pulumi.input_type
class EventSubscriptionAdvancedFilterNumberGreaterThanArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    value: pulumi.Input[float] = pulumi.input_property("value")
    """
    Specifies a single value to compare to when using a single value operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], value: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[float] value: Specifies a single value to compare to when using a single value operator.
        """
        __self__.key = key
        __self__.value = value

@pulumi.input_type
class EventSubscriptionAdvancedFilterNumberGreaterThanOrEqualArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    value: pulumi.Input[float] = pulumi.input_property("value")
    """
    Specifies a single value to compare to when using a single value operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], value: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[float] value: Specifies a single value to compare to when using a single value operator.
        """
        __self__.key = key
        __self__.value = value

@pulumi.input_type
class EventSubscriptionAdvancedFilterNumberInArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[float]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[float]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[float]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAdvancedFilterNumberLessThanArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    value: pulumi.Input[float] = pulumi.input_property("value")
    """
    Specifies a single value to compare to when using a single value operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], value: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[float] value: Specifies a single value to compare to when using a single value operator.
        """
        __self__.key = key
        __self__.value = value

@pulumi.input_type
class EventSubscriptionAdvancedFilterNumberLessThanOrEqualArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    value: pulumi.Input[float] = pulumi.input_property("value")
    """
    Specifies a single value to compare to when using a single value operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], value: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[float] value: Specifies a single value to compare to when using a single value operator.
        """
        __self__.key = key
        __self__.value = value

@pulumi.input_type
class EventSubscriptionAdvancedFilterNumberNotInArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[float]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[float]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[float]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAdvancedFilterStringBeginsWithArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[str]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[str]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAdvancedFilterStringContainArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[str]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[str]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAdvancedFilterStringEndsWithArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[str]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[str]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAdvancedFilterStringInArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[str]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[str]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAdvancedFilterStringNotInArgs:
    key: pulumi.Input[str] = pulumi.input_property("key")
    """
    Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
    """
    values: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("values")
    """
    Specifies an array of values to compare to when using a multiple values operator.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, key: pulumi.Input[str], values: pulumi.Input[List[pulumi.Input[str]]]) -> None:
        """
        :param pulumi.Input[str] key: Specifies the field within the event data that you want to use for filtering. Type of the field can be a number, boolean, or string.
        :param pulumi.Input[List[pulumi.Input[str]]] values: Specifies an array of values to compare to when using a multiple values operator.
        """
        __self__.key = key
        __self__.values = values

@pulumi.input_type
class EventSubscriptionAzureFunctionEndpointArgs:
    function_id: pulumi.Input[str] = pulumi.input_property("functionId")
    """
    Specifies the ID of the Function where the Event Subscription will receive events. This must be the functions ID in format {function_app.id}/functions/{name}.
    """
    max_events_per_batch: Optional[pulumi.Input[float]] = pulumi.input_property("maxEventsPerBatch")
    """
    Maximum number of events per batch.
    """
    preferred_batch_size_in_kilobytes: Optional[pulumi.Input[float]] = pulumi.input_property("preferredBatchSizeInKilobytes")
    """
    Preferred batch size in Kilobytes.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, function_id: pulumi.Input[str], max_events_per_batch: Optional[pulumi.Input[float]] = None, preferred_batch_size_in_kilobytes: Optional[pulumi.Input[float]] = None) -> None:
        """
        :param pulumi.Input[str] function_id: Specifies the ID of the Function where the Event Subscription will receive events. This must be the functions ID in format {function_app.id}/functions/{name}.
        :param pulumi.Input[float] max_events_per_batch: Maximum number of events per batch.
        :param pulumi.Input[float] preferred_batch_size_in_kilobytes: Preferred batch size in Kilobytes.
        """
        __self__.function_id = function_id
        __self__.max_events_per_batch = max_events_per_batch
        __self__.preferred_batch_size_in_kilobytes = preferred_batch_size_in_kilobytes

@pulumi.input_type
class EventSubscriptionEventhubEndpointArgs:
    eventhub_id: Optional[pulumi.Input[str]] = pulumi.input_property("eventhubId")
    """
    Specifies the id of the eventhub where the Event Subscription will receive events.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, eventhub_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] eventhub_id: Specifies the id of the eventhub where the Event Subscription will receive events.
        """
        __self__.eventhub_id = eventhub_id

@pulumi.input_type
class EventSubscriptionHybridConnectionEndpointArgs:
    hybrid_connection_id: Optional[pulumi.Input[str]] = pulumi.input_property("hybridConnectionId")
    """
    Specifies the id of the hybrid connection where the Event Subscription will receive events.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, hybrid_connection_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] hybrid_connection_id: Specifies the id of the hybrid connection where the Event Subscription will receive events.
        """
        __self__.hybrid_connection_id = hybrid_connection_id

@pulumi.input_type
class EventSubscriptionRetryPolicyArgs:
    event_time_to_live: pulumi.Input[float] = pulumi.input_property("eventTimeToLive")
    """
    Specifies the time to live (in minutes) for events.
    """
    max_delivery_attempts: pulumi.Input[float] = pulumi.input_property("maxDeliveryAttempts")
    """
    Specifies the maximum number of delivery retry attempts for events.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, event_time_to_live: pulumi.Input[float], max_delivery_attempts: pulumi.Input[float]) -> None:
        """
        :param pulumi.Input[float] event_time_to_live: Specifies the time to live (in minutes) for events.
        :param pulumi.Input[float] max_delivery_attempts: Specifies the maximum number of delivery retry attempts for events.
        """
        __self__.event_time_to_live = event_time_to_live
        __self__.max_delivery_attempts = max_delivery_attempts

@pulumi.input_type
class EventSubscriptionStorageBlobDeadLetterDestinationArgs:
    storage_account_id: pulumi.Input[str] = pulumi.input_property("storageAccountId")
    """
    Specifies the id of the storage account id where the storage blob is located.
    """
    storage_blob_container_name: pulumi.Input[str] = pulumi.input_property("storageBlobContainerName")
    """
    Specifies the name of the Storage blob container that is the destination of the deadletter events.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, storage_account_id: pulumi.Input[str], storage_blob_container_name: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] storage_account_id: Specifies the id of the storage account id where the storage blob is located.
        :param pulumi.Input[str] storage_blob_container_name: Specifies the name of the Storage blob container that is the destination of the deadletter events.
        """
        __self__.storage_account_id = storage_account_id
        __self__.storage_blob_container_name = storage_blob_container_name

@pulumi.input_type
class EventSubscriptionStorageQueueEndpointArgs:
    queue_name: pulumi.Input[str] = pulumi.input_property("queueName")
    """
    Specifies the name of the storage queue where the Event Subscription will receive events.
    """
    storage_account_id: pulumi.Input[str] = pulumi.input_property("storageAccountId")
    """
    Specifies the id of the storage account id where the storage queue is located.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, queue_name: pulumi.Input[str], storage_account_id: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] queue_name: Specifies the name of the storage queue where the Event Subscription will receive events.
        :param pulumi.Input[str] storage_account_id: Specifies the id of the storage account id where the storage queue is located.
        """
        __self__.queue_name = queue_name
        __self__.storage_account_id = storage_account_id

@pulumi.input_type
class EventSubscriptionSubjectFilterArgs:
    case_sensitive: Optional[pulumi.Input[bool]] = pulumi.input_property("caseSensitive")
    """
    Specifies if `subject_begins_with` and `subject_ends_with` case sensitive. This value defaults to `false`.
    """
    subject_begins_with: Optional[pulumi.Input[str]] = pulumi.input_property("subjectBeginsWith")
    """
    A string to filter events for an event subscription based on a resource path prefix.
    """
    subject_ends_with: Optional[pulumi.Input[str]] = pulumi.input_property("subjectEndsWith")
    """
    A string to filter events for an event subscription based on a resource path suffix.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, case_sensitive: Optional[pulumi.Input[bool]] = None, subject_begins_with: Optional[pulumi.Input[str]] = None, subject_ends_with: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[bool] case_sensitive: Specifies if `subject_begins_with` and `subject_ends_with` case sensitive. This value defaults to `false`.
        :param pulumi.Input[str] subject_begins_with: A string to filter events for an event subscription based on a resource path prefix.
        :param pulumi.Input[str] subject_ends_with: A string to filter events for an event subscription based on a resource path suffix.
        """
        __self__.case_sensitive = case_sensitive
        __self__.subject_begins_with = subject_begins_with
        __self__.subject_ends_with = subject_ends_with

@pulumi.input_type
class EventSubscriptionWebhookEndpointArgs:
    url: pulumi.Input[str] = pulumi.input_property("url")
    """
    Specifies the url of the webhook where the Event Subscription will receive events.
    """
    active_directory_app_id_or_uri: Optional[pulumi.Input[str]] = pulumi.input_property("activeDirectoryAppIdOrUri")
    """
    The Azure Active Directory Application ID or URI to get the access token that will be included as the bearer token in delivery requests.
    """
    active_directory_tenant_id: Optional[pulumi.Input[str]] = pulumi.input_property("activeDirectoryTenantId")
    """
    The Azure Active Directory Tenant ID to get the access token that will be included as the bearer token in delivery requests.
    """
    base_url: Optional[pulumi.Input[str]] = pulumi.input_property("baseUrl")
    """
    The base url of the webhook where the Event Subscription will receive events.
    """
    max_events_per_batch: Optional[pulumi.Input[float]] = pulumi.input_property("maxEventsPerBatch")
    """
    Maximum number of events per batch.
    """
    preferred_batch_size_in_kilobytes: Optional[pulumi.Input[float]] = pulumi.input_property("preferredBatchSizeInKilobytes")
    """
    Preferred batch size in Kilobytes.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, url: pulumi.Input[str], active_directory_app_id_or_uri: Optional[pulumi.Input[str]] = None, active_directory_tenant_id: Optional[pulumi.Input[str]] = None, base_url: Optional[pulumi.Input[str]] = None, max_events_per_batch: Optional[pulumi.Input[float]] = None, preferred_batch_size_in_kilobytes: Optional[pulumi.Input[float]] = None) -> None:
        """
        :param pulumi.Input[str] url: Specifies the url of the webhook where the Event Subscription will receive events.
        :param pulumi.Input[str] active_directory_app_id_or_uri: The Azure Active Directory Application ID or URI to get the access token that will be included as the bearer token in delivery requests.
        :param pulumi.Input[str] active_directory_tenant_id: The Azure Active Directory Tenant ID to get the access token that will be included as the bearer token in delivery requests.
        :param pulumi.Input[str] base_url: The base url of the webhook where the Event Subscription will receive events.
        :param pulumi.Input[float] max_events_per_batch: Maximum number of events per batch.
        :param pulumi.Input[float] preferred_batch_size_in_kilobytes: Preferred batch size in Kilobytes.
        """
        __self__.url = url
        __self__.active_directory_app_id_or_uri = active_directory_app_id_or_uri
        __self__.active_directory_tenant_id = active_directory_tenant_id
        __self__.base_url = base_url
        __self__.max_events_per_batch = max_events_per_batch
        __self__.preferred_batch_size_in_kilobytes = preferred_batch_size_in_kilobytes

@pulumi.input_type
class TopicInputMappingDefaultValuesArgs:
    data_version: Optional[pulumi.Input[str]] = pulumi.input_property("dataVersion")
    """
    Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    event_type: Optional[pulumi.Input[str]] = pulumi.input_property("eventType")
    """
    Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    subject: Optional[pulumi.Input[str]] = pulumi.input_property("subject")
    """
    Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, data_version: Optional[pulumi.Input[str]] = None, event_type: Optional[pulumi.Input[str]] = None, subject: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] data_version: Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] event_type: Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subject: Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        """
        __self__.data_version = data_version
        __self__.event_type = event_type
        __self__.subject = subject

@pulumi.input_type
class TopicInputMappingFieldsArgs:
    data_version: Optional[pulumi.Input[str]] = pulumi.input_property("dataVersion")
    """
    Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    event_time: Optional[pulumi.Input[str]] = pulumi.input_property("eventTime")
    """
    Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    event_type: Optional[pulumi.Input[str]] = pulumi.input_property("eventType")
    """
    Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    id: Optional[pulumi.Input[str]] = pulumi.input_property("id")
    """
    Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    subject: Optional[pulumi.Input[str]] = pulumi.input_property("subject")
    """
    Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """
    topic: Optional[pulumi.Input[str]] = pulumi.input_property("topic")
    """
    Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, data_version: Optional[pulumi.Input[str]] = None, event_time: Optional[pulumi.Input[str]] = None, event_type: Optional[pulumi.Input[str]] = None, id: Optional[pulumi.Input[str]] = None, subject: Optional[pulumi.Input[str]] = None, topic: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] data_version: Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] event_time: Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] event_type: Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] id: Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subject: Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic: Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        """
        __self__.data_version = data_version
        __self__.event_time = event_time
        __self__.event_type = event_type
        __self__.id = id
        __self__.subject = subject
        __self__.topic = topic
