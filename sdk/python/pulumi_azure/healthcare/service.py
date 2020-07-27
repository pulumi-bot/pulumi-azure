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

__all__ = ['Service']


class Service(pulumi.CustomResource):
    access_policy_object_ids: pulumi.Output[Optional[List[str]]] = pulumi.output_property("accessPolicyObjectIds")
    authentication_configuration: pulumi.Output['outputs.ServiceAuthenticationConfiguration'] = pulumi.output_property("authenticationConfiguration")
    """
    An `authentication_configuration` block as defined below.
    """
    cors_configuration: pulumi.Output['outputs.ServiceCorsConfiguration'] = pulumi.output_property("corsConfiguration")
    """
    A `cors_configuration` block as defined below.
    """
    cosmosdb_throughput: pulumi.Output[Optional[float]] = pulumi.output_property("cosmosdbThroughput")
    """
    The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
    """
    kind: pulumi.Output[Optional[str]] = pulumi.output_property("kind")
    """
    The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the supported Azure Region where the Service should be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the service instance. Used for service endpoint, must be unique within the audience.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the Resource Group in which to create the Service.
    """
    tags: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("tags")
    """
    A mapping of tags to assign to the resource.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, access_policy_object_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, authentication_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceAuthenticationConfigurationArgs']]] = None, cors_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceCorsConfigurationArgs']]] = None, cosmosdb_throughput: Optional[pulumi.Input[float]] = None, kind: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Healthcare Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.healthcare.Service("example",
            access_policy_object_ids=["xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"],
            authentication_configuration={
                "audience": "https://azurehealthcareapis.com/",
                "authority": "https://login.microsoftonline.com/$%7Bdata.azurerm_client_config.current.tenant_id%7D",
                "smartProxyEnabled": "true",
            },
            cors_configuration={
                "allowCredentials": "true",
                "allowedHeaders": [
                    "x-tempo-*",
                    "x-tempo2-*",
                ],
                "allowedMethods": [
                    "GET",
                    "PUT",
                ],
                "allowedOrigins": [
                    "http://www.example.com",
                    "http://www.example2.com",
                ],
                "maxAgeInSeconds": "500",
            },
            cosmosdb_throughput="2000",
            kind="fhir-R4",
            location="westus2",
            resource_group_name="sample-resource-group",
            tags={
                "environment": "testenv",
                "purpose": "AcceptanceTests",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServiceAuthenticationConfigurationArgs']] authentication_configuration: An `authentication_configuration` block as defined below.
        :param pulumi.Input[pulumi.InputType['ServiceCorsConfigurationArgs']] cors_configuration: A `cors_configuration` block as defined below.
        :param pulumi.Input[float] cosmosdb_throughput: The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
        :param pulumi.Input[str] kind: The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
        :param pulumi.Input[str] location: Specifies the supported Azure Region where the Service should be created.
        :param pulumi.Input[str] name: The name of the service instance. Used for service endpoint, must be unique within the audience.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the Service.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['access_policy_object_ids'] = access_policy_object_ids
            __props__['authentication_configuration'] = authentication_configuration
            __props__['cors_configuration'] = cors_configuration
            __props__['cosmosdb_throughput'] = cosmosdb_throughput
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Service, __self__).__init__(
            'azure:healthcare/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, access_policy_object_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, authentication_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceAuthenticationConfigurationArgs']]] = None, cors_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceCorsConfigurationArgs']]] = None, cosmosdb_throughput: Optional[pulumi.Input[float]] = None, kind: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, tags: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServiceAuthenticationConfigurationArgs']] authentication_configuration: An `authentication_configuration` block as defined below.
        :param pulumi.Input[pulumi.InputType['ServiceCorsConfigurationArgs']] cors_configuration: A `cors_configuration` block as defined below.
        :param pulumi.Input[float] cosmosdb_throughput: The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
        :param pulumi.Input[str] kind: The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
        :param pulumi.Input[str] location: Specifies the supported Azure Region where the Service should be created.
        :param pulumi.Input[str] name: The name of the service instance. Used for service endpoint, must be unique within the audience.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the Service.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_policy_object_ids"] = access_policy_object_ids
        __props__["authentication_configuration"] = authentication_configuration
        __props__["cors_configuration"] = cors_configuration
        __props__["cosmosdb_throughput"] = cosmosdb_throughput
        __props__["kind"] = kind
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Service(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

