# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 data_factory_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 activities_json: Optional[pulumi.Input[str]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
        :param pulumi.Input[str] activities_json: A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Pipeline.
        :param pulumi.Input[str] description: The description for the Data Factory Pipeline.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Pipeline.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of variables to associate with the Data Factory Pipeline.
        """
        pulumi.set(__self__, "data_factory_name", data_factory_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if activities_json is not None:
            pulumi.set(__self__, "activities_json", activities_json)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Input[str]:
        """
        The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="activitiesJson")
    def activities_json(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
        """
        return pulumi.get(self, "activities_json")

    @activities_json.setter
    def activities_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "activities_json", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Pipeline.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Data Factory Pipeline.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of parameters to associate with the Data Factory Pipeline.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of variables to associate with the Data Factory Pipeline.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "variables", value)


class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Pipeline inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pipeline = azure.datafactory.Pipeline("examplePipeline",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name)
        ```
        ### With Activities

        ```python
        import pulumi
        import pulumi_azure as azure

        test = azure.datafactory.Pipeline("test",
            resource_group_name=azurerm_resource_group["test"]["name"],
            data_factory_name=azurerm_data_factory["test"]["name"],
            variables={
                "bob": "item1",
            },
            activities_json=\"\"\"[
        	{
        		"name": "Append variable1",
        		"type": "AppendVariable",
        		"dependsOn": [],
        		"userProperties": [],
        		"typeProperties": {
        			"variableName": "bob",
        			"value": "something"
        		}
        	}
        ]
        \"\"\")
        ```

        ## Import

        Data Factory Pipeline's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/pipeline:Pipeline example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
        ```

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activities_json: Optional[pulumi.Input[str]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Pipeline inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pipeline = azure.datafactory.Pipeline("examplePipeline",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name)
        ```
        ### With Activities

        ```python
        import pulumi
        import pulumi_azure as azure

        test = azure.datafactory.Pipeline("test",
            resource_group_name=azurerm_resource_group["test"]["name"],
            data_factory_name=azurerm_data_factory["test"]["name"],
            variables={
                "bob": "item1",
            },
            activities_json=\"\"\"[
        	{
        		"name": "Append variable1",
        		"type": "AppendVariable",
        		"dependsOn": [],
        		"userProperties": [],
        		"typeProperties": {
        			"variableName": "bob",
        			"value": "something"
        		}
        	}
        ]
        \"\"\")
        ```

        ## Import

        Data Factory Pipeline's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/pipeline:Pipeline example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] activities_json: A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Pipeline.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Pipeline.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Pipeline.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of variables to associate with the Data Factory Pipeline.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activities_json: Optional[pulumi.Input[str]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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

            __props__['activities_json'] = activities_json
            __props__['annotations'] = annotations
            if data_factory_name is None and not opts.urn:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['variables'] = variables
        super(Pipeline, __self__).__init__(
            'azure:datafactory/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activities_json: Optional[pulumi.Input[str]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] activities_json: A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Pipeline.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Pipeline.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Pipeline.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map of variables to associate with the Data Factory Pipeline.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["activities_json"] = activities_json
        __props__["annotations"] = annotations
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["variables"] = variables
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activitiesJson")
    def activities_json(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
        """
        return pulumi.get(self, "activities_json")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Pipeline.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Output[str]:
        """
        The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the Data Factory Pipeline.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of parameters to associate with the Data Factory Pipeline.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of variables to associate with the Data Factory Pipeline.
        """
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

