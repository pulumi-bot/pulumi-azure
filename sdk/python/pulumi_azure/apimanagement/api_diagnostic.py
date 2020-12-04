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

__all__ = ['ApiDiagnostic']


class ApiDiagnostic(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 always_log_errors: Optional[pulumi.Input[bool]] = None,
                 api_management_logger_id: Optional[pulumi.Input[str]] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_name: Optional[pulumi.Input[str]] = None,
                 backend_request: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticBackendRequestArgs']]] = None,
                 backend_response: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticBackendResponseArgs']]] = None,
                 frontend_request: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendRequestArgs']]] = None,
                 frontend_response: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendResponseArgs']]] = None,
                 http_correlation_protocol: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 log_client_ip: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 verbosity: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a API Management Service API Diagnostics Logs.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_insights = azure.appinsights.Insights("exampleInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_type="web")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="My Company",
            publisher_email="company@mycompany.io",
            sku_name="Developer_1")
        example_api = azure.apimanagement.Api("exampleApi",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            revision="1",
            display_name="Example API",
            path="example",
            protocols=["https"],
            import_=azure.apimanagement.ApiImportArgs(
                content_format="swagger-link-json",
                content_value="http://conferenceapi.azurewebsites.net/?format=json",
            ))
        example_logger = azure.apimanagement.Logger("exampleLogger",
            api_management_name=example_service.name,
            resource_group_name=example_resource_group.name,
            application_insights=azure.apimanagement.LoggerApplicationInsightsArgs(
                instrumentation_key=example_insights.instrumentation_key,
            ))
        example_api_diagnostic = azure.apimanagement.ApiDiagnostic("exampleApiDiagnostic",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            api_name=example_api.name,
            api_management_logger_id=example_logger.id,
            always_log_errors=True,
            log_client_ip=True,
            verbosity="Verbose",
            http_correlation_protocol="W3C",
            frontend_request=azure.apimanagement.ApiDiagnosticFrontendRequestArgs(
                body_bytes=32,
                headers_to_logs=[
                    "content-type",
                    "accept",
                    "origin",
                ],
            ),
            frontend_response=azure.apimanagement.ApiDiagnosticFrontendResponseArgs(
                body_bytes=32,
                headers_to_logs=[
                    "content-type",
                    "content-length",
                    "origin",
                ],
            ),
            backend_request=azure.apimanagement.ApiDiagnosticBackendRequestArgs(
                body_bytes=32,
                headers_to_logs=[
                    "content-type",
                    "accept",
                    "origin",
                ],
            ),
            backend_response=azure.apimanagement.ApiDiagnosticBackendResponseArgs(
                body_bytes=32,
                headers_to_logs=[
                    "content-type",
                    "content-length",
                    "origin",
                ],
            ))
        ```

        ## Import

        API Management Service API Diagnostics Logs can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/apiDiagnostic:ApiDiagnostic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/diagnostics/diagnostic1/loggers/logger1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] always_log_errors: Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
        :param pulumi.Input[str] api_management_logger_id: The ID (name) of the Diagnostics Logger.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[str] api_name: The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticBackendRequestArgs']] backend_request: A `backend_request` block as defined below.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticBackendResponseArgs']] backend_response: A `backend_response` block as defined below.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendRequestArgs']] frontend_request: A `frontend_request` block as defined below.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendResponseArgs']] frontend_response: A `frontend_response` block as defined below.
        :param pulumi.Input[str] http_correlation_protocol: The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
        :param pulumi.Input[str] identifier: Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[bool] log_client_ip: Log client IP address.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[str] verbosity: Logging verbosity. Possible values are `verbose`, `information` or `error`.
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

            __props__['always_log_errors'] = always_log_errors
            if api_management_logger_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_logger_id'")
            __props__['api_management_logger_id'] = api_management_logger_id
            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if api_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_name'")
            __props__['api_name'] = api_name
            __props__['backend_request'] = backend_request
            __props__['backend_response'] = backend_response
            __props__['frontend_request'] = frontend_request
            __props__['frontend_response'] = frontend_response
            __props__['http_correlation_protocol'] = http_correlation_protocol
            if identifier is None and not opts.urn:
                raise TypeError("Missing required property 'identifier'")
            __props__['identifier'] = identifier
            __props__['log_client_ip'] = log_client_ip
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['verbosity'] = verbosity
        super(ApiDiagnostic, __self__).__init__(
            'azure:apimanagement/apiDiagnostic:ApiDiagnostic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            always_log_errors: Optional[pulumi.Input[bool]] = None,
            api_management_logger_id: Optional[pulumi.Input[str]] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            api_name: Optional[pulumi.Input[str]] = None,
            backend_request: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticBackendRequestArgs']]] = None,
            backend_response: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticBackendResponseArgs']]] = None,
            frontend_request: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendRequestArgs']]] = None,
            frontend_response: Optional[pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendResponseArgs']]] = None,
            http_correlation_protocol: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            log_client_ip: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            verbosity: Optional[pulumi.Input[str]] = None) -> 'ApiDiagnostic':
        """
        Get an existing ApiDiagnostic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] always_log_errors: Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
        :param pulumi.Input[str] api_management_logger_id: The ID (name) of the Diagnostics Logger.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[str] api_name: The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticBackendRequestArgs']] backend_request: A `backend_request` block as defined below.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticBackendResponseArgs']] backend_response: A `backend_response` block as defined below.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendRequestArgs']] frontend_request: A `frontend_request` block as defined below.
        :param pulumi.Input[pulumi.InputType['ApiDiagnosticFrontendResponseArgs']] frontend_response: A `frontend_response` block as defined below.
        :param pulumi.Input[str] http_correlation_protocol: The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
        :param pulumi.Input[str] identifier: Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[bool] log_client_ip: Log client IP address.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        :param pulumi.Input[str] verbosity: Logging verbosity. Possible values are `verbose`, `information` or `error`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["always_log_errors"] = always_log_errors
        __props__["api_management_logger_id"] = api_management_logger_id
        __props__["api_management_name"] = api_management_name
        __props__["api_name"] = api_name
        __props__["backend_request"] = backend_request
        __props__["backend_response"] = backend_response
        __props__["frontend_request"] = frontend_request
        __props__["frontend_response"] = frontend_response
        __props__["http_correlation_protocol"] = http_correlation_protocol
        __props__["identifier"] = identifier
        __props__["log_client_ip"] = log_client_ip
        __props__["resource_group_name"] = resource_group_name
        __props__["verbosity"] = verbosity
        return ApiDiagnostic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alwaysLogErrors")
    def always_log_errors(self) -> pulumi.Output[bool]:
        """
        Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
        """
        return pulumi.get(self, "always_log_errors")

    @property
    @pulumi.getter(name="apiManagementLoggerId")
    def api_management_logger_id(self) -> pulumi.Output[str]:
        """
        The ID (name) of the Diagnostics Logger.
        """
        return pulumi.get(self, "api_management_logger_id")

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> pulumi.Output[str]:
        """
        The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        """
        return pulumi.get(self, "api_name")

    @property
    @pulumi.getter(name="backendRequest")
    def backend_request(self) -> pulumi.Output['outputs.ApiDiagnosticBackendRequest']:
        """
        A `backend_request` block as defined below.
        """
        return pulumi.get(self, "backend_request")

    @property
    @pulumi.getter(name="backendResponse")
    def backend_response(self) -> pulumi.Output['outputs.ApiDiagnosticBackendResponse']:
        """
        A `backend_response` block as defined below.
        """
        return pulumi.get(self, "backend_response")

    @property
    @pulumi.getter(name="frontendRequest")
    def frontend_request(self) -> pulumi.Output['outputs.ApiDiagnosticFrontendRequest']:
        """
        A `frontend_request` block as defined below.
        """
        return pulumi.get(self, "frontend_request")

    @property
    @pulumi.getter(name="frontendResponse")
    def frontend_response(self) -> pulumi.Output['outputs.ApiDiagnosticFrontendResponse']:
        """
        A `frontend_response` block as defined below.
        """
        return pulumi.get(self, "frontend_response")

    @property
    @pulumi.getter(name="httpCorrelationProtocol")
    def http_correlation_protocol(self) -> pulumi.Output[str]:
        """
        The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
        """
        return pulumi.get(self, "http_correlation_protocol")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        """
        Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="logClientIp")
    def log_client_ip(self) -> pulumi.Output[bool]:
        """
        Log client IP address.
        """
        return pulumi.get(self, "log_client_ip")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def verbosity(self) -> pulumi.Output[str]:
        """
        Logging verbosity. Possible values are `verbose`, `information` or `error`.
        """
        return pulumi.get(self, "verbosity")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

