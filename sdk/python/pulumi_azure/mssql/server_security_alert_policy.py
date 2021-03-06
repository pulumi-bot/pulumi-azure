# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServerSecurityAlertPolicy(pulumi.CustomResource):
    disabled_alerts: pulumi.Output[list]
    """
    Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
    """
    email_account_admins: pulumi.Output[bool]
    """
    Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
    """
    email_addresses: pulumi.Output[list]
    """
    Specifies an array of e-mail addresses to which the alert is sent.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
    """
    retention_days: pulumi.Output[float]
    """
    Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
    """
    server_name: pulumi.Output[str]
    """
    Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
    """
    state: pulumi.Output[str]
    """
    Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`, `New`.
    """
    storage_account_access_key: pulumi.Output[str]
    """
    Specifies the identifier key of the Threat Detection audit storage account.
    """
    storage_endpoint: pulumi.Output[str]
    """
    Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
    """
    def __init__(__self__, resource_name, opts=None, disabled_alerts=None, email_account_admins=None, email_addresses=None, resource_group_name=None, retention_days=None, server_name=None, state=None, storage_account_access_key=None, storage_endpoint=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Security Alert Policy for a MSSQL Server.

        > **NOTE** Security Alert Policy is currently only available for MS SQL databases.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_sql_server = azure.sql.SqlServer("exampleSqlServer",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            version="12.0",
            administrator_login="4dm1n157r470r",
            administrator_login_password="4-v3ry-53cr37-p455w0rd")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_server_security_alert_policy = azure.mssql.ServerSecurityAlertPolicy("exampleServerSecurityAlertPolicy",
            resource_group_name=example_resource_group.name,
            server_name=example_sql_server.name,
            state="Enabled",
            storage_endpoint=example_account.primary_blob_endpoint,
            storage_account_access_key=example_account.primary_access_key,
            disabled_alerts=[
                "Sql_Injection",
                "Data_Exfiltration",
            ],
            retention_days=20)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] disabled_alerts: Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
        :param pulumi.Input[bool] email_account_admins: Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
        :param pulumi.Input[list] email_addresses: Specifies an array of e-mail addresses to which the alert is sent.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[float] retention_days: Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
        :param pulumi.Input[str] server_name: Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] state: Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`, `New`.
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
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

            __props__['disabled_alerts'] = disabled_alerts
            __props__['email_account_admins'] = email_account_admins
            __props__['email_addresses'] = email_addresses
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retention_days'] = retention_days
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            if state is None:
                raise TypeError("Missing required property 'state'")
            __props__['state'] = state
            __props__['storage_account_access_key'] = storage_account_access_key
            __props__['storage_endpoint'] = storage_endpoint
        super(ServerSecurityAlertPolicy, __self__).__init__(
            'azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, disabled_alerts=None, email_account_admins=None, email_addresses=None, resource_group_name=None, retention_days=None, server_name=None, state=None, storage_account_access_key=None, storage_endpoint=None):
        """
        Get an existing ServerSecurityAlertPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] disabled_alerts: Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
        :param pulumi.Input[bool] email_account_admins: Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
        :param pulumi.Input[list] email_addresses: Specifies an array of e-mail addresses to which the alert is sent.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[float] retention_days: Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
        :param pulumi.Input[str] server_name: Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] state: Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`, `New`.
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disabled_alerts"] = disabled_alerts
        __props__["email_account_admins"] = email_account_admins
        __props__["email_addresses"] = email_addresses
        __props__["resource_group_name"] = resource_group_name
        __props__["retention_days"] = retention_days
        __props__["server_name"] = server_name
        __props__["state"] = state
        __props__["storage_account_access_key"] = storage_account_access_key
        __props__["storage_endpoint"] = storage_endpoint
        return ServerSecurityAlertPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

