# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Endpoint']

warnings.warn("azure.trafficmanager.Endpoint has been deprecated in favor of azure.network.TrafficManagerEndpoint", DeprecationWarning)


class Endpoint(pulumi.CustomResource):
    warnings.warn("azure.trafficmanager.Endpoint has been deprecated in favor of azure.network.TrafficManagerEndpoint", DeprecationWarning)

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_headers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointCustomHeaderArgs']]]]] = None,
                 endpoint_location: Optional[pulumi.Input[str]] = None,
                 endpoint_status: Optional[pulumi.Input[str]] = None,
                 geo_mappings: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 min_child_endpoints: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointSubnetArgs']]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Traffic Manager Endpoint.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        server = random.RandomId("server",
            keepers={
                "azi_id": 1,
            },
            byte_length=8)
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_traffic_manager_profile = azure.network.TrafficManagerProfile("exampleTrafficManagerProfile",
            resource_group_name=example_resource_group.name,
            traffic_routing_method="Weighted",
            dns_config={
                "relativeName": server.hex,
                "ttl": 100,
            },
            monitor_config={
                "protocol": "http",
                "port": 80,
                "path": "/",
                "interval_in_seconds": 30,
                "timeoutInSeconds": 9,
                "toleratedNumberOfFailures": 3,
            },
            tags={
                "environment": "Production",
            })
        example_traffic_manager_endpoint = azure.network.TrafficManagerEndpoint("exampleTrafficManagerEndpoint",
            resource_group_name=example_resource_group.name,
            profile_name=example_traffic_manager_profile.name,
            type="externalEndpoints",
            weight=100)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointCustomHeaderArgs']]]] custom_headers: One or more `custom_header` blocks as defined below
        :param pulumi.Input[str] endpoint_location: Specifies the Azure location of the Endpoint,
               this must be specified for Profiles using the `Performance` routing method
               if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
               For Endpoints of type `azureEndpoints` the value will be taken from the
               location of the Azure target resource.
        :param pulumi.Input[str] endpoint_status: The status of the Endpoint, can be set to
               either `Enabled` or `Disabled`. Defaults to `Enabled`.
        :param pulumi.Input[List[pulumi.Input[str]]] geo_mappings: A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
        :param pulumi.Input[float] min_child_endpoints: This argument specifies the minimum number
               of endpoints that must be ‘online’ in the child profile in order for the
               parent profile to direct traffic to any of the endpoints in that child
               profile. This argument only applies to Endpoints of type `nestedEndpoints`
               and defaults to `1`.
        :param pulumi.Input[str] name: The name of the Traffic Manager endpoint. Changing this forces a
               new resource to be created.
        :param pulumi.Input[float] priority: Specifies the priority of this Endpoint, this must be
               specified for Profiles using the `Priority` traffic routing method. Supports
               values between 1 and 1000, with no Endpoints sharing the same value. If
               omitted the value will be computed in order of creation.
        :param pulumi.Input[str] profile_name: The name of the Traffic Manager Profile to attach
               create the Traffic Manager endpoint.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the Traffic Manager Profile exists.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointSubnetArgs']]]] subnets: One or more `subnet` blocks as defined below
        :param pulumi.Input[str] target: The FQDN DNS name of the target. This argument must be
               provided for an endpoint of type `externalEndpoints`, for other types it
               will be computed.
        :param pulumi.Input[str] target_resource_id: The resource id of an Azure resource to
               target. This argument must be provided for an endpoint of type
               `azureEndpoints` or `nestedEndpoints`.
        :param pulumi.Input[str] type: The Endpoint type, must be one of:
               - `azureEndpoints`
               - `externalEndpoints`
               - `nestedEndpoints`
        :param pulumi.Input[float] weight: Specifies how much traffic should be distributed to this
               endpoint, this must be specified for Profiles using the  `Weighted` traffic
               routing method. Supports values between 1 and 1000.
        """
        pulumi.log.warn("Endpoint is deprecated: azure.trafficmanager.Endpoint has been deprecated in favor of azure.network.TrafficManagerEndpoint")
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

            __props__['custom_headers'] = custom_headers
            __props__['endpoint_location'] = endpoint_location
            __props__['endpoint_status'] = endpoint_status
            __props__['geo_mappings'] = geo_mappings
            __props__['min_child_endpoints'] = min_child_endpoints
            __props__['name'] = name
            __props__['priority'] = priority
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['subnets'] = subnets
            __props__['target'] = target
            __props__['target_resource_id'] = target_resource_id
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['weight'] = weight
            __props__['endpoint_monitor_status'] = None
        super(Endpoint, __self__).__init__(
            'azure:trafficmanager/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_headers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointCustomHeaderArgs']]]]] = None,
            endpoint_location: Optional[pulumi.Input[str]] = None,
            endpoint_monitor_status: Optional[pulumi.Input[str]] = None,
            endpoint_status: Optional[pulumi.Input[str]] = None,
            geo_mappings: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            min_child_endpoints: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            profile_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            subnets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointSubnetArgs']]]]] = None,
            target: Optional[pulumi.Input[str]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[float]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointCustomHeaderArgs']]]] custom_headers: One or more `custom_header` blocks as defined below
        :param pulumi.Input[str] endpoint_location: Specifies the Azure location of the Endpoint,
               this must be specified for Profiles using the `Performance` routing method
               if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
               For Endpoints of type `azureEndpoints` the value will be taken from the
               location of the Azure target resource.
        :param pulumi.Input[str] endpoint_status: The status of the Endpoint, can be set to
               either `Enabled` or `Disabled`. Defaults to `Enabled`.
        :param pulumi.Input[List[pulumi.Input[str]]] geo_mappings: A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
        :param pulumi.Input[float] min_child_endpoints: This argument specifies the minimum number
               of endpoints that must be ‘online’ in the child profile in order for the
               parent profile to direct traffic to any of the endpoints in that child
               profile. This argument only applies to Endpoints of type `nestedEndpoints`
               and defaults to `1`.
        :param pulumi.Input[str] name: The name of the Traffic Manager endpoint. Changing this forces a
               new resource to be created.
        :param pulumi.Input[float] priority: Specifies the priority of this Endpoint, this must be
               specified for Profiles using the `Priority` traffic routing method. Supports
               values between 1 and 1000, with no Endpoints sharing the same value. If
               omitted the value will be computed in order of creation.
        :param pulumi.Input[str] profile_name: The name of the Traffic Manager Profile to attach
               create the Traffic Manager endpoint.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the Traffic Manager Profile exists.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointSubnetArgs']]]] subnets: One or more `subnet` blocks as defined below
        :param pulumi.Input[str] target: The FQDN DNS name of the target. This argument must be
               provided for an endpoint of type `externalEndpoints`, for other types it
               will be computed.
        :param pulumi.Input[str] target_resource_id: The resource id of an Azure resource to
               target. This argument must be provided for an endpoint of type
               `azureEndpoints` or `nestedEndpoints`.
        :param pulumi.Input[str] type: The Endpoint type, must be one of:
               - `azureEndpoints`
               - `externalEndpoints`
               - `nestedEndpoints`
        :param pulumi.Input[float] weight: Specifies how much traffic should be distributed to this
               endpoint, this must be specified for Profiles using the  `Weighted` traffic
               routing method. Supports values between 1 and 1000.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_headers"] = custom_headers
        __props__["endpoint_location"] = endpoint_location
        __props__["endpoint_monitor_status"] = endpoint_monitor_status
        __props__["endpoint_status"] = endpoint_status
        __props__["geo_mappings"] = geo_mappings
        __props__["min_child_endpoints"] = min_child_endpoints
        __props__["name"] = name
        __props__["priority"] = priority
        __props__["profile_name"] = profile_name
        __props__["resource_group_name"] = resource_group_name
        __props__["subnets"] = subnets
        __props__["target"] = target
        __props__["target_resource_id"] = target_resource_id
        __props__["type"] = type
        __props__["weight"] = weight
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customHeaders")
    def custom_headers(self) -> Optional[List['outputs.EndpointCustomHeader']]:
        """
        One or more `custom_header` blocks as defined below
        """
        return pulumi.get(self, "custom_headers")

    @property
    @pulumi.getter(name="endpointLocation")
    def endpoint_location(self) -> str:
        """
        Specifies the Azure location of the Endpoint,
        this must be specified for Profiles using the `Performance` routing method
        if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
        For Endpoints of type `azureEndpoints` the value will be taken from the
        location of the Azure target resource.
        """
        return pulumi.get(self, "endpoint_location")

    @property
    @pulumi.getter(name="endpointMonitorStatus")
    def endpoint_monitor_status(self) -> str:
        return pulumi.get(self, "endpoint_monitor_status")

    @property
    @pulumi.getter(name="endpointStatus")
    def endpoint_status(self) -> str:
        """
        The status of the Endpoint, can be set to
        either `Enabled` or `Disabled`. Defaults to `Enabled`.
        """
        return pulumi.get(self, "endpoint_status")

    @property
    @pulumi.getter(name="geoMappings")
    def geo_mappings(self) -> Optional[List[str]]:
        """
        A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
        """
        return pulumi.get(self, "geo_mappings")

    @property
    @pulumi.getter(name="minChildEndpoints")
    def min_child_endpoints(self) -> Optional[float]:
        """
        This argument specifies the minimum number
        of endpoints that must be ‘online’ in the child profile in order for the
        parent profile to direct traffic to any of the endpoints in that child
        profile. This argument only applies to Endpoints of type `nestedEndpoints`
        and defaults to `1`.
        """
        return pulumi.get(self, "min_child_endpoints")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Traffic Manager endpoint. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> float:
        """
        Specifies the priority of this Endpoint, this must be
        specified for Profiles using the `Priority` traffic routing method. Supports
        values between 1 and 1000, with no Endpoints sharing the same value. If
        omitted the value will be computed in order of creation.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> str:
        """
        The name of the Traffic Manager Profile to attach
        create the Traffic Manager endpoint.
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group where the Traffic Manager Profile exists.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def subnets(self) -> Optional[List['outputs.EndpointSubnet']]:
        """
        One or more `subnet` blocks as defined below
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        The FQDN DNS name of the target. This argument must be
        provided for an endpoint of type `externalEndpoints`, for other types it
        will be computed.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> Optional[str]:
        """
        The resource id of an Azure resource to
        target. This argument must be provided for an endpoint of type
        `azureEndpoints` or `nestedEndpoints`.
        """
        return pulumi.get(self, "target_resource_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Endpoint type, must be one of:
        - `azureEndpoints`
        - `externalEndpoints`
        - `nestedEndpoints`
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def weight(self) -> float:
        """
        Specifies how much traffic should be distributed to this
        endpoint, this must be specified for Profiles using the  `Weighted` traffic
        routing method. Supports values between 1 and 1000.
        """
        return pulumi.get(self, "weight")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

