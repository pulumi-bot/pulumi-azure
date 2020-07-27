# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ServiceQueryKey',
]

@pulumi.output_type
class ServiceQueryKey(dict):
    key: Optional[str] = pulumi.output_property("key")
    """
    The value of this Query Key.
    """
    name: Optional[str] = pulumi.output_property("name")
    """
    The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


