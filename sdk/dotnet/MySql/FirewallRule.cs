// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql
{
    /// <summary>
    /// Manages a Firewall Rule for a MySQL Server
    /// 
    /// ## Example Usage
    /// ### Single IP Address)
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
    ///         var exampleServer = new Azure.MySql.Server("exampleServer", new Azure.MySql.ServerArgs
    ///         {
    ///         });
    ///         // ...
    ///         var exampleFirewallRule = new Azure.MySql.FirewallRule("exampleFirewallRule", new Azure.MySql.FirewallRuleArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ServerName = exampleServer.Name,
    ///             StartIpAddress = "40.112.8.12",
    ///             EndIpAddress = "40.112.8.12",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### IP Range)
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
    ///         var exampleServer = new Azure.MySql.Server("exampleServer", new Azure.MySql.ServerArgs
    ///         {
    ///         });
    ///         // ...
    ///         var exampleFirewallRule = new Azure.MySql.FirewallRule("exampleFirewallRule", new Azure.MySql.FirewallRuleArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ServerName = exampleServer.Name,
    ///             StartIpAddress = "40.112.0.0",
    ///             EndIpAddress = "40.112.255.255",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class FirewallRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("endIpAddress")]
        public Output<string> EndIpAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("startIpAddress")]
        public Output<string> StartIpAddress { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallRule(string name, FirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:mysql/firewallRule:FirewallRule", name, args ?? new FirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallRule(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:mysql/firewallRule:FirewallRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallRule Get(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallRule(name, id, state, options);
        }
    }

    public sealed class FirewallRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("endIpAddress", required: true)]
        public Input<string> EndIpAddress { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("startIpAddress", required: true)]
        public Input<string> StartIpAddress { get; set; } = null!;

        public FirewallRuleArgs()
        {
        }
    }

    public sealed class FirewallRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("endIpAddress")]
        public Input<string>? EndIpAddress { get; set; }

        /// <summary>
        /// Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("startIpAddress")]
        public Input<string>? StartIpAddress { get; set; }

        public FirewallRuleState()
        {
        }
    }
}
