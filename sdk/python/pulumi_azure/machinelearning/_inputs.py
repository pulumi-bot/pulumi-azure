# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'WorkspaceIdentityArgs',
]

@pulumi.input_type
class WorkspaceIdentityArgs:
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
    """
    principal_id: Optional[pulumi.Input[str]] = pulumi.input_property("principalId")
    """
    The (Client) ID of the Service Principal.
    """
    tenant_id: Optional[pulumi.Input[str]] = pulumi.input_property("tenantId")
    """
    The ID of the Tenant the Service Principal is assigned in.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, type: pulumi.Input[str], principal_id: Optional[pulumi.Input[str]] = None, tenant_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] type: The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The (Client) ID of the Service Principal.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant the Service Principal is assigned in.
        """
        __self__.type = type
        __self__.principal_id = principal_id
        __self__.tenant_id = tenant_id

