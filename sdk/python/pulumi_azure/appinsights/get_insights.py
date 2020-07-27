# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetInsightsResult',
    'AwaitableGetInsightsResult',
    'get_insights',
]


class GetInsightsResult:
    """
    A collection of values returned by getInsights.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, app_id=None, application_type=None, id=None, instrumentation_key=None, location=None, name=None, resource_group_name=None, retention_in_days=None, tags=None) -> None:
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        __self__.app_id = app_id
        """
        The App ID associated with this Application Insights component.
        """
        if application_type and not isinstance(application_type, str):
            raise TypeError("Expected argument 'application_type' to be a str")
        __self__.application_type = application_type
        """
        The type of the component.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instrumentation_key and not isinstance(instrumentation_key, str):
            raise TypeError("Expected argument 'instrumentation_key' to be a str")
        __self__.instrumentation_key = instrumentation_key
        """
        The instrumentation key of the Application Insights component.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure location where the component exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if retention_in_days and not isinstance(retention_in_days, float):
            raise TypeError("Expected argument 'retention_in_days' to be a float")
        __self__.retention_in_days = retention_in_days
        """
        The retention period in days.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Tags applied to the component.
        """


class AwaitableGetInsightsResult(GetInsightsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInsightsResult(
            app_id=self.app_id,
            application_type=self.application_type,
            id=self.id,
            instrumentation_key=self.instrumentation_key,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            retention_in_days=self.retention_in_days,
            tags=self.tags)


def get_insights(name: Optional[str] = None, resource_group_name: Optional[str] = None, opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInsightsResult:
    """
    Use this data source to access information about an existing Application Insights component.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.appinsights.get_insights(name="production",
        resource_group_name="networking")
    pulumi.export("applicationInsightsInstrumentationKey", example.instrumentation_key)
    ```


    :param str name: Specifies the name of the Application Insights component.
    :param str resource_group_name: Specifies the name of the resource group the Application Insights component is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:appinsights/getInsights:getInsights', __args__, opts=opts).value

    return AwaitableGetInsightsResult(
        app_id=__ret__.get('appId'),
        application_type=__ret__.get('applicationType'),
        id=__ret__.get('id'),
        instrumentation_key=__ret__.get('instrumentationKey'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        retention_in_days=__ret__.get('retentionInDays'),
        tags=__ret__.get('tags'))
