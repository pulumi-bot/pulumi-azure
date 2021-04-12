# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 sku: pulumi.Input['ServiceSkuArgs'],
                 cors: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceCorArgs']]]] = None,
                 features: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceFeatureArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstream_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceUpstreamEndpointArgs']]]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input['ServiceSkuArgs'] sku: A `sku` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceCorArgs']]] cors: A `cors` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceFeatureArgs']]] features: A `features` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceUpstreamEndpointArgs']]] upstream_endpoints: An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if cors is not None:
            pulumi.set(__self__, "cors", cors)
        if features is not None:
            pulumi.set(__self__, "features", features)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if upstream_endpoints is not None:
            pulumi.set(__self__, "upstream_endpoints", upstream_endpoints)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input['ServiceSkuArgs']:
        """
        A `sku` block as documented below.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['ServiceSkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def cors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceCorArgs']]]]:
        """
        A `cors` block as documented below.
        """
        return pulumi.get(self, "cors")

    @cors.setter
    def cors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceCorArgs']]]]):
        pulumi.set(self, "cors", value)

    @property
    @pulumi.getter
    def features(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceFeatureArgs']]]]:
        """
        A `features` block as documented below.
        """
        return pulumi.get(self, "features")

    @features.setter
    def features(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceFeatureArgs']]]]):
        pulumi.set(self, "features", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SignalR service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="upstreamEndpoints")
    def upstream_endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceUpstreamEndpointArgs']]]]:
        """
        An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        """
        return pulumi.get(self, "upstream_endpoints")

    @upstream_endpoints.setter
    def upstream_endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceUpstreamEndpointArgs']]]]):
        pulumi.set(self, "upstream_endpoints", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCorArgs']]]]] = None,
                 features: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceFeatureArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ServiceSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstream_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceUpstreamEndpointArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure SignalR service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_service = azure.signalr.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=azure.signalr.ServiceSkuArgs(
                name="Free_F1",
                capacity=1,
            ),
            cors=[azure.signalr.ServiceCorArgs(
                allowed_origins=["http://www.example.com"],
            )],
            features=[azure.signalr.ServiceFeatureArgs(
                flag="ServiceMode",
                value="Default",
            )],
            upstream_endpoints=[azure.signalr.ServiceUpstreamEndpointArgs(
                category_patterns=[
                    "connections",
                    "messages",
                ],
                event_patterns=["*"],
                hub_patterns=["hub1"],
                url_template="http://foo.com",
            )])
        ```

        ## Import

        SignalR services can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:signalr/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/SignalR/tfex-signalr
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCorArgs']]]] cors: A `cors` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceFeatureArgs']]]] features: A `features` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ServiceSkuArgs']] sku: A `sku` block as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceUpstreamEndpointArgs']]]] upstream_endpoints: An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure SignalR service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_service = azure.signalr.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=azure.signalr.ServiceSkuArgs(
                name="Free_F1",
                capacity=1,
            ),
            cors=[azure.signalr.ServiceCorArgs(
                allowed_origins=["http://www.example.com"],
            )],
            features=[azure.signalr.ServiceFeatureArgs(
                flag="ServiceMode",
                value="Default",
            )],
            upstream_endpoints=[azure.signalr.ServiceUpstreamEndpointArgs(
                category_patterns=[
                    "connections",
                    "messages",
                ],
                event_patterns=["*"],
                hub_patterns=["hub1"],
                url_template="http://foo.com",
            )])
        ```

        ## Import

        SignalR services can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:signalr/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/SignalR/tfex-signalr
        ```

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCorArgs']]]]] = None,
                 features: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceFeatureArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ServiceSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstream_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceUpstreamEndpointArgs']]]]] = None,
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

            __props__['cors'] = cors
            __props__['features'] = features
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['upstream_endpoints'] = upstream_endpoints
            __props__['hostname'] = None
            __props__['ip_address'] = None
            __props__['primary_access_key'] = None
            __props__['primary_connection_string'] = None
            __props__['public_port'] = None
            __props__['secondary_access_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['server_port'] = None
        super(Service, __self__).__init__(
            'azure:signalr/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCorArgs']]]]] = None,
            features: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceFeatureArgs']]]]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            primary_access_key: Optional[pulumi.Input[str]] = None,
            primary_connection_string: Optional[pulumi.Input[str]] = None,
            public_port: Optional[pulumi.Input[int]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_access_key: Optional[pulumi.Input[str]] = None,
            secondary_connection_string: Optional[pulumi.Input[str]] = None,
            server_port: Optional[pulumi.Input[int]] = None,
            sku: Optional[pulumi.Input[pulumi.InputType['ServiceSkuArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            upstream_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceUpstreamEndpointArgs']]]]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCorArgs']]]] cors: A `cors` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceFeatureArgs']]]] features: A `features` block as documented below.
        :param pulumi.Input[str] hostname: The FQDN of the SignalR service.
        :param pulumi.Input[str] ip_address: The publicly accessible IP of the SignalR service.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The primary access key for the SignalR service.
        :param pulumi.Input[str] primary_connection_string: The primary connection string for the SignalR service.
        :param pulumi.Input[int] public_port: The publicly accessible port of the SignalR service which is designed for browser/client use.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The secondary access key for the SignalR service.
        :param pulumi.Input[str] secondary_connection_string: The secondary connection string for the SignalR service.
        :param pulumi.Input[int] server_port: The publicly accessible port of the SignalR service which is designed for customer server side use.
        :param pulumi.Input[pulumi.InputType['ServiceSkuArgs']] sku: A `sku` block as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceUpstreamEndpointArgs']]]] upstream_endpoints: An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cors"] = cors
        __props__["features"] = features
        __props__["hostname"] = hostname
        __props__["ip_address"] = ip_address
        __props__["location"] = location
        __props__["name"] = name
        __props__["primary_access_key"] = primary_access_key
        __props__["primary_connection_string"] = primary_connection_string
        __props__["public_port"] = public_port
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["server_port"] = server_port
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["upstream_endpoints"] = upstream_endpoints
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cors(self) -> pulumi.Output[Sequence['outputs.ServiceCor']]:
        """
        A `cors` block as documented below.
        """
        return pulumi.get(self, "cors")

    @property
    @pulumi.getter
    def features(self) -> pulumi.Output[Sequence['outputs.ServiceFeature']]:
        """
        A `features` block as documented below.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The FQDN of the SignalR service.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The publicly accessible IP of the SignalR service.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the SignalR service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> pulumi.Output[str]:
        """
        The primary access key for the SignalR service.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> pulumi.Output[str]:
        """
        The primary connection string for the SignalR service.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> pulumi.Output[int]:
        """
        The publicly accessible port of the SignalR service which is designed for browser/client use.
        """
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> pulumi.Output[str]:
        """
        The secondary access key for the SignalR service.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> pulumi.Output[str]:
        """
        The secondary connection string for the SignalR service.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> pulumi.Output[int]:
        """
        The publicly accessible port of the SignalR service which is designed for customer server side use.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.ServiceSku']:
        """
        A `sku` block as documented below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="upstreamEndpoints")
    def upstream_endpoints(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceUpstreamEndpoint']]]:
        """
        An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        """
        return pulumi.get(self, "upstream_endpoints")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

