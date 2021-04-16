# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DatabasePrincipalAssignmentArgs', 'DatabasePrincipalAssignment']

@pulumi.input_type
class DatabasePrincipalAssignmentArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 database_name: pulumi.Input[str],
                 principal_id: pulumi.Input[str],
                 principal_type: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 role: pulumi.Input[str],
                 tenant_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabasePrincipalAssignment resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the database in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_id: The object id of the principal. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_type: The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The tenant id in which the principal resides. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "principal_type", principal_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        The name of the database in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Input[str]:
        """
        The object id of the principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Input[str]:
        """
        The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_type")

    @principal_type.setter
    def principal_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_type", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The tenant id in which the principal resides. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DatabasePrincipalAssignmentState:
    def __init__(__self__, *,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 principal_name: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tenant_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabasePrincipalAssignment resources.
        :param pulumi.Input[str] cluster_name: The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the database in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_id: The object id of the principal. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_name: The name of the principal.
        :param pulumi.Input[str] principal_type: The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The tenant id in which the principal resides. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_name: The name of the tenant.
        """
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if principal_name is not None:
            pulumi.set(__self__, "principal_name", principal_name)
        if principal_type is not None:
            pulumi.set(__self__, "principal_type", principal_type)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if tenant_name is not None:
            pulumi.set(__self__, "tenant_name", tenant_name)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object id of the principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the principal.
        """
        return pulumi.get(self, "principal_name")

    @principal_name.setter
    def principal_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_name", value)

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_type")

    @principal_type.setter
    def principal_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_type", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant id in which the principal resides. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the tenant.
        """
        return pulumi.get(self, "tenant_name")

    @tenant_name.setter
    def tenant_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_name", value)


class DatabasePrincipalAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Database Principal Assignment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        rg = azure.core.ResourceGroup("rg", location="West Europe")
        example_cluster = azure.kusto.Cluster("exampleCluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ))
        example_database = azure.kusto.Database("exampleDatabase",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=example_cluster.name,
            hot_cache_period="P7D",
            soft_delete_period="P31D")
        example_database_principal_assignment = azure.kusto.DatabasePrincipalAssignment("exampleDatabasePrincipalAssignment",
            resource_group_name=rg.name,
            cluster_name=example_cluster.name,
            database_name=example_database.name,
            tenant_id=current.tenant_id,
            principal_id=current.client_id,
            principal_type="App",
            role="Viewer")
        ```

        ## Import

        Kusto Database Principal Assignment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/PrincipalAssignments/assignment1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the database in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_id: The object id of the principal. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_type: The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The tenant id in which the principal resides. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasePrincipalAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Database Principal Assignment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        rg = azure.core.ResourceGroup("rg", location="West Europe")
        example_cluster = azure.kusto.Cluster("exampleCluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ))
        example_database = azure.kusto.Database("exampleDatabase",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=example_cluster.name,
            hot_cache_period="P7D",
            soft_delete_period="P31D")
        example_database_principal_assignment = azure.kusto.DatabasePrincipalAssignment("exampleDatabasePrincipalAssignment",
            resource_group_name=rg.name,
            cluster_name=example_cluster.name,
            database_name=example_database.name,
            tenant_id=current.tenant_id,
            principal_id=current.client_id,
            principal_type="App",
            role="Viewer")
        ```

        ## Import

        Kusto Database Principal Assignment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/PrincipalAssignments/assignment1
        ```

        :param str resource_name: The name of the resource.
        :param DatabasePrincipalAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasePrincipalAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasePrincipalAssignmentArgs.__new__(DatabasePrincipalAssignmentArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["name"] = name
            if principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'principal_id'")
            __props__.__dict__["principal_id"] = principal_id
            if principal_type is None and not opts.urn:
                raise TypeError("Missing required property 'principal_type'")
            __props__.__dict__["principal_type"] = principal_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["principal_name"] = None
            __props__.__dict__["tenant_name"] = None
        super(DatabasePrincipalAssignment, __self__).__init__(
            'azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            principal_id: Optional[pulumi.Input[str]] = None,
            principal_name: Optional[pulumi.Input[str]] = None,
            principal_type: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            tenant_name: Optional[pulumi.Input[str]] = None) -> 'DatabasePrincipalAssignment':
        """
        Get an existing DatabasePrincipalAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the database in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_id: The object id of the principal. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_name: The name of the principal.
        :param pulumi.Input[str] principal_type: The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The tenant id in which the principal resides. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_name: The name of the tenant.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabasePrincipalAssignmentState.__new__(_DatabasePrincipalAssignmentState)

        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["name"] = name
        __props__.__dict__["principal_id"] = principal_id
        __props__.__dict__["principal_name"] = principal_name
        __props__.__dict__["principal_type"] = principal_type
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["role"] = role
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["tenant_name"] = tenant_name
        return DatabasePrincipalAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the database in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[str]:
        """
        The object id of the principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> pulumi.Output[str]:
        """
        The name of the principal.
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Output[str]:
        """
        The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The tenant id in which the principal resides. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> pulumi.Output[str]:
        """
        The name of the tenant.
        """
        return pulumi.get(self, "tenant_name")

