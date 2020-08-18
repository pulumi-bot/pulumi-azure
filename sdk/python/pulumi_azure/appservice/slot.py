# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Slot']


class Slot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service_name: Optional[pulumi.Input[str]] = None,
                 app_service_plan_id: Optional[pulumi.Input[str]] = None,
                 app_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 auth_settings: Optional[pulumi.Input[pulumi.InputType['SlotAuthSettingsArgs']]] = None,
                 client_affinity_enabled: Optional[pulumi.Input[bool]] = None,
                 connection_strings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SlotConnectionStringArgs']]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['SlotIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logs: Optional[pulumi.Input[pulumi.InputType['SlotLogsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 site_config: Optional[pulumi.Input[pulumi.InputType['SlotSiteConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an App Service Slot (within an App Service).

        > **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.

        ## Example Usage
        ### Net 4.X)

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        server = random.RandomId("server",
            keepers={
                "azi_id": 1,
            },
            byte_length=8)
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_plan = azure.appservice.Plan("examplePlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku={
                "tier": "Standard",
                "size": "S1",
            })
        example_app_service = azure.appservice.AppService("exampleAppService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id,
            site_config={
                "dotnetFrameworkVersion": "v4.0",
            },
            app_settings={
                "SOME_KEY": "some-value",
            },
            connection_strings=[{
                "name": "Database",
                "type": "SQLServer",
                "value": "Server=some-server.mydomain.com;Integrated Security=SSPI",
            }])
        example_slot = azure.appservice.Slot("exampleSlot",
            app_service_name=example_app_service.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id,
            site_config={
                "dotnetFrameworkVersion": "v4.0",
            },
            app_settings={
                "SOME_KEY": "some-value",
            },
            connection_strings=[{
                "name": "Database",
                "type": "SQLServer",
                "value": "Server=some-server.mydomain.com;Integrated Security=SSPI",
            }])
        ```
        ### Java 1.8)

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        server = random.RandomId("server",
            keepers={
                "azi_id": 1,
            },
            byte_length=8)
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_plan = azure.appservice.Plan("examplePlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku={
                "tier": "Standard",
                "size": "S1",
            })
        example_app_service = azure.appservice.AppService("exampleAppService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id,
            site_config={
                "javaVersion": "1.8",
                "javaContainer": "JETTY",
                "javaContainerVersion": "9.3",
            })
        example_slot = azure.appservice.Slot("exampleSlot",
            app_service_name=example_app_service.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id,
            site_config={
                "javaVersion": "1.8",
                "javaContainer": "JETTY",
                "javaContainerVersion": "9.3",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_name: The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_service_plan_id: The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] app_settings: A key-value pair of App Settings.
        :param pulumi.Input[pulumi.InputType['SlotAuthSettingsArgs']] auth_settings: A `auth_settings` block as defined below.
        :param pulumi.Input[bool] client_affinity_enabled: Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SlotConnectionStringArgs']]]] connection_strings: An `connection_string` block as defined below.
        :param pulumi.Input[bool] enabled: Is the App Service Slot Enabled?
        :param pulumi.Input[bool] https_only: Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['SlotIdentityArgs']] identity: A Managed Service Identity block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the App Service Slot component.
        :param pulumi.Input[pulumi.InputType['SlotSiteConfigArgs']] site_config: A `site_config` object as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            if app_service_name is None:
                raise TypeError("Missing required property 'app_service_name'")
            __props__['app_service_name'] = app_service_name
            if app_service_plan_id is None:
                raise TypeError("Missing required property 'app_service_plan_id'")
            __props__['app_service_plan_id'] = app_service_plan_id
            __props__['app_settings'] = app_settings
            __props__['auth_settings'] = auth_settings
            __props__['client_affinity_enabled'] = client_affinity_enabled
            __props__['connection_strings'] = connection_strings
            __props__['enabled'] = enabled
            __props__['https_only'] = https_only
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['logs'] = logs
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['site_config'] = site_config
            __props__['tags'] = tags
            __props__['default_site_hostname'] = None
            __props__['site_credentials'] = None
        super(Slot, __self__).__init__(
            'azure:appservice/slot:Slot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service_name: Optional[pulumi.Input[str]] = None,
            app_service_plan_id: Optional[pulumi.Input[str]] = None,
            app_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            auth_settings: Optional[pulumi.Input[pulumi.InputType['SlotAuthSettingsArgs']]] = None,
            client_affinity_enabled: Optional[pulumi.Input[bool]] = None,
            connection_strings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SlotConnectionStringArgs']]]]] = None,
            default_site_hostname: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            https_only: Optional[pulumi.Input[bool]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['SlotIdentityArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            logs: Optional[pulumi.Input[pulumi.InputType['SlotLogsArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            site_config: Optional[pulumi.Input[pulumi.InputType['SlotSiteConfigArgs']]] = None,
            site_credentials: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SlotSiteCredentialArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Slot':
        """
        Get an existing Slot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_name: The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_service_plan_id: The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] app_settings: A key-value pair of App Settings.
        :param pulumi.Input[pulumi.InputType['SlotAuthSettingsArgs']] auth_settings: A `auth_settings` block as defined below.
        :param pulumi.Input[bool] client_affinity_enabled: Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SlotConnectionStringArgs']]]] connection_strings: An `connection_string` block as defined below.
        :param pulumi.Input[str] default_site_hostname: The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
        :param pulumi.Input[bool] enabled: Is the App Service Slot Enabled?
        :param pulumi.Input[bool] https_only: Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['SlotIdentityArgs']] identity: A Managed Service Identity block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the App Service Slot component.
        :param pulumi.Input[pulumi.InputType['SlotSiteConfigArgs']] site_config: A `site_config` object as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SlotSiteCredentialArgs']]]] site_credentials: A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_service_name"] = app_service_name
        __props__["app_service_plan_id"] = app_service_plan_id
        __props__["app_settings"] = app_settings
        __props__["auth_settings"] = auth_settings
        __props__["client_affinity_enabled"] = client_affinity_enabled
        __props__["connection_strings"] = connection_strings
        __props__["default_site_hostname"] = default_site_hostname
        __props__["enabled"] = enabled
        __props__["https_only"] = https_only
        __props__["identity"] = identity
        __props__["location"] = location
        __props__["logs"] = logs
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["site_config"] = site_config
        __props__["site_credentials"] = site_credentials
        __props__["tags"] = tags
        return Slot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appServiceName")
    def app_service_name(self) -> str:
        """
        The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_service_name")

    @property
    @pulumi.getter(name="appServicePlanId")
    def app_service_plan_id(self) -> str:
        """
        The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
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
    @pulumi.getter(name="authSettings")
    def auth_settings(self) -> 'outputs.SlotAuthSettings':
        """
        A `auth_settings` block as defined below.
        """
        return pulumi.get(self, "auth_settings")

    @property
    @pulumi.getter(name="clientAffinityEnabled")
    def client_affinity_enabled(self) -> bool:
        """
        Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
        """
        return pulumi.get(self, "client_affinity_enabled")

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> List['outputs.SlotConnectionString']:
        """
        An `connection_string` block as defined below.
        """
        return pulumi.get(self, "connection_strings")

    @property
    @pulumi.getter(name="defaultSiteHostname")
    def default_site_hostname(self) -> str:
        """
        The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
        """
        return pulumi.get(self, "default_site_hostname")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Is the App Service Slot Enabled?
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[bool]:
        """
        Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
        """
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter
    def identity(self) -> 'outputs.SlotIdentity':
        """
        A Managed Service Identity block as defined below.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def logs(self) -> 'outputs.SlotLogs':
        return pulumi.get(self, "logs")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the App Service Slot component.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="siteConfig")
    def site_config(self) -> 'outputs.SlotSiteConfig':
        """
        A `site_config` object as defined below.
        """
        return pulumi.get(self, "site_config")

    @property
    @pulumi.getter(name="siteCredentials")
    def site_credentials(self) -> List['outputs.SlotSiteCredential']:
        """
        A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        """
        return pulumi.get(self, "site_credentials")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

