// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages network rules inside of a Azure Storage Account.
    /// 
    /// &gt; **NOTE:** Network Rules can be defined either directly on the `azure.storage.Account` resource, or using the `azure.storage.AccountNetworkRules` resource - but the two cannot be used together. Spurious changes will occur if both are used against the same Storage Account.
    /// 
    /// &gt; **NOTE:** Only one `azure.storage.AccountNetworkRules` can be tied to an `azure.storage.Account`. Spurious changes will occur if more than `azure.storage.AccountNetworkRules` is tied to the same `azure.storage.Account`.
    /// 
    /// &gt; **NOTE:** Deleting this resource updates the storage account back to the default values it had when the storage account was created.
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///             ServiceEndpoints = 
    ///             {
    ///                 "Microsoft.Storage",
    ///             },
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///             Tags = 
    ///             {
    ///                 { "environment", "staging" },
    ///             },
    ///         });
    ///         var test = new Azure.Storage.AccountNetworkRules("test", new Azure.Storage.AccountNetworkRulesArgs
    ///         {
    ///             ResourceGroupName = azurerm_resource_group.Test.Name,
    ///             StorageAccountName = azurerm_storage_account.Test.Name,
    ///             DefaultAction = "Allow",
    ///             IpRules = 
    ///             {
    ///                 "127.0.0.1",
    ///             },
    ///             VirtualNetworkSubnetIds = 
    ///             {
    ///                 azurerm_subnet.Test.Id,
    ///             },
    ///             Bypasses = 
    ///             {
    ///                 "Metrics",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Storage Account Network Rules can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class AccountNetworkRules : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
        /// </summary>
        [Output("bypasses")]
        public Output<ImmutableArray<string>> Bypasses { get; private set; } = null!;

        /// <summary>
        /// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
        /// </summary>
        [Output("defaultAction")]
        public Output<string> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
        /// </summary>
        [Output("ipRules")]
        public Output<ImmutableArray<string>> IpRules { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;

        /// <summary>
        /// A list of virtual network subnet ids to to secure the storage account.
        /// </summary>
        [Output("virtualNetworkSubnetIds")]
        public Output<ImmutableArray<string>> VirtualNetworkSubnetIds { get; private set; } = null!;


        /// <summary>
        /// Create a AccountNetworkRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountNetworkRules(string name, AccountNetworkRulesArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/accountNetworkRules:AccountNetworkRules", name, args ?? new AccountNetworkRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountNetworkRules(string name, Input<string> id, AccountNetworkRulesState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/accountNetworkRules:AccountNetworkRules", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountNetworkRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountNetworkRules Get(string name, Input<string> id, AccountNetworkRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountNetworkRules(name, id, state, options);
        }
    }

    public sealed class AccountNetworkRulesArgs : Pulumi.ResourceArgs
    {
        [Input("bypasses")]
        private InputList<string>? _bypasses;

        /// <summary>
        /// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
        /// </summary>
        public InputList<string> Bypasses
        {
            get => _bypasses ?? (_bypasses = new InputList<string>());
            set => _bypasses = value;
        }

        /// <summary>
        /// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        [Input("ipRules")]
        private InputList<string>? _ipRules;

        /// <summary>
        /// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
        /// </summary>
        public InputList<string> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<string>());
            set => _ipRules = value;
        }

        /// <summary>
        /// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        [Input("virtualNetworkSubnetIds")]
        private InputList<string>? _virtualNetworkSubnetIds;

        /// <summary>
        /// A list of virtual network subnet ids to to secure the storage account.
        /// </summary>
        public InputList<string> VirtualNetworkSubnetIds
        {
            get => _virtualNetworkSubnetIds ?? (_virtualNetworkSubnetIds = new InputList<string>());
            set => _virtualNetworkSubnetIds = value;
        }

        public AccountNetworkRulesArgs()
        {
        }
    }

    public sealed class AccountNetworkRulesState : Pulumi.ResourceArgs
    {
        [Input("bypasses")]
        private InputList<string>? _bypasses;

        /// <summary>
        /// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
        /// </summary>
        public InputList<string> Bypasses
        {
            get => _bypasses ?? (_bypasses = new InputList<string>());
            set => _bypasses = value;
        }

        /// <summary>
        /// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        [Input("ipRules")]
        private InputList<string>? _ipRules;

        /// <summary>
        /// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
        /// </summary>
        public InputList<string> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<string>());
            set => _ipRules = value;
        }

        /// <summary>
        /// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        [Input("virtualNetworkSubnetIds")]
        private InputList<string>? _virtualNetworkSubnetIds;

        /// <summary>
        /// A list of virtual network subnet ids to to secure the storage account.
        /// </summary>
        public InputList<string> VirtualNetworkSubnetIds
        {
            get => _virtualNetworkSubnetIds ?? (_virtualNetworkSubnetIds = new InputList<string>());
            set => _virtualNetworkSubnetIds = value;
        }

        public AccountNetworkRulesState()
        {
        }
    }
}
