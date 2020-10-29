// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppInsights
{
    /// <summary>
    /// Manages an Application Insights component.
    /// </summary>
    public partial class Insights : Pulumi.CustomResource
    {
        /// <summary>
        /// The App ID associated with this Application Insights component.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationType")]
        public Output<string> ApplicationType { get; private set; } = null!;

        /// <summary>
        /// The Connection String for this Application Insights component. (Sensitive)
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// Specifies the Application Insights component daily data volume cap in GB.
        /// </summary>
        [Output("dailyDataCapInGb")]
        public Output<double> DailyDataCapInGb { get; private set; } = null!;

        /// <summary>
        /// Specifies if a notification email will be send when the daily data volume cap is met.
        /// </summary>
        [Output("dailyDataCapNotificationsDisabled")]
        public Output<bool> DailyDataCapNotificationsDisabled { get; private set; } = null!;

        /// <summary>
        /// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
        /// </summary>
        [Output("disableIpMasking")]
        public Output<bool?> DisableIpMasking { get; private set; } = null!;

        /// <summary>
        /// The Instrumentation Key for this Application Insights component.
        /// </summary>
        [Output("instrumentationKey")]
        public Output<string> InstrumentationKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Application Insights component. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Application Insights component.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
        /// </summary>
        [Output("retentionInDays")]
        public Output<int?> RetentionInDays { get; private set; } = null!;

        /// <summary>
        /// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
        /// </summary>
        [Output("samplingPercentage")]
        public Output<double?> SamplingPercentage { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Insights resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Insights(string name, InsightsArgs args, CustomResourceOptions? options = null)
            : base("azure:appinsights/insights:Insights", name, args ?? new InsightsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Insights(string name, Input<string> id, InsightsState? state = null, CustomResourceOptions? options = null)
            : base("azure:appinsights/insights:Insights", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Insights resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Insights Get(string name, Input<string> id, InsightsState? state = null, CustomResourceOptions? options = null)
        {
            return new Insights(name, id, state, options);
        }
    }

    public sealed class InsightsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationType", required: true)]
        public Input<string> ApplicationType { get; set; } = null!;

        /// <summary>
        /// Specifies the Application Insights component daily data volume cap in GB.
        /// </summary>
        [Input("dailyDataCapInGb")]
        public Input<double>? DailyDataCapInGb { get; set; }

        /// <summary>
        /// Specifies if a notification email will be send when the daily data volume cap is met.
        /// </summary>
        [Input("dailyDataCapNotificationsDisabled")]
        public Input<bool>? DailyDataCapNotificationsDisabled { get; set; }

        /// <summary>
        /// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
        /// </summary>
        [Input("disableIpMasking")]
        public Input<bool>? DisableIpMasking { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Application Insights component. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Application Insights component.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
        /// </summary>
        [Input("samplingPercentage")]
        public Input<double>? SamplingPercentage { get; set; }

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

        public InsightsArgs()
        {
        }
    }

    public sealed class InsightsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The App ID associated with this Application Insights component.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationType")]
        public Input<string>? ApplicationType { get; set; }

        /// <summary>
        /// The Connection String for this Application Insights component. (Sensitive)
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// Specifies the Application Insights component daily data volume cap in GB.
        /// </summary>
        [Input("dailyDataCapInGb")]
        public Input<double>? DailyDataCapInGb { get; set; }

        /// <summary>
        /// Specifies if a notification email will be send when the daily data volume cap is met.
        /// </summary>
        [Input("dailyDataCapNotificationsDisabled")]
        public Input<bool>? DailyDataCapNotificationsDisabled { get; set; }

        /// <summary>
        /// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
        /// </summary>
        [Input("disableIpMasking")]
        public Input<bool>? DisableIpMasking { get; set; }

        /// <summary>
        /// The Instrumentation Key for this Application Insights component.
        /// </summary>
        [Input("instrumentationKey")]
        public Input<string>? InstrumentationKey { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Application Insights component. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Application Insights component.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
        /// </summary>
        [Input("samplingPercentage")]
        public Input<double>? SamplingPercentage { get; set; }

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

        public InsightsState()
        {
        }
    }
}
