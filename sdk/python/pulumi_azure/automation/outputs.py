# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ModuleModuleLink',
    'ModuleModuleLinkHash',
    'RunBookPublishContentLink',
    'RunBookPublishContentLinkHash',
    'ScheduleMonthlyOccurrence',
]

@pulumi.output_type
class ModuleModuleLink(dict):
    hash: Optional['outputs.ModuleModuleLinkHash'] = pulumi.output_property("hash")
    uri: str = pulumi.output_property("uri")
    """
    The uri of the module content (zip or nupkg).
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ModuleModuleLinkHash(dict):
    algorithm: str = pulumi.output_property("algorithm")
    value: str = pulumi.output_property("value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunBookPublishContentLink(dict):
    hash: Optional['outputs.RunBookPublishContentLinkHash'] = pulumi.output_property("hash")
    uri: str = pulumi.output_property("uri")
    """
    The uri of the runbook content.
    """
    version: Optional[str] = pulumi.output_property("version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunBookPublishContentLinkHash(dict):
    algorithm: str = pulumi.output_property("algorithm")
    value: str = pulumi.output_property("value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleMonthlyOccurrence(dict):
    day: str = pulumi.output_property("day")
    """
    Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
    """
    occurrence: float = pulumi.output_property("occurrence")
    """
    Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


