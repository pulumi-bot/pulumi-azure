# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'FunctionJavaScriptUDFInput',
    'FunctionJavaScriptUDFOutput',
    'OutputBlobSerialization',
    'OutputEventHubSerialization',
    'OutputServiceBusQueueSerialization',
    'OutputServicebusTopicSerialization',
    'ReferenceInputBlobSerialization',
    'StreamInputBlobSerialization',
    'StreamInputEventHubSerialization',
    'StreamInputIotHubSerialization',
]

@pulumi.output_type
class FunctionJavaScriptUDFInput(dict):
    def __init__(__self__, *,
                 type: str):
        """
        :param str type: The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FunctionJavaScriptUDFOutput(dict):
    def __init__(__self__, *,
                 type: str):
        """
        :param str type: The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OutputBlobSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None,
                 format: Optional[str] = None):
        """
        :param str type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param str format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OutputEventHubSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None,
                 format: Optional[str] = None):
        """
        :param str type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param str format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OutputServiceBusQueueSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None,
                 format: Optional[str] = None):
        """
        :param str type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param str format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OutputServicebusTopicSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None,
                 format: Optional[str] = None):
        """
        :param str type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param str format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReferenceInputBlobSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None):
        """
        :param str type: The serialization format used for the reference data. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `	` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for the reference data. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `	` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StreamInputBlobSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None):
        """
        :param str type: The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StreamInputEventHubSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None):
        """
        :param str type: The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StreamInputIotHubSerialization(dict):
    def __init__(__self__, *,
                 type: str,
                 encoding: Optional[str] = None,
                 field_delimiter: Optional[str] = None):
        """
        :param str type: The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param str encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param str field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[str]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


