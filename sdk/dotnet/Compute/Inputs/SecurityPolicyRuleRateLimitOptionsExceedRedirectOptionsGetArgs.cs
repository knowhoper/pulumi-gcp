// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The type indicates the intended use of the security policy.
        /// * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
        /// They filter requests before they hit the origin servers.
        /// * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
        /// (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
        /// They filter requests before the request is served from Google's cache.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsGetArgs()
        {
        }
    }
}
