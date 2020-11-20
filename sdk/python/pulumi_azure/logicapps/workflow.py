# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Workflow']


class Workflow(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 integration_service_environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workflow_schema: Optional[pulumi.Input[str]] = None,
                 workflow_version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Logic App Workflow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="East US")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        ```

        ## Import

        Logic App Workflows can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:logicapps/workflow:Workflow workflow1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_service_environment_id: The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] logic_app_integration_account_id: The ID of the integration account linked by this Logic App Workflow.
        :param pulumi.Input[str] name: Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of Key-Value pairs.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workflow_schema: Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workflow_version: Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
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

            __props__['integration_service_environment_id'] = integration_service_environment_id
            __props__['location'] = location
            __props__['logic_app_integration_account_id'] = logic_app_integration_account_id
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['workflow_schema'] = workflow_schema
            __props__['workflow_version'] = workflow_version
            __props__['access_endpoint'] = None
            __props__['connector_endpoint_ip_addresses'] = None
            __props__['connector_outbound_ip_addresses'] = None
            __props__['workflow_endpoint_ip_addresses'] = None
            __props__['workflow_outbound_ip_addresses'] = None
        super(Workflow, __self__).__init__(
            'azure:logicapps/workflow:Workflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_endpoint: Optional[pulumi.Input[str]] = None,
            connector_endpoint_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            connector_outbound_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            integration_service_environment_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            workflow_endpoint_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            workflow_outbound_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            workflow_schema: Optional[pulumi.Input[str]] = None,
            workflow_version: Optional[pulumi.Input[str]] = None) -> 'Workflow':
        """
        Get an existing Workflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_endpoint: The Access Endpoint for the Logic App Workflow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connector_endpoint_ip_addresses: The list of access endpoint ip addresses of connector.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connector_outbound_ip_addresses: The list of outgoing ip addresses of connector.
        :param pulumi.Input[str] integration_service_environment_id: The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] logic_app_integration_account_id: The ID of the integration account linked by this Logic App Workflow.
        :param pulumi.Input[str] name: Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of Key-Value pairs.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workflow_endpoint_ip_addresses: The list of access endpoint ip addresses of workflow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workflow_outbound_ip_addresses: The list of outgoing ip addresses of workflow.
        :param pulumi.Input[str] workflow_schema: Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workflow_version: Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_endpoint"] = access_endpoint
        __props__["connector_endpoint_ip_addresses"] = connector_endpoint_ip_addresses
        __props__["connector_outbound_ip_addresses"] = connector_outbound_ip_addresses
        __props__["integration_service_environment_id"] = integration_service_environment_id
        __props__["location"] = location
        __props__["logic_app_integration_account_id"] = logic_app_integration_account_id
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["workflow_endpoint_ip_addresses"] = workflow_endpoint_ip_addresses
        __props__["workflow_outbound_ip_addresses"] = workflow_outbound_ip_addresses
        __props__["workflow_schema"] = workflow_schema
        __props__["workflow_version"] = workflow_version
        return Workflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessEndpoint")
    def access_endpoint(self) -> pulumi.Output[str]:
        """
        The Access Endpoint for the Logic App Workflow.
        """
        return pulumi.get(self, "access_endpoint")

    @property
    @pulumi.getter(name="connectorEndpointIpAddresses")
    def connector_endpoint_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of access endpoint ip addresses of connector.
        """
        return pulumi.get(self, "connector_endpoint_ip_addresses")

    @property
    @pulumi.getter(name="connectorOutboundIpAddresses")
    def connector_outbound_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of outgoing ip addresses of connector.
        """
        return pulumi.get(self, "connector_outbound_ip_addresses")

    @property
    @pulumi.getter(name="integrationServiceEnvironmentId")
    def integration_service_environment_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        """
        return pulumi.get(self, "integration_service_environment_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logicAppIntegrationAccountId")
    def logic_app_integration_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the integration account linked by this Logic App Workflow.
        """
        return pulumi.get(self, "logic_app_integration_account_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of Key-Value pairs.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workflowEndpointIpAddresses")
    def workflow_endpoint_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of access endpoint ip addresses of workflow.
        """
        return pulumi.get(self, "workflow_endpoint_ip_addresses")

    @property
    @pulumi.getter(name="workflowOutboundIpAddresses")
    def workflow_outbound_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of outgoing ip addresses of workflow.
        """
        return pulumi.get(self, "workflow_outbound_ip_addresses")

    @property
    @pulumi.getter(name="workflowSchema")
    def workflow_schema(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_schema")

    @property
    @pulumi.getter(name="workflowVersion")
    def workflow_version(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

