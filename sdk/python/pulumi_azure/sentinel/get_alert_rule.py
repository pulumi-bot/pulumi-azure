# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAlertRuleResult:
    """
    A collection of values returned by getAlertRule.
    """
    def __init__(__self__, id=None, log_analytics_workspace_id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if log_analytics_workspace_id and not isinstance(log_analytics_workspace_id, str):
            raise TypeError("Expected argument 'log_analytics_workspace_id' to be a str")
        __self__.log_analytics_workspace_id = log_analytics_workspace_id
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetAlertRuleResult(GetAlertRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlertRuleResult(
            id=self.id,
            log_analytics_workspace_id=self.log_analytics_workspace_id,
            name=self.name)

def get_alert_rule(log_analytics_workspace_id=None,name=None,opts=None):
    """
    Use this data source to access information about an existing Sentinel Alert Rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_analytics_workspace = azure.operationalinsights.get_analytics_workspace(name="example",
        resource_group_name="example-resources")
    example_alert_rule = azure.sentinel.get_alert_rule(name="existing",
        log_analytics_workspace_id=example_analytics_workspace.id)
    pulumi.export("id", example_alert_rule.id)
    ```


    :param str log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel Alert Rule belongs to.
    :param str name: The name which should be used for this Sentinel Alert Rule.
    """
    __args__ = dict()


    __args__['logAnalyticsWorkspaceId'] = log_analytics_workspace_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:sentinel/getAlertRule:getAlertRule', __args__, opts=opts).value

    return AwaitableGetAlertRuleResult(
        id=__ret__.get('id'),
        log_analytics_workspace_id=__ret__.get('logAnalyticsWorkspaceId'),
        name=__ret__.get('name'))
