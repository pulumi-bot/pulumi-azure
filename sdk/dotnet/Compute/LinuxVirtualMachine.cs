// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages a Linux Virtual Machine.
    /// 
    /// ## Disclaimers
    /// 
    /// &gt; **Note** This provider will automatically remove the OS Disk by default - this behaviour can be configured using the `features` configuration within the Provider configuration block.
    /// 
    /// &gt; **Note** This resource does not support Unmanaged Disks. If you need to use Unmanaged Disks you can continue to use the `azure.compute.VirtualMachine` resource instead.
    /// 
    /// &gt; **Note** This resource does not support attaching existing OS Disks. You can instead capture an image of the OS Disk or continue to use the `azure.compute.VirtualMachine` resource instead.
    /// 
    /// &gt; In this release there's a known issue where the `public_ip_address` and `public_ip_addresses` fields may not be fully populated for Dynamic Public IP's.
    /// </summary>
    public partial class LinuxVirtualMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// A `additional_capabilities` block as defined below.
        /// </summary>
        [Output("additionalCapabilities")]
        public Output<Outputs.LinuxVirtualMachineAdditionalCapabilities?> AdditionalCapabilities { get; private set; } = null!;

        /// <summary>
        /// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Output("adminPassword")]
        public Output<string?> AdminPassword { get; private set; } = null!;

        /// <summary>
        /// One or more `admin_ssh_key` blocks as defined below.
        /// </summary>
        [Output("adminSshKeys")]
        public Output<ImmutableArray<Outputs.LinuxVirtualMachineAdminSshKey>> AdminSshKeys { get; private set; } = null!;

        /// <summary>
        /// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Output("adminUsername")]
        public Output<string> AdminUsername { get; private set; } = null!;

        /// <summary>
        /// Should Extension Operations be allowed on this Virtual Machine?
        /// </summary>
        [Output("allowExtensionOperations")]
        public Output<bool?> AllowExtensionOperations { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("availabilitySetId")]
        public Output<string?> AvailabilitySetId { get; private set; } = null!;

        /// <summary>
        /// A `boot_diagnostics` block as defined below.
        /// </summary>
        [Output("bootDiagnostics")]
        public Output<Outputs.LinuxVirtualMachineBootDiagnostics?> BootDiagnostics { get; private set; } = null!;

        /// <summary>
        /// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computer_name`, then you must specify `computer_name`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("computerName")]
        public Output<string> ComputerName { get; private set; } = null!;

        /// <summary>
        /// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Output("customData")]
        public Output<string?> CustomData { get; private set; } = null!;

        /// <summary>
        /// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dedicatedHostId")]
        public Output<string?> DedicatedHostId { get; private set; } = null!;

        /// <summary>
        /// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("disablePasswordAuthentication")]
        public Output<bool?> DisablePasswordAuthentication { get; private set; } = null!;

        /// <summary>
        /// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
        /// </summary>
        [Output("encryptionAtHostEnabled")]
        public Output<bool?> EncryptionAtHostEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("evictionPolicy")]
        public Output<string?> EvictionPolicy { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.LinuxVirtualMachineIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `eviction_policy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
        /// </summary>
        [Output("maxBidPrice")]
        public Output<double?> MaxBidPrice { get; private set; } = null!;

        /// <summary>
        /// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
        /// </summary>
        [Output("networkInterfaceIds")]
        public Output<ImmutableArray<string>> NetworkInterfaceIds { get; private set; } = null!;

        /// <summary>
        /// A `os_disk` block as defined below.
        /// </summary>
        [Output("osDisk")]
        public Output<Outputs.LinuxVirtualMachineOsDisk> OsDisk { get; private set; } = null!;

        /// <summary>
        /// A `plan` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.LinuxVirtualMachinePlan?> Plan { get; private set; } = null!;

        /// <summary>
        /// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("priority")]
        public Output<string?> Priority { get; private set; } = null!;

        /// <summary>
        /// The Primary Private IP Address assigned to this Virtual Machine.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// A list of Private IP Addresses assigned to this Virtual Machine.
        /// </summary>
        [Output("privateIpAddresses")]
        public Output<ImmutableArray<string>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("provisionVmAgent")]
        public Output<bool?> ProvisionVmAgent { get; private set; } = null!;

        /// <summary>
        /// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("proximityPlacementGroupId")]
        public Output<string?> ProximityPlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// The Primary Public IP Address assigned to this Virtual Machine.
        /// </summary>
        [Output("publicIpAddress")]
        public Output<string> PublicIpAddress { get; private set; } = null!;

        /// <summary>
        /// A list of the Public IP Addresses assigned to this Virtual Machine.
        /// </summary>
        [Output("publicIpAddresses")]
        public Output<ImmutableArray<string>> PublicIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// One or more `secret` blocks as defined below.
        /// </summary>
        [Output("secrets")]
        public Output<ImmutableArray<Outputs.LinuxVirtualMachineSecret>> Secrets { get; private set; } = null!;

        /// <summary>
        /// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
        /// </summary>
        [Output("size")]
        public Output<string> Size { get; private set; } = null!;

        /// <summary>
        /// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sourceImageId")]
        public Output<string?> SourceImageId { get; private set; } = null!;

        /// <summary>
        /// A `source_image_reference` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sourceImageReference")]
        public Output<Outputs.LinuxVirtualMachineSourceImageReference?> SourceImageReference { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to this Virtual Machine.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A 128-bit identifier which uniquely identifies this Virtual Machine.
        /// </summary>
        [Output("virtualMachineId")]
        public Output<string> VirtualMachineId { get; private set; } = null!;

        /// <summary>
        /// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualMachineScaleSetId")]
        public Output<string?> VirtualMachineScaleSetId { get; private set; } = null!;

        /// <summary>
        /// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a LinuxVirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinuxVirtualMachine(string name, LinuxVirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/linuxVirtualMachine:LinuxVirtualMachine", name, args ?? new LinuxVirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinuxVirtualMachine(string name, Input<string> id, LinuxVirtualMachineState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/linuxVirtualMachine:LinuxVirtualMachine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LinuxVirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinuxVirtualMachine Get(string name, Input<string> id, LinuxVirtualMachineState? state = null, CustomResourceOptions? options = null)
        {
            return new LinuxVirtualMachine(name, id, state, options);
        }
    }

    public sealed class LinuxVirtualMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `additional_capabilities` block as defined below.
        /// </summary>
        [Input("additionalCapabilities")]
        public Input<Inputs.LinuxVirtualMachineAdditionalCapabilitiesArgs>? AdditionalCapabilities { get; set; }

        /// <summary>
        /// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        [Input("adminSshKeys")]
        private InputList<Inputs.LinuxVirtualMachineAdminSshKeyArgs>? _adminSshKeys;

        /// <summary>
        /// One or more `admin_ssh_key` blocks as defined below.
        /// </summary>
        public InputList<Inputs.LinuxVirtualMachineAdminSshKeyArgs> AdminSshKeys
        {
            get => _adminSshKeys ?? (_adminSshKeys = new InputList<Inputs.LinuxVirtualMachineAdminSshKeyArgs>());
            set => _adminSshKeys = value;
        }

        /// <summary>
        /// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("adminUsername", required: true)]
        public Input<string> AdminUsername { get; set; } = null!;

        /// <summary>
        /// Should Extension Operations be allowed on this Virtual Machine?
        /// </summary>
        [Input("allowExtensionOperations")]
        public Input<bool>? AllowExtensionOperations { get; set; }

        /// <summary>
        /// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("availabilitySetId")]
        public Input<string>? AvailabilitySetId { get; set; }

        /// <summary>
        /// A `boot_diagnostics` block as defined below.
        /// </summary>
        [Input("bootDiagnostics")]
        public Input<Inputs.LinuxVirtualMachineBootDiagnosticsArgs>? BootDiagnostics { get; set; }

        /// <summary>
        /// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computer_name`, then you must specify `computer_name`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("computerName")]
        public Input<string>? ComputerName { get; set; }

        /// <summary>
        /// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        /// <summary>
        /// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostId")]
        public Input<string>? DedicatedHostId { get; set; }

        /// <summary>
        /// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("disablePasswordAuthentication")]
        public Input<bool>? DisablePasswordAuthentication { get; set; }

        /// <summary>
        /// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
        /// </summary>
        [Input("encryptionAtHostEnabled")]
        public Input<bool>? EncryptionAtHostEnabled { get; set; }

        /// <summary>
        /// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.LinuxVirtualMachineIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `eviction_policy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
        /// </summary>
        [Input("maxBidPrice")]
        public Input<double>? MaxBidPrice { get; set; }

        /// <summary>
        /// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaceIds", required: true)]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// A `os_disk` block as defined below.
        /// </summary>
        [Input("osDisk", required: true)]
        public Input<Inputs.LinuxVirtualMachineOsDiskArgs> OsDisk { get; set; } = null!;

        /// <summary>
        /// A `plan` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.LinuxVirtualMachinePlanArgs>? Plan { get; set; }

        /// <summary>
        /// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("provisionVmAgent")]
        public Input<bool>? ProvisionVmAgent { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("secrets")]
        private InputList<Inputs.LinuxVirtualMachineSecretArgs>? _secrets;

        /// <summary>
        /// One or more `secret` blocks as defined below.
        /// </summary>
        public InputList<Inputs.LinuxVirtualMachineSecretArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.LinuxVirtualMachineSecretArgs>());
            set => _secrets = value;
        }

        /// <summary>
        /// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
        /// </summary>
        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        /// <summary>
        /// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceImageId")]
        public Input<string>? SourceImageId { get; set; }

        /// <summary>
        /// A `source_image_reference` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceImageReference")]
        public Input<Inputs.LinuxVirtualMachineSourceImageReferenceArgs>? SourceImageReference { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to this Virtual Machine.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineScaleSetId")]
        public Input<string>? VirtualMachineScaleSetId { get; set; }

        /// <summary>
        /// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LinuxVirtualMachineArgs()
        {
        }
    }

    public sealed class LinuxVirtualMachineState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `additional_capabilities` block as defined below.
        /// </summary>
        [Input("additionalCapabilities")]
        public Input<Inputs.LinuxVirtualMachineAdditionalCapabilitiesGetArgs>? AdditionalCapabilities { get; set; }

        /// <summary>
        /// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        [Input("adminSshKeys")]
        private InputList<Inputs.LinuxVirtualMachineAdminSshKeyGetArgs>? _adminSshKeys;

        /// <summary>
        /// One or more `admin_ssh_key` blocks as defined below.
        /// </summary>
        public InputList<Inputs.LinuxVirtualMachineAdminSshKeyGetArgs> AdminSshKeys
        {
            get => _adminSshKeys ?? (_adminSshKeys = new InputList<Inputs.LinuxVirtualMachineAdminSshKeyGetArgs>());
            set => _adminSshKeys = value;
        }

        /// <summary>
        /// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("adminUsername")]
        public Input<string>? AdminUsername { get; set; }

        /// <summary>
        /// Should Extension Operations be allowed on this Virtual Machine?
        /// </summary>
        [Input("allowExtensionOperations")]
        public Input<bool>? AllowExtensionOperations { get; set; }

        /// <summary>
        /// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("availabilitySetId")]
        public Input<string>? AvailabilitySetId { get; set; }

        /// <summary>
        /// A `boot_diagnostics` block as defined below.
        /// </summary>
        [Input("bootDiagnostics")]
        public Input<Inputs.LinuxVirtualMachineBootDiagnosticsGetArgs>? BootDiagnostics { get; set; }

        /// <summary>
        /// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computer_name`, then you must specify `computer_name`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("computerName")]
        public Input<string>? ComputerName { get; set; }

        /// <summary>
        /// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        /// <summary>
        /// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostId")]
        public Input<string>? DedicatedHostId { get; set; }

        /// <summary>
        /// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("disablePasswordAuthentication")]
        public Input<bool>? DisablePasswordAuthentication { get; set; }

        /// <summary>
        /// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
        /// </summary>
        [Input("encryptionAtHostEnabled")]
        public Input<bool>? EncryptionAtHostEnabled { get; set; }

        /// <summary>
        /// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.LinuxVirtualMachineIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `eviction_policy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
        /// </summary>
        [Input("maxBidPrice")]
        public Input<double>? MaxBidPrice { get; set; }

        /// <summary>
        /// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaceIds")]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// A `os_disk` block as defined below.
        /// </summary>
        [Input("osDisk")]
        public Input<Inputs.LinuxVirtualMachineOsDiskGetArgs>? OsDisk { get; set; }

        /// <summary>
        /// A `plan` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.LinuxVirtualMachinePlanGetArgs>? Plan { get; set; }

        /// <summary>
        /// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The Primary Private IP Address assigned to this Virtual Machine.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;

        /// <summary>
        /// A list of Private IP Addresses assigned to this Virtual Machine.
        /// </summary>
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        /// <summary>
        /// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("provisionVmAgent")]
        public Input<bool>? ProvisionVmAgent { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The Primary Public IP Address assigned to this Virtual Machine.
        /// </summary>
        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        [Input("publicIpAddresses")]
        private InputList<string>? _publicIpAddresses;

        /// <summary>
        /// A list of the Public IP Addresses assigned to this Virtual Machine.
        /// </summary>
        public InputList<string> PublicIpAddresses
        {
            get => _publicIpAddresses ?? (_publicIpAddresses = new InputList<string>());
            set => _publicIpAddresses = value;
        }

        /// <summary>
        /// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("secrets")]
        private InputList<Inputs.LinuxVirtualMachineSecretGetArgs>? _secrets;

        /// <summary>
        /// One or more `secret` blocks as defined below.
        /// </summary>
        public InputList<Inputs.LinuxVirtualMachineSecretGetArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.LinuxVirtualMachineSecretGetArgs>());
            set => _secrets = value;
        }

        /// <summary>
        /// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceImageId")]
        public Input<string>? SourceImageId { get; set; }

        /// <summary>
        /// A `source_image_reference` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceImageReference")]
        public Input<Inputs.LinuxVirtualMachineSourceImageReferenceGetArgs>? SourceImageReference { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to this Virtual Machine.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A 128-bit identifier which uniquely identifies this Virtual Machine.
        /// </summary>
        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        /// <summary>
        /// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineScaleSetId")]
        public Input<string>? VirtualMachineScaleSetId { get; set; }

        /// <summary>
        /// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LinuxVirtualMachineState()
        {
        }
    }
}
