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

__all__ = ['ActionHttp']


class ActionHttp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 logic_app_id: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 run_afters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an HTTP Action within a Logic App Workflow

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="East US")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_action_http = azure.logicapps.ActionHttp("exampleActionHttp",
            logic_app_id=example_workflow.id,
            method="GET",
            uri="http://example.com/some-webhook")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.
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

            __props__['body'] = body
            __props__['headers'] = headers
            if logic_app_id is None:
                raise TypeError("Missing required property 'logic_app_id'")
            __props__['logic_app_id'] = logic_app_id
            if method is None:
                raise TypeError("Missing required property 'method'")
            __props__['method'] = method
            __props__['name'] = name
            __props__['run_afters'] = run_afters
            if uri is None:
                raise TypeError("Missing required property 'uri'")
            __props__['uri'] = uri
        super(ActionHttp, __self__).__init__(
            'azure:logicapps/actionHttp:ActionHttp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            body: Optional[pulumi.Input[str]] = None,
            headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            logic_app_id: Optional[pulumi.Input[str]] = None,
            method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            run_afters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'ActionHttp':
        """
        Get an existing ActionHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["body"] = body
        __props__["headers"] = headers
        __props__["logic_app_id"] = logic_app_id
        __props__["method"] = method
        __props__["name"] = name
        __props__["run_afters"] = run_afters
        __props__["uri"] = uri
        return ActionHttp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def body(self) -> Optional[str]:
        """
        Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def headers(self) -> Optional[Mapping[str, str]]:
        """
        Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> str:
        """
        Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "logic_app_id")

    @property
    @pulumi.getter
    def method(self) -> str:
        """
        Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="runAfters")
    def run_afters(self) -> Optional[List['outputs.ActionHttpRunAfter']]:
        """
        Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        """
        return pulumi.get(self, "run_afters")

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        Specifies the URI which will be called when this HTTP Action is triggered.
        """
        return pulumi.get(self, "uri")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

