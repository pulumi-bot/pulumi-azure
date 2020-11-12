// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    /// <summary>
    /// Manages an Azure IoT Time Series Insights Reference Data Set.
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
    ///             Location = "northeurope",
    ///         });
    ///         var exampleTimeSeriesInsightsStandardEnvironment = new Azure.Iot.TimeSeriesInsightsStandardEnvironment("exampleTimeSeriesInsightsStandardEnvironment", new Azure.Iot.TimeSeriesInsightsStandardEnvironmentArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SkuName = "S1_1",
    ///             DataRetentionTime = "P30D",
    ///         });
    ///         var exampleTimeSeriesInsightsReferenceDataSet = new Azure.Iot.TimeSeriesInsightsReferenceDataSet("exampleTimeSeriesInsightsReferenceDataSet", new Azure.Iot.TimeSeriesInsightsReferenceDataSetArgs
    ///         {
    ///             TimeSeriesInsightsEnvironmentId = exampleTimeSeriesInsightsStandardEnvironment.Id,
    ///             Location = exampleResourceGroup.Location,
    ///             KeyProperties = 
    ///             {
    ///                 new Azure.Iot.Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyArgs
    ///                 {
    ///                     Name = "keyProperty1",
    ///                     Type = "String",
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
    /// Azure IoT Time Series Insights Reference Data Set can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example/referenceDataSets/example
    /// ```
    /// </summary>
    public partial class TimeSeriesInsightsReferenceDataSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
        /// </summary>
        [Output("dataStringComparisonBehavior")]
        public Output<string?> DataStringComparisonBehavior { get; private set; } = null!;

        /// <summary>
        /// A `key_property` block as defined below.
        /// </summary>
        [Output("keyProperties")]
        public Output<ImmutableArray<Outputs.TimeSeriesInsightsReferenceDataSetKeyProperty>> KeyProperties { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("timeSeriesInsightsEnvironmentId")]
        public Output<string> TimeSeriesInsightsEnvironmentId { get; private set; } = null!;


        /// <summary>
        /// Create a TimeSeriesInsightsReferenceDataSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TimeSeriesInsightsReferenceDataSet(string name, TimeSeriesInsightsReferenceDataSetArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet", name, args ?? new TimeSeriesInsightsReferenceDataSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TimeSeriesInsightsReferenceDataSet(string name, Input<string> id, TimeSeriesInsightsReferenceDataSetState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TimeSeriesInsightsReferenceDataSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TimeSeriesInsightsReferenceDataSet Get(string name, Input<string> id, TimeSeriesInsightsReferenceDataSetState? state = null, CustomResourceOptions? options = null)
        {
            return new TimeSeriesInsightsReferenceDataSet(name, id, state, options);
        }
    }

    public sealed class TimeSeriesInsightsReferenceDataSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
        /// </summary>
        [Input("dataStringComparisonBehavior")]
        public Input<string>? DataStringComparisonBehavior { get; set; }

        [Input("keyProperties", required: true)]
        private InputList<Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyArgs>? _keyProperties;

        /// <summary>
        /// A `key_property` block as defined below.
        /// </summary>
        public InputList<Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyArgs> KeyProperties
        {
            get => _keyProperties ?? (_keyProperties = new InputList<Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyArgs>());
            set => _keyProperties = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("timeSeriesInsightsEnvironmentId", required: true)]
        public Input<string> TimeSeriesInsightsEnvironmentId { get; set; } = null!;

        public TimeSeriesInsightsReferenceDataSetArgs()
        {
        }
    }

    public sealed class TimeSeriesInsightsReferenceDataSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
        /// </summary>
        [Input("dataStringComparisonBehavior")]
        public Input<string>? DataStringComparisonBehavior { get; set; }

        [Input("keyProperties")]
        private InputList<Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyGetArgs>? _keyProperties;

        /// <summary>
        /// A `key_property` block as defined below.
        /// </summary>
        public InputList<Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyGetArgs> KeyProperties
        {
            get => _keyProperties ?? (_keyProperties = new InputList<Inputs.TimeSeriesInsightsReferenceDataSetKeyPropertyGetArgs>());
            set => _keyProperties = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("timeSeriesInsightsEnvironmentId")]
        public Input<string>? TimeSeriesInsightsEnvironmentId { get; set; }

        public TimeSeriesInsightsReferenceDataSetState()
        {
        }
    }
}
