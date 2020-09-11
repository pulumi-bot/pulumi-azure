# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetAnalyticsWorkspaceResult',
    'AwaitableGetAnalyticsWorkspaceResult',
    'get_analytics_workspace',
]

@pulumi.output_type
class GetAnalyticsWorkspaceResult:
    """
    A collection of values returned by getAnalyticsWorkspace.
    """
    def __init__(__self__, id=None, location=None, name=None, portal_url=None, primary_shared_key=None, resource_group_name=None, retention_in_days=None, secondary_shared_key=None, sku=None, tags=None, workspace_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if portal_url and not isinstance(portal_url, str):
            raise TypeError("Expected argument 'portal_url' to be a str")
        if portal_url is not None:
            warnings.warn("this property has been removed from the API and will be removed in version 3.0 of the provider", DeprecationWarning)
            pulumi.log.warn("portal_url is deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider")

        pulumi.set(__self__, "portal_url", portal_url)
        if primary_shared_key and not isinstance(primary_shared_key, str):
            raise TypeError("Expected argument 'primary_shared_key' to be a str")
        pulumi.set(__self__, "primary_shared_key", primary_shared_key)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if retention_in_days and not isinstance(retention_in_days, int):
            raise TypeError("Expected argument 'retention_in_days' to be a int")
        pulumi.set(__self__, "retention_in_days", retention_in_days)
        if secondary_shared_key and not isinstance(secondary_shared_key, str):
            raise TypeError("Expected argument 'secondary_shared_key' to be a str")
        pulumi.set(__self__, "secondary_shared_key", secondary_shared_key)
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

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
    @pulumi.getter(name="portalUrl")
    def portal_url(self) -> str:
        return pulumi.get(self, "portal_url")

    @property
    @pulumi.getter(name="primarySharedKey")
    def primary_shared_key(self) -> str:
        """
        The Primary shared key for the Log Analytics Workspace.
        """
        return pulumi.get(self, "primary_shared_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> int:
        """
        The workspace data retention in days.
        """
        return pulumi.get(self, "retention_in_days")

    @property
    @pulumi.getter(name="secondarySharedKey")
    def secondary_shared_key(self) -> str:
        """
        The Secondary shared key for the Log Analytics Workspace.
        """
        return pulumi.get(self, "secondary_shared_key")

    @property
    @pulumi.getter
    def sku(self) -> str:
        """
        The Sku of the Log Analytics Workspace.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        """
        The Workspace (or Customer) ID for the Log Analytics Workspace.
        """
        return pulumi.get(self, "workspace_id")


class AwaitableGetAnalyticsWorkspaceResult(GetAnalyticsWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAnalyticsWorkspaceResult(
            id=self.id,
            location=self.location,
            name=self.name,
            portal_url=self.portal_url,
            primary_shared_key=self.primary_shared_key,
            resource_group_name=self.resource_group_name,
            retention_in_days=self.retention_in_days,
            secondary_shared_key=self.secondary_shared_key,
            sku=self.sku,
            tags=self.tags,
            workspace_id=self.workspace_id)


def get_analytics_workspace(name: Optional[str] = None,
                            resource_group_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAnalyticsWorkspaceResult:
    """
    Use this data source to access information about an existing Log Analytics (formally Operational Insights) Workspace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.operationalinsights.get_analytics_workspace(name="acctest-01",
        resource_group_name="acctest")
    pulumi.export("logAnalyticsWorkspaceId", example.workspace_id)
    ```


    :param str name: Specifies the name of the Log Analytics Workspace.
    :param str resource_group_name: The name of the resource group in which the Log Analytics workspace is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:operationalinsights/getAnalyticsWorkspace:getAnalyticsWorkspace', __args__, opts=opts, typ=GetAnalyticsWorkspaceResult).value

    return AwaitableGetAnalyticsWorkspaceResult(
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        portal_url=__ret__.portal_url,
        primary_shared_key=__ret__.primary_shared_key,
        resource_group_name=__ret__.resource_group_name,
        retention_in_days=__ret__.retention_in_days,
        secondary_shared_key=__ret__.secondary_shared_key,
        sku=__ret__.sku,
        tags=__ret__.tags,
        workspace_id=__ret__.workspace_id)
