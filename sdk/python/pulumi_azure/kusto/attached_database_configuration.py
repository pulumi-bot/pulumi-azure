# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['AttachedDatabaseConfiguration']


class AttachedDatabaseConfiguration(pulumi.CustomResource):
    attached_database_names: pulumi.Output[List[str]] = pulumi.output_property("attachedDatabaseNames")
    """
    The list of databases from the `cluster_resource_id` which are currently attached to the cluster.
    """
    cluster_name: pulumi.Output[str] = pulumi.output_property("clusterName")
    """
    Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
    """
    cluster_resource_id: pulumi.Output[str] = pulumi.output_property("clusterResourceId")
    """
    The resource id of the cluster where the databases you would like to attach reside.
    """
    database_name: pulumi.Output[str] = pulumi.output_property("databaseName")
    """
    The name of the database which you would like to attach, use * if you want to follow all current and future databases.
    """
    default_principal_modification_kind: pulumi.Output[Optional[str]] = pulumi.output_property("defaultPrincipalModificationKind")
    """
    The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, cluster_name: Optional[pulumi.Input[str]] = None, cluster_resource_id: Optional[pulumi.Input[str]] = None, database_name: Optional[pulumi.Input[str]] = None, default_principal_modification_kind: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Kusto (also known as Azure Data Explorer) Attached Database Configuration

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="East US")
        follower_cluster = azure.kusto.Cluster("followerCluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku={
                "name": "Dev(No SLA)_Standard_D11_v2",
                "capacity": 1,
            })
        followed_cluster = azure.kusto.Cluster("followedCluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku={
                "name": "Dev(No SLA)_Standard_D11_v2",
                "capacity": 1,
            })
        followed_database = azure.kusto.Database("followedDatabase",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=azurerm_kusto_cluster["cluster2"]["name"])
        example = azure.kusto.AttachedDatabaseConfiguration("example",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=follower_cluster.name,
            cluster_resource_id=followed_cluster.id,
            database_name="*",
            default_principal_modifications_kind="None")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] cluster_resource_id: The resource id of the cluster where the databases you would like to attach reside.
        :param pulumi.Input[str] database_name: The name of the database which you would like to attach, use * if you want to follow all current and future databases.
        :param pulumi.Input[str] default_principal_modification_kind: The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
        :param pulumi.Input[str] location: Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
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

            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            if cluster_resource_id is None:
                raise TypeError("Missing required property 'cluster_resource_id'")
            __props__['cluster_resource_id'] = cluster_resource_id
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['default_principal_modification_kind'] = default_principal_modification_kind
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['attached_database_names'] = None
        super(AttachedDatabaseConfiguration, __self__).__init__(
            'azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, attached_database_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, cluster_name: Optional[pulumi.Input[str]] = None, cluster_resource_id: Optional[pulumi.Input[str]] = None, database_name: Optional[pulumi.Input[str]] = None, default_principal_modification_kind: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None) -> 'AttachedDatabaseConfiguration':
        """
        Get an existing AttachedDatabaseConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] attached_database_names: The list of databases from the `cluster_resource_id` which are currently attached to the cluster.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] cluster_resource_id: The resource id of the cluster where the databases you would like to attach reside.
        :param pulumi.Input[str] database_name: The name of the database which you would like to attach, use * if you want to follow all current and future databases.
        :param pulumi.Input[str] default_principal_modification_kind: The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
        :param pulumi.Input[str] location: Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attached_database_names"] = attached_database_names
        __props__["cluster_name"] = cluster_name
        __props__["cluster_resource_id"] = cluster_resource_id
        __props__["database_name"] = database_name
        __props__["default_principal_modification_kind"] = default_principal_modification_kind
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        return AttachedDatabaseConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

