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
    /// &gt; **NOTE:** This resource has been deprecated in favour of the `azure.policy.VirtualMachineConfigurationAssignment` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, information on migrating can be found in this guide.
    /// 
    /// Applies a Configuration Policy to a Virtual Machine.
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
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///         });
    ///         var exampleNetworkInterface = new Azure.Network.NetworkInterface("exampleNetworkInterface", new Azure.Network.NetworkInterfaceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
    ///                 {
    ///                     Name = "internal",
    ///                     SubnetId = exampleSubnet.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                 },
    ///             },
    ///         });
    ///         var exampleWindowsVirtualMachine = new Azure.Compute.WindowsVirtualMachine("exampleWindowsVirtualMachine", new Azure.Compute.WindowsVirtualMachineArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Size = "Standard_F2",
    ///             AdminUsername = "adminuser",
    ///             AdminPassword = "P@$$w0rd1234!",
    ///             NetworkInterfaceIds = 
    ///             {
    ///                 exampleNetworkInterface.Id,
    ///             },
    ///             Identity = new Azure.Compute.Inputs.WindowsVirtualMachineIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///             OsDisk = new Azure.Compute.Inputs.WindowsVirtualMachineOsDiskArgs
    ///             {
    ///                 Caching = "ReadWrite",
    ///                 StorageAccountType = "Standard_LRS",
    ///             },
    ///             SourceImageReference = new Azure.Compute.Inputs.WindowsVirtualMachineSourceImageReferenceArgs
    ///             {
    ///                 Publisher = "MicrosoftWindowsServer",
    ///                 Offer = "WindowsServer",
    ///                 Sku = "2019-Datacenter",
    ///                 Version = "latest",
    ///             },
    ///         });
    ///         var exampleExtension = new Azure.Compute.Extension("exampleExtension", new Azure.Compute.ExtensionArgs
    ///         {
    ///             VirtualMachineId = exampleWindowsVirtualMachine.Id,
    ///             Publisher = "Microsoft.GuestConfiguration",
    ///             Type = "ConfigurationforWindows",
    ///             TypeHandlerVersion = "1.0",
    ///             AutoUpgradeMinorVersion = true,
    ///         });
    ///         var exampleConfigurationPolicyAssignment = new Azure.Compute.ConfigurationPolicyAssignment("exampleConfigurationPolicyAssignment", new Azure.Compute.ConfigurationPolicyAssignmentArgs
    ///         {
    ///             Location = exampleWindowsVirtualMachine.Location,
    ///             VirtualMachineId = exampleWindowsVirtualMachine.Id,
    ///             Configuration = new Azure.Compute.Inputs.ConfigurationPolicyAssignmentConfigurationArgs
    ///             {
    ///                 Name = "AzureWindowsBaseline",
    ///                 Version = "1.*",
    ///                 Parameters = 
    ///                 {
    ///                     new Azure.Compute.Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs
    ///                     {
    ///                         Name = "Minimum Password Length;ExpectedValue",
    ///                         Value = "16",
    ///                     },
    ///                     new Azure.Compute.Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs
    ///                     {
    ///                         Name = "Minimum Password Age;ExpectedValue",
    ///                         Value = "0",
    ///                     },
    ///                     new Azure.Compute.Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs
    ///                     {
    ///                         Name = "Maximum Password Age;ExpectedValue",
    ///                         Value = "30,45",
    ///                     },
    ///                     new Azure.Compute.Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs
    ///                     {
    ///                         Name = "Enforce Password History;ExpectedValue",
    ///                         Value = "10",
    ///                     },
    ///                     new Azure.Compute.Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs
    ///                     {
    ///                         Name = "Password Must Meet Complexity Requirements;ExpectedValue",
    ///                         Value = "1",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Machine Configuration Policy Assignments can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
    /// ```
    /// </summary>
    [AzureResourceType("azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment")]
    public partial class ConfigurationPolicyAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// A `configuration` block as defined below.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.ConfigurationPolicyAssignmentConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualMachineId")]
        public Output<string> VirtualMachineId { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationPolicyAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationPolicyAssignment(string name, ConfigurationPolicyAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment", name, args ?? new ConfigurationPolicyAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationPolicyAssignment(string name, Input<string> id, ConfigurationPolicyAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationPolicyAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationPolicyAssignment Get(string name, Input<string> id, ConfigurationPolicyAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigurationPolicyAssignment(name, id, state, options);
        }
    }

    public sealed class ConfigurationPolicyAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `configuration` block as defined below.
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.ConfigurationPolicyAssignmentConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineId", required: true)]
        public Input<string> VirtualMachineId { get; set; } = null!;

        public ConfigurationPolicyAssignmentArgs()
        {
        }
    }

    public sealed class ConfigurationPolicyAssignmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `configuration` block as defined below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ConfigurationPolicyAssignmentConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        public ConfigurationPolicyAssignmentState()
        {
        }
    }
}
