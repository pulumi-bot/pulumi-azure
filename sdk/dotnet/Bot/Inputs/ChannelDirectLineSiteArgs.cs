// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Bot.Inputs
{

    public sealed class ChannelDirectLineSiteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables/Disables this site. Enabled by default
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
        /// </summary>
        [Input("enhancedAuthenticationEnabled")]
        public Input<bool>? EnhancedAuthenticationEnabled { get; set; }

        /// <summary>
        /// Id for the site
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Primary key for accessing this site
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Secondary key for accessing this site
        /// </summary>
        [Input("key2")]
        public Input<string>? Key2 { get; set; }

        /// <summary>
        /// The name of the site
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("trustedOrigins")]
        private InputList<string>? _trustedOrigins;

        /// <summary>
        /// This field is required when `is_secure_site_enabled` is enabled. Determines which origins can establish a Directline conversation for this site.
        /// </summary>
        public InputList<string> TrustedOrigins
        {
            get => _trustedOrigins ?? (_trustedOrigins = new InputList<string>());
            set => _trustedOrigins = value;
        }

        /// <summary>
        /// Enables v1 of the Directline protocol for this site. Enabled by default
        /// </summary>
        [Input("v1Allowed")]
        public Input<bool>? V1Allowed { get; set; }

        /// <summary>
        /// Enables v3 of the Directline protocol for this site. Enabled by default
        /// </summary>
        [Input("v3Allowed")]
        public Input<bool>? V3Allowed { get; set; }

        public ChannelDirectLineSiteArgs()
        {
        }
    }
}
