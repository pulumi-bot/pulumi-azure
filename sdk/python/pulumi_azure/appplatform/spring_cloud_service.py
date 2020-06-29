# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SpringCloudService(pulumi.CustomResource):
    config_server_git_setting: pulumi.Output[dict]
    """
    A `config_server_git_setting` block as defined below.

      * `httpBasicAuth` (`dict`) - A `http_basic_auth` block as defined below.
        * `password` (`str`) - The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
        * `username` (`str`) - The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.

      * `label` (`str`) - The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
      * `repositories` (`list`) - One or more `repository` blocks as defined below.
        * `httpBasicAuth` (`dict`) - A `http_basic_auth` block as defined below.
          * `password` (`str`) - The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
          * `username` (`str`) - The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.

        * `label` (`str`) - The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
        * `name` (`str`) - A name to identify on the Git repository, required only if repos exists.
        * `patterns` (`list`) - An array of strings used to match an application name. For each pattern, use the `{application}/{profile}` format with wildcards.
        * `searchPaths` (`list`) - An array of strings used to search subdirectories of the Git repository.
        * `sshAuth` (`dict`) - A `ssh_auth` block as defined below.
          * `hostKey` (`str`) - The host key of the Git repository server, should not include the algorithm prefix as covered by `host-key-algorithm`.
          * `hostKeyAlgorithm` (`str`) - The host key algorithm, should be `ssh-dss`, `ssh-rsa`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`. Required only if `host-key` exists.
          * `privateKey` (`str`) - The SSH private key to access the Git repository, required when the URI starts with `git@` or `ssh://`.
          * `strictHostKeyCheckingEnabled` (`bool`) - Indicates whether the Config Server instance will fail to start if the host_key does not match.

        * `uri` (`str`) - The URI of the Git repository that's used as the Config Server back end should be started with `http://`, `https://`, `git@`, or `ssh://`.

      * `searchPaths` (`list`) - An array of strings used to search subdirectories of the Git repository.
      * `sshAuth` (`dict`) - A `ssh_auth` block as defined below.
        * `hostKey` (`str`) - The host key of the Git repository server, should not include the algorithm prefix as covered by `host-key-algorithm`.
        * `hostKeyAlgorithm` (`str`) - The host key algorithm, should be `ssh-dss`, `ssh-rsa`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`. Required only if `host-key` exists.
        * `privateKey` (`str`) - The SSH private key to access the Git repository, required when the URI starts with `git@` or `ssh://`.
        * `strictHostKeyCheckingEnabled` (`bool`) - Indicates whether the Config Server instance will fail to start if the host_key does not match.

      * `uri` (`str`) - The URI of the default Git repository used as the Config Server back end, should be started with `http://`, `https://`, `git@`, or `ssh://`.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, config_server_git_setting=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Spring Cloud Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="Southeast Asia")
        example_spring_cloud_service = azure.appplatform.SpringCloudService("exampleSpringCloudService",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            config_server_git_setting={
                "uri": "https://github.com/Azure-Samples/piggymetrics",
                "label": "config",
                "searchPaths": [
                    "dir1",
                    "dir2",
                ],
            },
            tags={
                "Env": "staging",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] config_server_git_setting: A `config_server_git_setting` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **config_server_git_setting** object supports the following:

          * `httpBasicAuth` (`pulumi.Input[dict]`) - A `http_basic_auth` block as defined below.
            * `password` (`pulumi.Input[str]`) - The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
            * `username` (`pulumi.Input[str]`) - The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.

          * `label` (`pulumi.Input[str]`) - The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
          * `repositories` (`pulumi.Input[list]`) - One or more `repository` blocks as defined below.
            * `httpBasicAuth` (`pulumi.Input[dict]`) - A `http_basic_auth` block as defined below.
              * `password` (`pulumi.Input[str]`) - The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
              * `username` (`pulumi.Input[str]`) - The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.

            * `label` (`pulumi.Input[str]`) - The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
            * `name` (`pulumi.Input[str]`) - A name to identify on the Git repository, required only if repos exists.
            * `patterns` (`pulumi.Input[list]`) - An array of strings used to match an application name. For each pattern, use the `{application}/{profile}` format with wildcards.
            * `searchPaths` (`pulumi.Input[list]`) - An array of strings used to search subdirectories of the Git repository.
            * `sshAuth` (`pulumi.Input[dict]`) - A `ssh_auth` block as defined below.
              * `hostKey` (`pulumi.Input[str]`) - The host key of the Git repository server, should not include the algorithm prefix as covered by `host-key-algorithm`.
              * `hostKeyAlgorithm` (`pulumi.Input[str]`) - The host key algorithm, should be `ssh-dss`, `ssh-rsa`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`. Required only if `host-key` exists.
              * `privateKey` (`pulumi.Input[str]`) - The SSH private key to access the Git repository, required when the URI starts with `git@` or `ssh://`.
              * `strictHostKeyCheckingEnabled` (`pulumi.Input[bool]`) - Indicates whether the Config Server instance will fail to start if the host_key does not match.

            * `uri` (`pulumi.Input[str]`) - The URI of the Git repository that's used as the Config Server back end should be started with `http://`, `https://`, `git@`, or `ssh://`.

          * `searchPaths` (`pulumi.Input[list]`) - An array of strings used to search subdirectories of the Git repository.
          * `sshAuth` (`pulumi.Input[dict]`) - A `ssh_auth` block as defined below.
            * `hostKey` (`pulumi.Input[str]`) - The host key of the Git repository server, should not include the algorithm prefix as covered by `host-key-algorithm`.
            * `hostKeyAlgorithm` (`pulumi.Input[str]`) - The host key algorithm, should be `ssh-dss`, `ssh-rsa`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`. Required only if `host-key` exists.
            * `privateKey` (`pulumi.Input[str]`) - The SSH private key to access the Git repository, required when the URI starts with `git@` or `ssh://`.
            * `strictHostKeyCheckingEnabled` (`pulumi.Input[bool]`) - Indicates whether the Config Server instance will fail to start if the host_key does not match.

          * `uri` (`pulumi.Input[str]`) - The URI of the default Git repository used as the Config Server back end, should be started with `http://`, `https://`, `git@`, or `ssh://`.
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

            __props__['config_server_git_setting'] = config_server_git_setting
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(SpringCloudService, __self__).__init__(
            'azure:appplatform/springCloudService:SpringCloudService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, config_server_git_setting=None, location=None, name=None, resource_group_name=None, tags=None):
        """
        Get an existing SpringCloudService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] config_server_git_setting: A `config_server_git_setting` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **config_server_git_setting** object supports the following:

          * `httpBasicAuth` (`pulumi.Input[dict]`) - A `http_basic_auth` block as defined below.
            * `password` (`pulumi.Input[str]`) - The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
            * `username` (`pulumi.Input[str]`) - The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.

          * `label` (`pulumi.Input[str]`) - The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
          * `repositories` (`pulumi.Input[list]`) - One or more `repository` blocks as defined below.
            * `httpBasicAuth` (`pulumi.Input[dict]`) - A `http_basic_auth` block as defined below.
              * `password` (`pulumi.Input[str]`) - The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
              * `username` (`pulumi.Input[str]`) - The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.

            * `label` (`pulumi.Input[str]`) - The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
            * `name` (`pulumi.Input[str]`) - A name to identify on the Git repository, required only if repos exists.
            * `patterns` (`pulumi.Input[list]`) - An array of strings used to match an application name. For each pattern, use the `{application}/{profile}` format with wildcards.
            * `searchPaths` (`pulumi.Input[list]`) - An array of strings used to search subdirectories of the Git repository.
            * `sshAuth` (`pulumi.Input[dict]`) - A `ssh_auth` block as defined below.
              * `hostKey` (`pulumi.Input[str]`) - The host key of the Git repository server, should not include the algorithm prefix as covered by `host-key-algorithm`.
              * `hostKeyAlgorithm` (`pulumi.Input[str]`) - The host key algorithm, should be `ssh-dss`, `ssh-rsa`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`. Required only if `host-key` exists.
              * `privateKey` (`pulumi.Input[str]`) - The SSH private key to access the Git repository, required when the URI starts with `git@` or `ssh://`.
              * `strictHostKeyCheckingEnabled` (`pulumi.Input[bool]`) - Indicates whether the Config Server instance will fail to start if the host_key does not match.

            * `uri` (`pulumi.Input[str]`) - The URI of the Git repository that's used as the Config Server back end should be started with `http://`, `https://`, `git@`, or `ssh://`.

          * `searchPaths` (`pulumi.Input[list]`) - An array of strings used to search subdirectories of the Git repository.
          * `sshAuth` (`pulumi.Input[dict]`) - A `ssh_auth` block as defined below.
            * `hostKey` (`pulumi.Input[str]`) - The host key of the Git repository server, should not include the algorithm prefix as covered by `host-key-algorithm`.
            * `hostKeyAlgorithm` (`pulumi.Input[str]`) - The host key algorithm, should be `ssh-dss`, `ssh-rsa`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`. Required only if `host-key` exists.
            * `privateKey` (`pulumi.Input[str]`) - The SSH private key to access the Git repository, required when the URI starts with `git@` or `ssh://`.
            * `strictHostKeyCheckingEnabled` (`pulumi.Input[bool]`) - Indicates whether the Config Server instance will fail to start if the host_key does not match.

          * `uri` (`pulumi.Input[str]`) - The URI of the default Git repository used as the Config Server back end, should be started with `http://`, `https://`, `git@`, or `ssh://`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config_server_git_setting"] = config_server_git_setting
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return SpringCloudService(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
