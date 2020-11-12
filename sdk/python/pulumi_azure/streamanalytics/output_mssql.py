# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['OutputMssql']


class OutputMssql(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Stream Analytics Output to Microsoft SQL Server Database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_sql_server = azure.sql.SqlServer("exampleSqlServer",
            resource_group_name=azurerm_resource_group["example"]["name"],
            location=azurerm_resource_group["example"]["location"],
            version="12.0",
            administrator_login="dbadmin",
            administrator_login_password="example-password")
        example_database = azure.sql.Database("exampleDatabase",
            resource_group_name=azurerm_resource_group["example"]["name"],
            location=azurerm_resource_group["example"]["location"],
            server_name=example_sql_server.name,
            requested_service_objective_name="S0",
            collation="SQL_LATIN1_GENERAL_CP1_CI_AS",
            max_size_bytes="268435456000",
            create_mode="Default")
        example_output_mssql = azure.streamanalytics.OutputMssql("exampleOutputMssql",
            stream_analytics_job_name=azurerm_stream_analytics_job["example"]["name"],
            resource_group_name=azurerm_stream_analytics_job["example"]["resource_group_name"],
            server=example_sql_server.fully_qualified_domain_name,
            user=example_sql_server.administrator_login,
            password=example_sql_server.administrator_login_password,
            database=example_database.name,
            table="ExampleTable")
        ```

        ## Import

        Stream Analytics Outputs to Microsoft SQL Server Database can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/outputMssql:OutputMssql example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] password: Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server: The SQL server url. Changing this forces a new resource to be created.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] table: Table in the database that the output points to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user: Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
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

            if database is None:
                raise TypeError("Missing required property 'database'")
            __props__['database'] = database
            __props__['name'] = name
            if password is None:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server is None:
                raise TypeError("Missing required property 'server'")
            __props__['server'] = server
            if stream_analytics_job_name is None:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
            if table is None:
                raise TypeError("Missing required property 'table'")
            __props__['table'] = table
            if user is None:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
        super(OutputMssql, __self__).__init__(
            'azure:streamanalytics/outputMssql:OutputMssql',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
            table: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'OutputMssql':
        """
        Get an existing OutputMssql resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] password: Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server: The SQL server url. Changing this forces a new resource to be created.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] table: Table in the database that the output points to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user: Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["database"] = database
        __props__["name"] = name
        __props__["password"] = password
        __props__["resource_group_name"] = resource_group_name
        __props__["server"] = server
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        __props__["table"] = table
        __props__["user"] = user
        return OutputMssql(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Output. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        The SQL server url. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @property
    @pulumi.getter
    def table(self) -> pulumi.Output[str]:
        """
        Table in the database that the output points to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "table")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "user")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

