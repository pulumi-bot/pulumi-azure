# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FunctionJavaScriptUDF']


class FunctionJavaScriptUDF(pulumi.CustomResource):
    inputs: pulumi.Output[List['outputs.FunctionJavaScriptUDFInput']] = pulumi.output_property("inputs")
    """
    One or more `input` blocks as defined below.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
    """
    output: pulumi.Output['outputs.FunctionJavaScriptUDFOutput'] = pulumi.output_property("output")
    """
    An `output` blocks as defined below.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
    """
    script: pulumi.Output[str] = pulumi.output_property("script")
    """
    The JavaScript of this UDF Function.
    """
    stream_analytics_job_name: pulumi.Output[str] = pulumi.output_property("streamAnalyticsJobName")
    """
    The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, inputs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFInputArgs']]]]] = None, name: Optional[pulumi.Input[str]] = None, output: Optional[pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFOutputArgs']]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, script: Optional[pulumi.Input[str]] = None, stream_analytics_job_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a JavaScript UDF Function within Stream Analytics Streaming Job.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_function_java_script_udf = azure.streamanalytics.FunctionJavaScriptUDF("exampleFunctionJavaScriptUDF",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            script=\"\"\"function getRandomNumber(in) {
          return in;
        }
        \"\"\",
            inputs=[{
                "type": "bigint",
            }],
            output={
                "type": "bigint",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFInputArgs']]]] inputs: One or more `input` blocks as defined below.
        :param pulumi.Input[str] name: The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFOutputArgs']] output: An `output` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] script: The JavaScript of this UDF Function.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
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

            if inputs is None:
                raise TypeError("Missing required property 'inputs'")
            __props__['inputs'] = inputs
            __props__['name'] = name
            if output is None:
                raise TypeError("Missing required property 'output'")
            __props__['output'] = output
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if script is None:
                raise TypeError("Missing required property 'script'")
            __props__['script'] = script
            if stream_analytics_job_name is None:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
        super(FunctionJavaScriptUDF, __self__).__init__(
            'azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, inputs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFInputArgs']]]]] = None, name: Optional[pulumi.Input[str]] = None, output: Optional[pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFOutputArgs']]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, script: Optional[pulumi.Input[str]] = None, stream_analytics_job_name: Optional[pulumi.Input[str]] = None) -> 'FunctionJavaScriptUDF':
        """
        Get an existing FunctionJavaScriptUDF resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFInputArgs']]]] inputs: One or more `input` blocks as defined below.
        :param pulumi.Input[str] name: The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['FunctionJavaScriptUDFOutputArgs']] output: An `output` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] script: The JavaScript of this UDF Function.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["inputs"] = inputs
        __props__["name"] = name
        __props__["output"] = output
        __props__["resource_group_name"] = resource_group_name
        __props__["script"] = script
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        return FunctionJavaScriptUDF(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

