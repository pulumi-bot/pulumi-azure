# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['TemplateDeployment']


class TemplateDeployment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 parameters_body: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 template_body: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a template deployment of resources

        > **Note on ARM Template Deployments:** Due to the way the underlying Azure API is designed, this provider can only manage the deployment of the ARM Template - and not any resources which are created by it.
        This means that when deleting the `core.TemplateDeployment` resource, this provider will only remove the reference to the deployment, whilst leaving any resources created by that ARM Template Deployment.
        One workaround for this is to use a unique Resource Group for each ARM Template Deployment, which means deleting the Resource Group would contain any resources created within it - however this isn't ideal. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).

        ## Note

        This provider does not know about the individual resources created by Azure using a deployment template and therefore cannot delete these resources during a destroy. Destroying a template deployment removes the associated deployment operations, but will not delete the Azure resources created by the deployment. In order to delete these resources, the containing resource group must also be destroyed. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_mode: Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
               Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
               specified within the template, and this provider will not be aware of this.
        :param pulumi.Input[str] name: Specifies the name of the template deployment. Changing this forces a
               new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Specifies the name and value pairs that define the deployment parameters for the template.
        :param pulumi.Input[str] parameters_body: Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the template deployment.
        :param pulumi.Input[str] template_body: Specifies the JSON definition for the template.
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

            if deployment_mode is None:
                raise TypeError("Missing required property 'deployment_mode'")
            __props__['deployment_mode'] = deployment_mode
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['parameters_body'] = parameters_body
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['template_body'] = template_body
            __props__['outputs'] = None
        super(TemplateDeployment, __self__).__init__(
            'azure:core/templateDeployment:TemplateDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deployment_mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            outputs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            parameters_body: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            template_body: Optional[pulumi.Input[str]] = None) -> 'TemplateDeployment':
        """
        Get an existing TemplateDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_mode: Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
               Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
               specified within the template, and this provider will not be aware of this.
        :param pulumi.Input[str] name: Specifies the name of the template deployment. Changing this forces a
               new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] outputs: A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Specifies the name and value pairs that define the deployment parameters for the template.
        :param pulumi.Input[str] parameters_body: Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the template deployment.
        :param pulumi.Input[str] template_body: Specifies the JSON definition for the template.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["deployment_mode"] = deployment_mode
        __props__["name"] = name
        __props__["outputs"] = outputs
        __props__["parameters"] = parameters
        __props__["parameters_body"] = parameters_body
        __props__["resource_group_name"] = resource_group_name
        __props__["template_body"] = template_body
        return TemplateDeployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deploymentMode")
    def deployment_mode(self) -> pulumi.Output[str]:
        """
        Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
        Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
        specified within the template, and this provider will not be aware of this.
        """
        return pulumi.get(self, "deployment_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the template deployment. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the name and value pairs that define the deployment parameters for the template.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="parametersBody")
    def parameters_body(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
        """
        return pulumi.get(self, "parameters_body")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to
        create the template deployment.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> pulumi.Output[str]:
        """
        Specifies the JSON definition for the template.
        """
        return pulumi.get(self, "template_body")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

