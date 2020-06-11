# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FailoverGroup(pulumi.CustomResource):
    databases: pulumi.Output[list]
    """
    A list of database ids to add to the failover group
    """
    location: pulumi.Output[str]
    """
    the location of the failover group.
    """
    name: pulumi.Output[str]
    """
    The name of the failover group. Changing this forces a new resource to be created.
    """
    partner_servers: pulumi.Output[list]
    """
    A list of secondary servers as documented below

      * `id` (`str`) - the SQL server ID
      * `location` (`str`) - the location of the failover group.
      * `role` (`str`) - local replication role of the failover group instance.
    """
    read_write_endpoint_failover_policy: pulumi.Output[dict]
    """
    A read/write policy as documented below

      * `graceMinutes` (`float`) - Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
      * `mode` (`str`) - the failover mode. Possible values are `Manual`, `Automatic`
    """
    readonly_endpoint_failover_policy: pulumi.Output[dict]
    """
    a read-only policy as documented below

      * `mode` (`str`) - Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group containing the SQL server
    """
    role: pulumi.Output[str]
    """
    local replication role of the failover group instance.
    """
    server_name: pulumi.Output[str]
    """
    The name of the primary SQL server. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, databases=None, name=None, partner_servers=None, read_write_endpoint_failover_policy=None, readonly_endpoint_failover_policy=None, resource_group_name=None, server_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a failover group of databases on a collection of Azure SQL servers.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="uksouth")
        primary = azure.sql.SqlServer("primary",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            version="12.0",
            administrator_login="sqladmin",
            administrator_login_password="pa$$w0rd")
        secondary = azure.sql.SqlServer("secondary",
            resource_group_name=example_resource_group.name,
            location="northeurope",
            version="12.0",
            administrator_login="sqladmin",
            administrator_login_password="pa$$w0rd")
        db1 = azure.sql.Database("db1",
            resource_group_name=primary.resource_group_name,
            location=primary.location,
            server_name=primary.name)
        example_failover_group = azure.sql.FailoverGroup("exampleFailoverGroup",
            resource_group_name=primary.resource_group_name,
            server_name=primary.name,
            databases=[db1.id],
            partner_servers=[{
                "id": secondary.id,
            }],
            read_write_endpoint_failover_policy={
                "mode": "Automatic",
                "graceMinutes": 60,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] databases: A list of database ids to add to the failover group
        :param pulumi.Input[str] name: The name of the failover group. Changing this forces a new resource to be created.
        :param pulumi.Input[list] partner_servers: A list of secondary servers as documented below
        :param pulumi.Input[dict] read_write_endpoint_failover_policy: A read/write policy as documented below
        :param pulumi.Input[dict] readonly_endpoint_failover_policy: a read-only policy as documented below
        :param pulumi.Input[str] resource_group_name: The name of the resource group containing the SQL server
        :param pulumi.Input[str] server_name: The name of the primary SQL server. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **partner_servers** object supports the following:

          * `id` (`pulumi.Input[str]`) - the SQL server ID
          * `location` (`pulumi.Input[str]`) - the location of the failover group.
          * `role` (`pulumi.Input[str]`) - local replication role of the failover group instance.

        The **read_write_endpoint_failover_policy** object supports the following:

          * `graceMinutes` (`pulumi.Input[float]`) - Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
          * `mode` (`pulumi.Input[str]`) - the failover mode. Possible values are `Manual`, `Automatic`

        The **readonly_endpoint_failover_policy** object supports the following:

          * `mode` (`pulumi.Input[str]`) - Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['databases'] = databases
            __props__['name'] = name
            if partner_servers is None:
                raise TypeError("Missing required property 'partner_servers'")
            __props__['partner_servers'] = partner_servers
            if read_write_endpoint_failover_policy is None:
                raise TypeError("Missing required property 'read_write_endpoint_failover_policy'")
            __props__['read_write_endpoint_failover_policy'] = read_write_endpoint_failover_policy
            __props__['readonly_endpoint_failover_policy'] = readonly_endpoint_failover_policy
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['tags'] = tags
            __props__['location'] = None
            __props__['role'] = None
        super(FailoverGroup, __self__).__init__(
            'azure:sql/failoverGroup:FailoverGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, databases=None, location=None, name=None, partner_servers=None, read_write_endpoint_failover_policy=None, readonly_endpoint_failover_policy=None, resource_group_name=None, role=None, server_name=None, tags=None):
        """
        Get an existing FailoverGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] databases: A list of database ids to add to the failover group
        :param pulumi.Input[str] location: the location of the failover group.
        :param pulumi.Input[str] name: The name of the failover group. Changing this forces a new resource to be created.
        :param pulumi.Input[list] partner_servers: A list of secondary servers as documented below
        :param pulumi.Input[dict] read_write_endpoint_failover_policy: A read/write policy as documented below
        :param pulumi.Input[dict] readonly_endpoint_failover_policy: a read-only policy as documented below
        :param pulumi.Input[str] resource_group_name: The name of the resource group containing the SQL server
        :param pulumi.Input[str] role: local replication role of the failover group instance.
        :param pulumi.Input[str] server_name: The name of the primary SQL server. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **partner_servers** object supports the following:

          * `id` (`pulumi.Input[str]`) - the SQL server ID
          * `location` (`pulumi.Input[str]`) - the location of the failover group.
          * `role` (`pulumi.Input[str]`) - local replication role of the failover group instance.

        The **read_write_endpoint_failover_policy** object supports the following:

          * `graceMinutes` (`pulumi.Input[float]`) - Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
          * `mode` (`pulumi.Input[str]`) - the failover mode. Possible values are `Manual`, `Automatic`

        The **readonly_endpoint_failover_policy** object supports the following:

          * `mode` (`pulumi.Input[str]`) - Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["databases"] = databases
        __props__["location"] = location
        __props__["name"] = name
        __props__["partner_servers"] = partner_servers
        __props__["read_write_endpoint_failover_policy"] = read_write_endpoint_failover_policy
        __props__["readonly_endpoint_failover_policy"] = readonly_endpoint_failover_policy
        __props__["resource_group_name"] = resource_group_name
        __props__["role"] = role
        __props__["server_name"] = server_name
        __props__["tags"] = tags
        return FailoverGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

