# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['TriggerCustom']


class TriggerCustom(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 logic_app_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Custom Trigger within a Logic App Workflow

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="East US")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_trigger_custom = azure.logicapps.TriggerCustom("exampleTriggerCustom",
            logic_app_id=example_workflow.id,
            body=\"\"\"{
          "recurrence": {
            "frequency": "Day",
            "interval": 1
          },
          "type": "Recurrence"
        }
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the JSON Blob defining the Body of this Custom Trigger.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
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

            if body is None:
                raise TypeError("Missing required property 'body'")
            __props__['body'] = body
            if logic_app_id is None:
                raise TypeError("Missing required property 'logic_app_id'")
            __props__['logic_app_id'] = logic_app_id
            __props__['name'] = name
        super(TriggerCustom, __self__).__init__(
            'azure:logicapps/triggerCustom:TriggerCustom',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            body: Optional[pulumi.Input[str]] = None,
            logic_app_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'TriggerCustom':
        """
        Get an existing TriggerCustom resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the JSON Blob defining the Body of this Custom Trigger.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["body"] = body
        __props__["logic_app_id"] = logic_app_id
        __props__["name"] = name
        return TriggerCustom(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def body(self) -> str:
        """
        Specifies the JSON Blob defining the Body of this Custom Trigger.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> str:
        """
        Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "logic_app_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

