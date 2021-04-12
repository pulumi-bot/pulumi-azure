# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SpringCloudJavaDeploymentArgs', 'SpringCloudJavaDeployment']

@pulumi.input_type
class SpringCloudJavaDeploymentArgs:
    def __init__(__self__, *,
                 spring_cloud_app_id: pulumi.Input[str],
                 cpu: Optional[pulumi.Input[int]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 jvm_options: Optional[pulumi.Input[str]] = None,
                 memory_in_gb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 runtime_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SpringCloudJavaDeployment resource.
        :param pulumi.Input[str] spring_cloud_app_id: Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
        :param pulumi.Input[int] cpu: Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
        :param pulumi.Input[int] instance_count: Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
        :param pulumi.Input[str] jvm_options: Specifies the jvm option of the Spring Cloud Deployment.
        :param pulumi.Input[int] memory_in_gb: Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runtime_version: Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
        """
        pulumi.set(__self__, "spring_cloud_app_id", spring_cloud_app_id)
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)
        if environment_variables is not None:
            pulumi.set(__self__, "environment_variables", environment_variables)
        if instance_count is not None:
            pulumi.set(__self__, "instance_count", instance_count)
        if jvm_options is not None:
            pulumi.set(__self__, "jvm_options", jvm_options)
        if memory_in_gb is not None:
            pulumi.set(__self__, "memory_in_gb", memory_in_gb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if runtime_version is not None:
            pulumi.set(__self__, "runtime_version", runtime_version)

    @property
    @pulumi.getter(name="springCloudAppId")
    def spring_cloud_app_id(self) -> pulumi.Input[str]:
        """
        Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "spring_cloud_app_id")

    @spring_cloud_app_id.setter
    def spring_cloud_app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "spring_cloud_app_id", value)

    @property
    @pulumi.getter
    def cpu(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
        """
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu", value)

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
        """
        return pulumi.get(self, "environment_variables")

    @environment_variables.setter
    def environment_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "environment_variables", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="jvmOptions")
    def jvm_options(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the jvm option of the Spring Cloud Deployment.
        """
        return pulumi.get(self, "jvm_options")

    @jvm_options.setter
    def jvm_options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "jvm_options", value)

    @property
    @pulumi.getter(name="memoryInGb")
    def memory_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
        """
        return pulumi.get(self, "memory_in_gb")

    @memory_in_gb.setter
    def memory_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_in_gb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="runtimeVersion")
    def runtime_version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
        """
        return pulumi.get(self, "runtime_version")

    @runtime_version.setter
    def runtime_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "runtime_version", value)


class SpringCloudJavaDeployment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu: Optional[pulumi.Input[int]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 jvm_options: Optional[pulumi.Input[str]] = None,
                 memory_in_gb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 runtime_version: Optional[pulumi.Input[str]] = None,
                 spring_cloud_app_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Spring Cloud Deployment with a Java runtime.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_spring_cloud_service = azure.appplatform.SpringCloudService("exampleSpringCloudService",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_spring_cloud_app = azure.appplatform.SpringCloudApp("exampleSpringCloudApp",
            resource_group_name=example_resource_group.name,
            service_name=example_spring_cloud_service.name,
            identity=azure.appplatform.SpringCloudAppIdentityArgs(
                type="SystemAssigned",
            ))
        example_spring_cloud_java_deployment = azure.appplatform.SpringCloudJavaDeployment("exampleSpringCloudJavaDeployment",
            spring_cloud_app_id=example_spring_cloud_app.id,
            cpu=2,
            instance_count=2,
            jvm_options="-XX:+PrintGC",
            memory_in_gb=4,
            runtime_version="Java_11",
            environment_variables={
                "Foo": "Bar",
                "Env": "Staging",
            })
        ```

        ## Import

        Spring Cloud Deployment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/service1/apps/app1/deployments/deploy1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cpu: Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
        :param pulumi.Input[int] instance_count: Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
        :param pulumi.Input[str] jvm_options: Specifies the jvm option of the Spring Cloud Deployment.
        :param pulumi.Input[int] memory_in_gb: Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runtime_version: Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
        :param pulumi.Input[str] spring_cloud_app_id: Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SpringCloudJavaDeploymentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure Spring Cloud Deployment with a Java runtime.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_spring_cloud_service = azure.appplatform.SpringCloudService("exampleSpringCloudService",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_spring_cloud_app = azure.appplatform.SpringCloudApp("exampleSpringCloudApp",
            resource_group_name=example_resource_group.name,
            service_name=example_spring_cloud_service.name,
            identity=azure.appplatform.SpringCloudAppIdentityArgs(
                type="SystemAssigned",
            ))
        example_spring_cloud_java_deployment = azure.appplatform.SpringCloudJavaDeployment("exampleSpringCloudJavaDeployment",
            spring_cloud_app_id=example_spring_cloud_app.id,
            cpu=2,
            instance_count=2,
            jvm_options="-XX:+PrintGC",
            memory_in_gb=4,
            runtime_version="Java_11",
            environment_variables={
                "Foo": "Bar",
                "Env": "Staging",
            })
        ```

        ## Import

        Spring Cloud Deployment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/service1/apps/app1/deployments/deploy1
        ```

        :param str resource_name: The name of the resource.
        :param SpringCloudJavaDeploymentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SpringCloudJavaDeploymentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu: Optional[pulumi.Input[int]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 jvm_options: Optional[pulumi.Input[str]] = None,
                 memory_in_gb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 runtime_version: Optional[pulumi.Input[str]] = None,
                 spring_cloud_app_id: Optional[pulumi.Input[str]] = None,
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

            __props__['cpu'] = cpu
            __props__['environment_variables'] = environment_variables
            __props__['instance_count'] = instance_count
            __props__['jvm_options'] = jvm_options
            __props__['memory_in_gb'] = memory_in_gb
            __props__['name'] = name
            __props__['runtime_version'] = runtime_version
            if spring_cloud_app_id is None and not opts.urn:
                raise TypeError("Missing required property 'spring_cloud_app_id'")
            __props__['spring_cloud_app_id'] = spring_cloud_app_id
        super(SpringCloudJavaDeployment, __self__).__init__(
            'azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cpu: Optional[pulumi.Input[int]] = None,
            environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            instance_count: Optional[pulumi.Input[int]] = None,
            jvm_options: Optional[pulumi.Input[str]] = None,
            memory_in_gb: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            runtime_version: Optional[pulumi.Input[str]] = None,
            spring_cloud_app_id: Optional[pulumi.Input[str]] = None) -> 'SpringCloudJavaDeployment':
        """
        Get an existing SpringCloudJavaDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cpu: Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
        :param pulumi.Input[int] instance_count: Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
        :param pulumi.Input[str] jvm_options: Specifies the jvm option of the Spring Cloud Deployment.
        :param pulumi.Input[int] memory_in_gb: Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runtime_version: Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
        :param pulumi.Input[str] spring_cloud_app_id: Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cpu"] = cpu
        __props__["environment_variables"] = environment_variables
        __props__["instance_count"] = instance_count
        __props__["jvm_options"] = jvm_options
        __props__["memory_in_gb"] = memory_in_gb
        __props__["name"] = name
        __props__["runtime_version"] = runtime_version
        __props__["spring_cloud_app_id"] = spring_cloud_app_id
        return SpringCloudJavaDeployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cpu(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
        """
        return pulumi.get(self, "environment_variables")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="jvmOptions")
    def jvm_options(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the jvm option of the Spring Cloud Deployment.
        """
        return pulumi.get(self, "jvm_options")

    @property
    @pulumi.getter(name="memoryInGb")
    def memory_in_gb(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
        """
        return pulumi.get(self, "memory_in_gb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="runtimeVersion")
    def runtime_version(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
        """
        return pulumi.get(self, "runtime_version")

    @property
    @pulumi.getter(name="springCloudAppId")
    def spring_cloud_app_id(self) -> pulumi.Output[str]:
        """
        Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "spring_cloud_app_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

