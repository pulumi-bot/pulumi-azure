// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SiteRecovery
{
    /// <summary>
    /// Manages a Azure recovery vault protection container mapping. A protection container mapping decides how to translate the protection container when a VM is migrated from one region to another.
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
    ///         var primaryResourceGroup = new Azure.Core.ResourceGroup("primaryResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US",
    ///         });
    ///         var secondaryResourceGroup = new Azure.Core.ResourceGroup("secondaryResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "East US",
    ///         });
    ///         var vault = new Azure.RecoveryServices.Vault("vault", new Azure.RecoveryServices.VaultArgs
    ///         {
    ///             Location = secondaryResourceGroup.Location,
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             Sku = "Standard",
    ///         });
    ///         var primaryFabric = new Azure.SiteRecovery.Fabric("primaryFabric", new Azure.SiteRecovery.FabricArgs
    ///         {
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             RecoveryVaultName = vault.Name,
    ///             Location = primaryResourceGroup.Location,
    ///         });
    ///         var secondaryFabric = new Azure.SiteRecovery.Fabric("secondaryFabric", new Azure.SiteRecovery.FabricArgs
    ///         {
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             RecoveryVaultName = vault.Name,
    ///             Location = secondaryResourceGroup.Location,
    ///         });
    ///         var primaryProtectionContainer = new Azure.SiteRecovery.ProtectionContainer("primaryProtectionContainer", new Azure.SiteRecovery.ProtectionContainerArgs
    ///         {
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             RecoveryVaultName = vault.Name,
    ///             RecoveryFabricName = primaryFabric.Name,
    ///         });
    ///         var secondaryProtectionContainer = new Azure.SiteRecovery.ProtectionContainer("secondaryProtectionContainer", new Azure.SiteRecovery.ProtectionContainerArgs
    ///         {
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             RecoveryVaultName = vault.Name,
    ///             RecoveryFabricName = secondaryFabric.Name,
    ///         });
    ///         var policy = new Azure.SiteRecovery.ReplicationPolicy("policy", new Azure.SiteRecovery.ReplicationPolicyArgs
    ///         {
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             RecoveryVaultName = vault.Name,
    ///             RecoveryPointRetentionInMinutes = 24 * 60,
    ///             ApplicationConsistentSnapshotFrequencyInMinutes = 4 * 60,
    ///         });
    ///         var container_mapping = new Azure.SiteRecovery.ProtectionContainerMapping("container-mapping", new Azure.SiteRecovery.ProtectionContainerMappingArgs
    ///         {
    ///             ResourceGroupName = secondaryResourceGroup.Name,
    ///             RecoveryVaultName = vault.Name,
    ///             RecoveryFabricName = primaryFabric.Name,
    ///             RecoverySourceProtectionContainerName = primaryProtectionContainer.Name,
    ///             RecoveryTargetProtectionContainerId = secondaryProtectionContainer.Id,
    ///             RecoveryReplicationPolicyId = policy.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Site Recovery Protection Container Mappings can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:siterecovery/protectionContainerMapping:ProtectionContainerMapping mymapping /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name
    /// ```
    /// </summary>
    public partial class ProtectionContainerMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the network mapping.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of fabric that should contains the protection container to map.
        /// </summary>
        [Output("recoveryFabricName")]
        public Output<string> RecoveryFabricName { get; private set; } = null!;

        /// <summary>
        /// Id of the policy to use for this mapping.
        /// </summary>
        [Output("recoveryReplicationPolicyId")]
        public Output<string> RecoveryReplicationPolicyId { get; private set; } = null!;

        /// <summary>
        /// Name of the source protection container to map.
        /// </summary>
        [Output("recoverySourceProtectionContainerName")]
        public Output<string> RecoverySourceProtectionContainerName { get; private set; } = null!;

        /// <summary>
        /// Id of target protection container to map to.
        /// </summary>
        [Output("recoveryTargetProtectionContainerId")]
        public Output<string> RecoveryTargetProtectionContainerId { get; private set; } = null!;

        /// <summary>
        /// The name of the vault that should be updated.
        /// </summary>
        [Output("recoveryVaultName")]
        public Output<string> RecoveryVaultName { get; private set; } = null!;

        /// <summary>
        /// Name of the resource group where the vault that should be updated is located.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a ProtectionContainerMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProtectionContainerMapping(string name, ProtectionContainerMappingArgs args, CustomResourceOptions? options = null)
            : base("azure:siterecovery/protectionContainerMapping:ProtectionContainerMapping", name, args ?? new ProtectionContainerMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProtectionContainerMapping(string name, Input<string> id, ProtectionContainerMappingState? state = null, CustomResourceOptions? options = null)
            : base("azure:siterecovery/protectionContainerMapping:ProtectionContainerMapping", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProtectionContainerMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProtectionContainerMapping Get(string name, Input<string> id, ProtectionContainerMappingState? state = null, CustomResourceOptions? options = null)
        {
            return new ProtectionContainerMapping(name, id, state, options);
        }
    }

    public sealed class ProtectionContainerMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the network mapping.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of fabric that should contains the protection container to map.
        /// </summary>
        [Input("recoveryFabricName", required: true)]
        public Input<string> RecoveryFabricName { get; set; } = null!;

        /// <summary>
        /// Id of the policy to use for this mapping.
        /// </summary>
        [Input("recoveryReplicationPolicyId", required: true)]
        public Input<string> RecoveryReplicationPolicyId { get; set; } = null!;

        /// <summary>
        /// Name of the source protection container to map.
        /// </summary>
        [Input("recoverySourceProtectionContainerName", required: true)]
        public Input<string> RecoverySourceProtectionContainerName { get; set; } = null!;

        /// <summary>
        /// Id of target protection container to map to.
        /// </summary>
        [Input("recoveryTargetProtectionContainerId", required: true)]
        public Input<string> RecoveryTargetProtectionContainerId { get; set; } = null!;

        /// <summary>
        /// The name of the vault that should be updated.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group where the vault that should be updated is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ProtectionContainerMappingArgs()
        {
        }
    }

    public sealed class ProtectionContainerMappingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the network mapping.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of fabric that should contains the protection container to map.
        /// </summary>
        [Input("recoveryFabricName")]
        public Input<string>? RecoveryFabricName { get; set; }

        /// <summary>
        /// Id of the policy to use for this mapping.
        /// </summary>
        [Input("recoveryReplicationPolicyId")]
        public Input<string>? RecoveryReplicationPolicyId { get; set; }

        /// <summary>
        /// Name of the source protection container to map.
        /// </summary>
        [Input("recoverySourceProtectionContainerName")]
        public Input<string>? RecoverySourceProtectionContainerName { get; set; }

        /// <summary>
        /// Id of target protection container to map to.
        /// </summary>
        [Input("recoveryTargetProtectionContainerId")]
        public Input<string>? RecoveryTargetProtectionContainerId { get; set; }

        /// <summary>
        /// The name of the vault that should be updated.
        /// </summary>
        [Input("recoveryVaultName")]
        public Input<string>? RecoveryVaultName { get; set; }

        /// <summary>
        /// Name of the resource group where the vault that should be updated is located.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ProtectionContainerMappingState()
        {
        }
    }
}
