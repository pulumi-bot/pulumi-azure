// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media
{
    /// <summary>
    /// Manages a Media Job.
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
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleServiceAccount = new Azure.Media.ServiceAccount("exampleServiceAccount", new Azure.Media.ServiceAccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             StorageAccounts = 
    ///             {
    ///                 new Azure.Media.Inputs.ServiceAccountStorageAccountArgs
    ///                 {
    ///                     Id = exampleAccount.Id,
    ///                     IsPrimary = true,
    ///                 },
    ///             },
    ///         });
    ///         var exampleTransform = new Azure.Media.Transform("exampleTransform", new Azure.Media.TransformArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             MediaServicesAccountName = exampleServiceAccount.Name,
    ///             Description = "My transform description",
    ///             Outputs = 
    ///             {
    ///                 new Azure.Media.Inputs.TransformOutputArgs
    ///                 {
    ///                     RelativePriority = "Normal",
    ///                     OnErrorAction = "ContinueJob",
    ///                     BuiltinPreset = new Azure.Media.Inputs.TransformOutputBuiltinPresetArgs
    ///                     {
    ///                         PresetName = "AACGoodQualityAudio",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var input = new Azure.Media.Asset("input", new Azure.Media.AssetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             MediaServicesAccountName = exampleServiceAccount.Name,
    ///             Description = "Input Asset description",
    ///         });
    ///         var output = new Azure.Media.Asset("output", new Azure.Media.AssetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             MediaServicesAccountName = exampleServiceAccount.Name,
    ///             Description = "Output Asset description",
    ///         });
    ///         var exampleJob = new Azure.Media.Job("exampleJob", new Azure.Media.JobArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             MediaServicesAccountName = exampleServiceAccount.Name,
    ///             TransformName = exampleTransform.Name,
    ///             Description = "My Job description",
    ///             Priority = "Normal",
    ///             InputAsset = new Azure.Media.Inputs.JobInputAssetArgs
    ///             {
    ///                 Name = input.Name,
    ///             },
    ///             OutputAssets = 
    ///             {
    ///                 new Azure.Media.Inputs.JobOutputAssetArgs
    ///                 {
    ///                     Name = output.Name,
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
    /// Media Jobs can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:media/job:Job example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Media/mediaservices/account1/transforms/transform1/jobs/job1
    /// ```
    /// </summary>
    [AzureResourceType("azure:media/job:Job")]
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional customer supplied description of the Job.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A `input_asset` block as defined below. Changing this forces a new Media Job to be created.
        /// </summary>
        [Output("inputAsset")]
        public Output<Outputs.JobInputAsset> InputAsset { get; private set; } = null!;

        /// <summary>
        /// The Media Services account name. Changing this forces a new Transform to be created.
        /// </summary>
        [Output("mediaServicesAccountName")]
        public Output<string> MediaServicesAccountName { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Media Job. Changing this forces a new Media Job to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more `output_asset` blocks as defined below. Changing this forces a new Media Job to be created.
        /// </summary>
        [Output("outputAssets")]
        public Output<ImmutableArray<Outputs.JobOutputAsset>> OutputAssets { get; private set; } = null!;

        /// <summary>
        /// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal. Changing this forces a new Media Job to be created.
        /// </summary>
        [Output("priority")]
        public Output<string?> Priority { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Media Job should exist. Changing this forces a new Media Job to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Transform name. Changing this forces a new Media Job to be created.
        /// </summary>
        [Output("transformName")]
        public Output<string> TransformName { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("azure:media/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("azure:media/job:Job", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional customer supplied description of the Job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A `input_asset` block as defined below. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("inputAsset", required: true)]
        public Input<Inputs.JobInputAssetArgs> InputAsset { get; set; } = null!;

        /// <summary>
        /// The Media Services account name. Changing this forces a new Transform to be created.
        /// </summary>
        [Input("mediaServicesAccountName", required: true)]
        public Input<string> MediaServicesAccountName { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Media Job. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outputAssets", required: true)]
        private InputList<Inputs.JobOutputAssetArgs>? _outputAssets;

        /// <summary>
        /// One or more `output_asset` blocks as defined below. Changing this forces a new Media Job to be created.
        /// </summary>
        public InputList<Inputs.JobOutputAssetArgs> OutputAssets
        {
            get => _outputAssets ?? (_outputAssets = new InputList<Inputs.JobOutputAssetArgs>());
            set => _outputAssets = value;
        }

        /// <summary>
        /// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Media Job should exist. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Transform name. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("transformName", required: true)]
        public Input<string> TransformName { get; set; } = null!;

        public JobArgs()
        {
        }
    }

    public sealed class JobState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional customer supplied description of the Job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A `input_asset` block as defined below. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("inputAsset")]
        public Input<Inputs.JobInputAssetGetArgs>? InputAsset { get; set; }

        /// <summary>
        /// The Media Services account name. Changing this forces a new Transform to be created.
        /// </summary>
        [Input("mediaServicesAccountName")]
        public Input<string>? MediaServicesAccountName { get; set; }

        /// <summary>
        /// The name which should be used for this Media Job. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outputAssets")]
        private InputList<Inputs.JobOutputAssetGetArgs>? _outputAssets;

        /// <summary>
        /// One or more `output_asset` blocks as defined below. Changing this forces a new Media Job to be created.
        /// </summary>
        public InputList<Inputs.JobOutputAssetGetArgs> OutputAssets
        {
            get => _outputAssets ?? (_outputAssets = new InputList<Inputs.JobOutputAssetGetArgs>());
            set => _outputAssets = value;
        }

        /// <summary>
        /// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Media Job should exist. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Transform name. Changing this forces a new Media Job to be created.
        /// </summary>
        [Input("transformName")]
        public Input<string>? TransformName { get; set; }

        public JobState()
        {
        }
    }
}
