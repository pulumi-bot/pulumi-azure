# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ActionHttpArgs', 'ActionHttp']

@pulumi.input_type
class ActionHttpArgs:
    def __init__(__self__, *,
                 logic_app_id: pulumi.Input[str],
                 method: pulumi.Input[str],
                 uri: pulumi.Input[str],
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 run_afters: Optional[pulumi.Input[Sequence[pulumi.Input['ActionHttpRunAfterArgs']]]] = None):
        """
        The set of arguments for constructing a ActionHttp resource.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['ActionHttpRunAfterArgs']]] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        """
        pulumi.set(__self__, "logic_app_id", logic_app_id)
        pulumi.set(__self__, "method", method)
        pulumi.set(__self__, "uri", uri)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if run_afters is not None:
            pulumi.set(__self__, "run_afters", run_afters)

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "logic_app_id")

    @logic_app_id.setter
    def logic_app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "logic_app_id", value)

    @property
    @pulumi.getter
    def method(self) -> pulumi.Input[str]:
        """
        Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: pulumi.Input[str]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        """
        Specifies the URI which will be called when this HTTP Action is triggered.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="runAfters")
    def run_afters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ActionHttpRunAfterArgs']]]]:
        """
        Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        """
        return pulumi.get(self, "run_afters")

    @run_afters.setter
    def run_afters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ActionHttpRunAfterArgs']]]]):
        pulumi.set(self, "run_afters", value)


class ActionHttp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 logic_app_id: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 run_afters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]]] = None,
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

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_action_http = azure.logicapps.ActionHttp("exampleActionHttp",
            logic_app_id=example_workflow.id,
            method="GET",
            uri="http://example.com/some-webhook")
        ```

        ## Import

        Logic App HTTP Actions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:logicapps/actionHttp:ActionHttp webhook1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/webhook1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionHttpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an HTTP Action within a Logic App Workflow

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_action_http = azure.logicapps.ActionHttp("exampleActionHttp",
            logic_app_id=example_workflow.id,
            method="GET",
            uri="http://example.com/some-webhook")
        ```

        ## Import

        Logic App HTTP Actions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:logicapps/actionHttp:ActionHttp webhook1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/webhook1
        ```

        :param str resource_name: The name of the resource.
        :param ActionHttpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionHttpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 logic_app_id: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 run_afters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if logic_app_id is None and not opts.urn:
                raise TypeError("Missing required property 'logic_app_id'")
            __props__['logic_app_id'] = logic_app_id
            if method is None and not opts.urn:
                raise TypeError("Missing required property 'method'")
            __props__['method'] = method
            __props__['name'] = name
            __props__['run_afters'] = run_afters
            if uri is None and not opts.urn:
                raise TypeError("Missing required property 'uri'")
            __props__['uri'] = uri
        super(ActionHttp, __self__).__init__(
            'azure:logicapps/actionHttp:ActionHttp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            body: Optional[pulumi.Input[str]] = None,
            headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            logic_app_id: Optional[pulumi.Input[str]] = None,
            method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            run_afters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'ActionHttp':
        """
        Get an existing ActionHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ActionHttpRunAfterArgs']]]] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
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
    def body(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def headers(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "logic_app_id")

    @property
    @pulumi.getter
    def method(self) -> pulumi.Output[str]:
        """
        Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="runAfters")
    def run_afters(self) -> pulumi.Output[Optional[Sequence['outputs.ActionHttpRunAfter']]]:
        """
        Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        """
        return pulumi.get(self, "run_afters")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        Specifies the URI which will be called when this HTTP Action is triggered.
        """
        return pulumi.get(self, "uri")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

