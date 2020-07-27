# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ActionHttp(pulumi.CustomResource):
    body: pulumi.Output[str]
    """
    Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
    """
    headers: pulumi.Output[dict]
    """
    Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
    """
    logic_app_id: pulumi.Output[str]
    """
    Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
    """
    method: pulumi.Output[str]
    """
    Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
    """
    run_afters: pulumi.Output[list]
    """
    Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.

      * `actionName` (`str`) - Specifies the name of the precedent HTTP Action.
      * `actionResult` (`str`) - Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
    """
    uri: pulumi.Output[str]
    """
    Specifies the URI which will be called when this HTTP Action is triggered.
    """
    def __init__(__self__, resource_name, opts=None, body=None, headers=None, logic_app_id=None, method=None, name=None, run_afters=None, uri=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[dict] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[list] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.

        The **run_afters** object supports the following:

          * `actionName` (`pulumi.Input[str]`) - Specifies the name of the precedent HTTP Action.
          * `actionResult` (`pulumi.Input[str]`) - Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
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
    def get(resource_name, id, opts=None, body=None, headers=None, logic_app_id=None, method=None, name=None, run_afters=None, uri=None):
        """
        Get an existing ActionHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[dict] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[list] run_afters: Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `run_after` block is as defined below.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.

        The **run_afters** object supports the following:

          * `actionName` (`pulumi.Input[str]`) - Specifies the name of the precedent HTTP Action.
          * `actionResult` (`pulumi.Input[str]`) - Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
