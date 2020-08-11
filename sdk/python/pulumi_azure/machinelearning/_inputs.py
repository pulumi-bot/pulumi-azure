# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'WorkspaceIdentityArgs',
]

@pulumi.input_type
class WorkspaceIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The (Client) ID of the Service Principal.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant the Service Principal is assigned in.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "principalId", principal_id)
        pulumi.set(__self__, "tenantId", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        """
        ...

    @type.setter
    def type(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The (Client) ID of the Service Principal.
        """
        ...

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Tenant the Service Principal is assigned in.
        """
        ...

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        ...


