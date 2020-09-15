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
    'GetFunctionAppResult',
    'AwaitableGetFunctionAppResult',
    'get_function_app',
]

@pulumi.output_type
class GetFunctionAppResult:
    """
    A collection of values returned by getFunctionApp.
    """
    def __init__(__self__, app_service_plan_id=None, app_settings=None, connection_strings=None, default_hostname=None, enabled=None, id=None, location=None, name=None, os_type=None, outbound_ip_addresses=None, possible_outbound_ip_addresses=None, resource_group_name=None, site_configs=None, site_credentials=None, source_controls=None, tags=None):
        if app_service_plan_id and not isinstance(app_service_plan_id, str):
            raise TypeError("Expected argument 'app_service_plan_id' to be a str")
        pulumi.set(__self__, "app_service_plan_id", app_service_plan_id)
        if app_settings and not isinstance(app_settings, dict):
            raise TypeError("Expected argument 'app_settings' to be a dict")
        pulumi.set(__self__, "app_settings", app_settings)
        if connection_strings and not isinstance(connection_strings, list):
            raise TypeError("Expected argument 'connection_strings' to be a list")
        pulumi.set(__self__, "connection_strings", connection_strings)
        if default_hostname and not isinstance(default_hostname, str):
            raise TypeError("Expected argument 'default_hostname' to be a str")
        pulumi.set(__self__, "default_hostname", default_hostname)
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
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if outbound_ip_addresses and not isinstance(outbound_ip_addresses, str):
            raise TypeError("Expected argument 'outbound_ip_addresses' to be a str")
        pulumi.set(__self__, "outbound_ip_addresses", outbound_ip_addresses)
        if possible_outbound_ip_addresses and not isinstance(possible_outbound_ip_addresses, str):
            raise TypeError("Expected argument 'possible_outbound_ip_addresses' to be a str")
        pulumi.set(__self__, "possible_outbound_ip_addresses", possible_outbound_ip_addresses)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if site_configs and not isinstance(site_configs, list):
            raise TypeError("Expected argument 'site_configs' to be a list")
        pulumi.set(__self__, "site_configs", site_configs)
        if site_credentials and not isinstance(site_credentials, list):
            raise TypeError("Expected argument 'site_credentials' to be a list")
        pulumi.set(__self__, "site_credentials", site_credentials)
        if source_controls and not isinstance(source_controls, list):
            raise TypeError("Expected argument 'source_controls' to be a list")
        pulumi.set(__self__, "source_controls", source_controls)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="appServicePlanId")
    def app_service_plan_id(self) -> str:
        """
        The ID of the App Service Plan within which to create this Function App.
        """
        return pulumi.get(self, "app_service_plan_id")

    @property
    @pulumi.getter(name="appSettings")
    def app_settings(self) -> Mapping[str, str]:
        """
        A key-value pair of App Settings.
        """
        return pulumi.get(self, "app_settings")

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> Sequence['outputs.GetFunctionAppConnectionStringResult']:
        """
        An `connection_string` block as defined below.
        """
        return pulumi.get(self, "connection_strings")

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> str:
        """
        The default hostname associated with the Function App.
        """
        return pulumi.get(self, "default_hostname")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Is the Function App enabled?
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
        The name for this IP Restriction.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> str:
        """
        A string indicating the Operating System type for this function app.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="outboundIpAddresses")
    def outbound_ip_addresses(self) -> str:
        """
        A comma separated list of outbound IP addresses.
        """
        return pulumi.get(self, "outbound_ip_addresses")

    @property
    @pulumi.getter(name="possibleOutboundIpAddresses")
    def possible_outbound_ip_addresses(self) -> str:
        """
        A comma separated list of outbound IP addresses, not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        """
        return pulumi.get(self, "possible_outbound_ip_addresses")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="siteConfigs")
    def site_configs(self) -> Sequence['outputs.GetFunctionAppSiteConfigResult']:
        return pulumi.get(self, "site_configs")

    @property
    @pulumi.getter(name="siteCredentials")
    def site_credentials(self) -> Sequence['outputs.GetFunctionAppSiteCredentialResult']:
        """
        A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        """
        return pulumi.get(self, "site_credentials")

    @property
    @pulumi.getter(name="sourceControls")
    def source_controls(self) -> Sequence['outputs.GetFunctionAppSourceControlResult']:
        """
        A `source_control` block as defined below.
        """
        return pulumi.get(self, "source_controls")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")


class AwaitableGetFunctionAppResult(GetFunctionAppResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionAppResult(
            app_service_plan_id=self.app_service_plan_id,
            app_settings=self.app_settings,
            connection_strings=self.connection_strings,
            default_hostname=self.default_hostname,
            enabled=self.enabled,
            id=self.id,
            location=self.location,
            name=self.name,
            os_type=self.os_type,
            outbound_ip_addresses=self.outbound_ip_addresses,
            possible_outbound_ip_addresses=self.possible_outbound_ip_addresses,
            resource_group_name=self.resource_group_name,
            site_configs=self.site_configs,
            site_credentials=self.site_credentials,
            source_controls=self.source_controls,
            tags=self.tags)


def get_function_app(name: Optional[str] = None,
                     resource_group_name: Optional[str] = None,
                     tags: Optional[Mapping[str, str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionAppResult:
    """
    Use this data source to access information about a Function App.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.appservice.get_function_app(name="test-azure-functions",
        resource_group_name=azurerm_resource_group["example"]["name"])
    ```


    :param str name: The name of the Function App resource.
    :param str resource_group_name: The name of the Resource Group where the Function App exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:appservice/getFunctionApp:getFunctionApp', __args__, opts=opts, typ=GetFunctionAppResult).value

    return AwaitableGetFunctionAppResult(
        app_service_plan_id=__ret__.app_service_plan_id,
        app_settings=__ret__.app_settings,
        connection_strings=__ret__.connection_strings,
        default_hostname=__ret__.default_hostname,
        enabled=__ret__.enabled,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        os_type=__ret__.os_type,
        outbound_ip_addresses=__ret__.outbound_ip_addresses,
        possible_outbound_ip_addresses=__ret__.possible_outbound_ip_addresses,
        resource_group_name=__ret__.resource_group_name,
        site_configs=__ret__.site_configs,
        site_credentials=__ret__.site_credentials,
        source_controls=__ret__.source_controls,
        tags=__ret__.tags)
