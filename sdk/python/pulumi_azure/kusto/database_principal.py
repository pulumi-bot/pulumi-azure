# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['DatabasePrincipal']


class DatabasePrincipal(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Database Principal

        > **NOTE:** This resource is being **deprecated** due to API updates and should no longer be used.  Please use kusto.DatabasePrincipalAssignment instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        rg = azure.core.ResourceGroup("rg", location="East US")
        cluster = azure.kusto.Cluster("cluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku={
                "name": "Standard_D13_v2",
                "capacity": 2,
            })
        database = azure.kusto.Database("database",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=cluster.name,
            hot_cache_period="P7D",
            soft_delete_period="P31D")
        principal = azure.kusto.DatabasePrincipal("principal",
            resource_group_name=rg.name,
            cluster_name=cluster.name,
            database_name=azurerm_kusto_database["test"]["name"],
            role="Viewer",
            type="User",
            client_id=current.tenant_id,
            object_id=current.client_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The Client ID that owns the specified `object_id`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] object_id: An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
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

            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            if object_id is None:
                raise TypeError("Missing required property 'object_id'")
            __props__['object_id'] = object_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['app_id'] = None
            __props__['email'] = None
            __props__['fully_qualified_name'] = None
            __props__['name'] = None
        super(DatabasePrincipal, __self__).__init__(
            'azure:kusto/databasePrincipal:DatabasePrincipal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            fully_qualified_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'DatabasePrincipal':
        """
        Get an existing DatabasePrincipal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The app id, if not empty, of the principal.
        :param pulumi.Input[str] client_id: The Client ID that owns the specified `object_id`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] email: The email, if not empty, of the principal.
        :param pulumi.Input[str] fully_qualified_name: The fully qualified name of the principal.
        :param pulumi.Input[str] name: The name of the Kusto Database Principal.
        :param pulumi.Input[str] object_id: An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["client_id"] = client_id
        __props__["cluster_name"] = cluster_name
        __props__["database_name"] = database_name
        __props__["email"] = email
        __props__["fully_qualified_name"] = fully_qualified_name
        __props__["name"] = name
        __props__["object_id"] = object_id
        __props__["resource_group_name"] = resource_group_name
        __props__["role"] = role
        __props__["type"] = type
        return DatabasePrincipal(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        The app id, if not empty, of the principal.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The Client ID that owns the specified `object_id`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email, if not empty, of the principal.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="fullyQualifiedName")
    def fully_qualified_name(self) -> str:
        """
        The fully qualified name of the principal.
        """
        return pulumi.get(self, "fully_qualified_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Kusto Database Principal.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        """
        An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

