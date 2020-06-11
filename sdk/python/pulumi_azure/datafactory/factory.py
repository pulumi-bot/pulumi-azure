# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Factory(pulumi.CustomResource):
    github_configuration: pulumi.Output[dict]
    """
    A `github_configuration` block as defined below.

      * `account_name` (`str`) - Specifies the GitHub account name.
      * `branchName` (`str`) - Specifies the branch of the repository to get code from.
      * `gitUrl` (`str`) - Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
      * `repositoryName` (`str`) - Specifies the name of the git repository.
      * `rootFolder` (`str`) - Specifies the root folder within the repository. Set to `/` for the top level.
    """
    identity: pulumi.Output[dict]
    """
    An `identity` block as defined below.

      * `principal_id` (`str`) - The ID of the Principal (Client) in Azure Active Directory
      * `tenant_id` (`str`) - Specifies the Tenant ID associated with the VSTS account.
      * `type` (`str`) - Specifies the identity type of the Data Factory. At this time the only allowed value is `SystemAssigned`.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Data Factory.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vsts_configuration: pulumi.Output[dict]
    """
    A `vsts_configuration` block as defined below.

      * `account_name` (`str`) - Specifies the VSTS account name.
      * `branchName` (`str`) - Specifies the branch of the repository to get code from.
      * `projectName` (`str`) - Specifies the name of the VSTS project.
      * `repositoryName` (`str`) - Specifies the name of the git repository.
      * `rootFolder` (`str`) - Specifies the root folder within the repository. Set to `/` for the top level.
      * `tenant_id` (`str`) - Specifies the Tenant ID associated with the VSTS account.
    """
    def __init__(__self__, resource_name, opts=None, github_configuration=None, identity=None, location=None, name=None, resource_group_name=None, tags=None, vsts_configuration=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[dict] github_configuration: A `github_configuration` block as defined below.
        :param pulumi.Input[dict] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] vsts_configuration: A `vsts_configuration` block as defined below.

        The **github_configuration** object supports the following:

          * `account_name` (`pulumi.Input[str]`) - Specifies the GitHub account name.
          * `branchName` (`pulumi.Input[str]`) - Specifies the branch of the repository to get code from.
          * `gitUrl` (`pulumi.Input[str]`) - Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
          * `repositoryName` (`pulumi.Input[str]`) - Specifies the name of the git repository.
          * `rootFolder` (`pulumi.Input[str]`) - Specifies the root folder within the repository. Set to `/` for the top level.

        The **identity** object supports the following:

          * `principal_id` (`pulumi.Input[str]`) - The ID of the Principal (Client) in Azure Active Directory
          * `tenant_id` (`pulumi.Input[str]`) - Specifies the Tenant ID associated with the VSTS account.
          * `type` (`pulumi.Input[str]`) - Specifies the identity type of the Data Factory. At this time the only allowed value is `SystemAssigned`.

        The **vsts_configuration** object supports the following:

          * `account_name` (`pulumi.Input[str]`) - Specifies the VSTS account name.
          * `branchName` (`pulumi.Input[str]`) - Specifies the branch of the repository to get code from.
          * `projectName` (`pulumi.Input[str]`) - Specifies the name of the VSTS project.
          * `repositoryName` (`pulumi.Input[str]`) - Specifies the name of the git repository.
          * `rootFolder` (`pulumi.Input[str]`) - Specifies the root folder within the repository. Set to `/` for the top level.
          * `tenant_id` (`pulumi.Input[str]`) - Specifies the Tenant ID associated with the VSTS account.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, github_configuration=None, identity=None, location=None, name=None, resource_group_name=None, tags=None, vsts_configuration=None):
        """
        Get an existing Factory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] github_configuration: A `github_configuration` block as defined below.
        :param pulumi.Input[dict] identity: An `identity` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] vsts_configuration: A `vsts_configuration` block as defined below.

        The **github_configuration** object supports the following:

          * `account_name` (`pulumi.Input[str]`) - Specifies the GitHub account name.
          * `branchName` (`pulumi.Input[str]`) - Specifies the branch of the repository to get code from.
          * `gitUrl` (`pulumi.Input[str]`) - Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
          * `repositoryName` (`pulumi.Input[str]`) - Specifies the name of the git repository.
          * `rootFolder` (`pulumi.Input[str]`) - Specifies the root folder within the repository. Set to `/` for the top level.

        The **identity** object supports the following:

          * `principal_id` (`pulumi.Input[str]`) - The ID of the Principal (Client) in Azure Active Directory
          * `tenant_id` (`pulumi.Input[str]`) - Specifies the Tenant ID associated with the VSTS account.
          * `type` (`pulumi.Input[str]`) - Specifies the identity type of the Data Factory. At this time the only allowed value is `SystemAssigned`.

        The **vsts_configuration** object supports the following:

          * `account_name` (`pulumi.Input[str]`) - Specifies the VSTS account name.
          * `branchName` (`pulumi.Input[str]`) - Specifies the branch of the repository to get code from.
          * `projectName` (`pulumi.Input[str]`) - Specifies the name of the VSTS project.
          * `repositoryName` (`pulumi.Input[str]`) - Specifies the name of the git repository.
          * `rootFolder` (`pulumi.Input[str]`) - Specifies the root folder within the repository. Set to `/` for the top level.
          * `tenant_id` (`pulumi.Input[str]`) - Specifies the Tenant ID associated with the VSTS account.
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
