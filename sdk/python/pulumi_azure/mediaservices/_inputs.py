# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AccountStorageAccountArgs',
]

@pulumi.input_type
class AccountStorageAccountArgs:
    id: pulumi.Input[str] = pulumi.input_property("id")
    """
    Specifies the ID of the Storage Account that will be associated with the Media Services instance.
    """
    is_primary: Optional[pulumi.Input[bool]] = pulumi.input_property("isPrimary")
    """
    Specifies whether the storage account should be the primary account or not. Defaults to `false`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, id: pulumi.Input[str], is_primary: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[str] id: Specifies the ID of the Storage Account that will be associated with the Media Services instance.
        :param pulumi.Input[bool] is_primary: Specifies whether the storage account should be the primary account or not. Defaults to `false`.
        """
        __self__.id = id
        __self__.is_primary = is_primary

