// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Manages a Security Alert Policy for a MSSQL Server.
    /// 
    /// &gt; **NOTE** Security Alert Policy is currently only available for MS SQL databases.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US",
    ///         });
    ///         var exampleSqlServer = new Azure.Sql.SqlServer("exampleSqlServer", new Azure.Sql.SqlServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "4dm1n157r470r",
    ///             AdministratorLoginPassword = "4-v3ry-53cr37-p455w0rd",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleServerSecurityAlertPolicy = new Azure.MSSql.ServerSecurityAlertPolicy("exampleServerSecurityAlertPolicy", new Azure.MSSql.ServerSecurityAlertPolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ServerName = exampleSqlServer.Name,
    ///             State = "Enabled",
    ///             StorageEndpoint = exampleAccount.PrimaryBlobEndpoint,
    ///             StorageAccountAccessKey = exampleAccount.PrimaryAccessKey,
    ///             DisabledAlerts = 
    ///             {
    ///                 "Sql_Injection",
    ///                 "Data_Exfiltration",
    ///             },
    ///             RetentionDays = 20,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// MS SQL Server Security Alert Policy can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/securityAlertPolicies/Default
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy")]
    public partial class ServerSecurityAlertPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
        /// </summary>
        [Output("disabledAlerts")]
        public Output<ImmutableArray<string>> DisabledAlerts { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
        /// </summary>
        [Output("emailAccountAdmins")]
        public Output<bool?> EmailAccountAdmins { get; private set; } = null!;

        /// <summary>
        /// Specifies an array of e-mail addresses to which the alert is sent.
        /// </summary>
        [Output("emailAddresses")]
        public Output<ImmutableArray<string>> EmailAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Specifies the identifier key of the Threat Detection audit storage account.
        /// </summary>
        [Output("storageAccountAccessKey")]
        public Output<string?> StorageAccountAccessKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
        /// </summary>
        [Output("storageEndpoint")]
        public Output<string?> StorageEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a ServerSecurityAlertPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerSecurityAlertPolicy(string name, ServerSecurityAlertPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy", name, args ?? new ServerSecurityAlertPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerSecurityAlertPolicy(string name, Input<string> id, ServerSecurityAlertPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerSecurityAlertPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerSecurityAlertPolicy Get(string name, Input<string> id, ServerSecurityAlertPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerSecurityAlertPolicy(name, id, state, options);
        }
    }

    public sealed class ServerSecurityAlertPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("disabledAlerts")]
        private InputList<string>? _disabledAlerts;

        /// <summary>
        /// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
        /// </summary>
        public InputList<string> DisabledAlerts
        {
            get => _disabledAlerts ?? (_disabledAlerts = new InputList<string>());
            set => _disabledAlerts = value;
        }

        /// <summary>
        /// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
        /// </summary>
        [Input("emailAccountAdmins")]
        public Input<bool>? EmailAccountAdmins { get; set; }

        [Input("emailAddresses")]
        private InputList<string>? _emailAddresses;

        /// <summary>
        /// Specifies an array of e-mail addresses to which the alert is sent.
        /// </summary>
        public InputList<string> EmailAddresses
        {
            get => _emailAddresses ?? (_emailAddresses = new InputList<string>());
            set => _emailAddresses = value;
        }

        /// <summary>
        /// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
        /// </summary>
        [Input("state", required: true)]
        public Input<string> State { get; set; } = null!;

        /// <summary>
        /// Specifies the identifier key of the Threat Detection audit storage account.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        public ServerSecurityAlertPolicyArgs()
        {
        }
    }

    public sealed class ServerSecurityAlertPolicyState : Pulumi.ResourceArgs
    {
        [Input("disabledAlerts")]
        private InputList<string>? _disabledAlerts;

        /// <summary>
        /// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
        /// </summary>
        public InputList<string> DisabledAlerts
        {
            get => _disabledAlerts ?? (_disabledAlerts = new InputList<string>());
            set => _disabledAlerts = value;
        }

        /// <summary>
        /// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
        /// </summary>
        [Input("emailAccountAdmins")]
        public Input<bool>? EmailAccountAdmins { get; set; }

        [Input("emailAddresses")]
        private InputList<string>? _emailAddresses;

        /// <summary>
        /// Specifies an array of e-mail addresses to which the alert is sent.
        /// </summary>
        public InputList<string> EmailAddresses
        {
            get => _emailAddresses ?? (_emailAddresses = new InputList<string>());
            set => _emailAddresses = value;
        }

        /// <summary>
        /// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Specifies the identifier key of the Threat Detection audit storage account.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        public ServerSecurityAlertPolicyState()
        {
        }
    }
}
