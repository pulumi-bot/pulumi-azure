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

__all__ = ['Factory']


class Factory(pulumi.CustomResource):
    github_configuration: pulumi.Output[Optional['outputs.FactoryGithubConfiguration']] = pulumi.output_property("githubConfiguration")
    """
    A `github_configuration` block as defined below.
    """
    identity: pulumi.Output['outputs.FactoryIdentity'] = pulumi.output_property("identity")
    """
    An `identity` block as defined below.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which to create the Data Factory.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    vsts_configuration: pulumi.Output[Optional['outputs.FactoryVstsConfiguration']] = pulumi.output_property("vstsConfiguration")
    """
    A `vsts_configuration` block as defined below.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, github_configuration: Optional[pulumi.Input[pulumi.InputType['FactoryGithubConfigurationArgs']]] = None, identity: Optional[pulumi.Input[pulumi.InputType['FactoryIdentityArgs']]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, vsts_configuration: Optional[pulumi.Input[pulumi.InputType['FactoryVstsConfigurationArgs']]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an Azure Data Factory (Version 2).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FactoryGithubConfigurationArgs']] github_configuration: A `github_configuration` block as defined below.
        :param pulumi.Input[pulumi.InputType['FactoryIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['FactoryVstsConfigurationArgs']] vsts_configuration: A `vsts_configuration` block as defined below.
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

            __props__['github_configuration'] = github_configuration
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['vsts_configuration'] = vsts_configuration
        super(Factory, __self__).__init__(
            'azure:datafactory/factory:Factory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, github_configuration: Optional[pulumi.Input[pulumi.InputType['FactoryGithubConfigurationArgs']]] = None, identity: Optional[pulumi.Input[pulumi.InputType['FactoryIdentityArgs']]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, vsts_configuration: Optional[pulumi.Input[pulumi.InputType['FactoryVstsConfigurationArgs']]] = None) -> 'Factory':
        """
        Get an existing Factory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FactoryGithubConfigurationArgs']] github_configuration: A `github_configuration` block as defined below.
        :param pulumi.Input[pulumi.InputType['FactoryIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['FactoryVstsConfigurationArgs']] vsts_configuration: A `vsts_configuration` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["github_configuration"] = github_configuration
        __props__["identity"] = identity
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["vsts_configuration"] = vsts_configuration
        return Factory(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

