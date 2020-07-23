# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.output_type
class ServiceCor(dict):
    allowed_origins: List[str] = pulumi.output_property("allowedOrigins")
    """
    A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceFeature(dict):
    flag: str = pulumi.output_property("flag")
    """
    The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
    """
    value: str = pulumi.output_property("value")
    """
    A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceSku(dict):
    capacity: float = pulumi.output_property("capacity")
    """
    Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
    """
    name: str = pulumi.output_property("name")
    """
    Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

