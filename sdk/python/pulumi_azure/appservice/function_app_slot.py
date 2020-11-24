# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FunctionAppSlot']


class FunctionAppSlot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service_plan_id: Optional[pulumi.Input[str]] = None,
                 app_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 auth_settings: Optional[pulumi.Input[pulumi.InputType['FunctionAppSlotAuthSettingsArgs']]] = None,
                 client_affinity_enabled: Optional[pulumi.Input[bool]] = None,
                 connection_strings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionAppSlotConnectionStringArgs']]]]] = None,
                 daily_memory_time_quota: Optional[pulumi.Input[int]] = None,
                 enable_builtin_logging: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 function_app_name: Optional[pulumi.Input[str]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['FunctionAppSlotIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 site_config: Optional[pulumi.Input[pulumi.InputType['FunctionAppSlotSiteConfigArgs']]] = None,
                 storage_account_access_key: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Function App deployment Slot.

        ## Example Usage
        ### With App Service Plan)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westus2")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_plan = azure.appservice.Plan("examplePlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=azure.appservice.PlanSkuArgs(
                tier="Standard",
                size="S1",
            ))
        example_function_app = azure.appservice.FunctionApp("exampleFunctionApp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id,
            storage_account_name=example_account.name,
            storage_account_access_key=example_account.primary_access_key)
        example_function_app_slot = azure.appservice.FunctionAppSlot("exampleFunctionAppSlot",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id,
            function_app_name=example_function_app.name,
            storage_account_name=example_account.name,
            storage_account_access_key=example_account.primary_access_key)
        ```

        ## Import

        Function Apps Deployment Slots can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appservice/functionAppSlot:FunctionAppSlot functionapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1/slots/staging
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_plan_id: The ID of the App Service Plan within which to create this Function App Slot.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] app_settings: A key-value pair of App Settings.
        :param pulumi.Input[pulumi.InputType['FunctionAppSlotAuthSettingsArgs']] auth_settings: An `auth_settings` block as defined below.
        :param pulumi.Input[bool] client_affinity_enabled: Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionAppSlotConnectionStringArgs']]]] connection_strings: A `connection_string` block as defined below.
        :param pulumi.Input[int] daily_memory_time_quota: The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
        :param pulumi.Input[bool] enable_builtin_logging: Should the built-in logging of the Function App be enabled? Defaults to `true`.
        :param pulumi.Input[bool] enabled: Is the Function App enabled?
        :param pulumi.Input[str] function_app_name: The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] https_only: Can the Function App only be accessed via HTTPS? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['FunctionAppSlotIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Function App. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: A string indicating the Operating System type for this function app.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Function App Slot.
        :param pulumi.Input[pulumi.InputType['FunctionAppSlotSiteConfigArgs']] site_config: A `site_config` object as defined below.
        :param pulumi.Input[str] storage_account_access_key: The access key which will be used to access the backend storage account for the Function App.
        :param pulumi.Input[str] storage_account_name: The backend storage account name which will be used by the Function App (such as the dashboard, logs).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] version: The runtime version associated with the Function App. Defaults to `~1`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if app_service_plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_service_plan_id'")
            __props__['app_service_plan_id'] = app_service_plan_id
            __props__['app_settings'] = app_settings
            __props__['auth_settings'] = auth_settings
            __props__['client_affinity_enabled'] = client_affinity_enabled
            __props__['connection_strings'] = connection_strings
            __props__['daily_memory_time_quota'] = daily_memory_time_quota
            __props__['enable_builtin_logging'] = enable_builtin_logging
            __props__['enabled'] = enabled
            if function_app_name is None and not opts.urn:
                raise TypeError("Missing required property 'function_app_name'")
            __props__['function_app_name'] = function_app_name
            __props__['https_only'] = https_only
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['name'] = name
            __props__['os_type'] = os_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['site_config'] = site_config
            if storage_account_access_key is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_access_key'")
            __props__['storage_account_access_key'] = storage_account_access_key
            if storage_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            __props__['tags'] = tags
            __props__['version'] = version
            __props__['default_hostname'] = None
            __props__['kind'] = None
            __props__['outbound_ip_addresses'] = None
            __props__['possible_outbound_ip_addresses'] = None
            __props__['site_credentials'] = None
        super(FunctionAppSlot, __self__).__init__(
            'azure:appservice/functionAppSlot:FunctionAppSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service_plan_id: Optional[pulumi.Input[str]] = None,
            app_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            auth_settings: Optional[pulumi.Input[pulumi.InputType['FunctionAppSlotAuthSettingsArgs']]] = None,
            client_affinity_enabled: Optional[pulumi.Input[bool]] = None,
            connection_strings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionAppSlotConnectionStringArgs']]]]] = None,
            daily_memory_time_quota: Optional[pulumi.Input[int]] = None,
            default_hostname: Optional[pulumi.Input[str]] = None,
            enable_builtin_logging: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            function_app_name: Optional[pulumi.Input[str]] = None,
            https_only: Optional[pulumi.Input[bool]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['FunctionAppSlotIdentityArgs']]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            os_type: Optional[pulumi.Input[str]] = None,
            outbound_ip_addresses: Optional[pulumi.Input[str]] = None,
            possible_outbound_ip_addresses: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            site_config: Optional[pulumi.Input[pulumi.InputType['FunctionAppSlotSiteConfigArgs']]] = None,
            site_credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionAppSlotSiteCredentialArgs']]]]] = None,
            storage_account_access_key: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'FunctionAppSlot':
        """
        Get an existing FunctionAppSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_plan_id: The ID of the App Service Plan within which to create this Function App Slot.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] app_settings: A key-value pair of App Settings.
        :param pulumi.Input[pulumi.InputType['FunctionAppSlotAuthSettingsArgs']] auth_settings: An `auth_settings` block as defined below.
        :param pulumi.Input[bool] client_affinity_enabled: Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionAppSlotConnectionStringArgs']]]] connection_strings: A `connection_string` block as defined below.
        :param pulumi.Input[int] daily_memory_time_quota: The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
        :param pulumi.Input[str] default_hostname: The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
        :param pulumi.Input[bool] enable_builtin_logging: Should the built-in logging of the Function App be enabled? Defaults to `true`.
        :param pulumi.Input[bool] enabled: Is the Function App enabled?
        :param pulumi.Input[str] function_app_name: The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] https_only: Can the Function App only be accessed via HTTPS? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['FunctionAppSlotIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[str] kind: The Function App kind - such as `functionapp,linux,container`
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Function App. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: A string indicating the Operating System type for this function app.
        :param pulumi.Input[str] outbound_ip_addresses: A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        :param pulumi.Input[str] possible_outbound_ip_addresses: A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Function App Slot.
        :param pulumi.Input[pulumi.InputType['FunctionAppSlotSiteConfigArgs']] site_config: A `site_config` object as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionAppSlotSiteCredentialArgs']]]] site_credentials: A `site_credential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
        :param pulumi.Input[str] storage_account_access_key: The access key which will be used to access the backend storage account for the Function App.
        :param pulumi.Input[str] storage_account_name: The backend storage account name which will be used by the Function App (such as the dashboard, logs).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] version: The runtime version associated with the Function App. Defaults to `~1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_service_plan_id"] = app_service_plan_id
        __props__["app_settings"] = app_settings
        __props__["auth_settings"] = auth_settings
        __props__["client_affinity_enabled"] = client_affinity_enabled
        __props__["connection_strings"] = connection_strings
        __props__["daily_memory_time_quota"] = daily_memory_time_quota
        __props__["default_hostname"] = default_hostname
        __props__["enable_builtin_logging"] = enable_builtin_logging
        __props__["enabled"] = enabled
        __props__["function_app_name"] = function_app_name
        __props__["https_only"] = https_only
        __props__["identity"] = identity
        __props__["kind"] = kind
        __props__["location"] = location
        __props__["name"] = name
        __props__["os_type"] = os_type
        __props__["outbound_ip_addresses"] = outbound_ip_addresses
        __props__["possible_outbound_ip_addresses"] = possible_outbound_ip_addresses
        __props__["resource_group_name"] = resource_group_name
        __props__["site_config"] = site_config
        __props__["site_credentials"] = site_credentials
        __props__["storage_account_access_key"] = storage_account_access_key
        __props__["storage_account_name"] = storage_account_name
        __props__["tags"] = tags
        __props__["version"] = version
        return FunctionAppSlot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appServicePlanId")
    def app_service_plan_id(self) -> pulumi.Output[str]:
        """
        The ID of the App Service Plan within which to create this Function App Slot.
        """
        return pulumi.get(self, "app_service_plan_id")

    @property
    @pulumi.getter(name="appSettings")
    def app_settings(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A key-value pair of App Settings.
        """
        return pulumi.get(self, "app_settings")

    @property
    @pulumi.getter(name="authSettings")
    def auth_settings(self) -> pulumi.Output['outputs.FunctionAppSlotAuthSettings']:
        """
        An `auth_settings` block as defined below.
        """
        return pulumi.get(self, "auth_settings")

    @property
    @pulumi.getter(name="clientAffinityEnabled")
    def client_affinity_enabled(self) -> pulumi.Output[bool]:
        """
        Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
        """
        return pulumi.get(self, "client_affinity_enabled")

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> pulumi.Output[Sequence['outputs.FunctionAppSlotConnectionString']]:
        """
        A `connection_string` block as defined below.
        """
        return pulumi.get(self, "connection_strings")

    @property
    @pulumi.getter(name="dailyMemoryTimeQuota")
    def daily_memory_time_quota(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
        """
        return pulumi.get(self, "daily_memory_time_quota")

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> pulumi.Output[str]:
        """
        The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
        """
        return pulumi.get(self, "default_hostname")

    @property
    @pulumi.getter(name="enableBuiltinLogging")
    def enable_builtin_logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Should the built-in logging of the Function App be enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enable_builtin_logging")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the Function App enabled?
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="functionAppName")
    def function_app_name(self) -> pulumi.Output[str]:
        """
        The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "function_app_name")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> pulumi.Output[Optional[bool]]:
        """
        Can the Function App only be accessed via HTTPS? Defaults to `false`.
        """
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output['outputs.FunctionAppSlotIdentity']:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The Function App kind - such as `functionapp,linux,container`
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Function App. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[Optional[str]]:
        """
        A string indicating the Operating System type for this function app.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="outboundIpAddresses")
    def outbound_ip_addresses(self) -> pulumi.Output[str]:
        """
        A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        """
        return pulumi.get(self, "outbound_ip_addresses")

    @property
    @pulumi.getter(name="possibleOutboundIpAddresses")
    def possible_outbound_ip_addresses(self) -> pulumi.Output[str]:
        """
        A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        """
        return pulumi.get(self, "possible_outbound_ip_addresses")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Function App Slot.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="siteConfig")
    def site_config(self) -> pulumi.Output['outputs.FunctionAppSlotSiteConfig']:
        """
        A `site_config` object as defined below.
        """
        return pulumi.get(self, "site_config")

    @property
    @pulumi.getter(name="siteCredentials")
    def site_credentials(self) -> pulumi.Output[Sequence['outputs.FunctionAppSlotSiteCredential']]:
        """
        A `site_credential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
        """
        return pulumi.get(self, "site_credentials")

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> pulumi.Output[str]:
        """
        The access key which will be used to access the backend storage account for the Function App.
        """
        return pulumi.get(self, "storage_account_access_key")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Output[str]:
        """
        The backend storage account name which will be used by the Function App (such as the dashboard, logs).
        """
        return pulumi.get(self, "storage_account_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        The runtime version associated with the Function App. Defaults to `~1`.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

