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

__all__ = ['ApiOperation']


class ApiOperation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 operation_id: Optional[pulumi.Input[str]] = None,
                 request: Optional[pulumi.Input[pulumi.InputType['ApiOperationRequestArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 responses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationResponseArgs']]]]] = None,
                 template_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationTemplateParameterArgs']]]]] = None,
                 url_template: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an API Operation within an API Management Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_api = azure.apimanagement.get_api(name="search-api",
            api_management_name="search-api-management",
            resource_group_name="search-service",
            revision="2")
        example_api_operation = azure.apimanagement.ApiOperation("exampleApiOperation",
            operation_id="user-delete",
            api_name=example_api.name,
            api_management_name=example_api.api_management_name,
            resource_group_name=example_api.resource_group_name,
            display_name="Delete User Operation",
            method="DELETE",
            url_template="/users/{id}/delete",
            description="This can only be done by the logged in user.",
            responses=[azure.apimanagement.ApiOperationResponseArgs(
                status_code=200,
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_name: The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: A description for this API Operation, which may include HTML formatting tags.
        :param pulumi.Input[str] display_name: The Display Name for this API Management Operation.
        :param pulumi.Input[str] method: The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
        :param pulumi.Input[str] operation_id: A unique identifier for this API Operation. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ApiOperationRequestArgs']] request: A `request` block as defined below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationResponseArgs']]]] responses: One or more `response` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationTemplateParameterArgs']]]] template_parameters: One or more `template_parameter` blocks as defined below.
        :param pulumi.Input[str] url_template: The relative URL Template identifying the target resource for this operation, which may include parameters.
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

            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if api_name is None:
                raise TypeError("Missing required property 'api_name'")
            __props__['api_name'] = api_name
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if method is None:
                raise TypeError("Missing required property 'method'")
            __props__['method'] = method
            if operation_id is None:
                raise TypeError("Missing required property 'operation_id'")
            __props__['operation_id'] = operation_id
            __props__['request'] = request
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['responses'] = responses
            __props__['template_parameters'] = template_parameters
            if url_template is None:
                raise TypeError("Missing required property 'url_template'")
            __props__['url_template'] = url_template
        super(ApiOperation, __self__).__init__(
            'azure:apimanagement/apiOperation:ApiOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            api_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            method: Optional[pulumi.Input[str]] = None,
            operation_id: Optional[pulumi.Input[str]] = None,
            request: Optional[pulumi.Input[pulumi.InputType['ApiOperationRequestArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            responses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationResponseArgs']]]]] = None,
            template_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationTemplateParameterArgs']]]]] = None,
            url_template: Optional[pulumi.Input[str]] = None) -> 'ApiOperation':
        """
        Get an existing ApiOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_name: The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: A description for this API Operation, which may include HTML formatting tags.
        :param pulumi.Input[str] display_name: The Display Name for this API Management Operation.
        :param pulumi.Input[str] method: The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
        :param pulumi.Input[str] operation_id: A unique identifier for this API Operation. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ApiOperationRequestArgs']] request: A `request` block as defined below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationResponseArgs']]]] responses: One or more `response` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiOperationTemplateParameterArgs']]]] template_parameters: One or more `template_parameter` blocks as defined below.
        :param pulumi.Input[str] url_template: The relative URL Template identifying the target resource for this operation, which may include parameters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["api_name"] = api_name
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["method"] = method
        __props__["operation_id"] = operation_id
        __props__["request"] = request
        __props__["resource_group_name"] = resource_group_name
        __props__["responses"] = responses
        __props__["template_parameters"] = template_parameters
        __props__["url_template"] = url_template
        return ApiOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> pulumi.Output[str]:
        """
        The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for this API Operation, which may include HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The Display Name for this API Management Operation.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def method(self) -> pulumi.Output[str]:
        """
        The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter(name="operationId")
    def operation_id(self) -> pulumi.Output[str]:
        """
        A unique identifier for this API Operation. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "operation_id")

    @property
    @pulumi.getter
    def request(self) -> pulumi.Output['outputs.ApiOperationRequest']:
        """
        A `request` block as defined below.
        """
        return pulumi.get(self, "request")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def responses(self) -> pulumi.Output[Optional[Sequence['outputs.ApiOperationResponse']]]:
        """
        One or more `response` blocks as defined below.
        """
        return pulumi.get(self, "responses")

    @property
    @pulumi.getter(name="templateParameters")
    def template_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.ApiOperationTemplateParameter']]]:
        """
        One or more `template_parameter` blocks as defined below.
        """
        return pulumi.get(self, "template_parameters")

    @property
    @pulumi.getter(name="urlTemplate")
    def url_template(self) -> pulumi.Output[str]:
        """
        The relative URL Template identifying the target resource for this operation, which may include parameters.
        """
        return pulumi.get(self, "url_template")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

