# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetScheduledQueryRulesLogResult',
    'AwaitableGetScheduledQueryRulesLogResult',
    'get_scheduled_query_rules_log',
]

@pulumi.output_type
class GetScheduledQueryRulesLogResult:
    """
    A collection of values returned by getScheduledQueryRulesLog.
    """
    def __init__(__self__, authorized_resource_ids=None, criterias=None, data_source_id=None, description=None, enabled=None, id=None, location=None, name=None, resource_group_name=None, tags=None):
        if authorized_resource_ids and not isinstance(authorized_resource_ids, list):
            raise TypeError("Expected argument 'authorized_resource_ids' to be a list")
        pulumi.set(__self__, "authorized_resource_ids", authorized_resource_ids)
        if criterias and not isinstance(criterias, list):
            raise TypeError("Expected argument 'criterias' to be a list")
        pulumi.set(__self__, "criterias", criterias)
        if data_source_id and not isinstance(data_source_id, str):
            raise TypeError("Expected argument 'data_source_id' to be a str")
        pulumi.set(__self__, "data_source_id", data_source_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="authorizedResourceIds")
    def authorized_resource_ids(self) -> Sequence[str]:
        return pulumi.get(self, "authorized_resource_ids")

    @property
    @pulumi.getter
    def criterias(self) -> Sequence['outputs.GetScheduledQueryRulesLogCriteriaResult']:
        """
        A `criteria` block as defined below.
        """
        return pulumi.get(self, "criterias")

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
        """
        Name of the dimension.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")


class AwaitableGetScheduledQueryRulesLogResult(GetScheduledQueryRulesLogResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScheduledQueryRulesLogResult(
            authorized_resource_ids=self.authorized_resource_ids,
            criterias=self.criterias,
            data_source_id=self.data_source_id,
            description=self.description,
            enabled=self.enabled,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_scheduled_query_rules_log(name: Optional[str] = None,
                                  resource_group_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScheduledQueryRulesLogResult:
    """
    Use this data source to access the properties of a LogToMetricAction scheduled query rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.monitoring.get_scheduled_query_rules_log(name="tfex-queryrule",
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
    __ret__ = pulumi.runtime.invoke('azure:monitoring/getScheduledQueryRulesLog:getScheduledQueryRulesLog', __args__, opts=opts, typ=GetScheduledQueryRulesLogResult).value

    return AwaitableGetScheduledQueryRulesLogResult(
        authorized_resource_ids=__ret__.authorized_resource_ids,
        criterias=__ret__.criterias,
        data_source_id=__ret__.data_source_id,
        description=__ret__.description,
        enabled=__ret__.enabled,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags)
