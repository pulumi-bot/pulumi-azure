# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetDiagnosticCategoriesResult',
    'AwaitableGetDiagnosticCategoriesResult',
    'get_diagnostic_categories',
]

@pulumi.output_type
class GetDiagnosticCategoriesResult:
    """
    A collection of values returned by getDiagnosticCategories.
    """
    def __init__(__self__, id=None, logs=None, metrics=None, resource_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logs and not isinstance(logs, list):
            raise TypeError("Expected argument 'logs' to be a list")
        pulumi.set(__self__, "logs", logs)
        if metrics and not isinstance(metrics, list):
            raise TypeError("Expected argument 'metrics' to be a list")
        pulumi.set(__self__, "metrics", metrics)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def logs(self) -> Sequence[str]:
        """
        A list of the Log Categories supported for this Resource.
        """
        return pulumi.get(self, "logs")

    @property
    @pulumi.getter
    def metrics(self) -> Sequence[str]:
        """
        A list of the Metric Categories supported for this Resource.
        """
        return pulumi.get(self, "metrics")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        return pulumi.get(self, "resource_id")


class AwaitableGetDiagnosticCategoriesResult(GetDiagnosticCategoriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiagnosticCategoriesResult(
            id=self.id,
            logs=self.logs,
            metrics=self.metrics,
            resource_id=self.resource_id)


def get_diagnostic_categories(resource_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiagnosticCategoriesResult:
    """
    Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_key_vault = azure.keyvault.get_key_vault(name=azurerm_key_vault["example"]["name"],
        resource_group_name=azurerm_key_vault["example"]["resource_group_name"])
    example_diagnostic_categories = azure.monitoring.get_diagnostic_categories(resource_id=example_key_vault.id)
    ```


    :param str resource_id: The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
    """
    __args__ = dict()
    __args__['resourceId'] = resource_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:monitoring/getDiagnosticCategories:getDiagnosticCategories', __args__, opts=opts, typ=GetDiagnosticCategoriesResult).value

    return AwaitableGetDiagnosticCategoriesResult(
        id=__ret__.id,
        logs=__ret__.logs,
        metrics=__ret__.metrics,
        resource_id=__ret__.resource_id)
