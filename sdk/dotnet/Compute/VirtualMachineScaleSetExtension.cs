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
    /// Manages an Extension for a Virtual Machine Scale Set.
    /// 
    /// &gt; **NOTE:** This resource is not intended to be used with the `azure.compute.ScaleSet` resource - instead it's intended for this to be used with the `azure.compute.LinuxVirtualMachineScaleSet` and `azure.compute.WindowsVirtualMachineScaleSet` resources.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleLinuxVirtualMachineScaleSet = new Azure.Compute.LinuxVirtualMachineScaleSet("exampleLinuxVirtualMachineScaleSet", new Azure.Compute.LinuxVirtualMachineScaleSetArgs
    ///         {
    ///         });
    ///         //...
    ///         var exampleVirtualMachineScaleSetExtension = new Azure.Compute.VirtualMachineScaleSetExtension("exampleVirtualMachineScaleSetExtension", new Azure.Compute.VirtualMachineScaleSetExtensionArgs
    ///         {
    ///             VirtualMachineScaleSetId = exampleLinuxVirtualMachineScaleSet.Id,
    ///             Publisher = "Microsoft.Azure.Extensions",
    ///             Type = "CustomScript",
    ///             TypeHandlerVersion = "2.0",
    ///             Settings = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "commandToExecute", "echo $HOSTNAME" },
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Machine Scale Set Extensions can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class VirtualMachineScaleSetExtension : Pulumi.CustomResource
    {
        /// <summary>
        /// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        /// </summary>
        [Output("autoUpgradeMinorVersion")]
        public Output<bool?> AutoUpgradeMinorVersion { get; private set; } = null!;

        /// <summary>
        /// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        /// </summary>
        [Output("forceUpdateTag")]
        public Output<string?> ForceUpdateTag { get; private set; } = null!;

        /// <summary>
        /// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        /// </summary>
        [Output("protectedSettings")]
        public Output<string?> ProtectedSettings { get; private set; } = null!;

        /// <summary>
        /// An ordered list of Extension names which this should be provisioned after.
        /// </summary>
        [Output("provisionAfterExtensions")]
        public Output<ImmutableArray<string>> ProvisionAfterExtensions { get; private set; } = null!;

        /// <summary>
        /// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Output("publisher")]
        public Output<string> Publisher { get; private set; } = null!;

        /// <summary>
        /// A JSON String which specifies Settings for the Extension.
        /// </summary>
        [Output("settings")]
        public Output<string?> Settings { get; private set; } = null!;

        /// <summary>
        /// Specifies the Type of the Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        /// </summary>
        [Output("typeHandlerVersion")]
        public Output<string> TypeHandlerVersion { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualMachineScaleSetId")]
        public Output<string> VirtualMachineScaleSetId { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachineScaleSetExtension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachineScaleSetExtension(string name, VirtualMachineScaleSetExtensionArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension", name, args ?? new VirtualMachineScaleSetExtensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachineScaleSetExtension(string name, Input<string> id, VirtualMachineScaleSetExtensionState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachineScaleSetExtension Get(string name, Input<string> id, VirtualMachineScaleSetExtensionState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualMachineScaleSetExtension(name, id, state, options);
        }
    }

    public sealed class VirtualMachineScaleSetExtensionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<bool>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        /// </summary>
        [Input("forceUpdateTag")]
        public Input<string>? ForceUpdateTag { get; set; }

        /// <summary>
        /// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        /// </summary>
        [Input("protectedSettings")]
        public Input<string>? ProtectedSettings { get; set; }

        [Input("provisionAfterExtensions")]
        private InputList<string>? _provisionAfterExtensions;

        /// <summary>
        /// An ordered list of Extension names which this should be provisioned after.
        /// </summary>
        public InputList<string> ProvisionAfterExtensions
        {
            get => _provisionAfterExtensions ?? (_provisionAfterExtensions = new InputList<string>());
            set => _provisionAfterExtensions = value;
        }

        /// <summary>
        /// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        /// <summary>
        /// A JSON String which specifies Settings for the Extension.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        /// <summary>
        /// Specifies the Type of the Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        /// </summary>
        [Input("typeHandlerVersion", required: true)]
        public Input<string> TypeHandlerVersion { get; set; } = null!;

        /// <summary>
        /// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineScaleSetId", required: true)]
        public Input<string> VirtualMachineScaleSetId { get; set; } = null!;

        public VirtualMachineScaleSetExtensionArgs()
        {
        }
    }

    public sealed class VirtualMachineScaleSetExtensionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<bool>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        /// </summary>
        [Input("forceUpdateTag")]
        public Input<string>? ForceUpdateTag { get; set; }

        /// <summary>
        /// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        /// </summary>
        [Input("protectedSettings")]
        public Input<string>? ProtectedSettings { get; set; }

        [Input("provisionAfterExtensions")]
        private InputList<string>? _provisionAfterExtensions;

        /// <summary>
        /// An ordered list of Extension names which this should be provisioned after.
        /// </summary>
        public InputList<string> ProvisionAfterExtensions
        {
            get => _provisionAfterExtensions ?? (_provisionAfterExtensions = new InputList<string>());
            set => _provisionAfterExtensions = value;
        }

        /// <summary>
        /// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        /// <summary>
        /// A JSON String which specifies Settings for the Extension.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        /// <summary>
        /// Specifies the Type of the Extension. Changing this forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        /// </summary>
        [Input("typeHandlerVersion")]
        public Input<string>? TypeHandlerVersion { get; set; }

        /// <summary>
        /// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineScaleSetId")]
        public Input<string>? VirtualMachineScaleSetId { get; set; }

        public VirtualMachineScaleSetExtensionState()
        {
        }
    }
}
