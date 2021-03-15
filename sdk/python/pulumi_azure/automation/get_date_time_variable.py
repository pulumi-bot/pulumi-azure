# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetDateTimeVariableResult',
    'AwaitableGetDateTimeVariableResult',
    'get_date_time_variable',
]

@pulumi.output_type
class GetDateTimeVariableResult:
    """
    A collection of values returned by getDateTimeVariable.
    """
    def __init__(__self__, automation_account_name=None, description=None, encrypted=None, id=None, name=None, resource_group_name=None, value=None):
        if automation_account_name and not isinstance(automation_account_name, str):
            raise TypeError("Expected argument 'automation_account_name' to be a str")
        pulumi.set(__self__, "automation_account_name", automation_account_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> str:
        return pulumi.get(self, "automation_account_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the Automation Variable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        """
        Specifies if the Automation Variable is encrypted. Defaults to `false`.
        """
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
        """
        return pulumi.get(self, "value")


class AwaitableGetDateTimeVariableResult(GetDateTimeVariableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDateTimeVariableResult(
            automation_account_name=self.automation_account_name,
            description=self.description,
            encrypted=self.encrypted,
            id=self.id,
            name=self.name,
            resource_group_name=self.resource_group_name,
            value=self.value)


def get_date_time_variable(automation_account_name: Optional[str] = None,
                           name: Optional[str] = None,
                           resource_group_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDateTimeVariableResult:
    """
    Use this data source to access information about an existing Automation Datetime Variable.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.automation.get_date_time_variable(name="tfex-example-var",
        resource_group_name="tfex-example-rg",
        automation_account_name="tfex-example-account")
    pulumi.export("variableId", example.id)
    ```


    :param str automation_account_name: The name of the automation account in which the Automation Variable exists.
    :param str name: The name of the Automation Variable.
    :param str resource_group_name: The Name of the Resource Group where the automation account exists.
    """
    __args__ = dict()
    __args__['automationAccountName'] = automation_account_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:automation/getDateTimeVariable:getDateTimeVariable', __args__, opts=opts, typ=GetDateTimeVariableResult).value

    return AwaitableGetDateTimeVariableResult(
        automation_account_name=__ret__.automation_account_name,
        description=__ret__.description,
        encrypted=__ret__.encrypted,
        id=__ret__.id,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        value=__ret__.value)
