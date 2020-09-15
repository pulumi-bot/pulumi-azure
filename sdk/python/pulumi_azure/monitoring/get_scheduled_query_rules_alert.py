# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetScheduledQueryRulesAlertResult',
    'AwaitableGetScheduledQueryRulesAlertResult',
    'get_scheduled_query_rules_alert',
]

@pulumi.output_type
class GetScheduledQueryRulesAlertResult:
    """
    A collection of values returned by getScheduledQueryRulesAlert.
    """
    def __init__(__self__, actions=None, authorized_resource_ids=None, data_source_id=None, description=None, enabled=None, frequency=None, id=None, location=None, name=None, query=None, query_type=None, resource_group_name=None, severity=None, tags=None, throttling=None, time_window=None, triggers=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if authorized_resource_ids and not isinstance(authorized_resource_ids, list):
            raise TypeError("Expected argument 'authorized_resource_ids' to be a list")
        pulumi.set(__self__, "authorized_resource_ids", authorized_resource_ids)
        if data_source_id and not isinstance(data_source_id, str):
            raise TypeError("Expected argument 'data_source_id' to be a str")
        pulumi.set(__self__, "data_source_id", data_source_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if frequency and not isinstance(frequency, int):
            raise TypeError("Expected argument 'frequency' to be a int")
        pulumi.set(__self__, "frequency", frequency)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        pulumi.set(__self__, "query", query)
        if query_type and not isinstance(query_type, str):
            raise TypeError("Expected argument 'query_type' to be a str")
        pulumi.set(__self__, "query_type", query_type)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if severity and not isinstance(severity, int):
            raise TypeError("Expected argument 'severity' to be a int")
        pulumi.set(__self__, "severity", severity)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if throttling and not isinstance(throttling, int):
            raise TypeError("Expected argument 'throttling' to be a int")
        pulumi.set(__self__, "throttling", throttling)
        if time_window and not isinstance(time_window, int):
            raise TypeError("Expected argument 'time_window' to be a int")
        pulumi.set(__self__, "time_window", time_window)
        if triggers and not isinstance(triggers, list):
            raise TypeError("Expected argument 'triggers' to be a list")
        pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def actions(self) -> Sequence['outputs.GetScheduledQueryRulesAlertActionResult']:
        """
        An `action` block as defined below.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="authorizedResourceIds")
    def authorized_resource_ids(self) -> Sequence[str]:
        """
        The list of Resource IDs referred into query.
        """
        return pulumi.get(self, "authorized_resource_ids")

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> str:
        """
        The resource URI over which log search query is to be run.
        """
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the scheduled query rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether this scheduled query rule is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def frequency(self) -> int:
        """
        Frequency at which rule condition should be evaluated.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> str:
        """
        Log search query.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="queryType")
    def query_type(self) -> str:
        return pulumi.get(self, "query_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def severity(self) -> int:
        """
        Severity of the alert.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def throttling(self) -> int:
        """
        Time for which alerts should be throttled or suppressed.
        """
        return pulumi.get(self, "throttling")

    @property
    @pulumi.getter(name="timeWindow")
    def time_window(self) -> int:
        """
        Time window for which data needs to be fetched for query.
        """
        return pulumi.get(self, "time_window")

    @property
    @pulumi.getter
    def triggers(self) -> Sequence['outputs.GetScheduledQueryRulesAlertTriggerResult']:
        """
        A `trigger` block as defined below.
        """
        return pulumi.get(self, "triggers")


class AwaitableGetScheduledQueryRulesAlertResult(GetScheduledQueryRulesAlertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScheduledQueryRulesAlertResult(
            actions=self.actions,
            authorized_resource_ids=self.authorized_resource_ids,
            data_source_id=self.data_source_id,
            description=self.description,
            enabled=self.enabled,
            frequency=self.frequency,
            id=self.id,
            location=self.location,
            name=self.name,
            query=self.query,
            query_type=self.query_type,
            resource_group_name=self.resource_group_name,
            severity=self.severity,
            tags=self.tags,
            throttling=self.throttling,
            time_window=self.time_window,
            triggers=self.triggers)


def get_scheduled_query_rules_alert(name: Optional[str] = None,
                                    resource_group_name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScheduledQueryRulesAlertResult:
    """
    Use this data source to access the properties of an AlertingAction scheduled query rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.monitoring.get_scheduled_query_rules_alert(name="tfex-queryrule",
        resource_group_name="example-rg")
    pulumi.export("queryRuleId", example.id)
    ```


    :param str name: Specifies the name of the scheduled query rule.
    :param str resource_group_name: Specifies the name of the resource group where the scheduled query rule is located.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:monitoring/getScheduledQueryRulesAlert:getScheduledQueryRulesAlert', __args__, opts=opts, typ=GetScheduledQueryRulesAlertResult).value

    return AwaitableGetScheduledQueryRulesAlertResult(
        actions=__ret__.actions,
        authorized_resource_ids=__ret__.authorized_resource_ids,
        data_source_id=__ret__.data_source_id,
        description=__ret__.description,
        enabled=__ret__.enabled,
        frequency=__ret__.frequency,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        query=__ret__.query,
        query_type=__ret__.query_type,
        resource_group_name=__ret__.resource_group_name,
        severity=__ret__.severity,
        tags=__ret__.tags,
        throttling=__ret__.throttling,
        time_window=__ret__.time_window,
        triggers=__ret__.triggers)
