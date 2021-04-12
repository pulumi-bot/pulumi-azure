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

__all__ = ['SparkPoolArgs', 'SparkPool']

@pulumi.input_type
class SparkPoolArgs:
    def __init__(__self__, *,
                 node_size: pulumi.Input[str],
                 node_size_family: pulumi.Input[str],
                 synapse_workspace_id: pulumi.Input[str],
                 auto_pause: Optional[pulumi.Input['SparkPoolAutoPauseArgs']] = None,
                 auto_scale: Optional[pulumi.Input['SparkPoolAutoScaleArgs']] = None,
                 library_requirement: Optional[pulumi.Input['SparkPoolLibraryRequirementArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 spark_events_folder: Optional[pulumi.Input[str]] = None,
                 spark_log_folder: Optional[pulumi.Input[str]] = None,
                 spark_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SparkPool resource.
        :param pulumi.Input[str] node_size: The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
        :param pulumi.Input[str] node_size_family: The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
        :param pulumi.Input['SparkPoolAutoPauseArgs'] auto_pause: An `auto_pause` block as defined below.
        :param pulumi.Input['SparkPoolAutoScaleArgs'] auto_scale: An `auto_scale` block as defined below. Exactly one of `node_count` or `auto_scale` must be specified.
        :param pulumi.Input['SparkPoolLibraryRequirementArgs'] library_requirement: A `library_requirement` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
        :param pulumi.Input[int] node_count: The number of nodes in the Spark Pool. Exactly one of `node_count` or `auto_scale` must be specified.
        :param pulumi.Input[str] spark_events_folder: The Spark events folder. Defaults to `/events`.
        :param pulumi.Input[str] spark_log_folder: The default folder where Spark logs will be written. Defaults to `/logs`.
        :param pulumi.Input[str] spark_version: The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Synapse Spark Pool.
        """
        pulumi.set(__self__, "node_size", node_size)
        pulumi.set(__self__, "node_size_family", node_size_family)
        pulumi.set(__self__, "synapse_workspace_id", synapse_workspace_id)
        if auto_pause is not None:
            pulumi.set(__self__, "auto_pause", auto_pause)
        if auto_scale is not None:
            pulumi.set(__self__, "auto_scale", auto_scale)
        if library_requirement is not None:
            pulumi.set(__self__, "library_requirement", library_requirement)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_count is not None:
            pulumi.set(__self__, "node_count", node_count)
        if spark_events_folder is not None:
            pulumi.set(__self__, "spark_events_folder", spark_events_folder)
        if spark_log_folder is not None:
            pulumi.set(__self__, "spark_log_folder", spark_log_folder)
        if spark_version is not None:
            pulumi.set(__self__, "spark_version", spark_version)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="nodeSize")
    def node_size(self) -> pulumi.Input[str]:
        """
        The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
        """
        return pulumi.get(self, "node_size")

    @node_size.setter
    def node_size(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_size", value)

    @property
    @pulumi.getter(name="nodeSizeFamily")
    def node_size_family(self) -> pulumi.Input[str]:
        """
        The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
        """
        return pulumi.get(self, "node_size_family")

    @node_size_family.setter
    def node_size_family(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_size_family", value)

    @property
    @pulumi.getter(name="synapseWorkspaceId")
    def synapse_workspace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
        """
        return pulumi.get(self, "synapse_workspace_id")

    @synapse_workspace_id.setter
    def synapse_workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "synapse_workspace_id", value)

    @property
    @pulumi.getter(name="autoPause")
    def auto_pause(self) -> Optional[pulumi.Input['SparkPoolAutoPauseArgs']]:
        """
        An `auto_pause` block as defined below.
        """
        return pulumi.get(self, "auto_pause")

    @auto_pause.setter
    def auto_pause(self, value: Optional[pulumi.Input['SparkPoolAutoPauseArgs']]):
        pulumi.set(self, "auto_pause", value)

    @property
    @pulumi.getter(name="autoScale")
    def auto_scale(self) -> Optional[pulumi.Input['SparkPoolAutoScaleArgs']]:
        """
        An `auto_scale` block as defined below. Exactly one of `node_count` or `auto_scale` must be specified.
        """
        return pulumi.get(self, "auto_scale")

    @auto_scale.setter
    def auto_scale(self, value: Optional[pulumi.Input['SparkPoolAutoScaleArgs']]):
        pulumi.set(self, "auto_scale", value)

    @property
    @pulumi.getter(name="libraryRequirement")
    def library_requirement(self) -> Optional[pulumi.Input['SparkPoolLibraryRequirementArgs']]:
        """
        A `library_requirement` block as defined below.
        """
        return pulumi.get(self, "library_requirement")

    @library_requirement.setter
    def library_requirement(self, value: Optional[pulumi.Input['SparkPoolLibraryRequirementArgs']]):
        pulumi.set(self, "library_requirement", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of nodes in the Spark Pool. Exactly one of `node_count` or `auto_scale` must be specified.
        """
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter(name="sparkEventsFolder")
    def spark_events_folder(self) -> Optional[pulumi.Input[str]]:
        """
        The Spark events folder. Defaults to `/events`.
        """
        return pulumi.get(self, "spark_events_folder")

    @spark_events_folder.setter
    def spark_events_folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spark_events_folder", value)

    @property
    @pulumi.getter(name="sparkLogFolder")
    def spark_log_folder(self) -> Optional[pulumi.Input[str]]:
        """
        The default folder where Spark logs will be written. Defaults to `/logs`.
        """
        return pulumi.get(self, "spark_log_folder")

    @spark_log_folder.setter
    def spark_log_folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spark_log_folder", value)

    @property
    @pulumi.getter(name="sparkVersion")
    def spark_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
        """
        return pulumi.get(self, "spark_version")

    @spark_version.setter
    def spark_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spark_version", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the Synapse Spark Pool.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class SparkPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_pause: Optional[pulumi.Input[pulumi.InputType['SparkPoolAutoPauseArgs']]] = None,
                 auto_scale: Optional[pulumi.Input[pulumi.InputType['SparkPoolAutoScaleArgs']]] = None,
                 library_requirement: Optional[pulumi.Input[pulumi.InputType['SparkPoolLibraryRequirementArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 node_size: Optional[pulumi.Input[str]] = None,
                 node_size_family: Optional[pulumi.Input[str]] = None,
                 spark_events_folder: Optional[pulumi.Input[str]] = None,
                 spark_log_folder: Optional[pulumi.Input[str]] = None,
                 spark_version: Optional[pulumi.Input[str]] = None,
                 synapse_workspace_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Synapse Spark Pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            account_kind="StorageV2",
            is_hns_enabled=True)
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        example_workspace = azure.synapse.Workspace("exampleWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!")
        example_spark_pool = azure.synapse.SparkPool("exampleSparkPool",
            synapse_workspace_id=example_workspace.id,
            node_size_family="MemoryOptimized",
            node_size="Small",
            auto_scale=azure.synapse.SparkPoolAutoScaleArgs(
                max_node_count=50,
                min_node_count=3,
            ),
            auto_pause=azure.synapse.SparkPoolAutoPauseArgs(
                delay_in_minutes=15,
            ),
            tags={
                "ENV": "Production",
            })
        ```

        ## Import

        Synapse Spark Pool can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:synapse/sparkPool:SparkPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SparkPoolAutoPauseArgs']] auto_pause: An `auto_pause` block as defined below.
        :param pulumi.Input[pulumi.InputType['SparkPoolAutoScaleArgs']] auto_scale: An `auto_scale` block as defined below. Exactly one of `node_count` or `auto_scale` must be specified.
        :param pulumi.Input[pulumi.InputType['SparkPoolLibraryRequirementArgs']] library_requirement: A `library_requirement` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
        :param pulumi.Input[int] node_count: The number of nodes in the Spark Pool. Exactly one of `node_count` or `auto_scale` must be specified.
        :param pulumi.Input[str] node_size: The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
        :param pulumi.Input[str] node_size_family: The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
        :param pulumi.Input[str] spark_events_folder: The Spark events folder. Defaults to `/events`.
        :param pulumi.Input[str] spark_log_folder: The default folder where Spark logs will be written. Defaults to `/logs`.
        :param pulumi.Input[str] spark_version: The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Synapse Spark Pool.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SparkPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Synapse Spark Pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            account_kind="StorageV2",
            is_hns_enabled=True)
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        example_workspace = azure.synapse.Workspace("exampleWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!")
        example_spark_pool = azure.synapse.SparkPool("exampleSparkPool",
            synapse_workspace_id=example_workspace.id,
            node_size_family="MemoryOptimized",
            node_size="Small",
            auto_scale=azure.synapse.SparkPoolAutoScaleArgs(
                max_node_count=50,
                min_node_count=3,
            ),
            auto_pause=azure.synapse.SparkPoolAutoPauseArgs(
                delay_in_minutes=15,
            ),
            tags={
                "ENV": "Production",
            })
        ```

        ## Import

        Synapse Spark Pool can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:synapse/sparkPool:SparkPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
        ```

        :param str resource_name: The name of the resource.
        :param SparkPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SparkPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_pause: Optional[pulumi.Input[pulumi.InputType['SparkPoolAutoPauseArgs']]] = None,
                 auto_scale: Optional[pulumi.Input[pulumi.InputType['SparkPoolAutoScaleArgs']]] = None,
                 library_requirement: Optional[pulumi.Input[pulumi.InputType['SparkPoolLibraryRequirementArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 node_size: Optional[pulumi.Input[str]] = None,
                 node_size_family: Optional[pulumi.Input[str]] = None,
                 spark_events_folder: Optional[pulumi.Input[str]] = None,
                 spark_log_folder: Optional[pulumi.Input[str]] = None,
                 spark_version: Optional[pulumi.Input[str]] = None,
                 synapse_workspace_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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

            __props__['auto_pause'] = auto_pause
            __props__['auto_scale'] = auto_scale
            __props__['library_requirement'] = library_requirement
            __props__['name'] = name
            __props__['node_count'] = node_count
            if node_size is None and not opts.urn:
                raise TypeError("Missing required property 'node_size'")
            __props__['node_size'] = node_size
            if node_size_family is None and not opts.urn:
                raise TypeError("Missing required property 'node_size_family'")
            __props__['node_size_family'] = node_size_family
            __props__['spark_events_folder'] = spark_events_folder
            __props__['spark_log_folder'] = spark_log_folder
            __props__['spark_version'] = spark_version
            if synapse_workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'synapse_workspace_id'")
            __props__['synapse_workspace_id'] = synapse_workspace_id
            __props__['tags'] = tags
        super(SparkPool, __self__).__init__(
            'azure:synapse/sparkPool:SparkPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_pause: Optional[pulumi.Input[pulumi.InputType['SparkPoolAutoPauseArgs']]] = None,
            auto_scale: Optional[pulumi.Input[pulumi.InputType['SparkPoolAutoScaleArgs']]] = None,
            library_requirement: Optional[pulumi.Input[pulumi.InputType['SparkPoolLibraryRequirementArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_count: Optional[pulumi.Input[int]] = None,
            node_size: Optional[pulumi.Input[str]] = None,
            node_size_family: Optional[pulumi.Input[str]] = None,
            spark_events_folder: Optional[pulumi.Input[str]] = None,
            spark_log_folder: Optional[pulumi.Input[str]] = None,
            spark_version: Optional[pulumi.Input[str]] = None,
            synapse_workspace_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SparkPool':
        """
        Get an existing SparkPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SparkPoolAutoPauseArgs']] auto_pause: An `auto_pause` block as defined below.
        :param pulumi.Input[pulumi.InputType['SparkPoolAutoScaleArgs']] auto_scale: An `auto_scale` block as defined below. Exactly one of `node_count` or `auto_scale` must be specified.
        :param pulumi.Input[pulumi.InputType['SparkPoolLibraryRequirementArgs']] library_requirement: A `library_requirement` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
        :param pulumi.Input[int] node_count: The number of nodes in the Spark Pool. Exactly one of `node_count` or `auto_scale` must be specified.
        :param pulumi.Input[str] node_size: The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
        :param pulumi.Input[str] node_size_family: The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
        :param pulumi.Input[str] spark_events_folder: The Spark events folder. Defaults to `/events`.
        :param pulumi.Input[str] spark_log_folder: The default folder where Spark logs will be written. Defaults to `/logs`.
        :param pulumi.Input[str] spark_version: The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Synapse Spark Pool.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_pause"] = auto_pause
        __props__["auto_scale"] = auto_scale
        __props__["library_requirement"] = library_requirement
        __props__["name"] = name
        __props__["node_count"] = node_count
        __props__["node_size"] = node_size
        __props__["node_size_family"] = node_size_family
        __props__["spark_events_folder"] = spark_events_folder
        __props__["spark_log_folder"] = spark_log_folder
        __props__["spark_version"] = spark_version
        __props__["synapse_workspace_id"] = synapse_workspace_id
        __props__["tags"] = tags
        return SparkPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoPause")
    def auto_pause(self) -> pulumi.Output[Optional['outputs.SparkPoolAutoPause']]:
        """
        An `auto_pause` block as defined below.
        """
        return pulumi.get(self, "auto_pause")

    @property
    @pulumi.getter(name="autoScale")
    def auto_scale(self) -> pulumi.Output[Optional['outputs.SparkPoolAutoScale']]:
        """
        An `auto_scale` block as defined below. Exactly one of `node_count` or `auto_scale` must be specified.
        """
        return pulumi.get(self, "auto_scale")

    @property
    @pulumi.getter(name="libraryRequirement")
    def library_requirement(self) -> pulumi.Output[Optional['outputs.SparkPoolLibraryRequirement']]:
        """
        A `library_requirement` block as defined below.
        """
        return pulumi.get(self, "library_requirement")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of nodes in the Spark Pool. Exactly one of `node_count` or `auto_scale` must be specified.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="nodeSize")
    def node_size(self) -> pulumi.Output[str]:
        """
        The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
        """
        return pulumi.get(self, "node_size")

    @property
    @pulumi.getter(name="nodeSizeFamily")
    def node_size_family(self) -> pulumi.Output[str]:
        """
        The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
        """
        return pulumi.get(self, "node_size_family")

    @property
    @pulumi.getter(name="sparkEventsFolder")
    def spark_events_folder(self) -> pulumi.Output[Optional[str]]:
        """
        The Spark events folder. Defaults to `/events`.
        """
        return pulumi.get(self, "spark_events_folder")

    @property
    @pulumi.getter(name="sparkLogFolder")
    def spark_log_folder(self) -> pulumi.Output[Optional[str]]:
        """
        The default folder where Spark logs will be written. Defaults to `/logs`.
        """
        return pulumi.get(self, "spark_log_folder")

    @property
    @pulumi.getter(name="sparkVersion")
    def spark_version(self) -> pulumi.Output[Optional[str]]:
        """
        The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
        """
        return pulumi.get(self, "spark_version")

    @property
    @pulumi.getter(name="synapseWorkspaceId")
    def synapse_workspace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
        """
        return pulumi.get(self, "synapse_workspace_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Synapse Spark Pool.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

