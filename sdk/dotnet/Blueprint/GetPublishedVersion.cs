// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Blueprint
{
    public static class GetPublishedVersion
    {
        /// <summary>
        /// Use this data source to access information about an existing Blueprint Published Version
        /// 
        /// &gt; **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
        /// </summary>
        public static Task<GetPublishedVersionResult> InvokeAsync(GetPublishedVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublishedVersionResult>("azure:blueprint/getPublishedVersion:getPublishedVersion", args ?? new GetPublishedVersionArgs(), options.WithVersion());
    }


    public sealed class GetPublishedVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Blueprint Definition
        /// </summary>
        [Input("blueprintName", required: true)]
        public string BlueprintName { get; set; } = null!;

        /// <summary>
        /// The ID of the Management Group / Subscription where this Blueprint Definition is stored.
        /// </summary>
        [Input("scopeId", required: true)]
        public string ScopeId { get; set; } = null!;

        /// <summary>
        /// The Version name of the Published Version of the Blueprint Definition
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetPublishedVersionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublishedVersionResult
    {
        public readonly string BlueprintName;
        /// <summary>
        /// The description of the Blueprint Published Version
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the Blueprint Published Version
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastModified;
        public readonly string ScopeId;
        /// <summary>
        /// The target scope
        /// </summary>
        public readonly string TargetScope;
        public readonly string TimeCreated;
        /// <summary>
        /// The type of the Blueprint
        /// </summary>
        public readonly string Type;
        public readonly string Version;

        [OutputConstructor]
        private GetPublishedVersionResult(
            string blueprintName,

            string description,

            string displayName,

            string id,

            string lastModified,

            string scopeId,

            string targetScope,

            string timeCreated,

            string type,

            string version)
        {
            BlueprintName = blueprintName;
            Description = description;
            DisplayName = displayName;
            Id = id;
            LastModified = lastModified;
            ScopeId = scopeId;
            TargetScope = targetScope;
            TimeCreated = timeCreated;
            Type = type;
            Version = version;
        }
    }
}
