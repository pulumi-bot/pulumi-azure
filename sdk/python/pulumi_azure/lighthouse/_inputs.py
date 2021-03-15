# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'DefinitionAuthorizationArgs',
]

@pulumi.input_type
class DefinitionAuthorizationArgs:
    def __init__(__self__, *,
                 principal_id: pulumi.Input[str],
                 role_definition_id: pulumi.Input[str],
                 principal_display_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] principal_id: Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
        :param pulumi.Input[str] role_definition_id: The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
        :param pulumi.Input[str] principal_display_name: The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "role_definition_id", role_definition_id)
        if principal_display_name is not None:
            pulumi.set(__self__, "principal_display_name", principal_display_name)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Input[str]:
        """
        Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> pulumi.Input[str]:
        """
        The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
        """
        return pulumi.get(self, "role_definition_id")

    @role_definition_id.setter
    def role_definition_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_definition_id", value)

    @property
    @pulumi.getter(name="principalDisplayName")
    def principal_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
        """
        return pulumi.get(self, "principal_display_name")

    @principal_display_name.setter
    def principal_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_display_name", value)


