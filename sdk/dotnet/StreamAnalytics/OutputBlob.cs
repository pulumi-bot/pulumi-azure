// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.StreamAnalytics
{
    /// <summary>
    /// Manages a Stream Analytics Output to Blob Storage.
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
    ///         var exampleResourceGroup = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
    ///         {
    ///             Name = "example-resources",
    ///         }));
    ///         var exampleJob = Output.Create(Azure.StreamAnalytics.GetJob.InvokeAsync(new Azure.StreamAnalytics.GetJobArgs
    ///         {
    ///             Name = "example-job",
    ///             ResourceGroupName = azurerm_resource_group.Example.Name,
    ///         }));
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Apply(exampleResourceGroup =&gt; exampleResourceGroup.Name),
    ///             Location = exampleResourceGroup.Apply(exampleResourceGroup =&gt; exampleResourceGroup.Location),
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var exampleContainer = new Azure.Storage.Container("exampleContainer", new Azure.Storage.ContainerArgs
    ///         {
    ///             StorageAccountName = exampleAccount.Name,
    ///             ContainerAccessType = "private",
    ///         });
    ///         var exampleOutputBlob = new Azure.StreamAnalytics.OutputBlob("exampleOutputBlob", new Azure.StreamAnalytics.OutputBlobArgs
    ///         {
    ///             StreamAnalyticsJobName = exampleJob.Apply(exampleJob =&gt; exampleJob.Name),
    ///             ResourceGroupName = exampleJob.Apply(exampleJob =&gt; exampleJob.ResourceGroupName),
    ///             StorageAccountName = exampleAccount.Name,
    ///             StorageAccountKey = exampleAccount.PrimaryAccessKey,
    ///             StorageContainerName = exampleContainer.Name,
    ///             PathPattern = "some-pattern",
    ///             DateFormat = "yyyy-MM-dd",
    ///             TimeFormat = "HH",
    ///             Serialization = new Azure.StreamAnalytics.Inputs.OutputBlobSerializationArgs
    ///             {
    ///                 Type = "Csv",
    ///                 Encoding = "UTF8",
    ///                 FieldDelimiter = ",",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Stream Analytics Outputs to Blob Storage can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class OutputBlob : Pulumi.CustomResource
    {
        /// <summary>
        /// The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        /// </summary>
        [Output("dateFormat")]
        public Output<string> DateFormat { get; private set; } = null!;

        /// <summary>
        /// The name of the Stream Output. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        /// </summary>
        [Output("pathPattern")]
        public Output<string> PathPattern { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `serialization` block as defined below.
        /// </summary>
        [Output("serialization")]
        public Output<Outputs.OutputBlobSerialization> Serialization { get; private set; } = null!;

        /// <summary>
        /// The Access Key which should be used to connect to this Storage Account.
        /// </summary>
        [Output("storageAccountKey")]
        public Output<string> StorageAccountKey { get; private set; } = null!;

        /// <summary>
        /// The name of the Storage Account.
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;

        /// <summary>
        /// The name of the Container within the Storage Account.
        /// </summary>
        [Output("storageContainerName")]
        public Output<string> StorageContainerName { get; private set; } = null!;

        /// <summary>
        /// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        /// </summary>
        [Output("streamAnalyticsJobName")]
        public Output<string> StreamAnalyticsJobName { get; private set; } = null!;

        /// <summary>
        /// The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        /// </summary>
        [Output("timeFormat")]
        public Output<string> TimeFormat { get; private set; } = null!;


        /// <summary>
        /// Create a OutputBlob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutputBlob(string name, OutputBlobArgs args, CustomResourceOptions? options = null)
            : base("azure:streamanalytics/outputBlob:OutputBlob", name, args ?? new OutputBlobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OutputBlob(string name, Input<string> id, OutputBlobState? state = null, CustomResourceOptions? options = null)
            : base("azure:streamanalytics/outputBlob:OutputBlob", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OutputBlob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutputBlob Get(string name, Input<string> id, OutputBlobState? state = null, CustomResourceOptions? options = null)
        {
            return new OutputBlob(name, id, state, options);
        }
    }

    public sealed class OutputBlobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        /// </summary>
        [Input("dateFormat", required: true)]
        public Input<string> DateFormat { get; set; } = null!;

        /// <summary>
        /// The name of the Stream Output. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        /// </summary>
        [Input("pathPattern", required: true)]
        public Input<string> PathPattern { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `serialization` block as defined below.
        /// </summary>
        [Input("serialization", required: true)]
        public Input<Inputs.OutputBlobSerializationArgs> Serialization { get; set; } = null!;

        /// <summary>
        /// The Access Key which should be used to connect to this Storage Account.
        /// </summary>
        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Account.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Container within the Storage Account.
        /// </summary>
        [Input("storageContainerName", required: true)]
        public Input<string> StorageContainerName { get; set; } = null!;

        /// <summary>
        /// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        /// </summary>
        [Input("streamAnalyticsJobName", required: true)]
        public Input<string> StreamAnalyticsJobName { get; set; } = null!;

        /// <summary>
        /// The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        /// </summary>
        [Input("timeFormat", required: true)]
        public Input<string> TimeFormat { get; set; } = null!;

        public OutputBlobArgs()
        {
        }
    }

    public sealed class OutputBlobState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        /// </summary>
        [Input("dateFormat")]
        public Input<string>? DateFormat { get; set; }

        /// <summary>
        /// The name of the Stream Output. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        /// </summary>
        [Input("pathPattern")]
        public Input<string>? PathPattern { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `serialization` block as defined below.
        /// </summary>
        [Input("serialization")]
        public Input<Inputs.OutputBlobSerializationGetArgs>? Serialization { get; set; }

        /// <summary>
        /// The Access Key which should be used to connect to this Storage Account.
        /// </summary>
        [Input("storageAccountKey")]
        public Input<string>? StorageAccountKey { get; set; }

        /// <summary>
        /// The name of the Storage Account.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        /// <summary>
        /// The name of the Container within the Storage Account.
        /// </summary>
        [Input("storageContainerName")]
        public Input<string>? StorageContainerName { get; set; }

        /// <summary>
        /// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        /// </summary>
        [Input("streamAnalyticsJobName")]
        public Input<string>? StreamAnalyticsJobName { get; set; }

        /// <summary>
        /// The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        public OutputBlobState()
        {
        }
    }
}
