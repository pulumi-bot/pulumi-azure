// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.RecoveryServices
{
    /// <summary>
    /// Manages a Azure recovery replicated vms (Azure to Azure). An replicated VM keeps a copiously updated image of the vm in another region in order to be able to start the VM in that region in case of a disaster. 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/recovery_replicated_vm.html.markdown.
    /// </summary>
    public partial class ReplicatedVm : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `managed_disk` block.
        /// </summary>
        [Output("managedDisks")]
        public Output<ImmutableArray<Outputs.ReplicatedVmManagedDisks>> ManagedDisks { get; private set; } = null!;

        /// <summary>
        /// The name of the network mapping.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("recoveryReplicationPolicyId")]
        public Output<string> RecoveryReplicationPolicyId { get; private set; } = null!;

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
        /// Name of fabric that should contains this replication.
        /// </summary>
        [Output("sourceRecoveryFabricName")]
        public Output<string> SourceRecoveryFabricName { get; private set; } = null!;

        /// <summary>
        /// Name of the protection container to use.
        /// </summary>
        [Output("sourceRecoveryProtectionContainerName")]
        public Output<string> SourceRecoveryProtectionContainerName { get; private set; } = null!;

        /// <summary>
        /// Id of the VM to replicate
        /// </summary>
        [Output("sourceVmId")]
        public Output<string> SourceVmId { get; private set; } = null!;

        /// <summary>
        /// Id of availability set that the new VM should belong to when a failover is done.
        /// </summary>
        [Output("targetAvailabilitySetId")]
        public Output<string?> TargetAvailabilitySetId { get; private set; } = null!;

        /// <summary>
        /// Id of fabric where the VM replication should be handled when a failover is done.
        /// </summary>
        [Output("targetRecoveryFabricId")]
        public Output<string> TargetRecoveryFabricId { get; private set; } = null!;

        /// <summary>
        /// Id of protection container where the VM replication should be created when a failover is done.
        /// </summary>
        [Output("targetRecoveryProtectionContainerId")]
        public Output<string> TargetRecoveryProtectionContainerId { get; private set; } = null!;

        /// <summary>
        /// Id of resource group where the VM should be created when a failover is done.
        /// </summary>
        [Output("targetResourceGroupId")]
        public Output<string> TargetResourceGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicatedVm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicatedVm(string name, ReplicatedVmArgs args, CustomResourceOptions? options = null)
            : base("azure:recoveryservices/replicatedVm:ReplicatedVm", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ReplicatedVm(string name, Input<string> id, ReplicatedVmState? state = null, CustomResourceOptions? options = null)
            : base("azure:recoveryservices/replicatedVm:ReplicatedVm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicatedVm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicatedVm Get(string name, Input<string> id, ReplicatedVmState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicatedVm(name, id, state, options);
        }
    }

    public sealed class ReplicatedVmArgs : Pulumi.ResourceArgs
    {
        [Input("managedDisks")]
        private InputList<Inputs.ReplicatedVmManagedDisksArgs>? _managedDisks;

        /// <summary>
        /// One or more `managed_disk` block.
        /// </summary>
        public InputList<Inputs.ReplicatedVmManagedDisksArgs> ManagedDisks
        {
            get => _managedDisks ?? (_managedDisks = new InputList<Inputs.ReplicatedVmManagedDisksArgs>());
            set => _managedDisks = value;
        }

        /// <summary>
        /// The name of the network mapping.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("recoveryReplicationPolicyId", required: true)]
        public Input<string> RecoveryReplicationPolicyId { get; set; } = null!;

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

        /// <summary>
        /// Name of fabric that should contains this replication.
        /// </summary>
        [Input("sourceRecoveryFabricName", required: true)]
        public Input<string> SourceRecoveryFabricName { get; set; } = null!;

        /// <summary>
        /// Name of the protection container to use.
        /// </summary>
        [Input("sourceRecoveryProtectionContainerName", required: true)]
        public Input<string> SourceRecoveryProtectionContainerName { get; set; } = null!;

        /// <summary>
        /// Id of the VM to replicate
        /// </summary>
        [Input("sourceVmId", required: true)]
        public Input<string> SourceVmId { get; set; } = null!;

        /// <summary>
        /// Id of availability set that the new VM should belong to when a failover is done.
        /// </summary>
        [Input("targetAvailabilitySetId")]
        public Input<string>? TargetAvailabilitySetId { get; set; }

        /// <summary>
        /// Id of fabric where the VM replication should be handled when a failover is done.
        /// </summary>
        [Input("targetRecoveryFabricId", required: true)]
        public Input<string> TargetRecoveryFabricId { get; set; } = null!;

        /// <summary>
        /// Id of protection container where the VM replication should be created when a failover is done.
        /// </summary>
        [Input("targetRecoveryProtectionContainerId", required: true)]
        public Input<string> TargetRecoveryProtectionContainerId { get; set; } = null!;

        /// <summary>
        /// Id of resource group where the VM should be created when a failover is done.
        /// </summary>
        [Input("targetResourceGroupId", required: true)]
        public Input<string> TargetResourceGroupId { get; set; } = null!;

        public ReplicatedVmArgs()
        {
        }
    }

    public sealed class ReplicatedVmState : Pulumi.ResourceArgs
    {
        [Input("managedDisks")]
        private InputList<Inputs.ReplicatedVmManagedDisksGetArgs>? _managedDisks;

        /// <summary>
        /// One or more `managed_disk` block.
        /// </summary>
        public InputList<Inputs.ReplicatedVmManagedDisksGetArgs> ManagedDisks
        {
            get => _managedDisks ?? (_managedDisks = new InputList<Inputs.ReplicatedVmManagedDisksGetArgs>());
            set => _managedDisks = value;
        }

        /// <summary>
        /// The name of the network mapping.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("recoveryReplicationPolicyId")]
        public Input<string>? RecoveryReplicationPolicyId { get; set; }

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

        /// <summary>
        /// Name of fabric that should contains this replication.
        /// </summary>
        [Input("sourceRecoveryFabricName")]
        public Input<string>? SourceRecoveryFabricName { get; set; }

        /// <summary>
        /// Name of the protection container to use.
        /// </summary>
        [Input("sourceRecoveryProtectionContainerName")]
        public Input<string>? SourceRecoveryProtectionContainerName { get; set; }

        /// <summary>
        /// Id of the VM to replicate
        /// </summary>
        [Input("sourceVmId")]
        public Input<string>? SourceVmId { get; set; }

        /// <summary>
        /// Id of availability set that the new VM should belong to when a failover is done.
        /// </summary>
        [Input("targetAvailabilitySetId")]
        public Input<string>? TargetAvailabilitySetId { get; set; }

        /// <summary>
        /// Id of fabric where the VM replication should be handled when a failover is done.
        /// </summary>
        [Input("targetRecoveryFabricId")]
        public Input<string>? TargetRecoveryFabricId { get; set; }

        /// <summary>
        /// Id of protection container where the VM replication should be created when a failover is done.
        /// </summary>
        [Input("targetRecoveryProtectionContainerId")]
        public Input<string>? TargetRecoveryProtectionContainerId { get; set; }

        /// <summary>
        /// Id of resource group where the VM should be created when a failover is done.
        /// </summary>
        [Input("targetResourceGroupId")]
        public Input<string>? TargetResourceGroupId { get; set; }

        public ReplicatedVmState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ReplicatedVmManagedDisksArgs : Pulumi.ResourceArgs
    {
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        [Input("stagingStorageAccountId", required: true)]
        public Input<string> StagingStorageAccountId { get; set; } = null!;

        [Input("targetDiskType", required: true)]
        public Input<string> TargetDiskType { get; set; } = null!;

        [Input("targetReplicaDiskType", required: true)]
        public Input<string> TargetReplicaDiskType { get; set; } = null!;

        /// <summary>
        /// Id of resource group where the VM should be created when a failover is done.
        /// </summary>
        [Input("targetResourceGroupId", required: true)]
        public Input<string> TargetResourceGroupId { get; set; } = null!;

        public ReplicatedVmManagedDisksArgs()
        {
        }
    }

    public sealed class ReplicatedVmManagedDisksGetArgs : Pulumi.ResourceArgs
    {
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        [Input("stagingStorageAccountId", required: true)]
        public Input<string> StagingStorageAccountId { get; set; } = null!;

        [Input("targetDiskType", required: true)]
        public Input<string> TargetDiskType { get; set; } = null!;

        [Input("targetReplicaDiskType", required: true)]
        public Input<string> TargetReplicaDiskType { get; set; } = null!;

        /// <summary>
        /// Id of resource group where the VM should be created when a failover is done.
        /// </summary>
        [Input("targetResourceGroupId", required: true)]
        public Input<string> TargetResourceGroupId { get; set; } = null!;

        public ReplicatedVmManagedDisksGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ReplicatedVmManagedDisks
    {
        public readonly string DiskId;
        public readonly string StagingStorageAccountId;
        public readonly string TargetDiskType;
        public readonly string TargetReplicaDiskType;
        /// <summary>
        /// Id of resource group where the VM should be created when a failover is done.
        /// </summary>
        public readonly string TargetResourceGroupId;

        [OutputConstructor]
        private ReplicatedVmManagedDisks(
            string diskId,
            string stagingStorageAccountId,
            string targetDiskType,
            string targetReplicaDiskType,
            string targetResourceGroupId)
        {
            DiskId = diskId;
            StagingStorageAccountId = stagingStorageAccountId;
            TargetDiskType = targetDiskType;
            TargetReplicaDiskType = targetReplicaDiskType;
            TargetResourceGroupId = targetResourceGroupId;
        }
    }
    }
}
