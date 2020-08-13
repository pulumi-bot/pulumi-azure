# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetFunctionAppResult',
    'AwaitableGetFunctionAppResult',
    'get_function_app',
]


@pulumi.output_type
class _GetFunctionAppResult(dict):
    app_service_plan_id: str = pulumi.property("appServicePlanId")
    app_settings: Mapping[str, str] = pulumi.property("appSettings")
    connection_strings: List['outputs.GetFunctionAppConnectionStringResult'] = pulumi.property("connectionStrings")
    default_hostname: str = pulumi.property("defaultHostname")
    enabled: bool = pulumi.property("enabled")
    id: str = pulumi.property("id")
    location: str = pulumi.property("location")
    name: str = pulumi.property("name")
    os_type: str = pulumi.property("osType")
    outbound_ip_addresses: str = pulumi.property("outboundIpAddresses")
    possible_outbound_ip_addresses: str = pulumi.property("possibleOutboundIpAddresses")
    resource_group_name: str = pulumi.property("resourceGroupName")
    site_credentials: List['outputs.GetFunctionAppSiteCredentialResult'] = pulumi.property("siteCredentials")
    tags: Optional[Mapping[str, str]] = pulumi.property("tags")


class GetFunctionAppResult:
    """
    A collection of values returned by getFunctionApp.
    """
    def __init__(__self__, app_service_plan_id=None, app_settings=None, connection_strings=None, default_hostname=None, enabled=None, id=None, location=None, name=None, os_type=None, outbound_ip_addresses=None, possible_outbound_ip_addresses=None, resource_group_name=None, site_credentials=None, tags=None):
        if app_service_plan_id and not isinstance(app_service_plan_id, str):
            raise TypeError("Expected argument 'app_service_plan_id' to be a str")
        __self__.app_service_plan_id = app_service_plan_id
        """
        The ID of the App Service Plan within which to create this Function App.
        """
        if app_settings and not isinstance(app_settings, dict):
            raise TypeError("Expected argument 'app_settings' to be a dict")
        __self__.app_settings = app_settings
        """
        A key-value pair of App Settings.
        """
        if connection_strings and not isinstance(connection_strings, list):
            raise TypeError("Expected argument 'connection_strings' to be a list")
        __self__.connection_strings = connection_strings
        """
        An `connection_string` block as defined below.
        """
        if default_hostname and not isinstance(default_hostname, str):
            raise TypeError("Expected argument 'default_hostname' to be a str")
        __self__.default_hostname = default_hostname
        """
        The default hostname associated with the Function App.
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        """
        Is the Function App enabled?
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the Connection String.
        """
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        __self__.os_type = os_type
        """
        A string indicating the Operating System type for this function app.
        """
        if outbound_ip_addresses and not isinstance(outbound_ip_addresses, str):
            raise TypeError("Expected argument 'outbound_ip_addresses' to be a str")
        __self__.outbound_ip_addresses = outbound_ip_addresses
        """
        A comma separated list of outbound IP addresses.
        """
        if possible_outbound_ip_addresses and not isinstance(possible_outbound_ip_addresses, str):
            raise TypeError("Expected argument 'possible_outbound_ip_addresses' to be a str")
        __self__.possible_outbound_ip_addresses = possible_outbound_ip_addresses
        """
        A comma separated list of outbound IP addresses, not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if site_credentials and not isinstance(site_credentials, list):
            raise TypeError("Expected argument 'site_credentials' to be a list")
        __self__.site_credentials = site_credentials
        """
        A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags


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
            site_credentials=self.site_credentials,
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
    __ret__ = pulumi.runtime.invoke('azure:appservice/getFunctionApp:getFunctionApp', __args__, opts=opts, typ=_GetFunctionAppResult).value

    return AwaitableGetFunctionAppResult(
        app_service_plan_id=_utilities.get_dict_value(__ret__, 'appServicePlanId'),
        app_settings=_utilities.get_dict_value(__ret__, 'appSettings'),
        connection_strings=_utilities.get_dict_value(__ret__, 'connectionStrings'),
        default_hostname=_utilities.get_dict_value(__ret__, 'defaultHostname'),
        enabled=_utilities.get_dict_value(__ret__, 'enabled'),
        id=_utilities.get_dict_value(__ret__, 'id'),
        location=_utilities.get_dict_value(__ret__, 'location'),
        name=_utilities.get_dict_value(__ret__, 'name'),
        os_type=_utilities.get_dict_value(__ret__, 'osType'),
        outbound_ip_addresses=_utilities.get_dict_value(__ret__, 'outboundIpAddresses'),
        possible_outbound_ip_addresses=_utilities.get_dict_value(__ret__, 'possibleOutboundIpAddresses'),
        resource_group_name=_utilities.get_dict_value(__ret__, 'resourceGroupName'),
        site_credentials=_utilities.get_dict_value(__ret__, 'siteCredentials'),
        tags=_utilities.get_dict_value(__ret__, 'tags'))
