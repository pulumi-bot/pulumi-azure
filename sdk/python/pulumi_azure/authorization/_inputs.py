# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'RoleDefinitionPermissionArgs',
]

@pulumi.input_type
class RoleDefinitionPermissionArgs:
    def __init__(__self__, *,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 not_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 not_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: One or more Allowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_actions: One or more Allowed Data Actions, such as `*`, `Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] not_actions: One or more Disallowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] not_data_actions: One or more Disallowed Data Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if data_actions is not None:
            pulumi.set(__self__, "data_actions", data_actions)
        if not_actions is not None:
            pulumi.set(__self__, "not_actions", not_actions)
        if not_data_actions is not None:
            pulumi.set(__self__, "not_data_actions", not_data_actions)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more Allowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="dataActions")
    def data_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more Allowed Data Actions, such as `*`, `Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        return pulumi.get(self, "data_actions")

    @data_actions.setter
    def data_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "data_actions", value)

    @property
    @pulumi.getter(name="notActions")
    def not_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more Disallowed Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        return pulumi.get(self, "not_actions")

    @not_actions.setter
    def not_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "not_actions", value)

    @property
    @pulumi.getter(name="notDataActions")
    def not_data_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more Disallowed Data Actions, such as `*`, `Microsoft.Resources/subscriptions/resourceGroups/read`. See ['Azure Resource Manager resource provider operations'](https://docs.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for details.
        """
        return pulumi.get(self, "not_data_actions")

    @not_data_actions.setter
    def not_data_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "not_data_actions", value)


